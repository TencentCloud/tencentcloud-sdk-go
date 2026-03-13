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
	// 左上顶点x坐标
	X *int64 `json:"X,omitnil,omitempty" name:"X"`

	// 左上顶点y坐标
	Y *int64 `json:"Y,omitnil,omitempty" name:"Y"`

	// 宽
	Width *int64 `json:"Width,omitnil,omitempty" name:"Width"`

	// 高
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
}

type ImageTranslateLLMRequest struct {
	*tchttp.BaseRequest
	
	// <p>图片数据的Base64字符串，经Base64编码后不超过 9M，分辨率建议600*800以上，支持PNG、JPG、JPEG格式。</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// <p>目标语言，支持语言列表：</p><ul><li>中文：zh</li><li>繁体（中国台湾）：zh-TW</li><li>繁体（中国香港）：zh-HK</li><li>英文：en</li><li>日语：ja</li><li>韩语：ko</li><li>泰语：th</li><li>越南语：vi</li><li>俄语：ru</li><li>德语：de</li><li>法语：fr</li><li>阿拉伯语：ar</li><li>西班牙语：es</li><li>意大利语：it</li><li>印度尼西亚语：id</li><li>马来西亚语：ms</li><li>葡萄牙语：pt</li><li>土耳其语：tr<br>-</li></ul>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// <p>输入图 Url。 使用Url的时候，Data参数需要传入&quot;&quot;。 图片限制：小于 10MB，分辨率建议600*800以上，格式支持 jpg、jpeg、png。</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
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

// Predefined struct for user
type TextTranslateRequestParams struct {
	// 待翻译的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度需要低于6000字符。
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 源语言，支持：
	// auto：自动识别（识别为一种语言）
	// zh：简体中文
	// zh-TW：繁体中文
	// en：英语
	// ja：日语
	// ko：韩语
	// fr：法语
	// es：西班牙语
	// it：意大利语
	// de：德语
	// tr：土耳其语
	// ru：俄语
	// pt：葡萄牙语
	// vi：越南语
	// id：印尼语
	// th：泰语
	// ms：马来西亚语
	// ar：阿拉伯语
	// hi：印地语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：zh-TW（繁体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>zh-TW（繁体中文）：zh（简体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>en（英语）：zh（中文）、zh-TW（繁体中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、zh-TW（繁体中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用来标记不希望被翻译的文本内容，如句子中的特殊符号、人名、地名等；每次请求只支持配置一个不被翻译的单词；仅支持配置人名、地名等名词，不要配置动词或短语，否则会影响翻译结果。
	UntranslatedText *string `json:"UntranslatedText,omitnil,omitempty" name:"UntranslatedText"`

	// 需要使用的术语库列表，通过 [术语库操作指南](https://cloud.tencent.com/document/product/551/107926) 自行创建术语库获取。
	TermRepoIDList []*string `json:"TermRepoIDList,omitnil,omitempty" name:"TermRepoIDList"`

	// 需要使用的例句库列表，通过 [例句库操作指南](https://cloud.tencent.com/document/product/551/107927) 自行创建例句库获取。
	SentRepoIDList []*string `json:"SentRepoIDList,omitnil,omitempty" name:"SentRepoIDList"`
}

type TextTranslateRequest struct {
	*tchttp.BaseRequest
	
	// 待翻译的文本，文本统一使用utf-8格式编码，非utf-8格式编码字符会翻译失败，请传入有效文本，html标记等非常规翻译文本可能会翻译失败。单次请求的文本长度需要低于6000字符。
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 源语言，支持：
	// auto：自动识别（识别为一种语言）
	// zh：简体中文
	// zh-TW：繁体中文
	// en：英语
	// ja：日语
	// ko：韩语
	// fr：法语
	// es：西班牙语
	// it：意大利语
	// de：德语
	// tr：土耳其语
	// ru：俄语
	// pt：葡萄牙语
	// vi：越南语
	// id：印尼语
	// th：泰语
	// ms：马来西亚语
	// ar：阿拉伯语
	// hi：印地语
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，各源语言的目标语言支持列表如下
	// 
	// <li> zh（简体中文）：zh-TW（繁体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>zh-TW（繁体中文）：zh（简体中文）、en（英语）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）</li>
	// <li>en（英语）：zh（中文）、zh-TW（繁体中文）、ja（日语）、ko（韩语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）、vi（越南语）、id（印尼语）、th（泰语）、ms（马来语）、ar（阿拉伯语）、hi（印地语）</li>
	// <li>ja（日语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ko（韩语）</li>
	// <li>ko（韩语）：zh（中文）、zh-TW（繁体中文）、en（英语）、ja（日语）</li>
	// <li>fr（法语）：zh（中文）、zh-TW（繁体中文）、en（英语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>es（西班牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>it（意大利语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、de（德语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>de（德语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、tr（土耳其语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>tr（土耳其语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、ru（俄语）、pt（葡萄牙语）</li>
	// <li>ru（俄语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、pt（葡萄牙语）</li>
	// <li>pt（葡萄牙语）：zh（中文）、zh-TW（繁体中文）、en（英语）、fr（法语）、es（西班牙语）、it（意大利语）、de（德语）、tr（土耳其语）、ru（俄语）</li>
	// <li>vi（越南语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>id（印尼语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>th（泰语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ms（马来语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>ar（阿拉伯语）：zh（中文）、zh-TW（繁体中文）、en（英语）</li>
	// <li>hi（印地语）：en（英语）</li>
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 项目ID，可以根据控制台-账号中心-项目管理中的配置填写，如无配置请填写默认项目ID:0
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用来标记不希望被翻译的文本内容，如句子中的特殊符号、人名、地名等；每次请求只支持配置一个不被翻译的单词；仅支持配置人名、地名等名词，不要配置动词或短语，否则会影响翻译结果。
	UntranslatedText *string `json:"UntranslatedText,omitnil,omitempty" name:"UntranslatedText"`

	// 需要使用的术语库列表，通过 [术语库操作指南](https://cloud.tencent.com/document/product/551/107926) 自行创建术语库获取。
	TermRepoIDList []*string `json:"TermRepoIDList,omitnil,omitempty" name:"TermRepoIDList"`

	// 需要使用的例句库列表，通过 [例句库操作指南](https://cloud.tencent.com/document/product/551/107927) 自行创建例句库获取。
	SentRepoIDList []*string `json:"SentRepoIDList,omitnil,omitempty" name:"SentRepoIDList"`
}

func (r *TextTranslateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceText")
	delete(f, "Source")
	delete(f, "Target")
	delete(f, "ProjectId")
	delete(f, "UntranslatedText")
	delete(f, "TermRepoIDList")
	delete(f, "SentRepoIDList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextTranslateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextTranslateResponseParams struct {
	// 翻译后的文本
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`

	// 源语言，详见入参Source
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 目标语言，详见入参Target
	Target *string `json:"Target,omitnil,omitempty" name:"Target"`

	// 本次翻译消耗的字符数
	UsedAmount *int64 `json:"UsedAmount,omitnil,omitempty" name:"UsedAmount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TextTranslateResponse struct {
	*tchttp.BaseResponse
	Response *TextTranslateResponseParams `json:"Response"`
}

func (r *TextTranslateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextTranslateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TransDetail struct {
	// 当前行的原文本
	SourceLineText *string `json:"SourceLineText,omitnil,omitempty" name:"SourceLineText"`

	// 当前行的译文
	TargetLineText *string `json:"TargetLineText,omitnil,omitempty" name:"TargetLineText"`

	// 段落文本框位置
	BoundingBox *BoundingBox `json:"BoundingBox,omitnil,omitempty" name:"BoundingBox"`

	// 行数
	LinesCount *int64 `json:"LinesCount,omitnil,omitempty" name:"LinesCount"`

	// 行高
	LineHeight *int64 `json:"LineHeight,omitnil,omitempty" name:"LineHeight"`

	// 正常段落spam_code字段为0；如果存在spam_code字段且值大于0（1: 命中垃圾检查；2: 命中安全策略；3: 其他。），则命中安全检查被过滤。
	SpamCode *int64 `json:"SpamCode,omitnil,omitempty" name:"SpamCode"`

	// 段落文本旋转信息，只在valid为true时表示坐标有效
	RotateParagraphRect *RotateParagraphRect `json:"RotateParagraphRect,omitnil,omitempty" name:"RotateParagraphRect"`
}