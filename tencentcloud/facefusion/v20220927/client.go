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
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
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
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
//  INVALIDPARAMETERVALUE_PARAMETERVALUEERROR = "InvalidParameterValue.ParameterValueError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
func (c *Client) DescribeMaterialListWithContext(ctx context.Context, request *DescribeMaterialListRequest) (response *DescribeMaterialListResponse, err error) {
    if request == nil {
        request = NewDescribeMaterialListRequest()
    }
    
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
// 请求频率限制为20次/秒。
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//  FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//  FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//  FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//  FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//  FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
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
func (c *Client) FuseFace(request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
    return c.FuseFaceWithContext(context.Background(), request)
}

// FuseFace
// 本接口用于单脸、多脸、选脸融合，上传人脸图片，得到与素材模板融合后的人脸图片。支持为融合结果图添加标识。查看 <a href="https://cloud.tencent.com/document/product/670/38247" target="_blank">融合接入指引</a>。
//
// 
//
// 请求频率限制为20次/秒。
//
// >
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_FACEBORDERCHECKFAILED = "FailedOperation.FaceBorderCheckFailed"
//  FAILEDOPERATION_FACEEXCEEDBORDER = "FailedOperation.FaceExceedBorder"
//  FAILEDOPERATION_FACEFUSIONERROR = "FailedOperation.FaceFusionError"
//  FAILEDOPERATION_FACERECTINVALID = "FailedOperation.FaceRectInvalid"
//  FAILEDOPERATION_FACESIZETOOSMALL = "FailedOperation.FaceSizeTooSmall"
//  FAILEDOPERATION_FUSEFREQCTRL = "FailedOperation.FuseFreqCtrl"
//  FAILEDOPERATION_FUSEIMAGEERROR = "FailedOperation.FuseImageError"
//  FAILEDOPERATION_FUSEINNERERROR = "FailedOperation.FuseInnerError"
//  FAILEDOPERATION_FUSEMATERIALNOTAUTH = "FailedOperation.FuseMaterialNotAuth"
//  FAILEDOPERATION_FUSEMATERIALNOTEXIST = "FailedOperation.FuseMaterialNotExist"
//  FAILEDOPERATION_FUSESAVEPHOTOFAIL = "FailedOperation.FuseSavePhotoFail"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEPIXELEXCEED = "FailedOperation.ImagePixelExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONTOOSMALL = "FailedOperation.ImageResolutionTooSmall"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGESIZEINVALID = "FailedOperation.ImageSizeInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOFACEDETECTED = "FailedOperation.NoFaceDetected"
//  FAILEDOPERATION_PARAMETERVALUEERROR = "FailedOperation.ParameterValueError"
//  FAILEDOPERATION_PROJECTNOTAUTH = "FailedOperation.ProjectNotAuth"
//  FAILEDOPERATION_REQUESTENTITYTOOLARGE = "FailedOperation.RequestEntityTooLarge"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TEMPLATEFACEIDNOTEXIST = "FailedOperation.TemplateFaceIDNotExist"
//  INVALIDPARAMETERVALUE_ACTIVITYIDNOTFOUND = "InvalidParameterValue.ActivityIdNotFound"
//  INVALIDPARAMETERVALUE_FACERECTPARAMETERVALUEERROR = "InvalidParameterValue.FaceRectParameterValueError"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  INVALIDPARAMETERVALUE_MATERIALIDNOTFOUND = "InvalidParameterValue.MaterialIdNotFound"
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
func (c *Client) FuseFaceWithContext(ctx context.Context, request *FuseFaceRequest) (response *FuseFaceResponse, err error) {
    if request == nil {
        request = NewFuseFaceRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("FuseFace require credential")
    }

    request.SetContext(ctx)
    
    response = NewFuseFaceResponse()
    err = c.Send(request, response)
    return
}
