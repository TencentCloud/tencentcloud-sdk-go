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

package v20181119

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Coord struct {

	// 横坐标
	X *int64 `json:"X,omitempty" name:"X"`

	// 纵坐标
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type GeneralBasicOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的BASE64值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过3M。图片下载时间不超过3秒。
	// 图片的 ImageUrl、ImageBase64必须提供一个，如果都提供，只使用ImageBase64。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的URL地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过3M。图片下载时间不超过3秒。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 保留字段。
	Scene *string `json:"Scene,omitempty" name:"Scene"`
}

func (r *GeneralBasicOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralBasicOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 检测到的语言，目前支持的语种范围为：简体中文、繁体中文、英文、日文、韩文。未来将陆续新增对更多语种的支持。
	// 返回结果含义为：zh-中英混合，jap-日文，kor-韩文。
		Language *string `json:"Language,omitempty" name:"Language"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralBasicOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralBasicOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的BASE64值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过3M。图片下载时间不超过3秒。
	// 图片的 ImageUrl、ImageBase64必须提供一个，如果都提供，只使用ImageBase64。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片的URL地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过3M。图片下载时间不超过3秒。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`
}

func (r *GeneralFastOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralFastOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GeneralFastOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的文本信息，具体内容请点击左侧链接
		TextDetections []*TextDetection `json:"TextDetections,omitempty" name:"TextDetections" list`

		// 检测到的语言，目前支持的语种范围为：简体中文、繁体中文、英文、日文、韩文。未来将陆续新增对更多语种的支持。
	// 返回结果含义为：zh-中英混合，jap-日文，kor-韩文。
		Language *string `json:"Language,omitempty" name:"Language"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GeneralFastOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GeneralFastOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IDCardOCRRequest struct {
	*tchttp.BaseRequest

	// 图片的BASE64值。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过6M。图片下载时间不超过3秒。
	// 图片的 ImageUrl、ImageBase64必须提供一个，如果都提供，只使用ImageBase64。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 图片URL地址。
	// 支持的图片格式：PNG、JPG、JPEG，暂不支持GIF格式。
	// 支持的图片大小：所下载图片经Base64编码后不超过3M。图片下载时间不超过3秒。
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。非腾讯云存储的Url速度和稳定性可能受一定影响。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// FRONT为身份证有照片的一面（正面）
	// BACK为身份证有国徽的一面（反面）
	CardSide *string `json:"CardSide,omitempty" name:"CardSide"`

	// 可选字段，根据需要选择是否请求对应字段。
	// 目前包含的字段为：
	// CropIdCard-身份证照片裁剪，bool类型，
	// CropPortrait-人像照片裁剪，bool类型，
	// CopyWarn-复印件告警，bool类型，
	// ReshootWarn-翻拍告警，bool类型。
	// 
	// SDK设置方式参考：
	// Config = Json.stringify({"CropIdCard":true,"CropPortrait":true})
	// API 3.0 Explorer设置方式参考：
	// Config = {"CropIdCard":true,"CropPortrait":true}
	Config *string `json:"Config,omitempty" name:"Config"`
}

func (r *IDCardOCRRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IDCardOCRRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type IDCardOCRResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 姓名（正面）
		Name *string `json:"Name,omitempty" name:"Name"`

		// 性别（正面）
		Sex *string `json:"Sex,omitempty" name:"Sex"`

		// 民族（正面）
		Nation *string `json:"Nation,omitempty" name:"Nation"`

		// 出生日期（正面）
		Birth *string `json:"Birth,omitempty" name:"Birth"`

		// 地址（正面）
		Address *string `json:"Address,omitempty" name:"Address"`

		// 身份证号（正面）
		IdNum *string `json:"IdNum,omitempty" name:"IdNum"`

		// 发证机关（反面）
		Authority *string `json:"Authority,omitempty" name:"Authority"`

		// 证件有效期（反面）
		ValidDate *string `json:"ValidDate,omitempty" name:"ValidDate"`

		// 扩展信息，根据请求的可选字段返回对应内容，不请求则不返回，具体输入参考示例3。目前支持的扩展字段为：
	// IdCard身份证照片，请求CropIdCard时返回；
	// Portrait人像照片，请求CropPortrait时返回；
	// WarnInfos告警信息（Code告警码，Msg告警信息），识别出翻拍件或复印件时返回。
		AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IDCardOCRResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *IDCardOCRResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextDetection struct {

	// 识别出的文本行内容
	DetectedText *string `json:"DetectedText,omitempty" name:"DetectedText"`

	// 置信度 0 ~100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 文本行坐标，以四个顶点坐标表示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Polygon []*Coord `json:"Polygon,omitempty" name:"Polygon" list`

	// 此字段为扩展字段。
	// GeneralBasicOcr接口返回段落信息Parag，包含ParagNo。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}
