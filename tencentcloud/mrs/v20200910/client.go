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

package v20200910

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-09-10"

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


func NewImageMaskRequest() (request *ImageMaskRequest) {
    request = &ImageMaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "ImageMask")
    
    
    return
}

func NewImageMaskResponse() (response *ImageMaskResponse) {
    response = &ImageMaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageMask
// 医疗报告图片脱敏接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) ImageMask(request *ImageMaskRequest) (response *ImageMaskResponse, err error) {
    return c.ImageMaskWithContext(context.Background(), request)
}

// ImageMask
// 医疗报告图片脱敏接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) ImageMaskWithContext(ctx context.Context, request *ImageMaskRequest) (response *ImageMaskResponse, err error) {
    if request == nil {
        request = NewImageMaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageMask require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageMaskResponse()
    err = c.Send(request, response)
    return
}

func NewImageMaskAsyncRequest() (request *ImageMaskAsyncRequest) {
    request = &ImageMaskAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "ImageMaskAsync")
    
    
    return
}

func NewImageMaskAsyncResponse() (response *ImageMaskAsyncResponse) {
    response = &ImageMaskAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageMaskAsync
// 图片脱敏-异步接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCQUEUEFULLERROR = "FailedOperation.AsyncQueueFullError"
func (c *Client) ImageMaskAsync(request *ImageMaskAsyncRequest) (response *ImageMaskAsyncResponse, err error) {
    return c.ImageMaskAsyncWithContext(context.Background(), request)
}

// ImageMaskAsync
// 图片脱敏-异步接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCQUEUEFULLERROR = "FailedOperation.AsyncQueueFullError"
func (c *Client) ImageMaskAsyncWithContext(ctx context.Context, request *ImageMaskAsyncRequest) (response *ImageMaskAsyncResponse, err error) {
    if request == nil {
        request = NewImageMaskAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageMaskAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageMaskAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewImageMaskAsyncGetResultRequest() (request *ImageMaskAsyncGetResultRequest) {
    request = &ImageMaskAsyncGetResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "ImageMaskAsyncGetResult")
    
    
    return
}

func NewImageMaskAsyncGetResultResponse() (response *ImageMaskAsyncGetResultResponse) {
    response = &ImageMaskAsyncGetResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageMaskAsyncGetResult
// 图片脱敏-异步获取结果接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCTASKERROR = "FailedOperation.AsyncTaskError"
//  FAILEDOPERATION_ASYNCTASKHANDLING = "FailedOperation.AsyncTaskHandling"
func (c *Client) ImageMaskAsyncGetResult(request *ImageMaskAsyncGetResultRequest) (response *ImageMaskAsyncGetResultResponse, err error) {
    return c.ImageMaskAsyncGetResultWithContext(context.Background(), request)
}

// ImageMaskAsyncGetResult
// 图片脱敏-异步获取结果接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCTASKERROR = "FailedOperation.AsyncTaskError"
//  FAILEDOPERATION_ASYNCTASKHANDLING = "FailedOperation.AsyncTaskHandling"
func (c *Client) ImageMaskAsyncGetResultWithContext(ctx context.Context, request *ImageMaskAsyncGetResultRequest) (response *ImageMaskAsyncGetResultResponse, err error) {
    if request == nil {
        request = NewImageMaskAsyncGetResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageMaskAsyncGetResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageMaskAsyncGetResultResponse()
    err = c.Send(request, response)
    return
}

func NewImageToClassRequest() (request *ImageToClassRequest) {
    request = &ImageToClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "ImageToClass")
    
    
    return
}

func NewImageToClassResponse() (response *ImageToClassResponse) {
    response = &ImageToClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageToClass
// 图片分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  OPERATIONDENIED_UNSUPPORTTHISTYPE = "OperationDenied.UnSupportThisType"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) ImageToClass(request *ImageToClassRequest) (response *ImageToClassResponse, err error) {
    return c.ImageToClassWithContext(context.Background(), request)
}

// ImageToClass
// 图片分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  OPERATIONDENIED_UNSUPPORTTHISTYPE = "OperationDenied.UnSupportThisType"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) ImageToClassWithContext(ctx context.Context, request *ImageToClassRequest) (response *ImageToClassResponse, err error) {
    if request == nil {
        request = NewImageToClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageToClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageToClassResponse()
    err = c.Send(request, response)
    return
}

func NewImageToObjectRequest() (request *ImageToObjectRequest) {
    request = &ImageToObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "ImageToObject")
    
    
    return
}

func NewImageToObjectResponse() (response *ImageToObjectResponse) {
    response = &ImageToObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageToObject
// 图片转结构化对象
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEISNOTEXT = "InvalidParameterValue.ImageIsNoText"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  LIMITEXCEEDED_TEXTSIZELIMITEXCEEDED = "LimitExceeded.TextSizeLimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) ImageToObject(request *ImageToObjectRequest) (response *ImageToObjectResponse, err error) {
    return c.ImageToObjectWithContext(context.Background(), request)
}

// ImageToObject
// 图片转结构化对象
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEISNOTEXT = "InvalidParameterValue.ImageIsNoText"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  LIMITEXCEEDED_TEXTSIZELIMITEXCEEDED = "LimitExceeded.TextSizeLimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) ImageToObjectWithContext(ctx context.Context, request *ImageToObjectRequest) (response *ImageToObjectResponse, err error) {
    if request == nil {
        request = NewImageToObjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageToObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageToObjectResponse()
    err = c.Send(request, response)
    return
}

func NewTextToClassRequest() (request *TextToClassRequest) {
    request = &TextToClassRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "TextToClass")
    
    
    return
}

func NewTextToClassResponse() (response *TextToClassResponse) {
    response = &TextToClassResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToClass
// 文本分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) TextToClass(request *TextToClassRequest) (response *TextToClassResponse, err error) {
    return c.TextToClassWithContext(context.Background(), request)
}

// TextToClass
// 文本分类
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) TextToClassWithContext(ctx context.Context, request *TextToClassRequest) (response *TextToClassResponse, err error) {
    if request == nil {
        request = NewTextToClassRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToClass require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToClassResponse()
    err = c.Send(request, response)
    return
}

func NewTextToObjectRequest() (request *TextToObjectRequest) {
    request = &TextToObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "TextToObject")
    
    
    return
}

