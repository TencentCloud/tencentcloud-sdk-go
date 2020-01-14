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

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
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

// 评估输入图片在视觉上的质量，从多个方面评估，并同时给出综合的、客观的清晰度评分，和主观的美观度评分。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) AssessQuality(request *AssessQualityRequest) (response *AssessQualityResponse, err error) {
    if request == nil {
        request = NewAssessQualityRequest()
    }
    response = NewAssessQualityResponse()
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

// 根据输入的裁剪比例，智能判断一张图片的最佳裁剪区域，确保原图的主体区域不受影响。
// 
// 可以自动裁剪图片，适应不同平台、设备的展示要求，避免简单拉伸带来的变形。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) CropImage(request *CropImageRequest) (response *CropImageResponse, err error) {
    if request == nil {
        request = NewCropImageRequest()
    }
    response = NewCropImageResponse()
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

// 传入一张图片，可以识别图片中包含的人物是否为公众人物，如果是，输出人物的姓名、基本信息、脸部坐标。
// 
// 支持识别一张图片中存在的多个人脸，针对每个人脸，会给出与之最相似的公众人物。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) DetectCelebrity(request *DetectCelebrityRequest) (response *DetectCelebrityResponse, err error) {
    if request == nil {
        request = NewDetectCelebrityRequest()
    }
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

// 输入一张图片，返回AI针对一张图片是否是恶心的一系列判断值。
// 
// 通过恶心图片识别, 可以判断一张图片是否令人恶心, 同时给出它属于的潜在类别, 让您能够过滤掉使人不愉快的图片。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) DetectDisgust(request *DetectDisgustRequest) (response *DetectDisgustResponse, err error) {
    if request == nil {
        request = NewDetectDisgustRequest()
    }
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

// 图像标签利用深度学习技术、海量训练数据，可以对图片进行智能分类、物体识别等。
// 
// 目前支持8个大类、六十多个子类、数千个标签。涵盖各种日常场景、动植物、物品、美食、卡证等。具体分类请见图像分析常见问题功能和限制部分（https://cloud.tencent.com/document/product/865/17629 ）
// 
// 图像标签提供三个版本供选择：
// 
// • 摄像头版：针对搜索、手机摄像头照片进行优化，涵盖大量卡证、日常物品、二维码条形码。
// 
// • 相册版：针对手机相册、网盘进行优化，去除相册和网盘中不常见的标签，针对相册常见图片类型（人像、日常活动、日常物品等）识别效果更好。
// 
// • 网络版：针对网络图片进行优化，涵盖标签更多，满足长尾识别需求。
// 
// 每个产品的图像类型都有独特性，建议在接入初期，对三个版本进行对比评估后选择合适的版本使用。
// 
// 为了方便使用、减少图片传输次数，图像标签包装成多合一接口，实际上是多个服务。
// 
// 图像标签按照服务的实际使用数量进行收费。例如一张图片同时调用相册版、摄像头版两个服务，那么此次调用按照两次计费。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) DetectLabel(request *DetectLabelRequest) (response *DetectLabelResponse, err error) {
    if request == nil {
        request = NewDetectLabelRequest()
    }
    response = NewDetectLabelResponse()
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

// 可以识别输入的图片中是否包含不良行为，例如打架斗殴、赌博、抽烟等，可以应用于广告图、直播截图、短视频截图等审核，减少不良行为对平台内容质量的影响，维护健康向上的互联网环境。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) DetectMisbehavior(request *DetectMisbehaviorRequest) (response *DetectMisbehaviorResponse, err error) {
    if request == nil {
        request = NewDetectMisbehaviorRequest()
    }
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

// 本接口支持识别图片中包含的商品，能够输出商品的品类名称、类别，还可以输出商品在图片中的位置。支持一张图片多个商品的识别。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) DetectProduct(request *DetectProductRequest) (response *DetectProductResponse, err error) {
    if request == nil {
        request = NewDetectProductRequest()
    }
    response = NewDetectProductResponse()
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

// 传入一张图片，输出清晰度提升后的图片。
// 
// 可以消除图片有损压缩导致的噪声，和使用滤镜、拍摄失焦导致的模糊。让图片的边缘和细节更加清晰自然。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) EnhanceImage(request *EnhanceImageRequest) (response *EnhanceImageResponse, err error) {
    if request == nil {
        request = NewEnhanceImageRequest()
    }
    response = NewEnhanceImageResponse()
    err = c.Send(request, response)
    return
}

func NewImageModerationRequest() (request *ImageModerationRequest) {
    request = &ImageModerationRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tiia", APIVersion, "ImageModeration")
    return
}

func NewImageModerationResponse() (response *ImageModerationResponse) {
    response = &ImageModerationResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 图像审核包含鉴黄、政治敏感识别、暴恐识别、图文审核等服务。
// 
// • 鉴黄：识别图片中是否存在涉黄内容，将图片分为正常、性感、色情三类，并输出分类的置信度。
// 
// • 政治敏感识别：识别图片中是否存在政治人物、政治恶搞、著名政治事件照片等内容。
// 
// • 暴恐识别：识别图片中的暴力恐怖元素，例如武器、管制刀具、血腥、恐怖主义等。
// 
// • 图文审核：结合OCR、NLP和二维码识别算法，分析图片中的文字是否违规，或是否存在二维码等营销内容，并输出违规内容的类别（涉黄、涉政、谩骂、广告营销、暴恐等）。
// 
// 为了方便使用、减少图片传输次数，图像审核包装成多合一接口，实际上是多个服务。
// 
// 图像审核根据服务的调用次数收费。例如一张图片同时调用鉴黄、图文审核两个服务，那么此次调用按照两次计费。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) ImageModeration(request *ImageModerationRequest) (response *ImageModerationResponse, err error) {
    if request == nil {
        request = NewImageModerationRequest()
    }
    response = NewImageModerationResponse()
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

// 腾讯云车辆属性识别可对汽车车身及车辆属性进行检测与识别，目前支持11种车身颜色、20多种车型、300多种品牌、4000多种车系+年款的识别，同时支持对车辆的位置进行检测。如果图片中存在多辆车，会分别输出每辆车的车型和坐标。
// >     
// - 公共参数中的签名方式必须指定为V3版本，即配置SignatureMethod参数为TC3-HMAC-SHA256。
func (c *Client) RecognizeCar(request *RecognizeCarRequest) (response *RecognizeCarResponse, err error) {
    if request == nil {
        request = NewRecognizeCarRequest()
    }
    response = NewRecognizeCarResponse()
    err = c.Send(request, response)
    return
}
