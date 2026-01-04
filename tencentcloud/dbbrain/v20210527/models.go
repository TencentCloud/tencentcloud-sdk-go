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

package v20210527

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

// Predefined struct for user
type AddUserContactRequestParams struct {
	// 联系人姓名，由中英文、数字、空格、!@#$%^&*()_+-=（）组成，不能以下划线开头，长度在20以内。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 邮箱地址，支持大小写字母、数字、下划线、连字符及@字符， 只能以数字或字母开头，邮箱地址不可重复。
	ContactInfo *string `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 服务产品类型，固定值："mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type AddUserContactRequest struct {
	*tchttp.BaseRequest
	
	// 联系人姓名，由中英文、数字、空格、!@#$%^&*()_+-=（）组成，不能以下划线开头，长度在20以内。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 邮箱地址，支持大小写字母、数字、下划线、连字符及@字符， 只能以数字或字母开头，邮箱地址不可重复。
	ContactInfo *string `json:"ContactInfo,omitnil,omitempty" name:"ContactInfo"`

	// 服务产品类型，固定值："mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *AddUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ContactInfo")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddUserContactResponseParams struct {
	// 添加成功的联系人id。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddUserContactResponse struct {
	*tchttp.BaseResponse
	Response *AddUserContactResponseParams `json:"Response"`
}

func (r *AddUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Aggregation struct {
	// 平均执行时间（ms）。
	AvgExecTime *int64 `json:"AvgExecTime,omitnil,omitempty" name:"AvgExecTime"`

	// 平均扫描行数。
	AvgDocsExamined *int64 `json:"AvgDocsExamined,omitnil,omitempty" name:"AvgDocsExamined"`

	// 产生慢查次数（/天）。
	SlowLogCount *int64 `json:"SlowLogCount,omitnil,omitempty" name:"SlowLogCount"`

	// 内存排序次数。
	SortCount *int64 `json:"SortCount,omitnil,omitempty" name:"SortCount"`

	// 慢查模板概览。
	SlowLogs []*string `json:"SlowLogs,omitnil,omitempty" name:"SlowLogs"`
}

type AlarmProfileList struct {
	// 0-不是 1-是
	IsWebHook *int64 `json:"IsWebHook,omitnil,omitempty" name:"IsWebHook"`

	// 接收告警用户数量
	ReceiveUinCount *int64 `json:"ReceiveUinCount,omitnil,omitempty" name:"ReceiveUinCount"`

	// 语言
	Lang *string `json:"Lang,omitnil,omitempty" name:"Lang"`

	// 模板类型
	TemplateType *string `json:"TemplateType,omitnil,omitempty" name:"TemplateType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 接收组数量
	ReceiveGroupCount *int64 `json:"ReceiveGroupCount,omitnil,omitempty" name:"ReceiveGroupCount"`

	// 更新用户的uin
	UpdateUin *int64 `json:"UpdateUin,omitnil,omitempty" name:"UpdateUin"`

	// 接收类型
	ReceiveType []*int64 `json:"ReceiveType,omitnil,omitempty" name:"ReceiveType"`

	// 接收用户信息
	ReceiveInfo []*ReceiveInfo `json:"ReceiveInfo,omitnil,omitempty" name:"ReceiveInfo"`

	// 更新时间，格式: "yyyy-MM-dd HH:mm:ss"
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 模板名
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 发送渠道
	SendChannel []*int64 `json:"SendChannel,omitnil,omitempty" name:"SendChannel"`

	// 模板id
	TemplateId *int64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// webhook数量
	WebHookCount *int64 `json:"WebHookCount,omitnil,omitempty" name:"WebHookCount"`
}

