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
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CheckAnimateImageJobWithContext(ctx context.Context, request *CheckAnimateImageJobRequest) (response *CheckAnimateImageJobResponse, err error) {
    if request == nil {
        request = NewCheckAnimateImageJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CheckAnimateImageJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCheckAnimateImageJobResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImageAnimateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImageAnimateJobResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribePortraitSingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribePortraitSingJobResponse()
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
func (c *Client) DescribeTemplateToVideoJob(request *DescribeTemplateToVideoJobRequest) (response *DescribeTemplateToVideoJobResponse, err error) {
    return c.DescribeTemplateToVideoJobWithContext(context.Background(), request)
}

// DescribeTemplateToVideoJob
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
func (c *Client) DescribeTemplateToVideoJobWithContext(ctx context.Context, request *DescribeTemplateToVideoJobRequest) (response *DescribeTemplateToVideoJobResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateToVideoJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTemplateToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTemplateToVideoJobResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoStylizationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoStylizationJobResponse()
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
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) SubmitImageAnimateJobWithContext(ctx context.Context, request *SubmitImageAnimateJobRequest) (response *SubmitImageAnimateJobResponse, err error) {
    if request == nil {
        request = NewSubmitImageAnimateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitImageAnimateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitImageAnimateJobResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitPortraitSingJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitPortraitSingJobResponse()
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
func (c *Client) SubmitTemplateToVideoJob(request *SubmitTemplateToVideoJobRequest) (response *SubmitTemplateToVideoJobResponse, err error) {
    return c.SubmitTemplateToVideoJobWithContext(context.Background(), request)
}

// SubmitTemplateToVideoJob
// 提交视频特效任务接口
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
func (c *Client) SubmitTemplateToVideoJobWithContext(ctx context.Context, request *SubmitTemplateToVideoJobRequest) (response *SubmitTemplateToVideoJobResponse, err error) {
    if request == nil {
        request = NewSubmitTemplateToVideoJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTemplateToVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTemplateToVideoJobResponse()
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoStylizationJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoStylizationJobResponse()
    err = c.Send(request, response)
    return
}
