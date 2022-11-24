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
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type CheckStep struct {
	// 步骤编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// 步骤Id， 如：ConnectDBCheck、VersionCheck、SrcPrivilegeCheck等，具体校验项和源目标实例相关
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// 步骤名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 此检查步骤的结果，pass(校验通过)、failed(校验失败)、notStarted(校验还未开始进行)、blocked(检验阻塞)、warning(校验有告警，但仍通过)
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepStatus *string `json:"StepStatus,omitempty" name:"StepStatus"`

	// 此检查步骤的错误消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepMessage *string `json:"StepMessage,omitempty" name:"StepMessage"`

	// 每个检查步骤里的具体检查项
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailCheckItems []*DetailCheckItem `json:"DetailCheckItems,omitempty" name:"DetailCheckItems"`

	// 是否已跳过
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasSkipped *bool `json:"HasSkipped,omitempty" name:"HasSkipped"`
}

type CheckStepInfo struct {
	// 任务开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`

	// 任务步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *ProcessProgress `json:"Progress,omitempty" name:"Progress"`
}

type CompareAbstractInfo struct {
	// 对比结论: same,different
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conclusion *string `json:"Conclusion,omitempty" name:"Conclusion"`

	// 任务状态: success,failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 总的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTables *uint64 `json:"TotalTables,omitempty" name:"TotalTables"`

	// 已校验的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckedTables *uint64 `json:"CheckedTables,omitempty" name:"CheckedTables"`

	// 不一致的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DifferentTables *uint64 `json:"DifferentTables,omitempty" name:"DifferentTables"`

	// 跳过校验的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkippedTables *uint64 `json:"SkippedTables,omitempty" name:"SkippedTables"`

	// 不一致的数据行数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DifferentRows *uint64 `json:"DifferentRows,omitempty" name:"DifferentRows"`
}

type CompareDetailInfo struct {
	// 数据不一致的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Difference *DifferenceDetail `json:"Difference,omitempty" name:"Difference"`

	// 跳过校验的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Skipped *SkippedDetail `json:"Skipped,omitempty" name:"Skipped"`
}

type CompareObject struct {
	// 迁移对象模式 all(所有迁移对象)，partial(部分对象迁移)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectMode *string `json:"ObjectMode,omitempty" name:"ObjectMode"`

	// 迁移对象库表配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectItems []*CompareObjectItem `json:"ObjectItems,omitempty" name:"ObjectItems"`
}

type CompareObjectItem struct {
	// 迁移的库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 数据库选择模式: all 为当前对象下的所有对象,partial 为部分对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 迁移的 schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 表选择模式: all 为当前对象下的所有表对象,partial 为部分表对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMode *string `json:"TableMode,omitempty" name:"TableMode"`

	// 用于一致性校验的表配置，当 TableMode 为 partial 时，需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*CompareTableItem `json:"Tables,omitempty" name:"Tables"`

	// 视图选择模式: all 为当前对象下的所有视图对象,partial 为部分视图对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewMode *string `json:"ViewMode,omitempty" name:"ViewMode"`

	// 用于一致性校验的视图配置，当 ViewMode 为 partial 时， 需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Views []*CompareViewItem `json:"Views,omitempty" name:"Views"`
}

type CompareTableItem struct {
	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

type CompareTaskInfo struct {
	// 一致性校验任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 一致性校验结果，包括：unstart(未启动)、running(校验中)、canceled(已终止)、failed(校验任务失败)、inconsistent(不一致)、consistent(一致)、notexist(不存在校验任务)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`
}

type CompareTaskItem struct {
	// 迁移任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 对比任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 对比任务状态, 可能的值：created - 创建完成；readyRun - 等待运行；running - 运行中；success - 成功；stopping - 结束中；failed - 失败；canceled - 已终止
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 对比任务配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *CompareObject `json:"Config,omitempty" name:"Config"`

	// 对比任务校验详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckProcess *ProcessProgress `json:"CheckProcess,omitempty" name:"CheckProcess"`

	// 对比任务运行详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareProcess *ProcessProgress `json:"CompareProcess,omitempty" name:"CompareProcess"`

	// 对比结果, 可能的值：same - 一致；different - 不一致；skipAll - 跳过
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conclusion *string `json:"Conclusion,omitempty" name:"Conclusion"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`

	// 任务启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartedAt *string `json:"StartedAt,omitempty" name:"StartedAt"`

	// 对比结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedAt *string `json:"FinishedAt,omitempty" name:"FinishedAt"`
}

type CompareViewItem struct {
	// 视图名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

// Predefined struct for user
type CompleteMigrateJobRequestParams struct {
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 完成任务的方式,仅支持旧版MySQL迁移任务。waitForSync-等待主从差距为0才停止,immediately-立即完成，不会等待主从差距一致。默认为waitForSync
	CompleteMode *string `json:"CompleteMode,omitempty" name:"CompleteMode"`
}

type CompleteMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 完成任务的方式,仅支持旧版MySQL迁移任务。waitForSync-等待主从差距为0才停止,immediately-立即完成，不会等待主从差距一致。默认为waitForSync
	CompleteMode *string `json:"CompleteMode,omitempty" name:"CompleteMode"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 源端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)、noProxy,注意具体可选值依赖当前链路
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源端信息
	SrcInfo *Endpoint `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)、noProxy,注意具体可选值依赖当前链路
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标端信息
	DstInfo *Endpoint `json:"DstInfo,omitempty" name:"DstInfo"`

	// 同步任务选项
	Options *Options `json:"Options,omitempty" name:"Options"`

	// 同步库表对象信息
	Objects *Objects `json:"Objects,omitempty" name:"Objects"`

	// 同步任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 运行模式，取值如：Immediate(表示立即运行，默认为此项值)、Timed(表示定时运行)
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 期待启动时间，当RunMode取值为Timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`
}

type ConfigureSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 源端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)、noProxy,注意具体可选值依赖当前链路
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源端信息
	SrcInfo *Endpoint `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标端接入类型，cdb(云数据库)、cvm(云主机自建)、vpc(私有网络)、extranet(外网)、vpncloud(vpn接入)、dcg(专线接入)、ccn(云联网)、intranet(自研上云)、noProxy,注意具体可选值依赖当前链路
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标端信息
	DstInfo *Endpoint `json:"DstInfo,omitempty" name:"DstInfo"`

	// 同步任务选项
	Options *Options `json:"Options,omitempty" name:"Options"`

	// 同步库表对象信息
	Objects *Objects `json:"Objects,omitempty" name:"Objects"`

	// 同步任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 运行模式，取值如：Immediate(表示立即运行，默认为此项值)、Timed(表示定时运行)
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 期待启动时间，当RunMode取值为Timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`
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
	delete(f, "SrcInfo")
	delete(f, "DstAccessType")
	delete(f, "DstInfo")
	delete(f, "Options")
	delete(f, "Objects")
	delete(f, "JobName")
	delete(f, "RunMode")
	delete(f, "ExpectRunTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ConfigureSyncJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ConfigureSyncJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	ConditionColumn *string `json:"ConditionColumn,omitempty" name:"ConditionColumn"`

	// 条件覆盖操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionOperator *string `json:"ConditionOperator,omitempty" name:"ConditionOperator"`

	// 条件覆盖优先级处理
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionOrderInSrcAndDst *string `json:"ConditionOrderInSrcAndDst,omitempty" name:"ConditionOrderInSrcAndDst"`
}

type ConsistencyOption struct {
	// 一致性检测类型: full(全量检测迁移对象)、noCheck(不检测)、notConfigured(未配置)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`
}

