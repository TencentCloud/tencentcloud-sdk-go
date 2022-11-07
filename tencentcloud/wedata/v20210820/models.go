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

package v20210820

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlarmEventInfo struct {
	// 告警ID
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 告警时间
	AlarmTime *string `json:"AlarmTime,omitempty" name:"AlarmTime"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 规则名称
	RegularName *string `json:"RegularName,omitempty" name:"RegularName"`

	// 告警级别,0表示普通，1表示重要，2表示紧急
	AlarmLevel *uint64 `json:"AlarmLevel,omitempty" name:"AlarmLevel"`

	// 告警指标,0表示任务失败，1表示任务运行超时，2表示任务停止，3表示任务暂停
	AlarmIndicator *uint64 `json:"AlarmIndicator,omitempty" name:"AlarmIndicator"`

	// 告警方式,多个用逗号隔开（1:邮件，2:短信，3:微信，4:语音，5:代表企业微信，6:http）
	AlarmWay *uint64 `json:"AlarmWay,omitempty" name:"AlarmWay"`

	// 告警接收人Id，多个用逗号隔开
	AlarmRecipientId *string `json:"AlarmRecipientId,omitempty" name:"AlarmRecipientId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 告警指标描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicatorDesc *string `json:"AlarmIndicatorDesc,omitempty" name:"AlarmIndicatorDesc"`

	// 指标阈值，1表示离线任务第一次运行失败，2表示离线任务所有重试完成后失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *uint64 `json:"TriggerType,omitempty" name:"TriggerType"`

	// 预计的超时时间，分钟级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	EstimatedTime *uint64 `json:"EstimatedTime,omitempty" name:"EstimatedTime"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 0：部分成功，1：全部成功，2：全部失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSendSuccess *uint64 `json:"IsSendSuccess,omitempty" name:"IsSendSuccess"`

	// 消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// 阈值计算算子，1 : 大于 2 ：小于
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *int64 `json:"Operator,omitempty" name:"Operator"`
}

type AlarmInfo struct {
	// 关联任务id
	TaskIds *string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 告警类别；failure表示失败告警；overtime表示超时告警
	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`

	// 告警方式；SMS表示短信；Email表示邮件；HTTP 表示接口方式；Wechat表示微信方式
	AlarmWay *string `json:"AlarmWay,omitempty" name:"AlarmWay"`

	// 告警接收人，多个告警接收人以;分割
	AlarmRecipient *string `json:"AlarmRecipient,omitempty" name:"AlarmRecipient"`

	// 告警接收人id，多个告警接收人id以;分割
	AlarmRecipientId *string `json:"AlarmRecipientId,omitempty" name:"AlarmRecipientId"`

	// 预计运行的小时，取值范围0-23
	Hours *uint64 `json:"Hours,omitempty" name:"Hours"`

	// 预计运行分钟，取值范围0-59
	Minutes *uint64 `json:"Minutes,omitempty" name:"Minutes"`

	// 告警出发时机；1表示第一次运行失败；2表示所有重试完成后失败；
	TriggerType *uint64 `json:"TriggerType,omitempty" name:"TriggerType"`

	// 告警信息id
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 告警状态设置；1表示可用；0表示不可用，默认可用
	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

type AlarmReceiverInfo struct {
	// 告警ID
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 告警接收人ID
	AlarmReceiver *string `json:"AlarmReceiver,omitempty" name:"AlarmReceiver"`

	// 邮件，0：未设置，1：成功，2：失败
	Email *uint64 `json:"Email,omitempty" name:"Email"`

	// 短信，0：未设置，1：成功，2：失败
	Sms *uint64 `json:"Sms,omitempty" name:"Sms"`

	// 微信，0：未设置，1：成功，2：失败
	Wechat *uint64 `json:"Wechat,omitempty" name:"Wechat"`

	// 电话，0：未设置，1：成功，2：失败
	Voice *uint64 `json:"Voice,omitempty" name:"Voice"`

	// 企业微信，0：未设置，1：成功，2：失败
	Wecom *uint64 `json:"Wecom,omitempty" name:"Wecom"`

	// http，0：未设置，1：成功，2：失败
	Http *uint64 `json:"Http,omitempty" name:"Http"`
}

// Predefined struct for user
type BatchCreateIntegrationTaskAlarmsRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitempty" name:"TaskAlarmInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchCreateIntegrationTaskAlarmsRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitempty" name:"TaskAlarmInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchCreateIntegrationTaskAlarmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateIntegrationTaskAlarmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskAlarmInfo")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateIntegrationTaskAlarmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateIntegrationTaskAlarmsResponseParams struct {
	// 操作成功的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchCreateIntegrationTaskAlarmsResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateIntegrationTaskAlarmsResponseParams `json:"Response"`
}

func (r *BatchCreateIntegrationTaskAlarmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateIntegrationTaskAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchDeleteIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchDeleteIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteIntegrationTasksResponseParams `json:"Response"`
}

func (r *BatchDeleteIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteTasksNewRequestParams struct {
	// 批量删除的任务TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchDeleteTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 批量删除的任务TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchDeleteTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "DeleteMode")
	delete(f, "EnableNotify")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteTasksNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchDeleteTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteTasksNewResponseParams `json:"Response"`
}

func (r *BatchDeleteTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchForceSuccessIntegrationTaskInstancesRequestParams struct {
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchForceSuccessIntegrationTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchForceSuccessIntegrationTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchForceSuccessIntegrationTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchForceSuccessIntegrationTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchForceSuccessIntegrationTaskInstancesResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchForceSuccessIntegrationTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *BatchForceSuccessIntegrationTaskInstancesResponseParams `json:"Response"`
}

func (r *BatchForceSuccessIntegrationTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchForceSuccessIntegrationTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchKillIntegrationTaskInstancesRequestParams struct {
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchKillIntegrationTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchKillIntegrationTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchKillIntegrationTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchKillIntegrationTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchKillIntegrationTaskInstancesResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchKillIntegrationTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *BatchKillIntegrationTaskInstancesResponseParams `json:"Response"`
}

func (r *BatchKillIntegrationTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchKillIntegrationTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchMakeUpIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 补数据开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补数据结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchMakeUpIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 补数据开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补数据结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchMakeUpIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchMakeUpIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchMakeUpIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchMakeUpIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchMakeUpIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchMakeUpIntegrationTasksResponseParams `json:"Response"`
}

func (r *BatchMakeUpIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchMakeUpIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyOwnersNewRequestParams struct {
	// 需要更新责任人的TaskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 需要更新的责任人
	Owners *string `json:"Owners,omitempty" name:"Owners"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchModifyOwnersNewRequest struct {
	*tchttp.BaseRequest
	
	// 需要更新责任人的TaskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 需要更新的责任人
	Owners *string `json:"Owners,omitempty" name:"Owners"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchModifyOwnersNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyOwnersNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "Owners")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyOwnersNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyOwnersNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchModifyOwnersNewResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyOwnersNewResponseParams `json:"Response"`
}

func (r *BatchModifyOwnersNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyOwnersNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchOperateResult struct {
	// 批量操作成功数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 批量操作失败数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 批量操作的总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type BatchRerunIntegrationTaskInstancesRequestParams struct {
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchRerunIntegrationTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchRerunIntegrationTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRerunIntegrationTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRerunIntegrationTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRerunIntegrationTaskInstancesResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchRerunIntegrationTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *BatchRerunIntegrationTaskInstancesResponseParams `json:"Response"`
}

func (r *BatchRerunIntegrationTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRerunIntegrationTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchResult struct {
	// 正在运行的任务数
	Running *int64 `json:"Running,omitempty" name:"Running"`

	// 执行成功的任务数
	Success *int64 `json:"Success,omitempty" name:"Success"`

	// 执行失败的任务数
	Failed *int64 `json:"Failed,omitempty" name:"Failed"`

	// 总任务数
	Total *int64 `json:"Total,omitempty" name:"Total"`
}

// Predefined struct for user
type BatchResumeIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchResumeIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchResumeIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchResumeIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchResumeIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchResumeIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchResumeIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchResumeIntegrationTasksResponseParams `json:"Response"`
}

func (r *BatchResumeIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchResumeIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchReturn struct {
	// 执行结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 执行情况备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitempty" name:"ErrorDesc"`

	// 执行情况id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitempty" name:"ErrorId"`
}

// Predefined struct for user
type BatchStartIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchStartIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchStartIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStartIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStartIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStartIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchStartIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchStartIntegrationTasksResponseParams `json:"Response"`
}

func (r *BatchStartIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStartIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchStopIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchStopIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStopIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchStopIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchStopIntegrationTasksResponseParams `json:"Response"`
}

