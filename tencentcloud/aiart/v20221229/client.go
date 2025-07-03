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

package v20221229

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-12-29"

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


func NewChangeClothesRequest() (request *ChangeClothesRequest) {
    request = &ChangeClothesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "ChangeClothes")
    
    
    return
}

func NewChangeClothesResponse() (response *ChangeClothesResponse) {
    response = &ChangeClothesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ChangeClothes
// 上传正面全身模特照和服装平铺图，生成模特换装后的图片。
//
// 生成的换装图片分辨率和模特照分辨率一致。
//
// 模特换装默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ChangeClothes(request *ChangeClothesRequest) (response *ChangeClothesResponse, err error) {
    return c.ChangeClothesWithContext(context.Background(), request)
}

// ChangeClothes
// 上传正面全身模特照和服装平铺图，生成模特换装后的图片。
//
// 生成的换装图片分辨率和模特照分辨率一致。
//
// 模特换装默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ChangeClothesWithContext(ctx context.Context, request *ChangeClothesRequest) (response *ChangeClothesResponse, err error) {
    if request == nil {
        request = NewChangeClothesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ChangeClothes require credential")
    }

    request.SetContext(ctx)
    
    response = NewChangeClothesResponse()
    err = c.Send(request, response)
    return
}

func NewGenerateAvatarRequest() (request *GenerateAvatarRequest) {
    request = &GenerateAvatarRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "GenerateAvatar")
    
    
    return
}

func NewGenerateAvatarResponse() (response *GenerateAvatarResponse) {
    response = &GenerateAvatarResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// GenerateAvatar
// 百变头像接口将根据输入的人像照片，生成风格百变的头像。
//
// 百变头像默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GenerateAvatar(request *GenerateAvatarRequest) (response *GenerateAvatarResponse, err error) {
    return c.GenerateAvatarWithContext(context.Background(), request)
}

// GenerateAvatar
// 百变头像接口将根据输入的人像照片，生成风格百变的头像。
//
// 百变头像默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) GenerateAvatarWithContext(ctx context.Context, request *GenerateAvatarRequest) (response *GenerateAvatarResponse, err error) {
    if request == nil {
        request = NewGenerateAvatarRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("GenerateAvatar require credential")
    }

    request.SetContext(ctx)
    
    response = NewGenerateAvatarResponse()
    err = c.Send(request, response)
    return
}

func NewImageInpaintingRemovalRequest() (request *ImageInpaintingRemovalRequest) {
    request = &ImageInpaintingRemovalRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "ImageInpaintingRemoval")
    
    
    return
}

func NewImageInpaintingRemovalResponse() (response *ImageInpaintingRemovalResponse) {
    response = &ImageInpaintingRemovalResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageInpaintingRemoval
// 局部消除接口通过图像 mask 指定需要消除的人、物、文字等区域，在选定区域对图像内容进行消除与重绘补全。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageInpaintingRemoval(request *ImageInpaintingRemovalRequest) (response *ImageInpaintingRemovalResponse, err error) {
    return c.ImageInpaintingRemovalWithContext(context.Background(), request)
}

// ImageInpaintingRemoval
// 局部消除接口通过图像 mask 指定需要消除的人、物、文字等区域，在选定区域对图像内容进行消除与重绘补全。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageInpaintingRemovalWithContext(ctx context.Context, request *ImageInpaintingRemovalRequest) (response *ImageInpaintingRemovalResponse, err error) {
    if request == nil {
        request = NewImageInpaintingRemovalRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageInpaintingRemoval require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageInpaintingRemovalResponse()
    err = c.Send(request, response)
    return
}

func NewImageOutpaintingRequest() (request *ImageOutpaintingRequest) {
    request = &ImageOutpaintingRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "ImageOutpainting")
    
    
    return
}

func NewImageOutpaintingResponse() (response *ImageOutpaintingResponse) {
    response = &ImageOutpaintingResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageOutpainting
// 扩图接口支持对输入图像按指定宽高比实现智能扩图。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageOutpainting(request *ImageOutpaintingRequest) (response *ImageOutpaintingResponse, err error) {
    return c.ImageOutpaintingWithContext(context.Background(), request)
}

