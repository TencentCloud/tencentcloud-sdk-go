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

package v20200304

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgeInfo struct {

	// 变化到的人脸年龄 [10,80]。
	Age *int64 `json:"Age,omitempty" name:"Age"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。  
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`
}

type ChangeAgePicRequest struct {
	*tchttp.BaseRequest

	// 人脸变老变年轻信息。 
	// 您可以输入最多3个 AgeInfo 来实现给一张图中的最多3张人脸变老变年轻。
	AgeInfos []*AgeInfo `json:"AgeInfos,omitempty" name:"AgeInfos" list`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *ChangeAgePicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeAgePicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ChangeAgePicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// base64编码图片
		ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ChangeAgePicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ChangeAgePicResponse) FromJsonString(s string) error {
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

type GenderInfo struct {

	// 选择转换方向，0：男变女，1：女变男。
	Gender *int64 `json:"Gender,omitempty" name:"Gender"`

	// 人脸框位置。若不输入则选择 Image 或 Url 中面积最大的人脸。  
	// 您可以通过 [人脸检测与分析](https://cloud.tencent.com/document/api/867/32800)  接口获取人脸框位置信息。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`
}

type SwapGenderPicRequest struct {
	*tchttp.BaseRequest

	// 人脸转化性别信息。 
	// 您可以输入最多3个 GenderInfo 来实现给一张图中的最多3张人脸转换性别。
	GenderInfos []*GenderInfo `json:"GenderInfos,omitempty" name:"GenderInfos" list`

	// 图片 base64 数据，base64 编码后大小不可超过5M。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Image *string `json:"Image,omitempty" name:"Image"`

	// 图片的 Url ，对应图片 base64 编码后大小不可超过5M。 
	// 图片的 Url、Image必须提供一个，如果都提供，只使用 Url。 
	// 图片存储于腾讯云的 Url 可保障更高下载速度和稳定性，建议图片存储于腾讯云。 
	// 非腾讯云存储的Url速度和稳定性可能受一定影响。 
	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。
	Url *string `json:"Url,omitempty" name:"Url"`
}

func (r *SwapGenderPicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwapGenderPicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SwapGenderPicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果图片Base64信息。
		ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwapGenderPicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SwapGenderPicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