func NewTextToObjectResponse() (response *TextToObjectResponse) {
    response = &TextToObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TextToObject
// 文本转结构化对象。
//
// 
//
// 适用场景：经过腾讯医疗专用 OCR 从图片识别之后的文本，可以调用此接口。通过其它 OCR 识别的文本可能不适配。医院的 XML 格式文本也不适配，XML 文件需要经过特殊转换才能直接调用此接口。单次调用传入的文本不宜超过 2000 字。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  LIMITEXCEEDED_TEXTSIZELIMITEXCEEDED = "LimitExceeded.TextSizeLimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) TextToObject(request *TextToObjectRequest) (response *TextToObjectResponse, err error) {
    return c.TextToObjectWithContext(context.Background(), request)
}

// TextToObject
// 文本转结构化对象。
//
// 
//
// 适用场景：经过腾讯医疗专用 OCR 从图片识别之后的文本，可以调用此接口。通过其它 OCR 识别的文本可能不适配。医院的 XML 格式文本也不适配，XML 文件需要经过特殊转换才能直接调用此接口。单次调用传入的文本不宜超过 2000 字。
//
// 可能返回的错误码:
//  AUTHFAILURE = "AuthFailure"
//  FAILEDOPERATION_PUSHUSAGEMESSAGEERROR = "FailedOperation.PushUsageMessageError"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_IMAGEOCRERROR = "InternalError.ImageOcrError"
//  INTERNALERROR_IMAGEPROCESSERROR = "InternalError.ImageProcessError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
//  INTERNALERROR_STRUCTIONERROR = "InternalError.StructionError"
//  INTERNALERROR_TEXTCLASSIFYERROR = "InternalError.TextClassifyError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_AUTOFITDIRECTION = "InvalidParameter.AutoFitDirection"
//  INVALIDPARAMETER_IMAGEINFOLIST = "InvalidParameter.ImageInfoList"
//  INVALIDPARAMETER_IMAGEORIGINALSIZE = "InvalidParameter.ImageOriginalSize"
//  INVALIDPARAMETER_INVALIDACTION = "InvalidParameter.InvalidAction"
//  INVALIDPARAMETER_OCRENGINETYPE = "InvalidParameter.OcrEngineType"
//  INVALIDPARAMETER_ROTATETHEANGLE = "InvalidParameter.RotateTheAngle"
//  INVALIDPARAMETER_TEXT = "InvalidParameter.Text"
//  INVALIDPARAMETERVALUE_IMAGECODEINVALID = "InvalidParameterValue.ImageCodeInvalid"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  LIMITEXCEEDED_TEXTSIZELIMITEXCEEDED = "LimitExceeded.TextSizeLimitExceeded"
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) TextToObjectWithContext(ctx context.Context, request *TextToObjectRequest) (response *TextToObjectResponse, err error) {
    if request == nil {
        request = NewTextToObjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TextToObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewTextToObjectResponse()
    err = c.Send(request, response)
    return
}

func NewTurnPDFToObjectRequest() (request *TurnPDFToObjectRequest) {
    request = &TurnPDFToObjectRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "TurnPDFToObject")
    
    
    return
}

