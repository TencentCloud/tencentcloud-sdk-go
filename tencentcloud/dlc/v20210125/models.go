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

package v20210125

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CSV struct {

	// 压缩格式，["Snappy", "Gzip", "None"选一]。
	CodeCompress *string `json:"CodeCompress,omitempty" name:"CodeCompress"`

	// CSV序列化及反序列化数据结构。
	CSVSerde *CSVSerde `json:"CSVSerde,omitempty" name:"CSVSerde"`

	// 标题行，默认为0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadLines *int64 `json:"HeadLines,omitempty" name:"HeadLines"`

	// 格式，默认值为CSV
	// 注意：此字段可能返回 null，表示取不到有效值。
	Format *string `json:"Format,omitempty" name:"Format"`
}

type CSVSerde struct {

	// CSV序列化转义符，默认为"\\"，最长8个字符，如 Escape: "/\"
	Escape *string `json:"Escape,omitempty" name:"Escape"`

	// CSV序列化字段域符，默认为"'"，最长8个字符, 如 Quote: "\""
	Quote *string `json:"Quote,omitempty" name:"Quote"`

	// CSV序列化分隔符，默认为"\t"，最长8个字符, 如 Separator: "\t"
	Separator *string `json:"Separator,omitempty" name:"Separator"`
}

type Column struct {

	// 列名称，不区分大小写，最大支持25个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 列类型，支持如下类型定义:
	// string|tinyint|smallint|int|bigint|boolean|float|double|decimal|timestamp|date|binary|array<data_type>|map<primitive_type, data_type>|struct<col_name : data_type [COMMENT col_comment], ...>|uniontype<data_type, data_type, ...>。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 对该类的注释。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest

	// 数据库基础信息
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitempty" name:"DatabaseInfo"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDatabaseRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成的建库执行语句对象。
		Execution *Execution `json:"Execution,omitempty" name:"Execution"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateDatabaseResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScriptRequest struct {
	*tchttp.BaseRequest

	// 脚本名称，最大不能超过255个字符。
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitempty" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitempty" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *CreateScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScriptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScriptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTableRequest struct {
	*tchttp.BaseRequest

	// 数据表配置信息
	TableInfo *TableInfo `json:"TableInfo,omitempty" name:"TableInfo"`
}

func (r *CreateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成的建表执行语句对象。
		Execution *Execution `json:"Execution,omitempty" name:"Execution"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest

	// 计算任务，该参数中包含任务类型及其相关配置信息
	Task *Task `json:"Task,omitempty" name:"Task"`

	// 数据库名称。任务在执行前均会USE该数据库， 除了首次建库时，其他情况建议均添加上。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataFormat struct {

	// 文本格式，TextFile。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextFile *TextFile `json:"TextFile,omitempty" name:"TextFile"`

	// 文本格式，CSV。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CSV *CSV `json:"CSV,omitempty" name:"CSV"`

	// 文本格式，Json。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *Other `json:"Json,omitempty" name:"Json"`

	// Parquet格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parquet *Other `json:"Parquet,omitempty" name:"Parquet"`

	// ORC格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ORC *Other `json:"ORC,omitempty" name:"ORC"`

	// AVRO格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AVRO *Other `json:"AVRO,omitempty" name:"AVRO"`
}

type DatabaseInfo struct {

	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~256。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 数据库属性列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties" list`
}

