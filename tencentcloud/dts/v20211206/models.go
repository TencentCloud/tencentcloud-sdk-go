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

package v20211206

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type CheckStep struct {
	// 步骤编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNo *uint64 `json:"StepNo,omitnil" name:"StepNo"`

	// 步骤Id， 如：ConnectDBCheck、VersionCheck、SrcPrivilegeCheck等，具体校验项和源目标实例相关
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepId *string `json:"StepId,omitnil" name:"StepId"`

	// 步骤名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepName *string `json:"StepName,omitnil" name:"StepName"`

	// 此检查步骤的结果，pass(校验通过)、failed(校验失败)、notStarted(校验还未开始进行)、blocked(检验阻塞)、warning(校验有告警，但仍通过)
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepStatus *string `json:"StepStatus,omitnil" name:"StepStatus"`

	// 此检查步骤的错误消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepMessage *string `json:"StepMessage,omitnil" name:"StepMessage"`

	// 每个检查步骤里的具体检查项
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailCheckItems []*DetailCheckItem `json:"DetailCheckItems,omitnil" name:"DetailCheckItems"`

	// 是否已跳过
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasSkipped *bool `json:"HasSkipped,omitnil" name:"HasSkipped"`
}

type CheckStepInfo struct {
	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartAt *string `json:"StartAt,omitnil" name:"StartAt"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndAt *string `json:"EndAt,omitnil" name:"EndAt"`

	// 任务步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *ProcessProgress `json:"Progress,omitnil" name:"Progress"`
}

type Column struct {
	// 列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnName *string `json:"ColumnName,omitnil" name:"ColumnName"`

	// 新列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewColumnName *string `json:"NewColumnName,omitnil" name:"NewColumnName"`
}

type CompareAbstractInfo struct {
	// 校验配置参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Options *CompareOptions `json:"Options,omitnil" name:"Options"`

	// 一致性校验对比对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Objects *CompareObject `json:"Objects,omitnil" name:"Objects"`

	// 对比结论: same,different
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conclusion *string `json:"Conclusion,omitnil" name:"Conclusion"`

	// 任务状态: success,failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 总的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTables *uint64 `json:"TotalTables,omitnil" name:"TotalTables"`

	// 已校验的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckedTables *uint64 `json:"CheckedTables,omitnil" name:"CheckedTables"`

	// 不一致的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DifferentTables *uint64 `json:"DifferentTables,omitnil" name:"DifferentTables"`

	// 跳过校验的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkippedTables *uint64 `json:"SkippedTables,omitnil" name:"SkippedTables"`

	// 预估表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NearlyTableCount *uint64 `json:"NearlyTableCount,omitnil" name:"NearlyTableCount"`

	// 不一致的数据行数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DifferentRows *uint64 `json:"DifferentRows,omitnil" name:"DifferentRows"`

	// 源库行数，当对比类型为**行数对比**时此项有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcSampleRows *uint64 `json:"SrcSampleRows,omitnil" name:"SrcSampleRows"`

	// 目标库行数，当对比类型为**行数对比**时此项有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstSampleRows *uint64 `json:"DstSampleRows,omitnil" name:"DstSampleRows"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartedAt *string `json:"StartedAt,omitnil" name:"StartedAt"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedAt *string `json:"FinishedAt,omitnil" name:"FinishedAt"`
}

type CompareColumnItem struct {
	// 列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnName *string `json:"ColumnName,omitnil" name:"ColumnName"`
}

type CompareDetailInfo struct {
	// 数据不一致的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Difference *DifferenceDetail `json:"Difference,omitnil" name:"Difference"`

	// 跳过校验的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Skipped *SkippedDetail `json:"Skipped,omitnil" name:"Skipped"`
}

type CompareObject struct {
	// 对象模式 整实例-all,部分对象-partial
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectMode *string `json:"ObjectMode,omitnil" name:"ObjectMode"`

	// 对象列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectItems []*CompareObjectItem `json:"ObjectItems,omitnil" name:"ObjectItems"`

	// 高级对象类型，如account(账号),index(索引),shardkey(片键，后面可能会调整),schema(库表结构)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedObjects []*string `json:"AdvancedObjects,omitnil" name:"AdvancedObjects"`
}

type CompareObjectItem struct {
	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 数据库选择模式: all 为当前对象下的所有对象,partial 为部分对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbMode *string `json:"DbMode,omitnil" name:"DbMode"`

	// schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 表选择模式: all 为当前对象下的所有表对象,partial 为部分表对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMode *string `json:"TableMode,omitnil" name:"TableMode"`

	// 用于一致性校验的表配置，当 TableMode 为 partial 时，需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*CompareTableItem `json:"Tables,omitnil" name:"Tables"`

	// 视图选择模式: all 为当前对象下的所有视图对象,partial 为部分视图对象(一致性校验不校验视图，当前参数未启作用)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewMode *string `json:"ViewMode,omitnil" name:"ViewMode"`

	// 用于一致性校验的视图配置，当 ViewMode 为 partial 时， 需要填写(一致性校验不校验视图，当前参数未启作用)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Views []*CompareViewItem `json:"Views,omitnil" name:"Views"`
}

type CompareOptions struct {
	// 对比类型：dataCheck(完整数据对比)、sampleDataCheck(抽样数据对比)、rowsCount(行数对比)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitnil" name:"Method"`

	// 抽样比例;范围0,100
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleRate *int64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 线程数，取值1-5，默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThreadCount *int64 `json:"ThreadCount,omitnil" name:"ThreadCount"`
}

type CompareTableItem struct {
	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// column 模式，all 为全部，partial 表示部分(该参数仅对数据同步任务有效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnMode *string `json:"ColumnMode,omitnil" name:"ColumnMode"`

	// 当 ColumnMode 为 partial 时必填(该参数仅对数据同步任务有效)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*CompareColumnItem `json:"Columns,omitnil" name:"Columns"`
}

type CompareTaskInfo struct {
	// 一致性校验任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 一致性校验结果，包括：unstart(未启动)、running(校验中)、canceled(已终止)、failed(校验任务失败)、inconsistent(不一致)、consistent(一致)、notexist(不存在校验任务)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type CompareTaskItem struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 对比任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 对比任务状态, 可能的值：created - 创建完成；readyRun - 等待运行；running - 运行中；success - 成功；stopping - 结束中；failed - 失败；canceled - 已终止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 对比任务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *CompareObject `json:"Config,omitnil" name:"Config"`

	// 对比任务校验详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckProcess *ProcessProgress `json:"CheckProcess,omitnil" name:"CheckProcess"`

	// 对比任务运行详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareProcess *ProcessProgress `json:"CompareProcess,omitnil" name:"CompareProcess"`

	// 对比结果, 可能的值：same - 一致；different - 不一致；skipAll - 跳过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conclusion *string `json:"Conclusion,omitnil" name:"Conclusion"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil" name:"CreatedAt"`

	// 任务启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartedAt *string `json:"StartedAt,omitnil" name:"StartedAt"`

	// 对比结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedAt *string `json:"FinishedAt,omitnil" name:"FinishedAt"`

	// 对比类型，dataCheck(完整数据对比)、sampleDataCheck(抽样数据对比)、rowsCount(行数对比)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitnil" name:"Method"`

	// 对比配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Options *CompareOptions `json:"Options,omitnil" name:"Options"`

	// 一致性校验提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`
}

type CompareViewItem struct {
	// 视图名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitnil" name:"ViewName"`
}

// Predefined struct for user
type CompleteMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 完成任务的方式,仅支持旧版MySQL迁移任务。waitForSync-等待主从差距为0才停止,immediately-立即完成，不会等待主从差距一致。默认为waitForSync
	CompleteMode *string `json:"CompleteMode,omitnil" name:"CompleteMode"`
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 完成任务的方式,仅支持旧版MySQL迁移任务。waitForSync-等待主从差距为0才停止,immediately-立即完成，不会等待主从差距一致。默认为waitForSync
	CompleteMode *string `json:"CompleteMode,omitnil" name:"CompleteMode"`
}

func (r *CompleteMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CompleteMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *CompleteMigrateJobResponseParams `json:"Response"`
}

func (r *CompleteMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureSyncJobRequestParams struct {
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 源端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云),注意具体可选值依赖当前链路
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// 目标端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)、ckafka(CKafka实例),注意具体可选值依赖当前链路
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// 同步库表对象信息
	Objects *Objects `json:"Objects,omitnil" name:"Objects"`

	// 同步任务名称
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 枚举值是 liteMode 和 fullMode ，分别对应精简模式或正常模式
	JobMode *string `json:"JobMode,omitnil" name:"JobMode"`

	// 运行模式，取值如：Immediate(表示立即运行，默认为此项值)、Timed(表示定时运行)
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 期待启动时间，当RunMode取值为Timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitnil" name:"ExpectRunTime"`

	// 源端信息，单节点数据库使用，且SrcNodeType传single
	SrcInfo *Endpoint `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// 源端信息，多节点数据库使用，且SrcNodeType传cluster
	SrcInfos *SyncDBEndpointInfos `json:"SrcInfos,omitnil" name:"SrcInfos"`

	// 枚举值：cluster、single。源库为单节点数据库使用single，多节点使用cluster
	SrcNodeType *string `json:"SrcNodeType,omitnil" name:"SrcNodeType"`

	// 目标端信息，单节点数据库使用
	DstInfo *Endpoint `json:"DstInfo,omitnil" name:"DstInfo"`

	// 目标端信息，多节点数据库使用，且DstNodeType传cluster
	DstInfos *SyncDBEndpointInfos `json:"DstInfos,omitnil" name:"DstInfos"`

	// 枚举值：cluster、single。目标库为单节点数据库使用single，多节点使用cluster
	DstNodeType *string `json:"DstNodeType,omitnil" name:"DstNodeType"`

	// 同步任务选项；该字段下的RateLimitOption暂时无法生效、如果需要修改限速、可通过ModifySyncRateLimit接口完成限速
	Options *Options `json:"Options,omitnil" name:"Options"`

	// 自动重试的时间段、可设置5至720分钟、0表示不重试
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil" name:"AutoRetryTimeRangeMinutes"`
}

type ConfigureSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 源端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云),注意具体可选值依赖当前链路
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// 目标端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)、ckafka(CKafka实例),注意具体可选值依赖当前链路
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// 同步库表对象信息
	Objects *Objects `json:"Objects,omitnil" name:"Objects"`

	// 同步任务名称
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 枚举值是 liteMode 和 fullMode ，分别对应精简模式或正常模式
	JobMode *string `json:"JobMode,omitnil" name:"JobMode"`

	// 运行模式，取值如：Immediate(表示立即运行，默认为此项值)、Timed(表示定时运行)
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 期待启动时间，当RunMode取值为Timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitnil" name:"ExpectRunTime"`

	// 源端信息，单节点数据库使用，且SrcNodeType传single
	SrcInfo *Endpoint `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// 源端信息，多节点数据库使用，且SrcNodeType传cluster
	SrcInfos *SyncDBEndpointInfos `json:"SrcInfos,omitnil" name:"SrcInfos"`

	// 枚举值：cluster、single。源库为单节点数据库使用single，多节点使用cluster
	SrcNodeType *string `json:"SrcNodeType,omitnil" name:"SrcNodeType"`

	// 目标端信息，单节点数据库使用
	DstInfo *Endpoint `json:"DstInfo,omitnil" name:"DstInfo"`

	// 目标端信息，多节点数据库使用，且DstNodeType传cluster
	DstInfos *SyncDBEndpointInfos `json:"DstInfos,omitnil" name:"DstInfos"`

	// 枚举值：cluster、single。目标库为单节点数据库使用single，多节点使用cluster
	DstNodeType *string `json:"DstNodeType,omitnil" name:"DstNodeType"`

	// 同步任务选项；该字段下的RateLimitOption暂时无法生效、如果需要修改限速、可通过ModifySyncRateLimit接口完成限速
	Options *Options `json:"Options,omitnil" name:"Options"`

	// 自动重试的时间段、可设置5至720分钟、0表示不重试
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil" name:"AutoRetryTimeRangeMinutes"`
}

