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

package v20191213

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-12-13"

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


func NewBeautifyPicRequest() (request *BeautifyPicRequest) {
    request = &BeautifyPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "BeautifyPic")
    
    
    return
}

func NewBeautifyPicResponse() (response *BeautifyPicResponse) {
    response = &BeautifyPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BeautifyPic
// 用户上传一张人脸图片，精准定位五官，实现美肤、亮肤、祛痘等美颜功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BEAUTIFYFAILED = "FailedOperation.BeautifyFailed"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOFACEINPHOTO = "InvalidParameter.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_EYEENLARGINGILLEGAL = "InvalidParameterValue.EyeEnlargingIllegal"
//  INVALIDPARAMETERVALUE_FACELIFTINGILLEGAL = "InvalidParameterValue.FaceLiftingIllegal"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_SMOOTHINGILLEGAL = "InvalidParameterValue.SmoothingIllegal"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_WHITENINGILLEGAL = "InvalidParameterValue.WhiteningIllegal"
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
func (c *Client) BeautifyPic(request *BeautifyPicRequest) (response *BeautifyPicResponse, err error) {
    return c.BeautifyPicWithContext(context.Background(), request)
}

// BeautifyPic
// 用户上传一张人脸图片，精准定位五官，实现美肤、亮肤、祛痘等美颜功能。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BEAUTIFYFAILED = "FailedOperation.BeautifyFailed"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETER_NOFACEINPHOTO = "InvalidParameter.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_EYEENLARGINGILLEGAL = "InvalidParameterValue.EyeEnlargingIllegal"
//  INVALIDPARAMETERVALUE_FACELIFTINGILLEGAL = "InvalidParameterValue.FaceLiftingIllegal"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGEINVALID = "InvalidParameterValue.ImageInvalid"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_SMOOTHINGILLEGAL = "InvalidParameterValue.SmoothingIllegal"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  INVALIDPARAMETERVALUE_WHITENINGILLEGAL = "InvalidParameterValue.WhiteningIllegal"
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
func (c *Client) BeautifyPicWithContext(ctx context.Context, request *BeautifyPicRequest) (response *BeautifyPicResponse, err error) {
    if request == nil {
        request = NewBeautifyPicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BeautifyPic require credential")
    }

    request.SetContext(ctx)
    
    response = NewBeautifyPicResponse()
    err = c.Send(request, response)
    return
}

func NewBeautifyVideoRequest() (request *BeautifyVideoRequest) {
    request = &BeautifyVideoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "BeautifyVideo")
    
    
    return
}

func NewBeautifyVideoResponse() (response *BeautifyVideoResponse) {
    response = &BeautifyVideoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// BeautifyVideo
// 视频美颜
//
// 可能返回的错误码:
//  FAILEDOPERATION_EFFECTINNERERROR = "FailedOperation.EffectInnerError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) BeautifyVideo(request *BeautifyVideoRequest) (response *BeautifyVideoResponse, err error) {
    return c.BeautifyVideoWithContext(context.Background(), request)
}

// BeautifyVideo
// 视频美颜
//
// 可能返回的错误码:
//  FAILEDOPERATION_EFFECTINNERERROR = "FailedOperation.EffectInnerError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_JOBCONFLICT = "FailedOperation.JobConflict"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) BeautifyVideoWithContext(ctx context.Context, request *BeautifyVideoRequest) (response *BeautifyVideoResponse, err error) {
    if request == nil {
        request = NewBeautifyVideoRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("BeautifyVideo require credential")
    }

    request.SetContext(ctx)
    
    response = NewBeautifyVideoResponse()
    err = c.Send(request, response)
    return
}

func NewCancelBeautifyVideoJobRequest() (request *CancelBeautifyVideoJobRequest) {
    request = &CancelBeautifyVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "CancelBeautifyVideoJob")
    
    
    return
}