// ImageOutpainting
// 扩图接口支持对输入图像按指定宽高比实现智能扩图。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageOutpaintingWithContext(ctx context.Context, request *ImageOutpaintingRequest) (response *ImageOutpaintingResponse, err error) {
    if request == nil {
        request = NewImageOutpaintingRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageOutpainting require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageOutpaintingResponse()
    err = c.Send(request, response)
    return
}

func NewImageToImageRequest() (request *ImageToImageRequest) {
    request = &ImageToImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "ImageToImage")
    
    
    return
}

func NewImageToImageResponse() (response *ImageToImageResponse) {
    response = &ImageToImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageToImage
// 图像风格化（图生图）接口提供生成式的图生图风格转化能力，将根据输入的图像及文本描述，智能生成风格转化后的图像。建议避免输入人像过小、姿势复杂、人数较多的人像图片。
//
// 图像风格化（图生图）默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageToImage(request *ImageToImageRequest) (response *ImageToImageResponse, err error) {
    return c.ImageToImageWithContext(context.Background(), request)
}

// ImageToImage
// 图像风格化（图生图）接口提供生成式的图生图风格转化能力，将根据输入的图像及文本描述，智能生成风格转化后的图像。建议避免输入人像过小、姿势复杂、人数较多的人像图片。
//
// 图像风格化（图生图）默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ImageToImageWithContext(ctx context.Context, request *ImageToImageRequest) (response *ImageToImageResponse, err error) {
    if request == nil {
        request = NewImageToImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageToImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageToImageResponse()
    err = c.Send(request, response)
    return
}

func NewQueryDrawPortraitJobRequest() (request *QueryDrawPortraitJobRequest) {
    request = &QueryDrawPortraitJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "QueryDrawPortraitJob")
    
    
    return
}

