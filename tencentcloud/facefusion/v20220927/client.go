// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20220927

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2022-09-27"

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


func NewDescribeMaterialListRequest() (request *DescribeMaterialListRequest) {
    request = &DescribeMaterialListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("facefusion", APIVersion, "DescribeMaterialList")
    
    
    return
}

func NewDescribeMaterialListResponse() (response *DescribeMaterialListResponse) {
    response = &DescribeMaterialListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// DescribeMaterialList
// 通常通过腾讯云人脸融合的控制台可以查看到素材相关的参数数据，可以满足使用。本接口返回活动的素材数据，包括素材状态等。用于用户通过Api查看素材相关数据，方便使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeMaterialList(request *DescribeMaterialListRequest) (response *DescribeMaterialListResponse, err error) {
    return c.DescribeMaterialListWithContext(context.Background(), request)
}

// DescribeMaterialList
// 通常通过腾讯云人脸融合的控制台可以查看到素材相关的参数数据，可以满足使用。本接口返回活动的素材数据，包括素材状态等。用于用户通过Api查看素材相关数据，方便使用。
//
// 可能返回的错误码:
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeMaterialListWithContext(ctx context.Context, request *DescribeMaterialListRequest) (response *DescribeMaterialListResponse, err error) {
    if request == nil {
        request = NewDescribeMaterialListRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "facefusion", APIVersion, "DescribeMaterialList")
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeMaterialList require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeMaterialListResponse()
    err = c.Send(request, response)
    return
}

func NewFuseFaceRequest() (request *FuseFaceRequest) {
    request = &FuseFaceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("facefusion", APIVersion, "FuseFace")
    
    
    return
}

func NewFuseFaceResponse() (response *FuseFaceResponse) {
    response = &FuseFaceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FuseFace
// 本接口用于单脸、多脸、选脸融合，上传人脸图片，得到与素材模板融合后的人脸图片。支持为融合结果图添加标识。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">融合接入指引</a>。
//
// 
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) FuseFace(request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
    return c.FuseFaceWithContext(context.Background(), request)
}

// FuseFace
// 本接口用于单脸、多脸、选脸融合，上传人脸图片，得到与素材模板融合后的人脸图片。支持为融合结果图添加标识。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">融合接入指引</a>。
//
// 
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) FuseFaceWithContext(ctx context.Context, request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
    if request == nil {
        request = NewFuseFaceRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "facefusion", APIVersion, "FuseFace")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FuseFace require credential")
    }

    request.SetContext(ctx)
    
    response = NewFuseFaceResponse()
    err = c.Send(request, response)
    return
}

func NewFuseFaceUltraRequest() (request *FuseFaceUltraRequest) {
    request = &FuseFaceUltraRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("facefusion", APIVersion, "FuseFaceUltra")
    
    
    return
}

func NewFuseFaceUltraResponse() (response *FuseFaceUltraResponse) {
    response = &FuseFaceUltraResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// FuseFaceUltra
// 图片人脸融合（专业版）为同步接口，支持自定义美颜、人脸增强、牙齿增强、拉脸等参数，最高支持8K分辨率，有多个模型类型供选择。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">融合接入指引</a>。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MEDIADATAERROR = "FailedOperation.MediaDataError"
//  FAILEDOPERATION_MODERATIONCONFIGERROR = "FailedOperation.ModerationConfigError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_MODERATIONRESULTCONFIGERROR = "FailedOperation.ModerationResultConfigError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) FuseFaceUltra(request *FuseFaceUltraRequest) (response *FuseFaceUltraResponse, err error) {
    return c.FuseFaceUltraWithContext(context.Background(), request)
}

// FuseFaceUltra
// 图片人脸融合（专业版）为同步接口，支持自定义美颜、人脸增强、牙齿增强、拉脸等参数，最高支持8K分辨率，有多个模型类型供选择。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">融合接入指引</a>。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_MEDIADATAERROR = "FailedOperation.MediaDataError"
//  FAILEDOPERATION_MODERATIONCONFIGERROR = "FailedOperation.ModerationConfigError"
//  FAILEDOPERATION_MODERATIONFAILED = "FailedOperation.ModerationFailed"
//  FAILEDOPERATION_MODERATIONRESPONSEERROR = "FailedOperation.ModerationResponseError"
//  FAILEDOPERATION_MODERATIONRESULTCONFIGERROR = "FailedOperation.ModerationResultConfigError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE_FREEZE = "ResourceUnavailable.Freeze"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCEUNAVAILABLE_RECOVER = "ResourceUnavailable.Recover"
//  RESOURCEUNAVAILABLE_STOPUSING = "ResourceUnavailable.StopUsing"
func (c *Client) FuseFaceUltraWithContext(ctx context.Context, request *FuseFaceUltraRequest) (response *FuseFaceUltraResponse, err error) {
    if request == nil {
        request = NewFuseFaceUltraRequest()
    }
    c.InitBaseRequest(&request.BaseRequest, "facefusion", APIVersion, "FuseFaceUltra")
    
    if c.GetCredential() == nil {
        return nil, errors.New("FuseFaceUltra require credential")
    }

    request.SetContext(ctx)
    
    response = NewFuseFaceUltraResponse()
    err = c.Send(request, response)
    return
}