func NewCancelBeautifyVideoJobResponse() (response *CancelBeautifyVideoJobResponse) {
    response = &CancelBeautifyVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CancelBeautifyVideoJob
// 撤销视频美颜任务请求
//
// 可能返回的错误码:
//  FAILEDOPERATION_CANCELJOBFAILURE = "FailedOperation.CancelJobFailure"
//  FAILEDOPERATION_EFFECTINNERERROR = "FailedOperation.EffectInnerError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CancelBeautifyVideoJob(request *CancelBeautifyVideoJobRequest) (response *CancelBeautifyVideoJobResponse, err error) {
    return c.CancelBeautifyVideoJobWithContext(context.Background(), request)
}

// CancelBeautifyVideoJob
// 撤销视频美颜任务请求
//
// 可能返回的错误码:
//  FAILEDOPERATION_CANCELJOBFAILURE = "FailedOperation.CancelJobFailure"
//  FAILEDOPERATION_EFFECTINNERERROR = "FailedOperation.EffectInnerError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CancelBeautifyVideoJobWithContext(ctx context.Context, request *CancelBeautifyVideoJobRequest) (response *CancelBeautifyVideoJobResponse, err error) {
    if request == nil {
        request = NewCancelBeautifyVideoJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CancelBeautifyVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewCancelBeautifyVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewCreateModelRequest() (request *CreateModelRequest) {
    request = &CreateModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "CreateModel")
    
    
    return
}

func NewCreateModelResponse() (response *CreateModelResponse) {
    response = &CreateModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateModel
// 在使用LUT素材的modelid实现试唇色前，您需要先上传 LUT 格式的cube文件注册唇色ID。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。
//
// 
//
// 注：您也可以直接使用 [试唇色接口](https://cloud.tencent.com/document/product/1172/40706)，通过输入RGBA模型数值的方式指定唇色，更简单易用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUPLOADFAILED = "FailedOperation.ImageUploadFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODELVALUEEXCEED = "FailedOperation.ModelValueExceed"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  INVALIDPARAMETERVALUE_LUTIMAGEINVALID = "InvalidParameterValue.LutImageInvalid"
//  INVALIDPARAMETERVALUE_LUTIMAGESIZEINVALID = "InvalidParameterValue.LutImageSizeInvalid"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) CreateModel(request *CreateModelRequest) (response *CreateModelResponse, err error) {
    return c.CreateModelWithContext(context.Background(), request)
}

// CreateModel
// 在使用LUT素材的modelid实现试唇色前，您需要先上传 LUT 格式的cube文件注册唇色ID。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。
//
// 
//
// 注：您也可以直接使用 [试唇色接口](https://cloud.tencent.com/document/product/1172/40706)，通过输入RGBA模型数值的方式指定唇色，更简单易用。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUPLOADFAILED = "FailedOperation.ImageUploadFailed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODELVALUEEXCEED = "FailedOperation.ModelValueExceed"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  INVALIDPARAMETERVALUE_LUTIMAGEINVALID = "InvalidParameterValue.LutImageInvalid"
//  INVALIDPARAMETERVALUE_LUTIMAGESIZEINVALID = "InvalidParameterValue.LutImageSizeInvalid"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) CreateModelWithContext(ctx context.Context, request *CreateModelRequest) (response *CreateModelResponse, err error) {
    if request == nil {
        request = NewCreateModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateModelResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteModelRequest() (request *DeleteModelRequest) {
    request = &DeleteModelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "DeleteModel")
    
    
    return
}

func NewDeleteModelResponse() (response *DeleteModelResponse) {
    response = &DeleteModelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteModel
// 删除已注册的唇色素材。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  INVALIDPARAMETERVALUE_MODELIDNOTFOUND = "InvalidParameterValue.ModelIdNotFound"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) DeleteModel(request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
    return c.DeleteModelWithContext(context.Background(), request)
}

// DeleteModel
// 删除已注册的唇色素材。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  INVALIDPARAMETERVALUE_MODELIDNOTFOUND = "InvalidParameterValue.ModelIdNotFound"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) DeleteModelWithContext(ctx context.Context, request *DeleteModelRequest) (response *DeleteModelResponse, err error) {
    if request == nil {
        request = NewDeleteModelRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteModel require credential")
    }

    request.SetContext(ctx)
    
    response = NewDeleteModelResponse()
    err = c.Send(request, response)
    return
}