func NewQueryDrawPortraitJobResponse() (response *QueryDrawPortraitJobResponse) {
    response = &QueryDrawPortraitJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryDrawPortraitJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 生成图片分为提交任务和查询任务2个接口：
//
// 
//
// - 提交生成写真图片任务：选择风格模板，提交一个生成写真图片异步任务，根据写真模型 ID 生成写真图片，获得任务 ID。
//
// - 查询生成写真图片任务：根据任务 ID 查询生成图片任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
func (c *Client) QueryDrawPortraitJob(request *QueryDrawPortraitJobRequest) (response *QueryDrawPortraitJobResponse, err error) {
    return c.QueryDrawPortraitJobWithContext(context.Background(), request)
}

// QueryDrawPortraitJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 生成图片分为提交任务和查询任务2个接口：
//
// 
//
// - 提交生成写真图片任务：选择风格模板，提交一个生成写真图片异步任务，根据写真模型 ID 生成写真图片，获得任务 ID。
//
// - 查询生成写真图片任务：根据任务 ID 查询生成图片任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
func (c *Client) QueryDrawPortraitJobWithContext(ctx context.Context, request *QueryDrawPortraitJobRequest) (response *QueryDrawPortraitJobResponse, err error) {
    if request == nil {
        request = NewQueryDrawPortraitJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryDrawPortraitJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryDrawPortraitJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryGlamPicJobRequest() (request *QueryGlamPicJobRequest) {
    request = &QueryGlamPicJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "QueryGlamPicJob")
    
    
    return
}

func NewQueryGlamPicJobResponse() (response *QueryGlamPicJobResponse) {
    response = &QueryGlamPicJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryGlamPicJob
// AI 美照接口将根据模板为用户生成精美照片。分为提交任务和查询任务2个接口。
//
// - 提交任务：提交一个美照生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// AI 美照默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) QueryGlamPicJob(request *QueryGlamPicJobRequest) (response *QueryGlamPicJobResponse, err error) {
    return c.QueryGlamPicJobWithContext(context.Background(), request)
}

// QueryGlamPicJob
// AI 美照接口将根据模板为用户生成精美照片。分为提交任务和查询任务2个接口。
//
// - 提交任务：提交一个美照生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// AI 美照默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) QueryGlamPicJobWithContext(ctx context.Context, request *QueryGlamPicJobRequest) (response *QueryGlamPicJobResponse, err error) {
    if request == nil {
        request = NewQueryGlamPicJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryGlamPicJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryGlamPicJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryMemeJobRequest() (request *QueryMemeJobRequest) {
    request = &QueryMemeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "QueryMemeJob")
    
    
    return
}

func NewQueryMemeJobResponse() (response *QueryMemeJobResponse) {
    response = &QueryMemeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryMemeJob
// 表情动图生成接口将静态照片制作成动态的表情包。分为提交任务和查询任务2个接口。
//
// - 提交任务：提交一个表情动图生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// 表情动图生成默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) QueryMemeJob(request *QueryMemeJobRequest) (response *QueryMemeJobResponse, err error) {
    return c.QueryMemeJobWithContext(context.Background(), request)
}

// QueryMemeJob
// 表情动图生成接口将静态照片制作成动态的表情包。分为提交任务和查询任务2个接口。
//
// - 提交任务：提交一个表情动图生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// 表情动图生成默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) QueryMemeJobWithContext(ctx context.Context, request *QueryMemeJobRequest) (response *QueryMemeJobResponse, err error) {
    if request == nil {
        request = NewQueryMemeJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryMemeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryMemeJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTextToImageProJobRequest() (request *QueryTextToImageProJobRequest) {
    request = &QueryTextToImageProJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "QueryTextToImageProJob")
    
    
    return
}

func NewQueryTextToImageProJobResponse() (response *QueryTextToImageProJobResponse) {
    response = &QueryTextToImageProJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryTextToImageProJob
// 本接口已迁移至腾讯混元大模型-混元生图，即将停止此处维护，可切换至 [混元生图 API](https://cloud.tencent.com/document/product/1729/105970) 继续使用。
//
// 文生图（高级版）接口基于高级版文生图大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个文生图（高级版）异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。文生图（高级版）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) QueryTextToImageProJob(request *QueryTextToImageProJobRequest) (response *QueryTextToImageProJobResponse, err error) {
    return c.QueryTextToImageProJobWithContext(context.Background(), request)
}

// QueryTextToImageProJob
// 本接口已迁移至腾讯混元大模型-混元生图，即将停止此处维护，可切换至 [混元生图 API](https://cloud.tencent.com/document/product/1729/105970) 继续使用。
//
// 文生图（高级版）接口基于高级版文生图大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个文生图（高级版）异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。文生图（高级版）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  FAILEDOPERATION_JOBNOTEXIST = "FailedOperation.JobNotExist"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
func (c *Client) QueryTextToImageProJobWithContext(ctx context.Context, request *QueryTextToImageProJobRequest) (response *QueryTextToImageProJobResponse, err error) {
    if request == nil {
        request = NewQueryTextToImageProJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryTextToImageProJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryTextToImageProJobResponse()
    err = c.Send(request, response)
    return
}

func NewQueryTrainPortraitModelJobRequest() (request *QueryTrainPortraitModelJobRequest) {
    request = &QueryTrainPortraitModelJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "QueryTrainPortraitModelJob")
    
    
    return
}

func NewQueryTrainPortraitModelJobResponse() (response *QueryTrainPortraitModelJobResponse) {
    response = &QueryTrainPortraitModelJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// QueryTrainPortraitModelJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 如果选择免训练模式无需调用本接口。
//
// 训练模型分为提交任务和查询任务2个接口：
//
// 
//
// - 提交训练写真模型任务：完成上传图片后，提交一个训练写真模型异步任务，根据写真模型 ID 开始训练模型。
//
// - 查询训练写真模型任务：根据写真模型 ID 查询训练任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) QueryTrainPortraitModelJob(request *QueryTrainPortraitModelJobRequest) (response *QueryTrainPortraitModelJobResponse, err error) {
    return c.QueryTrainPortraitModelJobWithContext(context.Background(), request)
}

// QueryTrainPortraitModelJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 如果选择免训练模式无需调用本接口。
//
// 训练模型分为提交任务和查询任务2个接口：
//
// 
//
// - 提交训练写真模型任务：完成上传图片后，提交一个训练写真模型异步任务，根据写真模型 ID 开始训练模型。
//
// - 查询训练写真模型任务：根据写真模型 ID 查询训练任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) QueryTrainPortraitModelJobWithContext(ctx context.Context, request *QueryTrainPortraitModelJobRequest) (response *QueryTrainPortraitModelJobResponse, err error) {
    if request == nil {
        request = NewQueryTrainPortraitModelJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("QueryTrainPortraitModelJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewQueryTrainPortraitModelJobResponse()
    err = c.Send(request, response)
    return
}

func NewRefineImageRequest() (request *RefineImageRequest) {
    request = &RefineImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "RefineImage")
    
    
    return
}

func NewRefineImageResponse() (response *RefineImageResponse) {
    response = &RefineImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// RefineImage
// 将图像变清晰，增强图像细节。变清晰后的图片将保持原图比例，长边为2048。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) RefineImage(request *RefineImageRequest) (response *RefineImageResponse, err error) {
    return c.RefineImageWithContext(context.Background(), request)
}

// RefineImage
// 将图像变清晰，增强图像细节。变清晰后的图片将保持原图比例，长边为2048。
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
func (c *Client) RefineImageWithContext(ctx context.Context, request *RefineImageRequest) (response *RefineImageResponse, err error) {
    if request == nil {
        request = NewRefineImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RefineImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewRefineImageResponse()
    err = c.Send(request, response)
    return
}

func NewReplaceBackgroundRequest() (request *ReplaceBackgroundRequest) {
    request = &ReplaceBackgroundRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "ReplaceBackground")
    
    
    return
}

func NewReplaceBackgroundResponse() (response *ReplaceBackgroundResponse) {
    response = &ReplaceBackgroundResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ReplaceBackground
// 商品背景生成接口根据指定的背景描述 Prompt，将商品图中的原背景替换为自定义的新背景并保留商品主体形象，实现商品背景的自由生成与更换。
//
// 
//
// 商品背景生成默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReplaceBackground(request *ReplaceBackgroundRequest) (response *ReplaceBackgroundResponse, err error) {
    return c.ReplaceBackgroundWithContext(context.Background(), request)
}

// ReplaceBackground
// 商品背景生成接口根据指定的背景描述 Prompt，将商品图中的原背景替换为自定义的新背景并保留商品主体形象，实现商品背景的自由生成与更换。
//
// 
//
// 商品背景生成默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) ReplaceBackgroundWithContext(ctx context.Context, request *ReplaceBackgroundRequest) (response *ReplaceBackgroundResponse, err error) {
    if request == nil {
        request = NewReplaceBackgroundRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ReplaceBackground require credential")
    }

    request.SetContext(ctx)
    
    response = NewReplaceBackgroundResponse()
    err = c.Send(request, response)
    return
}

func NewSketchToImageRequest() (request *SketchToImageRequest) {
    request = &SketchToImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "SketchToImage")
    
    
    return
}

