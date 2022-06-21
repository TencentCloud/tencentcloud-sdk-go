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

package v20200417

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Location struct {
	// 位置矩形框的左上角横坐标
	XMin *int64 `json:"XMin,omitempty" name:"XMin"`

	// 位置矩形框的左上角纵坐标
	YMin *int64 `json:"YMin,omitempty" name:"YMin"`

	// 位置矩形框的右下角横坐标
	XMax *int64 `json:"XMax,omitempty" name:"XMax"`

	// 位置矩形框的右下角纵坐标
	YMax *int64 `json:"YMax,omitempty" name:"YMax"`
}

type ProductInfo struct {
	// 1表示找到同款商品，以下字段为同款商品信息； 
	// 0表示未找到同款商品， 具体商品信息为空（参考价格、名称、品牌等），仅提供商品类目。  
	// 是否找到同款的判断依据为Score分值，分值越大则同款的可能性越大。
	FindSKU *int64 `json:"FindSKU,omitempty" name:"FindSKU"`

	// 本商品在图片中的坐标，表示为矩形框的四个顶点坐标。
	Location *Location `json:"Location,omitempty" name:"Location"`

	// 商品名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 商品品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 参考价格，综合多个信息源，仅供参考。
	Price *string `json:"Price,omitempty" name:"Price"`

	// 识别结果的商品类目。 
	// 包含：鞋、图书音像、箱包、美妆个护、服饰、家电数码、玩具乐器、食品饮料、珠宝、家居家装、药品、酒水、绿植园艺、其他商品、非商品等。 
	// 当类别为“非商品”时，除Location、Score和本字段之外的商品信息为空。
	ProductCategory *string `json:"ProductCategory,omitempty" name:"ProductCategory"`

	// 输入图片中的主体物品和输出结果的相似度。分值越大，输出结果与输入图片是同款的可能性越高。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 搜索到的商品配图URL
	Image *string `json:"Image,omitempty" name:"Image"`
}

// Predefined struct for user
type RecognizeProductRequestParams struct {
	// 图片限制：内测版仅支持jpg、jpeg，图片大小不超过1M，分辨率在25万到100万之间。 
	// 建议先对图片进行压缩，以便提升处理速度。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过1M，分辨率在25万到100万之间。 
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

type RecognizeProductRequest struct {
	*tchttp.BaseRequest
	
	// 图片限制：内测版仅支持jpg、jpeg，图片大小不超过1M，分辨率在25万到100万之间。 
	// 建议先对图片进行压缩，以便提升处理速度。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过1M，分辨率在25万到100万之间。 
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *RecognizeProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ImageUrl")
	delete(f, "ImageBase64")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecognizeProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecognizeProductResponseParams struct {
	// 检测到的图片中的商品位置和品类预测。 
	// 当图片中存在多个商品时，输出多组坐标，按照__显著性__排序（综合考虑面积、是否在中心、检测算法置信度）。 
	// 最多可以输出__3组__检测结果。
	RegionDetected []*RegionDetected `json:"RegionDetected,omitempty" name:"RegionDetected"`

	// 图像识别出的商品的详细信息。 
	// 当图像中检测到多个物品时，会对显著性最高的进行识别。
	ProductInfo *ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RecognizeProductResponse struct {
	*tchttp.BaseResponse
	Response *RecognizeProductResponseParams `json:"Response"`
}

func (r *RecognizeProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecognizeProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionDetected struct {
	// 商品的品类预测结果。 
	// 包含：鞋、图书音像、箱包、美妆个护、服饰、家电数码、玩具乐器、食品饮料、珠宝、家居家装、药品、酒水、绿植园艺、其他商品、非商品等。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 商品品类预测的置信度
	CategoryScore *float64 `json:"CategoryScore,omitempty" name:"CategoryScore"`

	// 检测到的主体在图片中的坐标，表示为矩形框的四个顶点坐标
	Location *Location `json:"Location,omitempty" name:"Location"`
}