// Predefined struct for user
type CreateCheckSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type CreateCheckSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据对比任务名称，若为空则默认给CompareTaskId相同值
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， 默认为此项配置)，custom(自定义模式)
	ObjectMode *string `json:"ObjectMode,omitempty" name:"ObjectMode"`

	// 一致性对比对象配置
	Objects *CompareObject `json:"Objects,omitempty" name:"Objects"`
}

type CreateCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据对比任务名称，若为空则默认给CompareTaskId相同值
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， 默认为此项配置)，custom(自定义模式)
	ObjectMode *string `json:"ObjectMode,omitempty" name:"ObjectMode"`

	// 一致性对比对象配置
	Objects *CompareObject `json:"Objects,omitempty" name:"Objects"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCompareTaskResponseParams struct {
	// 数据对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type CreateMigrateCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 源实例数据库类型，mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 目标实例数据库类型，mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`

	// 目标实例地域，如：ap-guangzhou。注意，目标地域必须和API请求地域保持一致。
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 实例规格，包括：micro、small、medium、large、xlarge、2xlarge
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 购买数量，范围为[1,15]，默认为1
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 迁移服务名称，最大长度128
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
}

type CreateMigrationServiceRequest struct {
	*tchttp.BaseRequest
	
	// 源实例数据库类型，mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 目标实例数据库类型，mysql,redis,percona,mongodb,postgresql,sqlserver,mariadb
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`

	// 目标实例地域，如：ap-guangzhou。注意，目标地域必须和API请求地域保持一致。
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 实例规格，包括：micro、small、medium、large、xlarge、2xlarge
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 购买数量，范围为[1,15]，默认为1
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 迁移服务名称，最大长度128
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
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
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateSyncJobRequestParams struct {
	// 支付类型，PrePay：包年包月  PostPay：按时按量
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 源端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源端数据库所在地域,如ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`

	// 目标端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标端数据库所在地域,如ap-guangzhou
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 同步任务规格，Standard:标准版
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 无
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 同步任务数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 自动续费标识
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 同步链路规格
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 同步链路名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 创建类似任务的现有任务Id
	ExistedJobId *string `json:"ExistedJobId,omitempty" name:"ExistedJobId"`
}

type CreateSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 支付类型，PrePay：包年包月  PostPay：按时按量
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 源端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源端数据库所在地域,如ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`

	// 目标端数据库类型,如mysql,cynosdbmysql,tdapg,tdpg,tdsqlmysql等
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标端数据库所在地域,如ap-guangzhou
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 同步任务规格，Standard:标准版
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 无
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 同步任务数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 自动续费标识
	AutoRenew *uint64 `json:"AutoRenew,omitempty" name:"AutoRenew"`

	// 同步链路规格
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 同步链路名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 创建类似任务的现有任务Id
	ExistedJobId *string `json:"ExistedJobId,omitempty" name:"ExistedJobId"`
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
	JobIds []*string `json:"JobIds,omitempty" name:"JobIds"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	Region *string `json:"Region,omitempty" name:"Region"`

	// 实例网络接入类型，如：extranet(外网)、ipv6(公网ipv6)、cvm(云主机自建)、dcg(专线接入)、vpncloud(vpn接入的实例)、cdb(云数据库)、ccn(云联网)、intranet(自研上云)、vpc(私有网络)等，注意具体可选值依赖当前链路
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`

	// 实例数据库类型，如：mysql,redis,mongodb,postgresql,mariadb,percona 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`

	// 节点类型，为空或者"simple":表示普通节点，"cluster": 集群节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 数据库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info []*DBInfo `json:"Info,omitempty" name:"Info"`

	// 实例服务提供商，如:"aliyun","others"
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`

	// MongoDB可定义如下的参数: 	['AuthDatabase':'admin', 
	// 'AuthFlag': "1",	'AuthMechanism':"SCRAM-SHA-1"]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
}

type DBInfo struct {
	// 表示节点角色，针对分布式数据库，如mongodb中的mongos节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitempty" name:"Role"`

	// 内核版本，针对mariadb的不同内核版本等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbKernel *string `json:"DbKernel,omitempty" name:"DbKernel"`

	// 实例的IP地址，对于公网、专线、VPN、云联网、自研上云、VPC等接入方式此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitempty" name:"Host"`

	// 实例的端口，对于公网、云主机自建、专线、VPN、云联网、自研上云、VPC等接入方式此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 实例的用户名
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitempty" name:"User"`

	// 实例的密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// CVM实例短ID，格式如：ins-olgl39y8；与云服务器控制台页面显示的实例ID相同；如果接入类型为云主机自建的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`

	// VPN网关ID，格式如：vpngw-9ghexg7q；如果接入类型为vpncloud的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`

	// 专线网关ID，格式如：dcg-0rxtqqxb；如果接入类型为专线接入的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`

	// 数据库实例ID，格式如：cdb-powiqx8q；如果接入类型为云数据库的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云联网ID，如：ccn-afp6kltc 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnGwId *string `json:"CcnGwId,omitempty" name:"CcnGwId"`

	// 私有网络ID，格式如：vpc-92jblxto；如果接入类型为vpc、vpncloud、ccn、dcg的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络下的子网ID，格式如：subnet-3paxmkdz；如果接入类型为vpc、vpncloud、ccn、dcg的方式，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 数据库版本，当实例为RDS实例时才有效，格式如：5.6或者5.7，默认为5.6
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 实例所属账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitempty" name:"Account"`

	// 跨账号迁移时的角色,只允许[a-zA-Z0-9\-\_]+
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountRole *string `json:"AccountRole,omitempty" name:"AccountRole"`

	// 资源所属账号 为空或self(表示本账号内资源)、other(表示其他账户资源)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountMode *string `json:"AccountMode,omitempty" name:"AccountMode"`

	// 临时密钥Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时Token
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpToken *string `json:"TmpToken,omitempty" name:"TmpToken"`
}

