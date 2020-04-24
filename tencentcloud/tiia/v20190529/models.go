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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AssessQualityRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *AssessQualityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssessQualityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AssessQualityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 取值为TRUE或FALSE，TRUE为长图，FALSE为正常图，长图定义为长宽比大于等于3或小于等于1/3的图片。
		LongImage *bool `json:"LongImage,omitempty" name:"LongImage"`

		// 取值为TRUE或FALSE，TRUE为黑白图，FALSE为否。黑白图即灰度图，指红绿蓝三个通道都是以灰度色阶显示的图片，并非视觉上的“黑白图片”。
		BlackAndWhite *bool `json:"BlackAndWhite,omitempty" name:"BlackAndWhite"`

		// 取值为TRUE或FALSE，TRUE为小图，FALSE为否, 小图定义为最长边小于179像素的图片。当一张图片被判断为小图时，不建议做推荐和展示，其他字段统一输出为0或FALSE。
		SmallImage *bool `json:"SmallImage,omitempty" name:"SmallImage"`

		// 取值为TRUE或FALSE，TRUE为大图，FALSE为否，定义为最短边大于1000像素的图片
		BigImage *bool `json:"BigImage,omitempty" name:"BigImage"`

		// 取值为TRUE或FALSE，TRUE为纯色图或纯文字图，即没有内容或只有简单内容的图片，FALSE为正常图片。
		PureImage *bool `json:"PureImage,omitempty" name:"PureImage"`

		// 综合评分。图像清晰度的得分，对图片的噪声、曝光、模糊、压缩等因素进行综合评估，取值为[0, 100]，值越大，越清晰。一般大于50为较清晰图片，标准可以自行把握。
		ClarityScore *int64 `json:"ClarityScore,omitempty" name:"ClarityScore"`

		// 综合评分。图像美观度得分， 从构图、色彩等多个艺术性维度评价图片，取值为[0, 100]，值越大，越美观。一般大于50为较美观图片，标准可以自行把握。
		AestheticScore *int64 `json:"AestheticScore,omitempty" name:"AestheticScore"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AssessQualityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AssessQualityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CarTagItem struct {

	// 车系
	Serial *string `json:"Serial,omitempty" name:"Serial"`

	// 车辆品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 车辆类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 车辆颜色
	Color *string `json:"Color,omitempty" name:"Color"`

	// 置信度，0-100
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 年份，没识别出年份的时候返回0
	Year *int64 `json:"Year,omitempty" name:"Year"`

	// 车辆在图片中的坐标信息
	CarLocation []*Coord `json:"CarLocation,omitempty" name:"CarLocation" list`
}

type Coord struct {

	// 横坐标x
	X *int64 `json:"X,omitempty" name:"X"`

	// 纵坐标y
	Y *int64 `json:"Y,omitempty" name:"Y"`
}

type CropImageRequest struct {
	*tchttp.BaseRequest

	// 需要裁剪区域的宽度，与Height共同组成所需裁剪的图片宽高比例；
	// 输入数字请大于0、小于图片宽度的像素值；
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 需要裁剪区域的高度，与Width共同组成所需裁剪的图片宽高比例；
	// 输入数字请请大于0、小于图片高度的像素值；
	// 宽高比例（Width : Height）会简化为最简分数，即如果Width输入10、Height输入20，会简化为1：2。
	// Width : Height建议取值在[1, 2.5]之间，超过这个范围可能会影响效果；
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *CropImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CropImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CropImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 裁剪区域左上角X坐标值
		X *int64 `json:"X,omitempty" name:"X"`

		// 裁剪区域左上角Y坐标值
		Y *int64 `json:"Y,omitempty" name:"Y"`

		// 裁剪区域的宽度，单位为像素
		Width *int64 `json:"Width,omitempty" name:"Width"`

		// 裁剪区域的高度，单位为像素
		Height *int64 `json:"Height,omitempty" name:"Height"`

		// 原图宽度，单位为像素
		OriginalWidth *int64 `json:"OriginalWidth,omitempty" name:"OriginalWidth"`

		// 原图高度，单位为像素
		OriginalHeight *int64 `json:"OriginalHeight,omitempty" name:"OriginalHeight"`

		// 0：抠图正常；
	// 1：原图过长，指原图的高度是宽度的1.8倍以上；
	// 2：原图过宽，指原图的宽度是高度的1.8倍以上；
	// 3：抠图区域过长，指抠图的高度是主体备选框高度的1.6倍以上；
	// 4：抠图区域过宽，指当没有检测到人脸时，抠图区域宽度是检测出的原图主体区域宽度的1.6倍以上；
	// 5：纯色图，指裁剪区域视觉较为单一、缺乏主体部分 ；
	// 6：宽高比异常，指Width : Height取值超出[1, 2.5]的范围；
	// 
	// 以上是辅助决策的参考建议，可以根据业务需求选择采纳或忽视。
		CropResult *int64 `json:"CropResult,omitempty" name:"CropResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CropImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CropImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectCelebrityRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectCelebrityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectCelebrityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectCelebrityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 公众人物识别结果数组。如果检测不到人脸，返回为空；最多可以返回10个人脸识别结果。
		Faces []*Face `json:"Faces,omitempty" name:"Faces" list`

		// 本服务在不同误识率水平下（将图片中的人物识别错误的比例）的推荐阈值，可以用于控制识别结果的精度。 
	// FalseRate1Percent, FalseRate5Permil, FalseRate1Permil分别代表误识率在百分之一、千分之五、千分之一情况下的推荐阈值。 
	// 因为阈值会存在变动，请勿将此处输出的固定值处理，而是每次取值与confidence对比，来判断本次的识别结果是否可信。
	//  例如，如果您业务中可以接受的误识率是1%，则可以将所有confidence>=FalseRate1Percent的结论认为是正确的。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Threshold *Threshold `json:"Threshold,omitempty" name:"Threshold"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectCelebrityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectCelebrityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectDisgustRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectDisgustRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectDisgustRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectDisgustResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 对于图片中包含恶心内容的置信度，取值[0,1]，一般超过0.5则表明可能是恶心图片。
		Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

		// 与图像内容最相似的恶心内容的类别，包含腐烂、密集、畸形、血腥、蛇、虫子、牙齿等。
		Type *string `json:"Type,omitempty" name:"Type"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectDisgustResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectDisgustResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectLabelItem struct {

	// 图片中的物体名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 算法对于Name的置信度，0-100之间，值越高，表示对于Name越确定。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 标签的一级分类
	FirstCategory *string `json:"FirstCategory,omitempty" name:"FirstCategory"`

	// 标签的二级分类
	SecondCategory *string `json:"SecondCategory,omitempty" name:"SecondCategory"`
}

type DetectLabelRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`

	// 本次调用支持的识别场景，可选值如下：
	// WEB，针对网络图片优化;
	// CAMERA，针对手机摄像头拍摄图片优化;
	// ALBUM，针对手机相册、网盘产品优化;
	// 如果不传此参数，则默认为WEB。
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["WEB", "CAMERA"]，即对一张图片使用两个模型同时检测，输出两套识别结果。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes" list`
}