func NewSketchToImageResponse() (response *SketchToImageResponse) {
    response = &SketchToImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SketchToImage
// 线稿生图接口支持上传一张黑白线稿图，按照指定的主体对象以及样式、颜色、材质、风格等的文本描述prompt ，对线稿图进行色彩填充与细节描绘，得到一张完整绘制的图像。生成图分辨率默认为1024:1024。
//
// 线稿生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SketchToImage(request *SketchToImageRequest) (response *SketchToImageResponse, err error) {
    return c.SketchToImageWithContext(context.Background(), request)
}

// SketchToImage
// 线稿生图接口支持上传一张黑白线稿图，按照指定的主体对象以及样式、颜色、材质、风格等的文本描述prompt ，对线稿图进行色彩填充与细节描绘，得到一张完整绘制的图像。生成图分辨率默认为1024:1024。
//
// 线稿生图默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_STYLECONFLICT = "InvalidParameterValue.StyleConflict"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_DELIVERING = "ResourceUnavailable.Delivering"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SketchToImageWithContext(ctx context.Context, request *SketchToImageRequest) (response *SketchToImageResponse, err error) {
    if request == nil {
        request = NewSketchToImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SketchToImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSketchToImageResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitDrawPortraitJobRequest() (request *SubmitDrawPortraitJobRequest) {
    request = &SubmitDrawPortraitJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "SubmitDrawPortraitJob")
    
    
    return
}