func (r *BatchStopIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopTasksNewRequestParams struct {
	// 批量停止任务的TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchStopTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 批量停止任务的TaskId
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchStopTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStopTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopTasksNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchStopTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *BatchStopTasksNewResponseParams `json:"Response"`
}

func (r *BatchStopTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchSuspendIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchSuspendIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchSuspendIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchSuspendIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "TaskType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchSuspendIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchSuspendIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchSuspendIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchSuspendIntegrationTasksResponseParams `json:"Response"`
}

func (r *BatchSuspendIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchSuspendIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchUpdateIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 责任人（多个责任人用小写分号隔开；离线任务传入的是账号名，实时任务传入的是账号id）
	Incharge *string `json:"Incharge,omitempty" name:"Incharge"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type BatchUpdateIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 责任人（多个责任人用小写分号隔开；离线任务传入的是账号名，实时任务传入的是账号id）
	Incharge *string `json:"Incharge,omitempty" name:"Incharge"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *BatchUpdateIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchUpdateIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "Incharge")
	delete(f, "TaskType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchUpdateIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchUpdateIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BatchUpdateIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchUpdateIntegrationTasksResponseParams `json:"Response"`
}

func (r *BatchUpdateIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchUpdateIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BytesSpeed struct {
	// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 速度值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*SpeedValue `json:"Values,omitempty" name:"Values"`
}

type CanvasInfo struct {
	// 画布任务信息
	TasksList []*TaskCanvasInfo `json:"TasksList,omitempty" name:"TasksList"`

	// 画布任务链接信息
	LinksList []*TaskLinkInfo `json:"LinksList,omitempty" name:"LinksList"`
}

// Predefined struct for user
type CheckAlarmRegularNameExistRequestParams struct {
	// 项目名称
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 规则名称
	AlarmRegularName *string `json:"AlarmRegularName,omitempty" name:"AlarmRegularName"`

	// 主键ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

type CheckAlarmRegularNameExistRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 规则名称
	AlarmRegularName *string `json:"AlarmRegularName,omitempty" name:"AlarmRegularName"`

	// 主键ID
	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *CheckAlarmRegularNameExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAlarmRegularNameExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "AlarmRegularName")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAlarmRegularNameExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAlarmRegularNameExistResponseParams struct {
	// 是否重名
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckAlarmRegularNameExistResponse struct {
	*tchttp.BaseResponse
	Response *CheckAlarmRegularNameExistResponseParams `json:"Response"`
}

func (r *CheckAlarmRegularNameExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAlarmRegularNameExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIntegrationNodeNameExistsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 节点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 节点ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type CheckIntegrationNodeNameExistsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 节点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 节点ID
	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *CheckIntegrationNodeNameExistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIntegrationNodeNameExistsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Name")
	delete(f, "ProjectId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIntegrationNodeNameExistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIntegrationNodeNameExistsResponseParams struct {
	// 返回true代表存在，返回false代表不存在
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckIntegrationNodeNameExistsResponse struct {
	*tchttp.BaseResponse
	Response *CheckIntegrationNodeNameExistsResponseParams `json:"Response"`
}

func (r *CheckIntegrationNodeNameExistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIntegrationNodeNameExistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIntegrationTaskNameExistsRequestParams struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 同步类型1.单表同步，2.解决方案
	SyncType *int64 `json:"SyncType,omitempty" name:"SyncType"`
}

type CheckIntegrationTaskNameExistsRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 同步类型1.单表同步，2.解决方案
	SyncType *int64 `json:"SyncType,omitempty" name:"SyncType"`
}

func (r *CheckIntegrationTaskNameExistsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIntegrationTaskNameExistsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "SyncType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckIntegrationTaskNameExistsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckIntegrationTaskNameExistsResponseParams struct {
	// true表示存在，false表示不存在
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 任务名重复类型（0:未重复, 1:开发态重复, 2:生产态重复）
	ExistsType *int64 `json:"ExistsType,omitempty" name:"ExistsType"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckIntegrationTaskNameExistsResponse struct {
	*tchttp.BaseResponse
	Response *CheckIntegrationTaskNameExistsResponseParams `json:"Response"`
}

func (r *CheckIntegrationTaskNameExistsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckIntegrationTaskNameExistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTaskNameExistRequestParams struct {
	// 项目id/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型（跟调度传参保持一致27）
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type CheckTaskNameExistRequest struct {
	*tchttp.BaseRequest
	
	// 项目id/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型（跟调度传参保持一致27）
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

func (r *CheckTaskNameExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTaskNameExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TypeId")
	delete(f, "TaskName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckTaskNameExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckTaskNameExistResponseParams struct {
	// 结果
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CheckTaskNameExistResponse struct {
	*tchttp.BaseResponse
	Response *CheckTaskNameExistResponseParams `json:"Response"`
}

func (r *CheckTaskNameExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckTaskNameExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 0.仅提交，1.立即启动，2.停止线上作业，丢弃作业状态数据，重新启动运行，3.暂停线上作业，保留作业状态数据，继续运行，4.保留作业状态数据，继续运行
	CommitType *int64 `json:"CommitType,omitempty" name:"CommitType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type CommitIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 0.仅提交，1.立即启动，2.停止线上作业，丢弃作业状态数据，重新启动运行，3.暂停线上作业，保留作业状态数据，继续运行，4.保留作业状态数据，继续运行
	CommitType *int64 `json:"CommitType,omitempty" name:"CommitType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *CommitIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "CommitType")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CommitIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CommitIntegrationTaskResponseParams `json:"Response"`
}

func (r *CommitIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonContent struct {
	// 详情内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type CommonId struct {
	// Id值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`
}

// Predefined struct for user
type CreateCustomFunctionRequestParams struct {
	// 类型：HIVE、SPARK
	Type *string `json:"Type,omitempty" name:"Type"`

	// 分类：窗口函数、聚合函数、日期函数......
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 函数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集群实例引擎 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type CreateCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 类型：HIVE、SPARK
	Type *string `json:"Type,omitempty" name:"Type"`

	// 分类：窗口函数、聚合函数、日期函数......
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 函数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集群实例引擎 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 数据库名称
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateCustomFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Kind")
	delete(f, "Name")
	delete(f, "ClusterIdentifier")
	delete(f, "DbName")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomFunctionResponseParams struct {
	// 函数唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCustomFunctionResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomFunctionResponseParams `json:"Response"`
}

func (r *CreateCustomFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataSourceRequestParams struct {
	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitempty" name:"Type"`

	// 归属项目ID
	OwnerProjectId *string `json:"OwnerProjectId,omitempty" name:"OwnerProjectId"`

	// 归属项目Name
	OwnerProjectName *string `json:"OwnerProjectName,omitempty" name:"OwnerProjectName"`

	// 归属项目Name中文
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitempty" name:"OwnerProjectIdent"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitempty" name:"Collect"`

	// cos桶信息
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitempty" name:"COSRegion"`
}

type CreateDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitempty" name:"Type"`

	// 归属项目ID
	OwnerProjectId *string `json:"OwnerProjectId,omitempty" name:"OwnerProjectId"`

	// 归属项目Name
	OwnerProjectName *string `json:"OwnerProjectName,omitempty" name:"OwnerProjectName"`

	// 归属项目Name中文
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitempty" name:"OwnerProjectIdent"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitempty" name:"Collect"`

	// cos桶信息
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitempty" name:"COSRegion"`
}

func (r *CreateDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Category")
	delete(f, "Type")
	delete(f, "OwnerProjectId")
	delete(f, "OwnerProjectName")
	delete(f, "OwnerProjectIdent")
	delete(f, "BizParams")
	delete(f, "Params")
	delete(f, "Description")
	delete(f, "Display")
	delete(f, "DatabaseName")
	delete(f, "Instance")
	delete(f, "Status")
	delete(f, "ClusterId")
	delete(f, "Collect")
	delete(f, "COSBucket")
	delete(f, "COSRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataSourceResponseParams struct {
	// 主键ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *uint64 `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateDataSourceResponseParams `json:"Response"`
}

func (r *CreateDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

type CreateFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

func (r *CreateFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "ParentsFolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateFolderResponseParams struct {
	// 文件夹Id，null则创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonId `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateFolderResponseParams `json:"Response"`
}

func (r *CreateFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHiveTableByDDLRequestParams struct {
	// 数据源ID
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// 建hive表ddl
	DDLSql *string `json:"DDLSql,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitempty" name:"Privilege"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 目标表类型(HIVE或GBASE)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 责任人
	Incharge *string `json:"Incharge,omitempty" name:"Incharge"`
}

type CreateHiveTableByDDLRequest struct {
	*tchttp.BaseRequest
	
	// 数据源ID
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// 建hive表ddl
	DDLSql *string `json:"DDLSql,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitempty" name:"Privilege"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 目标表类型(HIVE或GBASE)
	Type *string `json:"Type,omitempty" name:"Type"`

	// 责任人
	Incharge *string `json:"Incharge,omitempty" name:"Incharge"`
}

func (r *CreateHiveTableByDDLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHiveTableByDDLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceId")
	delete(f, "Database")
	delete(f, "DDLSql")
	delete(f, "Privilege")
	delete(f, "ProjectId")
	delete(f, "Type")
	delete(f, "Incharge")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHiveTableByDDLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHiveTableByDDLResponseParams struct {
	// 表名称
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHiveTableByDDLResponse struct {
	*tchttp.BaseResponse
	Response *CreateHiveTableByDDLResponseParams `json:"Response"`
}

func (r *CreateHiveTableByDDLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHiveTableByDDLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHiveTableRequestParams struct {
	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// 建hive表ddl
	DDLSql *string `json:"DDLSql,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitempty" name:"Privilege"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 责任人
	Incharge *string `json:"Incharge,omitempty" name:"Incharge"`
}

type CreateHiveTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 数据库
	Database *string `json:"Database,omitempty" name:"Database"`

	// 建hive表ddl
	DDLSql *string `json:"DDLSql,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitempty" name:"Privilege"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 责任人
	Incharge *string `json:"Incharge,omitempty" name:"Incharge"`
}

func (r *CreateHiveTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHiveTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceId")
	delete(f, "Database")
	delete(f, "DDLSql")
	delete(f, "Privilege")
	delete(f, "ProjectId")
	delete(f, "Incharge")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHiveTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHiveTableResponseParams struct {
	// 建表是否成功
	IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateHiveTableResponse struct {
	*tchttp.BaseResponse
	Response *CreateHiveTableResponseParams `json:"Response"`
}

func (r *CreateHiveTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateHiveTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInLongAgentRequestParams struct {
	// 采集器类型，1：TKE Agent，2：BOSS SDK，默认：1
	AgentType *uint64 `json:"AgentType,omitempty" name:"AgentType"`

	// 采集器名称
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 集成资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// TKE集群的地域
	TkeRegion *string `json:"TkeRegion,omitempty" name:"TkeRegion"`

	// 当AgentType为1时，必填。当AgentType为2时，不用填
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type CreateInLongAgentRequest struct {
	*tchttp.BaseRequest
	
	// 采集器类型，1：TKE Agent，2：BOSS SDK，默认：1
	AgentType *uint64 `json:"AgentType,omitempty" name:"AgentType"`

	// 采集器名称
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 集成资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// TKE集群的地域
	TkeRegion *string `json:"TkeRegion,omitempty" name:"TkeRegion"`

	// 当AgentType为1时，必填。当AgentType为2时，不用填
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateInLongAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInLongAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentType")
	delete(f, "AgentName")
	delete(f, "ExecutorGroupId")
	delete(f, "ProjectId")
	delete(f, "TkeRegion")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateInLongAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateInLongAgentResponseParams struct {
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateInLongAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateInLongAgentResponseParams `json:"Response"`
}

func (r *CreateInLongAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateInLongAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationNodeRequestParams struct {
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type CreateIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *CreateIntegrationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeInfo")
	delete(f, "ProjectId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationNodeResponseParams struct {
	// 节点
	Id *string `json:"Id,omitempty" name:"Id"`

	// 当前任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIntegrationNodeResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationNodeResponseParams `json:"Response"`
}

func (r *CreateIntegrationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationTaskRequestParams struct {
	// 任务信息
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type CreateIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务信息
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskInfo")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateIntegrationTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateIntegrationTaskResponseParams `json:"Response"`
}

func (r *CreateIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOfflineTaskRequestParams struct {
	// 项目/工作
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 0
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 2099-12-31 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 当前日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 跟之前调用调度接口保持一致27
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`

	// 默认 ""
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// 区分画布和表单
	TaskMode *string `json:"TaskMode,omitempty" name:"TaskMode"`
}

type CreateOfflineTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目/工作
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 0
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 2099-12-31 00:00:00
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 当前日期
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 跟之前调用调度接口保持一致27
	TypeId *int64 `json:"TypeId,omitempty" name:"TypeId"`

	// 默认 ""
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// 区分画布和表单
	TaskMode *string `json:"TaskMode,omitempty" name:"TaskMode"`
}

func (r *CreateOfflineTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOfflineTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CycleStep")
	delete(f, "DelayTime")
	delete(f, "EndTime")
	delete(f, "Notes")
	delete(f, "StartTime")
	delete(f, "TaskName")
	delete(f, "TypeId")
	delete(f, "TaskAction")
	delete(f, "TaskMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOfflineTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOfflineTaskResponseParams struct {
	// 结果
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOfflineTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateOfflineTaskResponseParams `json:"Response"`
}

func (r *CreateOfflineTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOfflineTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrUpdateResourceRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件名
	Files []*string `json:"Files,omitempty" name:"Files"`

	// 文件所属路径，资源管理根路径为 /datastudio/resouce
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// cos存储桶名字
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// cos所属地域
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 是否为新文件，新增为 true，更新为 false
	NewFile *bool `json:"NewFile,omitempty" name:"NewFile"`

	// 文件大小
	FilesSize []*string `json:"FilesSize,omitempty" name:"FilesSize"`
}

type CreateOrUpdateResourceRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件名
	Files []*string `json:"Files,omitempty" name:"Files"`

	// 文件所属路径，资源管理根路径为 /datastudio/resouce
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// cos存储桶名字
	CosBucketName *string `json:"CosBucketName,omitempty" name:"CosBucketName"`

	// cos所属地域
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 是否为新文件，新增为 true，更新为 false
	NewFile *bool `json:"NewFile,omitempty" name:"NewFile"`

	// 文件大小
	FilesSize []*string `json:"FilesSize,omitempty" name:"FilesSize"`
}

func (r *CreateOrUpdateResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrUpdateResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Files")
	delete(f, "FilePath")
	delete(f, "CosBucketName")
	delete(f, "CosRegion")
	delete(f, "NewFile")
	delete(f, "FilesSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrUpdateResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrUpdateResourceResponseParams struct {
	// 响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*UserFileDTO `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateOrUpdateResourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrUpdateResourceResponseParams `json:"Response"`
}

func (r *CreateOrUpdateResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrUpdateResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskAlarmRegularRequestParams struct {
	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type CreateTaskAlarmRegularRequest struct {
	*tchttp.BaseRequest
	
	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateTaskAlarmRegularRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskAlarmRegularRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskAlarmInfo")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskAlarmRegularRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskAlarmRegularResponseParams struct {
	// 告警ID
	AlarmId *int64 `json:"AlarmId,omitempty" name:"AlarmId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateTaskAlarmRegularResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskAlarmRegularResponseParams `json:"Response"`
}

func (r *CreateTaskAlarmRegularResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskAlarmRegularResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 扩展属性
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 扩展属性
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`
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
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "TaskName")
	delete(f, "TaskType")
	delete(f, "TaskExt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskResponseParams struct {
	// 返回任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonId `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
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
type CreateWorkflowRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

type CreateWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *CreateWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowName")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowResponseParams struct {
	// 返回工作流Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonId `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowResponseParams `json:"Response"`
}

func (r *CreateWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataSourceInfo struct {
	// 若数据源列表为绑定数据库，则为db名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据源描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 数据源引擎的实例ID，如CDB实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源引擎所属区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 数据源类型:枚举值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据源所属的集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 应用ID AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// 业务侧数据源的配置信息扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizParams *string `json:"BizParams,omitempty" name:"BizParams"`

	// 数据源类别：绑定引擎、绑定数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 数据源展示名，为了可视化查看
	// 注意：此字段可能返回 null，表示取不到有效值。
	Display *string `json:"Display,omitempty" name:"Display"`

	// 数据源责任人账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerAccount *string `json:"OwnerAccount,omitempty" name:"OwnerAccount"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitempty" name:"Params"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 数据源责任人账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerAccountName *string `json:"OwnerAccountName,omitempty" name:"OwnerAccountName"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 归属项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerProjectId *string `json:"OwnerProjectId,omitempty" name:"OwnerProjectId"`

	// 归属项目Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerProjectName *string `json:"OwnerProjectName,omitempty" name:"OwnerProjectName"`

	// 归属项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitempty" name:"OwnerProjectIdent"`

	// 授权项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorityProjectName *string `json:"AuthorityProjectName,omitempty" name:"AuthorityProjectName"`

	// 授权用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorityUserName *string `json:"AuthorityUserName,omitempty" name:"AuthorityUserName"`

	// 是否有编辑权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edit *bool `json:"Edit,omitempty" name:"Edit"`

	// 是否有授权权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Author *bool `json:"Author,omitempty" name:"Author"`

	// 是否有转交权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deliver *bool `json:"Deliver,omitempty" name:"Deliver"`

	// 数据源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceStatus *string `json:"DataSourceStatus,omitempty" name:"DataSourceStatus"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Params json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamsString *string `json:"ParamsString,omitempty" name:"ParamsString"`

	// BizParams json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizParamsString *string `json:"BizParamsString,omitempty" name:"BizParamsString"`
}

type DataSourceInfoPage struct {
	// 分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 数据源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*DataSourceInfo `json:"Rows,omitempty" name:"Rows"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 总分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitempty" name:"TotalPageNumber"`
}

type DatasourceBaseInfo struct {
	// 若数据源列表为绑定数据库，则为db名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseNames []*string `json:"DatabaseNames,omitempty" name:"DatabaseNames"`

	// 数据源描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 数据源引擎的实例ID，如CDB实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源引擎所属区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 数据源类型:枚举值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据源所属的集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

// Predefined struct for user
type DeleteCustomFunctionRequestParams struct {
	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 函数 ID
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 项目ID，必须填
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DeleteCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 函数 ID
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 项目ID，必须填
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteCustomFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIdentifier")
	delete(f, "FunctionId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomFunctionResponseParams struct {
	// 函数 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCustomFunctionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomFunctionResponseParams `json:"Response"`
}

func (r *DeleteCustomFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataSourcesRequestParams struct {
	// id列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

type DeleteDataSourcesRequest struct {
	*tchttp.BaseRequest
	
	// id列表
	Ids []*uint64 `json:"Ids,omitempty" name:"Ids"`
}

func (r *DeleteDataSourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataSourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataSourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataSourcesResponseParams struct {
	// 是否删除成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteDataSourcesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDataSourcesResponseParams `json:"Response"`
}

func (r *DeleteDataSourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

type DeleteFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *DeleteFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFolderResponseParams struct {
	// true代表删除成功，false代表删除失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFolderResponseParams `json:"Response"`
}

func (r *DeleteFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInLongAgentRequestParams struct {
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DeleteInLongAgentRequest struct {
	*tchttp.BaseRequest
	
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteInLongAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInLongAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteInLongAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteInLongAgentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteInLongAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteInLongAgentResponseParams `json:"Response"`
}

func (r *DeleteInLongAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteInLongAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationNodeRequestParams struct {
	// 节点id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DeleteIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 节点id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteIntegrationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationNodeResponseParams struct {
	// 删除返回是否成功标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIntegrationNodeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationNodeResponseParams `json:"Response"`
}

func (r *DeleteIntegrationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DeleteIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationTaskResponseParams struct {
	// 任务删除成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteIntegrationTaskResponseParams `json:"Response"`
}

func (r *DeleteIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOfflineTaskRequestParams struct {
	// 操作者name
	OperatorName *string `json:"OperatorName,omitempty" name:"OperatorName"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 虚拟任务标记(跟之前调度接口保持一致默认false)
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`
}

type DeleteOfflineTaskRequest struct {
	*tchttp.BaseRequest
	
	// 操作者name
	OperatorName *string `json:"OperatorName,omitempty" name:"OperatorName"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 虚拟任务标记(跟之前调度接口保持一致默认false)
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`
}

func (r *DeleteOfflineTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOfflineTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OperatorName")
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "VirtualFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteOfflineTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteOfflineTaskResponseParams struct {
	// 结果
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteOfflineTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteOfflineTaskResponseParams `json:"Response"`
}

func (r *DeleteOfflineTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteOfflineTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type DeleteResourceRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DeleteResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteResourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceResponseParams `json:"Response"`
}

func (r *DeleteResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskAlarmRegularRequestParams struct {
	// 主键ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型(201表示实时任务，202表示离线任务)
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
}

type DeleteTaskAlarmRegularRequest struct {
	*tchttp.BaseRequest
	
	// 主键ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型(201表示实时任务，202表示离线任务)
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DeleteTaskAlarmRegularRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskAlarmRegularRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskAlarmRegularRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskAlarmRegularResponseParams struct {
	// 删除结果(true表示删除成功，false表示删除失败)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteTaskAlarmRegularResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskAlarmRegularResponseParams `json:"Response"`
}

func (r *DeleteTaskAlarmRegularResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskAlarmRegularResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowNewRequestParams struct {
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DeleteWorkflowNewRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DeleteWorkflowNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkFlowId")
	delete(f, "DeleteMode")
	delete(f, "EnableNotify")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowNewResponseParams struct {
	// 返回删除结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteWorkflowNewResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowNewResponseParams `json:"Response"`
}

func (r *DeleteWorkflowNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependencyConfig struct {
	// 仅五种周期运行依赖配置： HOUR,DAY,WEEK,MONTH,YEAR,CRONTAB,MINUTE
	DependConfType *string `json:"DependConfType,omitempty" name:"DependConfType"`

	// 依赖配置从属周期类型，CURRENT_HOUR，PREVIOUS_HOUR，CURRENT_DAY，PREVIOUS_DAY，PREVIOUS_WEEK，PREVIOUS_FRIDAY，PREVIOUS_WEEKEND，CURRENT_MONTH，PREVIOUS_MONTH，PREVIOUS_END_OF_MONTH
	//      * PREVIOUS_BEGIN_OF_MONTH，ALL_MONTH_OF_YEAR，ALL_DAY_OF_YEAR，CURRENT_YEAR，CURRENT，CURRENT_MINUTE，PREVIOUS_MINUTE_CYCLE，PREVIOUS_HOUR_CYCLE
	SubordinateCyclicType *string `json:"SubordinateCyclicType,omitempty" name:"SubordinateCyclicType"`

	// WAITING，等待（默认策略）EXECUTING:执行
	DependencyStrategy *string `json:"DependencyStrategy,omitempty" name:"DependencyStrategy"`

	// 父任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentTask *TaskInnerInfo `json:"ParentTask,omitempty" name:"ParentTask"`

	// 子任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SonTask *TaskInnerInfo `json:"SonTask,omitempty" name:"SonTask"`
}

// Predefined struct for user
type DescribeAlarmEventsRequestParams struct {
	// 过滤条件(key可以是：AlarmLevel,AlarmIndicator,KeyWord)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段（AlarmTime）
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 类型(201表示实时，202表示离线)
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeAlarmEventsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件(key可以是：AlarmLevel,AlarmIndicator,KeyWord)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段（AlarmTime）
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 类型(201表示实时，202表示离线)
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeAlarmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmEventsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "TaskType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmEventsResponseParams struct {
	// 告警事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmEventInfoList []*AlarmEventInfo `json:"AlarmEventInfoList,omitempty" name:"AlarmEventInfoList"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmEventsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmEventsResponseParams `json:"Response"`
}

func (r *DescribeAlarmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmReceiverRequestParams struct {
	// 告警ID
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 消息ID
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// 类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 告警接收人ID(逗号分隔)
	AlarmRecipient *string `json:"AlarmRecipient,omitempty" name:"AlarmRecipient"`

	// 告警接收人姓名(逗号分隔)
	AlarmRecipientName *string `json:"AlarmRecipientName,omitempty" name:"AlarmRecipientName"`

	// 告警时间
	AlarmTime *string `json:"AlarmTime,omitempty" name:"AlarmTime"`
}

type DescribeAlarmReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 告警ID
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 消息ID
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// 类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 告警接收人ID(逗号分隔)
	AlarmRecipient *string `json:"AlarmRecipient,omitempty" name:"AlarmRecipient"`

	// 告警接收人姓名(逗号分隔)
	AlarmRecipientName *string `json:"AlarmRecipientName,omitempty" name:"AlarmRecipientName"`

	// 告警时间
	AlarmTime *string `json:"AlarmTime,omitempty" name:"AlarmTime"`
}

func (r *DescribeAlarmReceiverRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmReceiverRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "MessageId")
	delete(f, "TaskType")
	delete(f, "AlarmRecipient")
	delete(f, "AlarmRecipientName")
	delete(f, "AlarmTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmReceiverResponseParams struct {
	// 告警接收人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmReceiverInfoList []*AlarmReceiverInfo `json:"AlarmReceiverInfoList,omitempty" name:"AlarmReceiverInfoList"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAlarmReceiverResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmReceiverResponseParams `json:"Response"`
}

func (r *DescribeAlarmReceiverResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmReceiverResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNamespaceListRequestParams struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeClusterNamespaceListRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeClusterNamespaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNamespaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterNamespaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterNamespaceListResponseParams struct {
	// 命名空间
	Namespaces []*Namespace `json:"Namespaces,omitempty" name:"Namespaces"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterNamespaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterNamespaceListResponseParams `json:"Response"`
}

func (r *DescribeClusterNamespaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterNamespaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceInfoListRequestParams struct {
	// 工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 可选过滤条件，Filter可选配置(参考): "Name": { "type": "string", "description": "数据源名称" }, "Type": { "type": "string", "description": "类型" }, "ClusterId": { "type": "string", "description": "集群id" }, "CategoryId": { "type": "string", "description": "分类，项目或空间id" }
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序配置
	OrderFields *OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 数据源类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据源名称过滤用
	DatasourceName *string `json:"DatasourceName,omitempty" name:"DatasourceName"`
}

type DescribeDataSourceInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 可选过滤条件，Filter可选配置(参考): "Name": { "type": "string", "description": "数据源名称" }, "Type": { "type": "string", "description": "类型" }, "ClusterId": { "type": "string", "description": "集群id" }, "CategoryId": { "type": "string", "description": "分类，项目或空间id" }
	Filters *Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序配置
	OrderFields *OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 数据源类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据源名称过滤用
	DatasourceName *string `json:"DatasourceName,omitempty" name:"DatasourceName"`
}

func (r *DescribeDataSourceInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "Type")
	delete(f, "DatasourceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSourceInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceInfoListResponseParams struct {
	// 总条数。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 数据源信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceSet []*DatasourceBaseInfo `json:"DatasourceSet,omitempty" name:"DatasourceSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataSourceInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSourceInfoListResponseParams `json:"Response"`
}

func (r *DescribeDataSourceInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceListRequestParams struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 返回数量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 排序配置
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 可选过滤条件，Filter可选配置(参考): "Name": { "type": "string", "description": "数据源名称" }, "Type": { "type": "string", "description": "类型" }, "ClusterId": { "type": "string", "description": "集群id" }, "CategoryId": { "type": "string", "description": "分类，项目或空间id" }
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDataSourceListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 返回数量
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 排序配置
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 可选过滤条件，Filter可选配置(参考): "Name": { "type": "string", "description": "数据源名称" }, "Type": { "type": "string", "description": "类型" }, "ClusterId": { "type": "string", "description": "集群id" }, "CategoryId": { "type": "string", "description": "分类，项目或空间id" }
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDataSourceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "OrderFields")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSourceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceListResponseParams struct {
	// 数据源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataSourceInfoPage `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataSourceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSourceListResponseParams `json:"Response"`
}

func (r *DescribeDataSourceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceWithoutInfoRequestParams struct {
	// 1
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

type DescribeDataSourceWithoutInfoRequest struct {
	*tchttp.BaseRequest
	
	// 1
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 1
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDataSourceWithoutInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceWithoutInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderFields")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataSourceWithoutInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceWithoutInfoResponseParams struct {
	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DataSourceInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataSourceWithoutInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataSourceWithoutInfoResponseParams `json:"Response"`
}

func (r *DescribeDataSourceWithoutInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataSourceWithoutInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataTypesRequestParams struct {
	// 数据源类型，MYSQL|KAFKA等
	DatasourceType *string `json:"DatasourceType,omitempty" name:"DatasourceType"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeDataTypesRequest struct {
	*tchttp.BaseRequest
	
	// 数据源类型，MYSQL|KAFKA等
	DatasourceType *string `json:"DatasourceType,omitempty" name:"DatasourceType"`

	// 项目ID。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeDataTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceType")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataTypesResponseParams struct {
	// 字段类型列表。
	TypeInfoSet []*Label `json:"TypeInfoSet,omitempty" name:"TypeInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDataTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataTypesResponseParams `json:"Response"`
}

func (r *DescribeDataTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseInfoListRequestParams struct {
	// 如果是hive这里写rpc，如果是其他类型不传
	ConnectionType *string `json:"ConnectionType,omitempty" name:"ConnectionType"`
}

type DescribeDatabaseInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 如果是hive这里写rpc，如果是其他类型不传
	ConnectionType *string `json:"ConnectionType,omitempty" name:"ConnectionType"`
}

func (r *DescribeDatabaseInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConnectionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseInfoListResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDatabaseInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseInfoListResponseParams `json:"Response"`
}

func (r *DescribeDatabaseInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceRequestParams struct {
	// 对象唯一ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type DescribeDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// 对象唯一ID
	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeDatasourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceResponseParams struct {
	// 数据源对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataSourceInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDatasourceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatasourceResponseParams `json:"Response"`
}

func (r *DescribeDatasourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatasourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDependTasksNewRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 上游/下游层级1-6级
	Deep *uint64 `json:"Deep,omitempty" name:"Deep"`

	// 1: 表示查询上游节点；0:表示查询下游节点；2：表示查询上游和下游节点
	Up *uint64 `json:"Up,omitempty" name:"Up"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

type DescribeDependTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 上游/下游层级1-6级
	Deep *uint64 `json:"Deep,omitempty" name:"Deep"`

	// 1: 表示查询上游节点；0:表示查询下游节点；2：表示查询上游和下游节点
	Up *uint64 `json:"Up,omitempty" name:"Up"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`
}

func (r *DescribeDependTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "Deep")
	delete(f, "Up")
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDependTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDependTasksNewResponseParams struct {
	// 画布任务和链接信息
	Data *CanvasInfo `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeDependTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDependTasksNewResponseParams `json:"Response"`
}

func (r *DescribeDependTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFolderListData struct {
	// 文件夹信息列表
	Items []*Folder `json:"Items,omitempty" name:"Items"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeFolderListRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeFolderListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeFolderListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentsFolderId")
	delete(f, "KeyWords")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFolderListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFolderListResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeFolderListData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFolderListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFolderListResponseParams `json:"Response"`
}

func (r *DescribeFolderListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFolderWorkflowListData struct {
	// 工作流信息列表
	Items []*Workflow `json:"Items,omitempty" name:"Items"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeFolderWorkflowListRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeFolderWorkflowListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeFolderWorkflowListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderWorkflowListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ParentsFolderId")
	delete(f, "KeyWords")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFolderWorkflowListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFolderWorkflowListResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeFolderWorkflowListData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFolderWorkflowListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFolderWorkflowListResponseParams `json:"Response"`
}

func (r *DescribeFolderWorkflowListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFolderWorkflowListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionKindsRequestParams struct {

}

type DescribeFunctionKindsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFunctionKindsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionKindsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFunctionKindsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionKindsResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kinds []*FunctionTypeOrKind `json:"Kinds,omitempty" name:"Kinds"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFunctionKindsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFunctionKindsResponseParams `json:"Response"`
}

func (r *DescribeFunctionKindsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionKindsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionTypesRequestParams struct {

}

type DescribeFunctionTypesRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeFunctionTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionTypesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFunctionTypesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFunctionTypesResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Types []*FunctionTypeOrKind `json:"Types,omitempty" name:"Types"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeFunctionTypesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFunctionTypesResponseParams `json:"Response"`
}

func (r *DescribeFunctionTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFunctionTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongAgentListRequestParams struct {
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// Agent Name
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 集群类型，1：TKE Agent，2：BOSS SDK，默认：1
	AgentType *uint64 `json:"AgentType,omitempty" name:"AgentType"`

	// Agent状态(running运行中，initializing 操作中，failed心跳异常)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 分页页码，从1开始，默认：1
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 分页每页记录数，默认10
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 名称搜索是否开启模糊匹配，1：开启，0：不开启（精确匹配）
	Like *uint64 `json:"Like,omitempty" name:"Like"`
}

type DescribeInLongAgentListRequest struct {
	*tchttp.BaseRequest
	
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// Agent Name
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 集群类型，1：TKE Agent，2：BOSS SDK，默认：1
	AgentType *uint64 `json:"AgentType,omitempty" name:"AgentType"`

	// Agent状态(running运行中，initializing 操作中，failed心跳异常)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Vpc Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 分页页码，从1开始，默认：1
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 分页每页记录数，默认10
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 名称搜索是否开启模糊匹配，1：开启，0：不开启（精确匹配）
	Like *uint64 `json:"Like,omitempty" name:"Like"`
}

func (r *DescribeInLongAgentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongAgentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "AgentId")
	delete(f, "AgentName")
	delete(f, "AgentType")
	delete(f, "Status")
	delete(f, "VpcId")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Like")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInLongAgentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongAgentListResponseParams struct {
	// 采集器信息列表
	Items []*InLongAgentDetail `json:"Items,omitempty" name:"Items"`

	// 页码
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 总页数
	TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInLongAgentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInLongAgentListResponseParams `json:"Response"`
}

func (r *DescribeInLongAgentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongAgentTaskListRequestParams struct {
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeInLongAgentTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeInLongAgentTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongAgentTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInLongAgentTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongAgentTaskListResponseParams struct {
	// 采集器关联的集成任务列表
	Items []*InLongAgentTask `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInLongAgentTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInLongAgentTaskListResponseParams `json:"Response"`
}

func (r *DescribeInLongAgentTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongAgentTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongAgentVpcListRequestParams struct {
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeInLongAgentVpcListRequest struct {
	*tchttp.BaseRequest
	
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeInLongAgentVpcListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongAgentVpcListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInLongAgentVpcListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongAgentVpcListResponseParams struct {
	// VPC列表
	VpcList []*string `json:"VpcList,omitempty" name:"VpcList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInLongAgentVpcListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInLongAgentVpcListResponseParams `json:"Response"`
}

func (r *DescribeInLongAgentVpcListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongAgentVpcListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongTkeClusterListRequestParams struct {
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// TKE集群地域
	TkeRegion *string `json:"TkeRegion,omitempty" name:"TkeRegion"`

	// 集群名称。
	// 多个名称用逗号连接。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// TKE集群状态 (Running 运行中 Creating 创建中 Initializing 创建中 Idling 闲置中 Abnormal 异常 Failed 异常 Terminating 删除中 Deleting 删除中 Scaling 规模调整中 Upgrading 升级中 Isolated 欠费隔离中 NodeUpgrading 节点升级中 Recovering 唤醒中 Activating 激活中 MasterScaling Master扩缩容中 Waiting 等待注册 ClusterLevelUpgrading 调整规格中 ResourceIsolate 隔离中 ResourceIsolated 已隔离 ResourceReverse 冲正中 Trading 集群开通中 ResourceReversal 集群冲正 ClusterLevelTrading 集群变配交易中)
	// 多个状态用逗号连接。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否安装Agent，true: 是，false: 否
	HasAgent *bool `json:"HasAgent,omitempty" name:"HasAgent"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	// 多个集群用逗号连接。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 分页页码，从1开始，默认：1
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 分页每页记录数，默认10
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeInLongTkeClusterListRequest struct {
	*tchttp.BaseRequest
	
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// TKE集群地域
	TkeRegion *string `json:"TkeRegion,omitempty" name:"TkeRegion"`

	// 集群名称。
	// 多个名称用逗号连接。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// TKE集群状态 (Running 运行中 Creating 创建中 Initializing 创建中 Idling 闲置中 Abnormal 异常 Failed 异常 Terminating 删除中 Deleting 删除中 Scaling 规模调整中 Upgrading 升级中 Isolated 欠费隔离中 NodeUpgrading 节点升级中 Recovering 唤醒中 Activating 激活中 MasterScaling Master扩缩容中 Waiting 等待注册 ClusterLevelUpgrading 调整规格中 ResourceIsolate 隔离中 ResourceIsolated 已隔离 ResourceReverse 冲正中 Trading 集群开通中 ResourceReversal 集群冲正 ClusterLevelTrading 集群变配交易中)
	// 多个状态用逗号连接。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否安装Agent，true: 是，false: 否
	HasAgent *bool `json:"HasAgent,omitempty" name:"HasAgent"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。
	// 多个集群用逗号连接。
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`

	// 分页页码，从1开始，默认：1
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 分页每页记录数，默认10
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeInLongTkeClusterListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongTkeClusterListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TkeRegion")
	delete(f, "ClusterName")
	delete(f, "Status")
	delete(f, "HasAgent")
	delete(f, "ClusterType")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInLongTkeClusterListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInLongTkeClusterListResponseParams struct {
	// TKE集群信息
	Items []*InLongTkeDetail `json:"Items,omitempty" name:"Items"`

	// 页码
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 总页数
	TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInLongTkeClusterListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInLongTkeClusterListResponseParams `json:"Response"`
}

func (r *DescribeInLongTkeClusterListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInLongTkeClusterListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLastLogRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

type DescribeInstanceLastLogRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

func (r *DescribeInstanceLastLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLastLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLastLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLastLogResponseParams struct {
	// 日志
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceLastLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLastLogResponseParams `json:"Response"`
}

func (r *DescribeInstanceLastLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLastLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListRequestParams struct {
	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 周期列表（如天，一次性），可选
	CycleList []*string `json:"CycleList,omitempty" name:"CycleList"`

	// 责任人
	OwnerList []*string `json:"OwnerList,omitempty" name:"OwnerList"`

	// 跟之前保持一致
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitempty" name:"SortCol"`

	// 类型列表（如35 shell任务），可选
	TaskTypeList []*int64 `json:"TaskTypeList,omitempty" name:"TaskTypeList"`

	// 状态列表（如成功 2，正在执行 1），可选
	StateList []*int64 `json:"StateList,omitempty" name:"StateList"`

	// 任务名称
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 周期列表（如天，一次性），可选
	CycleList []*string `json:"CycleList,omitempty" name:"CycleList"`

	// 责任人
	OwnerList []*string `json:"OwnerList,omitempty" name:"OwnerList"`

	// 跟之前保持一致
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitempty" name:"SortCol"`

	// 类型列表（如35 shell任务），可选
	TaskTypeList []*int64 `json:"TaskTypeList,omitempty" name:"TaskTypeList"`

	// 状态列表（如成功 2，正在执行 1），可选
	StateList []*int64 `json:"StateList,omitempty" name:"StateList"`

	// 任务名称
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
}

func (r *DescribeInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "CycleList")
	delete(f, "OwnerList")
	delete(f, "InstanceType")
	delete(f, "Sort")
	delete(f, "SortCol")
	delete(f, "TaskTypeList")
	delete(f, "StateList")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceListResponseParams struct {
	// 结果
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceListResponseParams `json:"Response"`
}

func (r *DescribeInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogListRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

type DescribeInstanceLogListRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

func (r *DescribeInstanceLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogListResponseParams struct {
	// 日志列表
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogListResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 服务器Ip
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 文件Name
	OriginFileName *string `json:"OriginFileName,omitempty" name:"OriginFileName"`
}

type DescribeInstanceLogRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 服务器Ip
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 文件Name
	OriginFileName *string `json:"OriginFileName,omitempty" name:"OriginFileName"`
}

func (r *DescribeInstanceLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	delete(f, "BrokerIp")
	delete(f, "OriginFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogResponseParams struct {
	// 返回结果
	Data *string `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

type DescribeInstanceLogsRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

func (r *DescribeInstanceLogsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogsResponseParams struct {
	// 返回日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*InstanceLog `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstanceLogsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogsResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationNodeRequestParams struct {
	// 节点id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type DescribeIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 节点id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DescribeIntegrationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationNodeResponseParams struct {
	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`

	// 上游节点是否已经修改。true 已修改，需要提示；false 没有修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCheckFlag *bool `json:"SourceCheckFlag,omitempty" name:"SourceCheckFlag"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationNodeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationNodeResponseParams `json:"Response"`
}

func (r *DescribeIntegrationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsAgentStatusRequestParams struct {
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

type DescribeIntegrationStatisticsAgentStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

func (r *DescribeIntegrationStatisticsAgentStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsAgentStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "ProjectId")
	delete(f, "QueryDate")
	delete(f, "ExecutorGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationStatisticsAgentStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsAgentStatusResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusData *string `json:"StatusData,omitempty" name:"StatusData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationStatisticsAgentStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationStatisticsAgentStatusResponseParams `json:"Response"`
}

func (r *DescribeIntegrationStatisticsAgentStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsInstanceTrendRequestParams struct {
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

type DescribeIntegrationStatisticsInstanceTrendRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

func (r *DescribeIntegrationStatisticsInstanceTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsInstanceTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "ProjectId")
	delete(f, "QueryDate")
	delete(f, "ExecutorGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationStatisticsInstanceTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsInstanceTrendResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrendsData []*IntegrationStatisticsTrendResult `json:"TrendsData,omitempty" name:"TrendsData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationStatisticsInstanceTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationStatisticsInstanceTrendResponseParams `json:"Response"`
}

func (r *DescribeIntegrationStatisticsInstanceTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsInstanceTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsRecordsTrendRequestParams struct {
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`
}

type DescribeIntegrationStatisticsRecordsTrendRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`
}

func (r *DescribeIntegrationStatisticsRecordsTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsRecordsTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "ProjectId")
	delete(f, "QueryDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationStatisticsRecordsTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsRecordsTrendResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrendsData []*IntegrationStatisticsTrendResult `json:"TrendsData,omitempty" name:"TrendsData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationStatisticsRecordsTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationStatisticsRecordsTrendResponseParams `json:"Response"`
}

func (r *DescribeIntegrationStatisticsRecordsTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsRecordsTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsRequestParams struct {
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`
}

type DescribeIntegrationStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`
}

func (r *DescribeIntegrationStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "ProjectId")
	delete(f, "QueryDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsResponseParams struct {
	// 总任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTask *int64 `json:"TotalTask,omitempty" name:"TotalTask"`

	// 生产态任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProdTask *int64 `json:"ProdTask,omitempty" name:"ProdTask"`

	// 开发态任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevTask *int64 `json:"DevTask,omitempty" name:"DevTask"`

	// 总读取条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalReadRecords *int64 `json:"TotalReadRecords,omitempty" name:"TotalReadRecords"`

	// 总写入条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalWriteRecords *int64 `json:"TotalWriteRecords,omitempty" name:"TotalWriteRecords"`

	// 总脏数据条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalErrorRecords *int64 `json:"TotalErrorRecords,omitempty" name:"TotalErrorRecords"`

	// 总告警事件数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalAlarmEvent *int64 `json:"TotalAlarmEvent,omitempty" name:"TotalAlarmEvent"`

	// 当天读取增长条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseReadRecords *int64 `json:"IncreaseReadRecords,omitempty" name:"IncreaseReadRecords"`

	// 当天写入增长条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseWriteRecords *int64 `json:"IncreaseWriteRecords,omitempty" name:"IncreaseWriteRecords"`

	// 当天脏数据增长条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseErrorRecords *int64 `json:"IncreaseErrorRecords,omitempty" name:"IncreaseErrorRecords"`

	// 当天告警事件增长数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseAlarmEvent *int64 `json:"IncreaseAlarmEvent,omitempty" name:"IncreaseAlarmEvent"`

	// 告警事件统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmEvent *string `json:"AlarmEvent,omitempty" name:"AlarmEvent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationStatisticsResponseParams `json:"Response"`
}

func (r *DescribeIntegrationStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsTaskStatusRequestParams struct {
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

type DescribeIntegrationStatisticsTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

func (r *DescribeIntegrationStatisticsTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "ProjectId")
	delete(f, "QueryDate")
	delete(f, "ExecutorGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationStatisticsTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsTaskStatusResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusData *string `json:"StatusData,omitempty" name:"StatusData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationStatisticsTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationStatisticsTaskStatusResponseParams `json:"Response"`
}

func (r *DescribeIntegrationStatisticsTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsTaskStatusTrendRequestParams struct {
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

type DescribeIntegrationStatisticsTaskStatusTrendRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

func (r *DescribeIntegrationStatisticsTaskStatusTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsTaskStatusTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "ProjectId")
	delete(f, "QueryDate")
	delete(f, "ExecutorGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationStatisticsTaskStatusTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationStatisticsTaskStatusTrendResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrendsData []*IntegrationStatisticsTrendResult `json:"TrendsData,omitempty" name:"TrendsData"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationStatisticsTaskStatusTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationStatisticsTaskStatusTrendResponseParams `json:"Response"`
}

func (r *DescribeIntegrationStatisticsTaskStatusTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationStatisticsTaskStatusTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type DescribeIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DescribeIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationTaskResponseParams struct {
	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationTaskResponseParams `json:"Response"`
}

func (r *DescribeIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationTasksRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分页第n页
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段信息
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 201. stream, 202. offline 默认实时
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type DescribeIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 分页第n页
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 查询filter
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序字段信息
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 201. stream, 202. offline 默认实时
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DescribeIntegrationTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationTasksResponseParams struct {
	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfoSet []*IntegrationTaskInfo `json:"TaskInfoSet,omitempty" name:"TaskInfoSet"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationTasksResponseParams `json:"Response"`
}

func (r *DescribeIntegrationTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationVersionNodesInfoRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// task version path
	TaskVersionPath *string `json:"TaskVersionPath,omitempty" name:"TaskVersionPath"`

	// task version
	TaskVersion *string `json:"TaskVersion,omitempty" name:"TaskVersion"`
}

type DescribeIntegrationVersionNodesInfoRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// task version path
	TaskVersionPath *string `json:"TaskVersionPath,omitempty" name:"TaskVersionPath"`

	// task version
	TaskVersion *string `json:"TaskVersion,omitempty" name:"TaskVersion"`
}

func (r *DescribeIntegrationVersionNodesInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationVersionNodesInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "TaskVersionPath")
	delete(f, "TaskVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationVersionNodesInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationVersionNodesInfoResponseParams struct {
	// 任务节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*IntegrationNodeInfo `json:"Nodes,omitempty" name:"Nodes"`

	// 任务映射信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mappings []*IntegrationNodeMapping `json:"Mappings,omitempty" name:"Mappings"`

	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeIntegrationVersionNodesInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeIntegrationVersionNodesInfoResponseParams `json:"Response"`
}

func (r *DescribeIntegrationVersionNodesInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIntegrationVersionNodesInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaTopicInfoRequestParams struct {
	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 数据源类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeKafkaTopicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 数据源类型
	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeKafkaTopicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaTopicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatasourceId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKafkaTopicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaTopicInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeKafkaTopicInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKafkaTopicInfoResponseParams `json:"Response"`
}

func (r *DescribeKafkaTopicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaTopicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineTaskTokenRequestParams struct {

}

type DescribeOfflineTaskTokenRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeOfflineTaskTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineTaskTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOfflineTaskTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOfflineTaskTokenResponseParams struct {
	// 长连接临时token
	Token *string `json:"Token,omitempty" name:"Token"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOfflineTaskTokenResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOfflineTaskTokenResponseParams `json:"Response"`
}

func (r *DescribeOfflineTaskTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOfflineTaskTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationalFunctionsRequestParams struct {
	// 场景类型：开发、使用
	Type *string `json:"Type,omitempty" name:"Type"`

	// 项目 ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 函数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}

type DescribeOrganizationalFunctionsRequest struct {
	*tchttp.BaseRequest
	
	// 场景类型：开发、使用
	Type *string `json:"Type,omitempty" name:"Type"`

	// 项目 ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 函数名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
}

func (r *DescribeOrganizationalFunctionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationalFunctionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "DisplayName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationalFunctionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationalFunctionsResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*OrganizationalFunction `json:"Content,omitempty" name:"Content"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeOrganizationalFunctionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOrganizationalFunctionsResponseParams `json:"Response"`
}

func (r *DescribeOrganizationalFunctionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOrganizationalFunctionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectRequestParams struct {
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否展示关联集群信息
	DescribeClusters *bool `json:"DescribeClusters,omitempty" name:"DescribeClusters"`

	// 是否展示关联执行组的信息，仅部分信息。
	DescribeExecutors *bool `json:"DescribeExecutors,omitempty" name:"DescribeExecutors"`

	// 默认不展示项目管理员信息
	DescribeAdminUsers *bool `json:"DescribeAdminUsers,omitempty" name:"DescribeAdminUsers"`

	// 默认不统计项目人员数量
	DescribeMemberCount *bool `json:"DescribeMemberCount,omitempty" name:"DescribeMemberCount"`

	// 默认不查询创建者的信息
	DescribeCreator *bool `json:"DescribeCreator,omitempty" name:"DescribeCreator"`

	// 项目名只在租户内唯一，一般用来转化为项目ID。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 是否展示关联集群信息
	DescribeClusters *bool `json:"DescribeClusters,omitempty" name:"DescribeClusters"`

	// 是否展示关联执行组的信息，仅部分信息。
	DescribeExecutors *bool `json:"DescribeExecutors,omitempty" name:"DescribeExecutors"`

	// 默认不展示项目管理员信息
	DescribeAdminUsers *bool `json:"DescribeAdminUsers,omitempty" name:"DescribeAdminUsers"`

	// 默认不统计项目人员数量
	DescribeMemberCount *bool `json:"DescribeMemberCount,omitempty" name:"DescribeMemberCount"`

	// 默认不查询创建者的信息
	DescribeCreator *bool `json:"DescribeCreator,omitempty" name:"DescribeCreator"`

	// 项目名只在租户内唯一，一般用来转化为项目ID。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

func (r *DescribeProjectRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DescribeClusters")
	delete(f, "DescribeExecutors")
	delete(f, "DescribeAdminUsers")
	delete(f, "DescribeMemberCount")
	delete(f, "DescribeCreator")
	delete(f, "ProjectName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeProjectResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectResponseParams `json:"Response"`
}

func (r *DescribeProjectResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskInstanceNodeInfoRequestParams struct {
	// 实时任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 工程id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeRealTimeTaskInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实时任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 工程id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeRealTimeTaskInstanceNodeInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealTimeTaskInstanceNodeInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealTimeTaskInstanceNodeInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskInstanceNodeInfoResponseParams struct {
	// 实时任务实例节点相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealTimeTaskInstanceNodeInfo *RealTimeTaskInstanceNodeInfo `json:"RealTimeTaskInstanceNodeInfo,omitempty" name:"RealTimeTaskInstanceNodeInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRealTimeTaskInstanceNodeInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRealTimeTaskInstanceNodeInfoResponseParams `json:"Response"`
}

func (r *DescribeRealTimeTaskInstanceNodeInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealTimeTaskInstanceNodeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskMetricOverviewRequestParams struct {
	// 无
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 无
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeRealTimeTaskMetricOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 无
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 无
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeRealTimeTaskMetricOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealTimeTaskMetricOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealTimeTaskMetricOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskMetricOverviewResponseParams struct {
	// 总读取记录数
	TotalRecordNumOfRead *uint64 `json:"TotalRecordNumOfRead,omitempty" name:"TotalRecordNumOfRead"`

	// 总读取字节数
	TotalRecordByteNumOfRead *uint64 `json:"TotalRecordByteNumOfRead,omitempty" name:"TotalRecordByteNumOfRead"`

	// 总写入记录数
	TotalRecordNumOfWrite *uint64 `json:"TotalRecordNumOfWrite,omitempty" name:"TotalRecordNumOfWrite"`

	// 总写入字节数 单位字节
	TotalRecordByteNumOfWrite *uint64 `json:"TotalRecordByteNumOfWrite,omitempty" name:"TotalRecordByteNumOfWrite"`

	// 总的脏记录数据
	TotalDirtyRecordNum *uint64 `json:"TotalDirtyRecordNum,omitempty" name:"TotalDirtyRecordNum"`

	// 总的脏字节数 单位字节
	TotalDirtyRecordByte *uint64 `json:"TotalDirtyRecordByte,omitempty" name:"TotalDirtyRecordByte"`

	// 运行时长 单位s
	TotalDuration *uint64 `json:"TotalDuration,omitempty" name:"TotalDuration"`

	// 开始运行时间
	BeginRunTime *string `json:"BeginRunTime,omitempty" name:"BeginRunTime"`

	// 目前运行到的时间
	EndRunTime *string `json:"EndRunTime,omitempty" name:"EndRunTime"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRealTimeTaskMetricOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRealTimeTaskMetricOverviewResponseParams `json:"Response"`
}

func (r *DescribeRealTimeTaskMetricOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealTimeTaskMetricOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskSpeedRequestParams struct {
	// 无
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 带毫秒的时间戳
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 带毫秒的时间戳
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 粒度，1或者5
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// 无
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeRealTimeTaskSpeedRequest struct {
	*tchttp.BaseRequest
	
	// 无
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 带毫秒的时间戳
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 带毫秒的时间戳
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 粒度，1或者5
	Granularity *uint64 `json:"Granularity,omitempty" name:"Granularity"`

	// 无
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeRealTimeTaskSpeedRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealTimeTaskSpeedRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Granularity")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealTimeTaskSpeedRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskSpeedResponseParams struct {
	// 同步速度条/s列表
	RecordsSpeedList []*RecordsSpeed `json:"RecordsSpeedList,omitempty" name:"RecordsSpeedList"`

	// 同步速度字节/s列表
	BytesSpeedList []*BytesSpeed `json:"BytesSpeedList,omitempty" name:"BytesSpeedList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRealTimeTaskSpeedResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRealTimeTaskSpeedResponseParams `json:"Response"`
}

func (r *DescribeRealTimeTaskSpeedResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealTimeTaskSpeedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedInstancesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 数据时间，格式yyyy-MM-dd HH:mm:ss
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 距离当前任务的层级距离，-1表示取父节点，1表示子节点
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeRelatedInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 数据时间，格式yyyy-MM-dd HH:mm:ss
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务id
	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`

	// 距离当前任务的层级距离，-1表示取父节点，1表示子节点
	Depth *int64 `json:"Depth,omitempty" name:"Depth"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeRelatedInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CurRunDate")
	delete(f, "TaskId")
	delete(f, "Depth")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRelatedInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRelatedInstancesResponseParams struct {
	// 无
	Data *DescribeTaskInstancesData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRelatedInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRelatedInstancesResponseParams `json:"Response"`
}

func (r *DescribeRelatedInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRelatedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceManagePathTreesRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 名字，供搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 文件夹类型
	DirType *string `json:"DirType,omitempty" name:"DirType"`
}

type DescribeResourceManagePathTreesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 名字，供搜索
	Name *string `json:"Name,omitempty" name:"Name"`

	// 文件类型
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 文件路径
	FilePath *string `json:"FilePath,omitempty" name:"FilePath"`

	// 文件夹类型
	DirType *string `json:"DirType,omitempty" name:"DirType"`
}

func (r *DescribeResourceManagePathTreesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceManagePathTreesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "FileType")
	delete(f, "FilePath")
	delete(f, "DirType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeResourceManagePathTreesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeResourceManagePathTreesResponseParams struct {
	// 响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ResourcePathTree `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeResourceManagePathTreesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeResourceManagePathTreesResponseParams `json:"Response"`
}

func (r *DescribeResourceManagePathTreesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeResourceManagePathTreesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandardRuleDetailInfoListRequestParams struct {
	// 空间、项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标准分类11编码映射 12数据过滤 13字符串转换 14数据元定义 15正则表达 16术语词典
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type DescribeStandardRuleDetailInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 空间、项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 标准分类11编码映射 12数据过滤 13字符串转换 14数据元定义 15正则表达 16术语词典
	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeStandardRuleDetailInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandardRuleDetailInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStandardRuleDetailInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStandardRuleDetailInfoListResponseParams struct {
	// 返回值
	// 注意：此字段可能返回 null，表示取不到有效值。
	StandardRuleDetailList *string `json:"StandardRuleDetailList,omitempty" name:"StandardRuleDetailList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStandardRuleDetailInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStandardRuleDetailInfoListResponseParams `json:"Response"`
}

func (r *DescribeStandardRuleDetailInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStandardRuleDetailInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamTaskLogListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// container名字
	Container *string `json:"Container,omitempty" name:"Container"`

	// 条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序类型 desc asc
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 作业运行的实例ID
	RunningOrderId *uint64 `json:"RunningOrderId,omitempty" name:"RunningOrderId"`
}

type DescribeStreamTaskLogListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 作业ID
	JobId *string `json:"JobId,omitempty" name:"JobId"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// container名字
	Container *string `json:"Container,omitempty" name:"Container"`

	// 条数
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 排序类型 desc asc
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 作业运行的实例ID
	RunningOrderId *uint64 `json:"RunningOrderId,omitempty" name:"RunningOrderId"`
}

func (r *DescribeStreamTaskLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamTaskLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "JobId")
	delete(f, "EndTime")
	delete(f, "StartTime")
	delete(f, "Container")
	delete(f, "Limit")
	delete(f, "OrderType")
	delete(f, "RunningOrderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamTaskLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamTaskLogListResponseParams struct {
	// 是否是全量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

	// 日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogContentList []*LogContentInfo `json:"LogContentList,omitempty" name:"LogContentList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeStreamTaskLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStreamTaskLogListResponseParams `json:"Response"`
}

func (r *DescribeStreamTaskLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStreamTaskLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableInfoListRequestParams struct {
	// 表名
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 如果是hive这里写rpc，如果是其他类型不传
	ConnectionType *string `json:"ConnectionType,omitempty" name:"ConnectionType"`

	// 数据库源类型
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`
}

type DescribeTableInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 表名
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 如果是hive这里写rpc，如果是其他类型不传
	ConnectionType *string `json:"ConnectionType,omitempty" name:"ConnectionType"`

	// 数据库源类型
	Catalog *string `json:"Catalog,omitempty" name:"Catalog"`
}

func (r *DescribeTableInfoListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableInfoListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "ConnectionType")
	delete(f, "Catalog")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableInfoListResponseParams struct {
	// 表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInfo []*TableInfo `json:"TableInfo,omitempty" name:"TableInfo"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableInfoListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableInfoListResponseParams `json:"Response"`
}

func (r *DescribeTableInfoListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableInfoListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableSchemaInfoRequestParams struct {
	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 表类型
	MsType *string `json:"MsType,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// HIVE传rpc
	ConnectionType *string `json:"ConnectionType,omitempty" name:"ConnectionType"`

	// 元数据Database下的Schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`
}

type DescribeTableSchemaInfoRequest struct {
	*tchttp.BaseRequest
	
	// 表名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 表类型
	MsType *string `json:"MsType,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// HIVE传rpc
	ConnectionType *string `json:"ConnectionType,omitempty" name:"ConnectionType"`

	// 元数据Database下的Schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`
}

func (r *DescribeTableSchemaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableSchemaInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "DatabaseName")
	delete(f, "MsType")
	delete(f, "DatasourceId")
	delete(f, "ConnectionType")
	delete(f, "SchemaName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableSchemaInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableSchemaInfoResponseParams struct {
	// 123
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaInfoList []*SchemaDetail `json:"SchemaInfoList,omitempty" name:"SchemaInfoList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTableSchemaInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableSchemaInfoResponseParams `json:"Response"`
}

func (r *DescribeTableSchemaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableSchemaInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskAlarmRegulationsRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型(201代表实时任务，202代表离线任务)
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 当前页
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 过滤条件(name有RegularStatus、AlarmLevel、AlarmIndicator、RegularName)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序条件(RegularId)
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`
}

type DescribeTaskAlarmRegulationsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型(201代表实时任务，202代表离线任务)
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 当前页
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 过滤条件(name有RegularStatus、AlarmLevel、AlarmIndicator、RegularName)
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 排序条件(RegularId)
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`
}

func (r *DescribeTaskAlarmRegulationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskAlarmRegulationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "TaskType")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskAlarmRegulationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskAlarmRegulationsResponseParams struct {
	// 任务告警规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAlarmInfos []*TaskAlarmInfo `json:"TaskAlarmInfos,omitempty" name:"TaskAlarmInfos"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskAlarmRegulationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskAlarmRegulationsResponseParams `json:"Response"`
}

func (r *DescribeTaskAlarmRegulationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskAlarmRegulationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务告警状态
	TaskAlarmStatus *int64 `json:"TaskAlarmStatus,omitempty" name:"TaskAlarmStatus"`
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务告警状态
	TaskAlarmStatus *int64 `json:"TaskAlarmStatus,omitempty" name:"TaskAlarmStatus"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "TaskAlarmStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskDetailResponseParams struct {
	// 任务详情1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TaskInfoData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskDetailResponseParams `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInstanceReportDetailRequestParams struct {
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务实例数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务实例运行时间
	IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`
}

type DescribeTaskInstanceReportDetailRequest struct {
	*tchttp.BaseRequest
	
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务实例数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务实例运行时间
	IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`
}

func (r *DescribeTaskInstanceReportDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstanceReportDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	delete(f, "IssueDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInstanceReportDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInstanceReportDetailResponseParams struct {
	// 任务实例运行指标概览
	Summary *InstanceReportSummary `json:"Summary,omitempty" name:"Summary"`

	// 任务实例读取节点运行指标
	ReadNode *InstanceReportReadNode `json:"ReadNode,omitempty" name:"ReadNode"`

	// 任务实例写入节点运行指标
	WriteNode *InstanceReportWriteNode `json:"WriteNode,omitempty" name:"WriteNode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskInstanceReportDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInstanceReportDetailResponseParams `json:"Response"`
}

func (r *DescribeTaskInstanceReportDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstanceReportDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInstanceRequestParams struct {
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务实例数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务实例运行时间
	IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`
}

type DescribeTaskInstanceRequest struct {
	*tchttp.BaseRequest
	
	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务实例数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 任务实例运行时间
	IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`
}

func (r *DescribeTaskInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	delete(f, "IssueDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInstanceResponseParams struct {
	// 任务实例详情
	TaskInstanceDetail *TaskInstanceDetail `json:"TaskInstanceDetail,omitempty" name:"TaskInstanceDetail"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInstanceResponseParams `json:"Response"`
}

func (r *DescribeTaskInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskInstancesData struct {
	// 实例列表
	Items []*TaskInstanceInfo `json:"Items,omitempty" name:"Items"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeTaskInstancesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 工作流id列表
	WorkflowIdList []*string `json:"WorkflowIdList,omitempty" name:"WorkflowIdList"`

	// 工作流名称列表，支持模糊搜索
	WorkflowNameList []*string `json:"WorkflowNameList,omitempty" name:"WorkflowNameList"`

	// 起始数据时间，格式yyyy-MM-dd HH:mm:ss
	DateFrom *string `json:"DateFrom,omitempty" name:"DateFrom"`

	// 结束数据时间，格式yyyy-MM-dd HH:mm:ss
	DateTo *string `json:"DateTo,omitempty" name:"DateTo"`

	// 任务id列表
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 任务名称列表，支持模糊搜索
	TaskNameList []*string `json:"TaskNameList,omitempty" name:"TaskNameList"`

	// 责任人名称列表
	InChargeList []*string `json:"InChargeList,omitempty" name:"InChargeList"`

	// 任务类型码列表，26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskTypeIdList []*int64 `json:"TaskTypeIdList,omitempty" name:"TaskTypeIdList"`

	// 实例状态列表，0等待事件，1等待上游，2等待运行，3运行中，4正在终止，5失败重试，6失败，7成功
	StateList []*string `json:"StateList,omitempty" name:"StateList"`

	// 周期类型列表，I分钟，H小时，D天，W周，M月，Y年，O一次性，C crontab
	TaskCycleUnitList []*string `json:"TaskCycleUnitList,omitempty" name:"TaskCycleUnitList"`

	// 实例类型，0补录实例，1周期实例，2非周期实例
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 排序字段信息列表，ScheduleDateTime / CostTime / StartTime / EndTime
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`
}

type DescribeTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 页号，默认为1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认为10，最大不超过200
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`

	// 工作流id列表
	WorkflowIdList []*string `json:"WorkflowIdList,omitempty" name:"WorkflowIdList"`

	// 工作流名称列表，支持模糊搜索
	WorkflowNameList []*string `json:"WorkflowNameList,omitempty" name:"WorkflowNameList"`

	// 起始数据时间，格式yyyy-MM-dd HH:mm:ss
	DateFrom *string `json:"DateFrom,omitempty" name:"DateFrom"`

	// 结束数据时间，格式yyyy-MM-dd HH:mm:ss
	DateTo *string `json:"DateTo,omitempty" name:"DateTo"`

	// 任务id列表
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 任务名称列表，支持模糊搜索
	TaskNameList []*string `json:"TaskNameList,omitempty" name:"TaskNameList"`

	// 责任人名称列表
	InChargeList []*string `json:"InChargeList,omitempty" name:"InChargeList"`

	// 任务类型码列表，26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	TaskTypeIdList []*int64 `json:"TaskTypeIdList,omitempty" name:"TaskTypeIdList"`

	// 实例状态列表，0等待事件，1等待上游，2等待运行，3运行中，4正在终止，5失败重试，6失败，7成功
	StateList []*string `json:"StateList,omitempty" name:"StateList"`

	// 周期类型列表，I分钟，H小时，D天，W周，M月，Y年，O一次性，C crontab
	TaskCycleUnitList []*string `json:"TaskCycleUnitList,omitempty" name:"TaskCycleUnitList"`

	// 实例类型，0补录实例，1周期实例，2非周期实例
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 排序字段信息列表，ScheduleDateTime / CostTime / StartTime / EndTime
	OrderFields []*OrderField `json:"OrderFields,omitempty" name:"OrderFields"`
}

func (r *DescribeTaskInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "WorkflowIdList")
	delete(f, "WorkflowNameList")
	delete(f, "DateFrom")
	delete(f, "DateTo")
	delete(f, "TaskIdList")
	delete(f, "TaskNameList")
	delete(f, "InChargeList")
	delete(f, "TaskTypeIdList")
	delete(f, "StateList")
	delete(f, "TaskCycleUnitList")
	delete(f, "InstanceType")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskInstancesResponseParams struct {
	// 无
	Data *DescribeTaskInstancesData `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskInstancesResponseParams `json:"Response"`
}

func (r *DescribeTaskInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLockStatusRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type DescribeTaskLockStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *DescribeTaskLockStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLockStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLockStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLockStatusResponseParams struct {
	// 任务锁状态信息
	TaskLockStatus *TaskLockStatus `json:"TaskLockStatus,omitempty" name:"TaskLockStatus"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskLockStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLockStatusResponseParams `json:"Response"`
}

func (r *DescribeTaskLockStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLockStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskReportDetailListRequestParams struct {
	// WeData项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 统计周期的开始日期，格式为 yyyy-MM-dd
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 统计周期的结束日期，格式为 yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 任务状态，多个状态用逗号连接
	StateList *string `json:"StateList,omitempty" name:"StateList"`

	// 排序字段名
	SortItem *string `json:"SortItem,omitempty" name:"SortItem"`

	// 升序或降序，传ASC或DESC
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// 页数，从1开始
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页的记录条数，默认10条
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeTaskReportDetailListRequest struct {
	*tchttp.BaseRequest
	
	// WeData项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 统计周期的开始日期，格式为 yyyy-MM-dd
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 统计周期的结束日期，格式为 yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 任务状态，多个状态用逗号连接
	StateList *string `json:"StateList,omitempty" name:"StateList"`

	// 排序字段名
	SortItem *string `json:"SortItem,omitempty" name:"SortItem"`

	// 升序或降序，传ASC或DESC
	SortType *string `json:"SortType,omitempty" name:"SortType"`

	// 页数，从1开始
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页的记录条数，默认10条
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeTaskReportDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskReportDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	delete(f, "StateList")
	delete(f, "SortItem")
	delete(f, "SortType")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskReportDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskReportDetailListResponseParams struct {
	// 页码，从1开始
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页的记录数
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 总页数
	TotalPage *uint64 `json:"TotalPage,omitempty" name:"TotalPage"`

	// 任务运行指标
	Items []*TaskReportDetail `json:"Items,omitempty" name:"Items"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskReportDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskReportDetailListResponseParams `json:"Response"`
}

func (r *DescribeTaskReportDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskReportDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskReportRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 统计周期的开始日期，格式为 yyyy-MM-dd
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 统计周期的结束日期，格式为 yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// WeData项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type DescribeTaskReportRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 统计周期的开始日期，格式为 yyyy-MM-dd
	BeginDate *string `json:"BeginDate,omitempty" name:"BeginDate"`

	// 统计周期的结束日期，格式为 yyyy-MM-dd
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// WeData项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeTaskReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskReportResponseParams struct {
	// 总读取条数
	TotalReadRecords *uint64 `json:"TotalReadRecords,omitempty" name:"TotalReadRecords"`

	// 总读取字节数，单位为Byte
	TotalReadBytes *uint64 `json:"TotalReadBytes,omitempty" name:"TotalReadBytes"`

	// 总写入条数
	TotalWriteRecords *uint64 `json:"TotalWriteRecords,omitempty" name:"TotalWriteRecords"`

	// 总写入字节数，单位为Byte
	TotalWriteBytes *uint64 `json:"TotalWriteBytes,omitempty" name:"TotalWriteBytes"`

	// 总脏数据条数
	TotalErrorRecords *uint64 `json:"TotalErrorRecords,omitempty" name:"TotalErrorRecords"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskReportResponseParams `json:"Response"`
}

func (r *DescribeTaskReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskScriptRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeTaskScriptRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeTaskScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskScriptResponseParams struct {
	// 任务脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TaskScriptContent `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTaskScriptResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskScriptResponseParams `json:"Response"`
}

func (r *DescribeTaskScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksByPageRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

type DescribeTasksByPageRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
}

func (r *DescribeTasksByPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksByPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTasksByPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTasksByPageResponseParams struct {
	// 无1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TaskInfoDataPage `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeTasksByPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTasksByPageResponseParams `json:"Response"`
}

func (r *DescribeTasksByPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTasksByPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DryRunDIOfflineTaskRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 资源组Id
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 默认 27
	TaskTypeId *uint64 `json:"TaskTypeId,omitempty" name:"TaskTypeId"`
}

type DryRunDIOfflineTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 资源组Id
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 默认 27
	TaskTypeId *uint64 `json:"TaskTypeId,omitempty" name:"TaskTypeId"`
}

func (r *DryRunDIOfflineTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DryRunDIOfflineTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "ResourceGroup")
	delete(f, "TaskTypeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DryRunDIOfflineTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DryRunDIOfflineTaskResponseParams struct {
	// 数据时间
	CurrentRunDate *string `json:"CurrentRunDate,omitempty" name:"CurrentRunDate"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务实例唯一key
	TaskInstanceKey *string `json:"TaskInstanceKey,omitempty" name:"TaskInstanceKey"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DryRunDIOfflineTaskResponse struct {
	*tchttp.BaseResponse
	Response *DryRunDIOfflineTaskResponseParams `json:"Response"`
}

func (r *DryRunDIOfflineTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DryRunDIOfflineTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 过滤值列表
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type Folder struct {
	// 文件ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 文件夹名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 所属项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type ForceSucInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

type ForceSucInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

func (r *ForceSucInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForceSucInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ForceSucInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ForceSucInstancesResponseParams struct {
	// 返回实例批量终止结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ForceSucInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ForceSucInstancesResponseParams `json:"Response"`
}

func (r *ForceSucInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ForceSucInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksByMultiWorkflowRequestParams struct {
	// 工作流Id集合
	WorkFlowIds []*string `json:"WorkFlowIds,omitempty" name:"WorkFlowIds"`
}

type FreezeTasksByMultiWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 工作流Id集合
	WorkFlowIds []*string `json:"WorkFlowIds,omitempty" name:"WorkFlowIds"`
}

func (r *FreezeTasksByMultiWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksByMultiWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkFlowIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeTasksByMultiWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksByMultiWorkflowResponseParams struct {
	// 操作结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FreezeTasksByMultiWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *FreezeTasksByMultiWorkflowResponseParams `json:"Response"`
}

func (r *FreezeTasksByMultiWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksByMultiWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksRequestParams struct {
	// 任务列表
	Tasks []*SimpleTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 任务操作是否消息通知下游任务责任人
	OperateIsInform *bool `json:"OperateIsInform,omitempty" name:"OperateIsInform"`
}

type FreezeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务列表
	Tasks []*SimpleTaskInfo `json:"Tasks,omitempty" name:"Tasks"`

	// 任务操作是否消息通知下游任务责任人
	OperateIsInform *bool `json:"OperateIsInform,omitempty" name:"OperateIsInform"`
}

func (r *FreezeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tasks")
	delete(f, "OperateIsInform")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksResponseParams struct {
	// 操作结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type FreezeTasksResponse struct {
	*tchttp.BaseResponse
	Response *FreezeTasksResponseParams `json:"Response"`
}

func (r *FreezeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionResource struct {
	// 资源路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitempty" name:"Path"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源唯一标识
	Id *string `json:"Id,omitempty" name:"Id"`

	// 资源 MD5 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 默认是 hdfs
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type FunctionTypeOrKind struct {
	// 无
	Name *string `json:"Name,omitempty" name:"Name"`

	// 无
	ZhName *string `json:"ZhName,omitempty" name:"ZhName"`

	// 无
	EnName *string `json:"EnName,omitempty" name:"EnName"`
}

type FunctionVersion struct {
	// 版本号：V0 V1 V2
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 提交人 ID
	UserId *string `json:"UserId,omitempty" name:"UserId"`

	// 变更类型：ADD、MODIFY
	Type *string `json:"Type,omitempty" name:"Type"`

	// 备注
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 提交时间: UTC 秒数
	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`

	// 提交人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitempty" name:"UserName"`

	// 版本内容：json string 格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitempty" name:"Content"`
}

// Predefined struct for user
type GenHiveTableDDLSqlRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 目标数据库
	SinkDatabase *string `json:"SinkDatabase,omitempty" name:"SinkDatabase"`

	// 节点id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 元数据类型(MYSQL、ORACLE)
	MsType *string `json:"MsType,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 来源库
	SourceDatabase *string `json:"SourceDatabase,omitempty" name:"SourceDatabase"`

	// 来源表
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 目标表元数据类型(HIVE、GBASE)
	SinkType *string `json:"SinkType,omitempty" name:"SinkType"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`
}

type GenHiveTableDDLSqlRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 目标数据库
	SinkDatabase *string `json:"SinkDatabase,omitempty" name:"SinkDatabase"`

	// 节点id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 元数据类型(MYSQL、ORACLE)
	MsType *string `json:"MsType,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 来源库
	SourceDatabase *string `json:"SourceDatabase,omitempty" name:"SourceDatabase"`

	// 来源表
	TableName *string `json:"TableName,omitempty" name:"TableName"`

	// 目标表元数据类型(HIVE、GBASE)
	SinkType *string `json:"SinkType,omitempty" name:"SinkType"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitempty" name:"SchemaName"`
}

func (r *GenHiveTableDDLSqlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenHiveTableDDLSqlRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SinkDatabase")
	delete(f, "Id")
	delete(f, "MsType")
	delete(f, "DatasourceId")
	delete(f, "SourceDatabase")
	delete(f, "TableName")
	delete(f, "SinkType")
	delete(f, "SchemaName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenHiveTableDDLSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenHiveTableDDLSqlResponseParams struct {
	// 生成的ddl语句
	DDLSql *string `json:"DDLSql,omitempty" name:"DDLSql"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GenHiveTableDDLSqlResponse struct {
	*tchttp.BaseResponse
	Response *GenHiveTableDDLSqlResponseParams `json:"Response"`
}

func (r *GenHiveTableDDLSqlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenHiveTableDDLSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralTaskParam struct {
	// 通用任务参数类型,例：SPARK_SQL
	Type *string `json:"Type,omitempty" name:"Type"`

	// 通用任务参数内容,直接作用于任务的参数。不同参数用;
	// 分割
	Value *string `json:"Value,omitempty" name:"Value"`
}

// Predefined struct for user
type GetIntegrationNodeColumnSchemaRequestParams struct {
	// 字段示例（json格式）
	ColumnContent *string `json:"ColumnContent,omitempty" name:"ColumnContent"`

	// 数据源类型
	DatasourceType *string `json:"DatasourceType,omitempty" name:"DatasourceType"`
}

type GetIntegrationNodeColumnSchemaRequest struct {
	*tchttp.BaseRequest
	
	// 字段示例（json格式）
	ColumnContent *string `json:"ColumnContent,omitempty" name:"ColumnContent"`

	// 数据源类型
	DatasourceType *string `json:"DatasourceType,omitempty" name:"DatasourceType"`
}

func (r *GetIntegrationNodeColumnSchemaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetIntegrationNodeColumnSchemaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ColumnContent")
	delete(f, "DatasourceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetIntegrationNodeColumnSchemaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetIntegrationNodeColumnSchemaResponseParams struct {
	// 字段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schemas []*IntegrationNodeSchema `json:"Schemas,omitempty" name:"Schemas"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetIntegrationNodeColumnSchemaResponse struct {
	*tchttp.BaseResponse
	Response *GetIntegrationNodeColumnSchemaResponseParams `json:"Response"`
}

func (r *GetIntegrationNodeColumnSchemaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetIntegrationNodeColumnSchemaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOfflineDIInstanceListRequestParams struct {
	// 第几页
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchConditionNew `json:"SearchCondition,omitempty" name:"SearchCondition"`
}

type GetOfflineDIInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 第几页
	PageIndex *uint64 `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchConditionNew `json:"SearchCondition,omitempty" name:"SearchCondition"`
}

func (r *GetOfflineDIInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOfflineDIInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "SearchCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOfflineDIInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOfflineDIInstanceListResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 实例详情
	List []*OfflineInstance `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetOfflineDIInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *GetOfflineDIInstanceListResponseParams `json:"Response"`
}

func (r *GetOfflineDIInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOfflineDIInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOfflineInstanceListRequestParams struct {
	// 第几页
	PageIndex *string `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchCondition `json:"SearchCondition,omitempty" name:"SearchCondition"`
}

type GetOfflineInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 第几页
	PageIndex *string `json:"PageIndex,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchCondition `json:"SearchCondition,omitempty" name:"SearchCondition"`
}

func (r *GetOfflineInstanceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOfflineInstanceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "SearchCondition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetOfflineInstanceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetOfflineInstanceListResponseParams struct {
	// 总条数
	Total *uint64 `json:"Total,omitempty" name:"Total"`

	// 实例详情
	List []*OfflineInstance `json:"List,omitempty" name:"List"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetOfflineInstanceListResponse struct {
	*tchttp.BaseResponse
	Response *GetOfflineInstanceListResponseParams `json:"Response"`
}

func (r *GetOfflineInstanceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetOfflineInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InLongAgentDetail struct {
	// Agent ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// Agent Name
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// Agent状态(running运行中，initializing 操作中，failed心跳异常)
	Status *string `json:"Status,omitempty" name:"Status"`

	// Agent状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`

	// 集群类型，1：TKE Agent，2：BOSS SDK，默认：1
	AgentType *uint64 `json:"AgentType,omitempty" name:"AgentType"`

	// 采集来源
	Source *string `json:"Source,omitempty" name:"Source"`

	// VPC
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 集成资源组Id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`

	// 集成资源组名称
	ExecutorGroupName *string `json:"ExecutorGroupName,omitempty" name:"ExecutorGroupName"`

	// 关联任务数
	TaskCount *uint64 `json:"TaskCount,omitempty" name:"TaskCount"`
}

type InLongAgentTask struct {
	// 集成任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 集成任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 集成任务状态
	TaskStatus *string `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type InLongTkeDetail struct {
	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// TKE集群状态 (Running 运行中 Creating 创建中 Idling 闲置中 Abnormal 异常)
	Status *string `json:"Status,omitempty" name:"Status"`

	// 是否安装Agent，true: 是，false: 否
	HasAgent *bool `json:"HasAgent,omitempty" name:"HasAgent"`

	// 采集器ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// TKE集群区域ID
	TkeRegion *string `json:"TkeRegion,omitempty" name:"TkeRegion"`

	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER
	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type InstanceInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

type InstanceLog struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 尝试运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *string `json:"Tries,omitempty" name:"Tries"`

	// 日志更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitempty" name:"LastUpdate"`

	// 日志所在节点
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 文件名  含全路径
	OriginFileName *string `json:"OriginFileName,omitempty" name:"OriginFileName"`

	// 日志创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实例日志类型, run: 运行; kill: 终止
	InstanceLogType *string `json:"InstanceLogType,omitempty" name:"InstanceLogType"`

	// 运行耗时
	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`
}

type InstanceNodeInfo struct {
	// 读取节点SOURCE 写入节点SINK
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type InstanceReportReadNode struct {
	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 数据来源
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 总条数
	TotalReadRecords *uint64 `json:"TotalReadRecords,omitempty" name:"TotalReadRecords"`

	// 总字节数
	TotalReadBytes *uint64 `json:"TotalReadBytes,omitempty" name:"TotalReadBytes"`

	// 速度（条/秒）
	RecordSpeed *uint64 `json:"RecordSpeed,omitempty" name:"RecordSpeed"`

	// 吞吐（Byte/秒）
	ByteSpeed *float64 `json:"ByteSpeed,omitempty" name:"ByteSpeed"`

	// 脏数据条数
	TotalErrorRecords *uint64 `json:"TotalErrorRecords,omitempty" name:"TotalErrorRecords"`
}

type InstanceReportSummary struct {
	// 总读取记录数
	TotalReadRecords *uint64 `json:"TotalReadRecords,omitempty" name:"TotalReadRecords"`

	// 总读取字节数
	TotalReadBytes *uint64 `json:"TotalReadBytes,omitempty" name:"TotalReadBytes"`

	// 总写入记录数
	TotalWriteRecords *uint64 `json:"TotalWriteRecords,omitempty" name:"TotalWriteRecords"`

	// 总写入字节数
	TotalWriteBytes *uint64 `json:"TotalWriteBytes,omitempty" name:"TotalWriteBytes"`

	// 速率（条/秒）
	RecordSpeed *uint64 `json:"RecordSpeed,omitempty" name:"RecordSpeed"`

	// 流量（Byte/秒）
	ByteSpeed *float64 `json:"ByteSpeed,omitempty" name:"ByteSpeed"`

	// 脏数据记录数
	TotalErrorRecords *uint64 `json:"TotalErrorRecords,omitempty" name:"TotalErrorRecords"`

	// 脏数据字节数
	TotalErrorBytes *uint64 `json:"TotalErrorBytes,omitempty" name:"TotalErrorBytes"`

	// 任务运行总时长
	TotalRunDuration *uint64 `json:"TotalRunDuration,omitempty" name:"TotalRunDuration"`

	// 任务开始运行时间
	BeginRunTime *string `json:"BeginRunTime,omitempty" name:"BeginRunTime"`

	// 任务结束运行时间
	EndRunTime *string `json:"EndRunTime,omitempty" name:"EndRunTime"`
}

type InstanceReportWriteNode struct {
	// 节点名称
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 数据来源
	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`

	// 总条数
	TotalWriteRecords *uint64 `json:"TotalWriteRecords,omitempty" name:"TotalWriteRecords"`

	// 总字节数
	TotalWriteBytes *uint64 `json:"TotalWriteBytes,omitempty" name:"TotalWriteBytes"`

	// 速度（条/秒）
	RecordSpeed *uint64 `json:"RecordSpeed,omitempty" name:"RecordSpeed"`

	// 吞吐（Byte/秒）
	ByteSpeed *float64 `json:"ByteSpeed,omitempty" name:"ByteSpeed"`

	// 脏数据条数
	TotalErrorRecords *uint64 `json:"TotalErrorRecords,omitempty" name:"TotalErrorRecords"`
}

type IntegrationNodeDetail struct {
	// 集成节点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集成节点类型
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 节点配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitempty" name:"Config"`

	// 节点扩展配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitempty" name:"ExtConfig"`

	// 节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema []*IntegrationNodeSchema `json:"Schema,omitempty" name:"Schema"`

	// 节点映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeMapping *IntegrationNodeMapping `json:"NodeMapping,omitempty" name:"NodeMapping"`

	// owner uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`
}

type IntegrationNodeInfo struct {
	// 集成节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 集成节点所属任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 集成节点名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 集成节点类型
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" name:"DataSourceType"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitempty" name:"DatasourceId"`

	// 节点配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitempty" name:"Config"`

	// 节点扩展配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitempty" name:"ExtConfig"`

	// 节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema []*IntegrationNodeSchema `json:"Schema,omitempty" name:"Schema"`

	// 节点映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeMapping *IntegrationNodeMapping `json:"NodeMapping,omitempty" name:"NodeMapping"`

	// 应用id
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 操作人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// owner uin
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type IntegrationNodeMapping struct {
	// 源节点id
	SourceId *string `json:"SourceId,omitempty" name:"SourceId"`

	// 目标节点id
	SinkId *string `json:"SinkId,omitempty" name:"SinkId"`

	// 源节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceSchema []*IntegrationNodeSchema `json:"SourceSchema,omitempty" name:"SourceSchema"`

	// 节点schema映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaMappings []*IntegrationNodeSchemaMapping `json:"SchemaMappings,omitempty" name:"SchemaMappings"`

	// 节点映射扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitempty" name:"ExtConfig"`
}

type IntegrationNodeSchema struct {
	// schema id
	Id *string `json:"Id,omitempty" name:"Id"`

	// schema名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// schema类型
	Type *string `json:"Type,omitempty" name:"Type"`

	// schema值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// schema拓展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*RecordField `json:"Properties,omitempty" name:"Properties"`
}

type IntegrationNodeSchemaMapping struct {
	// 源schema id
	SourceSchemaId *string `json:"SourceSchemaId,omitempty" name:"SourceSchemaId"`

	// 目标schema id
	SinkSchemaId *string `json:"SinkSchemaId,omitempty" name:"SinkSchemaId"`
}

type IntegrationStatisticsTrendResult struct {
	// 统计属性名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticName []*string `json:"StatisticName,omitempty" name:"StatisticName"`

	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticValue []*int64 `json:"StatisticValue,omitempty" name:"StatisticValue"`

	// 统计项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticType *string `json:"StatisticType,omitempty" name:"StatisticType"`
}

type IntegrationTaskInfo struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 同步类型1.解决方案(整库迁移),2.单表同步
	SyncType *int64 `json:"SyncType,omitempty" name:"SyncType"`

	// 201.实时,202.离线
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 任务所属工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务调度id(oceanus or us等作业id)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleTaskId *string `json:"ScheduleTaskId,omitempty" name:"ScheduleTaskId"`

	// 任务组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupId *string `json:"TaskGroupId,omitempty" name:"TaskGroupId"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitempty" name:"CreatorUin"`

	// 操作人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// owner uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 任务状态1.初始化,2.操作中,3.运行中,4.暂停,5.任务停止中,6.停止,7.执行失败,8.已删除,9.已锁定,10.配置过期,11.提交中,12.提交成功,13.提交失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*IntegrationNodeInfo `json:"Nodes,omitempty" name:"Nodes"`

	// 执行资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorId *string `json:"ExecutorId,omitempty" name:"ExecutorId"`

	// 任务配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitempty" name:"Config"`

	// 任务扩展配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitempty" name:"ExtConfig"`

	// 任务执行context信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteContext []*RecordField `json:"ExecuteContext,omitempty" name:"ExecuteContext"`

	// 节点映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mappings []*IntegrationNodeMapping `json:"Mappings,omitempty" name:"Mappings"`

	// 任务模式：1.画布模式，2.flink jar
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskMode *string `json:"TaskMode,omitempty" name:"TaskMode"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Incharge *string `json:"Incharge,omitempty" name:"Incharge"`

	// 离线新增参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTaskAddEntity *OfflineTaskAddParam `json:"OfflineTaskAddEntity,omitempty" name:"OfflineTaskAddEntity"`

	// group name
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitempty" name:"ExecutorGroupName"`

	// url
	// 注意：此字段可能返回 null，表示取不到有效值。
	InLongManagerUrl *string `json:"InLongManagerUrl,omitempty" name:"InLongManagerUrl"`

	// stream id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InLongStreamId *string `json:"InLongStreamId,omitempty" name:"InLongStreamId"`

	// version
	// 注意：此字段可能返回 null，表示取不到有效值。
	InLongManagerVersion *string `json:"InLongManagerVersion,omitempty" name:"InLongManagerVersion"`
}

// Predefined struct for user
type KillInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

type KillInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`
}

func (r *KillInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Instances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillInstancesResponseParams struct {
	// 返回实例批量终止结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type KillInstancesResponse struct {
	*tchttp.BaseResponse
	Response *KillInstancesResponseParams `json:"Response"`
}

func (r *KillInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Label struct {
	// 类型值。
	Value *string `json:"Value,omitempty" name:"Value"`

	// 类型名称。
	Text *string `json:"Text,omitempty" name:"Text"`
}

// Predefined struct for user
type LockIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type LockIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *LockIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "LockIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type LockIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type LockIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *LockIntegrationTaskResponseParams `json:"Response"`
}

func (r *LockIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *LockIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogContent struct {
	// 日志时间戳，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 日志包id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 日志内容
	Log *string `json:"Log,omitempty" name:"Log"`
}

type LogContentInfo struct {
	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitempty" name:"Log"`

	// 日志组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 日志Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 日志所属的容器名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitempty" name:"ContainerName"`
}

// Predefined struct for user
type MakeUpTasksNewRequestParams struct {
	// 补录的当前任务的taskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 补录选项标识，1表示当前任务；2表示当前及下游任务；3表示下游任务
	MakeUpType *uint64 `json:"MakeUpType,omitempty" name:"MakeUpType"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// true: 检查父任务实例状态；false: 不检查父任务实例状态
	CheckParent *bool `json:"CheckParent,omitempty" name:"CheckParent"`
}

type MakeUpTasksNewRequest struct {
	*tchttp.BaseRequest
	
	// 补录的当前任务的taskId数组
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 补录选项标识，1表示当前任务；2表示当前及下游任务；3表示下游任务
	MakeUpType *uint64 `json:"MakeUpType,omitempty" name:"MakeUpType"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// true: 检查父任务实例状态；false: 不检查父任务实例状态
	CheckParent *bool `json:"CheckParent,omitempty" name:"CheckParent"`
}

func (r *MakeUpTasksNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpTasksNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MakeUpType")
	delete(f, "ProjectId")
	delete(f, "CheckParent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MakeUpTasksNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MakeUpTasksNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MakeUpTasksNewResponse struct {
	*tchttp.BaseResponse
	Response *MakeUpTasksNewResponseParams `json:"Response"`
}

func (r *MakeUpTasksNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpTasksNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MakeUpWorkflowNewRequestParams struct {
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type MakeUpWorkflowNewRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitempty" name:"WorkFlowId"`

	// 补录开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 补录结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *MakeUpWorkflowNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpWorkflowNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkFlowId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MakeUpWorkflowNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MakeUpWorkflowNewResponseParams struct {
	// 返回补录成功或失败的任务数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type MakeUpWorkflowNewResponse struct {
	*tchttp.BaseResponse
	Response *MakeUpWorkflowNewResponseParams `json:"Response"`
}

func (r *MakeUpWorkflowNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MakeUpWorkflowNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataSourceRequestParams struct {
	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据源ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitempty" name:"Collect"`

	// 项目id
	OwnerProjectId *string `json:"OwnerProjectId,omitempty" name:"OwnerProjectId"`

	// 项目名称
	OwnerProjectName *string `json:"OwnerProjectName,omitempty" name:"OwnerProjectName"`

	// 项目中文名
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitempty" name:"OwnerProjectIdent"`

	// cos bucket
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitempty" name:"COSRegion"`
}

type ModifyDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitempty" name:"Type"`

	// 数据源ID
	ID *uint64 `json:"ID,omitempty" name:"ID"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitempty" name:"Collect"`

	// 项目id
	OwnerProjectId *string `json:"OwnerProjectId,omitempty" name:"OwnerProjectId"`

	// 项目名称
	OwnerProjectName *string `json:"OwnerProjectName,omitempty" name:"OwnerProjectName"`

	// 项目中文名
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitempty" name:"OwnerProjectIdent"`

	// cos bucket
	COSBucket *string `json:"COSBucket,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitempty" name:"COSRegion"`
}

func (r *ModifyDataSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Category")
	delete(f, "Type")
	delete(f, "ID")
	delete(f, "BizParams")
	delete(f, "Params")
	delete(f, "Description")
	delete(f, "Display")
	delete(f, "DatabaseName")
	delete(f, "Instance")
	delete(f, "Status")
	delete(f, "ClusterId")
	delete(f, "Collect")
	delete(f, "OwnerProjectId")
	delete(f, "OwnerProjectName")
	delete(f, "OwnerProjectIdent")
	delete(f, "COSBucket")
	delete(f, "COSRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataSourceResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyDataSourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDataSourceResponseParams `json:"Response"`
}

func (r *ModifyDataSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

type ModifyFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitempty" name:"ParentsFolderId"`
}

func (r *ModifyFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "FolderId")
	delete(f, "ParentsFolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyFolderResponseParams struct {
	// true代表成功，false代表失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyFolderResponse struct {
	*tchttp.BaseResponse
	Response *ModifyFolderResponseParams `json:"Response"`
}

func (r *ModifyFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationNodeRequestParams struct {
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 区分画布模式和表单模式
	TaskMode *uint64 `json:"TaskMode,omitempty" name:"TaskMode"`
}

type ModifyIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 区分画布模式和表单模式
	TaskMode *uint64 `json:"TaskMode,omitempty" name:"TaskMode"`
}

func (r *ModifyIntegrationNodeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationNodeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeInfo")
	delete(f, "ProjectId")
	delete(f, "TaskType")
	delete(f, "TaskMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntegrationNodeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationNodeResponseParams struct {
	// 节点id
	Id *string `json:"Id,omitempty" name:"Id"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIntegrationNodeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntegrationNodeResponseParams `json:"Response"`
}

func (r *ModifyIntegrationNodeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationTaskRequestParams struct {
	// 任务信息
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 默认false . 为true时表示走回滚节点逻辑
	RollbackFlag *bool `json:"RollbackFlag,omitempty" name:"RollbackFlag"`
}

type ModifyIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务信息
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 默认false . 为true时表示走回滚节点逻辑
	RollbackFlag *bool `json:"RollbackFlag,omitempty" name:"RollbackFlag"`
}

func (r *ModifyIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskInfo")
	delete(f, "ProjectId")
	delete(f, "RollbackFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationTaskResponseParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *ModifyIntegrationTaskResponseParams `json:"Response"`
}

func (r *ModifyIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskAlarmRegularRequestParams struct {
	// 主键ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ModifyTaskAlarmRegularRequest struct {
	*tchttp.BaseRequest
	
	// 主键ID
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ModifyTaskAlarmRegularRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskAlarmRegularRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "TaskAlarmInfo")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskAlarmRegularRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskAlarmRegularResponseParams struct {
	// 判断是否修改成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskAlarmRegularResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskAlarmRegularResponseParams `json:"Response"`
}

func (r *ModifyTaskAlarmRegularResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskAlarmRegularResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskInfoRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 执行时间，单位分钟，天/周/月/年调度才有。比如天调度，每天的02:00点执行一次，delayTime就是120分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 新的任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 失败重试间隔,单位分钟，创建任务的时候已经给了默认值
	RetryWait *int64 `json:"RetryWait,omitempty" name:"RetryWait"`

	// 失败重试次数，创建任务的时候已经给了默认值
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 是否可重试，1代表可以重试
	Retriable *int64 `json:"Retriable,omitempty" name:"Retriable"`

	// 运行优先级，4高 5中 6低
	RunPriority *int64 `json:"RunPriority,omitempty" name:"RunPriority"`

	// 任务的扩展配置
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`

	// 执行资源组id，需要去资源管理服务上创建调度资源组，并且绑定cvm机器
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 资源池队列名称
	YarnQueue *string `json:"YarnQueue,omitempty" name:"YarnQueue"`

	// 资源组下具体执行机，any 表示可以跑在任意一台。
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 责任人
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 任务备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 任务参数
	TaskParamInfos []*ParamInfo `json:"TaskParamInfos,omitempty" name:"TaskParamInfos"`

	// 源数据源
	SourceServer *string `json:"SourceServer,omitempty" name:"SourceServer"`

	// 目标数据源
	TargetServer *string `json:"TargetServer,omitempty" name:"TargetServer"`

	// 是否支持工作流依赖 yes / no 默认 no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`

	// 依赖配置
	DependencyConfigDTOs []*DependencyConfig `json:"DependencyConfigDTOs,omitempty" name:"DependencyConfigDTOs"`
}

type ModifyTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 执行时间，单位分钟，天/周/月/年调度才有。比如天调度，每天的02:00点执行一次，delayTime就是120分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 新的任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 失败重试间隔,单位分钟，创建任务的时候已经给了默认值
	RetryWait *int64 `json:"RetryWait,omitempty" name:"RetryWait"`

	// 失败重试次数，创建任务的时候已经给了默认值
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 是否可重试，1代表可以重试
	Retriable *int64 `json:"Retriable,omitempty" name:"Retriable"`

	// 运行优先级，4高 5中 6低
	RunPriority *int64 `json:"RunPriority,omitempty" name:"RunPriority"`

	// 任务的扩展配置
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`

	// 执行资源组id，需要去资源管理服务上创建调度资源组，并且绑定cvm机器
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 资源池队列名称
	YarnQueue *string `json:"YarnQueue,omitempty" name:"YarnQueue"`

	// 资源组下具体执行机，any 表示可以跑在任意一台。
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 责任人
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 任务备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 任务参数
	TaskParamInfos []*ParamInfo `json:"TaskParamInfos,omitempty" name:"TaskParamInfos"`

	// 源数据源
	SourceServer *string `json:"SourceServer,omitempty" name:"SourceServer"`

	// 目标数据源
	TargetServer *string `json:"TargetServer,omitempty" name:"TargetServer"`

	// 是否支持工作流依赖 yes / no 默认 no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`

	// 依赖配置
	DependencyConfigDTOs []*DependencyConfig `json:"DependencyConfigDTOs,omitempty" name:"DependencyConfigDTOs"`
}

func (r *ModifyTaskInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "DelayTime")
	delete(f, "StartupTime")
	delete(f, "SelfDepend")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskAction")
	delete(f, "CycleType")
	delete(f, "CycleStep")
	delete(f, "CrontabExpression")
	delete(f, "ExecutionStartTime")
	delete(f, "ExecutionEndTime")
	delete(f, "TaskName")
	delete(f, "RetryWait")
	delete(f, "TryLimit")
	delete(f, "Retriable")
	delete(f, "RunPriority")
	delete(f, "TaskExt")
	delete(f, "ResourceGroup")
	delete(f, "YarnQueue")
	delete(f, "BrokerIp")
	delete(f, "InCharge")
	delete(f, "Notes")
	delete(f, "TaskParamInfos")
	delete(f, "SourceServer")
	delete(f, "TargetServer")
	delete(f, "DependencyWorkflow")
	delete(f, "DependencyConfigDTOs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskInfoResponseParams struct {
	// 执行结果
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskInfoResponseParams `json:"Response"`
}

func (r *ModifyTaskInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskLinksRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父任务ID
	TaskFrom *string `json:"TaskFrom,omitempty" name:"TaskFrom"`

	// 子任务ID
	TaskTo *string `json:"TaskTo,omitempty" name:"TaskTo"`

	// 子任务工作流
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 父任务工作流
	RealFromWorkflowId *string `json:"RealFromWorkflowId,omitempty" name:"RealFromWorkflowId"`

	// 父子任务之间的依赖关系
	LinkDependencyType *string `json:"LinkDependencyType,omitempty" name:"LinkDependencyType"`
}

type ModifyTaskLinksRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 父任务ID
	TaskFrom *string `json:"TaskFrom,omitempty" name:"TaskFrom"`

	// 子任务ID
	TaskTo *string `json:"TaskTo,omitempty" name:"TaskTo"`

	// 子任务工作流
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 父任务工作流
	RealFromWorkflowId *string `json:"RealFromWorkflowId,omitempty" name:"RealFromWorkflowId"`

	// 父子任务之间的依赖关系
	LinkDependencyType *string `json:"LinkDependencyType,omitempty" name:"LinkDependencyType"`
}

func (r *ModifyTaskLinksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskLinksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskFrom")
	delete(f, "TaskTo")
	delete(f, "WorkflowId")
	delete(f, "RealFromWorkflowId")
	delete(f, "LinkDependencyType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskLinksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskLinksResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskLinksResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskLinksResponseParams `json:"Response"`
}

func (r *ModifyTaskLinksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskLinksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskNameRequestParams struct {
	// 名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

type ModifyTaskNameRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 备注
	Notes *string `json:"Notes,omitempty" name:"Notes"`
}

func (r *ModifyTaskNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskNameRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskName")
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "Notes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskNameRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskNameResponseParams struct {
	// 结果
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskNameResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskNameResponseParams `json:"Response"`
}

func (r *ModifyTaskNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskScriptRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 脚本内容 base64编码
	ScriptContent *string `json:"ScriptContent,omitempty" name:"ScriptContent"`

	// 集成任务脚本配置
	IntegrationNodeDetails []*IntegrationNodeDetail `json:"IntegrationNodeDetails,omitempty" name:"IntegrationNodeDetails"`
}

type ModifyTaskScriptRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 脚本内容 base64编码
	ScriptContent *string `json:"ScriptContent,omitempty" name:"ScriptContent"`

	// 集成任务脚本配置
	IntegrationNodeDetails []*IntegrationNodeDetail `json:"IntegrationNodeDetails,omitempty" name:"IntegrationNodeDetails"`
}

func (r *ModifyTaskScriptRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskScriptRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "ScriptContent")
	delete(f, "IntegrationNodeDetails")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskScriptRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskScriptResponseParams struct {
	// 详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *CommonContent `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyTaskScriptResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTaskScriptResponseParams `json:"Response"`
}

func (r *ModifyTaskScriptResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTaskScriptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowInfoRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 责任人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 责任人id
	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作流所属用户分组id  若有多个,分号隔开: a;b;c
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitempty" name:"WorkflowParams"`

	// 用于配置优化参数（线程、内存、CPU核数等），仅作用于Spark SQL节点。多个参数用英文分号分隔。
	GeneralTaskParams []*GeneralTaskParam `json:"GeneralTaskParams,omitempty" name:"GeneralTaskParams"`
}

type ModifyWorkflowInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 责任人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 责任人id
	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作流所属用户分组id  若有多个,分号隔开: a;b;c
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitempty" name:"WorkflowParams"`

	// 用于配置优化参数（线程、内存、CPU核数等），仅作用于Spark SQL节点。多个参数用英文分号分隔。
	GeneralTaskParams []*GeneralTaskParam `json:"GeneralTaskParams,omitempty" name:"GeneralTaskParams"`
}

func (r *ModifyWorkflowInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "Owner")
	delete(f, "OwnerId")
	delete(f, "WorkflowDesc")
	delete(f, "WorkflowName")
	delete(f, "FolderId")
	delete(f, "UserGroupId")
	delete(f, "UserGroupName")
	delete(f, "WorkflowParams")
	delete(f, "GeneralTaskParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkflowInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowInfoResponseParams struct {
	// true代表成功，false代表失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWorkflowInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkflowInfoResponseParams `json:"Response"`
}

func (r *ModifyWorkflowInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowScheduleRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 延迟时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 工作流依赖 ,yes 或者no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`
}

type ModifyWorkflowScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 延迟时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 工作流依赖 ,yes 或者no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`
}

func (r *ModifyWorkflowScheduleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowScheduleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "DelayTime")
	delete(f, "StartupTime")
	delete(f, "SelfDepend")
	delete(f, "CycleType")
	delete(f, "CycleStep")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskAction")
	delete(f, "CrontabExpression")
	delete(f, "ExecutionStartTime")
	delete(f, "ExecutionEndTime")
	delete(f, "DependencyWorkflow")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyWorkflowScheduleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyWorkflowScheduleResponseParams struct {
	// 执行结果
	Data *BatchResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyWorkflowScheduleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyWorkflowScheduleResponseParams `json:"Response"`
}

func (r *ModifyWorkflowScheduleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyWorkflowScheduleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Namespace struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 当前状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type OfflineInstance struct {
	// 创建账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`

	// 操作账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// 主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// 账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitempty" name:"AppId"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitempty" name:"WorkspaceId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 下发时间
	IssueId *string `json:"IssueId,omitempty" name:"IssueId"`

	// 资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InlongTaskId *string `json:"InlongTaskId,omitempty" name:"InlongTaskId"`

	// 资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 实例类型
	TaskRunType *uint64 `json:"TaskRunType,omitempty" name:"TaskRunType"`

	// 实例状态
	State *string `json:"State,omitempty" name:"State"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 唯一key
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
}

type OfflineTaskAddParam struct {
	// 名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 依赖
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 周期
	CycleType *uint64 `json:"CycleType,omitempty" name:"CycleType"`

	// 周期间隔
	CycleStep *uint64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 延迟时间
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// crontab
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 重试等待
	RetryWait *uint64 `json:"RetryWait,omitempty" name:"RetryWait"`

	// 是否可以重试
	Retriable *uint64 `json:"Retriable,omitempty" name:"Retriable"`

	// 重试限制
	TryLimit *uint64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 优先级
	RunPriority *uint64 `json:"RunPriority,omitempty" name:"RunPriority"`

	// 产品名称
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 1 有序串行 一次一个，排队 orderly 
	// 2 无序串行 一次一个，不排队 serial  
	// 3 并行 一次多个 parallel
	SelfDepend *uint64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// 周任务：1是周天，2是周1，7是周6 。
	// 月任务：如具体1，3号则写 "1,3"，指定月末不可和具体号数一起输入，仅能为 "L"
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// 调度执行结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 调度执行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`
}

type OperateResult struct {
	// 操作结果；true表示成功；false表示失败
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 错误编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitempty" name:"ErrorId"`

	// 操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitempty" name:"ErrorDesc"`
}

type OrderField struct {
	// 排序字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 排序方向：ASC|DESC
	Direction *string `json:"Direction,omitempty" name:"Direction"`
}

type OrganizationalFunction struct {
	// 名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 层级路径
	LayerPath *string `json:"LayerPath,omitempty" name:"LayerPath"`

	// 上级层级路径
	ParentLayerPath *string `json:"ParentLayerPath,omitempty" name:"ParentLayerPath"`

	// 函数类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 函数分类：窗口函数、聚合函数、日期函数......
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 函数种类：系统函数、自定义函数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitempty" name:"Category"`

	// 函数状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 函数说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 函数用法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Usage *string `json:"Usage,omitempty" name:"Usage"`

	// 函数参数说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamDesc *string `json:"ParamDesc,omitempty" name:"ParamDesc"`

	// 函数返回值说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnDesc *string `json:"ReturnDesc,omitempty" name:"ReturnDesc"`

	// 函数示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	Example *string `json:"Example,omitempty" name:"Example"`

	// 集群实例引擎 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 函数 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FuncId *string `json:"FuncId,omitempty" name:"FuncId"`

	// 函数类名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 函数资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceList []*FunctionVersion `json:"ResourceList,omitempty" name:"ResourceList"`

	// 操作人 ID 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserIds []*int64 `json:"OperatorUserIds,omitempty" name:"OperatorUserIds"`

	// 公有云 Owner ID 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserIds []*int64 `json:"OwnerUserIds,omitempty" name:"OwnerUserIds"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitempty" name:"DbName"`

	// 提交失败错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitErrorMsg *string `json:"SubmitErrorMsg,omitempty" name:"SubmitErrorMsg"`
}

type ParamInfo struct {
	// 参数名
	ParamKey *string `json:"ParamKey,omitempty" name:"ParamKey"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitempty" name:"ParamValue"`
}

type RealTimeTaskInstanceNodeInfo struct {
	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 实时任务实例节点信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNodeInfoList []*InstanceNodeInfo `json:"InstanceNodeInfoList,omitempty" name:"InstanceNodeInfoList"`
}

type RecordField struct {
	// 字段名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type RecordsSpeed struct {
	// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`

	// 速度值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*SpeedValue `json:"Values,omitempty" name:"Values"`
}

// Predefined struct for user
type RegisterEventListenerRequestParams struct {
	// 关键字，如果是任务，则传任务Id
	Key *string `json:"Key,omitempty" name:"Key"`

	// 事件名称
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件类型，默认 REST_API
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置信息，比如最长等待时间1天配置json：{"maxWaitEventTime":1,"maxWaitEventTimeUnit":"DAYS"}
	Properties *string `json:"Properties,omitempty" name:"Properties"`
}

type RegisterEventListenerRequest struct {
	*tchttp.BaseRequest
	
	// 关键字，如果是任务，则传任务Id
	Key *string `json:"Key,omitempty" name:"Key"`

	// 事件名称
	EventName *string `json:"EventName,omitempty" name:"EventName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件类型，默认 REST_API
	Type *string `json:"Type,omitempty" name:"Type"`

	// 配置信息，比如最长等待时间1天配置json：{"maxWaitEventTime":1,"maxWaitEventTimeUnit":"DAYS"}
	Properties *string `json:"Properties,omitempty" name:"Properties"`
}

func (r *RegisterEventListenerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventListenerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Key")
	delete(f, "EventName")
	delete(f, "ProjectId")
	delete(f, "Type")
	delete(f, "Properties")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterEventListenerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterEventListenerResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchReturn `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterEventListenerResponse struct {
	*tchttp.BaseResponse
	Response *RegisterEventListenerResponseParams `json:"Response"`
}

func (r *RegisterEventListenerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventListenerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterEventRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件名称，支持英文、数字和下划线，最长20个字符, 不能以数字下划线开头。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 事件分割类型，周期类型: DAY，HOUR，MIN，SECOND
	EventSubType *string `json:"EventSubType,omitempty" name:"EventSubType"`

	// 广播：BROADCAST,单播：SINGLE
	EventBroadcastType *string `json:"EventBroadcastType,omitempty" name:"EventBroadcastType"`

	// 周期类型为天和小时为HOURS ，周期类型为分钟 ：MINUTES,周期类型为秒：SECONDS
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// TBDS 事件所属人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 事件类型，默认值：TIME_SERIES
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 对应day： yyyyMMdd，对应HOUR：yyyyMMddHH，对应MIN：yyyyMMddHHmm，对应SECOND：yyyyMMddHHmmss
	DimensionFormat *string `json:"DimensionFormat,omitempty" name:"DimensionFormat"`

	// 存活时间
	TimeToLive *int64 `json:"TimeToLive,omitempty" name:"TimeToLive"`

	// 事件描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

type RegisterEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 事件名称，支持英文、数字和下划线，最长20个字符, 不能以数字下划线开头。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 事件分割类型，周期类型: DAY，HOUR，MIN，SECOND
	EventSubType *string `json:"EventSubType,omitempty" name:"EventSubType"`

	// 广播：BROADCAST,单播：SINGLE
	EventBroadcastType *string `json:"EventBroadcastType,omitempty" name:"EventBroadcastType"`

	// 周期类型为天和小时为HOURS ，周期类型为分钟 ：MINUTES,周期类型为秒：SECONDS
	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`

	// TBDS 事件所属人
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 事件类型，默认值：TIME_SERIES
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 对应day： yyyyMMdd，对应HOUR：yyyyMMddHH，对应MIN：yyyyMMddHHmm，对应SECOND：yyyyMMddHHmmss
	DimensionFormat *string `json:"DimensionFormat,omitempty" name:"DimensionFormat"`

	// 存活时间
	TimeToLive *int64 `json:"TimeToLive,omitempty" name:"TimeToLive"`

	// 事件描述
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *RegisterEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "EventSubType")
	delete(f, "EventBroadcastType")
	delete(f, "TimeUnit")
	delete(f, "Owner")
	delete(f, "EventType")
	delete(f, "DimensionFormat")
	delete(f, "TimeToLive")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RegisterEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RegisterEventResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchReturn `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RegisterEventResponse struct {
	*tchttp.BaseResponse
	Response *RegisterEventResponseParams `json:"Response"`
}

func (r *RegisterEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RegisterEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunInstancesRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务
	CheckFather *bool `json:"CheckFather,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子
	RerunType *string `json:"RerunType,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖
	DependentWay *string `json:"DependentWay,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否
	SkipEventListening *bool `json:"SkipEventListening,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitempty" name:"SonInstanceType"`
}

type RerunInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例嵌套集合
	Instances []*InstanceInfo `json:"Instances,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务
	CheckFather *bool `json:"CheckFather,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子
	RerunType *string `json:"RerunType,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖
	DependentWay *string `json:"DependentWay,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否
	SkipEventListening *bool `json:"SkipEventListening,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitempty" name:"SonInstanceType"`
}

func (r *RerunInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RerunInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RerunInstancesResponseParams struct {
	// 返回实例批量终止结果
	Data *OperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RerunInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RerunInstancesResponseParams `json:"Response"`
}

func (r *RerunInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RerunInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcePathTree struct {
	// 资源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 是否为叶子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLeaf *bool `json:"IsLeaf,omitempty" name:"IsLeaf"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 本地路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// 远程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePath *string `json:"RemotePath,omitempty" name:"RemotePath"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensionType *string `json:"FileExtensionType,omitempty" name:"FileExtensionType"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 文件MD5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Value *string `json:"Md5Value,omitempty" name:"Md5Value"`

	// 文件拥有者名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUser *string `json:"UpdateUser,omitempty" name:"UpdateUser"`

	// 文件更新人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserId *string `json:"UpdateUserId,omitempty" name:"UpdateUserId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// Cos存储桶名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucket *string `json:"CosBucket,omitempty" name:"CosBucket"`

	// Cos地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosRegion *string `json:"CosRegion,omitempty" name:"CosRegion"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
}

// Predefined struct for user
type RestartInLongAgentRequestParams struct {
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type RestartInLongAgentRequest struct {
	*tchttp.BaseRequest
	
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *RestartInLongAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInLongAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartInLongAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartInLongAgentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RestartInLongAgentResponse struct {
	*tchttp.BaseResponse
	Response *RestartInLongAgentResponseParams `json:"Response"`
}

func (r *RestartInLongAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartInLongAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type ResumeIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *ResumeIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ResumeIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *ResumeIntegrationTaskResponseParams `json:"Response"`
}

func (r *ResumeIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RobAndLockIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type RobAndLockIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *RobAndLockIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RobAndLockIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RobAndLockIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RobAndLockIntegrationTaskResponseParams struct {
	// 抢锁状态
	RobLockState *RobLockState `json:"RobLockState,omitempty" name:"RobLockState"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RobAndLockIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *RobAndLockIntegrationTaskResponseParams `json:"Response"`
}

func (r *RobAndLockIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RobAndLockIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RobLockState struct {
	// 是否可以抢锁
	IsRob *bool `json:"IsRob,omitempty" name:"IsRob"`

	// 当前持锁人
	Locker *string `json:"Locker,omitempty" name:"Locker"`
}

// Predefined struct for user
type RunTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type RunTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RunTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunTaskResponseParams struct {
	// 运行成功或者失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type RunTaskResponse struct {
	*tchttp.BaseResponse
	Response *RunTaskResponseParams `json:"Response"`
}

func (r *RunTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveCustomFunctionRequestParams struct {
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 分类：窗口函数、聚合函数、日期函数......
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集群引擎实例
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 类名
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 资源列表
	ResourceList []*FunctionResource `json:"ResourceList,omitempty" name:"ResourceList"`

	// 函数说明
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用法
	Usage *string `json:"Usage,omitempty" name:"Usage"`

	// 参数说明
	ParamDesc *string `json:"ParamDesc,omitempty" name:"ParamDesc"`

	// 返回值说明
	ReturnDesc *string `json:"ReturnDesc,omitempty" name:"ReturnDesc"`

	// 示例
	Example *string `json:"Example,omitempty" name:"Example"`
}

type SaveCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 分类：窗口函数、聚合函数、日期函数......
	Kind *string `json:"Kind,omitempty" name:"Kind"`

	// 集群引擎实例
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 类名
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 资源列表
	ResourceList []*FunctionResource `json:"ResourceList,omitempty" name:"ResourceList"`

	// 函数说明
	Description *string `json:"Description,omitempty" name:"Description"`

	// 用法
	Usage *string `json:"Usage,omitempty" name:"Usage"`

	// 参数说明
	ParamDesc *string `json:"ParamDesc,omitempty" name:"ParamDesc"`

	// 返回值说明
	ReturnDesc *string `json:"ReturnDesc,omitempty" name:"ReturnDesc"`

	// 示例
	Example *string `json:"Example,omitempty" name:"Example"`
}

func (r *SaveCustomFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveCustomFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionId")
	delete(f, "Kind")
	delete(f, "ClusterIdentifier")
	delete(f, "ClassName")
	delete(f, "ResourceList")
	delete(f, "Description")
	delete(f, "Usage")
	delete(f, "ParamDesc")
	delete(f, "ReturnDesc")
	delete(f, "Example")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveCustomFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveCustomFunctionResponseParams struct {
	// 函数唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SaveCustomFunctionResponse struct {
	*tchttp.BaseResponse
	Response *SaveCustomFunctionResponseParams `json:"Response"`
}

func (r *SaveCustomFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveCustomFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SchedulerTaskInstanceInfo struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 实例运行时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`
}

type SchemaDetail struct {
	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnKey *string `json:"ColumnKey,omitempty" name:"ColumnKey"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type SearchCondition struct {
	// 查询框架，必选
	Instance *SearchConditionInstance `json:"Instance,omitempty" name:"Instance"`

	// 查询关键字（任务Id精确匹配，任务名称模糊匹配），可选
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitempty" name:"SortCol"`
}

type SearchConditionInstance struct {
	// 执行空间 "DRY_RUN"
	ExecutionSpace *uint64 `json:"ExecutionSpace,omitempty" name:"ExecutionSpace"`

	// 产品名称，可选
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *uint64 `json:"ProductName,omitempty" name:"ProductName"`

	// 资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *uint64 `json:"ResourceGroup,omitempty" name:"ResourceGroup"`
}

type SearchConditionInstanceNew struct {
	// 执行空间 "DRY_RUN"
	ExecutionSpace *string `json:"ExecutionSpace,omitempty" name:"ExecutionSpace"`

	// 产品名称，可选
	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`

	// 资源组
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`
}

type SearchConditionNew struct {
	// 查询框架，必选
	Instance *SearchConditionInstanceNew `json:"Instance,omitempty" name:"Instance"`

	// 查询关键字（任务Id精确匹配，任务名称模糊匹配），可选
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitempty" name:"SortCol"`
}

// Predefined struct for user
type SetTaskAlarmNewRequestParams struct {
	// 设置任务超时告警和失败告警信息
	AlarmInfoList []*AlarmInfo `json:"AlarmInfoList,omitempty" name:"AlarmInfoList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type SetTaskAlarmNewRequest struct {
	*tchttp.BaseRequest
	
	// 设置任务超时告警和失败告警信息
	AlarmInfoList []*AlarmInfo `json:"AlarmInfoList,omitempty" name:"AlarmInfoList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *SetTaskAlarmNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTaskAlarmNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmInfoList")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetTaskAlarmNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetTaskAlarmNewResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperateResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetTaskAlarmNewResponse struct {
	*tchttp.BaseResponse
	Response *SetTaskAlarmNewResponseParams `json:"Response"`
}

func (r *SetTaskAlarmNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetTaskAlarmNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SimpleTaskInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type SpeedValue struct {
	// 带毫秒的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *uint64 `json:"Time,omitempty" name:"Time"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Speed *float64 `json:"Speed,omitempty" name:"Speed"`
}

// Predefined struct for user
type StartIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type StartIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *StartIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StartIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *StartIntegrationTaskResponseParams `json:"Response"`
}

func (r *StartIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type StopIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *StopIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type StopIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *StopIntegrationTaskResponseParams `json:"Response"`
}

func (r *StopIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitCustomFunctionRequestParams struct {
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 备注信息
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type SubmitCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitempty" name:"ClusterIdentifier"`

	// 备注信息
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *SubmitCustomFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCustomFunctionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "FunctionId")
	delete(f, "ClusterIdentifier")
	delete(f, "Comment")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitCustomFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitCustomFunctionResponseParams struct {
	// 函数唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionId *string `json:"FunctionId,omitempty" name:"FunctionId"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubmitCustomFunctionResponse struct {
	*tchttp.BaseResponse
	Response *SubmitCustomFunctionResponseParams `json:"Response"`
}

func (r *SubmitCustomFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitCustomFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

type SubmitTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

func (r *SubmitTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "VersionRemark")
	delete(f, "StartScheduling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskResponseParams struct {
	// 成功或者失败
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubmitTaskResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTaskResponseParams `json:"Response"`
}

func (r *SubmitTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitWorkflow struct {
	// 被提交的任务id集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`

	// 执行结果
	Result *bool `json:"Result,omitempty" name:"Result"`

	// 执行情况备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitempty" name:"ErrorDesc"`

	// 执行情况id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitempty" name:"ErrorId"`
}

// Predefined struct for user
type SubmitWorkflowRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 提交的版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

type SubmitWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 提交的版本备注
	VersionRemark *string `json:"VersionRemark,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitempty" name:"StartScheduling"`
}

func (r *SubmitWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "VersionRemark")
	delete(f, "StartScheduling")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitWorkflowResponseParams struct {
	// 执行结果
	Data *SubmitWorkflow `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SubmitWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *SubmitWorkflowResponseParams `json:"Response"`
}

func (r *SubmitWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SuspendIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type SuspendIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *SuspendIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SuspendIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SuspendIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SuspendIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *SuspendIntegrationTaskResponseParams `json:"Response"`
}

func (r *SuspendIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SuspendIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TableInfo struct {
	// 表Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitempty" name:"TableId"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitempty" name:"TableName"`
}

type TaskAlarmInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 规则名称
	RegularName *string `json:"RegularName,omitempty" name:"RegularName"`

	// 规则状态(0表示关闭，1表示打开)
	RegularStatus *uint64 `json:"RegularStatus,omitempty" name:"RegularStatus"`

	// 告警级别(0表示普通，1表示重要，2表示紧急)
	AlarmLevel *uint64 `json:"AlarmLevel,omitempty" name:"AlarmLevel"`

	// 告警指标,0表示任务失败，1表示任务运行超时，2表示任务停止，3表示任务暂停
	// ，4写入速度，5读取速度，6读取吞吐，7写入吞吐, 8脏数据字节数，9脏数据条数
	AlarmIndicator *uint64 `json:"AlarmIndicator,omitempty" name:"AlarmIndicator"`

	// 告警方式,多个用逗号隔开（1:邮件，2:短信，3:微信，4:语音，5:代表企业微信，6:http）
	AlarmWay *string `json:"AlarmWay,omitempty" name:"AlarmWay"`

	// 告警接收人ID，多个用逗号隔开
	AlarmRecipientId *string `json:"AlarmRecipientId,omitempty" name:"AlarmRecipientId"`

	// 任务类型(201表示实时，202表示离线)
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`

	// 告警接收人昵称，多个用逗号隔开
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmRecipientName *string `json:"AlarmRecipientName,omitempty" name:"AlarmRecipientName"`

	// 主键ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 规则ID
	RegularId *string `json:"RegularId,omitempty" name:"RegularId"`

	// 指标阈值(1表示离线任务第一次运行失败，2表示离线任务所有重试完成后失败)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *uint64 `json:"TriggerType,omitempty" name:"TriggerType"`

	// 预计的超时时间(分钟级别)
	// 注意：此字段可能返回 null，表示取不到有效值。
	EstimatedTime *uint64 `json:"EstimatedTime,omitempty" name:"EstimatedTime"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// 告警指标描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicatorDesc *string `json:"AlarmIndicatorDesc,omitempty" name:"AlarmIndicatorDesc"`

	// 实时任务告警需要的参数，1是大于2是小于
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *uint64 `json:"Operator,omitempty" name:"Operator"`

	// 节点id，多个逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`

	// 节点名称，多个逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type TaskCanvasInfo struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitempty" name:"ProjectIdent"`

	// 任务状态，'Y','F','O','T','INVALID' 分别表示调度中、已停止、已暂停、停止中、已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 任务类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitempty" name:"TaskTypeId"`

	// 任务类型描述，其中任务类型id和任务类型描述的对应的关系为
	// 20	通用数据同步任务
	// 21	JDBC SQL
	// 22	Tbase
	// 25	数据ETL
	// 30	Python
	// 31	PySpark
	// 34	Hive SQL
	// 35	Shell
	// 36	Spark SQL
	// 37	HDFS到HBase
	// 38	SHELL
	// 39	Spark
	// 45	DATA_QUALITY
	// 55	THIVE到MYSQL
	// 56	THIVE到PG
	// 66	HDFS到PG
	// 67	HDFS到Oracle
	// 68	HDFS到MYSQL
	// 69	FTP到HDFS
	// 70	HIVE SQL
	// 72	HIVE到HDFS
	// 75	HDFS到HIVE
	// 81	PYTHONSQL脚本
	// 82	SPARKSCALA计算
	// 83	虫洞任务
	// 84	校验对账文件
	// 85	HDFS到THIVE
	// 86	TDW到HDFS
	// 87	HDFS到TDW
	// 88	校验对账文件
	// 91	FLINK任务
	// 92	MapReduce
	// 98	custom topology
	// 99	kafkatoHDFS
	// 100	kafkatoHbase
	// 101	MYSQL导入至HIVE(DX)
	// 104	MYSQL到HIVE
	// 105	HIVE到MYSQL
	// 106	SQL SERVER到HIVE
	// 107	HIVE到SQL SERVER
	// 108	ORACLE到HIVE
	// 109	HIVE到ORACLE
	// 111	HIVE到MYSQL(NEW)
	// 112	HIVE到PG
	// 113	HIVE到PHOENIX
	// 118	MYSQL到HDFS
	// 119	PG到HDFS
	// 120	ORACLE到HDFS
	// 121	数据质量
	// 10000	自定义业务
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitempty" name:"TaskTypeDesc"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitempty" name:"FirstSubmitTime"`

	// 首次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstRunTime *string `json:"FirstRunTime,omitempty" name:"FirstRunTime"`

	// 调度计划展示描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitempty" name:"ScheduleDesc"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 调度周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitempty" name:"CycleUnit"`

	// 画布x轴坐标点
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitempty" name:"LeftCoordinate"`

	// 画布y轴坐标点
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitempty" name:"TopCoordinate"`

	// 跨工作流虚拟任务标识；true标识跨工作流任务；false标识本工作流任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`

	// 弹性周期配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// 延迟时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`
}

type TaskExtInfo struct {
	// 键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TaskInfoData struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitempty" name:"ProjectIdent"`

	// 任务状态，'Y','F','O','T','INVALID' 分别表示调度中、已停止、已暂停、停止中、已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 跨工作流虚拟任务标识；true标识跨工作流任务；false标识本工作流任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`

	// 延时实例生成时间(延时调度)，转换为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitempty" name:"DelayTime"`

	// crontab表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitempty" name:"CrontabExpression"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitempty" name:"LastUpdate"`

	// 生效日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 执行时间左闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitempty" name:"ExecutionEndTime"`

	// 周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 步长
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *int64 `json:"CycleStep,omitempty" name:"CycleStep"`

	// 延时执行时间（延时执行) 对应为 开始时间 状态为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupTime *int64 `json:"StartupTime,omitempty" name:"StartupTime"`

	// 重试等待时间,单位分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryWait *int64 `json:"RetryWait,omitempty" name:"RetryWait"`

	// 是否可重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retriable *int64 `json:"Retriable,omitempty" name:"Retriable"`

	// 调度扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitempty" name:"TaskAction"`

	// 运行次数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 运行优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunPriority *int64 `json:"RunPriority,omitempty" name:"RunPriority"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`

	// 指定的运行节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 集群
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 最小数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinDateTime *string `json:"MinDateTime,omitempty" name:"MinDateTime"`

	// 最大数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDateTime *string `json:"MaxDateTime,omitempty" name:"MaxDateTime"`

	// 是否自身依赖 是1 否2 并行3
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDepend *int64 `json:"SelfDepend,omitempty" name:"SelfDepend"`

	// 扩展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskExt []*TaskExtInfo `json:"TaskExt,omitempty" name:"TaskExt"`

	// 任务备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notes *string `json:"Notes,omitempty" name:"Notes"`

	// 队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	YarnQueue *string `json:"YarnQueue,omitempty" name:"YarnQueue"`

	// 任务版本是否已提交
	// 注意：此字段可能返回 null，表示取不到有效值。
	Submit *bool `json:"Submit,omitempty" name:"Submit"`

	// 最新调度计划变更时间 仅生产态
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSchedulerCommitTime *string `json:"LastSchedulerCommitTime,omitempty" name:"LastSchedulerCommitTime"`

	// 仅生产态存储于生产态序列化任务信息, 减少base CPU重复密集计算
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalizedJobStartTime *string `json:"NormalizedJobStartTime,omitempty" name:"NormalizedJobStartTime"`

	// 源数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServer *string `json:"SourceServer,omitempty" name:"SourceServer"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creater *string `json:"Creater,omitempty" name:"Creater"`

	// 分支，依赖关系，and/or, 默认and
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyRel *string `json:"DependencyRel,omitempty" name:"DependencyRel"`

	// 是否支持工作流依赖 yes / no 默认 no
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyWorkflow *string `json:"DependencyWorkflow,omitempty" name:"DependencyWorkflow"`

	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*ParamInfo `json:"Params,omitempty" name:"Params"`

	// 最后修改的人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUser *string `json:"UpdateUser,omitempty" name:"UpdateUser"`

	// 最后修改的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 最后修改的人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserId *string `json:"UpdateUserId,omitempty" name:"UpdateUserId"`

	// 调度计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitempty" name:"SchedulerDesc"`

	// 资源组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *string `json:"ResourceGroup,omitempty" name:"ResourceGroup"`

	// 版本提交说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionDesc *string `json:"VersionDesc,omitempty" name:"VersionDesc"`

	// 真实工作流Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealWorkflowId *string `json:"RealWorkflowId,omitempty" name:"RealWorkflowId"`

	// 目标数据源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServer *string `json:"TargetServer,omitempty" name:"TargetServer"`

	// 依赖配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyConfigs []*DependencyConfig `json:"DependencyConfigs,omitempty" name:"DependencyConfigs"`

	// 虚拟任务状态1
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskStatus *string `json:"VirtualTaskStatus,omitempty" name:"VirtualTaskStatus"`

	// 虚拟任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskId *string `json:"VirtualTaskId,omitempty" name:"VirtualTaskId"`
}

type TaskInfoDataPage struct {
	// 页号
	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`

	// 任务集合信息
	Items []*TaskInfoData `json:"Items,omitempty" name:"Items"`

	// 总页数1
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TaskInnerInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitempty" name:"CycleType"`

	// 虚拟任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskId *string `json:"VirtualTaskId,omitempty" name:"VirtualTaskId"`

	// 虚拟任务标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitempty" name:"VirtualFlag"`

	// 真实任务工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealWorkflowId *string `json:"RealWorkflowId,omitempty" name:"RealWorkflowId"`
}

type TaskInstanceDetail struct {
	// 实例id
	TaskRunId *string `json:"TaskRunId,omitempty" name:"TaskRunId"`

	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 实例数据运行时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 实例实际运行时间
	IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`

	// InLong任务Id
	InlongTaskId *string `json:"InlongTaskId,omitempty" name:"InlongTaskId"`

	// 执行资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`

	// 任务类型(1 调试运行,2 调度执行)
	TaskRunType *uint64 `json:"TaskRunType,omitempty" name:"TaskRunType"`

	// 任务状态(1 正在执行,2 成功,3 失败,4 等待终止,5 正在终止,6 已终止,7 终止失败,9 等待执行)
	State *uint64 `json:"State,omitempty" name:"State"`

	// 实例开始运行时间，格式：yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 实例结束运行时间，格式：yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// Broker IP
	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`

	// 运行实例的EKS Pod名称
	PodName *string `json:"PodName,omitempty" name:"PodName"`

	// 下一个调度周期的数据运行时间
	NextRunDate *string `json:"NextRunDate,omitempty" name:"NextRunDate"`

	// 创建者的账号Id
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// 操作者的账号Id
	OperatorUin *uint64 `json:"OperatorUin,omitempty" name:"OperatorUin"`

	// 拥有者的账号Id
	OwnerUin *uint64 `json:"OwnerUin,omitempty" name:"OwnerUin"`

	// App Id
	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`

	// WeData项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
}

type TaskInstanceInfo struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitempty" name:"ProjectIdent"`

	// 实例状态，0等待事件，1等待上游，2等待运行，3运行中，4正在终止，5失败重试，6失败，7成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *int64 `json:"State,omitempty" name:"State"`

	// 任务类型id，26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40CDW PG，92MapReduce
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitempty" name:"TaskTypeId"`

	// 任务类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitempty" name:"TaskTypeDesc"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 调度计划展示描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitempty" name:"SchedulerDesc"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitempty" name:"InCharge"`

	// 调度周期类型，I分钟，H小时，D天，W周，M月，Y年，O一次性，C crontab
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitempty" name:"CycleType"`

	// 实例开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 实例结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 实例类型，0补录实例，1周期实例，2非周期实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *int64 `json:"InstanceType,omitempty" name:"InstanceType"`

	// 最大重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *int64 `json:"TryLimit,omitempty" name:"TryLimit"`

	// 当前重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *int64 `json:"Tries,omitempty" name:"Tries"`

	// 计划调度时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDateTime *string `json:"SchedulerDateTime,omitempty" name:"SchedulerDateTime"`

	// 运行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *string `json:"CostTime,omitempty" name:"CostTime"`
}

type TaskLinkInfo struct {
	// 下游任务id
	TaskTo *string `json:"TaskTo,omitempty" name:"TaskTo"`

	// 上游任务id
	TaskFrom *string `json:"TaskFrom,omitempty" name:"TaskFrom"`

	// 依赖边类型 1、“real_real”表示任务->任务；2、"virtual_real" 跨工作流任务->任务
	LinkType *string `json:"LinkType,omitempty" name:"LinkType"`

	// 依赖边id
	LinkId *string `json:"LinkId,omitempty" name:"LinkId"`
}

type TaskLockStatus struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 持锁者
	Locker *string `json:"Locker,omitempty" name:"Locker"`

	// 当前操作用户是否为持锁者，1表示为持锁者，0表示为不为持锁者
	IsLocker *int64 `json:"IsLocker,omitempty" name:"IsLocker"`

	// 是否可以抢锁，1表示可以抢锁，0表示不可以抢锁
	IsRob *int64 `json:"IsRob,omitempty" name:"IsRob"`
}

// Predefined struct for user
type TaskLogRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 起始时间戳，单位毫秒
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳，单位毫秒
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 拉取日志数量，默认100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 日志排序 desc 倒序 asc 顺序
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

type TaskLogRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 起始时间戳，单位毫秒
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间戳，单位毫秒
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 拉取日志数量，默认100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 日志排序 desc 倒序 asc 顺序
	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
}

func (r *TaskLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TaskLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ProjectId")
	delete(f, "Limit")
	delete(f, "OrderType")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TaskLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TaskLogResponseParams struct {
	// 详细日志
	LogContentList []*LogContent `json:"LogContentList,omitempty" name:"LogContentList"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TaskLogResponse struct {
	*tchttp.BaseResponse
	Response *TaskLogResponseParams `json:"Response"`
}

func (r *TaskLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TaskLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskReportDetail struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例数据运行时间
	CurRunDate *string `json:"CurRunDate,omitempty" name:"CurRunDate"`

	// 实例实际下发时间
	IssueDate *string `json:"IssueDate,omitempty" name:"IssueDate"`

	// 任务状态码。1 正在执行,2 成功,3 失败,4 等待终止,5 正在终止,6 已终止,7 终止失败,9 等待执行。
	TaskState *string `json:"TaskState,omitempty" name:"TaskState"`

	// 总读取条数
	TotalReadRecords *uint64 `json:"TotalReadRecords,omitempty" name:"TotalReadRecords"`

	// 总读取字节数
	TotalReadBytes *uint64 `json:"TotalReadBytes,omitempty" name:"TotalReadBytes"`

	// 总写入条数
	TotalWriteRecords *uint64 `json:"TotalWriteRecords,omitempty" name:"TotalWriteRecords"`

	// 总写入字节数
	TotalWriteBytes *uint64 `json:"TotalWriteBytes,omitempty" name:"TotalWriteBytes"`

	// 写入速度（条/秒）
	RecordSpeed *uint64 `json:"RecordSpeed,omitempty" name:"RecordSpeed"`

	// 吞吐（Byte/秒）
	ByteSpeed *float64 `json:"ByteSpeed,omitempty" name:"ByteSpeed"`

	// 脏数据条数
	TotalErrorRecords *uint64 `json:"TotalErrorRecords,omitempty" name:"TotalErrorRecords"`
}

type TaskScriptContent struct {
	// 脚本内容 base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitempty" name:"ScriptContent"`
}

// Predefined struct for user
type TriggerEventRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 案例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间格式：如果选择触发时间：2022年6月21，则设置为20220621
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

type TriggerEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 案例名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 时间格式：如果选择触发时间：2022年6月21，则设置为20220621
	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`

	// 描述信息
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *TriggerEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Name")
	delete(f, "Dimension")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerEventResponseParams struct {
	// 成功或者失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchReturn `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type TriggerEventResponse struct {
	*tchttp.BaseResponse
	Response *TriggerEventResponseParams `json:"Response"`
}

func (r *TriggerEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

type UnlockIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *UnlockIntegrationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockIntegrationTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnlockIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnlockIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UnlockIntegrationTaskResponse struct {
	*tchttp.BaseResponse
	Response *UnlockIntegrationTaskResponseParams `json:"Response"`
}

func (r *UnlockIntegrationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnlockIntegrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateInLongAgentRequestParams struct {
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 采集器名称
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 集成资源组ID
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

type UpdateInLongAgentRequest struct {
	*tchttp.BaseRequest
	
	// 采集器ID
	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`

	// WeData项目ID
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 采集器名称
	AgentName *string `json:"AgentName,omitempty" name:"AgentName"`

	// 集成资源组ID
	ExecutorGroupId *string `json:"ExecutorGroupId,omitempty" name:"ExecutorGroupId"`
}

func (r *UpdateInLongAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateInLongAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "ProjectId")
	delete(f, "AgentName")
	delete(f, "ExecutorGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateInLongAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateInLongAgentResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateInLongAgentResponse struct {
	*tchttp.BaseResponse
	Response *UpdateInLongAgentResponseParams `json:"Response"`
}

func (r *UpdateInLongAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateInLongAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserFileDTO struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 文件类型，如 jar zip 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensionType *string `json:"FileExtensionType,omitempty" name:"FileExtensionType"`

	// 文件上传类型，资源管理为 resource
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUploadType *string `json:"FileUploadType,omitempty" name:"FileUploadType"`

	// 文件MD5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Value *string `json:"Md5Value,omitempty" name:"Md5Value"`

	// 创建时间，秒级别的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 更新时间，秒级别的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 文件大小，单位为字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitempty" name:"Size"`

	// 本地路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// 本地临时路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalTmpPath *string `json:"LocalTmpPath,omitempty" name:"LocalTmpPath"`

	// 远程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePath *string `json:"RemotePath,omitempty" name:"RemotePath"`

	// 文件拥有者名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`

	// 文件拥有者uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 文件深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathDepth *string `json:"PathDepth,omitempty" name:"PathDepth"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`

	// 本地临时压缩文件绝对路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZipPath *string `json:"ZipPath,omitempty" name:"ZipPath"`

	// 文件所属存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 文件所属存储桶的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitempty" name:"Region"`
}

type Workflow struct {
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitempty" name:"WorkflowId"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitempty" name:"Owner"`

	// 责任人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerId *string `json:"OwnerId,omitempty" name:"OwnerId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`

	// 项目标识
	ProjectIdent *string `json:"ProjectIdent,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`

	// 工作流所属用户分组id 若有多个,分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupId *string `json:"UserGroupId,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupName *string `json:"UserGroupName,omitempty" name:"UserGroupName"`
}