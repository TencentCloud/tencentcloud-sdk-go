// Copyright (c) 2017-2025 Tencent. All Rights Reserved.
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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccessInfo struct {
	// 访问引擎的方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *string `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// 访问引擎的url，内部的部分参数需要根据实际情况替换
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessConnectionInfos []*string `json:"AccessConnectionInfos,omitnil,omitempty" name:"AccessConnectionInfos"`
}

// Predefined struct for user
type AddDMSPartitionsRequestParams struct {
	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

type AddDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`
}

func (r *AddDMSPartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDMSPartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Partitions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddDMSPartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddDMSPartitionsResponseParams struct {
	// 成功数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 分区值
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddDMSPartitionsResponse struct {
	*tchttp.BaseResponse
	Response *AddDMSPartitionsResponseParams `json:"Response"`
}

func (r *AddDMSPartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddDMSPartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOptimizerEnginesRequestParams struct {
	// 数据目录名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 引擎信息列表
	Engines []*OptimizerEngineInfo `json:"Engines,omitnil,omitempty" name:"Engines"`

	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据表名称
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type AddOptimizerEnginesRequest struct {
	*tchttp.BaseRequest
	
	// 数据目录名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 引擎信息列表
	Engines []*OptimizerEngineInfo `json:"Engines,omitnil,omitempty" name:"Engines"`

	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据表名称
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

func (r *AddOptimizerEnginesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOptimizerEnginesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Catalog")
	delete(f, "Engines")
	delete(f, "Database")
	delete(f, "Table")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddOptimizerEnginesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddOptimizerEnginesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddOptimizerEnginesResponse struct {
	*tchttp.BaseResponse
	Response *AddOptimizerEnginesResponseParams `json:"Response"`
}

func (r *AddOptimizerEnginesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddOptimizerEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersToWorkGroupRequestParams struct {
	// 要操作的工作组和用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

type AddUsersToWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要操作的工作组和用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

func (r *AddUsersToWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersToWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUsersToWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUsersToWorkGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUsersToWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddUsersToWorkGroupResponseParams `json:"Response"`
}

func (r *AddUsersToWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUsersToWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSDatabaseRequestParams struct {
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil,omitempty" name:"CurrentName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 路径
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type AlterDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil,omitempty" name:"CurrentName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 路径
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *AlterDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrentName")
	delete(f, "SchemaName")
	delete(f, "Location")
	delete(f, "Asset")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AlterDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSDatabaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AlterDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *AlterDMSDatabaseResponseParams `json:"Response"`
}

func (r *AlterDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSPartitionRequestParams struct {
	// 当前名称，变更前db名称
	CurrentDbName *string `json:"CurrentDbName,omitnil,omitempty" name:"CurrentDbName"`

	// 当前名称，变更前table名称
	CurrentTableName *string `json:"CurrentTableName,omitnil,omitempty" name:"CurrentTableName"`

	// 当前名称，变更前Part名称
	CurrentValues *string `json:"CurrentValues,omitnil,omitempty" name:"CurrentValues"`

	// 分区
	Partition *DMSPartition `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type AlterDMSPartitionRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称，变更前db名称
	CurrentDbName *string `json:"CurrentDbName,omitnil,omitempty" name:"CurrentDbName"`

	// 当前名称，变更前table名称
	CurrentTableName *string `json:"CurrentTableName,omitnil,omitempty" name:"CurrentTableName"`

	// 当前名称，变更前Part名称
	CurrentValues *string `json:"CurrentValues,omitnil,omitempty" name:"CurrentValues"`

	// 分区
	Partition *DMSPartition `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *AlterDMSPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrentDbName")
	delete(f, "CurrentTableName")
	delete(f, "CurrentValues")
	delete(f, "Partition")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AlterDMSPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSPartitionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AlterDMSPartitionResponse struct {
	*tchttp.BaseResponse
	Response *AlterDMSPartitionResponseParams `json:"Response"`
}

func (r *AlterDMSPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSTableRequestParams struct {
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil,omitempty" name:"CurrentName"`

	// 当前数据库名称
	CurrentDbName *string `json:"CurrentDbName,omitnil,omitempty" name:"CurrentDbName"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 当前表名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type AlterDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 当前名称
	CurrentName *string `json:"CurrentName,omitnil,omitempty" name:"CurrentName"`

	// 当前数据库名称
	CurrentDbName *string `json:"CurrentDbName,omitnil,omitempty" name:"CurrentDbName"`

	// 基础对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 表类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 当前表名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *AlterDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CurrentName")
	delete(f, "CurrentDbName")
	delete(f, "Asset")
	delete(f, "Type")
	delete(f, "DbName")
	delete(f, "StorageSize")
	delete(f, "RecordCount")
	delete(f, "LifeTime")
	delete(f, "DataUpdateTime")
	delete(f, "StructUpdateTime")
	delete(f, "LastAccessTime")
	delete(f, "Sds")
	delete(f, "Columns")
	delete(f, "PartitionKeys")
	delete(f, "ViewOriginalText")
	delete(f, "ViewExpandedText")
	delete(f, "Partitions")
	delete(f, "Name")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AlterDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AlterDMSTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AlterDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *AlterDMSTableResponseParams `json:"Response"`
}

func (r *AlterDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AlterDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalysisTaskResults struct {
	// 任务Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务创建时间，毫秒时间戳
	InstanceStartTime *int64 `json:"InstanceStartTime,omitnil,omitempty" name:"InstanceStartTime"`

	// 任务结束时间，毫秒时间戳
	InstanceCompleteTime *int64 `json:"InstanceCompleteTime,omitnil,omitempty" name:"InstanceCompleteTime"`

	// 任务状态：0 初始化， 1 执行中， 2 执行成功，3 数据写入中，4 排队中。-1 执行失败，-3 已取消。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 任务SQL语句
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`

	// 计算资源名字
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 单位毫秒，引擎内执行耗时, 反映真正用于计算所需的耗时，即从  Spark 任务第一个 Task  开始执行到任务结束之间的耗时。
	// 具体的：会统计任务的每个 Spark Stage 第一个 Task 到最后一个 Task 完成时长之和，不包含任务开始的排队耗时（即剔除从任务提交到 Spark Task 开始执行之间的调度等其他耗时），也不包含任务执行过程中多个 Spark Stage 之间因 executor 资源不足而等待执行 Task 所消耗的时间。
	JobTimeSum *int64 `json:"JobTimeSum,omitnil,omitempty" name:"JobTimeSum"`

	// 单位秒，累计 CPU* 秒 ( 累计 CPU * 时 = 累计 CPU* 秒/ 3600)，统计参与计算所用 Spark Executor 每个 core 的 CPU 执行时长总和
	TaskTimeSum *int64 `json:"TaskTimeSum,omitnil,omitempty" name:"TaskTimeSum"`

	// 数据扫描总行数
	InputRecordsSum *int64 `json:"InputRecordsSum,omitnil,omitempty" name:"InputRecordsSum"`

	// 数据扫描总 bytes
	InputBytesSum *int64 `json:"InputBytesSum,omitnil,omitempty" name:"InputBytesSum"`

	// 输出总行数
	OutputRecordsSum *int64 `json:"OutputRecordsSum,omitnil,omitempty" name:"OutputRecordsSum"`

	// 输出总 bytes
	OutputBytesSum *int64 `json:"OutputBytesSum,omitnil,omitempty" name:"OutputBytesSum"`

	// shuffle read 总 bytes
	ShuffleReadBytesSum *int64 `json:"ShuffleReadBytesSum,omitnil,omitempty" name:"ShuffleReadBytesSum"`

	// shuffle read 总行数
	ShuffleReadRecordsSum *int64 `json:"ShuffleReadRecordsSum,omitnil,omitempty" name:"ShuffleReadRecordsSum"`

	// 洞察结果类型分类，一个 json 数组，有如下几种类型：SPARK-StageScheduleDelay（资源抢占）, SPARK-ShuffleFailure（Shuffle异常）, SPARK-SlowTask（慢task）, SPARK-DataSkew（数据倾斜）, SPARK-InsufficientResource（磁盘或内存不足）
	AnalysisStatus *string `json:"AnalysisStatus,omitnil,omitempty" name:"AnalysisStatus"`

	// 任务输出文件总数
	OutputFilesNum *int64 `json:"OutputFilesNum,omitnil,omitempty" name:"OutputFilesNum"`

	// 任务输出小文件总数
	OutputSmallFilesNum *int64 `json:"OutputSmallFilesNum,omitnil,omitempty" name:"OutputSmallFilesNum"`
}

type Asset struct {
	// 主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 对象GUID值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Guid *string `json:"Guid,omitnil,omitempty" name:"Guid"`

	// 数据目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 对象owner
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 对象owner账户
	OwnerAccount *string `json:"OwnerAccount,omitnil,omitempty" name:"OwnerAccount"`

	// 权限
	PermValues []*KVPair `json:"PermValues,omitnil,omitempty" name:"PermValues"`

	// 附加属性
	Params []*KVPair `json:"Params,omitnil,omitempty" name:"Params"`

	// 附加业务属性
	BizParams []*KVPair `json:"BizParams,omitnil,omitempty" name:"BizParams"`

	// 数据版本
	DataVersion *int64 `json:"DataVersion,omitnil,omitempty" name:"DataVersion"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 数据源主键
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`
}

// Predefined struct for user
type AssignMangedTablePropertiesRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// V2 upsert表 upsert键
	UpsertKeys []*string `json:"UpsertKeys,omitnil,omitempty" name:"UpsertKeys"`
}

type AssignMangedTablePropertiesRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// V2 upsert表 upsert键
	UpsertKeys []*string `json:"UpsertKeys,omitnil,omitempty" name:"UpsertKeys"`
}

