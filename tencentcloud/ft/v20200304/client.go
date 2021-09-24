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

package v20200304

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-04"

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


func NewCancelFaceMorphJobRequest() (request *CancelFaceMorphJobRequest) {
    request = &CancelFaceMorphJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "CancelFaceMorphJob")
    return
}

func NewCancelFaceMorphJobResponse() (response *CancelFaceMorphJobResponse) {
    response = &CancelFaceMorphJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelFaceMorphJob
// 撤销人像渐变任务请求
//
// 可能返回的错误码:
//  FAILEDOPERATION_CANCELJOBFAILURE = "FailedOperation.CancelJobFailure"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"
//  FAILEDOPERATION_JOBSTOPPROCESSING = "FailedOperation.JobStopProcessing"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
func (c *Client) CancelFaceMorphJob(request *CancelFaceMorphJobRequest) (response *CancelFaceMorphJobResponse, err error) {
    if request == nil {
        request = NewCancelFaceMorphJobRequest()
    }
    response = NewCancelFaceMorphJobResponse()
    err = c.Send(request, response)
    return
}

func NewChangeAgePicRequest() (request *ChangeAgePicRequest) {
    request = &ChangeAgePicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "ChangeAgePic")
    return
}

func NewChangeAgePicResponse() (response *ChangeAgePicResponse) {
    response = &ChangeAgePicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ChangeAgePic
// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸变老或变年轻的图片，支持实现人脸不同年龄的变化。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//  FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ChangeAgePic(request *ChangeAgePicRequest) (response *ChangeAgePicResponse, err error) {
    if request == nil {
        request = NewChangeAgePicRequest()
    }
    response = NewChangeAgePicResponse()
    err = c.Send(request, response)
    return
}

func NewFaceCartoonPicRequest() (request *FaceCartoonPicRequest) {
    request = &FaceCartoonPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "FaceCartoonPic")
    return
}

func NewFaceCartoonPicResponse() (response *FaceCartoonPicResponse) {
    response = &FaceCartoonPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// FaceCartoonPic
// 输入一张人脸照片，生成个性化的二次元动漫形象，可用于打造个性头像、趣味活动、特效类应用等场景，提升社交娱乐的体验。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  UNSUPPORTEDOPERATION_UNKNOWMETHOD = "UnsupportedOperation.UnknowMethod"
func (c *Client) FaceCartoonPic(request *FaceCartoonPicRequest) (response *FaceCartoonPicResponse, err error) {
    if request == nil {
        request = NewFaceCartoonPicRequest()
    }
    response = NewFaceCartoonPicResponse()
    err = c.Send(request, response)
    return
}

func NewMorphFaceRequest() (request *MorphFaceRequest) {
    request = &MorphFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "MorphFace")
    return
}

func NewMorphFaceResponse() (response *MorphFaceResponse) {
    response = &MorphFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// MorphFace
// 输入2-5张人脸照片，生成一段以人脸为焦点的渐变视频或GIF图，支持自定义图片播放速度、视频每秒传输帧数，可用于短视频、表情包、创意H5等应用场景，丰富静态图片的玩法。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) MorphFace(request *MorphFaceRequest) (response *MorphFaceResponse, err error) {
    if request == nil {
        request = NewMorphFaceRequest()
    }
    response = NewMorphFaceResponse()
    err = c.Send(request, response)
    return
}

func NewQueryFaceMorphJobRequest() (request *QueryFaceMorphJobRequest) {
    request = &QueryFaceMorphJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "QueryFaceMorphJob")
    return
}

func NewQueryFaceMorphJobResponse() (response *QueryFaceMorphJobResponse) {
    response = &QueryFaceMorphJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryFaceMorphJob
// 查询人像渐变处理进度
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
func (c *Client) QueryFaceMorphJob(request *QueryFaceMorphJobRequest) (response *QueryFaceMorphJobResponse, err error) {
    if request == nil {
        request = NewQueryFaceMorphJobRequest()
    }
    response = NewQueryFaceMorphJobResponse()
    err = c.Send(request, response)
    return
}

func NewSwapGenderPicRequest() (request *SwapGenderPicRequest) {
    request = &SwapGenderPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ft", APIVersion, "SwapGenderPic")
    return
}

func NewSwapGenderPicResponse() (response *SwapGenderPicResponse) {
    response = &SwapGenderPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SwapGenderPic
// 用户上传一张人脸图片，基于人脸编辑与生成算法，输出一张人脸性别转换的图片。男变女可实现美颜、淡妆、加刘海和长发的效果；女变男可实现加胡须、变短发的效果。 
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SwapGenderPic(request *SwapGenderPicRequest) (response *SwapGenderPicResponse, err error) {
    if request == nil {
        request = NewSwapGenderPicRequest()
    }
    response = NewSwapGenderPicResponse()
    err = c.Send(request, response)
    return
}
