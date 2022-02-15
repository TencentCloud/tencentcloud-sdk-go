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

package v20190529

import (
    "context"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-05-29"

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


func NewAssessQualityRequest() (request *AssessQualityRequest) {
    request = &AssessQualityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "AssessQuality")
    
    
    return
}

func NewAssessQualityResponse() (response *AssessQualityResponse) {
    response = &AssessQualityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// AssessQuality
// 评估输入图片在视觉上的质量，从多个方面评估，并同时给出综合的、客观的清晰度评分，和主观的美观度评分。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) AssessQuality(request *AssessQualityRequest) (response *AssessQualityResponse, err error) {
    if request == nil {
        request = NewAssessQualityRequest()
    }
    
    response = NewAssessQualityResponse()
    err = c.Send(request, response)
    return
}

// AssessQuality
// 评估输入图片在视觉上的质量，从多个方面评估，并同时给出综合的、客观的清晰度评分，和主观的美观度评分。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) AssessQualityWithContext(ctx context.Context, request *AssessQualityRequest) (response *AssessQualityResponse, err error) {
    if request == nil {
        request = NewAssessQualityRequest()
    }
    request.SetContext(ctx)
    
    response = NewAssessQualityResponse()
    err = c.Send(request, response)
    return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
    request = &CreateGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "CreateGroup")
    
    
    return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
    response = &CreateGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateGroup
// 用于创建一个空的图片库，如果图片库已存在则返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_PARAMETEREMPTY = "FailedOperation.ParameterEmpty"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_BRIEFTOOLONG = "InvalidParameterValue.BriefTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDALREADYEXIST = "InvalidParameterValue.ImageGroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPNAMEEMPTY = "InvalidParameterValue.ImageGroupNameEmpty"
//  INVALIDPARAMETERVALUE_IMAGEGROUPNAMETOOLONG = "InvalidParameterValue.ImageGroupNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

// CreateGroup
// 用于创建一个空的图片库，如果图片库已存在则返回错误。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_PARAMETEREMPTY = "FailedOperation.ParameterEmpty"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_BRIEFTOOLONG = "InvalidParameterValue.BriefTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDALREADYEXIST = "InvalidParameterValue.ImageGroupIdAlreadyExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPNAMEEMPTY = "InvalidParameterValue.ImageGroupNameEmpty"
//  INVALIDPARAMETERVALUE_IMAGEGROUPNAMETOOLONG = "InvalidParameterValue.ImageGroupNameTooLong"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateImageRequest() (request *CreateImageRequest) {
    request = &CreateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "CreateImage")
    
    
    return
}

func NewCreateImageResponse() (response *CreateImageResponse) {
    response = &CreateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CreateImage
// 创建图片，并添加对应图片的自定义信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEENTITYCOUNTEXCEED = "FailedOperation.ImageEntityCountExceed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGENUMEXCEED = "FailedOperation.ImageNumExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_CUSTOMCONTENTTOOLONG = "InvalidParameterValue.CustomContentTooLong"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMEALREADYEXIST = "InvalidParameterValue.PicNameAlreadyExist"
//  INVALIDPARAMETERVALUE_PICNAMEEMPTY = "InvalidParameterValue.PicNameEmpty"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
//  INVALIDPARAMETERVALUE_TAGSKEYSEXCEED = "InvalidParameterValue.TagsKeysExceed"
//  INVALIDPARAMETERVALUE_TAGSVALUEILLEGAL = "InvalidParameterValue.TagsValueIllegal"
//  INVALIDPARAMETERVALUE_TAGSVALUESIZEEXCEED = "InvalidParameterValue.TagsValueSizeExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

// CreateImage
// 创建图片，并添加对应图片的自定义信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEENTITYCOUNTEXCEED = "FailedOperation.ImageEntityCountExceed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGENUMEXCEED = "FailedOperation.ImageNumExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_CUSTOMCONTENTTOOLONG = "InvalidParameterValue.CustomContentTooLong"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMEALREADYEXIST = "InvalidParameterValue.PicNameAlreadyExist"
//  INVALIDPARAMETERVALUE_PICNAMEEMPTY = "InvalidParameterValue.PicNameEmpty"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
//  INVALIDPARAMETERVALUE_TAGSKEYSEXCEED = "InvalidParameterValue.TagsKeysExceed"
//  INVALIDPARAMETERVALUE_TAGSVALUEILLEGAL = "InvalidParameterValue.TagsValueIllegal"
//  INVALIDPARAMETERVALUE_TAGSVALUESIZEEXCEED = "InvalidParameterValue.TagsValueSizeExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewCreateImageResponse()
    err = c.Send(request, response)
    return
}

