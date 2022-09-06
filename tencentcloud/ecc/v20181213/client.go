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

package v20181213

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-12-13"

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


func NewCorrectMultiImageRequest() (request *CorrectMultiImageRequest) {
    request = &CorrectMultiImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecc", APIVersion, "CorrectMultiImage")
    
    
    return
}

func NewCorrectMultiImageResponse() (response *CorrectMultiImageResponse) {
    response = &CorrectMultiImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CorrectMultiImage
// https://ecc.tencentcloudapi.com/?Action=CorrectMultiImage
//
// 多图像识别批改接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INTERNALERROR_OCRERROR = "InternalError.OcrError"
//  INTERNALERROR_OCRSERVERINTERNERROR = "InternalError.OcrServerInternError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_OVERLOADERROR = "InternalError.OverLoadError"
//  INTERNALERROR_RECOGNIZEERROR = "InternalError.RecognizeError"
//  INTERNALERROR_SERVERCONNECTDOWNLOADERROR = "InternalError.ServerConnectDownloadError"
//  INTERNALERROR_SPLITERROR = "InternalError.SplitError"
//  INVALIDPARAMETER_EMPTYPARAMETERERROR = "InvalidParameter.EmptyParameterError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETERVALUE_APPIDINVALIDERROR = "InvalidParameterValue.AppidInvalidError"
//  INVALIDPARAMETERVALUE_DECODEIMAGEERROR = "InvalidParameterValue.DecodeImageError"
//  INVALIDPARAMETERVALUE_DOWNLOADIMAGEFAILERROR = "InvalidParameterValue.DownloadImageFailError"
//  INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"
//  INVALIDPARAMETERVALUE_IMAGEDOWNLOADFAILERROR = "InvalidParameterValue.ImageDownloadFailError"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEEDERROR = "InvalidParameterValue.ImageSizeExceedError"
//  INVALIDPARAMETERVALUE_IMAGETOOBIGERROR = "InvalidParameterValue.ImageTooBigError"
//  INVALIDPARAMETERVALUE_INPUTTYPEVALUEERROR = "InvalidParameterValue.InputTypeValueError"
//  INVALIDPARAMETERVALUE_SESSIONERROR = "InvalidParameterValue.SessionError"
//  INVALIDPARAMETERVALUE_URLFROMATIVADLIDERROR = "InvalidParameterValue.UrlFromatIvadlidError"
//  LIMITEXCEEDED_FREQLIMITFORBIDDENACCESSERROR = "LimitExceeded.FreqLimitForbiddenAccessError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
//  RESOURCENOTFOUND_SERVERNAMENOTEXISTINLICENSEERROR = "ResourceNotFound.ServerNameNotExistInLicenseError"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  UNAUTHORIZEDOPERATION_LICENSEINVALIDFORBIDDENACCESSERROR = "UnauthorizedOperation.LicenseInvalidForbiddenAccessError"
//  UNAUTHORIZEDOPERATION_SERVERNAMEUNAUTHORIZEDINERROR = "UnauthorizedOperation.ServerNameUnauthorizedInError"
func (c *Client) CorrectMultiImage(request *CorrectMultiImageRequest) (response *CorrectMultiImageResponse, err error) {
    return c.CorrectMultiImageWithContext(context.Background(), request)
}

// CorrectMultiImage
// https://ecc.tencentcloudapi.com/?Action=CorrectMultiImage
//
// 多图像识别批改接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INTERNALERROR_OCRERROR = "InternalError.OcrError"
//  INTERNALERROR_OCRSERVERINTERNERROR = "InternalError.OcrServerInternError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_OVERLOADERROR = "InternalError.OverLoadError"
//  INTERNALERROR_RECOGNIZEERROR = "InternalError.RecognizeError"
//  INTERNALERROR_SERVERCONNECTDOWNLOADERROR = "InternalError.ServerConnectDownloadError"
//  INTERNALERROR_SPLITERROR = "InternalError.SplitError"
//  INVALIDPARAMETER_EMPTYPARAMETERERROR = "InvalidParameter.EmptyParameterError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETERVALUE_APPIDINVALIDERROR = "InvalidParameterValue.AppidInvalidError"
//  INVALIDPARAMETERVALUE_DECODEIMAGEERROR = "InvalidParameterValue.DecodeImageError"
//  INVALIDPARAMETERVALUE_DOWNLOADIMAGEFAILERROR = "InvalidParameterValue.DownloadImageFailError"
//  INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"
//  INVALIDPARAMETERVALUE_IMAGEDOWNLOADFAILERROR = "InvalidParameterValue.ImageDownloadFailError"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEEDERROR = "InvalidParameterValue.ImageSizeExceedError"
//  INVALIDPARAMETERVALUE_IMAGETOOBIGERROR = "InvalidParameterValue.ImageTooBigError"
//  INVALIDPARAMETERVALUE_INPUTTYPEVALUEERROR = "InvalidParameterValue.InputTypeValueError"
//  INVALIDPARAMETERVALUE_SESSIONERROR = "InvalidParameterValue.SessionError"
//  INVALIDPARAMETERVALUE_URLFROMATIVADLIDERROR = "InvalidParameterValue.UrlFromatIvadlidError"
//  LIMITEXCEEDED_FREQLIMITFORBIDDENACCESSERROR = "LimitExceeded.FreqLimitForbiddenAccessError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
//  RESOURCENOTFOUND_SERVERNAMENOTEXISTINLICENSEERROR = "ResourceNotFound.ServerNameNotExistInLicenseError"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  UNAUTHORIZEDOPERATION_LICENSEINVALIDFORBIDDENACCESSERROR = "UnauthorizedOperation.LicenseInvalidForbiddenAccessError"
//  UNAUTHORIZEDOPERATION_SERVERNAMEUNAUTHORIZEDINERROR = "UnauthorizedOperation.ServerNameUnauthorizedInError"
func (c *Client) CorrectMultiImageWithContext(ctx context.Context, request *CorrectMultiImageRequest) (response *CorrectMultiImageResponse, err error) {
    if request == nil {
        request = NewCorrectMultiImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CorrectMultiImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewCorrectMultiImageResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTaskRequest() (request *DescribeTaskRequest) {
    request = &DescribeTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecc", APIVersion, "DescribeTask")
    
    
    return
}

func NewDescribeTaskResponse() (response *DescribeTaskResponse) {
    response = &DescribeTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeTask
// 异步任务结果查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATAFORMATABNORMALERROR = "InternalError.DataFormatAbnormalError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INTERNALERROR_OCRERROR = "InternalError.OcrError"
//  INTERNALERROR_OCRSERVERINTERNERROR = "InternalError.OcrServerInternError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_OVERLOADERROR = "InternalError.OverLoadError"
//  INTERNALERROR_RECOGNIZEERROR = "InternalError.RecognizeError"
//  INTERNALERROR_SERVERCONNECTDOWNLOADERROR = "InternalError.ServerConnectDownloadError"
//  INTERNALERROR_SPLITERROR = "InternalError.SplitError"
//  INVALIDPARAMETER_EMPTYPARAMETERERROR = "InvalidParameter.EmptyParameterError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETERVALUE_APPIDINVALIDERROR = "InvalidParameterValue.AppidInvalidError"
//  INVALIDPARAMETERVALUE_DECODEIMAGEERROR = "InvalidParameterValue.DecodeImageError"
//  INVALIDPARAMETERVALUE_DOWNLOADIMAGEFAILERROR = "InvalidParameterValue.DownloadImageFailError"
//  INVALIDPARAMETERVALUE_IMAGEDOWNLOADFAILERROR = "InvalidParameterValue.ImageDownloadFailError"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEEDERROR = "InvalidParameterValue.ImageSizeExceedError"
//  INVALIDPARAMETERVALUE_URLFROMATIVADLIDERROR = "InvalidParameterValue.UrlFromatIvadlidError"
//  LIMITEXCEEDED_FREQLIMITFORBIDDENACCESSERROR = "LimitExceeded.FreqLimitForbiddenAccessError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
//  RESOURCENOTFOUND_SERVERNAMENOTEXISTINLICENSEERROR = "ResourceNotFound.ServerNameNotExistInLicenseError"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  UNAUTHORIZEDOPERATION_LICENSEINVALIDFORBIDDENACCESSERROR = "UnauthorizedOperation.LicenseInvalidForbiddenAccessError"
//  UNAUTHORIZEDOPERATION_SERVERNAMEUNAUTHORIZEDINERROR = "UnauthorizedOperation.ServerNameUnauthorizedInError"
func (c *Client) DescribeTask(request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    return c.DescribeTaskWithContext(context.Background(), request)
}

// DescribeTask
// 异步任务结果查询接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATAFORMATABNORMALERROR = "InternalError.DataFormatAbnormalError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INTERNALERROR_OCRERROR = "InternalError.OcrError"
//  INTERNALERROR_OCRSERVERINTERNERROR = "InternalError.OcrServerInternError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_OVERLOADERROR = "InternalError.OverLoadError"
//  INTERNALERROR_RECOGNIZEERROR = "InternalError.RecognizeError"
//  INTERNALERROR_SERVERCONNECTDOWNLOADERROR = "InternalError.ServerConnectDownloadError"
//  INTERNALERROR_SPLITERROR = "InternalError.SplitError"
//  INVALIDPARAMETER_EMPTYPARAMETERERROR = "InvalidParameter.EmptyParameterError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETER_TASKNOTFOUND = "InvalidParameter.TaskNotFound"
//  INVALIDPARAMETERVALUE_APPIDINVALIDERROR = "InvalidParameterValue.AppidInvalidError"
//  INVALIDPARAMETERVALUE_DECODEIMAGEERROR = "InvalidParameterValue.DecodeImageError"
//  INVALIDPARAMETERVALUE_DOWNLOADIMAGEFAILERROR = "InvalidParameterValue.DownloadImageFailError"
//  INVALIDPARAMETERVALUE_IMAGEDOWNLOADFAILERROR = "InvalidParameterValue.ImageDownloadFailError"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEEDERROR = "InvalidParameterValue.ImageSizeExceedError"
//  INVALIDPARAMETERVALUE_URLFROMATIVADLIDERROR = "InvalidParameterValue.UrlFromatIvadlidError"
//  LIMITEXCEEDED_FREQLIMITFORBIDDENACCESSERROR = "LimitExceeded.FreqLimitForbiddenAccessError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
//  RESOURCENOTFOUND_SERVERNAMENOTEXISTINLICENSEERROR = "ResourceNotFound.ServerNameNotExistInLicenseError"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  UNAUTHORIZEDOPERATION_LICENSEINVALIDFORBIDDENACCESSERROR = "UnauthorizedOperation.LicenseInvalidForbiddenAccessError"
//  UNAUTHORIZEDOPERATION_SERVERNAMEUNAUTHORIZEDINERROR = "UnauthorizedOperation.ServerNameUnauthorizedInError"
func (c *Client) DescribeTaskWithContext(ctx context.Context, request *DescribeTaskRequest) (response *DescribeTaskResponse, err error) {
    if request == nil {
        request = NewDescribeTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeTaskResponse()
    err = c.Send(request, response)
    return
}

func NewECCRequest() (request *ECCRequest) {
    request = &ECCRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecc", APIVersion, "ECC")
    
    
    return
}

func NewECCResponse() (response *ECCResponse) {
    response = &ECCResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// ECC
// 接口请求域名： ecc.tencentcloudapi.com 
//
// 纯文本英语作文批改
//
// 可能返回的错误码:
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETERVALUE_SESSIONERROR = "InvalidParameterValue.SessionError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
func (c *Client) ECC(request *ECCRequest) (response *ECCResponse, err error) {
    return c.ECCWithContext(context.Background(), request)
}

// ECC
// 接口请求域名： ecc.tencentcloudapi.com 
//
// 纯文本英语作文批改
//
// 可能返回的错误码:
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETERVALUE_SESSIONERROR = "InvalidParameterValue.SessionError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
func (c *Client) ECCWithContext(ctx context.Context, request *ECCRequest) (response *ECCResponse, err error) {
    if request == nil {
        request = NewECCRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ECC require credential")
    }

    request.SetContext(ctx)
    
    response = NewECCResponse()
    err = c.Send(request, response)
    return
}

func NewEHOCRRequest() (request *EHOCRRequest) {
    request = &EHOCRRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ecc", APIVersion, "EHOCR")
    
    
    return
}

func NewEHOCRResponse() (response *EHOCRResponse) {
    response = &EHOCRResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EHOCR
// https://ecc.tencentcloudapi.com/?Action=EHOCR
//
// 图像识别批改接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INTERNALERROR_OCRERROR = "InternalError.OcrError"
//  INTERNALERROR_OCRSERVERINTERNERROR = "InternalError.OcrServerInternError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_OVERLOADERROR = "InternalError.OverLoadError"
//  INTERNALERROR_RECOGNIZEERROR = "InternalError.RecognizeError"
//  INTERNALERROR_SERVERCONNECTDOWNLOADERROR = "InternalError.ServerConnectDownloadError"
//  INTERNALERROR_SPLITERROR = "InternalError.SplitError"
//  INVALIDPARAMETER_EMPTYPARAMETERERROR = "InvalidParameter.EmptyParameterError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETERVALUE_APPIDINVALIDERROR = "InvalidParameterValue.AppidInvalidError"
//  INVALIDPARAMETERVALUE_DECODEIMAGEERROR = "InvalidParameterValue.DecodeImageError"
//  INVALIDPARAMETERVALUE_DOWNLOADIMAGEFAILERROR = "InvalidParameterValue.DownloadImageFailError"
//  INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"
//  INVALIDPARAMETERVALUE_IMAGEDOWNLOADFAILERROR = "InvalidParameterValue.ImageDownloadFailError"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEEDERROR = "InvalidParameterValue.ImageSizeExceedError"
//  INVALIDPARAMETERVALUE_IMAGETOOBIGERROR = "InvalidParameterValue.ImageTooBigError"
//  INVALIDPARAMETERVALUE_INPUTTYPEVALUEERROR = "InvalidParameterValue.InputTypeValueError"
//  INVALIDPARAMETERVALUE_SESSIONERROR = "InvalidParameterValue.SessionError"
//  INVALIDPARAMETERVALUE_URLFROMATIVADLIDERROR = "InvalidParameterValue.UrlFromatIvadlidError"
//  LIMITEXCEEDED_FREQLIMITFORBIDDENACCESSERROR = "LimitExceeded.FreqLimitForbiddenAccessError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
//  RESOURCENOTFOUND_SERVERNAMENOTEXISTINLICENSEERROR = "ResourceNotFound.ServerNameNotExistInLicenseError"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  UNAUTHORIZEDOPERATION_LICENSEINVALIDFORBIDDENACCESSERROR = "UnauthorizedOperation.LicenseInvalidForbiddenAccessError"
//  UNAUTHORIZEDOPERATION_SERVERNAMEUNAUTHORIZEDINERROR = "UnauthorizedOperation.ServerNameUnauthorizedInError"
func (c *Client) EHOCR(request *EHOCRRequest) (response *EHOCRResponse, err error) {
    return c.EHOCRWithContext(context.Background(), request)
}

// EHOCR
// https://ecc.tencentcloudapi.com/?Action=EHOCR
//
// 图像识别批改接口
//
// 可能返回的错误码:
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_CORRECTERROR = "InternalError.CorrectError"
//  INTERNALERROR_DATABASESERVERINTERNALERROR = "InternalError.DatabaseServerInternalError"
//  INTERNALERROR_OCRERROR = "InternalError.OcrError"
//  INTERNALERROR_OCRSERVERINTERNERROR = "InternalError.OcrServerInternError"
//  INTERNALERROR_OTHERERROR = "InternalError.OtherError"
//  INTERNALERROR_OVERLOADERROR = "InternalError.OverLoadError"
//  INTERNALERROR_RECOGNIZEERROR = "InternalError.RecognizeError"
//  INTERNALERROR_SERVERCONNECTDOWNLOADERROR = "InternalError.ServerConnectDownloadError"
//  INTERNALERROR_SPLITERROR = "InternalError.SplitError"
//  INVALIDPARAMETER_EMPTYPARAMETERERROR = "InvalidParameter.EmptyParameterError"
//  INVALIDPARAMETER_INPUTERROR = "InvalidParameter.InputError"
//  INVALIDPARAMETERVALUE_APPIDINVALIDERROR = "InvalidParameterValue.AppidInvalidError"
//  INVALIDPARAMETERVALUE_DECODEIMAGEERROR = "InvalidParameterValue.DecodeImageError"
//  INVALIDPARAMETERVALUE_DOWNLOADIMAGEFAILERROR = "InvalidParameterValue.DownloadImageFailError"
//  INVALIDPARAMETERVALUE_EMPTYIMAGEERROR = "InvalidParameterValue.EmptyImageError"
//  INVALIDPARAMETERVALUE_IMAGEDOWNLOADFAILERROR = "InvalidParameterValue.ImageDownloadFailError"
//  INVALIDPARAMETERVALUE_IMAGESIZEEXCEEDERROR = "InvalidParameterValue.ImageSizeExceedError"
//  INVALIDPARAMETERVALUE_IMAGETOOBIGERROR = "InvalidParameterValue.ImageTooBigError"
//  INVALIDPARAMETERVALUE_INPUTTYPEVALUEERROR = "InvalidParameterValue.InputTypeValueError"
//  INVALIDPARAMETERVALUE_SESSIONERROR = "InvalidParameterValue.SessionError"
//  INVALIDPARAMETERVALUE_URLFROMATIVADLIDERROR = "InvalidParameterValue.UrlFromatIvadlidError"
//  LIMITEXCEEDED_FREQLIMITFORBIDDENACCESSERROR = "LimitExceeded.FreqLimitForbiddenAccessError"
//  RESOURCENOTFOUND_CANNOTFINDUSER = "ResourceNotFound.CannotFindUser"
//  RESOURCENOTFOUND_SERVERNAMENOTEXISTINLICENSEERROR = "ResourceNotFound.ServerNameNotExistInLicenseError"
//  RESOURCEUNAVAILABLE_AUTHORIZEERROR = "ResourceUnavailable.AuthorizeError"
//  UNAUTHORIZEDOPERATION_LICENSEINVALIDFORBIDDENACCESSERROR = "UnauthorizedOperation.LicenseInvalidForbiddenAccessError"
//  UNAUTHORIZEDOPERATION_SERVERNAMEUNAUTHORIZEDINERROR = "UnauthorizedOperation.ServerNameUnauthorizedInError"
func (c *Client) EHOCRWithContext(ctx context.Context, request *EHOCRRequest) (response *EHOCRResponse, err error) {
    if request == nil {
        request = NewEHOCRRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EHOCR require credential")
    }

    request.SetContext(ctx)
    
    response = NewEHOCRResponse()
    err = c.Send(request, response)
    return
}
