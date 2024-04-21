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
// 智能图生图接口提供基础版图生图能力，将根据输入的图像及文本描述，智能生成风格转化后的图像。建议避免输入人像过小、姿势复杂、人数较多的人像图片。
//
// 智能图生图默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
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
// 智能图生图接口提供基础版图生图能力，将根据输入的图像及文本描述，智能生成风格转化后的图像。建议避免输入人像过小、姿势复杂、人数较多的人像图片。
//
// 智能图生图默认提供3个并发任务数，代表最多能同时处理3个已提交的任务，上一个任务处理完毕后才能开始处理下一个任务。
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