func NewCropImageRequest() (request *CropImageRequest) {
    request = &CropImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "CropImage")
    
    
    return
}

func NewCropImageResponse() (response *CropImageResponse) {
    response = &CropImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// CropImage
// 根据输入的裁剪比例，智能判断一张图片的最佳裁剪区域，确保原图的主体区域不受影响。
//
// 
//
// 可以自动裁剪图片，适应不同平台、设备的展示要求，避免简单拉伸带来的变形。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CropImage(request *CropImageRequest) (response *CropImageResponse, err error) {
    if request == nil {
        request = NewCropImageRequest()
    }
    
    response = NewCropImageResponse()
    err = c.Send(request, response)
    return
}

// CropImage
// 根据输入的裁剪比例，智能判断一张图片的最佳裁剪区域，确保原图的主体区域不受影响。
//
// 
//
// 可以自动裁剪图片，适应不同平台、设备的展示要求，避免简单拉伸带来的变形。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CropImageWithContext(ctx context.Context, request *CropImageRequest) (response *CropImageResponse, err error) {
    if request == nil {
        request = NewCropImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewCropImageResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteImagesRequest() (request *DeleteImagesRequest) {
    request = &DeleteImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DeleteImages")
    
    
    return
}

func NewDeleteImagesResponse() (response *DeleteImagesResponse) {
    response = &DeleteImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DeleteImages
// 删除图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDELETEFAILED = "FailedOperation.ImageDeleteFailed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    
    response = NewDeleteImagesResponse()
    err = c.Send(request, response)
    return
}

// DeleteImages
// 删除图片。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDELETEFAILED = "FailedOperation.ImageDeleteFailed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
func (c *Client) DeleteImagesWithContext(ctx context.Context, request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDeleteImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
    request = &DescribeGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DescribeGroups")
    
    
    return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
    response = &DescribeGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeGroups
// 查询所有的图库信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

// DescribeGroups
// 查询所有的图库信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeGroupsWithContext(ctx context.Context, request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
    request = &DescribeImagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DescribeImages")
    
    
    return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
    response = &DescribeImagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DescribeImages
// 获取指定图片库中的图片列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMEEMPTY = "InvalidParameterValue.PicNameEmpty"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

// DescribeImages
// 获取指定图片库中的图片列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMEEMPTY = "InvalidParameterValue.PicNameEmpty"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDetectCelebrityRequest() (request *DetectCelebrityRequest) {
    request = &DetectCelebrityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DetectCelebrity")
    
    
    return
}

func NewDetectCelebrityResponse() (response *DetectCelebrityResponse) {
    response = &DetectCelebrityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectCelebrity
// 传入一张图片，可以识别图片中包含的人物是否为公众人物，如果是，输出人物的姓名、基本信息、脸部坐标。
//
// 
//
// 支持识别一张图片中存在的多个人脸，针对每个人脸，会给出与之最相似的公众人物。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectCelebrity(request *DetectCelebrityRequest) (response *DetectCelebrityResponse, err error) {
    if request == nil {
        request = NewDetectCelebrityRequest()
    }
    
    response = NewDetectCelebrityResponse()
    err = c.Send(request, response)
    return
}

// DetectCelebrity
// 传入一张图片，可以识别图片中包含的人物是否为公众人物，如果是，输出人物的姓名、基本信息、脸部坐标。
//
// 
//
// 支持识别一张图片中存在的多个人脸，针对每个人脸，会给出与之最相似的公众人物。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_NOFACEINPHOTO = "InvalidParameterValue.NoFaceInPhoto"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectCelebrityWithContext(ctx context.Context, request *DetectCelebrityRequest) (response *DetectCelebrityResponse, err error) {
    if request == nil {
        request = NewDetectCelebrityRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetectCelebrityResponse()
    err = c.Send(request, response)
    return
}

func NewDetectDisgustRequest() (request *DetectDisgustRequest) {
    request = &DetectDisgustRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DetectDisgust")
    
    
    return
}

func NewDetectDisgustResponse() (response *DetectDisgustResponse) {
    response = &DetectDisgustResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectDisgust
// 输入一张图片，返回AI针对一张图片是否是恶心的一系列判断值。
//
// 
//
// 通过恶心图片识别, 可以判断一张图片是否令人恶心, 同时给出它属于的潜在类别, 让您能够过滤掉使人不愉快的图片。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectDisgust(request *DetectDisgustRequest) (response *DetectDisgustResponse, err error) {
    if request == nil {
        request = NewDetectDisgustRequest()
    }
    
    response = NewDetectDisgustResponse()
    err = c.Send(request, response)
    return
}

// DetectDisgust
// 输入一张图片，返回AI针对一张图片是否是恶心的一系列判断值。
//
// 
//
// 通过恶心图片识别, 可以判断一张图片是否令人恶心, 同时给出它属于的潜在类别, 让您能够过滤掉使人不愉快的图片。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectDisgustWithContext(ctx context.Context, request *DetectDisgustRequest) (response *DetectDisgustResponse, err error) {
    if request == nil {
        request = NewDetectDisgustRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetectDisgustResponse()
    err = c.Send(request, response)
    return
}

func NewDetectLabelRequest() (request *DetectLabelRequest) {
    request = &DetectLabelRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DetectLabel")
    
    
    return
}

func NewDetectLabelResponse() (response *DetectLabelResponse) {
    response = &DetectLabelResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectLabel
// 图像标签利用深度学习技术、海量训练数据，可以对图片进行智能分类、物体识别等。
//
// 
//
// 目前支持8个大类、六十多个子类、数千个标签。涵盖各种日常场景、动植物、物品、美食、卡证等。具体分类请见[图像分析常见问题功能与限制相关](https://cloud.tencent.com/document/product/865/39164)。
//
// 
//
// 图像标签提供四个版本供选择：
//
// 
//
// • 摄像头版：针对搜索、手机摄像头照片进行优化，涵盖大量卡证、日常物品、二维码条形码。
//
// 
//
// • 相册版：针对手机相册、网盘进行优化，去除相册和网盘中不常见的标签，针对相册常见图片类型（人像、日常活动、日常物品等）识别效果更好。
//
// 
//
// • 网络版：针对网络图片进行优化，涵盖标签更多，满足长尾识别需求。
//
// 
//
// • 新闻版：针对新闻、资讯、广电等行业进行优化，增加定制识别，支持万级图像标签。
//
// 
//
// 每个产品的图像类型都有独特性，建议在接入初期，对四个版本进行对比评估后选择合适的版本使用。
//
// 
//
// 为了方便使用、减少图片传输次数，图像标签包装成多合一接口，实际上是多个服务。
//
// 
//
// 图像标签按照服务的实际使用数量进行收费。例如一张图片同时调用相册版、摄像头版两个服务，那么此次调用按照两次计费。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectLabel(request *DetectLabelRequest) (response *DetectLabelResponse, err error) {
    if request == nil {
        request = NewDetectLabelRequest()
    }
    
    response = NewDetectLabelResponse()
    err = c.Send(request, response)
    return
}

// DetectLabel
// 图像标签利用深度学习技术、海量训练数据，可以对图片进行智能分类、物体识别等。
//
// 
//
// 目前支持8个大类、六十多个子类、数千个标签。涵盖各种日常场景、动植物、物品、美食、卡证等。具体分类请见[图像分析常见问题功能与限制相关](https://cloud.tencent.com/document/product/865/39164)。
//
// 
//
// 图像标签提供四个版本供选择：
//
// 
//
// • 摄像头版：针对搜索、手机摄像头照片进行优化，涵盖大量卡证、日常物品、二维码条形码。
//
// 
//
// • 相册版：针对手机相册、网盘进行优化，去除相册和网盘中不常见的标签，针对相册常见图片类型（人像、日常活动、日常物品等）识别效果更好。
//
// 
//
// • 网络版：针对网络图片进行优化，涵盖标签更多，满足长尾识别需求。
//
// 
//
// • 新闻版：针对新闻、资讯、广电等行业进行优化，增加定制识别，支持万级图像标签。
//
// 
//
// 每个产品的图像类型都有独特性，建议在接入初期，对四个版本进行对比评估后选择合适的版本使用。
//
// 
//
// 为了方便使用、减少图片传输次数，图像标签包装成多合一接口，实际上是多个服务。
//
// 
//
// 图像标签按照服务的实际使用数量进行收费。例如一张图片同时调用相册版、摄像头版两个服务，那么此次调用按照两次计费。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectLabelWithContext(ctx context.Context, request *DetectLabelRequest) (response *DetectLabelResponse, err error) {
    if request == nil {
        request = NewDetectLabelRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetectLabelResponse()
    err = c.Send(request, response)
    return
}

func NewDetectLabelBetaRequest() (request *DetectLabelBetaRequest) {
    request = &DetectLabelBetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DetectLabelBeta")
    
    
    return
}

func NewDetectLabelBetaResponse() (response *DetectLabelBetaResponse) {
    response = &DetectLabelBetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectLabelBeta
// 图像标签测试接口
//
// 
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectLabelBeta(request *DetectLabelBetaRequest) (response *DetectLabelBetaResponse, err error) {
    if request == nil {
        request = NewDetectLabelBetaRequest()
    }
    
    response = NewDetectLabelBetaResponse()
    err = c.Send(request, response)
    return
}

// DetectLabelBeta
// 图像标签测试接口
//
// 
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectLabelBetaWithContext(ctx context.Context, request *DetectLabelBetaRequest) (response *DetectLabelBetaResponse, err error) {
    if request == nil {
        request = NewDetectLabelBetaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetectLabelBetaResponse()
    err = c.Send(request, response)
    return
}

func NewDetectMisbehaviorRequest() (request *DetectMisbehaviorRequest) {
    request = &DetectMisbehaviorRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DetectMisbehavior")
    
    
    return
}

func NewDetectMisbehaviorResponse() (response *DetectMisbehaviorResponse) {
    response = &DetectMisbehaviorResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectMisbehavior
// 可以识别输入的图片中是否包含不良行为，例如打架斗殴、赌博、抽烟等，可以应用于广告图、直播截图、短视频截图等审核，减少不良行为对平台内容质量的影响，维护健康向上的互联网环境。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectMisbehavior(request *DetectMisbehaviorRequest) (response *DetectMisbehaviorResponse, err error) {
    if request == nil {
        request = NewDetectMisbehaviorRequest()
    }
    
    response = NewDetectMisbehaviorResponse()
    err = c.Send(request, response)
    return
}

// DetectMisbehavior
// 可以识别输入的图片中是否包含不良行为，例如打架斗殴、赌博、抽烟等，可以应用于广告图、直播截图、短视频截图等审核，减少不良行为对平台内容质量的影响，维护健康向上的互联网环境。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectMisbehaviorWithContext(ctx context.Context, request *DetectMisbehaviorRequest) (response *DetectMisbehaviorResponse, err error) {
    if request == nil {
        request = NewDetectMisbehaviorRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetectMisbehaviorResponse()
    err = c.Send(request, response)
    return
}

func NewDetectProductRequest() (request *DetectProductRequest) {
    request = &DetectProductRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DetectProduct")
    
    
    return
}

func NewDetectProductResponse() (response *DetectProductResponse) {
    response = &DetectProductResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectProduct
// 本接口支持识别图片中包含的商品，能够输出商品的品类名称、类别，还可以输出商品在图片中的位置。支持一张图片多个商品的识别。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectProduct(request *DetectProductRequest) (response *DetectProductResponse, err error) {
    if request == nil {
        request = NewDetectProductRequest()
    }
    
    response = NewDetectProductResponse()
    err = c.Send(request, response)
    return
}

// DetectProduct
// 本接口支持识别图片中包含的商品，能够输出商品的品类名称、类别，还可以输出商品在图片中的位置。支持一张图片多个商品的识别。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectProductWithContext(ctx context.Context, request *DetectProductRequest) (response *DetectProductResponse, err error) {
    if request == nil {
        request = NewDetectProductRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetectProductResponse()
    err = c.Send(request, response)
    return
}

func NewDetectProductBetaRequest() (request *DetectProductBetaRequest) {
    request = &DetectProductBetaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "DetectProductBeta")
    
    
    return
}

func NewDetectProductBetaResponse() (response *DetectProductBetaResponse) {
    response = &DetectProductBetaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectProductBeta
// 商品识别-微信识物版，基于人工智能技术、海量训练图片、亿级商品库，可以实现全覆盖、细粒度、高准确率的商品识别和商品推荐功能。
//
// 本服务可以识别出图片中的主体位置、主体商品类型，覆盖亿级SKU，输出具体商品的价格、型号等详细信息。
//
// 客户无需自建商品库，即可快速实现商品识别、拍照搜商品等功能。
//
// >   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectProductBeta(request *DetectProductBetaRequest) (response *DetectProductBetaResponse, err error) {
    if request == nil {
        request = NewDetectProductBetaRequest()
    }
    
    response = NewDetectProductBetaResponse()
    err = c.Send(request, response)
    return
}

// DetectProductBeta
// 商品识别-微信识物版，基于人工智能技术、海量训练图片、亿级商品库，可以实现全覆盖、细粒度、高准确率的商品识别和商品推荐功能。
//
// 本服务可以识别出图片中的主体位置、主体商品类型，覆盖亿级SKU，输出具体商品的价格、型号等详细信息。
//
// 客户无需自建商品库，即可快速实现商品识别、拍照搜商品等功能。
//
// >   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectProductBetaWithContext(ctx context.Context, request *DetectProductBetaRequest) (response *DetectProductBetaResponse, err error) {
    if request == nil {
        request = NewDetectProductBetaRequest()
    }
    request.SetContext(ctx)
    
    response = NewDetectProductBetaResponse()
    err = c.Send(request, response)
    return
}

func NewEnhanceImageRequest() (request *EnhanceImageRequest) {
    request = &EnhanceImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "EnhanceImage")
    
    
    return
}

