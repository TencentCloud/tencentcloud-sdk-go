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

package v20191115

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DataManipulationRequestParams struct {
	// 操作类型，add或del
	OpType *string `json:"OpType,omitempty" name:"OpType"`

	// 数据编码类型
	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`

	// 数据
	Contents *string `json:"Contents,omitempty" name:"Contents"`

	// 应用Id
	ResourceId *uint64 `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DataManipulationRequest struct {
	*tchttp.BaseRequest
	
	// 操作类型，add或del
	OpType *string `json:"OpType,omitempty" name:"OpType"`

	// 数据编码类型
	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`

	// 数据
	Contents *string `json:"Contents,omitempty" name:"Contents"`

	// 应用Id
	ResourceId *uint64 `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DataManipulationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DataManipulationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OpType")
	delete(f, "Encoding")
	delete(f, "Contents")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DataManipulationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DataManipulationResponseParams struct {
	// 数据操作结果
	Data *DataManipulationResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DataManipulationResponse struct {
	*tchttp.BaseResponse
	Response *DataManipulationResponseParams `json:"Response"`
}

func (r *DataManipulationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DataManipulationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataManipulationResult struct {
	// 应用ID
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 序号
	Seq *int64 `json:"Seq,omitempty" name:"Seq"`

	// 结果
	TotalResult *string `json:"TotalResult,omitempty" name:"TotalResult"`

	// 操作结果明细
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result []*DataManipulationResultItem `json:"Result,omitempty" name:"Result"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorResult *string `json:"ErrorResult,omitempty" name:"ErrorResult"`
}

type DataManipulationResultItem struct {
	// 结果
	Result *string `json:"Result,omitempty" name:"Result"`

	// 文档ID
	DocId *string `json:"DocId,omitempty" name:"DocId"`

	// 错误码
	Errno *int64 `json:"Errno,omitempty" name:"Errno"`
}

// Predefined struct for user
type DataSearchRequestParams struct {
	// 云搜的业务ID，用以表明当前数据请求的业务
	ResourceId *uint64 `json:"ResourceId,omitempty" name:"ResourceId"`

	// 检索串
	SearchQuery *string `json:"SearchQuery,omitempty" name:"SearchQuery"`

	// 当前页，从第0页开始计算
	PageId *uint64 `json:"PageId,omitempty" name:"PageId"`

	// 每页结果数
	NumPerPage *uint64 `json:"NumPerPage,omitempty" name:"NumPerPage"`

	// 当前检索号，用于定位问题，建议指定并且全局唯一
	SearchId *string `json:"SearchId,omitempty" name:"SearchId"`

	// 请求编码，0表示utf8，1表示gbk，建议指定
	QueryEncode *uint64 `json:"QueryEncode,omitempty" name:"QueryEncode"`

	// 排序类型
	RankType *uint64 `json:"RankType,omitempty" name:"RankType"`

	// 数值过滤，结果中按属性过滤
	NumFilter *string `json:"NumFilter,omitempty" name:"NumFilter"`

	// 分类过滤，导航类检索请求
	ClFilter *string `json:"ClFilter,omitempty" name:"ClFilter"`

	// 检索用户相关字段
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 检索来源
	SourceId *uint64 `json:"SourceId,omitempty" name:"SourceId"`

	// 是否进行二次检索，0关闭，1打开
	SecondSearch *uint64 `json:"SecondSearch,omitempty" name:"SecondSearch"`

	// 指定返回最大篇数，无特殊原因不建议指定
	MaxDocReturn *uint64 `json:"MaxDocReturn,omitempty" name:"MaxDocReturn"`

	// 是否smartbox检索，0关闭，1打开
	IsSmartbox *uint64 `json:"IsSmartbox,omitempty" name:"IsSmartbox"`

	// 是否打开高红标亮，0关闭，1打开
	EnableAbsHighlight *uint64 `json:"EnableAbsHighlight,omitempty" name:"EnableAbsHighlight"`

	// 指定访问QC纠错业务ID
	QcBid *uint64 `json:"QcBid,omitempty" name:"QcBid"`

	// 按指定字段进行group by，只能对数值字段进行操作
	GroupBy *string `json:"GroupBy,omitempty" name:"GroupBy"`

	// 按指定字段进行distinct，只能对数值字段进行操作
	Distinct *string `json:"Distinct,omitempty" name:"Distinct"`

	// 高级排序参数，具体参见高级排序说明
	L4RankExpression *string `json:"L4RankExpression,omitempty" name:"L4RankExpression"`

	// 高级排序参数，具体参见高级排序说明
	MatchValue *string `json:"MatchValue,omitempty" name:"MatchValue"`

	// 经度信息
	Longitude *float64 `json:"Longitude,omitempty" name:"Longitude"`

	// 纬度信息
	Latitude *float64 `json:"Latitude,omitempty" name:"Latitude"`

	// 分类过滤并集
	MultiFilter []*string `json:"MultiFilter,omitempty" name:"MultiFilter"`
}

type DataSearchRequest struct {
	*tchttp.BaseRequest
	
	// 云搜的业务ID，用以表明当前数据请求的业务
	ResourceId *uint64 `json:"ResourceId,omitempty" name:"ResourceId"`

	// 检索串
	SearchQuery *string `json:"SearchQuery,omitempty" name:"SearchQuery"`

	// 当前页，从第0页开始计算
	PageId *uint64 `json:"PageId,omitempty" name:"PageId"`

	// 每页结果数
	NumPerPage *uint64 `json:"NumPerPage,omitempty" name:"NumPerPage"`

	// 当前检索号，用于定位问题，建议指定并且全局唯一
	SearchId *string `json:"SearchId,omitempty" name:"SearchId"`

	// 请求编码，0表示utf8，1表示gbk，建议指定
	QueryEncode *uint64 `json:"QueryEncode,omitempty" name:"QueryEncode"`

	// 排序类型
	RankType *uint64 `json:"RankType,omitempty" name:"RankType"`

	// 数值过滤，结果中按属性过滤
	NumFilter *string `json:"NumFilter,omitempty" name:"NumFilter"`

	// 分类过滤，导航类检索请求
	ClFilter *string `json:"ClFilter,omitempty" name:"ClFilter"`

	// 检索用户相关字段
	Extra *string `json:"Extra,omitempty" name:"Extra"`

	// 检索来源
	SourceId *uint64 `json:"SourceId,omitempty" name:"SourceId"`

	// 是否进行二次检索，0关闭，1打开
	SecondSearch *uint64 `json:"SecondSearch,omitempty" name:"SecondSearch"`

	// 指定返回最大篇数，无特殊原因不建议指定
	MaxDocReturn *uint64 `json:"MaxDocReturn,omitempty" name:"MaxDocReturn"`

	// 是否smartbox检索，0关闭，1打开
	IsSmartbox *uint64 `json:"IsSmartbox,omitempty" name:"IsSmartbox"`

	// 是否打开高红标亮，0关闭，1打开
	EnableAbsHighlight *uint64 `json:"EnableAbsHighlight,omitempty" name:"EnableAbsHighlight"`

	// 指定访问QC纠错业务ID
	QcBid *uint64 `json:"QcBid,omitempty" name:"QcBid"`

	// 按指定字段进行group by，只能对数值字段进行操作
	GroupBy *string `json:"GroupBy,omitempty" name:"GroupBy"`

	// 按指定字段进行distinct，只能对数值字段进行操作
	Distinct *string `json:"Distinct,omitempty" name:"Distinct"`

	// 高级排序参数，具体参见高级排序说明
	L4RankExpression *string `json:"L4RankExpression,omitempty" name:"L4RankExpression"`

	// 高级排序参数，具体参见高级排序说明
	MatchValue *string `json:"MatchValue,omitempty" name:"MatchValue"`

	// 经度信息
	Longitude *float64 `json:"Longitude,omitempty" name:"Longitude"`

	// 纬度信息
	Latitude *float64 `json:"Latitude,omitempty" name:"Latitude"`

	// 分类过滤并集
	MultiFilter []*string `json:"MultiFilter,omitempty" name:"MultiFilter"`
}

func (r *DataSearchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DataSearchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceId")
	delete(f, "SearchQuery")
	delete(f, "PageId")
	delete(f, "NumPerPage")
	delete(f, "SearchId")
	delete(f, "QueryEncode")
	delete(f, "RankType")
	delete(f, "NumFilter")
	delete(f, "ClFilter")
	delete(f, "Extra")
	delete(f, "SourceId")
	delete(f, "SecondSearch")
	delete(f, "MaxDocReturn")
	delete(f, "IsSmartbox")
	delete(f, "EnableAbsHighlight")
	delete(f, "QcBid")
	delete(f, "GroupBy")
	delete(f, "Distinct")
	delete(f, "L4RankExpression")
	delete(f, "MatchValue")
	delete(f, "Longitude")
	delete(f, "Latitude")
	delete(f, "MultiFilter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DataSearchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DataSearchResponseParams struct {
	// 检索结果
	Data *SearchResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DataSearchResponse struct {
	*tchttp.BaseResponse
	Response *DataSearchResponseParams `json:"Response"`
}

func (r *DataSearchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DataSearchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchResult struct {
	// 检索耗时，单位ms
	CostTime *uint64 `json:"CostTime,omitempty" name:"CostTime"`

	// 搜索最多可以展示的结果数，多页
	DisplayNum *uint64 `json:"DisplayNum,omitempty" name:"DisplayNum"`

	// 和检索请求中的echo相对应
	Echo *string `json:"Echo,omitempty" name:"Echo"`

	// 检索结果的估算篇数，由索引平台估算
	EResultNum *uint64 `json:"EResultNum,omitempty" name:"EResultNum"`

	// 检索返回的当前页码结果数
	ResultNum *uint64 `json:"ResultNum,omitempty" name:"ResultNum"`

	// 检索结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultList []*SearchResultItem `json:"ResultList,omitempty" name:"ResultList"`

	// 检索的分词结果，array类型，可包含多个
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegList []*SearchResultSeg `json:"SegList,omitempty" name:"SegList"`
}

type SearchResultItem struct {
	// 动态摘要信息
	DocAbs *string `json:"DocAbs,omitempty" name:"DocAbs"`

	// 检索文档id
	DocId *string `json:"DocId,omitempty" name:"DocId"`

	// 原始文档信息
	DocMeta *string `json:"DocMeta,omitempty" name:"DocMeta"`

	// 精计算得分
	L2Score *float64 `json:"L2Score,omitempty" name:"L2Score"`

	// 文档级回传信息
	SearchDebuginfo *string `json:"SearchDebuginfo,omitempty" name:"SearchDebuginfo"`
}

type SearchResultSeg struct {
	// 分词
	SegStr *string `json:"SegStr,omitempty" name:"SegStr"`
}