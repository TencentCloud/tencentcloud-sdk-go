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

package v20190408

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AnalyzeSentimentRequestParams struct {
	// 待分析的文本（仅支持UTF-8格式，不超过200字）。
	Text *string `json:"Text,omitnil" name:"Text"`
}

type AnalyzeSentimentRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（仅支持UTF-8格式，不超过200字）。
	Text *string `json:"Text,omitnil" name:"Text"`
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
	Positive *float64 `json:"Positive,omitnil" name:"Positive"`

	// 中性情感概率。
	Neutral *float64 `json:"Neutral,omitnil" name:"Neutral"`

	// 负面情感概率。
	Negative *float64 `json:"Negative,omitnil" name:"Negative"`

	// 情感分类结果：
	// positive：正面情感
	// negative：负面情感
	// neutral：中性、无情感
	Sentiment *string `json:"Sentiment,omitnil" name:"Sentiment"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Word *string `json:"Word,omitnil" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitnil" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitnil" name:"Length"`

	// 词性。
	Pos *string `json:"Pos,omitnil" name:"Pos"`
}

type Category struct {
	// 分类id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil" name:"Id"`

	// 分类英文名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitnil" name:"Label"`

	// 分类中文名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 分类置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitnil" name:"Score"`
}

// Predefined struct for user
type ClassifyContentRequestParams struct {
	// 待分类的文章的标题（仅支持UTF-8格式，不超过100字符）。
	Title *string `json:"Title,omitnil" name:"Title"`

	// 待分类文章的内容, 每个元素对应一个段落。（仅支持UTF-8格式，文章内容长度总和不超过2000字符）
	Content []*string `json:"Content,omitnil" name:"Content"`
}

type ClassifyContentRequest struct {
	*tchttp.BaseRequest
	
	// 待分类的文章的标题（仅支持UTF-8格式，不超过100字符）。
	Title *string `json:"Title,omitnil" name:"Title"`

	// 待分类文章的内容, 每个元素对应一个段落。（仅支持UTF-8格式，文章内容长度总和不超过2000字符）
	Content []*string `json:"Content,omitnil" name:"Content"`
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
	FirstClassification *Category `json:"FirstClassification,omitnil" name:"FirstClassification"`

	// 二级分类。（请参见附录[三级分类体系表](https://cloud.tencent.com/document/product/271/94286)）
	SecondClassification *Category `json:"SecondClassification,omitnil" name:"SecondClassification"`

	// 三级分类。（请参见附录[三级分类体系表](https://cloud.tencent.com/document/product/271/94286)）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdClassification *Category `json:"ThirdClassification,omitnil" name:"ThirdClassification"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	Text *string `json:"Text,omitnil" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitnil" name:"TargetType"`
}

