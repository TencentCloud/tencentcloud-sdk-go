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

package v20221229

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type ImageToImageRequestParams struct {
	// 输入图 Base64 数据。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	InputImage *string `json:"InputImage,omitempty" name:"InputImage"`

	// 输入图 Url。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	InputUrl *string `json:"InputUrl,omitempty" name:"InputUrl"`

	// 文本描述。
	// 用于在输入图的基础上引导生成图效果，建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。推荐使用中文。最多支持512个 utf-8 字符。
	// 注意：如果不输入任何文本描述，可能导致较差的效果，建议根据期望的效果输入相应的文本描述。
	Prompt *string `json:"Prompt,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传512个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在  [智能图生图风格列表](https://cloud.tencent.com/document/product/1668/86250) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	// 如果想要探索风格列表之外的风格，也可以尝试在 Prompt 中输入其他的风格描述。
	Styles []*string `json:"Styles,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	ResultConfig *ResultConfig `json:"ResultConfig,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitempty" name:"LogoParam"`

	// 生成自由度。
	// Strength 值越小，生成图和原图越接近。取值范围0~1，不传默认为0.6。
	Strength *float64 `json:"Strength,omitempty" name:"Strength"`
}

type ImageToImageRequest struct {
	*tchttp.BaseRequest
	
	// 输入图 Base64 数据。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	InputImage *string `json:"InputImage,omitempty" name:"InputImage"`

	// 输入图 Url。
	// 算法将根据输入的图片，结合文本描述智能生成与之相关的图像。
	// Base64 和 Url 必须提供一个，如果都提供以 Base64 为准。
	// 图片限制：单边分辨率小于2000，转成 Base64 字符串后小于 5MB。
	InputUrl *string `json:"InputUrl,omitempty" name:"InputUrl"`

	// 文本描述。
	// 用于在输入图的基础上引导生成图效果，建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。推荐使用中文。最多支持512个 utf-8 字符。
	// 注意：如果不输入任何文本描述，可能导致较差的效果，建议根据期望的效果输入相应的文本描述。
	Prompt *string `json:"Prompt,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传512个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在  [智能图生图风格列表](https://cloud.tencent.com/document/product/1668/86250) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	// 如果想要探索风格列表之外的风格，也可以尝试在 Prompt 中输入其他的风格描述。
	Styles []*string `json:"Styles,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	ResultConfig *ResultConfig `json:"ResultConfig,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitempty" name:"LogoParam"`

	// 生成自由度。
	// Strength 值越小，生成图和原图越接近。取值范围0~1，不传默认为0.6。
	Strength *float64 `json:"Strength,omitempty" name:"Strength"`
}

func (r *ImageToImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputImage")
	delete(f, "InputUrl")
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Styles")
	delete(f, "ResultConfig")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	delete(f, "Strength")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageToImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageToImageResponseParams struct {
	// 返回的生成图 Base64 编码。
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ImageToImageResponse struct {
	*tchttp.BaseResponse
	Response *ImageToImageResponseParams `json:"Response"`
}

func (r *ImageToImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageToImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogoParam struct {
	// 水印url
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogoUrl *string `json:"LogoUrl,omitempty" name:"LogoUrl"`

	// 水印base64，url和base64二选一传入
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogoImage *string `json:"LogoImage,omitempty" name:"LogoImage"`

	// 水印图片位于融合结果图中的坐标，将按照坐标对标识图片进行位置和大小的拉伸匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogoRect *LogoRect `json:"LogoRect,omitempty" name:"LogoRect"`
}

type LogoRect struct {
	// 左上角X坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *int64 `json:"X,omitempty" name:"X"`

	// 左上角Y坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *int64 `json:"Y,omitempty" name:"Y"`

	// 方框宽度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 方框高度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type ResultConfig struct {
	// 生成图分辨率
	// 支持生成以下不同分辨率的图片，对应1:1方图、3:4竖图、4:3横图三种尺寸规格，不传默认为"768:768"
	// 取值：
	// ● 768:768
	// ● 768:1024
	// ● 1024:768
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resolution *string `json:"Resolution,omitempty" name:"Resolution"`
}

// Predefined struct for user
type TextToImageRequestParams struct {
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传512个 utf-8 字符。
	Prompt *string `json:"Prompt,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传512个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [智能文生图风格列表](https://cloud.tencent.com/document/product/1668/86249) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	// 如果想要探索风格列表之外的风格，也可以尝试在 Prompt 中输入其他的风格描述。
	Styles []*string `json:"Styles,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	ResultConfig *ResultConfig `json:"ResultConfig,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitempty" name:"LogoParam"`
}

type TextToImageRequest struct {
	*tchttp.BaseRequest
	
	// 文本描述。
	// 算法将根据输入的文本智能生成与之相关的图像。建议详细描述画面主体、细节、场景等，文本描述越丰富，生成效果越精美。
	// 不能为空，推荐使用中文。最多可传512个 utf-8 字符。
	Prompt *string `json:"Prompt,omitempty" name:"Prompt"`

	// 反向文本描述。
	// 用于一定程度上从反面引导模型生成的走向，减少生成结果中出现描述内容的可能，但不能完全杜绝。
	// 推荐使用中文。最多可传512个 utf-8 字符。
	NegativePrompt *string `json:"NegativePrompt,omitempty" name:"NegativePrompt"`

	// 绘画风格。
	// 请在 [智能文生图风格列表](https://cloud.tencent.com/document/product/1668/86249) 中选择期望的风格，传入风格编号。
	// 推荐使用且只使用一种风格。不传默认使用201（日系动漫风格）。
	// 如果想要探索风格列表之外的风格，也可以尝试在 Prompt 中输入其他的风格描述。
	Styles []*string `json:"Styles,omitempty" name:"Styles"`

	// 生成图结果的配置，包括输出图片分辨率和尺寸等。
	ResultConfig *ResultConfig `json:"ResultConfig,omitempty" name:"ResultConfig"`

	// 为生成结果图添加标识的开关，默认为1。
	// 1：添加标识。
	// 0：不添加标识。
	// 其他数值：默认按1处理。
	// 建议您使用显著标识来提示结果图使用了 AI 绘画技术，是 AI 生成的图片。
	LogoAdd *int64 `json:"LogoAdd,omitempty" name:"LogoAdd"`

	// 标识内容设置。
	// 默认在生成结果图右下角添加“图片由 AI 生成”字样，您可根据自身需要替换为其他的标识图片。
	LogoParam *LogoParam `json:"LogoParam,omitempty" name:"LogoParam"`
}

func (r *TextToImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Prompt")
	delete(f, "NegativePrompt")
	delete(f, "Styles")
	delete(f, "ResultConfig")
	delete(f, "LogoAdd")
	delete(f, "LogoParam")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextToImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextToImageResponseParams struct {
	// 返回的生成图 Base64 编码。
	ResultImage *string `json:"ResultImage,omitempty" name:"ResultImage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextToImageResponse struct {
	*tchttp.BaseResponse
	Response *TextToImageResponseParams `json:"Response"`
}

func (r *TextToImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextToImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}