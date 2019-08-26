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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AutoSummarizationRequest struct {
	*tchttp.BaseRequest

	// 待处理的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 指定摘要的长度（默认值为200）
	// 注：为保证摘要的可读性，最终生成的摘要长度并不会严格遵循这个值，会有略微的浮动
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

func (r *AutoSummarizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AutoSummarizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AutoSummarizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文本摘要结果
		Summary *string `json:"Summary,omitempty" name:"Summary"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AutoSummarizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AutoSummarizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CCIToken struct {

	// 错别字的起始位置，从0开始
	BeginOffset *uint64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 错别字纠错结果
	CorrectWord *string `json:"CorrectWord,omitempty" name:"CorrectWord"`

	// 错别字内容
	Word *string `json:"Word,omitempty" name:"Word"`
}

type ClassificationResult struct {

	// 一级分类名称
	FirstClassName *string `json:"FirstClassName,omitempty" name:"FirstClassName"`

	// 一级分类概率
	FirstClassProbability *float64 `json:"FirstClassProbability,omitempty" name:"FirstClassProbability"`

	// 二级分类名称
	SecondClassName *string `json:"SecondClassName,omitempty" name:"SecondClassName"`

	// 二级分类概率
	SecondClassProbability *float64 `json:"SecondClassProbability,omitempty" name:"SecondClassProbability"`
}

type ContentApprovalRequest struct {
	*tchttp.BaseRequest

	// 待审核的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *ContentApprovalRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ContentApprovalRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ContentApprovalResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文本是否恶意：
	// 0、正常；
	// 1、恶意；
	// 2、可疑送审
		EvilFlag *uint64 `json:"EvilFlag,omitempty" name:"EvilFlag"`

		// 恶意关键词组
		EvilKeywords []*string `json:"EvilKeywords,omitempty" name:"EvilKeywords" list`

		// 文本恶意类型：
	// 0、正常；
	// 1、政治；
	// 2、色情；
	// 3、辱骂/低俗；
	// 4、暴恐/毒品；
	// 5、广告/灌水；
	// 6、迷信/邪教；
	// 7、其他违法（如跨站追杀/恶意竞争等）；
	// 8、综合
		EvilType *uint64 `json:"EvilType,omitempty" name:"EvilType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ContentApprovalResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ContentApprovalResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DependencyParsingRequest struct {
	*tchttp.BaseRequest

	// 待分析的文本（仅支持UTF-8格式，不超过200字）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *DependencyParsingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DependencyParsingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DependencyParsingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 句法依存分析结果
		DpTokens []*DpToken `json:"DpTokens,omitempty" name:"DpTokens" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DependencyParsingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DependencyParsingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DpToken struct {

	// 当前词父节点的序号
	HeadId *uint64 `json:"HeadId,omitempty" name:"HeadId"`

	// 基础词的序号
	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// 句法依存关系的类型
	Relation *string `json:"Relation,omitempty" name:"Relation"`

	// 基础词
	Word *string `json:"Word,omitempty" name:"Word"`
}

type Keyword struct {

	// 权重
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 关键词
	Word *string `json:"Word,omitempty" name:"Word"`
}

type KeywordsExtractionRequest struct {
	*tchttp.BaseRequest

	// 待处理的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *KeywordsExtractionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *KeywordsExtractionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type KeywordsExtractionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 关键词提取结果
		Keywords []*Keyword `json:"Keywords,omitempty" name:"Keywords" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *KeywordsExtractionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *KeywordsExtractionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LexicalAnalysisRequest struct {
	*tchttp.BaseRequest

	// 待分析的文本（仅支持UTF-8格式，不超过500字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 词法分析模式（默认取1值）：
	// 1、高精度；
	// 2、高性能；
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *LexicalAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LexicalAnalysisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LexicalAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 命名实体识别结果。取值范围：
	// <li>PER：表示人名</li>
	// <li>LOC：表示地名</li>
	// <li>ORG：表示机构团体名</li>
		NerTokens []*NerToken `json:"NerTokens,omitempty" name:"NerTokens" list`

		// 分词&词性标注结果（词性表请参见附录）
		PosTokens []*PosToken `json:"PosTokens,omitempty" name:"PosTokens" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LexicalAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LexicalAnalysisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NerToken struct {

	// 起始位置
	BeginOffset *uint64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// 命名实体类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 基础词
	Word *string `json:"Word,omitempty" name:"Word"`
}

type PosToken struct {

	// 起始位置
	BeginOffset *uint64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 长度
	Length *uint64 `json:"Length,omitempty" name:"Length"`

	// 词性
	Pos *string `json:"Pos,omitempty" name:"Pos"`

	// 基础词
	Word *string `json:"Word,omitempty" name:"Word"`
}

type SensitiveWordsRecognitionRequest struct {
	*tchttp.BaseRequest

	// 待识别的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *SensitiveWordsRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SensitiveWordsRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SensitiveWordsRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 敏感词数组
		SensitiveWords []*string `json:"SensitiveWords,omitempty" name:"SensitiveWords" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SensitiveWordsRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SensitiveWordsRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceEmbeddingRequest struct {
	*tchttp.BaseRequest

	// 输入的文本（仅支持UTF-8格式，不超过500字）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *SentenceEmbeddingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceEmbeddingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 句向量的维度
		Dimension *uint64 `json:"Dimension,omitempty" name:"Dimension"`

		// 句向量数组
		Vector []*float64 `json:"Vector,omitempty" name:"Vector" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SentenceEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceEmbeddingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceSimilarityRequest struct {
	*tchttp.BaseRequest

	// 计算相似度的源句子（仅支持UTF-8格式，不超过500字）
	SrcText *string `json:"SrcText,omitempty" name:"SrcText"`

	// 计算相似度的目标句子（仅支持UTF-8格式，不超过500字）
	TargetText *string `json:"TargetText,omitempty" name:"TargetText"`
}

func (r *SentenceSimilarityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceSimilarityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceSimilarityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 两个文本的相似度
		Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SentenceSimilarityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceSimilarityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentimentAnalysisRequest struct {
	*tchttp.BaseRequest

	// 待分析的文本（仅支持UTF-8格式，不超过200字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 文本所属类型（默认取4值）：
	// 1、电商
	// 2、APP
	// 3、美食
	// 4、酒店和其他
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *SentimentAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentimentAnalysisRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentimentAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 负面情感概率
		Negative *float64 `json:"Negative,omitempty" name:"Negative"`

		// 正面情感概率
		Positive *float64 `json:"Positive,omitempty" name:"Positive"`

		// 情感属性
		Sentiment *string `json:"Sentiment,omitempty" name:"Sentiment"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SentimentAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentimentAnalysisResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimilarWordsRequest struct {
	*tchttp.BaseRequest

	// 输入的词语（仅支持UTF-8格式，不超过20字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 相似词个数；取值范围：1-200，默认为10；
	WordNumber *uint64 `json:"WordNumber,omitempty" name:"WordNumber"`
}

func (r *SimilarWordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimilarWordsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimilarWordsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 相似词数组
		SimilarWords []*string `json:"SimilarWords,omitempty" name:"SimilarWords" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SimilarWordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimilarWordsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextClassificationRequest struct {
	*tchttp.BaseRequest

	// 待分类的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 领域分类体系（默认取1值）：
	// 1、通用领域
	// 2、新闻领域
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *TextClassificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextClassificationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextClassificationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 文本分类结果（文本分类映射表请参见附录）
		Classes []*ClassificationResult `json:"Classes,omitempty" name:"Classes" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextClassificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextClassificationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextCorrectionRequest struct {
	*tchttp.BaseRequest

	// 待纠错的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *TextCorrectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextCorrectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextCorrectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 纠错详情
		CCITokens []*CCIToken `json:"CCITokens,omitempty" name:"CCITokens" list`

		// 纠错后的文本
		ResultText *string `json:"ResultText,omitempty" name:"ResultText"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextCorrectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextCorrectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WordEmbeddingRequest struct {
	*tchttp.BaseRequest

	// 输入的词语（仅支持UTF-8格式，不超过20字）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *WordEmbeddingRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WordEmbeddingRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WordEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 词向量的维度
		Dimension *uint64 `json:"Dimension,omitempty" name:"Dimension"`

		// 词向量数组
		Vector []*float64 `json:"Vector,omitempty" name:"Vector" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WordEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WordEmbeddingResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WordSimilarityRequest struct {
	*tchttp.BaseRequest

	// 计算相似度的源词（仅支持UTF-8格式，不超过20字）
	SrcWord *string `json:"SrcWord,omitempty" name:"SrcWord"`

	// 计算相似度的目标词（仅支持UTF-8格式，不超过20字）
	TargetWord *string `json:"TargetWord,omitempty" name:"TargetWord"`
}

func (r *WordSimilarityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WordSimilarityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WordSimilarityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 两个词语的相似度
		Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WordSimilarityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WordSimilarityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
