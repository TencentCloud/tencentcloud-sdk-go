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


func NewConfirmVideoTranslateJobRequest() (request *ConfirmVideoTranslateJobRequest) {
    request = &ConfirmVideoTranslateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "ConfirmVideoTranslateJob")
    
    
    return
}

func NewConfirmVideoTranslateJobResponse() (response *ConfirmVideoTranslateJobResponse) {
    response = &ConfirmVideoTranslateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ConfirmVideoTranslateJob
// 确认视频转译结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSNOTFINISHED = "FailedOperation.AudioProcessNotFinished"
//  FAILEDOPERATION_CONFIRMTASKEXCEPTION = "FailedOperation.ConfirmTaskException"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_TEXTMODERATIONNOTPASS = "FailedOperation.TextModerationNotPass"
//  FAILEDOPERATION_TRANSLATIONCONFIRMHASFINISHED = "FailedOperation.TranslationConfirmHasFinished"
//  FAILEDOPERATION_TRANSLATIONNOTNEEDCONFIRM = "FailedOperation.TranslationNotNeedConfirm"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) ConfirmVideoTranslateJob(request *ConfirmVideoTranslateJobRequest) (response *ConfirmVideoTranslateJobResponse, err error) {
    return c.ConfirmVideoTranslateJobWithContext(context.Background(), request)
}

// ConfirmVideoTranslateJob
// 确认视频转译结果
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSNOTFINISHED = "FailedOperation.AudioProcessNotFinished"
//  FAILEDOPERATION_CONFIRMTASKEXCEPTION = "FailedOperation.ConfirmTaskException"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_TEXTMODERATIONNOTPASS = "FailedOperation.TextModerationNotPass"
//  FAILEDOPERATION_TRANSLATIONCONFIRMHASFINISHED = "FailedOperation.TranslationConfirmHasFinished"
//  FAILEDOPERATION_TRANSLATIONNOTNEEDCONFIRM = "FailedOperation.TranslationNotNeedConfirm"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) ConfirmVideoTranslateJobWithContext(ctx context.Context, request *ConfirmVideoTranslateJobRequest) (response *ConfirmVideoTranslateJobResponse, err error) {
    if request == nil {
        request = NewConfirmVideoTranslateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ConfirmVideoTranslateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewConfirmVideoTranslateJobResponse()
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
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

func NewDescribeVideoTranslateJobRequest() (request *DescribeVideoTranslateJobRequest) {
    request = &DescribeVideoTranslateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "DescribeVideoTranslateJob")
    
    
    return
}

func NewDescribeVideoTranslateJobResponse() (response *DescribeVideoTranslateJobResponse) {
    response = &DescribeVideoTranslateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeVideoTranslateJob
// 查询视频翻译任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSFAILED = "FailedOperation.AudioProcessFailed"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DescribeVideoTranslateJob(request *DescribeVideoTranslateJobRequest) (response *DescribeVideoTranslateJobResponse, err error) {
    return c.DescribeVideoTranslateJobWithContext(context.Background(), request)
}

