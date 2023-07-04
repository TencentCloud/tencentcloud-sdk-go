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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AnalyzeSentimentRequestParams struct {
	// 待分析的文本（仅支持UTF-8格式，不超过200字）。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AnalyzeSentimentRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（仅支持UTF-8格式，不超过200字）。
	Text *string `json:"Text,omitempty" name:"Text"`
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
	Positive *float64 `json:"Positive,omitempty" name:"Positive"`

	// 中性情感概率。
	Neutral *float64 `json:"Neutral,omitempty" name:"Neutral"`

	// 负面情感概率。
	Negative *float64 `json:"Negative,omitempty" name:"Negative"`

	// 情感分类结果：
	// positive：正面情感
	// negative：负面情感
	// neutral：中性、无情感
	Sentiment *string `json:"Sentiment,omitempty" name:"Sentiment"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type AutoSummarizationRequestParams struct {
	// 待处理的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 指定摘要的长度上限（默认值为200）
	// 注：为保证摘要的可读性，最终生成的摘要长度会低于指定的长度上限。
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

type AutoSummarizationRequest struct {
	*tchttp.BaseRequest
	
	// 待处理的文本（仅支持UTF-8格式，不超过2000字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 指定摘要的长度上限（默认值为200）
	// 注：为保证摘要的可读性，最终生成的摘要长度会低于指定的长度上限。
	Length *uint64 `json:"Length,omitempty" name:"Length"`
}

func (r *AutoSummarizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AutoSummarizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Length")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AutoSummarizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AutoSummarizationResponseParams struct {
	// 文本摘要结果
	Summary *string `json:"Summary,omitempty" name:"Summary"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AutoSummarizationResponse struct {
	*tchttp.BaseResponse
	Response *AutoSummarizationResponseParams `json:"Response"`
}

func (r *AutoSummarizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AutoSummarizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BasicParticiple struct {
	// 基础词。
	Word *string `json:"Word,omitempty" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 词性。
	Pos *string `json:"Pos,omitempty" name:"Pos"`
}

type CCIToken struct {
	// 错别字内容
	Word *string `json:"Word,omitempty" name:"Word"`

	// 错别字的起始位置，从0开始
	BeginOffset *uint64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 错别字纠错结果
	CorrectWord *string `json:"CorrectWord,omitempty" name:"CorrectWord"`
}

type Category struct {
	// 分类id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitempty" name:"Id"`

	// 分类英文名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 分类中文名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分类置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

// Predefined struct for user
type ChatBotRequestParams struct {
	// 用户请求的query
	Query *string `json:"Query,omitempty" name:"Query"`

	// 服务的id,  主要用于儿童闲聊接口，比如手Q的openid
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 0: 通用闲聊, 1:儿童闲聊, 默认是通用闲聊
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

type ChatBotRequest struct {
	*tchttp.BaseRequest
	
	// 用户请求的query
	Query *string `json:"Query,omitempty" name:"Query"`

	// 服务的id,  主要用于儿童闲聊接口，比如手Q的openid
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 0: 通用闲聊, 1:儿童闲聊, 默认是通用闲聊
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *ChatBotRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatBotRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	delete(f, "OpenId")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ChatBotRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ChatBotResponseParams struct {
	// 闲聊回复
	Reply *string `json:"Reply,omitempty" name:"Reply"`

	// 对于当前输出回复的自信度
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ChatBotResponse struct {
	*tchttp.BaseResponse
	Response *ChatBotResponseParams `json:"Response"`
}

func (r *ChatBotResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ChatBotResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassificationResult struct {
	// 一级分类名称
	FirstClassName *string `json:"FirstClassName,omitempty" name:"FirstClassName"`

	// 二级分类名称
	SecondClassName *string `json:"SecondClassName,omitempty" name:"SecondClassName"`

	// 一级分类概率
	FirstClassProbability *float64 `json:"FirstClassProbability,omitempty" name:"FirstClassProbability"`

	// 二级分类概率
	SecondClassProbability *float64 `json:"SecondClassProbability,omitempty" name:"SecondClassProbability"`

	// 三级分类名称，仅有当新闻领域五分类可能出现，详情见文本分类文档
	ThirdClassName *string `json:"ThirdClassName,omitempty" name:"ThirdClassName"`

	// 三级分类概率，仅有当新闻领域五分类可能出现，详情见文本分类文档
	ThirdClassProbability *float64 `json:"ThirdClassProbability,omitempty" name:"ThirdClassProbability"`

	// 四级分类名称，仅有当新闻领域五分类可能出现，详情见文本分类文档
	FourthClassName *string `json:"FourthClassName,omitempty" name:"FourthClassName"`

	// 四级分类概率，仅有当新闻领域五分类可能出现，详情见文本分类文档
	FourthClassProbability *float64 `json:"FourthClassProbability,omitempty" name:"FourthClassProbability"`

	// 五级分类名称，仅有当新闻领域五分类可能出现，详情见文本分类文档
	FifthClassName *string `json:"FifthClassName,omitempty" name:"FifthClassName"`

	// 五级分类概率，仅有当新闻领域五分类可能出现，详情见文本分类文档
	FifthClassProbability *float64 `json:"FifthClassProbability,omitempty" name:"FifthClassProbability"`
}

// Predefined struct for user
type ClassifyContentRequestParams struct {
	// 待分类的文章的标题（仅支持UTF-8格式，不超过100字符）。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 待分类文章的内容, 每个元素对应一个段落。（仅支持UTF-8格式，文章内容长度总和不超过2000字符）
	Content []*string `json:"Content,omitempty" name:"Content"`
}

type ClassifyContentRequest struct {
	*tchttp.BaseRequest
	
	// 待分类的文章的标题（仅支持UTF-8格式，不超过100字符）。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 待分类文章的内容, 每个元素对应一个段落。（仅支持UTF-8格式，文章内容长度总和不超过2000字符）
	Content []*string `json:"Content,omitempty" name:"Content"`
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
	// 一级分类。分类详情见附录-三级分类体系表。
	FirstClassification *Category `json:"FirstClassification,omitempty" name:"FirstClassification"`

	// 二级分类。分类详情见附录-三级分类体系表。
	SecondClassification *Category `json:"SecondClassification,omitempty" name:"SecondClassification"`

	// 三级分类。分类详情见附录-三级分类体系表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdClassification *Category `json:"ThirdClassification,omitempty" name:"ThirdClassification"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Text *string `json:"Text,omitempty" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitempty" name:"TargetType"`
}

type ComposeCoupletRequest struct {
	*tchttp.BaseRequest
	
	// 生成对联的关键词。长度需>=2，当长度>2时，自动截取前两个字作为关键字。内容需为常用汉字（不含有数字、英文、韩语、日语、符号等等其他）。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitempty" name:"TargetType"`
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
	TopScroll *string `json:"TopScroll,omitempty" name:"TopScroll"`

	// 上联与下联。
	Content []*string `json:"Content,omitempty" name:"Content"`

	// 当对联随机生成时，展示随机生成原因。
	RandomCause *string `json:"RandomCause,omitempty" name:"RandomCause"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Text *string `json:"Text,omitempty" name:"Text"`

	// 生成诗词的类型。0：藏头或藏身；1：藏头；2：藏身。默认为0。
	PoetryType *int64 `json:"PoetryType,omitempty" name:"PoetryType"`

	// 诗的体裁。0：五言律诗或七言律诗；5：五言律诗；7：七言律诗。默认为0。
	Genre *int64 `json:"Genre,omitempty" name:"Genre"`
}

type ComposePoetryRequest struct {
	*tchttp.BaseRequest
	
	// 生成诗词的关键词。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 生成诗词的类型。0：藏头或藏身；1：藏头；2：藏身。默认为0。
	PoetryType *int64 `json:"PoetryType,omitempty" name:"PoetryType"`

	// 诗的体裁。0：五言律诗或七言律诗；5：五言律诗；7：七言律诗。默认为0。
	Genre *int64 `json:"Genre,omitempty" name:"Genre"`
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
	Title *string `json:"Title,omitempty" name:"Title"`

	// 诗的内容。
	Content []*string `json:"Content,omitempty" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Word *string `json:"Word,omitempty" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 词性。
	Pos *string `json:"Pos,omitempty" name:"Pos"`
}

type CorrectionItem struct {
	// 纠错句子的序号。
	Order *int64 `json:"Order,omitempty" name:"Order"`

	// 错误的起始位置，从0开始。
	BeginOffset *int64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 错误内容长度。
	Len *int64 `json:"Len,omitempty" name:"Len"`

	// 错误内容。
	Word *string `json:"Word,omitempty" name:"Word"`

	// 纠错结果，当为删除类错误时，结果为null。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorrectWord []*string `json:"CorrectWord,omitempty" name:"CorrectWord"`

	// 纠错类型。0：替换；1：插入；2：删除。
	CorrectionType *int64 `json:"CorrectionType,omitempty" name:"CorrectionType"`

	// 纠错信息置信度。0：error；1：warning；error的置信度更高。（仅供参考）
	Confidence *int64 `json:"Confidence,omitempty" name:"Confidence"`

	// 纠错信息中文描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionZh *string `json:"DescriptionZh,omitempty" name:"DescriptionZh"`

	// 纠错信息英文描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DescriptionEn *string `json:"DescriptionEn,omitempty" name:"DescriptionEn"`
}

// Predefined struct for user
type CreateDictRequestParams struct {
	// 自定义词库名称，不超过20字。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义词库描述，不超过100字。
	Description *string `json:"Description,omitempty" name:"Description"`
}

type CreateDictRequest struct {
	*tchttp.BaseRequest
	
	// 自定义词库名称，不超过20字。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义词库描述，不超过100字。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateDictRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDictRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDictRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDictResponseParams struct {
	// 创建的自定义词库ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDictResponse struct {
	*tchttp.BaseResponse
	Response *CreateDictResponseParams `json:"Response"`
}

func (r *CreateDictResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWordItemsRequestParams struct {
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 待添加的词条集合。
	WordItems []*WordItem `json:"WordItems,omitempty" name:"WordItems"`
}

type CreateWordItemsRequest struct {
	*tchttp.BaseRequest
	
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 待添加的词条集合。
	WordItems []*WordItem `json:"WordItems,omitempty" name:"WordItems"`
}

func (r *CreateWordItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWordItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictId")
	delete(f, "WordItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWordItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWordItemsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWordItemsResponse struct {
	*tchttp.BaseResponse
	Response *CreateWordItemsResponseParams `json:"Response"`
}

func (r *CreateWordItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWordItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDictRequestParams struct {
	// 要删除的自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`
}

type DeleteDictRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`
}

func (r *DeleteDictRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDictRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDictRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDictResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDictResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDictResponseParams `json:"Response"`
}

func (r *DeleteDictResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWordItemsRequestParams struct {
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 待删除的词条集合。
	WordItems []*WordItem `json:"WordItems,omitempty" name:"WordItems"`
}

type DeleteWordItemsRequest struct {
	*tchttp.BaseRequest
	
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 待删除的词条集合。
	WordItems []*WordItem `json:"WordItems,omitempty" name:"WordItems"`
}

func (r *DeleteWordItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWordItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictId")
	delete(f, "WordItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWordItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWordItemsResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWordItemsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWordItemsResponseParams `json:"Response"`
}

func (r *DeleteWordItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWordItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DependencyParsingRequestParams struct {
	// 待分析的文本（仅支持UTF-8格式，不超过200字）
	Text *string `json:"Text,omitempty" name:"Text"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DependencyParsingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DependencyParsingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DependencyParsingResponseParams struct {
	// 句法依存分析结果，其中句法依存关系的类型包括：
	// <li>主谓关系，eg: 我送她一束花 (我 <-- 送)
	// <li>动宾关系，eg: 我送她一束花 (送 --> 花)
	// <li>间宾关系，eg: 我送她一束花 (送 --> 她)
	// <li>前置宾语，eg: 他什么书都读 (书 <-- 读)
	// <li>兼语，eg: 他请我吃饭 (请 --> 我)
	// <li>定中关系，eg: 红苹果 (红 <-- 苹果)
	// <li>状中结构，eg: 非常美丽 (非常 <-- 美丽)
	// <li>动补结构，eg: 做完了作业 (做 --> 完)
	// <li>并列关系，eg: 大山和大海 (大山 --> 大海)
	// <li>介宾关系，eg: 在贸易区内 (在 --> 内)
	// <li>左附加关系，eg: 大山和大海 (和 <-- 大海)
	// <li>右附加关系，eg: 孩子们 (孩子 --> 们)
	// <li>独立结构，eg: 两个单句在结构上彼此独立
	// <li>标点符号，eg: 。
	// <li>核心关系，eg: 整个句子的核心
	DpTokens []*DpToken `json:"DpTokens,omitempty" name:"DpTokens"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DependencyParsingResponse struct {
	*tchttp.BaseResponse
	Response *DependencyParsingResponseParams `json:"Response"`
}

func (r *DependencyParsingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DependencyParsingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDictRequestParams struct {
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 自定义词库名称，模糊搜索。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeDictRequest struct {
	*tchttp.BaseRequest
	
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 自定义词库名称，模糊搜索。
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeDictRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDictRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDictRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDictResponseParams struct {
	// 查询到的词库信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dicts []*DictInfo `json:"Dicts,omitempty" name:"Dicts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDictResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDictResponseParams `json:"Response"`
}

func (r *DescribeDictResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDictsRequestParams struct {
	// 每页数据量，范围为1~100，默认为10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeDictsRequest struct {
	*tchttp.BaseRequest
	
	// 每页数据量，范围为1~100，默认为10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDictsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDictsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDictsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDictsResponseParams struct {
	// 记录总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 自定义词库信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dicts []*DictInfo `json:"Dicts,omitempty" name:"Dicts"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDictsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDictsResponseParams `json:"Response"`
}

func (r *DescribeDictsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWordItemsRequestParams struct {
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 分页偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数据量，范围为1~100，默认为10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 待检索的词条文本，支持模糊匹配。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type DescribeWordItemsRequest struct {
	*tchttp.BaseRequest
	
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 分页偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 每页数据量，范围为1~100，默认为10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 待检索的词条文本，支持模糊匹配。
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *DescribeWordItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWordItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWordItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWordItemsResponseParams struct {
	// 词条记录总条数。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 词条信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WordItems []*WordItem `json:"WordItems,omitempty" name:"WordItems"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeWordItemsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWordItemsResponseParams `json:"Response"`
}

func (r *DescribeWordItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWordItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DictInfo struct {
	// 自定义词库名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 自定义词库ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 自定义词库描述信息。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义词库修改时间，形式为:yyyy-mm-dd hh:mm:ss。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 自定义词库创建时间，形式为:yyyy-mm-dd hh:mm:ss。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DpToken struct {
	// 句法依存关系的类型
	Relation *string `json:"Relation,omitempty" name:"Relation"`

	// 当前词父节点的序号
	HeadId *uint64 `json:"HeadId,omitempty" name:"HeadId"`

	// 基础词
	Word *string `json:"Word,omitempty" name:"Word"`

	// 基础词的序号
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type Embellish struct {
	// 润色后的文本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 润色类型。类型如下：
	// expansion：扩写
	// rewriting：改写
	// translation_m2a：从现代文改写为古文
	// translation_a2m：从古文改写为现代文
	// 
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmbellishType *string `json:"EmbellishType,omitempty" name:"EmbellishType"`
}

type Entity struct {
	// 基础词。
	Word *string `json:"Word,omitempty" name:"Word"`

	// 基础词在NormalText中的起始位置。
	BeginOffset *int64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 基础词的长度。
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 实体类型的标准名字。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 类型名字的自然语言表达。（中文或英文）
	Name *string `json:"Name,omitempty" name:"Name"`
}

// Predefined struct for user
type EvaluateSentenceSimilarityRequestParams struct {
	// 待分析的句子对数组。句子对应不超过1对，仅支持中文文本，原句子与目标句子均应不超过64字符。
	SentencePairList []*SentencePair `json:"SentencePairList,omitempty" name:"SentencePairList"`
}

type EvaluateSentenceSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的句子对数组。句子对应不超过1对，仅支持中文文本，原句子与目标句子均应不超过64字符。
	SentencePairList []*SentencePair `json:"SentencePairList,omitempty" name:"SentencePairList"`
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
	ScoreList []*float64 `json:"ScoreList,omitempty" name:"ScoreList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	SourceWord *string `json:"SourceWord,omitempty" name:"SourceWord"`

	// 计算相似度的目标词。（仅支持UTF-8格式，不超过10字符）
	TargetWord *string `json:"TargetWord,omitempty" name:"TargetWord"`
}

type EvaluateWordSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// 计算相似度的源词。（仅支持UTF-8格式，不超过10字符）
	SourceWord *string `json:"SourceWord,omitempty" name:"SourceWord"`

	// 计算相似度的目标词。（仅支持UTF-8格式，不超过10字符）
	TargetWord *string `json:"TargetWord,omitempty" name:"TargetWord"`
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
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type GenerateCoupletRequestParams struct {
	// 生成对联的关键词。长度需>=2，当长度>2时，自动截取前两个字作为关键字。内容需为常用汉字（不含有数字、英文、韩语、日语、符号等等其他）。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitempty" name:"TargetType"`
}

type GenerateCoupletRequest struct {
	*tchttp.BaseRequest
	
	// 生成对联的关键词。长度需>=2，当长度>2时，自动截取前两个字作为关键字。内容需为常用汉字（不含有数字、英文、韩语、日语、符号等等其他）。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 返回的文本结果为繁体还是简体。0：简体；1：繁体。默认为0。
	TargetType *int64 `json:"TargetType,omitempty" name:"TargetType"`
}

func (r *GenerateCoupletRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCoupletRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "TargetType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateCoupletRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateCoupletResponseParams struct {
	// 横批。
	TopScroll *string `json:"TopScroll,omitempty" name:"TopScroll"`

	// 上联与下联。
	Content []*string `json:"Content,omitempty" name:"Content"`

	// 当对联随机生成时，展示随机生成原因。
	RandomCause *string `json:"RandomCause,omitempty" name:"RandomCause"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenerateCoupletResponse struct {
	*tchttp.BaseResponse
	Response *GenerateCoupletResponseParams `json:"Response"`
}

func (r *GenerateCoupletResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCoupletResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateKeywordSentenceRequestParams struct {
	// 生成句子的关键词，关键词个数需不超过4个，中文关键词长度应不超过10字符，英文关键词长度不超过3个单词。关键词中不可包含标点符号。
	WordList []*string `json:"WordList,omitempty" name:"WordList"`

	// 返回生成句子的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 指定生成句子的领域，支持领域如下：
	// general：通用领域，支持中英文
	// academic：学术领域，仅支持英文
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
}

type GenerateKeywordSentenceRequest struct {
	*tchttp.BaseRequest
	
	// 生成句子的关键词，关键词个数需不超过4个，中文关键词长度应不超过10字符，英文关键词长度不超过3个单词。关键词中不可包含标点符号。
	WordList []*string `json:"WordList,omitempty" name:"WordList"`

	// 返回生成句子的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 指定生成句子的领域，支持领域如下：
	// general：通用领域，支持中英文
	// academic：学术领域，仅支持英文
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitempty" name:"Domain"`
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
	KeywordSentenceList []*KeywordSentence `json:"KeywordSentenceList,omitempty" name:"KeywordSentenceList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type GeneratePoetryRequestParams struct {
	// 生成诗词的关键词。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 生成诗词的类型。0：藏头或藏身；1：藏头；2：藏身。默认为0。
	PoetryType *int64 `json:"PoetryType,omitempty" name:"PoetryType"`

	// 诗的体裁。0：五言律诗或七言律诗；5：五言律诗；7：七言律诗。默认为0。
	Genre *int64 `json:"Genre,omitempty" name:"Genre"`
}

type GeneratePoetryRequest struct {
	*tchttp.BaseRequest
	
	// 生成诗词的关键词。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 生成诗词的类型。0：藏头或藏身；1：藏头；2：藏身。默认为0。
	PoetryType *int64 `json:"PoetryType,omitempty" name:"PoetryType"`

	// 诗的体裁。0：五言律诗或七言律诗；5：五言律诗；7：七言律诗。默认为0。
	Genre *int64 `json:"Genre,omitempty" name:"Genre"`
}

func (r *GeneratePoetryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneratePoetryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "PoetryType")
	delete(f, "Genre")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GeneratePoetryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GeneratePoetryResponseParams struct {
	// 诗题，即输入的生成诗词的关键词。
	Title *string `json:"Title,omitempty" name:"Title"`

	// 诗的内容。
	Content []*string `json:"Content,omitempty" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GeneratePoetryResponse struct {
	*tchttp.BaseResponse
	Response *GeneratePoetryResponseParams `json:"Response"`
}

func (r *GeneratePoetryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GeneratePoetryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Keyword struct {
	// 权重
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 关键词
	Word *string `json:"Word,omitempty" name:"Word"`
}

type KeywordSentence struct {
	// 通过关键词生成的句子。
	TargetText *string `json:"TargetText,omitempty" name:"TargetText"`
}

// Predefined struct for user
type KeywordsExtractionRequestParams struct {
	// 待处理的文本（仅支持UTF-8格式，不超过10000字符）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 指定关键词个数上限（默认值为5）
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

type KeywordsExtractionRequest struct {
	*tchttp.BaseRequest
	
	// 待处理的文本（仅支持UTF-8格式，不超过10000字符）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 指定关键词个数上限（默认值为5）
	Num *uint64 `json:"Num,omitempty" name:"Num"`
}

func (r *KeywordsExtractionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KeywordsExtractionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Num")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KeywordsExtractionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KeywordsExtractionResponseParams struct {
	// 关键词提取结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Keywords []*Keyword `json:"Keywords,omitempty" name:"Keywords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type KeywordsExtractionResponse struct {
	*tchttp.BaseResponse
	Response *KeywordsExtractionResponseParams `json:"Response"`
}

func (r *KeywordsExtractionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KeywordsExtractionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LexicalAnalysisRequestParams struct {
	// 待分析的文本（仅支持UTF-8格式，不超过500字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 指定要加载的自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 词法分析模式（默认取2值）：
	// 1、高精度（混合粒度分词能力）；
	// 2、高性能（单粒度分词能力）；
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

type LexicalAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（仅支持UTF-8格式，不超过500字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 指定要加载的自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 词法分析模式（默认取2值）：
	// 1、高精度（混合粒度分词能力）；
	// 2、高性能（单粒度分词能力）；
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *LexicalAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LexicalAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "DictId")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LexicalAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LexicalAnalysisResponseParams struct {
	// 命名实体识别结果。取值范围：
	// <li>PER：表示人名，如刘德华、贝克汉姆</li>
	// <li>LOC：表示地名，如北京、华山</li>
	// <li>ORG：表示机构团体名，如腾讯、最高人民法院、人大附中</li>
	// <li>PRODUCTION：表示产品名，如QQ、微信、iPhone</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NerTokens []*NerToken `json:"NerTokens,omitempty" name:"NerTokens"`

	// 分词&词性标注结果（词性表请参见附录）
	PosTokens []*PosToken `json:"PosTokens,omitempty" name:"PosTokens"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LexicalAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *LexicalAnalysisResponseParams `json:"Response"`
}

func (r *LexicalAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LexicalAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NerToken struct {
	// 基础词
	Word *string `json:"Word,omitempty" name:"Word"`

	// 长度
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 起始位置
	BeginOffset *int64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 命名实体类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

// Predefined struct for user
type ParseWordsRequestParams struct {
	// 待分析的文本（支持中英文文本，不超过500字符）
	Text *string `json:"Text,omitempty" name:"Text"`
}

type ParseWordsRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（支持中英文文本，不超过500字符）
	Text *string `json:"Text,omitempty" name:"Text"`
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
	NormalText *string `json:"NormalText,omitempty" name:"NormalText"`

	// 基础粒度分词和词性标注的结果。（词性表请参见附录）
	BasicParticiples []*BasicParticiple `json:"BasicParticiples,omitempty" name:"BasicParticiples"`

	// 复合粒度分词和词性标注的结果。（词性表请参见附录）
	CompoundParticiples []*CompoundParticiple `json:"CompoundParticiples,omitempty" name:"CompoundParticiples"`

	// 实体识别结果。（实体类型数据请参见附录）
	Entities []*Entity `json:"Entities,omitempty" name:"Entities"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type PosToken struct {
	// 基础词
	Word *string `json:"Word,omitempty" name:"Word"`

	// 长度
	Length *int64 `json:"Length,omitempty" name:"Length"`

	// 起始位置
	BeginOffset *int64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 词性
	Pos *string `json:"Pos,omitempty" name:"Pos"`
}

// Predefined struct for user
type RetrieveSimilarWordsRequestParams struct {
	// 输入的词语。（仅支持UTF-8格式，不超过10字符）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 召回的相似词个数，取值范围为1-20。
	Number *int64 `json:"Number,omitempty" name:"Number"`
}

type RetrieveSimilarWordsRequest struct {
	*tchttp.BaseRequest
	
	// 输入的词语。（仅支持UTF-8格式，不超过10字符）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 召回的相似词个数，取值范围为1-20。
	Number *int64 `json:"Number,omitempty" name:"Number"`
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
	WordList []*string `json:"WordList,omitempty" name:"WordList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SearchResult struct {
	// 被搜索的词条文本。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 0表示词条不存在，1表示存在。
	IsExist *uint64 `json:"IsExist,omitempty" name:"IsExist"`

	// 匹配到的词条文本。
	MatchText *string `json:"MatchText,omitempty" name:"MatchText"`

	// 词条的词性。
	Pos *string `json:"Pos,omitempty" name:"Pos"`
}

// Predefined struct for user
type SearchWordItemsRequestParams struct {
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 待检索的词条集合。
	WordItems []*WordItem `json:"WordItems,omitempty" name:"WordItems"`
}

type SearchWordItemsRequest struct {
	*tchttp.BaseRequest
	
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 待检索的词条集合。
	WordItems []*WordItem `json:"WordItems,omitempty" name:"WordItems"`
}

func (r *SearchWordItemsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchWordItemsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictId")
	delete(f, "WordItems")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchWordItemsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SearchWordItemsResponseParams struct {
	// 词条检索结果集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Results []*SearchResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SearchWordItemsResponse struct {
	*tchttp.BaseResponse
	Response *SearchWordItemsResponseParams `json:"Response"`
}

func (r *SearchWordItemsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchWordItemsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceCorrectionRequestParams struct {
	// 待纠错的句子列表。可以以数组方式在一次请求中填写多个待纠错的句子。文本统一使用utf-8格式编码，每个中文句子的长度不超过150字符，每个英文句子的长度不超过100个单词，且数组长度需小于150，即句子总数需少于150句。
	TextList []*string `json:"TextList,omitempty" name:"TextList"`
}

type SentenceCorrectionRequest struct {
	*tchttp.BaseRequest
	
	// 待纠错的句子列表。可以以数组方式在一次请求中填写多个待纠错的句子。文本统一使用utf-8格式编码，每个中文句子的长度不超过150字符，每个英文句子的长度不超过100个单词，且数组长度需小于150，即句子总数需少于150句。
	TextList []*string `json:"TextList,omitempty" name:"TextList"`
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
	CorrectionList []*CorrectionItem `json:"CorrectionList,omitempty" name:"CorrectionList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type SentenceEmbeddingRequestParams struct {
	// 输入的文本（仅支持UTF-8格式，不超过500字）
	Text *string `json:"Text,omitempty" name:"Text"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceEmbeddingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SentenceEmbeddingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentenceEmbeddingResponseParams struct {
	// 句向量数组
	Vector []*float64 `json:"Vector,omitempty" name:"Vector"`

	// 句向量的维度
	Dimension *uint64 `json:"Dimension,omitempty" name:"Dimension"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SentenceEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *SentenceEmbeddingResponseParams `json:"Response"`
}

func (r *SentenceEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentenceEmbeddingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SentencePair struct {
	// 需要与目标句子计算相似度的源句子。（仅支持UTF-8格式，不超过500字符）
	SourceText *string `json:"SourceText,omitempty" name:"SourceText"`

	// 目标句子。（仅支持UTF-8格式，不超过500字符）
	TargetText *string `json:"TargetText,omitempty" name:"TargetText"`
}

// Predefined struct for user
type SentimentAnalysisRequestParams struct {
	// 待分析的文本（仅支持UTF-8格式，不超过200字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 待分析文本所属的类型，仅当输入参数Mode取值为2class时有效（默认取4值）：
	// 1、商品评论类
	// 2、社交类
	// 3、美食酒店类
	// 4、通用领域类
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`

	// 情感分类模式选项，可取2class或3class（默认值为2class）
	// 1、2class：返回正负面二分类情感结果
	// 2、3class：返回正负面及中性三分类情感结果
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

type SentimentAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 待分析的文本（仅支持UTF-8格式，不超过200字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 待分析文本所属的类型，仅当输入参数Mode取值为2class时有效（默认取4值）：
	// 1、商品评论类
	// 2、社交类
	// 3、美食酒店类
	// 4、通用领域类
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`

	// 情感分类模式选项，可取2class或3class（默认值为2class）
	// 1、2class：返回正负面二分类情感结果
	// 2、3class：返回正负面及中性三分类情感结果
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

func (r *SentimentAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentimentAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Flag")
	delete(f, "Mode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SentimentAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SentimentAnalysisResponseParams struct {
	// 正面情感概率
	Positive *float64 `json:"Positive,omitempty" name:"Positive"`

	// 中性情感概率，当输入参数Mode取值为3class时有效，否则值为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Neutral *float64 `json:"Neutral,omitempty" name:"Neutral"`

	// 负面情感概率
	Negative *float64 `json:"Negative,omitempty" name:"Negative"`

	// 情感分类结果：
	// 1、positive，表示正面情感
	// 2、negative，表示负面情感
	// 3、neutral，表示中性、无情感
	Sentiment *string `json:"Sentiment,omitempty" name:"Sentiment"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SentimentAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *SentimentAnalysisResponseParams `json:"Response"`
}

func (r *SentimentAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SentimentAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SimilarWordsRequestParams struct {
	// 输入的词语（仅支持UTF-8格式，不超过20字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 相似词个数；取值范围：1-200，默认为10；
	WordNumber *uint64 `json:"WordNumber,omitempty" name:"WordNumber"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SimilarWordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "WordNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SimilarWordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SimilarWordsResponseParams struct {
	// 相似词数组
	SimilarWords []*string `json:"SimilarWords,omitempty" name:"SimilarWords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SimilarWordsResponse struct {
	*tchttp.BaseResponse
	Response *SimilarWordsResponseParams `json:"Response"`
}

func (r *SimilarWordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SimilarWordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Similarity struct {
	// 目标文本句子
	Text *string `json:"Text,omitempty" name:"Text"`

	// 相似度分数
	Score *float64 `json:"Score,omitempty" name:"Score"`
}

// Predefined struct for user
type TextClassificationRequestParams struct {
	// 待分类的文本（仅支持UTF-8格式，不超过10000字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 领域分类体系（默认取1值）：
	// 1、通用领域，二分类
	// 2、新闻领域，五分类。类别数据不一定全部返回，详情见类目映射表（注意：目前五分类已下线不可用）
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

type TextClassificationRequest struct {
	*tchttp.BaseRequest
	
	// 待分类的文本（仅支持UTF-8格式，不超过10000字）
	Text *string `json:"Text,omitempty" name:"Text"`

	// 领域分类体系（默认取1值）：
	// 1、通用领域，二分类
	// 2、新闻领域，五分类。类别数据不一定全部返回，详情见类目映射表（注意：目前五分类已下线不可用）
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
}

func (r *TextClassificationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextClassificationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	delete(f, "Flag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextClassificationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextClassificationResponseParams struct {
	// 文本分类结果（文本分类映射表请参见附录）
	Classes []*ClassificationResult `json:"Classes,omitempty" name:"Classes"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextClassificationResponse struct {
	*tchttp.BaseResponse
	Response *TextClassificationResponseParams `json:"Response"`
}

func (r *TextClassificationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextClassificationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextCorrectionProRequestParams struct {
	// 待纠错的文本（仅支持UTF-8格式，不超过128字符）
	Text *string `json:"Text,omitempty" name:"Text"`
}

type TextCorrectionProRequest struct {
	*tchttp.BaseRequest
	
	// 待纠错的文本（仅支持UTF-8格式，不超过128字符）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *TextCorrectionProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextCorrectionProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextCorrectionProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextCorrectionProResponseParams struct {
	// 纠错详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CCITokens []*CCIToken `json:"CCITokens,omitempty" name:"CCITokens"`

	// 纠错后的文本
	ResultText *string `json:"ResultText,omitempty" name:"ResultText"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextCorrectionProResponse struct {
	*tchttp.BaseResponse
	Response *TextCorrectionProResponseParams `json:"Response"`
}

func (r *TextCorrectionProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextCorrectionProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextCorrectionRequestParams struct {
	// 待纠错的文本（仅支持UTF-8格式，不超过2000字符）
	Text *string `json:"Text,omitempty" name:"Text"`
}

type TextCorrectionRequest struct {
	*tchttp.BaseRequest
	
	// 待纠错的文本（仅支持UTF-8格式，不超过2000字符）
	Text *string `json:"Text,omitempty" name:"Text"`
}

func (r *TextCorrectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextCorrectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextCorrectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextCorrectionResponseParams struct {
	// 纠错详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CCITokens []*CCIToken `json:"CCITokens,omitempty" name:"CCITokens"`

	// 纠错后的文本
	ResultText *string `json:"ResultText,omitempty" name:"ResultText"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextCorrectionResponse struct {
	*tchttp.BaseResponse
	Response *TextCorrectionResponseParams `json:"Response"`
}

func (r *TextCorrectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextCorrectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextEmbellishRequestParams struct {
	// 待润色的文本。中文文本长度需<=50字符；英文文本长度需<=30个单词。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 待润色文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitempty" name:"SourceLang"`

	// 返回润色结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 控制润色类型，类型如下：
	// both：同时返回改写和扩写
	// expansion：扩写
	// rewriting：改写
	// m2a：从现代文改写为古文
	// a2m：从古文改写为现代文
	// 默认为both。
	Style *string `json:"Style,omitempty" name:"Style"`
}

type TextEmbellishRequest struct {
	*tchttp.BaseRequest
	
	// 待润色的文本。中文文本长度需<=50字符；英文文本长度需<=30个单词。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 待润色文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitempty" name:"SourceLang"`

	// 返回润色结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 控制润色类型，类型如下：
	// both：同时返回改写和扩写
	// expansion：扩写
	// rewriting：改写
	// m2a：从现代文改写为古文
	// a2m：从古文改写为现代文
	// 默认为both。
	Style *string `json:"Style,omitempty" name:"Style"`
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
	EmbellishList []*Embellish `json:"EmbellishList,omitempty" name:"EmbellishList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type TextSimilarityProRequestParams struct {
	// 需要与目标句子计算相似度的源句子（仅支持UTF-8格式，不超过128字符）
	SrcText *string `json:"SrcText,omitempty" name:"SrcText"`

	// 目标句子（仅支持UTF-8格式，不超过128字符）
	TargetText []*string `json:"TargetText,omitempty" name:"TargetText"`
}

type TextSimilarityProRequest struct {
	*tchttp.BaseRequest
	
	// 需要与目标句子计算相似度的源句子（仅支持UTF-8格式，不超过128字符）
	SrcText *string `json:"SrcText,omitempty" name:"SrcText"`

	// 目标句子（仅支持UTF-8格式，不超过128字符）
	TargetText []*string `json:"TargetText,omitempty" name:"TargetText"`
}

func (r *TextSimilarityProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextSimilarityProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcText")
	delete(f, "TargetText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextSimilarityProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextSimilarityProResponseParams struct {
	// 每个目标句子与源句子的相似度分值，按照分值降序排列
	Similarity []*Similarity `json:"Similarity,omitempty" name:"Similarity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextSimilarityProResponse struct {
	*tchttp.BaseResponse
	Response *TextSimilarityProResponseParams `json:"Response"`
}

func (r *TextSimilarityProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextSimilarityProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextSimilarityRequestParams struct {
	// 需要与目标句子计算相似度的源句子（仅支持UTF-8格式，不超过500字符）
	SrcText *string `json:"SrcText,omitempty" name:"SrcText"`

	// 目标句子（以句子数量为单位消耗资源包）
	TargetText []*string `json:"TargetText,omitempty" name:"TargetText"`
}

type TextSimilarityRequest struct {
	*tchttp.BaseRequest
	
	// 需要与目标句子计算相似度的源句子（仅支持UTF-8格式，不超过500字符）
	SrcText *string `json:"SrcText,omitempty" name:"SrcText"`

	// 目标句子（以句子数量为单位消耗资源包）
	TargetText []*string `json:"TargetText,omitempty" name:"TargetText"`
}

func (r *TextSimilarityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextSimilarityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcText")
	delete(f, "TargetText")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TextSimilarityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextSimilarityResponseParams struct {
	// 每个目标句子与源句子的相似度分值，按照分值降序排列
	Similarity []*Similarity `json:"Similarity,omitempty" name:"Similarity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TextSimilarityResponse struct {
	*tchttp.BaseResponse
	Response *TextSimilarityResponseParams `json:"Response"`
}

func (r *TextSimilarityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TextSimilarityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TextWritingRequestParams struct {
	// 待续写的句子，文本统一使用utf-8格式编码，长度不超过200字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 待续写文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitempty" name:"SourceLang"`

	// 返回续写结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 指定续写领域，支持领域如下：
	// general：通用领域，支持中英文补全
	// academic：学术领域，仅支持英文补全
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 指定续写风格，支持风格如下：
	// science_fiction：科幻
	// military_history：军事
	// xuanhuan_wuxia：武侠
	// urban_officialdom：职场
	// 默认为xuanhuan_wuxia（武侠）。
	Style *string `json:"Style,omitempty" name:"Style"`
}

type TextWritingRequest struct {
	*tchttp.BaseRequest
	
	// 待续写的句子，文本统一使用utf-8格式编码，长度不超过200字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 待续写文本的语言类型，支持语言如下：
	// zh：中文
	// en：英文
	SourceLang *string `json:"SourceLang,omitempty" name:"SourceLang"`

	// 返回续写结果的个数。数量需>=1且<=5。
	// （注意实际结果可能小于指定个数）
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 指定续写领域，支持领域如下：
	// general：通用领域，支持中英文补全
	// academic：学术领域，仅支持英文补全
	// 默认为general（通用领域）。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 指定续写风格，支持风格如下：
	// science_fiction：科幻
	// military_history：军事
	// xuanhuan_wuxia：武侠
	// urban_officialdom：职场
	// 默认为xuanhuan_wuxia（武侠）。
	Style *string `json:"Style,omitempty" name:"Style"`
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
	WritingList []*Writing `json:"WritingList,omitempty" name:"WritingList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

// Predefined struct for user
type UpdateDictRequestParams struct {
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 词库描述，不超过100字。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 词库名称，不超过20字。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type UpdateDictRequest struct {
	*tchttp.BaseRequest
	
	// 自定义词库ID。
	DictId *string `json:"DictId,omitempty" name:"DictId"`

	// 词库描述，不超过100字。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 词库名称，不超过20字。
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *UpdateDictRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDictRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DictId")
	delete(f, "Description")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDictRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDictResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateDictResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDictResponseParams `json:"Response"`
}

func (r *UpdateDictResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDictResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WordEmbeddingRequestParams struct {
	// 输入的词语（仅支持UTF-8格式，不超过20字）
	Text *string `json:"Text,omitempty" name:"Text"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WordEmbeddingRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Text")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WordEmbeddingRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WordEmbeddingResponseParams struct {
	// 词向量数组
	Vector []*float64 `json:"Vector,omitempty" name:"Vector"`

	// 词向量的维度
	Dimension *uint64 `json:"Dimension,omitempty" name:"Dimension"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type WordEmbeddingResponse struct {
	*tchttp.BaseResponse
	Response *WordEmbeddingResponseParams `json:"Response"`
}

func (r *WordEmbeddingResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WordEmbeddingResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WordItem struct {
	// 词条文本内容。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 词条创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 词条的词性。
	Pos *string `json:"Pos,omitempty" name:"Pos"`
}

// Predefined struct for user
type WordSimilarityRequestParams struct {
	// 计算相似度的源词（仅支持UTF-8格式，不超过20字）
	SrcWord *string `json:"SrcWord,omitempty" name:"SrcWord"`

	// 计算相似度的目标词（仅支持UTF-8格式，不超过20字）
	TargetWord *string `json:"TargetWord,omitempty" name:"TargetWord"`
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

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WordSimilarityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcWord")
	delete(f, "TargetWord")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "WordSimilarityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type WordSimilarityResponseParams struct {
	// 两个词语的相似度
	Similarity *float64 `json:"Similarity,omitempty" name:"Similarity"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type WordSimilarityResponse struct {
	*tchttp.BaseResponse
	Response *WordSimilarityResponseParams `json:"Response"`
}

func (r *WordSimilarityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *WordSimilarityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Writing struct {
	// 续写的文本。
	TargetText *string `json:"TargetText,omitempty" name:"TargetText"`

	// 续写的前缀。
	PrefixText *string `json:"PrefixText,omitempty" name:"PrefixText"`
}