type DatabaseResponseInfo struct {

	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~256。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 数据库属性列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties" list`

	// 数据库创建时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据库更新时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}

type DeleteScriptRequest struct {
	*tchttp.BaseRequest

	// 脚本id，其可以通过DescribeScripts接口提取
	ScriptIds []*string `json:"ScriptIds,omitempty" name:"ScriptIds" list`
}

func (r *DeleteScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScriptRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScriptResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 删除的脚本数量
		ScriptsAffected *int64 `json:"ScriptsAffected,omitempty" name:"ScriptsAffected"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScriptResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据库对象列表。
		DatabaseList []*DatabaseResponseInfo `json:"DatabaseList,omitempty" name:"DatabaseList" list`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScriptsRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScriptsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScriptsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Script列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Scripts []*Script `json:"Scripts,omitempty" name:"Scripts" list`

		// 实例总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScriptsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest

	// 查询对象表名称
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
}

func (r *DescribeTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTableResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据表对象
		Table *TableResponseInfo `json:"Table,omitempty" name:"Table"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTableResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest

	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 数据表对象列表。
		TableList []*TableResponseInfo `json:"TableList,omitempty" name:"TableList" list`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTablesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,每个过滤参数支持的过滤值不超过5个。
	// task-id - String - （任务ID过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字）取值形如：DROP TABLE。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitempty" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务对象列表。
		TaskList []*TaskResponseInfo `json:"TaskList,omitempty" name:"TaskList" list`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest

	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeViewsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视图对象列表。
		ViewList []*ViewResponseInfo `json:"ViewList,omitempty" name:"ViewList" list`

		// 实例总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeViewsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Execution struct {

	// 自动生成SQL语句。
	SQL *string `json:"SQL,omitempty" name:"SQL"`
}

type Filter struct {

	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑或（OR）关系。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type Other struct {

	// 枚举类型，默认值为Json，可选值为[Json, Parquet, ORC, AVRD]之一。
	Format *string `json:"Format,omitempty" name:"Format"`
}

type Partition struct {

	// 分区列名。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 分区类型。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 对分区的描述。
	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type Property struct {

	// 属性key名称。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 属性key对应的value。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type SQLTask struct {

	// base64加密后的SQL语句
	SQL *string `json:"SQL,omitempty" name:"SQL"`
}

type Script struct {

	// 脚本Id，长度36字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptId *string `json:"ScriptId,omitempty" name:"ScriptId"`

	// 脚本名称，长度0-25。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptName *string `json:"ScriptName,omitempty" name:"ScriptName"`

	// 脚本描述，长度0-50。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDesc *string `json:"ScriptDesc,omitempty" name:"ScriptDesc"`

	// 默认关联数据库。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// SQL描述，长度0-10000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLStatement *string `json:"SQLStatement,omitempty" name:"SQLStatement"`

	// 更新时间戳， 单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TableBaseInfo struct {

	// 该数据表所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据表名字
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

type TableInfo struct {

	// 数据表配置信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 数据表格式。每次入参可选如下其一的KV结构，[TextFile，CSV，Json, Parquet, ORC, AVRD]。
	DataFormat *DataFormat `json:"DataFormat,omitempty" name:"DataFormat"`

	// 数据表列信息。
	Columns []*Column `json:"Columns,omitempty" name:"Columns" list`

	// 数据表分块信息。
	Partitions []*Partition `json:"Partitions,omitempty" name:"Partitions" list`

	// 数据存储路径。当前仅支持cos路径，格式如下：cosn://bucket-name/filepath。
	Location *string `json:"Location,omitempty" name:"Location"`
}

type TableResponseInfo struct {

	// 数据表基本信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitempty" name:"TableBaseInfo"`

	// 数据表列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitempty" name:"Columns" list`

	// 数据表分块信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*Partition `json:"Partitions,omitempty" name:"Partitions" list`

	// 数据存储路径。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitempty" name:"Location"`

	// 数据表属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties" list`

	// 数据表更新时间, 单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`

	// 数据表创建时间,单位: ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 数据格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputFormat *string `json:"InputFormat,omitempty" name:"InputFormat"`
}

type Task struct {

	// SQL查询任务
	SQLTask *SQLTask `json:"SQLTask,omitempty" name:"SQLTask"`
}

type TaskResponseInfo struct {

	// 任务所属Database的名称。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 任务数据量。
	DataAmount *int64 `json:"DataAmount,omitempty" name:"DataAmount"`

	// 任务Id。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 计算时长，单位： ms。
	UsedTime *int64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// 任务输出路径。
	OutputPath *string `json:"OutputPath,omitempty" name:"OutputPath"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务状态, 0 初始化， 1 执行中， 2 执行成功，3 数据写入中，-1 执行失败。
	State *int64 `json:"State,omitempty" name:"State"`

	// 任务SQL类型，DDL|DML等
	SQLType *string `json:"SQLType,omitempty" name:"SQLType"`

	// 任务SQL语句
	SQL *string `json:"SQL,omitempty" name:"SQL"`

	// 结果是否过期。
	ResultExpired *bool `json:"ResultExpired,omitempty" name:"ResultExpired"`

	// 数据影响统计信息。
	RowAffectInfo *string `json:"RowAffectInfo,omitempty" name:"RowAffectInfo"`

	// 任务结果数据表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSet *string `json:"DataSet,omitempty" name:"DataSet"`

	// 失败信息, 例如：errorMessage。该字段已废弃。
	Error *string `json:"Error,omitempty" name:"Error"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitempty" name:"Percentage"`

	// 任务执行输出信息。
	OutputMessage *string `json:"OutputMessage,omitempty" name:"OutputMessage"`
}

type TextFile struct {

	// 文本类型，本参数取值为TextFile。
	Format *string `json:"Format,omitempty" name:"Format"`

	// 处理文本用的正则表达式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type ViewBaseInfo struct {

	// 该视图所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 视图名称
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

type ViewResponseInfo struct {

	// 视图基本信息。
	ViewBaseInfo *ViewBaseInfo `json:"ViewBaseInfo,omitempty" name:"ViewBaseInfo"`

	// 视图列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitempty" name:"Columns" list`

	// 视图属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitempty" name:"Properties" list`

	// 视图创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 视图更新时间。
	ModifiedTime *string `json:"ModifiedTime,omitempty" name:"ModifiedTime"`
}