type DBItem struct {
	// 需要迁移或同步的库名，当ObjectMode为partial时，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 迁移或同步后的库名，默认与源库相同
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewDbName *string `json:"NewDbName,omitempty" name:"NewDbName"`

	// 迁移或同步的 schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 迁移或同步后的 schema name
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewSchemaName *string `json:"NewSchemaName,omitempty" name:"NewSchemaName"`

	// DB选择模式: all(为当前对象下的所有对象)，partial(部分对象)，当ObjectMode为partial时，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DBMode *string `json:"DBMode,omitempty" name:"DBMode"`

	// schema选择模式: all(为当前对象下的所有对象)，partial(部分对象)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaMode *string `json:"SchemaMode,omitempty" name:"SchemaMode"`

	// 表选择模式: all(为当前对象下的所有对象)，partial(部分对象)，当DBMode为partial时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMode *string `json:"TableMode,omitempty" name:"TableMode"`

	// 表图对象集合，当 TableMode 为 partial 时，此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*TableItem `json:"Tables,omitempty" name:"Tables"`

	// 视图选择模式: all 为当前对象下的所有视图对象,partial 为部分视图对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewMode *string `json:"ViewMode,omitempty" name:"ViewMode"`

	// 视图对象集合，当 ViewMode 为 partial 时， 此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Views []*ViewItem `json:"Views,omitempty" name:"Views"`

	// postgresql独有参数，角色选择模式: all 为当前对象下的所有角色对象,partial 为部分角色对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleMode *string `json:"RoleMode,omitempty" name:"RoleMode"`

	// postgresql独有参数，当 RoleMode 为 partial 时， 此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*RoleItem `json:"Roles,omitempty" name:"Roles"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionMode *string `json:"FunctionMode,omitempty" name:"FunctionMode"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerMode *string `json:"TriggerMode,omitempty" name:"TriggerMode"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventMode *string `json:"EventMode,omitempty" name:"EventMode"`

	// 选择要同步的模式，partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureMode *string `json:"ProcedureMode,omitempty" name:"ProcedureMode"`

	// FunctionMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Functions []*string `json:"Functions,omitempty" name:"Functions"`

	// ProcedureMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedures []*string `json:"Procedures,omitempty" name:"Procedures"`

	// EventMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Events []*string `json:"Events,omitempty" name:"Events"`

	// TriggerMode取值为partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Triggers []*string `json:"Triggers,omitempty" name:"Triggers"`
}

type Database struct {
	// 需要迁移或同步的库名，当ObjectMode为Partial时，此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 迁移或同步后的库名，默认与源库相同
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewDbName *string `json:"NewDbName,omitempty" name:"NewDbName"`

	// DB选择模式: All(为当前对象下的所有对象)，Partial(部分对象)，当Mode为Partial时，此项必填。注意，高级对象的同步不依赖此值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbMode *string `json:"DbMode,omitempty" name:"DbMode"`

	// 迁移或同步的 schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`

	// 迁移或同步后的 schema name
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewSchemaName *string `json:"NewSchemaName,omitempty" name:"NewSchemaName"`

	// 表选择模式: All(为当前对象下的所有对象)，Partial(部分对象)，当DBMode为Partial时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMode *string `json:"TableMode,omitempty" name:"TableMode"`

	// 表图对象集合，当 TableMode 为 Partial 时，此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tables []*Table `json:"Tables,omitempty" name:"Tables"`

	// 视图选择模式: All 为当前对象下的所有视图对象,Partial 为部分视图对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewMode *string `json:"ViewMode,omitempty" name:"ViewMode"`

	// 视图对象集合，当 ViewMode 为 Partial 时， 此项需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Views []*View `json:"Views,omitempty" name:"Views"`

	// 选择要同步的模式，Partial为部分，all为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionMode *string `json:"FunctionMode,omitempty" name:"FunctionMode"`

	// FunctionMode取值为Partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Functions []*string `json:"Functions,omitempty" name:"Functions"`

	// 选择要同步的模式，Partial为部分，All为整选
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureMode *string `json:"ProcedureMode,omitempty" name:"ProcedureMode"`

	// ProcedureMode取值为Partial时需要填写
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedures []*string `json:"Procedures,omitempty" name:"Procedures"`
}

type DatabaseTableObject struct {
	// 迁移对象类型 all(全实例)，partial(部分对象)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectMode *string `json:"ObjectMode,omitempty" name:"ObjectMode"`

	// 迁移对象，当 ObjectMode 为 partial 时，不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*DBItem `json:"Databases,omitempty" name:"Databases"`

	// 高级对象类型，如trigger、function、procedure、event
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedObjects []*string `json:"AdvancedObjects,omitempty" name:"AdvancedObjects"`
}

type DdlOption struct {
	// ddl类型，如Database,Table,View,Index等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlObject *string `json:"DdlObject,omitempty" name:"DdlObject"`

	// ddl具体值，对于Database可取值[Create,Drop,Alter]<br>对于Table可取值[Create,Drop,Alter,Truncate,Rename]<br/>对于View可取值[Create,Drop]<br/>对于Index可取值[Create,Drop]
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlValue []*string `json:"DdlValue,omitempty" name:"DdlValue"`
}

// Predefined struct for user
type DeleteCompareTaskRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`
}

type DeleteCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DescribeCheckSyncJobResultRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	// 校验结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 步骤总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepCount *uint64 `json:"StepCount,omitempty" name:"StepCount"`

	// 当前所在步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepCur *uint64 `json:"StepCur,omitempty" name:"StepCur"`

	// 总体进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfos []*StepInfo `json:"StepInfos,omitempty" name:"StepInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 校验任务 Id
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 校验不一致结果的 limit
	DifferenceLimit *uint64 `json:"DifferenceLimit,omitempty" name:"DifferenceLimit"`

	// 不一致的 Offset
	DifferenceOffset *uint64 `json:"DifferenceOffset,omitempty" name:"DifferenceOffset"`

	// 搜索条件，不一致的库名
	DifferenceDB *string `json:"DifferenceDB,omitempty" name:"DifferenceDB"`

	// 搜索条件，不一致的表名
	DifferenceTable *string `json:"DifferenceTable,omitempty" name:"DifferenceTable"`

	// 未校验的 Limit
	SkippedLimit *uint64 `json:"SkippedLimit,omitempty" name:"SkippedLimit"`

	// 未校验的 Offset
	SkippedOffset *uint64 `json:"SkippedOffset,omitempty" name:"SkippedOffset"`

	// 搜索条件，未校验的库名
	SkippedDB *string `json:"SkippedDB,omitempty" name:"SkippedDB"`

	// 搜索条件，未校验的表名
	SkippedTable *string `json:"SkippedTable,omitempty" name:"SkippedTable"`
}

type DescribeCompareReportRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 校验任务 Id
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 校验不一致结果的 limit
	DifferenceLimit *uint64 `json:"DifferenceLimit,omitempty" name:"DifferenceLimit"`

	// 不一致的 Offset
	DifferenceOffset *uint64 `json:"DifferenceOffset,omitempty" name:"DifferenceOffset"`

	// 搜索条件，不一致的库名
	DifferenceDB *string `json:"DifferenceDB,omitempty" name:"DifferenceDB"`

	// 搜索条件，不一致的表名
	DifferenceTable *string `json:"DifferenceTable,omitempty" name:"DifferenceTable"`

	// 未校验的 Limit
	SkippedLimit *uint64 `json:"SkippedLimit,omitempty" name:"SkippedLimit"`

	// 未校验的 Offset
	SkippedOffset *uint64 `json:"SkippedOffset,omitempty" name:"SkippedOffset"`

	// 搜索条件，未校验的库名
	SkippedDB *string `json:"SkippedDB,omitempty" name:"SkippedDB"`

	// 搜索条件，未校验的表名
	SkippedTable *string `json:"SkippedTable,omitempty" name:"SkippedTable"`
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
	Abstract *CompareAbstractInfo `json:"Abstract,omitempty" name:"Abstract"`

	// 一致性校验详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *CompareDetailInfo `json:"Detail,omitempty" name:"Detail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 分页设置，表示每页显示多少条任务，默认为 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeCompareTasksRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 分页设置，表示每页显示多少条任务，默认为 20
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCompareTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCompareTasksResponseParams struct {
	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 一致性校验列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*CompareTaskItem `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`

	// 实例作为迁移的源还是目标,src(表示源)，dst(表示目标)
	MigrateRole *string `json:"MigrateRole,omitempty" name:"MigrateRole"`

	// 云数据库实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 资源所属账号 为空值或self(表示本账号内资源)、other(表示其他账户资源)
	AccountMode *string `json:"AccountMode,omitempty" name:"AccountMode"`

	// 临时密钥Id，若为跨账号资源此项必填
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥Key，若为跨账号资源此项必填
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时密钥Token，若为跨账号资源此项必填
	TmpToken *string `json:"TmpToken,omitempty" name:"TmpToken"`
}

type DescribeMigrateDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 数据库类型，如mysql
	DatabaseType *string `json:"DatabaseType,omitempty" name:"DatabaseType"`

	// 实例作为迁移的源还是目标,src(表示源)，dst(表示目标)
	MigrateRole *string `json:"MigrateRole,omitempty" name:"MigrateRole"`

	// 云数据库实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 云数据库名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 返回数量限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 资源所属账号 为空值或self(表示本账号内资源)、other(表示其他账户资源)
	AccountMode *string `json:"AccountMode,omitempty" name:"AccountMode"`

	// 临时密钥Id，若为跨账号资源此项必填
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥Key，若为跨账号资源此项必填
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时密钥Token，若为跨账号资源此项必填
	TmpToken *string `json:"TmpToken,omitempty" name:"TmpToken"`
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
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*MigrateDBItem `json:"Instances,omitempty" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DescribeMigrationCheckJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	Status *string `json:"Status,omitempty" name:"Status"`

	// 校验任务结果输出简要信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefMsg *string `json:"BriefMsg,omitempty" name:"BriefMsg"`

	// 检查步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo []*CheckStep `json:"StepInfo,omitempty" name:"StepInfo"`

	// 校验结果，如：checkPass(校验通过)、checkNotPass(校验未通过)
	CheckFlag *string `json:"CheckFlag,omitempty" name:"CheckFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DescribeMigrationDetailRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 任务创建(提交)时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务更新时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 任务开始执行时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务执行结束时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 迁移任务简要错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefMsg *string `json:"BriefMsg,omitempty" name:"BriefMsg"`

	// 任务状态，取值为：created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行中)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *MigrateAction `json:"Action,omitempty" name:"Action"`

	// 迁移执行过程信息，在校验阶段显示校验过程步骤信息，在迁移阶段会显示迁移步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo *MigrateDetailInfo `json:"StepInfo,omitempty" name:"StepInfo"`

	// 源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfo *DBEndpointInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 数据一致性校验结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTask *CompareTaskInfo `json:"CompareTask,omitempty" name:"CompareTask"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 运行模式，取值如：immediate(表示立即运行)、timed(表示定时运行)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如："2006-01-02 15:04:05"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`

	// 迁移选项，描述任务如何执行迁移等一系列配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 校验任务运行详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckStepInfo *CheckStepInfo `json:"CheckStepInfo,omitempty" name:"CheckStepInfo"`

	// 描述计费相关的信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeInfo *TradeInfo `json:"TradeInfo,omitempty" name:"TradeInfo"`

	// 任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo []*ErrorInfoItem `json:"ErrorInfo,omitempty" name:"ErrorInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 数据迁移任务状态，可取值包括：created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行中)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 源实例ID，格式如：cdb-c1nl9rpv
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`

	// 源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	SrcDatabaseType []*string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	SrcAccessType []*string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 目标实例ID，格式如：cdb-c1nl9rpv
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// 目标实例地域，如：ap-guangzhou
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 目标源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	DstDatabaseType []*string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	DstAccessType []*string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 任务运行模式，值包括：immediate(立即运行)，timed(定时运行)
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 排序方式，可能取值为asc、desc，默认按照创建时间倒序
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeMigrationJobsRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID，如：dts-amm1jw5q
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 数据迁移任务状态，可取值包括：created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行中)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 源实例ID，格式如：cdb-c1nl9rpv
	SrcInstanceId *string `json:"SrcInstanceId,omitempty" name:"SrcInstanceId"`

	// 源实例地域，如：ap-guangzhou
	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`

	// 源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	SrcDatabaseType []*string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	SrcAccessType []*string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 目标实例ID，格式如：cdb-c1nl9rpv
	DstInstanceId *string `json:"DstInstanceId,omitempty" name:"DstInstanceId"`

	// 目标实例地域，如：ap-guangzhou
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 目标源实例数据库类型，如：sqlserver,mysql,mongodb,redis,tendis,keewidb,clickhouse,cynosdbmysql,percona,tdsqlpercona,mariadb,tdsqlmysql,postgresql
	DstDatabaseType []*string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标实例接入类型，值包括：extranet(外网)、vpncloud(云vpn接入的实例)、dcg(专线接入的实例)、ccn(云联网接入的实例)、cdb(云上cdb实例)、cvm(cvm自建实例)
	DstAccessType []*string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 任务运行模式，值包括：immediate(立即运行)，timed(定时运行)
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 排序方式，可能取值为asc、desc，默认按照创建时间倒序
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// 返回实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 标签过滤
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 迁移任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobList []*JobItem `json:"JobList,omitempty" name:"JobList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type DescribeSyncJobsRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 同步任务名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 排序字段，可以取值为CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序方式，升序为ASC，降序为DESC，默认为CreateTime降序
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回同步任务实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 状态集合，如Initialized,CheckPass,Running,ResumableErr,Stopped
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 运行模式，如Immediate:立即运行，Timed:定时运行
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 任务类型，如mysql2mysql：msyql同步到mysql
	JobType *string `json:"JobType,omitempty" name:"JobType"`

	// 付费类型，PrePay：预付费，PostPay：后付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// tag
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
}

type DescribeSyncJobsRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 同步任务名
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 排序字段，可以取值为CreateTime
	Order *string `json:"Order,omitempty" name:"Order"`

	// 排序方式，升序为ASC，降序为DESC，默认为CreateTime降序
	OrderSeq *string `json:"OrderSeq,omitempty" name:"OrderSeq"`

	// 偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回同步任务实例数量，默认20，有效区间[1,100]
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 状态集合，如Initialized,CheckPass,Running,ResumableErr,Stopped
	Status []*string `json:"Status,omitempty" name:"Status"`

	// 运行模式，如Immediate:立即运行，Timed:定时运行
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 任务类型，如mysql2mysql：msyql同步到mysql
	JobType *string `json:"JobType,omitempty" name:"JobType"`

	// 付费类型，PrePay：预付费，PostPay：后付费
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// tag
	TagFilters []*TagFilter `json:"TagFilters,omitempty" name:"TagFilters"`
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
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 任务详情数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobList []*SyncJobInfo `json:"JobList,omitempty" name:"JobList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DestroyMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type DestroySyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	CheckItemName *string `json:"CheckItemName,omitempty" name:"CheckItemName"`

	// 检查项详细内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// pass(通过)，failed(失败), warning(校验有警告，但仍通过)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckResult *string `json:"CheckResult,omitempty" name:"CheckResult"`

	// 检查项失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailureReason *string `json:"FailureReason,omitempty" name:"FailureReason"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 运行报错日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorLog []*string `json:"ErrorLog,omitempty" name:"ErrorLog"`

	// 详细帮助的文档链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc []*string `json:"HelpDoc,omitempty" name:"HelpDoc"`

	// 跳过风险文案
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkipInfo *string `json:"SkipInfo,omitempty" name:"SkipInfo"`
}