func NewTurnPDFToObjectResponse() (response *TurnPDFToObjectResponse) {
    response = &TurnPDFToObjectResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TurnPDFToObject
// 将PDF格式的体检报告文件结构化，解析关键信息。
//
// 注意：该接口是按照体检报告 PDF 页面数量统计次数，不是按照 PDF 文件数量统计次数。通过该接口传入的报告必须是体检报告，非体检报告可能无法正确解析。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) TurnPDFToObject(request *TurnPDFToObjectRequest) (response *TurnPDFToObjectResponse, err error) {
    return c.TurnPDFToObjectWithContext(context.Background(), request)
}

// TurnPDFToObject
// 将PDF格式的体检报告文件结构化，解析关键信息。
//
// 注意：该接口是按照体检报告 PDF 页面数量统计次数，不是按照 PDF 文件数量统计次数。通过该接口传入的报告必须是体检报告，非体检报告可能无法正确解析。
//
// 可能返回的错误码:
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) TurnPDFToObjectWithContext(ctx context.Context, request *TurnPDFToObjectRequest) (response *TurnPDFToObjectResponse, err error) {
    if request == nil {
        request = NewTurnPDFToObjectRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TurnPDFToObject require credential")
    }

    request.SetContext(ctx)
    
    response = NewTurnPDFToObjectResponse()
    err = c.Send(request, response)
    return
}

func NewTurnPDFToObjectAsyncRequest() (request *TurnPDFToObjectAsyncRequest) {
    request = &TurnPDFToObjectAsyncRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "TurnPDFToObjectAsync")
    
    
    return
}

func NewTurnPDFToObjectAsyncResponse() (response *TurnPDFToObjectAsyncResponse) {
    response = &TurnPDFToObjectAsyncResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TurnPDFToObjectAsync
// 体检报告PDF文件结构化-异步接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCQUEUEFULLERROR = "FailedOperation.AsyncQueueFullError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) TurnPDFToObjectAsync(request *TurnPDFToObjectAsyncRequest) (response *TurnPDFToObjectAsyncResponse, err error) {
    return c.TurnPDFToObjectAsyncWithContext(context.Background(), request)
}

// TurnPDFToObjectAsync
// 体检报告PDF文件结构化-异步接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_ASYNCQUEUEFULLERROR = "FailedOperation.AsyncQueueFullError"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) TurnPDFToObjectAsyncWithContext(ctx context.Context, request *TurnPDFToObjectAsyncRequest) (response *TurnPDFToObjectAsyncResponse, err error) {
    if request == nil {
        request = NewTurnPDFToObjectAsyncRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TurnPDFToObjectAsync require credential")
    }

    request.SetContext(ctx)
    
    response = NewTurnPDFToObjectAsyncResponse()
    err = c.Send(request, response)
    return
}

func NewTurnPDFToObjectAsyncGetResultRequest() (request *TurnPDFToObjectAsyncGetResultRequest) {
    request = &TurnPDFToObjectAsyncGetResultRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("mrs", APIVersion, "TurnPDFToObjectAsyncGetResult")
    
    
    return
}

func NewTurnPDFToObjectAsyncGetResultResponse() (response *TurnPDFToObjectAsyncGetResultResponse) {
    response = &TurnPDFToObjectAsyncGetResultResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// TurnPDFToObjectAsyncGetResult
// 体检报告PDF文件结构化异步获取结果接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYRESULT = "FailedOperation.EmptyResult"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) TurnPDFToObjectAsyncGetResult(request *TurnPDFToObjectAsyncGetResultRequest) (response *TurnPDFToObjectAsyncGetResultResponse, err error) {
    return c.TurnPDFToObjectAsyncGetResultWithContext(context.Background(), request)
}

// TurnPDFToObjectAsyncGetResult
// 体检报告PDF文件结构化异步获取结果接口
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYRESULT = "FailedOperation.EmptyResult"
//  FAILEDOPERATION_SERVICENOTOPEN = "FailedOperation.ServiceNotOpen"
//  INTERNALERROR_SERVERTIMEOUTERROR = "InternalError.ServerTimeOutError"
func (c *Client) TurnPDFToObjectAsyncGetResultWithContext(ctx context.Context, request *TurnPDFToObjectAsyncGetResultRequest) (response *TurnPDFToObjectAsyncGetResultResponse, err error) {
    if request == nil {
        request = NewTurnPDFToObjectAsyncGetResultRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("TurnPDFToObjectAsyncGetResult require credential")
    }

    request.SetContext(ctx)
    
    response = NewTurnPDFToObjectAsyncGetResultResponse()
    err = c.Send(request, response)
    return
}
