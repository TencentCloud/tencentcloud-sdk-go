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

package v20240523

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2024-05-23"

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


func NewCheckAnimateImageJobRequest() (request *CheckAnimateImageJobRequest) {
    request = &CheckAnimateImageJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "CheckAnimateImageJob")
    
    
    return
}

func NewCheckAnimateImageJobResponse() (response *CheckAnimateImageJobResponse) {
    response = &CheckAnimateImageJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CheckAnimateImageJob
// 检查图片跳舞输入图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckAnimateImageJob(request *CheckAnimateImageJobRequest) (response *CheckAnimateImageJobResponse, err error) {
    return c.CheckAnimateImageJobWithContext(context.Background(), request)
}

// CheckAnimateImageJob
// 检查图片跳舞输入图
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckAnimateImageJobWithContext(ctx context.Context, request *CheckAnimateImageJobRequest) (response *CheckAnimateImageJobResponse, err error) {
    if request == nil {
        request = NewCheckAnimateImageJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "CheckAnimateImageJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAnimateImageJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAnimateImageJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAigcElementRequest() (request *CreateAigcElementRequest) {
    request = &CreateAigcElementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "CreateAigcElement")
    
    
    return
}

func NewCreateAigcElementResponse() (response *CreateAigcElementResponse) {
    response = &CreateAigcElementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateAigcElement
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAigcElement(request *CreateAigcElementRequest) (response *CreateAigcElementResponse, err error) {
    return c.CreateAigcElementWithContext(context.Background(), request)
}

// CreateAigcElement
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateAigcElementWithContext(ctx context.Context, request *CreateAigcElementRequest) (response *CreateAigcElementResponse, err error) {
    if request == nil {
        request = NewCreateAigcElementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "CreateAigcElement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateAigcElement require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateAigcElementResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAigcElementRequest() (request *DeleteAigcElementRequest) {
    request = &DeleteAigcElementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DeleteAigcElement")
    
    
    return
}

func NewDeleteAigcElementResponse() (response *DeleteAigcElementResponse) {
    response = &DeleteAigcElementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DeleteAigcElement
// 删除主体库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAigcElement(request *DeleteAigcElementRequest) (response *DeleteAigcElementResponse, err error) {
    return c.DeleteAigcElementWithContext(context.Background(), request)
}

// DeleteAigcElement
// 删除主体库
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DeleteAigcElementWithContext(ctx context.Context, request *DeleteAigcElementRequest) (response *DeleteAigcElementResponse, err error) {
    if request == nil {
        request = NewDeleteAigcElementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DeleteAigcElement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteAigcElement require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteAigcElementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcElementRequest() (request *DescribeAigcElementRequest) {
    request = &DescribeAigcElementRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeAigcElement")
    
    
    return
}

func NewDescribeAigcElementResponse() (response *DescribeAigcElementResponse) {
    response = &DescribeAigcElementResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcElement
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAigcElement(request *DescribeAigcElementRequest) (response *DescribeAigcElementResponse, err error) {
    return c.DescribeAigcElementWithContext(context.Background(), request)
}

// DescribeAigcElement
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAigcElementWithContext(ctx context.Context, request *DescribeAigcElementRequest) (response *DescribeAigcElementResponse, err error) {
    if request == nil {
        request = NewDescribeAigcElementRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeAigcElement")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcElement require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcElementResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAigcVideoJobRequest() (request *DescribeAigcVideoJobRequest) {
    request = &DescribeAigcVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeAigcVideoJob")
    
    
    return
}

func NewDescribeAigcVideoJobResponse() (response *DescribeAigcVideoJobResponse) {
    response = &DescribeAigcVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeAigcVideoJob
// 查询生视频任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAigcVideoJob(request *DescribeAigcVideoJobRequest) (response *DescribeAigcVideoJobResponse, err error) {
    return c.DescribeAigcVideoJobWithContext(context.Background(), request)
}

// DescribeAigcVideoJob
// 查询生视频任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) DescribeAigcVideoJobWithContext(ctx context.Context, request *DescribeAigcVideoJobRequest) (response *DescribeAigcVideoJobResponse, err error) {
    if request == nil {
        request = NewDescribeAigcVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeAigcVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeAigcVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeAigcVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHumanActorJobRequest() (request *DescribeHumanActorJobRequest) {
    request = &DescribeHumanActorJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeHumanActorJob")
    
    
    return
}

func NewDescribeHumanActorJobResponse() (response *DescribeHumanActorJobResponse) {
    response = &DescribeHumanActorJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHumanActorJob
// 通过JobId提交请求，获取人像驱动任务的结果信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeHumanActorJob(request *DescribeHumanActorJobRequest) (response *DescribeHumanActorJobResponse, err error) {
    return c.DescribeHumanActorJobWithContext(context.Background(), request)
}

// DescribeHumanActorJob
// 通过JobId提交请求，获取人像驱动任务的结果信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeHumanActorJobWithContext(ctx context.Context, request *DescribeHumanActorJobRequest) (response *DescribeHumanActorJobResponse, err error) {
    if request == nil {
        request = NewDescribeHumanActorJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeHumanActorJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHumanActorJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHumanActorJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeHunyuanToVideoJobRequest() (request *DescribeHunyuanToVideoJobRequest) {
    request = &DescribeHunyuanToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeHunyuanToVideoJob")
    
    
    return
}

func NewDescribeHunyuanToVideoJobResponse() (response *DescribeHunyuanToVideoJobResponse) {
    response = &DescribeHunyuanToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeHunyuanToVideoJob
// 查询混元生视频任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeHunyuanToVideoJob(request *DescribeHunyuanToVideoJobRequest) (response *DescribeHunyuanToVideoJobResponse, err error) {
    return c.DescribeHunyuanToVideoJobWithContext(context.Background(), request)
}

// DescribeHunyuanToVideoJob
// 查询混元生视频任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeHunyuanToVideoJobWithContext(ctx context.Context, request *DescribeHunyuanToVideoJobRequest) (response *DescribeHunyuanToVideoJobResponse, err error) {
    if request == nil {
        request = NewDescribeHunyuanToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeHunyuanToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeHunyuanToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeHunyuanToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageAnimateJobRequest() (request *DescribeImageAnimateJobRequest) {
    request = &DescribeImageAnimateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeImageAnimateJob")
    
    
    return
}

func NewDescribeImageAnimateJobResponse() (response *DescribeImageAnimateJobResponse) {
    response = &DescribeImageAnimateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageAnimateJob
// 用于查询图片跳舞任务。图片跳舞能力支持舞蹈动作结合图片生成跳舞视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageAnimateJob(request *DescribeImageAnimateJobRequest) (response *DescribeImageAnimateJobResponse, err error) {
    return c.DescribeImageAnimateJobWithContext(context.Background(), request)
}

// DescribeImageAnimateJob
// 用于查询图片跳舞任务。图片跳舞能力支持舞蹈动作结合图片生成跳舞视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageAnimateJobWithContext(ctx context.Context, request *DescribeImageAnimateJobRequest) (response *DescribeImageAnimateJobResponse, err error) {
    if request == nil {
        request = NewDescribeImageAnimateJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeImageAnimateJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAnimateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAnimateJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageToVideoGeneralJobRequest() (request *DescribeImageToVideoGeneralJobRequest) {
    request = &DescribeImageToVideoGeneralJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeImageToVideoGeneralJob")
    
    
    return
}

func NewDescribeImageToVideoGeneralJobResponse() (response *DescribeImageToVideoGeneralJobResponse) {
    response = &DescribeImageToVideoGeneralJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageToVideoGeneralJob
// 查询图生视频通用能力任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageToVideoGeneralJob(request *DescribeImageToVideoGeneralJobRequest) (response *DescribeImageToVideoGeneralJobResponse, err error) {
    return c.DescribeImageToVideoGeneralJobWithContext(context.Background(), request)
}

// DescribeImageToVideoGeneralJob
// 查询图生视频通用能力任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageToVideoGeneralJobWithContext(ctx context.Context, request *DescribeImageToVideoGeneralJobRequest) (response *DescribeImageToVideoGeneralJobResponse, err error) {
    if request == nil {
        request = NewDescribeImageToVideoGeneralJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeImageToVideoGeneralJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageToVideoGeneralJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageToVideoGeneralJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageToVideoJobRequest() (request *DescribeImageToVideoJobRequest) {
    request = &DescribeImageToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeImageToVideoJob")
    
    
    return
}

func NewDescribeImageToVideoJobResponse() (response *DescribeImageToVideoJobResponse) {
    response = &DescribeImageToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageToVideoJob
// 用于查询视频特效任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageToVideoJob(request *DescribeImageToVideoJobRequest) (response *DescribeImageToVideoJobResponse, err error) {
    return c.DescribeImageToVideoJobWithContext(context.Background(), request)
}

// DescribeImageToVideoJob
// 用于查询视频特效任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageToVideoJobWithContext(ctx context.Context, request *DescribeImageToVideoJobRequest) (response *DescribeImageToVideoJobResponse, err error) {
    if request == nil {
        request = NewDescribeImageToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeImageToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImageToVideoViduJobRequest() (request *DescribeImageToVideoViduJobRequest) {
    request = &DescribeImageToVideoViduJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeImageToVideoViduJob")
    
    
    return
}

func NewDescribeImageToVideoViduJobResponse() (response *DescribeImageToVideoViduJobResponse) {
    response = &DescribeImageToVideoViduJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeImageToVideoViduJob
// 查询Vidu图生视频任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageToVideoViduJob(request *DescribeImageToVideoViduJobRequest) (response *DescribeImageToVideoViduJobResponse, err error) {
    return c.DescribeImageToVideoViduJobWithContext(context.Background(), request)
}

// DescribeImageToVideoViduJob
// 查询Vidu图生视频任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeImageToVideoViduJobWithContext(ctx context.Context, request *DescribeImageToVideoViduJobRequest) (response *DescribeImageToVideoViduJobResponse, err error) {
    if request == nil {
        request = NewDescribeImageToVideoViduJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeImageToVideoViduJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageToVideoViduJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageToVideoViduJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMotionControlKlingJobRequest() (request *DescribeMotionControlKlingJobRequest) {
    request = &DescribeMotionControlKlingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeMotionControlKlingJob")
    
    
    return
}

func NewDescribeMotionControlKlingJobResponse() (response *DescribeMotionControlKlingJobResponse) {
    response = &DescribeMotionControlKlingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMotionControlKlingJob
// 查询Kling动作控制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMotionControlKlingJob(request *DescribeMotionControlKlingJobRequest) (response *DescribeMotionControlKlingJobResponse, err error) {
    return c.DescribeMotionControlKlingJobWithContext(context.Background(), request)
}

// DescribeMotionControlKlingJob
// 查询Kling动作控制任务
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeMotionControlKlingJobWithContext(ctx context.Context, request *DescribeMotionControlKlingJobRequest) (response *DescribeMotionControlKlingJobResponse, err error) {
    if request == nil {
        request = NewDescribeMotionControlKlingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeMotionControlKlingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMotionControlKlingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMotionControlKlingJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePortraitSingJobRequest() (request *DescribePortraitSingJobRequest) {
    request = &DescribePortraitSingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribePortraitSingJob")
    
    
    return
}

func NewDescribePortraitSingJobResponse() (response *DescribePortraitSingJobResponse) {
    response = &DescribePortraitSingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribePortraitSingJob
// 用于查询图片唱演任务。
//
// 支持提交音频和图片生成唱演视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
func (c *Client) DescribePortraitSingJob(request *DescribePortraitSingJobRequest) (response *DescribePortraitSingJobResponse, err error) {
    return c.DescribePortraitSingJobWithContext(context.Background(), request)
}

// DescribePortraitSingJob
// 用于查询图片唱演任务。
//
// 支持提交音频和图片生成唱演视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
func (c *Client) DescribePortraitSingJobWithContext(ctx context.Context, request *DescribePortraitSingJobRequest) (response *DescribePortraitSingJobResponse, err error) {
    if request == nil {
        request = NewDescribePortraitSingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribePortraitSingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePortraitSingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePortraitSingJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeReferenceToVideoViduJobRequest() (request *DescribeReferenceToVideoViduJobRequest) {
    request = &DescribeReferenceToVideoViduJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeReferenceToVideoViduJob")
    
    
    return
}

func NewDescribeReferenceToVideoViduJobResponse() (response *DescribeReferenceToVideoViduJobResponse) {
    response = &DescribeReferenceToVideoViduJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeReferenceToVideoViduJob
// 查询Vidu参考生视频任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
func (c *Client) DescribeReferenceToVideoViduJob(request *DescribeReferenceToVideoViduJobRequest) (response *DescribeReferenceToVideoViduJobResponse, err error) {
    return c.DescribeReferenceToVideoViduJobWithContext(context.Background(), request)
}

// DescribeReferenceToVideoViduJob
// 查询Vidu参考生视频任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
func (c *Client) DescribeReferenceToVideoViduJobWithContext(ctx context.Context, request *DescribeReferenceToVideoViduJobRequest) (response *DescribeReferenceToVideoViduJobResponse, err error) {
    if request == nil {
        request = NewDescribeReferenceToVideoViduJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeReferenceToVideoViduJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeReferenceToVideoViduJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeReferenceToVideoViduJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateToVideoJobRequest() (request *DescribeTemplateToVideoJobRequest) {
    request = &DescribeTemplateToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeTemplateToVideoJob")
    
    
    return
}

func NewDescribeTemplateToVideoJobResponse() (response *DescribeTemplateToVideoJobResponse) {
    response = &DescribeTemplateToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTemplateToVideoJob
// 用于查询视频特效任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTemplateToVideoJob(request *DescribeTemplateToVideoJobRequest) (response *DescribeTemplateToVideoJobResponse, err error) {
    return c.DescribeTemplateToVideoJobWithContext(context.Background(), request)
}

// DescribeTemplateToVideoJob
// 用于查询视频特效任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTemplateToVideoJobWithContext(ctx context.Context, request *DescribeTemplateToVideoJobRequest) (response *DescribeTemplateToVideoJobResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeTemplateToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplateToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextToVideoJobRequest() (request *DescribeTextToVideoJobRequest) {
    request = &DescribeTextToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeTextToVideoJob")
    
    
    return
}

func NewDescribeTextToVideoJobResponse() (response *DescribeTextToVideoJobResponse) {
    response = &DescribeTextToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTextToVideoJob
// 用于查询文生视频任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTextToVideoJob(request *DescribeTextToVideoJobRequest) (response *DescribeTextToVideoJobResponse, err error) {
    return c.DescribeTextToVideoJobWithContext(context.Background(), request)
}

// DescribeTextToVideoJob
// 用于查询文生视频任务。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTextToVideoJobWithContext(ctx context.Context, request *DescribeTextToVideoJobRequest) (response *DescribeTextToVideoJobResponse, err error) {
    if request == nil {
        request = NewDescribeTextToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeTextToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTextToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTextToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTextToVideoViduJobRequest() (request *DescribeTextToVideoViduJobRequest) {
    request = &DescribeTextToVideoViduJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeTextToVideoViduJob")
    
    
    return
}

func NewDescribeTextToVideoViduJobResponse() (response *DescribeTextToVideoViduJobResponse) {
    response = &DescribeTextToVideoViduJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeTextToVideoViduJob
// 查询Vidu文生视频任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTextToVideoViduJob(request *DescribeTextToVideoViduJobRequest) (response *DescribeTextToVideoViduJobResponse, err error) {
    return c.DescribeTextToVideoViduJobWithContext(context.Background(), request)
}

// DescribeTextToVideoViduJob
// 查询Vidu文生视频任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_IMAGEBODYSMALL = "FailedOperation.ImageBodySmall"
//  FAILEDOPERATION_IMAGECHECKNOBODY = "FailedOperation.ImageCheckNoBody"
//  FAILEDOPERATION_IMAGEDETECTFACEFAILED = "FailedOperation.ImageDetectFaceFailed"
//  FAILEDOPERATION_IMAGEMANYPEOPLE = "FailedOperation.ImageManyPeople"
//  FAILEDOPERATION_IMAGERADIOEXCCEED = "FailedOperation.ImageRadioExcceed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  FAILEDOPERATION_KEYPOINTUNDETECTED = "FailedOperation.KeyPointUndetected"
//  FAILEDOPERATION_TEMPLATEFIRSTFRAMENOTDETECTFACE = "FailedOperation.TemplateFirstFrameNotDetectFace"
//  FAILEDOPERATION_TEMPLATENOTDETECTBODY = "FailedOperation.TemplateNotDetectBody"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DescribeTextToVideoViduJobWithContext(ctx context.Context, request *DescribeTextToVideoViduJobRequest) (response *DescribeTextToVideoViduJobResponse, err error) {
    if request == nil {
        request = NewDescribeTextToVideoViduJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeTextToVideoViduJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTextToVideoViduJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTextToVideoViduJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoEditJobRequest() (request *DescribeVideoEditJobRequest) {
    request = &DescribeVideoEditJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeVideoEditJob")
    
    
    return
}

func NewDescribeVideoEditJobResponse() (response *DescribeVideoEditJobResponse) {
    response = &DescribeVideoEditJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoEditJob
// 用于提交视频编辑任务，支持上传视频、文本及图片素材开展编辑操作，涵盖风格迁移、元素替换、内容增减等核心能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoEditJob(request *DescribeVideoEditJobRequest) (response *DescribeVideoEditJobResponse, err error) {
    return c.DescribeVideoEditJobWithContext(context.Background(), request)
}

// DescribeVideoEditJob
// 用于提交视频编辑任务，支持上传视频、文本及图片素材开展编辑操作，涵盖风格迁移、元素替换、内容增减等核心能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoEditJobWithContext(ctx context.Context, request *DescribeVideoEditJobRequest) (response *DescribeVideoEditJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoEditJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeVideoEditJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoEditJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoEditJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoEditKlingJobRequest() (request *DescribeVideoEditKlingJobRequest) {
    request = &DescribeVideoEditKlingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeVideoEditKlingJob")
    
    
    return
}

func NewDescribeVideoEditKlingJobResponse() (response *DescribeVideoEditKlingJobResponse) {
    response = &DescribeVideoEditKlingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoEditKlingJob
// 查询Kling多模态编辑任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoEditKlingJob(request *DescribeVideoEditKlingJobRequest) (response *DescribeVideoEditKlingJobResponse, err error) {
    return c.DescribeVideoEditKlingJobWithContext(context.Background(), request)
}

// DescribeVideoEditKlingJob
// 查询Kling多模态编辑任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoEditKlingJobWithContext(ctx context.Context, request *DescribeVideoEditKlingJobRequest) (response *DescribeVideoEditKlingJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoEditKlingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeVideoEditKlingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoEditKlingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoEditKlingJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoExtendKlingJobRequest() (request *DescribeVideoExtendKlingJobRequest) {
    request = &DescribeVideoExtendKlingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeVideoExtendKlingJob")
    
    
    return
}

func NewDescribeVideoExtendKlingJobResponse() (response *DescribeVideoExtendKlingJobResponse) {
    response = &DescribeVideoExtendKlingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoExtendKlingJob
// 查询视频延长任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoExtendKlingJob(request *DescribeVideoExtendKlingJobRequest) (response *DescribeVideoExtendKlingJobResponse, err error) {
    return c.DescribeVideoExtendKlingJobWithContext(context.Background(), request)
}

// DescribeVideoExtendKlingJob
// 查询视频延长任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoExtendKlingJobWithContext(ctx context.Context, request *DescribeVideoExtendKlingJobRequest) (response *DescribeVideoExtendKlingJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoExtendKlingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeVideoExtendKlingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoExtendKlingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoExtendKlingJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoFaceFusionJobRequest() (request *DescribeVideoFaceFusionJobRequest) {
    request = &DescribeVideoFaceFusionJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeVideoFaceFusionJob")
    
    
    return
}

func NewDescribeVideoFaceFusionJobResponse() (response *DescribeVideoFaceFusionJobResponse) {
    response = &DescribeVideoFaceFusionJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoFaceFusionJob
// 查询视频人脸融合任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoFaceFusionJob(request *DescribeVideoFaceFusionJobRequest) (response *DescribeVideoFaceFusionJobResponse, err error) {
    return c.DescribeVideoFaceFusionJobWithContext(context.Background(), request)
}

// DescribeVideoFaceFusionJob
// 查询视频人脸融合任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoFaceFusionJobWithContext(ctx context.Context, request *DescribeVideoFaceFusionJobRequest) (response *DescribeVideoFaceFusionJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoFaceFusionJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeVideoFaceFusionJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoFaceFusionJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoFaceFusionJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoStylizationJobRequest() (request *DescribeVideoStylizationJobRequest) {
    request = &DescribeVideoStylizationJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeVideoStylizationJob")
    
    
    return
}

func NewDescribeVideoStylizationJobResponse() (response *DescribeVideoStylizationJobResponse) {
    response = &DescribeVideoStylizationJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoStylizationJob
// 用于查询视频风格化任务。视频风格化支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_TASKSTATUSERROR = "FailedOperation.TaskStatusError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeVideoStylizationJob(request *DescribeVideoStylizationJobRequest) (response *DescribeVideoStylizationJobResponse, err error) {
    return c.DescribeVideoStylizationJobWithContext(context.Background(), request)
}

// DescribeVideoStylizationJob
// 用于查询视频风格化任务。视频风格化支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_TASKNOTEXIST = "FailedOperation.TaskNotExist"
//  FAILEDOPERATION_TASKSTATUSERROR = "FailedOperation.TaskStatusError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeVideoStylizationJobWithContext(ctx context.Context, request *DescribeVideoStylizationJobRequest) (response *DescribeVideoStylizationJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoStylizationJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeVideoStylizationJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoStylizationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoStylizationJobResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeVideoVoiceJobRequest() (request *DescribeVideoVoiceJobRequest) {
    request = &DescribeVideoVoiceJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeVideoVoiceJob")
    
    
    return
}

func NewDescribeVideoVoiceJobResponse() (response *DescribeVideoVoiceJobResponse) {
    response = &DescribeVideoVoiceJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoVoiceJob
// 通过JobId提交请求，获取视频配音频任务的结果信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoVoiceJob(request *DescribeVideoVoiceJobRequest) (response *DescribeVideoVoiceJobResponse, err error) {
    return c.DescribeVideoVoiceJobWithContext(context.Background(), request)
}

// DescribeVideoVoiceJob
// 通过JobId提交请求，获取视频配音频任务的结果信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) DescribeVideoVoiceJobWithContext(ctx context.Context, request *DescribeVideoVoiceJobRequest) (response *DescribeVideoVoiceJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoVoiceJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "DescribeVideoVoiceJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoVoiceJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoVoiceJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitAigcVideoJobRequest() (request *SubmitAigcVideoJobRequest) {
    request = &SubmitAigcVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitAigcVideoJob")
    
    
    return
}

func NewSubmitAigcVideoJobResponse() (response *SubmitAigcVideoJobResponse) {
    response = &SubmitAigcVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitAigcVideoJob
// 提交生视频任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitAigcVideoJob(request *SubmitAigcVideoJobRequest) (response *SubmitAigcVideoJobResponse, err error) {
    return c.SubmitAigcVideoJobWithContext(context.Background(), request)
}

// SubmitAigcVideoJob
// 提交生视频任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitAigcVideoJobWithContext(ctx context.Context, request *SubmitAigcVideoJobRequest) (response *SubmitAigcVideoJobResponse, err error) {
    if request == nil {
        request = NewSubmitAigcVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitAigcVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitAigcVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitAigcVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHumanActorJobRequest() (request *SubmitHumanActorJobRequest) {
    request = &SubmitHumanActorJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitHumanActorJob")
    
    
    return
}

func NewSubmitHumanActorJobResponse() (response *SubmitHumanActorJobResponse) {
    response = &SubmitHumanActorJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHumanActorJob
// 用于提交人像驱动任务
//
// 支持提交音频和图文来生成对应视频，满足动态交互、内容生产等场景需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitHumanActorJob(request *SubmitHumanActorJobRequest) (response *SubmitHumanActorJobResponse, err error) {
    return c.SubmitHumanActorJobWithContext(context.Background(), request)
}

// SubmitHumanActorJob
// 用于提交人像驱动任务
//
// 支持提交音频和图文来生成对应视频，满足动态交互、内容生产等场景需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitHumanActorJobWithContext(ctx context.Context, request *SubmitHumanActorJobRequest) (response *SubmitHumanActorJobResponse, err error) {
    if request == nil {
        request = NewSubmitHumanActorJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitHumanActorJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHumanActorJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHumanActorJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitHunyuanToVideoJobRequest() (request *SubmitHunyuanToVideoJobRequest) {
    request = &SubmitHunyuanToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitHunyuanToVideoJob")
    
    
    return
}

func NewSubmitHunyuanToVideoJobResponse() (response *SubmitHunyuanToVideoJobResponse) {
    response = &SubmitHunyuanToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitHunyuanToVideoJob
// ●混元生视频接口，基于混元大模型，根据输入的文本或图片智能生成视频。
//
// 
//
// ●默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitHunyuanToVideoJob(request *SubmitHunyuanToVideoJobRequest) (response *SubmitHunyuanToVideoJobResponse, err error) {
    return c.SubmitHunyuanToVideoJobWithContext(context.Background(), request)
}

// SubmitHunyuanToVideoJob
// ●混元生视频接口，基于混元大模型，根据输入的文本或图片智能生成视频。
//
// 
//
// ●默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DRIVERFAILED = "FailedOperation.DriverFailed"
//  FAILEDOPERATION_JOBNOTFOUND = "FailedOperation.JobNotFound"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_VIDEOURLINVALID = "InvalidParameter.VideoUrlInvalid"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitHunyuanToVideoJobWithContext(ctx context.Context, request *SubmitHunyuanToVideoJobRequest) (response *SubmitHunyuanToVideoJobResponse, err error) {
    if request == nil {
        request = NewSubmitHunyuanToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitHunyuanToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitHunyuanToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitHunyuanToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitImageAnimateJobRequest() (request *SubmitImageAnimateJobRequest) {
    request = &SubmitImageAnimateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitImageAnimateJob")
    
    
    return
}

func NewSubmitImageAnimateJobResponse() (response *SubmitImageAnimateJobResponse) {
    response = &SubmitImageAnimateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitImageAnimateJob
// 用于提交图片跳舞任务。图片跳舞能力支持舞蹈动作结合图片生成跳舞视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageAnimateJob(request *SubmitImageAnimateJobRequest) (response *SubmitImageAnimateJobResponse, err error) {
    return c.SubmitImageAnimateJobWithContext(context.Background(), request)
}

// SubmitImageAnimateJob
// 用于提交图片跳舞任务。图片跳舞能力支持舞蹈动作结合图片生成跳舞视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageAnimateJobWithContext(ctx context.Context, request *SubmitImageAnimateJobRequest) (response *SubmitImageAnimateJobResponse, err error) {
    if request == nil {
        request = NewSubmitImageAnimateJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitImageAnimateJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageAnimateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageAnimateJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitImageToVideoGeneralJobRequest() (request *SubmitImageToVideoGeneralJobRequest) {
    request = &SubmitImageToVideoGeneralJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitImageToVideoGeneralJob")
    
    
    return
}

func NewSubmitImageToVideoGeneralJobResponse() (response *SubmitImageToVideoGeneralJobResponse) {
    response = &SubmitImageToVideoGeneralJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitImageToVideoGeneralJob
// 图生视频通用能力接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageToVideoGeneralJob(request *SubmitImageToVideoGeneralJobRequest) (response *SubmitImageToVideoGeneralJobResponse, err error) {
    return c.SubmitImageToVideoGeneralJobWithContext(context.Background(), request)
}

// SubmitImageToVideoGeneralJob
// 图生视频通用能力接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageToVideoGeneralJobWithContext(ctx context.Context, request *SubmitImageToVideoGeneralJobRequest) (response *SubmitImageToVideoGeneralJobResponse, err error) {
    if request == nil {
        request = NewSubmitImageToVideoGeneralJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitImageToVideoGeneralJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageToVideoGeneralJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageToVideoGeneralJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitImageToVideoJobRequest() (request *SubmitImageToVideoJobRequest) {
    request = &SubmitImageToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitImageToVideoJob")
    
    
    return
}

func NewSubmitImageToVideoJobResponse() (response *SubmitImageToVideoJobResponse) {
    response = &SubmitImageToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitImageToVideoJob
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageToVideoJob(request *SubmitImageToVideoJobRequest) (response *SubmitImageToVideoJobResponse, err error) {
    return c.SubmitImageToVideoJobWithContext(context.Background(), request)
}

// SubmitImageToVideoJob
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageToVideoJobWithContext(ctx context.Context, request *SubmitImageToVideoJobRequest) (response *SubmitImageToVideoJobResponse, err error) {
    if request == nil {
        request = NewSubmitImageToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitImageToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitImageToVideoViduJobRequest() (request *SubmitImageToVideoViduJobRequest) {
    request = &SubmitImageToVideoViduJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitImageToVideoViduJob")
    
    
    return
}

func NewSubmitImageToVideoViduJobResponse() (response *SubmitImageToVideoViduJobResponse) {
    response = &SubmitImageToVideoViduJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitImageToVideoViduJob
// 提交Vidu图生视频任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageToVideoViduJob(request *SubmitImageToVideoViduJobRequest) (response *SubmitImageToVideoViduJobResponse, err error) {
    return c.SubmitImageToVideoViduJobWithContext(context.Background(), request)
}

// SubmitImageToVideoViduJob
// 提交Vidu图生视频任务接口
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageToVideoViduJobWithContext(ctx context.Context, request *SubmitImageToVideoViduJobRequest) (response *SubmitImageToVideoViduJobResponse, err error) {
    if request == nil {
        request = NewSubmitImageToVideoViduJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitImageToVideoViduJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageToVideoViduJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageToVideoViduJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitMotionControlKlingJobRequest() (request *SubmitMotionControlKlingJobRequest) {
    request = &SubmitMotionControlKlingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitMotionControlKlingJob")
    
    
    return
}

func NewSubmitMotionControlKlingJobResponse() (response *SubmitMotionControlKlingJobResponse) {
    response = &SubmitMotionControlKlingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitMotionControlKlingJob
// 提交动作控制(Kling)任务并发
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitMotionControlKlingJob(request *SubmitMotionControlKlingJobRequest) (response *SubmitMotionControlKlingJobResponse, err error) {
    return c.SubmitMotionControlKlingJobWithContext(context.Background(), request)
}

// SubmitMotionControlKlingJob
// 提交动作控制(Kling)任务并发
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_BODYJOINTSFAIL = "FailedOperation.BodyJointsFail"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEBODYJOINSUNDETECTED = "FailedOperation.ImageBodyJoinsUndetected"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERATIOEXCCEED = "FailedOperation.ImageRatioExcceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBQUEUEFULL = "FailedOperation.JobQueueFull"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  LIMITEXCEEDED = "LimitExceeded"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitMotionControlKlingJobWithContext(ctx context.Context, request *SubmitMotionControlKlingJobRequest) (response *SubmitMotionControlKlingJobResponse, err error) {
    if request == nil {
        request = NewSubmitMotionControlKlingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitMotionControlKlingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitMotionControlKlingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitMotionControlKlingJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitPortraitSingJobRequest() (request *SubmitPortraitSingJobRequest) {
    request = &SubmitPortraitSingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitPortraitSingJob")
    
    
    return
}

func NewSubmitPortraitSingJobResponse() (response *SubmitPortraitSingJobResponse) {
    response = &SubmitPortraitSingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitPortraitSingJob
// 用于提交图片唱演任务。
//
// 支持提交音频和图片生成唱演视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_SUBMITASSPFAILED = "FailedOperation.SubmitAsspFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDAUDIOFORMAT = "InvalidParameterValue.InvalidAudioFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESIZE = "InvalidParameterValue.InvalidImageSize"
//  INVALIDPARAMETERVALUE_INVALIDMODEL = "InvalidParameterValue.InvalidModel"
//  INVALIDPARAMETERVALUE_TOOLARGEFACEANGLE = "InvalidParameterValue.TooLargeFaceAngle"
//  INVALIDPARAMETERVALUE_TOOLOWFACEQUALITY = "InvalidParameterValue.TooLowFaceQuality"
//  INVALIDPARAMETERVALUE_TOOMANYFACES = "InvalidParameterValue.TooManyFaces"
//  INVALIDPARAMETERVALUE_TOOMUCHFACEOCCLUSION = "InvalidParameterValue.TooMuchFaceOcclusion"
//  INVALIDPARAMETERVALUE_TOOSMALLFACESIZE = "InvalidParameterValue.TooSmallFaceSize"
func (c *Client) SubmitPortraitSingJob(request *SubmitPortraitSingJobRequest) (response *SubmitPortraitSingJobResponse, err error) {
    return c.SubmitPortraitSingJobWithContext(context.Background(), request)
}

// SubmitPortraitSingJob
// 用于提交图片唱演任务。
//
// 支持提交音频和图片生成唱演视频，满足社交娱乐、互动营销等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_SUBMITASSPFAILED = "FailedOperation.SubmitAsspFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDAUDIOFORMAT = "InvalidParameterValue.InvalidAudioFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESIZE = "InvalidParameterValue.InvalidImageSize"
//  INVALIDPARAMETERVALUE_INVALIDMODEL = "InvalidParameterValue.InvalidModel"
//  INVALIDPARAMETERVALUE_TOOLARGEFACEANGLE = "InvalidParameterValue.TooLargeFaceAngle"
//  INVALIDPARAMETERVALUE_TOOLOWFACEQUALITY = "InvalidParameterValue.TooLowFaceQuality"
//  INVALIDPARAMETERVALUE_TOOMANYFACES = "InvalidParameterValue.TooManyFaces"
//  INVALIDPARAMETERVALUE_TOOMUCHFACEOCCLUSION = "InvalidParameterValue.TooMuchFaceOcclusion"
//  INVALIDPARAMETERVALUE_TOOSMALLFACESIZE = "InvalidParameterValue.TooSmallFaceSize"
func (c *Client) SubmitPortraitSingJobWithContext(ctx context.Context, request *SubmitPortraitSingJobRequest) (response *SubmitPortraitSingJobResponse, err error) {
    if request == nil {
        request = NewSubmitPortraitSingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitPortraitSingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitPortraitSingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitPortraitSingJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitReferenceToVideoViduJobRequest() (request *SubmitReferenceToVideoViduJobRequest) {
    request = &SubmitReferenceToVideoViduJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitReferenceToVideoViduJob")
    
    
    return
}

func NewSubmitReferenceToVideoViduJobResponse() (response *SubmitReferenceToVideoViduJobResponse) {
    response = &SubmitReferenceToVideoViduJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitReferenceToVideoViduJob
// 提交Vidu参考生视频任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_SUBMITASSPFAILED = "FailedOperation.SubmitAsspFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDAUDIOFORMAT = "InvalidParameterValue.InvalidAudioFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESIZE = "InvalidParameterValue.InvalidImageSize"
//  INVALIDPARAMETERVALUE_INVALIDMODEL = "InvalidParameterValue.InvalidModel"
//  INVALIDPARAMETERVALUE_TOOLARGEFACEANGLE = "InvalidParameterValue.TooLargeFaceAngle"
//  INVALIDPARAMETERVALUE_TOOLOWFACEQUALITY = "InvalidParameterValue.TooLowFaceQuality"
//  INVALIDPARAMETERVALUE_TOOMANYFACES = "InvalidParameterValue.TooManyFaces"
//  INVALIDPARAMETERVALUE_TOOMUCHFACEOCCLUSION = "InvalidParameterValue.TooMuchFaceOcclusion"
//  INVALIDPARAMETERVALUE_TOOSMALLFACESIZE = "InvalidParameterValue.TooSmallFaceSize"
func (c *Client) SubmitReferenceToVideoViduJob(request *SubmitReferenceToVideoViduJobRequest) (response *SubmitReferenceToVideoViduJobResponse, err error) {
    return c.SubmitReferenceToVideoViduJobWithContext(context.Background(), request)
}

// SubmitReferenceToVideoViduJob
// 提交Vidu参考生视频任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_IMAGEFACEDETECTFAILED = "FailedOperation.ImageFaceDetectFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONAUDIOFAILED = "FailedOperation.ModerationAudioFailed"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_SUBMITASSPFAILED = "FailedOperation.SubmitAsspFailed"
//  INVALIDPARAMETERVALUE_INVALIDAUDIODURATION = "InvalidParameterValue.InvalidAudioDuration"
//  INVALIDPARAMETERVALUE_INVALIDAUDIOFORMAT = "InvalidParameterValue.InvalidAudioFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEASPECTRATIO = "InvalidParameterValue.InvalidImageAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDIMAGEFORMAT = "InvalidParameterValue.InvalidImageFormat"
//  INVALIDPARAMETERVALUE_INVALIDIMAGERESOLUTION = "InvalidParameterValue.InvalidImageResolution"
//  INVALIDPARAMETERVALUE_INVALIDIMAGESIZE = "InvalidParameterValue.InvalidImageSize"
//  INVALIDPARAMETERVALUE_INVALIDMODEL = "InvalidParameterValue.InvalidModel"
//  INVALIDPARAMETERVALUE_TOOLARGEFACEANGLE = "InvalidParameterValue.TooLargeFaceAngle"
//  INVALIDPARAMETERVALUE_TOOLOWFACEQUALITY = "InvalidParameterValue.TooLowFaceQuality"
//  INVALIDPARAMETERVALUE_TOOMANYFACES = "InvalidParameterValue.TooManyFaces"
//  INVALIDPARAMETERVALUE_TOOMUCHFACEOCCLUSION = "InvalidParameterValue.TooMuchFaceOcclusion"
//  INVALIDPARAMETERVALUE_TOOSMALLFACESIZE = "InvalidParameterValue.TooSmallFaceSize"
func (c *Client) SubmitReferenceToVideoViduJobWithContext(ctx context.Context, request *SubmitReferenceToVideoViduJobRequest) (response *SubmitReferenceToVideoViduJobResponse, err error) {
    if request == nil {
        request = NewSubmitReferenceToVideoViduJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitReferenceToVideoViduJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitReferenceToVideoViduJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitReferenceToVideoViduJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTemplateToVideoJobRequest() (request *SubmitTemplateToVideoJobRequest) {
    request = &SubmitTemplateToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitTemplateToVideoJob")
    
    
    return
}

func NewSubmitTemplateToVideoJobResponse() (response *SubmitTemplateToVideoJobResponse) {
    response = &SubmitTemplateToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTemplateToVideoJob
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTemplateToVideoJob(request *SubmitTemplateToVideoJobRequest) (response *SubmitTemplateToVideoJobResponse, err error) {
    return c.SubmitTemplateToVideoJobWithContext(context.Background(), request)
}

// SubmitTemplateToVideoJob
// 提交视频特效任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTemplateToVideoJobWithContext(ctx context.Context, request *SubmitTemplateToVideoJobRequest) (response *SubmitTemplateToVideoJobResponse, err error) {
    if request == nil {
        request = NewSubmitTemplateToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitTemplateToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTemplateToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTemplateToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTextToVideoJobRequest() (request *SubmitTextToVideoJobRequest) {
    request = &SubmitTextToVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitTextToVideoJob")
    
    
    return
}

func NewSubmitTextToVideoJobResponse() (response *SubmitTextToVideoJobResponse) {
    response = &SubmitTextToVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTextToVideoJob
// 通过提交对视频内容的描述文本生成一个短视频。文生视频为异步处理任务，成功提交任务后返回任务的JobId。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTextToVideoJob(request *SubmitTextToVideoJobRequest) (response *SubmitTextToVideoJobResponse, err error) {
    return c.SubmitTextToVideoJobWithContext(context.Background(), request)
}

// SubmitTextToVideoJob
// 通过提交对视频内容的描述文本生成一个短视频。文生视频为异步处理任务，成功提交任务后返回任务的JobId。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTextToVideoJobWithContext(ctx context.Context, request *SubmitTextToVideoJobRequest) (response *SubmitTextToVideoJobResponse, err error) {
    if request == nil {
        request = NewSubmitTextToVideoJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitTextToVideoJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTextToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTextToVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTextToVideoViduJobRequest() (request *SubmitTextToVideoViduJobRequest) {
    request = &SubmitTextToVideoViduJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitTextToVideoViduJob")
    
    
    return
}

func NewSubmitTextToVideoViduJobResponse() (response *SubmitTextToVideoViduJobResponse) {
    response = &SubmitTextToVideoViduJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTextToVideoViduJob
// 提交Vidu文生视频任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTextToVideoViduJob(request *SubmitTextToVideoViduJobRequest) (response *SubmitTextToVideoViduJobResponse, err error) {
    return c.SubmitTextToVideoViduJobWithContext(context.Background(), request)
}

// SubmitTextToVideoViduJob
// 提交Vidu文生视频任务接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTextToVideoViduJobWithContext(ctx context.Context, request *SubmitTextToVideoViduJobRequest) (response *SubmitTextToVideoViduJobResponse, err error) {
    if request == nil {
        request = NewSubmitTextToVideoViduJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitTextToVideoViduJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTextToVideoViduJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTextToVideoViduJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoEditJobRequest() (request *SubmitVideoEditJobRequest) {
    request = &SubmitVideoEditJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitVideoEditJob")
    
    
    return
}

func NewSubmitVideoEditJobResponse() (response *SubmitVideoEditJobResponse) {
    response = &SubmitVideoEditJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoEditJob
// 用于提交视频编辑任务，支持上传视频、文本及图片素材开展编辑操作，涵盖风格迁移、元素替换、内容增减等核心能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoEditJob(request *SubmitVideoEditJobRequest) (response *SubmitVideoEditJobResponse, err error) {
    return c.SubmitVideoEditJobWithContext(context.Background(), request)
}

// SubmitVideoEditJob
// 用于提交视频编辑任务，支持上传视频、文本及图片素材开展编辑操作，涵盖风格迁移、元素替换、内容增减等核心能力。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoEditJobWithContext(ctx context.Context, request *SubmitVideoEditJobRequest) (response *SubmitVideoEditJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoEditJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitVideoEditJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoEditJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoEditJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoEditKlingJobRequest() (request *SubmitVideoEditKlingJobRequest) {
    request = &SubmitVideoEditKlingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitVideoEditKlingJob")
    
    
    return
}

func NewSubmitVideoEditKlingJobResponse() (response *SubmitVideoEditKlingJobResponse) {
    response = &SubmitVideoEditKlingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoEditKlingJob
// 提交Kling多模态编辑任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoEditKlingJob(request *SubmitVideoEditKlingJobRequest) (response *SubmitVideoEditKlingJobResponse, err error) {
    return c.SubmitVideoEditKlingJobWithContext(context.Background(), request)
}

// SubmitVideoEditKlingJob
// 提交Kling多模态编辑任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoEditKlingJobWithContext(ctx context.Context, request *SubmitVideoEditKlingJobRequest) (response *SubmitVideoEditKlingJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoEditKlingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitVideoEditKlingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoEditKlingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoEditKlingJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoExtendKlingJobRequest() (request *SubmitVideoExtendKlingJobRequest) {
    request = &SubmitVideoExtendKlingJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitVideoExtendKlingJob")
    
    
    return
}

func NewSubmitVideoExtendKlingJobResponse() (response *SubmitVideoExtendKlingJobResponse) {
    response = &SubmitVideoExtendKlingJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoExtendKlingJob
// 用于提交视频延长任务接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoExtendKlingJob(request *SubmitVideoExtendKlingJobRequest) (response *SubmitVideoExtendKlingJobResponse, err error) {
    return c.SubmitVideoExtendKlingJobWithContext(context.Background(), request)
}

// SubmitVideoExtendKlingJob
// 用于提交视频延长任务接口。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  MISSINGPARAMETER = "MissingParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoExtendKlingJobWithContext(ctx context.Context, request *SubmitVideoExtendKlingJobRequest) (response *SubmitVideoExtendKlingJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoExtendKlingJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitVideoExtendKlingJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoExtendKlingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoExtendKlingJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoFaceFusionJobRequest() (request *SubmitVideoFaceFusionJobRequest) {
    request = &SubmitVideoFaceFusionJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitVideoFaceFusionJob")
    
    
    return
}

func NewSubmitVideoFaceFusionJobResponse() (response *SubmitVideoFaceFusionJobResponse) {
    response = &SubmitVideoFaceFusionJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoFaceFusionJob
// 提交视频人脸融合任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
func (c *Client) SubmitVideoFaceFusionJob(request *SubmitVideoFaceFusionJobRequest) (response *SubmitVideoFaceFusionJobResponse, err error) {
    return c.SubmitVideoFaceFusionJobWithContext(context.Background(), request)
}

// SubmitVideoFaceFusionJob
// 提交视频人脸融合任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
func (c *Client) SubmitVideoFaceFusionJobWithContext(ctx context.Context, request *SubmitVideoFaceFusionJobRequest) (response *SubmitVideoFaceFusionJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoFaceFusionJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitVideoFaceFusionJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoFaceFusionJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoFaceFusionJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoStylizationJobRequest() (request *SubmitVideoStylizationJobRequest) {
    request = &SubmitVideoStylizationJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitVideoStylizationJob")
    
    
    return
}

func NewSubmitVideoStylizationJobResponse() (response *SubmitVideoStylizationJobResponse) {
    response = &SubmitVideoStylizationJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoStylizationJob
// 用于提交视频风格化任务。支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOASPECTRATIO = "InvalidParameterValue.InvalidVideoAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDVIDEODURATION = "InvalidParameterValue.InvalidVideoDuration"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFPS = "InvalidParameterValue.InvalidVideoFPS"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFORMAT = "InvalidParameterValue.InvalidVideoFormat"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLENOTEXIST = "InvalidParameterValue.StyleNotExist"
//  INVALIDPARAMETERVALUE_STYLESTRENGTHNOTEXIST = "InvalidParameterValue.StyleStrengthNotExist"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_VIDEOSIZEEXCEED = "InvalidParameterValue.VideoSizeExceed"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitVideoStylizationJob(request *SubmitVideoStylizationJobRequest) (response *SubmitVideoStylizationJobResponse, err error) {
    return c.SubmitVideoStylizationJobWithContext(context.Background(), request)
}

// SubmitVideoStylizationJob
// 用于提交视频风格化任务。支持将输入视频生成特定风格的视频。生成后的视频画面风格多样、流畅自然，能够满足社交娱乐、互动营销、视频素材制作等场景的需求。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOASPECTRATIO = "InvalidParameterValue.InvalidVideoAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDVIDEODURATION = "InvalidParameterValue.InvalidVideoDuration"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFPS = "InvalidParameterValue.InvalidVideoFPS"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFORMAT = "InvalidParameterValue.InvalidVideoFormat"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLENOTEXIST = "InvalidParameterValue.StyleNotExist"
//  INVALIDPARAMETERVALUE_STYLESTRENGTHNOTEXIST = "InvalidParameterValue.StyleStrengthNotExist"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_VIDEOSIZEEXCEED = "InvalidParameterValue.VideoSizeExceed"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitVideoStylizationJobWithContext(ctx context.Context, request *SubmitVideoStylizationJobRequest) (response *SubmitVideoStylizationJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoStylizationJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitVideoStylizationJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoStylizationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoStylizationJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitVideoVoiceJobRequest() (request *SubmitVideoVoiceJobRequest) {
    request = &SubmitVideoVoiceJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitVideoVoiceJob")
    
    
    return
}

func NewSubmitVideoVoiceJobResponse() (response *SubmitVideoVoiceJobResponse) {
    response = &SubmitVideoVoiceJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoVoiceJob
// 提交视频配音效任务，输入视频后提交请求，会返回一个JobId，用于查询视频配音效的处理进度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOASPECTRATIO = "InvalidParameterValue.InvalidVideoAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDVIDEODURATION = "InvalidParameterValue.InvalidVideoDuration"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFPS = "InvalidParameterValue.InvalidVideoFPS"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFORMAT = "InvalidParameterValue.InvalidVideoFormat"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLENOTEXIST = "InvalidParameterValue.StyleNotExist"
//  INVALIDPARAMETERVALUE_STYLESTRENGTHNOTEXIST = "InvalidParameterValue.StyleStrengthNotExist"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_VIDEOSIZEEXCEED = "InvalidParameterValue.VideoSizeExceed"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitVideoVoiceJob(request *SubmitVideoVoiceJobRequest) (response *SubmitVideoVoiceJobResponse, err error) {
    return c.SubmitVideoVoiceJobWithContext(context.Background(), request)
}

// SubmitVideoVoiceJob
// 提交视频配音效任务，输入视频后提交请求，会返回一个JobId，用于查询视频配音效的处理进度。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_VIDEODECODEFAILED = "FailedOperation.VideoDecodeFailed"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOASPECTRATIO = "InvalidParameterValue.InvalidVideoAspectRatio"
//  INVALIDPARAMETERVALUE_INVALIDVIDEODURATION = "InvalidParameterValue.InvalidVideoDuration"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFPS = "InvalidParameterValue.InvalidVideoFPS"
//  INVALIDPARAMETERVALUE_INVALIDVIDEOFORMAT = "InvalidParameterValue.InvalidVideoFormat"
//  INVALIDPARAMETERVALUE_INVALIDVIDEORESOLUTION = "InvalidParameterValue.InvalidVideoResolution"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLENOTEXIST = "InvalidParameterValue.StyleNotExist"
//  INVALIDPARAMETERVALUE_STYLESTRENGTHNOTEXIST = "InvalidParameterValue.StyleStrengthNotExist"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_VIDEOSIZEEXCEED = "InvalidParameterValue.VideoSizeExceed"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SubmitVideoVoiceJobWithContext(ctx context.Context, request *SubmitVideoVoiceJobRequest) (response *SubmitVideoVoiceJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoVoiceJobRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "vclm", APIVersion, "SubmitVideoVoiceJob")
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoVoiceJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoVoiceJobResponse()
    err = c.Send(request, response)
    return
}
