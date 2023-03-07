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
    "errors"
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
// 
//
// >   
//
// - 可前往 [图像处理](https://cloud.tencent.com/document/product/1590) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) AssessQuality(request *AssessQualityRequest) (response *AssessQualityResponse, err error) {
    return c.AssessQualityWithContext(context.Background(), request)
}

// AssessQuality
// 评估输入图片在视觉上的质量，从多个方面评估，并同时给出综合的、客观的清晰度评分，和主观的美观度评分。
//
// 
//
// >   
//
// - 可前往 [图像处理](https://cloud.tencent.com/document/product/1590) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("AssessQuality require credential")
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
// 本接口用于创建一个空的图片库，图片库主要用于存储在创建图片时提取的图片特征数据，如果图片库已存在则返回错误。不同的图片库类型对应不同的图像搜索服务类型，根据输入参数GroupType区分。
//
// 
//
// <table>
//
//     <th>服务类型</th><th>GroupType</th><th>功能描述</th>
//
//     <tr>
//
//         <td>通用图像搜索</td>
//
//         <td>4</td>
//
//         <td>通用图像搜索1.0版。<br>在自建图片库中搜索相同原图或相似图片集，并给出相似度打分，可支持裁剪、翻转、调色、加水印等二次编辑后的图片搜索。适用于图片版权保护、原图查询等场景。</td>
//
//     </tr>
//
//     <tr>
//
//         <td rowspan="2">商品图像搜索</td>
//
//         <td>7</td><td>商品图像搜索2.0升级版。<br>
//
//         在自建图库中搜索同款或相似商品，并给出相似度打分。对于服饰类商品可支持识别服饰类别、属性等信息。适用于商品分类、检索、推荐等电商场景。</td>
//
//     </tr>
//
//      <tr>
//
//         <td>5</td>
//
//         <td>商品图像搜索1.0版。<br>
//
//         功能和2.0升级版类似。</td>
//
//     </tr>
//
//     <tr>
//
//     <td>图案花纹搜索</td><td>6</td><td>图案花纹搜索1.0版。<br>
//
//     在自建图库中搜索相似的图案、logo、纹理等图像元素或主体，并给出相似度打分。</td>
//
//     </tr>
//
// </table> 
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETEREMPTY = "FailedOperation.ParameterEmpty"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
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
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    return c.CreateGroupWithContext(context.Background(), request)
}

