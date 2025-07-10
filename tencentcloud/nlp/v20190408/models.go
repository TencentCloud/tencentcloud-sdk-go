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

package v20190408

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AnalyzeSentimentRequestParams struct {
	// 待分析的文本（仅支持UTF-8格式，不超过200字）。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type AnalyzeSentimentRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（仅支持UTF-8格式，不超过200字）。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

func (r *AnalyzeSentimentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeSentimentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AnalyzeSentimentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AnalyzeSentimentResponseParams struct {
	// 正面情感概率。
	Positive *float64 `json:"Positive,omitnil,omitempty" name:"Positive"`

	// 中性情感概率。
	Neutral *float64 `json:"Neutral,omitnil,omitempty" name:"Neutral"`

	// 负面情感概率。
	Negative *float64 `json:"Negative,omitnil,omitempty" name:"Negative"`

	// 情感分类结果：
	// positive：正面情感
	// negative：负面情感
	// neutral：中性、无情感
	Sentiment *string `json:"Sentiment,omitnil,omitempty" name:"Sentiment"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AnalyzeSentimentResponse struct {
	*tchttp.BaseResponse
	Response *AnalyzeSentimentResponseParams `json:"Response"`
}

func (r *AnalyzeSentimentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AnalyzeSentimentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicParticiple struct {
	// 基础词。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitnil,omitempty" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 词性。
	Pos *string `json:"Pos,omitnil,omitempty" name:"Pos"`
}

type Category struct {
	// 分类id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 分类英文名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// 分类中文名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`
}

// Predefined struct for user
type ClassifyContentRequestParams struct {
	// 待分类的文章的标题（仅支持UTF-8格式，不超过100字符）。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 待分类文章的内容, 每个元素对应一个段落。（仅支持UTF-8格式，文章内容长度总和不超过2000字符）
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`
}

type ClassifyContentRequest struct {
	*tchttp.BaseRequest
	
	// 待分类的文章的标题（仅支持UTF-8格式，不超过100字符）。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 待分类文章的内容, 每个元素对应一个段落。（仅支持UTF-8格式，文章内容长度总和不超过2000字符）
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`
}

func (r *ClassifyContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClassifyContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Title")
	delete(f, "Content")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClassifyContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClassifyContentResponseParams struct {
	// 一级分类。（请参见附录[三级分类体系表](https://cloud.tencent.com/document/product/271/94286)）
	FirstClassification *Category `json:"FirstClassification,omitnil,omitempty" name:"FirstClassification"`

	// 二级分类。（请参见附录[三级分类体系表](https://cloud.tencent.com/document/product/271/94286)）
	SecondClassification *Category `json:"SecondClassification,omitnil,omitempty" name:"SecondClassification"`

	// 三级分类。（请参见附录[三级分类体系表](https://cloud.tencent.com/document/product/271/94286)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdClassification *Category `json:"ThirdClassification,omitnil,omitempty" name:"ThirdClassification"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClassifyContentResponse struct {
	*tchttp.BaseResponse
	Response *ClassifyContentResponseParams `json:"Response"`
}

func (r *ClassifyContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClassifyContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ComposeCoupletRequestParams struct {
	// 生成对联的关键词。长度需>=2，当长度>2时，自动截取前两个字作为关键字。内容需为常用汉字（不含有数字、英文、韩语、日语、符号等等其他）。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

type ComposeCoupletRequest struct {
	*tchttp.BaseRequest
	
	// 生成对联的关键词。长度需>=2，当长度>2时，自动截取前两个字作为关键字。内容需为常用汉字（不含有数字、英文、韩语、日语、符号等等其他）。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitnil,omitempty" name:"TargetType"`
}

func (r *ComposeCoupletRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ComposeCoupletRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "TargetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ComposeCoupletRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ComposeCoupletResponseParams struct {
	// 横批。
	TopScroll *string `json:"TopScroll,omitnil,omitempty" name:"TopScroll"`

	// 上联与下联。
	Content []*string `json:"Content,omitnil,omitempty" name:"Content"`

	// 当对联随机生成时，展示随机生成原因。
	RandomCause *string `json:"RandomCause,omitnil,omitempty" name:"RandomCause"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ComposeCoupletResponse struct {
	*tchttp.BaseResponse
	Response *ComposeCoupletResponseParams `json:"Response"`
}

func (r *ComposeCoupletResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ComposeCoupletResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompoundParticiple struct {
	// 基础词。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitnil,omitempty" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 词性。
	Pos *string `json:"Pos,omitnil,omitempty" name:"Pos"`
}

type CorrectionItem struct {
	// 纠错句子的序号。
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`

	// 错误的起始位置，从0开始。
	BeginOffset *int64 `json:"BeginOffset,omitnil,omitempty" name:"BeginOffset"`

	// 错误内容长度。
	Len *int64 `json:"Len,omitnil,omitempty" name:"Len"`

	// 错误内容。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 纠错结果，当为删除类错误时，结果为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorrectWord []*string `json:"CorrectWord,omitnil,omitempty" name:"CorrectWord"`

	// 纠错类型。0：替换；1：插入；2：删除。
	CorrectionType *int64 `json:"CorrectionType,omitnil,omitempty" name:"CorrectionType"`

	// 纠错信息置信度。0：error；1：warning；error的置信度更高。（仅供参考）
	Confidence *int64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 纠错信息中文描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionZh *string `json:"DescriptionZh,omitnil,omitempty" name:"DescriptionZh"`

	// 纠错信息英文描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionEn *string `json:"DescriptionEn,omitnil,omitempty" name:"DescriptionEn"`
}

type Entity struct {
	// 基础词。
	Word *string `json:"Word,omitnil,omitempty" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitnil,omitempty" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 实体类型的标准名字。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 类型名字的自然语言表达。（中文或英文）
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

// Predefined struct for user
type EvaluateSentenceSimilarityRequestParams struct {
	// 待分析的句子对数组。句子对应不超过1对，仅支持中文文本，原句子与目标句子均应不超过500字符。
	SentencePairList []*SentencePair `json:"SentencePairList,omitnil,omitempty" name:"SentencePairList"`
}

type EvaluateSentenceSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的句子对数组。句子对应不超过1对，仅支持中文文本，原句子与目标句子均应不超过500字符。
	SentencePairList []*SentencePair `json:"SentencePairList,omitnil,omitempty" name:"SentencePairList"`
}

func (r *EvaluateSentenceSimilarityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluateSentenceSimilarityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SentencePairList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EvaluateSentenceSimilarityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EvaluateSentenceSimilarityResponseParams struct {
	// 每个句子对的相似度分值。
	ScoreList []*float64 `json:"ScoreList,omitnil,omitempty" name:"ScoreList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EvaluateSentenceSimilarityResponse struct {
	*tchttp.BaseResponse
	Response *EvaluateSentenceSimilarityResponseParams `json:"Response"`
}

func (r *EvaluateSentenceSimilarityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluateSentenceSimilarityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseWordsRequestParams struct {
	// 待分析的文本（支持中英文文本，不超过500字符）
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type ParseWordsRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（支持中英文文本，不超过500字符）
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

func (r *ParseWordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseWordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ParseWordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ParseWordsResponseParams struct {
	// 输入文本正则化的结果。（包括对英文文本中的开头和实体进行大写等）
	NormalText *string `json:"NormalText,omitnil,omitempty" name:"NormalText"`

	// 基础粒度分词和词性标注的结果。（请参见附录[词性表](https://cloud.tencent.com/document/product/271/36460)）
	BasicParticiples []*BasicParticiple `json:"BasicParticiples,omitnil,omitempty" name:"BasicParticiples"`

	// 复合粒度分词和词性标注的结果。（请参见附录[词性表](https://cloud.tencent.com/document/product/271/36460)）
	CompoundParticiples []*CompoundParticiple `json:"CompoundParticiples,omitnil,omitempty" name:"CompoundParticiples"`

	// 实体识别结果。（请参见附录[实体类型数据](https://cloud.tencent.com/document/product/271/90592)）
	Entities []*Entity `json:"Entities,omitnil,omitempty" name:"Entities"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ParseWordsResponse struct {
	*tchttp.BaseResponse
	Response *ParseWordsResponseParams `json:"Response"`
}

func (r *ParseWordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ParseWordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceCorrectionRequestParams struct {
	// 待纠错的句子列表。可以以数组方式在一次请求中填写多个待纠错的句子。文本统一使用utf-8格式编码，每个中文句子的长度不超过150字符，每个英文句子的长度不超过100个单词，且数组长度需小于30，即句子总数需少于30句。
	TextList []*string `json:"TextList,omitnil,omitempty" name:"TextList"`
}

type SentenceCorrectionRequest struct {
	*tchttp.BaseRequest
	
	// 待纠错的句子列表。可以以数组方式在一次请求中填写多个待纠错的句子。文本统一使用utf-8格式编码，每个中文句子的长度不超过150字符，每个英文句子的长度不超过100个单词，且数组长度需小于30，即句子总数需少于30句。
	TextList []*string `json:"TextList,omitnil,omitempty" name:"TextList"`
}

func (r *SentenceCorrectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceCorrectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TextList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SentenceCorrectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceCorrectionResponseParams struct {
	// 纠错结果列表。
	// （注意仅展示错误句子的纠错结果，若句子无错则不展示，若全部待纠错句子都被认为无错，则可能返回数组为空）
	CorrectionList []*CorrectionItem `json:"CorrectionList,omitnil,omitempty" name:"CorrectionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SentenceCorrectionResponse struct {
	*tchttp.BaseResponse
	Response *SentenceCorrectionResponseParams `json:"Response"`
}

func (r *SentenceCorrectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceCorrectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SentencePair struct {
	// 需要与目标句子计算相似度的源句子。（仅支持UTF-8格式，不超过500字符）
	SourceText *string `json:"SourceText,omitnil,omitempty" name:"SourceText"`

	// 目标句子。（仅支持UTF-8格式，不超过500字符）
	TargetText *string `json:"TargetText,omitnil,omitempty" name:"TargetText"`
}