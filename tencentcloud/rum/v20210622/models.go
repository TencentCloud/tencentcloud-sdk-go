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

	// 项目对应实例 id
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
