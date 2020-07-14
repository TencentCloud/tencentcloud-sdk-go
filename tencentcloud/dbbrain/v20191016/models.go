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

package v20191016

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 事件 ID 。通过“获取实例诊断历史DescribeDBDiagHistory”获取。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
}

func (r *DescribeDBDiagEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagEventRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 诊断项。
		DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

		// 诊断类型。
		DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

		// 事件 ID 。
		EventId *int64 `json:"EventId,omitempty" name:"EventId"`

		// 事件详情。
		Explanation *string `json:"Explanation,omitempty" name:"Explanation"`

		// 概要。
		Outline *string `json:"Outline,omitempty" name:"Outline"`

		// 诊断出的问题。
		Problem *string `json:"Problem,omitempty" name:"Problem"`

		// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
		Severity *int64 `json:"Severity,omitempty" name:"Severity"`

		// 开始时间
		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

		// 建议。
		Suggestions *string `json:"Suggestions,omitempty" name:"Suggestions"`

		// 保留字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Metric *string `json:"Metric,omitempty" name:"Metric"`

		// 结束时间。
		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDiagEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagEventResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间。如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。如“2019-09-11 12:13:14”。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeDBDiagHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagHistoryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件描述。
		Events []*DiagHistoryEventItem `json:"Events,omitempty" name:"Events" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBDiagHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBDiagHistoryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 时间段天数，截止日期为当日，默认为7天。
	RangeDays *int64 `json:"RangeDays,omitempty" name:"RangeDays"`
}

func (r *DescribeDBSpaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSpaceStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 磁盘增长量(MB)。
		Growth *int64 `json:"Growth,omitempty" name:"Growth"`

		// 磁盘剩余(MB)。
		Remain *int64 `json:"Remain,omitempty" name:"Remain"`

		// 磁盘总量(MB)。
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 预计可用天数。
		AvailableDays *int64 `json:"AvailableDays,omitempty" name:"AvailableDays"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSpaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDBSpaceStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 柱间单位时间间隔，单位为秒。
		Period *int64 `json:"Period,omitempty" name:"Period"`

		// 单位时间间隔内慢日志数量统计。
		TimeSeries []*TimeSlice `json:"TimeSeries,omitempty" name:"TimeSeries" list`

		// 单位时间间隔内的实例 cpu 利用率监控数据。
		SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 截止时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序）。
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTopSqlsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 慢日志 top sql 列表
		Rows []*SlowLogTopSqlItem `json:"Rows,omitempty" name:"Rows" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSlowLogTopSqlsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为20，默认为最大值。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 开始日期，最早为当日的前第6天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 截止日期，最早为当日的前第6天，默认为当日。
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Top表空间统计信息的时序数据列表。
		TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitempty" name:"TopSpaceTableTimeSeries" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为20，默认为最大值。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeTopSpaceTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的Top表空间统计信息列表。
		TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitempty" name:"TopSpaceTables" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopSpaceTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopSpaceTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {

	// 诊断类型。
	DiagType *string `json:"DiagType,omitempty" name:"DiagType"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 事件 ID 。
	EventId *int64 `json:"EventId,omitempty" name:"EventId"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitempty" name:"Severity"`

	// 概要。
	Outline *string `json:"Outline,omitempty" name:"Outline"`

	// 诊断项。
	DiagItem *string `json:"DiagItem,omitempty" name:"DiagItem"`

	// 实例 ID 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 保留字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type MonitorFloatMetric struct {

	// 指标名称。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*float64 `json:"Values,omitempty" name:"Values" list`
}

type MonitorFloatMetricSeriesData struct {

	// 监控指标。
	Series []*MonitorFloatMetric `json:"Series,omitempty" name:"Series" list`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp" list`
}

type MonitorMetric struct {

	// 指标名称。
	Metric *string `json:"Metric,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitempty" name:"Unit"`

	// 指标值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*int64 `json:"Values,omitempty" name:"Values" list`
}

type MonitorMetricSeriesData struct {

	// 监控指标。
	Series []*MonitorMetric `json:"Series,omitempty" name:"Series" list`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitempty" name:"Timestamp" list`
}

type SlowLogTopSqlItem struct {

	// sql总锁等待时间
	LockTime *float64 `json:"LockTime,omitempty" name:"LockTime"`

	// 最大锁等待时间
	LockTimeMax *float64 `json:"LockTimeMax,omitempty" name:"LockTimeMax"`

	// 最小锁等待时间
	LockTimeMin *float64 `json:"LockTimeMin,omitempty" name:"LockTimeMin"`

	// 总扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitempty" name:"RowsExamined"`

	// 最大扫描行数
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitempty" name:"RowsExaminedMax"`

	// 最小扫描行数
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitempty" name:"RowsExaminedMin"`

	// 总耗时
	QueryTime *float64 `json:"QueryTime,omitempty" name:"QueryTime"`

	// 最大执行时间
	QueryTimeMax *float64 `json:"QueryTimeMax,omitempty" name:"QueryTimeMax"`

	// 最小执行时间
	QueryTimeMin *float64 `json:"QueryTimeMin,omitempty" name:"QueryTimeMin"`

	// 总返回行数
	RowsSent *int64 `json:"RowsSent,omitempty" name:"RowsSent"`

	// 最大返回行数
	RowsSentMax *int64 `json:"RowsSentMax,omitempty" name:"RowsSentMax"`

	// 最小返回行数
	RowsSentMin *int64 `json:"RowsSentMin,omitempty" name:"RowsSentMin"`

	// 执行次数
	ExecTimes *int64 `json:"ExecTimes,omitempty" name:"ExecTimes"`

	// sql模板
	SqlTemplate *string `json:"SqlTemplate,omitempty" name:"SqlTemplate"`

	// 带参数SQL（随机）
	SqlText *string `json:"SqlText,omitempty" name:"SqlText"`

	// schema
	Schema *string `json:"Schema,omitempty" name:"Schema"`

	// 总耗时占比
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitempty" name:"QueryTimeRatio"`

	// sql总锁等待时间占比
	LockTimeRatio *float64 `json:"LockTimeRatio,omitempty" name:"LockTimeRatio"`

	// 总扫描行数占比
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitempty" name:"RowsExaminedRatio"`

	// 总返回行数占比
	RowsSentRatio *float64 `json:"RowsSentRatio,omitempty" name:"RowsSentRatio"`
}

type TableSpaceData struct {

	// 表名。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 数据空间（MB）。
	DataLength *float64 `json:"DataLength,omitempty" name:"DataLength"`

	// 索引空间（MB）。
	IndexLength *float64 `json:"IndexLength,omitempty" name:"IndexLength"`

	// 碎片空间（MB）。
	DataFree *float64 `json:"DataFree,omitempty" name:"DataFree"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitempty" name:"TotalLength"`

	// 碎片率（%）。
	FragRatio *float64 `json:"FragRatio,omitempty" name:"FragRatio"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitempty" name:"TableRows"`

	// 表对应的独立物理文件大小（MB）。
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitempty" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {

	// 表名。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitempty" name:"Engine"`

	// 单位时间间隔内的空间指标数据。
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitempty" name:"SeriesData"`
}

type TimeSlice struct {

	// 总数
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 统计开始时间
	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
}