func NewSubmitDrawPortraitJobResponse() (response *SubmitDrawPortraitJobResponse) {
    response = &SubmitDrawPortraitJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitDrawPortraitJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 生成图片分为提交任务和查询任务2个接口：
//
// 
//
// - 提交生成写真图片任务：选择风格模板，提交一个生成写真图片异步任务，根据写真模型 ID 生成写真图片，获得任务 ID。
//
// - 查询生成写真图片任务：根据任务 ID 查询生成图片任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 提交生成写真图片任务默认提供1个并发。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitDrawPortraitJob(request *SubmitDrawPortraitJobRequest) (response *SubmitDrawPortraitJobResponse, err error) {
    return c.SubmitDrawPortraitJobWithContext(context.Background(), request)
}

// SubmitDrawPortraitJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 生成图片分为提交任务和查询任务2个接口：
//
// 
//
// - 提交生成写真图片任务：选择风格模板，提交一个生成写真图片异步任务，根据写真模型 ID 生成写真图片，获得任务 ID。
//
// - 查询生成写真图片任务：根据任务 ID 查询生成图片任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 提交生成写真图片任务默认提供1个并发。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitDrawPortraitJobWithContext(ctx context.Context, request *SubmitDrawPortraitJobRequest) (response *SubmitDrawPortraitJobResponse, err error) {
    if request == nil {
        request = NewSubmitDrawPortraitJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitDrawPortraitJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitDrawPortraitJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitGlamPicJobRequest() (request *SubmitGlamPicJobRequest) {
    request = &SubmitGlamPicJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "SubmitGlamPicJob")
    
    
    return
}

func NewSubmitGlamPicJobResponse() (response *SubmitGlamPicJobResponse) {
    response = &SubmitGlamPicJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitGlamPicJob
// AI 美照接口将根据模板为用户生成精美照片。分为提交任务和查询任务2个接口。
//
// - 提交任务：提交一个美照生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// AI 美照默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitGlamPicJob(request *SubmitGlamPicJobRequest) (response *SubmitGlamPicJobResponse, err error) {
    return c.SubmitGlamPicJobWithContext(context.Background(), request)
}

// SubmitGlamPicJob
// AI 美照接口将根据模板为用户生成精美照片。分为提交任务和查询任务2个接口。
//
// - 提交任务：提交一个美照生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// AI 美照默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitGlamPicJobWithContext(ctx context.Context, request *SubmitGlamPicJobRequest) (response *SubmitGlamPicJobResponse, err error) {
    if request == nil {
        request = NewSubmitGlamPicJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitGlamPicJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitGlamPicJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitMemeJobRequest() (request *SubmitMemeJobRequest) {
    request = &SubmitMemeJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "SubmitMemeJob")
    
    
    return
}

func NewSubmitMemeJobResponse() (response *SubmitMemeJobResponse) {
    response = &SubmitMemeJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitMemeJob
// 表情动图生成接口将静态照片制作成动态的表情包。分为提交任务和查询任务2个接口。
//
// 
//
// - 提交任务：提交一个表情动图生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// 表情动图生成默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitMemeJob(request *SubmitMemeJobRequest) (response *SubmitMemeJobResponse, err error) {
    return c.SubmitMemeJobWithContext(context.Background(), request)
}