func (r *ConfigureSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "SrcAccessType")
	delete(f, "DstAccessType")
	delete(f, "Objects")
	delete(f, "JobName")
	delete(f, "JobMode")
	delete(f, "RunMode")
	delete(f, "ExpectRunTime")
	delete(f, "SrcInfo")
	delete(f, "SrcInfos")
	delete(f, "SrcNodeType")
	delete(f, "DstInfo")
	delete(f, "DstInfos")
	delete(f, "DstNodeType")
	delete(f, "Options")
	delete(f, "AutoRetryTimeRangeMinutes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ConfigureSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ConfigureSyncJobResponseParams `json:"Response"`
}

func (r *ConfigureSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ConfigureSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConflictHandleOption struct {
	// 条件覆盖的列
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionColumn *string `json:"ConditionColumn,omitnil" name:"ConditionColumn"`

	// 条件覆盖操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionOperator *string `json:"ConditionOperator,omitnil" name:"ConditionOperator"`

	// 条件覆盖优先级处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionOrderInSrcAndDst *string `json:"ConditionOrderInSrcAndDst,omitnil" name:"ConditionOrderInSrcAndDst"`
}

type ConsistencyOption struct {
	// 一致性检测类型: full(全量检测迁移对象)、noCheck(不检测)、notConfigured(未配置)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil" name:"Mode"`
}

// Predefined struct for user
type ContinueMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type ContinueMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *ContinueMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ContinueMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ContinueMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *ContinueMigrateJobResponseParams `json:"Response"`
}

func (r *ContinueMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type ContinueSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *ContinueSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ContinueSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ContinueSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ContinueSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ContinueSyncJobResponseParams `json:"Response"`
}

func (r *ContinueSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ContinueSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCheckSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type CreateCheckSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *CreateCheckSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCheckSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCheckSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCheckSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCheckSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateCheckSyncJobResponseParams `json:"Response"`
}

func (r *CreateCheckSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCheckSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCompareTaskRequestParams struct {
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 数据对比任务名称，若为空则默认给CompareTaskId相同值
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， 默认为此项配置)，custom(自定义模式)
	ObjectMode *string `json:"ObjectMode,omitnil" name:"ObjectMode"`

	// 一致性对比对象配置
	Objects *CompareObject `json:"Objects,omitnil" name:"Objects"`

	// 一致性校验选项
	Options *CompareOptions `json:"Options,omitnil" name:"Options"`
}

type CreateCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 数据对比任务名称，若为空则默认给CompareTaskId相同值
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， 默认为此项配置)，custom(自定义模式)
	ObjectMode *string `json:"ObjectMode,omitnil" name:"ObjectMode"`

	// 一致性对比对象配置
	Objects *CompareObject `json:"Objects,omitnil" name:"Objects"`

	// 一致性校验选项
	Options *CompareOptions `json:"Options,omitnil" name:"Options"`
}

func (r *CreateCompareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCompareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "TaskName")
	delete(f, "ObjectMode")
	delete(f, "Objects")
	delete(f, "Options")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCompareTaskResponseParams struct {
	// 数据对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCompareTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateCompareTaskResponseParams `json:"Response"`
}

func (r *CreateCompareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCompareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateCheckJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *CreateMigrateCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrateCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrateCheckJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateMigrateCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrateCheckJobResponseParams `json:"Response"`
}

