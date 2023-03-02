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