// SubmitMemeJob
// 表情动图生成接口将静态照片制作成动态的表情包。分为提交任务和查询任务2个接口。
//
// 
//
// - 提交任务：提交一个表情动图生成异步任务，获得任务 ID。
//
// - 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 
//
// 表情动图生成默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitMemeJobWithContext(ctx context.Context, request *SubmitMemeJobRequest) (response *SubmitMemeJobResponse, err error) {
    if request == nil {
        request = NewSubmitMemeJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitMemeJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitMemeJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTextToImageProJobRequest() (request *SubmitTextToImageProJobRequest) {
    request = &SubmitTextToImageProJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "SubmitTextToImageProJob")
    
    
    return
}

func NewSubmitTextToImageProJobResponse() (response *SubmitTextToImageProJobResponse) {
    response = &SubmitTextToImageProJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTextToImageProJob
// 本接口已迁移至腾讯混元大模型-混元生图，即将停止此处维护，可切换至 [混元生图 API](https://cloud.tencent.com/document/product/1729/105969) 继续使用。
//
// 文生图（高级版）接口基于高级版文生图大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个文生图（高级版）异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。文生图（高级版）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LOGOPARAMERR = "InvalidParameterValue.LogoParamErr"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTextToImageProJob(request *SubmitTextToImageProJobRequest) (response *SubmitTextToImageProJobResponse, err error) {
    return c.SubmitTextToImageProJobWithContext(context.Background(), request)
}

// SubmitTextToImageProJob
// 本接口已迁移至腾讯混元大模型-混元生图，即将停止此处维护，可切换至 [混元生图 API](https://cloud.tencent.com/document/product/1729/105969) 继续使用。
//
// 文生图（高级版）接口基于高级版文生图大模型，将根据输入的文本描述，智能生成与之相关的结果图。分为提交任务和查询任务2个接口。
//
// 提交任务：输入文本等，提交一个文生图（高级版）异步任务，获得任务 ID。
//
// 查询任务：根据任务 ID 查询任务的处理状态、处理结果，任务处理完成后可获得生成图像结果。
//
// 并发任务数（并发）说明：并发任务数指能同时处理的任务数量。文生图（高级版）默认提供1个并发任务数，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_LOGOPARAMERR = "InvalidParameterValue.LogoParamErr"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) SubmitTextToImageProJobWithContext(ctx context.Context, request *SubmitTextToImageProJobRequest) (response *SubmitTextToImageProJobResponse, err error) {
    if request == nil {
        request = NewSubmitTextToImageProJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTextToImageProJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTextToImageProJobResponse()
    err = c.Send(request, response)
    return
}

func NewSubmitTrainPortraitModelJobRequest() (request *SubmitTrainPortraitModelJobRequest) {
    request = &SubmitTrainPortraitModelJobRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "SubmitTrainPortraitModelJob")
    
    
    return
}

func NewSubmitTrainPortraitModelJobResponse() (response *SubmitTrainPortraitModelJobResponse) {
    response = &SubmitTrainPortraitModelJobResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// SubmitTrainPortraitModelJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 如果选择免训练模式无需调用本接口。
//
// 训练模型分为提交任务和查询任务2个接口：
//
// - 提交训练写真模型任务：完成上传图片后，提交一个训练写真模型异步任务，根据写真模型 ID 开始训练模型。
//
// - 查询训练写真模型任务：根据写真模型 ID 查询训练任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 提交训练写真模型任务按并发计费，无默认并发额度。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitTrainPortraitModelJob(request *SubmitTrainPortraitModelJobRequest) (response *SubmitTrainPortraitModelJobResponse, err error) {
    return c.SubmitTrainPortraitModelJobWithContext(context.Background(), request)
}