func NewGetModelListRequest() (request *GetModelListRequest) {
    request = &GetModelListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "GetModelList")
    
    
    return
}

func NewGetModelListResponse() (response *GetModelListResponse) {
    response = &GetModelListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// GetModelList
// 查询已注册的唇色素材。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) GetModelList(request *GetModelListRequest) (response *GetModelListResponse, err error) {
    return c.GetModelListWithContext(context.Background(), request)
}

// GetModelList
// 查询已注册的唇色素材。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
func (c *Client) GetModelListWithContext(ctx context.Context, request *GetModelListRequest) (response *GetModelListResponse, err error) {
    if request == nil {
        request = NewGetModelListRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GetModelList require credential")
    }

    request.SetContext(ctx)
    
    response = NewGetModelListResponse()
    err = c.Send(request, response)
    return
}

func NewQueryBeautifyVideoJobRequest() (request *QueryBeautifyVideoJobRequest) {
    request = &QueryBeautifyVideoJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "QueryBeautifyVideoJob")
    
    
    return
}

func NewQueryBeautifyVideoJobResponse() (response *QueryBeautifyVideoJobResponse) {
    response = &QueryBeautifyVideoJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// QueryBeautifyVideoJob
// 查询视频美颜处理进度
//
// 可能返回的错误码:
//  FAILEDOPERATION_EFFECTFREQCTRL = "FailedOperation.EffectFreqCtrl"
//  FAILEDOPERATION_EFFECTINNERERROR = "FailedOperation.EffectInnerError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"
//  FAILEDOPERATION_JOBSTOPPROCESSING = "FailedOperation.JobStopProcessing"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) QueryBeautifyVideoJob(request *QueryBeautifyVideoJobRequest) (response *QueryBeautifyVideoJobResponse, err error) {
    return c.QueryBeautifyVideoJobWithContext(context.Background(), request)
}

// QueryBeautifyVideoJob
// 查询视频美颜处理进度
//
// 可能返回的错误码:
//  FAILEDOPERATION_EFFECTFREQCTRL = "FailedOperation.EffectFreqCtrl"
//  FAILEDOPERATION_EFFECTINNERERROR = "FailedOperation.EffectInnerError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INTERNALERROR = "FailedOperation.InternalError"
//  FAILEDOPERATION_JOBHASBEENCANCELED = "FailedOperation.JobHasBeenCanceled"
//  FAILEDOPERATION_JOBSTOPPROCESSING = "FailedOperation.JobStopProcessing"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) QueryBeautifyVideoJobWithContext(ctx context.Context, request *QueryBeautifyVideoJobRequest) (response *QueryBeautifyVideoJobResponse, err error) {
    if request == nil {
        request = NewQueryBeautifyVideoJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryBeautifyVideoJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryBeautifyVideoJobResponse()
    err = c.Send(request, response)
    return
}

func NewStyleImageRequest() (request *StyleImageRequest) {
    request = &StyleImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "StyleImage")
    
    
    return
}

func NewStyleImageResponse() (response *StyleImageResponse) {
    response = &StyleImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StyleImage
// 上传一张照片，输出滤镜处理后的图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) StyleImage(request *StyleImageRequest) (response *StyleImageResponse, err error) {
    return c.StyleImageWithContext(context.Background(), request)
}

// StyleImage
// 上传一张照片，输出滤镜处理后的图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) StyleImageWithContext(ctx context.Context, request *StyleImageRequest) (response *StyleImageResponse, err error) {
    if request == nil {
        request = NewStyleImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StyleImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewStyleImageResponse()
    err = c.Send(request, response)
    return
}

func NewStyleImageProRequest() (request *StyleImageProRequest) {
    request = &StyleImageProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "StyleImagePro")
    
    
    return
}

func NewStyleImageProResponse() (response *StyleImageProResponse) {
    response = &StyleImageProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// StyleImagePro
// 上传一张照片，输出滤镜处理后的图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) StyleImagePro(request *StyleImageProRequest) (response *StyleImageProResponse, err error) {
    return c.StyleImageProWithContext(context.Background(), request)
}

