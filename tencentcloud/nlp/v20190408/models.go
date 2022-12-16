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

type CCIToken struct {
	// 错别字内容
	Word *string `json:"Word,omitempty" name:"Word"`

	// 错别字的起始位置，从0开始
	BeginOffset *uint64 `json:"BeginOffset,omitempty" name:"BeginOffset"`

	// 错别字纠错结果
	CorrectWord *string `json:"CorrectWord,omitempty" name:"CorrectWord"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstClassProbability *float64 `json:"FirstClassProbability,omitempty" name:"FirstClassProbability"`

	// 二级分类概率
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondClassProbability *float64 `json:"SecondClassProbability,omitempty" name:"SecondClassProbability"`

	// 三级分类名称，仅有当新闻领域五分类可能出现，详情见文本分类文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdClassName *string `json:"ThirdClassName,omitempty" name:"ThirdClassName"`

	// 三级分类概率，仅有当新闻领域五分类可能出现，详情见文本分类文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdClassProbability *float64 `json:"ThirdClassProbability,omitempty" name:"ThirdClassProbability"`

	// 四级分类名称，仅有当新闻领域五分类可能出现，详情见文本分类文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	FourthClassName *string `json:"FourthClassName,omitempty" name:"FourthClassName"`

	// 四级分类概率，仅有当新闻领域五分类可能出现，详情见文本分类文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	FourthClassProbability *float64 `json:"FourthClassProbability,omitempty" name:"FourthClassProbability"`

	// 五级分类名称，仅有当新闻领域五分类可能出现，详情见文本分类文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	FifthClassName *string `json:"FifthClassName,omitempty" name:"FifthClassName"`

	// 五级分类概率，仅有当新闻领域五分类可能出现，详情见文本分类文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	FifthClassProbability *float64 `json:"FifthClassProbability,omitempty" name:"FifthClassProbability"`
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 自定义词库修改时间，形式为:yyyy-mm-dd hh:mm:ss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 自定义词库创建时间，形式为:yyyy-mm-dd hh:mm:ss。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DpToken struct {
	// 句法依存关系的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Relation *string `json:"Relation,omitempty" name:"Relation"`

	// 当前词父节点的序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadId *uint64 `json:"HeadId,omitempty" name:"HeadId"`

	// 基础词
	// 注意：此字段可能返回 null，表示取不到有效值。
	Word *string `json:"Word,omitempty" name:"Word"`

	// 基础词的序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitempty" name:"Id"`
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

type SearchResult struct {
	// 被搜索的词条文本。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 0表示词条不存在，1表示存在。
	IsExist *uint64 `json:"IsExist,omitempty" name:"IsExist"`

	// 匹配到的词条文本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchText *string `json:"MatchText,omitempty" name:"MatchText"`

	// 词条的词性。
	// 注意：此字段可能返回 null，表示取不到有效值。
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
	// 注意：此字段可能返回 null，表示取不到有效值。
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