func (r *CreateMigrateCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrateCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationServiceRequestParams struct {
	// 源实例数据库类型，如mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb,cynosdbmysql
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// 目标实例数据库类型，如mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb,cynosdbmysql
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil" name:"SrcRegion"`

	// 目标实例地域，如：ap-guangzhou。注意，目标地域必须和API请求地域保持一致。
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 实例规格，包括：small、medium、large、xlarge、2xlarge
	InstanceClass *string `json:"InstanceClass,omitnil" name:"InstanceClass"`

	// 购买数量，范围为[1,15]，默认为1
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 迁移服务名称，最大长度128
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`
}

type CreateMigrationServiceRequest struct {
	*tchttp.BaseRequest
	
	// 源实例数据库类型，如mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb,cynosdbmysql
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// 目标实例数据库类型，如mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb,cynosdbmysql
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil" name:"SrcRegion"`

	// 目标实例地域，如：ap-guangzhou。注意，目标地域必须和API请求地域保持一致。
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 实例规格，包括：small、medium、large、xlarge、2xlarge
	InstanceClass *string `json:"InstanceClass,omitnil" name:"InstanceClass"`

	// 购买数量，范围为[1,15]，默认为1
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 迁移服务名称，最大长度128
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`
}

func (r *CreateMigrationServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SrcDatabaseType")
	delete(f, "DstDatabaseType")
	delete(f, "SrcRegion")
	delete(f, "DstRegion")
	delete(f, "InstanceClass")
	delete(f, "Count")
	delete(f, "JobName")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMigrationServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMigrationServiceResponseParams struct {
	// 下单成功随机生成的迁移任务id列表，形如：dts-c1f6rs21
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobIds []*string `json:"JobIds,omitnil" name:"JobIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateMigrationServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateMigrationServiceResponseParams `json:"Response"`
}

func (r *CreateMigrationServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMigrationServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModifyCheckSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type CreateModifyCheckSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *CreateModifyCheckSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModifyCheckSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateModifyCheckSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateModifyCheckSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateModifyCheckSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateModifyCheckSyncJobResponseParams `json:"Response"`
}

func (r *CreateModifyCheckSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateModifyCheckSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSyncJobRequestParams struct {
	// 付款类型, 如：PrePay(表示包年包月)、PostPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 源端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// 源端数据库所在地域,如ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil" name:"SrcRegion"`

	// 目标端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql,kafka等
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// 目标端数据库所在地域,如ap-guangzhou
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 同步任务规格，Standard:标准版
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// 一次购买的同步任务数量，取值范围为[1, 10]，默认为1
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 自动续费标识，当PayMode值为PrePay则此项配置有意义，取值为：1（表示自动续费）、0（不自动续费，默认为此值）
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 同步链路规格，如micro,small,medium,large，默认为medium
	InstanceClass *string `json:"InstanceClass,omitnil" name:"InstanceClass"`

	// 同步任务名称
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 创建类似任务的现有任务Id
	ExistedJobId *string `json:"ExistedJobId,omitnil" name:"ExistedJobId"`
}

type CreateSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 付款类型, 如：PrePay(表示包年包月)、PostPay(表示按时按量)
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 源端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// 源端数据库所在地域,如ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil" name:"SrcRegion"`

	// 目标端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql,kafka等
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// 目标端数据库所在地域,如ap-guangzhou
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 同步任务规格，Standard:标准版
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// 一次购买的同步任务数量，取值范围为[1, 10]，默认为1
	Count *uint64 `json:"Count,omitnil" name:"Count"`

	// 自动续费标识，当PayMode值为PrePay则此项配置有意义，取值为：1（表示自动续费）、0（不自动续费，默认为此值）
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 同步链路规格，如micro,small,medium,large，默认为medium
	InstanceClass *string `json:"InstanceClass,omitnil" name:"InstanceClass"`

	// 同步任务名称
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 创建类似任务的现有任务Id
	ExistedJobId *string `json:"ExistedJobId,omitnil" name:"ExistedJobId"`
}

func (r *CreateSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PayMode")
	delete(f, "SrcDatabaseType")
	delete(f, "SrcRegion")
	delete(f, "DstDatabaseType")
	delete(f, "DstRegion")
	delete(f, "Specification")
	delete(f, "Tags")
	delete(f, "Count")
	delete(f, "AutoRenew")
	delete(f, "InstanceClass")
	delete(f, "JobName")
	delete(f, "ExistedJobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSyncJobResponseParams struct {
	// 同步任务ids
	JobIds []*string `json:"JobIds,omitnil" name:"JobIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *CreateSyncJobResponseParams `json:"Response"`
}

func (r *CreateSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DBEndpointInfo struct {
	// 实例所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 实例网络接入类型，如：extranet(外网)、ipv6(公网ipv6)、cvm(云主机自建)、dcg(专线接入)、vpncloud(vpn接入的实例)、cdb(云数据库)、ccn(云联网)、intranet(自研上云)、vpc(私有网络)等，注意具体可选值依赖当前链路
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *string `json:"AccessType,omitnil" name:"AccessType"`

	// 实例数据库类型，如：mysql,redis,mongodb,postgresql,mariadb,percona 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseType *string `json:"DatabaseType,omitnil" name:"DatabaseType"`

	// 节点类型，为空或者simple表示普通节点、cluster表示集群节点；对于mongo业务，取值为replicaset(mongodb副本集)、standalone(mongodb单节点)、cluster(mongodb集群)；对于redis实例，为空或simple(单节点)、cluster(集群)、cluster-cache(cache集群)、cluster-proxy(代理集群)
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil" name:"NodeType"`

	// 数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info []*DBInfo `json:"Info,omitnil" name:"Info"`

	// 实例服务提供商，如:"aliyun","others"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supplier *string `json:"Supplier,omitnil" name:"Supplier"`

	// MongoDB可定义如下的参数: 	['AuthDatabase':'admin', 
	// 'AuthFlag': "1",	'AuthMechanism':"SCRAM-SHA-1"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitnil" name:"ExtraAttr"`

	// 数据库所属网络环境，AccessType为云联网(ccn)时必填， UserIDC表示用户IDC、TencentVPC表示腾讯云VPC；
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseNetEnv *string `json:"DatabaseNetEnv,omitnil" name:"DatabaseNetEnv"`
}

type DBInfo struct {
	// 表示节点角色，针对分布式数据库，如mongodb中的mongos节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil" name:"Role"`

	// 内核版本，针对mariadb的不同内核版本等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbKernel *string `json:"DbKernel,omitnil" name:"DbKernel"`

	// 实例的IP地址，对于公网、专线、VPN、云联网、自研上云、VPC等接入方式此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil" name:"Host"`

	// 实例的端口，对于公网、云主机自建、专线、VPN、云联网、自研上云、VPC等接入方式此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 实例的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 实例的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// CVM实例短ID，格式如：ins-olgl39y8；与云服务器控制台页面显示的实例ID相同；如果接入类型为云主机自建的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmInstanceId *string `json:"CvmInstanceId,omitnil" name:"CvmInstanceId"`

	// VPN网关ID，格式如：vpngw-9ghexg7q；如果接入类型为vpncloud的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil" name:"UniqVpnGwId"`

	// 专线网关ID，格式如：dcg-0rxtqqxb；如果接入类型为专线接入的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqDcgId *string `json:"UniqDcgId,omitnil" name:"UniqDcgId"`

	// 数据库实例ID，格式如：cdb-powiqx8q；如果接入类型为云数据库的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 云联网ID，如：ccn-afp6kltc 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnGwId *string `json:"CcnGwId,omitnil" name:"CcnGwId"`

	// 私有网络ID，格式如：vpc-92jblxto；如果接入类型为vpc、vpncloud、ccn、dcg的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络下的子网ID，格式如：subnet-3paxmkdz；如果接入类型为vpc、vpncloud、ccn、dcg的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 数据库版本，当实例为RDS实例时才有效，格式如：5.6或者5.7，默认为5.6
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`

	// 实例所属账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitnil" name:"Account"`

	// 跨账号迁移时的角色,只允许[a-zA-Z0-9\-\_]+
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountRole *string `json:"AccountRole,omitnil" name:"AccountRole"`

	// 资源所属账号 为空或self(表示本账号内资源)、other(表示其他账户资源)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountMode *string `json:"AccountMode,omitnil" name:"AccountMode"`

	// 临时密钥Id，可通过 获取联合身份临时访问凭证获取临时密钥https://cloud.tencent.com/document/product/1312/48195
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitnil" name:"TmpSecretId"`

	// 临时密钥Key，可通过 获取联合身份临时访问凭证获取临时密钥https://cloud.tencent.com/document/product/1312/48195
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitnil" name:"TmpSecretKey"`

	// 临时Token，可通过 获取联合身份临时访问凭证获取临时密钥https://cloud.tencent.com/document/product/1312/48195
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpToken *string `json:"TmpToken,omitnil" name:"TmpToken"`
}

type DBItem struct {
	// 需要迁移或同步的库名，当ObjectMode为partial时，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 迁移或同步后的库名，默认与源库相同
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewDbName *string `json:"NewDbName,omitnil" name:"NewDbName"`

	// 迁移或同步的 schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 迁移或同步后的 schema name
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewSchemaName *string `json:"NewSchemaName,omitnil" name:"NewSchemaName"`

	// DB选择模式: all(为当前对象下的所有对象)，partial(部分对象)，当ObjectMode为partial时，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBMode *string `json:"DBMode,omitnil" name:"DBMode"`

	// schema选择模式: all(为当前对象下的所有对象)，partial(部分对象)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaMode *string `json:"SchemaMode,omitnil" name:"SchemaMode"`

	// 表选择模式: all(为当前对象下的所有对象)，partial(部分对象)，当DBMode为partial时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMode *string `json:"TableMode,omitnil" name:"TableMode"`

	// 表图对象集合，当 TableMode 为 partial 时，此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*TableItem `json:"Tables,omitnil" name:"Tables"`

	// 视图选择模式: all 为当前对象下的所有视图对象,partial 为部分视图对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewMode *string `json:"ViewMode,omitnil" name:"ViewMode"`

	// 视图对象集合，当 ViewMode 为 partial 时， 此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Views []*ViewItem `json:"Views,omitnil" name:"Views"`

	// postgresql独有参数，角色选择模式: all 为当前对象下的所有角色对象,partial 为部分角色对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleMode *string `json:"RoleMode,omitnil" name:"RoleMode"`

	// postgresql独有参数，当 RoleMode 为 partial 时， 此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*RoleItem `json:"Roles,omitnil" name:"Roles"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionMode *string `json:"FunctionMode,omitnil" name:"FunctionMode"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerMode *string `json:"TriggerMode,omitnil" name:"TriggerMode"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventMode *string `json:"EventMode,omitnil" name:"EventMode"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureMode *string `json:"ProcedureMode,omitnil" name:"ProcedureMode"`

	// FunctionMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Functions []*string `json:"Functions,omitnil" name:"Functions"`

	// ProcedureMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedures []*string `json:"Procedures,omitnil" name:"Procedures"`

	// EventMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*string `json:"Events,omitnil" name:"Events"`

	// TriggerMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Triggers []*string `json:"Triggers,omitnil" name:"Triggers"`
}

type Database struct {
	// 需要迁移或同步的库名，当ObjectMode为Partial时，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 迁移或同步后的库名，默认与源库相同
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewDbName *string `json:"NewDbName,omitnil" name:"NewDbName"`

	// DB选择模式: All(为当前对象下的所有对象)，Partial(部分对象)，当Mode为Partial时，此项必填。注意，高级对象的同步不依赖此值，如果整库同步此处应该为All。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbMode *string `json:"DbMode,omitnil" name:"DbMode"`

	// 迁移或同步的 schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil" name:"SchemaName"`

	// 迁移或同步后的 schema name
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewSchemaName *string `json:"NewSchemaName,omitnil" name:"NewSchemaName"`

	// 表选择模式: All(为当前对象下的所有对象)，Partial(部分对象)，当DBMode为Partial时此项必填，如果整库同步此处应该为All。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMode *string `json:"TableMode,omitnil" name:"TableMode"`

	// 表图对象集合，当 TableMode 为 Partial 时，此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*Table `json:"Tables,omitnil" name:"Tables"`

	// 视图选择模式: All 为当前对象下的所有视图对象,Partial 为部分视图对象，如果整库同步此处应该为All。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewMode *string `json:"ViewMode,omitnil" name:"ViewMode"`

	// 视图对象集合，当 ViewMode 为 Partial 时， 此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Views []*View `json:"Views,omitnil" name:"Views"`

	// 选择要同步的模式，Partial为部分，All为整选，如果整库同步此处应该为All。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionMode *string `json:"FunctionMode,omitnil" name:"FunctionMode"`

	// FunctionMode取值为Partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Functions []*string `json:"Functions,omitnil" name:"Functions"`

	// 选择要同步的模式，Partial为部分，All为整选，如果整库同步此处应该为All。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureMode *string `json:"ProcedureMode,omitnil" name:"ProcedureMode"`

	// ProcedureMode取值为Partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedures []*string `json:"Procedures,omitnil" name:"Procedures"`

	// 触发器迁移模式，All(为当前对象下的所有对象)，Partial(部分对象)，如果整库同步此处应该为All。数据同步暂不支持此高级对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerMode *string `json:"TriggerMode,omitnil" name:"TriggerMode"`

	// 当TriggerMode为partial，指定要迁移的触发器名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Triggers []*string `json:"Triggers,omitnil" name:"Triggers"`

	// 事件迁移模式，All(为当前对象下的所有对象)，Partial(部分对象)，如果整库同步此处应该为All。数据同步暂不支持此高级对象。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventMode *string `json:"EventMode,omitnil" name:"EventMode"`

	// 当EventMode为partial，指定要迁移的事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*string `json:"Events,omitnil" name:"Events"`
}

type DatabaseTableObject struct {
	// 迁移对象类型 all(全实例)，partial(部分对象)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectMode *string `json:"ObjectMode,omitnil" name:"ObjectMode"`

	// 迁移对象，当 ObjectMode 为 partial 时，不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*DBItem `json:"Databases,omitnil" name:"Databases"`

	// 高级对象类型，如trigger、function、procedure、event。注意：如果要迁移同步高级对象，此配置中应该包含对应的高级对象类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedObjects []*string `json:"AdvancedObjects,omitnil" name:"AdvancedObjects"`
}

type DdlOption struct {
	// ddl类型，如Database,Table,View,Index等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlObject *string `json:"DdlObject,omitnil" name:"DdlObject"`

	// ddl具体值，对于Database可取值[Create,Drop,Alter]<br>对于Table可取值[Create,Drop,Alter,Truncate,Rename]<br/>对于View可取值[Create,Drop]<br/>对于Index可取值[Create,Drop]
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlValue []*string `json:"DdlValue,omitnil" name:"DdlValue"`
}

// Predefined struct for user
type DeleteCompareTaskRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`
}

type DeleteCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`
}

func (r *DeleteCompareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCompareTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCompareTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCompareTaskResponseParams `json:"Response"`
}

func (r *DeleteCompareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCompareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckSyncJobResultRequestParams struct {
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23，此值必填
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DescribeCheckSyncJobResultRequest struct {
	*tchttp.BaseRequest
	
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23，此值必填
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DescribeCheckSyncJobResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckSyncJobResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCheckSyncJobResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCheckSyncJobResultResponseParams struct {
	// 校验任务执行状态，如：notStarted(未开始)、running(校验中)、failed(校验任务失败)、success(任务成功)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 步骤总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepCount *uint64 `json:"StepCount,omitnil" name:"StepCount"`

	// 当前所在步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepCur *uint64 `json:"StepCur,omitnil" name:"StepCur"`

	// 总体进度，范围为[0,100]
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitnil" name:"Progress"`

	// 步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfos []*StepInfo `json:"StepInfos,omitnil" name:"StepInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCheckSyncJobResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCheckSyncJobResultResponseParams `json:"Response"`
}

func (r *DescribeCheckSyncJobResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCheckSyncJobResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareReportRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 校验任务 Id
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 校验不一致结果的 limit
	DifferenceLimit *uint64 `json:"DifferenceLimit,omitnil" name:"DifferenceLimit"`

	// 不一致的 Offset
	DifferenceOffset *uint64 `json:"DifferenceOffset,omitnil" name:"DifferenceOffset"`

	// 搜索条件，不一致的库名
	DifferenceDB *string `json:"DifferenceDB,omitnil" name:"DifferenceDB"`

	// 搜索条件，不一致的表名
	DifferenceTable *string `json:"DifferenceTable,omitnil" name:"DifferenceTable"`

	// 未校验的 Limit
	SkippedLimit *uint64 `json:"SkippedLimit,omitnil" name:"SkippedLimit"`

	// 未校验的 Offset
	SkippedOffset *uint64 `json:"SkippedOffset,omitnil" name:"SkippedOffset"`

	// 搜索条件，未校验的库名
	SkippedDB *string `json:"SkippedDB,omitnil" name:"SkippedDB"`

	// 搜索条件，未校验的表名
	SkippedTable *string `json:"SkippedTable,omitnil" name:"SkippedTable"`
}

type DescribeCompareReportRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 校验任务 Id
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 校验不一致结果的 limit
	DifferenceLimit *uint64 `json:"DifferenceLimit,omitnil" name:"DifferenceLimit"`

	// 不一致的 Offset
	DifferenceOffset *uint64 `json:"DifferenceOffset,omitnil" name:"DifferenceOffset"`

	// 搜索条件，不一致的库名
	DifferenceDB *string `json:"DifferenceDB,omitnil" name:"DifferenceDB"`

	// 搜索条件，不一致的表名
	DifferenceTable *string `json:"DifferenceTable,omitnil" name:"DifferenceTable"`

	// 未校验的 Limit
	SkippedLimit *uint64 `json:"SkippedLimit,omitnil" name:"SkippedLimit"`

	// 未校验的 Offset
	SkippedOffset *uint64 `json:"SkippedOffset,omitnil" name:"SkippedOffset"`

	// 搜索条件，未校验的库名
	SkippedDB *string `json:"SkippedDB,omitnil" name:"SkippedDB"`

	// 搜索条件，未校验的表名
	SkippedTable *string `json:"SkippedTable,omitnil" name:"SkippedTable"`
}

func (r *DescribeCompareReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	delete(f, "DifferenceLimit")
	delete(f, "DifferenceOffset")
	delete(f, "DifferenceDB")
	delete(f, "DifferenceTable")
	delete(f, "SkippedLimit")
	delete(f, "SkippedOffset")
	delete(f, "SkippedDB")
	delete(f, "SkippedTable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompareReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareReportResponseParams struct {
	// 一致性校验摘要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abstract *CompareAbstractInfo `json:"Abstract,omitnil" name:"Abstract"`

	// 一致性校验详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *CompareDetailInfo `json:"Detail,omitnil" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCompareReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompareReportResponseParams `json:"Response"`
}

func (r *DescribeCompareReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareTasksRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 分页设置，表示每页显示多少条任务，默认为 20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 校验任务 ID
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 任务状态过滤，可能的值：created - 创建完成；readyRun - 等待运行；running - 运行中；success - 成功；stopping - 结束中；failed - 失败；canceled - 已终止
	Status []*string `json:"Status,omitnil" name:"Status"`
}

type DescribeCompareTasksRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 分页设置，表示每页显示多少条任务，默认为 20
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 校验任务 ID
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 任务状态过滤，可能的值：created - 创建完成；readyRun - 等待运行；running - 运行中；success - 成功；stopping - 结束中；failed - 失败；canceled - 已终止
	Status []*string `json:"Status,omitnil" name:"Status"`
}

func (r *DescribeCompareTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "CompareTaskId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompareTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareTasksResponseParams struct {
	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 一致性校验列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*CompareTaskItem `json:"Items,omitnil" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCompareTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCompareTasksResponseParams `json:"Response"`
}

func (r *DescribeCompareTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCompareTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateDBInstancesRequestParams struct {
	// 数据库类型，如mysql
	DatabaseType *string `json:"DatabaseType,omitnil" name:"DatabaseType"`

	// 实例作为迁移的源还是目标,src(表示源)，dst(表示目标)
	MigrateRole *string `json:"MigrateRole,omitnil" name:"MigrateRole"`

	// 云数据库实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 云数据库名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 资源所属账号 为空值或self(表示本账号内资源)、other(表示其他账户资源)
	AccountMode *string `json:"AccountMode,omitnil" name:"AccountMode"`

	// 临时密钥Id，若为跨账号资源此项必填
	TmpSecretId *string `json:"TmpSecretId,omitnil" name:"TmpSecretId"`

	// 临时密钥Key，若为跨账号资源此项必填
	TmpSecretKey *string `json:"TmpSecretKey,omitnil" name:"TmpSecretKey"`

	// 临时密钥Token，若为跨账号资源此项必填
	TmpToken *string `json:"TmpToken,omitnil" name:"TmpToken"`
}

type DescribeMigrateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库类型，如mysql
	DatabaseType *string `json:"DatabaseType,omitnil" name:"DatabaseType"`

	// 实例作为迁移的源还是目标,src(表示源)，dst(表示目标)
	MigrateRole *string `json:"MigrateRole,omitnil" name:"MigrateRole"`

	// 云数据库实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 云数据库名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 资源所属账号 为空值或self(表示本账号内资源)、other(表示其他账户资源)
	AccountMode *string `json:"AccountMode,omitnil" name:"AccountMode"`

	// 临时密钥Id，若为跨账号资源此项必填
	TmpSecretId *string `json:"TmpSecretId,omitnil" name:"TmpSecretId"`

	// 临时密钥Key，若为跨账号资源此项必填
	TmpSecretKey *string `json:"TmpSecretKey,omitnil" name:"TmpSecretKey"`

	// 临时密钥Token，若为跨账号资源此项必填
	TmpToken *string `json:"TmpToken,omitnil" name:"TmpToken"`
}

func (r *DescribeMigrateDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseType")
	delete(f, "MigrateRole")
	delete(f, "InstanceId")
	delete(f, "InstanceName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "AccountMode")
	delete(f, "TmpSecretId")
	delete(f, "TmpSecretKey")
	delete(f, "TmpToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrateDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrateDBInstancesResponseParams struct {
	// 符合筛选条件的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*MigrateDBItem `json:"Instances,omitnil" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMigrateDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrateDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeMigrateDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrateDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationCheckJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DescribeMigrationCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DescribeMigrationCheckJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationCheckJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationCheckJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationCheckJobResponseParams struct {
	// 校验任务执行状态，如：notStarted(未开始)、running(校验中)、failed(校验任务失败)、success(任务成功)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 校验任务结果输出简要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefMsg *string `json:"BriefMsg,omitnil" name:"BriefMsg"`

	// 检查步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo []*CheckStep `json:"StepInfo,omitnil" name:"StepInfo"`

	// 校验结果，如：checkPass(校验通过)、checkNotPass(校验未通过)
	CheckFlag *string `json:"CheckFlag,omitnil" name:"CheckFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMigrationCheckJobResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationCheckJobResponseParams `json:"Response"`
}

func (r *DescribeMigrationCheckJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationCheckJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDetailRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DescribeMigrationDetailRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DescribeMigrationDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationDetailResponseParams struct {
	// 数据迁移任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 数据迁移任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 任务创建(提交)时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务更新时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 任务开始执行时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 任务执行结束时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 迁移任务简要错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefMsg *string `json:"BriefMsg,omitnil" name:"BriefMsg"`

	// 任务状态，取值为：created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行中)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)、
	// pausing(暂停中)、
	// manualPaused(已暂停)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *MigrateAction `json:"Action,omitnil" name:"Action"`

	// 迁移执行过程信息，在校验阶段显示校验过程步骤信息，在迁移阶段会显示迁移步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo *MigrateDetailInfo `json:"StepInfo,omitnil" name:"StepInfo"`

	// 源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// 目标端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// 数据一致性校验结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTask *CompareTaskInfo `json:"CompareTask,omitnil" name:"CompareTask"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// 运行模式，取值如：immediate(表示立即运行)、timed(表示定时运行)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如："2006-01-02 15:04:05"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectRunTime *string `json:"ExpectRunTime,omitnil" name:"ExpectRunTime"`

	// 迁移选项，描述任务如何执行迁移等一系列配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// 校验任务运行详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckStepInfo *CheckStepInfo `json:"CheckStepInfo,omitnil" name:"CheckStepInfo"`

	// 描述计费相关的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeInfo *TradeInfo `json:"TradeInfo,omitnil" name:"TradeInfo"`

	// 任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo []*ErrorInfoItem `json:"ErrorInfo,omitnil" name:"ErrorInfo"`

	// 全量导出可重入标识：enum::"yes"/"no"。yes表示当前任务可重入、no表示当前任务处于全量导出且不可重入阶段；如果在该值为no时重启任务导出流程不支持断点续传
	DumperResumeCtrl *string `json:"DumperResumeCtrl,omitnil" name:"DumperResumeCtrl"`

	// 任务的限速信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitOption *RateLimitOption `json:"RateLimitOption,omitnil" name:"RateLimitOption"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMigrationDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationDetailResponseParams `json:"Response"`
}

func (r *DescribeMigrationDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationJobsRequestParams struct {
	// 数据迁移任务ID，如：dts-amm1jw5q
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 数据迁移任务状态，可取值包括：created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行中)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)
	Status []*string `json:"Status,omitnil" name:"Status"`

	// 源实例ID，格式如：cdb-c1nl9rpv
	SrcInstanceId *string `json:"SrcInstanceId,omitnil" name:"SrcInstanceId"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil" name:"SrcRegion"`

	// 源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	SrcDatabaseType []*string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	SrcAccessType []*string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// 目标实例ID，格式如：cdb-c1nl9rpv
	DstInstanceId *string `json:"DstInstanceId,omitnil" name:"DstInstanceId"`

	// 目标实例地域，如：ap-guangzhou
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 目标源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	DstDatabaseType []*string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// 目标实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	DstAccessType []*string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// 任务运行模式，值包括：immediate(立即运行)，timed(定时运行)
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 排序方式，可能取值为asc、desc，默认按照创建时间倒序
	OrderSeq *string `json:"OrderSeq,omitnil" name:"OrderSeq"`

	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

type DescribeMigrationJobsRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID，如：dts-amm1jw5q
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 数据迁移任务状态，可取值包括：created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行中)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)
	Status []*string `json:"Status,omitnil" name:"Status"`

	// 源实例ID，格式如：cdb-c1nl9rpv
	SrcInstanceId *string `json:"SrcInstanceId,omitnil" name:"SrcInstanceId"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitnil" name:"SrcRegion"`

	// 源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	SrcDatabaseType []*string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	SrcAccessType []*string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// 目标实例ID，格式如：cdb-c1nl9rpv
	DstInstanceId *string `json:"DstInstanceId,omitnil" name:"DstInstanceId"`

	// 目标实例地域，如：ap-guangzhou
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 目标源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	DstDatabaseType []*string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// 目标实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	DstAccessType []*string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// 任务运行模式，值包括：immediate(立即运行)，timed(定时运行)
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 排序方式，可能取值为asc、desc，默认按照创建时间倒序
	OrderSeq *string `json:"OrderSeq,omitnil" name:"OrderSeq"`

	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

func (r *DescribeMigrationJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "Status")
	delete(f, "SrcInstanceId")
	delete(f, "SrcRegion")
	delete(f, "SrcDatabaseType")
	delete(f, "SrcAccessType")
	delete(f, "DstInstanceId")
	delete(f, "DstRegion")
	delete(f, "DstDatabaseType")
	delete(f, "DstAccessType")
	delete(f, "RunMode")
	delete(f, "OrderSeq")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMigrationJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMigrationJobsResponseParams struct {
	// 迁移任务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 迁移任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobList []*JobItem `json:"JobList,omitnil" name:"JobList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeMigrationJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMigrationJobsResponseParams `json:"Response"`
}

func (r *DescribeMigrationJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMigrationJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyCheckSyncJobResultRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DescribeModifyCheckSyncJobResultRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DescribeModifyCheckSyncJobResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyCheckSyncJobResultRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModifyCheckSyncJobResultRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModifyCheckSyncJobResultResponseParams struct {
	// 校验任务执行状态，如：notStarted(未开始)、running(校验中)、failed(校验任务失败)、success(任务成功)
	Status *string `json:"Status,omitnil" name:"Status"`

	// 校验的步骤总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepCount *uint64 `json:"StepCount,omitnil" name:"StepCount"`

	// 当前所在步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepCur *uint64 `json:"StepCur,omitnil" name:"StepCur"`

	// 总体进度，范围为[0,100]	
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitnil" name:"Progress"`

	// 步骤详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfos []*StepInfo `json:"StepInfos,omitnil" name:"StepInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeModifyCheckSyncJobResultResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModifyCheckSyncJobResultResponseParams `json:"Response"`
}

func (r *DescribeModifyCheckSyncJobResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModifyCheckSyncJobResultResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSyncJobsRequestParams struct {
	// 同步任务id，如sync-werwfs23
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 同步任务名
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 排序字段，可以取值为CreateTime
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式，升序为ASC，降序为DESC，默认为CreateTime降序
	OrderSeq *string `json:"OrderSeq,omitnil" name:"OrderSeq"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回同步任务实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 状态集合，如Initialized,CheckPass,Running,ResumableErr,Stopped
	Status []*string `json:"Status,omitnil" name:"Status"`

	// 运行模式，如Immediate:立即运行，Timed:定时运行
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 任务类型，如mysql2mysql：msyql同步到mysql
	JobType *string `json:"JobType,omitnil" name:"JobType"`

	// 付费类型，PrePay：预付费，PostPay：后付费
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// tag
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

type DescribeSyncJobsRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id，如sync-werwfs23
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 同步任务名
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 排序字段，可以取值为CreateTime
	Order *string `json:"Order,omitnil" name:"Order"`

	// 排序方式，升序为ASC，降序为DESC，默认为CreateTime降序
	OrderSeq *string `json:"OrderSeq,omitnil" name:"OrderSeq"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 返回同步任务实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 状态集合，如Initialized,CheckPass,Running,ResumableErr,Stopped
	Status []*string `json:"Status,omitnil" name:"Status"`

	// 运行模式，如Immediate:立即运行，Timed:定时运行
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 任务类型，如mysql2mysql：msyql同步到mysql
	JobType *string `json:"JobType,omitnil" name:"JobType"`

	// 付费类型，PrePay：预付费，PostPay：后付费
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// tag
	TagFilters []*TagFilter `json:"TagFilters,omitnil" name:"TagFilters"`
}

func (r *DescribeSyncJobsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncJobsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	delete(f, "Order")
	delete(f, "OrderSeq")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Status")
	delete(f, "RunMode")
	delete(f, "JobType")
	delete(f, "PayMode")
	delete(f, "TagFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSyncJobsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSyncJobsResponseParams struct {
	// 任务数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 任务详情数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobList []*SyncJobInfo `json:"JobList,omitnil" name:"JobList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSyncJobsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSyncJobsResponseParams `json:"Response"`
}

func (r *DescribeSyncJobsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSyncJobsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMigrateJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DestroyMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DestroyMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroyMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroyMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroyMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *DestroyMigrateJobResponseParams `json:"Response"`
}

func (r *DestroyMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroyMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type DestroySyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *DestroySyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DestroySyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DestroySyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DestroySyncJobResponse struct {
	*tchttp.BaseResponse
	Response *DestroySyncJobResponseParams `json:"Response"`
}

func (r *DestroySyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DestroySyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailCheckItem struct {
	// 检查项的名称，如：源实例权限检查
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckItemName *string `json:"CheckItemName,omitnil" name:"CheckItemName"`

	// 检查项详细内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// pass(通过)，failed(失败), warning(校验有警告，但仍通过)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckResult *string `json:"CheckResult,omitnil" name:"CheckResult"`

	// 检查项失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitnil" name:"FailureReason"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitnil" name:"Solution"`

	// 运行报错日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorLog []*string `json:"ErrorLog,omitnil" name:"ErrorLog"`

	// 详细帮助的文档链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc []*string `json:"HelpDoc,omitnil" name:"HelpDoc"`

	// 跳过风险文案
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipInfo *string `json:"SkipInfo,omitnil" name:"SkipInfo"`
}

type DifferenceDetail struct {
	// 数据不一致的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 校验不一致的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DifferenceItem `json:"Items,omitnil" name:"Items"`
}

type DifferenceItem struct {
	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Db *string `json:"Db,omitnil" name:"Db"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitnil" name:"Table"`

	// 分块号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Chunk *int64 `json:"Chunk,omitnil" name:"Chunk"`

	// 源库数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcItem *string `json:"SrcItem,omitnil" name:"SrcItem"`

	// 目标库数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstItem *string `json:"DstItem,omitnil" name:"DstItem"`

	// 索引名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitnil" name:"IndexName"`

	// 索引下边界
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowerBoundary *string `json:"LowerBoundary,omitnil" name:"LowerBoundary"`

	// 索引上边界
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpperBoundary *string `json:"UpperBoundary,omitnil" name:"UpperBoundary"`

	// 对比消耗时间,单位为 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *float64 `json:"CostTime,omitnil" name:"CostTime"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedAt *string `json:"FinishedAt,omitnil" name:"FinishedAt"`
}

type DynamicOptions struct {
	// 所要同步的DML和DDL的选项，Insert(插入操作)、Update(更新操作)、Delete(删除操作)、DDL(结构同步)，PartialDDL(自定义,和DdlOptions一起起作用 )；必填、dts会用该值覆盖原有的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpTypes []*string `json:"OpTypes,omitnil" name:"OpTypes"`

	// DDL同步选项，具体描述要同步那些DDL; 当OpTypes取值PartialDDL时、字段不能为空；必填、dts会用该值覆盖原有的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlOptions []*DdlOption `json:"DdlOptions,omitnil" name:"DdlOptions"`

	// 冲突处理选项，ReportError(报错)、Ignore(忽略)、Cover(覆盖)、ConditionCover(条件覆盖); 目前目标端为kafka的链路不支持修改该配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictHandleType *string `json:"ConflictHandleType,omitnil" name:"ConflictHandleType"`

	// 冲突处理的详细选项，如条件覆盖中的条件行和条件操作；不能部分更新该选项的内部字段；有更新时、需要全量更新该字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictHandleOption *ConflictHandleOption `json:"ConflictHandleOption,omitnil" name:"ConflictHandleOption"`
}

type Endpoint struct {
	// 地域英文名，如：ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// tdsql mysql版的节点类型，枚举值为proxy、set
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil" name:"Role"`

	// 数据库内核类型，tdsql中用于区分不同内核：percona,mariadb,mysql
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbKernel *string `json:"DbKernel,omitnil" name:"DbKernel"`

	// 数据库实例ID，格式如：cdb-powiqx8q
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例的IP地址，接入类型为非cdb时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// 实例端口，接入类型为非cdb时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// 用户名，对于访问需要用户名密码认证的实例必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil" name:"User"`

	// 密码，对于访问需要用户名密码认证的实例必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitnil" name:"Password"`

	// 数据库名，数据库为cdwpg时，需要提供
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 私有网络ID，对于私有网络、专线、VPN的接入方式此项必填，格式如：vpc-92jblxto
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 私有网络下的子网ID，对于私有网络、专线、VPN的接入方式此项必填，格式如：subnet-3paxmkdz
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// CVM实例短ID，格式如：ins-olgl39y8，与云服务器控制台页面显示的实例ID相同。如果是CVM自建实例，需要传递此字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmInstanceId *string `json:"CvmInstanceId,omitnil" name:"CvmInstanceId"`

	// 专线网关ID，对于专线接入类型此项必填，格式如：dcg-0rxtqqxb
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqDcgId *string `json:"UniqDcgId,omitnil" name:"UniqDcgId"`

	// VPN网关ID，对于vpn接入类型此项必填，格式如：vpngw-9ghexg7q
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpnGwId *string `json:"UniqVpnGwId,omitnil" name:"UniqVpnGwId"`

	// 云联网ID，对于云联网接入类型此项必填，如：ccn-afp6kltc
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitnil" name:"CcnId"`

	// 云厂商类型，当实例为RDS实例时，填写为aliyun, 其他情况均填写others，默认为others
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supplier *string `json:"Supplier,omitnil" name:"Supplier"`

	// 数据库版本，当实例为RDS实例时才有效，其他实例忽略，格式如：5.6或者5.7，默认为5.6
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`

	// 实例所属账号，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitnil" name:"Account"`

	// 资源所属账号 为空或self(表示本账号内资源)、other(表示跨账号资源)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountMode *string `json:"AccountMode,omitnil" name:"AccountMode"`

	// 跨账号同步时的角色，只允许[a-zA-Z0-9\-\_]+，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountRole *string `json:"AccountRole,omitnil" name:"AccountRole"`

	// 外部角色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleExternalId *string `json:"RoleExternalId,omitnil" name:"RoleExternalId"`

	// 临时密钥Id，可通过获取联合身份临时访问凭证获取临时密钥https://cloud.tencent.com/document/product/1312/48195，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitnil" name:"TmpSecretId"`

	// 临时密钥Key，可通过获取联合身份临时访问凭证获取临时密钥https://cloud.tencent.com/document/product/1312/48195，，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitnil" name:"TmpSecretKey"`

	// 临时Token，可通过获取联合身份临时访问凭证获取临时密钥https://cloud.tencent.com/document/product/1312/48195，，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpToken *string `json:"TmpToken,omitnil" name:"TmpToken"`

	// 是否走加密传输、UnEncrypted表示不走加密传输，Encrypted表示走加密传输，默认UnEncrypted
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncryptConn *string `json:"EncryptConn,omitnil" name:"EncryptConn"`

	// 数据库所属网络环境，AccessType为云联网(ccn)时必填， UserIDC表示用户IDC、TencentVPC表示腾讯云VPC；
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseNetEnv *string `json:"DatabaseNetEnv,omitnil" name:"DatabaseNetEnv"`

	// 数据库为跨账号云联网下的实例时、表示云联网所属主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnOwnerUin *string `json:"CcnOwnerUin,omitnil" name:"CcnOwnerUin"`
}

type ErrInfo struct {
	// 错误原因
	Reason *string `json:"Reason,omitnil" name:"Reason"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitnil" name:"Solution"`
}

type ErrorInfoItem struct {
	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitnil" name:"Solution"`

	// 错误日志信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorLog *string `json:"ErrorLog,omitnil" name:"ErrorLog"`

	// 文档提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc *string `json:"HelpDoc,omitnil" name:"HelpDoc"`
}

// Predefined struct for user
type IsolateMigrateJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type IsolateMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *IsolateMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *IsolateMigrateJobResponseParams `json:"Response"`
}

func (r *IsolateMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type IsolateSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *IsolateSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsolateSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsolateSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type IsolateSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *IsolateSyncJobResponseParams `json:"Response"`
}

func (r *IsolateSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsolateSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JobItem struct {
	// 数据迁移任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 数据迁移任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 任务创建(提交)时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 任务更新时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// 任务开始执行时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 任务执行结束时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 迁移任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefMsg *string `json:"BriefMsg,omitnil" name:"BriefMsg"`

	// 任务状态，取值为：creating(创建中)、created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)、
	// pausing(暂停中)、
	// manualPaused(已暂停)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务运行模式，值包括：immediate(立即运行)，timed(定时运行)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如：2022-07-11 16:20:49
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectRunTime *string `json:"ExpectRunTime,omitnil" name:"ExpectRunTime"`

	// 任务操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *MigrateAction `json:"Action,omitnil" name:"Action"`

	// 迁移执行过程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo *MigrateDetailInfo `json:"StepInfo,omitnil" name:"StepInfo"`

	// 源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// 目标端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// 数据一致性校验结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTask *CompareTaskInfo `json:"CompareTask,omitnil" name:"CompareTask"`

	// 计费状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeInfo *TradeInfo `json:"TradeInfo,omitnil" name:"TradeInfo"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// 自动重试时间段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil" name:"AutoRetryTimeRangeMinutes"`

	// 全量导出可重入标识：enum::"yes"/"no"。yes表示当前任务可重入、no表示当前任务处于全量导出且不可重入阶段；如果在该值为no时重启任务导出流程不支持断点续传
	// 注意：此字段可能返回 null，表示取不到有效值。
	DumperResumeCtrl *string `json:"DumperResumeCtrl,omitnil" name:"DumperResumeCtrl"`
}

type KafkaOption struct {
	// 投递到kafka的数据类型，如Avro,Json,canal-pb,canal-json
	DataType *string `json:"DataType,omitnil" name:"DataType"`

	// 同步topic策略，如Single（集中投递到单topic）,Multi (自定义topic名称)
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// 用于存储ddl的topic
	DDLTopicName *string `json:"DDLTopicName,omitnil" name:"DDLTopicName"`

	// 单topic和自定义topic的描述
	TopicRules []*TopicRule `json:"TopicRules,omitnil" name:"TopicRules"`
}

type KeyValuePairOption struct {
	// 选项key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 选项value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type MigrateAction struct {
	// 任务的所有操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllAction []*string `json:"AllAction,omitnil" name:"AllAction"`

	// 任务在当前状态下允许的操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowedAction []*string `json:"AllowedAction,omitnil" name:"AllowedAction"`
}

type MigrateDBItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 实例Vip
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 实例Vport
	Vport *int64 `json:"Vport,omitnil" name:"Vport"`

	// 是否可以作为迁移对象，1-可以，0-不可以
	Usable *int64 `json:"Usable,omitnil" name:"Usable"`

	// 不可以作为迁移对象的原因
	Hint *string `json:"Hint,omitnil" name:"Hint"`
}

type MigrateDetailInfo struct {
	// 总步骤数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepAll *uint64 `json:"StepAll,omitnil" name:"StepAll"`

	// 当前步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNow *uint64 `json:"StepNow,omitnil" name:"StepNow"`

	// 主从差距，MB；只在任务正常，迁移或者同步的最后一步（追Binlog的阶段才有校），如果是非法值，返回-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitnil" name:"MasterSlaveDistance"`

	// 主从差距，秒；只在任务正常，迁移或者同步的最后一步（追Binlog的阶段才有校），如果是非法值，返回-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitnil" name:"SecondsBehindMaster"`

	// 步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo []*StepDetailInfo `json:"StepInfo,omitnil" name:"StepInfo"`
}

type MigrateOption struct {
	// 迁移对象选项，需要告知迁移服务迁移哪些库表对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseTable *DatabaseTableObject `json:"DatabaseTable,omitnil" name:"DatabaseTable"`

	// 迁移类型，full(全量迁移)，structure(结构迁移)，fullAndIncrement(全量加增量迁移)， 默认为fullAndIncrement
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrateType *string `json:"MigrateType,omitnil" name:"MigrateType"`

	// 数据一致性校验选项， 默认为不开启一致性校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *ConsistencyOption `json:"Consistency,omitnil" name:"Consistency"`

	// 是否迁移账号，yes(迁移账号)，no(不迁移账号)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMigrateAccount *bool `json:"IsMigrateAccount,omitnil" name:"IsMigrateAccount"`

	// 是否用源库Root账户覆盖目标库，值包括：false-不覆盖，true-覆盖，选择库表或者结构迁移时应该为false，注意只对旧版迁移有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOverrideRoot *bool `json:"IsOverrideRoot,omitnil" name:"IsOverrideRoot"`

	// 是否在迁移时设置目标库只读(仅对mysql有效)，true(设置只读)、false(不设置只读，默认此值)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDstReadOnly *bool `json:"IsDstReadOnly,omitnil" name:"IsDstReadOnly"`

	// 其他附加信息，对于特定库可设置额外参数，Redis可定义如下的参数: 
	// ["DstWriteMode":normal, 	目标库写入模式,可取值clearData(清空目标实例数据)、overwrite(以覆盖写的方式执行任务)、normal(跟正常流程一样，不做额外动作) 	"IsDstReadOnly":true, 	是否在迁移时设置目标库只读,true(设置只读)、false(不设置只读) 	"ClientOutputBufferHardLimit":512, 	从机缓冲区的硬性容量限制(MB) 	"ClientOutputBufferSoftLimit":512, 	从机缓冲区的软性容量限制(MB) 	"ClientOutputBufferPersistTime":60, 从机缓冲区的软性限制持续时间(秒) 	"ReplBacklogSize":512, 	环形缓冲区容量限制(MB) 	"ReplTimeout":120，		复制超时时间(秒) ]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitnil" name:"ExtraAttr"`
}

// Predefined struct for user
type ModifyCompareTaskNameRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 一致性校验任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`
}

type ModifyCompareTaskNameRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 一致性校验任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`
}

func (r *ModifyCompareTaskNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCompareTaskNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompareTaskNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCompareTaskNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCompareTaskNameResponseParams `json:"Response"`
}

func (r *ModifyCompareTaskNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompareTaskRequestParams struct {
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， 默认为此项配置)、custom(自定义)，注意自定义对比对象必须是迁移对象的子集
	ObjectMode *string `json:"ObjectMode,omitnil" name:"ObjectMode"`

	// 对比对象，若CompareObjectMode取值为custom，则此项必填
	Objects *CompareObject `json:"Objects,omitnil" name:"Objects"`

	// 一致性校验选项
	Options *CompareOptions `json:"Options,omitnil" name:"Options"`
}

type ModifyCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， 默认为此项配置)、custom(自定义)，注意自定义对比对象必须是迁移对象的子集
	ObjectMode *string `json:"ObjectMode,omitnil" name:"ObjectMode"`

	// 对比对象，若CompareObjectMode取值为custom，则此项必填
	Objects *CompareObject `json:"Objects,omitnil" name:"Objects"`

	// 一致性校验选项
	Options *CompareOptions `json:"Options,omitnil" name:"Options"`
}

func (r *ModifyCompareTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	delete(f, "TaskName")
	delete(f, "ObjectMode")
	delete(f, "Objects")
	delete(f, "Options")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompareTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCompareTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCompareTaskResponseParams `json:"Response"`
}

func (r *ModifyCompareTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCompareTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateJobSpecRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 新实例规格大小，包括：micro、small、medium、large、xlarge、2xlarge
	NewInstanceClass *string `json:"NewInstanceClass,omitnil" name:"NewInstanceClass"`
}

type ModifyMigrateJobSpecRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 新实例规格大小，包括：micro、small、medium、large、xlarge、2xlarge
	NewInstanceClass *string `json:"NewInstanceClass,omitnil" name:"NewInstanceClass"`
}

func (r *ModifyMigrateJobSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "NewInstanceClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateJobSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateJobSpecResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMigrateJobSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateJobSpecResponseParams `json:"Response"`
}

func (r *ModifyMigrateJobSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateJobSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateNameRequestParams struct {
	// 迁移任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 修改后的迁移任务名
	JobName *string `json:"JobName,omitnil" name:"JobName"`
}

type ModifyMigrateNameRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 修改后的迁移任务名
	JobName *string `json:"JobName,omitnil" name:"JobName"`
}

func (r *ModifyMigrateNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "JobName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateNameResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMigrateNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateNameResponseParams `json:"Response"`
}

func (r *ModifyMigrateNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateRateLimitRequestParams struct {
	// 迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 迁移任务全量导出线程数、有效值为 1-16
	DumpThread *int64 `json:"DumpThread,omitnil" name:"DumpThread"`

	// 迁移全量导出的 Rps 限制、需要大于 0
	DumpRps *int64 `json:"DumpRps,omitnil" name:"DumpRps"`

	// 迁移任务全量导入线程数、有效值为 1-16
	LoadThread *int64 `json:"LoadThread,omitnil" name:"LoadThread"`

	// 迁移任务增量导入线程数、有效值为 1-128
	SinkerThread *int64 `json:"SinkerThread,omitnil" name:"SinkerThread"`

	// 全量导入Rps限制
	LoadRps *int64 `json:"LoadRps,omitnil" name:"LoadRps"`
}

type ModifyMigrateRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 迁移任务全量导出线程数、有效值为 1-16
	DumpThread *int64 `json:"DumpThread,omitnil" name:"DumpThread"`

	// 迁移全量导出的 Rps 限制、需要大于 0
	DumpRps *int64 `json:"DumpRps,omitnil" name:"DumpRps"`

	// 迁移任务全量导入线程数、有效值为 1-16
	LoadThread *int64 `json:"LoadThread,omitnil" name:"LoadThread"`

	// 迁移任务增量导入线程数、有效值为 1-128
	SinkerThread *int64 `json:"SinkerThread,omitnil" name:"SinkerThread"`

	// 全量导入Rps限制
	LoadRps *int64 `json:"LoadRps,omitnil" name:"LoadRps"`
}

func (r *ModifyMigrateRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "DumpThread")
	delete(f, "DumpRps")
	delete(f, "LoadThread")
	delete(f, "SinkerThread")
	delete(f, "LoadRps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMigrateRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateRateLimitResponseParams `json:"Response"`
}

func (r *ModifyMigrateRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateRuntimeAttributeRequestParams struct {
	// 迁移任务id，如：dts-2rgv0f09
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 需要修改的属性，此结构设计为通用结构，用于屏蔽多个业务的定制属性。<br>例如对于Redis:<br>{<br>	 "Key": "DstWriteMode",	//目标库写入模式<br> 	"Value": "normal"	          //clearData(清空目标实例数据)、overwrite(以覆盖写的方式执行任务)、normal(跟正常流程一样，不做额外动作，默认为此值) <br>},<br>{<br/>	 "Key": "IsDstReadOnly",	//是否在迁移时设置目标库只读<br/> 	"Value": "true"	          //true(设置只读)、false(不设置只读) <br/>} 
	OtherOptions []*KeyValuePairOption `json:"OtherOptions,omitnil" name:"OtherOptions"`
}

type ModifyMigrateRuntimeAttributeRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务id，如：dts-2rgv0f09
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 需要修改的属性，此结构设计为通用结构，用于屏蔽多个业务的定制属性。<br>例如对于Redis:<br>{<br>	 "Key": "DstWriteMode",	//目标库写入模式<br> 	"Value": "normal"	          //clearData(清空目标实例数据)、overwrite(以覆盖写的方式执行任务)、normal(跟正常流程一样，不做额外动作，默认为此值) <br>},<br>{<br/>	 "Key": "IsDstReadOnly",	//是否在迁移时设置目标库只读<br/> 	"Value": "true"	          //true(设置只读)、false(不设置只读) <br/>} 
	OtherOptions []*KeyValuePairOption `json:"OtherOptions,omitnil" name:"OtherOptions"`
}

func (r *ModifyMigrateRuntimeAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateRuntimeAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "OtherOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrateRuntimeAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrateRuntimeAttributeResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMigrateRuntimeAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrateRuntimeAttributeResponseParams `json:"Response"`
}

func (r *ModifyMigrateRuntimeAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrateRuntimeAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 运行模式，取值如：immediate(表示立即运行)、timed(表示定时运行)
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 迁移任务配置选项，描述任务如何执行迁移等一系列配置信息；字段下的RateLimitOption不可配置、如果需要修改任务的限速信息、请在任务运行后通过ModifyMigrateRateLimit接口修改
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// 源实例信息
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// 目标实例信息
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// 迁移任务名称，最大长度128
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitnil" name:"ExpectRunTime"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// 自动重试的时间段、可设置5至720分钟、0表示不重试
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil" name:"AutoRetryTimeRangeMinutes"`
}

type ModifyMigrationJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 运行模式，取值如：immediate(表示立即运行)、timed(表示定时运行)
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 迁移任务配置选项，描述任务如何执行迁移等一系列配置信息；字段下的RateLimitOption不可配置、如果需要修改任务的限速信息、请在任务运行后通过ModifyMigrateRateLimit接口修改
	MigrateOption *MigrateOption `json:"MigrateOption,omitnil" name:"MigrateOption"`

	// 源实例信息
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// 目标实例信息
	DstInfo *DBEndpointInfo `json:"DstInfo,omitnil" name:"DstInfo"`

	// 迁移任务名称，最大长度128
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitnil" name:"ExpectRunTime"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// 自动重试的时间段、可设置5至720分钟、0表示不重试
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil" name:"AutoRetryTimeRangeMinutes"`
}

func (r *ModifyMigrationJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "RunMode")
	delete(f, "MigrateOption")
	delete(f, "SrcInfo")
	delete(f, "DstInfo")
	delete(f, "JobName")
	delete(f, "ExpectRunTime")
	delete(f, "Tags")
	delete(f, "AutoRetryTimeRangeMinutes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyMigrationJobResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMigrationJobResponseParams `json:"Response"`
}

func (r *ModifyMigrationJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMigrationJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncJobConfigRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 修改后的同步对象
	DynamicObjects *Objects `json:"DynamicObjects,omitnil" name:"DynamicObjects"`

	// 修改后的同步任务选项
	DynamicOptions *DynamicOptions `json:"DynamicOptions,omitnil" name:"DynamicOptions"`
}

type ModifySyncJobConfigRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 修改后的同步对象
	DynamicObjects *Objects `json:"DynamicObjects,omitnil" name:"DynamicObjects"`

	// 修改后的同步任务选项
	DynamicOptions *DynamicOptions `json:"DynamicOptions,omitnil" name:"DynamicOptions"`
}

func (r *ModifySyncJobConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncJobConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "DynamicObjects")
	delete(f, "DynamicOptions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySyncJobConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncJobConfigResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySyncJobConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifySyncJobConfigResponseParams `json:"Response"`
}

func (r *ModifySyncJobConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncJobConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncRateLimitRequestParams struct {
	// 迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 同步任务全量导出线程数、有效值为 1-16
	DumpThread *int64 `json:"DumpThread,omitnil" name:"DumpThread"`

	// 同步任务全量导出的 Rps 限制、需要大于 0
	DumpRps *int64 `json:"DumpRps,omitnil" name:"DumpRps"`

	// 同步任务全量导入线程数、有效值为 1-16
	LoadThread *int64 `json:"LoadThread,omitnil" name:"LoadThread"`

	// 同步任务增量导入线程数、有效值为 1-128
	SinkerThread *int64 `json:"SinkerThread,omitnil" name:"SinkerThread"`

	// 同步任务全量导入的Rps
	LoadRps *int64 `json:"LoadRps,omitnil" name:"LoadRps"`
}

type ModifySyncRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 同步任务全量导出线程数、有效值为 1-16
	DumpThread *int64 `json:"DumpThread,omitnil" name:"DumpThread"`

	// 同步任务全量导出的 Rps 限制、需要大于 0
	DumpRps *int64 `json:"DumpRps,omitnil" name:"DumpRps"`

	// 同步任务全量导入线程数、有效值为 1-16
	LoadThread *int64 `json:"LoadThread,omitnil" name:"LoadThread"`

	// 同步任务增量导入线程数、有效值为 1-128
	SinkerThread *int64 `json:"SinkerThread,omitnil" name:"SinkerThread"`

	// 同步任务全量导入的Rps
	LoadRps *int64 `json:"LoadRps,omitnil" name:"LoadRps"`
}

func (r *ModifySyncRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "DumpThread")
	delete(f, "DumpRps")
	delete(f, "LoadThread")
	delete(f, "SinkerThread")
	delete(f, "LoadRps")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySyncRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySyncRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifySyncRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifySyncRateLimitResponseParams `json:"Response"`
}

func (r *ModifySyncRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySyncRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Objects struct {
	// 同步对象类型 Partial(部分对象)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 同步对象，当 Mode 为 Partial 时，不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*Database `json:"Databases,omitnil" name:"Databases"`

	// 高级对象类型，如function、procedure。注意：如果要迁移同步高级对象，此配置中应该包含对应的高级对象类型。当需要同步高级对象时，初始化类型必须包含结构初始化类型，即任务的Options.InitType字段值为Structure或Full
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedObjects []*string `json:"AdvancedObjects,omitnil" name:"AdvancedObjects"`

	// OnlineDDL类型，冗余字段不做配置用途
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnlineDDL *OnlineDDL `json:"OnlineDDL,omitnil" name:"OnlineDDL"`
}

type OnlineDDL struct {
	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type Options struct {
	// 同步初始化选项，Data(全量数据初始化)、Structure(结构初始化)、Full(全量数据且结构初始化，默认)、None(仅增量)
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitType *string `json:"InitType,omitnil" name:"InitType"`

	// 同名表的处理，ReportErrorAfterCheck(前置校验并报错，默认)、InitializeAfterDelete(删除并重新初始化)、ExecuteAfterIgnore(忽略并继续执行)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealOfExistSameTable *string `json:"DealOfExistSameTable,omitnil" name:"DealOfExistSameTable"`

	// 冲突处理选项，ReportError(报错，默认为该值)、Ignore(忽略)、Cover(覆盖)、ConditionCover(条件覆盖)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictHandleType *string `json:"ConflictHandleType,omitnil" name:"ConflictHandleType"`

	// 是否添加附加列
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddAdditionalColumn *bool `json:"AddAdditionalColumn,omitnil" name:"AddAdditionalColumn"`

	// 所要同步的DML和DDL的选项，Insert(插入操作)、Update(更新操作)、Delete(删除操作)、DDL(结构同步)， 不填（不选），PartialDDL(自定义,和DdlOptions一起起作用 )
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpTypes []*string `json:"OpTypes,omitnil" name:"OpTypes"`

	// 冲突处理的详细选项，如条件覆盖中的条件行和条件操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictHandleOption *ConflictHandleOption `json:"ConflictHandleOption,omitnil" name:"ConflictHandleOption"`

	// DDL同步选项，具体描述要同步那些DDL
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlOptions []*DdlOption `json:"DdlOptions,omitnil" name:"DdlOptions"`

	// kafka同步选项
	// 注意：此字段可能返回 null，表示取不到有效值。
	KafkaOption *KafkaOption `json:"KafkaOption,omitnil" name:"KafkaOption"`

	// 任务限速信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitOption *RateLimitOption `json:"RateLimitOption,omitnil" name:"RateLimitOption"`

	// 自动重试的时间窗口设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil" name:"AutoRetryTimeRangeMinutes"`
}

// Predefined struct for user
type PauseMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type PauseMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *PauseMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PauseMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *PauseMigrateJobResponseParams `json:"Response"`
}

func (r *PauseMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type PauseSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *PauseSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PauseSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *PauseSyncJobResponseParams `json:"Response"`
}

func (r *PauseSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProcessProgress struct {
	// 步骤的状态， 包括：notStarted(未开始)、running(运行中)、success(成功)、failed(失败)等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 进度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitnil" name:"Percent"`

	// 总的步骤数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepAll *uint64 `json:"StepAll,omitnil" name:"StepAll"`

	// 当前进行的步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNow *uint64 `json:"StepNow,omitnil" name:"StepNow"`

	// 当前步骤输出提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Steps []*StepDetailInfo `json:"Steps,omitnil" name:"Steps"`
}

type ProcessStepTip struct {
	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitnil" name:"Solution"`

	// 文档提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc *string `json:"HelpDoc,omitnil" name:"HelpDoc"`
}

type RateLimitOption struct {
	// 当前生效的全量导出线程数，配置任务时可调整该字段值，注意：如果不设置或设置为0则表示保持当前值，最大值为16
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentDumpThread *int64 `json:"CurrentDumpThread,omitnil" name:"CurrentDumpThread"`

	// 默认的全量导出线程数，该字段仅在出参有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDumpThread *int64 `json:"DefaultDumpThread,omitnil" name:"DefaultDumpThread"`

	// 当前生效的全量导出Rps，配置任务时可调整该字段值，注意：如果不设置或设置为0则表示保持当前值，最大值为50000000
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentDumpRps *int64 `json:"CurrentDumpRps,omitnil" name:"CurrentDumpRps"`

	// 默认的全量导出Rps，该字段仅在出参有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultDumpRps *int64 `json:"DefaultDumpRps,omitnil" name:"DefaultDumpRps"`

	// 当前生效的全量导入线程数，配置任务时可调整该字段值，注意：如果不设置或设置为0则表示保持当前值，最大值为16
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentLoadThread *int64 `json:"CurrentLoadThread,omitnil" name:"CurrentLoadThread"`

	// 默认的全量导入线程数，该字段仅在出参有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultLoadThread *int64 `json:"DefaultLoadThread,omitnil" name:"DefaultLoadThread"`

	// 当前生效的全量导入Rps，配置任务时可调整该字段值，注意：如果不设置或设置为0则表示保持当前值，最大值为50000000	
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentLoadRps *int64 `json:"CurrentLoadRps,omitnil" name:"CurrentLoadRps"`

	// 默认的全量导入Rps，该字段仅在出参有意义	
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultLoadRps *int64 `json:"DefaultLoadRps,omitnil" name:"DefaultLoadRps"`

	// 当前生效的增量导入线程数，配置任务时可调整该字段值，注意：如果不设置或设置为0则表示保持当前值，最大值为128
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentSinkerThread *int64 `json:"CurrentSinkerThread,omitnil" name:"CurrentSinkerThread"`

	// 默认的增量导入线程数，该字段仅在出参有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultSinkerThread *int64 `json:"DefaultSinkerThread,omitnil" name:"DefaultSinkerThread"`

	// enum:"no"/"yes"、no表示用户未设置过限速、yes表示设置过限速，该字段仅在出参有意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasUserSetRateLimit *string `json:"HasUserSetRateLimit,omitnil" name:"HasUserSetRateLimit"`
}

// Predefined struct for user
type RecoverMigrateJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type RecoverMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *RecoverMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecoverMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *RecoverMigrateJobResponseParams `json:"Response"`
}

func (r *RecoverMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverSyncJobRequestParams struct {
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type RecoverSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *RecoverSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RecoverSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RecoverSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RecoverSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *RecoverSyncJobResponseParams `json:"Response"`
}

func (r *RecoverSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RecoverSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 任务规格
	NewInstanceClass *string `json:"NewInstanceClass,omitnil" name:"NewInstanceClass"`
}

type ResizeSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 任务规格
	NewInstanceClass *string `json:"NewInstanceClass,omitnil" name:"NewInstanceClass"`
}

func (r *ResizeSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "NewInstanceClass")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResizeSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResizeSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResizeSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ResizeSyncJobResponseParams `json:"Response"`
}

func (r *ResizeSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResizeSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 恢复任务的模式，目前的取值有：clearData 清空目标实例数据，overwrite 以覆盖写的方式执行任务，normal 跟正常流程一样，不做额外动作；注意，clearData、overwrite仅对redis生效，normal仅针对非redis链路生效
	ResumeOption *string `json:"ResumeOption,omitnil" name:"ResumeOption"`
}

type ResumeMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 恢复任务的模式，目前的取值有：clearData 清空目标实例数据，overwrite 以覆盖写的方式执行任务，normal 跟正常流程一样，不做额外动作；注意，clearData、overwrite仅对redis生效，normal仅针对非redis链路生效
	ResumeOption *string `json:"ResumeOption,omitnil" name:"ResumeOption"`
}