// CreateGroup
// 本接口用于创建一个空的图片库，图片库主要用于存储在创建图片时提取的图片特征数据，如果图片库已存在则返回错误。不同的图片库类型对应不同的图像搜索服务类型，根据输入参数GroupType区分。
//
// 
//
// <table>
//
//     <th>服务类型</th><th>GroupType</th><th>功能描述</th>
//
//     <tr>
//
//         <td>通用图像搜索</td>
//
//         <td>4</td>
//
//         <td>通用图像搜索1.0版。<br>在自建图片库中搜索相同原图或相似图片集，并给出相似度打分，可支持裁剪、翻转、调色、加水印等二次编辑后的图片搜索。适用于图片版权保护、原图查询等场景。</td>
//
//     </tr>
//
//     <tr>
//
//         <td rowspan="2">商品图像搜索</td>
//
//         <td>7</td><td>商品图像搜索2.0升级版。<br>
//
//         在自建图库中搜索同款或相似商品，并给出相似度打分。对于服饰类商品可支持识别服饰类别、属性等信息。适用于商品分类、检索、推荐等电商场景。</td>
//
//     </tr>
//
//      <tr>
//
//         <td>5</td>
//
//         <td>商品图像搜索1.0版。<br>
//
//         功能和2.0升级版类似。</td>
//
//     </tr>
//
//     <tr>
//
//     <td>图案花纹搜索</td><td>6</td><td>图案花纹搜索1.0版。<br>
//
//     在自建图库中搜索相似的图案、logo、纹理等图像元素或主体，并给出相似度打分。</td>
//
//     </tr>
//
// </table> 
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_PARAMETEREMPTY = "FailedOperation.ParameterEmpty"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
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
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CreateGroupWithContext(ctx context.Context, request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
    if request == nil {
        request = NewCreateGroupRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateGroup require credential")
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
// 创建图片，并添加对应图片的自定义信息。模型将在创建图片时自动提取图像特征并存储到指定的图片库中，每创建一张图片会对应提取和存储一条图片特征数据。
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEENTITYCOUNTEXCEED = "FailedOperation.ImageEntityCountExceed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGENUMEXCEED = "FailedOperation.ImageNumExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_IMAGEFORMATNOTSUPPORT = "InvalidParameter.ImageFormatNotSupport"
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
    return c.CreateImageWithContext(context.Background(), request)
}

// CreateImage
// 创建图片，并添加对应图片的自定义信息。模型将在创建图片时自动提取图像特征并存储到指定的图片库中，每创建一张图片会对应提取和存储一条图片特征数据。
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEDOWNLOADERROR = "FailedOperation.ImageDownloadError"
//  FAILEDOPERATION_IMAGEENTITYCOUNTEXCEED = "FailedOperation.ImageEntityCountExceed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGENUMEXCEED = "FailedOperation.ImageNumExceed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_IMAGEFORMATNOTSUPPORT = "InvalidParameter.ImageFormatNotSupport"
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) CreateImageWithContext(ctx context.Context, request *CreateImageRequest) (response *CreateImageResponse, err error) {
    if request == nil {
        request = NewCreateImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImage require credential")
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
// 根据输入的裁剪比例，智能判断一张图片的最佳裁剪区域，确保原图的主体区域不受影响，以适应不同平台、设备的展示要求，避免简单拉伸带来的变形。
//
// 
//
// >   
//
// - 可前往 [图像处理](https://cloud.tencent.com/document/product/1590) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) CropImage(request *CropImageRequest) (response *CropImageResponse, err error) {
    return c.CropImageWithContext(context.Background(), request)
}

// CropImage
// 根据输入的裁剪比例，智能判断一张图片的最佳裁剪区域，确保原图的主体区域不受影响，以适应不同平台、设备的展示要求，避免简单拉伸带来的变形。
//
// 
//
// >   
//
// - 可前往 [图像处理](https://cloud.tencent.com/document/product/1590) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("CropImage require credential")
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
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDELETEFAILED = "FailedOperation.ImageDeleteFailed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DeleteImages(request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    return c.DeleteImagesWithContext(context.Background(), request)
}