type AlarmsRules struct {
	// 间隔
	Interval *int64 `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 告警名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 指标
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 操作符
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 等级 
	// fatal-致命
	// critical-严重
	// warning-告警
	// information-通知
	Severity *string `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 指标值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type AuditInstance struct {
	// 审计状态，已开通审计为：YES，未开通审计为：ON。
	AuditStatus *string `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 审计日志大小，为兼容老版本用。
	BillingAmount *int64 `json:"BillingAmount,omitnil,omitempty" name:"BillingAmount"`

	// 计费确认状态，0-未确认；1-已确认。
	BillingConfirmed *int64 `json:"BillingConfirmed,omitnil,omitempty" name:"BillingConfirmed"`

	// 低频存储时长。
	ColdLogExpireDay *int64 `json:"ColdLogExpireDay,omitnil,omitempty" name:"ColdLogExpireDay"`

	// 低频日志存储量单位MB。
	ColdLogSize *int64 `json:"ColdLogSize,omitnil,omitempty" name:"ColdLogSize"`

	// 高频日志存储天数。
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`

	// 高频日志存储量，单位MB。
	HotLogSize *int64 `json:"HotLogSize,omitnil,omitempty" name:"HotLogSize"`

	// 实例Id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保存总天数，为高频存储时长+低频存储时长。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 实例创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例详细信息。
	InstanceInfo *AuditInstanceInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`
}

type AuditInstanceFilter struct {
	// 搜索条件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 要搜索的条件的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AuditInstanceInfo struct {
	// appId。
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 审计状态，0-未开通审计；1-已开通审计。
	AuditStatus *int64 `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`

	// 实例Id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 项目Id。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例所在地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 资源Tags。
	ResourceTags []*string `json:"ResourceTags,omitnil,omitempty" name:"ResourceTags"`
}

type AuditLogFile struct {
	// 审计日志文件生成异步任务ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 审计日志文件名称。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 审计日志文件创建时间。格式为 : "2019-03-20 17:09:13"。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件状态值。可能返回的值为：
	// "creating" - 生成中;
	// "failed" - 创建失败;
	// "success" - 已生成;
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 文件大小，单位为 KB。
	FileSize *float64 `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 审计日志下载地址。
	DownloadUrl *string `json:"DownloadUrl,omitnil,omitempty" name:"DownloadUrl"`

	// 错误信息。
	ErrMsg *string `json:"ErrMsg,omitnil,omitempty" name:"ErrMsg"`

	// 文件生成进度。（单位：%）
	Progress *float64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 文件生成成功时间。格式: "yyyy-MM-dd HH:mm:ss"
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AuditLogFilter struct {
	// 客户端地址。
	Host []*string `json:"Host,omitnil,omitempty" name:"Host"`

	// 数据库名称。
	DBName []*string `json:"DBName,omitnil,omitempty" name:"DBName"`

	// 用户名。
	User []*string `json:"User,omitnil,omitempty" name:"User"`

	// 返回行数。表示筛选返回行数大于该值的审计日志。
	SentRows *int64 `json:"SentRows,omitnil,omitempty" name:"SentRows"`

	// 影响行数。表示筛选影响行数大于该值的审计日志。
	AffectRows *int64 `json:"AffectRows,omitnil,omitempty" name:"AffectRows"`

	// 执行时间。单位为：µs。表示筛选执行时间大于该值的审计日志。
	ExecTime *int64 `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`
}

type AutonomyActionVo struct {
	// 自治任务ID。
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 自治事件ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 类型：支持RedisAutoScaleUp
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自治任务触发时间。格式: "yyyy-MM-dd HH:mm:ss"
	TriggerTime *string `json:"TriggerTime,omitnil,omitempty" name:"TriggerTime"`

	// 自治任务创建时间。格式: "yyyy-MM-dd HH:mm:ss"
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 自治任务更新时间，格式: "yyyy-MM-dd HH:mm:ss"
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 自治任务完成时间。格式: "yyyy-MM-dd HH:mm:ss"
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 剩余时间，单位：秒。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 触发原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 自治任务状态：RUNNING，FINISHED，TERMINATED，CANCELLED
	// 其中：
	// RUNNING    - 运行中  
	// FINISHED   - 已完成  
	// TERMINATED - 已终止  
	// CANCELLED  - 已取消  
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type AutonomyEventVo struct {
	// 自治事件ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 自治事件类型：支持RunningAutoRecovery，RedisAutoScale
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自治事件状态：支持 RUNNING，FINISHED，TERMINATED
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 触发原因。	
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 自治任务触发时间。
	TriggerTime *int64 `json:"TriggerTime,omitnil,omitempty" name:"TriggerTime"`

	// 自治任务最后触发时间。
	LastTriggerTime *int64 `json:"LastTriggerTime,omitnil,omitempty" name:"LastTriggerTime"`

	// 自治任务创建时间。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 自治任务更新时间。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 自治任务完成时间；非结束状态的时候，该值无意义。
	FinishTime *int64 `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`
}

type AutonomyUserProfileInfo struct {
	// 是否开启自治。枚举值：true，false。
	// 其中：
	// true - 开启
	// false - 关闭
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 用户Uin。
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`

	// 内存上限。
	MemoryUpperLimit *int64 `json:"MemoryUpperLimit,omitnil,omitempty" name:"MemoryUpperLimit"`

	// 指标阈值规则。
	ThresholdRule *MetricThreshold `json:"ThresholdRule,omitnil,omitempty" name:"ThresholdRule"`

	// 自治功能类型。
	EnabledItems []*string `json:"EnabledItems,omitnil,omitempty" name:"EnabledItems"`
}

// Predefined struct for user
type CancelDBAutonomyActionRequestParams struct {
	// 自治任务ID。
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CancelDBAutonomyActionRequest struct {
	*tchttp.BaseRequest
	
	// 自治任务ID。
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CancelDBAutonomyActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDBAutonomyActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActionId")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDBAutonomyActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDBAutonomyActionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelDBAutonomyActionResponse struct {
	*tchttp.BaseResponse
	Response *CancelDBAutonomyActionResponseParams `json:"Response"`
}

func (r *CancelDBAutonomyActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDBAutonomyActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDBAutonomyEventRequestParams struct {
	// 自治事件ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CancelDBAutonomyEventRequest struct {
	*tchttp.BaseRequest
	
	// 自治事件ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CancelDBAutonomyEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDBAutonomyEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelDBAutonomyEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelDBAutonomyEventResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelDBAutonomyEventResponse struct {
	*tchttp.BaseResponse
	Response *CancelDBAutonomyEventResponseParams `json:"Response"`
}

func (r *CancelDBAutonomyEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelDBAutonomyEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKillTaskRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CancelKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CancelKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelKillTaskResponseParams struct {
	// kill会话任务终止成功返回1。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CancelKillTaskResponseParams `json:"Response"`
}

func (r *CancelKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelRedisBigKeyAnalysisTasksRequestParams struct {
	// 自治任务ID。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CancelRedisBigKeyAnalysisTasksRequest struct {
	*tchttp.BaseRequest
	
	// 自治任务ID。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CancelRedisBigKeyAnalysisTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelRedisBigKeyAnalysisTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestIds")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelRedisBigKeyAnalysisTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelRedisBigKeyAnalysisTasksResponseParams struct {
	// 终止大Key任务结果；0-成功。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelRedisBigKeyAnalysisTasksResponse struct {
	*tchttp.BaseResponse
	Response *CancelRedisBigKeyAnalysisTasksResponseParams `json:"Response"`
}

func (r *CancelRedisBigKeyAnalysisTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelRedisBigKeyAnalysisTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceRequestParams struct {
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// NodeRequestType主要标识数据库产品类型，与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CloseAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// NodeRequestType主要标识数据库产品类型，与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CloseAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseAuditServiceResponseParams struct {
	// 0-关闭审计成功，非0关闭审计失败。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *CloseAuditServiceResponseParams `json:"Response"`
}

func (r *CloseAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmdCostGroup struct {
	// 该延迟区间内命令数占总命令数百分比
	Percent *float64 `json:"Percent,omitnil,omitempty" name:"Percent"`

	// 延迟区间上界，单位ms
	CostUpperLimit *string `json:"CostUpperLimit,omitnil,omitempty" name:"CostUpperLimit"`

	// 延迟区间下界，单位ms
	CostLowerLimit *string `json:"CostLowerLimit,omitnil,omitempty" name:"CostLowerLimit"`

	// 该延迟区间内命令次数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type CmdPerfInfo struct {
	// redis命令
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 监控数据
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`
}

type ContactItem struct {
	// 联系人id。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 联系人姓名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 联系人绑定的邮箱。
	Mail *string `json:"Mail,omitnil,omitempty" name:"Mail"`
}

// Predefined struct for user
type CreateAuditLogFileRequestParams struct {
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB for MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 与Product保持一致。如："dcdb" ,"mariadb"
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤条件。可按设置的过滤条件过滤日志。
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type CreateAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB for MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 与Product保持一致。如："dcdb" ,"mariadb"
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤条件。可按设置的过滤条件过滤日志。
	Filter *AuditLogFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

func (r *CreateAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filter")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAuditLogFileResponseParams struct {
	// 审计日志文件下载的任务ID
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateAuditLogFileResponseParams `json:"Response"`
}

func (r *CreateAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportTaskRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2020-11-08T14:00:00+08:00”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2020-11-09T14:00:00+08:00”。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否发送邮件: 0 - 否，1 - 是。
	SendMailFlag *int64 `json:"SendMailFlag,omitnil,omitempty" name:"SendMailFlag"`

	// 接收邮件的联系人ID数组。
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 接收邮件的联系组ID数组。
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateDBDiagReportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2020-11-08T14:00:00+08:00”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2020-11-09T14:00:00+08:00”。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否发送邮件: 0 - 否，1 - 是。
	SendMailFlag *int64 `json:"SendMailFlag,omitnil,omitempty" name:"SendMailFlag"`

	// 接收邮件的联系人ID数组。
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 接收邮件的联系组ID数组。
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SendMailFlag")
	delete(f, "ContactPerson")
	delete(f, "ContactGroup")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportTaskResponseParams struct {
	// 异步任务的请求 ID，可使用此 ID 查询异步任务的执行结果。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBDiagReportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportTaskResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 示例值：cdb-dctw4edd
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 健康报告相应的任务ID，可通过[DescribeDBDiagReportTasks](https://cloud.tencent.com/document/product/1130/54873)查询。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateDBDiagReportUrlRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 示例值：cdb-dctw4edd
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 健康报告相应的任务ID，可通过[DescribeDBDiagReportTasks](https://cloud.tencent.com/document/product/1130/54873)查询。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateDBDiagReportUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDBDiagReportUrlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDBDiagReportUrlResponseParams struct {
	// 健康报告浏览地址。
	ReportUrl *string `json:"ReportUrl,omitnil,omitempty" name:"ReportUrl"`

	// 健康报告浏览地址到期时间戳（秒）。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDBDiagReportUrlResponse struct {
	*tchttp.BaseResponse
	Response *CreateDBDiagReportUrlResponseParams `json:"Response"`
}

func (r *CreateDBDiagReportUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDBDiagReportUrlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKillTaskRequestParams struct {
	// kill会话任务的关联实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务持续时间，单位秒，手动关闭任务传-1。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 任务过滤条件，客户端IP。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 任务过滤条件，数据库库名,多个","隔开。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 任务过滤条件，相关命令，多个","隔开。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 任务过滤条件，支持单条件前缀匹配。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 任务过滤条件，支持多个关键字匹配，与Info参数互斥。
	Infos []*string `json:"Infos,omitnil,omitempty" name:"Infos"`

	// 任务过滤条件，用户类型。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 任务过滤条件，会话持续时长，单位秒。
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// kill会话任务的关联实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务持续时间，单位秒，手动关闭任务传-1。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 任务过滤条件，客户端IP。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 任务过滤条件，数据库库名,多个","隔开。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 任务过滤条件，相关命令，多个","隔开。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 任务过滤条件，支持单条件前缀匹配。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 任务过滤条件，支持多个关键字匹配，与Info参数互斥。
	Infos []*string `json:"Infos,omitnil,omitempty" name:"Infos"`

	// 任务过滤条件，用户类型。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 任务过滤条件，会话持续时长，单位秒。
	Time *int64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Duration")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "Command")
	delete(f, "Info")
	delete(f, "Infos")
	delete(f, "User")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateKillTaskResponseParams struct {
	// kill会话任务创建成功返回1
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateKillTaskResponseParams `json:"Response"`
}

func (r *CreateKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileRequestParams struct {
	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// 配置名称，需要保持唯一性，数据库巡检邮件配置名称自拟；定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"scheduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 配置绑定的实例ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。当配置级别为"Instance"时需要传入且只能为一个实例；当配置级别为“User”时，此参数不填。
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil,omitempty" name:"BindInstanceIds"`
}

type CreateMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// 配置名称，需要保持唯一性，数据库巡检邮件配置名称自拟；定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"scheduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 配置绑定的实例ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。当配置级别为"Instance"时需要传入且只能为一个实例；当配置级别为“User”时，此参数不填。
	BindInstanceIds []*string `json:"BindInstanceIds,omitnil,omitempty" name:"BindInstanceIds"`
}

func (r *CreateMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileInfo")
	delete(f, "ProfileLevel")
	delete(f, "ProfileName")
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "BindInstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMailProfileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateMailProfileResponseParams `json:"Response"`
}

func (r *CreateMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxySessionKillTaskRequestParams struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实列代理ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`
}

type CreateProxySessionKillTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实列代理ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`
}

func (r *CreateProxySessionKillTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "InstanceProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProxySessionKillTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProxySessionKillTaskResponseParams struct {
	// 创建 kill 会话任务返回的异步任务 id
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProxySessionKillTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProxySessionKillTaskResponseParams `json:"Response"`
}

func (r *CreateProxySessionKillTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProxySessionKillTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRedisBigKeyAnalysisTaskRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分片节点序号列表。当列表为空时，选择所有分片节点。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`

	// Top Key前缀的分隔符列表。
	// 目前仅支持以下分割符：[",", ";", ":", "_", "-", "+", "@", "=", "|", "#", "."]，当列表为空时，默认选择所有分隔符。
	KeyDelimiterList []*string `json:"KeyDelimiterList,omitnil,omitempty" name:"KeyDelimiterList"`
}

type CreateRedisBigKeyAnalysisTaskRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分片节点序号列表。当列表为空时，选择所有分片节点。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`

	// Top Key前缀的分隔符列表。
	// 目前仅支持以下分割符：[",", ";", ":", "_", "-", "+", "@", "=", "|", "#", "."]，当列表为空时，默认选择所有分隔符。
	KeyDelimiterList []*string `json:"KeyDelimiterList,omitnil,omitempty" name:"KeyDelimiterList"`
}

func (r *CreateRedisBigKeyAnalysisTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedisBigKeyAnalysisTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "ShardIds")
	delete(f, "KeyDelimiterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRedisBigKeyAnalysisTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRedisBigKeyAnalysisTaskResponseParams struct {
	// 异步任务ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRedisBigKeyAnalysisTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateRedisBigKeyAnalysisTaskResponseParams `json:"Response"`
}

func (r *CreateRedisBigKeyAnalysisTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRedisBigKeyAnalysisTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileRequestParams struct {
	// 取值范围1-7，分别代表周一至周日。
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil,omitempty" name:"WeekConfiguration"`

	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置名称，需要保持唯一性，定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"scheduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置订阅的实例ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	BindInstanceId *string `json:"BindInstanceId,omitnil,omitempty" name:"BindInstanceId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateSchedulerMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 取值范围1-7，分别代表周一至周日。
	WeekConfiguration []*int64 `json:"WeekConfiguration,omitnil,omitempty" name:"WeekConfiguration"`

	// 邮件配置内容。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 配置名称，需要保持唯一性，定期生成邮件配置命名格式："scheduler_" + {instanceId}，如"scheduler_cdb-test"。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置订阅的实例ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	BindInstanceId *string `json:"BindInstanceId,omitnil,omitempty" name:"BindInstanceId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，"mariadb" - 云数据库 MariaDB，"dcdb" - 云数据库 TDSQL MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateSchedulerMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WeekConfiguration")
	delete(f, "ProfileInfo")
	delete(f, "ProfileName")
	delete(f, "BindInstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSchedulerMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSchedulerMailProfileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSchedulerMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateSchedulerMailProfileResponseParams `json:"Response"`
}

func (r *CreateSchedulerMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSchedulerMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAuditLogExportTaskRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 导出日志开始时间，例如2020-12-28 00:00:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 导出日志结束时间，例如2020-12-28 01:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志风险等级列表，支持值包括：0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

type CreateSecurityAuditLogExportTaskRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 导出日志开始时间，例如2020-12-28 00:00:00。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 导出日志结束时间，例如2020-12-28 01:00:00。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志风险等级列表，支持值包括：0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*int64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

func (r *CreateSecurityAuditLogExportTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "DangerLevels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSecurityAuditLogExportTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSecurityAuditLogExportTaskResponseParams struct {
	// 日志导出任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSecurityAuditLogExportTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateSecurityAuditLogExportTaskResponseParams `json:"Response"`
}

func (r *CreateSecurityAuditLogExportTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSecurityAuditLogExportTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSqlFilterRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL语句的类型，取值包括SELECT, UPDATE, DELETE, INSERT, REPLACE。
	// 其中：
	// SELECT   - 查询  
	// UPDATE   - 更新  
	// DELETE   - 删除  
	// INSERT   - 插入  
	// REPLACE  - 替换
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 关键字，用于筛选SQL语句，多个关键字用英文逗号分隔，逗号不能作为关键词，多个关键词之间的关系为“逻辑与”。
	FilterKey *string `json:"FilterKey,omitnil,omitempty" name:"FilterKey"`

	// 最大并发度，取值不能小于0，如果该值设为 0，则表示限制所有匹配的SQL执行。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 限流时长，单位秒，支持-1和小于2147483647的正整数，-1表示永不过期。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 通过 [VerifyUserAccount](https://cloud.tencent.com/document/product/1130/72828) 获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type CreateSqlFilterRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL语句的类型，取值包括SELECT, UPDATE, DELETE, INSERT, REPLACE。
	// 其中：
	// SELECT   - 查询  
	// UPDATE   - 更新  
	// DELETE   - 删除  
	// INSERT   - 插入  
	// REPLACE  - 替换
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 关键字，用于筛选SQL语句，多个关键字用英文逗号分隔，逗号不能作为关键词，多个关键词之间的关系为“逻辑与”。
	FilterKey *string `json:"FilterKey,omitnil,omitempty" name:"FilterKey"`

	// 最大并发度，取值不能小于0，如果该值设为 0，则表示限制所有匹配的SQL执行。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 限流时长，单位秒，支持-1和小于2147483647的正整数，-1表示永不过期。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 通过 [VerifyUserAccount](https://cloud.tencent.com/document/product/1130/72828) 获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *CreateSqlFilterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSqlFilterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SqlType")
	delete(f, "FilterKey")
	delete(f, "MaxConcurrency")
	delete(f, "Duration")
	delete(f, "SessionToken")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSqlFilterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSqlFilterResponseParams struct {
	// 限流任务ID。
	FilterId *int64 `json:"FilterId,omitnil,omitempty" name:"FilterId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSqlFilterResponse struct {
	*tchttp.BaseResponse
	Response *CreateSqlFilterResponseParams `json:"Response"`
}

func (r *CreateSqlFilterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSqlFilterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAutonomyProfileRequestParams struct {
	// 配置类型，为需要配置的功能枚举值，目前包含以下枚举值：AutonomyGlobal（自治功能全局配置）、RedisAutoScaleUp（Redis自治扩容配置）
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 自治功能相关配置，标准JSON字符串格式。
	ProfileInfo *string `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`
}

type CreateUserAutonomyProfileRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型，为需要配置的功能枚举值，目前包含以下枚举值：AutonomyGlobal（自治功能全局配置）、RedisAutoScaleUp（Redis自治扩容配置）
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 自治功能相关配置，标准JSON字符串格式。
	ProfileInfo *string `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`
}

func (r *CreateUserAutonomyProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutonomyProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "ProfileInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateUserAutonomyProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateUserAutonomyProfileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateUserAutonomyProfileResponse struct {
	*tchttp.BaseResponse
	Response *CreateUserAutonomyProfileResponseParams `json:"Response"`
}

func (r *CreateUserAutonomyProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateUserAutonomyProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileRequestParams struct {
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB for MariaDB， "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB for MySQL， "postgres" - 云数据库 PostgreSQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// NodeRequestType主要标识数据库产品类型，与Product保持一致。该字段规则如下： 当product为"dcdb"则输入"dcdb"， 当product为"mariadb"则输入"mariadb"， 当product为"mysql"则输入"mysql"， 当product为"cynosdb"则输入"mysql"， 当product为"postgres"则输入"postgres"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志文件生成异步任务ID。可通过[查询审计日志文件](https://cloud.tencent.com/document/product/1130/90371)获得。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

type DeleteAuditLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB for MariaDB， "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB for MySQL， "postgres" - 云数据库 PostgreSQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// NodeRequestType主要标识数据库产品类型，与Product保持一致。该字段规则如下： 当product为"dcdb"则输入"dcdb"， 当product为"mariadb"则输入"mariadb"， 当product为"mysql"则输入"mysql"， 当product为"cynosdb"则输入"mysql"， 当product为"postgres"则输入"postgres"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 审计日志文件生成异步任务ID。可通过[查询审计日志文件](https://cloud.tencent.com/document/product/1130/90371)获得。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`
}

func (r *DeleteAuditLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAuditLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAuditLogFileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAuditLogFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAuditLogFileResponseParams `json:"Response"`
}

func (r *DeleteAuditLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAuditLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBDiagReportTasksRequestParams struct {
	// 需要删除的任务id列表。可通过[查询健康报告生成任务列表](https://cloud.tencent.com/document/product/1130/57805)获取
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"dbbrain-mysql" - 自建 MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的任务id列表。可通过[查询健康报告生成任务列表](https://cloud.tencent.com/document/product/1130/57805)获取
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"dbbrain-mysql" - 自建 MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DeleteDBDiagReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBDiagReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AsyncRequestIds")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDBDiagReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDBDiagReportTasksResponseParams struct {
	// 任务删除状态, 0-删除成功
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDBDiagReportTasksResponseParams `json:"Response"`
}

func (r *DeleteDBDiagReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDBDiagReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRedisBigKeyAnalysisTasksRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待删除的异步任务ID列表。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteRedisBigKeyAnalysisTasksRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待删除的异步任务ID列表。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DeleteRedisBigKeyAnalysisTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRedisBigKeyAnalysisTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRedisBigKeyAnalysisTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRedisBigKeyAnalysisTasksResponseParams struct {
	// 状态值，为0时代表正常处理。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRedisBigKeyAnalysisTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRedisBigKeyAnalysisTasksResponseParams `json:"Response"`
}

func (r *DeleteRedisBigKeyAnalysisTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRedisBigKeyAnalysisTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 日志导出任务Id列表，接口会忽略不存在或已删除的任务Id。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值： "mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 日志导出任务Id列表，接口会忽略不存在或已删除的任务Id。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值： "mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DeleteSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSecurityAuditLogExportTasksResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSecurityAuditLogExportTasksResponseParams `json:"Response"`
}

func (r *DeleteSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSqlFiltersRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 限流任务ID列表。通过接口[创建实例SQL限流任务](https://cloud.tencent.com/document/product/1130/72835)获得。
	FilterIds []*int64 `json:"FilterIds,omitnil,omitempty" name:"FilterIds"`

	// 通过 [VerifyUserAccount](https://cloud.tencent.com/document/product/1130/72828) 获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DeleteSqlFiltersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 限流任务ID列表。通过接口[创建实例SQL限流任务](https://cloud.tencent.com/document/product/1130/72835)获得。
	FilterIds []*int64 `json:"FilterIds,omitnil,omitempty" name:"FilterIds"`

	// 通过 [VerifyUserAccount](https://cloud.tencent.com/document/product/1130/72828) 获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DeleteSqlFiltersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSqlFiltersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilterIds")
	delete(f, "SessionToken")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSqlFiltersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSqlFiltersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSqlFiltersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSqlFiltersResponseParams `json:"Response"`
}

func (r *DeleteSqlFiltersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSqlFiltersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmTemplateRequestParams struct {
	// 搜索字段
	TemplateNameRegexp *string `json:"TemplateNameRegexp,omitnil,omitempty" name:"TemplateNameRegexp"`

	// 返回限制长度，最大值: 50，默认值: 50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏置，最大值: 无限制，默认值: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeAlarmTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 搜索字段
	TemplateNameRegexp *string `json:"TemplateNameRegexp,omitnil,omitempty" name:"TemplateNameRegexp"`

	// 返回限制长度，最大值: 50，默认值: 50
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏置，最大值: 无限制，默认值: 0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeAlarmTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateNameRegexp")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmTemplateResponseParams struct {
	// 模板列表
	ProfileList []*AlarmProfileList `json:"ProfileList,omitnil,omitempty" name:"ProfileList"`

	// 模板总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmTemplateResponseParams `json:"Response"`
}

func (r *DescribeAlarmTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactRequestParams struct {
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系人名数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeAllUserContactRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系人名数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DescribeAllUserContactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserContactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserContactResponseParams struct {
	// 联系人的总数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 联系人的信息。
	Contacts []*ContactItem `json:"Contacts,omitnil,omitempty" name:"Contacts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllUserContactResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserContactResponseParams `json:"Response"`
}

func (r *DescribeAllUserContactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserContactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupRequestParams struct {
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系组名称数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

type DescribeAllUserGroupRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，固定值：mysql。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 联系组名称数组，支持模糊搜索。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`
}

func (r *DescribeAllUserGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "Names")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllUserGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllUserGroupResponseParams struct {
	// 组总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 组信息。
	Groups []*GroupItem `json:"Groups,omitnil,omitempty" name:"Groups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllUserGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllUserGroupResponseParams `json:"Response"`
}

func (r *DescribeAllUserGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllUserGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListRequestParams struct {
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// NodeRequestType主要标识数据库产品类型，与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 审计状态标识，0-未开通审计；1-已开通审计，默认为0。
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询实例的搜索条件。Name包括：InstanceId， InstanceName。
	Filters []*AuditInstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeAuditInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// NodeRequestType主要标识数据库产品类型，与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 审计状态标识，0-未开通审计；1-已开通审计，默认为0。
	AuditSwitch *int64 `json:"AuditSwitch,omitnil,omitempty" name:"AuditSwitch"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 查询实例的搜索条件。Name包括：InstanceId， InstanceName。
	Filters []*AuditInstanceFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeAuditInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "AuditSwitch")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditInstanceListResponseParams struct {
	// 符合条件的实例个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例详情。
	Items []*AuditInstance `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditInstanceListResponseParams `json:"Response"`
}

func (r *DescribeAuditInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogFilesRequestParams struct {
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB for MariaDB， "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB for MySQL， "postgres" - 云数据库 PostgreSQL
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 该字段规则如下： 当product为"dcdb"则输入"dcdb"， 当product为"mariadb"则输入"mariadb"， 当product为"mysql"则输入"mysql"， 当product为"cynosdb"则输入"mysql"， 当product为"postgres"则输入"postgres"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAuditLogFilesRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB for MariaDB， "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB for MySQL， "postgres" - 云数据库 PostgreSQL
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 该字段规则如下： 当product为"dcdb"则输入"dcdb"， 当product为"mariadb"则输入"mariadb"， 当product为"mysql"则输入"mysql"， 当product为"cynosdb"则输入"mysql"， 当product为"postgres"则输入"postgres"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAuditLogFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogFilesResponseParams struct {
	// 符合条件的审计日志文件个数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 审计日志文件详情。
	Items []*AuditLogFile `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogFilesResponseParams `json:"Response"`
}

func (r *DescribeAuditLogFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBAutonomyActionRequestParams struct {
	// 自治任务ID。可通过 [DescribeDBAutonomyActions](https://cloud.tencent.com/document/product/1130/116974) 接口获取。
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBAutonomyActionRequest struct {
	*tchttp.BaseRequest
	
	// 自治任务ID。可通过 [DescribeDBAutonomyActions](https://cloud.tencent.com/document/product/1130/116974) 接口获取。
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBAutonomyActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBAutonomyActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActionId")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBAutonomyActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBAutonomyActionResponseParams struct {
	// 自治任务ID。
	ActionId *int64 `json:"ActionId,omitnil,omitempty" name:"ActionId"`

	// 自治事件ID。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 任务ID。
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 类型：支持RedisAutoScaleUp
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 自治任务触发时间。格式: "yyyy-MM-dd HH:mm:ss"
	TriggerTime *string `json:"TriggerTime,omitnil,omitempty" name:"TriggerTime"`

	// 自治任务创建时间。格式: "yyyy-MM-dd HH:mm:ss"
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 自治任务更新时间。格式: "yyyy-MM-dd HH:mm:ss"
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 自治任务完成时间。格式: "yyyy-MM-dd HH:mm:ss"
	FinishTime *string `json:"FinishTime,omitnil,omitempty" name:"FinishTime"`

	// 剩余时间，单位：秒。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 触发原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 自治任务状态：支持 RUNNING，FINISHED，TERMINATED，CANCELLED
	// 其中：
	// RUNNING    - 运行中  
	// FINISHED   - 已完成  
	// TERMINATED - 已终止  
	// CANCELLED  - 已取消  
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务相关的图表等信息。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBAutonomyActionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBAutonomyActionResponseParams `json:"Response"`
}

func (r *DescribeDBAutonomyActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBAutonomyActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBAutonomyActionsRequestParams struct {
	// 事件ID。可通过 [DescribeDBDiagHistory](https://cloud.tencent.com/document/product/1130/39559) 接口获取。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBAutonomyActionsRequest struct {
	*tchttp.BaseRequest
	
	// 事件ID。可通过 [DescribeDBDiagHistory](https://cloud.tencent.com/document/product/1130/39559) 接口获取。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBAutonomyActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBAutonomyActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventId")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBAutonomyActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBAutonomyActionsResponseParams struct {
	// 自治事件总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 自治事件列表。
	Actions []*AutonomyActionVo `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBAutonomyActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBAutonomyActionsResponseParams `json:"Response"`
}

func (r *DescribeDBAutonomyActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBAutonomyActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBAutonomyEventsRequestParams struct {
	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBAutonomyEventsRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 分页参数，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，默认值为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBAutonomyEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBAutonomyEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBAutonomyEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBAutonomyEventsResponseParams struct {
	// 自治事件列表总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 自治事件列表。
	Events []*AutonomyEventVo `json:"Events,omitnil,omitempty" name:"Events"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBAutonomyEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBAutonomyEventsResponseParams `json:"Response"`
}

func (r *DescribeDBAutonomyEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBAutonomyEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 
	// 查询TDSQL MySQL分布式实例:Instanceld：填写集群ID&Shard实例ID，如：dcdbt-157xxxk&shard-qxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件 ID 。通过“获取实例诊断历史[DescribeDBDiagHistory](https://cloud.tencent.com/document/product/1130/39559) ”获取。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"mariadb"-mariadb;"cynosdb"-TDSQL-C for MySQL ;"dcdb"-TDSQL MySQL ;"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagEventRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 
	// 查询TDSQL MySQL分布式实例:Instanceld：填写集群ID&Shard实例ID，如：dcdbt-157xxxk&shard-qxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 事件 ID 。通过“获取实例诊断历史[DescribeDBDiagHistory](https://cloud.tencent.com/document/product/1130/39559) ”获取。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"mariadb"-mariadb;"cynosdb"-TDSQL-C for MySQL ;"dcdb"-TDSQL MySQL ;"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBDiagEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EventId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventResponseParams struct {
	// 诊断项。
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// 诊断类型。支持值包括"高危账号","自增键耗尽","连接性检查","CPU利用率","死锁","全表扫描","高并发/压力请求","预编译语句过多","内存利用率","Metadata lock","磁盘超限","内存超限","只读锁","只读实例剔除","行锁","活跃会话","慢SQL","数据库快照","磁盘空间利用率","执行计划变化","主从切换","Table open cache命中率低","大表","事务未提交","事务导致复制延迟"等。
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// 事件 ID 。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 诊断事件详情，若无附加解释信息则输出为空。
	Explanation *string `json:"Explanation,omitnil,omitempty" name:"Explanation"`

	// 诊断概要。
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 诊断出的问题。
	Problem *string `json:"Problem,omitnil,omitempty" name:"Problem"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 开始时间。格式: "yyyy-MM-dd HH:mm:ss"
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 诊断建议，若无建议则输出为空。
	Suggestions *string `json:"Suggestions,omitnil,omitempty" name:"Suggestions"`

	// 保留字段。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 结束时间。格式: "yyyy-MM-dd HH:mm:ss"
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventsRequestParams struct {
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 风险等级列表，取值按影响程度从高至低分别为：1 - 致命、2 -严重、3 - 告警、4 - 提示、5 -健康。
	Severities []*int64 `json:"Severities,omitnil,omitempty" name:"Severities"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 查询TDSQL MySQL分布式实例:Instanceld：填写集群ID&Shard实例ID，如：dcdbt-157xxxk&shard-qxxxx
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeDBDiagEventsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 风险等级列表，取值按影响程度从高至低分别为：1 - 致命、2 -严重、3 - 告警、4 - 提示、5 -健康。
	Severities []*int64 `json:"Severities,omitnil,omitempty" name:"Severities"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 查询TDSQL MySQL分布式实例:Instanceld：填写集群ID&Shard实例ID，如：dcdbt-157xxxk&shard-qxxxx
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeDBDiagEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Severities")
	delete(f, "InstanceIds")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagEventsResponseParams struct {
	// 诊断事件的总数目。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 诊断事件的列表。
	Items []*DiagHistoryEventItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagEventsResponseParams `json:"Response"`
}

func (r *DescribeDBDiagEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 
	// 查询TDSQL MySQL分布式实例:Instanceld：填写集群ID&Shard实例ID，如：dcdbt-157xxxk&shard-qxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。结束时间与开始时间的间隔最大可为2天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-11 12:13:14”，结束时间与开始时间的间隔最大可为2天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"mariadb"-mariadb;"cynosdb"-TDSQL-C for MySQL ;"dcdb"-TDSQL MySQL ;"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	// 
	// 查询TDSQL MySQL分布式实例:Instanceld：填写集群ID&Shard实例ID，如：dcdbt-157xxxk&shard-qxxxx
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。结束时间与开始时间的间隔最大可为2天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-11 12:13:14”，结束时间与开始时间的间隔最大可为2天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"mariadb"-mariadb;"cynosdb"-TDSQL-C for MySQL ;"dcdb"-TDSQL MySQL ;"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBDiagHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagHistoryResponseParams struct {
	// 事件描述。
	Events []*DiagHistoryEventItem `json:"Events,omitnil,omitempty" name:"Events"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagHistoryResponseParams `json:"Response"`
}

func (r *DescribeDBDiagHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportContentRequestParams struct {
	// 实例名
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异步任务ID
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagReportContentRequest struct {
	*tchttp.BaseRequest
	
	// 实例名
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 异步任务ID
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBDiagReportContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagReportContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportContentResponseParams struct {
	// 报告内容。
	Report *string `json:"Report,omitnil,omitempty" name:"Report"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagReportContentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagReportContentResponseParams `json:"Response"`
}

func (r *DescribeDBDiagReportContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksRequestParams struct {
	// 第一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID数组，用于筛选指定实例的任务列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 计划任务；"MANUAL" - 手动触发。
	Sources []*string `json:"Sources,omitnil,omitempty" name:"Sources"`

	// 报告的健康等级，支持的取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK" - 危险；"HIGH_RISK" - 高危。
	HealthLevels *string `json:"HealthLevels,omitnil,omitempty" name:"HealthLevels"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatuses *string `json:"TaskStatuses,omitnil,omitempty" name:"TaskStatuses"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBDiagReportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 第一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 最后一个任务的开始时间，用于范围查询，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID数组，用于筛选指定实例的任务列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 计划任务；"MANUAL" - 手动触发。
	Sources []*string `json:"Sources,omitnil,omitempty" name:"Sources"`

	// 报告的健康等级，支持的取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK" - 危险；"HIGH_RISK" - 高危。
	HealthLevels *string `json:"HealthLevels,omitnil,omitempty" name:"HealthLevels"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatuses *string `json:"TaskStatuses,omitnil,omitempty" name:"TaskStatuses"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBDiagReportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceIds")
	delete(f, "Sources")
	delete(f, "HealthLevels")
	delete(f, "TaskStatuses")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBDiagReportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBDiagReportTasksResponseParams struct {
	// 任务总数目。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表。
	Tasks []*HealthReportTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBDiagReportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBDiagReportTasksResponseParams `json:"Response"`
}

func (r *DescribeDBDiagReportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBDiagReportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBPerfTimeSeriesRequestParams struct {
	// 需要获取性能趋势的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标名称，多个指标之间用逗号分隔。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 需要获取性能趋势的集群ID。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 性能数据统计粒度。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实列节点ID。
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 实列代理ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 代理节点ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

type DescribeDBPerfTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取性能趋势的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指标名称，多个指标之间用逗号分隔。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 需要获取性能趋势的集群ID。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 性能数据统计粒度。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 实列节点ID。
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 实列代理ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 代理节点ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

func (r *DescribeDBPerfTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBPerfTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Metric")
	delete(f, "Product")
	delete(f, "ClusterId")
	delete(f, "Period")
	delete(f, "InstanceNodeId")
	delete(f, "InstanceProxyId")
	delete(f, "ProxyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBPerfTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBPerfTimeSeriesResponseParams struct {
	// 实列性能数据。
	Data *MonitorMetricSeriesData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBPerfTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBPerfTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeDBPerfTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBPerfTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间段天数，截止日期为当日，默认为7天。
	RangeDays *int64 `json:"RangeDays,omitnil,omitempty" name:"RangeDays"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeDBSpaceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 时间段天数，截止日期为当日，默认为7天。
	RangeDays *int64 `json:"RangeDays,omitnil,omitempty" name:"RangeDays"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeDBSpaceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RangeDays")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDBSpaceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDBSpaceStatusResponseParams struct {
	// 磁盘增长量(MB)。
	Growth *int64 `json:"Growth,omitnil,omitempty" name:"Growth"`

	// 磁盘剩余(MB)。
	Remain *int64 `json:"Remain,omitnil,omitempty" name:"Remain"`

	// 磁盘总量(MB)。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 预计可用天数。
	AvailableDays *int64 `json:"AvailableDays,omitnil,omitempty" name:"AvailableDays"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDBSpaceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDBSpaceStatusResponseParams `json:"Response"`
}

func (r *DescribeDBSpaceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDBSpaceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesRequestParams struct {
	// 是否是DBbrain支持的实例，固定传 true。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"cynosdb" - 云数据库 TDSQL-C for MySQL，"dbbrain-mysql" - 自建 MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页参数，偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页值，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据实例名称条件查询。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 根据实例ID条件查询。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 根据地域条件查询。
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

type DescribeDiagDBInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 是否是DBbrain支持的实例，固定传 true。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"cynosdb" - 云数据库 TDSQL-C for MySQL，"dbbrain-mysql" - 自建 MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页参数，偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页参数，分页值，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据实例名称条件查询。
	InstanceNames []*string `json:"InstanceNames,omitnil,omitempty" name:"InstanceNames"`

	// 根据实例ID条件查询。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 根据地域条件查询。
	Regions []*string `json:"Regions,omitnil,omitempty" name:"Regions"`
}

func (r *DescribeDiagDBInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "IsSupported")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceNames")
	delete(f, "InstanceIds")
	delete(f, "Regions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDiagDBInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDiagDBInstancesResponseParams struct {
	// 实例总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 全实例巡检状态：0：开启全实例巡检；1：未开启全实例巡检。
	DbScanStatus *int64 `json:"DbScanStatus,omitnil,omitempty" name:"DbScanStatus"`

	// 实例相关信息。
	Items []*InstanceInfo `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDiagDBInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDiagDBInstancesResponseParams `json:"Response"`
}

func (r *DescribeDiagDBInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDiagDBInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreRequestParams struct {
	// 需要获取健康得分的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取健康得分的时间，时间格式如：2019-09-10 12:13:14。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeHealthScoreRequest struct {
	*tchttp.BaseRequest
	
	// 需要获取健康得分的实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 获取健康得分的时间，时间格式如：2019-09-10 12:13:14。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeHealthScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Time")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHealthScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreResponseParams struct {
	// 健康得分以及异常扣分项。
	Data *HealthScoreInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHealthScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHealthScoreResponseParams `json:"Response"`
}

func (r *DescribeHealthScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreTimeSeriesRequestParams struct {
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeHealthScoreTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeHealthScoreTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeHealthScoreTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeHealthScoreTimeSeriesResponseParams struct {
	// 健康得分趋势数据
	Data *HealthScoreTimeSeriesData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeHealthScoreTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeHealthScoreTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeHealthScoreTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeHealthScoreTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexRecommendAggregationSlowLogsRequestParams struct {
	// 服务产品类型，支持值包括："mongodb" - 云数据库 。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称。
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 表名。
	Collection *string `json:"Collection,omitnil,omitempty" name:"Collection"`

	// 签名。这个值是 [DescribeIndexRecommendInfo](https://cloud.tencent.com/document/product/1130/98911) 接口返回
	Signs []*string `json:"Signs,omitnil,omitempty" name:"Signs"`
}

type DescribeIndexRecommendAggregationSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括："mongodb" - 云数据库 。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名称。
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 表名。
	Collection *string `json:"Collection,omitnil,omitempty" name:"Collection"`

	// 签名。这个值是 [DescribeIndexRecommendInfo](https://cloud.tencent.com/document/product/1130/98911) 接口返回
	Signs []*string `json:"Signs,omitnil,omitempty" name:"Signs"`
}

func (r *DescribeIndexRecommendAggregationSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRecommendAggregationSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "Db")
	delete(f, "Collection")
	delete(f, "Signs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexRecommendAggregationSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexRecommendAggregationSlowLogsResponseParams struct {
	// 查询实例慢查询聚合结果。
	Aggregation *Aggregation `json:"Aggregation,omitnil,omitempty" name:"Aggregation"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexRecommendAggregationSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexRecommendAggregationSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeIndexRecommendAggregationSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRecommendAggregationSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexRecommendInfoRequestParams struct {
	// 服务产品类型，支持值包括："mongodb" - 云数据库 。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeIndexRecommendInfoRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括："mongodb" - 云数据库 。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeIndexRecommendInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRecommendInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexRecommendInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIndexRecommendInfoResponseParams struct {
	// 索引推荐的集合数量。
	CollectionNum *int64 `json:"CollectionNum,omitnil,omitempty" name:"CollectionNum"`

	// 索引推荐的索引数量。
	IndexNum *int64 `json:"IndexNum,omitnil,omitempty" name:"IndexNum"`

	// 索引项。
	Items []*MongoDBIndex `json:"Items,omitnil,omitempty" name:"Items"`

	// 优化级别，1-4，优先级从高到低。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 历史优化数。
	Optimized *int64 `json:"Optimized,omitnil,omitempty" name:"Optimized"`

	// 累计优化条数。
	OptimizedCount *int64 `json:"OptimizedCount,omitnil,omitempty" name:"OptimizedCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeIndexRecommendInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIndexRecommendInfoResponseParams `json:"Response"`
}

func (r *DescribeIndexRecommendInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRecommendInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileRequestParams struct {
	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单位，最大支持50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据邮件配置名称查询，定期发送的邮件配置名称遵循："scheduler_"+{instanceId}的规则。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`
}

type DescribeMailProfileRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 分页偏移量。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页单位，最大支持50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 根据邮件配置名称查询，定期发送的邮件配置名称遵循："scheduler_"+{instanceId}的规则。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`
}

func (r *DescribeMailProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ProfileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMailProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMailProfileResponseParams struct {
	// 邮件配置详情。
	ProfileList []*UserProfile `json:"ProfileList,omitnil,omitempty" name:"ProfileList"`

	// 邮件配置总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMailProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMailProfileResponseParams `json:"Response"`
}

func (r *DescribeMailProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMailProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricTopProxiesRequestParams struct {
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指标 eg: cpu_util, connections
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 默认前20条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMetricTopProxiesRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指标 eg: cpu_util, connections
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 默认前20条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMetricTopProxiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricTopProxiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "Metric")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMetricTopProxiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricTopProxiesResponseParams struct {
	// 命令列表
	Data []*RedisMetricTopProxiesData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMetricTopProxiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMetricTopProxiesResponseParams `json:"Response"`
}

func (r *DescribeMetricTopProxiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricTopProxiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMongoDBProcessListRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值：mongodb
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表,如果是多个 使用 ','  分割
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 命令类型 ,如果是多个 使用 ','  分割
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeMongoDBProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值：mongodb
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表,如果是多个 使用 ','  分割
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 命令类型 ,如果是多个 使用 ','  分割
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeMongoDBProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMongoDBProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "ID")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "Type")
	delete(f, "Time")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMongoDBProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMongoDBProcessListResponseParams struct {
	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcessList *MongoDBProcessList `json:"ProcessList,omitnil,omitempty" name:"ProcessList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMongoDBProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMongoDBProcessListResponseParams `json:"Response"`
}

func (r *DescribeMongoDBProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMongoDBProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作账号名，用于筛选线程列表。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 线程的操作状态。包含以下枚举值：Sending data​-线程正在处理查询结果， ​Sorting result​-线程正在对查询结果进行排序​，Creating tmp table​-线程正在创建临时表，Altering table​-线程正在执行表结构变更，Updating-线程执行更新中。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 线程的执行类型。包含以下枚举值：Sleep-线程处于空闲状态，Query-线程正在执行一个查询，Connect-从服务器连接到主服务器，Execute-线程正在执行预处理语句。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 线程的操作语句，用于筛选线程列表。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"mariadb"-mariadb;"cynosdb"-TDSQL-C for MySQL ;"dcdb"-TDSQL MySQL 默认为"mysql"。
	// 
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 会话统计的维度信息,可以多个维度。
	StatDimensions []*StatDimension `json:"StatDimensions,omitnil,omitempty" name:"StatDimensions"`
}

type DescribeMySqlProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 线程的ID，用于筛选线程列表。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作账号名，用于筛选线程列表。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 线程的操作主机地址，用于筛选线程列表。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库，用于筛选线程列表。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 线程的操作状态。包含以下枚举值：Sending data​-线程正在处理查询结果， ​Sorting result​-线程正在对查询结果进行排序​，Creating tmp table​-线程正在创建临时表，Altering table​-线程正在执行表结构变更，Updating-线程执行更新中。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 线程的执行类型。包含以下枚举值：Sleep-线程处于空闲状态，Query-线程正在执行一个查询，Connect-从服务器连接到主服务器，Execute-线程正在执行预处理语句。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 线程的操作时长最小值，单位秒，用于筛选操作时长大于该值的线程列表。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 线程的操作语句，用于筛选线程列表。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// 返回数量，默认20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"mariadb"-mariadb;"cynosdb"-TDSQL-C for MySQL ;"dcdb"-TDSQL MySQL 默认为"mysql"。
	// 
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 会话统计的维度信息,可以多个维度。
	StatDimensions []*StatDimension `json:"StatDimensions,omitnil,omitempty" name:"StatDimensions"`
}

func (r *DescribeMySqlProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ID")
	delete(f, "User")
	delete(f, "Host")
	delete(f, "DB")
	delete(f, "State")
	delete(f, "Command")
	delete(f, "Time")
	delete(f, "Info")
	delete(f, "Limit")
	delete(f, "Product")
	delete(f, "StatDimensions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMySqlProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMySqlProcessListResponseParams struct {
	// 实时线程列表。
	ProcessList []*MySqlProcess `json:"ProcessList,omitnil,omitempty" name:"ProcessList"`

	// sql会话统计信息。如果请求参数中包含StatDimensions，该参数则可能返回，否则不返回。
	Statistics []*StatisticInfo `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMySqlProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMySqlProcessListResponseParams `json:"Response"`
}

func (r *DescribeMySqlProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMySqlProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoPrimaryKeyTablesRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早为30天前的日期。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeNoPrimaryKeyTablesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早为30天前的日期。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeNoPrimaryKeyTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoPrimaryKeyTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNoPrimaryKeyTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoPrimaryKeyTablesResponseParams struct {
	// 无主键表总数。
	NoPrimaryKeyTableCount *int64 `json:"NoPrimaryKeyTableCount,omitnil,omitempty" name:"NoPrimaryKeyTableCount"`

	// 与昨日扫描无主键表的差值，正数为增加，负数为减少，0为无变化。
	NoPrimaryKeyTableCountDiff *int64 `json:"NoPrimaryKeyTableCountDiff,omitnil,omitempty" name:"NoPrimaryKeyTableCountDiff"`

	// 记录的无主键表总数（不超过无主键表总数），可用于分页查询。
	NoPrimaryKeyTableRecordCount *int64 `json:"NoPrimaryKeyTableRecordCount,omitnil,omitempty" name:"NoPrimaryKeyTableRecordCount"`

	// 无主键表列表。
	NoPrimaryKeyTables []*Table `json:"NoPrimaryKeyTables,omitnil,omitempty" name:"NoPrimaryKeyTables"`

	// 采集时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNoPrimaryKeyTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNoPrimaryKeyTablesResponseParams `json:"Response"`
}

func (r *DescribeNoPrimaryKeyTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoPrimaryKeyTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyProcessStatisticsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 该实例下需要查询的某一个 ProxyID 。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 返回数量。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照某字段排序。支持值包括："AllConn"，"ActiveConn"，"Ip"。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方向。支持值包括："DESC"，"ASC"。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type DescribeProxyProcessStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 该实例下需要查询的某一个 ProxyID 。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 返回数量。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 偏移量，默认0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 按照某字段排序。支持值包括："AllConn"，"ActiveConn"，"Ip"。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方向。支持值包括："DESC"，"ASC"。
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *DescribeProxyProcessStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyProcessStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "InstanceProxyId")
	delete(f, "Limit")
	delete(f, "Product")
	delete(f, "Offset")
	delete(f, "SortBy")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxyProcessStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxyProcessStatisticsResponseParams struct {
	// 实时会话统计详情。
	ProcessStatistics *ProcessStatistic `json:"ProcessStatistics,omitnil,omitempty" name:"ProcessStatistics"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxyProcessStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxyProcessStatisticsResponseParams `json:"Response"`
}

func (r *DescribeProxyProcessStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxyProcessStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// kill 会话异步任务 ID,  接口 CreateProxySessionKillTask 调用成功后获取。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeProxySessionKillTasksRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// kill 会话异步任务 ID,  接口 CreateProxySessionKillTask 调用成功后获取。
	AsyncRequestIds []*int64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeProxySessionKillTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AsyncRequestIds")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProxySessionKillTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProxySessionKillTasksResponseParams struct {
	// kill 任务的详情。
	Tasks []*TaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProxySessionKillTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProxySessionKillTasksResponseParams `json:"Response"`
}

func (r *DescribeProxySessionKillTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProxySessionKillTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisBigKeyAnalysisTasksRequestParams struct {
	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRedisBigKeyAnalysisTasksRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRedisBigKeyAnalysisTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisBigKeyAnalysisTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisBigKeyAnalysisTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisBigKeyAnalysisTasksResponseParams struct {
	// 任务总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 任务列表。
	Tasks []*RedisBigKeyTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisBigKeyAnalysisTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisBigKeyAnalysisTasksResponseParams `json:"Response"`
}

func (r *DescribeRedisBigKeyAnalysisTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisBigKeyAnalysisTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisCmdPerfTimeSeriesRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 需要分析的redis命令
	CommandList []*string `json:"CommandList,omitnil,omitempty" name:"CommandList"`

	// 监控指标，包括：qps,latency_p99,latency_avg,latency_max，以逗号分隔
	// 其中：
	// qps          - 每秒查询率  
	// latency_p99  - 99分位延迟  
	// latency_avg  - 平均延迟  
	// latency_max  - 最大延迟  
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Period，监控指标时间粒度，单位秒，若不提供则根据开始时间和结束时间取默认值
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

type DescribeRedisCmdPerfTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 需要分析的redis命令
	CommandList []*string `json:"CommandList,omitnil,omitempty" name:"CommandList"`

	// 监控指标，包括：qps,latency_p99,latency_avg,latency_max，以逗号分隔
	// 其中：
	// qps          - 每秒查询率  
	// latency_p99  - 99分位延迟  
	// latency_avg  - 平均延迟  
	// latency_max  - 最大延迟  
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// Period，监控指标时间粒度，单位秒，若不提供则根据开始时间和结束时间取默认值
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`
}

func (r *DescribeRedisCmdPerfTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisCmdPerfTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "CommandList")
	delete(f, "Metric")
	delete(f, "Period")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisCmdPerfTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisCmdPerfTimeSeriesResponseParams struct {
	// redis命令延迟趋势
	CmdPerfList []*CmdPerfInfo `json:"CmdPerfList,omitnil,omitempty" name:"CmdPerfList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisCmdPerfTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisCmdPerfTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeRedisCmdPerfTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisCmdPerfTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisCommandCostStatisticsRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeRedisCommandCostStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeRedisCommandCostStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisCommandCostStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisCommandCostStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisCommandCostStatisticsResponseParams struct {
	// redis延迟分布区间
	CmdCostGroups []*CmdCostGroup `json:"CmdCostGroups,omitnil,omitempty" name:"CmdCostGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisCommandCostStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisCommandCostStatisticsResponseParams `json:"Response"`
}

func (r *DescribeRedisCommandCostStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisCommandCostStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisCommandOverviewRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeRedisCommandOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2025-03-17T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2025-03-17T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeRedisCommandOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisCommandOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisCommandOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisCommandOverviewResponseParams struct {
	// redis访问命令统计
	CmdList []*RedisCmdInfo `json:"CmdList,omitnil,omitempty" name:"CmdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisCommandOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisCommandOverviewResponseParams `json:"Response"`
}

func (r *DescribeRedisCommandOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisCommandOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisProcessListRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询的Proxy节点数量上限，默认值为20，最大值为50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Proxy节点的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRedisProcessListRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询的Proxy节点数量上限，默认值为20，最大值为50。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// Proxy节点的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRedisProcessListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisProcessListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisProcessListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisProcessListResponseParams struct {
	// 该实例的Proxy节点数量，可用于分页查询。
	ProxyCount *int64 `json:"ProxyCount,omitnil,omitempty" name:"ProxyCount"`

	// 实时会话详情列表。
	Processes []*Process `json:"Processes,omitnil,omitempty" name:"Processes"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisProcessListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisProcessListResponseParams `json:"Response"`
}

func (r *DescribeRedisProcessListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisProcessListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisSlowLogTopSqlsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Redis Proxy节点ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 排序键，支持ExecTimes,QueryTime,QueryTimeMax,QueryTimeAvg等排序键，默认为QueryTime。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序），默认为DESC。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRedisSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Redis Proxy节点ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 排序键，支持ExecTimes,QueryTime,QueryTimeMax,QueryTimeAvg等排序键，默认为QueryTime。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序），默认为DESC。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRedisSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisSlowLogTopSqlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "InstanceProxyId")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisSlowLogTopSqlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisSlowLogTopSqlsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢日志 top sql 列表。
	Rows []*SlowLogAgg `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisSlowLogTopSqlsResponseParams `json:"Response"`
}

func (r *DescribeRedisSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisSlowLogTopSqlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopBigKeysRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询某个日期最新的任务，如2021-05-27，最早可为前30天的日期。该参数与AsyncRequestId参数不可同时为空。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 排序字段，取值包括Capacity - 内存，ItemCount - 元素数量，默认为Capacity。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// key类型筛选条件，默认为不进行筛选，取值包括string, list, set, hash, sortedset, stream。
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 异步任务ID。当为空时，选择最近任务的ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 分片节点序号列表。当列表为空时，选择所有分片节点。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`

	// 是否仅查询未设置过期时间的大Key。
	// 当为true时，仅查询未设置过期时间的大Key，默认为false。
	UnExpireKey *bool `json:"UnExpireKey,omitnil,omitempty" name:"UnExpireKey"`
}

type DescribeRedisTopBigKeysRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询某个日期最新的任务，如2021-05-27，最早可为前30天的日期。该参数与AsyncRequestId参数不可同时为空。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 排序字段，取值包括Capacity - 内存，ItemCount - 元素数量，默认为Capacity。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// key类型筛选条件，默认为不进行筛选，取值包括string, list, set, hash, sortedset, stream。
	KeyType *string `json:"KeyType,omitnil,omitempty" name:"KeyType"`

	// 查询数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 异步任务ID。当为空时，选择最近任务的ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 分片节点序号列表。当列表为空时，选择所有分片节点。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`

	// 是否仅查询未设置过期时间的大Key。
	// 当为true时，仅查询未设置过期时间的大Key，默认为false。
	UnExpireKey *bool `json:"UnExpireKey,omitnil,omitempty" name:"UnExpireKey"`
}

func (r *DescribeRedisTopBigKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopBigKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "Date")
	delete(f, "SortBy")
	delete(f, "KeyType")
	delete(f, "Limit")
	delete(f, "AsyncRequestId")
	delete(f, "ShardIds")
	delete(f, "UnExpireKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopBigKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopBigKeysResponseParams struct {
	// top key列表。
	TopKeys []*RedisKeySpaceData `json:"TopKeys,omitnil,omitempty" name:"TopKeys"`

	// 采集时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisTopBigKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopBigKeysResponseParams `json:"Response"`
}

func (r *DescribeRedisTopBigKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopBigKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopCostCommandsRequestParams struct {
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 默认前20条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeRedisTopCostCommandsRequest struct {
	*tchttp.BaseRequest
	
	// 开始时间，如“2021-05-27 00:00:00”，支持的最早查询时间为当前时间的前30天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2021-05-27 01:00:00”，支持的最早查询时间为当前时间的前30天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 实例ID列表。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括："mysql" - 云数据库 MySQL，"redis" - 云数据库 Redis，"mariadb"-数据库mariadb    默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 默认前20条
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeRedisTopCostCommandsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopCostCommandsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopCostCommandsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopCostCommandsResponseParams struct {
	// 命令列表
	TopCostCmdList []*RedisCostCmd `json:"TopCostCmdList,omitnil,omitempty" name:"TopCostCmdList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisTopCostCommandsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopCostCommandsResponseParams `json:"Response"`
}

func (r *DescribeRedisTopCostCommandsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopCostCommandsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopHotKeysRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2024-09-22T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2024-09-22T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Redis 节点数组。
	InstanceNodeIds []*string `json:"InstanceNodeIds,omitnil,omitempty" name:"InstanceNodeIds"`

	// top 数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRedisTopHotKeysRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2024-09-22T00:00:00+00:00”。0天 < 当前服务器时间 - 开始时间 <= 10天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2024-09-22T01:00:00+00:00”，0天 < 结束时间 - 开始时间 <= 10天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，仅仅支持值 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Redis 节点数组。
	InstanceNodeIds []*string `json:"InstanceNodeIds,omitnil,omitempty" name:"InstanceNodeIds"`

	// top 数目，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRedisTopHotKeysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopHotKeysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "InstanceNodeIds")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopHotKeysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopHotKeysResponseParams struct {
	// 热Key分析结果
	TopHotKeys []*TopHotKeys `json:"TopHotKeys,omitnil,omitempty" name:"TopHotKeys"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisTopHotKeysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopHotKeysResponseParams `json:"Response"`
}

func (r *DescribeRedisTopHotKeysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopHotKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早可为前30天的日期。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询数目，默认为20，最大值为500。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分片ID数组。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

type DescribeRedisTopKeyPrefixListRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询日期，如2021-05-27，最早可为前30天的日期。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询数目，默认为20，最大值为500。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分片ID数组。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

func (r *DescribeRedisTopKeyPrefixListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Date")
	delete(f, "Product")
	delete(f, "Limit")
	delete(f, "ShardIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisTopKeyPrefixListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisTopKeyPrefixListResponseParams struct {
	// top key前缀列表。
	Items []*RedisPreKeySpaceData `json:"Items,omitnil,omitempty" name:"Items"`

	// 采集时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisTopKeyPrefixListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisTopKeyPrefixListResponseParams `json:"Response"`
}

func (r *DescribeRedisTopKeyPrefixListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisTopKeyPrefixListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisUnExpiredKeyStatisticsRequestParams struct {
	// 实例 ID。可通过接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询某个日期最新的任务，如2021-05-27，最早可为前30天的日期。该参数与AsyncRequestId参数不可同时为空。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 异步任务ID。当为空时，选择最近任务的ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 分片节点序号列表。当列表为空时，选择所有分片节点。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

type DescribeRedisUnExpiredKeyStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括 "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 查询某个日期最新的任务，如2021-05-27，最早可为前30天的日期。该参数与AsyncRequestId参数不可同时为空。
	Date *string `json:"Date,omitnil,omitempty" name:"Date"`

	// 异步任务ID。当为空时，选择最近任务的ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 分片节点序号列表。当列表为空时，选择所有分片节点。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

func (r *DescribeRedisUnExpiredKeyStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisUnExpiredKeyStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "Date")
	delete(f, "AsyncRequestId")
	delete(f, "ShardIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRedisUnExpiredKeyStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRedisUnExpiredKeyStatisticsResponseParams struct {
	// 全量Key的聚合分布信息列表。
	SeriesData []*RedisGlobalKeyInfo `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRedisUnExpiredKeyStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRedisUnExpiredKeyStatisticsResponseParams `json:"Response"`
}

func (r *DescribeRedisUnExpiredKeyStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRedisUnExpiredKeyStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSecurityAuditLogDownloadUrlsRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "AsyncRequestId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogDownloadUrlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogDownloadUrlsResponseParams struct {
	// 导出结果的COS链接列表。当结果集很大时，可能会切分为多个url下载。
	Urls []*string `json:"Urls,omitnil,omitempty" name:"Urls"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogDownloadUrlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogDownloadUrlsResponseParams `json:"Response"`
}

func (r *DescribeSecurityAuditLogDownloadUrlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogDownloadUrlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksRequestParams struct {
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志导出任务Id列表。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeSecurityAuditLogExportTasksRequest struct {
	*tchttp.BaseRequest
	
	// 安全审计组Id。
	SecAuditGroupId *string `json:"SecAuditGroupId,omitnil,omitempty" name:"SecAuditGroupId"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 日志导出任务Id列表。
	AsyncRequestIds []*uint64 `json:"AsyncRequestIds,omitnil,omitempty" name:"AsyncRequestIds"`

	// 偏移量，默认0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeSecurityAuditLogExportTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SecAuditGroupId")
	delete(f, "Product")
	delete(f, "AsyncRequestIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSecurityAuditLogExportTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSecurityAuditLogExportTasksResponseParams struct {
	// 安全审计日志导出任务列表。
	Tasks []*SecLogExportTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 安全审计日志导出任务总数。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSecurityAuditLogExportTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSecurityAuditLogExportTasksResponseParams `json:"Response"`
}

func (r *DescribeSecurityAuditLogExportTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSecurityAuditLogExportTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogQueryTimeStatsRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”，截止时间与开始时间的间隔小于7天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Proxy节点ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 实例节点ID。
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 查询类型，目前支持值：mongod，mongos。
	// 其中：
	// mongod - MongoDB的数据存储节点
	// mongos - MongoDB的路由节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeSlowLogQueryTimeStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”，截止时间与开始时间的间隔小于7天。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 TDSQL-C for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Proxy节点ID。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 实例节点ID。
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 查询类型，目前支持值：mongod，mongos。
	// 其中：
	// mongod - MongoDB的数据存储节点
	// mongos - MongoDB的路由节点
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeSlowLogQueryTimeStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogQueryTimeStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "InstanceProxyId")
	delete(f, "InstanceNodeId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogQueryTimeStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogQueryTimeStatsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢日志 top sql 列表。
	Items []*SqlCostDistribution `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogQueryTimeStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogQueryTimeStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogQueryTimeStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogQueryTimeStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-10 12:13:14”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Proxy节点ID。	
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 实列节点ID。	
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 查询类型，目前支持值：mongod，mongos。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeSlowLogTimeSeriesStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间，如“2019-09-10 12:13:14”，结束时间与开始时间的间隔最大可为7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis，"mongodb" - 云数据库 MongoDB，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// Proxy节点ID。	
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 实列节点ID。	
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 查询类型，目前支持值：mongod，mongos。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeSlowLogTimeSeriesStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "InstanceProxyId")
	delete(f, "InstanceNodeId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTimeSeriesStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTimeSeriesStatsResponseParams struct {
	// 柱间单位时间间隔，单位为秒。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 单位时间间隔内慢日志数量统计。
	TimeSeries []*TimeSlice `json:"TimeSeries,omitnil,omitempty" name:"TimeSeries"`

	// 单位时间间隔内的实例 cpu 利用率监控数据。
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogTimeSeriesStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTimeSeriesStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTimeSeriesStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTimeSeriesStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键，默认为QueryTime。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序），默认为DESC。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据库名称数组。
	SchemaList []*SchemaItem `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSlowLogTopSqlsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序键，目前支持 QueryTime,ExecTimes,RowsSent,LockTime以及RowsExamined 等排序键，默认为QueryTime。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 排序方式，支持ASC（升序）以及DESC（降序），默认为DESC。
	OrderBy *string `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数据库名称数组。
	SchemaList []*SchemaItem `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSlowLogTopSqlsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "SortBy")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SchemaList")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogTopSqlsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogTopSqlsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢日志 top sql 列表
	Rows []*SlowLogTopSqlItem `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogTopSqlsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogTopSqlsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogTopSqlsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogTopSqlsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询范围的开始时间，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// SQL模板的MD5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type DescribeSlowLogUserHostStatsRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 查询范围的开始时间，时间格式如：2019-09-10 12:13:14。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 查询范围的结束时间，时间格式如：2019-09-10 12:13:14。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// SQL模板的MD5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

func (r *DescribeSlowLogUserHostStatsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Product")
	delete(f, "Md5")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogUserHostStatsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogUserHostStatsResponseParams struct {
	// 来源地址数目。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 各来源地址的慢日志占比详情列表。
	Items []*SlowLogHost `json:"Items,omitnil,omitempty" name:"Items"`

	// 各来源用户名的慢日志占比详情列表。
	UserNameItems []*SlowLogUser `json:"UserNameItems,omitnil,omitempty" name:"UserNameItems"`

	// 来源用户数目。
	UserTotalCount *int64 `json:"UserTotalCount,omitnil,omitempty" name:"UserTotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogUserHostStatsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogUserHostStatsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogUserHostStatsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogUserHostStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsRequestParams struct {
	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// sql模板的md5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据库列表
	DB []*string `json:"DB,omitnil,omitempty" name:"DB"`

	// 关键字
	Key []*string `json:"Key,omitnil,omitempty" name:"Key"`

	// 用户
	User []*string `json:"User,omitnil,omitempty" name:"User"`

	// IP
	Ip []*string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 耗时区间,耗时区间的左右边界分别对应数组的第0个元素和第一个元素
	Time []*int64 `json:"Time,omitnil,omitempty" name:"Time"`
}

type DescribeSlowLogsRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// sql模板的md5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 开始时间，如“2019-09-10 12:13:14”。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 截止时间，如“2019-09-11 10:13:14”，截止时间与开始时间的间隔小于7天。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数目，默认为20，最大为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 数据库列表
	DB []*string `json:"DB,omitnil,omitempty" name:"DB"`

	// 关键字
	Key []*string `json:"Key,omitnil,omitempty" name:"Key"`

	// 用户
	User []*string `json:"User,omitnil,omitempty" name:"User"`

	// IP
	Ip []*string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 耗时区间,耗时区间的左右边界分别对应数组的第0个元素和第一个元素
	Time []*int64 `json:"Time,omitnil,omitempty" name:"Time"`
}

func (r *DescribeSlowLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "InstanceId")
	delete(f, "Md5")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "DB")
	delete(f, "Key")
	delete(f, "User")
	delete(f, "Ip")
	delete(f, "Time")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSlowLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSlowLogsResponseParams struct {
	// 符合条件的记录总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 慢日志明细
	Rows []*SlowLogInfoItem `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSlowLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSlowLogsResponseParams `json:"Response"`
}

func (r *DescribeSlowLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSlowLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlFiltersRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务ID列表，用于筛选任务列表。
	FilterIds []*int64 `json:"FilterIds,omitnil,omitempty" name:"FilterIds"`

	// 任务状态列表，用于筛选任务列表，取值包括RUNNING - 运行中, FINISHED - 已完成, TERMINATED - 已终止。
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSqlFiltersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务ID列表，用于筛选任务列表。
	FilterIds []*int64 `json:"FilterIds,omitnil,omitempty" name:"FilterIds"`

	// 任务状态列表，用于筛选任务列表，取值包括RUNNING - 运行中, FINISHED - 已完成, TERMINATED - 已终止。
	Statuses []*string `json:"Statuses,omitnil,omitempty" name:"Statuses"`

	// 偏移量，默认为0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSqlFiltersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlFiltersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilterIds")
	delete(f, "Statuses")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSqlFiltersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlFiltersResponseParams struct {
	// 限流任务总数目。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 限流任务列表。
	Items []*SQLFilter `json:"Items,omitnil,omitempty" name:"Items"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSqlFiltersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSqlFiltersResponseParams `json:"Response"`
}

func (r *DescribeSqlFiltersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlFiltersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlTemplateRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeSqlTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeSqlTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Schema")
	delete(f, "SqlText")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSqlTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSqlTemplateResponseParams struct {
	// 数据库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// SQL类型。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// SQL模板内容。
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// SQL模板ID。
	SqlId *int64 `json:"SqlId,omitnil,omitempty" name:"SqlId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSqlTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSqlTemplateResponseParams `json:"Response"`
}

func (r *DescribeSqlTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSqlTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesRequestParams struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemaTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemaTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemaTimeSeriesResponseParams struct {
	// 返回的Top库空间统计信息的时序数据列表。
	TopSpaceSchemaTimeSeries []*SchemaSpaceTimeSeries `json:"TopSpaceSchemaTimeSeries,omitnil,omitempty" name:"TopSpaceSchemaTimeSeries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemaTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemaTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemaTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemaTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	// 其中：
	// DataLength       - 数据长度  
	// IndexLength      - 索引长度  
	// TotalLength      - 总长度  
	// DataFree         - 空闲空间  
	// FragRatio        - 碎片率  
	// TableRows        - 表行数  
	// PhysicalFileSize - 物理文件大小
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceSchemasRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top库数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top库所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	// 其中：
	// DataLength       - 数据长度  
	// IndexLength      - 索引长度  
	// TotalLength      - 总长度  
	// DataFree         - 空闲空间  
	// FragRatio        - 碎片率  
	// TableRows        - 表行数  
	// PhysicalFileSize - 物理文件大小
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceSchemasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceSchemasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceSchemasResponseParams struct {
	// 返回的Top库空间统计信息列表。
	TopSpaceSchemas []*SchemaSpaceData `json:"TopSpaceSchemas,omitnil,omitempty" name:"TopSpaceSchemas"`

	// 采集库空间数据的时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceSchemasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceSchemasResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceSchemasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceSchemasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesRequestParams struct {
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceTableTimeSeriesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize，默认为 PhysicalFileSize。
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 开始日期，如“2021-01-01”，最早为当日的前第29天，默认为截止日期的前第6天。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 截止日期，如“2021-01-01”，最早为当日的前第29天，默认为当日。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTableTimeSeriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTableTimeSeriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTableTimeSeriesResponseParams struct {
	// 返回的Top表空间统计信息的时序数据列表。
	TopSpaceTableTimeSeries []*TableSpaceTimeSeries `json:"TopSpaceTableTimeSeries,omitnil,omitempty" name:"TopSpaceTableTimeSeries"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTableTimeSeriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTableTimeSeriesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTableTimeSeriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTableTimeSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	// 其中：
	// DataLength       - 数据长度  
	// IndexLength      - 索引长度  
	// TotalLength      - 总长度  
	// DataFree         - 空闲空间  
	// FragRatio        - 碎片率  
	// TableRows        - 表行数  
	// PhysicalFileSize - 物理文件大小
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeTopSpaceTablesRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的Top表数量，最大值为100，默认为20。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 筛选Top表所用的排序字段，可选字段包含DataLength、IndexLength、TotalLength、DataFree、FragRatio、TableRows、PhysicalFileSize（仅云数据库 MySQL实例支持），云数据库 MySQL实例默认为 PhysicalFileSize，其他产品实例默认为TotalLength。
	// 其中：
	// DataLength       - 数据长度  
	// IndexLength      - 索引长度  
	// TotalLength      - 总长度  
	// DataFree         - 空闲空间  
	// FragRatio        - 碎片率  
	// TableRows        - 表行数  
	// PhysicalFileSize - 物理文件大小
	SortBy *string `json:"SortBy,omitnil,omitempty" name:"SortBy"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeTopSpaceTablesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "SortBy")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopSpaceTablesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopSpaceTablesResponseParams struct {
	// 返回的Top表空间统计信息列表。
	TopSpaceTables []*TableSpaceData `json:"TopSpaceTables,omitnil,omitempty" name:"TopSpaceTables"`

	// 采集表空间数据的时间戳（秒）。
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopSpaceTablesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopSpaceTablesResponseParams `json:"Response"`
}

func (r *DescribeTopSpaceTablesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopSpaceTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserAutonomyProfileRequestParams struct {
	// 配置类型，为需要配置的功能枚举值，目前包含一下枚举值：AutonomyGlobal（自治功能全局配置）、RedisAutoScaleUp（Redis自治扩容配置）。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeUserAutonomyProfileRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型，为需要配置的功能枚举值，目前包含一下枚举值：AutonomyGlobal（自治功能全局配置）、RedisAutoScaleUp（Redis自治扩容配置）。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 实列ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeUserAutonomyProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAutonomyProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserAutonomyProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserAutonomyProfileResponseParams struct {
	// 配置类型，为需要配置的功能枚举值，目前包含一下枚举值：AutonomyGlobal（自治功能全局配置）、RedisAutoScaleUp（Redis自治扩容配置）。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 更新时间。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 自治用户配置。
	ProfileInfo *AutonomyUserProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserAutonomyProfileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserAutonomyProfileResponseParams `json:"Response"`
}

func (r *DescribeUserAutonomyProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserAutonomyProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL；"dbbrain-mysql" - 自建 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type DescribeUserSqlAdviceRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL；"dbbrain-mysql" - 自建 MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *DescribeUserSqlAdviceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SqlText")
	delete(f, "Schema")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserSqlAdviceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserSqlAdviceResponseParams struct {
	// SQL优化建议，可解析为JSON数组，无需优化时输出为空。
	Advices *string `json:"Advices,omitnil,omitempty" name:"Advices"`

	// SQL优化建议备注，可解析为String数组，无需优化时输出为空。
	Comments *string `json:"Comments,omitnil,omitempty" name:"Comments"`

	// SQL语句。
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 库名。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 相关表的DDL信息，可解析为JSON数组。
	Tables *string `json:"Tables,omitnil,omitempty" name:"Tables"`

	// SQL执行计划，可解析为JSON，无需优化时输出为空。
	SqlPlan *string `json:"SqlPlan,omitnil,omitempty" name:"SqlPlan"`

	// SQL优化后的成本节约详情，可解析为JSON，无需优化时输出为空。
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUserSqlAdviceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserSqlAdviceResponseParams `json:"Response"`
}

func (r *DescribeUserSqlAdviceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserSqlAdviceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagHistoryEventItem struct {
	// 诊断类型。支持值包括"高危账号","自增键耗尽","连接性检查","CPU利用率","死锁","全表扫描","高并发/压力请求","预编译语句过多","内存利用率","Metadata lock","磁盘超限","内存超限","只读锁","只读实例剔除","行锁","活跃会话","慢SQL","数据库快照","磁盘空间利用率","执行计划变化","主从切换","Table open cache命中率低","大表","事务未提交","事务导致复制延迟"等。
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 事件唯一ID 。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 诊断概要。
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 诊断项说明。
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// 实例 ID 。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 保留字段。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 集群ID。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// vip地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// vip端口。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`
}

type EventInfo struct {
	// 事件 ID 。
	EventId *int64 `json:"EventId,omitnil,omitempty" name:"EventId"`

	// 诊断类型。
	DiagType *string `json:"DiagType,omitnil,omitempty" name:"DiagType"`

	// 开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 概要。
	Outline *string `json:"Outline,omitnil,omitempty" name:"Outline"`

	// 严重程度。严重程度分为5级，按影响程度从高至低分别为：1：致命，2：严重，3：告警，4：提示，5：健康。
	Severity *int64 `json:"Severity,omitnil,omitempty" name:"Severity"`

	// 扣分。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// 保留字段。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 告警数目。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type GroupItem struct {
	// 组id。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 组名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 组成员数量。
	MemberCount *int64 `json:"MemberCount,omitnil,omitempty" name:"MemberCount"`
}

type HealthReportTask struct {
	// 异步任务请求 ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 任务的触发来源，支持的取值包括："DAILY_INSPECTION" - 实例巡检；"SCHEDULED" - 定时生成；"MANUAL" - 手动触发。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 任务完成进度，单位%。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始执行时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务完成执行时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务所属实例的基础信息。
	InstanceInfo *InstanceBasicInfo `json:"InstanceInfo,omitnil,omitempty" name:"InstanceInfo"`

	// 健康报告中的健康信息。
	HealthStatus *HealthStatus `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`
}

type HealthScoreInfo struct {
	// 异常详情。
	IssueTypes []*IssueTypeInfo `json:"IssueTypes,omitnil,omitempty" name:"IssueTypes"`

	// 异常事件总数。
	EventsTotalCount *int64 `json:"EventsTotalCount,omitnil,omitempty" name:"EventsTotalCount"`

	// 健康得分。
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 健康等级, 如："HEALTH", "SUB_HEALTH", "RISK", "HIGH_RISK"。
	HealthLevel *string `json:"HealthLevel,omitnil,omitempty" name:"HealthLevel"`
}

type HealthScoreTimeSeriesData struct {
	// 平均得分
	Avg *float64 `json:"Avg,omitnil,omitempty" name:"Avg"`

	// 健康状态
	// 1-health
	// 2-warning
	// 3-critical
	HealthStatus *uint64 `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 指标名称
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 得分序列
	Series []*uint64 `json:"Series,omitnil,omitempty" name:"Series"`

	// 时间序列，单位：毫秒数
	Timestamp []*uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`
}

type HealthStatus struct {
	// 健康分数，满分100。
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 健康等级，取值包括："HEALTH" - 健康；"SUB_HEALTH" - 亚健康；"RISK"- 危险；"HIGH_RISK" - 高危。
	HealthLevel *string `json:"HealthLevel,omitnil,omitempty" name:"HealthLevel"`

	// 总扣分分数。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// 扣分详情。
	ScoreDetails []*ScoreDetail `json:"ScoreDetails,omitnil,omitempty" name:"ScoreDetails"`

	// 健康等级版本，默认为"V1"
	HealthLevelVersion *string `json:"HealthLevelVersion,omitnil,omitempty" name:"HealthLevelVersion"`
}

type IndexesToBuild struct {
	// 索引id，唯一标识一个索引。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 创建索引命令。
	IndexCommand *string `json:"IndexCommand,omitnil,omitempty" name:"IndexCommand"`

	// 索引字符串。
	IndexStr *string `json:"IndexStr,omitnil,omitempty" name:"IndexStr"`

	// 优化级别，1-4，优先级从高到低。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 索引得分。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 签名。
	Signs []*string `json:"Signs,omitnil,omitempty" name:"Signs"`

	// 0-待创建；1-创建中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type IndexesToDrop struct {
	// 索引字符串。
	IndexStr *string `json:"IndexStr,omitnil,omitempty" name:"IndexStr"`

	// 索引得分。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 无效原因。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 删除索引命令。
	IndexCommand *string `json:"IndexCommand,omitnil,omitempty" name:"IndexCommand"`

	// 索引名。
	IndexName *string `json:"IndexName,omitnil,omitempty" name:"IndexName"`
}

type InstanceBasicInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例内网IP。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 实例内网Port。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 实例产品。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 实例引擎版本。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// CPU数量，对于Redis为0。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 实例部署模式。
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 实例内存配置。
	InstanceConf *RedisInstanceConf `json:"InstanceConf,omitnil,omitempty" name:"InstanceConf"`

	// DBbrain是否支持该实例。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 实例内存，单位MB。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 实例地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 实例子网统一ID，对于redis为空字符串。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// 实例私有网络统一ID，对于redis为空字符串。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 实例磁盘容量，对于Redis为0。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`
}

type InstanceConfs struct {
	// 数据库巡检开关, Yes/No。
	DailyInspection *string `json:"DailyInspection,omitnil,omitempty" name:"DailyInspection"`

	// 实例概览开关，Yes/No。
	OverviewDisplay *string `json:"OverviewDisplay,omitnil,omitempty" name:"OverviewDisplay"`

	// redis大key分析的自定义分割符，仅redis使用
	KeyDelimiters []*string `json:"KeyDelimiters,omitnil,omitempty" name:"KeyDelimiters"`

	// 分片节点数量。
	ShardNum *string `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 是否开启大key周期性分析，仅redis产品有效。
	AnalysisTopKey *string `json:"AnalysisTopKey,omitnil,omitempty" name:"AnalysisTopKey"`
}

type InstanceID struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type InstanceInfo struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例名称。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 实例所属地域。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 健康得分。
	HealthScore *int64 `json:"HealthScore,omitnil,omitempty" name:"HealthScore"`

	// 所属产品。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 异常事件数量。
	EventCount *int64 `json:"EventCount,omitnil,omitempty" name:"EventCount"`

	// 实例类型：1:MASTER；2:DR，3：RO，4:SDR。
	InstanceType *int64 `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 核心数。
	Cpu *int64 `json:"Cpu,omitnil,omitempty" name:"Cpu"`

	// 内存，单位MB。
	Memory *int64 `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 硬盘存储，单位GB。
	Volume *int64 `json:"Volume,omitnil,omitempty" name:"Volume"`

	// 数据库版本。
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 内网地址。
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 内网端口。
	Vport *int64 `json:"Vport,omitnil,omitempty" name:"Vport"`

	// 接入来源。
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 分组ID。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组组名。
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 实例状态：0：发货中；1：运行正常；4：销毁中；5：隔离中。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 子网统一ID。
	UniqSubnetId *string `json:"UniqSubnetId,omitnil,omitempty" name:"UniqSubnetId"`

	// cdb类型。
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// cdb实例初始化标志：0：未初始化；1：已初始化。
	InitFlag *int64 `json:"InitFlag,omitnil,omitempty" name:"InitFlag"`

	// 任务状态。
	TaskStatus *int64 `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 私有网络统一ID。
	UniqVpcId *string `json:"UniqVpcId,omitnil,omitempty" name:"UniqVpcId"`

	// 实例巡检/概览的状态。
	InstanceConf *InstanceConfs `json:"InstanceConf,omitnil,omitempty" name:"InstanceConf"`

	// 资源到期时间。
	DeadlineTime *string `json:"DeadlineTime,omitnil,omitempty" name:"DeadlineTime"`

	// 是否是DBbrain支持的实例。
	IsSupported *bool `json:"IsSupported,omitnil,omitempty" name:"IsSupported"`

	// 实例安全审计日志开启状态：ON： 安全审计开启；OFF： 未开启安全审计。
	SecAuditStatus *string `json:"SecAuditStatus,omitnil,omitempty" name:"SecAuditStatus"`

	// 实例审计日志开启状态，ALL_AUDIT： 开启全审计；RULE_AUDIT： 开启规则审计；UNBOUND： 未开启审计。
	AuditPolicyStatus *string `json:"AuditPolicyStatus,omitnil,omitempty" name:"AuditPolicyStatus"`

	// 实例审计日志运行状态：normal： 运行中； paused： 欠费暂停。
	AuditRunningStatus *string `json:"AuditRunningStatus,omitnil,omitempty" name:"AuditRunningStatus"`

	// 内网vip。
	InternalVip *string `json:"InternalVip,omitnil,omitempty" name:"InternalVip"`

	// 内网port。
	InternalVport *int64 `json:"InternalVport,omitnil,omitempty" name:"InternalVport"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 所属集群ID（仅对集群数据库产品该字段非空，如TDSQL-C）。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 所属集群名称（仅对集群数据库产品该字段非空，如TDSQL-C）。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 自建MySQL的Agent状态，"not_deployed" - 未部署，"deploying" - 部署中，"connected" - 连接正常，"deploy_failed" - 连接失败，"monitoring" - 连接正常，"stopped" - 暂停连接，"connect_failed" - 连接失败，unknown - 未知。
	AgentStatus *string `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// 自建MySQL的实例状态，"not_attached" - 未连接，"attached" - 连接正常，"failed" - 连接失败，"stopped" - 停止监控，unknown- 未知。
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

type IssueTypeInfo struct {
	// 指标分类：AVAILABILITY：可用性，MAINTAINABILITY：可维护性，PERFORMANCE，性能，RELIABILITY可靠性。
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 异常事件。
	Events []*EventInfo `json:"Events,omitnil,omitempty" name:"Events"`

	// 异常事件总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type KillMySqlThreadsRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// kill会话任务的阶段，取值包括："Prepare"-准备阶段，"Commit"-提交阶段。
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 需要kill的sql会话ID列表，通过接口[查询实时线程列表](https://cloud.tencent.com/document/product/1130/57824)
	// 此参数用于Prepare阶段。
	Threads []*int64 `json:"Threads,omitnil,omitempty" name:"Threads"`

	// 执行ID，此参数用于Commit阶段。
	SqlExecId *string `json:"SqlExecId,omitnil,omitempty" name:"SqlExecId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 默认是true，会记录下kill的记录；该参数为true, 则在kill操作前校验目标会话是否存在，存在则继续kill，否则取消kill。为了加快kill速度，可设置为false。
	RecordHistory *bool `json:"RecordHistory,omitnil,omitempty" name:"RecordHistory"`
}

type KillMySqlThreadsRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// kill会话任务的阶段，取值包括："Prepare"-准备阶段，"Commit"-提交阶段。
	Stage *string `json:"Stage,omitnil,omitempty" name:"Stage"`

	// 需要kill的sql会话ID列表，通过接口[查询实时线程列表](https://cloud.tencent.com/document/product/1130/57824)
	// 此参数用于Prepare阶段。
	Threads []*int64 `json:"Threads,omitnil,omitempty" name:"Threads"`

	// 执行ID，此参数用于Commit阶段。
	SqlExecId *string `json:"SqlExecId,omitnil,omitempty" name:"SqlExecId"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 默认是true，会记录下kill的记录；该参数为true, 则在kill操作前校验目标会话是否存在，存在则继续kill，否则取消kill。为了加快kill速度，可设置为false。
	RecordHistory *bool `json:"RecordHistory,omitnil,omitempty" name:"RecordHistory"`
}

func (r *KillMySqlThreadsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMySqlThreadsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Stage")
	delete(f, "Threads")
	delete(f, "SqlExecId")
	delete(f, "Product")
	delete(f, "RecordHistory")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillMySqlThreadsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillMySqlThreadsResponseParams struct {
	// kill完成的sql会话ID列表。
	Threads []*int64 `json:"Threads,omitnil,omitempty" name:"Threads"`

	// 执行ID， Prepare阶段的任务输出，用于Commit阶段中指定执行kill操作的会话ID。
	SqlExecId *string `json:"SqlExecId,omitnil,omitempty" name:"SqlExecId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillMySqlThreadsResponse struct {
	*tchttp.BaseResponse
	Response *KillMySqlThreadsResponseParams `json:"Response"`
}

func (r *KillMySqlThreadsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillMySqlThreadsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MailConfiguration struct {
	// 是否开启邮件发送: 0, 否; 1, 是。
	SendMail *int64 `json:"SendMail,omitnil,omitempty" name:"SendMail"`

	// 地域配置, 如["ap-guangzhou", "ap-shanghai"]。巡检的邮件发送模板，配置需要发送巡检邮件的地域；订阅的邮件发送模板，配置当前订阅实例的所属地域。
	Region []*string `json:"Region,omitnil,omitempty" name:"Region"`

	// 包含的健康等级，包括值：HEALTH-健康，SUB_HEALTH-亚健康，RISK-风险，HIGH_RISK-高危。
	HealthStatus []*string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 联系人id, 联系人/联系组不能都为空。
	ContactPerson []*int64 `json:"ContactPerson,omitnil,omitempty" name:"ContactPerson"`

	// 联系组id, 联系人/联系组不能都为空。
	ContactGroup []*int64 `json:"ContactGroup,omitnil,omitempty" name:"ContactGroup"`
}

type MetricThreshold struct {
	// 指标。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 阈值。
	Threshold *int64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// 时间间隔。
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`
}

// Predefined struct for user
type ModifyAlarmPolicyRequestParams struct {
	// 策略类型，固定值：instance， instance-代表实例类型策略
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 策略是否开启，0-代表开启，1-代表关闭
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 实例id列表，可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceIds []*InstanceID `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// User-动态关联该用户所有实例
	// Instance-关联实例列表的实例
	NewProfileLevel *string `json:"NewProfileLevel,omitnil,omitempty" name:"NewProfileLevel"`

	// 新策略名，包含中文英文数字，长度在60个字符以内，不能以“_”开头。
	NewProfileName *string `json:"NewProfileName,omitnil,omitempty" name:"NewProfileName"`

	// 旧策略名，包含中文英文数字，长度在60个字符以内，不能以“_”开头。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 策略类型，固定值：alarm_policy
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 规则类型 0-快速，1-自定义 若值为0，则QuickRule不能为空，若值为1，则Rules 不能为空
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 接受模板
	TemplateInfo []*TemplateInfo `json:"TemplateInfo,omitnil,omitempty" name:"TemplateInfo"`

	// 快速规则  支持包括fatal-致命, critical-严重,
	// warning-告警,
	// information-通知，与Rules互斥
	QuickRule *string `json:"QuickRule,omitnil,omitempty" name:"QuickRule"`

	// 自定义规则，与QuickRule互斥。
	Rules []*AlarmsRules `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type ModifyAlarmPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 策略类型，固定值：instance， instance-代表实例类型策略
	ApplyType *string `json:"ApplyType,omitnil,omitempty" name:"ApplyType"`

	// 策略是否开启，0-代表开启，1-代表关闭
	Enable *int64 `json:"Enable,omitnil,omitempty" name:"Enable"`

	// 实例id列表，可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceIds []*InstanceID `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// User-动态关联该用户所有实例
	// Instance-关联实例列表的实例
	NewProfileLevel *string `json:"NewProfileLevel,omitnil,omitempty" name:"NewProfileLevel"`

	// 新策略名，包含中文英文数字，长度在60个字符以内，不能以“_”开头。
	NewProfileName *string `json:"NewProfileName,omitnil,omitempty" name:"NewProfileName"`

	// 旧策略名，包含中文英文数字，长度在60个字符以内，不能以“_”开头。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 策略类型，固定值：alarm_policy
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 规则类型 0-快速，1-自定义 若值为0，则QuickRule不能为空，若值为1，则Rules 不能为空
	RuleType *int64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 接受模板
	TemplateInfo []*TemplateInfo `json:"TemplateInfo,omitnil,omitempty" name:"TemplateInfo"`

	// 快速规则  支持包括fatal-致命, critical-严重,
	// warning-告警,
	// information-通知，与Rules互斥
	QuickRule *string `json:"QuickRule,omitnil,omitempty" name:"QuickRule"`

	// 自定义规则，与QuickRule互斥。
	Rules []*AlarmsRules `json:"Rules,omitnil,omitempty" name:"Rules"`
}

func (r *ModifyAlarmPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyType")
	delete(f, "Enable")
	delete(f, "InstanceIds")
	delete(f, "NewProfileLevel")
	delete(f, "NewProfileName")
	delete(f, "ProfileName")
	delete(f, "ProfileType")
	delete(f, "Remark")
	delete(f, "RuleType")
	delete(f, "TemplateInfo")
	delete(f, "QuickRule")
	delete(f, "Rules")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAlarmPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAlarmPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAlarmPolicyResponseParams `json:"Response"`
}

func (r *ModifyAlarmPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceRequestParams struct {
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保存总时长，只能是7,30,90,180,365,1095,1825。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保存时长，只能是7,30,90,180,365,1095,1825。
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

type ModifyAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保存总时长，只能是7,30,90,180,365,1095,1825。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保存时长，只能是7,30,90,180,365,1095,1825。
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

func (r *ModifyAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HotLogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAuditServiceResponseParams struct {
	// 审计配置修改结果，0-修改成功,非0-修改失败。
	Success *int64 `json:"Success,omitnil,omitempty" name:"Success"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAuditServiceResponseParams `json:"Response"`
}

func (r *ModifyAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiagDBInstanceConfRequestParams struct {
	// 实例配置，包括巡检、概览开关等。
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil,omitempty" name:"InstanceConfs"`

	// 生效实例地域，固定为"All"，代表全地域。
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定更改巡检状态的实例ID。
	// 可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ModifyDiagDBInstanceConfRequest struct {
	*tchttp.BaseRequest
	
	// 实例配置，包括巡检、概览开关等。
	InstanceConfs *InstanceConfs `json:"InstanceConfs,omitnil,omitempty" name:"InstanceConfs"`

	// 生效实例地域，固定为"All"，代表全地域。
	Regions *string `json:"Regions,omitnil,omitempty" name:"Regions"`

	// 服务产品类型，支持值包括： "mysql" - 云数据库 MySQL， "cynosdb" - 云数据库 CynosDB  for MySQL，"redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 指定更改巡检状态的实例ID。
	// 可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *ModifyDiagDBInstanceConfRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceConfs")
	delete(f, "Regions")
	delete(f, "Product")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDiagDBInstanceConfRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDiagDBInstanceConfResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDiagDBInstanceConfResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDiagDBInstanceConfResponseParams `json:"Response"`
}

func (r *ModifyDiagDBInstanceConfResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDiagDBInstanceConfResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySqlFiltersRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL限流任务ID列表。可通过接口[查询实例SQL限流任务列表
	// ](https://cloud.tencent.com/document/product/1130/72831)获得。
	FilterIds []*int64 `json:"FilterIds,omitnil,omitempty" name:"FilterIds"`

	// 限流任务状态，仅支持传参TERMINATED - 终止。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 通过 [VerifyUserAccount](https://cloud.tencent.com/document/product/1130/72828) 获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type ModifySqlFiltersRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// SQL限流任务ID列表。可通过接口[查询实例SQL限流任务列表
	// ](https://cloud.tencent.com/document/product/1130/72831)获得。
	FilterIds []*int64 `json:"FilterIds,omitnil,omitempty" name:"FilterIds"`

	// 限流任务状态，仅支持传参TERMINATED - 终止。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 通过 [VerifyUserAccount](https://cloud.tencent.com/document/product/1130/72828) 获取有效期为5分钟的会话token，使用后会自动延长token有效期至五分钟后。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *ModifySqlFiltersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySqlFiltersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "FilterIds")
	delete(f, "Status")
	delete(f, "SessionToken")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySqlFiltersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySqlFiltersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySqlFiltersResponse struct {
	*tchttp.BaseResponse
	Response *ModifySqlFiltersResponseParams `json:"Response"`
}

func (r *ModifySqlFiltersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySqlFiltersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserAutonomyProfileRequestParams struct {
	// 配置类型，为需要配置的功能枚举值，目前包含一下枚举值：AutonomyGlobal（自治功能全局配置）、RedisAutoScaleUp（Redis自治扩容配置）
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 自治功能相关配置，标准JSON字符串格式。
	NewProfileInfo *string `json:"NewProfileInfo,omitnil,omitempty" name:"NewProfileInfo"`
}

type ModifyUserAutonomyProfileRequest struct {
	*tchttp.BaseRequest
	
	// 配置类型，为需要配置的功能枚举值，目前包含一下枚举值：AutonomyGlobal（自治功能全局配置）、RedisAutoScaleUp（Redis自治扩容配置）
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，支持值包括： "redis" - 云数据库 Redis。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 自治功能相关配置，标准JSON字符串格式。
	NewProfileInfo *string `json:"NewProfileInfo,omitnil,omitempty" name:"NewProfileInfo"`
}

func (r *ModifyUserAutonomyProfileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserAutonomyProfileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProfileType")
	delete(f, "InstanceId")
	delete(f, "Product")
	delete(f, "NewProfileInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserAutonomyProfileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserAutonomyProfileResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUserAutonomyProfileResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserAutonomyProfileResponseParams `json:"Response"`
}

func (r *ModifyUserAutonomyProfileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserAutonomyProfileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MongoDBIndex struct {
	// 实例id。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 表名。
	Collection *string `json:"Collection,omitnil,omitempty" name:"Collection"`

	// 库名。
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// 优化级别，1-4，优先级从高到低。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 得分。
	Score *int64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 推荐索引列表。
	IndexesToBuild []*IndexesToBuild `json:"IndexesToBuild,omitnil,omitempty" name:"IndexesToBuild"`

	// 无效索引列表。
	IndexesToDrop []*IndexesToDrop `json:"IndexesToDrop,omitnil,omitempty" name:"IndexesToDrop"`
}

type MongoDBProcessItem struct {
	// 是否内部IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsInternalIp *bool `json:"IsInternalIp,omitnil,omitempty" name:"IsInternalIp"`

	// 语句类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 语句详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 客户端ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *float64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 会话ID
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 分片名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShardName *string `json:"ShardName,omitnil,omitempty" name:"ShardName"`

	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`
}

type MongoDBProcessList struct {
	// 列名字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Names []*string `json:"Names,omitnil,omitempty" name:"Names"`

	// 接口返回数据详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*MongoDBProcessItem `json:"Data,omitnil,omitempty" name:"Data"`
}

type MonitorFloatMetric struct {
	// 指标名称。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 指标值。
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type MonitorFloatMetricSeriesData struct {
	// 监控指标。
	Series []*MonitorFloatMetric `json:"Series,omitnil,omitempty" name:"Series"`

	// 监控指标对应的时间戳。
	Timestamp []*int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type MonitorMetric struct {
	// 指标名称。
	Metric *string `json:"Metric,omitnil,omitempty" name:"Metric"`

	// 指标单位。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 指标值。
	Values []*float64 `json:"Values,omitnil,omitempty" name:"Values"`
}

type MonitorMetricSeriesData struct {
	// 监控指标。
	Series []*MonitorMetric `json:"Series,omitnil,omitempty" name:"Series"`

	// 监控指标对应的时间戳。（精度：秒）
	Timestamp []*int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type MySqlProcess struct {
	// 线程ID。
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 线程的操作账号名。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 线程的操作主机地址。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 线程的操作数据库。
	DB *string `json:"DB,omitnil,omitempty" name:"DB"`

	// 线程的操作状态。包含以下枚举值：Sending data​-线程正在处理查询结果， ​Sorting result​-线程正在对查询结果进行排序​，Creating tmp table​-线程正在创建临时表，Altering table​-线程正在执行表结构变更，Updating-线程执行更新中。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 线程的执行类型。包含以下枚举值：Sleep-线程处于空闲状态，Query-线程正在执行一个查询，Connect-从服务器连接到主服务器，Execute-线程正在执行预处理语句。
	Command *string `json:"Command,omitnil,omitempty" name:"Command"`

	// 线程的操作时长，单位秒。
	Time *string `json:"Time,omitnil,omitempty" name:"Time"`

	// 线程的操作语句。
	Info *string `json:"Info,omitnil,omitempty" name:"Info"`

	// sql类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`
}

// Predefined struct for user
type OpenAuditServiceRequestParams struct {
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保存总时长，只能是7,30,90,180,365,1095,1825。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保存时长，只能是7,30,90,180,365,1095,1825。
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

type OpenAuditServiceRequest struct {
	*tchttp.BaseRequest
	
	// 服务产品类型，支持值包括： "dcdb" - 云数据库 Tdsql， "mariadb" - 云数据库 MariaDB。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`

	// 与Product保持一致。如："dcdb" ,"mariadb"。
	NodeRequestType *string `json:"NodeRequestType,omitnil,omitempty" name:"NodeRequestType"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 日志保存总时长，只能是7,30,90,180,365,1095,1825。
	LogExpireDay *int64 `json:"LogExpireDay,omitnil,omitempty" name:"LogExpireDay"`

	// 高频日志保存时长，只能是7,30,90,180,365,1095,1825。
	HotLogExpireDay *int64 `json:"HotLogExpireDay,omitnil,omitempty" name:"HotLogExpireDay"`
}

func (r *OpenAuditServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Product")
	delete(f, "NodeRequestType")
	delete(f, "InstanceId")
	delete(f, "LogExpireDay")
	delete(f, "HotLogExpireDay")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenAuditServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenAuditServiceResponseParams struct {
	// taskId 为0表示开通审计成功，否则开通失败
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenAuditServiceResponse struct {
	*tchttp.BaseResponse
	Response *OpenAuditServiceResponseParams `json:"Response"`
}

func (r *OpenAuditServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenAuditServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Process struct {
	// 会话 ID。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 访问来源，IP 地址和端口号。
	// 格式：IP:Port
	Address *string `json:"Address,omitnil,omitempty" name:"Address"`

	// 文件描述符。
	FileDescriptor *int64 `json:"FileDescriptor,omitnil,omitempty" name:"FileDescriptor"`

	// 会话名称，使用 CLIENT SETNAME 命令设置。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 最后一次执行的命令。
	LastCommand *string `json:"LastCommand,omitnil,omitempty" name:"LastCommand"`

	// 会话存活时间，单位：秒。
	Age *int64 `json:"Age,omitnil,omitempty" name:"Age"`

	// 最后一次执行命令后空闲的时间，单位：秒。
	Idle *int64 `json:"Idle,omitnil,omitempty" name:"Idle"`

	// 会话所属的 Proxy节点 ID。
	ProxyId *string `json:"ProxyId,omitnil,omitempty" name:"ProxyId"`
}

type ProcessStatistic struct {
	// 会话详情数组。
	Items []*SessionItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 总连接数。
	AllConnSum *int64 `json:"AllConnSum,omitnil,omitempty" name:"AllConnSum"`

	// 总活跃连接数。
	ActiveConnSum *int64 `json:"ActiveConnSum,omitnil,omitempty" name:"ActiveConnSum"`
}

type ProfileInfo struct {
	// 语言类型, 包含“zh”-中文，“en”-英文。
	Language *string `json:"Language,omitnil,omitempty" name:"Language"`

	// 邮件模板的内容。
	MailConfiguration *MailConfiguration `json:"MailConfiguration,omitnil,omitempty" name:"MailConfiguration"`
}

type ReceiveInfo struct {
	// 接收组
	ReceiveGroup []*int64 `json:"ReceiveGroup,omitnil,omitempty" name:"ReceiveGroup"`

	// 最后接收时间，格式: "HH:mm:ss"
	EndReceiveTime *string `json:"EndReceiveTime,omitnil,omitempty" name:"EndReceiveTime"`

	// 接收名
	ReceiveName *string `json:"ReceiveName,omitnil,omitempty" name:"ReceiveName"`

	// 推送渠道
	SendChannel []*int64 `json:"SendChannel,omitnil,omitempty" name:"SendChannel"`

	// 开始时间，格式: "HH:mm:ss"
	StartReceiveTime *string `json:"StartReceiveTime,omitnil,omitempty" name:"StartReceiveTime"`

	// 接收用户列表
	ReceiveUin []*ReceiveUin `json:"ReceiveUin,omitnil,omitempty" name:"ReceiveUin"`
}

type ReceiveUin struct {
	// 用户名
	UinName *string `json:"UinName,omitnil,omitempty" name:"UinName"`

	// 用户id
	Uin *string `json:"Uin,omitnil,omitempty" name:"Uin"`
}

type RedisBigKeyTask struct {
	// 异步任务请求 ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务状态。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 任务执行进度。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 任务包含的分片节点序号列表。
	ShardIds []*int64 `json:"ShardIds,omitnil,omitempty" name:"ShardIds"`
}

type RedisCmdInfo struct {
	// redis命令
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 命令次数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type RedisCostCmd struct {
	// 命令
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 最大cost
	MaxCost *uint64 `json:"MaxCost,omitnil,omitempty" name:"MaxCost"`
}

type RedisGlobalKeyInfo struct {
	// 占用内存大小，单位Byte。
	Capacity *int64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`

	// Key个数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 剩余过期时间范围的结束时间，当小于0时，代表已过期时间，单位：小时。当RangeMin与RangeMax同时为空时，代表未设置过期时间。当RangeMax为空时，代表剩余过期时间大于等于RangeMin小时。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeMax *int64 `json:"RangeMax,omitnil,omitempty" name:"RangeMax"`

	// 剩余过期时间范围的起始时间，当小于0时，代表已过期时间，单位：小时。当RangeMin与RangeMax同时为空时，代表未设置过期时间。当RangeMin为空时，代表已过期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangeMin *int64 `json:"RangeMin,omitnil,omitempty" name:"RangeMin"`
}

type RedisInstanceConf struct {
	// 副本数量
	ReplicasNum *string `json:"ReplicasNum,omitnil,omitempty" name:"ReplicasNum"`

	// 分片数量
	ShardNum *string `json:"ShardNum,omitnil,omitempty" name:"ShardNum"`

	// 分片内存大小，单位MB
	ShardSize *string `json:"ShardSize,omitnil,omitempty" name:"ShardSize"`
}

type RedisKeySpaceData struct {
	// key名。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// key类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// key编码方式。包括 int、string、linkedlist、hashtable、skiplist、zipmap、ziplist、intset、quicklist、listpack。
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// key过期时间戳（毫秒），0代表未设置过期时间。
	ExpireTime *int64 `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// key内存大小，单位Byte。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 元素个数。
	ItemCount *int64 `json:"ItemCount,omitnil,omitempty" name:"ItemCount"`

	// 最大元素长度。
	MaxElementSize *int64 `json:"MaxElementSize,omitnil,omitempty" name:"MaxElementSize"`

	// 平均元素长度。
	AveElementSize *int64 `json:"AveElementSize,omitnil,omitempty" name:"AveElementSize"`

	// 所属分片序号。
	ShardId *string `json:"ShardId,omitnil,omitempty" name:"ShardId"`

	// key所属数据库编号。
	Db *int64 `json:"Db,omitnil,omitempty" name:"Db"`
}

type RedisMetricTopProxiesData struct {
	// host
	// 注意：此字段可能返回 null，表示取不到有效值。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// Proxy Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceProxyId *string `json:"InstanceProxyId,omitnil,omitempty" name:"InstanceProxyId"`

	// 端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 最新的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 时间（秒）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp []*uint64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 序列数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Series []*MonitorFloatMetric `json:"Series,omitnil,omitempty" name:"Series"`
}

type RedisPreKeySpaceData struct {
	// 平均元素长度。
	AveElementSize *int64 `json:"AveElementSize,omitnil,omitempty" name:"AveElementSize"`

	// 总占用内存（Byte）。
	Length *int64 `json:"Length,omitnil,omitempty" name:"Length"`

	// key前缀。
	KeyPreIndex *string `json:"KeyPreIndex,omitnil,omitempty" name:"KeyPreIndex"`

	// 元素数量。
	ItemCount *int64 `json:"ItemCount,omitnil,omitempty" name:"ItemCount"`

	// key个数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 最大元素长度。
	MaxElementSize *int64 `json:"MaxElementSize,omitnil,omitempty" name:"MaxElementSize"`
}

type SQLFilter struct {
	// 任务ID。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务状态，取值包括RUNNING - 运行中, FINISHED - 已完成, TERMINATED - 已终止。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// SQL类型，取值包括SELECT, UPDATE, DELETE, INSERT, REPLACE。
	SqlType *string `json:"SqlType,omitnil,omitempty" name:"SqlType"`

	// 筛选SQL的关键词，多个关键词用英文逗号拼接。
	OriginKeys *string `json:"OriginKeys,omitnil,omitempty" name:"OriginKeys"`

	// 筛选SQL的规则。
	OriginRule *string `json:"OriginRule,omitnil,omitempty" name:"OriginRule"`

	// 已拒绝SQL数目。
	RejectedSqlCount *int64 `json:"RejectedSqlCount,omitnil,omitempty" name:"RejectedSqlCount"`

	// 当前并发数。
	CurrentConcurrency *int64 `json:"CurrentConcurrency,omitnil,omitempty" name:"CurrentConcurrency"`

	// 最大并发数。
	MaxConcurrency *int64 `json:"MaxConcurrency,omitnil,omitempty" name:"MaxConcurrency"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 当前时间。
	CurrentTime *string `json:"CurrentTime,omitnil,omitempty" name:"CurrentTime"`

	// 限流过期时间。
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`
}

type SchemaItem struct {
	// 数据库名称
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

type SchemaSpaceData struct {
	// 库名。
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 数据空间（MB）。
	DataLength *float64 `json:"DataLength,omitnil,omitempty" name:"DataLength"`

	// 索引空间（MB）。
	IndexLength *float64 `json:"IndexLength,omitnil,omitempty" name:"IndexLength"`

	// 碎片空间（MB）。
	DataFree *float64 `json:"DataFree,omitnil,omitempty" name:"DataFree"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitnil,omitempty" name:"TotalLength"`

	// 碎片率（%）。
	FragRatio *float64 `json:"FragRatio,omitnil,omitempty" name:"FragRatio"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitnil,omitempty" name:"TableRows"`

	// 库中所有表对应的独立物理文件大小加和（MB）。
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil,omitempty" name:"PhysicalFileSize"`
}

type SchemaSpaceTimeSeries struct {
	// 库名
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 单位时间间隔内的空间指标数据。
	SeriesData *MonitorMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`
}

type ScoreDetail struct {
	// 扣分项分类，取值包括：可用性、可维护性、性能及可靠性。
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 扣分总分。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`

	// 扣分总分上限。
	ScoreLostMax *int64 `json:"ScoreLostMax,omitnil,omitempty" name:"ScoreLostMax"`

	// 扣分项列表。
	Items []*ScoreItem `json:"Items,omitnil,omitempty" name:"Items"`
}

type ScoreItem struct {
	// 异常诊断项名称。
	DiagItem *string `json:"DiagItem,omitnil,omitempty" name:"DiagItem"`

	// 诊断项分类，取值包括：可用性、可维护性、性能及可靠性。
	IssueType *string `json:"IssueType,omitnil,omitempty" name:"IssueType"`

	// 健康等级，取值包括：信息、提示、告警、严重、致命。
	TopSeverity *string `json:"TopSeverity,omitnil,omitempty" name:"TopSeverity"`

	// 该异常诊断项出现次数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 扣分分数。
	ScoreLost *int64 `json:"ScoreLost,omitnil,omitempty" name:"ScoreLost"`
}

type SecLogExportTaskInfo struct {
	// 异步任务Id。
	AsyncRequestId *uint64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 任务开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务执行进度。
	Progress *uint64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 导出日志开始时间。
	LogStartTime *string `json:"LogStartTime,omitnil,omitempty" name:"LogStartTime"`

	// 导出日志结束时间。
	LogEndTime *string `json:"LogEndTime,omitnil,omitempty" name:"LogEndTime"`

	// 日志文件总大小，单位KB。
	TotalSize *uint64 `json:"TotalSize,omitnil,omitempty" name:"TotalSize"`

	// 风险等级列表。0 无风险；1 低风险；2 中风险；3 高风险。
	DangerLevels []*uint64 `json:"DangerLevels,omitnil,omitempty" name:"DangerLevels"`
}

type SessionItem struct {
	// 访问来源。
	Ip *string `json:"Ip,omitnil,omitempty" name:"Ip"`

	// 当前访问来源活跃连接数
	ActiveConn *string `json:"ActiveConn,omitnil,omitempty" name:"ActiveConn"`

	// 当前访问来源总连接数
	AllConn *int64 `json:"AllConn,omitnil,omitempty" name:"AllConn"`
}

type SlowLogAgg struct {
	// 命令模板。
	Cmd *string `json:"Cmd,omitnil,omitempty" name:"Cmd"`

	// 命令详情。
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 执行次数。
	ExecTimes *int64 `json:"ExecTimes,omitnil,omitempty" name:"ExecTimes"`

	// 总耗时。单位：s
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 平均执行时间。单位：s
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitnil,omitempty" name:"QueryTimeAvg"`

	// 最大执行时间。单位：s
	QueryTimeMax *float64 `json:"QueryTimeMax,omitnil,omitempty" name:"QueryTimeMax"`

	// 最小执行时间。单位：s
	QueryTimeMin *float64 `json:"QueryTimeMin,omitnil,omitempty" name:"QueryTimeMin"`

	// 总耗时占比。单位：%
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitnil,omitempty" name:"QueryTimeRatio"`
}

type SlowLogHost struct {
	// 来源地址。
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// 该来源地址的慢日志数目占总数目的比例，单位%。
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 该来源地址的慢日志数目。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type SlowLogInfoItem struct {
	// 慢日志开始时间
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// sql语句
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 数据库
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// User来源
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// IP来源
	UserHost *string `json:"UserHost,omitnil,omitempty" name:"UserHost"`

	// 执行时间,单位秒
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 锁时间,单位秒
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// 扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// 返回行数
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`
}

type SlowLogTopSqlItem struct {
	// sql总锁等待时间，单位秒
	LockTime *float64 `json:"LockTime,omitnil,omitempty" name:"LockTime"`

	// 最大锁等待时间，单位秒
	LockTimeMax *float64 `json:"LockTimeMax,omitnil,omitempty" name:"LockTimeMax"`

	// 最小锁等待时间，单位秒
	LockTimeMin *float64 `json:"LockTimeMin,omitnil,omitempty" name:"LockTimeMin"`

	// 总扫描行数
	RowsExamined *int64 `json:"RowsExamined,omitnil,omitempty" name:"RowsExamined"`

	// 最大扫描行数
	RowsExaminedMax *int64 `json:"RowsExaminedMax,omitnil,omitempty" name:"RowsExaminedMax"`

	// 最小扫描行数
	RowsExaminedMin *int64 `json:"RowsExaminedMin,omitnil,omitempty" name:"RowsExaminedMin"`

	// 总耗时，单位秒
	QueryTime *float64 `json:"QueryTime,omitnil,omitempty" name:"QueryTime"`

	// 最大执行时间，单位秒
	QueryTimeMax *float64 `json:"QueryTimeMax,omitnil,omitempty" name:"QueryTimeMax"`

	// 最小执行时间，单位秒
	QueryTimeMin *float64 `json:"QueryTimeMin,omitnil,omitempty" name:"QueryTimeMin"`

	// 总返回行数
	RowsSent *int64 `json:"RowsSent,omitnil,omitempty" name:"RowsSent"`

	// 最大返回行数
	RowsSentMax *int64 `json:"RowsSentMax,omitnil,omitempty" name:"RowsSentMax"`

	// 最小返回行数
	RowsSentMin *int64 `json:"RowsSentMin,omitnil,omitempty" name:"RowsSentMin"`

	// 执行次数
	ExecTimes *int64 `json:"ExecTimes,omitnil,omitempty" name:"ExecTimes"`

	// sql模板
	SqlTemplate *string `json:"SqlTemplate,omitnil,omitempty" name:"SqlTemplate"`

	// 带参数SQL（随机）
	SqlText *string `json:"SqlText,omitnil,omitempty" name:"SqlText"`

	// 数据库名
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 总耗时占比，单位%
	QueryTimeRatio *float64 `json:"QueryTimeRatio,omitnil,omitempty" name:"QueryTimeRatio"`

	// sql总锁等待时间占比，单位%
	LockTimeRatio *float64 `json:"LockTimeRatio,omitnil,omitempty" name:"LockTimeRatio"`

	// 总扫描行数占比，单位%
	RowsExaminedRatio *float64 `json:"RowsExaminedRatio,omitnil,omitempty" name:"RowsExaminedRatio"`

	// 总返回行数占比，单位%
	RowsSentRatio *float64 `json:"RowsSentRatio,omitnil,omitempty" name:"RowsSentRatio"`

	// 平均执行时间，单位秒
	QueryTimeAvg *float64 `json:"QueryTimeAvg,omitnil,omitempty" name:"QueryTimeAvg"`

	// 平均返回行数
	RowsSentAvg *float64 `json:"RowsSentAvg,omitnil,omitempty" name:"RowsSentAvg"`

	// 平均锁等待时间，单位秒
	LockTimeAvg *float64 `json:"LockTimeAvg,omitnil,omitempty" name:"LockTimeAvg"`

	// 平均扫描行数
	RowsExaminedAvg *float64 `json:"RowsExaminedAvg,omitnil,omitempty" name:"RowsExaminedAvg"`

	// SQL模板的MD5值
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`
}

type SlowLogUser struct {
	// 来源用户名。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 该来源用户名的慢日志数目占总数目的比例，单位%。
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`

	// 该来源用户名的慢日志数目。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type SqlCostDistribution struct {
	// sql条数。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 分段耗时下边界，单位是秒。
	From *float64 `json:"From,omitnil,omitempty" name:"From"`

	// 分段耗时上边界，单位是秒。
	To *float64 `json:"To,omitnil,omitempty" name:"To"`

	// 耗时占比。单位（%）
	Ratio *float64 `json:"Ratio,omitnil,omitempty" name:"Ratio"`
}

type StatDimension struct {
	// 维度名称，目前仅支持：SqlTag。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// SQL 标签过滤与统计信息
	// 示例：
	// 
	// 示例 1：[p=position] 统计包含 p=position 标签的 SQL 会话。
	// 示例 2：[p] 统计包含 p 标签的 SQL 会话。
	// 示例 3：[p=position, c=idCard] 统计同时包含 p=position 标签和 c=idCard 标签的 SQL 会话。
	Data []*string `json:"Data,omitnil,omitempty" name:"Data"`
}

type StatisticDataInfo struct {
	// 统计维度的值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 平均时间。
	TimeAvg *float64 `json:"TimeAvg,omitnil,omitempty" name:"TimeAvg"`

	// 总时间。
	TimeSum *float64 `json:"TimeSum,omitnil,omitempty" name:"TimeSum"`

	// 数量。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type StatisticInfo struct {
	// 统计分析的维度。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 统计分析的维度下的统计数据详情。
	Data []*StatisticDataInfo `json:"Data,omitnil,omitempty" name:"Data"`
}

type Table struct {
	// 库名。
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 表名。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitnil,omitempty" name:"TableRows"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitnil,omitempty" name:"TotalLength"`
}

type TableSpaceData struct {
	// 表名。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 数据空间（MB）。
	DataLength *float64 `json:"DataLength,omitnil,omitempty" name:"DataLength"`

	// 索引空间（MB）。
	IndexLength *float64 `json:"IndexLength,omitnil,omitempty" name:"IndexLength"`

	// 碎片空间（MB）。
	DataFree *float64 `json:"DataFree,omitnil,omitempty" name:"DataFree"`

	// 总使用空间（MB）。
	TotalLength *float64 `json:"TotalLength,omitnil,omitempty" name:"TotalLength"`

	// 碎片率（%）。
	FragRatio *float64 `json:"FragRatio,omitnil,omitempty" name:"FragRatio"`

	// 行数。
	TableRows *int64 `json:"TableRows,omitnil,omitempty" name:"TableRows"`

	// 表对应的独立物理文件大小（MB）。
	PhysicalFileSize *float64 `json:"PhysicalFileSize,omitnil,omitempty" name:"PhysicalFileSize"`
}

type TableSpaceTimeSeries struct {
	// 表名。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 库名。
	TableSchema *string `json:"TableSchema,omitnil,omitempty" name:"TableSchema"`

	// 库表的存储引擎。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 单位时间间隔内的空间指标数据。
	SeriesData *MonitorFloatMetricSeriesData `json:"SeriesData,omitnil,omitempty" name:"SeriesData"`
}

type TaskInfo struct {
	// 异步任务 ID。
	AsyncRequestId *int64 `json:"AsyncRequestId,omitnil,omitempty" name:"AsyncRequestId"`

	// 当前实例所有 proxy 列表。
	InstProxyList []*string `json:"InstProxyList,omitnil,omitempty" name:"InstProxyList"`

	// 当前实例所有 proxy 数量。
	InstProxyCount *int64 `json:"InstProxyCount,omitnil,omitempty" name:"InstProxyCount"`

	// 任务创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务启动时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务的状态，支持的取值包括："created" - 新建；"chosen" - 待执行； "running" - 执行中；"failed" - 失败；"finished" - 已完成。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 完成 kill 任务的 proxyId。
	FinishedProxyList []*string `json:"FinishedProxyList,omitnil,omitempty" name:"FinishedProxyList"`

	// kill 任务实行失败的 proxyId。
	FailedProxyList []*string `json:"FailedProxyList,omitnil,omitempty" name:"FailedProxyList"`

	// 任务结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务执行进度。
	Progress *int64 `json:"Progress,omitnil,omitempty" name:"Progress"`

	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TemplateInfo struct {
	// 模板id,通过接口[通知模板查询](https://cloud.tencent.com/document/product/1130/97726)获得。
	TemplateId *string `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板名,通过接口[通知模板查询](https://cloud.tencent.com/document/product/1130/97726)获得。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`
}

type TimeSlice struct {
	// 总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 统计开始时间
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`
}

type TopHotKeys struct {
	// 访问频次。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 热Key所属数据库。
	Db *string `json:"Db,omitnil,omitempty" name:"Db"`

	// Redis节点。
	InstanceNodeId *string `json:"InstanceNodeId,omitnil,omitempty" name:"InstanceNodeId"`

	// 热Key。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 数据类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type UpdateAgentSwitchRequestParams struct {
	// Agent标识。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 停止或重连Agent，支持值包括："on" - 重连Agent， "off" - 停止Agent。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 服务产品类型，仅支持 "dbbrain-mysql" - 自建MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type UpdateAgentSwitchRequest struct {
	*tchttp.BaseRequest
	
	// Agent标识。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 停止或重连Agent，支持值包括："on" - 重连Agent， "off" - 停止Agent。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 服务产品类型，仅支持 "dbbrain-mysql" - 自建MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *UpdateAgentSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAgentSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "Switch")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAgentSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAgentSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAgentSwitchResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAgentSwitchResponseParams `json:"Response"`
}

func (r *UpdateAgentSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAgentSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateMonitorSwitchRequestParams struct {
	// 停止或重连Agent实例，支持值包括："on" - 重连实例， "off" - 停止实例。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，仅支持 "dbbrain-mysql" - 自建MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type UpdateMonitorSwitchRequest struct {
	*tchttp.BaseRequest
	
	// 停止或重连Agent实例，支持值包括："on" - 重连实例， "off" - 停止实例。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务产品类型，仅支持 "dbbrain-mysql" - 自建MySQL。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *UpdateMonitorSwitchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateMonitorSwitchRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Switch")
	delete(f, "InstanceId")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateMonitorSwitchRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateMonitorSwitchResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateMonitorSwitchResponse struct {
	*tchttp.BaseResponse
	Response *UpdateMonitorSwitchResponseParams `json:"Response"`
}

func (r *UpdateMonitorSwitchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateMonitorSwitchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserProfile struct {
	// 配置的id。
	ProfileId *string `json:"ProfileId,omitnil,omitempty" name:"ProfileId"`

	// 配置类型，支持值包括："dbScan_mail_configuration" - 数据库巡检邮件配置，"scheduler_mail_configuration" - 定期生成邮件配置。
	ProfileType *string `json:"ProfileType,omitnil,omitempty" name:"ProfileType"`

	// 配置级别，支持值包括："User" - 用户级别，"Instance" - 实例级别，其中数据库巡检邮件配置为用户级别，定期生成邮件配置为实例级别。
	ProfileLevel *string `json:"ProfileLevel,omitnil,omitempty" name:"ProfileLevel"`

	// 配置名称。
	ProfileName *string `json:"ProfileName,omitnil,omitempty" name:"ProfileName"`

	// 配置详情。
	ProfileInfo *ProfileInfo `json:"ProfileInfo,omitnil,omitempty" name:"ProfileInfo"`
}

// Predefined struct for user
type VerifyUserAccountRequestParams struct {
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库账号名。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库账号密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

type VerifyUserAccountRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID。可通过 [DescribeDiagDBInstances](https://cloud.tencent.com/document/api/1130/57798) 接口获取。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库账号名。
	User *string `json:"User,omitnil,omitempty" name:"User"`

	// 数据库账号密码。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 服务产品类型，支持值："mysql" - 云数据库 MySQL；"cynosdb" - 云数据库 TDSQL-C for MySQL，默认为"mysql"。
	Product *string `json:"Product,omitnil,omitempty" name:"Product"`
}

func (r *VerifyUserAccountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyUserAccountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "Product")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyUserAccountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyUserAccountResponseParams struct {
	// 会话token，有效期为5分钟。
	SessionToken *string `json:"SessionToken,omitnil,omitempty" name:"SessionToken"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyUserAccountResponse struct {
	*tchttp.BaseResponse
	Response *VerifyUserAccountResponseParams `json:"Response"`
}

func (r *VerifyUserAccountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyUserAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}