type DifferenceDetail struct {
	// 数据不一致的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 校验不一致的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DifferenceItem `json:"Items,omitempty" name:"Items"`
}

type DifferenceItem struct {
	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Db *string `json:"Db,omitempty" name:"Db"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitempty" name:"Table"`

	// 分块号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Chunk *int64 `json:"Chunk,omitempty" name:"Chunk"`

	// 源库数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcItem *string `json:"SrcItem,omitempty" name:"SrcItem"`

	// 目标库数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstItem *string `json:"DstItem,omitempty" name:"DstItem"`

	// 索引名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IndexName *string `json:"IndexName,omitempty" name:"IndexName"`

	// 索引下边界
	// 注意：此字段可能返回 null，表示取不到有效值。
	LowerBoundary *string `json:"LowerBoundary,omitempty" name:"LowerBoundary"`

	// 索引上边界
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpperBoundary *string `json:"UpperBoundary,omitempty" name:"UpperBoundary"`

	// 对比消耗时间,单位为 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`

	// 完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishedAt *string `json:"FinishedAt,omitempty" name:"FinishedAt"`
}

type Endpoint struct {
	// 地域英文名，如：ap-guangzhou
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 数据库内核类型，tdsql中用于区分不同内核：percona,mariadb,mysql
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbKernel *string `json:"DbKernel,omitempty" name:"DbKernel"`

	// 数据库实例ID，格式如：cdb-powiqx8q
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例的IP地址，接入类型为非cdb时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 实例端口，接入类型为非cdb时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 用户名，对于访问需要用户名密码认证的实例必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitempty" name:"User"`

	// 密码，对于访问需要用户名密码认证的实例必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 数据库名，数据库为cdwpg时，需要提供
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 私有网络ID，对于私有网络、专线、VPN的接入方式此项必填，格式如：vpc-92jblxto
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络下的子网ID，对于私有网络、专线、VPN的接入方式此项必填，格式如：subnet-3paxmkdz
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// CVM实例短ID，格式如：ins-olgl39y8，与云服务器控制台页面显示的实例ID相同。如果是CVM自建实例，需要传递此字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	CvmInstanceId *string `json:"CvmInstanceId,omitempty" name:"CvmInstanceId"`

	// 专线网关ID，对于专线接入类型此项必填，格式如：dcg-0rxtqqxb
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqDcgId *string `json:"UniqDcgId,omitempty" name:"UniqDcgId"`

	// VPN网关ID，对于vpn接入类型此项必填，格式如：vpngw-9ghexg7q
	// 注意：此字段可能返回 null，表示取不到有效值。
	UniqVpnGwId *string `json:"UniqVpnGwId,omitempty" name:"UniqVpnGwId"`

	// 云联网ID，对于云联网接入类型此项必填，如：ccn-afp6kltc
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 云厂商类型，当实例为RDS实例时，填写为aliyun, 其他情况均填写others，默认为others
	// 注意：此字段可能返回 null，表示取不到有效值。
	Supplier *string `json:"Supplier,omitempty" name:"Supplier"`

	// 数据库版本，当实例为RDS实例时才有效，其他实例忽略，格式如：5.6或者5.7，默认为5.6
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`

	// 资源所属账号 为空或self(表示本账号内资源)、other(表示跨账号资源)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountMode *string `json:"AccountMode,omitempty" name:"AccountMode"`

	// 实例所属账号，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	Account *string `json:"Account,omitempty" name:"Account"`

	// 跨账号同步时的角色，只允许[a-zA-Z0-9\-\_]+，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountRole *string `json:"AccountRole,omitempty" name:"AccountRole"`

	// 临时密钥Id，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时密钥Key，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时Token，如果为跨账号实例此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpToken *string `json:"TmpToken,omitempty" name:"TmpToken"`
}

type ErrorInfoItem struct {
	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 错误日志信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorLog *string `json:"ErrorLog,omitempty" name:"ErrorLog"`

	// 文档提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

// Predefined struct for user
type IsolateMigrateJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type IsolateMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type IsolateSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 数据迁移任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 任务创建(提交)时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 任务更新时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 任务开始执行时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务执行结束时间，格式为 yyyy-mm-dd hh:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 迁移任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BriefMsg *string `json:"BriefMsg,omitempty" name:"BriefMsg"`

	// 任务状态，取值为：creating(创建中)、created(创建完成)、checking(校验中)、checkPass(校验通过)、checkNotPass(校验不通过)、readyRun(准备运行)、running(任务运行)、readyComplete(准备完成)、success(任务成功)、failed(任务失败)、stopping(中止中)、completing(完成中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务运行模式，值包括：immediate(立即运行)，timed(定时运行)
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如：2022-07-11 16:20:49
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`

	// 任务操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *MigrateAction `json:"Action,omitempty" name:"Action"`

	// 迁移执行过程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo *MigrateDetailInfo `json:"StepInfo,omitempty" name:"StepInfo"`

	// 源实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfo *DBEndpointInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 数据一致性校验结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareTask *CompareTaskInfo `json:"CompareTask,omitempty" name:"CompareTask"`

	// 计费状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeInfo *TradeInfo `json:"TradeInfo,omitempty" name:"TradeInfo"`

	// 标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
}

