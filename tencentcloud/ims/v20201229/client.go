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

package v20201229

import (
    "context"
    "errors"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-12-29"

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


func NewCreateImageModerationAsyncTaskRequest() (request *CreateImageModerationAsyncTaskRequest) {
    request = &CreateImageModerationAsyncTaskRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ims", APIVersion, "CreateImageModerationAsyncTask")
    
    
    return
}

func NewCreateImageModerationAsyncTaskResponse() (response *CreateImageModerationAsyncTaskResponse) {
    response = &CreateImageModerationAsyncTaskResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// CreateImageModerationAsyncTask
// 本接口用于提交图片文件进行异步智能审核任务。使用前请您使用腾讯云主账号登录控制台 [开通图片内容安全服务](https://console.cloud.tencent.com/cms/image/package) 并调整好对应的业务配置。
//
// ### 接口使用说明：
//
// - 前往“[内容安全控制台-图片内容安全](https://console.cloud.tencent.com/cms/image/package)”开启使用图片内容安全服务，首次开通服务的用户可免费领用试用套餐包，包含**3000张图片**识别额度，有效期为15天。
//
// - 该接口为收费接口，计费方式敬请参见 [腾讯云图片内容安全定价](https://cloud.tencent.com/product/ims/pricing)。
//
// 
//
// ### 接口功能说明：
//
// - 支持对图片文件或链接进行检测，通过深度学习技术，识别可能令人反感、不安全或不适宜的违规图片内容；
//
// - 支持对长图进行拆分检测；
//
// - 支持识别多种违规场景，包括：低俗、违法违规、色情、广告等场景；
//
// - 支持多种物体检测（实体、广告台标、二维码等）及图片中文本的OCR文本识别；
//
// - 支持根据不同的业务场景配置自定义的审核策略；
//
// - 支持用户自定义选择图片风险库，打击自定义识别类型的违规图片（目前仅支持黑名单配置）；
//
// - 支持在审核图片内容时同时关联账号或设备信息，可识别违规风险账号或设备。
//
// ### 接口调用说明：
//
// - 图片文件大小限制：**Base64编码内容应小于10MB**；**图片资源应小于100MB**
//
// - 图片尺寸支持：**长和宽 需>50分辨率且<40000分辨率，并且图片长宽比<90:1；**
//
// - 图片文件分辨率支持：建议**分辨率大于256x256**，否则可能会影响识别效果；
//
// - 图片文件支持格式：PNG、JPG、JPEG、BMP、GIF、WEBP格式；
//
// - 图片文件链接支持的传输协议：HTTP、HTTPS；
//
// - 若传入图片文件的访问链接，则需要注意**图片异步接口下载超时时间为15秒，失败后重试1次为3 秒，共18秒下载时间**，为保障被检测图片的稳定性和可靠性，建议您使用腾讯云COS存储或者CDN缓存等；
//
// - 默认接口请求频率限制：**20次/秒**，超过此调用频率则会报错。
//
// - ### 接口回调格式：
//
// - 图片异步检测接口回调格式同[图片同步检测接口输出参数](https://cloud.tencent.com/document/product/1125/53273#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0)
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYIMAGECONTENT = "InvalidParameterValue.EmptyImageContent"
//  INVALIDPARAMETERVALUE_IMAGESIZETOOSMALL = "InvalidParameterValue.ImageSizeTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDDATAID = "InvalidParameterValue.InvalidDataId"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
//  RESOURCEUNAVAILABLE_INVALIDIMAGECONTENT = "ResourceUnavailable.InvalidImageContent"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateImageModerationAsyncTask(request *CreateImageModerationAsyncTaskRequest) (response *CreateImageModerationAsyncTaskResponse, err error) {
    return c.CreateImageModerationAsyncTaskWithContext(context.Background(), request)
}

// CreateImageModerationAsyncTask
// 本接口用于提交图片文件进行异步智能审核任务。使用前请您使用腾讯云主账号登录控制台 [开通图片内容安全服务](https://console.cloud.tencent.com/cms/image/package) 并调整好对应的业务配置。
//
// ### 接口使用说明：
//
// - 前往“[内容安全控制台-图片内容安全](https://console.cloud.tencent.com/cms/image/package)”开启使用图片内容安全服务，首次开通服务的用户可免费领用试用套餐包，包含**3000张图片**识别额度，有效期为15天。
//
// - 该接口为收费接口，计费方式敬请参见 [腾讯云图片内容安全定价](https://cloud.tencent.com/product/ims/pricing)。
//
// 
//
// ### 接口功能说明：
//
// - 支持对图片文件或链接进行检测，通过深度学习技术，识别可能令人反感、不安全或不适宜的违规图片内容；
//
// - 支持对长图进行拆分检测；
//
// - 支持识别多种违规场景，包括：低俗、违法违规、色情、广告等场景；
//
// - 支持多种物体检测（实体、广告台标、二维码等）及图片中文本的OCR文本识别；
//
// - 支持根据不同的业务场景配置自定义的审核策略；
//
// - 支持用户自定义选择图片风险库，打击自定义识别类型的违规图片（目前仅支持黑名单配置）；
//
// - 支持在审核图片内容时同时关联账号或设备信息，可识别违规风险账号或设备。
//
// ### 接口调用说明：
//
// - 图片文件大小限制：**Base64编码内容应小于10MB**；**图片资源应小于100MB**
//
// - 图片尺寸支持：**长和宽 需>50分辨率且<40000分辨率，并且图片长宽比<90:1；**
//
// - 图片文件分辨率支持：建议**分辨率大于256x256**，否则可能会影响识别效果；
//
// - 图片文件支持格式：PNG、JPG、JPEG、BMP、GIF、WEBP格式；
//
// - 图片文件链接支持的传输协议：HTTP、HTTPS；
//
// - 若传入图片文件的访问链接，则需要注意**图片异步接口下载超时时间为15秒，失败后重试1次为3 秒，共18秒下载时间**，为保障被检测图片的稳定性和可靠性，建议您使用腾讯云COS存储或者CDN缓存等；
//
// - 默认接口请求频率限制：**20次/秒**，超过此调用频率则会报错。
//
// - ### 接口回调格式：
//
// - 图片异步检测接口回调格式同[图片同步检测接口输出参数](https://cloud.tencent.com/document/product/1125/53273#3.-.E8.BE.93.E5.87.BA.E5.8F.82.E6.95.B0)
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYIMAGECONTENT = "InvalidParameterValue.EmptyImageContent"
//  INVALIDPARAMETERVALUE_IMAGESIZETOOSMALL = "InvalidParameterValue.ImageSizeTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDCALLBACKURL = "InvalidParameterValue.InvalidCallbackUrl"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDDATAID = "InvalidParameterValue.InvalidDataId"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
//  RESOURCEUNAVAILABLE_INVALIDIMAGECONTENT = "ResourceUnavailable.InvalidImageContent"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateImageModerationAsyncTaskWithContext(ctx context.Context, request *CreateImageModerationAsyncTaskRequest) (response *CreateImageModerationAsyncTaskResponse, err error) {
    if request == nil {
        request = NewCreateImageModerationAsyncTaskRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("CreateImageModerationAsyncTask require credential")
    }

    request.SetContext(ctx)
    
    response = NewCreateImageModerationAsyncTaskResponse()
    err = c.Send(request, response)
    return
}