// DeleteImages
// 删除图片。
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDELETEFAILED = "FailedOperation.ImageDeleteFailed"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DeleteImagesWithContext(ctx context.Context, request *DeleteImagesRequest) (response *DeleteImagesResponse, err error) {
    if request == nil {
        request = NewDeleteImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DeleteImages require credential")
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
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    return c.DescribeGroupsWithContext(context.Background(), request)
}

// DescribeGroups
// 查询所有的图库信息。
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 可能返回的错误码:
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeGroupsWithContext(ctx context.Context, request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeGroupsRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeGroups require credential")
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
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    return c.DescribeImagesWithContext(context.Background(), request)
}

// DescribeImages
// 获取指定图片库中的图片列表。
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
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
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DescribeImagesWithContext(ctx context.Context, request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
    if request == nil {
        request = NewDescribeImagesRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DescribeImages require credential")
    }

    request.SetContext(ctx)
    
    response = NewDescribeImagesResponse()
    err = c.Send(request, response)
    return
}

func NewDetectChefDressRequest() (request *DetectChefDressRequest) {
    request = &DetectChefDressRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiia", APIVersion, "DetectChefDress")
    
    
    return
}

func NewDetectChefDressResponse() (response *DetectChefDressResponse) {
    response = &DetectChefDressResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectChefDress
// 可对图片中厨师穿戴进行识别，支持厨师服识别，厨师帽识别，赤膊识别和口罩识别,可应用于明厨亮灶场景。
//
// "被优选过滤"标签值在人体优选开关开启时才会返回。
//
// 厨师服：厨师服定义为白色上衣
//
// 厨师服识别(酒店版)：厨师服定义为红色，白色，黑色上衣
//
// 
//
// |序号 | 标签名称 | 标签值 |
//
// | :-----|  :----------   |:-----------------  |
//
// | 1 | 厨师服识别<div style="width: 70pt"> |无厨师服、有厨师服、被优选过滤|
//
// | 2 | 厨师服识别（酒店版）<div style="width: 70pt"> |无厨师服、有厨师服、被优选过滤|
//
// | 3 | 厨师帽识别<div style="width: 70pt"> |无厨师帽、有厨师帽、被优选过滤	|
//
// | 4 | 赤膊识别<div style="width: 70pt"> |非赤膊、赤膊、被优选过滤|
//
// | 5 | 口罩识别<div style="width: 70pt"> |无口罩、有口罩、口罩不确定、被优选过滤	|
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOBODYINPHOTO = "FailedOperation.NoBodyInPhoto"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectChefDress(request *DetectChefDressRequest) (response *DetectChefDressResponse, err error) {
    return c.DetectChefDressWithContext(context.Background(), request)
}

// DetectChefDress
// 可对图片中厨师穿戴进行识别，支持厨师服识别，厨师帽识别，赤膊识别和口罩识别,可应用于明厨亮灶场景。
//
// "被优选过滤"标签值在人体优选开关开启时才会返回。
//
// 厨师服：厨师服定义为白色上衣
//
// 厨师服识别(酒店版)：厨师服定义为红色，白色，黑色上衣
//
// 
//
// |序号 | 标签名称 | 标签值 |
//
// | :-----|  :----------   |:-----------------  |
//
// | 1 | 厨师服识别<div style="width: 70pt"> |无厨师服、有厨师服、被优选过滤|
//
// | 2 | 厨师服识别（酒店版）<div style="width: 70pt"> |无厨师服、有厨师服、被优选过滤|
//
// | 3 | 厨师帽识别<div style="width: 70pt"> |无厨师帽、有厨师帽、被优选过滤	|
//
// | 4 | 赤膊识别<div style="width: 70pt"> |非赤膊、赤膊、被优选过滤|
//
// | 5 | 口罩识别<div style="width: 70pt"> |无口罩、有口罩、口罩不确定、被优选过滤	|
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOBODYINPHOTO = "FailedOperation.NoBodyInPhoto"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectChefDressWithContext(ctx context.Context, request *DetectChefDressRequest) (response *DetectChefDressResponse, err error) {
    if request == nil {
        request = NewDetectChefDressRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectChefDress require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectChefDressResponse()
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
func (c *Client) DetectDisgust(request *DetectDisgustRequest) (response *DetectDisgustResponse, err error) {
    return c.DetectDisgustWithContext(context.Background(), request)
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
func (c *Client) DetectDisgustWithContext(ctx context.Context, request *DetectDisgustRequest) (response *DetectDisgustResponse, err error) {
    if request == nil {
        request = NewDetectDisgustRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectDisgust require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectDisgustResponse()
    err = c.Send(request, response)
    return
}

func NewDetectEnvelopeRequest() (request *DetectEnvelopeRequest) {
    request = &DetectEnvelopeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiia", APIVersion, "DetectEnvelope")
    
    
    return
}

func NewDetectEnvelopeResponse() (response *DetectEnvelopeResponse) {
    response = &DetectEnvelopeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectEnvelope
// 文件封识别可检测图片中是否包含符合文件封（即文件、单据、资料等的袋状包装）特征的物品，覆盖顺丰快递文件封、文件袋、档案袋等多种文件封类型，可应用于物流行业对文件快递的包装审核等场景。
//
// 
//
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectEnvelope(request *DetectEnvelopeRequest) (response *DetectEnvelopeResponse, err error) {
    return c.DetectEnvelopeWithContext(context.Background(), request)
}

// DetectEnvelope
// 文件封识别可检测图片中是否包含符合文件封（即文件、单据、资料等的袋状包装）特征的物品，覆盖顺丰快递文件封、文件袋、档案袋等多种文件封类型，可应用于物流行业对文件快递的包装审核等场景。
//
// 
//
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectEnvelopeWithContext(ctx context.Context, request *DetectEnvelopeRequest) (response *DetectEnvelopeResponse, err error) {
    if request == nil {
        request = NewDetectEnvelopeRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectEnvelope require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectEnvelopeResponse()
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
// 图像标签利用深度学习技术，可以对图片进行智能分类、物体识别等。
//
// 
//
// 目前支持八个大类、六十多个子类、数千个标签，涵盖各种日常场景、动植物、物品、美食等。
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
// 为了方便使用、减少图片传输次数，图像标签将不同版本包装成多合一接口，实际上是多个服务，分别计费。建议在接入初期，对四个版本进行对比评估后选择合适的版本使用。
//
// 
//
// >?
//
// - 图像标签已升级服务，建议使用新版接口[通用图像标签](https://cloud.tencent.com/document/product/865/75196)。
//
// - 图像标签摄像头版、相册版、网络版、新闻版分别按照各自的实际使用次数进行收费，例如一张图片同时使用相册版、摄像头版，则按照两次调用计费。建议测试对比后从中选择一个最合适的版本使用即可。
//
// 
//
// >   
//
// - 可前往 [图像标签](https://cloud.tencent.com/document/product/1588) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
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
    return c.DetectLabelWithContext(context.Background(), request)
}

// DetectLabel
// 图像标签利用深度学习技术，可以对图片进行智能分类、物体识别等。
//
// 
//
// 目前支持八个大类、六十多个子类、数千个标签，涵盖各种日常场景、动植物、物品、美食等。
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
// 为了方便使用、减少图片传输次数，图像标签将不同版本包装成多合一接口，实际上是多个服务，分别计费。建议在接入初期，对四个版本进行对比评估后选择合适的版本使用。
//
// 
//
// >?
//
// - 图像标签已升级服务，建议使用新版接口[通用图像标签](https://cloud.tencent.com/document/product/865/75196)。
//
// - 图像标签摄像头版、相册版、网络版、新闻版分别按照各自的实际使用次数进行收费，例如一张图片同时使用相册版、摄像头版，则按照两次调用计费。建议测试对比后从中选择一个最合适的版本使用即可。
//
// 
//
// >   
//
// - 可前往 [图像标签](https://cloud.tencent.com/document/product/1588) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectLabel require credential")
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectLabelBeta(request *DetectLabelBetaRequest) (response *DetectLabelBetaResponse, err error) {
    return c.DetectLabelBetaWithContext(context.Background(), request)
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectLabelBeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectLabelBetaResponse()
    err = c.Send(request, response)
    return
}

func NewDetectLabelProRequest() (request *DetectLabelProRequest) {
    request = &DetectLabelProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiia", APIVersion, "DetectLabelPro")
    
    
    return
}

func NewDetectLabelProResponse() (response *DetectLabelProResponse) {
    response = &DetectLabelProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectLabelPro
// 通用图像标签可识别数千种常见物体或场景，覆盖日常物品、场景、动物、植物、食物、饮品、交通工具等多个大类，返回主体的标签名称和所属细分类目等。
//
// 
//
// >   
//
// - 可前往 [图像标签](https://cloud.tencent.com/document/product/1588) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
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
func (c *Client) DetectLabelPro(request *DetectLabelProRequest) (response *DetectLabelProResponse, err error) {
    return c.DetectLabelProWithContext(context.Background(), request)
}

// DetectLabelPro
// 通用图像标签可识别数千种常见物体或场景，覆盖日常物品、场景、动物、植物、食物、饮品、交通工具等多个大类，返回主体的标签名称和所属细分类目等。
//
// 
//
// >   
//
// - 可前往 [图像标签](https://cloud.tencent.com/document/product/1588) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  AUTHFAILURE_INVALIDAUTHORIZATION = "AuthFailure.InvalidAuthorization"
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
func (c *Client) DetectLabelProWithContext(ctx context.Context, request *DetectLabelProRequest) (response *DetectLabelProResponse, err error) {
    if request == nil {
        request = NewDetectLabelProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectLabelPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectLabelProResponse()
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectMisbehavior(request *DetectMisbehaviorRequest) (response *DetectMisbehaviorResponse, err error) {
    return c.DetectMisbehaviorWithContext(context.Background(), request)
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
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectMisbehaviorWithContext(ctx context.Context, request *DetectMisbehaviorRequest) (response *DetectMisbehaviorResponse, err error) {
    if request == nil {
        request = NewDetectMisbehaviorRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectMisbehavior require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectMisbehaviorResponse()
    err = c.Send(request, response)
    return
}

func NewDetectPetRequest() (request *DetectPetRequest) {
    request = &DetectPetRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiia", APIVersion, "DetectPet")
    
    
    return
}

func NewDetectPetResponse() (response *DetectPetResponse) {
    response = &DetectPetResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectPet
// 传入一张图片，识别出图片中是否存在宠物
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_NOOBJECTDETECTED = "FailedOperation.NoObjectDetected"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectPet(request *DetectPetRequest) (response *DetectPetResponse, err error) {
    return c.DetectPetWithContext(context.Background(), request)
}

// DetectPet
// 传入一张图片，识别出图片中是否存在宠物
//
// >     
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_INVOKECHARGEERROR = "FailedOperation.InvokeChargeError"
//  FAILEDOPERATION_NOOBJECTDETECTED = "FailedOperation.NoObjectDetected"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) DetectPetWithContext(ctx context.Context, request *DetectPetRequest) (response *DetectPetResponse, err error) {
    if request == nil {
        request = NewDetectPetRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectPet require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectPetResponse()
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
// >?    
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
    return c.DetectProductWithContext(context.Background(), request)
}

// DetectProduct
// 本接口支持识别图片中包含的商品，能够输出商品的品类名称、类别，还可以输出商品在图片中的位置。支持一张图片多个商品的识别。
//
// >?    
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectProduct require credential")
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
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectProductBeta(request *DetectProductBetaRequest) (response *DetectProductBetaResponse, err error) {
    return c.DetectProductBetaWithContext(context.Background(), request)
}

// DetectProductBeta
// 商品识别-微信识物版，基于人工智能技术、海量训练图片、亿级商品库，可以实现全覆盖、细粒度、高准确率的商品识别和商品推荐功能。
//
// 本服务可以识别出图片中的主体位置、主体商品类型，覆盖亿级SKU，输出具体商品的价格、型号等详细信息。
//
// 客户无需自建商品库，即可快速实现商品识别、拍照搜商品等功能。
//
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEUNQUALIFIED = "FailedOperation.ImageUnQualified"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNOPENERROR = "FailedOperation.UnOpenError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectProductBetaWithContext(ctx context.Context, request *DetectProductBetaRequest) (response *DetectProductBetaResponse, err error) {
    if request == nil {
        request = NewDetectProductBetaRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectProductBeta require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectProductBetaResponse()
    err = c.Send(request, response)
    return
}

func NewDetectSecurityRequest() (request *DetectSecurityRequest) {
    request = &DetectSecurityRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiia", APIVersion, "DetectSecurity")
    
    
    return
}

func NewDetectSecurityResponse() (response *DetectSecurityResponse) {
    response = &DetectSecurityResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// DetectSecurity
// 安全属性识别可对图片中人体安全防护属性进行识别，支持识别安全帽，反光衣，护目镜，工服，手套，工地安全带，口罩，抽烟，玩手机等多种属性。
//
// "被优选过滤"标签值在人体优选开关开启时才会返回。
//
// 
//
// |序号 | 标签名称 | 标签值 |
//
// | :-----|  :----------   |:-----------------  |
//
// | 1 | 安全帽识别<div style="width: 70pt"> |无安全帽、有安全帽、被优选过滤|
//
// | 2 | 玩手机识别<div style="width: 70pt"> |没有电话、打电话、玩手机、被优选过滤|
//
// | 3 | 抽烟识别<div style="width: 70pt"> |没有抽烟、抽烟、被优选过滤	|
//
// | 4 | 口罩识别<div style="width: 70pt"> |无口罩、有口罩、口罩不确定、被优选过滤|
//
// | 5 | 工地安全带识别<div style="width: 70pt"> |无工地安全带、工地安全带、被优选过滤	|
//
// | 6 | 手套识别<div style="width: 70pt"> |无手套、有手套、手套不确定、被优选过滤	|
//
// | 7 | 工服识别<div style="width: 70pt"> |无工服、有工服、被优选过滤|
//
// | 8 | 护目镜识别<div style="width: 70pt"> |无护目镜、有护目镜、被优选过滤|
//
// | 9 | 反光衣识别<div style="width: 70pt"> |无反光衣、有反光衣、被优选过滤|
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOBODYINPHOTO = "FailedOperation.NoBodyInPhoto"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectSecurity(request *DetectSecurityRequest) (response *DetectSecurityResponse, err error) {
    return c.DetectSecurityWithContext(context.Background(), request)
}

// DetectSecurity
// 安全属性识别可对图片中人体安全防护属性进行识别，支持识别安全帽，反光衣，护目镜，工服，手套，工地安全带，口罩，抽烟，玩手机等多种属性。
//
// "被优选过滤"标签值在人体优选开关开启时才会返回。
//
// 
//
// |序号 | 标签名称 | 标签值 |
//
// | :-----|  :----------   |:-----------------  |
//
// | 1 | 安全帽识别<div style="width: 70pt"> |无安全帽、有安全帽、被优选过滤|
//
// | 2 | 玩手机识别<div style="width: 70pt"> |没有电话、打电话、玩手机、被优选过滤|
//
// | 3 | 抽烟识别<div style="width: 70pt"> |没有抽烟、抽烟、被优选过滤	|
//
// | 4 | 口罩识别<div style="width: 70pt"> |无口罩、有口罩、口罩不确定、被优选过滤|
//
// | 5 | 工地安全带识别<div style="width: 70pt"> |无工地安全带、工地安全带、被优选过滤	|
//
// | 6 | 手套识别<div style="width: 70pt"> |无手套、有手套、手套不确定、被优选过滤	|
//
// | 7 | 工服识别<div style="width: 70pt"> |无工服、有工服、被优选过滤|
//
// | 8 | 护目镜识别<div style="width: 70pt"> |无护目镜、有护目镜、被优选过滤|
//
// | 9 | 反光衣识别<div style="width: 70pt"> |无反光衣、有反光衣、被优选过滤|
//
// 可能返回的错误码:
//  FAILEDOPERATION_DOWNLOADERROR = "FailedOperation.DownLoadError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_NOBODYINPHOTO = "FailedOperation.NoBodyInPhoto"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_TOOLARGEFILEERROR = "FailedOperation.TooLargeFileError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_IMAGEEMPTY = "InvalidParameterValue.ImageEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) DetectSecurityWithContext(ctx context.Context, request *DetectSecurityRequest) (response *DetectSecurityResponse, err error) {
    if request == nil {
        request = NewDetectSecurityRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("DetectSecurity require credential")
    }

    request.SetContext(ctx)
    
    response = NewDetectSecurityResponse()
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
// 
//
// >   
//
// - 可前往 [图像处理](https://cloud.tencent.com/document/product/1590) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnhanceImage(request *EnhanceImageRequest) (response *EnhanceImageResponse, err error) {
    return c.EnhanceImageWithContext(context.Background(), request)
}

// EnhanceImage
// 传入一张图片，输出清晰度提升后的图片。
//
// 
//
// 可以消除图片有损压缩导致的噪声，和使用滤镜、拍摄失焦导致的模糊。让图片的边缘和细节更加清晰自然。
//
// 
//
// >   
//
// - 可前往 [图像处理](https://cloud.tencent.com/document/product/1590) 产品文档中查看更多产品信息。
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
//  RESOURCEUNAVAILABLE_ISOPENING = "ResourceUnavailable.IsOpening"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) EnhanceImageWithContext(ctx context.Context, request *EnhanceImageRequest) (response *EnhanceImageResponse, err error) {
    if request == nil {
        request = NewEnhanceImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("EnhanceImage require credential")
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
// 车辆识别可对图片中车辆的车型进行识别，输出车辆的品牌（如路虎）、车系（如神行者2）、类型（如中型SUV）、颜色和车辆在图中的坐标等信息，覆盖轿车、SUV、大型客车等市面常见车，支持三千多种车辆型号。如果图片中存在多辆车，会分别输出每辆车的车型和坐标。
//
// 
//
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeCar(request *RecognizeCarRequest) (response *RecognizeCarResponse, err error) {
    return c.RecognizeCarWithContext(context.Background(), request)
}

// RecognizeCar
// 车辆识别可对图片中车辆的车型进行识别，输出车辆的品牌（如路虎）、车系（如神行者2）、类型（如中型SUV）、颜色和车辆在图中的坐标等信息，覆盖轿车、SUV、大型客车等市面常见车，支持三千多种车辆型号。如果图片中存在多辆车，会分别输出每辆车的车型和坐标。
//
// 
//
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
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
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeCar require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeCarResponse()
    err = c.Send(request, response)
    return
}

func NewRecognizeCarProRequest() (request *RecognizeCarProRequest) {
    request = &RecognizeCarProRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiia", APIVersion, "RecognizeCarPro")
    
    
    return
}

func NewRecognizeCarProResponse() (response *RecognizeCarProResponse) {
    response = &RecognizeCarProResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// RecognizeCarPro
// 车辆识别（增强版）在车辆识别的基础上**增加了车牌识别的功能，并升级了车型识别的效果**。可对图片中车辆的车型和车牌进行同时识别，输出车辆的车牌信息，以及车辆品牌（如路虎）、车系（如神行者2）、类型（如中型SUV）、颜色和车辆在图中的坐标等信息，覆盖轿车、SUV、大型客车等市面常见车，支持三千多种车辆型号。如果图片中存在多辆车，会分别输出每辆车的车型、车牌和坐标。
//
// 
//
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeCarPro(request *RecognizeCarProRequest) (response *RecognizeCarProResponse, err error) {
    return c.RecognizeCarProWithContext(context.Background(), request)
}

// RecognizeCarPro
// 车辆识别（增强版）在车辆识别的基础上**增加了车牌识别的功能，并升级了车型识别的效果**。可对图片中车辆的车型和车牌进行同时识别，输出车辆的车牌信息，以及车辆品牌（如路虎）、车系（如神行者2）、类型（如中型SUV）、颜色和车辆在图中的坐标等信息，覆盖轿车、SUV、大型客车等市面常见车，支持三千多种车辆型号。如果图片中存在多辆车，会分别输出每辆车的车型、车牌和坐标。
//
// 
//
// >?   
//
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
//
// 可能返回的错误码:
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
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETERVALUELIMIT = "InvalidParameterValue.InvalidParameterValueLimit"
//  INVALIDPARAMETERVALUE_URLILLEGAL = "InvalidParameterValue.UrlIllegal"
//  LIMITEXCEEDED_TOOLARGEFILEERROR = "LimitExceeded.TooLargeFileError"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
//  RESOURCESSOLDOUT_CHARGESTATUSEXCEPTION = "ResourcesSoldOut.ChargeStatusException"
func (c *Client) RecognizeCarProWithContext(ctx context.Context, request *RecognizeCarProRequest) (response *RecognizeCarProResponse, err error) {
    if request == nil {
        request = NewRecognizeCarProRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("RecognizeCarPro require credential")
    }

    request.SetContext(ctx)
    
    response = NewRecognizeCarProResponse()
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
// 本接口用于对一张图片，在指定图片库中检索出与之相似的图片列表。
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEGROUPCHARGESTATUSCLOSE = "FailedOperation.ImageGroupChargeStatusClose"
//  FAILEDOPERATION_IMAGEGROUPEMPTY = "FailedOperation.ImageGroupEmpty"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESEARCHINVALID = "FailedOperation.ImageSearchInvalid"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_IMAGEFORMATNOTSUPPORT = "InvalidParameter.ImageFormatNotSupport"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_FILTERINVALID = "InvalidParameterValue.FilterInvalid"
//  INVALIDPARAMETERVALUE_FILTERSIZEEXCEED = "InvalidParameterValue.FilterSizeExceed"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SearchImage(request *SearchImageRequest) (response *SearchImageResponse, err error) {
    return c.SearchImageWithContext(context.Background(), request)
}

// SearchImage
// 本接口用于对一张图片，在指定图片库中检索出与之相似的图片列表。
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_BALANCEINSUFFICIENT = "FailedOperation.BalanceInsufficient"
//  FAILEDOPERATION_EMPTYIMAGEERROR = "FailedOperation.EmptyImageError"
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEGROUPCHARGESTATUSCLOSE = "FailedOperation.ImageGroupChargeStatusClose"
//  FAILEDOPERATION_IMAGEGROUPEMPTY = "FailedOperation.ImageGroupEmpty"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGERESOLUTIONEXCEED = "FailedOperation.ImageResolutionExceed"
//  FAILEDOPERATION_IMAGESEARCHINVALID = "FailedOperation.ImageSearchInvalid"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_INNERERROR = "FailedOperation.InnerError"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  FAILEDOPERATION_UNKNOWERROR = "FailedOperation.UnKnowError"
//  FAILEDOPERATION_UNKNOWN = "FailedOperation.Unknown"
//  INVALIDPARAMETER_IMAGEFORMATNOTSUPPORT = "InvalidParameter.ImageFormatNotSupport"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ENTITYIDTOOLONG = "InvalidParameterValue.EntityIdTooLong"
//  INVALIDPARAMETERVALUE_FILTERINVALID = "InvalidParameterValue.FilterInvalid"
//  INVALIDPARAMETERVALUE_FILTERSIZEEXCEED = "InvalidParameterValue.FilterSizeExceed"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_LIMITEXCEED = "InvalidParameterValue.LimitExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) SearchImageWithContext(ctx context.Context, request *SearchImageRequest) (response *SearchImageResponse, err error) {
    if request == nil {
        request = NewSearchImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("SearchImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewSearchImageResponse()
    err = c.Send(request, response)
    return
}

func NewUpdateImageRequest() (request *UpdateImageRequest) {
    request = &UpdateImageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("tiia", APIVersion, "UpdateImage")
    
    
    return
}

func NewUpdateImageResponse() (response *UpdateImageResponse) {
    response = &UpdateImageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// UpdateImage
// 本接口支持根据图库ID、物品ID、图片名称来修改图片信息（暂仅支持修改Tags）
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEGROUPEMPTY = "FailedOperation.ImageGroupEmpty"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER_IMAGEFORMATNOTSUPPORT = "InvalidParameter.ImageFormatNotSupport"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMEEMPTY = "InvalidParameterValue.PicNameEmpty"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
//  INVALIDPARAMETERVALUE_TAGSKEYSEXCEED = "InvalidParameterValue.TagsKeysExceed"
//  INVALIDPARAMETERVALUE_TAGSVALUEILLEGAL = "InvalidParameterValue.TagsValueIllegal"
//  INVALIDPARAMETERVALUE_TAGSVALUESIZEEXCEED = "InvalidParameterValue.TagsValueSizeExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) UpdateImage(request *UpdateImageRequest) (response *UpdateImageResponse, err error) {
    return c.UpdateImageWithContext(context.Background(), request)
}

// UpdateImage
// 本接口支持根据图库ID、物品ID、图片名称来修改图片信息（暂仅支持修改Tags）
//
// 
//
// >   
//
// - 可前往 [图像搜索](https://cloud.tencent.com/document/product/1589) 产品文档中查看更多产品信息。
//
// 
//
// 可能返回的错误码:
//  FAILEDOPERATION_IMAGEDECODEFAILED = "FailedOperation.ImageDecodeFailed"
//  FAILEDOPERATION_IMAGEGROUPEMPTY = "FailedOperation.ImageGroupEmpty"
//  FAILEDOPERATION_IMAGENOTFOUNDINFO = "FailedOperation.ImageNotFoundInfo"
//  FAILEDOPERATION_IMAGENOTSUPPORTED = "FailedOperation.ImageNotSupported"
//  FAILEDOPERATION_IMAGESIZEEXCEED = "FailedOperation.ImageSizeExceed"
//  FAILEDOPERATION_IMAGEURLINVALID = "FailedOperation.ImageUrlInvalid"
//  FAILEDOPERATION_REQUESTERROR = "FailedOperation.RequestError"
//  FAILEDOPERATION_REQUESTTIMEOUT = "FailedOperation.RequestTimeout"
//  FAILEDOPERATION_RPCFAIL = "FailedOperation.RpcFail"
//  FAILEDOPERATION_SERVERERROR = "FailedOperation.ServerError"
//  INVALIDPARAMETER_IMAGEFORMATNOTSUPPORT = "InvalidParameter.ImageFormatNotSupport"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE_ENTITYIDEMPTY = "InvalidParameterValue.EntityIdEmpty"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDILLEGAL = "InvalidParameterValue.ImageGroupIdIllegal"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDNOTEXIST = "InvalidParameterValue.ImageGroupIdNotExist"
//  INVALIDPARAMETERVALUE_IMAGEGROUPIDTOOLONG = "InvalidParameterValue.ImageGroupIdTooLong"
//  INVALIDPARAMETERVALUE_PICNAMEEMPTY = "InvalidParameterValue.PicNameEmpty"
//  INVALIDPARAMETERVALUE_PICNAMETOOLONG = "InvalidParameterValue.PicNameTooLong"
//  INVALIDPARAMETERVALUE_TAGSKEYSEXCEED = "InvalidParameterValue.TagsKeysExceed"
//  INVALIDPARAMETERVALUE_TAGSVALUEILLEGAL = "InvalidParameterValue.TagsValueIllegal"
//  INVALIDPARAMETERVALUE_TAGSVALUESIZEEXCEED = "InvalidParameterValue.TagsValueSizeExceed"
//  MISSINGPARAMETER_ERRORPARAMETEREMPTY = "MissingParameter.ErrorParameterEmpty"
//  RESOURCEUNAVAILABLE_INARREARS = "ResourceUnavailable.InArrears"
//  RESOURCEUNAVAILABLE_NOTEXIST = "ResourceUnavailable.NotExist"
func (c *Client) UpdateImageWithContext(ctx context.Context, request *UpdateImageRequest) (response *UpdateImageResponse, err error) {
    if request == nil {
        request = NewUpdateImageRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("UpdateImage require credential")
    }

    request.SetContext(ctx)
    
    response = NewUpdateImageResponse()
    err = c.Send(request, response)
    return
}