func NewEnhanceImageResponse() (response *EnhanceImageResponse) {
    response = &EnhanceImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// EnhanceImage
// 传入一张图片，输出清晰度提升后的图片。
//
// 
//
// 可以消除图片有损压缩导致的噪声，和使用滤镜、拍摄失焦导致的模糊。让图片的边缘和细节更加清晰自然。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnhanceImage(request *EnhanceImageRequest) (response *EnhanceImageResponse, err error) {
    if request == nil {
        request = NewEnhanceImageRequest()
    }
    
    response = NewEnhanceImageResponse()
    err = c.Send(request, response)
    return
}

// EnhanceImage
// 传入一张图片，输出清晰度提升后的图片。
//
// 
//
// 可以消除图片有损压缩导致的噪声，和使用滤镜、拍摄失焦导致的模糊。让图片的边缘和细节更加清晰自然。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnhanceImageWithContext(ctx context.Context, request *EnhanceImageRequest) (response *EnhanceImageResponse, err error) {
    if request == nil {
        request = NewEnhanceImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewEnhanceImageResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeCarRequest() (request *RecognizeCarRequest) {
    request = &RecognizeCarRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "RecognizeCar")
    
    
    return
}

func NewRecognizeCarResponse() (response *RecognizeCarResponse) {
    response = &RecognizeCarResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeCar
// 腾讯云车辆属性识别可对汽车车身及车辆属性进行检测与识别，目前支持11种车身颜色、20多种车型、300多种品牌、4000多种车系+年款的识别，同时支持对车辆的位置进行检测。如果图片中存在多辆车，会分别输出每辆车的车型和坐标。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_RECOGNIZEFAILDED = "FailedOperation.RecognizeFailded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeCar(request *RecognizeCarRequest) (response *RecognizeCarResponse, err error) {
    if request == nil {
        request = NewRecognizeCarRequest()
    }
    
    response = NewRecognizeCarResponse()
    err = c.Send(request, response)
    return
}