type ComposeCoupletRequest struct {
	*tchttp.BaseRequest
	
	// 生成对联的关键词。长度需>=2，当长度>2时，自动截取前两个字作为关键字。内容需为常用汉字（不含有数字、英文、韩语、日语、符号等等其他）。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitnil" name:"TargetType"`
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
	TopScroll *string `json:"TopScroll,omitnil" name:"TopScroll"`

	// 上联与下联。
	Content []*string `json:"Content,omitnil" name:"Content"`

	// 当对联随机生成时，展示随机生成原因。
	RandomCause *string `json:"RandomCause,omitnil" name:"RandomCause"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

// Predefined struct for user
type ComposePoetryRequestParams struct {
	// 生成诗词的关键词。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 生成诗词的类型。0：藏头或藏身；1：藏头；2：藏身。默认为0。
	PoetryType *int64 `json:"PoetryType,omitnil" name:"PoetryType"`

	// 诗的体裁。0：五言律诗或七言律诗；5：五言律诗；7：七言律诗。默认为0。
	Genre *int64 `json:"Genre,omitnil" name:"Genre"`
}

type ComposePoetryRequest struct {
	*tchttp.BaseRequest
	
	// 生成诗词的关键词。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 生成诗词的类型。0：藏头或藏身；1：藏头；2：藏身。默认为0。
	PoetryType *int64 `json:"PoetryType,omitnil" name:"PoetryType"`

	// 诗的体裁。0：五言律诗或七言律诗；5：五言律诗；7：七言律诗。默认为0。
	Genre *int64 `json:"Genre,omitnil" name:"Genre"`
}

func (r *ComposePoetryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ComposePoetryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "PoetryType")
	delete(f, "Genre")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ComposePoetryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ComposePoetryResponseParams struct {
	// 诗题，即输入的生成诗词的关键词。
	Title *string `json:"Title,omitnil" name:"Title"`

	// 诗的内容。
	Content []*string `json:"Content,omitnil" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ComposePoetryResponse struct {
	*tchttp.BaseResponse
	Response *ComposePoetryResponseParams `json:"Response"`
}

func (r *ComposePoetryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ComposePoetryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompoundParticiple struct {
	// 基础词。
	Word *string `json:"Word,omitnil" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitnil" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitnil" name:"Length"`

	// 词性。
	Pos *string `json:"Pos,omitnil" name:"Pos"`
}

type CorrectionItem struct {
	// 纠错句子的序号。
	Order *int64 `json:"Order,omitnil" name:"Order"`

	// 错误的起始位置，从0开始。
	BeginOffset *int64 `json:"BeginOffset,omitnil" name:"BeginOffset"`

	// 错误内容长度。
	Len *int64 `json:"Len,omitnil" name:"Len"`

	// 错误内容。
	Word *string `json:"Word,omitnil" name:"Word"`

	// 纠错结果，当为删除类错误时，结果为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorrectWord []*string `json:"CorrectWord,omitnil" name:"CorrectWord"`

	// 纠错类型。0：替换；1：插入；2：删除。
	CorrectionType *int64 `json:"CorrectionType,omitnil" name:"CorrectionType"`

	// 纠错信息置信度。0：error；1：warning；error的置信度更高。（仅供参考）
	Confidence *int64 `json:"Confidence,omitnil" name:"Confidence"`

	// 纠错信息中文描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionZh *string `json:"DescriptionZh,omitnil" name:"DescriptionZh"`

	// 纠错信息英文描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionEn *string `json:"DescriptionEn,omitnil" name:"DescriptionEn"`
}

type Embellish struct {
	// 润色后的文本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 润色类型。类型如下：
	// expansion：扩写
	// rewriting：改写
	// translation_m2a：从现代文改写为古文
	// translation_a2m：从古文改写为现代文
	// 
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmbellishType *string `json:"EmbellishType,omitnil" name:"EmbellishType"`
}

type Entity struct {
	// 基础词。
	Word *string `json:"Word,omitnil" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitnil" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitnil" name:"Length"`

	// 实体类型的标准名字。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 类型名字的自然语言表达。（中文或英文）
	Name *string `json:"Name,omitnil" name:"Name"`
}

// Predefined struct for user
type EvaluateSentenceSimilarityRequestParams struct {
	// 待分析的句子对数组。句子对应不超过1对，仅支持中文文本，原句子与目标句子均应不超过500字符。
	SentencePairList []*SentencePair `json:"SentencePairList,omitnil" name:"SentencePairList"`
}

type EvaluateSentenceSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的句子对数组。句子对应不超过1对，仅支持中文文本，原句子与目标句子均应不超过500字符。
	SentencePairList []*SentencePair `json:"SentencePairList,omitnil" name:"SentencePairList"`
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
	ScoreList []*float64 `json:"ScoreList,omitnil" name:"ScoreList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type EvaluateWordSimilarityRequestParams struct {
	// 计算相似度的源词。（仅支持UTF-8格式，不超过10字符）
	SourceWord *string `json:"SourceWord,omitnil" name:"SourceWord"`

	// 计算相似度的目标词。（仅支持UTF-8格式，不超过10字符）
	TargetWord *string `json:"TargetWord,omitnil" name:"TargetWord"`
}

type EvaluateWordSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// 计算相似度的源词。（仅支持UTF-8格式，不超过10字符）
	SourceWord *string `json:"SourceWord,omitnil" name:"SourceWord"`

	// 计算相似度的目标词。（仅支持UTF-8格式，不超过10字符）
	TargetWord *string `json:"TargetWord,omitnil" name:"TargetWord"`
}

func (r *EvaluateWordSimilarityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluateWordSimilarityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceWord")
	delete(f, "TargetWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EvaluateWordSimilarityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EvaluateWordSimilarityResponseParams struct {
	// 词相似度分值。
	Similarity *float64 `json:"Similarity,omitnil" name:"Similarity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type EvaluateWordSimilarityResponse struct {
	*tchttp.BaseResponse
	Response *EvaluateWordSimilarityResponseParams `json:"Response"`
}

func (r *EvaluateWordSimilarityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EvaluateWordSimilarityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateKeywordSentenceRequestParams struct {
	// 生成句子的关键词，关键词个数需不超过4个，中文关键词长度应不超过10字符，英文关键词长度不超过3个单词。关键词中不可包含标点符号。
	WordList []*string `json:"WordList,omitnil" name:"WordList"`

	// 返回生成句子的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitnil" name:"Number"`

	// 指定生成句子的领域，支持领域如下：
	// general：通用领域，支持中英文
	// academic：学术领域，仅支持英文
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

type GenerateKeywordSentenceRequest struct {
	*tchttp.BaseRequest
	
	// 生成句子的关键词，关键词个数需不超过4个，中文关键词长度应不超过10字符，英文关键词长度不超过3个单词。关键词中不可包含标点符号。
	WordList []*string `json:"WordList,omitnil" name:"WordList"`

	// 返回生成句子的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitnil" name:"Number"`

	// 指定生成句子的领域，支持领域如下：
	// general：通用领域，支持中英文
	// academic：学术领域，仅支持英文
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitnil" name:"Domain"`
}

func (r *GenerateKeywordSentenceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateKeywordSentenceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WordList")
	delete(f, "Number")
	delete(f, "Domain")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateKeywordSentenceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateKeywordSentenceResponseParams struct {
	// 生成的句子列表。
	KeywordSentenceList []*KeywordSentence `json:"KeywordSentenceList,omitnil" name:"KeywordSentenceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type GenerateKeywordSentenceResponse struct {
	*tchttp.BaseResponse
	Response *GenerateKeywordSentenceResponseParams `json:"Response"`
}

func (r *GenerateKeywordSentenceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateKeywordSentenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeywordSentence struct {
	// 通过关键词生成的句子。
	TargetText *string `json:"TargetText,omitnil" name:"TargetText"`
}

// Predefined struct for user
type ParseWordsRequestParams struct {
	// 待分析的文本（支持中英文文本，不超过500字符）
	Text *string `json:"Text,omitnil" name:"Text"`
}

type ParseWordsRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（支持中英文文本，不超过500字符）
	Text *string `json:"Text,omitnil" name:"Text"`
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
	NormalText *string `json:"NormalText,omitnil" name:"NormalText"`

	// 基础粒度分词和词性标注的结果。（请参见附录[词性表](https://cloud.tencent.com/document/product/271/36460)）
	BasicParticiples []*BasicParticiple `json:"BasicParticiples,omitnil" name:"BasicParticiples"`

	// 复合粒度分词和词性标注的结果。（请参见附录[词性表](https://cloud.tencent.com/document/product/271/36460)）
	CompoundParticiples []*CompoundParticiple `json:"CompoundParticiples,omitnil" name:"CompoundParticiples"`

	// 实体识别结果。（请参见附录[实体类型数据](https://cloud.tencent.com/document/product/271/90592)）
	Entities []*Entity `json:"Entities,omitnil" name:"Entities"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type RetrieveSimilarWordsRequestParams struct {
	// 输入的词语。（仅支持UTF-8格式，不超过10字符）
	Text *string `json:"Text,omitnil" name:"Text"`

	// 召回的相似词个数，取值范围为1-20。
	Number *int64 `json:"Number,omitnil" name:"Number"`
}

type RetrieveSimilarWordsRequest struct {
	*tchttp.BaseRequest
	
	// 输入的词语。（仅支持UTF-8格式，不超过10字符）
	Text *string `json:"Text,omitnil" name:"Text"`

	// 召回的相似词个数，取值范围为1-20。
	Number *int64 `json:"Number,omitnil" name:"Number"`
}

func (r *RetrieveSimilarWordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetrieveSimilarWordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetrieveSimilarWordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetrieveSimilarWordsResponseParams struct {
	// 召回的相似词数组。
	WordList []*string `json:"WordList,omitnil" name:"WordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RetrieveSimilarWordsResponse struct {
	*tchttp.BaseResponse
	Response *RetrieveSimilarWordsResponseParams `json:"Response"`
}

func (r *RetrieveSimilarWordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetrieveSimilarWordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceCorrectionRequestParams struct {
	// 待纠错的句子列表。可以以数组方式在一次请求中填写多个待纠错的句子。文本统一使用utf-8格式编码，每个中文句子的长度不超过150字符，每个英文句子的长度不超过100个单词，且数组长度需小于30，即句子总数需少于30句。
	TextList []*string `json:"TextList,omitnil" name:"TextList"`
}

type SentenceCorrectionRequest struct {
	*tchttp.BaseRequest
	
	// 待纠错的句子列表。可以以数组方式在一次请求中填写多个待纠错的句子。文本统一使用utf-8格式编码，每个中文句子的长度不超过150字符，每个英文句子的长度不超过100个单词，且数组长度需小于30，即句子总数需少于30句。
	TextList []*string `json:"TextList,omitnil" name:"TextList"`
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
	CorrectionList []*CorrectionItem `json:"CorrectionList,omitnil" name:"CorrectionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	SourceText *string `json:"SourceText,omitnil" name:"SourceText"`

	// 目标句子。（仅支持UTF-8格式，不超过500字符）
	TargetText *string `json:"TargetText,omitnil" name:"TargetText"`
}

// Predefined struct for user
type TestingTextGenerationRequestParams struct {
	// 会话内容,按对话时间从旧到新在数组中排列。
	Messages []*TextGenerationMessage `json:"Messages,omitnil" name:"Messages"`

	// 模型名称，当前固定为TestingModel
	Model *string `json:"Model,omitnil" name:"Model"`

	// 会话id。
	QueryId *string `json:"QueryId,omitnil" name:"QueryId"`

	// 超参数temperature, 该参数用于控制生成文本中重复内容。取值区间为[0.0, 2.0], 非必要不建议使用, 不合理的取值会影响效果。默认为1.0。
	Temperature *float64 `json:"Temperature,omitnil" name:"Temperature"`

	// 超参数top_p, 该参数用于控制生成文本的多样性。较小的"top_p"值会限制模型的选择范围，使生成的文本更加一致和确定性。较大的"top_p"值会扩大选择范围，使生成的文本更加多样化和随机。取值区间为[0.0, 1.0], 非必要不建议使用, 不合理的取值会影响效果。默认值为1.0。
	TopP *float64 `json:"TopP,omitnil" name:"TopP"`

	// 超参数top_k,该参数用于控制生成文本的多样性和可控性，较小的"top_k"值会限制模型的选择范围，使生成的文本更加一致和确定性。较大的"top_k"值会扩大选择范围，使生成的文本更加多样化和随机。取值区间为[1, 100]，默认值 40。
	TopK *float64 `json:"TopK,omitnil" name:"TopK"`

	// 重复惩罚项,该参数用于用于控制生成文本中重复内容。建议范围[1.0, 1.05]非必要不建议使用, 不合理的取值会影响效果。默认为1。
	RepetitionPenalty *float64 `json:"RepetitionPenalty,omitnil" name:"RepetitionPenalty"`

	// 输出结果最大tokens数量。取值区间为(0, 1024]。默认值为768。
	OutputSeqLen *int64 `json:"OutputSeqLen,omitnil" name:"OutputSeqLen"`

	// 输入文本的最大 tokens 数量。取值区间为(0, 1024]。默认值为256。
	MaxInputSeqLen *int64 `json:"MaxInputSeqLen,omitnil" name:"MaxInputSeqLen"`
}

type TestingTextGenerationRequest struct {
	*tchttp.BaseRequest
	
	// 会话内容,按对话时间从旧到新在数组中排列。
	Messages []*TextGenerationMessage `json:"Messages,omitnil" name:"Messages"`

	// 模型名称，当前固定为TestingModel
	Model *string `json:"Model,omitnil" name:"Model"`

	// 会话id。
	QueryId *string `json:"QueryId,omitnil" name:"QueryId"`

	// 超参数temperature, 该参数用于控制生成文本中重复内容。取值区间为[0.0, 2.0], 非必要不建议使用, 不合理的取值会影响效果。默认为1.0。
	Temperature *float64 `json:"Temperature,omitnil" name:"Temperature"`

	// 超参数top_p, 该参数用于控制生成文本的多样性。较小的"top_p"值会限制模型的选择范围，使生成的文本更加一致和确定性。较大的"top_p"值会扩大选择范围，使生成的文本更加多样化和随机。取值区间为[0.0, 1.0], 非必要不建议使用, 不合理的取值会影响效果。默认值为1.0。
	TopP *float64 `json:"TopP,omitnil" name:"TopP"`

	// 超参数top_k,该参数用于控制生成文本的多样性和可控性，较小的"top_k"值会限制模型的选择范围，使生成的文本更加一致和确定性。较大的"top_k"值会扩大选择范围，使生成的文本更加多样化和随机。取值区间为[1, 100]，默认值 40。
	TopK *float64 `json:"TopK,omitnil" name:"TopK"`

	// 重复惩罚项,该参数用于用于控制生成文本中重复内容。建议范围[1.0, 1.05]非必要不建议使用, 不合理的取值会影响效果。默认为1。
	RepetitionPenalty *float64 `json:"RepetitionPenalty,omitnil" name:"RepetitionPenalty"`

	// 输出结果最大tokens数量。取值区间为(0, 1024]。默认值为768。
	OutputSeqLen *int64 `json:"OutputSeqLen,omitnil" name:"OutputSeqLen"`

	// 输入文本的最大 tokens 数量。取值区间为(0, 1024]。默认值为256。
	MaxInputSeqLen *int64 `json:"MaxInputSeqLen,omitnil" name:"MaxInputSeqLen"`
}

func (r *TestingTextGenerationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TestingTextGenerationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Messages")
	delete(f, "Model")
	delete(f, "QueryId")
	delete(f, "Temperature")
	delete(f, "TopP")
	delete(f, "TopK")
	delete(f, "RepetitionPenalty")
	delete(f, "OutputSeqLen")
	delete(f, "MaxInputSeqLen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TestingTextGenerationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TestingTextGenerationResponseParams struct {
	// 结果
	Choices []*TextGenerationChoices `json:"Choices,omitnil" name:"Choices"`

	// unix时间戳的字符串
	Created *int64 `json:"Created,omitnil" name:"Created"`

	// 会话id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 模型名
	Model *string `json:"Model,omitnil" name:"Model"`

	// token数量
	Usage *TextGenerationUsage `json:"Usage,omitnil" name:"Usage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TestingTextGenerationResponse struct {
	*tchttp.BaseResponse
	Response *TestingTextGenerationResponseParams `json:"Response"`
}

func (r *TestingTextGenerationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TestingTextGenerationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextEmbellishRequestParams struct {
	// 待润色的文本。中文文本长度需<=50字符；英文文本长度需<=30个单词。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 待润色文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitnil" name:"SourceLang"`

	// 返回润色结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitnil" name:"Number"`

	// 控制润色类型，类型如下：
	// both：同时返回改写和扩写
	// expansion：扩写
	// rewriting：改写
	// m2a：从现代文改写为古文
	// a2m：从古文改写为现代文
	// 默认为both。
	Style *string `json:"Style,omitnil" name:"Style"`
}

type TextEmbellishRequest struct {
	*tchttp.BaseRequest
	
	// 待润色的文本。中文文本长度需<=50字符；英文文本长度需<=30个单词。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 待润色文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitnil" name:"SourceLang"`

	// 返回润色结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitnil" name:"Number"`

	// 控制润色类型，类型如下：
	// both：同时返回改写和扩写
	// expansion：扩写
	// rewriting：改写
	// m2a：从现代文改写为古文
	// a2m：从古文改写为现代文
	// 默认为both。
	Style *string `json:"Style,omitnil" name:"Style"`
}

func (r *TextEmbellishRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextEmbellishRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "SourceLang")
	delete(f, "Number")
	delete(f, "Style")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextEmbellishRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextEmbellishResponseParams struct {
	// 润色结果列表。
	EmbellishList []*Embellish `json:"EmbellishList,omitnil" name:"EmbellishList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextEmbellishResponse struct {
	*tchttp.BaseResponse
	Response *TextEmbellishResponseParams `json:"Response"`
}

func (r *TextEmbellishResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextEmbellishResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TextGenerationChoices struct {
	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *TextGenerationMessage `json:"Message,omitnil" name:"Message"`
}

type TextGenerationMessage struct {
	// 角色支持 system, user, assistant。默认为user。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil" name:"Role"`

	// 消息的内容。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil" name:"Content"`
}

type TextGenerationUsage struct {
	// 输入tokens数量
	PromptTokens *int64 `json:"PromptTokens,omitnil" name:"PromptTokens"`

	// 输出tokens数量
	CompletionTokens *int64 `json:"CompletionTokens,omitnil" name:"CompletionTokens"`

	// 总token数量
	TotalTokens *int64 `json:"TotalTokens,omitnil" name:"TotalTokens"`
}

// Predefined struct for user
type TextWritingRequestParams struct {
	// 待续写的句子，文本统一使用utf-8格式编码，长度不超过200字符。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 待续写文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitnil" name:"SourceLang"`

	// 返回续写结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitnil" name:"Number"`

	// 指定续写领域，支持领域如下：
	// general：通用领域，支持中英文补全
	// academic：学术领域，仅支持英文补全
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定续写风格，支持风格如下：
	// science_fiction：科幻
	// military_history：军事
	// xuanhuan_wuxia：武侠
	// urban_officialdom：职场
	// 默认为xuanhuan_wuxia（武侠）。
	Style *string `json:"Style,omitnil" name:"Style"`
}

type TextWritingRequest struct {
	*tchttp.BaseRequest
	
	// 待续写的句子，文本统一使用utf-8格式编码，长度不超过200字符。
	Text *string `json:"Text,omitnil" name:"Text"`

	// 待续写文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitnil" name:"SourceLang"`

	// 返回续写结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitnil" name:"Number"`

	// 指定续写领域，支持领域如下：
	// general：通用领域，支持中英文补全
	// academic：学术领域，仅支持英文补全
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitnil" name:"Domain"`

	// 指定续写风格，支持风格如下：
	// science_fiction：科幻
	// military_history：军事
	// xuanhuan_wuxia：武侠
	// urban_officialdom：职场
	// 默认为xuanhuan_wuxia（武侠）。
	Style *string `json:"Style,omitnil" name:"Style"`
}

func (r *TextWritingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextWritingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "SourceLang")
	delete(f, "Number")
	delete(f, "Domain")
	delete(f, "Style")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextWritingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextWritingResponseParams struct {
	// 续写结果列表。
	WritingList []*Writing `json:"WritingList,omitnil" name:"WritingList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type TextWritingResponse struct {
	*tchttp.BaseResponse
	Response *TextWritingResponseParams `json:"Response"`
}

func (r *TextWritingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextWritingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Writing struct {
	// 续写的文本。
	TargetText *string `json:"TargetText,omitnil" name:"TargetText"`

	// 续写的前缀。
	PrefixText *string `json:"PrefixText,omitnil" name:"PrefixText"`
}