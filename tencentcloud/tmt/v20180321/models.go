// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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

package v20180321

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type BoundingBox struct {
	// <p>左上顶点x坐标</p>
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// <p>左上顶点y坐标</p>
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// <p>宽</p><p>单位：px</p>
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// <p>高</p><p>单位：px</p>
	Height *int64 `json:"Height,omitnil,omitempty" name:"Height"`
}

type Coord struct {
	// X坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// Y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`
}

// Predefined struct for user
type ImageTranslateLLMRequestParams struct {
	// <p>图片数据的Base64字符串，经Base64编码后不超过 9M，分辨率建议600*800以上，支持PNG、JPG、JPEG格式。</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>目标语言，支持语言列表：</p><ul><li>中文：zh</li><li>繁体（中国台湾）：zh-TW</li><li>繁体（中国香港）：zh-HK</li><li>英文：en</li><li>日语：ja</li><li>韩语：ko</li><li>泰语：th</li><li>越南语：vi</li><li>俄语：ru</li><li>德语：de</li><li>法语：fr</li><li>阿拉伯语：ar</li><li>西班牙语：es</li><li>意大利语：it</li><li>印度尼西亚语：id</li><li>马来西亚语：ms</li><li>葡萄牙语：pt</li><li>土耳其语：tr<br>-</li></ul>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>输入图 Url。 使用Url的时候，Data参数需要传入&quot;&quot;。 图片限制：小于 10MB，分辨率建议600*800以上，格式支持 jpg、jpeg、png。</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>调用模式。</p><p>枚举值：</p><ul><li>0： 端到端图片翻译大模型pro版</li><li>1： 端到端图片翻译大模型lite版</li></ul><p>默认值：0</p>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

type ImageTranslateLLMRequest struct {
	*tchttp.BaseRequest
	
	// <p>图片数据的Base64字符串，经Base64编码后不超过 9M，分辨率建议600*800以上，支持PNG、JPG、JPEG格式。</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>目标语言，支持语言列表：</p><ul><li>中文：zh</li><li>繁体（中国台湾）：zh-TW</li><li>繁体（中国香港）：zh-HK</li><li>英文：en</li><li>日语：ja</li><li>韩语：ko</li><li>泰语：th</li><li>越南语：vi</li><li>俄语：ru</li><li>德语：de</li><li>法语：fr</li><li>阿拉伯语：ar</li><li>西班牙语：es</li><li>意大利语：it</li><li>印度尼西亚语：id</li><li>马来西亚语：ms</li><li>葡萄牙语：pt</li><li>土耳其语：tr<br>-</li></ul>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>输入图 Url。 使用Url的时候，Data参数需要传入&quot;&quot;。 图片限制：小于 10MB，分辨率建议600*800以上，格式支持 jpg、jpeg、png。</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>调用模式。</p><p>枚举值：</p><ul><li>0： 端到端图片翻译大模型pro版</li><li>1： 端到端图片翻译大模型lite版</li></ul><p>默认值：0</p>
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

func (r *ImageTranslateLLMRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateLLMRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Data")
	delete(f, "Target")
	delete(f, "Url")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ImageTranslateLLMRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ImageTranslateLLMResponseParams struct {
	// <p>图片数据的Base64字符串，输出格式为JPG。</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>原文本主要源语言。</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>目标翻译语言。</p>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>图片中的全部原文本。</p>
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// <p>图片中全部译文。</p>
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// <p>逆时针图片角度，取值范围为0-359</p>
	Angle *float64 `json:"Angle,omitnil,omitempty" name:"Angle"`

	// <p>翻译详情信息</p>
	TransDetails []*TransDetail `json:"TransDetails,omitnil,omitempty" name:"TransDetails"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ImageTranslateLLMResponse struct {
	*tchttp.BaseResponse
	Response *ImageTranslateLLMResponseParams `json:"Response"`
}

func (r *ImageTranslateLLMResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ImageTranslateLLMResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RotateParagraphRect struct {
	// 段落文本坐标
	Coord []*Coord `json:"Coord,omitnil,omitempty" name:"Coord"`

	// 旋转角度
	TiltAngle *float64 `json:"TiltAngle,omitnil,omitempty" name:"TiltAngle"`

	// 段落文本信息是否有效
	Valid *bool `json:"Valid,omitnil,omitempty" name:"Valid"`
}

type TransDetail struct {
	// <p>当前行的原文本</p>
	SourceLineText *string `json:"SourceLineText,omitnil,omitempty" name:"SourceLineText"`

	// <p>当前行的译文</p>
	TargetLineText *string `json:"TargetLineText,omitnil,omitempty" name:"TargetLineText"`

	// <p>段落文本框位置</p>
	BoundingBox *BoundingBox `json:"BoundingBox,omitnil,omitempty" name:"BoundingBox"`

	// <p>行数</p>
	LinesCount *int64 `json:"LinesCount,omitnil,omitempty" name:"LinesCount"`

	// <p>行高</p><p>单位：px</p>
	LineHeight *int64 `json:"LineHeight,omitnil,omitempty" name:"LineHeight"`

	// <p>正常段落spam_code字段为0；如果存在spam_code字段且值大于0（1: 命中垃圾检查；2: 命中安全策略；3: 其他。），则命中安全检查被过滤。</p>
	SpamCode *int64 `json:"SpamCode,omitnil,omitempty" name:"SpamCode"`

	// <p>段落文本旋转信息，只在valid为true时表示坐标有效</p>
	RotateParagraphRect *RotateParagraphRect `json:"RotateParagraphRect,omitnil,omitempty" name:"RotateParagraphRect"`
}