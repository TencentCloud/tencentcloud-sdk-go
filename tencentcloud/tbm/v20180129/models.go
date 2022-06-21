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

package v20180129

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgePortrait struct {
	// 年龄区间
	AgeRange *string `json:"AgeRange,omitempty" name:"AgeRange"`

	// 百分比
	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
}

type AgePortraitInfo struct {
	// 用户年龄画像数组
	PortraitSet []*AgePortrait `json:"PortraitSet,omitempty" name:"PortraitSet"`
}

type BrandReportArticle struct {
	// 文章标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 文章url地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 文章来源
	FromSite *string `json:"FromSite,omitempty" name:"FromSite"`

	// 文章发表日期
	PubTime *string `json:"PubTime,omitempty" name:"PubTime"`

	// 文章标识
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`

	// 文章热度值
	Hot *uint64 `json:"Hot,omitempty" name:"Hot"`

	// 文章来源等级
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 文章摘要
	Abstract *string `json:"Abstract,omitempty" name:"Abstract"`

	// 文章ID
	ArticleId *string `json:"ArticleId,omitempty" name:"ArticleId"`
}

type Comment struct {
	// 评论的日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 差评的个数
	NegCommentCount *uint64 `json:"NegCommentCount,omitempty" name:"NegCommentCount"`

	// 好评的个数
	PosCommentCount *uint64 `json:"PosCommentCount,omitempty" name:"PosCommentCount"`
}

type CommentInfo struct {
	// 用户评论内容
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 评论的时间
	Date *string `json:"Date,omitempty" name:"Date"`
}

type DateCount struct {
	// 统计日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 统计值
	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

// Predefined struct for user
type DescribeBrandCommentCountRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeBrandCommentCountRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始日期
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束日期
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeBrandCommentCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandCommentCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBrandCommentCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandCommentCountResponseParams struct {
	// 按天统计好评/差评数
	CommentSet []*Comment `json:"CommentSet,omitempty" name:"CommentSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBrandCommentCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBrandCommentCountResponseParams `json:"Response"`
}