// RecognizeCar
// 腾讯云车辆属性识别可对汽车车身及车辆属性进行检测与识别，目前支持11种车身颜色、20多种车型、300多种品牌、4000多种车系+年款的识别，同时支持对车辆的位置进行检测。如果图片中存在多辆车，会分别输出每辆车的车型和坐标。
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_RECOGNIZEFAILDED = "FailedOperation.RecognizeFailded"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeCarWithContext(ctx context.Context, request *RecognizeCarRequest) (response *RecognizeCarResponse, err error) {
    if request == nil {
        request = NewRecognizeCarRequest()
    }
    request.SetContext(ctx)
    
    response = NewRecognizeCarResponse()
    err = c.Send(request, response)
    return
}

func NewSearchImageRequest() (request *SearchImageRequest) {
    request = &SearchImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "SearchImage")
    
    
    return
}

func NewSearchImageResponse() (response *SearchImageResponse) {
    response = &SearchImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// SearchImage
// 本接口用于对一张待识别的商品图片，在指定图片库中检索出最相似的图片列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEGROUPEMPTY = "FailedOperation.ImageGroupEmpty"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESEARCHINVALID = "FailedOperation.ImageSearchInvalid"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_FILTERINVALID = "InvalidParameterValue.FilterInvalid"
//  INVALIDPARAMETERVALUE_FILTERSIZEEXCEED = "InvalidParameterValue.FilterSizeExceed"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SearchImage(request *SearchImageRequest) (response *SearchImageResponse, err error) {
    if request == nil {
        request = NewSearchImageRequest()
    }
    
    response = NewSearchImageResponse()
    err = c.Send(request, response)
    return
}

// SearchImage
// 本接口用于对一张待识别的商品图片，在指定图片库中检索出最相似的图片列表。
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEGROUPEMPTY = "FailedOperation.ImageGroupEmpty"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESEARCHINVALID = "FailedOperation.ImageSearchInvalid"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_FILTERINVALID = "InvalidParameterValue.FilterInvalid"
//  INVALIDPARAMETERVALUE_FILTERSIZEEXCEED = "InvalidParameterValue.FilterSizeExceed"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SearchImageWithContext(ctx context.Context, request *SearchImageRequest) (response *SearchImageResponse, err error) {
    if request == nil {
        request = NewSearchImageRequest()
    }
    request.SetContext(ctx)
    
    response = NewSearchImageResponse()
    err = c.Send(request, response)
    return
}
