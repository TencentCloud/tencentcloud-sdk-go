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

package v20210622

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreateProjectRequest struct {
	*tchttp.BaseRequest

	// 创建的项目名(不为空且最长为 200)
	Name *string `json:"Name,omitempty" name:"Name"`

	// 业务系统 ID
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 项目抽样率(大于等于 0)
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// 是否开启聚类
	EnableURLGroup *uint64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// 项目类型("web", "mp", "android", "ios", "node", "hippy", "weex", "viola", "rn")
	Type *string `json:"Type,omitempty" name:"Type"`

	// 项目对应仓库地址(可选，最长为 256)
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 项目对应网页地址(可选，最长为 256)
	URL *string `json:"URL,omitempty" name:"URL"`

	// 创建的项目描述(可选，最长为 1000)
	Desc *string `json:"Desc,omitempty" name:"Desc"`
}

func (r *CreateProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "InstanceID")
	delete(f, "Rate")
	delete(f, "EnableURLGroup")
	delete(f, "Type")
	delete(f, "Repo")
	delete(f, "URL")
	delete(f, "Desc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateProjectResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 项目 id
		ID *uint64 `json:"ID,omitempty" name:"ID"`

		// 项目唯一key
		Key *string `json:"Key,omitempty" name:"Key"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataEventUrlRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitempty" name:"Device"`

	// 是否海外
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// 筛选条件
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DescribeDataEventUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEventUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataEventUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回值
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataEventUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEventUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataLogUrlStatisticsRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// "analysis", "compare", "samp", "version", "ext3","nettype", "platform","isp","region","device","browser","ext1","ext2"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitempty" name:"Device"`

	// 是否海外
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitempty" name:"Browser"`
}