func (r *ResumeMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "ResumeOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResumeMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *ResumeMigrateJobResponseParams `json:"Response"`
}

func (r *ResumeMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type ResumeSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *ResumeSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResumeSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *ResumeSyncJobResponseParams `json:"Response"`
}

func (r *ResumeSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleItem struct {
	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// 迁移后的角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewRoleName *string `json:"NewRoleName,omitnil" name:"NewRoleName"`
}

// Predefined struct for user
type SkipCheckItemRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 需要跳过校验项的步骤id，需要通过DescribeMigrationCheckJob接口返回StepInfo[i].StepId字段获取，例如：["OptimizeCheck"]
	StepIds []*string `json:"StepIds,omitnil" name:"StepIds"`

	// 当出现外键依赖检查导致校验不通过时、可以通过该字段选择是否迁移外键依赖，当StepIds包含ConstraintCheck且该字段值为shield时表示不迁移外键依赖、当StepIds包含ConstraintCheck且值为migrate时表示迁移外键依赖
	ForeignKeyFlag *string `json:"ForeignKeyFlag,omitnil" name:"ForeignKeyFlag"`
}

type SkipCheckItemRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 需要跳过校验项的步骤id，需要通过DescribeMigrationCheckJob接口返回StepInfo[i].StepId字段获取，例如：["OptimizeCheck"]
	StepIds []*string `json:"StepIds,omitnil" name:"StepIds"`

	// 当出现外键依赖检查导致校验不通过时、可以通过该字段选择是否迁移外键依赖，当StepIds包含ConstraintCheck且该字段值为shield时表示不迁移外键依赖、当StepIds包含ConstraintCheck且值为migrate时表示迁移外键依赖
	ForeignKeyFlag *string `json:"ForeignKeyFlag,omitnil" name:"ForeignKeyFlag"`
}

