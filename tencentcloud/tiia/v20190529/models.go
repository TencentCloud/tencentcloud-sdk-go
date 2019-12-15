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

type Candidate struct {

	// 识别出人脸对应的候选人数组。当前返回相似度最高的候选人。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 相似度，0-100之间。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`
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

		// 0：抠图正常, 1：原图过长, 2：原图过宽, 3：抠图区域过长, 4：抠图区域过宽, 5：纯色图 6：宽高比异常
	// 原图过长是指原图的高度是宽度的1.8倍以上；
	// 原图过宽是指原图的宽度是高度的1.8倍以上；
	// 抠图区域过长是指抠图的高度是主体备选框高度的1.6倍以上；
	// 抠图过宽是指当没有检测到人脸时，抠图区域宽度是检测出的原图主体区域宽度的1.6倍以上；
	// 纯色图是指裁剪区域视觉较为单一、缺乏主体部分；
	// Width : Height取值超出[1, 2.5]的范围；
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

type DisgustResult struct {

	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 图像恶心的分数，0-100之间，分数越高恶心几率越大。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`
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
	// 当Confidence低于70分时，Name仅供参考。您可以根据业务实际情况调整阈值。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 人脸区域左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸区域左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸区域宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸区域高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type FaceRect struct {

	// 人脸区域左上角横坐标。
	X *int64 `json:"X,omitempty" name:"X"`

	// 人脸区域左上角纵坐标。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 人脸区域宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 人脸区域高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type FaceResult struct {

	// 检测出的人脸框位置。
	FaceRect *FaceRect `json:"FaceRect,omitempty" name:"FaceRect"`

	// 候选人列表。当前返回相似度最高的候选人。
	Candidates []*Candidate `json:"Candidates,omitempty" name:"Candidates" list`
}

type ImageModerationRequest struct {
	*tchttp.BaseRequest

	// 本次调用支持的识别场景，可选值如下：
	// 1. PORN，即色情识别
	// 2. TERRORISM，即暴恐识别
	// 3. POLITICS，即政治敏感识别
	// 4. TEXT, 即图像文本识别
	// 
	// 支持多场景（Scenes）一起检测。例如，使用 Scenes=["PORN", "TERRORISM"]，即对一张图片同时进行色情识别和暴恐识别。
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes" list`

	// 图片URL地址。 
	// 图片限制： 
	//  • 图片格式：PNG、JPG、JPEG。 
	//  • 图片大小：所下载图片经Base64编码后不超过4M。图片下载时间不超过3秒。 TEXT场景要求图片经Base64编码后不超过3M。
	//  • 图片像素：大于50*50像素，否则影响识别效果； 
	//  • 长宽比：长边：短边<5； 
	// 接口响应时间会受到图片下载时间的影响，建议使用更可靠的存储服务，推荐将图片存储在腾讯云COS。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 预留字段，后期用于展示更多识别信息。
	Config *string `json:"Config,omitempty" name:"Config"`

	// 透传字段，透传简单信息。
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 图片经过base64编码的内容。最大不超过4M。与ImageUrl同时存在时优先使用ImageUrl字段。
	// **注意：图片需要base64编码，并且要去掉编码头部。**
	ImageBase64 *string `json:"ImageBase64,omitempty" name:"ImageBase64"`
}

func (r *ImageModerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageModerationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ImageModerationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
		Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

		// 色情识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PornResult *PornResult `json:"PornResult,omitempty" name:"PornResult"`

		// 暴恐识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TerrorismResult *TerrorismResult `json:"TerrorismResult,omitempty" name:"TerrorismResult"`

		// 政治敏感识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PoliticsResult *PoliticsResult `json:"PoliticsResult,omitempty" name:"PoliticsResult"`

		// 透传字段，透传简单信息。
		Extra *string `json:"Extra,omitempty" name:"Extra"`

		// 恶心内容识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		DisgustResult *DisgustResult `json:"DisgustResult,omitempty" name:"DisgustResult"`

		// 文字识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TextResult *TextResult `json:"TextResult,omitempty" name:"TextResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ImageModerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ImageModerationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Labels struct {

	// 公众人物身份标签的一级分类，例如体育明星、娱乐明星、政治人物等；
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstLabel *string `json:"FirstLabel,omitempty" name:"FirstLabel"`

	// 公众人物身份标签的二级分类，例如歌手（对应一级标签为“娱乐明星”）；
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondLabel *string `json:"SecondLabel,omitempty" name:"SecondLabel"`
}

type PoliticsResult struct {

	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误，
	// -1400表示图片解码失败，
	// -1401表示图片不符合规范。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 图像涉政的分数，0-100之间，分数越高涉政几率越大。
	// Type为DNA时：
	// 0到75，Suggestion建议为PASS
	// 75到90，Suggestion建议为REVIEW
	// 90到100，Suggestion建议为BLOCK
	// Type为FACE时：
	// 0到70，Suggestion建议为PASS
	// 70到80，Suggestion建议为REVIEW
	// 80到100，Suggestion建议为BLOCK
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// Type取值为‘FACE’时，人脸识别的结果列表。基于图片中实际检测到的人脸数，返回数组最大值不超过5个。
	FaceResults []*FaceResult `json:"FaceResults,omitempty" name:"FaceResults" list`

	// 取值'DNA' 或‘FACE’。DNA表示结论和置信度来自图像指纹，FACE表示结论和置信度来自人脸识别。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 鉴政识别返回的详细标签后期开放。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}

type PornResult struct {

	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误，
	// -1400表示图片解码失败。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 算法对于Suggestion的置信度，0-100之间，值越高，表示对于Suggestion越确定。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 预留字段，后期用于展示更多识别信息。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 色情识别类型：
	// PORN：色情
	// HOT：性感
	// NORMAL：正常
	// FAIL：识别失败
	Type *string `json:"Type,omitempty" name:"Type"`
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

		// 汽车的四个矩形顶点坐标
		CarCoords []*Coord `json:"CarCoords,omitempty" name:"CarCoords" list`

		// 车辆属性识别的结果数组
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

type TerrorismResult struct {

	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误，
	// -1400表示图片解码失败。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 图像涉恐的分数，0-100之间，分数越高涉恐几率越大。
	// Type为LABEL时：
	// 0到86，Suggestion建议为PASS
	// 86到91，Suggestion建议为REVIEW
	// 91到100，Suggestion建议为BLOCK
	// Type为FACE时：
	// 0到70，Suggestion建议为PASS
	// 70到80，Suggestion建议为REVIEW
	// 80到100，Suggestion建议为BLOCK
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// Type取值为‘FACE’时，人脸识别的结果列表。基于图片中实际检测到的人脸数，返回数组最大值不超过5个。
	FaceResults []*FaceResult `json:"FaceResults,omitempty" name:"FaceResults" list`

	// 暴恐识别返回的详细标签后期开放。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`

	// 取值'LABEL' 或‘FACE’，LABEL表示结论和置信度来自标签分类，FACE表示结论和置信度来自人脸识别。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type TextResult struct {

	// 该识别场景的错误码：
	// 0表示成功，
	// -1表示系统错误，
	// -2表示引擎错误，
	// -1400表示图片解码失败，
	// -1401表示图片不符合规范。
	// -1402表示图片文件太大。
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误码描述信息。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 识别场景的审核结论：
	// PASS：正常
	// REVIEW：疑似
	// BLOCK：违规
	// 
	// Suggestion由Type决定：
	// Type为 NOTEXT/NORMAL 时，Suggestion为PASS；
	// Type为 POLITICS/PORN/TERRORISM/ADS 时，Suggestion为BLOCK；
	// 其他情况下Suggestion为REVIEW。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 算法对于识别结果的置信度，0-100之间，值越高，表示对于结论越确定。
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别到的关键词数组
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 图片中是否包含敏感文本内容。
	// 包含：
	// NOTEXT：无文本
	// NORMAL：内容正常
	// ADS：广告推广
	// POLITICS：政治
	// PORN：色情
	// DRUGS：涉毒
	// CURSE：谩骂
	// TERRORISM：暴恐
	// OTHERS：其他
	// 
	// 本服务利用微信团队提供的算法，可以准确识别图片中是否包含二维码。当图片中存在二维码时，分类为ADS，关键词输出为“二维码”。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 预留字段，后期用于展示更多识别信息。
	AdvancedInfo *string `json:"AdvancedInfo,omitempty" name:"AdvancedInfo"`
}