func (r *DescribeBrandCommentCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandCommentCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandExposureRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeBrandExposureRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeBrandExposureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandExposureRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBrandExposureRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandExposureResponseParams struct {
	// 累计曝光量
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 按天计算的统计数据
	DateCountSet []*DateCount `json:"DateCountSet,omitempty" name:"DateCountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBrandExposureResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBrandExposureResponseParams `json:"Response"`
}

func (r *DescribeBrandExposureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandExposureResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandMediaReportRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeBrandMediaReportRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeBrandMediaReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandMediaReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBrandMediaReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandMediaReportResponseParams struct {
	// 查询范围内文章总数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 按天计算的每天文章数
	DateCountSet []*DateCount `json:"DateCountSet,omitempty" name:"DateCountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBrandMediaReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBrandMediaReportResponseParams `json:"Response"`
}

func (r *DescribeBrandMediaReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandMediaReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandNegCommentsRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 查询条数上限，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移，默认从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeBrandNegCommentsRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 查询条数上限，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移，默认从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBrandNegCommentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandNegCommentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBrandNegCommentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandNegCommentsResponseParams struct {
	// 评论列表
	BrandCommentSet []*CommentInfo `json:"BrandCommentSet,omitempty" name:"BrandCommentSet"`

	// 总的差评个数
	TotalComments *uint64 `json:"TotalComments,omitempty" name:"TotalComments"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBrandNegCommentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBrandNegCommentsResponseParams `json:"Response"`
}

func (r *DescribeBrandNegCommentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandNegCommentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandPosCommentsRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 查询条数上限，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeBrandPosCommentsRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 查询条数上限，默认20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 查询偏移，从0开始
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBrandPosCommentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandPosCommentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBrandPosCommentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandPosCommentsResponseParams struct {
	// 评论列表
	BrandCommentSet []*CommentInfo `json:"BrandCommentSet,omitempty" name:"BrandCommentSet"`

	// 总的好评个数
	TotalComments *uint64 `json:"TotalComments,omitempty" name:"TotalComments"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBrandPosCommentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBrandPosCommentsResponseParams `json:"Response"`
}

func (r *DescribeBrandPosCommentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandPosCommentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandSocialOpinionRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 检索开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 检索结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 查询偏移，默认从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询条数上限，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 列表显示标记，若为true，则返回文章列表详情
	ShowList *bool `json:"ShowList,omitempty" name:"ShowList"`
}

type DescribeBrandSocialOpinionRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 检索开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 检索结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 查询偏移，默认从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询条数上限，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 列表显示标记，若为true，则返回文章列表详情
	ShowList *bool `json:"ShowList,omitempty" name:"ShowList"`
}

func (r *DescribeBrandSocialOpinionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandSocialOpinionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ShowList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBrandSocialOpinionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandSocialOpinionResponseParams struct {
	// 文章总数
	ArticleCount *uint64 `json:"ArticleCount,omitempty" name:"ArticleCount"`

	// 来源统计总数
	FromCount *uint64 `json:"FromCount,omitempty" name:"FromCount"`

	// 疑似负面报道总数
	AdverseCount *uint64 `json:"AdverseCount,omitempty" name:"AdverseCount"`

	// 文章列表详情
	ArticleSet []*BrandReportArticle `json:"ArticleSet,omitempty" name:"ArticleSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBrandSocialOpinionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBrandSocialOpinionResponseParams `json:"Response"`
}

func (r *DescribeBrandSocialOpinionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandSocialOpinionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandSocialReportRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeBrandSocialReportRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeBrandSocialReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandSocialReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBrandSocialReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBrandSocialReportResponseParams struct {
	// 累计统计数据
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 按天计算的统计数据
	DateCountSet []*DateCount `json:"DateCountSet,omitempty" name:"DateCountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeBrandSocialReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBrandSocialReportResponseParams `json:"Response"`
}

func (r *DescribeBrandSocialReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBrandSocialReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndustryNewsRequestParams struct {
	// 行业ID
	IndustryId *string `json:"IndustryId,omitempty" name:"IndustryId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 是否显示列表，若为 true，则返回文章列表
	ShowList *bool `json:"ShowList,omitempty" name:"ShowList"`

	// 查询偏移，默认从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询条数上限，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeIndustryNewsRequest struct {
	*tchttp.BaseRequest
	
	// 行业ID
	IndustryId *string `json:"IndustryId,omitempty" name:"IndustryId"`

	// 查询开始时间
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 查询结束时间
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 是否显示列表，若为 true，则返回文章列表
	ShowList *bool `json:"ShowList,omitempty" name:"ShowList"`

	// 查询偏移，默认从0开始
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询条数上限，默认20
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIndustryNewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndustryNewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IndustryId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "ShowList")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndustryNewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndustryNewsResponseParams struct {
	// 总计文章数量
	NewsCount *uint64 `json:"NewsCount,omitempty" name:"NewsCount"`

	// 总计来源数量
	FromCount *uint64 `json:"FromCount,omitempty" name:"FromCount"`

	// 总计疑似负面数量
	AdverseCount *uint64 `json:"AdverseCount,omitempty" name:"AdverseCount"`

	// 文章列表
	NewsSet []*IndustryNews `json:"NewsSet,omitempty" name:"NewsSet"`

	// 按天统计的数量列表
	DateCountSet []*DateCount `json:"DateCountSet,omitempty" name:"DateCountSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIndustryNewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndustryNewsResponseParams `json:"Response"`
}

func (r *DescribeIndustryNewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndustryNewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPortraitRequestParams struct {
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`
}

type DescribeUserPortraitRequest struct {
	*tchttp.BaseRequest
	
	// 品牌ID
	BrandId *string `json:"BrandId,omitempty" name:"BrandId"`
}

func (r *DescribeUserPortraitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPortraitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BrandId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserPortraitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserPortraitResponseParams struct {
	// 年龄画像
	Age *AgePortraitInfo `json:"Age,omitempty" name:"Age"`

	// 性别画像
	Gender *GenderPortraitInfo `json:"Gender,omitempty" name:"Gender"`

	// 省份画像
	Province *ProvincePortraitInfo `json:"Province,omitempty" name:"Province"`

	// 电影喜好画像
	Movie *MoviePortraitInfo `json:"Movie,omitempty" name:"Movie"`

	// 明星喜好画像
	Star *StarPortraitInfo `json:"Star,omitempty" name:"Star"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserPortraitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserPortraitResponseParams `json:"Response"`
}

func (r *DescribeUserPortraitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserPortraitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenderPortrait struct {
	// 性别
	Gender *string `json:"Gender,omitempty" name:"Gender"`

	// 百分比
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`
}

type GenderPortraitInfo struct {
	// 用户性别画像数组
	PortraitSet []*GenderPortrait `json:"PortraitSet,omitempty" name:"PortraitSet"`
}

type IndustryNews struct {
	// 行业报道ID
	IndustryId *string `json:"IndustryId,omitempty" name:"IndustryId"`

	// 报道发表时间
	PubTime *string `json:"PubTime,omitempty" name:"PubTime"`

	// 报道来源
	FromSite *string `json:"FromSite,omitempty" name:"FromSite"`

	// 报道标题
	Title *string `json:"Title,omitempty" name:"Title"`

	// 报道来源url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 报道来源等级
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 热度值
	Hot *uint64 `json:"Hot,omitempty" name:"Hot"`

	// 报道标识
	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`

	// 报道摘要
	Abstract *string `json:"Abstract,omitempty" name:"Abstract"`
}

type MoviePortrait struct {
	// 电影名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 百分比
	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
}

type MoviePortraitInfo struct {
	// 用户喜好电影画像数组
	PortraitSet []*MoviePortrait `json:"PortraitSet,omitempty" name:"PortraitSet"`
}

type ProvincePortrait struct {
	// 省份名称
	Province *string `json:"Province,omitempty" name:"Province"`

	// 百分比
	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
}

type ProvincePortraitInfo struct {
	// 用户省份画像数组
	PortraitSet []*ProvincePortrait `json:"PortraitSet,omitempty" name:"PortraitSet"`
}

type StarPortrait struct {
	// 喜欢的明星名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 百分比
	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
}

type StarPortraitInfo struct {
	// 用户喜好的明星画像数组
	PortraitSet []*StarPortrait `json:"PortraitSet,omitempty" name:"PortraitSet"`
}