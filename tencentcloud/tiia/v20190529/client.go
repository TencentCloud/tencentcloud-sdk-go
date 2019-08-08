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
func (c *Client) AssessQuality(request *AssessQualityRequest) (response *AssessQualityResponse, err error) {
    if request == nil {
        request = NewAssessQualityRequest()
    }
    response = NewAssessQualityResponse()
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
func (c *Client) DetectCelebrity(request *DetectCelebrityRequest) (response *DetectCelebrityResponse, err error) {
    if request == nil {
        request = NewDetectCelebrityRequest()
    }
    response = NewDetectCelebrityResponse()
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

// 传入一张图片，识别出图片中存在的物体，并返回物体的名称（分类）、置信度，一张图片会给出多个可能的标签。
func (c *Client) DetectLabel(request *DetectLabelRequest) (response *DetectLabelResponse, err error) {
    if request == nil {
        request = NewDetectLabelRequest()
    }
    response = NewDetectLabelResponse()
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
// 
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

// 本接口提供多种维度的图像审核能力，支持色情和性感内容识别，政治人物和涉政敏感场景识别，以及暴恐人物、场景、旗帜标识等违禁内容的识别。
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

// 腾讯云车辆属性识别可对汽车车身及车辆属性进行检测与识别，目前支持11种车身颜色、20多种车型、300多种品牌、4000多种车系+年款的识别，同时支持对车标的位置进行检测。
func (c *Client) RecognizeCar(request *RecognizeCarRequest) (response *RecognizeCarResponse, err error) {
    if request == nil {
        request = NewRecognizeCarRequest()
    }
    response = NewRecognizeCarResponse()
    err = c.Send(request, response)
    return
}