func (r *DescribeDataLogUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataLogUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataLogUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回值
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataLogUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataLogUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataPerformancePageRequest struct {
	*tchttp.BaseRequest

	// 项目ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// ["pagepv", "allcount"]
	Type *string `json:"Type,omitempty" name:"Type"`

	// 日志等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 运营商
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 地区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 网络类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 机型
	Device *string `json:"Device,omitempty" name:"Device"`

	// 版本
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// 是否海外
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// 浏览器
	Browser *string `json:"Browser,omitempty" name:"Browser"`

	// 操作系统
	Os *string `json:"Os,omitempty" name:"Os"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 来源页面
	From *string `json:"From,omitempty" name:"From"`

	// 耗时计算方式
	CostType *string `json:"CostType,omitempty" name:"CostType"`
}

func (r *DescribeDataPerformancePageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Type")
	delete(f, "Level")
	delete(f, "Isp")
	delete(f, "Area")
	delete(f, "NetType")
	delete(f, "Platform")
	delete(f, "Device")
	delete(f, "VersionNum")
	delete(f, "ExtFirst")
	delete(f, "ExtSecond")
	delete(f, "ExtThird")
	delete(f, "IsAbroad")
	delete(f, "Browser")
	delete(f, "Os")
	delete(f, "Engine")
	delete(f, "Brand")
	delete(f, "From")
	delete(f, "CostType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPerformancePageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataPerformancePageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回值
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataPerformancePageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPerformancePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataPvUrlStatisticsRequest struct {
	*tchttp.BaseRequest

	// 开始时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 类型:"allcount", "falls", "samp", "version", "ext3","nettype", "platform","isp","region","device","browser","ext1","ext2"
	Type *string `json:"Type,omitempty" name:"Type"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`

	// 项目ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 自定义2
	ExtSecond *string `json:"ExtSecond,omitempty" name:"ExtSecond"`

	// 浏览器引擎
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 运营商
	Isp *string `json:"Isp,omitempty" name:"Isp"`

	// 来源页面
	From *string `json:"From,omitempty" name:"From"`

	// 日志等级
	Level *string `json:"Level,omitempty" name:"Level"`

	// 品牌
	Brand *string `json:"Brand,omitempty" name:"Brand"`

	// 地区
	Area *string `json:"Area,omitempty" name:"Area"`

	// 版本
	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`

	// 平台
	Platform *string `json:"Platform,omitempty" name:"Platform"`

	// 自定义3
	ExtThird *string `json:"ExtThird,omitempty" name:"ExtThird"`

	// 自定义1
	ExtFirst *string `json:"ExtFirst,omitempty" name:"ExtFirst"`

	// 网络类型
	NetType *string `json:"NetType,omitempty" name:"NetType"`

	// 机型
	Device *string `json:"Device,omitempty" name:"Device"`

	// 是否海外
	IsAbroad *string `json:"IsAbroad,omitempty" name:"IsAbroad"`

	// 操作系统
	Os *string `json:"Os,omitempty" name:"Os"`

	// 浏览器
	Browser *string `json:"Browser,omitempty" name:"Browser"`
}

func (r *DescribeDataPvUrlStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "Type")
	delete(f, "EndTime")
	delete(f, "ID")
	delete(f, "ExtSecond")
	delete(f, "Engine")
	delete(f, "Isp")
	delete(f, "From")
	delete(f, "Level")
	delete(f, "Brand")
	delete(f, "Area")
	delete(f, "VersionNum")
	delete(f, "Platform")
	delete(f, "ExtThird")
	delete(f, "ExtFirst")
	delete(f, "NetType")
	delete(f, "Device")
	delete(f, "IsAbroad")
	delete(f, "Os")
	delete(f, "Browser")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataPvUrlStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataPvUrlStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回值
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataPvUrlStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataPvUrlStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorRequest struct {
	*tchttp.BaseRequest

	// 日期
	Date *string `json:"Date,omitempty" name:"Date"`

	// 项目ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeErrorRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Date")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeErrorRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeErrorResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内容
		Content *string `json:"Content,omitempty" name:"Content"`

		// 项目ID
		ID *int64 `json:"ID,omitempty" name:"ID"`

		// 时间
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeErrorResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeErrorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogListRequest struct {
	*tchttp.BaseRequest

	// 排序方式  desc  asc
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// searchlog   histogram
	ActionType *string `json:"ActionType,omitempty" name:"ActionType"`

	// 项目ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 上下文
	Context *string `json:"Context,omitempty" name:"Context"`

	// 查询语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sort")
	delete(f, "ActionType")
	delete(f, "ID")
	delete(f, "StartTime")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Query")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回字符串
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsRequest struct {
	*tchttp.BaseRequest

	// 分页每页数目，整型
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页页码，整型
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeProjectsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 列表总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 项目列表
		ProjectSet []*RumProject `json:"ProjectSet,omitempty" name:"ProjectSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScoresRequest struct {
	*tchttp.BaseRequest

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 项目ID
	ID *int64 `json:"ID,omitempty" name:"ID"`
}

func (r *DescribeScoresRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScoresRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScoresResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数组
		ScoreSet []*ScoreInfo `json:"ScoreSet,omitempty" name:"ScoreSet"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScoresResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScoresResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 一个或者多个过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 过滤键的名称。
	Name *string `json:"Name,omitempty" name:"Name"`
}

type RumProject struct {

	// 项目名
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建者 id
	Creator *string `json:"Creator,omitempty" name:"Creator"`

	// 实例 id
	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`

	// 项目类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 项目仓库地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Repo *string `json:"Repo,omitempty" name:"Repo"`

	// 项目网址地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	URL *string `json:"URL,omitempty" name:"URL"`

	// 项目采样频率
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// 项目唯一key（长度 12 位）
	Key *string `json:"Key,omitempty" name:"Key"`

	// 是否开启url聚类
	EnableURLGroup *int64 `json:"EnableURLGroup,omitempty" name:"EnableURLGroup"`

	// 实例名
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 项目 ID
	ID *int64 `json:"ID,omitempty" name:"ID"`

	// 实例 key
	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`

	// 项目描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitempty" name:"Desc"`

	// 是否星标  1:是 0:否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsStar *int64 `json:"IsStar,omitempty" name:"IsStar"`
}

type ScoreInfo struct {

	// duration
	StaticDuration *string `json:"StaticDuration,omitempty" name:"StaticDuration"`

	// pv
	PagePv *string `json:"PagePv,omitempty" name:"PagePv"`

	// 失败
	ApiFail *string `json:"ApiFail,omitempty" name:"ApiFail"`

	// 请求
	ApiNum *string `json:"ApiNum,omitempty" name:"ApiNum"`

	// fail
	StaticFail *string `json:"StaticFail,omitempty" name:"StaticFail"`

	// 项目id
	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`

	// uv
	PageUv *string `json:"PageUv,omitempty" name:"PageUv"`

	// 请求次数
	ApiDuration *string `json:"ApiDuration,omitempty" name:"ApiDuration"`

	// 分数
	Score *string `json:"Score,omitempty" name:"Score"`

	// error
	PageError *string `json:"PageError,omitempty" name:"PageError"`

	// num
	StaticNum *string `json:"StaticNum,omitempty" name:"StaticNum"`

	// num
	RecordNum *int64 `json:"RecordNum,omitempty" name:"RecordNum"`

	// Duration
	PageDuration *string `json:"PageDuration,omitempty" name:"PageDuration"`
}
