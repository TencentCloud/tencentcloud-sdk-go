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

type CreatePrefetchTaskRequest struct {
	*tchttp.BaseRequest

	// Zone ID
	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`

	// 要预热的资源列表，每个元素格式类似如下:
	// http://www.example.com/example.txt
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 是否对url进行encode
	// 若内容含有非 ASCII 字符集的字符，请开启此开关，编码转换（编码规则遵循 RFC3986）
	EncodeUrl *bool `json:"EncodeUrl,omitempty" name:"EncodeUrl"`

	// 附带的http头部信息
	Headers []*Header `json:"Headers,omitempty" name:"Headers"`
}

func (r *CreatePrefetchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneId")
	delete(f, "Targets")
	delete(f, "EncodeUrl")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePrefetchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreatePrefetchTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
		JobId *string `json:"JobId,omitempty" name:"JobId"`

		// 失败的任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailedList []*FailReason `json:"FailedList,omitempty" name:"FailedList"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePrefetchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePrefetchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

	// 要刷新的资源列表，每个元素格式依据Type而定
	// 1) Type = purge_host 时
	// 形如：www.example.com 或 foo.bar.example.com
	// 2) Type = purge_prefix 时
	// 形如：http://www.example.com/example
	// 3) Type = purge_url 时
	// 形如：https://www.example.com/example.jpg
	// 4）Type = purge_all 时
	// Targets可为空，不需要填写
	Targets []*string `json:"Targets,omitempty" name:"Targets"`

	// 若有编码转换，仅清除编码转换后匹配的资源
	// 若内容含有非 ASCII 字符集的字符，请开启此开关，编码转换（编码规则遵循 RFC3986）
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

type DescribePrefetchTasksRequest struct {
	*tchttp.BaseRequest

	// 任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

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

	// 查询的资源
	Target *string `json:"Target,omitempty" name:"Target"`
}

func (r *DescribePrefetchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Statuses")
	delete(f, "ZoneId")
	delete(f, "Domains")
	delete(f, "Target")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePrefetchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePrefetchTasksResponse struct {
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

func (r *DescribePrefetchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePrefetchTasksResponse) FromJsonString(s string) error {
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

type DescribeZonesRequest struct {
	*tchttp.BaseRequest

	// 分页参数，页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页参数，每页返回的站点个数
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 查询条件过滤器，复杂类型
	Filters []*ZoneFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZonesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZonesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZonesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的站点数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 站点详细信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Zones []*Zone `json:"Zones,omitempty" name:"Zones"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZonesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DownloadL7LogsRequest struct {
	*tchttp.BaseRequest

	// 起始时间(需严格按照RFC3339标准传参)
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间(需严格按照RFC3339标准传参)
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 每页展示条数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 当前页
	PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

	// 站点集合
	Zones []*string `json:"Zones,omitempty" name:"Zones"`

	// 域名集合
	Domains []*string `json:"Domains,omitempty" name:"Domains"`
}

func (r *DownloadL7LogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageSize")
	delete(f, "PageNo")
	delete(f, "Zones")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DownloadL7LogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DownloadL7LogsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 七层离线日志data
	// 注意：此字段可能返回 null，表示取不到有效值。
		Data []*L7OfflineLog `json:"Data,omitempty" name:"Data"`

		// 页面大小
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

		// 页号
	// 注意：此字段可能返回 null，表示取不到有效值。
		PageNo *int64 `json:"PageNo,omitempty" name:"PageNo"`

		// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
		Pages *int64 `json:"Pages,omitempty" name:"Pages"`

		// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DownloadL7LogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DownloadL7LogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FailReason struct {

	// 失败原因
	Reason *string `json:"Reason,omitempty" name:"Reason"`

	// 处理失败的资源列表。
	// 该列表元素来源于输入参数中的Targets，因此格式和入参中的Targets保持一致
	Targets []*string `json:"Targets,omitempty" name:"Targets"`
}

type Header struct {

	// HTTP头部
	Name *string `json:"Name,omitempty" name:"Name"`

	// HTTP头部值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type L7OfflineLog struct {

	// 日志打包开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTime *int64 `json:"LogTime,omitempty" name:"LogTime"`

	// 站点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domain *string `json:"Domain,omitempty" name:"Domain"`

	// 原始大小 单位byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 日志数据包名
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogPacketName *string `json:"LogPacketName,omitempty" name:"LogPacketName"`
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

type Zone struct {

	// 站点ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 站点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 站点当前使用的 NS 列表
	OriginalNameServers []*string `json:"OriginalNameServers,omitempty" name:"OriginalNameServers"`

	// 腾讯云分配的 NS 列表
	NameServers []*string `json:"NameServers,omitempty" name:"NameServers"`

	// 站点状态
	// - active：NS 已切换
	// - pending：NS 未切换
	// - moved：NS 已切走
	// - deactivated：被封禁
	Status *string `json:"Status,omitempty" name:"Status"`

	// 站点接入方式
	// - full：NS 接入
	// - partial：CNAME 接入
	Type *string `json:"Type,omitempty" name:"Type"`

	// 站点是否关闭
	Paused *bool `json:"Paused,omitempty" name:"Paused"`

	// 站点创建时间
	CreatedOn *string `json:"CreatedOn,omitempty" name:"CreatedOn"`

	// 站点修改时间
	ModifiedOn *string `json:"ModifiedOn,omitempty" name:"ModifiedOn"`
}

type ZoneFilter struct {

	// 过滤字段名，支持的列表如下：
	// - name: 站点名。
	// - status: 站点状态
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤字段值
	Values []*string `json:"Values,omitempty" name:"Values"`

	// 是否启用模糊查询，仅支持过滤字段名为name。模糊查询时，Values长度最大为1
	Fuzzy *bool `json:"Fuzzy,omitempty" name:"Fuzzy"`
}