func NewImageModerationRequest() (request *ImageModerationRequest) {
    request = &ImageModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    
    request.Init().WithApiInfo("ims", APIVersion, "ImageModeration")
    
    
    return
}

func NewImageModerationResponse() (response *ImageModerationResponse) {
    response = &ImageModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    } 
    return

}

// ImageModeration
// 本接口（Image Moderation, IM）用于提交图片文件进行同步智能审核任务。使用前请您使用腾讯云主账号登录控制台 [开通图片内容安全服务](https://console.cloud.tencent.com/cms) 并调整好对应的业务配置。
//
// ### 接口使用说明：
//
// - 前往“[内容安全控制台-图片内容安全](https://console.cloud.tencent.com/cms)”开启使用图片内容安全服务，首次开通服务的用户可免费领用试用套餐包，包含**3000张图片**识别额度，有效期为15天。
//
// - 该接口为收费接口，计费方式敬请参见 [腾讯云图片内容安全定价](https://cloud.tencent.com/product/ims/pricing)。
//
// 
//
// ### 接口功能说明：
//
// - 支持对图片文件或链接进行检测，通过深度学习技术，识别可能令人反感、不安全或不适宜的违规图片内容；
//
// - 支持对GIF图/长图进行截帧或拆分检测；
//
// - 支持识别多种违规场景，包括：低俗、违法违规、色情、广告等场景；
//
// - 支持多种物体检测（实体、广告台标、二维码等）及图片中文本的OCR文本识别；
//
// - 支持根据不同的业务场景配置自定义的审核策略；
//
// - 支持用户自定义选择图片风险库，打击自定义识别类型的违规图片（目前仅支持黑名单配置）；
//
// - 支持在审核图片内容时同时关联账号或设备信息，可识别违规风险账号或设备。
//
// 
//
// ### 接口调用说明：
//
// - 图片文件大小限制：图片文件须大于256字节，Base64编码后的FileContent内容应小于10MB，FileURL的源图应小于30MB；
//
// - 图片默认尺寸支持：**长和宽 需>50分辨率且<10000分辨率，并且图片长宽比<90:1；**
//
// - 图片文件分辨率支持：建议**分辨率大于256x256**，否则可能会影响识别效果；
//
// - 图片文件支持格式（无备注默认只支持静态图）：BMP、ICO、JPEG、JNG、PNG、TIFF、RAW、SVG、GIF（静、动）、WEBP（静、动）、HEIC（静、动），动图在处理时，默认最多抽取5帧图像，并且每隔5帧进行一次采样；
//
// - 图片文件链接支持的传输协议：HTTP、HTTPS；
//
// - 图片审核仅支持单次传输一条URL进行审核，支持多并发请求。默认并发为100 QPS，超过此限制将返回`RequestLimitExceeded`错误；
//
// - 若传入图片文件的访问链接时，为保障被检测图片的稳定性和可靠性，建议您使用腾讯云COS存储或者CDN缓存等。
//
// 
//
// ### 图片下载说明：
//
// - 下载时间限制：图片首次下载限制为3秒。如超过此时间，将重试一次，重试时间为3秒。如果重试仍超时，将返回 `ImageDownloadError`（下载超时）；
//
// - 网络安全策略：由于网络安全策略，带有重定向的链接（如HTTP返回302状态码）可能导致下载失败。请尽量避免使用此类链接，以确保下载顺利进行，可能返回 `ResourceUnavailable.ImageDownloadError`；
//
// - 异步回源支持：目前不支持通过异步回源方式获取图片源。请确保使用直接链接，以避免下载失败。
//
// 
//
// <div class="rno-api-explorer" style="margin-bottom:20px">
//
//     <div class="rno-api-explorer-inner">
//
//         <div class="rno-api-explorer-hd">
//
//             <div class="rno-api-explorer-title">
//
//                 关于版本迭代的描述
//
//             </div>
//
//         </div>
//
//         <div class="rno-api-explorer-body">
//
//             <div class="rno-api-explorer-cont">
//
//                 <p>当前页面版本为图片内容安全2020版本，2020.11.3日前接入的图片内容安全接口为2019版本，在此时间前接入的用户可直接访问以下链接进行维护操作：<a href="https://cloud.tencent.com/document/product/1125/38206" target="_blank">图片内容安全-2019版本</a></p>
//
//                 <p>2020版本相对2019版本进行了升级，支持更灵活的多场景业务策略配置以及更丰富的识别回调信息，满足不同业务的识别需求，建议按照2020版本接入指引进行接口升级；同时，2019版本也会持续维护直至用户不再使用为止。</p>
//
//             </div>
//
//         </div>
//
//     </div>
//
// </div>
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEASPECTRATIOTOOLARGE = "InvalidParameter.ImageAspectRatioTooLarge"
//  INVALIDPARAMETER_IMAGEDATATOOSMALL = "InvalidParameter.ImageDataTooSmall"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYIMAGECONTENT = "InvalidParameterValue.EmptyImageContent"
//  INVALIDPARAMETERVALUE_IMAGESIZETOOSMALL = "InvalidParameterValue.ImageSizeTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDDATAID = "InvalidParameterValue.InvalidDataId"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
//  RESOURCEUNAVAILABLE_INVALIDIMAGECONTENT = "ResourceUnavailable.InvalidImageContent"
//  RESOURCEUNAVAILABLE_MODELCALLFAILED = "ResourceUnavailable.ModelCallFailed"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImageModeration(request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    return c.ImageModerationWithContext(context.Background(), request)
}

