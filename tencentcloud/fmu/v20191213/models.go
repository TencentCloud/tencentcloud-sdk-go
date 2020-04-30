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

package v20191213

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type BeautifyPicRequest struct {
	*tchttp.BaseRequest

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url 。对应图片 base64 编码后大小不可超过5M。 
	// Url、Image必须提供一个，如果都提供，只使用 Url。  
	// 图片存储于腾讯云的Url可保障更高下载速度和稳定性，建议图片存储于腾讯云。  
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 美白程度，取值范围[0,100]。0不美白，100代表最高程度。默认值30。
	Whitening *uint64 `json:"Whitening,omitempty" name:"Whitening"`

	// 磨皮程度，取值范围[0,100]。0不磨皮，100代表最高程度。默认值10。
	Smoothing *uint64 `json:"Smoothing,omitempty" name:"Smoothing"`

	// 瘦脸程度，取值范围[0,100]。0不瘦脸，100代表最高程度。默认值70。
	FaceLifting *uint64 `json:"FaceLifting,omitempty" name:"FaceLifting"`

	// 大眼程度，取值范围[0,100]。0不大眼，100代表最高程度。默认值70。
	EyeEnlarging *uint64 `json:"EyeEnlarging,omitempty" name:"EyeEnlarging"`
}

func (r *BeautifyPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BeautifyPicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type BeautifyPicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 处理后的图片 base64 数据。
		ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BeautifyPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *BeautifyPicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModelRequest struct {
	*tchttp.BaseRequest

	// 用于试唇色，要求必须是LUT 格式的cube文件转换成512*512的PNG图片。查看 [LUT文件的使用说明](https://cloud.tencent.com/document/product/1172/41701)。了解 [cube文件转png图片小工具](http://yyb.gtimg.com/aiplat/static/qcloud-cube-to-png.html)。
	LUTFile *string `json:"LUTFile,omitempty" name:"LUTFile"`

	// 文件描述信息，可用于备注。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateModelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唇色素材ID。
		ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateModelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModelRequest struct {
	*tchttp.BaseRequest

	// 素材ID。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *DeleteModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteModelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteModelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FaceRect struct {

	// 人脸框左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸框左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸框宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸框高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type GetModelListRequest struct {
	*tchttp.BaseRequest

	// 起始序号，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认值为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *GetModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModelListRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetModelListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唇色素材总数量。
		ModelIdNum *int64 `json:"ModelIdNum,omitempty" name:"ModelIdNum"`

		// 素材数据
	// 注意：此字段可能返回 null，表示取不到有效值。
		ModelInfos []*ModelInfo `json:"ModelInfos,omitempty" name:"ModelInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetModelListResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LipColorInfo struct {

	// 使用RGBA模型试唇色。
	RGBA *RGBAInfo `json:"RGBA,omitempty" name:"RGBA"`

	// 使用已注册的 LUT 文件试唇色。  
	// ModelId 和 RGBA 两个参数只需提供一个，若都提供只使用 ModelId。
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。  
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 涂妆浓淡[0,100]。建议取值50。本参数仅控制ModelId对应的涂妆浓淡。
	ModelAlpha *int64 `json:"ModelAlpha,omitempty" name:"ModelAlpha"`
}

type ModelInfo struct {

	// 唇色素材ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 唇色素材 url 。 LUT 文件 url 5分钟有效。
	LUTFileUrl *string `json:"LUTFileUrl,omitempty" name:"LUTFileUrl"`

	// 文件描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type RGBAInfo struct {

	// R通道数值。[0,255]。
	R *int64 `json:"R,omitempty" name:"R"`

	// G通道数值。[0,255]。
	G *int64 `json:"G,omitempty" name:"G"`

	// B通道数值。[0,255]。
	B *int64 `json:"B,omitempty" name:"B"`

	// A通道数值。[0,100]。建议取值50。
	A *int64 `json:"A,omitempty" name:"A"`
}

type TryLipstickPicRequest struct {
	*tchttp.BaseRequest

	// 唇色信息。 
	// 您可以输入最多3个 LipColorInfo 来实现给一张图中的最多3张人脸试唇色。
	LipColorInfos []*LipColorInfo `json:"LipColorInfos,omitempty" name:"LipColorInfos" list`

	// 图片 base64 数据，base64 编码后大小不可超过6M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过6M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *TryLipstickPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TryLipstickPicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TryLipstickPicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果图片Base64信息。
		ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TryLipstickPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TryLipstickPicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