func (r *SkipCheckItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipCheckItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StepIds")
	delete(f, "ForeignKeyFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SkipCheckItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SkipCheckItemResponseParams struct {
	// 跳过的提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SkipCheckItemResponse struct {
	*tchttp.BaseResponse
	Response *SkipCheckItemResponseParams `json:"Response"`
}

func (r *SkipCheckItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipCheckItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SkipSyncCheckItemRequestParams struct {
	// 任务id，如：sync-4ddgid2
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 需要跳过校验项的步骤id，需要通过`DescribeCheckSyncJobResult`接口返回StepInfos[i].StepId字段获取，例如：["OptimizeCheck"]
	StepIds []*string `json:"StepIds,omitnil" name:"StepIds"`
}

type SkipSyncCheckItemRequest struct {
	*tchttp.BaseRequest
	
	// 任务id，如：sync-4ddgid2
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 需要跳过校验项的步骤id，需要通过`DescribeCheckSyncJobResult`接口返回StepInfos[i].StepId字段获取，例如：["OptimizeCheck"]
	StepIds []*string `json:"StepIds,omitnil" name:"StepIds"`
}

func (r *SkipSyncCheckItemRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipSyncCheckItemRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "StepIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SkipSyncCheckItemRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SkipSyncCheckItemResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SkipSyncCheckItemResponse struct {
	*tchttp.BaseResponse
	Response *SkipSyncCheckItemResponseParams `json:"Response"`
}

func (r *SkipSyncCheckItemResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SkipSyncCheckItemResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SkippedDetail struct {
	// 跳过的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 跳过校验的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SkippedItem `json:"Items,omitnil" name:"Items"`
}

type SkippedItem struct {
	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Db *string `json:"Db,omitnil" name:"Db"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitnil" name:"Table"`

	// 未发起检查的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil" name:"Reason"`
}

// Predefined struct for user
type StartCompareRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`
}

type StartCompareRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`
}

func (r *StartCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartCompareResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartCompareResponse struct {
	*tchttp.BaseResponse
	Response *StartCompareResponseParams `json:"Response"`
}

func (r *StartCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *StartMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StartMigrateJobResponseParams `json:"Response"`
}

func (r *StartMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartModifySyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type StartModifySyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *StartModifySyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartModifySyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartModifySyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartModifySyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartModifySyncJobResponse struct {
	*tchttp.BaseResponse
	Response *StartModifySyncJobResponseParams `json:"Response"`
}

func (r *StartModifySyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartModifySyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type StartSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *StartSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StartSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *StartSyncJobResponseParams `json:"Response"`
}

func (r *StartSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StepDetailInfo struct {
	// 步骤序列
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNo *uint64 `json:"StepNo,omitnil" name:"StepNo"`

	// 步骤展现名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepName *string `json:"StepName,omitnil" name:"StepName"`

	// 步骤英文标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepId *string `json:"StepId,omitnil" name:"StepId"`

	// 步骤状态:success(成功)、failed(失败)、running(执行中)、notStarted(未执行)、默认为notStarted
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 当前步骤开始的时间，格式为"yyyy-mm-dd hh:mm:ss"，该字段不存在或者为空是无意义 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 步骤错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepMessage *string `json:"StepMessage,omitnil" name:"StepMessage"`

	// 执行进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitnil" name:"Percent"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errors []*ProcessStepTip `json:"Errors,omitnil" name:"Errors"`

	// 告警提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warnings []*ProcessStepTip `json:"Warnings,omitnil" name:"Warnings"`
}

type StepInfo struct {
	// 步骤编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNo *uint64 `json:"StepNo,omitnil" name:"StepNo"`

	// 步骤名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepName *string `json:"StepName,omitnil" name:"StepName"`

	// 步骤标号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepId *string `json:"StepId,omitnil" name:"StepId"`

	// 当前步骤状态,可能返回有 notStarted(未开始)、running(校验中)、failed(校验任务失败)、finished(完成)、skipped(跳过)、paused(暂停)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 步骤开始时间，可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errors []*StepTip `json:"Errors,omitnil" name:"Errors"`

	// 警告信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warnings []*StepTip `json:"Warnings,omitnil" name:"Warnings"`

	// 当前步骤进度，范围为[0-100]，若为-1表示当前步骤不支持查看进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitnil" name:"Progress"`
}

type StepTip struct {
	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitnil" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 解决方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitnil" name:"Solution"`

	// 帮助文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc *string `json:"HelpDoc,omitnil" name:"HelpDoc"`

	// 当前步骤跳过信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipInfo *string `json:"SkipInfo,omitnil" name:"SkipInfo"`
}

// Predefined struct for user
type StopCompareRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`
}

type StopCompareRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitnil" name:"CompareTaskId"`
}

func (r *StopCompareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCompareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	delete(f, "CompareTaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopCompareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopCompareResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopCompareResponse struct {
	*tchttp.BaseResponse
	Response *StopCompareResponseParams `json:"Response"`
}

func (r *StopCompareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopCompareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *StopMigrateJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopMigrateJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopMigrateJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopMigrateJobResponse struct {
	*tchttp.BaseResponse
	Response *StopMigrateJobResponseParams `json:"Response"`
}

func (r *StopMigrateJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopMigrateJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

type StopSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitnil" name:"JobId"`
}

func (r *StopSyncJobRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSyncJobRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "JobId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type StopSyncJobResponse struct {
	*tchttp.BaseResponse
	Response *StopSyncJobResponseParams `json:"Response"`
}

func (r *StopSyncJobResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopSyncJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncDBEndpointInfos struct {
	// 数据库所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 实例网络接入类型，如：extranet(外网)、ipv6(公网ipv6)、cvm(云主机自建)、dcg(专线接入)、vpncloud(vpn接入的实例)、cdb(云数据库)、ccn(云联网)、intranet(自研上云)、vpc(私有网络)等，注意具体可选值依赖当前链路
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *string `json:"AccessType,omitnil" name:"AccessType"`

	// 实例数据库类型，如：mysql,redis,mongodb,postgresql,mariadb,percona 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseType *string `json:"DatabaseType,omitnil" name:"DatabaseType"`

	// 数据库信息。注意：如果数据类型为tdsqlmysql，此处Endpoint数组的顺序应该与set顺序对应，第一个分片（shardkey范围起始为0的分片）必须要输入在第一个位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info []*Endpoint `json:"Info,omitnil" name:"Info"`
}

type SyncDetailInfo struct {
	// 总步骤数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepAll *uint64 `json:"StepAll,omitnil" name:"StepAll"`

	// 当前步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNow *uint64 `json:"StepNow,omitnil" name:"StepNow"`

	// 总体进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitnil" name:"Progress"`

	// 当前步骤进度，范围为[0-100]，若为-1表示当前步骤不支持查看进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentStepProgress *int64 `json:"CurrentStepProgress,omitnil" name:"CurrentStepProgress"`

	// 同步两端数据量差距
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitnil" name:"MasterSlaveDistance"`

	// 同步两端时间差距
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitnil" name:"SecondsBehindMaster"`

	// 总体描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil" name:"Message"`

	// 详细步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfos []*StepInfo `json:"StepInfos,omitnil" name:"StepInfos"`

	// 不能发起一致性校验的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	CauseOfCompareDisable *string `json:"CauseOfCompareDisable,omitnil" name:"CauseOfCompareDisable"`

	// 任务的错误和解决方案信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrInfo *ErrInfo `json:"ErrInfo,omitnil" name:"ErrInfo"`
}

type SyncJobInfo struct {
	// 同步任务id，如：sync-btso140
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil" name:"JobId"`

	// 同步任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitnil" name:"JobName"`

	// 付款方式，PostPay(按量付费)、PrePay(包年包月)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitnil" name:"PayMode"`

	// 运行模式，Immediate(表示立即运行，默认为此项值)、Timed(表示定时运行)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunMode *string `json:"RunMode,omitnil" name:"RunMode"`

	// 期待运行时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectRunTime *string `json:"ExpectRunTime,omitnil" name:"ExpectRunTime"`

	// 支持的所有操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllActions []*string `json:"AllActions,omitnil" name:"AllActions"`

	// 当前状态能进行的操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Actions []*string `json:"Actions,omitnil" name:"Actions"`

	// 同步选项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Options *Options `json:"Options,omitnil" name:"Options"`

	// 同步库表对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Objects *Objects `json:"Objects,omitnil" name:"Objects"`

	// 任务规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 过期时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 源端地域，如：ap-guangzhou等
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcRegion *string `json:"SrcRegion,omitnil" name:"SrcRegion"`

	// 源端数据库类型，mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcDatabaseType *string `json:"SrcDatabaseType,omitnil" name:"SrcDatabaseType"`

	// 源端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcAccessType *string `json:"SrcAccessType,omitnil" name:"SrcAccessType"`

	// 源端信息，单节点数据库使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfo *Endpoint `json:"SrcInfo,omitnil" name:"SrcInfo"`

	// 枚举值：cluster、single。源库为单节点数据库使用single，多节点使用cluster
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcNodeType *string `json:"SrcNodeType,omitnil" name:"SrcNodeType"`

	// 源端信息，多节点数据库使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfos *SyncDBEndpointInfos `json:"SrcInfos,omitnil" name:"SrcInfos"`

	// 目标端地域，如：ap-guangzhou等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstRegion *string `json:"DstRegion,omitnil" name:"DstRegion"`

	// 目标端数据库类型，mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstDatabaseType *string `json:"DstDatabaseType,omitnil" name:"DstDatabaseType"`

	// 目标端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstAccessType *string `json:"DstAccessType,omitnil" name:"DstAccessType"`

	// 目标端信息，单节点数据库使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfo *Endpoint `json:"DstInfo,omitnil" name:"DstInfo"`

	// 枚举值：cluster、single。目标库为单节点数据库使用single，多节点使用cluster
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstNodeType *string `json:"DstNodeType,omitnil" name:"DstNodeType"`

	// 目标端信息，多节点数据库使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfos *SyncDBEndpointInfos `json:"DstInfos,omitnil" name:"DstInfos"`

	// 创建时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 开始时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil" name:"StartTime"`

	// 任务状态，UnInitialized(未初始化)、Initialized(已初始化)、Checking(校验中)、CheckPass(校验通过)、CheckNotPass(校验不通过)、ReadyRunning(准备运行)、Running(运行中)、Pausing(暂停中)、Paused(已暂停)、Stopping(停止中)、Stopped(已结束)、ResumableErr(任务错误)、Resuming(恢复中)、Failed(失败)、Released(已释放)、Resetting(重置中)、Unknown(未知)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 结束时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil" name:"EndTime"`

	// 标签相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitnil" name:"Tags"`

	// 同步任务运行步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *SyncDetailInfo `json:"Detail,omitnil" name:"Detail"`

	// 用于计费的状态，可能取值有：Normal(正常状态)、Resizing(变配中)、Renewing(续费中)、Isolating(隔离中)、Isolated(已隔离)、Offlining(下线中)、Offlined(已下线)、NotBilled(未计费)、Recovering(解隔离)、PostPay2Prepaying(按量计费转包年包月中)、PrePay2Postpaying(包年包月转按量计费中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeStatus *string `json:"TradeStatus,omitnil" name:"TradeStatus"`

	// 同步链路规格，如micro,small,medium,large
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceClass *string `json:"InstanceClass,omitnil" name:"InstanceClass"`

	// 自动续费标识，当PayMode值为PrePay则此项配置有意义，取值为：1（表示自动续费）、0（不自动续费）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenew *uint64 `json:"AutoRenew,omitnil" name:"AutoRenew"`

	// 下线时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`

	// 自动重试时间段设置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRetryTimeRangeMinutes *int64 `json:"AutoRetryTimeRangeMinutes,omitnil" name:"AutoRetryTimeRangeMinutes"`

	// 全量导出可重入标识：enum::"yes"/"no"。yes表示当前任务可重入、no表示当前任务处于全量导出且不可重入阶段；如果在该值为no时重启任务导出流程不支持断点续传
	// 注意：此字段可能返回 null，表示取不到有效值。
	DumperResumeCtrl *string `json:"DumperResumeCtrl,omitnil" name:"DumperResumeCtrl"`
}

type Table struct {
	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 新表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewTableName *string `json:"NewTableName,omitnil" name:"NewTableName"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterCondition *string `json:"FilterCondition,omitnil" name:"FilterCondition"`

	// 是否同步表中所有列，All：当前表下的所有列,Partial(ModifySyncJobConfig接口里的对应字段ColumnMode暂不支持Partial)：当前表下的部分列，通过填充Columns字段详细表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnMode *string `json:"ColumnMode,omitnil" name:"ColumnMode"`

	// 同步的列信息，当ColumnMode为Partial时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*Column `json:"Columns,omitnil" name:"Columns"`

	// 同步临时表，注意此配置与NewTableName互斥，只能使用其中一种。当配置的同步对象为表级别且TableEditMode为pt时此项有意义，针对pt-osc等工具在同步过程中产生的临时表进行同步，需要提前将可能的临时表配置在这里，否则不会同步任何临时表。示例，如要对t1进行pt-osc操作，此项配置应该为["\_t1\_new","\_t1\_old"]；如要对t1进行gh-ost操作，此项配置应该为["\_t1\_ghc","\_t1\_gho","\_t1\_del"]，pt-osc与gh-ost产生的临时表可同时配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpTables []*string `json:"TmpTables,omitnil" name:"TmpTables"`

	// 编辑表类型，rename(表映射)，pt(同步附加表)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableEditMode *string `json:"TableEditMode,omitnil" name:"TableEditMode"`
}

type TableItem struct {
	// 迁移的表名，大小写敏感
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil" name:"TableName"`

	// 迁移后的表名，当TableEditMode为rename时此项必填，注意此配置与TmpTables互斥，只能使用其中一种
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewTableName *string `json:"NewTableName,omitnil" name:"NewTableName"`

	// 迁移临时表，注意此配置与NewTableName互斥，只能使用其中一种。当配置的同步对象为表级别且TableEditMode为pt时此项有意义，针对pt-osc等工具在迁移过程中产生的临时表进行同步，需要提前将可能的临时表配置在这里，否则不会同步任何临时表。示例，如要对t1进行pt-osc操作，此项配置应该为["\_t1\_new","\_t1\_old"]；如要对t1进行gh-ost操作，此项配置应该为["\_t1\_ghc","\_t1\_gho","\_t1\_del"]，pt-osc与gh-ost产生的临时表可同时配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpTables []*string `json:"TmpTables,omitnil" name:"TmpTables"`

	// 编辑表类型，rename(表映射)，pt(同步附加表)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableEditMode *string `json:"TableEditMode,omitnil" name:"TableEditMode"`
}

type TagFilter struct {
	// 标签键值
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitnil" name:"TagValue"`
}

type TagItem struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type TopicRule struct {
	// topic名
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// topic分区策略，如 自定义topic：Random（随机投递），集中投递到单Topic：AllInPartitionZero（全部投递至partition0）、PartitionByTable(按表名分区)、PartitionByTableAndKey(按表名加主键分区)
	PartitionType *string `json:"PartitionType,omitnil" name:"PartitionType"`

	// 库名匹配规则，仅“自定义topic”生效，如Regular（正则匹配）, Default(不符合匹配规则的剩余库)，数组中必须有一项为‘Default’
	DbMatchMode *string `json:"DbMatchMode,omitnil" name:"DbMatchMode"`

	// 库名，仅“自定义topic”时，DbMatchMode=Regular生效
	DbName *string `json:"DbName,omitnil" name:"DbName"`

	// 表名匹配规则，仅“自定义topic”生效，如Regular（正则匹配）, Default(不符合匹配规则的剩余表)，数组中必须有一项为‘Default’
	TableMatchMode *string `json:"TableMatchMode,omitnil" name:"TableMatchMode"`

	// 表名，仅“自定义topic”时，TableMatchMode=Regular生效
	TableName *string `json:"TableName,omitnil" name:"TableName"`
}

type TradeInfo struct {
	// 交易订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitnil" name:"DealName"`

	// 上一次交易订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastDealName *string `json:"LastDealName,omitnil" name:"LastDealName"`

	// 实例规格，包括：micro、small、medium、large、xlarge、2xlarge等
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceClass *string `json:"InstanceClass,omitnil" name:"InstanceClass"`

	// 计费任务状态， normal(计费或待计费)、resizing(变配中)、reversing(冲正中，比较短暂的状态)、isolating(隔离中，比较短暂的状态)、isolated(已隔离)、offlining(下线中)、offlined(已下线)、notBilled(未计费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeStatus *string `json:"TradeStatus,omitnil" name:"TradeStatus"`

	// 到期时间，格式为"yyyy-mm-dd hh:mm:ss"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 下线时间，格式为"yyyy-mm-dd hh:mm:ss"
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTime *string `json:"OfflineTime,omitnil" name:"OfflineTime"`

	// 隔离时间，格式为"yyyy-mm-dd hh:mm:ss"
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitnil" name:"IsolateTime"`

	// 下线原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineReason *string `json:"OfflineReason,omitnil" name:"OfflineReason"`

	// 隔离原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateReason *string `json:"IsolateReason,omitnil" name:"IsolateReason"`

	// 付费类型，包括：postpay(后付费)、prepay(预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayType *string `json:"PayType,omitnil" name:"PayType"`

	// 任务计费类型，包括：billing(计费)、notBilling(不计费)、 promotions(促销活动中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingType *string `json:"BillingType,omitnil" name:"BillingType"`
}

type View struct {
	// view名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitnil" name:"ViewName"`

	// 预留字段、目前暂时不支持view的重命名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewViewName *string `json:"NewViewName,omitnil" name:"NewViewName"`
}

type ViewItem struct {
	// 视图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitnil" name:"ViewName"`

	// 迁移后的视图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewViewName *string `json:"NewViewName,omitnil" name:"NewViewName"`
}