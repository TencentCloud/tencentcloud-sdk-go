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

package v20201016

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AlarmInfo struct {

	// 告警策略名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监控对象列表。
	AlarmTargets []*AlarmTargetInfo `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// 触发条件。
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是min。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知模板列表。
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// 开启状态。
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近更新时间。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 自定义通知模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`

	// 自定义回调模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`
}

type AlarmNotice struct {

	// 告警通知模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警模板的类型。可选值：
	// <br><li> Trigger - 告警触发
	// <br><li> Recovery - 告警恢复
	// <br><li> All - 告警触发和告警恢复
	Type *string `json:"Type,omitempty" name:"Type"`

	// 告警通知模板接收者信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// 告警通知模板回调信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`

	// 告警通知模板ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近更新时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AlarmTarget struct {

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 查询语句。
	Query *string `json:"Query,omitempty" name:"Query"`

	// 告警对象序号；从1开始递增。
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 查询范围起始时间相对当前的历史时间，单位非分钟，取值为非正，最大值为0，最小值为-1440。
	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 查询范围终止时间相对当前的历史时间，单位非分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。
	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 日志集ID。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

type AlarmTargetInfo struct {

	// 日志集ID。
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志集名称。
	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`

	// 日志主题ID。
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 查询语句。
	Query *string `json:"Query,omitempty" name:"Query"`

	// 告警对象序号。
	Number *int64 `json:"Number,omitempty" name:"Number"`

	// 查询范围起始时间相对当前的历史时间，取值为非正，最大值为0，最小值为-1440。
	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 查询范围终止时间相对当前的历史时间，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。
	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type CallBackInfo struct {

	// 回调时的Body
	Body *string `json:"Body,omitempty" name:"Body"`

	// 回调时的Headers
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*string `json:"Headers,omitempty" name:"Headers"`
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 告警模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警模板的类型。可选值：
	// <br><li> Trigger - 告警触发
	// <br><li> Recovery - 告警恢复
	// <br><li> All - 告警触发和告警恢复
	Type *string `json:"Type,omitempty" name:"Type"`

	// 告警模板接收者信息。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// 告警模板回调信息。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警模板ID
		AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// 触发条件。
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知模板列表。
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// 是否开启告警策略。默认值为true
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *CreateAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AlarmTargets")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略ID。
		AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportRequest struct {
	*tchttp.BaseRequest

	// 日志主题
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志导出检索语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 日志导出数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 日志导出起始时间，毫秒时间戳
	From *int64 `json:"From,omitempty" name:"From"`

	// 日志导出结束时间，毫秒时间戳
	To *int64 `json:"To,omitempty" name:"To"`

	// 日志导出时间排序。desc，asc，默认为desc
	Order *string `json:"Order,omitempty" name:"Order"`

	// 日志导出数据格式。json，csv，默认为json
	Format *string `json:"Format,omitempty" name:"Format"`
}

func (r *CreateExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Query")
	delete(f, "Count")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Order")
	delete(f, "Format")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志导出ID。
		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 索引规则
	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`

	// 是否生效，默认为true
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *CreateIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Rule")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组名字，不能重复
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 创建机器组类型，Type为ip，Values中为Ip字符串列表创建机器组，Type为label， Values中为标签字符串列表创建机器组
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的机器组。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

func (r *CreateMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器组ID
		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志主题分区个数。默认创建1个，最大支持创建10个分区。
	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启自动分裂，默认值为false
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// 开启自动分裂后，每个主题能够允许的最大分区数，默认值为50
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// 日志主题的存储类型，可选值 hot（热存储），cold（冷存储）默认为hot
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LogsetId")
	delete(f, "TopicName")
	delete(f, "PartitionCount")
	delete(f, "Tags")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	delete(f, "StorageType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志主题ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 告警通知模板
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
}

func (r *DeleteAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExportId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteExportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// <br><li> name
	// 
	// 按照【告警通知模板名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> alarmNoticeId
	// 
	// 按照【告警通知模板ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> uid
	// 
	// 按照【接收用户ID】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> groupId
	// 
	// 按照【用户组ID】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNoticesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警通知模板列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		AlarmNotices []*AlarmNotice `json:"AlarmNotices,omitempty" name:"AlarmNotices"`

		// 符合条件的告警通知模板总数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest

	// <br><li> name
	// 
	// 按照【告警策略名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> alarmId
	// 
	// 按照【告警策略ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> topicId
	// 
	// 按照【监控对象的日志主题ID】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> enable
	// 
	// 按照【启用状态】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警策略列表。
		Alarms []*AlarmInfo `json:"Alarms,omitempty" name:"Alarms"`

		// 符合查询条件的告警策略数目。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExportsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExportsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志导出列表
		Exports []*ExportInfo `json:"Exports,omitempty" name:"Exports"`

		// 总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志主题ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// 是否生效
		Status *bool `json:"Status,omitempty" name:"Status"`

		// 索引配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`

		// 索引修改时间，初始值为索引创建时间。
		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志时间,  格式: YYYY-mm-dd HH:MM:SS
	BTime *string `json:"BTime,omitempty" name:"BTime"`

	// 日志包序号
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 日志包内一条日志的序号
	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 上文日志条数,  默认值10
	PrevLogs *int64 `json:"PrevLogs,omitempty" name:"PrevLogs"`

	// 下文日志条数,  默认值10
	NextLogs *int64 `json:"NextLogs,omitempty" name:"NextLogs"`
}

func (r *DescribeLogContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "BTime")
	delete(f, "PkgId")
	delete(f, "PkgLogId")
	delete(f, "PrevLogs")
	delete(f, "NextLogs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLogContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志上下文信息集合
		LogContextInfos []*LogContextInfo `json:"LogContextInfos,omitempty" name:"LogContextInfos"`

		// 上文日志是否已经返回
		PrevOver *bool `json:"PrevOver,omitempty" name:"PrevOver"`

		// 下文日志是否已经返回
		NextOver *bool `json:"NextOver,omitempty" name:"NextOver"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLogContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// <br><li> machineGroupName
	// 
	// 按照【机器组名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> machineGroupId
	// 
	// 按照【机器组ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tagKey
	// 
	// 按照【标签键】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tag:tagKey
	// 
	// 按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。
	// 类型：String
	// 
	// 必选：否
	// 
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页的限制数目，默认值为20，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMachineGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachineGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器组信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`

		// 分页的总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 查询的机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMachinesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 机器状态信息组
		Machines []*MachineInfo `json:"Machines,omitempty" name:"Machines"`

		// 机器组是否开启自动升级功能
		AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

		// 机器组自动升级功能预设开始时间
		UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

		// 机器组自动升级功能预设结束时间
		UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

		// 当前用户可用最新的Loglistener版本
		LatestAgentVersion *string `json:"LatestAgentVersion,omitempty" name:"LatestAgentVersion"`

		// 是否开启服务日志
		ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribePartitionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分区列表
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePartitionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// <br><li> topicName
	// 
	// 按照【日志主题名称】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> topicId
	// 
	// 按照【日志主题ID】进行过滤。
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> logsetId
	// 
	// 按照【日志集ID】进行过滤，可通过调用DescribeLogsets查询已创建的日志集列表或登录控制台进行查看；也可以调用CreateLogset创建新的日志集。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tagKey
	// 
	// 按照【标签键】进行过滤。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// <br><li> tag:tagKey
	// 
	// 按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。
	// 
	// 类型：String
	// 
	// 必选：否
	// 
	// 
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页的偏移量，默认值为0。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 分页单页限制数目，默认值为20，最大值100。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志主题列表
		Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics"`

		// 总数目
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportInfo struct {

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志导出任务ID
	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`

	// 日志导出查询语句
	Query *string `json:"Query,omitempty" name:"Query"`

	// 日志导出文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 日志文件大小
	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`

	// 日志导出时间排序
	Order *string `json:"Order,omitempty" name:"Order"`

	// 日志导出格式
	Format *string `json:"Format,omitempty" name:"Format"`

	// 日志导出数量
	Count *uint64 `json:"Count,omitempty" name:"Count"`

	// 日志下载状态。Processing:导出正在进行中，Complete:导出完成，Failed:导出失败，Expired:日志导出已过期（三天有效期）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 日志导出起始时间
	From *int64 `json:"From,omitempty" name:"From"`

	// 日志导出结束时间
	To *int64 `json:"To,omitempty" name:"To"`

	// 日志导出路径
	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`

	// 日志导出创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type Filter struct {

	// 需要过滤的字段。
	Key *string `json:"Key,omitempty" name:"Key"`

	// 需要过滤的值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type FullTextInfo struct {

	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 全文索引的分词符，字符串中每个字符代表一个分词符
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type GetAlarmLogRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为1024
	Query *string `json:"Query,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *GetAlarmLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAlarmLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type GetAlarmLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加载后续内容的Context
		Context *string `json:"Context,omitempty" name:"Context"`

		// 日志查询结果是否全部返回
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// 返回的是否为分析结果
		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`

		// 如果Analysis为True，则返回分析结果的列名，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`

		// 日志查询结果；当Analysis为True时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// 日志分析结果；当Analysis为False时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlarmLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAlarmLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValueInfo struct {

	// 需要配置键值或者元字段索引的字段
	Key *string `json:"Key,omitempty" name:"Key"`

	// 字段的索引描述信息
	Value *ValueInfo `json:"Value,omitempty" name:"Value"`
}

type LogContextInfo struct {

	// 日志来源设备
	Source *string `json:"Source,omitempty" name:"Source"`

	// 采集路径
	Filename *string `json:"Filename,omitempty" name:"Filename"`

	// 日志内容
	Content *string `json:"Content,omitempty" name:"Content"`

	// 日志包序号
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 日志包内一条日志的序号
	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 日志时间戳
	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
}

type LogInfo struct {

	// 日志时间，单位ms
	Time *int64 `json:"Time,omitempty" name:"Time"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 日志来源IP
	Source *string `json:"Source,omitempty" name:"Source"`

	// 日志文件名称
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 日志上报请求包的ID
	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`

	// 请求包内日志的ID
	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`

	// 日志内容的Json序列化字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
}

type LogItem struct {

	// 日志Key
	Key *string `json:"Key,omitempty" name:"Key"`

	// 日志Value
	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogItems struct {

	// 分析结果返回的KV数据对
	Data []*LogItem `json:"Data,omitempty" name:"Data"`
}

type MachineGroupInfo struct {

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 机器组类型
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 机器组绑定的标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

type MachineGroupTypeInfo struct {

	// 机器组类型，ip表示该机器组Values中存的是采集机器的IP地址，label表示该机器组Values中存储的是机器的标签
	Type *string `json:"Type,omitempty" name:"Type"`

	// 机器描述列表
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type MachineInfo struct {

	// 机器的IP
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// 机器状态，0:异常，1:正常
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 机器离线时间，空为正常，异常返回具体时间
	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`

	// 机器是否开启自动升级。0:关闭，1:开启
	AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 机器当前版本号。
	Version *string `json:"Version,omitempty" name:"Version"`

	// 机器升级功能状态。
	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`

	// 机器升级结果标识。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 机器升级结果信息。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type MergePartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 合并的PartitionId
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

func (r *MergePartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MergePartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type MergePartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 合并结果集
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MergePartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MergePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 告警通知模板ID。
	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`

	// 告警模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 告警模板的类型。可选值：
	// <br><li> Trigger - 告警触发
	// <br><li> Recovery - 告警恢复
	// <br><li> All - 告警触发和告警恢复
	Type *string `json:"Type,omitempty" name:"Type"`

	// 告警模板接收者信息。
	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`

	// 告警模板回调信息。
	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmNoticeId")
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "NoticeReceivers")
	delete(f, "WebCallbacks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmNoticeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。
	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`

	// 告警策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 监控任务运行时间点。
	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`

	// 触发条件。
	Condition *string `json:"Condition,omitempty" name:"Condition"`

	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。
	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`

	// 告警重复的周期。单位是分钟。取值范围是0~1440。
	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`

	// 关联的告警通知模板列表。
	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`

	// 监控对象列表。
	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`

	// 是否开启告警策略。
	Status *bool `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAlarmRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AlarmId")
	delete(f, "Name")
	delete(f, "MonitorTime")
	delete(f, "Condition")
	delete(f, "TriggerCount")
	delete(f, "AlarmPeriod")
	delete(f, "AlarmNoticeIds")
	delete(f, "AlarmTargets")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAlarmRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 默认不生效
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 索引规则，Rule和Effective两个必须有一个参数存在
	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
}

func (r *ModifyIndexRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "Status")
	delete(f, "Rule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyIndexRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIndexResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// 机器组名称
	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`

	// 机器组类型
	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`

	// 标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 是否开启机器组自动更新
	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`

	// 升级开始时间，建议业务低峰期升级LogListener
	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`

	// 升级结束时间，建议业务低峰期升级LogListener
	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`

	// 是否开启服务日志，用于记录因Loglistener 服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费
	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
}

func (r *ModifyMachineGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "GroupName")
	delete(f, "MachineGroupType")
	delete(f, "Tags")
	delete(f, "AutoUpdate")
	delete(f, "UpdateStartTime")
	delete(f, "UpdateEndTime")
	delete(f, "ServiceLogging")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMachineGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachineGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，并且不能有重复的键值对。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 该日志主题是否开始采集
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 是否开启自动分裂
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// 若开启最大分裂，该主题能够能够允许的最大分区数
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "TopicName")
	delete(f, "Tags")
	delete(f, "Status")
	delete(f, "AutoSplit")
	delete(f, "MaxSplitPartitions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorTime struct {

	// 可选值：
	// <br><li> Period - 周期执行
	// <br><li> Fixed - 定期执行
	Type *string `json:"Type,omitempty" name:"Type"`

	// 执行的周期，或者定制执行的时间节点。单位为分钟，取值范围为1~1440。
	Time *int64 `json:"Time,omitempty" name:"Time"`
}

type NoticeReceiver struct {

	// 接受者类型。可选值：
	// <br><li> Uin - 用户ID
	// <br><li> Group - 用户组ID
	// 暂不支持其余接收者类型。
	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`

	// 接收者。
	ReceiverIds []*int64 `json:"ReceiverIds,omitempty" name:"ReceiverIds"`

	// 通知接收渠道。
	// <br><li> Email - 邮件
	// <br><li> Sms - 短信
	// <br><li> WeChat - 微信
	// <br><li> Phone - 电话
	ReceiverChannels []*string `json:"ReceiverChannels,omitempty" name:"ReceiverChannels"`

	// 允许接收信息的开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 允许接收信息的结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 位序
	Index *int64 `json:"Index,omitempty" name:"Index"`
}

type PartitionInfo struct {

	// 分区ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`

	// 分区的状态（readwrite或者是readonly）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分区哈希键起始key
	InclusiveBeginKey *string `json:"InclusiveBeginKey,omitempty" name:"InclusiveBeginKey"`

	// 分区哈希键结束key
	ExclusiveEndKey *string `json:"ExclusiveEndKey,omitempty" name:"ExclusiveEndKey"`

	// 分区创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 只读分区数据停止写入时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastWriteTime *string `json:"LastWriteTime,omitempty" name:"LastWriteTime"`
}

type RuleInfo struct {

	// 全文索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullText *FullTextInfo `json:"FullText,omitempty" name:"FullText"`

	// 键值索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyValue *RuleKeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`

	// 元字段索引配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *RuleTagInfo `json:"Tag,omitempty" name:"Tag"`
}

type RuleKeyValueInfo struct {

	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 需要建立索引的键值对信息；最大只能配置100个键值对
	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type RuleTagInfo struct {

	// 是否大小写敏感
	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`

	// 标签索引配置中的字段信息
	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 要查询的日志的起始时间，Unix时间戳，单位ms
	From *int64 `json:"From,omitempty" name:"From"`

	// 要查询的日志的结束时间，Unix时间戳，单位ms
	To *int64 `json:"To,omitempty" name:"To"`

	// 查询语句，语句长度最大为4096
	Query *string `json:"Query,omitempty" name:"Query"`

	// 单次查询返回的日志条数，最大值为100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 加载更多日志时使用，透传上次返回的Context值，获取后续的日志内容
	Context *string `json:"Context,omitempty" name:"Context"`

	// 日志接口是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为 desc
	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *SearchLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "From")
	delete(f, "To")
	delete(f, "Query")
	delete(f, "Limit")
	delete(f, "Context")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SearchLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 加载后续内容的Context
		Context *string `json:"Context,omitempty" name:"Context"`

		// 日志查询结果是否全部返回
		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`

		// 返回的是否为分析结果
		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`

		// 如果Analysis为True，则返回分析结果的列名，否则为空
	// 注意：此字段可能返回 null，表示取不到有效值。
		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`

		// 日志查询结果；当Analysis为True时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		Results []*LogInfo `json:"Results,omitempty" name:"Results"`

		// 日志分析结果；当Analysis为False时，可能返回为null
	// 注意：此字段可能返回 null，表示取不到有效值。
		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 待分裂分区ID
	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`

	// 分区切分的哈希key的位置，只在Number=2时有意义
	SplitKey *string `json:"SplitKey,omitempty" name:"SplitKey"`

	// 分区分裂个数(可选)，默认等于2
	Number *int64 `json:"Number,omitempty" name:"Number"`
}

func (r *SplitPartitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicId")
	delete(f, "PartitionId")
	delete(f, "SplitKey")
	delete(f, "Number")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SplitPartitionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分裂结果集
		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SplitPartitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SplitPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type TopicInfo struct {

	// 日志集ID
	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 主题分区个数
	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`

	// 是否开启索引
	Index *bool `json:"Index,omitempty" name:"Index"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 日主主题是否开启采集
	Status *bool `json:"Status,omitempty" name:"Status"`

	// 日志主题绑定的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 该主题是否开启自动分裂
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`

	// 若开启自动分裂的话，该主题能够允许的最大分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`

	// 日主题的存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

type ValueInfo struct {

	// 字段类型，目前支持的类型有：long、text、double
	Type *string `json:"Type,omitempty" name:"Type"`

	// 字段的分词符，只有当字段类型为text时才有意义；输入字符串中的每个字符代表一个分词符
	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`

	// 字段是否开启分析功能
	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`

	// 是否包含中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type WebCallback struct {

	// 回调地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 回调的类型。可选值：
	// <br><li> WeCom
	// <br><li> Http
	CallbackType *string `json:"CallbackType,omitempty" name:"CallbackType"`

	// 回调方法。可选值：
	// <br><li> POST
	// <br><li> PUT
	// 默认值为POST。CallbackType为Http时为必选。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitempty" name:"Method"`

	// 请求头。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*string `json:"Headers,omitempty" name:"Headers"`

	// 请求内容。CallbackType为Http时为必选。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Body *string `json:"Body,omitempty" name:"Body"`

	// 序号
	Index *int64 `json:"Index,omitempty" name:"Index"`
}