type KeyValuePairOption struct {
	// 选项key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 选项value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`
}

type MigrateAction struct {
	// 任务的所有操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllAction []*string `json:"AllAction,omitempty" name:"AllAction"`

	// 任务在当前状态下允许的操作列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllowedAction []*string `json:"AllowedAction,omitempty" name:"AllowedAction"`
}

type MigrateDBItem struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例Vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// 实例Vport
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 是否可以作为迁移对象，1-可以，0-不可以
	Usable *int64 `json:"Usable,omitempty" name:"Usable"`

	// 不可以作为迁移对象的原因
	Hint *string `json:"Hint,omitempty" name:"Hint"`
}

type MigrateDetailInfo struct {
	// 总步骤数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepAll *uint64 `json:"StepAll,omitempty" name:"StepAll"`

	// 当前步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNow *uint64 `json:"StepNow,omitempty" name:"StepNow"`

	// 主从差距，MB；只在任务正常，迁移或者同步的最后一步（追Binlog的阶段才有校），如果是非法值，返回-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// 主从差距，秒；只在任务正常，迁移或者同步的最后一步（追Binlog的阶段才有校），如果是非法值，返回-1
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// 步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfo []*StepDetailInfo `json:"StepInfo,omitempty" name:"StepInfo"`
}

type MigrateOption struct {
	// 迁移对象选项，需要告知迁移服务迁移哪些库表对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseTable *DatabaseTableObject `json:"DatabaseTable,omitempty" name:"DatabaseTable"`

	// 迁移类型，full(全量迁移)，structure(结构迁移)，fullAndIncrement(全量加增量迁移)， 默认为fullAndIncrement
	// 注意：此字段可能返回 null，表示取不到有效值。
	MigrateType *string `json:"MigrateType,omitempty" name:"MigrateType"`

	// 数据一致性校验选项， 默认为不开启一致性校验
	// 注意：此字段可能返回 null，表示取不到有效值。
	Consistency *ConsistencyOption `json:"Consistency,omitempty" name:"Consistency"`

	// 是否迁移账号，yes(迁移账号)，no(不迁移账号)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMigrateAccount *bool `json:"IsMigrateAccount,omitempty" name:"IsMigrateAccount"`

	// 是否用源库Root账户覆盖目标库，值包括：false-不覆盖，true-覆盖，选择库表或者结构迁移时应该为false，注意只对旧版迁移有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOverrideRoot *bool `json:"IsOverrideRoot,omitempty" name:"IsOverrideRoot"`

	// 是否在迁移时设置目标库只读(仅对mysql有效)，true(设置只读)、false(不设置只读，默认此值)
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDstReadOnly *bool `json:"IsDstReadOnly,omitempty" name:"IsDstReadOnly"`

	// 其他附加信息，对于特定库可设置额外参数，Redis可定义如下的参数: 
	// ["ClientOutputBufferHardLimit":512, 	从机缓冲区的硬性容量限制(MB) 	"ClientOutputBufferSoftLimit":512, 	从机缓冲区的软性容量限制(MB) 	"ClientOutputBufferPersistTime":60, 从机缓冲区的软性限制持续时间(秒) 	"ReplBacklogSize":512, 	环形缓冲区容量限制(MB) 	"ReplTimeout":120，		复制超时时间(秒) ]
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraAttr []*KeyValuePairOption `json:"ExtraAttr,omitempty" name:"ExtraAttr"`
}

// Predefined struct for user
type ModifyCompareTaskNameRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 一致性校验任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type ModifyCompareTaskNameRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 一致性校验任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， **默认为此项配置**)，custom(自定义模式)
	ObjectMode *string `json:"ObjectMode,omitempty" name:"ObjectMode"`

	// 对比对象，若CompareObjectMode取值为custom，则此项必填
	Objects *CompareObject `json:"Objects,omitempty" name:"Objects"`
}

type ModifyCompareTaskRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 数据对比对象模式，sameAsMigrate(全部迁移对象， **默认为此项配置**)，custom(自定义模式)
	ObjectMode *string `json:"ObjectMode,omitempty" name:"ObjectMode"`

	// 对比对象，若CompareObjectMode取值为custom，则此项必填
	Objects *CompareObject `json:"Objects,omitempty" name:"Objects"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCompareTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCompareTaskResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 新实例规格大小，包括：micro、small、medium、large、xlarge、2xlarge
	NewInstanceClass *string `json:"NewInstanceClass,omitempty" name:"NewInstanceClass"`
}

type ModifyMigrateJobSpecRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 新实例规格大小，包括：micro、small、medium、large、xlarge、2xlarge
	NewInstanceClass *string `json:"NewInstanceClass,omitempty" name:"NewInstanceClass"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 修改后的迁移任务名
	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

type ModifyMigrateNameRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 修改后的迁移任务名
	JobName *string `json:"JobName,omitempty" name:"JobName"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type ModifyMigrationJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 运行模式，取值如：immediate(表示立即运行)、timed(表示定时运行)
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 迁移任务配置选项，描述任务如何执行迁移等一系列配置信息
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例信息
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例信息
	DstInfo *DBEndpointInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 迁移任务名称，最大长度128
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
}

type ModifyMigrationJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 运行模式，取值如：immediate(表示立即运行)、timed(表示定时运行)
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 迁移任务配置选项，描述任务如何执行迁移等一系列配置信息
	MigrateOption *MigrateOption `json:"MigrateOption,omitempty" name:"MigrateOption"`

	// 源实例信息
	SrcInfo *DBEndpointInfo `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标实例信息
	DstInfo *DBEndpointInfo `json:"DstInfo,omitempty" name:"DstInfo"`

	// 迁移任务名称，最大长度128
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 期待启动时间，当RunMode取值为timed时，此值必填，形如："2006-01-02 15:04:05"
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`

	// 标签信息
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMigrationJobRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMigrationJobResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type Objects struct {
	// 迁移对象类型 Partial(部分对象)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *string `json:"Mode,omitempty" name:"Mode"`

	// 同步对象，当 Mode 为 Partial 时，不为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Databases []*Database `json:"Databases,omitempty" name:"Databases"`

	// 高级对象类型，如function、procedure，当需要同步高级对象时，初始化类型必须包含结构初始化类型，即Options.InitType字段值为Structure或Full
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdvancedObjects []*string `json:"AdvancedObjects,omitempty" name:"AdvancedObjects"`
}

type Options struct {
	// 同步初始化选项，Data(全量数据初始化)、Structure(结构初始化)、Full(全量数据且结构初始化，默认)、None(仅增量)
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitType *string `json:"InitType,omitempty" name:"InitType"`

	// 同名表的处理，ReportErrorAfterCheck(前置校验并报错，默认)、InitializeAfterDelete(删除并重新初始化)、ExecuteAfterIgnore(忽略并继续执行)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealOfExistSameTable *string `json:"DealOfExistSameTable,omitempty" name:"DealOfExistSameTable"`

	// 冲突处理选项，ReportError(报错，默认为该值)、Ignore(忽略)、Cover(覆盖)、ConditionCover(条件覆盖)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictHandleType *string `json:"ConflictHandleType,omitempty" name:"ConflictHandleType"`

	// 是否添加附加列
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddAdditionalColumn *bool `json:"AddAdditionalColumn,omitempty" name:"AddAdditionalColumn"`

	// 所要同步的DML和DDL的选项，Insert(插入操作)、Update(更新操作)、Delete(删除操作)、DDL(结构同步)， 不填（不选），PartialDDL(自定义,和DdlOptions一起起作用 )
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpTypes []*string `json:"OpTypes,omitempty" name:"OpTypes"`

	// 冲突处理的详细选项，如条件覆盖中的条件行和条件操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConflictHandleOption *ConflictHandleOption `json:"ConflictHandleOption,omitempty" name:"ConflictHandleOption"`

	// DDL同步选项，具体描述要同步那些DDL
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlOptions []*DdlOption `json:"DdlOptions,omitempty" name:"DdlOptions"`
}

type ProcessProgress struct {
	// 步骤的状态， 包括：notStarted(未开始)、running(运行中)、success(成功)、failed(失败)等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 进度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// 总的步骤数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepAll *uint64 `json:"StepAll,omitempty" name:"StepAll"`

	// 当前进行的步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNow *uint64 `json:"StepNow,omitempty" name:"StepNow"`

	// 当前步骤输出提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Steps []*StepDetailInfo `json:"Steps,omitempty" name:"Steps"`
}

type ProcessStepTip struct {
	// 提示信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 解决方案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 文档提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

// Predefined struct for user
type RecoverMigrateJobRequestParams struct {
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type RecoverMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type RecoverSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步实例id（即标识一个同步作业），形如sync-werwfs23
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 任务规格
	NewInstanceClass *string `json:"NewInstanceClass,omitempty" name:"NewInstanceClass"`
}

type ResizeSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 任务规格
	NewInstanceClass *string `json:"NewInstanceClass,omitempty" name:"NewInstanceClass"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 恢复任务的模式，目前的取值有：clearData 清空目标实例数据，overwrite 以覆盖写的方式执行任务，normal 跟正常流程一样，不做额外动作
	ResumeOption *string `json:"ResumeOption,omitempty" name:"ResumeOption"`
}

type ResumeMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 恢复任务的模式，目前的取值有：clearData 清空目标实例数据，overwrite 以覆盖写的方式执行任务，normal 跟正常流程一样，不做额外动作
	ResumeOption *string `json:"ResumeOption,omitempty" name:"ResumeOption"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type ResumeSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// 迁移后的角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewRoleName *string `json:"NewRoleName,omitempty" name:"NewRoleName"`
}

type SkippedDetail struct {
	// 跳过的表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 跳过校验的表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*SkippedItem `json:"Items,omitempty" name:"Items"`
}

type SkippedItem struct {
	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Db *string `json:"Db,omitempty" name:"Db"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *string `json:"Table,omitempty" name:"Table"`

	// 未发起检查的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

// Predefined struct for user
type StartCompareRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`
}

type StartCompareRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type StartMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type StartSyncJobRequestParams struct {
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type StartSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// 步骤展现名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 步骤英文标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// 步骤状态:success(成功)、failed(失败)、running(执行中)、notStarted(未执行)、默认为notStarted
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 当前步骤开始的时间，格式为"yyyy-mm-dd hh:mm:ss"，该字段不存在或者为空是无意义 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 步骤错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepMessage *string `json:"StepMessage,omitempty" name:"StepMessage"`

	// 执行进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *uint64 `json:"Percent,omitempty" name:"Percent"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errors []*ProcessStepTip `json:"Errors,omitempty" name:"Errors"`

	// 告警提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warnings []*ProcessStepTip `json:"Warnings,omitempty" name:"Warnings"`
}

type StepInfo struct {
	// 步骤编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNo *uint64 `json:"StepNo,omitempty" name:"StepNo"`

	// 步骤名
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepName *string `json:"StepName,omitempty" name:"StepName"`

	// 步骤标号
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepId *string `json:"StepId,omitempty" name:"StepId"`

	// 当前状态，是否完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 步骤开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Errors []*StepTip `json:"Errors,omitempty" name:"Errors"`

	// 警告信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Warnings []*StepTip `json:"Warnings,omitempty" name:"Warnings"`

	// 当前步骤进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`
}