// ImageModeration
// 本接口（Image Moderation, IM）用于提交图片文件进行同步智能审核任务。使用前请您使用腾讯云主账号登录控制台 [开通图片内容安全服务](https://console.cloud.tencent.com/cms) 并调整好对应的业务配置。
//
// ### 接口使用说明：
//
// - 前往“[内容安全控制台-图片内容安全](https://console.cloud.tencent.com/cms)”开启使用图片内容安全服务，首次开通服务的用户可免费领用试用套餐包，包含**3000张图片**识别额度，有效期为15天。
//
// - 该接口为收费接口，计费方式敬请参见 [腾讯云图片内容安全定价](https://cloud.tencent.com/product/ims/pricing)。
//
// 
//
// ### 接口功能说明：
//
// - 支持对图片文件或链接进行检测，通过深度学习技术，识别可能令人反感、不安全或不适宜的违规图片内容；
//
// - 支持对GIF图/长图进行截帧或拆分检测；
//
// - 支持识别多种违规场景，包括：低俗、违法违规、色情、广告等场景；
//
// - 支持多种物体检测（实体、广告台标、二维码等）及图片中文本的OCR文本识别；
//
// - 支持根据不同的业务场景配置自定义的审核策略；
//
// - 支持用户自定义选择图片风险库，打击自定义识别类型的违规图片（目前仅支持黑名单配置）；
//
// - 支持在审核图片内容时同时关联账号或设备信息，可识别违规风险账号或设备。
//
// 
//
// ### 接口调用说明：
//
// - 图片文件大小限制：图片文件须大于256字节，Base64编码后的FileContent内容应小于10MB，FileURL的源图应小于30MB；
//
// - 图片默认尺寸支持：**长和宽 需>50分辨率且<10000分辨率，并且图片长宽比<90:1；**
//
// - 图片文件分辨率支持：建议**分辨率大于256x256**，否则可能会影响识别效果；
//
// - 图片文件支持格式（无备注默认只支持静态图）：BMP、ICO、JPEG、JNG、PNG、TIFF、RAW、SVG、GIF（静、动）、WEBP（静、动）、HEIC（静、动），动图在处理时，默认最多抽取5帧图像，并且每隔5帧进行一次采样；
//
// - 图片文件链接支持的传输协议：HTTP、HTTPS；
//
// - 图片审核仅支持单次传输一条URL进行审核，支持多并发请求。默认并发为100 QPS，超过此限制将返回`RequestLimitExceeded`错误；
//
// - 若传入图片文件的访问链接时，为保障被检测图片的稳定性和可靠性，建议您使用腾讯云COS存储或者CDN缓存等。
//
// 
//
// ### 图片下载说明：
//
// - 下载时间限制：图片首次下载限制为3秒。如超过此时间，将重试一次，重试时间为3秒。如果重试仍超时，将返回 `ImageDownloadError`（下载超时）；
//
// - 网络安全策略：由于网络安全策略，带有重定向的链接（如HTTP返回302状态码）可能导致下载失败。请尽量避免使用此类链接，以确保下载顺利进行，可能返回 `ResourceUnavailable.ImageDownloadError`；
//
// - 异步回源支持：目前不支持通过异步回源方式获取图片源。请确保使用直接链接，以避免下载失败。
//
// 
//
// <div class="rno-api-explorer" style="margin-bottom:20px">
//
//     <div class="rno-api-explorer-inner">
//
//         <div class="rno-api-explorer-hd">
//
//             <div class="rno-api-explorer-title">
//
//                 关于版本迭代的描述
//
//             </div>
//
//         </div>
//
//         <div class="rno-api-explorer-body">
//
//             <div class="rno-api-explorer-cont">
//
//                 <p>当前页面版本为图片内容安全2020版本，2020.11.3日前接入的图片内容安全接口为2019版本，在此时间前接入的用户可直接访问以下链接进行维护操作：<a href="https://cloud.tencent.com/document/product/1125/38206" target="_blank">图片内容安全-2019版本</a></p>
//
//                 <p>2020版本相对2019版本进行了升级，支持更灵活的多场景业务策略配置以及更丰富的识别回调信息，满足不同业务的识别需求，建议按照2020版本接入指引进行接口升级；同时，2019版本也会持续维护直至用户不再使用为止。</p>
//
//             </div>
//
//         </div>
//
//     </div>
//
// </div>
//
// 可能返回的错误码:
//  DRYRUNOPERATION = "DryRunOperation"
//  FAILEDOPERATION = "FailedOperation"
//  INTERNALERROR = "InternalError"
//  INTERNALERROR_INTERNALERROR = "InternalError.InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_IMAGEASPECTRATIOTOOLARGE = "InvalidParameter.ImageAspectRatioTooLarge"
//  INVALIDPARAMETER_IMAGEDATATOOSMALL = "InvalidParameter.ImageDataTooSmall"
//  INVALIDPARAMETER_IMAGESIZETOOSMALL = "InvalidParameter.ImageSizeTooSmall"
//  INVALIDPARAMETER_INVALIDIMAGECONTENT = "InvalidParameter.InvalidImageContent"
//  INVALIDPARAMETER_INVALIDPARAMETER = "InvalidParameter.InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_EMPTYIMAGECONTENT = "InvalidParameterValue.EmptyImageContent"
//  INVALIDPARAMETERVALUE_IMAGESIZETOOSMALL = "InvalidParameterValue.ImageSizeTooSmall"
//  INVALIDPARAMETERVALUE_INVALIDCONTENT = "InvalidParameterValue.InvalidContent"
//  INVALIDPARAMETERVALUE_INVALIDDATAID = "InvalidParameterValue.InvalidDataId"
//  INVALIDPARAMETERVALUE_INVALIDFILECONTENTSIZE = "InvalidParameterValue.InvalidFileContentSize"
//  INVALIDPARAMETERVALUE_INVALIDIMAGECONTENT = "InvalidParameterValue.InvalidImageContent"
//  INVALIDPARAMETERVALUE_INVALIDPARAMETER = "InvalidParameterValue.InvalidParameter"
//  LIMITEXCEEDED = "LimitExceeded"
//  MISSINGPARAMETER = "MissingParameter"
//  OPERATIONDENIED = "OperationDenied"
//  REQUESTLIMITEXCEEDED = "RequestLimitExceeded"
//  RESOURCEINUSE = "ResourceInUse"
//  RESOURCEINSUFFICIENT = "ResourceInsufficient"
//  RESOURCENOTFOUND = "ResourceNotFound"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
//  RESOURCEUNAVAILABLE_IMAGEDOWNLOADERROR = "ResourceUnavailable.ImageDownloadError"
//  RESOURCEUNAVAILABLE_INVALIDIMAGECONTENT = "ResourceUnavailable.InvalidImageContent"
//  RESOURCEUNAVAILABLE_MODELCALLFAILED = "ResourceUnavailable.ModelCallFailed"
//  RESOURCESSOLDOUT = "ResourcesSoldOut"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
//  UNAUTHORIZEDOPERATION_UNAUTHORIZED = "UnauthorizedOperation.Unauthorized"
//  UNKNOWNPARAMETER = "UnknownParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) ImageModerationWithContext(ctx context.Context, request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    if request == nil {
        request = NewImageModerationRequest()
    }
    
    if c.GetCredential() == nil {
        return nil, errors.New("ImageModeration require credential")
    }

    request.SetContext(ctx)
    
    response = NewImageModerationResponse()
    err = c.Send(request, response)
    return
}