func (r *AssignMangedTablePropertiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignMangedTablePropertiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "Partitions")
	delete(f, "Properties")
	delete(f, "UpsertKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssignMangedTablePropertiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssignMangedTablePropertiesResponseParams struct {
	// 分配的原生表表属性
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssignMangedTablePropertiesResponse struct {
	*tchttp.BaseResponse
	Response *AssignMangedTablePropertiesResponseParams `json:"Response"`
}

func (r *AssignMangedTablePropertiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssignMangedTablePropertiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDatasourceHouseRequestParams struct {
	// 网络配置名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 数据源类型
	DatasourceConnectionType *string `json:"DatasourceConnectionType,omitnil,omitempty" name:"DatasourceConnectionType"`

	// 数据源网络配置
	DatasourceConnectionConfig *DatasourceConnectionConfig `json:"DatasourceConnectionConfig,omitnil,omitempty" name:"DatasourceConnectionConfig"`

	// 引擎名称，只允许绑定一个引擎
	DataEngineNames []*string `json:"DataEngineNames,omitnil,omitempty" name:"DataEngineNames"`

	// 网络类型，2-跨源型，4-增强型
	NetworkConnectionType *int64 `json:"NetworkConnectionType,omitnil,omitempty" name:"NetworkConnectionType"`

	// 网络配置描述
	NetworkConnectionDesc *string `json:"NetworkConnectionDesc,omitnil,omitempty" name:"NetworkConnectionDesc"`
}

type AssociateDatasourceHouseRequest struct {
	*tchttp.BaseRequest
	
	// 网络配置名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 数据源类型
	DatasourceConnectionType *string `json:"DatasourceConnectionType,omitnil,omitempty" name:"DatasourceConnectionType"`

	// 数据源网络配置
	DatasourceConnectionConfig *DatasourceConnectionConfig `json:"DatasourceConnectionConfig,omitnil,omitempty" name:"DatasourceConnectionConfig"`

	// 引擎名称，只允许绑定一个引擎
	DataEngineNames []*string `json:"DataEngineNames,omitnil,omitempty" name:"DataEngineNames"`

	// 网络类型，2-跨源型，4-增强型
	NetworkConnectionType *int64 `json:"NetworkConnectionType,omitnil,omitempty" name:"NetworkConnectionType"`

	// 网络配置描述
	NetworkConnectionDesc *string `json:"NetworkConnectionDesc,omitnil,omitempty" name:"NetworkConnectionDesc"`
}

func (r *AssociateDatasourceHouseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDatasourceHouseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceConnectionName")
	delete(f, "DatasourceConnectionType")
	delete(f, "DatasourceConnectionConfig")
	delete(f, "DataEngineNames")
	delete(f, "NetworkConnectionType")
	delete(f, "NetworkConnectionDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AssociateDatasourceHouseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AssociateDatasourceHouseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AssociateDatasourceHouseResponse struct {
	*tchttp.BaseResponse
	Response *AssociateDatasourceHouseResponseParams `json:"Response"`
}

func (r *AssociateDatasourceHouseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AssociateDatasourceHouseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDataMaskPolicyRequestParams struct {
	// 要绑定的数据脱敏策略权限对象集合
	DataMaskStrategyPolicySet []*DataMaskStrategyPolicy `json:"DataMaskStrategyPolicySet,omitnil,omitempty" name:"DataMaskStrategyPolicySet"`
}

type AttachDataMaskPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 要绑定的数据脱敏策略权限对象集合
	DataMaskStrategyPolicySet []*DataMaskStrategyPolicy `json:"DataMaskStrategyPolicySet,omitnil,omitempty" name:"DataMaskStrategyPolicySet"`
}

func (r *AttachDataMaskPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDataMaskPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataMaskStrategyPolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachDataMaskPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachDataMaskPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachDataMaskPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachDataMaskPolicyResponseParams `json:"Response"`
}

func (r *AttachDataMaskPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachDataMaskPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachUserPolicyRequestParams struct {
	// 用户Id，和子用户uin相同，需要先使用CreateUser接口创建用户。可以使用DescribeUsers接口查看。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

type AttachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和子用户uin相同，需要先使用CreateUser接口创建用户。可以使用DescribeUsers接口查看。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

func (r *AttachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachUserPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachUserPolicyResponseParams `json:"Response"`
}

func (r *AttachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachWorkGroupPolicyRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 要绑定的策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

type AttachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 要绑定的策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

func (r *AttachWorkGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachWorkGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachWorkGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachWorkGroupPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachWorkGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *AttachWorkGroupPolicyResponseParams `json:"Response"`
}

func (r *AttachWorkGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachWorkGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchSQLCostInfo struct {
	// 任务id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 本次消耗，单位cu
	Cost *float64 `json:"Cost,omitnil,omitempty" name:"Cost"`

	// 时间开销，秒
	TimeCost *int64 `json:"TimeCost,omitnil,omitempty" name:"TimeCost"`

	// 操作者
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type BatchSqlTask struct {
	// SQL子任务唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 运行SQL
	ExecuteSQL *string `json:"ExecuteSQL,omitnil,omitempty" name:"ExecuteSQL"`

	// 任务信息，成功则返回：Task Success!，失败则返回异常信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type BindWorkGroupsToUserRequestParams struct {
	// 绑定的用户和工作组信息
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

type BindWorkGroupsToUserRequest struct {
	*tchttp.BaseRequest
	
	// 绑定的用户和工作组信息
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

func (r *BindWorkGroupsToUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindWorkGroupsToUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindWorkGroupsToUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindWorkGroupsToUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindWorkGroupsToUserResponse struct {
	*tchttp.BaseResponse
	Response *BindWorkGroupsToUserResponseParams `json:"Response"`
}

func (r *BindWorkGroupsToUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindWorkGroupsToUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CHDFSProductVpcInfo struct {
	// vpc id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc名称
	VpcName *string `json:"VpcName,omitnil,omitempty" name:"VpcName"`

	// vpc子网信息列表
	VpcCidrBlock []*VpcCidrBlock `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 权限组Id
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

type CSV struct {
	// 压缩格式，["Snappy", "Gzip", "None"选一]。
	CodeCompress *string `json:"CodeCompress,omitnil,omitempty" name:"CodeCompress"`

	// CSV序列化及反序列化数据结构。
	CSVSerde *CSVSerde `json:"CSVSerde,omitnil,omitempty" name:"CSVSerde"`

	// 标题行，默认为0。
	HeadLines *int64 `json:"HeadLines,omitnil,omitempty" name:"HeadLines"`

	// 格式，默认值为CSV
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type CSVSerde struct {
	// CSV序列化转义符，默认为"\\"，最长8个字符，如 Escape: "/\"
	Escape *string `json:"Escape,omitnil,omitempty" name:"Escape"`

	// CSV序列化字段域符，默认为"'"，最长8个字符, 如 Quote: "\""
	Quote *string `json:"Quote,omitnil,omitempty" name:"Quote"`

	// CSV序列化分隔符，默认为"\t"，最长8个字符, 如 Separator: "\t"
	Separator *string `json:"Separator,omitnil,omitempty" name:"Separator"`
}

// Predefined struct for user
type CancelNotebookSessionStatementBatchRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

type CancelNotebookSessionStatementBatchRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

func (r *CancelNotebookSessionStatementBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelNotebookSessionStatementBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelNotebookSessionStatementBatchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelNotebookSessionStatementBatchResponse struct {
	*tchttp.BaseResponse
	Response *CancelNotebookSessionStatementBatchResponseParams `json:"Response"`
}

func (r *CancelNotebookSessionStatementBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelNotebookSessionStatementRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil,omitempty" name:"StatementId"`
}

type CancelNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil,omitempty" name:"StatementId"`
}

func (r *CancelNotebookSessionStatementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "StatementId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelNotebookSessionStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelNotebookSessionStatementResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelNotebookSessionStatementResponse struct {
	*tchttp.BaseResponse
	Response *CancelNotebookSessionStatementResponseParams `json:"Response"`
}

func (r *CancelNotebookSessionStatementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelNotebookSessionStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelSparkSessionBatchSQLRequestParams struct {
	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 用户自定义主键，若不为空，则使用该值进行查询
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`
}

type CancelSparkSessionBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 用户自定义主键，若不为空，则使用该值进行查询
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`
}

func (r *CancelSparkSessionBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelSparkSessionBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "CustomKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelSparkSessionBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelSparkSessionBatchSQLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelSparkSessionBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *CancelSparkSessionBatchSQLResponseParams `json:"Response"`
}

func (r *CancelSparkSessionBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelSparkSessionBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTaskRequestParams struct {
	// 任务Id，全局唯一
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type CancelTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id，全局唯一
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *CancelTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelTaskResponseParams `json:"Response"`
}

func (r *CancelTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTasksRequestParams struct {
	// 任务Id数组，全局唯一
	TaskId []*string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 配置信息，key-value数组，对外不可见。key1：AuthorityRole（鉴权角色，默认传SubUin，base64加密，仅在jdbc提交任务时使用）
	Config []*KVPair `json:"Config,omitnil,omitempty" name:"Config"`
}

type CancelTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id数组，全局唯一
	TaskId []*string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 配置信息，key-value数组，对外不可见。key1：AuthorityRole（鉴权角色，默认传SubUin，base64加密，仅在jdbc提交任务时使用）
	Config []*KVPair `json:"Config,omitnil,omitempty" name:"Config"`
}

func (r *CancelTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Config")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelTasksResponse struct {
	*tchttp.BaseResponse
	Response *CancelTasksResponseParams `json:"Response"`
}

func (r *CancelTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineConfigPairsValidityRequestParams struct {
	// 引擎小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`

	// 用户自定义参数
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil,omitempty" name:"DataEngineConfigPairs"`

	// 引擎大版本ID，存在小版本ID时仅需传入小版本ID，不存在时会获取当前大版本下最新的小版本ID。
	ImageVersionId *string `json:"ImageVersionId,omitnil,omitempty" name:"ImageVersionId"`
}

type CheckDataEngineConfigPairsValidityRequest struct {
	*tchttp.BaseRequest
	
	// 引擎小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`

	// 用户自定义参数
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil,omitempty" name:"DataEngineConfigPairs"`

	// 引擎大版本ID，存在小版本ID时仅需传入小版本ID，不存在时会获取当前大版本下最新的小版本ID。
	ImageVersionId *string `json:"ImageVersionId,omitnil,omitempty" name:"ImageVersionId"`
}

func (r *CheckDataEngineConfigPairsValidityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineConfigPairsValidityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChildImageVersionId")
	delete(f, "DataEngineConfigPairs")
	delete(f, "ImageVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDataEngineConfigPairsValidityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineConfigPairsValidityResponseParams struct {
	// 参数有效性：ture:有效，false:至少存在一个无效参数；
	IsAvailable *bool `json:"IsAvailable,omitnil,omitempty" name:"IsAvailable"`

	// 无效参数集合
	UnavailableConfig []*string `json:"UnavailableConfig,omitnil,omitempty" name:"UnavailableConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckDataEngineConfigPairsValidityResponse struct {
	*tchttp.BaseResponse
	Response *CheckDataEngineConfigPairsValidityResponseParams `json:"Response"`
}

func (r *CheckDataEngineConfigPairsValidityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineConfigPairsValidityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeRollbackRequestParams struct {
	// 引擎唯一id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

type CheckDataEngineImageCanBeRollbackRequest struct {
	*tchttp.BaseRequest
	
	// 引擎唯一id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

func (r *CheckDataEngineImageCanBeRollbackRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeRollbackRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDataEngineImageCanBeRollbackRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeRollbackResponseParams struct {
	// 回滚后日志记录id
	ToRecordId *string `json:"ToRecordId,omitnil,omitempty" name:"ToRecordId"`

	// 回滚前日志记录id
	FromRecordId *string `json:"FromRecordId,omitnil,omitempty" name:"FromRecordId"`

	// 是否能够回滚
	IsRollback *bool `json:"IsRollback,omitnil,omitempty" name:"IsRollback"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckDataEngineImageCanBeRollbackResponse struct {
	*tchttp.BaseResponse
	Response *CheckDataEngineImageCanBeRollbackResponseParams `json:"Response"`
}

func (r *CheckDataEngineImageCanBeRollbackResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeRollbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeUpgradeRequestParams struct {
	// 集群id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

type CheckDataEngineImageCanBeUpgradeRequest struct {
	*tchttp.BaseRequest
	
	// 集群id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

func (r *CheckDataEngineImageCanBeUpgradeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeUpgradeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckDataEngineImageCanBeUpgradeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckDataEngineImageCanBeUpgradeResponseParams struct {
	// 当前大版本下，可升级的集群镜像小版本id
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`

	// 是否能够升级
	IsUpgrade *bool `json:"IsUpgrade,omitnil,omitempty" name:"IsUpgrade"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckDataEngineImageCanBeUpgradeResponse struct {
	*tchttp.BaseResponse
	Response *CheckDataEngineImageCanBeUpgradeResponseParams `json:"Response"`
}

func (r *CheckDataEngineImageCanBeUpgradeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckDataEngineImageCanBeUpgradeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckLockMetaDataRequestParams struct {
	// 锁ID
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil,omitempty" name:"TxnId"`

	// 过期时间ms
	ElapsedMs *int64 `json:"ElapsedMs,omitnil,omitempty" name:"ElapsedMs"`
}

type CheckLockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 锁ID
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil,omitempty" name:"TxnId"`

	// 过期时间ms
	ElapsedMs *int64 `json:"ElapsedMs,omitnil,omitempty" name:"ElapsedMs"`
}

func (r *CheckLockMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckLockMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LockId")
	delete(f, "DatasourceConnectionName")
	delete(f, "TxnId")
	delete(f, "ElapsedMs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckLockMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckLockMetaDataResponseParams struct {
	// 锁ID
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 锁状态：ACQUIRED、WAITING、ABORT、NOT_ACQUIRED
	LockState *string `json:"LockState,omitnil,omitempty" name:"LockState"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckLockMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *CheckLockMetaDataResponseParams `json:"Response"`
}

func (r *CheckLockMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckLockMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {
	// 列名称，不区分大小写，最大支持25个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// string|tinyint|smallint|int|bigint|boolean|float|double|decimal|timestamp|date|binary|array|map|struct|uniontype
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 对该类的注释。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 表示整个 numeric 的长度
	Precision *int64 `json:"Precision,omitnil,omitempty" name:"Precision"`

	// 表示小数部分的长度
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// 是否为null
	Nullable *string `json:"Nullable,omitnil,omitempty" name:"Nullable"`

	// 字段位置，小的在前
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 字段创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 字段修改时间
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 是否为分区字段
	IsPartition *bool `json:"IsPartition,omitnil,omitempty" name:"IsPartition"`

	// 数据脱敏策略信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataMaskStrategyInfo *DataMaskStrategyInfo `json:"DataMaskStrategyInfo,omitnil,omitempty" name:"DataMaskStrategyInfo"`
}

type CommonMetrics struct {
	// 创建任务时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTaskTime *float64 `json:"CreateTaskTime,omitnil,omitempty" name:"CreateTaskTime"`

	// 预处理总时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessTime *float64 `json:"ProcessTime,omitnil,omitempty" name:"ProcessTime"`

	// 排队时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueTime *float64 `json:"QueueTime,omitnil,omitempty" name:"QueueTime"`

	// 执行时长，单位：ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionTime *float64 `json:"ExecutionTime,omitnil,omitempty" name:"ExecutionTime"`

	// 是否命中结果缓存
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsResultCacheHit *bool `json:"IsResultCacheHit,omitnil,omitempty" name:"IsResultCacheHit"`

	// 匹配物化视图数据量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchedMVBytes *int64 `json:"MatchedMVBytes,omitnil,omitempty" name:"MatchedMVBytes"`

	// 匹配物化视图列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchedMVs *string `json:"MatchedMVs,omitnil,omitempty" name:"MatchedMVs"`

	// 结果数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedBytes *string `json:"AffectedBytes,omitnil,omitempty" name:"AffectedBytes"`

	// 	结果行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AffectedRows *int64 `json:"AffectedRows,omitnil,omitempty" name:"AffectedRows"`

	// 扫描数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedBytes *int64 `json:"ProcessedBytes,omitnil,omitempty" name:"ProcessedBytes"`

	// 	扫描行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessedRows *int64 `json:"ProcessedRows,omitnil,omitempty" name:"ProcessedRows"`
}

type CoreInfo struct {
	// 时间戳(毫秒)数组
	Timestamp []*int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// core 用量
	CoreUsage []*int64 `json:"CoreUsage,omitnil,omitempty" name:"CoreUsage"`
}

type CosPermission struct {
	// cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// 权限【"read","write"】
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`
}

// Predefined struct for user
type CreateCHDFSBindingProductRequestParams struct {
	// 需要绑定的元数据加速桶名
	MountPoint *string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 桶的类型，分为cos和lakefs
	BucketType *string `json:"BucketType,omitnil,omitempty" name:"BucketType"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 引擎名称，ProductName选择DLC产品时，必传此参数。其他产品可不传
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// vpc信息，产品名称为other时必传此参数
	VpcInfo []*VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`
}

type CreateCHDFSBindingProductRequest struct {
	*tchttp.BaseRequest
	
	// 需要绑定的元数据加速桶名
	MountPoint *string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 桶的类型，分为cos和lakefs
	BucketType *string `json:"BucketType,omitnil,omitempty" name:"BucketType"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 引擎名称，ProductName选择DLC产品时，必传此参数。其他产品可不传
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// vpc信息，产品名称为other时必传此参数
	VpcInfo []*VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`
}

func (r *CreateCHDFSBindingProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCHDFSBindingProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPoint")
	delete(f, "BucketType")
	delete(f, "ProductName")
	delete(f, "EngineName")
	delete(f, "VpcInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCHDFSBindingProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCHDFSBindingProductResponseParams struct {
	// 绑定信息
	MountPointAssociates []*MountPointAssociates `json:"MountPointAssociates,omitnil,omitempty" name:"MountPointAssociates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCHDFSBindingProductResponse struct {
	*tchttp.BaseResponse
	Response *CreateCHDFSBindingProductResponseParams `json:"Response"`
}

func (r *CreateCHDFSBindingProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCHDFSBindingProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDMSDatabaseRequestParams struct {
	// 基础元数据对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// Schema目录
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// Db存储路径
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type CreateDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 基础元数据对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// Schema目录
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// Db存储路径
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *CreateDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Asset")
	delete(f, "SchemaName")
	delete(f, "Location")
	delete(f, "Name")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDMSDatabaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDMSDatabaseResponseParams `json:"Response"`
}

func (r *CreateDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDMSTableRequestParams struct {
	// 基础对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 表类型：EXTERNAL_TABLE, VIRTUAL_VIEW, MATERIALIZED_VIEW
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type CreateDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 基础对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 表类型：EXTERNAL_TABLE, VIRTUAL_VIEW, MATERIALIZED_VIEW
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil,omitempty" name:"LifeTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil,omitempty" name:"StructUpdateTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil,omitempty" name:"Sds"`

	// 列
	Columns []*DMSColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 分区键值
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil,omitempty" name:"PartitionKeys"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil,omitempty" name:"ViewExpandedText"`

	// 分区
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *CreateDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Asset")
	delete(f, "Type")
	delete(f, "DbName")
	delete(f, "StorageSize")
	delete(f, "RecordCount")
	delete(f, "LifeTime")
	delete(f, "DataUpdateTime")
	delete(f, "StructUpdateTime")
	delete(f, "LastAccessTime")
	delete(f, "Sds")
	delete(f, "Columns")
	delete(f, "PartitionKeys")
	delete(f, "ViewOriginalText")
	delete(f, "ViewExpandedText")
	delete(f, "Partitions")
	delete(f, "Name")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDMSTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateDMSTableResponseParams `json:"Response"`
}

func (r *CreateDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataEngineRequestParams struct {
	// 引擎类型spark/presto
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 集群类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 计费模式 0=共享模式 1=按量计费 2=包年包月
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 是否自动启动集群
	AutoResume *bool `json:"AutoResume,omitnil,omitempty" name:"AutoResume"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil,omitempty" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil,omitempty" name:"MaxClusters"`

	// 是否为默认虚拟集群
	//
	// Deprecated: DefaultDataEngine is deprecated.
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitnil,omitempty" name:"DefaultDataEngine"`

	// VPC网段
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 描述信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 集群规模
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 计费类型，后付费：0，预付费：1。当前只支持后付费，不填默认为后付费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 资源使用时长，后付费：固定填1，预付费：最少填1，代表购买资源一个月，最长不超过120。默认1
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 资源使用时长的单位，后付费：h，预付费：m。默认为h
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 资源的自动续费标志。后付费无需续费，固定填0；预付费下：0表示手动续费、1代表自动续费、2代表不续费，在0下如果是大客户，会自动帮大客户续费。默认为0
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 创建资源的时候需要绑定的标签信息
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，有效值：SQL/BATCH，标准引擎默认为BATCH
	EngineExecType *string `json:"EngineExecType,omitnil,omitempty" name:"EngineExecType"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil,omitempty" name:"TolerableQueueTime"`

	// 集群自动挂起时间，默认10分钟
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil,omitempty" name:"AutoSuspendTime"`

	// 资源类型。Standard_CU：标准型；Memory_CU：内存型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 集群高级配置
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil,omitempty" name:"DataEngineConfigPairs"`

	// 集群镜像版本名字。如SuperSQL-P 1.1;SuperSQL-S 3.2等,不传，默认创建最新镜像版本的集群
	ImageVersionName *string `json:"ImageVersionName,omitnil,omitempty" name:"ImageVersionName"`

	// 主集群名称，创建容灾集群时指定
	MainClusterName *string `json:"MainClusterName,omitnil,omitempty" name:"MainClusterName"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil,omitempty" name:"ElasticLimit"`

	// spark作业集群session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`

	// 自动授权
	AutoAuthorization *bool `json:"AutoAuthorization,omitnil,omitempty" name:"AutoAuthorization"`

	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 引擎世代，SuperSQL：代表supersql引擎，Native：代表标准引擎。默认值为SuperSQL
	EngineGeneration *string `json:"EngineGeneration,omitnil,omitempty" name:"EngineGeneration"`
}

type CreateDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型spark/presto
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 集群类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 计费模式 0=共享模式 1=按量计费 2=包年包月
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 是否自动启动集群
	AutoResume *bool `json:"AutoResume,omitnil,omitempty" name:"AutoResume"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil,omitempty" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil,omitempty" name:"MaxClusters"`

	// 是否为默认虚拟集群
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitnil,omitempty" name:"DefaultDataEngine"`

	// VPC网段
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 描述信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 集群规模
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 计费类型，后付费：0，预付费：1。当前只支持后付费，不填默认为后付费。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 资源使用时长，后付费：固定填1，预付费：最少填1，代表购买资源一个月，最长不超过120。默认1
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 资源使用时长的单位，后付费：h，预付费：m。默认为h
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 资源的自动续费标志。后付费无需续费，固定填0；预付费下：0表示手动续费、1代表自动续费、2代表不续费，在0下如果是大客户，会自动帮大客户续费。默认为0
	AutoRenew *int64 `json:"AutoRenew,omitnil,omitempty" name:"AutoRenew"`

	// 创建资源的时候需要绑定的标签信息
	Tags []*TagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，有效值：SQL/BATCH，标准引擎默认为BATCH
	EngineExecType *string `json:"EngineExecType,omitnil,omitempty" name:"EngineExecType"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil,omitempty" name:"TolerableQueueTime"`

	// 集群自动挂起时间，默认10分钟
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil,omitempty" name:"AutoSuspendTime"`

	// 资源类型。Standard_CU：标准型；Memory_CU：内存型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 集群高级配置
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil,omitempty" name:"DataEngineConfigPairs"`

	// 集群镜像版本名字。如SuperSQL-P 1.1;SuperSQL-S 3.2等,不传，默认创建最新镜像版本的集群
	ImageVersionName *string `json:"ImageVersionName,omitnil,omitempty" name:"ImageVersionName"`

	// 主集群名称，创建容灾集群时指定
	MainClusterName *string `json:"MainClusterName,omitnil,omitempty" name:"MainClusterName"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil,omitempty" name:"ElasticLimit"`

	// spark作业集群session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`

	// 自动授权
	AutoAuthorization *bool `json:"AutoAuthorization,omitnil,omitempty" name:"AutoAuthorization"`

	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 引擎世代，SuperSQL：代表supersql引擎，Native：代表标准引擎。默认值为SuperSQL
	EngineGeneration *string `json:"EngineGeneration,omitnil,omitempty" name:"EngineGeneration"`
}

func (r *CreateDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	delete(f, "DataEngineName")
	delete(f, "ClusterType")
	delete(f, "Mode")
	delete(f, "AutoResume")
	delete(f, "MinClusters")
	delete(f, "MaxClusters")
	delete(f, "DefaultDataEngine")
	delete(f, "CidrBlock")
	delete(f, "Message")
	delete(f, "Size")
	delete(f, "PayMode")
	delete(f, "TimeSpan")
	delete(f, "TimeUnit")
	delete(f, "AutoRenew")
	delete(f, "Tags")
	delete(f, "AutoSuspend")
	delete(f, "CrontabResumeSuspend")
	delete(f, "CrontabResumeSuspendStrategy")
	delete(f, "EngineExecType")
	delete(f, "MaxConcurrency")
	delete(f, "TolerableQueueTime")
	delete(f, "AutoSuspendTime")
	delete(f, "ResourceType")
	delete(f, "DataEngineConfigPairs")
	delete(f, "ImageVersionName")
	delete(f, "MainClusterName")
	delete(f, "ElasticSwitch")
	delete(f, "ElasticLimit")
	delete(f, "SessionResourceTemplate")
	delete(f, "AutoAuthorization")
	delete(f, "EngineNetworkId")
	delete(f, "EngineGeneration")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataEngineResponseParams struct {
	// 虚拟引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataEngineResponseParams `json:"Response"`
}

func (r *CreateDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataMaskStrategyRequestParams struct {
	// 数据脱敏策略详情
	Strategy *DataMaskStrategyInfo `json:"Strategy,omitnil,omitempty" name:"Strategy"`
}

type CreateDataMaskStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 数据脱敏策略详情
	Strategy *DataMaskStrategyInfo `json:"Strategy,omitnil,omitempty" name:"Strategy"`
}

func (r *CreateDataMaskStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataMaskStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Strategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataMaskStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataMaskStrategyResponseParams struct {
	// 策略id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDataMaskStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataMaskStrategyResponseParams `json:"Response"`
}

func (r *CreateDataMaskStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataMaskStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseRequestParams struct {
	// 数据库基础信息
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitnil,omitempty" name:"DatabaseInfo"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type CreateDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库基础信息
	DatabaseInfo *DatabaseInfo `json:"DatabaseInfo,omitnil,omitempty" name:"DatabaseInfo"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *CreateDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseInfo")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDatabaseResponseParams struct {
	// 生成的建库执行语句对象。
	Execution *Execution `json:"Execution,omitnil,omitempty" name:"Execution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateDatabaseResponseParams `json:"Response"`
}

func (r *CreateDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportTaskRequestParams struct {
	// 数据来源，lakefsStorage、taskResult
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 导出任务输入配置
	InputConf []*KVPair `json:"InputConf,omitnil,omitempty" name:"InputConf"`

	// 导出任务输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导出到cos
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`
}

type CreateExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据来源，lakefsStorage、taskResult
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 导出任务输入配置
	InputConf []*KVPair `json:"InputConf,omitnil,omitempty" name:"InputConf"`

	// 导出任务输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导出到cos
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`
}

func (r *CreateExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputType")
	delete(f, "InputConf")
	delete(f, "OutputConf")
	delete(f, "OutputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateExportTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateExportTaskResponseParams `json:"Response"`
}

func (r *CreateExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImportTaskRequestParams struct {
	// 数据来源，cos
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 输入配置
	InputConf []*KVPair `json:"InputConf,omitnil,omitempty" name:"InputConf"`

	// 输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导入到托管存储，即lakefsStorage
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`
}

type CreateImportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据来源，cos
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 输入配置
	InputConf []*KVPair `json:"InputConf,omitnil,omitempty" name:"InputConf"`

	// 输出配置
	OutputConf []*KVPair `json:"OutputConf,omitnil,omitempty" name:"OutputConf"`

	// 目标数据源的类型，目前支持导入到托管存储，即lakefsStorage
	OutputType *string `json:"OutputType,omitnil,omitempty" name:"OutputType"`
}

func (r *CreateImportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InputType")
	delete(f, "InputConf")
	delete(f, "OutputConf")
	delete(f, "OutputType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateImportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateImportTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateImportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateImportTaskResponseParams `json:"Response"`
}

func (r *CreateImportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateImportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalTableRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type CreateInternalTableRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

func (r *CreateInternalTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "Partitions")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInternalTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInternalTableResponseParams struct {
	// 创建托管存储内表sql语句描述
	Execution *string `json:"Execution,omitnil,omitempty" name:"Execution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateInternalTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateInternalTableResponseParams `json:"Response"`
}

func (r *CreateInternalTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInternalTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionRequestParams struct {
	// Session名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// session文件地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitnil,omitempty" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitnil,omitempty" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitnil,omitempty" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	ProgramArchives []*string `json:"ProgramArchives,omitnil,omitempty" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// Session相关配置，当前支持：
	// 1. dlc.eni: 用户配置的eni网关信息，可以通过该字段设置；
	// 2. dlc.role.arn: 用户配置的roleArn鉴权策略配置信息，可以通过该字段设置；
	// 3. dlc.sql.set.config: 用户配置的集群配置信息，可以通过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitnil,omitempty" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil,omitempty" name:"TimeoutInSecond"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 指定spark版本名称，当前任务使用该spark镜像运行
	SparkImage *string `json:"SparkImage,omitnil,omitempty" name:"SparkImage"`

	// 是否继承集群的资源类配置：0：自定义（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`
}

type CreateNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// session文件地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitnil,omitempty" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitnil,omitempty" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitnil,omitempty" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	ProgramArchives []*string `json:"ProgramArchives,omitnil,omitempty" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// Session相关配置，当前支持：
	// 1. dlc.eni: 用户配置的eni网关信息，可以通过该字段设置；
	// 2. dlc.role.arn: 用户配置的roleArn鉴权策略配置信息，可以通过该字段设置；
	// 3. dlc.sql.set.config: 用户配置的集群配置信息，可以通过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitnil,omitempty" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil,omitempty" name:"TimeoutInSecond"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 指定spark版本名称，当前任务使用该spark镜像运行
	SparkImage *string `json:"SparkImage,omitnil,omitempty" name:"SparkImage"`

	// 是否继承集群的资源类配置：0：自定义（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`
}

func (r *CreateNotebookSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Kind")
	delete(f, "DataEngineName")
	delete(f, "ProgramDependentFiles")
	delete(f, "ProgramDependentJars")
	delete(f, "ProgramDependentPython")
	delete(f, "ProgramArchives")
	delete(f, "DriverSize")
	delete(f, "ExecutorSize")
	delete(f, "ExecutorNumbers")
	delete(f, "Arguments")
	delete(f, "ProxyUser")
	delete(f, "TimeoutInSecond")
	delete(f, "ExecutorMaxNumbers")
	delete(f, "SparkImage")
	delete(f, "IsInherit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionResponseParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Spark任务返回的AppId
	SparkAppId *string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotebookSessionResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookSessionResponseParams `json:"Response"`
}

func (r *CreateNotebookSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`
}

type CreateNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`
}

func (r *CreateNotebookSessionStatementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Code")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookSessionStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementResponseParams struct {
	// Session Statement详情
	NotebookSessionStatement *NotebookSessionStatementInfo `json:"NotebookSessionStatement,omitnil,omitempty" name:"NotebookSessionStatement"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotebookSessionStatementResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookSessionStatementResponseParams `json:"Response"`
}

func (r *CreateNotebookSessionStatementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementSupportBatchSQLRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 类型，当前支持：sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 是否保存运行结果
	SaveResult *bool `json:"SaveResult,omitnil,omitempty" name:"SaveResult"`
}

type CreateNotebookSessionStatementSupportBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 执行的代码
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 类型，当前支持：sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 是否保存运行结果
	SaveResult *bool `json:"SaveResult,omitnil,omitempty" name:"SaveResult"`
}

func (r *CreateNotebookSessionStatementSupportBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementSupportBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Code")
	delete(f, "Kind")
	delete(f, "SaveResult")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotebookSessionStatementSupportBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotebookSessionStatementSupportBatchSQLResponseParams struct {
	// Session Statement详情
	NotebookSessionStatementBatches *NotebookSessionStatementBatchInformation `json:"NotebookSessionStatementBatches,omitnil,omitempty" name:"NotebookSessionStatementBatches"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotebookSessionStatementSupportBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotebookSessionStatementSupportBatchSQLResponseParams `json:"Response"`
}

func (r *CreateNotebookSessionStatementSupportBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotebookSessionStatementSupportBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResultDownloadRequestParams struct {
	// 查询结果任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 下载格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否重新生成下载文件，仅当之前任务状态为 timeout | error 时有效
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

type CreateResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 查询结果任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 下载格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 是否重新生成下载文件，仅当之前任务状态为 timeout | error 时有效
	Force *bool `json:"Force,omitnil,omitempty" name:"Force"`
}

func (r *CreateResultDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResultDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Format")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateResultDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateResultDownloadResponseParams struct {
	// 下载任务Id
	DownloadId *string `json:"DownloadId,omitnil,omitempty" name:"DownloadId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateResultDownloadResponse struct {
	*tchttp.BaseResponse
	Response *CreateResultDownloadResponseParams `json:"Response"`
}

func (r *CreateResultDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateResultDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScriptRequestParams struct {
	// 脚本名称，最大不能超过255个字符。
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitnil,omitempty" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitnil,omitempty" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`
}

type CreateScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本名称，最大不能超过255个字符。
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// base64编码后的sql语句
	SQLStatement *string `json:"SQLStatement,omitnil,omitempty" name:"SQLStatement"`

	// 脚本描述， 不能超过50个字符
	ScriptDesc *string `json:"ScriptDesc,omitnil,omitempty" name:"ScriptDesc"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`
}

func (r *CreateScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptName")
	delete(f, "SQLStatement")
	delete(f, "ScriptDesc")
	delete(f, "DatabaseName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScriptResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScriptResponse struct {
	*tchttp.BaseResponse
	Response *CreateScriptResponseParams `json:"Response"`
}

func (r *CreateScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppRequestParams struct {
	// spark作业名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil,omitempty" name:"AppFile"`

	// 数据访问策略，CAM Role arn，控制台通过数据作业—>作业配置获取，SDK通过DescribeUserRoles接口获取对应的值；
	RoleArn *int64 `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil,omitempty" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil,omitempty" name:"AppExecutorNums"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil,omitempty" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil,omitempty" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil,omitempty" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil,omitempty" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil,omitempty" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil,omitempty" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil,omitempty" name:"AppFiles"`

	// spark作业程序入参，空格分割
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil,omitempty" name:"MaxRetries"`

	// 数据源名称
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil,omitempty" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil,omitempty" name:"AppPythonFiles"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil,omitempty" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil,omitempty" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil,omitempty" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil,omitempty" name:"IsSessionStarted"`

	// 依赖包信息
	DependencyPackages []*DependencyPackage `json:"DependencyPackages,omitnil,omitempty" name:"DependencyPackages"`
}

type CreateSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil,omitempty" name:"AppFile"`

	// 数据访问策略，CAM Role arn，控制台通过数据作业—>作业配置获取，SDK通过DescribeUserRoles接口获取对应的值；
	RoleArn *int64 `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil,omitempty" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil,omitempty" name:"AppExecutorNums"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil,omitempty" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil,omitempty" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil,omitempty" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil,omitempty" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil,omitempty" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil,omitempty" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil,omitempty" name:"AppFiles"`

	// spark作业程序入参，空格分割
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil,omitempty" name:"MaxRetries"`

	// 数据源名称
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil,omitempty" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil,omitempty" name:"AppPythonFiles"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil,omitempty" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil,omitempty" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil,omitempty" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil,omitempty" name:"IsSessionStarted"`

	// 依赖包信息
	DependencyPackages []*DependencyPackage `json:"DependencyPackages,omitnil,omitempty" name:"DependencyPackages"`
}

func (r *CreateSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	delete(f, "IsLocalArchives")
	delete(f, "AppArchives")
	delete(f, "SparkImage")
	delete(f, "SparkImageVersion")
	delete(f, "AppExecutorMaxNumbers")
	delete(f, "SessionId")
	delete(f, "IsInherit")
	delete(f, "IsSessionStarted")
	delete(f, "DependencyPackages")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppResponseParams struct {
	// App唯一标识
	SparkAppId *string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkAppResponseParams `json:"Response"`
}

func (r *CreateSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppTaskRequestParams struct {
	// spark作业名
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// spark作业程序入参，以空格分隔；一般用于周期性调用使用
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

type CreateSparkAppTaskRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// spark作业程序入参，以空格分隔；一般用于周期性调用使用
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

func (r *CreateSparkAppTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobName")
	delete(f, "CmdArgs")
	delete(f, "SourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkAppTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkAppTaskResponseParams struct {
	// 批Id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSparkAppTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkAppTaskResponseParams `json:"Response"`
}

func (r *CreateSparkAppTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkAppTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkSessionBatchSQLRequestParams struct {
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 运行sql，需要base64编码。
	ExecuteSQL *string `json:"ExecuteSQL,omitnil,omitempty" name:"ExecuteSQL"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil,omitempty" name:"TimeoutInSecond"`

	// Session唯一标识，当指定sessionid，则使用该session运行任务。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 指定要创建的session名称
	SessionName *string `json:"SessionName,omitnil,omitempty" name:"SessionName"`

	// Session相关配置，当前支持：1.dlc.eni：用户配置的eni网关信息，可以用过该字段设置；
	// 2.dlc.role.arn：用户配置的roleArn鉴权策略配置信息，可以用过该字段设置；
	// 3.dlc.sql.set.config：用户配置的集群配置信息，可以用过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 是否继承集群的资源类配置：0：不继承（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// 用户自定义主键，需唯一
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

type CreateSparkSessionBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 运行sql，需要base64编码。
	ExecuteSQL *string `json:"ExecuteSQL,omitnil,omitempty" name:"ExecuteSQL"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil,omitempty" name:"TimeoutInSecond"`

	// Session唯一标识，当指定sessionid，则使用该session运行任务。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 指定要创建的session名称
	SessionName *string `json:"SessionName,omitnil,omitempty" name:"SessionName"`

	// Session相关配置，当前支持：1.dlc.eni：用户配置的eni网关信息，可以用过该字段设置；
	// 2.dlc.role.arn：用户配置的roleArn鉴权策略配置信息，可以用过该字段设置；
	// 3.dlc.sql.set.config：用户配置的集群配置信息，可以用过该字段设置；
	Arguments []*KVPair `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 是否继承集群的资源类配置：0：不继承（默认），1：继承集群；
	IsInherit *int64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// 用户自定义主键，需唯一
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

func (r *CreateSparkSessionBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkSessionBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "ExecuteSQL")
	delete(f, "DriverSize")
	delete(f, "ExecutorSize")
	delete(f, "ExecutorNumbers")
	delete(f, "ExecutorMaxNumbers")
	delete(f, "TimeoutInSecond")
	delete(f, "SessionId")
	delete(f, "SessionName")
	delete(f, "Arguments")
	delete(f, "IsInherit")
	delete(f, "CustomKey")
	delete(f, "SourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkSessionBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkSessionBatchSQLResponseParams struct {
	// 批任务唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// Statement任务列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Statements []*StatementInformation `json:"Statements,omitnil,omitempty" name:"Statements"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSparkSessionBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkSessionBatchSQLResponseParams `json:"Response"`
}

func (r *CreateSparkSessionBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkSessionBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkSubmitTaskRequestParams struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型：当前支持1: BatchType, 2: StreamingType, 4: SQLType
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 引擎名称，当前仅支持Spark批作业集群
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 指定运行的程序脚本路径，当前仅支持jar和py，对于SQLType该值设为空字符串
	PackagePath *string `json:"PackagePath,omitnil,omitempty" name:"PackagePath"`

	// 指定的鉴权信息
	RoleArn *int64 `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 运行任务所需资源是否继承自集群上配置资源信息，0（默认，不继承）、1（继承，当设置为该值，则任务级资源配置可不额外指定）
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// jar任务时需要指定主程序
	MainClass *string `json:"MainClass,omitnil,omitempty" name:"MainClass"`

	// 当前DriverSize规格仅支持（内存型集群则使用m前缀的枚举值）: small/medium/large/xlarge/m.small/m.medium/m.large/m.xlarge
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// 当前ExecutorSize规格仅支持（内存型集群则使用m前缀的枚举值）: small/medium/large/xlarge/m.small/m.medium/m.large/m.xlarge
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定使用的executor数量，最小为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// 指定使用的executor最大数量, 当该值大于ExecutorNums则自动开启动态
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 提交任务的附加配置集合，当前支持Key包含：MAINARGS：程序入口参数，空格分割(SqlType任务通过该值指定base64加密后的sql)、SPARKCONFIG：Spark配置，以换行符分隔、ENI：Eni连接信息、DEPENDENCYPACKAGEPATH：依赖的程序包（--jars、--py-files:支持py/zip/egg等归档格式），多文件以逗号分隔、DEPENDENCYFILEPATH：依赖文件资源（--files: 非jar、zip），多文件以逗号分隔、DEPENDENCYARCHIVESPATH：依赖archives资源（--archives: 支持tar.gz/tgz/tar等归档格式)，多文件以逗号分隔、MAXRETRIES：任务重试次数，非流任务默认为1、SPARKIMAGE：Spark镜像版本号，支持使用dlc镜像/用户自定的tcr镜像运行任务、SPARKIMAGEVERSION：Spark镜像版本名称，与SPARKIMAGE一一对应；SPARKPRESETCODE：base64后的notebook预置代码；SPARKENV：base64后的spark环境变量；SPARKGITINFO：base64后的git相关信息
	CmdArgs []*KVPair `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// ai资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

type CreateSparkSubmitTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务类型：当前支持1: BatchType, 2: StreamingType, 4: SQLType
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 引擎名称，当前仅支持Spark批作业集群
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 指定运行的程序脚本路径，当前仅支持jar和py，对于SQLType该值设为空字符串
	PackagePath *string `json:"PackagePath,omitnil,omitempty" name:"PackagePath"`

	// 指定的鉴权信息
	RoleArn *int64 `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 运行任务所需资源是否继承自集群上配置资源信息，0（默认，不继承）、1（继承，当设置为该值，则任务级资源配置可不额外指定）
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// jar任务时需要指定主程序
	MainClass *string `json:"MainClass,omitnil,omitempty" name:"MainClass"`

	// 当前DriverSize规格仅支持（内存型集群则使用m前缀的枚举值）: small/medium/large/xlarge/m.small/m.medium/m.large/m.xlarge
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// 当前ExecutorSize规格仅支持（内存型集群则使用m前缀的枚举值）: small/medium/large/xlarge/m.small/m.medium/m.large/m.xlarge
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定使用的executor数量，最小为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// 指定使用的executor最大数量, 当该值大于ExecutorNums则自动开启动态
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 提交任务的附加配置集合，当前支持Key包含：MAINARGS：程序入口参数，空格分割(SqlType任务通过该值指定base64加密后的sql)、SPARKCONFIG：Spark配置，以换行符分隔、ENI：Eni连接信息、DEPENDENCYPACKAGEPATH：依赖的程序包（--jars、--py-files:支持py/zip/egg等归档格式），多文件以逗号分隔、DEPENDENCYFILEPATH：依赖文件资源（--files: 非jar、zip），多文件以逗号分隔、DEPENDENCYARCHIVESPATH：依赖archives资源（--archives: 支持tar.gz/tgz/tar等归档格式)，多文件以逗号分隔、MAXRETRIES：任务重试次数，非流任务默认为1、SPARKIMAGE：Spark镜像版本号，支持使用dlc镜像/用户自定的tcr镜像运行任务、SPARKIMAGEVERSION：Spark镜像版本名称，与SPARKIMAGE一一对应；SPARKPRESETCODE：base64后的notebook预置代码；SPARKENV：base64后的spark环境变量；SPARKGITINFO：base64后的git相关信息
	CmdArgs []*KVPair `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// ai资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

func (r *CreateSparkSubmitTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkSubmitTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "TaskType")
	delete(f, "DataEngineName")
	delete(f, "PackagePath")
	delete(f, "RoleArn")
	delete(f, "IsInherit")
	delete(f, "MainClass")
	delete(f, "DriverSize")
	delete(f, "ExecutorSize")
	delete(f, "ExecutorNumbers")
	delete(f, "ExecutorMaxNumbers")
	delete(f, "CmdArgs")
	delete(f, "SourceInfo")
	delete(f, "ResourceGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSparkSubmitTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSparkSubmitTaskResponseParams struct {
	// 批作业ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 批任务ID，用改ID进行任务的查询与删除等
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSparkSubmitTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSparkSubmitTaskResponseParams `json:"Response"`
}

func (r *CreateSparkSubmitTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSparkSubmitTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStandardEngineResourceGroupRequestParams struct {
	// 标准引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 标准引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 自动启动，（任务提交自动拉起资源组）0-自动启动，1-不自动启动
	AutoLaunch *int64 `json:"AutoLaunch,omitnil,omitempty" name:"AutoLaunch"`

	// 自动挂起资源组。0-自动挂起，1-不自动挂起
	AutoPause *int64 `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// driver的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	DriverCuSpec *string `json:"DriverCuSpec,omitnil,omitempty" name:"DriverCuSpec"`

	// executor的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	ExecutorCuSpec *string `json:"ExecutorCuSpec,omitnil,omitempty" name:"ExecutorCuSpec"`

	// executor最小数量，
	MinExecutorNums *int64 `json:"MinExecutorNums,omitnil,omitempty" name:"MinExecutorNums"`

	// executor最大数量
	MaxExecutorNums *int64 `json:"MaxExecutorNums,omitnil,omitempty" name:"MaxExecutorNums"`

	// 创建资源组后是否直接拉起，0-拉起，1-不拉起
	IsLaunchNow *int64 `json:"IsLaunchNow,omitnil,omitempty" name:"IsLaunchNow"`

	// 自动挂起时间，单位分钟，取值范围在1-999（在无任务AutoPauseTime后，资源组自动挂起）
	AutoPauseTime *int64 `json:"AutoPauseTime,omitnil,omitempty" name:"AutoPauseTime"`

	// 资源组静态参数，需要重启资源组生效
	StaticConfigPairs []*EngineResourceGroupConfigPair `json:"StaticConfigPairs,omitnil,omitempty" name:"StaticConfigPairs"`

	// 资源组动态参数，下一个任务生效。
	DynamicConfigPairs []*EngineResourceGroupConfigPair `json:"DynamicConfigPairs,omitnil,omitempty" name:"DynamicConfigPairs"`

	// 任务并发数，默人是5个
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 网络配置名称
	NetworkConfigNames []*string `json:"NetworkConfigNames,omitnil,omitempty" name:"NetworkConfigNames"`

	// 自定义镜像域名
	PublicDomain *string `json:"PublicDomain,omitnil,omitempty" name:"PublicDomain"`

	// 自定义镜像实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// AI类型资源组的框架类型，machine-learning，python，spark-ml，不填默认为machine-learning
	FrameType *string `json:"FrameType,omitnil,omitempty" name:"FrameType"`

	// 镜像类型，bulit-in：内置，custom：自定义，不填默认为bulit-in
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像id
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// AI资源组有效，资源组可用资源上限，该值需要小于引擎资源上限
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 资源组场景
	ResourceGroupScene *string `json:"ResourceGroupScene,omitnil,omitempty" name:"ResourceGroupScene"`

	// 自定义镜像所在地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// python类型资源组python单机节点资源上限，该值要小于资源组的资源上限.small:1cu medium:2cu large:4cu xlarge:8cu 4xlarge:16cu 8xlarge:32cu 16xlarge:64cu，如果是高内存型资源，在类型前面加上m.
	PythonCuSpec *string `json:"PythonCuSpec,omitnil,omitempty" name:"PythonCuSpec"`

	// 仅SQL资源组资源配置模式，fast：快速模式，custom：自定义模式
	SparkSpecMode *string `json:"SparkSpecMode,omitnil,omitempty" name:"SparkSpecMode"`

	// 仅SQL资源组资源上限，仅用于快速模块
	SparkSize *int64 `json:"SparkSize,omitnil,omitempty" name:"SparkSize"`

	// GPUDriver规格
	DriverGPUSpec *int64 `json:"DriverGPUSpec,omitnil,omitempty" name:"DriverGPUSpec"`

	// GPUExecutor规格
	ExecutorGPUSpec *int64 `json:"ExecutorGPUSpec,omitnil,omitempty" name:"ExecutorGPUSpec"`

	// GPU上限
	GPULimitSize *int64 `json:"GPULimitSize,omitnil,omitempty" name:"GPULimitSize"`

	// GPU规格
	GPUSize *int64 `json:"GPUSize,omitnil,omitempty" name:"GPUSize"`

	// Pod GPU规格上限
	PythonGPUSpec *int64 `json:"PythonGPUSpec,omitnil,omitempty" name:"PythonGPUSpec"`
}

type CreateStandardEngineResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 标准引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 标准引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 自动启动，（任务提交自动拉起资源组）0-自动启动，1-不自动启动
	AutoLaunch *int64 `json:"AutoLaunch,omitnil,omitempty" name:"AutoLaunch"`

	// 自动挂起资源组。0-自动挂起，1-不自动挂起
	AutoPause *int64 `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// driver的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	DriverCuSpec *string `json:"DriverCuSpec,omitnil,omitempty" name:"DriverCuSpec"`

	// executor的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	ExecutorCuSpec *string `json:"ExecutorCuSpec,omitnil,omitempty" name:"ExecutorCuSpec"`

	// executor最小数量，
	MinExecutorNums *int64 `json:"MinExecutorNums,omitnil,omitempty" name:"MinExecutorNums"`

	// executor最大数量
	MaxExecutorNums *int64 `json:"MaxExecutorNums,omitnil,omitempty" name:"MaxExecutorNums"`

	// 创建资源组后是否直接拉起，0-拉起，1-不拉起
	IsLaunchNow *int64 `json:"IsLaunchNow,omitnil,omitempty" name:"IsLaunchNow"`

	// 自动挂起时间，单位分钟，取值范围在1-999（在无任务AutoPauseTime后，资源组自动挂起）
	AutoPauseTime *int64 `json:"AutoPauseTime,omitnil,omitempty" name:"AutoPauseTime"`

	// 资源组静态参数，需要重启资源组生效
	StaticConfigPairs []*EngineResourceGroupConfigPair `json:"StaticConfigPairs,omitnil,omitempty" name:"StaticConfigPairs"`

	// 资源组动态参数，下一个任务生效。
	DynamicConfigPairs []*EngineResourceGroupConfigPair `json:"DynamicConfigPairs,omitnil,omitempty" name:"DynamicConfigPairs"`

	// 任务并发数，默人是5个
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 网络配置名称
	NetworkConfigNames []*string `json:"NetworkConfigNames,omitnil,omitempty" name:"NetworkConfigNames"`

	// 自定义镜像域名
	PublicDomain *string `json:"PublicDomain,omitnil,omitempty" name:"PublicDomain"`

	// 自定义镜像实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// AI类型资源组的框架类型，machine-learning，python，spark-ml，不填默认为machine-learning
	FrameType *string `json:"FrameType,omitnil,omitempty" name:"FrameType"`

	// 镜像类型，bulit-in：内置，custom：自定义，不填默认为bulit-in
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像id
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// AI资源组有效，资源组可用资源上限，该值需要小于引擎资源上限
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 资源组场景
	ResourceGroupScene *string `json:"ResourceGroupScene,omitnil,omitempty" name:"ResourceGroupScene"`

	// 自定义镜像所在地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// python类型资源组python单机节点资源上限，该值要小于资源组的资源上限.small:1cu medium:2cu large:4cu xlarge:8cu 4xlarge:16cu 8xlarge:32cu 16xlarge:64cu，如果是高内存型资源，在类型前面加上m.
	PythonCuSpec *string `json:"PythonCuSpec,omitnil,omitempty" name:"PythonCuSpec"`

	// 仅SQL资源组资源配置模式，fast：快速模式，custom：自定义模式
	SparkSpecMode *string `json:"SparkSpecMode,omitnil,omitempty" name:"SparkSpecMode"`

	// 仅SQL资源组资源上限，仅用于快速模块
	SparkSize *int64 `json:"SparkSize,omitnil,omitempty" name:"SparkSize"`

	// GPUDriver规格
	DriverGPUSpec *int64 `json:"DriverGPUSpec,omitnil,omitempty" name:"DriverGPUSpec"`

	// GPUExecutor规格
	ExecutorGPUSpec *int64 `json:"ExecutorGPUSpec,omitnil,omitempty" name:"ExecutorGPUSpec"`

	// GPU上限
	GPULimitSize *int64 `json:"GPULimitSize,omitnil,omitempty" name:"GPULimitSize"`

	// GPU规格
	GPUSize *int64 `json:"GPUSize,omitnil,omitempty" name:"GPUSize"`

	// Pod GPU规格上限
	PythonGPUSpec *int64 `json:"PythonGPUSpec,omitnil,omitempty" name:"PythonGPUSpec"`
}

func (r *CreateStandardEngineResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStandardEngineResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupName")
	delete(f, "DataEngineName")
	delete(f, "AutoLaunch")
	delete(f, "AutoPause")
	delete(f, "DriverCuSpec")
	delete(f, "ExecutorCuSpec")
	delete(f, "MinExecutorNums")
	delete(f, "MaxExecutorNums")
	delete(f, "IsLaunchNow")
	delete(f, "AutoPauseTime")
	delete(f, "StaticConfigPairs")
	delete(f, "DynamicConfigPairs")
	delete(f, "MaxConcurrency")
	delete(f, "NetworkConfigNames")
	delete(f, "PublicDomain")
	delete(f, "RegistryId")
	delete(f, "FrameType")
	delete(f, "ImageType")
	delete(f, "ImageName")
	delete(f, "ImageVersion")
	delete(f, "Size")
	delete(f, "ResourceGroupScene")
	delete(f, "RegionName")
	delete(f, "PythonCuSpec")
	delete(f, "SparkSpecMode")
	delete(f, "SparkSize")
	delete(f, "DriverGPUSpec")
	delete(f, "ExecutorGPUSpec")
	delete(f, "GPULimitSize")
	delete(f, "GPUSize")
	delete(f, "PythonGPUSpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStandardEngineResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStandardEngineResourceGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStandardEngineResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateStandardEngineResourceGroupResponseParams `json:"Response"`
}

func (r *CreateStandardEngineResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStandardEngineResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStoreLocationRequestParams struct {
	// 计算结果存储cos路径，如：cosn://bucketname/
	StoreLocation *string `json:"StoreLocation,omitnil,omitempty" name:"StoreLocation"`
}

type CreateStoreLocationRequest struct {
	*tchttp.BaseRequest
	
	// 计算结果存储cos路径，如：cosn://bucketname/
	StoreLocation *string `json:"StoreLocation,omitnil,omitempty" name:"StoreLocation"`
}

func (r *CreateStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StoreLocation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateStoreLocationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *CreateStoreLocationResponseParams `json:"Response"`
}

func (r *CreateStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableRequestParams struct {
	// 数据表配置信息
	TableInfo *TableInfo `json:"TableInfo,omitnil,omitempty" name:"TableInfo"`
}

type CreateTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据表配置信息
	TableInfo *TableInfo `json:"TableInfo,omitnil,omitempty" name:"TableInfo"`
}

func (r *CreateTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTableResponseParams struct {
	// 生成的建表执行语句对象。
	Execution *Execution `json:"Execution,omitnil,omitempty" name:"Execution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateTableResponseParams `json:"Response"`
}

func (r *CreateTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 计算任务，该参数中包含任务类型及其相关配置信息
	Task *Task `json:"Task,omitnil,omitempty" name:"Task"`

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 默认数据源名称。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 数据引擎名称，不填提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 标准spark执行任务resourceGroupName
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 计算任务，该参数中包含任务类型及其相关配置信息
	Task *Task `json:"Task,omitnil,omitempty" name:"Task"`

	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 默认数据源名称。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 数据引擎名称，不填提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 标准spark执行任务resourceGroupName
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

func (r *CreateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Task")
	delete(f, "DatabaseName")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	delete(f, "ResourceGroupName")
	delete(f, "SourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskResponseParams `json:"Response"`
}

func (r *CreateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksInOrderRequestParams struct {
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 数据源名称，默认为COSDataCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type CreateTasksInOrderRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 数据源名称，默认为COSDataCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *CreateTasksInOrderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksInOrderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Tasks")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTasksInOrderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksInOrderResponseParams struct {
	// 本批次提交的任务的批次Id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 任务Id集合，按照执行顺序排列
	TaskIdSet []*string `json:"TaskIdSet,omitnil,omitempty" name:"TaskIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTasksInOrderResponse struct {
	*tchttp.BaseResponse
	Response *CreateTasksInOrderResponseParams `json:"Response"`
}

func (r *CreateTasksInOrderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksInOrderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksRequestParams struct {
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 计算引擎名称，不填任务提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// spark集群资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 是否使用multi- statement方式运行一批次任务，true: 是，false: 否
	IsMultiStatement *bool `json:"IsMultiStatement,omitnil,omitempty" name:"IsMultiStatement"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

type CreateTasksRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称。如果SQL语句中有数据库名称，优先使用SQL语句中的数据库，否则使用该参数指定的数据库（注：当提交建库sql时，该字段传空字符串）。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// SQL任务信息
	Tasks *TasksInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 计算引擎名称，不填任务提交到默认集群
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// spark集群资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 是否使用multi- statement方式运行一批次任务，true: 是，false: 否
	IsMultiStatement *bool `json:"IsMultiStatement,omitnil,omitempty" name:"IsMultiStatement"`

	// 任务来源信息
	SourceInfo []*KVPair `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

func (r *CreateTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Tasks")
	delete(f, "DatasourceConnectionName")
	delete(f, "DataEngineName")
	delete(f, "ResourceGroupName")
	delete(f, "IsMultiStatement")
	delete(f, "SourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTasksResponseParams struct {
	// 本批次提交的任务的批次Id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 任务Id集合，按照执行顺序排列
	TaskIdSet []*string `json:"TaskIdSet,omitnil,omitempty" name:"TaskIdSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTasksResponse struct {
	*tchttp.BaseResponse
	Response *CreateTasksResponseParams `json:"Response"`
}

func (r *CreateTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTcIcebergTableRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 为true时只获取sql而不执行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type CreateTcIcebergTableRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 为true时只获取sql而不执行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`
}

func (r *CreateTcIcebergTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTcIcebergTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "DryRun")
	delete(f, "Partitions")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTcIcebergTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTcIcebergTableResponseParams struct {
	// amoro的SessionId
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 执行的sql
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`

	// 为true时只返回sql而不实际执行
	DryRun *bool `json:"DryRun,omitnil,omitempty" name:"DryRun"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTcIcebergTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateTcIcebergTableResponseParams `json:"Response"`
}

func (r *CreateTcIcebergTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTcIcebergTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserRequestParams struct {
	// 需要授权的子用户uin，可以通过腾讯云控制台右上角 → “账号信息” → “账号ID进行查看”。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`

	// 绑定到用户的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 用户类型。ADMIN：管理员 COMMON：一般用户。当用户类型为管理员的时候，不能设置权限集合和绑定的工作组集合，管理员默认拥有所有权限。该参数不填默认为COMMON
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 绑定到用户的工作组ID集合。
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil,omitempty" name:"WorkGroupIds"`

	// 用户别名，字符长度小50
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`
}

type CreateUserRequest struct {
	*tchttp.BaseRequest
	
	// 需要授权的子用户uin，可以通过腾讯云控制台右上角 → “账号信息” → “账号ID进行查看”。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`

	// 绑定到用户的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 用户类型。ADMIN：管理员 COMMON：一般用户。当用户类型为管理员的时候，不能设置权限集合和绑定的工作组集合，管理员默认拥有所有权限。该参数不填默认为COMMON
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 绑定到用户的工作组ID集合。
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil,omitempty" name:"WorkGroupIds"`

	// 用户别名，字符长度小50
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`
}

func (r *CreateUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserDescription")
	delete(f, "PolicySet")
	delete(f, "UserType")
	delete(f, "WorkGroupIds")
	delete(f, "UserAlias")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserResponseParams `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserVpcConnectionRequestParams struct {
	// 用户vpcid
	UserVpcId *string `json:"UserVpcId,omitnil,omitempty" name:"UserVpcId"`

	// 用户子网
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 用户终端节点名称
	UserVpcEndpointName *string `json:"UserVpcEndpointName,omitnil,omitempty" name:"UserVpcEndpointName"`

	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 手动指定vip，不填自动分配子网下的一个ip
	UserVpcEndpointVip *string `json:"UserVpcEndpointVip,omitnil,omitempty" name:"UserVpcEndpointVip"`
}

type CreateUserVpcConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 用户vpcid
	UserVpcId *string `json:"UserVpcId,omitnil,omitempty" name:"UserVpcId"`

	// 用户子网
	UserSubnetId *string `json:"UserSubnetId,omitnil,omitempty" name:"UserSubnetId"`

	// 用户终端节点名称
	UserVpcEndpointName *string `json:"UserVpcEndpointName,omitnil,omitempty" name:"UserVpcEndpointName"`

	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 手动指定vip，不填自动分配子网下的一个ip
	UserVpcEndpointVip *string `json:"UserVpcEndpointVip,omitnil,omitempty" name:"UserVpcEndpointVip"`
}

func (r *CreateUserVpcConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserVpcConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserVpcId")
	delete(f, "UserSubnetId")
	delete(f, "UserVpcEndpointName")
	delete(f, "EngineNetworkId")
	delete(f, "UserVpcEndpointVip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserVpcConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserVpcConnectionResponseParams struct {
	// 终端节点IP
	UserVpcEndpointId *string `json:"UserVpcEndpointId,omitnil,omitempty" name:"UserVpcEndpointId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserVpcConnectionResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserVpcConnectionResponseParams `json:"Response"`
}

func (r *CreateUserVpcConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserVpcConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkGroupRequestParams struct {
	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil,omitempty" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil,omitempty" name:"WorkGroupDescription"`

	// 工作组绑定的鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 需要绑定到工作组的用户Id集合
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type CreateWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil,omitempty" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil,omitempty" name:"WorkGroupDescription"`

	// 工作组绑定的鉴权策略集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 需要绑定到工作组的用户Id集合
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

func (r *CreateWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupName")
	delete(f, "WorkGroupDescription")
	delete(f, "PolicySet")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkGroupResponseParams struct {
	// 工作组Id，全局唯一
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkGroupResponseParams `json:"Response"`
}

func (r *CreateWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CrontabResumeSuspendStrategy struct {
	// 定时拉起时间：如：周一&周三8点
	ResumeTime *string `json:"ResumeTime,omitnil,omitempty" name:"ResumeTime"`

	// 定时挂起时间：如：周一&周三20点
	SuspendTime *string `json:"SuspendTime,omitnil,omitempty" name:"SuspendTime"`

	// 挂起配置：0（默认）：等待任务结束后挂起、1：强制挂起
	SuspendStrategy *int64 `json:"SuspendStrategy,omitnil,omitempty" name:"SuspendStrategy"`
}

type CustomConfig struct {
	// 自定义参数名
	ConfigKey *string `json:"ConfigKey,omitnil,omitempty" name:"ConfigKey"`

	// 自定义参数值
	ConfigValue *string `json:"ConfigValue,omitnil,omitempty" name:"ConfigValue"`
}

type DLCCatalogAccess struct {
	// VPCID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 产品类型
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DMSColumn struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 附加参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*KVPair `json:"Params,omitnil,omitempty" name:"Params"`

	// 业务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizParams []*KVPair `json:"BizParams,omitnil,omitempty" name:"BizParams"`

	// 是否分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitnil,omitempty" name:"IsPartition"`
}

type DMSColumnOrder struct {
	// 列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Col *string `json:"Col,omitnil,omitempty" name:"Col"`

	// 排序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Order *int64 `json:"Order,omitnil,omitempty" name:"Order"`
}

type DMSPartition struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据目录名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 数据版本
	DataVersion *int64 `json:"DataVersion,omitnil,omitempty" name:"DataVersion"`

	// 分区名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 附件属性
	Params []*KVPair `json:"Params,omitnil,omitempty" name:"Params"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil,omitempty" name:"Sds"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DMSSds struct {
	// 存储地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 输入格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputFormat *string `json:"InputFormat,omitnil,omitempty" name:"InputFormat"`

	// 输出格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// bucket数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumBuckets *int64 `json:"NumBuckets,omitnil,omitempty" name:"NumBuckets"`

	// 是是否压缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Compressed *bool `json:"Compressed,omitnil,omitempty" name:"Compressed"`

	// 是否有子目录
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoredAsSubDirectories *bool `json:"StoredAsSubDirectories,omitnil,omitempty" name:"StoredAsSubDirectories"`

	// 序列化lib
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeLib *string `json:"SerdeLib,omitnil,omitempty" name:"SerdeLib"`

	// 序列化名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeName *string `json:"SerdeName,omitnil,omitempty" name:"SerdeName"`

	// 桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	BucketCols []*string `json:"BucketCols,omitnil,omitempty" name:"BucketCols"`

	// 序列化参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SerdeParams []*KVPair `json:"SerdeParams,omitnil,omitempty" name:"SerdeParams"`

	// 附加参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*KVPair `json:"Params,omitnil,omitempty" name:"Params"`

	// 列排序(Expired)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortCols *DMSColumnOrder `json:"SortCols,omitnil,omitempty" name:"SortCols"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cols []*DMSColumn `json:"Cols,omitnil,omitempty" name:"Cols"`

	// 列排序字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SortColumns []*DMSColumnOrder `json:"SortColumns,omitnil,omitempty" name:"SortColumns"`
}

type DMSTable struct {
	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewOriginalText *string `json:"ViewOriginalText,omitnil,omitempty" name:"ViewOriginalText"`

	// 视图文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewExpandedText *string `json:"ViewExpandedText,omitnil,omitempty" name:"ViewExpandedText"`

	// hive维护版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retention *int64 `json:"Retention,omitnil,omitempty" name:"Retention"`

	// 存储对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sds *DMSSds `json:"Sds,omitnil,omitempty" name:"Sds"`

	// 分区列
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil,omitempty" name:"PartitionKeys"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 存储大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 记录数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeTime *int64 `json:"LifeTime,omitnil,omitempty" name:"LifeTime"`

	// 最后访问时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 数据更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataUpdateTime *string `json:"DataUpdateTime,omitnil,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StructUpdateTime *string `json:"StructUpdateTime,omitnil,omitempty" name:"StructUpdateTime"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*DMSColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DMSTableInfo struct {
	// DMS表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *DMSTable `json:"Table,omitnil,omitempty" name:"Table"`

	// 基础对象信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`
}

type DataEngineBasicInfo struct {
	// DataEngine名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 数据引擎状态  -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 返回信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 引擎类型，有效值：PrestoSQL/SparkSQL/SparkBatch
	DataEngineType *string `json:"DataEngineType,omitnil,omitempty" name:"DataEngineType"`

	// 用户ID
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 账号ID
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`
}

type DataEngineConfigInstanceInfo struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 用户自定义配置项集合
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil,omitempty" name:"DataEngineConfigPairs"`

	// 作业集群资源参数配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`
}

type DataEngineConfigPair struct {
	// 配置项
	ConfigItem *string `json:"ConfigItem,omitnil,omitempty" name:"ConfigItem"`

	// 配置值
	ConfigValue *string `json:"ConfigValue,omitnil,omitempty" name:"ConfigValue"`
}

type DataEngineImageSessionParameter struct {
	// 配置id
	ParameterId *string `json:"ParameterId,omitnil,omitempty" name:"ParameterId"`

	// 小版本镜像ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`

	// 集群类型：SparkSQL/PrestoSQL/SparkBatch
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 参数key
	KeyName *string `json:"KeyName,omitnil,omitempty" name:"KeyName"`

	// Key描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyDescription *string `json:"KeyDescription,omitnil,omitempty" name:"KeyDescription"`

	// value类型
	ValueType *string `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// value长度限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueLengthLimit *string `json:"ValueLengthLimit,omitnil,omitempty" name:"ValueLengthLimit"`

	// value正则限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueRegexpLimit *string `json:"ValueRegexpLimit,omitnil,omitempty" name:"ValueRegexpLimit"`

	// value默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueDefault *string `json:"ValueDefault,omitnil,omitempty" name:"ValueDefault"`

	// 是否为公共版本：1：公共；2：私有
	IsPublic *uint64 `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 配置类型：1：session配置（默认）；2：common配置；3：cluster配置
	ParameterType *uint64 `json:"ParameterType,omitnil,omitempty" name:"ParameterType"`

	// 提交方式：User(用户)、BackGround（后台）
	SubmitMethod *string `json:"SubmitMethod,omitnil,omitempty" name:"SubmitMethod"`

	// 操作者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 插入时间
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DataEngineImageVersion struct {
	// 镜像大版本ID
	ImageVersionId *string `json:"ImageVersionId,omitnil,omitempty" name:"ImageVersionId"`

	// 镜像大版本名称
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// 镜像大版本描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否为公共版本：1：公共；2：私有
	IsPublic *uint64 `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 集群类型：SparkSQL/PrestoSQL/SparkBatch
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 版本状态：1：初始化；2：上线；3：下线
	IsSharedEngine *uint64 `json:"IsSharedEngine,omitnil,omitempty" name:"IsSharedEngine"`

	// 版本状态：1：初始化；2：上线；3：下线
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 插入时间
	InsertTime *string `json:"InsertTime,omitnil,omitempty" name:"InsertTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DataEngineInfo struct {
	// DataEngine名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 引擎类型 spark/presto
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 集群资源类型 spark_private/presto_private/presto_cu/spark_cu
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 引用ID
	QuotaId *string `json:"QuotaId,omitnil,omitempty" name:"QuotaId"`

	// 数据引擎状态  -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 集群规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 计费模式 0共享模式 1按量计费 2包年包月
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 最小集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinClusters *int64 `json:"MinClusters,omitnil,omitempty" name:"MinClusters"`

	// 最大集群数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxClusters *int64 `json:"MaxClusters,omitnil,omitempty" name:"MaxClusters"`

	// 是否自动恢复
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoResume *bool `json:"AutoResume,omitnil,omitempty" name:"AutoResume"`

	// 自动恢复时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpendAfter *int64 `json:"SpendAfter,omitnil,omitempty" name:"SpendAfter"`

	// 集群网段
	// 注意：此字段可能返回 null，表示取不到有效值。
	CidrBlock *string `json:"CidrBlock,omitnil,omitempty" name:"CidrBlock"`

	// 是否为默认引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDataEngine *bool `json:"DefaultDataEngine,omitnil,omitempty" name:"DefaultDataEngine"`

	// 返回信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 操作者
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 到期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 隔离时间
	IsolatedTime *string `json:"IsolatedTime,omitnil,omitempty" name:"IsolatedTime"`

	// 冲正时间
	ReversalTime *string `json:"ReversalTime,omitnil,omitempty" name:"ReversalTime"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 标签对集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagList []*TagInfo `json:"TagList,omitnil,omitempty" name:"TagList"`

	// 引擎拥有的权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permissions []*string `json:"Permissions,omitnil,omitempty" name:"Permissions"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSuspend *bool `json:"AutoSuspend,omitnil,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 引擎执行任务类型，有效值：SQL/BATCH
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineExecType *string `json:"EngineExecType,omitnil,omitempty" name:"EngineExecType"`

	// 自动续费标志，0，初始状态，默认不自动续费，若用户有预付费不停服特权，自动续费。1：自动续费。2：明确不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`

	// 集群自动挂起时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil,omitempty" name:"AutoSuspendTime"`

	// 网络连接配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionSet []*NetworkConnection `json:"NetworkConnectionSet,omitnil,omitempty" name:"NetworkConnectionSet"`

	// ui的跳转地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	UiURL *string `json:"UiURL,omitnil,omitempty" name:"UiURL"`

	// 引擎的资源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 集群镜像版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersionId *string `json:"ImageVersionId,omitnil,omitempty" name:"ImageVersionId"`

	// 集群镜像小版本ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`

	// 集群镜像版本名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersionName *string `json:"ImageVersionName,omitnil,omitempty" name:"ImageVersionName"`

	// 是否开启备集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitnil,omitempty" name:"StartStandbyCluster"`

	// spark jar 包年包月集群是否开启弹性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	ElasticLimit *int64 `json:"ElasticLimit,omitnil,omitempty" name:"ElasticLimit"`

	// 是否为默认引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultHouse *bool `json:"DefaultHouse,omitnil,omitempty" name:"DefaultHouse"`

	// 单个集群任务最大并发数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 任务排队上限时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil,omitempty" name:"TolerableQueueTime"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAppId *int64 `json:"UserAppId,omitnil,omitempty" name:"UserAppId"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserUin *string `json:"UserUin,omitnil,omitempty" name:"UserUin"`

	// SessionResourceTemplate
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`

	// 自动授权开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoAuthorization *bool `json:"AutoAuthorization,omitnil,omitempty" name:"AutoAuthorization"`

	// 引擎版本，支持Native/SuperSQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineGeneration *string `json:"EngineGeneration,omitnil,omitempty" name:"EngineGeneration"`

	// 引擎详细类型，支持：SparkSQL/SparkBatch/PrestoSQL/Kyuubi
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil,omitempty" name:"EngineTypeDetail"`

	// 引擎网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 标准引擎关联的资源组个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineResourceGroupCount *int64 `json:"EngineResourceGroupCount,omitnil,omitempty" name:"EngineResourceGroupCount"`

	// 引擎当前使用量（Cu）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineResourceUsedCU *int64 `json:"EngineResourceUsedCU,omitnil,omitempty" name:"EngineResourceUsedCU"`

	// 引擎的访问信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessInfos []*AccessInfo `json:"AccessInfos,omitnil,omitempty" name:"AccessInfos"`

	// 引擎所在网络名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkName *string `json:"EngineNetworkName,omitnil,omitempty" name:"EngineNetworkName"`

	// 是否使用预留池
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPoolMode *string `json:"IsPoolMode,omitnil,omitempty" name:"IsPoolMode"`

	// 是否支持AI，false: 不支持；true：支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSupportAI *bool `json:"IsSupportAI,omitnil,omitempty" name:"IsSupportAI"`

	// 网关id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayState *int64 `json:"GatewayState,omitnil,omitempty" name:"GatewayState"`

	// 是否能执行AI任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAIGateway *bool `json:"IsAIGateway,omitnil,omitempty" name:"IsAIGateway"`

	// 1:AI引擎，0:非AI引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAIEngine *int64 `json:"IsAIEngine,omitnil,omitempty" name:"IsAIEngine"`

	// 引擎资源弹性伸缩策略
	ScheduleElasticityConf *ScheduleElasticityConf `json:"ScheduleElasticityConf,omitnil,omitempty" name:"ScheduleElasticityConf"`
}

type DataEngineScaleInfo struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 引擎规格详情
	ScaleDetail []*DataEngineScaleInfoDetail `json:"ScaleDetail,omitnil,omitempty" name:"ScaleDetail"`
}

type DataEngineScaleInfoDetail struct {
	// 统计开始时间，格式为：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 统计结束时间，格式为：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 当前统计时间段，引擎规格
	CU *int64 `json:"CU,omitnil,omitempty" name:"CU"`
}

type DataFormat struct {
	// 文本格式，TextFile。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextFile *TextFile `json:"TextFile,omitnil,omitempty" name:"TextFile"`

	// 文本格式，CSV。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CSV *CSV `json:"CSV,omitnil,omitempty" name:"CSV"`

	// 文本格式，Json。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Json *Other `json:"Json,omitnil,omitempty" name:"Json"`

	// Parquet格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parquet *Other `json:"Parquet,omitnil,omitempty" name:"Parquet"`

	// ORC格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ORC *Other `json:"ORC,omitnil,omitempty" name:"ORC"`

	// AVRO格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	AVRO *Other `json:"AVRO,omitnil,omitempty" name:"AVRO"`
}

type DataGovernPolicy struct {
	// 治理规则类型，Customize: 自定义；Intelligence: 智能治理
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 治理引擎
	GovernEngine *string `json:"GovernEngine,omitnil,omitempty" name:"GovernEngine"`
}

type DataMaskStrategy struct {
	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 用户AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAppId *string `json:"UserAppId,omitnil,omitempty" name:"UserAppId"`

	// 用户Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 操作用户子账号uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// MASK_SHOW_FIRST_4; MASK_SHOW_LAST_4;MASK_HASH; MASK_DATE_SHOW_YEAR; MASK_NULL; MASK_DEFAULT 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyDesc *string `json:"StrategyDesc,omitnil,omitempty" name:"StrategyDesc"`

	// 用户组策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GroupInfo `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 用户子账号uin列表，按;拼接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Users *string `json:"Users,omitnil,omitempty" name:"Users"`

	// 1: 生效中； 0：已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 策略创建时间，毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 策略更新时间，毫秒时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type DataMaskStrategyInfo struct {
	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// MASK_SHOW_FIRST_4; MASK_SHOW_LAST_4;MASK_HASH; MASK_DATE_SHOW_YEAR; MASK_NULL; MASK_DEFAULT 等
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// 策略描述
	StrategyDesc *string `json:"StrategyDesc,omitnil,omitempty" name:"StrategyDesc"`

	// 用户组策略列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Groups []*GroupInfo `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 用户子账号uin列表，按;拼接
	Users *string `json:"Users,omitnil,omitempty" name:"Users"`

	// 策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DataMaskStrategyPolicy struct {
	// 数据脱敏权限对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyInfo *Policy `json:"PolicyInfo,omitnil,omitempty" name:"PolicyInfo"`

	// 数据脱敏策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataMaskStrategyId *string `json:"DataMaskStrategyId,omitnil,omitempty" name:"DataMaskStrategyId"`

	// 绑定字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnType *string `json:"ColumnType,omitnil,omitempty" name:"ColumnType"`
}

type DataSourceInfo struct {
	// 数据源实例的唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据源的名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据源的JDBC访问链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	JdbcUrl *string `json:"JdbcUrl,omitnil,omitempty" name:"JdbcUrl"`

	// 用于访问数据源的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据源访问密码，需要base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 数据源的VPC和子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *DatasourceConnectionLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 默认数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`
}

type DatabaseInfo struct {
	// 数据库名称，长度0~128，支持数字、字母下划线，不允许数字大头，统一转换为小写。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~500。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 数据库属性列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 数据库cos路径
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`
}

type DatabaseResponseInfo struct {
	// 数据库名称。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库描述信息，长度 0~256。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 允许针对数据库的属性元数据信息进行指定。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 数据库创建时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据库更新时间戳，单位：s。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// cos存储路径
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 建库用户昵称
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 建库用户ID
	UserSubUin *string `json:"UserSubUin,omitnil,omitempty" name:"UserSubUin"`

	// 数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernPolicy *DataGovernPolicy `json:"GovernPolicy,omitnil,omitempty" name:"GovernPolicy"`

	// 数据库ID（无效字段）
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`
}

type DatasourceConnectionConfig struct {
	// Mysql数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mysql *MysqlInfo `json:"Mysql,omitnil,omitempty" name:"Mysql"`

	// Hive数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hive *HiveInfo `json:"Hive,omitnil,omitempty" name:"Hive"`

	// Kafka数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kafka *KafkaInfo `json:"Kafka,omitnil,omitempty" name:"Kafka"`

	// 其他数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	OtherDatasourceConnection *OtherDatasourceConnection `json:"OtherDatasourceConnection,omitnil,omitempty" name:"OtherDatasourceConnection"`

	// PostgreSQL数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	PostgreSql *DataSourceInfo `json:"PostgreSql,omitnil,omitempty" name:"PostgreSql"`

	// SQLServer数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlServer *DataSourceInfo `json:"SqlServer,omitnil,omitempty" name:"SqlServer"`

	// ClickHouse数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClickHouse *DataSourceInfo `json:"ClickHouse,omitnil,omitempty" name:"ClickHouse"`

	// Elasticsearch数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elasticsearch *ElasticsearchInfo `json:"Elasticsearch,omitnil,omitempty" name:"Elasticsearch"`

	// TDSQL-PostgreSQL数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	TDSQLPostgreSql *DataSourceInfo `json:"TDSQLPostgreSql,omitnil,omitempty" name:"TDSQLPostgreSql"`

	// Doris数据源连接的属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	TCHouseD *TCHouseD `json:"TCHouseD,omitnil,omitempty" name:"TCHouseD"`

	// TccHive数据目录连接信息
	TccHive *TccHive `json:"TccHive,omitnil,omitempty" name:"TccHive"`
}

type DatasourceConnectionInfo struct {
	// 数据源数字Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据源字符串Id
	DatasourceConnectionId *string `json:"DatasourceConnectionId,omitnil,omitempty" name:"DatasourceConnectionId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 数据源描述
	DatasourceConnectionDesc *string `json:"DatasourceConnectionDesc,omitnil,omitempty" name:"DatasourceConnectionDesc"`

	// 数据源类型，支持DataLakeCatalog、IcebergCatalog、Result、Mysql、HiveCos、HiveHdfs
	DatasourceConnectionType *string `json:"DatasourceConnectionType,omitnil,omitempty" name:"DatasourceConnectionType"`

	// 数据源属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionConfig *DatasourceConnectionConfig `json:"DatasourceConnectionConfig,omitnil,omitempty" name:"DatasourceConnectionConfig"`

	// 数据源状态：0（初始化）、1（成功）、-1（已删除）、-2（失败）、-3（删除中）
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 用户AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 数据源创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据源最近一次更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 数据源同步失败原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 数据源绑定的计算引擎信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngines []*DataEngineInfo `json:"DataEngines,omitnil,omitempty" name:"DataEngines"`

	// 创建人
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 网络配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionSet []*NetworkConnection `json:"NetworkConnectionSet,omitnil,omitempty" name:"NetworkConnectionSet"`

	// 连通性状态：0（未测试，默认）、1（正常）、2（失败）
	ConnectivityState *uint64 `json:"ConnectivityState,omitnil,omitempty" name:"ConnectivityState"`

	// 连通性测试提示信息
	ConnectivityTips *string `json:"ConnectivityTips,omitnil,omitempty" name:"ConnectivityTips"`

	// 自定义参数
	CustomConfig []*CustomConfig `json:"CustomConfig,omitnil,omitempty" name:"CustomConfig"`

	// 是否允许回退
	AllowRollback *bool `json:"AllowRollback,omitnil,omitempty" name:"AllowRollback"`
}

type DatasourceConnectionLocation struct {
	// 数据连接所在Vpc实例Id，如“vpc-azd4dt1c”。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// Vpc的IPv4 CIDR
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 数据连接所在子网的实例Id，如“subnet-bthucmmy”
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Subnet的IPv4 CIDR
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitnil,omitempty" name:"SubnetCidrBlock"`
}

// Predefined struct for user
type DeleteCHDFSBindingProductRequestParams struct {
	// 需要解绑的元数据加速桶名
	MountPoint *string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 桶的类型，分为cos和lakefs
	BucketType *string `json:"BucketType,omitnil,omitempty" name:"BucketType"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 引擎名称，ProductName选择DLC产品时，必传此参数。其他产品可不传
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// vpc信息，ProductName选择other时，必传此参数
	VpcInfo []*VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`
}

type DeleteCHDFSBindingProductRequest struct {
	*tchttp.BaseRequest
	
	// 需要解绑的元数据加速桶名
	MountPoint *string `json:"MountPoint,omitnil,omitempty" name:"MountPoint"`

	// 桶的类型，分为cos和lakefs
	BucketType *string `json:"BucketType,omitnil,omitempty" name:"BucketType"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 引擎名称，ProductName选择DLC产品时，必传此参数。其他产品可不传
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// vpc信息，ProductName选择other时，必传此参数
	VpcInfo []*VpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`
}

func (r *DeleteCHDFSBindingProductRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCHDFSBindingProductRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MountPoint")
	delete(f, "BucketType")
	delete(f, "ProductName")
	delete(f, "EngineName")
	delete(f, "VpcInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCHDFSBindingProductRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCHDFSBindingProductResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCHDFSBindingProductResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCHDFSBindingProductResponseParams `json:"Response"`
}

func (r *DeleteCHDFSBindingProductResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCHDFSBindingProductResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataEngineRequestParams struct {
	// 删除虚拟集群的名称数组
	DataEngineNames []*string `json:"DataEngineNames,omitnil,omitempty" name:"DataEngineNames"`
}

type DeleteDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 删除虚拟集群的名称数组
	DataEngineNames []*string `json:"DataEngineNames,omitnil,omitempty" name:"DataEngineNames"`
}

func (r *DeleteDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataEngineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataEngineResponseParams `json:"Response"`
}

func (r *DeleteDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataMaskStrategyRequestParams struct {
	// 数据脱敏策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DeleteDataMaskStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 数据脱敏策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *DeleteDataMaskStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataMaskStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataMaskStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataMaskStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDataMaskStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataMaskStrategyResponseParams `json:"Response"`
}

func (r *DeleteDataMaskStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataMaskStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNativeSparkSessionRequestParams struct {
	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// spark session名称
	EngineSessionName *string `json:"EngineSessionName,omitnil,omitempty" name:"EngineSessionName"`
}

type DeleteNativeSparkSessionRequest struct {
	*tchttp.BaseRequest
	
	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// spark session名称
	EngineSessionName *string `json:"EngineSessionName,omitnil,omitempty" name:"EngineSessionName"`
}

func (r *DeleteNativeSparkSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNativeSparkSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "ResourceGroupId")
	delete(f, "EngineSessionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNativeSparkSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNativeSparkSessionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNativeSparkSessionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNativeSparkSessionResponseParams `json:"Response"`
}

func (r *DeleteNativeSparkSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNativeSparkSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookSessionRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DeleteNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DeleteNotebookSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotebookSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotebookSessionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNotebookSessionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotebookSessionResponseParams `json:"Response"`
}

func (r *DeleteNotebookSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotebookSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScriptRequestParams struct {
	// 脚本id，其可以通过DescribeScripts接口提取
	ScriptIds []*string `json:"ScriptIds,omitnil,omitempty" name:"ScriptIds"`
}

type DeleteScriptRequest struct {
	*tchttp.BaseRequest
	
	// 脚本id，其可以通过DescribeScripts接口提取
	ScriptIds []*string `json:"ScriptIds,omitnil,omitempty" name:"ScriptIds"`
}

func (r *DeleteScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScriptResponseParams struct {
	// 删除的脚本数量
	ScriptsAffected *int64 `json:"ScriptsAffected,omitnil,omitempty" name:"ScriptsAffected"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScriptResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScriptResponseParams `json:"Response"`
}

func (r *DeleteScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSparkAppRequestParams struct {
	// spark作业名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

type DeleteSparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

func (r *DeleteSparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSparkAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSparkAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSparkAppResponseParams `json:"Response"`
}

func (r *DeleteSparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStandardEngineResourceGroupRequestParams struct {
	// 标准引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`
}

type DeleteStandardEngineResourceGroupRequest struct {
	*tchttp.BaseRequest
	
	// 标准引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`
}

func (r *DeleteStandardEngineResourceGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStandardEngineResourceGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteStandardEngineResourceGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteStandardEngineResourceGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteStandardEngineResourceGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteStandardEngineResourceGroupResponseParams `json:"Response"`
}

func (r *DeleteStandardEngineResourceGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteStandardEngineResourceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`
}

type DeleteTableRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`
}

func (r *DeleteTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTableResponseParams `json:"Response"`
}

func (r *DeleteTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteThirdPartyAccessUserRequestParams struct {

}

type DeleteThirdPartyAccessUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DeleteThirdPartyAccessUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteThirdPartyAccessUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteThirdPartyAccessUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteThirdPartyAccessUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteThirdPartyAccessUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteThirdPartyAccessUserResponseParams `json:"Response"`
}

func (r *DeleteThirdPartyAccessUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteThirdPartyAccessUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserRequestParams struct {
	// 需要删除的用户的Id
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的用户的Id
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

func (r *DeleteUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserResponseParams `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserVpcConnectionRequestParams struct {
	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 终端节点ID
	UserVpcEndpointId *string `json:"UserVpcEndpointId,omitnil,omitempty" name:"UserVpcEndpointId"`
}

type DeleteUserVpcConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 终端节点ID
	UserVpcEndpointId *string `json:"UserVpcEndpointId,omitnil,omitempty" name:"UserVpcEndpointId"`
}

func (r *DeleteUserVpcConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserVpcConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineNetworkId")
	delete(f, "UserVpcEndpointId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUserVpcConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUserVpcConnectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUserVpcConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUserVpcConnectionResponseParams `json:"Response"`
}

func (r *DeleteUserVpcConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUserVpcConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersFromWorkGroupRequestParams struct {
	// 要删除的用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

type DeleteUsersFromWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的用户信息
	AddInfo *UserIdSetOfWorkGroupId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

func (r *DeleteUsersFromWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersFromWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteUsersFromWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteUsersFromWorkGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteUsersFromWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteUsersFromWorkGroupResponseParams `json:"Response"`
}

func (r *DeleteUsersFromWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteUsersFromWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkGroupRequestParams struct {
	// 要删除的工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil,omitempty" name:"WorkGroupIds"`
}

type DeleteWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil,omitempty" name:"WorkGroupIds"`
}

func (r *DeleteWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkGroupResponseParams `json:"Response"`
}

func (r *DeleteWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependencyPackage struct {
	// 依赖包类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageSource *string `json:"PackageSource,omitnil,omitempty" name:"PackageSource"`

	// 依赖包信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MavenPackage *string `json:"MavenPackage,omitnil,omitempty" name:"MavenPackage"`

	// 依赖包仓库
	// 注意：此字段可能返回 null，表示取不到有效值。
	MavenRepository *string `json:"MavenRepository,omitnil,omitempty" name:"MavenRepository"`

	// maven包exclusion信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MavenExclusion *string `json:"MavenExclusion,omitnil,omitempty" name:"MavenExclusion"`

	// pypi包信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PypiPackage *string `json:"PypiPackage,omitnil,omitempty" name:"PypiPackage"`

	// pypi索引地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PypiIndexUrl *string `json:"PypiIndexUrl,omitnil,omitempty" name:"PypiIndexUrl"`

	// 文件包的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackageType *string `json:"PackageType,omitnil,omitempty" name:"PackageType"`

	// 文件包的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PackagePath *string `json:"PackagePath,omitnil,omitempty" name:"PackagePath"`
}

// Predefined struct for user
type DescribeAdvancedStoreLocationRequestParams struct {

}

type DescribeAdvancedStoreLocationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAdvancedStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvancedStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAdvancedStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAdvancedStoreLocationResponseParams struct {
	// 是否启用高级设置：0-否，1-是
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 查询结果保存cos路径
	StoreLocation *string `json:"StoreLocation,omitnil,omitempty" name:"StoreLocation"`

	// 是否有托管存储权限
	HasLakeFs *bool `json:"HasLakeFs,omitnil,omitempty" name:"HasLakeFs"`

	// 托管存储状态，HasLakeFs等于true时，该值才有意义
	LakeFsStatus *string `json:"LakeFsStatus,omitnil,omitempty" name:"LakeFsStatus"`

	// 托管存储桶类型
	BucketType *string `json:"BucketType,omitnil,omitempty" name:"BucketType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAdvancedStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAdvancedStoreLocationResponseParams `json:"Response"`
}

func (r *DescribeAdvancedStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAdvancedStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterMonitorInfosRequestParams struct {
	// 引擎Id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 任务创建时间的起始时间
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// 任务创建时间的结束时间
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 指标名称
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`
}

type DescribeClusterMonitorInfosRequest struct {
	*tchttp.BaseRequest
	
	// 引擎Id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 任务创建时间的起始时间
	TimeStart *string `json:"TimeStart,omitnil,omitempty" name:"TimeStart"`

	// 任务创建时间的结束时间
	TimeEnd *string `json:"TimeEnd,omitnil,omitempty" name:"TimeEnd"`

	// 指标名称
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`
}

func (r *DescribeClusterMonitorInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterMonitorInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "TimeStart")
	delete(f, "TimeEnd")
	delete(f, "MetricName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterMonitorInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterMonitorInfosResponseParams struct {
	// 集群监控信息列表
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeClusterMonitorInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterMonitorInfosResponseParams `json:"Response"`
}

func (r *DescribeClusterMonitorInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterMonitorInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDLCCatalogAccessRequestParams struct {
	// 显示条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录数量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type DescribeDLCCatalogAccessRequest struct {
	*tchttp.BaseRequest
	
	// 显示条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 记录数量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件
	Filter *Filter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *DescribeDLCCatalogAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLCCatalogAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDLCCatalogAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDLCCatalogAccessResponseParams struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// DLCCatalog授权列表
	Rows []*DLCCatalogAccess `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDLCCatalogAccessResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDLCCatalogAccessResponseParams `json:"Response"`
}

func (r *DescribeDLCCatalogAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDLCCatalogAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSDatabaseRequestParams struct {
	// 数据库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 匹配规则，只支持填*
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DescribeDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 匹配规则，只支持填*
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DescribeDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "SchemaName")
	delete(f, "Pattern")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSDatabaseResponseParams struct {
	// 数据库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 存储地址
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSDatabaseResponseParams `json:"Response"`
}

func (r *DescribeDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSPartitionsRequestParams struct {
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 单个分区名称，精准匹配
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 多个分区名称，精准匹配
	PartitionNames []*string `json:"PartitionNames,omitnil,omitempty" name:"PartitionNames"`

	// 多个分区字段的匹配，模糊匹配
	PartValues []*string `json:"PartValues,omitnil,omitempty" name:"PartValues"`

	// 过滤SQL
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 最大分区数量
	MaxParts *int64 `json:"MaxParts,omitnil,omitempty" name:"MaxParts"`

	// 翻页跳过数量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 表达式
	Expression *string `json:"Expression,omitnil,omitempty" name:"Expression"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DescribeDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 单个分区名称，精准匹配
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 多个分区名称，精准匹配
	PartitionNames []*string `json:"PartitionNames,omitnil,omitempty" name:"PartitionNames"`

	// 多个分区字段的匹配，模糊匹配
	PartValues []*string `json:"PartValues,omitnil,omitempty" name:"PartValues"`

	// 过滤SQL
	Filter *string `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 最大分区数量
	MaxParts *int64 `json:"MaxParts,omitnil,omitempty" name:"MaxParts"`

	// 翻页跳过数量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页面数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 表达式
	Expression *string `json:"Expression,omitnil,omitempty" name:"Expression"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DescribeDMSPartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSPartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "TableName")
	delete(f, "SchemaName")
	delete(f, "Name")
	delete(f, "Values")
	delete(f, "PartitionNames")
	delete(f, "PartValues")
	delete(f, "Filter")
	delete(f, "MaxParts")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Expression")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSPartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSPartitionsResponseParams struct {
	// 分区信息
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDMSPartitionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSPartitionsResponseParams `json:"Response"`
}

func (r *DescribeDMSPartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSPartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTableRequestParams struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// catalog类型
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 查询模式，只支持填*
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DescribeDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// catalog类型
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 查询模式，只支持填*
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DescribeDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbName")
	delete(f, "SchemaName")
	delete(f, "Name")
	delete(f, "Catalog")
	delete(f, "Keyword")
	delete(f, "Pattern")
	delete(f, "Type")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTableResponseParams struct {
	// 基础对象
	Asset *Asset `json:"Asset,omitnil,omitempty" name:"Asset"`

	// 视图文本
	ViewOriginalText *string `json:"ViewOriginalText,omitnil,omitempty" name:"ViewOriginalText"`

	// 视图文本
	ViewExpandedText *string `json:"ViewExpandedText,omitnil,omitempty" name:"ViewExpandedText"`

	// hive维护版本
	Retention *int64 `json:"Retention,omitnil,omitempty" name:"Retention"`

	// 存储对象
	Sds *DMSSds `json:"Sds,omitnil,omitempty" name:"Sds"`

	// 分区列
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionKeys []*DMSColumn `json:"PartitionKeys,omitnil,omitempty" name:"PartitionKeys"`

	// 分区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*DMSPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// Schame名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 存储大小
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 记录数量
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 生命周期
	LifeTime *int64 `json:"LifeTime,omitnil,omitempty" name:"LifeTime"`

	// 最后访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 数据更新时间
	DataUpdateTime *string `json:"DataUpdateTime,omitnil,omitempty" name:"DataUpdateTime"`

	// 结构更新时间
	StructUpdateTime *string `json:"StructUpdateTime,omitnil,omitempty" name:"StructUpdateTime"`

	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*DMSColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSTableResponseParams `json:"Response"`
}

func (r *DescribeDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTablesRequestParams struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// catalog类型
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 查询模式，只支持填*
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 筛选参数：更新开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 筛选参数：更新结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段：create_time：创建时间
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段：true：升序（默认），false：降序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DescribeDMSTablesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// catalog类型
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 查询关键词
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 查询模式，只支持填*
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 表类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 筛选参数：更新开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 筛选参数：更新结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段：create_time：创建时间
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段：true：升序（默认），false：降序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DescribeDMSTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbName")
	delete(f, "SchemaName")
	delete(f, "Name")
	delete(f, "Catalog")
	delete(f, "Keyword")
	delete(f, "Pattern")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDMSTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDMSTablesResponseParams struct {
	// DMS元数据列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableList []*DMSTableInfo `json:"TableList,omitnil,omitempty" name:"TableList"`

	// 统计值
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDMSTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDMSTablesResponseParams `json:"Response"`
}

func (r *DescribeDMSTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDMSTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineEventsRequestParams struct {
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 返回数量，默认为10，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 资源组id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeDataEngineEventsRequest struct {
	*tchttp.BaseRequest
	
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 返回数量，默认为10，最大为100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 资源组id
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeDataEngineEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEngineEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineEventsResponseParams struct {
	// 事件详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*HouseEventsInfo `json:"Events,omitnil,omitempty" name:"Events"`

	// 分页号
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 分页大小
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 总页数
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEngineEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEngineEventsResponseParams `json:"Response"`
}

func (r *DescribeDataEngineEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineImageVersionsRequestParams struct {
	// 引擎类型：SparkSQL、PrestoSQL、SparkBatch、StandardSpark、StandardPresto
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 排序字段: InsertTime（插入时间，默认），UpdateTime（更新时间）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序方式：false（降序，默认），true（升序）
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`
}

type DescribeDataEngineImageVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型：SparkSQL、PrestoSQL、SparkBatch、StandardSpark、StandardPresto
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 排序字段: InsertTime（插入时间，默认），UpdateTime（更新时间）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序方式：false（降序，默认），true（升序）
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`
}

func (r *DescribeDataEngineImageVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineImageVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	delete(f, "Sort")
	delete(f, "Asc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEngineImageVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineImageVersionsResponseParams struct {
	// 集群大版本镜像信息列表
	ImageParentVersions []*DataEngineImageVersion `json:"ImageParentVersions,omitnil,omitempty" name:"ImageParentVersions"`

	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEngineImageVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEngineImageVersionsResponseParams `json:"Response"`
}

func (r *DescribeDataEngineImageVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineImageVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginePythonSparkImagesRequestParams struct {
	// 集群镜像小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`
}

type DescribeDataEnginePythonSparkImagesRequest struct {
	*tchttp.BaseRequest
	
	// 集群镜像小版本ID
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`
}

func (r *DescribeDataEnginePythonSparkImagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginePythonSparkImagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ChildImageVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEnginePythonSparkImagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginePythonSparkImagesResponseParams struct {
	// PYSPARK镜像信息列表
	PythonSparkImages []*PythonSparkImage `json:"PythonSparkImages,omitnil,omitempty" name:"PythonSparkImages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEnginePythonSparkImagesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEnginePythonSparkImagesResponseParams `json:"Response"`
}

func (r *DescribeDataEnginePythonSparkImagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginePythonSparkImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineRequestParams struct {
	// House名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`
}

type DescribeDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// House名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`
}

func (r *DescribeDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineResponseParams struct {
	// 数据引擎详细信息
	DataEngine *DataEngineInfo `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEngineResponseParams `json:"Response"`
}

func (r *DescribeDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineSessionParametersRequestParams struct {
	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 引擎名称，当指定引擎名称后优先使用名称获取配置
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`
}

type DescribeDataEngineSessionParametersRequest struct {
	*tchttp.BaseRequest
	
	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 引擎名称，当指定引擎名称后优先使用名称获取配置
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`
}

func (r *DescribeDataEngineSessionParametersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineSessionParametersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEngineSessionParametersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEngineSessionParametersResponseParams struct {
	// 集群Session配置列表
	DataEngineParameters []*DataEngineImageSessionParameter `json:"DataEngineParameters,omitnil,omitempty" name:"DataEngineParameters"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEngineSessionParametersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEngineSessionParametersResponseParams `json:"Response"`
}

func (r *DescribeDataEngineSessionParametersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEngineSessionParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginesRequestParams struct {
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤类型，支持如下的过滤类型，传参Name应为以下其中一个, data-engine-name - String（数据引擎名称）：engine-type - String（引擎类型：spark：spark 引擎，presto：presto引擎），state - String (数据引擎状态 -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中) ， mode - String（计费模式 0共享模式 1按量计费 2包年包月） ， create-time - String（创建时间，10位时间戳） message - String （描述信息），cluster-type - String (集群资源类型 spark_private/presto_private/presto_cu/spark_cu/kyuubi_cu)，engine-id - String（数据引擎ID），key-word - String（数据引擎名称或集群资源类型或描述信息模糊搜索），engine-exec-type - String（引擎执行任务类型，SQL/BATCH），engine-network-id - String（引擎网络Id）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 已废弃，请使用DatasourceConnectionNameSet
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 是否不返回共享引擎，true不返回共享引擎，false可以返回共享引擎
	ExcludePublicEngine *bool `json:"ExcludePublicEngine,omitnil,omitempty" name:"ExcludePublicEngine"`

	// 参数应该为引擎权限类型，有效类型："USE", "MODIFY", "OPERATE", "MONITOR", "DELETE"
	AccessTypes []*string `json:"AccessTypes,omitnil,omitempty" name:"AccessTypes"`

	// 引擎执行任务类型，有效值：SQL/BATCH，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitnil,omitempty" name:"EngineExecType"`

	// 引擎类型，有效值：spark/presto/kyuubi，为空时默认获取非kyuubi引擎（网关引擎）
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 网络配置列表，若传入该参数，则返回网络配置关联的计算引擎
	DatasourceConnectionNameSet []*string `json:"DatasourceConnectionNameSet,omitnil,omitempty" name:"DatasourceConnectionNameSet"`

	// 引擎版本，有效值：Native/SuperSQL，为空时默认获取SuperSQL引擎
	EngineGeneration *string `json:"EngineGeneration,omitnil,omitempty" name:"EngineGeneration"`

	// 引擎类型，支持：SparkSQL、SparkBatch、PrestoSQL、Kyuubi
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil,omitempty" name:"EngineTypeDetail"`

	// 默认 false, 为 true 时仅列出具有洞察 listener 的引擎
	ListHasListener *bool `json:"ListHasListener,omitnil,omitempty" name:"ListHasListener"`
}

type DescribeDataEnginesRequest struct {
	*tchttp.BaseRequest
	
	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤类型，支持如下的过滤类型，传参Name应为以下其中一个, data-engine-name - String（数据引擎名称）：engine-type - String（引擎类型：spark：spark 引擎，presto：presto引擎），state - String (数据引擎状态 -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中) ， mode - String（计费模式 0共享模式 1按量计费 2包年包月） ， create-time - String（创建时间，10位时间戳） message - String （描述信息），cluster-type - String (集群资源类型 spark_private/presto_private/presto_cu/spark_cu/kyuubi_cu)，engine-id - String（数据引擎ID），key-word - String（数据引擎名称或集群资源类型或描述信息模糊搜索），engine-exec-type - String（引擎执行任务类型，SQL/BATCH），engine-network-id - String（引擎网络Id）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 已废弃，请使用DatasourceConnectionNameSet
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 是否不返回共享引擎，true不返回共享引擎，false可以返回共享引擎
	ExcludePublicEngine *bool `json:"ExcludePublicEngine,omitnil,omitempty" name:"ExcludePublicEngine"`

	// 参数应该为引擎权限类型，有效类型："USE", "MODIFY", "OPERATE", "MONITOR", "DELETE"
	AccessTypes []*string `json:"AccessTypes,omitnil,omitempty" name:"AccessTypes"`

	// 引擎执行任务类型，有效值：SQL/BATCH，默认为SQL
	EngineExecType *string `json:"EngineExecType,omitnil,omitempty" name:"EngineExecType"`

	// 引擎类型，有效值：spark/presto/kyuubi，为空时默认获取非kyuubi引擎（网关引擎）
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 网络配置列表，若传入该参数，则返回网络配置关联的计算引擎
	DatasourceConnectionNameSet []*string `json:"DatasourceConnectionNameSet,omitnil,omitempty" name:"DatasourceConnectionNameSet"`

	// 引擎版本，有效值：Native/SuperSQL，为空时默认获取SuperSQL引擎
	EngineGeneration *string `json:"EngineGeneration,omitnil,omitempty" name:"EngineGeneration"`

	// 引擎类型，支持：SparkSQL、SparkBatch、PrestoSQL、Kyuubi
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil,omitempty" name:"EngineTypeDetail"`

	// 默认 false, 为 true 时仅列出具有洞察 listener 的引擎
	ListHasListener *bool `json:"ListHasListener,omitnil,omitempty" name:"ListHasListener"`
}

func (r *DescribeDataEnginesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "DatasourceConnectionName")
	delete(f, "ExcludePublicEngine")
	delete(f, "AccessTypes")
	delete(f, "EngineExecType")
	delete(f, "EngineType")
	delete(f, "DatasourceConnectionNameSet")
	delete(f, "EngineGeneration")
	delete(f, "EngineTypeDetail")
	delete(f, "ListHasListener")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEnginesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginesResponseParams struct {
	// 数据引擎列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngines []*DataEngineInfo `json:"DataEngines,omitnil,omitempty" name:"DataEngines"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEnginesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEnginesResponseParams `json:"Response"`
}

func (r *DescribeDataEnginesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginesScaleDetailRequestParams struct {
	// 引擎名称列表
	DataEngineNames []*string `json:"DataEngineNames,omitnil,omitempty" name:"DataEngineNames"`

	// 开始时间，时间格式：yyyy-MM-dd HH:mm:ss，最长查询一个月内的记录
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间格式：yyyy-MM-dd HH:mm:ss，最长查询一个月内的记录
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeDataEnginesScaleDetailRequest struct {
	*tchttp.BaseRequest
	
	// 引擎名称列表
	DataEngineNames []*string `json:"DataEngineNames,omitnil,omitempty" name:"DataEngineNames"`

	// 开始时间，时间格式：yyyy-MM-dd HH:mm:ss，最长查询一个月内的记录
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，时间格式：yyyy-MM-dd HH:mm:ss，最长查询一个月内的记录
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeDataEnginesScaleDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginesScaleDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineNames")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataEnginesScaleDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataEnginesScaleDetailResponseParams struct {
	// 引擎规格统计详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scales []*DataEngineScaleInfo `json:"Scales,omitnil,omitempty" name:"Scales"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataEnginesScaleDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataEnginesScaleDetailResponseParams `json:"Response"`
}

func (r *DescribeDataEnginesScaleDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataEnginesScaleDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataMaskStrategiesRequestParams struct {
	// 分页参数，单页返回数据量，默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，数据便偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤字段，strategy-name: 按策略名称搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDataMaskStrategiesRequest struct {
	*tchttp.BaseRequest
	
	// 分页参数，单页返回数据量，默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，数据便偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤字段，strategy-name: 按策略名称搜索
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDataMaskStrategiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataMaskStrategiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataMaskStrategiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataMaskStrategiesResponseParams struct {
	// 总数据脱敏策略数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据脱敏策略列表
	Strategies []*DataMaskStrategy `json:"Strategies,omitnil,omitempty" name:"Strategies"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataMaskStrategiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataMaskStrategiesResponseParams `json:"Response"`
}

func (r *DescribeDataMaskStrategiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataMaskStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 数据源唯名称，该名称可以通过DescribeDatasourceConnection接口查询到。默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 排序字段，CreateTime：创建时间，Name：数据库名称
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型：false：降序（默认）、true：升序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 查询类型：all：全部数据（默认）、permission：有权限的数据
	// 注意：此字段需要开启白名单使用，如果需要使用，请提交工单联系我们。
	DescribeType *string `json:"DescribeType,omitnil,omitempty" name:"DescribeType"`
}

type DescribeDatabasesRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 模糊匹配，库名关键字。
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 数据源唯名称，该名称可以通过DescribeDatasourceConnection接口查询到。默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 排序字段，CreateTime：创建时间，Name：数据库名称
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序类型：false：降序（默认）、true：升序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 查询类型：all：全部数据（默认）、permission：有权限的数据
	// 注意：此字段需要开启白名单使用，如果需要使用，请提交工单联系我们。
	DescribeType *string `json:"DescribeType,omitnil,omitempty" name:"DescribeType"`
}

func (r *DescribeDatabasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "KeyWord")
	delete(f, "DatasourceConnectionName")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "DescribeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabasesResponseParams struct {
	// 数据库对象列表。
	DatabaseList []*DatabaseResponseInfo `json:"DatabaseList,omitnil,omitempty" name:"DatabaseList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabasesResponseParams `json:"Response"`
}

func (r *DescribeDatabasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceConnectionRequestParams struct {
	// 连接ID列表，指定要查询的连接ID
	DatasourceConnectionIds []*string `json:"DatasourceConnectionIds,omitnil,omitempty" name:"DatasourceConnectionIds"`

	// 过滤条件，当前支持的过滤键为：DatasourceConnectionName（数据源连接名）。
	// DatasourceConnectionType   （数据源连接连接类型）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time（默认，创建时间）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为desc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 筛选字段：起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 筛选字段：截止时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 连接名称列表，指定要查询的连接名称
	DatasourceConnectionNames []*string `json:"DatasourceConnectionNames,omitnil,omitempty" name:"DatasourceConnectionNames"`

	// 连接类型，支持Mysql/HiveCos/Kafka/DataLakeCatalog
	DatasourceConnectionTypes []*string `json:"DatasourceConnectionTypes,omitnil,omitempty" name:"DatasourceConnectionTypes"`

	// 返回指定hive版本的数据源，该参数指定后，会过滤掉该参数指定版本以外的hive数据源，非hive数据源正常返回
	HiveVersion []*string `json:"HiveVersion,omitnil,omitempty" name:"HiveVersion"`
}

type DescribeDatasourceConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 连接ID列表，指定要查询的连接ID
	DatasourceConnectionIds []*string `json:"DatasourceConnectionIds,omitnil,omitempty" name:"DatasourceConnectionIds"`

	// 过滤条件，当前支持的过滤键为：DatasourceConnectionName（数据源连接名）。
	// DatasourceConnectionType   （数据源连接连接类型）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time（默认，创建时间）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为desc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 筛选字段：起始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 筛选字段：截止时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 连接名称列表，指定要查询的连接名称
	DatasourceConnectionNames []*string `json:"DatasourceConnectionNames,omitnil,omitempty" name:"DatasourceConnectionNames"`

	// 连接类型，支持Mysql/HiveCos/Kafka/DataLakeCatalog
	DatasourceConnectionTypes []*string `json:"DatasourceConnectionTypes,omitnil,omitempty" name:"DatasourceConnectionTypes"`

	// 返回指定hive版本的数据源，该参数指定后，会过滤掉该参数指定版本以外的hive数据源，非hive数据源正常返回
	HiveVersion []*string `json:"HiveVersion,omitnil,omitempty" name:"HiveVersion"`
}

func (r *DescribeDatasourceConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceConnectionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DatasourceConnectionNames")
	delete(f, "DatasourceConnectionTypes")
	delete(f, "HiveVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasourceConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceConnectionResponseParams struct {
	// 数据连接总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据连接对象集合
	ConnectionSet []*DatasourceConnectionInfo `json:"ConnectionSet,omitnil,omitempty" name:"ConnectionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatasourceConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasourceConnectionResponseParams `json:"Response"`
}

func (r *DescribeDatasourceConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineNetworksRequestParams struct {
	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序，降序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件可选，engine-network-id--引擎网络ID，engine-network-state--引擎网络状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeEngineNetworksRequest struct {
	*tchttp.BaseRequest
	
	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序，降序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件可选，engine-network-id--引擎网络ID，engine-network-state--引擎网络状态
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeEngineNetworksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineNetworksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEngineNetworksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineNetworksResponseParams struct {
	// 引擎网络信息
	EngineNetworkInfos []*EngineNetworkInfo `json:"EngineNetworkInfos,omitnil,omitempty" name:"EngineNetworkInfos"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEngineNetworksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEngineNetworksResponseParams `json:"Response"`
}

func (r *DescribeEngineNetworksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineNetworksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineNodeSpecRequestParams struct {
	// 引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`
}

type DescribeEngineNodeSpecRequest struct {
	*tchttp.BaseRequest
	
	// 引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`
}

func (r *DescribeEngineNodeSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineNodeSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEngineNodeSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineNodeSpecResponseParams struct {
	// driver可用的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	DriverSpec []*SpecInfo `json:"DriverSpec,omitnil,omitempty" name:"DriverSpec"`

	// executor可用的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorSpec []*SpecInfo `json:"ExecutorSpec,omitnil,omitempty" name:"ExecutorSpec"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEngineNodeSpecResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEngineNodeSpecResponseParams `json:"Response"`
}

func (r *DescribeEngineNodeSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineNodeSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineUsageInfoRequestParams struct {
	// 数据引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

type DescribeEngineUsageInfoRequest struct {
	*tchttp.BaseRequest
	
	// 数据引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

func (r *DescribeEngineUsageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineUsageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEngineUsageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEngineUsageInfoResponseParams struct {
	// 集群总规格
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 已占用集群规格
	Used *int64 `json:"Used,omitnil,omitempty" name:"Used"`

	// 剩余集群规格
	Available *int64 `json:"Available,omitnil,omitempty" name:"Available"`

	// 剩余集群规格百分比
	AvailPercent *int64 `json:"AvailPercent,omitnil,omitempty" name:"AvailPercent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEngineUsageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEngineUsageInfoResponseParams `json:"Response"`
}

func (r *DescribeEngineUsageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEngineUsageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForbiddenTableProRequestParams struct {

}

type DescribeForbiddenTableProRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeForbiddenTableProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForbiddenTableProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeForbiddenTableProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeForbiddenTableProResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeForbiddenTableProResponse struct {
	*tchttp.BaseResponse
	Response *DescribeForbiddenTableProResponseParams `json:"Response"`
}

func (r *DescribeForbiddenTableProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeForbiddenTableProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsDirSummaryRequestParams struct {

}

type DescribeLakeFsDirSummaryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLakeFsDirSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsDirSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsDirSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsDirSummaryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLakeFsDirSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsDirSummaryResponseParams `json:"Response"`
}

func (r *DescribeLakeFsDirSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsDirSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsInfoRequestParams struct {

}

type DescribeLakeFsInfoRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeLakeFsInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsInfoResponseParams struct {
	// 托管存储信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LakeFsInfos []*LakeFsInfo `json:"LakeFsInfos,omitnil,omitempty" name:"LakeFsInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLakeFsInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsInfoResponseParams `json:"Response"`
}

func (r *DescribeLakeFsInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsTaskResultRequestParams struct {
	// 需要访问的任务结果路径
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`
}

type DescribeLakeFsTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 需要访问的任务结果路径
	FsPath *string `json:"FsPath,omitnil,omitempty" name:"FsPath"`
}

func (r *DescribeLakeFsTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FsPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLakeFsTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLakeFsTaskResultResponseParams struct {
	// 路径的访问实例
	AccessToken *LakeFileSystemToken `json:"AccessToken,omitnil,omitempty" name:"AccessToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLakeFsTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLakeFsTaskResultResponseParams `json:"Response"`
}

func (r *DescribeLakeFsTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLakeFsTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNativeSparkSessionsRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`
}

type DescribeNativeSparkSessionsRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`
}

func (r *DescribeNativeSparkSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNativeSparkSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "ResourceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNativeSparkSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNativeSparkSessionsResponseParams struct {
	// spark session列表
	SparkSessionsList []*SparkSessionInfo `json:"SparkSessionsList,omitnil,omitempty" name:"SparkSessionsList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNativeSparkSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNativeSparkSessionsResponseParams `json:"Response"`
}

func (r *DescribeNativeSparkSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNativeSparkSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkConnectionsRequestParams struct {
	// 网络配置类型
	NetworkConnectionType *int64 `json:"NetworkConnectionType,omitnil,omitempty" name:"NetworkConnectionType"`

	// 计算引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 数据源vpcid
	DatasourceConnectionVpcId *string `json:"DatasourceConnectionVpcId,omitnil,omitempty" name:"DatasourceConnectionVpcId"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 网络配置名称
	NetworkConnectionName *string `json:"NetworkConnectionName,omitnil,omitempty" name:"NetworkConnectionName"`
}

type DescribeNetworkConnectionsRequest struct {
	*tchttp.BaseRequest
	
	// 网络配置类型
	NetworkConnectionType *int64 `json:"NetworkConnectionType,omitnil,omitempty" name:"NetworkConnectionType"`

	// 计算引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 数据源vpcid
	DatasourceConnectionVpcId *string `json:"DatasourceConnectionVpcId,omitnil,omitempty" name:"DatasourceConnectionVpcId"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 网络配置名称
	NetworkConnectionName *string `json:"NetworkConnectionName,omitnil,omitempty" name:"NetworkConnectionName"`
}

func (r *DescribeNetworkConnectionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkConnectionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkConnectionType")
	delete(f, "DataEngineName")
	delete(f, "DatasourceConnectionVpcId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "NetworkConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNetworkConnectionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNetworkConnectionsResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 网络配置列表
	NetworkConnectionSet []*NetworkConnection `json:"NetworkConnectionSet,omitnil,omitempty" name:"NetworkConnectionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNetworkConnectionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNetworkConnectionsResponseParams `json:"Response"`
}

func (r *DescribeNetworkConnectionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNetworkConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionLogRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeNotebookSessionLogRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeNotebookSessionLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionLogResponseParams struct {
	// 日志信息，默认获取最新的200条
	Logs []*string `json:"Logs,omitnil,omitempty" name:"Logs"`

	// 分页参数，默认200
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookSessionLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionLogResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeNotebookSessionRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeNotebookSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionResponseParams struct {
	// Session详情信息
	Session *NotebookSessionInfo `json:"Session,omitnil,omitempty" name:"Session"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil,omitempty" name:"StatementId"`

	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeNotebookSessionStatementRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil,omitempty" name:"StatementId"`

	// 任务唯一标识
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeNotebookSessionStatementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "StatementId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionStatementRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementResponseParams struct {
	// Session Statement详情
	NotebookSessionStatement *NotebookSessionStatementInfo `json:"NotebookSessionStatement,omitnil,omitempty" name:"NotebookSessionStatement"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookSessionStatementResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionStatementResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionStatementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementSqlResultRequestParams struct {
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 批次Id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 返回结果集中字段值长度截取，如果超过该长度则截取到该长度
	DataFieldCutLen *int64 `json:"DataFieldCutLen,omitnil,omitempty" name:"DataFieldCutLen"`
}

type DescribeNotebookSessionStatementSqlResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *uint64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 批次Id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 返回结果集中字段值长度截取，如果超过该长度则截取到该长度
	DataFieldCutLen *int64 `json:"DataFieldCutLen,omitnil,omitempty" name:"DataFieldCutLen"`
}

func (r *DescribeNotebookSessionStatementSqlResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementSqlResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "MaxResults")
	delete(f, "NextToken")
	delete(f, "BatchId")
	delete(f, "DataFieldCutLen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionStatementSqlResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementSqlResultResponseParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 结果数据
	ResultSet *string `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`

	// schema
	ResultSchema []*Column `json:"ResultSchema,omitnil,omitempty" name:"ResultSchema"`

	// 分页信息
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 存储结果地址
	OutputPath *string `json:"OutputPath,omitnil,omitempty" name:"OutputPath"`

	// 引擎计算耗时
	UseTime *int64 `json:"UseTime,omitnil,omitempty" name:"UseTime"`

	// 结果条数
	AffectRows *int64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// 数据扫描量
	DataAmount *int64 `json:"DataAmount,omitnil,omitempty" name:"DataAmount"`

	// spark ui地址
	UiUrl *string `json:"UiUrl,omitnil,omitempty" name:"UiUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookSessionStatementSqlResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionStatementSqlResultResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionStatementSqlResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementSqlResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementsRequestParams struct {
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 批任务id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

type DescribeNotebookSessionStatementsRequest struct {
	*tchttp.BaseRequest
	
	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 批任务id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

func (r *DescribeNotebookSessionStatementsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionStatementsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionStatementsResponseParams struct {
	// Session Statement详情
	NotebookSessionStatements *NotebookSessionStatementBatchInformation `json:"NotebookSessionStatements,omitnil,omitempty" name:"NotebookSessionStatements"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookSessionStatementsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionStatementsResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionStatementsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionStatementsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionsRequestParams struct {
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 排序字段（默认按创建时间）
	SortFields []*string `json:"SortFields,omitnil,omitempty" name:"SortFields"`

	// 排序字段：true：升序、false：降序（默认）
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 分页参数，默认10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤类型，支持如下的过滤类型，传参Name应为以下其中一个, engine-generation - String（引擎时代： supersql：supersql引擎，native：标准引擎）：notebook-keyword - String（数据引擎名称或sessionid或sessionname的模糊搜索）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeNotebookSessionsRequest struct {
	*tchttp.BaseRequest
	
	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State []*string `json:"State,omitnil,omitempty" name:"State"`

	// 排序字段（默认按创建时间）
	SortFields []*string `json:"SortFields,omitnil,omitempty" name:"SortFields"`

	// 排序字段：true：升序、false：降序（默认）
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 分页参数，默认10
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页参数，默认0
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤类型，支持如下的过滤类型，传参Name应为以下其中一个, engine-generation - String（引擎时代： supersql：supersql引擎，native：标准引擎）：notebook-keyword - String（数据引擎名称或sessionid或sessionname的模糊搜索）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeNotebookSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "State")
	delete(f, "SortFields")
	delete(f, "Asc")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotebookSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotebookSessionsResponseParams struct {
	// session总数量
	TotalElements *int64 `json:"TotalElements,omitnil,omitempty" name:"TotalElements"`

	// 总页数
	TotalPages *int64 `json:"TotalPages,omitnil,omitempty" name:"TotalPages"`

	// 当前页码
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 当前页数量
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// session列表信息
	Sessions []*NotebookSessions `json:"Sessions,omitnil,omitempty" name:"Sessions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotebookSessionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotebookSessionsResponseParams `json:"Response"`
}

func (r *DescribeNotebookSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotebookSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtherCHDFSBindingListRequestParams struct {
	// 桶名
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`
}

type DescribeOtherCHDFSBindingListRequest struct {
	*tchttp.BaseRequest
	
	// 桶名
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`
}

func (r *DescribeOtherCHDFSBindingListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtherCHDFSBindingListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BucketId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOtherCHDFSBindingListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOtherCHDFSBindingListResponseParams struct {
	// 非DLC 产品绑定列表
	OtherCHDFSBindingList []*OtherCHDFSBinding `json:"OtherCHDFSBindingList,omitnil,omitempty" name:"OtherCHDFSBindingList"`

	// 总记录数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOtherCHDFSBindingListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOtherCHDFSBindingListResponseParams `json:"Response"`
}

func (r *DescribeOtherCHDFSBindingListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOtherCHDFSBindingListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResultDownloadRequestParams struct {
	// 查询任务Id
	DownloadId *string `json:"DownloadId,omitnil,omitempty" name:"DownloadId"`
}

type DescribeResultDownloadRequest struct {
	*tchttp.BaseRequest
	
	// 查询任务Id
	DownloadId *string `json:"DownloadId,omitnil,omitempty" name:"DownloadId"`
}

func (r *DescribeResultDownloadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResultDownloadRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DownloadId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResultDownloadRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResultDownloadResponseParams struct {
	// 下载文件路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 任务状态 init | queue | format | compress | success|  timeout | error
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务异常原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 临时SecretId
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 临时SecretKey
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 临时Token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeResultDownloadResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResultDownloadResponseParams `json:"Response"`
}

func (r *DescribeResultDownloadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResultDownloadResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScriptsRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序，默认asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeScriptsRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按字段排序，支持如下字段类型，update-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序，默认asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// script-id - String - （过滤条件）script-id取值形如：157de0d1-26b4-4df2-a2d0-b64afc406c25。
	// script-name-keyword - String - （过滤条件）数据表名称,形如：script-test。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeScriptsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScriptsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScriptsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScriptsResponseParams struct {
	// Script列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scripts []*Script `json:"Scripts,omitnil,omitempty" name:"Scripts"`

	// 实例总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScriptsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScriptsResponseParams `json:"Response"`
}

func (r *DescribeScriptsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScriptsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionImageVersionRequestParams struct {
	// 引擎Id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 框架类型：machine-learning、python、spark-ml
	FrameworkType *string `json:"FrameworkType,omitnil,omitempty" name:"FrameworkType"`
}

type DescribeSessionImageVersionRequest struct {
	*tchttp.BaseRequest
	
	// 引擎Id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 框架类型：machine-learning、python、spark-ml
	FrameworkType *string `json:"FrameworkType,omitnil,omitempty" name:"FrameworkType"`
}

func (r *DescribeSessionImageVersionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionImageVersionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "FrameworkType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSessionImageVersionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSessionImageVersionResponseParams struct {
	// 扩展镜像列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineSessionImages []*EngineSessionImage `json:"EngineSessionImages,omitnil,omitempty" name:"EngineSessionImages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSessionImageVersionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSessionImageVersionResponseParams `json:"Response"`
}

func (r *DescribeSessionImageVersionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSessionImageVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobRequestParams struct {
	// spark作业Id，与JobName同时存在时，JobName无效，JobId与JobName至少存在一个
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`
}

type DescribeSparkAppJobRequest struct {
	*tchttp.BaseRequest
	
	// spark作业Id，与JobName同时存在时，JobName无效，JobId与JobName至少存在一个
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`
}

func (r *DescribeSparkAppJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobResponseParams struct {
	// spark作业详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Job *SparkJobInfo `json:"Job,omitnil,omitempty" name:"Job"`

	// 查询的spark作业是否存在
	IsExists *bool `json:"IsExists,omitnil,omitempty" name:"IsExists"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSparkAppJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppJobResponseParams `json:"Response"`
}

func (r *DescribeSparkAppJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobsRequestParams struct {
	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一:spark-job-name（作业名称），spark-job-id（作业id），spark-app-type（作业类型，1：批任务，2：流任务，4：SQL作业），user-name（创建人），key-word（作业名称或ID关键词模糊搜索）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询列表偏移量, 默认值0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询列表限制数量, 默认值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSparkAppJobsRequest struct {
	*tchttp.BaseRequest
	
	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一:spark-job-name（作业名称），spark-job-id（作业id），spark-app-type（作业类型，1：批任务，2：流任务，4：SQL作业），user-name（创建人），key-word（作业名称或ID关键词模糊搜索）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询列表偏移量, 默认值0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询列表限制数量, 默认值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSparkAppJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppJobsResponseParams struct {
	// spark作业列表详情
	SparkAppJobs []*SparkJobInfo `json:"SparkAppJobs,omitnil,omitempty" name:"SparkAppJobs"`

	// spark作业总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSparkAppJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppJobsResponseParams `json:"Response"`
}

func (r *DescribeSparkAppJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppTasksRequestParams struct {
	// spark作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行实例id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 按照该参数过滤,支持task-state
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeSparkAppTasksRequest struct {
	*tchttp.BaseRequest
	
	// spark作业Id
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 分页查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页查询Limit
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 执行实例id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 更新时间起始点，支持格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 更新时间截止点，支持格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 按照该参数过滤,支持task-state
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeSparkAppTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TaskId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkAppTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkAppTasksResponseParams struct {
	// 任务结果（该字段已废弃）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks *TaskResponseInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkAppTasks []*TaskResponseInfo `json:"SparkAppTasks,omitnil,omitempty" name:"SparkAppTasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSparkAppTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkAppTasksResponseParams `json:"Response"`
}

func (r *DescribeSparkAppTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkAppTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSQLCostRequestParams struct {
	// SparkSQL唯一标识
	BatchIds []*string `json:"BatchIds,omitnil,omitempty" name:"BatchIds"`
}

type DescribeSparkSessionBatchSQLCostRequest struct {
	*tchttp.BaseRequest
	
	// SparkSQL唯一标识
	BatchIds []*string `json:"BatchIds,omitnil,omitempty" name:"BatchIds"`
}

func (r *DescribeSparkSessionBatchSQLCostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSQLCostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkSessionBatchSQLCostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSQLCostResponseParams struct {
	// 任务消耗
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostInfo []*BatchSQLCostInfo `json:"CostInfo,omitnil,omitempty" name:"CostInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSparkSessionBatchSQLCostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkSessionBatchSQLCostResponseParams `json:"Response"`
}

func (r *DescribeSparkSessionBatchSQLCostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSQLCostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSQLRequestParams struct {
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 用户自定义主键, 若不为空，则按照该值查询
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`
}

type DescribeSparkSessionBatchSQLRequest struct {
	*tchttp.BaseRequest
	
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 用户自定义主键, 若不为空，则按照该值查询
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`
}

func (r *DescribeSparkSessionBatchSQLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSQLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "CustomKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkSessionBatchSQLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSQLResponseParams struct {
	// 状态：0：运行中、1：成功、2：失败、3：取消、4：超时；
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// SQL子任务列表，仅展示运行完成的子任务，若某个任务运行失败，后续其它子任务不返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*BatchSqlTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 非sql运行的异常事件信息，包含资源创建失败、调度异常，JOB超时等，正常运行下该Event值为空
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSparkSessionBatchSQLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkSessionBatchSQLResponseParams `json:"Response"`
}

func (r *DescribeSparkSessionBatchSQLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSQLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSqlLogRequestParams struct {
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 用户自定义主键，若不为空，则按照该值进行查询
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`
}

type DescribeSparkSessionBatchSqlLogRequest struct {
	*tchttp.BaseRequest
	
	// SparkSQL唯一标识
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 用户自定义主键，若不为空，则按照该值进行查询
	CustomKey *string `json:"CustomKey,omitnil,omitempty" name:"CustomKey"`
}

func (r *DescribeSparkSessionBatchSqlLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSqlLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BatchId")
	delete(f, "CustomKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSparkSessionBatchSqlLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSparkSessionBatchSqlLogResponseParams struct {
	// 状态：0：运行中、1：成功、2：失败、3：取消、4：超时；
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 日志信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSet []*SparkSessionBatchLog `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSparkSessionBatchSqlLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSparkSessionBatchSqlLogResponseParams `json:"Response"`
}

func (r *DescribeSparkSessionBatchSqlLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSparkSessionBatchSqlLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandardEngineResourceGroupConfigInfoRequestParams struct {
	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序，降序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件可选，engine-resource-group-id--引擎资源组ID，engine-id---引擎ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据条数，默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeStandardEngineResourceGroupConfigInfoRequest struct {
	*tchttp.BaseRequest
	
	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序，降序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件可选，engine-resource-group-id--引擎资源组ID，engine-id---引擎ID
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据条数，默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeStandardEngineResourceGroupConfigInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandardEngineResourceGroupConfigInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStandardEngineResourceGroupConfigInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandardEngineResourceGroupConfigInfoResponseParams struct {
	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 标准引擎资源组，配置相关信息
	StandardEngineResourceGroupConfigInfos []*StandardEngineResourceGroupConfigInfo `json:"StandardEngineResourceGroupConfigInfos,omitnil,omitempty" name:"StandardEngineResourceGroupConfigInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStandardEngineResourceGroupConfigInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStandardEngineResourceGroupConfigInfoResponseParams `json:"Response"`
}

func (r *DescribeStandardEngineResourceGroupConfigInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandardEngineResourceGroupConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandardEngineResourceGroupsRequestParams struct {
	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序，降序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件可选，app-id--用户appID，engine-resource-group-id--引擎资源组ID，data-engine-name--引擎名称，engine-resource-group-name---引擎资源组名称（模糊查询），engine-resource-group-state---引擎资源组状态engine-resource-group-name-unique --引擎资源组名称（完全匹配）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据条数，默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeStandardEngineResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 排序字段
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 升序，降序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件可选，app-id--用户appID，engine-resource-group-id--引擎资源组ID，data-engine-name--引擎名称，engine-resource-group-name---引擎资源组名称（模糊查询），engine-resource-group-state---引擎资源组状态engine-resource-group-name-unique --引擎资源组名称（完全匹配）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据条数，默认10
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeStandardEngineResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandardEngineResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStandardEngineResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandardEngineResourceGroupsResponseParams struct {
	// 标准引擎资源组信息
	UserEngineResourceGroupInfos []*StandardEngineResourceGroupInfo `json:"UserEngineResourceGroupInfos,omitnil,omitempty" name:"UserEngineResourceGroupInfos"`

	// 资源组总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStandardEngineResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStandardEngineResourceGroupsResponseParams `json:"Response"`
}

func (r *DescribeStandardEngineResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandardEngineResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStoreLocationRequestParams struct {

}

type DescribeStoreLocationRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStoreLocationResponseParams struct {
	// 返回用户设置的结果存储位置路径，如果未设置则返回空字符串：""
	StoreLocation *string `json:"StoreLocation,omitnil,omitempty" name:"StoreLocation"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStoreLocationResponseParams `json:"Response"`
}

func (r *DescribeStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubUserAccessPolicyRequestParams struct {

}

type DescribeSubUserAccessPolicyRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSubUserAccessPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubUserAccessPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubUserAccessPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubUserAccessPolicyResponseParams struct {
	// 子用户访问策略
	PolicyDocument *string `json:"PolicyDocument,omitnil,omitempty" name:"PolicyDocument"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSubUserAccessPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubUserAccessPolicyResponseParams `json:"Response"`
}

func (r *DescribeSubUserAccessPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubUserAccessPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablePartitionsRequestParams struct {
	// 数据目录名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据表名称
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 查询偏移位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 当次查询的数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊查询的分区名称
	FuzzyPartition *string `json:"FuzzyPartition,omitnil,omitempty" name:"FuzzyPartition"`

	// 排序信息
	Sorts []*Sort `json:"Sorts,omitnil,omitempty" name:"Sorts"`

	// 分页查询的游标信息
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

type DescribeTablePartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 数据目录名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 数据表名称
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 查询偏移位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 当次查询的数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 模糊查询的分区名称
	FuzzyPartition *string `json:"FuzzyPartition,omitnil,omitempty" name:"FuzzyPartition"`

	// 排序信息
	Sorts []*Sort `json:"Sorts,omitnil,omitempty" name:"Sorts"`

	// 分页查询的游标信息
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`
}

func (r *DescribeTablePartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablePartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Catalog")
	delete(f, "Database")
	delete(f, "Table")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FuzzyPartition")
	delete(f, "Sorts")
	delete(f, "Cursor")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablePartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablePartitionsResponseParams struct {
	// 分区信息值
	MixedPartitions *MixedTablePartitions `json:"MixedPartitions,omitnil,omitempty" name:"MixedPartitions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTablePartitionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablePartitionsResponseParams `json:"Response"`
}

func (r *DescribeTablePartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableRequestParams struct {
	// 查询对象表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 查询表所在的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DescribeTableRequest struct {
	*tchttp.BaseRequest
	
	// 查询对象表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 查询表所在的数据库名称。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 查询表所在的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DescribeTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableName")
	delete(f, "DatabaseName")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableResponseParams struct {
	// 数据表对象
	Table *TableResponseInfo `json:"Table,omitnil,omitempty" name:"Table"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableResponseParams `json:"Response"`
}

func (r *DescribeTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesNameRequestParams struct {
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`
}

type DescribeTablesNameRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`
}

func (r *DescribeTablesNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "TableType")
	delete(f, "TableFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesNameResponseParams struct {
	// 数据表名称对象列表。
	TableNameList []*string `json:"TableNameList,omitnil,omitempty" name:"TableNameList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTablesNameResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesNameResponseParams `json:"Response"`
}

func (r *DescribeTablesNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesRequestParams struct {
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`

	// 查询类型：all：全部数据（默认）、permission：有权限的数据
	// 注意：此字段需要开启白名单使用，如果需要使用，请提交工单联系我们。
	DescribeType *string `json:"DescribeType,omitnil,omitempty" name:"DescribeType"`
}

type DescribeTablesRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// table-name - String - （过滤条件）数据表名称,形如：table-001。
	// table-id - String - （过滤条件）table id形如：12342。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 指定查询的数据源名称，默认为DataLakeCatalog
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 起始时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 终止时间：用于对更新时间的筛选，格式为yyyy-mm-dd HH:MM:SS
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序字段，支持：CreateTime（创建时间）、UpdateTime（更新时间）、StorageSize（存储空间）、RecordCount（行数）、Name（表名称）（不传则默认按name升序）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段，false：降序（默认）；true：升序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// table type，表类型查询,可用值:EXTERNAL_TABLE,INDEX_TABLE,MANAGED_TABLE,MATERIALIZED_VIEW,TABLE,VIEW,VIRTUAL_VIEW
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 筛选字段-表格式：不传（默认）为查全部；LAKEFS：托管表；ICEBERG：非托管iceberg表；HIVE：非托管hive表；OTHER：非托管其它；
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`

	// 查询类型：all：全部数据（默认）、permission：有权限的数据
	// 注意：此字段需要开启白名单使用，如果需要使用，请提交工单联系我们。
	DescribeType *string `json:"DescribeType,omitnil,omitempty" name:"DescribeType"`
}

func (r *DescribeTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "TableType")
	delete(f, "TableFormat")
	delete(f, "DescribeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablesResponseParams struct {
	// 数据表对象列表。
	TableList []*TableResponseInfo `json:"TableList,omitnil,omitempty" name:"TableList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTablesResponseParams `json:"Response"`
}

func (r *DescribeTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogRequestParams struct {
	// 列表返回的Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下一次分页参数，第一次传空。透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// SparkSQL任务唯一ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

type DescribeTaskLogRequest struct {
	*tchttp.BaseRequest
	
	// 列表返回的Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下一次分页参数，第一次传空。透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// SparkSQL任务唯一ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

func (r *DescribeTaskLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Asc")
	delete(f, "Filters")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLogResponseParams struct {
	// 下一次分页参数
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 是否获取完结
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 日志详情
	Results []*JobLogResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 日志url
	LogUrl *string `json:"LogUrl,omitnil,omitempty" name:"LogUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLogResponseParams `json:"Response"`
}

func (r *DescribeTaskLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskMonitorInfosRequestParams struct {
	// 任务ID列表，上限50个
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 引擎名称
	HouseName *string `json:"HouseName,omitnil,omitempty" name:"HouseName"`

	// 任务创建时间的起始时间
	CreateTimeStart *string `json:"CreateTimeStart,omitnil,omitempty" name:"CreateTimeStart"`

	// 任务创建时间的结束时间
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`

	// 每一页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeTaskMonitorInfosRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID列表，上限50个
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 引擎名称
	HouseName *string `json:"HouseName,omitnil,omitempty" name:"HouseName"`

	// 任务创建时间的起始时间
	CreateTimeStart *string `json:"CreateTimeStart,omitnil,omitempty" name:"CreateTimeStart"`

	// 任务创建时间的结束时间
	CreateTimeEnd *string `json:"CreateTimeEnd,omitnil,omitempty" name:"CreateTimeEnd"`

	// 每一页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeTaskMonitorInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskMonitorInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "HouseName")
	delete(f, "CreateTimeStart")
	delete(f, "CreateTimeEnd")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskMonitorInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskMonitorInfosResponseParams struct {
	// 任务监控信息列表
	TaskMonitorInfoList []*TaskMonitorInfo `json:"TaskMonitorInfoList,omitnil,omitempty" name:"TaskMonitorInfoList"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskMonitorInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskMonitorInfosResponseParams `json:"Response"`
}

func (r *DescribeTaskMonitorInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskMonitorInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResourceUsageRequestParams struct {
	// 任务 id
	TaskInstanceId *string `json:"TaskInstanceId,omitnil,omitempty" name:"TaskInstanceId"`
}

type DescribeTaskResourceUsageRequest struct {
	*tchttp.BaseRequest
	
	// 任务 id
	TaskInstanceId *string `json:"TaskInstanceId,omitnil,omitempty" name:"TaskInstanceId"`
}

func (r *DescribeTaskResourceUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResourceUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskInstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResourceUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResourceUsageResponseParams struct {
	// core 用量信息
	CoreInfo *CoreInfo `json:"CoreInfo,omitnil,omitempty" name:"CoreInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskResourceUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResourceUsageResponseParams `json:"Response"`
}

func (r *DescribeTaskResourceUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResourceUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultRequestParams struct {
	// 任务唯一ID，仅支持30天内的任务
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否转化数据类型
	IsTransformDataType *bool `json:"IsTransformDataType,omitnil,omitempty" name:"IsTransformDataType"`

	// 返回结果集中字段长度截取，如果字段值长度超过该长度则截取到该长度
	DataFieldCutLen *int64 `json:"DataFieldCutLen,omitnil,omitempty" name:"DataFieldCutLen"`
}

type DescribeTaskResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务唯一ID，仅支持30天内的任务
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 上一次请求响应返回的分页信息。第一次可以不带，从头开始返回数据，每次返回MaxResults字段设置的数据量。
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 返回结果的最大行数，范围0~1000，默认为1000.
	MaxResults *int64 `json:"MaxResults,omitnil,omitempty" name:"MaxResults"`

	// 是否转化数据类型
	IsTransformDataType *bool `json:"IsTransformDataType,omitnil,omitempty" name:"IsTransformDataType"`

	// 返回结果集中字段长度截取，如果字段值长度超过该长度则截取到该长度
	DataFieldCutLen *int64 `json:"DataFieldCutLen,omitnil,omitempty" name:"DataFieldCutLen"`
}

func (r *DescribeTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "NextToken")
	delete(f, "MaxResults")
	delete(f, "IsTransformDataType")
	delete(f, "DataFieldCutLen")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskResultResponseParams struct {
	// 查询的任务信息，返回为空表示输入任务ID对应的任务不存在。只有当任务状态为成功（2）的时候，才会返回任务的结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfo *TaskResultInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskResultResponseParams `json:"Response"`
}

func (r *DescribeTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksAnalysisRequestParams struct {
	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个: task-id - String - （任务ID准确过滤）task-id 取值形如：e386471f-139a-4e59-877f-50ece8135b99。task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)，rule-id - String - （洞察类型）取值范围 SPARK-StageScheduleDelay（资源抢占）, SPARK-ShuffleFailure（Shuffle异常）, SPARK-SlowTask（慢task）, SPARK-DataSkew（数据倾斜）, SPARK-InsufficientResource（磁盘或内存不足）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，instance-start-time (任务开始时间）,job-time-sum （单位毫秒，引擎内执行耗时）,task-time-sum （CU资源消耗，单位秒）,input-bytes-sum（数据扫描总大小，单位bytes）,shuffle-read-bytes-sum（数据shuffle总大小，单位bytes）
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 任务开始时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近30天数据查询。默认为当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近30天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeTasksAnalysisRequest struct {
	*tchttp.BaseRequest
	
	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个: task-id - String - （任务ID准确过滤）task-id 取值形如：e386471f-139a-4e59-877f-50ece8135b99。task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)，rule-id - String - （洞察类型）取值范围 SPARK-StageScheduleDelay（资源抢占）, SPARK-ShuffleFailure（Shuffle异常）, SPARK-SlowTask（慢task）, SPARK-DataSkew（数据倾斜）, SPARK-InsufficientResource（磁盘或内存不足）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，instance-start-time (任务开始时间）,job-time-sum （单位毫秒，引擎内执行耗时）,task-time-sum （CU资源消耗，单位秒）,input-bytes-sum（数据扫描总大小，单位bytes）,shuffle-read-bytes-sum（数据shuffle总大小，单位bytes）
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 任务开始时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近30天数据查询。默认为当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近30天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeTasksAnalysisRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksAnalysisRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksAnalysisRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksAnalysisResponseParams struct {
	// 洞察结果分页列表
	TaskList []*AnalysisTaskResults `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 洞察结果总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksAnalysisResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksAnalysisResponseParams `json:"Response"`
}

func (r *DescribeTasksAnalysisResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksCostInfoRequestParams struct {
	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 下一页的标识
	SearchAfter *string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// 每页的大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeTasksCostInfoRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 下一页的标识
	SearchAfter *string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// 每页的大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeTasksCostInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksCostInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataEngineName")
	delete(f, "SearchAfter")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksCostInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksCostInfoResponseParams struct {
	// 下一页的标识
	SearchAfter *string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// 返回的数据，字符串类型的二维数组，首行为列中文名称
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksCostInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksCostInfoResponseParams `json:"Response"`
}

func (r *DescribeTasksCostInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksCostInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksOverviewRequestParams struct {
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 引擎名
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// DataEngine-dm8bjs29
	HouseIds []*string `json:"HouseIds,omitnil,omitempty" name:"HouseIds"`
}

type DescribeTasksOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 筛选条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 引擎名
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// DataEngine-dm8bjs29
	HouseIds []*string `json:"HouseIds,omitnil,omitempty" name:"HouseIds"`
}

func (r *DescribeTasksOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "DataEngineName")
	delete(f, "HouseIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksOverviewResponseParams struct {
	// 各类任务个数大于0
	TasksOverview *TasksOverview `json:"TasksOverview,omitnil,omitempty" name:"TasksOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksOverviewResponseParams `json:"Response"`
}

func (r *DescribeTasksOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksRequestParams struct {
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	// task-kind - string （任务类型过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time（创建时间，默认）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// spark引擎资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	// task-kind - string （任务类型过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，支持如下字段类型，create-time（创建时间，默认）、update-time（更新时间）
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc。
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// spark引擎资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataEngineName")
	delete(f, "ResourceGroupName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksResponseParams struct {
	// 任务对象列表。
	TaskList []*TaskResponseInfo `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务概览信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TasksOverview *TasksOverview `json:"TasksOverview,omitnil,omitempty" name:"TasksOverview"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksResponseParams `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAccessUserRequestParams struct {

}

type DescribeThirdPartyAccessUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeThirdPartyAccessUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAccessUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeThirdPartyAccessUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdPartyAccessUserResponseParams struct {
	// 用户信息
	UserInfo *OpendThirdAccessUserInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeThirdPartyAccessUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeThirdPartyAccessUserResponseParams `json:"Response"`
}

func (r *DescribeThirdPartyAccessUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdPartyAccessUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUDFPolicyRequestParams struct {
	// udf名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库名(全局UDF：global-function)
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

type DescribeUDFPolicyRequest struct {
	*tchttp.BaseRequest
	
	// udf名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库名(全局UDF：global-function)
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`
}

func (r *DescribeUDFPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUDFPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DatabaseName")
	delete(f, "CatalogName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUDFPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUDFPolicyResponseParams struct {
	// UDF权限信息
	UDFPolicyInfos []*UDFPolicyInfo `json:"UDFPolicyInfos,omitnil,omitempty" name:"UDFPolicyInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUDFPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUDFPolicyResponseParams `json:"Response"`
}

func (r *DescribeUDFPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUDFPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpdatableDataEnginesRequestParams struct {
	// 引擎配置操作命令，UpdateSparkSQLLakefsPath 更新托管表路径，UpdateSparkSQLResultPath 更新结果桶路径
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil,omitempty" name:"DataEngineConfigCommand"`

	// 是否使用托管存储作为结果存储
	UseLakeFs *bool `json:"UseLakeFs,omitnil,omitempty" name:"UseLakeFs"`

	// 用户自定义结果存储路径
	CustomResultPath *string `json:"CustomResultPath,omitnil,omitempty" name:"CustomResultPath"`
}

type DescribeUpdatableDataEnginesRequest struct {
	*tchttp.BaseRequest
	
	// 引擎配置操作命令，UpdateSparkSQLLakefsPath 更新托管表路径，UpdateSparkSQLResultPath 更新结果桶路径
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil,omitempty" name:"DataEngineConfigCommand"`

	// 是否使用托管存储作为结果存储
	UseLakeFs *bool `json:"UseLakeFs,omitnil,omitempty" name:"UseLakeFs"`

	// 用户自定义结果存储路径
	CustomResultPath *string `json:"CustomResultPath,omitnil,omitempty" name:"CustomResultPath"`
}

func (r *DescribeUpdatableDataEnginesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpdatableDataEnginesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineConfigCommand")
	delete(f, "UseLakeFs")
	delete(f, "CustomResultPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpdatableDataEnginesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpdatableDataEnginesResponseParams struct {
	// 集群基础信息
	DataEngineBasicInfos []*DataEngineBasicInfo `json:"DataEngineBasicInfos,omitnil,omitempty" name:"DataEngineBasicInfos"`

	// 集群个数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUpdatableDataEnginesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpdatableDataEnginesResponseParams `json:"Response"`
}

func (r *DescribeUpdatableDataEnginesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpdatableDataEnginesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDataEngineConfigRequestParams struct {
	// 排序方式，desc表示倒序，asc表示正序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,每种过滤参数支持的过滤值不超过5个。
	// app-id - String - （appid过滤）
	// engine-id - String - （引擎ID过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeUserDataEngineConfigRequest struct {
	*tchttp.BaseRequest
	
	// 排序方式，desc表示倒序，asc表示正序
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,每种过滤参数支持的过滤值不超过5个。
	// app-id - String - （appid过滤）
	// engine-id - String - （引擎ID过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeUserDataEngineConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDataEngineConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserDataEngineConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserDataEngineConfigResponseParams struct {
	// 用户引擎自定义配置项列表。
	DataEngineConfigInstanceInfos []*DataEngineConfigInstanceInfo `json:"DataEngineConfigInstanceInfos,omitnil,omitempty" name:"DataEngineConfigInstanceInfos"`

	// 配置项总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserDataEngineConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserDataEngineConfigResponseParams `json:"Response"`
}

func (r *DescribeUserDataEngineConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserDataEngineConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoRequestParams struct {
	// 用户Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 必传字段，查询的信息类型，Group：工作组 DataAuth：数据权限 EngineAuth:引擎权限 RowFilter：行级别权限
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为Group时，支持Key为workgroup-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为Group时，支持create-time、group-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 必传字段，查询的信息类型，Group：工作组 DataAuth：数据权限 EngineAuth:引擎权限 RowFilter：行级别权限
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为Group时，支持Key为workgroup-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为Group时，支持create-time、group-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "Type")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInfoResponseParams struct {
	// 用户详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserInfo *UserDetailInfo `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInfoResponseParams `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRegisterTimeRequestParams struct {

}

type DescribeUserRegisterTimeRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeUserRegisterTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRegisterTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRegisterTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRegisterTimeResponseParams struct {
	// 用户注册时间
	RegisterTime *int64 `json:"RegisterTime,omitnil,omitempty" name:"RegisterTime"`

	// 是否时老用户
	IsOldUser *bool `json:"IsOldUser,omitnil,omitempty" name:"IsOldUser"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserRegisterTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserRegisterTimeResponseParams `json:"Response"`
}

func (r *DescribeUserRegisterTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRegisterTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRolesRequestParams struct {
	// 列举的数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列举的偏移位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照arn模糊列举
	Fuzzy *string `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`

	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 是否设置为常驻：1非常驻（默认）、2常驻（仅能设置一个常驻）
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
}

type DescribeUserRolesRequest struct {
	*tchttp.BaseRequest
	
	// 列举的数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列举的偏移位置
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照arn模糊列举
	Fuzzy *string `json:"Fuzzy,omitnil,omitempty" name:"Fuzzy"`

	// 返回结果按照该字段排序
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 正序或者倒序，例如：desc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 是否设置为常驻：1非常驻（默认）、2常驻（仅能设置一个常驻）
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
}

func (r *DescribeUserRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Fuzzy")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "IsDefault")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserRolesResponseParams struct {
	// 符合列举条件的总数量
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 用户角色信息
	UserRoles []*UserRole `json:"UserRoles,omitnil,omitempty" name:"UserRoles"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserRolesResponseParams `json:"Response"`
}

func (r *DescribeUserRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserTypeRequestParams struct {
	// 用户ID（UIN），如果不填默认为调用方的子UIN
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeUserTypeRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID（UIN），如果不填默认为调用方的子UIN
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeUserTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserTypeResponseParams struct {
	// 用户类型。ADMIN：管理员 COMMON：普通用户
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserTypeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserTypeResponseParams `json:"Response"`
}

func (r *DescribeUserTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserVpcConnectionRequestParams struct {
	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 引擎ID集合
	DataEngineIds []*string `json:"DataEngineIds,omitnil,omitempty" name:"DataEngineIds"`

	// 终端节点ID集合
	UserVpcEndpointIds []*string `json:"UserVpcEndpointIds,omitnil,omitempty" name:"UserVpcEndpointIds"`
}

type DescribeUserVpcConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 引擎网络ID
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 引擎ID集合
	DataEngineIds []*string `json:"DataEngineIds,omitnil,omitempty" name:"DataEngineIds"`

	// 终端节点ID集合
	UserVpcEndpointIds []*string `json:"UserVpcEndpointIds,omitnil,omitempty" name:"UserVpcEndpointIds"`
}

func (r *DescribeUserVpcConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserVpcConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineNetworkId")
	delete(f, "DataEngineIds")
	delete(f, "UserVpcEndpointIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserVpcConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserVpcConnectionResponseParams struct {
	// 用户vpc连接信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVpcConnectionInfos []*UserVpcConnectionInfo `json:"UserVpcConnectionInfos,omitnil,omitempty" name:"UserVpcConnectionInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserVpcConnectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserVpcConnectionResponseParams `json:"Response"`
}

func (r *DescribeUserVpcConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserVpcConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersRequestParams struct {
	// 指定查询的子用户uin，用户需要通过CreateUser接口创建。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件，支持如下字段类型，user-type：根据用户类型过滤。user-keyword：根据用户名称过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeUsersRequest struct {
	*tchttp.BaseRequest
	
	// 指定查询的子用户uin，用户需要通过CreateUser接口创建。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 过滤条件，支持如下字段类型，user-type：根据用户类型过滤。user-keyword：根据用户名称过滤
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsersResponseParams struct {
	// 查询到的用户总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 查询到的授权用户信息集合
	UserSet []*UserInfo `json:"UserSet,omitnil,omitempty" name:"UserSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsersResponseParams `json:"Response"`
}

func (r *DescribeUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeViewsRequestParams struct {
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据库所属的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 排序字段
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序规则，true:升序；false:降序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 按视图更新时间筛选，开始时间，如2021-11-11 00:00:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 按视图更新时间筛选，结束时间，如2021-11-12 00:00:00
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询类型：all：全部数据（默认）、permission：有权限的数据
	// 注意：此字段需要开启白名单使用，如果需要使用，请提交工单联系我们。
	DescribeType *string `json:"DescribeType,omitnil,omitempty" name:"DescribeType"`
}

type DescribeViewsRequest struct {
	*tchttp.BaseRequest
	
	// 列出该数据库下所属数据表。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 返回数量，默认为10，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据偏移量，从0开始，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，如下支持的过滤类型，传参Name应为其一
	// view-name - String - （过滤条件）数据表名称,形如：view-001。
	// view-id - String - （过滤条件）view id形如：12342。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 数据库所属的数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 排序字段
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序规则，true:升序；false:降序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 按视图更新时间筛选，开始时间，如2021-11-11 00:00:00
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 按视图更新时间筛选，结束时间，如2021-11-12 00:00:00
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 查询类型：all：全部数据（默认）、permission：有权限的数据
	// 注意：此字段需要开启白名单使用，如果需要使用，请提交工单联系我们。
	DescribeType *string `json:"DescribeType,omitnil,omitempty" name:"DescribeType"`
}

func (r *DescribeViewsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "DatasourceConnectionName")
	delete(f, "Sort")
	delete(f, "Asc")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DescribeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeViewsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeViewsResponseParams struct {
	// 视图对象列表。
	ViewList []*ViewResponseInfo `json:"ViewList,omitnil,omitempty" name:"ViewList"`

	// 实例总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeViewsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeViewsResponseParams `json:"Response"`
}

func (r *DescribeViewsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupInfoRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 查询信息类型：User：用户信息 DataAuth：数据权限 EngineAuth：引擎权限
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为User时，支持Key为user-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为User时，支持create-time、user-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeWorkGroupInfoRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 查询信息类型：User：用户信息 DataAuth：数据权限 EngineAuth：引擎权限
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 查询的过滤条件。
	// 
	// 当Type为User时，支持Key为user-name的模糊搜索；
	// 
	// 当Type为DataAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// data-name：库表的模糊搜索。
	// 
	// 当Type为EngineAuth时，支持key：
	// 
	// policy-type：权限类型。
	// 
	// policy-source：数据来源。
	// 
	// engine-name：库表的模糊搜索。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段。
	// 
	// 当Type为User时，支持create-time、user-name
	// 
	// 当Type为DataAuth时，支持create-time
	// 
	// 当Type为EngineAuth时，支持create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeWorkGroupInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "Type")
	delete(f, "Filters")
	delete(f, "SortBy")
	delete(f, "Sorting")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkGroupInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupInfoResponseParams struct {
	// 工作组详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupInfo *WorkGroupDetailInfo `json:"WorkGroupInfo,omitnil,omitempty" name:"WorkGroupInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkGroupInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkGroupInfoResponseParams `json:"Response"`
}

func (r *DescribeWorkGroupInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupsRequestParams struct {
	// 查询的工作组Id，不填或填0表示不过滤。
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 过滤条件，当前仅支持按照工作组名称进行模糊搜索。Key为workgroup-name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`
}

type DescribeWorkGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 查询的工作组Id，不填或填0表示不过滤。
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 过滤条件，当前仅支持按照工作组名称进行模糊搜索。Key为workgroup-name
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值100
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序字段，支持如下字段类型，create-time
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，desc表示正序，asc表示反序， 默认为asc
	Sorting *string `json:"Sorting,omitnil,omitempty" name:"Sorting"`
}

func (r *DescribeWorkGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Sorting")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkGroupsResponseParams struct {
	// 工作组总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 工作组信息集合
	WorkGroupSet []*WorkGroupInfo `json:"WorkGroupSet,omitnil,omitempty" name:"WorkGroupSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkGroupsResponseParams `json:"Response"`
}

func (r *DescribeWorkGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyRequestParams struct {
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

type DetachUserPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

func (r *DetachUserPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachUserPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachUserPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachUserPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachUserPolicyResponseParams `json:"Response"`
}

func (r *DetachUserPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachUserPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachWorkGroupPolicyRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

type DetachWorkGroupPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 解绑的权限集合
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`
}

func (r *DetachWorkGroupPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachWorkGroupPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "PolicySet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachWorkGroupPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachWorkGroupPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachWorkGroupPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DetachWorkGroupPolicyResponseParams `json:"Response"`
}

func (r *DetachWorkGroupPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachWorkGroupPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSDatabaseRequestParams struct {
	// 数据库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil,omitempty" name:"DeleteData"`

	// 是否级联删除
	Cascade *bool `json:"Cascade,omitnil,omitempty" name:"Cascade"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DropDMSDatabaseRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil,omitempty" name:"DeleteData"`

	// 是否级联删除
	Cascade *bool `json:"Cascade,omitnil,omitempty" name:"Cascade"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DropDMSDatabaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSDatabaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DeleteData")
	delete(f, "Cascade")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropDMSDatabaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSDatabaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DropDMSDatabaseResponse struct {
	*tchttp.BaseResponse
	Response *DropDMSDatabaseResponseParams `json:"Response"`
}

func (r *DropDMSDatabaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSDatabaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSPartitionsRequestParams struct {
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 数据表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 分区名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 单个分区名称
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否删除分区数据
	DeleteData *bool `json:"DeleteData,omitnil,omitempty" name:"DeleteData"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DropDMSPartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库Schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 数据表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 分区名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 单个分区名称
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 是否删除分区数据
	DeleteData *bool `json:"DeleteData,omitnil,omitempty" name:"DeleteData"`

	// 数据源连接名
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DropDMSPartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSPartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "SchemaName")
	delete(f, "TableName")
	delete(f, "Name")
	delete(f, "Values")
	delete(f, "DeleteData")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropDMSPartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSPartitionsResponseParams struct {
	// 状态
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DropDMSPartitionsResponse struct {
	*tchttp.BaseResponse
	Response *DropDMSPartitionsResponseParams `json:"Response"`
}

func (r *DropDMSPartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSPartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSTableRequestParams struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil,omitempty" name:"DeleteData"`

	// 环境属性
	EnvProps *KVPair `json:"EnvProps,omitnil,omitempty" name:"EnvProps"`

	// 数据目录信息
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type DropDMSTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否删除数据
	DeleteData *bool `json:"DeleteData,omitnil,omitempty" name:"DeleteData"`

	// 环境属性
	EnvProps *KVPair `json:"EnvProps,omitnil,omitempty" name:"EnvProps"`

	// 数据目录信息
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *DropDMSTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DbName")
	delete(f, "Name")
	delete(f, "DeleteData")
	delete(f, "EnvProps")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DropDMSTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DropDMSTableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DropDMSTableResponse struct {
	*tchttp.BaseResponse
	Response *DropDMSTableResponseParams `json:"Response"`
}

func (r *DropDMSTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DropDMSTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ElasticPlan struct {
	// 最小集群数
	MinElasticClusters *int64 `json:"MinElasticClusters,omitnil,omitempty" name:"MinElasticClusters"`

	// 最大集群数
	MaxElasticClusters *int64 `json:"MaxElasticClusters,omitnil,omitempty" name:"MaxElasticClusters"`

	// 最大排队时间
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil,omitempty" name:"TolerableQueueTime"`

	// 开始时间，Once格式：yyyy-MM-dd HH:mm:ss; 非Once格式： HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，Once格式：yyyy-MM-dd HH:mm:ss; 非Once格式： HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ElasticsearchInfo struct {
	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 密码，需要base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 数据源的VPC和子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *DatasourceConnectionLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 默认数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 访问Elasticsearch的ip、端口信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceInfo []*IpPortPair `json:"ServiceInfo,omitnil,omitempty" name:"ServiceInfo"`
}

type EngineNetworkInfo struct {
	// 引擎网络名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkName *string `json:"EngineNetworkName,omitnil,omitempty" name:"EngineNetworkName"`

	// 引擎网络状态，0--初始化，2--可用，-1--已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkState *int64 `json:"EngineNetworkState,omitnil,omitempty" name:"EngineNetworkState"`

	// 引擎网络cidr
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkCidr *string `json:"EngineNetworkCidr,omitnil,omitempty" name:"EngineNetworkCidr"`

	// 引擎网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 私有连接个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateLinkNumber *int64 `json:"PrivateLinkNumber,omitnil,omitempty" name:"PrivateLinkNumber"`

	// 计算引擎个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNumber *int64 `json:"EngineNumber,omitnil,omitempty" name:"EngineNumber"`

	// 网关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GateWayInfo []*GatewayInfo `json:"GateWayInfo,omitnil,omitempty" name:"GateWayInfo"`
}

type EngineResourceGroupConfigPair struct {
	// 配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigItem *string `json:"ConfigItem,omitnil,omitempty" name:"ConfigItem"`

	// 配置项的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigValue *string `json:"ConfigValue,omitnil,omitempty" name:"ConfigValue"`
}

type EngineSessionImage struct {
	// Spark镜像唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImageId *string `json:"SparkImageId,omitnil,omitempty" name:"SparkImageId"`

	// Spark镜像版本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImageVersion *string `json:"SparkImageVersion,omitnil,omitempty" name:"SparkImageVersion"`

	// 小版本镜像类型.1:TensorFlow、2:Pytorch、3:SK-learn
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImageType *int64 `json:"SparkImageType,omitnil,omitempty" name:"SparkImageType"`

	// 镜像地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImageTag *string `json:"SparkImageTag,omitnil,omitempty" name:"SparkImageTag"`
}

type Execution struct {
	// 自动生成SQL语句。
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`
}

type FavorInfo struct {
	// 优先事项
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// Catalog名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// DataBase名称
	DataBase *string `json:"DataBase,omitnil,omitempty" name:"DataBase"`

	// Table名称
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`
}

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑或（OR）关系。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type GatewayInfo struct {
	// 网关ID，
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关名称，全局唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayName *string `json:"GatewayName,omitnil,omitempty" name:"GatewayName"`

	// 网关的规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// -2已删除 -1失败 0初始化中 1挂起 2运行中 3准备删除 4删除中 5挂起中 6启动中 7隔离中 8隔离 9续费中 10变配中 11冲正中
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 计费模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *int64 `json:"Mode,omitnil,omitempty" name:"Mode"`
}

// Predefined struct for user
type GenerateCreateMangedTableSqlRequestParams struct {
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// V2 upsert表 upsert键
	UpsertKeys []*string `json:"UpsertKeys,omitnil,omitempty" name:"UpsertKeys"`
}

type GenerateCreateMangedTableSqlRequest struct {
	*tchttp.BaseRequest
	
	// 表基本信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 表字段信息
	Columns []*TColumn `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表分区信息
	Partitions []*TPartition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 表属性信息
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// V2 upsert表 upsert键
	UpsertKeys []*string `json:"UpsertKeys,omitnil,omitempty" name:"UpsertKeys"`
}

func (r *GenerateCreateMangedTableSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCreateMangedTableSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableBaseInfo")
	delete(f, "Columns")
	delete(f, "Partitions")
	delete(f, "Properties")
	delete(f, "UpsertKeys")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateCreateMangedTableSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateCreateMangedTableSqlResponseParams struct {
	// 创建托管存储内表sql语句描述
	Execution *Execution `json:"Execution,omitnil,omitempty" name:"Execution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateCreateMangedTableSqlResponse struct {
	*tchttp.BaseResponse
	Response *GenerateCreateMangedTableSqlResponseParams `json:"Response"`
}

func (r *GenerateCreateMangedTableSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateCreateMangedTableSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOptimizerPolicyRequestParams struct {
	// 策略描述
	SmartPolicy *SmartPolicy `json:"SmartPolicy,omitnil,omitempty" name:"SmartPolicy"`
}

type GetOptimizerPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略描述
	SmartPolicy *SmartPolicy `json:"SmartPolicy,omitnil,omitempty" name:"SmartPolicy"`
}

func (r *GetOptimizerPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOptimizerPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SmartPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOptimizerPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOptimizerPolicyResponseParams struct {
	// 智能优化策略
	SmartOptimizerPolicy *SmartOptimizerPolicy `json:"SmartOptimizerPolicy,omitnil,omitempty" name:"SmartOptimizerPolicy"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetOptimizerPolicyResponse struct {
	*tchttp.BaseResponse
	Response *GetOptimizerPolicyResponseParams `json:"Response"`
}

func (r *GetOptimizerPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOptimizerPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantDLCCatalogAccessRequestParams struct {
	// 授权VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 产品(EMR|DLC|Doris|Inlong|Wedata)
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// VPC所属账号UIN
	VpcUin *string `json:"VpcUin,omitnil,omitempty" name:"VpcUin"`

	// VPC所属账号AppId
	VpcAppId *uint64 `json:"VpcAppId,omitnil,omitempty" name:"VpcAppId"`
}

type GrantDLCCatalogAccessRequest struct {
	*tchttp.BaseRequest
	
	// 授权VpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 产品(EMR|DLC|Doris|Inlong|Wedata)
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// VPC所属账号UIN
	VpcUin *string `json:"VpcUin,omitnil,omitempty" name:"VpcUin"`

	// VPC所属账号AppId
	VpcAppId *uint64 `json:"VpcAppId,omitnil,omitempty" name:"VpcAppId"`
}

func (r *GrantDLCCatalogAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantDLCCatalogAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	delete(f, "Product")
	delete(f, "Description")
	delete(f, "VpcUin")
	delete(f, "VpcAppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GrantDLCCatalogAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GrantDLCCatalogAccessResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GrantDLCCatalogAccessResponse struct {
	*tchttp.BaseResponse
	Response *GrantDLCCatalogAccessResponseParams `json:"Response"`
}

func (r *GrantDLCCatalogAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GrantDLCCatalogAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfo struct {
	// 用户组ID
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 策略类型
	StrategyType *string `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`
}

type HiveInfo struct {
	// hive metastore的地址
	MetaStoreUrl *string `json:"MetaStoreUrl,omitnil,omitempty" name:"MetaStoreUrl"`

	// hive数据源类型，代表数据储存的位置，COS或者HDFS
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源所在的私有网络信息
	Location *DatasourceConnectionLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 如果类型为HDFS，需要传一个用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 如果类型为HDFS，需要选择是否高可用
	HighAvailability *bool `json:"HighAvailability,omitnil,omitempty" name:"HighAvailability"`

	// 如果类型为COS，需要填写COS桶连接
	BucketUrl *string `json:"BucketUrl,omitnil,omitempty" name:"BucketUrl"`

	// json字符串。如果类型为HDFS，需要填写该字段
	HdfsProperties *string `json:"HdfsProperties,omitnil,omitempty" name:"HdfsProperties"`

	// Hive的元数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mysql *MysqlInfo `json:"Mysql,omitnil,omitempty" name:"Mysql"`

	// emr集群Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// emr集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// EMR集群中hive组件的版本号
	HiveVersion *string `json:"HiveVersion,omitnil,omitempty" name:"HiveVersion"`

	// Kerberos详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	KerberosInfo *KerberosInfo `json:"KerberosInfo,omitnil,omitempty" name:"KerberosInfo"`

	// 是否开启Kerberos
	// 注意：此字段可能返回 null，表示取不到有效值。
	KerberosEnable *bool `json:"KerberosEnable,omitnil,omitempty" name:"KerberosEnable"`
}

type HiveTablePartition struct {
	// 分区信息名称
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 分区记录数
	Records *int64 `json:"Records,omitnil,omitempty" name:"Records"`

	// 分区数据文件存储量
	DataFileStorage *int64 `json:"DataFileStorage,omitnil,omitempty" name:"DataFileStorage"`

	// 分区创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分区schema更新时间
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 最后一次分区更新的访问时间
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`
}

type HouseEventsInfo struct {
	// 事件时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time []*string `json:"Time,omitnil,omitempty" name:"Time"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventsAction []*string `json:"EventsAction,omitnil,omitempty" name:"EventsAction"`

	// 集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterInfo []*string `json:"ClusterInfo,omitnil,omitempty" name:"ClusterInfo"`
}

type IcebergTablePartition struct {
	// 分区信息名称
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 分区记录数
	Records *int64 `json:"Records,omitnil,omitempty" name:"Records"`

	// 分区数据文件数量
	DataFileSize *int64 `json:"DataFileSize,omitnil,omitempty" name:"DataFileSize"`

	// 分区数据文件存储量
	DataFileStorage *int64 `json:"DataFileStorage,omitnil,omitempty" name:"DataFileStorage"`

	// 分区创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分区更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最后一次分区更新的快照ID
	LastUpdateSnapshotId *string `json:"LastUpdateSnapshotId,omitnil,omitempty" name:"LastUpdateSnapshotId"`

	// 分区的location
	Location *LocationInfo `json:"Location,omitnil,omitempty" name:"Location"`
}

type IpPortPair struct {
	// ip信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 端口信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`
}

type JobLogResult struct {
	// 日志时间戳，毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志topic id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志topic name
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`

	// 日志内容，json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitnil,omitempty" name:"LogJson"`

	// 日志ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`
}

type KVPair struct {
	// 配置的key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 配置的value值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KafkaInfo struct {
	// kafka实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// kafka数据源的网络信息
	Location *DatasourceConnectionLocation `json:"Location,omitnil,omitempty" name:"Location"`
}

type KerberosInfo struct {
	// Krb5Conf文件值
	Krb5Conf *string `json:"Krb5Conf,omitnil,omitempty" name:"Krb5Conf"`

	// KeyTab文件值
	KeyTab *string `json:"KeyTab,omitnil,omitempty" name:"KeyTab"`

	// 服务主体
	ServicePrincipal *string `json:"ServicePrincipal,omitnil,omitempty" name:"ServicePrincipal"`
}

type LakeFileSystemToken struct {
	// Token使用的临时密钥的ID
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// Token使用的临时密钥
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// Token信息
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 过期时间
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 颁布时间
	IssueTime *int64 `json:"IssueTime,omitnil,omitempty" name:"IssueTime"`
}

type LakeFsInfo struct {
	// 托管存储名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 托管存储类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 容量
	SpaceUsedSize *float64 `json:"SpaceUsedSize,omitnil,omitempty" name:"SpaceUsedSize"`

	// 创建时候的时间戳
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitnil,omitempty" name:"CreateTimeStamp"`

	// 是否是用户默认桶，0：默认桶，1：非默认桶
	DefaultBucket *int64 `json:"DefaultBucket,omitnil,omitempty" name:"DefaultBucket"`

	// 托管存储short name
	ShortName *string `json:"ShortName,omitnil,omitempty" name:"ShortName"`

	// 桶描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 托管桶状态，当前取值为：creating、bind、readOnly、isolate
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type LaunchStandardEngineResourceGroupsRequestParams struct {
	// 标准引擎资源组名称
	EngineResourceGroupNames []*string `json:"EngineResourceGroupNames,omitnil,omitempty" name:"EngineResourceGroupNames"`
}

type LaunchStandardEngineResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 标准引擎资源组名称
	EngineResourceGroupNames []*string `json:"EngineResourceGroupNames,omitnil,omitempty" name:"EngineResourceGroupNames"`
}

func (r *LaunchStandardEngineResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LaunchStandardEngineResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LaunchStandardEngineResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LaunchStandardEngineResourceGroupsResponseParams struct {
	// 批量操作资源组时，操作失败的资源组相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateEngineResourceGroupFailMessages []*OperateEngineResourceGroupFailMessage `json:"OperateEngineResourceGroupFailMessages,omitnil,omitempty" name:"OperateEngineResourceGroupFailMessages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LaunchStandardEngineResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *LaunchStandardEngineResourceGroupsResponseParams `json:"Response"`
}

func (r *LaunchStandardEngineResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LaunchStandardEngineResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskJobLogDetailRequestParams struct {
	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下一次分页参数，第一次传空
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 列表返回的Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 最近1000条日志是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// SparkSQL任务唯一ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`
}

type ListTaskJobLogDetailRequest struct {
	*tchttp.BaseRequest
	
	// 开始运行时间，unix时间戳（毫秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束运行时间，unix时间戳（毫秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页大小，最大1000，配合Context一起使用
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 下一次分页参数，第一次传空
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 列表返回的Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 最近1000条日志是否升序排列，true:升序排序，false:倒序，默认false，倒序排列
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`

	// 预览日志的通用过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// SparkSQL任务唯一ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`
}

func (r *ListTaskJobLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskJobLogDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "TaskId")
	delete(f, "Asc")
	delete(f, "Filters")
	delete(f, "BatchId")
	delete(f, "DataEngineId")
	delete(f, "ResourceGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskJobLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskJobLogDetailResponseParams struct {
	// 下一次分页参数
	Context *string `json:"Context,omitnil,omitempty" name:"Context"`

	// 是否获取完结
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 日志详情
	Results []*JobLogResult `json:"Results,omitnil,omitempty" name:"Results"`

	// 日志url(字段已废弃)
	LogUrl *string `json:"LogUrl,omitnil,omitempty" name:"LogUrl"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskJobLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskJobLogDetailResponseParams `json:"Response"`
}

func (r *ListTaskJobLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskJobLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskJobLogNameRequestParams struct {
	// 查询的taskId
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// SparkSQL批任务唯一ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

type ListTaskJobLogNameRequest struct {
	*tchttp.BaseRequest
	
	// 查询的taskId
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// SparkSQL批任务唯一ID
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

func (r *ListTaskJobLogNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskJobLogNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "BatchId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListTaskJobLogNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListTaskJobLogNameResponseParams struct {
	// 日志名称列表
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListTaskJobLogNameResponse struct {
	*tchttp.BaseResponse
	Response *ListTaskJobLogNameResponseParams `json:"Response"`
}

func (r *ListTaskJobLogNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListTaskJobLogNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LocationInfo struct {
	// 桶名称
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// location路径
	DataLocation *string `json:"DataLocation,omitnil,omitempty" name:"DataLocation"`
}

type LockComponentInfo struct {
	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 表名称
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 分区
	Partition *string `json:"Partition,omitnil,omitempty" name:"Partition"`

	// 锁类型：SHARED_READ、SHARED_WRITE、EXCLUSIVE
	LockType *string `json:"LockType,omitnil,omitempty" name:"LockType"`

	// 锁级别：DB、TABLE、PARTITION
	LockLevel *string `json:"LockLevel,omitnil,omitempty" name:"LockLevel"`

	// 锁操作：SELECT,INSERT,UPDATE,DELETE,UNSET,NO_TXN
	DataOperationType *string `json:"DataOperationType,omitnil,omitempty" name:"DataOperationType"`

	// 是否保持Acid
	IsAcid *bool `json:"IsAcid,omitnil,omitempty" name:"IsAcid"`

	// 是否动态分区写
	IsDynamicPartitionWrite *bool `json:"IsDynamicPartitionWrite,omitnil,omitempty" name:"IsDynamicPartitionWrite"`
}

// Predefined struct for user
type LockMetaDataRequestParams struct {
	// 加锁内容
	LockComponentList []*LockComponentInfo `json:"LockComponentList,omitnil,omitempty" name:"LockComponentList"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 事务id
	TxnId *int64 `json:"TxnId,omitnil,omitempty" name:"TxnId"`

	// 客户端信息
	AgentInfo *string `json:"AgentInfo,omitnil,omitempty" name:"AgentInfo"`

	// 主机名
	Hostname *string `json:"Hostname,omitnil,omitempty" name:"Hostname"`
}

type LockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 加锁内容
	LockComponentList []*LockComponentInfo `json:"LockComponentList,omitnil,omitempty" name:"LockComponentList"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 事务id
	TxnId *int64 `json:"TxnId,omitnil,omitempty" name:"TxnId"`

	// 客户端信息
	AgentInfo *string `json:"AgentInfo,omitnil,omitempty" name:"AgentInfo"`

	// 主机名
	Hostname *string `json:"Hostname,omitnil,omitempty" name:"Hostname"`
}

func (r *LockMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LockComponentList")
	delete(f, "DatasourceConnectionName")
	delete(f, "TxnId")
	delete(f, "AgentInfo")
	delete(f, "Hostname")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LockMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockMetaDataResponseParams struct {
	// 锁id
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 锁状态：ACQUIRED、WAITING、ABORT、NOT_ACQUIRED
	LockState *string `json:"LockState,omitnil,omitempty" name:"LockState"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type LockMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *LockMetaDataResponseParams `json:"Response"`
}

func (r *LockMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MixedTablePartitions struct {
	// 数据表格式
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`

	// 分区总数
	TotalSize *int64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 分页查询的游标信息，在获取下一页信息时需要回传到服务端
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// iceberg表分区信息
	IcebergPartitions []*IcebergTablePartition `json:"IcebergPartitions,omitnil,omitempty" name:"IcebergPartitions"`

	// hive表分区信息
	HivePartitions []*HiveTablePartition `json:"HivePartitions,omitnil,omitempty" name:"HivePartitions"`
}

// Predefined struct for user
type ModifyAdvancedStoreLocationRequestParams struct {
	// 查询结果保存cos路径
	StoreLocation *string `json:"StoreLocation,omitnil,omitempty" name:"StoreLocation"`

	// 是否启用高级设置：0-否，1-是
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type ModifyAdvancedStoreLocationRequest struct {
	*tchttp.BaseRequest
	
	// 查询结果保存cos路径
	StoreLocation *string `json:"StoreLocation,omitnil,omitempty" name:"StoreLocation"`

	// 是否启用高级设置：0-否，1-是
	Enable *uint64 `json:"Enable,omitnil,omitempty" name:"Enable"`
}

func (r *ModifyAdvancedStoreLocationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdvancedStoreLocationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StoreLocation")
	delete(f, "Enable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAdvancedStoreLocationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAdvancedStoreLocationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAdvancedStoreLocationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAdvancedStoreLocationResponseParams `json:"Response"`
}

func (r *ModifyAdvancedStoreLocationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAdvancedStoreLocationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataEngineDescriptionRequestParams struct {
	// 要修改的引擎的名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 引擎的描述信息，最大长度为250
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ModifyDataEngineDescriptionRequest struct {
	*tchttp.BaseRequest
	
	// 要修改的引擎的名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 引擎的描述信息，最大长度为250
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

func (r *ModifyDataEngineDescriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataEngineDescriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Message")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataEngineDescriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataEngineDescriptionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDataEngineDescriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataEngineDescriptionResponseParams `json:"Response"`
}

func (r *ModifyDataEngineDescriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataEngineDescriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernEventRuleRequestParams struct {

}

type ModifyGovernEventRuleRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ModifyGovernEventRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernEventRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGovernEventRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernEventRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGovernEventRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGovernEventRuleResponseParams `json:"Response"`
}

func (r *ModifyGovernEventRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernEventRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppBatchRequestParams struct {
	// 需要批量修改的Spark作业任务ID列表
	SparkAppId []*string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// 引擎ID
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppDriverSize *string `json:"AppDriverSize,omitnil,omitempty" name:"AppDriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitnil,omitempty" name:"AppExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	AppExecutorNums *uint64 `json:"AppExecutorNums,omitnil,omitempty" name:"AppExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	AppExecutorMaxNumbers *uint64 `json:"AppExecutorMaxNumbers,omitnil,omitempty" name:"AppExecutorMaxNumbers"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`
}

type ModifySparkAppBatchRequest struct {
	*tchttp.BaseRequest
	
	// 需要批量修改的Spark作业任务ID列表
	SparkAppId []*string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// 引擎ID
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppDriverSize *string `json:"AppDriverSize,omitnil,omitempty" name:"AppDriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	AppExecutorSize *string `json:"AppExecutorSize,omitnil,omitempty" name:"AppExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	AppExecutorNums *uint64 `json:"AppExecutorNums,omitnil,omitempty" name:"AppExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	AppExecutorMaxNumbers *uint64 `json:"AppExecutorMaxNumbers,omitnil,omitempty" name:"AppExecutorMaxNumbers"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`
}

func (r *ModifySparkAppBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppBatchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SparkAppId")
	delete(f, "DataEngine")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "AppExecutorMaxNumbers")
	delete(f, "IsInherit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySparkAppBatchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppBatchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySparkAppBatchResponse struct {
	*tchttp.BaseResponse
	Response *ModifySparkAppBatchResponseParams `json:"Response"`
}

func (r *ModifySparkAppBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppRequestParams struct {
	// spark作业名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil,omitempty" name:"AppFile"`

	// 数据访问策略，CAM Role arn
	RoleArn *int64 `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil,omitempty" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil,omitempty" name:"AppExecutorNums"`

	// spark作业Id
	SparkAppId *string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil,omitempty" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil,omitempty" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil,omitempty" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil,omitempty" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil,omitempty" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil,omitempty" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil,omitempty" name:"AppFiles"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil,omitempty" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil,omitempty" name:"AppPythonFiles"`

	// spark作业程序入参
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil,omitempty" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil,omitempty" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil,omitempty" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务资源配置是否继承集群配置模板：0（默认）不继承、1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil,omitempty" name:"IsSessionStarted"`

	// 标准引擎依赖包
	DependencyPackages []*DependencyPackage `json:"DependencyPackages,omitnil,omitempty" name:"DependencyPackages"`
}

type ModifySparkAppRequest struct {
	*tchttp.BaseRequest
	
	// spark作业名
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// spark作业类型，1代表spark jar作业，2代表spark streaming作业
	AppType *int64 `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 执行spark作业的数据引擎名称
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// spark作业程序包文件路径
	AppFile *string `json:"AppFile,omitnil,omitempty" name:"AppFile"`

	// 数据访问策略，CAM Role arn
	RoleArn *int64 `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppDriverSize *string `json:"AppDriverSize,omitnil,omitempty" name:"AppDriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	AppExecutorSize *string `json:"AppExecutorSize,omitnil,omitempty" name:"AppExecutorSize"`

	// spark作业executor个数
	AppExecutorNums *int64 `json:"AppExecutorNums,omitnil,omitempty" name:"AppExecutorNums"`

	// spark作业Id
	SparkAppId *string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil,omitempty" name:"Eni"`

	// spark作业程序包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocal *string `json:"IsLocal,omitnil,omitempty" name:"IsLocal"`

	// spark作业主类
	MainClass *string `json:"MainClass,omitnil,omitempty" name:"MainClass"`

	// spark配置，以换行符分隔
	AppConf *string `json:"AppConf,omitnil,omitempty" name:"AppConf"`

	// spark 作业依赖jar包是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalJars *string `json:"IsLocalJars,omitnil,omitempty" name:"IsLocalJars"`

	// spark 作业依赖jar包（--jars），以逗号分隔
	AppJars *string `json:"AppJars,omitnil,omitempty" name:"AppJars"`

	// spark作业依赖文件资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalFiles *string `json:"IsLocalFiles,omitnil,omitempty" name:"IsLocalFiles"`

	// spark作业依赖文件资源（--files）（非jar、zip），以逗号分隔
	AppFiles *string `json:"AppFiles,omitnil,omitempty" name:"AppFiles"`

	// pyspark：依赖上传方式，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil,omitempty" name:"IsLocalPythonFiles"`

	// pyspark作业依赖python资源（--py-files），支持py/zip/egg等归档格式，多文件以逗号分隔
	AppPythonFiles *string `json:"AppPythonFiles,omitnil,omitempty" name:"AppPythonFiles"`

	// spark作业程序入参
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 最大重试次数，只对spark流任务生效
	MaxRetries *int64 `json:"MaxRetries,omitnil,omitempty" name:"MaxRetries"`

	// 数据源名
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// spark作业依赖archives资源是否本地上传，cos：存放与cos，lakefs：本地上传（控制台使用，该方式不支持直接接口调用）
	IsLocalArchives *string `json:"IsLocalArchives,omitnil,omitempty" name:"IsLocalArchives"`

	// spark作业依赖archives资源（--archives），支持tar.gz/tgz/tar等归档格式，以逗号分隔
	AppArchives *string `json:"AppArchives,omitnil,omitempty" name:"AppArchives"`

	// Spark Image 版本号
	SparkImage *string `json:"SparkImage,omitnil,omitempty" name:"SparkImage"`

	// Spark Image 版本名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil,omitempty" name:"SparkImageVersion"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于AppExecutorNums
	AppExecutorMaxNumbers *int64 `json:"AppExecutorMaxNumbers,omitnil,omitempty" name:"AppExecutorMaxNumbers"`

	// 关联dlc查询脚本
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 任务资源配置是否继承集群配置模板：0（默认）不继承、1：继承
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil,omitempty" name:"IsSessionStarted"`

	// 标准引擎依赖包
	DependencyPackages []*DependencyPackage `json:"DependencyPackages,omitnil,omitempty" name:"DependencyPackages"`
}

func (r *ModifySparkAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "AppType")
	delete(f, "DataEngine")
	delete(f, "AppFile")
	delete(f, "RoleArn")
	delete(f, "AppDriverSize")
	delete(f, "AppExecutorSize")
	delete(f, "AppExecutorNums")
	delete(f, "SparkAppId")
	delete(f, "Eni")
	delete(f, "IsLocal")
	delete(f, "MainClass")
	delete(f, "AppConf")
	delete(f, "IsLocalJars")
	delete(f, "AppJars")
	delete(f, "IsLocalFiles")
	delete(f, "AppFiles")
	delete(f, "IsLocalPythonFiles")
	delete(f, "AppPythonFiles")
	delete(f, "CmdArgs")
	delete(f, "MaxRetries")
	delete(f, "DataSource")
	delete(f, "IsLocalArchives")
	delete(f, "AppArchives")
	delete(f, "SparkImage")
	delete(f, "SparkImageVersion")
	delete(f, "AppExecutorMaxNumbers")
	delete(f, "SessionId")
	delete(f, "IsInherit")
	delete(f, "IsSessionStarted")
	delete(f, "DependencyPackages")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySparkAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySparkAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySparkAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifySparkAppResponseParams `json:"Response"`
}

func (r *ModifySparkAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySparkAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserRequestParams struct {
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户描述
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`
}

type ModifyUserRequest struct {
	*tchttp.BaseRequest
	
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户描述
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`
}

func (r *ModifyUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserResponseParams `json:"Response"`
}

func (r *ModifyUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserTypeRequestParams struct {
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户要修改到的类型，ADMIN：管理员，COMMON：一般用户。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

type ModifyUserTypeRequest struct {
	*tchttp.BaseRequest
	
	// 用户ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户要修改到的类型，ADMIN：管理员，COMMON：一般用户。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`
}

func (r *ModifyUserTypeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserTypeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UserId")
	delete(f, "UserType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserTypeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserTypeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserTypeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserTypeResponseParams `json:"Response"`
}

func (r *ModifyUserTypeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkGroupRequestParams struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 工作组描述，最大字符数限制50
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil,omitempty" name:"WorkGroupDescription"`
}

type ModifyWorkGroupRequest struct {
	*tchttp.BaseRequest
	
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 工作组描述，最大字符数限制50
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil,omitempty" name:"WorkGroupDescription"`
}

func (r *ModifyWorkGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkGroupId")
	delete(f, "WorkGroupDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyWorkGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkGroupResponseParams `json:"Response"`
}

func (r *ModifyWorkGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MountPointAssociates struct {
	// 桶Id
	BucketId *string `json:"BucketId,omitnil,omitempty" name:"BucketId"`

	// vpcId
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网地址
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 权限组Id
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`

	// 权限规则Id
	AccessRuleId *int64 `json:"AccessRuleId,omitnil,omitempty" name:"AccessRuleId"`
}

type MysqlInfo struct {
	// 连接mysql的jdbc url
	JdbcUrl *string `json:"JdbcUrl,omitnil,omitempty" name:"JdbcUrl"`

	// 用户名
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// mysql密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// mysql数据源的网络信息
	Location *DatasourceConnectionLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 数据库实例ID，和数据库侧保持一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库实例名称，和数据库侧保持一致
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type NetWork struct {
	// 服务clbip
	ClbIp *string `json:"ClbIp,omitnil,omitempty" name:"ClbIp"`

	// 服务clbPort
	ClbPort *string `json:"ClbPort,omitnil,omitempty" name:"ClbPort"`

	// vpc实例id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc网段
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 子网实例id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 子网网段
	SubnetCidrBlock *string `json:"SubnetCidrBlock,omitnil,omitempty" name:"SubnetCidrBlock"`
}

type NetworkConnection struct {
	// 网络配置id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 网络配置唯一标志符
	AssociateId *string `json:"AssociateId,omitnil,omitempty" name:"AssociateId"`

	// 计算引擎id
	HouseId *string `json:"HouseId,omitnil,omitempty" name:"HouseId"`

	// 数据源id(已废弃)
	DatasourceConnectionId *string `json:"DatasourceConnectionId,omitnil,omitempty" name:"DatasourceConnectionId"`

	// 网络配置状态（0-初始化，1-正常）
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创建用户Appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	Appid *int64 `json:"Appid,omitnil,omitempty" name:"Appid"`

	// 计算引擎名称
	HouseName *string `json:"HouseName,omitnil,omitempty" name:"HouseName"`

	// 网络配置名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 网络配置类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionType *int64 `json:"NetworkConnectionType,omitnil,omitempty" name:"NetworkConnectionType"`

	// 创建用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 创建用户SubAccountUin
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubAccountUin *string `json:"SubAccountUin,omitnil,omitempty" name:"SubAccountUin"`

	// 网络配置描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConnectionDesc *string `json:"NetworkConnectionDesc,omitnil,omitempty" name:"NetworkConnectionDesc"`

	// 数据源vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionVpcId *string `json:"DatasourceConnectionVpcId,omitnil,omitempty" name:"DatasourceConnectionVpcId"`

	// 数据源SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionSubnetId *string `json:"DatasourceConnectionSubnetId,omitnil,omitempty" name:"DatasourceConnectionSubnetId"`

	// 数据源SubnetId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionCidrBlock *string `json:"DatasourceConnectionCidrBlock,omitnil,omitempty" name:"DatasourceConnectionCidrBlock"`

	// 数据源SubnetCidrBlock
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionSubnetCidrBlock *string `json:"DatasourceConnectionSubnetCidrBlock,omitnil,omitempty" name:"DatasourceConnectionSubnetCidrBlock"`

	// 支持 eg
	EGSupport *int64 `json:"EGSupport,omitnil,omitempty" name:"EGSupport"`
}

type NotebookSessionInfo struct {
	// Session名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// DLC Spark作业引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// Session相关配置，当前支持：eni、roleArn以及用户指定的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Arguments []*KVPair `json:"Arguments,omitnil,omitempty" name:"Arguments"`

	// 运行程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentFiles []*string `json:"ProgramDependentFiles,omitnil,omitempty" name:"ProgramDependentFiles"`

	// 依赖的jar程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentJars []*string `json:"ProgramDependentJars,omitnil,omitempty" name:"ProgramDependentJars"`

	// 依赖的python程序地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramDependentPython []*string `json:"ProgramDependentPython,omitnil,omitempty" name:"ProgramDependentPython"`

	// 依赖的pyspark虚拟环境地址，当前支持：cosn://和lakefs://两种路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgramArchives []*string `json:"ProgramArchives,omitnil,omitempty" name:"ProgramArchives"`

	// 指定的Driver规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// 指定的Executor规格，当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu）
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定的Executor数量，默认为1
	ExecutorNumbers *uint64 `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitnil,omitempty" name:"ProxyUser"`

	// 指定的Session超时时间，单位秒，默认3600秒
	TimeoutInSecond *int64 `json:"TimeoutInSecond,omitnil,omitempty" name:"TimeoutInSecond"`

	// Spark任务返回的AppId
	SparkAppId *string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Session创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 其它信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppInfo []*KVPair `json:"AppInfo,omitnil,omitempty" name:"AppInfo"`

	// Spark ui地址
	SparkUiUrl *string `json:"SparkUiUrl,omitnil,omitempty" name:"SparkUiUrl"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于ExecutorNumbers
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// session类型，group：资源组下session independent：独立资源session， 不依赖资源组
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// session，pod大小
	PodSize *int64 `json:"PodSize,omitnil,omitempty" name:"PodSize"`

	// pod数量
	PodNumbers *int64 `json:"PodNumbers,omitnil,omitempty" name:"PodNumbers"`
}

type NotebookSessionStatementBatchInformation struct {
	// 任务详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotebookSessionStatementBatch []*NotebookSessionStatementInfo `json:"NotebookSessionStatementBatch,omitnil,omitempty" name:"NotebookSessionStatementBatch"`

	// 当前批任务是否运行完成
	IsAvailable *bool `json:"IsAvailable,omitnil,omitempty" name:"IsAvailable"`

	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// Batch唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`
}

type NotebookSessionStatementInfo struct {
	// 完成时间戳
	Completed *int64 `json:"Completed,omitnil,omitempty" name:"Completed"`

	// 开始时间戳
	Started *int64 `json:"Started,omitnil,omitempty" name:"Started"`

	// 完成进度，百分制
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// Session Statement唯一标识
	StatementId *string `json:"StatementId,omitnil,omitempty" name:"StatementId"`

	// Session Statement状态，包含：waiting（排队中）、running（运行中）、available（正常）、error（异常）、cancelling（取消中）、cancelled（已取消）
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Statement输出信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutPut *StatementOutput `json:"OutPut,omitnil,omitempty" name:"OutPut"`

	// 批任务id
	BatchId *string `json:"BatchId,omitnil,omitempty" name:"BatchId"`

	// 运行语句
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type NotebookSessions struct {
	// 类型，当前支持：spark、pyspark、sparkr、sql
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// Session唯一标识
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 代理用户，默认为root
	ProxyUser *string `json:"ProxyUser,omitnil,omitempty" name:"ProxyUser"`

	// Session状态，包含：not_started（未启动）、starting（已启动）、idle（等待输入）、busy(正在运行statement)、shutting_down（停止）、error（异常）、dead（已退出）、killed（被杀死）、success（正常停止）
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// Spark任务返回的AppId
	SparkAppId *string `json:"SparkAppId,omitnil,omitempty" name:"SparkAppId"`

	// Session名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Session创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 最新的运行时间
	LastRunningTime *string `json:"LastRunningTime,omitnil,omitempty" name:"LastRunningTime"`

	// 创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// spark ui地址
	SparkUiUrl *string `json:"SparkUiUrl,omitnil,omitempty" name:"SparkUiUrl"`

	// session类型，group：资源组session independent：独立资源session，不依赖资源组
	SessionType *string `json:"SessionType,omitnil,omitempty" name:"SessionType"`

	// 引擎id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 资源组名字
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

type OpendThirdAccessUserInfo struct {
	// id信息
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 用户主UIN
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 用户AppId
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 开通时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type OperateEngineResourceGroupFailMessage struct {
	// 引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 操作失败的提示信息
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`
}

type OptimizerEngineInfo struct {
	// 引擎资源名称
	HouseName *string `json:"HouseName,omitnil,omitempty" name:"HouseName"`

	// 引擎资源ID
	HouseId *string `json:"HouseId,omitnil,omitempty" name:"HouseId"`

	// 该参数仅针对spark作业引擎有效，用于执行数据优化任务的资源大小，不填时将采用该引擎所有资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	HouseSize *int64 `json:"HouseSize,omitnil,omitempty" name:"HouseSize"`
}

type Other struct {
	// 枚举类型，默认值为Json，可选值为[Json, Parquet, ORC, AVRD]之一。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`
}

type OtherCHDFSBinding struct {
	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 用户名称（该字段已废弃）
	SuperUser []*string `json:"SuperUser,omitnil,omitempty" name:"SuperUser"`

	// vpc配置信息
	VpcInfo []*CHDFSProductVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 是否与该桶绑定（该字段已废弃）
	IsBind *bool `json:"IsBind,omitnil,omitempty" name:"IsBind"`
}

type OtherDatasourceConnection struct {
	// 网络参数
	Location *DatasourceConnectionLocation `json:"Location,omitnil,omitempty" name:"Location"`
}

type Param struct {
	// 参数key，例如：
	ConfigItem *string `json:"ConfigItem,omitnil,omitempty" name:"ConfigItem"`

	// 参数值
	ConfigValue *string `json:"ConfigValue,omitnil,omitempty" name:"ConfigValue"`

	// 下发操作，支持：ADD、DELETE、MODIFY
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`
}

type Partition struct {
	// 分区列名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分区类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 对分区的描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 隐式分区转换策略
	Transform *string `json:"Transform,omitnil,omitempty" name:"Transform"`

	// 转换策略参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransformArgs []*string `json:"TransformArgs,omitnil,omitempty" name:"TransformArgs"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

// Predefined struct for user
type PauseStandardEngineResourceGroupsRequestParams struct {
	// 标准引擎资源组名称
	EngineResourceGroupNames []*string `json:"EngineResourceGroupNames,omitnil,omitempty" name:"EngineResourceGroupNames"`
}

type PauseStandardEngineResourceGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 标准引擎资源组名称
	EngineResourceGroupNames []*string `json:"EngineResourceGroupNames,omitnil,omitempty" name:"EngineResourceGroupNames"`
}

func (r *PauseStandardEngineResourceGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseStandardEngineResourceGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseStandardEngineResourceGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseStandardEngineResourceGroupsResponseParams struct {
	// 批量操作资源组时，操作失败的资源组相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperateEngineResourceGroupFailMessages []*OperateEngineResourceGroupFailMessage `json:"OperateEngineResourceGroupFailMessages,omitnil,omitempty" name:"OperateEngineResourceGroupFailMessages"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseStandardEngineResourceGroupsResponse struct {
	*tchttp.BaseResponse
	Response *PauseStandardEngineResourceGroupsResponseParams `json:"Response"`
}

func (r *PauseStandardEngineResourceGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseStandardEngineResourceGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Policy struct {
	// 需要授权的数据库名，填 * 代表当前Catalog下所有数据库。当授权类型为管理员级别时，只允许填 “*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定数据库。
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 需要授权的数据源名称，管理员级别下只支持填  * （代表该级别全部资源）；数据源级别和数据库级别鉴权的情况下，只支持填COSDataCatalog或者*；在数据表级别鉴权下可以填写用户自定义数据源。不填情况下默认为DataLakeCatalog。注意：如果是对用户自定义数据源进行鉴权，DLC能够管理的权限是用户接入数据源的时候提供的账户的子集。
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 需要授权的表名，填 * 代表当前Database下所有表。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定数据表。
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 授权的权限操作，对于不同级别的鉴权提供不同操作。管理员权限：ALL，不填默认为ALL；数据连接级鉴权：CREATE；数据库级别鉴权：ALL、CREATE、ALTER、DROP；数据表权限：ALL、SELECT、INSERT、ALTER、DELETE、DROP、UPDATE。注意：在数据表权限下，指定的数据源不为COSDataCatalog的时候，只支持SELECT操作。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 授权类型，现在支持八种授权类型：ADMIN:管理员级别鉴权 DATASOURCE：数据连接级别鉴权 DATABASE：数据库级别鉴权 TABLE：表级别鉴权 VIEW：视图级别鉴权 FUNCTION：函数级别鉴权 COLUMN：列级别鉴权 ENGINE：数据引擎鉴权。不填默认为管理员级别鉴权。
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// 需要授权的函数名，填 * 代表当前Catalog下所有函数。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别时只允许填空，其他类型下可以任意指定函数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Function *string `json:"Function,omitnil,omitempty" name:"Function"`

	// 需要授权的视图，填 * 代表当前Database下所有视图。当授权类型为管理员级别时，只允许填“*”，当授权类型为数据连接级别、数据库级别时只允许填空，其他类型下可以任意指定视图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	View *string `json:"View,omitnil,omitempty" name:"View"`

	// 需要授权的列，填 * 代表当前所有列。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	Column *string `json:"Column,omitnil,omitempty" name:"Column"`

	// 需要授权的数据引擎，填 * 代表当前所有引擎。当授权类型为管理员级别时，只允许填“*”
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// 用户是否可以进行二次授权。当为true的时候，被授权的用户可以将本次获取的权限再次授权给其他子用户。默认为false
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReAuth *bool `json:"ReAuth,omitnil,omitempty" name:"ReAuth"`

	// 权限来源，入参不填。USER：权限来自用户本身；WORKGROUP：权限来自绑定的工作组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 授权模式，入参不填。COMMON：普通模式；SENIOR：高级模式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 操作者，入参不填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 权限创建的时间，入参不填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 权限所属工作组的ID，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceId *int64 `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// 权限所属工作组的名称，只有当该权限的来源为工作组时才会有值。即仅当Source字段的值为WORKGROUP时该字段才有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 引擎类型
	EngineGeneration *string `json:"EngineGeneration,omitnil,omitempty" name:"EngineGeneration"`
}

type Policys struct {
	// 策略集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 策略总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type PrestoMonitorMetrics struct {
	// 	Alluxio本地缓存命中率
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalCacheHitRate *float64 `json:"LocalCacheHitRate,omitnil,omitempty" name:"LocalCacheHitRate"`

	// Fragment缓存命中率
	// 注意：此字段可能返回 null，表示取不到有效值。
	FragmentCacheHitRate *float64 `json:"FragmentCacheHitRate,omitnil,omitempty" name:"FragmentCacheHitRate"`
}

type Property struct {
	// 属性key名称。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 属性key对应的value。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PythonSparkImage struct {
	// spark镜像唯一id
	SparkImageId *string `json:"SparkImageId,omitnil,omitempty" name:"SparkImageId"`

	// 集群小版本镜像id
	ChildImageVersionId *string `json:"ChildImageVersionId,omitnil,omitempty" name:"ChildImageVersionId"`

	// spark镜像名称
	SparkImageVersion *string `json:"SparkImageVersion,omitnil,omitempty" name:"SparkImageVersion"`

	// spark镜像描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type QueryInternalTableWarehouseRequestParams struct {
	// 库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

type QueryInternalTableWarehouseRequest struct {
	*tchttp.BaseRequest
	
	// 库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`
}

func (r *QueryInternalTableWarehouseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInternalTableWarehouseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseName")
	delete(f, "TableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryInternalTableWarehouseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryInternalTableWarehouseResponseParams struct {
	// warehouse路径
	WarehousePath *string `json:"WarehousePath,omitnil,omitempty" name:"WarehousePath"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryInternalTableWarehouseResponse struct {
	*tchttp.BaseResponse
	Response *QueryInternalTableWarehouseResponseParams `json:"Response"`
}

func (r *QueryInternalTableWarehouseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryInternalTableWarehouseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResultRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// objectListMarker={marker}&lastReadFile={filename}&lastReadOffsetlastReadFile为上一次读取的文件，lastReadOffset为上一次读取到的位置
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

type QueryResultRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// objectListMarker={marker}&lastReadFile={filename}&lastReadOffsetlastReadFile为上一次读取的文件，lastReadOffset为上一次读取到的位置
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`
}

func (r *QueryResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "NextToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryResultResponseParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 结果数据
	ResultSet *string `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`

	// schema
	ResultSchema []*Column `json:"ResultSchema,omitnil,omitempty" name:"ResultSchema"`

	// 分页信息
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryResultResponse struct {
	*tchttp.BaseResponse
	Response *QueryResultResponseParams `json:"Response"`
}

func (r *QueryResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryTaskCostDetailRequestParams struct {
	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 下一页的标识
	SearchAfter *string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// 每页的大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type QueryTaskCostDetailRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件，如下支持的过滤类型，传参Name应为以下其中一个,其中task-id支持最大50个过滤个数，其他过滤参数支持的总数不超过5个。
	// task-id - String - （任务ID准确过滤）task-id取值形如：e386471f-139a-4e59-877f-50ece8135b99。
	// task-state - String - （任务状态过滤）取值范围 0(初始化)， 1(运行中)， 2(成功)， -1(失败)。
	// task-sql-keyword - String - （SQL语句关键字模糊过滤）取值形如：DROP TABLE。
	// task-operator- string （子uin过滤）
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 起始时间点，格式为yyyy-mm-dd HH:MM:SS。默认为45天前的当前时刻
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间点，格式为yyyy-mm-dd HH:MM:SS时间跨度在(0,30天]，支持最近45天数据查询。默认为当前时刻
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 数据引擎名称，用于筛选
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 下一页的标识
	SearchAfter *string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// 每页的大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *QueryTaskCostDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTaskCostDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "DataEngineName")
	delete(f, "SearchAfter")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "QueryTaskCostDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type QueryTaskCostDetailResponseParams struct {
	// 下一页的标识
	SearchAfter *string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// 返回的数据
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type QueryTaskCostDetailResponse struct {
	*tchttp.BaseResponse
	Response *QueryTaskCostDetailResponseParams `json:"Response"`
}

func (r *QueryTaskCostDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *QueryTaskCostDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterThirdPartyAccessUserRequestParams struct {

}

type RegisterThirdPartyAccessUserRequest struct {
	*tchttp.BaseRequest
	
}

func (r *RegisterThirdPartyAccessUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterThirdPartyAccessUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterThirdPartyAccessUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterThirdPartyAccessUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RegisterThirdPartyAccessUserResponse struct {
	*tchttp.BaseResponse
	Response *RegisterThirdPartyAccessUserResponseParams `json:"Response"`
}

func (r *RegisterThirdPartyAccessUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterThirdPartyAccessUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDataEngineRequestParams struct {
	// CU队列名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 续费时长，单位月，最少续费1一个月
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 付费类型，默认为1，预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 单位，默认m，仅能填m
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 自动续费标志，0，初始状态，默认不自动续费，若用户有预付费不停服特权，自动续费。1：自动续费。2：明确不自动续费。不传该参数默认为0
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type RenewDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// CU队列名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 续费时长，单位月，最少续费1一个月
	TimeSpan *int64 `json:"TimeSpan,omitnil,omitempty" name:"TimeSpan"`

	// 付费类型，默认为1，预付费
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 单位，默认m，仅能填m
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 自动续费标志，0，初始状态，默认不自动续费，若用户有预付费不停服特权，自动续费。1：自动续费。2：明确不自动续费。不传该参数默认为0
	RenewFlag *int64 `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

func (r *RenewDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "TimeSpan")
	delete(f, "PayMode")
	delete(f, "TimeUnit")
	delete(f, "RenewFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewDataEngineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *RenewDataEngineResponseParams `json:"Response"`
}

func (r *RenewDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportHeartbeatMetaDataRequestParams struct {
	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 锁ID
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil,omitempty" name:"TxnId"`
}

type ReportHeartbeatMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 锁ID
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 事务ID
	TxnId *int64 `json:"TxnId,omitnil,omitempty" name:"TxnId"`
}

func (r *ReportHeartbeatMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportHeartbeatMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceConnectionName")
	delete(f, "LockId")
	delete(f, "TxnId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReportHeartbeatMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReportHeartbeatMetaDataResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReportHeartbeatMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *ReportHeartbeatMetaDataResponseParams `json:"Response"`
}

func (r *ReportHeartbeatMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReportHeartbeatMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceInfo struct {
	// 归属类型
	AttributionType *string `json:"AttributionType,omitnil,omitempty" name:"AttributionType"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 引擎名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 如资源类型为spark-sql 取值为Name, 如为spark-batch 取值为session app_name
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 亲和性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Favor []*FavorInfo `json:"Favor,omitnil,omitempty" name:"Favor"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 标准引擎资源组信息
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`
}

// Predefined struct for user
type RestartDataEngineRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 是否强制重启，忽略任务
	ForcedOperation *bool `json:"ForcedOperation,omitnil,omitempty" name:"ForcedOperation"`
}

type RestartDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 是否强制重启，忽略任务
	ForcedOperation *bool `json:"ForcedOperation,omitnil,omitempty" name:"ForcedOperation"`
}

func (r *RestartDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "ForcedOperation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartDataEngineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *RestartDataEngineResponseParams `json:"Response"`
}

func (r *RestartDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeDLCCatalogAccessRequestParams struct {
	// VpcID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type RevokeDLCCatalogAccessRequest struct {
	*tchttp.BaseRequest
	
	// VpcID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

func (r *RevokeDLCCatalogAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeDLCCatalogAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "VpcId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RevokeDLCCatalogAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RevokeDLCCatalogAccessResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RevokeDLCCatalogAccessResponse struct {
	*tchttp.BaseResponse
	Response *RevokeDLCCatalogAccessResponseParams `json:"Response"`
}

func (r *RevokeDLCCatalogAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RevokeDLCCatalogAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackDataEngineImageRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 检查是否能回滚的接口返回的FromRecordId参数
	FromRecordId *string `json:"FromRecordId,omitnil,omitempty" name:"FromRecordId"`

	// 检查是否能回滚的接口返回的ToRecordId参数
	ToRecordId *string `json:"ToRecordId,omitnil,omitempty" name:"ToRecordId"`
}

type RollbackDataEngineImageRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 检查是否能回滚的接口返回的FromRecordId参数
	FromRecordId *string `json:"FromRecordId,omitnil,omitempty" name:"FromRecordId"`

	// 检查是否能回滚的接口返回的ToRecordId参数
	ToRecordId *string `json:"ToRecordId,omitnil,omitempty" name:"ToRecordId"`
}

func (r *RollbackDataEngineImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackDataEngineImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "FromRecordId")
	delete(f, "ToRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackDataEngineImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackDataEngineImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackDataEngineImageResponse struct {
	*tchttp.BaseResponse
	Response *RollbackDataEngineImageResponseParams `json:"Response"`
}

func (r *RollbackDataEngineImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackDataEngineImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SQLTask struct {
	// base64加密后的SQL语句
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`

	// 任务的配置信息
	Config []*KVPair `json:"Config,omitnil,omitempty" name:"Config"`
}

type ScheduleElasticityConf struct {
	// 是否开启弹性伸缩：true/false
	ScheduledElasticityEnabled *bool `json:"ScheduledElasticityEnabled,omitnil,omitempty" name:"ScheduledElasticityEnabled"`

	// 调度类型：ONCE（一次性调度），DAILY（每日调度），WEEKLY（每周调度），MONTHLY（每月调度）
	ScheduleType *string `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// 调度日期：WEEKLY传：1~7； MONTHLY传:1~31；其它类型不传
	ScheduleDays []*int64 `json:"ScheduleDays,omitnil,omitempty" name:"ScheduleDays"`

	// 调度时区
	TimeZone *string `json:"TimeZone,omitnil,omitempty" name:"TimeZone"`

	// 弹性伸缩计划
	ElasticPlans []*ElasticPlan `json:"ElasticPlans,omitnil,omitempty" name:"ElasticPlans"`
}

type Script struct {
	// 脚本Id，长度36字节。
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 脚本名称，长度0-25。
	ScriptName *string `json:"ScriptName,omitnil,omitempty" name:"ScriptName"`

	// 脚本描述，长度0-50。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptDesc *string `json:"ScriptDesc,omitnil,omitempty" name:"ScriptDesc"`

	// 默认关联数据库。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// SQL描述，长度0-10000。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SQLStatement *string `json:"SQLStatement,omitnil,omitempty" name:"SQLStatement"`

	// 更新时间戳， 单位：ms。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SessionResourceTemplate struct {
	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	ExecutorNums *uint64 `json:"ExecutorNums,omitnil,omitempty" name:"ExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 运行时参数
	RunningTimeParameters []*DataEngineConfigPair `json:"RunningTimeParameters,omitnil,omitempty" name:"RunningTimeParameters"`
}

type SmartOptimizerChangeTablePolicy struct {
	// change表的数据保存时间，单位为天
	DataRetentionTime *int64 `json:"DataRetentionTime,omitnil,omitempty" name:"DataRetentionTime"`
}

type SmartOptimizerIndexPolicy struct {
	// 开启索引
	IndexEnable *string `json:"IndexEnable,omitnil,omitempty" name:"IndexEnable"`
}

type SmartOptimizerLifecyclePolicy struct {
	// 生命周期启用
	LifecycleEnable *string `json:"LifecycleEnable,omitnil,omitempty" name:"LifecycleEnable"`

	// 过期时间
	Expiration *int64 `json:"Expiration,omitnil,omitempty" name:"Expiration"`

	// 是否删表
	DropTable *bool `json:"DropTable,omitnil,omitempty" name:"DropTable"`

	// 过期字段
	ExpiredField *string `json:"ExpiredField,omitnil,omitempty" name:"ExpiredField"`

	// 过期字段格式
	ExpiredFieldFormat *string `json:"ExpiredFieldFormat,omitnil,omitempty" name:"ExpiredFieldFormat"`
}

type SmartOptimizerPolicy struct {
	// 是否继承
	Inherit *string `json:"Inherit,omitnil,omitempty" name:"Inherit"`

	// ResourceInfo
	// 注意：此字段可能返回 null，表示取不到有效值。
	Resources []*ResourceInfo `json:"Resources,omitnil,omitempty" name:"Resources"`

	// SmartOptimizerWrittenPolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	Written *SmartOptimizerWrittenPolicy `json:"Written,omitnil,omitempty" name:"Written"`

	// SmartOptimizerLifecyclePolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lifecycle *SmartOptimizerLifecyclePolicy `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`

	// SmartOptimizerIndexPolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *SmartOptimizerIndexPolicy `json:"Index,omitnil,omitempty" name:"Index"`

	// SmartOptimizerChangeTablePolicy
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChangeTable *SmartOptimizerChangeTablePolicy `json:"ChangeTable,omitnil,omitempty" name:"ChangeTable"`
}

type SmartOptimizerWrittenPolicy struct {
	// none/enable/disable/default
	WrittenEnable *string `json:"WrittenEnable,omitnil,omitempty" name:"WrittenEnable"`
}

type SmartPolicy struct {
	// 基础信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseInfo *SmartPolicyBaseInfo `json:"BaseInfo,omitnil,omitempty" name:"BaseInfo"`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *SmartOptimizerPolicy `json:"Policy,omitnil,omitempty" name:"Policy"`
}

type SmartPolicyBaseInfo struct {
	// 用户uin
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// Catalog/Database/Table
	PolicyType *string `json:"PolicyType,omitnil,omitempty" name:"PolicyType"`

	// Catalog名称
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 表名称
	Table *string `json:"Table,omitnil,omitempty" name:"Table"`

	// 用户appid
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type Sort struct {
	// 排序字段
	Field *string `json:"Field,omitnil,omitempty" name:"Field"`

	// 是否按照ASC排序，否则DESC排序
	Asc *bool `json:"Asc,omitnil,omitempty" name:"Asc"`
}

type SparkJobInfo struct {
	// spark作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// spark作业名
	JobName *string `json:"JobName,omitnil,omitempty" name:"JobName"`

	// spark作业类型，可去1或者2，1表示batch作业， 2表示streaming作业
	JobType *int64 `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 引擎名
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// 该字段已下线，请使用字段Datasource
	Eni *string `json:"Eni,omitnil,omitempty" name:"Eni"`

	// 程序包是否本地上传，cos或者lakefs
	IsLocal *string `json:"IsLocal,omitnil,omitempty" name:"IsLocal"`

	// 程序包路径
	JobFile *string `json:"JobFile,omitnil,omitempty" name:"JobFile"`

	// 角色ID
	RoleArn *int64 `json:"RoleArn,omitnil,omitempty" name:"RoleArn"`

	// spark作业运行主类
	MainClass *string `json:"MainClass,omitnil,omitempty" name:"MainClass"`

	// 命令行参数，spark作业命令行参数，空格分隔
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// spark原生配置，换行符分隔
	JobConf *string `json:"JobConf,omitnil,omitempty" name:"JobConf"`

	// 依赖jars是否本地上传，cos或者lakefs
	IsLocalJars *string `json:"IsLocalJars,omitnil,omitempty" name:"IsLocalJars"`

	// spark作业依赖jars，逗号分隔
	JobJars *string `json:"JobJars,omitnil,omitempty" name:"JobJars"`

	// 依赖文件是否本地上传，cos或者lakefs
	IsLocalFiles *string `json:"IsLocalFiles,omitnil,omitempty" name:"IsLocalFiles"`

	// spark作业依赖文件，逗号分隔
	JobFiles *string `json:"JobFiles,omitnil,omitempty" name:"JobFiles"`

	// spark作业driver资源大小
	JobDriverSize *string `json:"JobDriverSize,omitnil,omitempty" name:"JobDriverSize"`

	// spark作业executor资源大小
	JobExecutorSize *string `json:"JobExecutorSize,omitnil,omitempty" name:"JobExecutorSize"`

	// spark作业executor个数
	JobExecutorNums *int64 `json:"JobExecutorNums,omitnil,omitempty" name:"JobExecutorNums"`

	// spark流任务最大重试次数
	JobMaxAttempts *int64 `json:"JobMaxAttempts,omitnil,omitempty" name:"JobMaxAttempts"`

	// spark作业创建者
	JobCreator *string `json:"JobCreator,omitnil,omitempty" name:"JobCreator"`

	// spark作业创建时间
	JobCreateTime *int64 `json:"JobCreateTime,omitnil,omitempty" name:"JobCreateTime"`

	// spark作业更新时间
	JobUpdateTime *uint64 `json:"JobUpdateTime,omitnil,omitempty" name:"JobUpdateTime"`

	// spark作业最近任务ID
	CurrentTaskId *string `json:"CurrentTaskId,omitnil,omitempty" name:"CurrentTaskId"`

	// spark作业最近运行状态，初始化：0，运行中：1，成功：2，数据写入中： 3， 排队中： 4， 失败： -1， 已删除： -3，已过期： -5
	JobStatus *int64 `json:"JobStatus,omitnil,omitempty" name:"JobStatus"`

	// spark流作业统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamingStat *StreamingStatistics `json:"StreamingStat,omitnil,omitempty" name:"StreamingStat"`

	// 数据源名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSource *string `json:"DataSource,omitnil,omitempty" name:"DataSource"`

	// pyspark：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLocalPythonFiles *string `json:"IsLocalPythonFiles,omitnil,omitempty" name:"IsLocalPythonFiles"`

	// 注：该返回值已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppPythonFiles *string `json:"AppPythonFiles,omitnil,omitempty" name:"AppPythonFiles"`

	// archives：依赖上传方式，1、cos；2、lakefs（控制台使用，该方式不支持直接接口调用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLocalArchives *string `json:"IsLocalArchives,omitnil,omitempty" name:"IsLocalArchives"`

	// archives：依赖资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobArchives *string `json:"JobArchives,omitnil,omitempty" name:"JobArchives"`

	// Spark Image 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImage *string `json:"SparkImage,omitnil,omitempty" name:"SparkImage"`

	// pyspark：python依赖, 除py文件外，还支持zip/egg等归档格式，多文件以逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobPythonFiles *string `json:"JobPythonFiles,omitnil,omitempty" name:"JobPythonFiles"`

	// 当前job正在运行或准备运行的任务个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNum *int64 `json:"TaskNum,omitnil,omitempty" name:"TaskNum"`

	// 引擎状态：-100（默认：未知状态），-2~11：引擎正常状态；
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineStatus *int64 `json:"DataEngineStatus,omitnil,omitempty" name:"DataEngineStatus"`

	// 指定的Executor数量（最大值），默认为1，当开启动态分配有效，若未开启，则该值等于JobExecutorNums
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobExecutorMaxNumbers *int64 `json:"JobExecutorMaxNumbers,omitnil,omitempty" name:"JobExecutorMaxNumbers"`

	// 镜像版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkImageVersion *string `json:"SparkImageVersion,omitnil,omitempty" name:"SparkImageVersion"`

	// 查询脚本关联id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// spark_emr_livy
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineClusterType *string `json:"DataEngineClusterType,omitnil,omitempty" name:"DataEngineClusterType"`

	// Spark 3.2-EMR
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineImageVersion *string `json:"DataEngineImageVersion,omitnil,omitempty" name:"DataEngineImageVersion"`

	// 任务资源配置是否继承集群模板，0（默认）不继承，1：继承
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsInherit *uint64 `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`

	// 是否使用session脚本的sql运行任务：false：否，true：是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSessionStarted *bool `json:"IsSessionStarted,omitnil,omitempty" name:"IsSessionStarted"`

	// 引擎详细类型：SparkSQL、PrestoSQL、SparkBatch、StandardSpark、StandardPresto
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil,omitempty" name:"EngineTypeDetail"`
}

type SparkMonitorMetrics struct {
	// shuffle写溢出到COS数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShuffleWriteBytesCos *int64 `json:"ShuffleWriteBytesCos,omitnil,omitempty" name:"ShuffleWriteBytesCos"`

	// shuffle写数据量，单位：byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShuffleWriteBytesTotal *int64 `json:"ShuffleWriteBytesTotal,omitnil,omitempty" name:"ShuffleWriteBytesTotal"`
}

type SparkSessionBatchLog struct {
	// 日志步骤：BEG/CS/DS/DSS/DSF/FINF/RTO/CANCEL/CT/DT/DTS/DTF/FINT/EXCE
	Step *string `json:"Step,omitnil,omitempty" name:"Step"`

	// 时间
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志提示
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 日志操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operate []*SparkSessionBatchLogOperate `json:"Operate,omitnil,omitempty" name:"Operate"`
}

type SparkSessionBatchLogOperate struct {
	// 操作提示
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 操作类型：COPY、LOG、UI、RESULT、List、TAB
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 补充信息：如：taskid、sessionid、sparkui等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supplement []*KVPair `json:"Supplement,omitnil,omitempty" name:"Supplement"`
}

type SparkSessionInfo struct {
	// spark session id
	SparkSessionId *string `json:"SparkSessionId,omitnil,omitempty" name:"SparkSessionId"`

	// spark session名称
	SparkSessionName *string `json:"SparkSessionName,omitnil,omitempty" name:"SparkSessionName"`

	// 资源组id
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// engine session id
	EngineSessionId *string `json:"EngineSessionId,omitnil,omitempty" name:"EngineSessionId"`

	// engine session   
	// name
	EngineSessionName *string `json:"EngineSessionName,omitnil,omitempty" name:"EngineSessionName"`

	// 自动销毁时间
	IdleTimeoutMin *int64 `json:"IdleTimeoutMin,omitnil,omitempty" name:"IdleTimeoutMin"`

	// driver规格
	DriverSpec *string `json:"DriverSpec,omitnil,omitempty" name:"DriverSpec"`

	// executor规格
	ExecutorSpec *string `json:"ExecutorSpec,omitnil,omitempty" name:"ExecutorSpec"`

	// executor最小数量
	ExecutorNumMin *int64 `json:"ExecutorNumMin,omitnil,omitempty" name:"ExecutorNumMin"`

	// executor最大数量
	ExecutorNumMax *int64 `json:"ExecutorNumMax,omitnil,omitempty" name:"ExecutorNumMax"`

	// 总规格最小
	TotalSpecMin *int64 `json:"TotalSpecMin,omitnil,omitempty" name:"TotalSpecMin"`

	// 总规格最大
	TotalSpecMax *int64 `json:"TotalSpecMax,omitnil,omitempty" name:"TotalSpecMax"`
}

type SpecInfo struct {
	// 规格名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 当前规格的cu数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cu *int64 `json:"Cu,omitnil,omitempty" name:"Cu"`

	// 当前规格的cpu数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 当前规格的内存数，单位G
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`
}

type StandardEngineResourceGroupConfigInfo struct {
	// 引擎资源组 ID
	ResourceGroupId *string `json:"ResourceGroupId,omitnil,omitempty" name:"ResourceGroupId"`

	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 资源组静态参数，需要重启资源组生效
	StaticConfigPairs []*EngineResourceGroupConfigPair `json:"StaticConfigPairs,omitnil,omitempty" name:"StaticConfigPairs"`

	// 资源组动态参数，下一个任务生效。
	DynamicConfigPairs []*EngineResourceGroupConfigPair `json:"DynamicConfigPairs,omitnil,omitempty" name:"DynamicConfigPairs"`

	// 创建时间
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type StandardEngineResourceGroupInfo struct {
	// 标准引擎资源组ID
	EngineResourceGroupId *string `json:"EngineResourceGroupId,omitnil,omitempty" name:"EngineResourceGroupId"`

	// 标准引擎资源组名称，支持1-50个英文、汉字、数字、连接线-或下划线_
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 资源组 状态，-1--删除、0--启动中、2--运行、3--暂停、4--暂停中、7--切换引擎中、8--配置修改中。9--资源组重启中，10--因为变配导致资源组启动、11--因为隔离导致资源组挂起、12- 资源配置下发中、 13-接入点隔离导致资源组挂起中
	ResourceGroupState *int64 `json:"ResourceGroupState,omitnil,omitempty" name:"ResourceGroupState"`

	// 自动启动，（任务提交自动拉起资源组）0-自动启动，1-不自动启动
	AutoLaunch *int64 `json:"AutoLaunch,omitnil,omitempty" name:"AutoLaunch"`

	// 自动挂起资源组。0-自动挂起，1-不自动挂起
	AutoPause *int64 `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 自动挂起时间，单位分钟，取值范围在1-999（在无任务AutoPauseTime后，资源组自动挂起）
	AutoPauseTime *int64 `json:"AutoPauseTime,omitnil,omitempty" name:"AutoPauseTime"`

	// driver的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	DriverCuSpec *string `json:"DriverCuSpec,omitnil,omitempty" name:"DriverCuSpec"`

	// executor的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	ExecutorCuSpec *string `json:"ExecutorCuSpec,omitnil,omitempty" name:"ExecutorCuSpec"`

	// 任务并发数
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// executor最小数量，
	MinExecutorNums *int64 `json:"MinExecutorNums,omitnil,omitempty" name:"MinExecutorNums"`

	// executor最大数量，
	MaxExecutorNums *int64 `json:"MaxExecutorNums,omitnil,omitempty" name:"MaxExecutorNums"`

	// 创建时间戳
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间戳
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否待重启，作为有资源参数，静态参数修改未重启生效的标识；0-- 不需要重启、1--因为资源参数待重启、2--因静态参数重启、3--因资源和静态参数而待重启、4--因网络配置而待重启、5--因网络配置和资源配置而待重启、6--因网络配置和静态参数而待重启、7--因网络配置，资源参数和静态参数而待重启、
	NeedRestart *int64 `json:"NeedRestart,omitnil,omitempty" name:"NeedRestart"`

	// 绑定的引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 绑定的引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 绑定的引擎状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngineState *int64 `json:"DataEngineState,omitnil,omitempty" name:"DataEngineState"`

	// 接入点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPointId *string `json:"AccessPointId,omitnil,omitempty" name:"AccessPointId"`

	// 接入点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPointName *string `json:"AccessPointName,omitnil,omitempty" name:"AccessPointName"`

	// 接入点状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessPointState *int64 `json:"AccessPointState,omitnil,omitempty" name:"AccessPointState"`

	// 资源组类型，console/ default
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupType *string `json:"ResourceGroupType,omitnil,omitempty" name:"ResourceGroupType"`

	// 引擎网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 网络配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkConfigNames []*string `json:"NetworkConfigNames,omitnil,omitempty" name:"NetworkConfigNames"`

	// AI类型资源组的框架类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameType *string `json:"FrameType,omitnil,omitempty" name:"FrameType"`

	// AI类型资源组的镜像类型，内置：bulit-in，自定义：custom
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// AI资源组的可用资源上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 是否是默认资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 资源组场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroupScene *string `json:"ResourceGroupScene,omitnil,omitempty" name:"ResourceGroupScene"`

	// python类型资源组python单机节点资源上限，该值要小于资源组的资源上限.small:1cu medium:2cu large:4cu xlarge:8cu 4xlarge:16cu 8xlarge:32cu 16xlarge:64cu，如果是高内存型资源，在类型前面加上m.
	// 注意：此字段可能返回 null，表示取不到有效值。
	PythonCuSpec *string `json:"PythonCuSpec,omitnil,omitempty" name:"PythonCuSpec"`

	// Spark类型资源组资源配置模式，fast：快速模式，custom：自定义模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkSpecMode *string `json:"SparkSpecMode,omitnil,omitempty" name:"SparkSpecMode"`

	// Spark类型资源组资源上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkSize *int64 `json:"SparkSize,omitnil,omitempty" name:"SparkSize"`

	// Spark类型资源组资源最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkMinSize *int64 `json:"SparkMinSize,omitnil,omitempty" name:"SparkMinSize"`

	// 自定义镜像容器镜像服务domain 名称
	PublicDomain *string `json:"PublicDomain,omitnil,omitempty" name:"PublicDomain"`

	// 自定义镜像容器镜像服务tcr实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 容器镜像服务tcr所在地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// 资源组启动耗时
	LaunchTime *string `json:"LaunchTime,omitnil,omitempty" name:"LaunchTime"`
}

type StatementInformation struct {
	// SQL任务唯一ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// SQL内容
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`
}

type StatementOutput struct {
	// 执行总数
	ExecutionCount *int64 `json:"ExecutionCount,omitnil,omitempty" name:"ExecutionCount"`

	// Statement数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*KVPair `json:"Data,omitnil,omitempty" name:"Data"`

	// Statement状态:ok,error
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 错误名称
	ErrorName *string `json:"ErrorName,omitnil,omitempty" name:"ErrorName"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorValue *string `json:"ErrorValue,omitnil,omitempty" name:"ErrorValue"`

	// 错误堆栈信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage []*string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// SQL类型任务结果返回
	SQLResult *string `json:"SQLResult,omitnil,omitempty" name:"SQLResult"`
}

type StreamingStatistics struct {
	// 任务开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 数据接收器数
	Receivers *int64 `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 运行中的接收器数
	NumActiveReceivers *int64 `json:"NumActiveReceivers,omitnil,omitempty" name:"NumActiveReceivers"`

	// 不活跃的接收器数
	NumInactiveReceivers *int64 `json:"NumInactiveReceivers,omitnil,omitempty" name:"NumInactiveReceivers"`

	// 运行中的批数
	NumActiveBatches *int64 `json:"NumActiveBatches,omitnil,omitempty" name:"NumActiveBatches"`

	// 待处理的批数
	NumRetainedCompletedBatches *int64 `json:"NumRetainedCompletedBatches,omitnil,omitempty" name:"NumRetainedCompletedBatches"`

	// 已完成的批数
	NumTotalCompletedBatches *int64 `json:"NumTotalCompletedBatches,omitnil,omitempty" name:"NumTotalCompletedBatches"`

	// 平均输入速率
	AverageInputRate *float64 `json:"AverageInputRate,omitnil,omitempty" name:"AverageInputRate"`

	// 平均等待时长
	AverageSchedulingDelay *float64 `json:"AverageSchedulingDelay,omitnil,omitempty" name:"AverageSchedulingDelay"`

	// 平均处理时长
	AverageProcessingTime *float64 `json:"AverageProcessingTime,omitnil,omitempty" name:"AverageProcessingTime"`

	// 平均延时
	AverageTotalDelay *float64 `json:"AverageTotalDelay,omitnil,omitempty" name:"AverageTotalDelay"`
}

// Predefined struct for user
type SuspendResumeDataEngineRequestParams struct {
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 操作类型 suspend/resume
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`
}

type SuspendResumeDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 虚拟集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 操作类型 suspend/resume
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`
}

func (r *SuspendResumeDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendResumeDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "Operate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SuspendResumeDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SuspendResumeDataEngineResponseParams struct {
	// 虚拟集群详细信息
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SuspendResumeDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *SuspendResumeDataEngineResponseParams `json:"Response"`
}

func (r *SuspendResumeDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendResumeDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineImageRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 新镜像版本ID
	NewImageVersionId *string `json:"NewImageVersionId,omitnil,omitempty" name:"NewImageVersionId"`
}

type SwitchDataEngineImageRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 新镜像版本ID
	NewImageVersionId *string `json:"NewImageVersionId,omitnil,omitempty" name:"NewImageVersionId"`
}

func (r *SwitchDataEngineImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "NewImageVersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDataEngineImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchDataEngineImageResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDataEngineImageResponseParams `json:"Response"`
}

func (r *SwitchDataEngineImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineRequestParams struct {
	// 主集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 是否开启备集群
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitnil,omitempty" name:"StartStandbyCluster"`
}

type SwitchDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 主集群名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 是否开启备集群
	StartStandbyCluster *bool `json:"StartStandbyCluster,omitnil,omitempty" name:"StartStandbyCluster"`
}

func (r *SwitchDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineName")
	delete(f, "StartStandbyCluster")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SwitchDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SwitchDataEngineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SwitchDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *SwitchDataEngineResponseParams `json:"Response"`
}

func (r *SwitchDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SwitchDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TCHouseD struct {
	// 数据源实例的唯一ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据源名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 数据源的JDBC
	JdbcUrl *string `json:"JdbcUrl,omitnil,omitempty" name:"JdbcUrl"`

	// 用于访问数据源的用户
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据源访问密码，需要base64编码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 数据源的VPC和子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *DatasourceConnectionLocation `json:"Location,omitnil,omitempty" name:"Location"`

	// 默认数据库名
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 访问信息
	AccessInfo *string `json:"AccessInfo,omitnil,omitempty" name:"AccessInfo"`
}

type TColumn struct {
	// 字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 字段默认值
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// 字段是否是非空
	NotNull *bool `json:"NotNull,omitnil,omitempty" name:"NotNull"`

	// 表示整个 numeric 的长度,取值1-38
	Precision *int64 `json:"Precision,omitnil,omitempty" name:"Precision"`

	// 表示小数部分的长度
	// Scale小于Precision
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// 字段位置，小的在前
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 是否为分区字段
	IsPartition *bool `json:"IsPartition,omitnil,omitempty" name:"IsPartition"`
}

type TPartition struct {
	// 字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 分区类型（已废弃）
	PartitionType *string `json:"PartitionType,omitnil,omitempty" name:"PartitionType"`

	// 分区格式（已废弃）
	PartitionFormat *string `json:"PartitionFormat,omitnil,omitempty" name:"PartitionFormat"`

	// 分区分隔数（已废弃）
	PartitionDot *int64 `json:"PartitionDot,omitnil,omitempty" name:"PartitionDot"`

	// 分区转换策略
	Transform *string `json:"Transform,omitnil,omitempty" name:"Transform"`

	// 策略参数
	TransformArgs []*string `json:"TransformArgs,omitnil,omitempty" name:"TransformArgs"`
}

type TableBaseInfo struct {
	// 该数据表所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据表名字
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 该数据表所属数据源名字
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 该数据表备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// 具体类型，表or视图
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据格式类型，hive，iceberg等
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`

	// 建表用户昵称
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 建表用户ID
	UserSubUin *string `json:"UserSubUin,omitnil,omitempty" name:"UserSubUin"`

	// 数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: GovernPolicy is deprecated.
	GovernPolicy *DataGovernPolicy `json:"GovernPolicy,omitnil,omitempty" name:"GovernPolicy"`

	// 库数据治理是否关闭，关闭：true，开启：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DbGovernPolicyIsDisable is deprecated.
	DbGovernPolicyIsDisable *string `json:"DbGovernPolicyIsDisable,omitnil,omitempty" name:"DbGovernPolicyIsDisable"`

	// 智能数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	SmartPolicy *SmartPolicy `json:"SmartPolicy,omitnil,omitempty" name:"SmartPolicy"`

	// T-ICEBERG表的主键
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrimaryKeys []*string `json:"PrimaryKeys,omitnil,omitempty" name:"PrimaryKeys"`
}

type TableInfo struct {
	// 数据表配置信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 数据表格式。每次入参可选如下其一的KV结构，[TextFile，CSV，Json, Parquet, ORC, AVRD]。
	DataFormat *DataFormat `json:"DataFormat,omitnil,omitempty" name:"DataFormat"`

	// 数据表列信息。
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 数据表分块信息。
	Partitions []*Partition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 数据存储路径。当前仅支持cos路径，格式如下：cosn://bucket-name/filepath。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`
}

type TableResponseInfo struct {
	// 数据表基本信息。
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 数据表列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 数据表分块信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions []*Partition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 数据存储路径。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 数据表属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 数据表更新时间, 单位: ms。
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 数据表创建时间,单位: ms。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据格式。
	InputFormat *string `json:"InputFormat,omitnil,omitempty" name:"InputFormat"`

	// 数据表存储大小（单位：Byte）
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 数据表行数
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// xxxx
	MapMaterializedViewName *string `json:"MapMaterializedViewName,omitnil,omitempty" name:"MapMaterializedViewName"`

	// 访问热点
	HeatValue *int64 `json:"HeatValue,omitnil,omitempty" name:"HeatValue"`

	// InputFormat的缩写
	InputFormatShort *string `json:"InputFormatShort,omitnil,omitempty" name:"InputFormatShort"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type Task struct {
	// SQL查询任务
	SQLTask *SQLTask `json:"SQLTask,omitnil,omitempty" name:"SQLTask"`

	// Spark SQL查询任务
	SparkSQLTask *SQLTask `json:"SparkSQLTask,omitnil,omitempty" name:"SparkSQLTask"`
}

type TaskMonitorInfo struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 引擎名称
	HouseName *string `json:"HouseName,omitnil,omitempty" name:"HouseName"`

	// sql语句
	QuerySQL *string `json:"QuerySQL,omitnil,omitempty" name:"QuerySQL"`

	// 任务时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 执行时间
	UsedTime *string `json:"UsedTime,omitnil,omitempty" name:"UsedTime"`

	// 数据扫描量
	DataAmount *string `json:"DataAmount,omitnil,omitempty" name:"DataAmount"`

	// 指标信息
	QueryStats *string `json:"QueryStats,omitnil,omitempty" name:"QueryStats"`
}

type TaskResponseInfo struct {
	// 任务所属Database的名称。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 任务数据量。
	DataAmount *int64 `json:"DataAmount,omitnil,omitempty" name:"DataAmount"`

	// 任务Id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 计算耗时，单位： ms
	UsedTime *int64 `json:"UsedTime,omitnil,omitempty" name:"UsedTime"`

	// 任务输出路径。
	OutputPath *string `json:"OutputPath,omitnil,omitempty" name:"OutputPath"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务状态：0 初始化， 1 执行中， 2 执行成功，3 数据写入中，4 排队中。-1 执行失败，-3 已取消。
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 任务SQL类型，DDL|DML等
	SQLType *string `json:"SQLType,omitnil,omitempty" name:"SQLType"`

	// 任务SQL语句
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`

	// 结果是否过期。
	ResultExpired *bool `json:"ResultExpired,omitnil,omitempty" name:"ResultExpired"`

	// 数据影响统计信息。
	RowAffectInfo *string `json:"RowAffectInfo,omitnil,omitempty" name:"RowAffectInfo"`

	// 任务结果数据表。
	DataSet *string `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 失败信息, 例如：errorMessage。该字段已废弃。
	Error *string `json:"Error,omitnil,omitempty" name:"Error"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitnil,omitempty" name:"Percentage"`

	// 任务执行输出信息。
	OutputMessage *string `json:"OutputMessage,omitnil,omitempty" name:"OutputMessage"`

	// 执行SQL的引擎类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务进度明细
	ProgressDetail *string `json:"ProgressDetail,omitnil,omitempty" name:"ProgressDetail"`

	// 任务结束时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 计算资源id
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 执行sql的子uin
	OperateUin *string `json:"OperateUin,omitnil,omitempty" name:"OperateUin"`

	// 计算资源名字
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 导入类型是本地导入还是cos
	InputType *string `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 导入配置
	InputConf *string `json:"InputConf,omitnil,omitempty" name:"InputConf"`

	// 数据条数
	DataNumber *int64 `json:"DataNumber,omitnil,omitempty" name:"DataNumber"`

	// 查询数据能不能下载
	CanDownload *bool `json:"CanDownload,omitnil,omitempty" name:"CanDownload"`

	// 用户别名
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// spark应用作业名
	SparkJobName *string `json:"SparkJobName,omitnil,omitempty" name:"SparkJobName"`

	// spark应用作业Id
	SparkJobId *string `json:"SparkJobId,omitnil,omitempty" name:"SparkJobId"`

	// spark应用入口jar文件
	SparkJobFile *string `json:"SparkJobFile,omitnil,omitempty" name:"SparkJobFile"`

	// spark ui url
	UiUrl *string `json:"UiUrl,omitnil,omitempty" name:"UiUrl"`

	// 任务耗时，单位： ms
	TotalTime *int64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// spark app job执行task的程序入口参数
	CmdArgs *string `json:"CmdArgs,omitnil,omitempty" name:"CmdArgs"`

	// 集群镜像大版本名称
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// driver规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// executor规格：small,medium,large,xlarge；内存型(引擎类型)：m.small,m.medium,m.large,m.xlarge
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// 指定executor数量，最小值为1，最大值小于集群规格
	ExecutorNums *uint64 `json:"ExecutorNums,omitnil,omitempty" name:"ExecutorNums"`

	// 指定executor max数量（动态配置场景下），最小值为1，最大值小于集群规格（当ExecutorMaxNumbers小于ExecutorNums时，改值设定为ExecutorNums）
	ExecutorMaxNumbers *uint64 `json:"ExecutorMaxNumbers,omitnil,omitempty" name:"ExecutorMaxNumbers"`

	// 任务公共指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommonMetrics *CommonMetrics `json:"CommonMetrics,omitnil,omitempty" name:"CommonMetrics"`

	// spark任务指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	SparkMonitorMetrics *SparkMonitorMetrics `json:"SparkMonitorMetrics,omitnil,omitempty" name:"SparkMonitorMetrics"`

	// presto任务指标数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrestoMonitorMetrics *PrestoMonitorMetrics `json:"PrestoMonitorMetrics,omitnil,omitempty" name:"PrestoMonitorMetrics"`

	// 结果文件格式：默认为csv
	ResultFormat *string `json:"ResultFormat,omitnil,omitempty" name:"ResultFormat"`

	// 引擎类型，SparkSQL：SparkSQL 引擎；SparkBatch：Spark作业引擎；PrestoSQL：Presto引擎
	EngineTypeDetail *string `json:"EngineTypeDetail,omitnil,omitempty" name:"EngineTypeDetail"`

	// spark引擎资源组名称
	ResourceGroupName *string `json:"ResourceGroupName,omitnil,omitempty" name:"ResourceGroupName"`

	// 任务执行耗时
	JobTimeSum *int64 `json:"JobTimeSum,omitnil,omitempty" name:"JobTimeSum"`
}

type TaskResultInfo struct {
	// 任务唯一ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据源名称，当前任务执行时候选中的默认数据源
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 数据库名称，当前任务执行时候选中的默认数据库
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 当前执行的SQL，一个任务包含一个SQL
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`

	// 执行任务的类型，现在分为DDL、DML、DQL
	SQLType *string `json:"SQLType,omitnil,omitempty" name:"SQLType"`

	// 任务当前的状态，0：初始化 1：任务运行中 2：任务执行成功  3：数据写入中 4：排队中 -1：任务执行失败 -3：用户手动终止 。只有任务执行成功的情况下，才会返回任务执行的结果
	State *int64 `json:"State,omitnil,omitempty" name:"State"`

	// 扫描的数据量，单位byte
	DataAmount *int64 `json:"DataAmount,omitnil,omitempty" name:"DataAmount"`

	// 计算耗时，单位： ms
	UsedTime *int64 `json:"UsedTime,omitnil,omitempty" name:"UsedTime"`

	// 任务结果输出的COS桶地址
	OutputPath *string `json:"OutputPath,omitnil,omitempty" name:"OutputPath"`

	// 任务创建时间，时间戳
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务执行信息，成功时返回success，失败时返回失败原因
	OutputMessage *string `json:"OutputMessage,omitnil,omitempty" name:"OutputMessage"`

	// 被影响的行数
	RowAffectInfo *string `json:"RowAffectInfo,omitnil,omitempty" name:"RowAffectInfo"`

	// 结果的schema信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultSchema []*Column `json:"ResultSchema,omitnil,omitempty" name:"ResultSchema"`

	// 结果信息，反转义后，外层数组的每个元素为一行数据
	ResultSet *string `json:"ResultSet,omitnil,omitempty" name:"ResultSet"`

	// 分页信息，如果没有更多结果数据，nextToken为空
	NextToken *string `json:"NextToken,omitnil,omitempty" name:"NextToken"`

	// 任务执行进度num/100(%)
	Percentage *int64 `json:"Percentage,omitnil,omitempty" name:"Percentage"`

	// 任务进度明细
	ProgressDetail *string `json:"ProgressDetail,omitnil,omitempty" name:"ProgressDetail"`

	// 控制台展示格式。table：表格展示 text：文本展示
	DisplayFormat *string `json:"DisplayFormat,omitnil,omitempty" name:"DisplayFormat"`

	// 任务耗时，单位： ms
	TotalTime *int64 `json:"TotalTime,omitnil,omitempty" name:"TotalTime"`

	// 获取结果消耗的时间
	QueryResultTime *float64 `json:"QueryResultTime,omitnil,omitempty" name:"QueryResultTime"`

	// base64 编码结果集
	ResultSetEncode *string `json:"ResultSetEncode,omitnil,omitempty" name:"ResultSetEncode"`
}

type TasksInfo struct {
	// 任务类型，SQLTask：SQL查询任务。SparkSQLTask：Spark SQL查询任务
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 容错策略。Proceed：前面任务出错/取消后继续执行后面的任务。Terminate：前面的任务出错/取消之后终止后面任务的执行，后面的任务全部标记为已取消。
	FailureTolerance *string `json:"FailureTolerance,omitnil,omitempty" name:"FailureTolerance"`

	// base64加密后的SQL语句，用";"号分隔每个SQL语句，一次最多提交50个任务。严格按照前后顺序执行
	SQL *string `json:"SQL,omitnil,omitempty" name:"SQL"`

	// 任务的配置信息，当前仅支持SparkSQLTask任务。
	Config []*KVPair `json:"Config,omitnil,omitempty" name:"Config"`

	// 任务的用户自定义参数信息
	Params []*KVPair `json:"Params,omitnil,omitempty" name:"Params"`
}

type TasksOverview struct {
	// 正在排队的任务个数
	TaskQueuedCount *int64 `json:"TaskQueuedCount,omitnil,omitempty" name:"TaskQueuedCount"`

	// 初始化的任务个数
	TaskInitCount *int64 `json:"TaskInitCount,omitnil,omitempty" name:"TaskInitCount"`

	// 正在执行的任务个数
	TaskRunningCount *int64 `json:"TaskRunningCount,omitnil,omitempty" name:"TaskRunningCount"`

	// 当前时间范围的总任务个数
	TotalTaskCount *int64 `json:"TotalTaskCount,omitnil,omitempty" name:"TotalTaskCount"`
}

type TccHive struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 终端节点服务ID
	EndpointServiceId *string `json:"EndpointServiceId,omitnil,omitempty" name:"EndpointServiceId"`

	// thrift连接地址
	MetaStoreUrl *string `json:"MetaStoreUrl,omitnil,omitempty" name:"MetaStoreUrl"`

	// hive版本
	HiveVersion *string `json:"HiveVersion,omitnil,omitempty" name:"HiveVersion"`

	// 网络信息
	TccConnection *NetWork `json:"TccConnection,omitnil,omitempty" name:"TccConnection"`

	// Hms终端节点服务ID
	HmsEndpointServiceId *string `json:"HmsEndpointServiceId,omitnil,omitempty" name:"HmsEndpointServiceId"`
}

type TextFile struct {
	// 文本类型，本参数取值为TextFile。
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 处理文本用的正则表达式。
	Regex *string `json:"Regex,omitnil,omitempty" name:"Regex"`
}

type UDFPolicyInfo struct {
	// 权限类型
	// 示例：select，alter，drop
	Accesses []*string `json:"Accesses,omitnil,omitempty" name:"Accesses"`

	// 拥有权限的用户信息
	Users []*string `json:"Users,omitnil,omitempty" name:"Users"`

	// 拥有权限的工作组的信息
	Groups []*string `json:"Groups,omitnil,omitempty" name:"Groups"`
}

// Predefined struct for user
type UnbindWorkGroupsFromUserRequestParams struct {
	// 解绑的工作组Id和用户Id的关联关系
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

type UnbindWorkGroupsFromUserRequest struct {
	*tchttp.BaseRequest
	
	// 解绑的工作组Id和用户Id的关联关系
	AddInfo *WorkGroupIdSetOfUserId `json:"AddInfo,omitnil,omitempty" name:"AddInfo"`
}

func (r *UnbindWorkGroupsFromUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkGroupsFromUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AddInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindWorkGroupsFromUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindWorkGroupsFromUserResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindWorkGroupsFromUserResponse struct {
	*tchttp.BaseResponse
	Response *UnbindWorkGroupsFromUserResponseParams `json:"Response"`
}

func (r *UnbindWorkGroupsFromUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindWorkGroupsFromUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnboundDatasourceHouseRequestParams struct {
	// 网络配置名称
	NetworkConnectionName *string `json:"NetworkConnectionName,omitnil,omitempty" name:"NetworkConnectionName"`
}

type UnboundDatasourceHouseRequest struct {
	*tchttp.BaseRequest
	
	// 网络配置名称
	NetworkConnectionName *string `json:"NetworkConnectionName,omitnil,omitempty" name:"NetworkConnectionName"`
}

func (r *UnboundDatasourceHouseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnboundDatasourceHouseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnboundDatasourceHouseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnboundDatasourceHouseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnboundDatasourceHouseResponse struct {
	*tchttp.BaseResponse
	Response *UnboundDatasourceHouseResponseParams `json:"Response"`
}

func (r *UnboundDatasourceHouseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnboundDatasourceHouseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockMetaDataRequestParams struct {
	// 锁ID
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

type UnlockMetaDataRequest struct {
	*tchttp.BaseRequest
	
	// 锁ID
	LockId *int64 `json:"LockId,omitnil,omitempty" name:"LockId"`

	// 数据源名称
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`
}

func (r *UnlockMetaDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockMetaDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LockId")
	delete(f, "DatasourceConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlockMetaDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockMetaDataResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnlockMetaDataResponse struct {
	*tchttp.BaseResponse
	Response *UnlockMetaDataResponseParams `json:"Response"`
}

func (r *UnlockMetaDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockMetaDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConfContext struct {
	// 参数类型，可选：StaticConfigType，DynamicConfigType
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 参数的配置数组
	Params []*Param `json:"Params,omitnil,omitempty" name:"Params"`
}

// Predefined struct for user
type UpdateDataEngineConfigRequestParams struct {
	// 引擎ID
	DataEngineIds []*string `json:"DataEngineIds,omitnil,omitempty" name:"DataEngineIds"`

	// 引擎配置命令，支持UpdateSparkSQLLakefsPath（更新原生表配置）、UpdateSparkSQLResultPath（更新结果路径配置）
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil,omitempty" name:"DataEngineConfigCommand"`

	// 是否使用lakefs作为结果存储
	UseLakeFs *bool `json:"UseLakeFs,omitnil,omitempty" name:"UseLakeFs"`

	// 用户自定义结果路径
	CustomResultPath *string `json:"CustomResultPath,omitnil,omitempty" name:"CustomResultPath"`
}

type UpdateDataEngineConfigRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineIds []*string `json:"DataEngineIds,omitnil,omitempty" name:"DataEngineIds"`

	// 引擎配置命令，支持UpdateSparkSQLLakefsPath（更新原生表配置）、UpdateSparkSQLResultPath（更新结果路径配置）
	DataEngineConfigCommand *string `json:"DataEngineConfigCommand,omitnil,omitempty" name:"DataEngineConfigCommand"`

	// 是否使用lakefs作为结果存储
	UseLakeFs *bool `json:"UseLakeFs,omitnil,omitempty" name:"UseLakeFs"`

	// 用户自定义结果路径
	CustomResultPath *string `json:"CustomResultPath,omitnil,omitempty" name:"CustomResultPath"`
}

func (r *UpdateDataEngineConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineIds")
	delete(f, "DataEngineConfigCommand")
	delete(f, "UseLakeFs")
	delete(f, "CustomResultPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataEngineConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataEngineConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDataEngineConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataEngineConfigResponseParams `json:"Response"`
}

func (r *UpdateDataEngineConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataEngineRequestParams struct {
	// 资源大小
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil,omitempty" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil,omitempty" name:"MaxClusters"`

	// 开启自动刷新：true：开启、false（默认）：关闭
	AutoResume *bool `json:"AutoResume,omitnil,omitempty" name:"AutoResume"`

	// 数据引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 相关信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil,omitempty" name:"TolerableQueueTime"`

	// 集群自动挂起时间
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil,omitempty" name:"AutoSuspendTime"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil,omitempty" name:"ElasticLimit"`

	// Spark批作业集群Session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`

	// 引擎资源弹性伸缩策略
	ScheduleElasticityConf *ScheduleElasticityConf `json:"ScheduleElasticityConf,omitnil,omitempty" name:"ScheduleElasticityConf"`
}

type UpdateDataEngineRequest struct {
	*tchttp.BaseRequest
	
	// 资源大小
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 最小资源
	MinClusters *int64 `json:"MinClusters,omitnil,omitempty" name:"MinClusters"`

	// 最大资源
	MaxClusters *int64 `json:"MaxClusters,omitnil,omitempty" name:"MaxClusters"`

	// 开启自动刷新：true：开启、false（默认）：关闭
	AutoResume *bool `json:"AutoResume,omitnil,omitempty" name:"AutoResume"`

	// 数据引擎名称
	DataEngineName *string `json:"DataEngineName,omitnil,omitempty" name:"DataEngineName"`

	// 相关信息
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 是否自定挂起集群：false（默认）：不自动挂起、true：自动挂起
	AutoSuspend *bool `json:"AutoSuspend,omitnil,omitempty" name:"AutoSuspend"`

	// 定时启停集群策略：0（默认）：关闭定时策略、1：开启定时策略（注：定时启停策略与自动挂起策略互斥）
	CrontabResumeSuspend *int64 `json:"CrontabResumeSuspend,omitnil,omitempty" name:"CrontabResumeSuspend"`

	// 定时启停策略，复杂类型：包含启停时间、挂起集群策略
	CrontabResumeSuspendStrategy *CrontabResumeSuspendStrategy `json:"CrontabResumeSuspendStrategy,omitnil,omitempty" name:"CrontabResumeSuspendStrategy"`

	// 单个集群最大并发任务数，默认5
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 可容忍的排队时间，默认0。当任务排队的时间超过可容忍的时间时可能会触发扩容。如果该参数为0，则表示一旦有任务排队就可能立即触发扩容。
	TolerableQueueTime *int64 `json:"TolerableQueueTime,omitnil,omitempty" name:"TolerableQueueTime"`

	// 集群自动挂起时间
	AutoSuspendTime *int64 `json:"AutoSuspendTime,omitnil,omitempty" name:"AutoSuspendTime"`

	// spark jar 包年包月集群是否开启弹性
	ElasticSwitch *bool `json:"ElasticSwitch,omitnil,omitempty" name:"ElasticSwitch"`

	// spark jar 包年包月集群弹性上限
	ElasticLimit *int64 `json:"ElasticLimit,omitnil,omitempty" name:"ElasticLimit"`

	// Spark批作业集群Session资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`

	// 引擎资源弹性伸缩策略
	ScheduleElasticityConf *ScheduleElasticityConf `json:"ScheduleElasticityConf,omitnil,omitempty" name:"ScheduleElasticityConf"`
}

func (r *UpdateDataEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Size")
	delete(f, "MinClusters")
	delete(f, "MaxClusters")
	delete(f, "AutoResume")
	delete(f, "DataEngineName")
	delete(f, "Message")
	delete(f, "AutoSuspend")
	delete(f, "CrontabResumeSuspend")
	delete(f, "CrontabResumeSuspendStrategy")
	delete(f, "MaxConcurrency")
	delete(f, "TolerableQueueTime")
	delete(f, "AutoSuspendTime")
	delete(f, "ElasticSwitch")
	delete(f, "ElasticLimit")
	delete(f, "SessionResourceTemplate")
	delete(f, "ScheduleElasticityConf")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataEngineResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDataEngineResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataEngineResponseParams `json:"Response"`
}

func (r *UpdateDataEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataMaskStrategyRequestParams struct {
	// 数据脱敏策略详情
	Strategy *DataMaskStrategyInfo `json:"Strategy,omitnil,omitempty" name:"Strategy"`
}

type UpdateDataMaskStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 数据脱敏策略详情
	Strategy *DataMaskStrategyInfo `json:"Strategy,omitnil,omitempty" name:"Strategy"`
}

func (r *UpdateDataMaskStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataMaskStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Strategy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateDataMaskStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateDataMaskStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateDataMaskStrategyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateDataMaskStrategyResponseParams `json:"Response"`
}

func (r *UpdateDataMaskStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateDataMaskStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEngineResourceGroupNetworkConfigInfoRequestParams struct {
	// 引擎资源组ID
	EngineResourceGroupId *string `json:"EngineResourceGroupId,omitnil,omitempty" name:"EngineResourceGroupId"`

	// 是否立即重启资源组生效，0--立即生效，1--只保持不重启生效
	IsEffectiveNow *int64 `json:"IsEffectiveNow,omitnil,omitempty" name:"IsEffectiveNow"`

	// 资源组绑定的网络配置名称集合
	NetworkConfigNames []*string `json:"NetworkConfigNames,omitnil,omitempty" name:"NetworkConfigNames"`
}

type UpdateEngineResourceGroupNetworkConfigInfoRequest struct {
	*tchttp.BaseRequest
	
	// 引擎资源组ID
	EngineResourceGroupId *string `json:"EngineResourceGroupId,omitnil,omitempty" name:"EngineResourceGroupId"`

	// 是否立即重启资源组生效，0--立即生效，1--只保持不重启生效
	IsEffectiveNow *int64 `json:"IsEffectiveNow,omitnil,omitempty" name:"IsEffectiveNow"`

	// 资源组绑定的网络配置名称集合
	NetworkConfigNames []*string `json:"NetworkConfigNames,omitnil,omitempty" name:"NetworkConfigNames"`
}

func (r *UpdateEngineResourceGroupNetworkConfigInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEngineResourceGroupNetworkConfigInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupId")
	delete(f, "IsEffectiveNow")
	delete(f, "NetworkConfigNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEngineResourceGroupNetworkConfigInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEngineResourceGroupNetworkConfigInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateEngineResourceGroupNetworkConfigInfoResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEngineResourceGroupNetworkConfigInfoResponseParams `json:"Response"`
}

func (r *UpdateEngineResourceGroupNetworkConfigInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEngineResourceGroupNetworkConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNetworkConnectionRequestParams struct {
	// 网络配置描述
	NetworkConnectionDesc *string `json:"NetworkConnectionDesc,omitnil,omitempty" name:"NetworkConnectionDesc"`

	// 网络配置名称
	NetworkConnectionName *string `json:"NetworkConnectionName,omitnil,omitempty" name:"NetworkConnectionName"`
}

type UpdateNetworkConnectionRequest struct {
	*tchttp.BaseRequest
	
	// 网络配置描述
	NetworkConnectionDesc *string `json:"NetworkConnectionDesc,omitnil,omitempty" name:"NetworkConnectionDesc"`

	// 网络配置名称
	NetworkConnectionName *string `json:"NetworkConnectionName,omitnil,omitempty" name:"NetworkConnectionName"`
}

func (r *UpdateNetworkConnectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNetworkConnectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NetworkConnectionDesc")
	delete(f, "NetworkConnectionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateNetworkConnectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateNetworkConnectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateNetworkConnectionResponse struct {
	*tchttp.BaseResponse
	Response *UpdateNetworkConnectionResponseParams `json:"Response"`
}

func (r *UpdateNetworkConnectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateNetworkConnectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRowFilterRequestParams struct {
	// 行过滤策略的id，此值可以通过DescribeUserInfo或者DescribeWorkGroupInfo接口获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 新的过滤策略。
	Policy *Policy `json:"Policy,omitnil,omitempty" name:"Policy"`
}

type UpdateRowFilterRequest struct {
	*tchttp.BaseRequest
	
	// 行过滤策略的id，此值可以通过DescribeUserInfo或者DescribeWorkGroupInfo接口获取
	PolicyId *int64 `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 新的过滤策略。
	Policy *Policy `json:"Policy,omitnil,omitempty" name:"Policy"`
}

func (r *UpdateRowFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRowFilterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PolicyId")
	delete(f, "Policy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateRowFilterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateRowFilterResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateRowFilterResponse struct {
	*tchttp.BaseResponse
	Response *UpdateRowFilterResponseParams `json:"Response"`
}

func (r *UpdateRowFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateRowFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStandardEngineResourceGroupBaseInfoRequestParams struct {
	// 引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 自动启动，（任务提交自动拉起资源组）0-自动启动，1-不自动启动
	AutoLaunch *int64 `json:"AutoLaunch,omitnil,omitempty" name:"AutoLaunch"`

	// 自动挂起资源组。0-自动挂起，1-不自动挂起
	AutoPause *int64 `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 自动挂起时间，单位分钟，取值范围在1-999（在无任务AutoPauseTime后，资源组自动挂起）
	AutoPauseTime *int64 `json:"AutoPauseTime,omitnil,omitempty" name:"AutoPauseTime"`

	// 任务并发数
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`
}

type UpdateStandardEngineResourceGroupBaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 自动启动，（任务提交自动拉起资源组）0-自动启动，1-不自动启动
	AutoLaunch *int64 `json:"AutoLaunch,omitnil,omitempty" name:"AutoLaunch"`

	// 自动挂起资源组。0-自动挂起，1-不自动挂起
	AutoPause *int64 `json:"AutoPause,omitnil,omitempty" name:"AutoPause"`

	// 自动挂起时间，单位分钟，取值范围在1-999（在无任务AutoPauseTime后，资源组自动挂起）
	AutoPauseTime *int64 `json:"AutoPauseTime,omitnil,omitempty" name:"AutoPauseTime"`

	// 任务并发数
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`
}

func (r *UpdateStandardEngineResourceGroupBaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStandardEngineResourceGroupBaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupName")
	delete(f, "AutoLaunch")
	delete(f, "AutoPause")
	delete(f, "AutoPauseTime")
	delete(f, "MaxConcurrency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateStandardEngineResourceGroupBaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStandardEngineResourceGroupBaseInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateStandardEngineResourceGroupBaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *UpdateStandardEngineResourceGroupBaseInfoResponseParams `json:"Response"`
}

func (r *UpdateStandardEngineResourceGroupBaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStandardEngineResourceGroupBaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStandardEngineResourceGroupConfigInfoRequestParams struct {
	// 引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 需要更新的配置
	UpdateConfContext []*UpdateConfContext `json:"UpdateConfContext,omitnil,omitempty" name:"UpdateConfContext"`

	// 是否立即重启资源组生效，0--立即生效，1--只保持不重启生效
	IsEffectiveNow *int64 `json:"IsEffectiveNow,omitnil,omitempty" name:"IsEffectiveNow"`
}

type UpdateStandardEngineResourceGroupConfigInfoRequest struct {
	*tchttp.BaseRequest
	
	// 引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// 需要更新的配置
	UpdateConfContext []*UpdateConfContext `json:"UpdateConfContext,omitnil,omitempty" name:"UpdateConfContext"`

	// 是否立即重启资源组生效，0--立即生效，1--只保持不重启生效
	IsEffectiveNow *int64 `json:"IsEffectiveNow,omitnil,omitempty" name:"IsEffectiveNow"`
}

func (r *UpdateStandardEngineResourceGroupConfigInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStandardEngineResourceGroupConfigInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupName")
	delete(f, "UpdateConfContext")
	delete(f, "IsEffectiveNow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateStandardEngineResourceGroupConfigInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStandardEngineResourceGroupConfigInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateStandardEngineResourceGroupConfigInfoResponse struct {
	*tchttp.BaseResponse
	Response *UpdateStandardEngineResourceGroupConfigInfoResponseParams `json:"Response"`
}

func (r *UpdateStandardEngineResourceGroupConfigInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStandardEngineResourceGroupConfigInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStandardEngineResourceGroupResourceInfoRequestParams struct {
	// 引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// driver的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	DriverCuSpec *string `json:"DriverCuSpec,omitnil,omitempty" name:"DriverCuSpec"`

	// executor的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	ExecutorCuSpec *string `json:"ExecutorCuSpec,omitnil,omitempty" name:"ExecutorCuSpec"`

	// executor最小数量，
	MinExecutorNums *int64 `json:"MinExecutorNums,omitnil,omitempty" name:"MinExecutorNums"`

	// executor最大数量
	MaxExecutorNums *int64 `json:"MaxExecutorNums,omitnil,omitempty" name:"MaxExecutorNums"`

	// 是否立即重启资源组生效，0--立即生效，1--只保持不重启生效
	IsEffectiveNow *int64 `json:"IsEffectiveNow,omitnil,omitempty" name:"IsEffectiveNow"`

	// AI资源组资源上限
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 镜像类型，内置镜像：built-in，自定义镜像：custom
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像版本，镜像id
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// 框架类型
	FrameType *string `json:"FrameType,omitnil,omitempty" name:"FrameType"`

	// 自定义镜像域名
	PublicDomain *string `json:"PublicDomain,omitnil,omitempty" name:"PublicDomain"`

	// 自定义镜像实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 自定义镜像所属地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// python类型资源组python单机节点资源上限，该值要小于资源组的资源上限.small:1cu medium:2cu large:4cu xlarge:8cu 4xlarge:16cu 8xlarge:32cu 16xlarge:64cu，如果是高内存型资源，在类型前面加上m.
	PythonCuSpec *string `json:"PythonCuSpec,omitnil,omitempty" name:"PythonCuSpec"`

	// 仅SQL资源组资源配置模式，fast：快速模式，custom：自定义模式
	SparkSpecMode *string `json:"SparkSpecMode,omitnil,omitempty" name:"SparkSpecMode"`

	// 仅SQL资源组资源上限，仅用于快速模式
	SparkSize *int64 `json:"SparkSize,omitnil,omitempty" name:"SparkSize"`

	// gpuDriver规格
	DriverGPUSpec *int64 `json:"DriverGPUSpec,omitnil,omitempty" name:"DriverGPUSpec"`

	// gpuExcutor 规格
	ExecutorGPUSpec *int64 `json:"ExecutorGPUSpec,omitnil,omitempty" name:"ExecutorGPUSpec"`

	// gpu 上限
	GPULimitSize *int64 `json:"GPULimitSize,omitnil,omitempty" name:"GPULimitSize"`

	// gpu 规格
	GPUSize *int64 `json:"GPUSize,omitnil,omitempty" name:"GPUSize"`

	// gpupod 规格
	PythonGPUSpec *int64 `json:"PythonGPUSpec,omitnil,omitempty" name:"PythonGPUSpec"`
}

type UpdateStandardEngineResourceGroupResourceInfoRequest struct {
	*tchttp.BaseRequest
	
	// 引擎资源组名称
	EngineResourceGroupName *string `json:"EngineResourceGroupName,omitnil,omitempty" name:"EngineResourceGroupName"`

	// driver的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	DriverCuSpec *string `json:"DriverCuSpec,omitnil,omitempty" name:"DriverCuSpec"`

	// executor的cu规格：
	// 当前支持：small（默认，1cu）、medium（2cu）、large（4cu）、xlarge（8cu），内存型cu为cpu：men=1:8，m.small（1cu内存型）、m.medium（2cu内存型）、m.large（4cu内存型）、m.xlarge（8cu内存型）
	ExecutorCuSpec *string `json:"ExecutorCuSpec,omitnil,omitempty" name:"ExecutorCuSpec"`

	// executor最小数量，
	MinExecutorNums *int64 `json:"MinExecutorNums,omitnil,omitempty" name:"MinExecutorNums"`

	// executor最大数量
	MaxExecutorNums *int64 `json:"MaxExecutorNums,omitnil,omitempty" name:"MaxExecutorNums"`

	// 是否立即重启资源组生效，0--立即生效，1--只保持不重启生效
	IsEffectiveNow *int64 `json:"IsEffectiveNow,omitnil,omitempty" name:"IsEffectiveNow"`

	// AI资源组资源上限
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 镜像类型，内置镜像：built-in，自定义镜像：custom
	ImageType *string `json:"ImageType,omitnil,omitempty" name:"ImageType"`

	// 镜像名称
	ImageName *string `json:"ImageName,omitnil,omitempty" name:"ImageName"`

	// 镜像版本，镜像id
	ImageVersion *string `json:"ImageVersion,omitnil,omitempty" name:"ImageVersion"`

	// 框架类型
	FrameType *string `json:"FrameType,omitnil,omitempty" name:"FrameType"`

	// 自定义镜像域名
	PublicDomain *string `json:"PublicDomain,omitnil,omitempty" name:"PublicDomain"`

	// 自定义镜像实例id
	RegistryId *string `json:"RegistryId,omitnil,omitempty" name:"RegistryId"`

	// 自定义镜像所属地域
	RegionName *string `json:"RegionName,omitnil,omitempty" name:"RegionName"`

	// python类型资源组python单机节点资源上限，该值要小于资源组的资源上限.small:1cu medium:2cu large:4cu xlarge:8cu 4xlarge:16cu 8xlarge:32cu 16xlarge:64cu，如果是高内存型资源，在类型前面加上m.
	PythonCuSpec *string `json:"PythonCuSpec,omitnil,omitempty" name:"PythonCuSpec"`

	// 仅SQL资源组资源配置模式，fast：快速模式，custom：自定义模式
	SparkSpecMode *string `json:"SparkSpecMode,omitnil,omitempty" name:"SparkSpecMode"`

	// 仅SQL资源组资源上限，仅用于快速模式
	SparkSize *int64 `json:"SparkSize,omitnil,omitempty" name:"SparkSize"`

	// gpuDriver规格
	DriverGPUSpec *int64 `json:"DriverGPUSpec,omitnil,omitempty" name:"DriverGPUSpec"`

	// gpuExcutor 规格
	ExecutorGPUSpec *int64 `json:"ExecutorGPUSpec,omitnil,omitempty" name:"ExecutorGPUSpec"`

	// gpu 上限
	GPULimitSize *int64 `json:"GPULimitSize,omitnil,omitempty" name:"GPULimitSize"`

	// gpu 规格
	GPUSize *int64 `json:"GPUSize,omitnil,omitempty" name:"GPUSize"`

	// gpupod 规格
	PythonGPUSpec *int64 `json:"PythonGPUSpec,omitnil,omitempty" name:"PythonGPUSpec"`
}

func (r *UpdateStandardEngineResourceGroupResourceInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStandardEngineResourceGroupResourceInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineResourceGroupName")
	delete(f, "DriverCuSpec")
	delete(f, "ExecutorCuSpec")
	delete(f, "MinExecutorNums")
	delete(f, "MaxExecutorNums")
	delete(f, "IsEffectiveNow")
	delete(f, "Size")
	delete(f, "ImageType")
	delete(f, "ImageName")
	delete(f, "ImageVersion")
	delete(f, "FrameType")
	delete(f, "PublicDomain")
	delete(f, "RegistryId")
	delete(f, "RegionName")
	delete(f, "PythonCuSpec")
	delete(f, "SparkSpecMode")
	delete(f, "SparkSize")
	delete(f, "DriverGPUSpec")
	delete(f, "ExecutorGPUSpec")
	delete(f, "GPULimitSize")
	delete(f, "GPUSize")
	delete(f, "PythonGPUSpec")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateStandardEngineResourceGroupResourceInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateStandardEngineResourceGroupResourceInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateStandardEngineResourceGroupResourceInfoResponse struct {
	*tchttp.BaseResponse
	Response *UpdateStandardEngineResourceGroupResourceInfoResponseParams `json:"Response"`
}

func (r *UpdateStandardEngineResourceGroupResourceInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateStandardEngineResourceGroupResourceInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUDFPolicyRequestParams struct {
	// UDF名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// UDF权限信息
	UDFPolicyInfos []*UDFPolicyInfo `json:"UDFPolicyInfos,omitnil,omitempty" name:"UDFPolicyInfos"`
}

type UpdateUDFPolicyRequest struct {
	*tchttp.BaseRequest
	
	// UDF名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据目录名
	CatalogName *string `json:"CatalogName,omitnil,omitempty" name:"CatalogName"`

	// UDF权限信息
	UDFPolicyInfos []*UDFPolicyInfo `json:"UDFPolicyInfos,omitnil,omitempty" name:"UDFPolicyInfos"`
}

func (r *UpdateUDFPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUDFPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DatabaseName")
	delete(f, "CatalogName")
	delete(f, "UDFPolicyInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUDFPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUDFPolicyResponseParams struct {
	// UDF权限信息
	UDFPolicyInfos []*UDFPolicyInfo `json:"UDFPolicyInfos,omitnil,omitempty" name:"UDFPolicyInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUDFPolicyResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUDFPolicyResponseParams `json:"Response"`
}

func (r *UpdateUDFPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUDFPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDataEngineConfigRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 用户自定义引擎配置项集合。该参数需要传用户需要添加的全部配置项，例如，已有配置项k1:v1，添加k2:v2，需要传[k1:v1,k2:v2]。
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil,omitempty" name:"DataEngineConfigPairs"`

	// 作业引擎资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`
}

type UpdateUserDataEngineConfigRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`

	// 用户自定义引擎配置项集合。该参数需要传用户需要添加的全部配置项，例如，已有配置项k1:v1，添加k2:v2，需要传[k1:v1,k2:v2]。
	DataEngineConfigPairs []*DataEngineConfigPair `json:"DataEngineConfigPairs,omitnil,omitempty" name:"DataEngineConfigPairs"`

	// 作业引擎资源配置模板
	SessionResourceTemplate *SessionResourceTemplate `json:"SessionResourceTemplate,omitnil,omitempty" name:"SessionResourceTemplate"`
}

func (r *UpdateUserDataEngineConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDataEngineConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	delete(f, "DataEngineConfigPairs")
	delete(f, "SessionResourceTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUserDataEngineConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUserDataEngineConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUserDataEngineConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUserDataEngineConfigResponseParams `json:"Response"`
}

func (r *UpdateUserDataEngineConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUserDataEngineConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDataEngineImageRequestParams struct {
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

type UpgradeDataEngineImageRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	DataEngineId *string `json:"DataEngineId,omitnil,omitempty" name:"DataEngineId"`
}

func (r *UpgradeDataEngineImageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDataEngineImageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DataEngineId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeDataEngineImageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeDataEngineImageResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeDataEngineImageResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeDataEngineImageResponseParams `json:"Response"`
}

func (r *UpgradeDataEngineImageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeDataEngineImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserDetailInfo struct {
	// 用户Id
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 返回的信息类型，Group：返回的当前用户的工作组信息；DataAuth：返回的当前用户的数据权限信息；EngineAuth：返回的当前用户的引擎权限信息
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 用户类型：ADMIN：管理员 COMMON：一般用户
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户描述信息
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`

	// 数据权限信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataPolicyInfo *Policys `json:"DataPolicyInfo,omitnil,omitempty" name:"DataPolicyInfo"`

	// 引擎权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnginePolicyInfo *Policys `json:"EnginePolicyInfo,omitnil,omitempty" name:"EnginePolicyInfo"`

	// 绑定到该用户的工作组集合信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupInfo *WorkGroups `json:"WorkGroupInfo,omitnil,omitempty" name:"WorkGroupInfo"`

	// 用户别名
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 行过滤集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	RowFilterInfo *Policys `json:"RowFilterInfo,omitnil,omitempty" name:"RowFilterInfo"`

	// 账号类型
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`

	// 数据源权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogPolicyInfo *Policys `json:"CatalogPolicyInfo,omitnil,omitempty" name:"CatalogPolicyInfo"`
}

type UserIdSetOfWorkGroupId struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 用户Id集合，和CAM侧Uin匹配
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type UserInfo struct {
	// 用户Id，和子用户uin相同
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户描述信息，方便区分不同用户
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`

	// 单独给用户绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间，格式如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 关联的工作组集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupSet []*WorkGroupMessage `json:"WorkGroupSet,omitnil,omitempty" name:"WorkGroupSet"`

	// 是否是主账号
	IsOwner *bool `json:"IsOwner,omitnil,omitempty" name:"IsOwner"`

	// 用户类型。ADMIN：管理员 COMMON：普通用户。
	UserType *string `json:"UserType,omitnil,omitempty" name:"UserType"`

	// 用户别名
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 账号类型
	AccountType *string `json:"AccountType,omitnil,omitempty" name:"AccountType"`
}

type UserMessage struct {
	// 用户Id，和CAM侧子用户Uin匹配
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDescription *string `json:"UserDescription,omitnil,omitempty" name:"UserDescription"`

	// 当前用户的创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 当前用户的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 用户别名
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`
}

type UserRole struct {
	// 角色ID
	RoleId *int64 `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 用户app ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 用户ID
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 角色权限
	Arn *string `json:"Arn,omitnil,omitempty" name:"Arn"`

	// 最近修改时间戳
	ModifyTime *int64 `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 角色描述信息
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 创建者UIN
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// cos授权路径列表
	CosPermissionList []*CosPermission `json:"CosPermissionList,omitnil,omitempty" name:"CosPermissionList"`

	// cam策略json
	PermissionJson *string `json:"PermissionJson,omitnil,omitempty" name:"PermissionJson"`

	// 是否设置为常驻：1非常驻（默认）、2常驻（仅能设置一个常驻）
	IsDefault *int64 `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`
}

type UserVpcConnectionInfo struct {
	// 引擎网络ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineNetworkId *string `json:"EngineNetworkId,omitnil,omitempty" name:"EngineNetworkId"`

	// 用户vpcid
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVpcId *string `json:"UserVpcId,omitnil,omitempty" name:"UserVpcId"`

	// 用户终端节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVpcEndpointId *string `json:"UserVpcEndpointId,omitnil,omitempty" name:"UserVpcEndpointId"`

	// 用户终端节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserVpcEndpointName *string `json:"UserVpcEndpointName,omitnil,omitempty" name:"UserVpcEndpointName"`

	// 接入点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessConnectionInfos []*string `json:"AccessConnectionInfos,omitnil,omitempty" name:"AccessConnectionInfos"`
}

type Users struct {
	// 用户信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSet []*UserMessage `json:"UserSet,omitnil,omitempty" name:"UserSet"`

	// 用户总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type ViewBaseInfo struct {
	// 该视图所属数据库名字
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 视图名称
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// 视图创建人昵称
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 视图创建人ID
	UserSubUin *string `json:"UserSubUin,omitnil,omitempty" name:"UserSubUin"`
}

type ViewResponseInfo struct {
	// 视图基本信息。
	ViewBaseInfo *ViewBaseInfo `json:"ViewBaseInfo,omitnil,omitempty" name:"ViewBaseInfo"`

	// 视图列信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 视图属性信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 视图创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 视图更新时间。
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
}

type VpcCidrBlock struct {
	// 子网Id
	CidrId *string `json:"CidrId,omitnil,omitempty" name:"CidrId"`

	// 子网网段
	CidrAddr *string `json:"CidrAddr,omitnil,omitempty" name:"CidrAddr"`
}

type VpcInfo struct {
	// vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// vpc子网
	VpcCidrBlock *string `json:"VpcCidrBlock,omitnil,omitempty" name:"VpcCidrBlock"`

	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 权限组Id
	AccessGroupId *string `json:"AccessGroupId,omitnil,omitempty" name:"AccessGroupId"`
}

type WorkGroupDetailInfo struct {
	// 工作组Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil,omitempty" name:"WorkGroupName"`

	// 包含的信息类型。User：用户信息；DataAuth：数据权限；EngineAuth:引擎权限
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 工作组上绑定的用户集合
	UserInfo *Users `json:"UserInfo,omitnil,omitempty" name:"UserInfo"`

	// 数据权限集合
	DataPolicyInfo *Policys `json:"DataPolicyInfo,omitnil,omitempty" name:"DataPolicyInfo"`

	// 引擎权限集合
	EnginePolicyInfo *Policys `json:"EnginePolicyInfo,omitnil,omitempty" name:"EnginePolicyInfo"`

	// 工作组描述信息
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil,omitempty" name:"WorkGroupDescription"`

	// 行过滤信息集合
	RowFilterInfo *Policys `json:"RowFilterInfo,omitnil,omitempty" name:"RowFilterInfo"`

	// 数据目录权限集
	// 注意：此字段可能返回 null，表示取不到有效值。
	CatalogPolicyInfo *Policy `json:"CatalogPolicyInfo,omitnil,omitempty" name:"CatalogPolicyInfo"`
}

type WorkGroupIdSetOfUserId struct {
	// 用户Id，和CAM侧Uin匹配
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 工作组Id集合
	WorkGroupIds []*int64 `json:"WorkGroupIds,omitnil,omitempty" name:"WorkGroupIds"`
}

type WorkGroupInfo struct {
	// 查询到的工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil,omitempty" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil,omitempty" name:"WorkGroupDescription"`

	// 工作组关联的用户数量
	UserNum *int64 `json:"UserNum,omitnil,omitempty" name:"UserNum"`

	// 工作组关联的用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSet []*UserMessage `json:"UserSet,omitnil,omitempty" name:"UserSet"`

	// 工作组绑定的权限集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicySet []*Policy `json:"PolicySet,omitnil,omitempty" name:"PolicySet"`

	// 工作组的创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 工作组的创建时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type WorkGroupMessage struct {
	// 工作组唯一Id
	WorkGroupId *int64 `json:"WorkGroupId,omitnil,omitempty" name:"WorkGroupId"`

	// 工作组名称
	WorkGroupName *string `json:"WorkGroupName,omitnil,omitempty" name:"WorkGroupName"`

	// 工作组描述
	WorkGroupDescription *string `json:"WorkGroupDescription,omitnil,omitempty" name:"WorkGroupDescription"`

	// 创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 工作组创建的时间，形如2021-07-28 16:19:32
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type WorkGroups struct {
	// 工作组信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkGroupSet []*WorkGroupMessage `json:"WorkGroupSet,omitnil,omitempty" name:"WorkGroupSet"`

	// 工作组总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}