// StyleImagePro
// 上传一张照片，输出滤镜处理后的图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) StyleImageProWithContext(ctx context.Context, request *StyleImageProRequest) (response *StyleImageProResponse, err error) {
    if request == nil {
        request = NewStyleImageProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("StyleImagePro require credential")
    }

    request.SetContext(ctx)
    
    response = NewStyleImageProResponse()
    err = c.Send(request, response)
    return
}

func NewTryLipstickPicRequest() (request *TryLipstickPicRequest) {
    request = &TryLipstickPicRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("fmu", APIVersion, "TryLipstickPic")
    
    
    return
}

func NewTryLipstickPicResponse() (response *TryLipstickPicResponse) {
    response = &TryLipstickPicResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TryLipstickPic
// 对图片中的人脸嘴唇进行着色，最多支持同时对一张图中的3张人脸进行试唇色。
//
// 
//
// 您可以通过事先注册在腾讯云的唇色素材（LUT文件）改变图片中的人脸唇色，也可以输入RGBA模型数值。
//
// 
//
// 为了更好的效果，建议您使用事先注册在腾讯云的唇色素材（LUT文件）。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEGRAYNOTSUPPORT = "FailedOperation.ImageGrayNotSupport"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODELVALUEEXCEED = "FailedOperation.ModelValueExceed"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_FACERECTINVALID = "InvalidParameterValue.FaceRectInvalid"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_LUTIMAGEINVALID = "InvalidParameterValue.LutImageInvalid"
//  INVALIDPARAMETERVALUE_LUTIMAGESIZEINVALID = "InvalidParameterValue.LutImageSizeInvalid"
//  INVALIDPARAMETERVALUE_MODELIDNOTFOUND = "InvalidParameterValue.ModelIdNotFound"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TryLipstickPic(request *TryLipstickPicRequest) (response *TryLipstickPicResponse, err error) {
    return c.TryLipstickPicWithContext(context.Background(), request)
}

// TryLipstickPic
// 对图片中的人脸嘴唇进行着色，最多支持同时对一张图中的3张人脸进行试唇色。
//
// 
//
// 您可以通过事先注册在腾讯云的唇色素材（LUT文件）改变图片中的人脸唇色，也可以输入RGBA模型数值。
//
// 
//
// 为了更好的效果，建议您使用事先注册在腾讯云的唇色素材（LUT文件）。
//
// 
//
// >     
//
// - 公共参数中的签名方式请使用V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DETECTNOFACE = "FailedOperation.DetectNoFace"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FREQCTRL = "FailedOperation.FreqCtrl"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEGRAYNOTSUPPORT = "FailedOperation.ImageGrayNotSupport"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODELVALUEEXCEED = "FailedOperation.ModelValueExceed"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_FACERECTINVALID = "InvalidParameterValue.FaceRectInvalid"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDFIRST = "InvalidParameterValue.FaceRectInvalidFirst"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDSECOND = "InvalidParameterValue.FaceRectInvalidSecond"
//  INVALIDPARAMETERVALUE_FACERECTINVALIDTHRID = "InvalidParameterValue.FaceRectInvalidThrid"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEED = "InvalidParameterValue.ImageSizeExceed"
//  INVALIDPARAMETERVALUE_LUTIMAGEINVALID = "InvalidParameterValue.LutImageInvalid"
//  INVALIDPARAMETERVALUE_LUTIMAGESIZEINVALID = "InvalidParameterValue.LutImageSizeInvalid"
//  INVALIDPARAMETERVALUE_MODELIDNOTFOUND = "InvalidParameterValue.ModelIdNotFound"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_GETAUTHINFOERROR = "ResourceUnavailable.GetAuthInfoError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_NOTREADY = "ResourceUnavailable.NotReady"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCEUNAVAILABLE_UNKNOWNSTATUS = "ResourceUnavailable.UnknownStatus"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TryLipstickPicWithContext(ctx context.Context, request *TryLipstickPicRequest) (response *TryLipstickPicResponse, err error) {
    if request == nil {
        request = NewTryLipstickPicRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TryLipstickPic require credential")
    }

    request.SetContext(ctx)
    
    response = NewTryLipstickPicResponse()
    err = c.Send(request, response)
    return
}