// DescribeVideoTranslateJob
// 查询视频翻译任务
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIOPROCESSFAILED = "FailedOperation.AudioProcessFailed"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_UINLIMITEXCEEDED = "RequestLimitExceeded.UinLimitExceeded"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DescribeVideoTranslateJobWithContext(ctx context.Context, request *DescribeVideoTranslateJobRequest) (response *DescribeVideoTranslateJobResponse, err error) {
    if request == nil {
        request = NewDescribeVideoTranslateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeVideoTranslateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeVideoTranslateJobResponse()
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
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
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
//  INVALIDPARAMETER_TEMPLATENOTEXISTED = "InvalidParameter.TemplateNotExisted"
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

func NewSubmitVideoTranslateJobRequest() (request *SubmitVideoTranslateJobRequest) {
    request = &SubmitVideoTranslateJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("vclm", APIVersion, "SubmitVideoTranslateJob")
    
    
    return
}

func NewSubmitVideoTranslateJobResponse() (response *SubmitVideoTranslateJobResponse) {
    response = &SubmitVideoTranslateJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitVideoTranslateJob
// ###### 支持音色种别列表
//
// | 音色名称                 | 性别 | 目标语言         | 音色ID |
//
// | ------------------------ | ---- | ---------------- | ------ |
//
// | Florian Multilingual     | 男 | 德语(德国)       | 701001 |
//
// | Seraphina                | 女  | 德语(德国)       | 701002 |
//
// | Ada Multilingual         | 女 | 英语(英国)       | 701003 |
//
// | Ollie Multilingual       | 男  | 英语(英国)       | 701004 |
//
// | Ava Multilingual         | 女 | 英语(美国)       | 701005 |
//
// | Andrew Multilingual      | 男  | 英语(美国)       | 701006 |
//
// | Emma Multilingual        | 女  | 英语(美国)       | 701007 |
//
// | Brian Multilingual       | 男  | 英语(美国)       | 701008 |
//
// | Jenny Multilingual       | 女  | 英语(美国)       | 701009 |
//
// | Ryan Multilingual        | 男  | 英语(美国)       | 701010 |
//
// | Adam Multilingual        | 男  | 英语(美国)       | 701011 |
//
// | AlloyTurbo Multilingual  | 男  | 英语(美国)       | 701012 |
//
// | Amanda Multilingual      | 女  | 英语(美国)       | 701013 |
//
// | Brandon Multilingual     | 男  | 英语(美国)       | 701014 |
//
// | Christopher Multilingual | 男  | 英语(美国)       | 701015 |
//
// | Cora Multilingual        | 女  | 英语(美国)       | 701016 |
//
// | Davis Multilingual       | 男  | 英语(美国)       | 701017 |
//
// | Derek Multilingual       | 男  | 英语(美国)       | 701018 |
//
// | Dustin Multilingual      | 男  | 英语(美国)       | 701019 |
//
// | Evelyn Multilingual      | 女  | 英语(美国)       | 701020 |
//
// | Lewis Multilingual       | 男  | 英语(美国)       | 701021 |
//
// | Lola Multilingual        | 女  | 英语(美国)       | 701022 |
//
// | Nancy Multilingual       | 女  | 英语(美国)       | 701023 |
//
// | NovaTurbo Multilingual   | 女   | 英语(美国)       | 701024 |
//
// | Phoebe Multilingual      | 女  | 英语(美国)       | 701025 |
//
// | Samuel Multilingual      | 男  | 英语(美国)       | 701026 |
//
// | Serena Multilingual      | 女  | 英语(美国)       | 701027 |
//
// | Steffan Multilingual     | 男  | 英语(美国)       | 701028 |
//
// | Arabella Multilingual    | 女  | 西班牙语(西班牙) | 701029 |
//
// | Isidora Multilingual     | 女  | 西班牙语(西班牙) | 701030 |
//
// | Tristan Multilingual     | 男  | 西班牙语(西班牙) | 701031 |
//
// | Ximena Multilingual      | 女  | 西班牙语(西班牙) | 701032 |
//
// | Remy Multilingual        | 男  | 法语(法国)       | 701033 |
//
// | Vivienne Multilingual    | 女  | 法语(法国)       | 701034 |
//
// | Lucien Multilingual      | 男  | 法语(法国)       | 701035 |
//
// | Alessio Multilingual     | 男  | 意大利语(意大利) | 701036 |
//
// | Giuseppe Multilingual    | 男  | 意大利语(意大利) | 701037 |
//
// | Isabella Multilingual    | 女  | 意大利语(意大利) | 701038 |
//
// | Marcello Multilingual    | 男  | 意大利语(意大利) | 701039 |
//
// | Masaru Multilingual      | 男  | 日语(日本)       | 701040 |
//
// | Hyunsu Multilingual      | 男  | 韩语(韩国)       | 701041 |
//
// | Macerio Multilingual     | 男  | 葡萄牙语(巴西)   | 701042 |
//
// | Thalita Multilingual     | 女  | 葡萄牙语(巴西)   | 701043 |
//
// | 晓辰 多语言              | 女  | 中文(普通话)     | 701044 |
//
// | 晓晓 多语言              | 女  | 中文(普通话)     | 701045 |
//
// | 晓宇 多语言              | 女  | 中文(普通话)     | 701046 |
//
// | 云逸 多语言              | 男 | 中文(普通话)     | 701047 |
//
// | Yunfan Multilingual      | 男  | 中文(普通话)     | 701048 |
//
// | Yunxiao Multilingual     | 男  | 中文(普通话)     | 701049 |
//
// | 晓晓 方言                | 女  | 中文(普通话)     | 701050 |
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_VIDEODURATIONEXCEED = "FailedOperation.VideoDurationExceed"
//  FAILEDOPERATION_VIDEOFPSEXCEED = "FailedOperation.VideoFpsExceed"
//  FAILEDOPERATION_VIDEORESOLUTIONEXCEED = "FailedOperation.VideoResolutionExceed"
//  FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoTranslateJob(request *SubmitVideoTranslateJobRequest) (response *SubmitVideoTranslateJobResponse, err error) {
    return c.SubmitVideoTranslateJobWithContext(context.Background(), request)
}

// SubmitVideoTranslateJob
// ###### 支持音色种别列表
//
// | 音色名称                 | 性别 | 目标语言         | 音色ID |
//
// | ------------------------ | ---- | ---------------- | ------ |
//
// | Florian Multilingual     | 男 | 德语(德国)       | 701001 |
//
// | Seraphina                | 女  | 德语(德国)       | 701002 |
//
// | Ada Multilingual         | 女 | 英语(英国)       | 701003 |
//
// | Ollie Multilingual       | 男  | 英语(英国)       | 701004 |
//
// | Ava Multilingual         | 女 | 英语(美国)       | 701005 |
//
// | Andrew Multilingual      | 男  | 英语(美国)       | 701006 |
//
// | Emma Multilingual        | 女  | 英语(美国)       | 701007 |
//
// | Brian Multilingual       | 男  | 英语(美国)       | 701008 |
//
// | Jenny Multilingual       | 女  | 英语(美国)       | 701009 |
//
// | Ryan Multilingual        | 男  | 英语(美国)       | 701010 |
//
// | Adam Multilingual        | 男  | 英语(美国)       | 701011 |
//
// | AlloyTurbo Multilingual  | 男  | 英语(美国)       | 701012 |
//
// | Amanda Multilingual      | 女  | 英语(美国)       | 701013 |
//
// | Brandon Multilingual     | 男  | 英语(美国)       | 701014 |
//
// | Christopher Multilingual | 男  | 英语(美国)       | 701015 |
//
// | Cora Multilingual        | 女  | 英语(美国)       | 701016 |
//
// | Davis Multilingual       | 男  | 英语(美国)       | 701017 |
//
// | Derek Multilingual       | 男  | 英语(美国)       | 701018 |
//
// | Dustin Multilingual      | 男  | 英语(美国)       | 701019 |
//
// | Evelyn Multilingual      | 女  | 英语(美国)       | 701020 |
//
// | Lewis Multilingual       | 男  | 英语(美国)       | 701021 |
//
// | Lola Multilingual        | 女  | 英语(美国)       | 701022 |
//
// | Nancy Multilingual       | 女  | 英语(美国)       | 701023 |
//
// | NovaTurbo Multilingual   | 女   | 英语(美国)       | 701024 |
//
// | Phoebe Multilingual      | 女  | 英语(美国)       | 701025 |
//
// | Samuel Multilingual      | 男  | 英语(美国)       | 701026 |
//
// | Serena Multilingual      | 女  | 英语(美国)       | 701027 |
//
// | Steffan Multilingual     | 男  | 英语(美国)       | 701028 |
//
// | Arabella Multilingual    | 女  | 西班牙语(西班牙) | 701029 |
//
// | Isidora Multilingual     | 女  | 西班牙语(西班牙) | 701030 |
//
// | Tristan Multilingual     | 男  | 西班牙语(西班牙) | 701031 |
//
// | Ximena Multilingual      | 女  | 西班牙语(西班牙) | 701032 |
//
// | Remy Multilingual        | 男  | 法语(法国)       | 701033 |
//
// | Vivienne Multilingual    | 女  | 法语(法国)       | 701034 |
//
// | Lucien Multilingual      | 男  | 法语(法国)       | 701035 |
//
// | Alessio Multilingual     | 男  | 意大利语(意大利) | 701036 |
//
// | Giuseppe Multilingual    | 男  | 意大利语(意大利) | 701037 |
//
// | Isabella Multilingual    | 女  | 意大利语(意大利) | 701038 |
//
// | Marcello Multilingual    | 男  | 意大利语(意大利) | 701039 |
//
// | Masaru Multilingual      | 男  | 日语(日本)       | 701040 |
//
// | Hyunsu Multilingual      | 男  | 韩语(韩国)       | 701041 |
//
// | Macerio Multilingual     | 男  | 葡萄牙语(巴西)   | 701042 |
//
// | Thalita Multilingual     | 女  | 葡萄牙语(巴西)   | 701043 |
//
// | 晓辰 多语言              | 女  | 中文(普通话)     | 701044 |
//
// | 晓晓 多语言              | 女  | 中文(普通话)     | 701045 |
//
// | 晓宇 多语言              | 女  | 中文(普通话)     | 701046 |
//
// | 云逸 多语言              | 男 | 中文(普通话)     | 701047 |
//
// | Yunfan Multilingual      | 男  | 中文(普通话)     | 701048 |
//
// | Yunxiao Multilingual     | 男  | 中文(普通话)     | 701049 |
//
// | 晓晓 方言                | 女  | 中文(普通话)     | 701050 |
//
// 可能返回的错误码:
//  FAILEDOPERATION_AUDIODECODEFAILED = "FailedOperation.AudioDecodeFailed"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownloadError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_VIDEODURATIONEXCEED = "FailedOperation.VideoDurationExceed"
//  FAILEDOPERATION_VIDEOFPSEXCEED = "FailedOperation.VideoFpsExceed"
//  FAILEDOPERATION_VIDEORESOLUTIONEXCEED = "FailedOperation.VideoResolutionExceed"
//  FAILEDOPERATION_VIDEOSIZEEXCEED = "FailedOperation.VideoSizeExceed"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitVideoTranslateJobWithContext(ctx context.Context, request *SubmitVideoTranslateJobRequest) (response *SubmitVideoTranslateJobResponse, err error) {
    if request == nil {
        request = NewSubmitVideoTranslateJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitVideoTranslateJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitVideoTranslateJobResponse()
    err = c.Send(request, response)
    return
}