type StepTip struct {
	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	Code *string `json:"Code,omitempty" name:"Code"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 解决方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Solution *string `json:"Solution,omitempty" name:"Solution"`

	// 帮助文档
	// 注意：此字段可能返回 null，表示取不到有效值。
	HelpDoc *string `json:"HelpDoc,omitempty" name:"HelpDoc"`
}

// Predefined struct for user
type StopCompareRequestParams struct {
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`
}

type StopCompareRequest struct {
	*tchttp.BaseRequest
	
	// 迁移任务 Id
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 对比任务 ID，形如：dts-8yv4w2i1-cmp-37skmii9
	CompareTaskId *string `json:"CompareTaskId,omitempty" name:"CompareTaskId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type StopMigrateJobRequest struct {
	*tchttp.BaseRequest
	
	// 数据迁移任务ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
	JobId *string `json:"JobId,omitempty" name:"JobId"`
}

type StopSyncJobRequest struct {
	*tchttp.BaseRequest
	
	// 同步任务id
	JobId *string `json:"JobId,omitempty" name:"JobId"`
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
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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

type SyncDetailInfo struct {
	// 总步骤数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepAll *uint64 `json:"StepAll,omitempty" name:"StepAll"`

	// 当前步骤
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepNow *uint64 `json:"StepNow,omitempty" name:"StepNow"`

	// 总体进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Progress *int64 `json:"Progress,omitempty" name:"Progress"`

	// 当前步骤进度
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentStepProgress *int64 `json:"CurrentStepProgress,omitempty" name:"CurrentStepProgress"`

	// 同步两端数据量差距
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterSlaveDistance *int64 `json:"MasterSlaveDistance,omitempty" name:"MasterSlaveDistance"`

	// 同步两端时间差距
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecondsBehindMaster *int64 `json:"SecondsBehindMaster,omitempty" name:"SecondsBehindMaster"`

	// 总体描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 详细步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StepInfos []*StepInfo `json:"StepInfos,omitempty" name:"StepInfos"`
}

type SyncJobInfo struct {
	// 同步任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 同步任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobName *string `json:"JobName,omitempty" name:"JobName"`

	// 付款方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`

	// 运行模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunMode *string `json:"RunMode,omitempty" name:"RunMode"`

	// 期待运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpectRunTime *string `json:"ExpectRunTime,omitempty" name:"ExpectRunTime"`

	// 支持的所有操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	AllActions []*string `json:"AllActions,omitempty" name:"AllActions"`

	// 当前状态能进行的操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Actions []*string `json:"Actions,omitempty" name:"Actions"`

	// 同步选项
	// 注意：此字段可能返回 null，表示取不到有效值。
	Options *Options `json:"Options,omitempty" name:"Options"`

	// 同步库表对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Objects *Objects `json:"Objects,omitempty" name:"Objects"`

	// 任务规格
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitempty" name:"Specification"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 源端地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcRegion *string `json:"SrcRegion,omitempty" name:"SrcRegion"`

	// 源端数据库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcDatabaseType *string `json:"SrcDatabaseType,omitempty" name:"SrcDatabaseType"`

	// 源端接入类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcAccessType *string `json:"SrcAccessType,omitempty" name:"SrcAccessType"`

	// 源端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SrcInfo *Endpoint `json:"SrcInfo,omitempty" name:"SrcInfo"`

	// 目标端地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstRegion *string `json:"DstRegion,omitempty" name:"DstRegion"`

	// 目标端数据库类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstDatabaseType *string `json:"DstDatabaseType,omitempty" name:"DstDatabaseType"`

	// 目标端接入类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstAccessType *string `json:"DstAccessType,omitempty" name:"DstAccessType"`

	// 目标端信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	DstInfo *Endpoint `json:"DstInfo,omitempty" name:"DstInfo"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 标签相关
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagItem `json:"Tags,omitempty" name:"Tags"`

	// 同步任务运行步骤信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Detail *SyncDetailInfo `json:"Detail,omitempty" name:"Detail"`
}

type Table struct {
	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 新表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`

	// 过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	FilterCondition *string `json:"FilterCondition,omitempty" name:"FilterCondition"`
}

type TableItem struct {
	// 迁移的表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 迁移后的表名，当TableEditMode为rename时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewTableName *string `json:"NewTableName,omitempty" name:"NewTableName"`

	// 迁移临时表，针对pt-osc等工具在迁移过程中产生的临时表同步，需要提前将可能的临时表配置在这里，当TableEditMode为pt时此项必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpTables []*string `json:"TmpTables,omitempty" name:"TmpTables"`

	// 编辑表类型，rename(表映射)，pt(同步附加表)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableEditMode *string `json:"TableEditMode,omitempty" name:"TableEditMode"`
}

type TagFilter struct {
	// 标签键值
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue []*string `json:"TagValue,omitempty" name:"TagValue"`
}

type TagItem struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type TradeInfo struct {
	// 交易订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	DealName *string `json:"DealName,omitempty" name:"DealName"`

	// 上一次交易订单号
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastDealName *string `json:"LastDealName,omitempty" name:"LastDealName"`

	// 实例规格，包括：micro、small、medium、large、xlarge、2xlarge等
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceClass *string `json:"InstanceClass,omitempty" name:"InstanceClass"`

	// 计费任务状态， normal(计费或待计费)、resizing(变配中)、reversing(冲正中，比较短暂的状态)、isolating(隔离中，比较短暂的状态)、isolated(已隔离)、offlining(下线中)、offlined(已下线)、notBilled(未计费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeStatus *string `json:"TradeStatus,omitempty" name:"TradeStatus"`

	// 到期时间，格式为"yyyy-mm-dd hh:mm:ss"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 下线时间，格式为"yyyy-mm-dd hh:mm:ss"
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 隔离时间，格式为"yyyy-mm-dd hh:mm:ss"
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitempty" name:"IsolateTime"`

	// 下线原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineReason *string `json:"OfflineReason,omitempty" name:"OfflineReason"`

	// 隔离原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateReason *string `json:"IsolateReason,omitempty" name:"IsolateReason"`

	// 付费类型，包括：postpay(后付费)、prepay(预付费)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayType *string `json:"PayType,omitempty" name:"PayType"`

	// 任务计费类型，包括：billing(计费)、notBilling(不计费)、 promotions(促销活动中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingType *string `json:"BillingType,omitempty" name:"BillingType"`
}

type View struct {
	// view名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 新view名
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewViewName *string `json:"NewViewName,omitempty" name:"NewViewName"`
}

type ViewItem struct {
	// 视图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`

	// 迁移后的视图名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewViewName *string `json:"NewViewName,omitempty" name:"NewViewName"`
}