func (r *DetectLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectLabelRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectLabelResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Web网络版标签结果数组。如未选择WEB场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Labels []*DetectLabelItem `json:"Labels,omitempty" name:"Labels" list`

		// Camera摄像头版标签结果数组。如未选择CAMERA场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CameraLabels []*DetectLabelItem `json:"CameraLabels,omitempty" name:"CameraLabels" list`

		// Album相册版标签结果数组。如未选择ALBUM场景，则为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AlbumLabels []*DetectLabelItem `json:"AlbumLabels,omitempty" name:"AlbumLabels" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectLabelResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectMisbehaviorRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectMisbehaviorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectMisbehaviorRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectMisbehaviorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 对于图片中包含不良行为的置信度，取值[0,1]，一般超过0.5则表明可能包含不良行为内容；
		Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

		// 图像中最可能包含的不良行为类别，包括赌博、打架斗殴、吸毒等。
		Type *string `json:"Type,omitempty" name:"Type"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectMisbehaviorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectMisbehaviorResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectProductBetaRequest struct {
	*tchttp.BaseRequest

	// 图片限制：内测版仅支持jpg、jpeg，图片大小不超过1M，分辨率在25万到100万之间。 
	// 建议先对图片进行压缩，以便提升处理速度。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过1M，分辨率在25万到100万之间。 
	// 与ImageUrl同时存在时优先使用ImageUrl字段。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectProductBetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectProductBetaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectProductBetaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 检测到的图片中的商品位置和品类预测。 
	// 当图片中存在多个商品时，输出多组坐标，按照__显著性__排序（综合考虑面积、是否在中心、检测算法置信度）。 
	// 最多可以输出__3组__检测结果。
		RegionDetected []*RegionDetected `json:"RegionDetected,omitempty" name:"RegionDetected" list`

		// 图像识别出的商品的详细信息。 
	// 当图像中检测到多个物品时，会对显著性最高的进行识别。
		ProductInfo *ProductInfo `json:"ProductInfo,omitempty" name:"ProductInfo"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectProductBetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectProductBetaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectProductRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *DetectProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectProductRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetectProductResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 商品识别结果数组
		Products []*Product `json:"Products,omitempty" name:"Products" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetectProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetectProductResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhanceImageRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 支持PNG、JPG、JPEG、BMP，不支持 GIF 图片。图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *EnhanceImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnhanceImageRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhanceImageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 增强后图片的base64编码。
		EnhancedImage *string `json:"EnhancedImage,omitempty" name:"EnhancedImage"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnhanceImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnhanceImageResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Face struct {

	// 与图片中人脸最相似的公众人物的名字。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 公众人物身份标签的数组，一个公众人物可能有多个身份标签。
	Labels []*Labels `json:"Labels,omitempty" name:"Labels" list`

	// 对人物的简介。
	BasicInfo *string `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 算法对于Name的置信度（图像中人脸与公众人物的相似度），0-100之间，值越高，表示对于Name越确定。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 人脸区域左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸区域左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸区域宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸区域高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 公众人物的唯一编号，可以用于区分同名人物、一个人物不同称呼等情况。唯一编号为8个字符构成的字符串。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitempty" name:"ID"`
}

type Labels struct {

	// 公众人物身份标签的一级分类，例如体育明星、娱乐明星、政治人物等；
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstLabel *string `json:"FirstLabel,omitempty" name:"FirstLabel"`

	// 公众人物身份标签的二级分类，例如歌手（对应一级标签为“娱乐明星”）；
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondLabel *string `json:"SecondLabel,omitempty" name:"SecondLabel"`
}

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

type Product struct {

	// 图片中商品的三级分类识别结果，选取所有三级分类中的置信度最大者
	Name *string `json:"Name,omitempty" name:"Name"`

	// 三级商品分类对应的一级分类和二级分类，两级之间用“-”（中划线）隔开，例如商品名称是“硬盘”，那么Parents输出为“电脑、办公-电脑配件”
	Parents *string `json:"Parents,omitempty" name:"Parents"`

	// 算法对于Name的置信度，0-100之间，值越高，表示对于Name越确定
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 商品坐标X轴的最小值
	XMin *int64 `json:"XMin,omitempty" name:"XMin"`

	// 商品坐标Y轴的最小值
	YMin *int64 `json:"YMin,omitempty" name:"YMin"`

	// 商品坐标X轴的最大值
	XMax *int64 `json:"XMax,omitempty" name:"XMax"`

	// 商品坐标Y轴的最大值
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

type RecognizeCarRequest struct {
	*tchttp.BaseRequest

	// 图片URL地址。 
	// 图片限制： 
	// • 图片格式：PNG、JPG、JPEG。 
	// • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 
	// 建议：
	// • 图片像素：大于50*50像素，否则影响识别效果； 
	// • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	// 支持的图片格式：PNG、JPG、JPEG、BMP，暂不支持GIF格式。支持的图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *RecognizeCarRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecognizeCarRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RecognizeCarResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 汽车的四个矩形顶点坐标，如果图片中存在多辆车，则输出最大车辆的坐标。
		CarCoords []*Coord `json:"CarCoords,omitempty" name:"CarCoords" list`

		// 车辆属性识别的结果数组，如果识别到多辆车，则会输出每辆车的top1结果。
		CarTags []*CarTagItem `json:"CarTags,omitempty" name:"CarTags" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RecognizeCarResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RecognizeCarResponse) FromJsonString(s string) error {
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

type Threshold struct {

	// 误识率在百分之一时的推荐阈值。
	FalseRate1Percent *int64 `json:"FalseRate1Percent,omitempty" name:"FalseRate1Percent"`

	// 误识率在千分之五时的推荐阈值。
	FalseRate5Permil *int64 `json:"FalseRate5Permil,omitempty" name:"FalseRate5Permil"`

	// 误识率在千分之一时的推荐阈值。
	FalseRate1Permil *int64 `json:"FalseRate1Permil,omitempty" name:"FalseRate1Permil"`
}