// SubmitTrainPortraitModelJob
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 如果选择免训练模式无需调用本接口。
//
// 训练模型分为提交任务和查询任务2个接口：
//
// - 提交训练写真模型任务：完成上传图片后，提交一个训练写真模型异步任务，根据写真模型 ID 开始训练模型。
//
// - 查询训练写真模型任务：根据写真模型 ID 查询训练任务的处理状态、处理结果。
//
// 
//
// 每个写真模型自训练完成起1年内有效，有效期内可使用写真模型 ID 生成图片，期满后需要重新训练。
//
// 提交训练写真模型任务按并发计费，无默认并发额度。
//
// 可能返回的错误码:
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) SubmitTrainPortraitModelJobWithContext(ctx context.Context, request *SubmitTrainPortraitModelJobRequest) (response *SubmitTrainPortraitModelJobResponse, err error) {
    if request == nil {
        request = NewSubmitTrainPortraitModelJobRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SubmitTrainPortraitModelJob require credential")
    }

    request.SetContext(ctx)
    
    response = NewSubmitTrainPortraitModelJobResponse()
    err = c.Send(request, response)
    return
}

func NewTextToImageRequest() (request *TextToImageRequest) {
    request = &TextToImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "TextToImage")
    
    
    return
}

func NewTextToImageResponse() (response *TextToImageResponse) {
    response = &TextToImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToImage
// **本接口已迁移至腾讯混元大模型-文生图轻量版，即将停止此处维护，可切换至 [文生图轻量版 API](https://cloud.tencent.com/document/product/1729/108738) 继续使用。**
//
// 
//
// 智能文生图接口基于文生图（标准版）模型，将根据输入的文本描述，智能生成与之相关的结果图。
//
// 
//
// 智能文生图默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextToImage(request *TextToImageRequest) (response *TextToImageResponse, err error) {
    return c.TextToImageWithContext(context.Background(), request)
}

// TextToImage
// **本接口已迁移至腾讯混元大模型-文生图轻量版，即将停止此处维护，可切换至 [文生图轻量版 API](https://cloud.tencent.com/document/product/1729/108738) 继续使用。**
//
// 
//
// 智能文生图接口基于文生图（标准版）模型，将根据输入的文本描述，智能生成与之相关的结果图。
//
// 
//
// 智能文生图默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextToImageWithContext(ctx context.Context, request *TextToImageRequest) (response *TextToImageResponse, err error) {
    if request == nil {
        request = NewTextToImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToImageResponse()
    err = c.Send(request, response)
    return
}

func NewTextToImageLiteRequest() (request *TextToImageLiteRequest) {
    request = &TextToImageLiteRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "TextToImageLite")
    
    
    return
}

