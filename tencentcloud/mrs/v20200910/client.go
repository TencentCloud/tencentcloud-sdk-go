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
    if request == nil {
        request = NewImageToClassRequest()
    }
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
//  UNAUTHORIZEDOPERATION_PERMISSIONDENIEDERROR = "UnauthorizedOperation.PermissionDeniedError"
//  UNSUPPORTEDOPERATION_UNSUPPORTTHISTYPE = "UnsupportedOperation.UnSupportThisType"
func (c *Client) ImageToObject(request *ImageToObjectRequest) (response *ImageToObjectResponse, err error) {
    if request == nil {
        request = NewImageToObjectRequest()
    }
    response = NewImageToObjectResponse()
    err = c.Send(request, response)
    return
}

func NewReportImageStructuredRequest() (request *ReportImageStructuredRequest) {
    request = &ReportImageStructuredRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mrs", APIVersion, "ReportImageStructured")
    return
}

func NewReportImageStructuredResponse() (response *ReportImageStructuredResponse) {
    response = &ReportImageStructuredResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportImageStructured
// 接口没有流量
//
// 
//
// 将输入的图片类型报告结构化
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INVALIDPARAMETERVALUE_IMAGECODEERROR = "InvalidParameterValue.ImageCodeError"
//  INVALIDPARAMETERVALUE_IMAGEISNOTEXT = "InvalidParameterValue.ImageIsNoText"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  INVALIDPARAMETERVALUE_TEXTISNEED = "InvalidParameterValue.TextIsNeed"
//  LIMITEXCEEDED_SERVERLIMITEXCEEDED = "LimitExceeded.ServerLimitExceeded"
//  MISSINGPARAMETER_IMAGEISNEED = "MissingParameter.ImageIsNeed"
//  OPERATIONDENIED_IMAGETOOLARGE = "OperationDenied.ImageTooLarge"
//  OPERATIONDENIED_UNSUPPORTTHISTYPE = "OperationDenied.UnSupportThisType"
//  REQUESTLIMITEXCEEDED_QPSLIMITEXCEEDED = "RequestLimitExceeded.QPSLimitExceeded"
func (c *Client) ReportImageStructured(request *ReportImageStructuredRequest) (response *ReportImageStructuredResponse, err error) {
    if request == nil {
        request = NewReportImageStructuredRequest()
    }
    response = NewReportImageStructuredResponse()
    err = c.Send(request, response)
    return
}

func NewReportTextStructuredRequest() (request *ReportTextStructuredRequest) {
    request = &ReportTextStructuredRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("mrs", APIVersion, "ReportTextStructured")
    return
}

func NewReportTextStructuredResponse() (response *ReportTextStructuredResponse) {
    response = &ReportTextStructuredResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ReportTextStructured
// 接口还未上线
//
// 
//
// 将输入的医疗报告文本内容进行结构化输出
//
// 可能返回的错误码:
//  FAILEDOPERATION_UNKNOWNERROR = "FailedOperation.UnknownError"
//  INVALIDPARAMETERVALUE_IMAGECODEERROR = "InvalidParameterValue.ImageCodeError"
//  INVALIDPARAMETERVALUE_IMAGEISNOTEXT = "InvalidParameterValue.ImageIsNoText"
//  INVALIDPARAMETERVALUE_IMAGEURLINVALID = "InvalidParameterValue.ImageURLInvalid"
//  INVALIDPARAMETERVALUE_TEXTISNEED = "InvalidParameterValue.TextIsNeed"
//  MISSINGPARAMETER_IMAGEISNEED = "MissingParameter.ImageIsNeed"
func (c *Client) ReportTextStructured(request *ReportTextStructuredRequest) (response *ReportTextStructuredResponse, err error) {
    if request == nil {
        request = NewReportTextStructuredRequest()
    }
    response = NewReportTextStructuredResponse()
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
    if request == nil {
        request = NewTextToClassRequest()
    }
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
// 文本转结构化对象
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
func (c *Client) TextToObject(request *TextToObjectRequest) (response *TextToObjectResponse, err error) {
    if request == nil {
        request = NewTextToObjectRequest()
    }
    response = NewTextToObjectResponse()
    err = c.Send(request, response)
    return
}
