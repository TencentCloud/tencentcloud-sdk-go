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

package v20220106

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CreatePurgeTaskRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 类型，当前支持的类型：
	// - purge_url：URL
	// - purge_prefix：前缀
	// - purge_host：Hostname
	// - purge_all：全部缓存
	Type *string `json:"Type,omitempty" name:"Type"`

	// 内容，一行一个
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源
	// 若内容含有非 ASCII 字符集的字符，请打开 URL Encode 开关，编码转换（编码规则遵循 RFC3986）
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`
}

func (r *CreatePurgeTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Type")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePurgeTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePurgeTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 失败的任务列表及原因
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePurgeTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePurgeTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 查询起始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 查询起始偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 查询最大返回的结果条数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询的状态
	// 允许的值为：processing、success、failed、timeout、invalid
	Statuses []*string `json:"Statuses,omitempty" name:"Statuses"`

	// zone id
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 查询的域名列表
	Domains []*string `json:"Domains,omitempty" name:"Domains"`

	// 查询内容
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePurgeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePurgeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePurgeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 该查询条件总共条目数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务结果列表
		Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePurgeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePurgeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailReason struct {

	// 失败原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 失败列表
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type Task struct {

	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 资源
	Target *string `json:"Target,omitempty" name:"Target"`

	// 任务类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 任务创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务完成时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}