func NewTextToImageLiteResponse() (response *TextToImageLiteResponse) {
    response = &TextToImageLiteResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToImageLite
// 混元文生图接口，基于混元大模型，根据输入的文本描述智能生成图片
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextToImageLite(request *TextToImageLiteRequest) (response *TextToImageLiteResponse, err error) {
    return c.TextToImageLiteWithContext(context.Background(), request)
}

// TextToImageLite
// 混元文生图接口，基于混元大模型，根据输入的文本描述智能生成图片
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextToImageLiteWithContext(ctx context.Context, request *TextToImageLiteRequest) (response *TextToImageLiteResponse, err error) {
    if request == nil {
        request = NewTextToImageLiteRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToImageLite require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToImageLiteResponse()
    err = c.Send(request, response)
    return
}

func NewTextToImageRapidRequest() (request *TextToImageRapidRequest) {
    request = &TextToImageRapidRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "TextToImageRapid")
    
    
    return
}

func NewTextToImageRapidResponse() (response *TextToImageRapidResponse) {
    response = &TextToImageRapidResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToImageRapid
// 混元文生图接口，基于混元大模型，根据输入的文本描述智能生成图片
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextToImageRapid(request *TextToImageRapidRequest) (response *TextToImageRapidResponse, err error) {
    return c.TextToImageRapidWithContext(context.Background(), request)
}

// TextToImageRapid
// 混元文生图接口，基于混元大模型，根据输入的文本描述智能生成图片
//
// 默认提供1个并发，代表最多能同时处理1个已提交的任务，上一个任务处理完毕后，才能开始处理下一个任务。
//
// 可能返回的错误码:
//  AUTHFAILURE_UNAUTHORIZEDOPERATION = "AuthFailure.UnauthorizedOperation"
//  FAILEDOPERATION_CONSOLESERVERERROR = "FailedOperation.ConsoleServerError"
//  FAILEDOPERATION_GENERATEIMAGEFAILED = "FailedOperation.GenerateImageFailed"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  INVALIDPARAMETERVALUE_TEXTLENGTHEXCEED = "InvalidParameterValue.TextLengthExceed"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  OPERATIONDENIED_IMAGEILLEGALDETECTED = "OperationDenied.ImageIllegalDetected"
//  OPERATIONDENIED_TEXTILLEGALDETECTED = "OperationDenied.TextIllegalDetected"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_LOWBALANCE = "ResourceUnavailable.LowBalance"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) TextToImageRapidWithContext(ctx context.Context, request *TextToImageRapidRequest) (response *TextToImageRapidResponse, err error) {
    if request == nil {
        request = NewTextToImageRapidRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToImageRapid require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToImageRapidResponse()
    err = c.Send(request, response)
    return
}

func NewUploadTrainPortraitImagesRequest() (request *UploadTrainPortraitImagesRequest) {
    request = &UploadTrainPortraitImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("aiart", APIVersion, "UploadTrainPortraitImages")
    
    
    return
}

func NewUploadTrainPortraitImagesResponse() (response *UploadTrainPortraitImagesResponse) {
    response = &UploadTrainPortraitImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// UploadTrainPortraitImages
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 本接口用于上传人像图片并指定对应的写真模型 ID。上传的图片要求是同一个人，建议上传单人、正脸、脸部区域占比较大、脸部清晰无遮挡、无大角度偏转、无夸张表情的图片。
//
// 可选模式：
//
// - 常规训练模式：上传20 - 25张图片用于模型训练，完成训练后可生成写真图片。
//
// - 快速训练模式：仅需上传1张图片用于模型训练，训练速度更快，完成训练后可生成写真图片。
//
// - 免训练模式：仅需上传1张图片，跳过训练环节，直接生成写真图片。
//
// 
//
// 上传写真训练图片默认提供1个并发。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) UploadTrainPortraitImages(request *UploadTrainPortraitImagesRequest) (response *UploadTrainPortraitImagesResponse, err error) {
    return c.UploadTrainPortraitImagesWithContext(context.Background(), request)
}

// UploadTrainPortraitImages
// AI 写真分为上传训练图片、训练写真模型（可选跳过）、生成写真图片3个环节，需要依次调用对应接口。
//
// 本接口用于上传人像图片并指定对应的写真模型 ID。上传的图片要求是同一个人，建议上传单人、正脸、脸部区域占比较大、脸部清晰无遮挡、无大角度偏转、无夸张表情的图片。
//
// 可选模式：
//
// - 常规训练模式：上传20 - 25张图片用于模型训练，完成训练后可生成写真图片。
//
// - 快速训练模式：仅需上传1张图片用于模型训练，训练速度更快，完成训练后可生成写真图片。
//
// - 免训练模式：仅需上传1张图片，跳过训练环节，直接生成写真图片。
//
// 
//
// 上传写真训练图片默认提供1个并发。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED_JOBNUMEXCEED = "RequestLimitExceeded.JobNumExceed"
func (c *Client) UploadTrainPortraitImagesWithContext(ctx context.Context, request *UploadTrainPortraitImagesRequest) (response *UploadTrainPortraitImagesResponse, err error) {
    if request == nil {
        request = NewUploadTrainPortraitImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UploadTrainPortraitImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewUploadTrainPortraitImagesResponse()
    err = c.Send(request, response)
    return
}
