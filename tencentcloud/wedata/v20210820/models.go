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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AdhocDetail struct {
	// 子任务记录Id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 脚本内容
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 任务启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 当前任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 提交任务id
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`
}

type AdhocRecord struct {
	// 任务提交记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 任务提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type AgentStatus struct {
	// 运行中的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Running *int64 `json:"Running,omitnil,omitempty" name:"Running"`

	// 异常的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Abnormal *int64 `json:"Abnormal,omitnil,omitempty" name:"Abnormal"`

	// 操作中的数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InOperation *int64 `json:"InOperation,omitnil,omitempty" name:"InOperation"`
}

type AiOpsEventListenerDTO struct {
	// 事件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 项目展示名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectDisplayName *string `json:"ProjectDisplayName,omitnil,omitempty" name:"ProjectDisplayName"`

	// 事件周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 事件项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 扩展名
	// 注意：此字段可能返回 null，表示取不到有效值。
	PropertiesList []*ParamInfoDs `json:"PropertiesList,omitnil,omitempty" name:"PropertiesList"`

	// 事件广播类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`
}

type AiopsDLCResourceConfigDto struct {
	// Driver资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	DriverSize *string `json:"DriverSize,omitnil,omitempty" name:"DriverSize"`

	// Executor资源
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorSize *string `json:"ExecutorSize,omitnil,omitempty" name:"ExecutorSize"`

	// Executor数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorNumbers *string `json:"ExecutorNumbers,omitnil,omitempty" name:"ExecutorNumbers"`

	// 资源配置方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsInherit *string `json:"IsInherit,omitnil,omitempty" name:"IsInherit"`
}

type AiopsScriptInfo struct {
	// 脚本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 脚本所在COS的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`

	// 脚本所在COS的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 脚本所在COS的桶名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucketName *string `json:"CosBucketName,omitnil,omitempty" name:"CosBucketName"`
}

type AiopsSimpleTaskDto struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 任务创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 配置策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigType *string `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 时间维度
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeDimension *string `json:"TimeDimension,omitnil,omitempty" name:"TimeDimension"`

	// 实例范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceScope *string `json:"InstanceScope,omitnil,omitempty" name:"InstanceScope"`

	// 执行策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExectuorPolicy *string `json:"ExectuorPolicy,omitnil,omitempty" name:"ExectuorPolicy"`
}

type AlarmEventInfo struct {
	// 告警ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警时间
	AlarmTime *string `json:"AlarmTime,omitnil,omitempty" name:"AlarmTime"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 规则名称
	RegularName *string `json:"RegularName,omitnil,omitempty" name:"RegularName"`

	// 告警级别,0表示普通，1表示重要，2表示紧急
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警方式,多个用逗号隔开（1:邮件，2:短信，3:微信，4:语音，5:代表企业微信，6:http）
	AlarmWay *uint64 `json:"AlarmWay,omitnil,omitempty" name:"AlarmWay"`

	// 告警接收人Id，多个用逗号隔开
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 告警指标,0表示任务失败，1表示任务运行超时，2表示任务停止，3表示任务暂停
	AlarmIndicator *uint64 `json:"AlarmIndicator,omitnil,omitempty" name:"AlarmIndicator"`

	// 告警指标描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicatorDesc *string `json:"AlarmIndicatorDesc,omitnil,omitempty" name:"AlarmIndicatorDesc"`

	// 指标阈值，1表示离线任务第一次运行失败，2表示离线任务所有重试完成后失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 预计的超时时间，分钟级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	EstimatedTime *uint64 `json:"EstimatedTime,omitnil,omitempty" name:"EstimatedTime"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 0：部分成功，1：全部成功，2：全部失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSendSuccess *uint64 `json:"IsSendSuccess,omitnil,omitempty" name:"IsSendSuccess"`

	// 是否在免打扰时间内，0:否, 1:是
	// 注意：此字段可能返回 null，表示取不到有效值。
	InQuitePeriods *uint64 `json:"InQuitePeriods,omitnil,omitempty" name:"InQuitePeriods"`

	// 告警记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 消息ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 阈值计算算子，1 : 大于 2 ：小于
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *int64 `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 告警规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegularId *string `json:"RegularId,omitnil,omitempty" name:"RegularId"`

	// 告警接收人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmRecipientName *string `json:"AlarmRecipientName,omitnil,omitempty" name:"AlarmRecipientName"`

	// 告警任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 发送结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendResult *string `json:"SendResult,omitnil,omitempty" name:"SendResult"`

	// 监控对象id
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorObjectId *string `json:"MonitorObjectId,omitnil,omitempty" name:"MonitorObjectId"`

	// 监控对象名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorObjectName *string `json:"MonitorObjectName,omitnil,omitempty" name:"MonitorObjectName"`

	// 指标阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *float64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// 告警原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmReason *string `json:"AlarmReason,omitnil,omitempty" name:"AlarmReason"`
}

type AlarmIndicatorInfo struct {
	// 指标id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 告警指标,0任务失败,1任务运行超时,2任务停止,3任务暂停, 4读取速度,5写入速度,6读取吞吐 7写入吞吐, 8脏数据字节数,9脏数据条数,10任务异常,11任务检测异常, 12重启次数, 13任务延时, 14近20分内的重启次数 15传输延迟,16业务延迟, 50离线包CPU使用率, 51离线包内存使用率, 52离线包并行度使用率, 53离线包排队中的实例数, 54实时包资源使用率, 55实时包运行中的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicator *uint64 `json:"AlarmIndicator,omitnil,omitempty" name:"AlarmIndicator"`

	// 告警指标描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicatorDesc *string `json:"AlarmIndicatorDesc,omitnil,omitempty" name:"AlarmIndicatorDesc"`

	// 指标阈值，1表示离线任务第一次运行失败，2表示离线任务所有重试完成后失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 预计的超时时间，分钟级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	EstimatedTime *uint64 `json:"EstimatedTime,omitnil,omitempty" name:"EstimatedTime"`

	// 告警阈值的算子,1 大于,2 小于
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *uint64 `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 告警指标阈值单位：ms(毫秒)、s(秒)、min(分钟)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicatorUnit *string `json:"AlarmIndicatorUnit,omitnil,omitempty" name:"AlarmIndicatorUnit"`

	// 告警周期
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 告警周期单位:hour,minute,day
	DurationUnit *string `json:"DurationUnit,omitnil,omitempty" name:"DurationUnit"`

	// 周期内最多告警次数
	MaxTimes *int64 `json:"MaxTimes,omitnil,omitempty" name:"MaxTimes"`

	// 指标阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *float64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`
}

type AlarmInfo struct {
	// 关联任务id
	TaskIds *string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 告警类别；failure表示失败告警；overtime表示超时告警
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// 告警方式；SMS表示短信；Email表示邮件；HTTP 表示接口方式；Wechat表示微信方式
	AlarmWay *string `json:"AlarmWay,omitnil,omitempty" name:"AlarmWay"`

	// 告警接收人，多个告警接收人以;分割
	AlarmRecipient *string `json:"AlarmRecipient,omitnil,omitempty" name:"AlarmRecipient"`

	// 告警接收人id，多个告警接收人id以;分割
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// 预计运行的小时，取值范围0-23
	Hours *uint64 `json:"Hours,omitnil,omitempty" name:"Hours"`

	// 预计运行分钟，取值范围0-59
	Minutes *uint64 `json:"Minutes,omitnil,omitempty" name:"Minutes"`

	// 告警出发时机；1表示第一次运行失败；2表示所有重试完成后失败；
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 告警信息id
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警状态设置；1表示可用；0表示不可用，默认可用
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type AlarmReceiverInfo struct {
	// 告警ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 告警接收人ID
	AlarmReceiver *string `json:"AlarmReceiver,omitnil,omitempty" name:"AlarmReceiver"`

	// 邮件，0：未设置，1：成功，2：失败
	Email *uint64 `json:"Email,omitnil,omitempty" name:"Email"`

	// 短信，0：未设置，1：成功，2：失败
	Sms *uint64 `json:"Sms,omitnil,omitempty" name:"Sms"`

	// 微信，0：未设置，1：成功，2：失败
	Wechat *uint64 `json:"Wechat,omitnil,omitempty" name:"Wechat"`

	// 电话，0：未设置，1：成功，2：失败
	Voice *uint64 `json:"Voice,omitnil,omitempty" name:"Voice"`

	// 企业微信，0：未设置，1：成功，2：失败
	Wecom *uint64 `json:"Wecom,omitnil,omitempty" name:"Wecom"`

	// http，0：未设置，1：成功，2：失败
	Http *uint64 `json:"Http,omitnil,omitempty" name:"Http"`

	// 企业微信群，0：未设置，1：成功，2：失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	WecomGroup *uint64 `json:"WecomGroup,omitnil,omitempty" name:"WecomGroup"`

	// 飞书群，0：未设置，1：成功，2：失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	LarkGroup *uint64 `json:"LarkGroup,omitnil,omitempty" name:"LarkGroup"`
}

type Apply struct {
	// 申请人id
	ApplicantId *string `json:"ApplicantId,omitnil,omitempty" name:"ApplicantId"`

	// 申请人名称
	ApplicantName *string `json:"ApplicantName,omitnil,omitempty" name:"ApplicantName"`

	// 审批备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 审批分类key
	ApproveClassification *string `json:"ApproveClassification,omitnil,omitempty" name:"ApproveClassification"`

	// 审批单id
	ApproveId *string `json:"ApproveId,omitnil,omitempty" name:"ApproveId"`

	// 审批类型key
	ApproveType *string `json:"ApproveType,omitnil,omitempty" name:"ApproveType"`

	// 申请原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 审批时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveTime *string `json:"ApproveTime,omitnil,omitempty" name:"ApproveTime"`

	// 审批分类名称
	ApproveClassificationName *string `json:"ApproveClassificationName,omitnil,omitempty" name:"ApproveClassificationName"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 审批类型名称
	ApproveTypeName *string `json:"ApproveTypeName,omitnil,omitempty" name:"ApproveTypeName"`

	// 审批异常或者失败信息
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 申请名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyName *string `json:"ApplyName,omitnil,omitempty" name:"ApplyName"`

	// 审批人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverId *string `json:"ApproverId,omitnil,omitempty" name:"ApproverId"`

	// 审批人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproverName *string `json:"ApproverName,omitnil,omitempty" name:"ApproverName"`

	// 审批所属项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApproveProjectName *string `json:"ApproveProjectName,omitnil,omitempty" name:"ApproveProjectName"`

	// 审批id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApplyId *string `json:"ApplyId,omitnil,omitempty" name:"ApplyId"`
}

type ApproveModify struct {
	// 审批单id
	ApproveId *string `json:"ApproveId,omitnil,omitempty" name:"ApproveId"`

	// 是否修改成功
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`
}

type ApproveType struct {
	// 申请分类key
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 类型名称
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 申请类型key
	Classification *string `json:"Classification,omitnil,omitempty" name:"Classification"`
}

type AttributeItemDTO struct {
	// key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type BaseClusterInfo struct {
	// 集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 集群类型，EMR/CDW等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterType *string `json:"ClusterType,omitnil,omitempty" name:"ClusterType"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 地域中文
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionCn *string `json:"RegionCn,omitnil,omitempty" name:"RegionCn"`

	// 地域英文
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionEn *string `json:"RegionEn,omitnil,omitempty" name:"RegionEn"`

	// 地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionArea *string `json:"RegionArea,omitnil,omitempty" name:"RegionArea"`

	// 集群是否使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	Used *bool `json:"Used,omitnil,omitempty" name:"Used"`

	// 集群状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 集群状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusInfo *string `json:"StatusInfo,omitnil,omitempty" name:"StatusInfo"`

	// 集群存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 集群计算类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeType *string `json:"ComputeType,omitnil,omitempty" name:"ComputeType"`

	// 集群资源量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterResource *string `json:"ClusterResource,omitnil,omitempty" name:"ClusterResource"`

	// 集群付费方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChargeType *string `json:"ChargeType,omitnil,omitempty" name:"ChargeType"`

	// 集群创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 额外配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraConf *string `json:"ExtraConf,omitnil,omitempty" name:"ExtraConf"`

	// ranger账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RangerUserName *string `json:"RangerUserName,omitnil,omitempty" name:"RangerUserName"`

	// cdw账号（用于展示）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CdwUserName *string `json:"CdwUserName,omitnil,omitempty" name:"CdwUserName"`
}

type BaseTenant struct {
	// 租户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 租户标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantName *string `json:"TenantName,omitnil,omitempty" name:"TenantName"`

	// 租户显示名称，一般是中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 租户主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserId *string `json:"OwnerUserId,omitnil,omitempty" name:"OwnerUserId"`

	// 租户的额外配置参数, json格式字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type BaseUser struct {
	// 有云的UIN，全局唯一
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 用户全局唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户显示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 电话号码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PhoneNum *string `json:"PhoneNum,omitnil,omitempty" name:"PhoneNum"`

	// 邮箱
	// 注意：此字段可能返回 null，表示取不到有效值。
	Email *string `json:"Email,omitnil,omitempty" name:"Email"`
}

// Predefined struct for user
type BatchCreateIntegrationTaskAlarmsRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitnil,omitempty" name:"TaskAlarmInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchCreateIntegrationTaskAlarmsRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitnil,omitempty" name:"TaskAlarmInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type BatchCreateTaskVersionAsyncRequestParams struct {
	// 任务信息
	Tasks []*BatchCreateTaskVersionDTO `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否自动运行
	AutoRun *bool `json:"AutoRun,omitnil,omitempty" name:"AutoRun"`

	// 告警方式:email-邮件;sms-短信;wecom-企业微信
	AlarmWays *string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// 告警对象:1-项目管理员，2-任务责任人
	AlarmRecipientTypes *string `json:"AlarmRecipientTypes,omitnil,omitempty" name:"AlarmRecipientTypes"`

	// 是否需要校验父任务已经提交到调度
	NeedCheckParentSubmitted *bool `json:"NeedCheckParentSubmitted,omitnil,omitempty" name:"NeedCheckParentSubmitted"`
}

type BatchCreateTaskVersionAsyncRequest struct {
	*tchttp.BaseRequest
	
	// 任务信息
	Tasks []*BatchCreateTaskVersionDTO `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否自动运行
	AutoRun *bool `json:"AutoRun,omitnil,omitempty" name:"AutoRun"`

	// 告警方式:email-邮件;sms-短信;wecom-企业微信
	AlarmWays *string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// 告警对象:1-项目管理员，2-任务责任人
	AlarmRecipientTypes *string `json:"AlarmRecipientTypes,omitnil,omitempty" name:"AlarmRecipientTypes"`

	// 是否需要校验父任务已经提交到调度
	NeedCheckParentSubmitted *bool `json:"NeedCheckParentSubmitted,omitnil,omitempty" name:"NeedCheckParentSubmitted"`
}

func (r *BatchCreateTaskVersionAsyncRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateTaskVersionAsyncRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tasks")
	delete(f, "ProjectId")
	delete(f, "AutoRun")
	delete(f, "AlarmWays")
	delete(f, "AlarmRecipientTypes")
	delete(f, "NeedCheckParentSubmitted")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchCreateTaskVersionAsyncRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchCreateTaskVersionAsyncResponseParams struct {
	// 批量操作返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchTaskOperateNew `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchCreateTaskVersionAsyncResponse struct {
	*tchttp.BaseResponse
	Response *BatchCreateTaskVersionAsyncResponseParams `json:"Response"`
}

func (r *BatchCreateTaskVersionAsyncResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchCreateTaskVersionAsyncResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchCreateTaskVersionDTO struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// VersionRemark
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

// Predefined struct for user
type BatchDeleteIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否删除开发态任务。默认不删除开发态，为 0 不删除 , 为 1 删除
	DeleteKFFlag *int64 `json:"DeleteKFFlag,omitnil,omitempty" name:"DeleteKFFlag"`

	// 操作名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 本次批量操作涉及任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`
}

type BatchDeleteIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否删除开发态任务。默认不删除开发态，为 0 不删除 , 为 1 删除
	DeleteKFFlag *int64 `json:"DeleteKFFlag,omitnil,omitempty" name:"DeleteKFFlag"`

	// 操作名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 本次批量操作涉及任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`
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
	delete(f, "DeleteKFFlag")
	delete(f, "Name")
	delete(f, "TaskNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type BatchDeleteOpsTasksRequestParams struct {
	// 批量删除的任务TaskId
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitnil,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchDeleteOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 批量删除的任务TaskId
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// true : 删除后下游任务可正常运行
	// false：删除后下游任务不可运行
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`

	// true：通知下游任务责任人
	// false:  不通知下游任务责任人
	EnableNotify *bool `json:"EnableNotify,omitnil,omitempty" name:"EnableNotify"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *BatchDeleteOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "DeleteMode")
	delete(f, "EnableNotify")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchDeleteOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchDeleteOpsTasksResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchDeleteOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchDeleteOpsTasksResponseParams `json:"Response"`
}

func (r *BatchDeleteOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchDeleteOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchForceSuccessIntegrationTaskInstancesRequestParams struct {
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchForceSuccessIntegrationTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchKillIntegrationTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实际传的为taskId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 补数据开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 补数据结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchMakeUpIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 补数据开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 补数据结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type BatchModifyOpsOwnersRequestParams struct {
	// 需要更新责任人的TaskId数组
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 需要更新的责任人userId信息，多个责任人用;连接
	Owners *string `json:"Owners,omitnil,omitempty" name:"Owners"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchModifyOpsOwnersRequest struct {
	*tchttp.BaseRequest
	
	// 需要更新责任人的TaskId数组
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 需要更新的责任人userId信息，多个责任人用;连接
	Owners *string `json:"Owners,omitnil,omitempty" name:"Owners"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *BatchModifyOpsOwnersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyOpsOwnersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "Owners")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchModifyOpsOwnersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchModifyOpsOwnersResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchOperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchModifyOpsOwnersResponse struct {
	*tchttp.BaseResponse
	Response *BatchModifyOpsOwnersResponseParams `json:"Response"`
}

func (r *BatchModifyOpsOwnersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchModifyOpsOwnersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BatchOperateResult struct {
	// 批量操作成功数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 批量操作失败数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 批量操作的总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type BatchOperateResultOpsDto struct {
	// 结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 错误id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitnil,omitempty" name:"ErrorId"`

	// 错误说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`
}

type BatchOperationOpsDto struct {
	// 批量操作成功数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 批量操作失败数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 批量操作的总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type BatchOpsDTO struct {
	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 成功数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 失败数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailMessageList []*FailMessage `json:"FailMessageList,omitnil,omitempty" name:"FailMessageList"`
}

// Predefined struct for user
type BatchRerunIntegrationTaskInstancesRequestParams struct {
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchRerunIntegrationTaskInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例信息
	Instances []*SchedulerTaskInstanceInfo `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实际传的为taskId
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Running *int64 `json:"Running,omitnil,omitempty" name:"Running"`

	// 执行成功的任务数
	Success *int64 `json:"Success,omitnil,omitempty" name:"Success"`

	// 执行失败的任务数
	Failed *int64 `json:"Failed,omitnil,omitempty" name:"Failed"`

	// 总任务数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type BatchResultDs struct {
	// 成功数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Success *int64 `json:"Success,omitnil,omitempty" name:"Success"`

	// 失败数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Failed *int64 `json:"Failed,omitnil,omitempty" name:"Failed"`

	// 总计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

// Predefined struct for user
type BatchResumeIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型, 201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchResumeIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型, 201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 本次批量操作涉及任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 执行情况备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`

	// 执行情况id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitnil,omitempty" name:"ErrorId"`
}

// Predefined struct for user
type BatchRunOpsTaskRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否补录中间实例,0不补录;1补录
	EnableMakeUp *int64 `json:"EnableMakeUp,omitnil,omitempty" name:"EnableMakeUp"`

	// 任务id列表
	Tasks []*string `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

type BatchRunOpsTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否补录中间实例,0不补录;1补录
	EnableMakeUp *int64 `json:"EnableMakeUp,omitnil,omitempty" name:"EnableMakeUp"`

	// 任务id列表
	Tasks []*string `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

func (r *BatchRunOpsTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRunOpsTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EnableMakeUp")
	delete(f, "Tasks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchRunOpsTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchRunOpsTaskResponseParams struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchRunOpsTaskResponse struct {
	*tchttp.BaseResponse
	Response *BatchRunOpsTaskResponseParams `json:"Response"`
}

func (r *BatchRunOpsTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchRunOpsTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStartIntegrationTasksRequestParams struct {
	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 批量运行集成任务，目前仅实时集成用到了这个参数
	StartTaskInfoSet []*StartTaskInfo `json:"StartTaskInfoSet,omitnil,omitempty" name:"StartTaskInfoSet"`
}

type BatchStartIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 批量运行集成任务，目前仅实时集成用到了这个参数
	StartTaskInfoSet []*StartTaskInfo `json:"StartTaskInfoSet,omitnil,omitempty" name:"StartTaskInfoSet"`
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
	delete(f, "TaskType")
	delete(f, "ProjectId")
	delete(f, "TaskIds")
	delete(f, "StartTaskInfoSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStartIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStartIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 本次批量操作成功任务id，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type BatchStopIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 本次批量操作涉及成功任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type BatchStopOpsTasksRequestParams struct {
	// 批量停止任务的TaskId
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type BatchStopOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 批量停止任务的TaskId
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *BatchStopOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIdList")
	delete(f, "ProjectId")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStopOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopOpsTasksResponseParams struct {
	// 返回批量操作成功个数、失败个数、操作总数
	Data *BatchOperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchStopOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *BatchStopOpsTasksResponseParams `json:"Response"`
}

func (r *BatchStopOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopWorkflowsByIdsRequestParams struct {
	// 工作流id列表
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type BatchStopWorkflowsByIdsRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id列表
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *BatchStopWorkflowsByIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopWorkflowsByIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowIds")
	delete(f, "ProjectId")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchStopWorkflowsByIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchStopWorkflowsByIdsResponseParams struct {
	// 操作返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BatchStopWorkflowsByIdsResponse struct {
	*tchttp.BaseResponse
	Response *BatchStopWorkflowsByIdsResponseParams `json:"Response"`
}

func (r *BatchStopWorkflowsByIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BatchStopWorkflowsByIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchSuspendIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, SUSPEND_WITHOUT_SP,RESUME, COMMIT, TIMESTAMP)	
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 本次批量操作涉及任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`
}

type BatchSuspendIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, SUSPEND_WITHOUT_SP,RESUME, COMMIT, TIMESTAMP)	
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 本次批量操作涉及任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`
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
	delete(f, "Event")
	delete(f, "TaskNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchSuspendIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchSuspendIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type BatchTaskOperateNew struct {
	// 操作Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *uint64 `json:"JobId,omitnil,omitempty" name:"JobId"`
}

// Predefined struct for user
type BatchUpdateIntegrationTasksRequestParams struct {
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 责任人（多个责任人用小写分号隔开；离线任务传入的是账号名，实时任务传入的是账号id）
	Incharge *string `json:"Incharge,omitnil,omitempty" name:"Incharge"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 责任人Id（多个责任人用小写分号隔开）
	InchargeIds *string `json:"InchargeIds,omitnil,omitempty" name:"InchargeIds"`

	// 本次批量操作涉及任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`
}

type BatchUpdateIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 责任人（多个责任人用小写分号隔开；离线任务传入的是账号名，实时任务传入的是账号id）
	Incharge *string `json:"Incharge,omitnil,omitempty" name:"Incharge"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 责任人Id（多个责任人用小写分号隔开）
	InchargeIds *string `json:"InchargeIds,omitnil,omitempty" name:"InchargeIds"`

	// 本次批量操作涉及任务，用于审计
	TaskNames []*string `json:"TaskNames,omitnil,omitempty" name:"TaskNames"`
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
	delete(f, "InchargeIds")
	delete(f, "TaskNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BatchUpdateIntegrationTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BatchUpdateIntegrationTasksResponseParams struct {
	// 操作成功的任务数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 操作失败的任务数
	FailedCount *int64 `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// 任务总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 速度值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*SpeedValue `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type CheckAlarmRegularNameExistRequestParams struct {
	// 项目名称
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则名称
	AlarmRegularName *string `json:"AlarmRegularName,omitnil,omitempty" name:"AlarmRegularName"`

	// 任务ID
	//
	// Deprecated: TaskId is deprecated.
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主键ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务类型:201.实时,202.离线
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 监控对象类型(1:所有任务,2:指定任务,3:指定责任人,4:指定资源组)
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

type CheckAlarmRegularNameExistRequest struct {
	*tchttp.BaseRequest
	
	// 项目名称
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则名称
	AlarmRegularName *string `json:"AlarmRegularName,omitnil,omitempty" name:"AlarmRegularName"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 主键ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务类型:201.实时,202.离线
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 监控对象类型(1:所有任务,2:指定任务,3:指定责任人,4:指定资源组)
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
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
	delete(f, "AlarmRegularName")
	delete(f, "TaskId")
	delete(f, "Id")
	delete(f, "TaskType")
	delete(f, "MonitorType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAlarmRegularNameExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAlarmRegularNameExistResponseParams struct {
	// 是否重名
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 节点ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type CheckIntegrationNodeNameExistsRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 节点ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 同步类型1.单表同步，2.解决方案
	SyncType *int64 `json:"SyncType,omitnil,omitempty" name:"SyncType"`
}

type CheckIntegrationTaskNameExistsRequest struct {
	*tchttp.BaseRequest
	
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 同步类型1.单表同步，2.解决方案
	SyncType *int64 `json:"SyncType,omitnil,omitempty" name:"SyncType"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 任务名重复类型（0:未重复, 1:开发态重复, 2:生产态重复）
	ExistsType *int64 `json:"ExistsType,omitnil,omitempty" name:"ExistsType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型（跟调度传参保持一致27）
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type CheckTaskNameExistRequest struct {
	*tchttp.BaseRequest
	
	// 项目id/工作空间id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型（跟调度传参保持一致27）
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CollectionFolderOpsDto struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页面数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前页面数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*FolderOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type CollectionInstanceOpsDto struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页面数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前页面数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InstanceOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type ColumnAggregationLineage struct {
	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 父节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 元数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreType *string `json:"MetastoreType,omitnil,omitempty" name:"MetastoreType"`

	// 字符串类型的父节点集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentSet *string `json:"ParentSet,omitnil,omitempty" name:"ParentSet"`

	// 字符串类型的子节点集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildSet *string `json:"ChildSet,omitnil,omitempty" name:"ChildSet"`

	// 列信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnInfoSet []*SimpleColumnInfo `json:"ColumnInfoSet,omitnil,omitempty" name:"ColumnInfoSet"`
}

type ColumnBasicInfo struct {
	// 表的全局唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 数据源全局唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据源名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 数据类型,string/int等
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataType *string `json:"DataType,omitnil,omitempty" name:"DataType"`

	// 字段类型, varchar(32)/int(10)等
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnType *string `json:"ColumnType,omitnil,omitempty" name:"ColumnType"`

	// 字段默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnDefault *string `json:"ColumnDefault,omitnil,omitempty" name:"ColumnDefault"`

	// 索引类型, PRI/MUL/PARTITION等,普通字段该值为空串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnKey *string `json:"ColumnKey,omitnil,omitempty" name:"ColumnKey"`

	// 字段顺序标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnPosition *int64 `json:"ColumnPosition,omitnil,omitempty" name:"ColumnPosition"`

	// 字段注释
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnComment *string `json:"ColumnComment,omitnil,omitempty" name:"ColumnComment"`

	// 数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoreType *string `json:"StoreType,omitnil,omitempty" name:"StoreType"`

	// 所属项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 所属项目英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 所属项目中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectDisplayName *string `json:"ProjectDisplayName,omitnil,omitempty" name:"ProjectDisplayName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 精度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`
}

type ColumnData struct {
	// ColumnName1
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// ColumnName1
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventProjectName *string `json:"EventProjectName,omitnil,omitempty" name:"EventProjectName"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ColumnItem struct {
	// ColumnName1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// ColumnName1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnRef *string `json:"ColumnRef,omitnil,omitempty" name:"ColumnRef"`
}

type ColumnLineageInfo struct {
	// 血缘id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 字段中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnNameCn *string `json:"ColumnNameCn,omitnil,omitempty" name:"ColumnNameCn"`

	// 字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnType *string `json:"ColumnType,omitnil,omitempty" name:"ColumnType"`

	// 关系参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationParams *string `json:"RelationParams,omitnil,omitempty" name:"RelationParams"`

	// 参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 父id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 元数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreType *string `json:"MetastoreType,omitnil,omitempty" name:"MetastoreType"`

	// 元数据类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreTypeName *string `json:"MetastoreTypeName,omitnil,omitempty" name:"MetastoreTypeName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 字段全名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualifiedName *string `json:"QualifiedName,omitnil,omitempty" name:"QualifiedName"`

	// 下游节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownStreamCount *int64 `json:"DownStreamCount,omitnil,omitempty" name:"DownStreamCount"`

	// 上游节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpStreamCount *int64 `json:"UpStreamCount,omitnil,omitempty" name:"UpStreamCount"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 由中心节点出发的路径信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrefixPath *string `json:"PrefixPath,omitnil,omitempty" name:"PrefixPath"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 任务id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*string `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 父节点列表字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentSet *string `json:"ParentSet,omitnil,omitempty" name:"ParentSet"`

	// 子节点列表字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildSet *string `json:"ChildSet,omitnil,omitempty" name:"ChildSet"`

	// 额外参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtParams []*LineageParamRecord `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 表ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`
}

type ColumnMeta struct {
	// 字段英文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameEn *string `json:"NameEn,omitnil,omitempty" name:"NameEn"`

	// 字段中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameCn *string `json:"NameCn,omitnil,omitempty" name:"NameCn"`

	// 字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 字段序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 是否为分区字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitnil,omitempty" name:"IsPartition"`

	// 字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// HBase列簇属性集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnFamiliesFieldSet []*Pair `json:"ColumnFamiliesFieldSet,omitnil,omitempty" name:"ColumnFamiliesFieldSet"`

	// 对应码表字典ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictionaryId *string `json:"DictionaryId,omitnil,omitempty" name:"DictionaryId"`

	// 对应码表字典名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DictionaryName *string `json:"DictionaryName,omitnil,omitempty" name:"DictionaryName"`

	// 安全等级：名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelName *string `json:"LevelName,omitnil,omitempty" name:"LevelName"`

	// 安全等级：值范围1-10
	// 注意：此字段可能返回 null，表示取不到有效值。
	LevelRank *int64 `json:"LevelRank,omitnil,omitempty" name:"LevelRank"`

	// influxdb字段类别
	// 注意：此字段可能返回 null，表示取不到有效值。
	InfluxCategory *string `json:"InfluxCategory,omitnil,omitempty" name:"InfluxCategory"`
}

// Predefined struct for user
type CommitIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 0.仅提交，1.立即启动，2.停止线上作业，丢弃作业状态数据，重新启动运行，3.暂停线上作业，保留作业状态数据，继续运行，4.保留作业状态数据，继续运行
	CommitType *int64 `json:"CommitType,omitnil,omitempty" name:"CommitType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 额外参数
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 提交版本描述
	VersionDesc *string `json:"VersionDesc,omitnil,omitempty" name:"VersionDesc"`

	// 提交版本号
	InstanceVersion *int64 `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 前端操作类型描述
	EventDesc *string `json:"EventDesc,omitnil,omitempty" name:"EventDesc"`
}

type CommitIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 0.仅提交，1.立即启动，2.停止线上作业，丢弃作业状态数据，重新启动运行，3.暂停线上作业，保留作业状态数据，继续运行，4.保留作业状态数据，继续运行
	CommitType *int64 `json:"CommitType,omitnil,omitempty" name:"CommitType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 额外参数
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 提交版本描述
	VersionDesc *string `json:"VersionDesc,omitnil,omitempty" name:"VersionDesc"`

	// 提交版本号
	InstanceVersion *int64 `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 前端操作类型描述
	EventDesc *string `json:"EventDesc,omitnil,omitempty" name:"EventDesc"`
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
	delete(f, "ExtConfig")
	delete(f, "VersionDesc")
	delete(f, "InstanceVersion")
	delete(f, "EventDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type CommitRuleGroupTaskRequestParams struct {
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 触发类型 1.手动触发 2.调度事中触发 3.周期调度触发
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 规则配置列表
	ExecRuleConfig []*RuleConfig `json:"ExecRuleConfig,omitnil,omitempty" name:"ExecRuleConfig"`

	// 执行配置
	ExecConfig *RuleExecConfig `json:"ExecConfig,omitnil,omitempty" name:"ExecConfig"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 该规则运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type CommitRuleGroupTaskRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 触发类型 1.手动触发 2.调度事中触发 3.周期调度触发
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 规则配置列表
	ExecRuleConfig []*RuleConfig `json:"ExecRuleConfig,omitnil,omitempty" name:"ExecRuleConfig"`

	// 执行配置
	ExecConfig *RuleExecConfig `json:"ExecConfig,omitnil,omitempty" name:"ExecConfig"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 该规则运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

func (r *CommitRuleGroupTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitRuleGroupTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleGroupId")
	delete(f, "TriggerType")
	delete(f, "ExecRuleConfig")
	delete(f, "ExecConfig")
	delete(f, "ProjectId")
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CommitRuleGroupTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CommitRuleGroupTaskResponseParams struct {
	// 规则组执行id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleGroupExecResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CommitRuleGroupTaskResponse struct {
	*tchttp.BaseResponse
	Response *CommitRuleGroupTaskResponseParams `json:"Response"`
}

func (r *CommitRuleGroupTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CommitRuleGroupTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonContent struct {
	// 详情内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type CommonId struct {
	// Id值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type CommonIdOpsDto struct {
	// 返回补录计划id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type CompareResult struct {
	// 对比结果项列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*CompareResultItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 检测总行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRows *uint64 `json:"TotalRows,omitnil,omitempty" name:"TotalRows"`

	// 检测通过行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PassRows *uint64 `json:"PassRows,omitnil,omitempty" name:"PassRows"`

	// 检测不通过行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerRows *uint64 `json:"TriggerRows,omitnil,omitempty" name:"TriggerRows"`
}

type CompareResultItem struct {
	// 对比结果， 1为真 2为假
	// 注意：此字段可能返回 null，表示取不到有效值。
	FixResult *uint64 `json:"FixResult,omitnil,omitempty" name:"FixResult"`

	// 质量sql执行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultValue *string `json:"ResultValue,omitnil,omitempty" name:"ResultValue"`

	// 阈值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*ThresholdValue `json:"Values,omitnil,omitempty" name:"Values"`

	// 比较操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 比较类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareType *uint64 `json:"CompareType,omitnil,omitempty" name:"CompareType"`

	// 值比较类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueComputeType *uint64 `json:"ValueComputeType,omitnil,omitempty" name:"ValueComputeType"`
}

type CompareRule struct {
	// 比较条件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*CompareRuleItem `json:"Items,omitnil,omitempty" name:"Items"`

	// 周期性模板默认周期，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *uint64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// o 表示 或，a 表示 且，数字表示items下标
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComputeExpression *string `json:"ComputeExpression,omitnil,omitempty" name:"ComputeExpression"`
}

type CompareRuleItem struct {
	// 比较类型 1.固定值  2.波动值  3.数值范围比较  4.枚举范围比较  5.不用比较
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareType *uint64 `json:"CompareType,omitnil,omitempty" name:"CompareType"`

	// 比较操作类型
	// <  <=  ==  =>  > !=
	// IRLCRO:在区间内(左闭右开)
	// IRLORC:在区间内(左开右闭)
	// IRLCRC:在区间内(左闭右闭)
	// IRLORO:在区间内(左开右开)
	// NRLCRO:不在区间内(左闭右开)
	// NRLORC:不在区间内(左开右闭)
	// NRLCRC:不在区间内(左闭右闭)
	// NRLORO:不在区间内(左开右开)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 质量统计值类型 1.绝对值  2.上升 3. 下降  4._C包含   5. N_C不包含
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueComputeType *uint64 `json:"ValueComputeType,omitnil,omitempty" name:"ValueComputeType"`

	// 比较阈值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueList []*ThresholdValue `json:"ValueList,omitnil,omitempty" name:"ValueList"`
}

type Content struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 诊断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Diagnose *string `json:"Diagnose,omitnil,omitempty" name:"Diagnose"`

	// 理由
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type CosTokenResponse struct {
	// token id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// token内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 密钥id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretId *string `json:"SecretId,omitnil,omitempty" name:"SecretId"`

	// 密钥内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretKey *string `json:"SecretKey,omitnil,omitempty" name:"SecretKey"`

	// 响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	Response *string `json:"Response,omitnil,omitempty" name:"Response"`

	// 用户uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredTime *uint64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 操作者uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`
}

// Predefined struct for user
type CountOpsInstanceStateRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type CountOpsInstanceStateRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *CountOpsInstanceStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CountOpsInstanceStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CountOpsInstanceStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CountOpsInstanceStateResponseParams struct {
	// 任务对应实例的状态统计
	Data *TaskInstanceCountDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CountOpsInstanceStateResponse struct {
	*tchttp.BaseResponse
	Response *CountOpsInstanceStateResponseParams `json:"Response"`
}

func (r *CountOpsInstanceStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CountOpsInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomFunctionRequestParams struct {
	// 枚举值：HIVE、SPARK、DLC
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 枚举值：ANALYSIS(函数)、ENCRYPTION(加密函数)、AGGREGATE(聚合函数)、LOGIC(逻辑函数)、DATE_AND_TIME(日期与时间函数)、MATH(数学函数)、CONVERSION(转换函数)、STRING(字符串函数)、IP_AND_DOMAIN(IP和域名函数)、WINDOW(窗口函数)、OTHER(其他函数)
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 函数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群实例引擎 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type CreateCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 枚举值：HIVE、SPARK、DLC
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 枚举值：ANALYSIS(函数)、ENCRYPTION(加密函数)、AGGREGATE(聚合函数)、LOGIC(逻辑函数)、DATE_AND_TIME(日期与时间函数)、MATH(数学函数)、CONVERSION(转换函数)、STRING(字符串函数)、IP_AND_DOMAIN(IP和域名函数)、WINDOW(窗口函数)、OTHER(其他函数)
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 函数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集群实例引擎 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 数据库名称
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 归属项目ID
	OwnerProjectId *string `json:"OwnerProjectId,omitnil,omitempty" name:"OwnerProjectId"`

	// 归属项目Name
	OwnerProjectName *string `json:"OwnerProjectName,omitnil,omitempty" name:"OwnerProjectName"`

	// 归属项目Name中文
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitnil,omitempty" name:"OwnerProjectIdent"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitnil,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitnil,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitnil,omitempty" name:"Collect"`

	// cos桶信息
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitnil,omitempty" name:"COSRegion"`

	// 连接测试结果
	ConnectResult *string `json:"ConnectResult,omitnil,omitempty" name:"ConnectResult"`

	// 开发环境数据源配置
	DevelopmentParams *string `json:"DevelopmentParams,omitnil,omitempty" name:"DevelopmentParams"`

	// 新建数据源的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type CreateDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 归属项目ID
	OwnerProjectId *string `json:"OwnerProjectId,omitnil,omitempty" name:"OwnerProjectId"`

	// 归属项目Name
	OwnerProjectName *string `json:"OwnerProjectName,omitnil,omitempty" name:"OwnerProjectName"`

	// 归属项目Name中文
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitnil,omitempty" name:"OwnerProjectIdent"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitnil,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitnil,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitnil,omitempty" name:"Collect"`

	// cos桶信息
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitnil,omitempty" name:"COSRegion"`

	// 连接测试结果
	ConnectResult *string `json:"ConnectResult,omitnil,omitempty" name:"ConnectResult"`

	// 开发环境数据源配置
	DevelopmentParams *string `json:"DevelopmentParams,omitnil,omitempty" name:"DevelopmentParams"`

	// 新建数据源的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	delete(f, "ConnectResult")
	delete(f, "DevelopmentParams")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDataSourceResponseParams struct {
	// 主键ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *uint64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateDsFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`
}

type CreateDsFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`
}

func (r *CreateDsFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDsFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "ParentsFolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDsFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDsFolderResponseParams struct {
	// 文件夹Id，null则创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDsFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateDsFolderResponseParams `json:"Response"`
}

func (r *CreateDsFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDsFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHiveTableByDDLRequestParams struct {
	// 数据源ID
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 建hive表ddl
	DDLSql *string `json:"DDLSql,omitnil,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标表类型(HIVE或GBASE)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 责任人
	Incharge *string `json:"Incharge,omitnil,omitempty" name:"Incharge"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`
}

type CreateHiveTableByDDLRequest struct {
	*tchttp.BaseRequest
	
	// 数据源ID
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// 建hive表ddl
	DDLSql *string `json:"DDLSql,omitnil,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标表类型(HIVE或GBASE)
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 责任人
	Incharge *string `json:"Incharge,omitnil,omitempty" name:"Incharge"`

	// schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`
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
	delete(f, "SchemaName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateHiveTableByDDLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateHiveTableByDDLResponseParams struct {
	// 表名称
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// base64转码之后的建表语句
	DDLSql *string `json:"DDLSql,omitnil,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 责任人
	Incharge *string `json:"Incharge,omitnil,omitempty" name:"Incharge"`
}

type CreateHiveTableRequest struct {
	*tchttp.BaseRequest
	
	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	Database *string `json:"Database,omitnil,omitempty" name:"Database"`

	// base64转码之后的建表语句
	DDLSql *string `json:"DDLSql,omitnil,omitempty" name:"DDLSql"`

	// 表权限 ，默认为0:项目共享;1:仅个人与管理员
	Privilege *int64 `json:"Privilege,omitnil,omitempty" name:"Privilege"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 责任人
	Incharge *string `json:"Incharge,omitnil,omitempty" name:"Incharge"`
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
	IsSuccess *bool `json:"IsSuccess,omitnil,omitempty" name:"IsSuccess"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateIntegrationNodeRequestParams struct {
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type CreateIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 当前任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type CreateIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务信息
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CreateMakeDatetimeInfo struct {
	// 开始日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

// Predefined struct for user
type CreateOfflineTaskRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 间隔，可选，默认1。非空。默认 1
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 延时执行时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 任务结束数据时间。非空。默认当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备注
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// 当前日期
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 跟之前调用调度接口保持一致27
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 时间指定，如月任务指定1，3号，则填入 1，3。非空。默认 "" 月任务：如具体1，3号则写 "1,3"，指定月末不可和具体号数一起输入，仅能为 "L"
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 区分画布和表单
	TaskMode *string `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`
}

type CreateOfflineTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 间隔，可选，默认1。非空。默认 1
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 延时执行时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 任务结束数据时间。非空。默认当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 备注
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// 当前日期
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 跟之前调用调度接口保持一致27
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 时间指定，如月任务指定1，3号，则填入 1，3。非空。默认 "" 月任务：如具体1，3号则写 "1,3"，指定月末不可和具体号数一起输入，仅能为 "L"
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 区分画布和表单
	TaskMode *string `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`
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
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 结果
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateOpsMakePlanRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划名称
	MakeName *string `json:"MakeName,omitnil,omitempty" name:"MakeName"`

	// 补录任务集合
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 补录计划日期范围
	MakeDatetimeList []*CreateMakeDatetimeInfo `json:"MakeDatetimeList,omitnil,omitempty" name:"MakeDatetimeList"`

	// 项目标识
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 补录是否检查父任务状态，默认不检查。不推荐使用，后续会废弃，推荐使用 CheckParentType。
	CheckParent *bool `json:"CheckParent,omitnil,omitempty" name:"CheckParent"`

	// 补录检查父任务类型。取值范围：
	// <li> NONE: 全部不检查 </li>
	// <li> ALL: 检查全部上游父任务 </li>
	// <li> MAKE_SCOPE: 只在（当前补录计划）选中任务中检查 </li>
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 已弃用。任务自依赖类型：parallel（并行），serial（无序串行），orderly（有序串行）
	SelfDependence *string `json:"SelfDependence,omitnil,omitempty" name:"SelfDependence"`

	// 并行度
	ParallelNum *int64 `json:"ParallelNum,omitnil,omitempty" name:"ParallelNum"`

	// 补录实例生成周期是否和原周期相同，默认为true
	SameCycle *bool `json:"SameCycle,omitnil,omitempty" name:"SameCycle"`

	// 补录实例目标周期类型
	TargetTaskCycle *string `json:"TargetTaskCycle,omitnil,omitempty" name:"TargetTaskCycle"`

	// 补录实例目标周期类型指定时间
	TargetTaskAction *int64 `json:"TargetTaskAction,omitnil,omitempty" name:"TargetTaskAction"`

	// 补录实例自定义参数
	MapParamList []*StrToStrMap `json:"MapParamList,omitnil,omitempty" name:"MapParamList"`

	// 创建人id
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 补录计划说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否使用任务原有自依赖配置，默认为true
	SameSelfDependType *bool `json:"SameSelfDependType,omitnil,omitempty" name:"SameSelfDependType"`

	// 补录实例原始周期类型
	SourceTaskCycle *string `json:"SourceTaskCycle,omitnil,omitempty" name:"SourceTaskCycle"`

	// 补录指定的调度资源组（ID） 为空则表示使用任务原有调度执行资源组
	SchedulerResourceGroup *string `json:"SchedulerResourceGroup,omitnil,omitempty" name:"SchedulerResourceGroup"`

	// 补录指定的集成资源组（ID） 为空则表示使用任务原有集成执行资源组
	IntegrationResourceGroup *string `json:"IntegrationResourceGroup,omitnil,omitempty" name:"IntegrationResourceGroup"`

	// 补录指定的调度资源组名称 为空则表示使用任务原有调度执行资源组
	SchedulerResourceGroupName *string `json:"SchedulerResourceGroupName,omitnil,omitempty" name:"SchedulerResourceGroupName"`

	// 补录指定的集成资源组名称 为空则表示使用任务原有集成执行资源组
	IntegrationResourceGroupName *string `json:"IntegrationResourceGroupName,omitnil,omitempty" name:"IntegrationResourceGroupName"`

	// 补录扩展属性
	MakeExtList []*StrToStrMap `json:"MakeExtList,omitnil,omitempty" name:"MakeExtList"`

	// 补录扩展属性
	SameSelfWorkflowDependType *bool `json:"SameSelfWorkflowDependType,omitnil,omitempty" name:"SameSelfWorkflowDependType"`

	// 补录扩展属性
	SelfWorkflowDependency *string `json:"SelfWorkflowDependency,omitnil,omitempty" name:"SelfWorkflowDependency"`
}

type CreateOpsMakePlanRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划名称
	MakeName *string `json:"MakeName,omitnil,omitempty" name:"MakeName"`

	// 补录任务集合
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 补录计划日期范围
	MakeDatetimeList []*CreateMakeDatetimeInfo `json:"MakeDatetimeList,omitnil,omitempty" name:"MakeDatetimeList"`

	// 项目标识
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 补录是否检查父任务状态，默认不检查。不推荐使用，后续会废弃，推荐使用 CheckParentType。
	CheckParent *bool `json:"CheckParent,omitnil,omitempty" name:"CheckParent"`

	// 补录检查父任务类型。取值范围：
	// <li> NONE: 全部不检查 </li>
	// <li> ALL: 检查全部上游父任务 </li>
	// <li> MAKE_SCOPE: 只在（当前补录计划）选中任务中检查 </li>
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 已弃用。任务自依赖类型：parallel（并行），serial（无序串行），orderly（有序串行）
	SelfDependence *string `json:"SelfDependence,omitnil,omitempty" name:"SelfDependence"`

	// 并行度
	ParallelNum *int64 `json:"ParallelNum,omitnil,omitempty" name:"ParallelNum"`

	// 补录实例生成周期是否和原周期相同，默认为true
	SameCycle *bool `json:"SameCycle,omitnil,omitempty" name:"SameCycle"`

	// 补录实例目标周期类型
	TargetTaskCycle *string `json:"TargetTaskCycle,omitnil,omitempty" name:"TargetTaskCycle"`

	// 补录实例目标周期类型指定时间
	TargetTaskAction *int64 `json:"TargetTaskAction,omitnil,omitempty" name:"TargetTaskAction"`

	// 补录实例自定义参数
	MapParamList []*StrToStrMap `json:"MapParamList,omitnil,omitempty" name:"MapParamList"`

	// 创建人id
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 补录计划说明
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 是否使用任务原有自依赖配置，默认为true
	SameSelfDependType *bool `json:"SameSelfDependType,omitnil,omitempty" name:"SameSelfDependType"`

	// 补录实例原始周期类型
	SourceTaskCycle *string `json:"SourceTaskCycle,omitnil,omitempty" name:"SourceTaskCycle"`

	// 补录指定的调度资源组（ID） 为空则表示使用任务原有调度执行资源组
	SchedulerResourceGroup *string `json:"SchedulerResourceGroup,omitnil,omitempty" name:"SchedulerResourceGroup"`

	// 补录指定的集成资源组（ID） 为空则表示使用任务原有集成执行资源组
	IntegrationResourceGroup *string `json:"IntegrationResourceGroup,omitnil,omitempty" name:"IntegrationResourceGroup"`

	// 补录指定的调度资源组名称 为空则表示使用任务原有调度执行资源组
	SchedulerResourceGroupName *string `json:"SchedulerResourceGroupName,omitnil,omitempty" name:"SchedulerResourceGroupName"`

	// 补录指定的集成资源组名称 为空则表示使用任务原有集成执行资源组
	IntegrationResourceGroupName *string `json:"IntegrationResourceGroupName,omitnil,omitempty" name:"IntegrationResourceGroupName"`

	// 补录扩展属性
	MakeExtList []*StrToStrMap `json:"MakeExtList,omitnil,omitempty" name:"MakeExtList"`

	// 补录扩展属性
	SameSelfWorkflowDependType *bool `json:"SameSelfWorkflowDependType,omitnil,omitempty" name:"SameSelfWorkflowDependType"`

	// 补录扩展属性
	SelfWorkflowDependency *string `json:"SelfWorkflowDependency,omitnil,omitempty" name:"SelfWorkflowDependency"`
}

func (r *CreateOpsMakePlanRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpsMakePlanRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "MakeName")
	delete(f, "TaskIdList")
	delete(f, "MakeDatetimeList")
	delete(f, "ProjectIdent")
	delete(f, "CheckParent")
	delete(f, "CheckParentType")
	delete(f, "ProjectName")
	delete(f, "SelfDependence")
	delete(f, "ParallelNum")
	delete(f, "SameCycle")
	delete(f, "TargetTaskCycle")
	delete(f, "TargetTaskAction")
	delete(f, "MapParamList")
	delete(f, "CreatorId")
	delete(f, "Creator")
	delete(f, "Remark")
	delete(f, "SameSelfDependType")
	delete(f, "SourceTaskCycle")
	delete(f, "SchedulerResourceGroup")
	delete(f, "IntegrationResourceGroup")
	delete(f, "SchedulerResourceGroupName")
	delete(f, "IntegrationResourceGroupName")
	delete(f, "MakeExtList")
	delete(f, "SameSelfWorkflowDependType")
	delete(f, "SelfWorkflowDependency")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOpsMakePlanRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOpsMakePlanResponseParams struct {
	// 结果
	Data *CommonIdOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOpsMakePlanResponse struct {
	*tchttp.BaseResponse
	Response *CreateOpsMakePlanResponseParams `json:"Response"`
}

func (r *CreateOpsMakePlanResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOpsMakePlanResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组Id
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 规则模板列表
	RuleTemplateId *uint64 `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则类型 1.系统模版, 2.自定义模版, 3.自定义SQL
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则所属质量维度（1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源字段详细类型，int、string
	SourceObjectDataTypeName *string `json:"SourceObjectDataTypeName,omitnil,omitempty" name:"SourceObjectDataTypeName"`

	// 源字段名称
	SourceObjectValue *string `json:"SourceObjectValue,omitnil,omitempty" name:"SourceObjectValue"`

	// 检测范围 1.全表   2.条件扫描
	ConditionType *uint64 `json:"ConditionType,omitnil,omitempty" name:"ConditionType"`

	// 条件扫描WHERE条件表达式
	ConditionExpression *string `json:"ConditionExpression,omitnil,omitempty" name:"ConditionExpression"`

	// 自定义SQL
	CustomSql *string `json:"CustomSql,omitnil,omitempty" name:"CustomSql"`

	// 报警触发条件
	CompareRule *CompareRule `json:"CompareRule,omitnil,omitempty" name:"CompareRule"`

	// 报警触发级别 1.低, 2.中, 3.高
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库Id
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 目标库Id
	TargetDatabaseId *string `json:"TargetDatabaseId,omitnil,omitempty" name:"TargetDatabaseId"`

	// 目标表Id
	TargetTableId *string `json:"TargetTableId,omitnil,omitempty" name:"TargetTableId"`

	// 目标过滤条件表达式
	TargetConditionExpr *string `json:"TargetConditionExpr,omitnil,omitempty" name:"TargetConditionExpr"`

	// 源字段与目标字段关联条件on表达式
	RelConditionExpr *string `json:"RelConditionExpr,omitnil,omitempty" name:"RelConditionExpr"`

	// 自定义模版sql表达式字段替换参数
	FieldConfig *RuleFieldConfig `json:"FieldConfig,omitnil,omitempty" name:"FieldConfig"`

	// 目标字段名称  CITY
	TargetObjectValue *string `json:"TargetObjectValue,omitnil,omitempty" name:"TargetObjectValue"`

	// 该规则支持的执行引擎列表
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`
}

type CreateRuleRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组Id
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 规则模板列表
	RuleTemplateId *uint64 `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则类型 1.系统模版, 2.自定义模版, 3.自定义SQL
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则所属质量维度（1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源字段详细类型，int、string
	SourceObjectDataTypeName *string `json:"SourceObjectDataTypeName,omitnil,omitempty" name:"SourceObjectDataTypeName"`

	// 源字段名称
	SourceObjectValue *string `json:"SourceObjectValue,omitnil,omitempty" name:"SourceObjectValue"`

	// 检测范围 1.全表   2.条件扫描
	ConditionType *uint64 `json:"ConditionType,omitnil,omitempty" name:"ConditionType"`

	// 条件扫描WHERE条件表达式
	ConditionExpression *string `json:"ConditionExpression,omitnil,omitempty" name:"ConditionExpression"`

	// 自定义SQL
	CustomSql *string `json:"CustomSql,omitnil,omitempty" name:"CustomSql"`

	// 报警触发条件
	CompareRule *CompareRule `json:"CompareRule,omitnil,omitempty" name:"CompareRule"`

	// 报警触发级别 1.低, 2.中, 3.高
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库Id
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 目标库Id
	TargetDatabaseId *string `json:"TargetDatabaseId,omitnil,omitempty" name:"TargetDatabaseId"`

	// 目标表Id
	TargetTableId *string `json:"TargetTableId,omitnil,omitempty" name:"TargetTableId"`

	// 目标过滤条件表达式
	TargetConditionExpr *string `json:"TargetConditionExpr,omitnil,omitempty" name:"TargetConditionExpr"`

	// 源字段与目标字段关联条件on表达式
	RelConditionExpr *string `json:"RelConditionExpr,omitnil,omitempty" name:"RelConditionExpr"`

	// 自定义模版sql表达式字段替换参数
	FieldConfig *RuleFieldConfig `json:"FieldConfig,omitnil,omitempty" name:"FieldConfig"`

	// 目标字段名称  CITY
	TargetObjectValue *string `json:"TargetObjectValue,omitnil,omitempty" name:"TargetObjectValue"`

	// 该规则支持的执行引擎列表
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`
}

func (r *CreateRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RuleGroupId")
	delete(f, "Name")
	delete(f, "TableId")
	delete(f, "RuleTemplateId")
	delete(f, "Type")
	delete(f, "QualityDim")
	delete(f, "SourceObjectDataTypeName")
	delete(f, "SourceObjectValue")
	delete(f, "ConditionType")
	delete(f, "ConditionExpression")
	delete(f, "CustomSql")
	delete(f, "CompareRule")
	delete(f, "AlarmLevel")
	delete(f, "Description")
	delete(f, "DatasourceId")
	delete(f, "DatabaseId")
	delete(f, "TargetDatabaseId")
	delete(f, "TargetTableId")
	delete(f, "TargetConditionExpr")
	delete(f, "RelConditionExpr")
	delete(f, "FieldConfig")
	delete(f, "TargetObjectValue")
	delete(f, "SourceEngineTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleResponseParams struct {
	// 规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *Rule `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleResponseParams `json:"Response"`
}

func (r *CreateRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleTemplateRequestParams struct {
	// 模板类型  1.系统模板   2.自定义模板
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 质量检测维度 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源端数据对象类型 1.常量  2.离线表级   2.离线字段级
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 模板描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 源端对应的引擎类型
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 是否关联其它库表
	MultiSourceFlag *bool `json:"MultiSourceFlag,omitnil,omitempty" name:"MultiSourceFlag"`

	// SQL 表达式
	SqlExpression *string `json:"SqlExpression,omitnil,omitempty" name:"SqlExpression"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否添加where参数
	WhereFlag *bool `json:"WhereFlag,omitnil,omitempty" name:"WhereFlag"`
}

type CreateRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板类型  1.系统模板   2.自定义模板
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 质量检测维度 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源端数据对象类型 1.常量  2.离线表级   2.离线字段级
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 模板描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 源端对应的引擎类型
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 是否关联其它库表
	MultiSourceFlag *bool `json:"MultiSourceFlag,omitnil,omitempty" name:"MultiSourceFlag"`

	// SQL 表达式
	SqlExpression *string `json:"SqlExpression,omitnil,omitempty" name:"SqlExpression"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否添加where参数
	WhereFlag *bool `json:"WhereFlag,omitnil,omitempty" name:"WhereFlag"`
}

func (r *CreateRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "QualityDim")
	delete(f, "SourceObjectType")
	delete(f, "Description")
	delete(f, "SourceEngineTypes")
	delete(f, "MultiSourceFlag")
	delete(f, "SqlExpression")
	delete(f, "ProjectId")
	delete(f, "WhereFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRuleTemplateResponseParams struct {
	// 模板Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *uint64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *CreateRuleTemplateResponseParams `json:"Response"`
}

func (r *CreateRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskAlarmRegularRequestParams struct {
	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitnil,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type CreateTaskAlarmRegularRequest struct {
	*tchttp.BaseRequest
	
	// 告警配置信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitnil,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	AlarmId *int64 `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateTaskFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 工作量ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 父文件夹ID
	ParentFolderId *string `json:"ParentFolderId,omitnil,omitempty" name:"ParentFolderId"`
}

type CreateTaskFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 工作量ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 父文件夹ID
	ParentFolderId *string `json:"ParentFolderId,omitnil,omitempty" name:"ParentFolderId"`
}

func (r *CreateTaskFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "WorkflowId")
	delete(f, "ParentFolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskFolderResponseParams struct {
	// 任务文件夹Id，null则创建失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskFolderResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskFolderResponseParams `json:"Response"`
}

func (r *CreateTaskFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40TCHouse-P，92MapReduce
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 扩展属性
	TaskExt []*TaskExtInfo `json:"TaskExt,omitnil,omitempty" name:"TaskExt"`
}

type CreateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 26离线同步，30Python，31PySpark，32DLC，33Impala，34Hive SQL，35Shell，36Spark SQL，39Spark，40TCHouse-P，92MapReduce
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 扩展属性
	TaskExt []*TaskExtInfo `json:"TaskExt,omitnil,omitempty" name:"TaskExt"`
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
	Data *CommonId `json:"Data,omitnil,omitempty" name:"Data"`

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
type CreateTaskVersionDsRequestParams struct {
	// 任务id
	Task *BatchCreateTaskVersionDTO `json:"Task,omitnil,omitempty" name:"Task"`

	// 是否需要校验父任务已经提交到调度
	NeedCheckParentSubmitted *bool `json:"NeedCheckParentSubmitted,omitnil,omitempty" name:"NeedCheckParentSubmitted"`

	// 是否自动运行
	AutoRun *bool `json:"AutoRun,omitnil,omitempty" name:"AutoRun"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求来源，WEB 前端；CLIENT 客户端
	RequestFromSource *string `json:"RequestFromSource,omitnil,omitempty" name:"RequestFromSource"`

	// 告警方式:email-邮件;sms-短信;wecom-企业微信
	AlarmWays *string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// 告警对象:1-项目管理员，2-任务责任人
	AlarmRecipientTypes *string `json:"AlarmRecipientTypes,omitnil,omitempty" name:"AlarmRecipientTypes"`

	// 是否需要校验循环依赖，默认为 true，如果使用了 CheckTaskCycleLink 和 CheckTaskCycleConfiguration 两个接口校验成功可以传 false，后台服务器不再做校验
	EnableCheckTaskCycleLink *bool `json:"EnableCheckTaskCycleLink,omitnil,omitempty" name:"EnableCheckTaskCycleLink"`
}

type CreateTaskVersionDsRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	Task *BatchCreateTaskVersionDTO `json:"Task,omitnil,omitempty" name:"Task"`

	// 是否需要校验父任务已经提交到调度
	NeedCheckParentSubmitted *bool `json:"NeedCheckParentSubmitted,omitnil,omitempty" name:"NeedCheckParentSubmitted"`

	// 是否自动运行
	AutoRun *bool `json:"AutoRun,omitnil,omitempty" name:"AutoRun"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求来源，WEB 前端；CLIENT 客户端
	RequestFromSource *string `json:"RequestFromSource,omitnil,omitempty" name:"RequestFromSource"`

	// 告警方式:email-邮件;sms-短信;wecom-企业微信
	AlarmWays *string `json:"AlarmWays,omitnil,omitempty" name:"AlarmWays"`

	// 告警对象:1-项目管理员，2-任务责任人
	AlarmRecipientTypes *string `json:"AlarmRecipientTypes,omitnil,omitempty" name:"AlarmRecipientTypes"`

	// 是否需要校验循环依赖，默认为 true，如果使用了 CheckTaskCycleLink 和 CheckTaskCycleConfiguration 两个接口校验成功可以传 false，后台服务器不再做校验
	EnableCheckTaskCycleLink *bool `json:"EnableCheckTaskCycleLink,omitnil,omitempty" name:"EnableCheckTaskCycleLink"`
}

func (r *CreateTaskVersionDsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskVersionDsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Task")
	delete(f, "NeedCheckParentSubmitted")
	delete(f, "AutoRun")
	delete(f, "ProjectId")
	delete(f, "RequestFromSource")
	delete(f, "AlarmWays")
	delete(f, "AlarmRecipientTypes")
	delete(f, "EnableCheckTaskCycleLink")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTaskVersionDsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateTaskVersionDsResponseParams struct {
	// 版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateTaskVersionDsResponse struct {
	*tchttp.BaseResponse
	Response *CreateTaskVersionDsResponseParams `json:"Response"`
}

func (r *CreateTaskVersionDsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTaskVersionDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowDsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流描述
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`
}

type CreateWorkflowDsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流描述
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`
}

func (r *CreateWorkflowDsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowDsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowName")
	delete(f, "FolderId")
	delete(f, "WorkflowDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowDsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowDsResponseParams struct {
	// 工作流ID
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowDsResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowDsResponseParams `json:"Response"`
}

func (r *CreateWorkflowDsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DagInstancesRequestParams struct {
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件,当前接口需要把要查询的示例信息放在该字段
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

type DagInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件,当前接口需要把要查询的示例信息放在该字段
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

func (r *DagInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DagInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	delete(f, "SearchCondition")
	delete(f, "OptType")
	delete(f, "OperatorName")
	delete(f, "OperatorId")
	delete(f, "ProjectId")
	delete(f, "ProjectIdent")
	delete(f, "ProjectName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Count")
	delete(f, "RequestBaseInfo")
	delete(f, "IsCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DagInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DagInstancesResponseParams struct {
	// 结果
	Data *CollectionInstanceOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DagInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DagInstancesResponseParams `json:"Response"`
}

func (r *DagInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DagInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DailyScoreInfo struct {
	// 统计日期 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticsDate *int64 `json:"StatisticsDate,omitnil,omitempty" name:"StatisticsDate"`

	// 评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`
}

type DataCheckStat struct {
	// 表总数
	TableTotal *uint64 `json:"TableTotal,omitnil,omitempty" name:"TableTotal"`

	// 字段总数
	ColumnTotal *uint64 `json:"ColumnTotal,omitnil,omitempty" name:"ColumnTotal"`

	// 表配置检测数
	TableConfig *uint64 `json:"TableConfig,omitnil,omitempty" name:"TableConfig"`

	// 字段配置检测数
	ColumnConfig *uint64 `json:"ColumnConfig,omitnil,omitempty" name:"ColumnConfig"`

	// 表实际检测数
	TableExec *uint64 `json:"TableExec,omitnil,omitempty" name:"TableExec"`

	// 字段实际检测数
	ColumnExec *uint64 `json:"ColumnExec,omitnil,omitempty" name:"ColumnExec"`
}

type DataServicePublishedApiListFilter struct {
	// 请求路径关键词筛选
	PathUrl *string `json:"PathUrl,omitnil,omitempty" name:"PathUrl"`

	// Api名称关键词筛选
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// Api认证方式筛选 0:免认证 1:应用认证
	AuthTypes []*uint64 `json:"AuthTypes,omitnil,omitempty" name:"AuthTypes"`

	// 服务Api状态 1:已上线  3:已下线
	ApiStatus []*uint64 `json:"ApiStatus,omitnil,omitempty" name:"ApiStatus"`

	// API配置方式 0:向导、1、脚本、2、注册Api
	ConfigTypes []*uint64 `json:"ConfigTypes,omitnil,omitempty" name:"ConfigTypes"`
}

type DataServiceRequestListOrder struct {
	// 排序参数名称
	// 取值： 
	//   CreateTime 表示按照创建时间排序
	//   ModifyTime 表示按照更新时间排序
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序参数顺序
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type DataServiceRequestParam struct {
	// 参数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 绑定字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindField *string `json:"BindField,omitnil,omitempty" name:"BindField"`

	// 参数类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 参数位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamPosition *string `json:"ParamPosition,omitnil,omitempty" name:"ParamPosition"`

	// 操作符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 是否为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	NonEmpty *uint64 `json:"NonEmpty,omitnil,omitempty" name:"NonEmpty"`

	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 示例值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExampleValue *string `json:"ExampleValue,omitnil,omitempty" name:"ExampleValue"`

	// 参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type DataServiceResponseParam struct {
	// 参数名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 绑定字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindField *string `json:"BindField,omitnil,omitempty" name:"BindField"`

	// 参数类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamType *string `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// 示例值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExampleValue *string `json:"ExampleValue,omitnil,omitempty" name:"ExampleValue"`

	// 参数描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type DataSourceInfo struct {
	// 若数据源列表为绑定数据库，则为db名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据源描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 数据源引擎的实例ID，如CDB实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源引擎所属区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 数据源类型:枚举值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源所属的集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 应用ID AppId
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 业务侧数据源的配置信息扩展
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizParams *string `json:"BizParams,omitnil,omitempty" name:"BizParams"`

	// 数据源类别：绑定引擎、绑定数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 数据源展示名，为了可视化查看
	// 注意：此字段可能返回 null，表示取不到有效值。
	Display *string `json:"Display,omitnil,omitempty" name:"Display"`

	// 数据源责任人账号ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerAccount *string `json:"OwnerAccount,omitnil,omitempty" name:"OwnerAccount"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源责任人账号名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerAccountName *string `json:"OwnerAccountName,omitnil,omitempty" name:"OwnerAccountName"`

	// 集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 归属项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerProjectId *string `json:"OwnerProjectId,omitnil,omitempty" name:"OwnerProjectId"`

	// 归属项目Name
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerProjectName *string `json:"OwnerProjectName,omitnil,omitempty" name:"OwnerProjectName"`

	// 归属项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitnil,omitempty" name:"OwnerProjectIdent"`

	// 授权项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorityProjectName *string `json:"AuthorityProjectName,omitnil,omitempty" name:"AuthorityProjectName"`

	// 授权用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthorityUserName *string `json:"AuthorityUserName,omitnil,omitempty" name:"AuthorityUserName"`

	// 是否有编辑权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Edit *bool `json:"Edit,omitnil,omitempty" name:"Edit"`

	// 是否有授权权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Author *bool `json:"Author,omitnil,omitempty" name:"Author"`

	// 是否有转交权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Deliver *bool `json:"Deliver,omitnil,omitempty" name:"Deliver"`

	// 数据源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceStatus *string `json:"DataSourceStatus,omitnil,omitempty" name:"DataSourceStatus"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Params json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamsString *string `json:"ParamsString,omitnil,omitempty" name:"ParamsString"`

	// BizParams json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizParamsString *string `json:"BizParamsString,omitnil,omitempty" name:"BizParamsString"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *int64 `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 数据源页面展示类型，与Type对应
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowType *string `json:"ShowType,omitnil,omitempty" name:"ShowType"`

	// 当前数据源生产源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductId *uint64 `json:"ProductId,omitnil,omitempty" name:"ProductId"`

	// 当前数据源开发源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevelopmentId *uint64 `json:"DevelopmentId,omitnil,omitempty" name:"DevelopmentId"`

	// 同params 内容为开发数据源的数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevelopmentParams *string `json:"DevelopmentParams,omitnil,omitempty" name:"DevelopmentParams"`
}

type DataSourceInfoPage struct {
	// 分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*DataSourceInfo `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type DatabaseInfo struct {
	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据库id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceType *uint64 `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`

	// 数据库原始名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginDatabaseName *string `json:"OriginDatabaseName,omitnil,omitempty" name:"OriginDatabaseName"`

	// schema名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginSchemaName *string `json:"OriginSchemaName,omitnil,omitempty" name:"OriginSchemaName"`

	// 0或者未返回.未定义，1.生产 2.开发
	// 注意：此字段可能返回 null，表示取不到有效值。
	DsEnvType *int64 `json:"DsEnvType,omitnil,omitempty" name:"DsEnvType"`

	// EMR引擎部署方式：CVM/TKE
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDeployType *string `json:"ClusterDeployType,omitnil,omitempty" name:"ClusterDeployType"`

	// 模式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`
}

type DatabaseMeta struct {
	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 技术类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreType *string `json:"MetastoreType,omitnil,omitempty" name:"MetastoreType"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 项目英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 数据源类别：绑定引擎、绑定数据库,可用值:DB,ENGINE
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 数据源描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源引擎的实例ID，如CDB实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 数据源引擎所属区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// db名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 项目中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectDisplayName *string `json:"ProjectDisplayName,omitnil,omitempty" name:"ProjectDisplayName"`

	// 责任人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerAccountName *string `json:"OwnerAccountName,omitnil,omitempty" name:"OwnerAccountName"`

	// 数据来源展示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 数据库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据来源类型：hive/mysql/hbase等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`

	// 存储量大小,单位为 byte
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 格式化后的存储量大小，带单位，如 12B
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSizeWithUnit *string `json:"StorageSizeWithUnit,omitnil,omitempty" name:"StorageSizeWithUnit"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type DatasourceBaseInfo struct {
	// 若数据源列表为绑定数据库，则为db名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseNames []*string `json:"DatabaseNames,omitnil,omitempty" name:"DatabaseNames"`

	// 数据源描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 数据源引擎的实例ID，如CDB实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源引擎所属区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 数据源类型:枚举值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源所属的集群id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 数据源版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 数据源附带参数信息Params json字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamsString *string `json:"ParamsString,omitnil,omitempty" name:"ParamsString"`

	// 区分数据源类型自定义源还是系统源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`
}

// Predefined struct for user
type DeleteCustomFunctionRequestParams struct {
	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 函数 ID
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 项目ID，必须填
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 函数名称
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 函数类型，HIVE，SPARK，DLC，CDW_POSTGRESQL
	FunctionType *string `json:"FunctionType,omitnil,omitempty" name:"FunctionType"`

	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 模式名
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 函数命令格式
	CommandFormat *string `json:"CommandFormat,omitnil,omitempty" name:"CommandFormat"`
}

type DeleteCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 函数 ID
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 项目ID，必须填
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 函数名称
	FunctionName *string `json:"FunctionName,omitnil,omitempty" name:"FunctionName"`

	// 函数类型，HIVE，SPARK，DLC，CDW_POSTGRESQL
	FunctionType *string `json:"FunctionType,omitnil,omitempty" name:"FunctionType"`

	// 数据库名
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 模式名
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 函数命令格式
	CommandFormat *string `json:"CommandFormat,omitnil,omitempty" name:"CommandFormat"`
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
	delete(f, "FunctionName")
	delete(f, "FunctionType")
	delete(f, "DatabaseName")
	delete(f, "SchemaName")
	delete(f, "CommandFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomFunctionResponseParams struct {
	// 函数 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteDataSourcesRequest struct {
	*tchttp.BaseRequest
	
	// id列表
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDataSourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDataSourcesResponseParams struct {
	// 是否删除成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteDsFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

type DeleteDsFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`
}

func (r *DeleteDsFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDsFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDsFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDsFolderResponseParams struct {
	// true代表删除成功，false代表删除失败
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDsFolderResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDsFolderResponseParams `json:"Response"`
}

func (r *DeleteDsFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDsFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFilePathRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 使用状态
	// - true
	// - false
	UseStatus *string `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// 文件路径
	FilePaths []*string `json:"FilePaths,omitnil,omitempty" name:"FilePaths"`
}

type DeleteFilePathRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源ID
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 使用状态
	// - true
	// - false
	UseStatus *string `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// 文件路径
	FilePaths []*string `json:"FilePaths,omitnil,omitempty" name:"FilePaths"`
}

func (r *DeleteFilePathRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFilePathRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceIds")
	delete(f, "UseStatus")
	delete(f, "FilePaths")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFilePathRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFilePathResponseParams struct {
	// 文件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserFileList []*UserFileInfo `json:"UserFileList,omitnil,omitempty" name:"UserFileList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFilePathResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFilePathResponseParams `json:"Response"`
}

func (r *DeleteFilePathResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFilePathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DeleteFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DeleteFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteFileResponseParams struct {
	// 删除结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteFileResponseParams `json:"Response"`
}

func (r *DeleteFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteIntegrationNodeRequestParams struct {
	// 节点id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 节点id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 任务删除成功与否标识
	// 0表示删除成功
	// 1 表示失败，失败原因见 DeleteErrInfo
	// 100 表示running or suspend task can't be deleted失败，失败原因也会写到DeleteErrInfo里面
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteFlag *int64 `json:"DeleteFlag,omitnil,omitempty" name:"DeleteFlag"`

	// 删除失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteErrInfo *string `json:"DeleteErrInfo,omitnil,omitempty" name:"DeleteErrInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 虚拟任务标记(跟之前调度接口保持一致默认false)
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`
}

type DeleteOfflineTaskRequest struct {
	*tchttp.BaseRequest
	
	// 操作者name
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 虚拟任务标记(跟之前调度接口保持一致默认false)
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteProjectParamDsRequestParams struct {
	// 参数名
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteProjectParamDsRequest struct {
	*tchttp.BaseRequest
	
	// 参数名
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteProjectParamDsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectParamDsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ParamKey")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectParamDsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectParamDsResponseParams struct {
	// 结果 true 删除成功
	// false 删除失败
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectParamDsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectParamDsResponseParams `json:"Response"`
}

func (r *DeleteProjectParamDsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectParamDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectUsersRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID列表
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

type DeleteProjectUsersRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 用户ID列表
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`
}

func (r *DeleteProjectUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UserIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteProjectUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteProjectUsersResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteProjectUsersResponse struct {
	*tchttp.BaseResponse
	Response *DeleteProjectUsersResponseParams `json:"Response"`
}

func (r *DeleteProjectUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteProjectUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFileRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DeleteResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源id
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

func (r *DeleteResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ResourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFileResponseParams struct {
	// 资源删除结果
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceFileResponseParams `json:"Response"`
}

func (r *DeleteResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFilesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 使用状态
	UseStatus *bool `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// 资源id列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 资源路径列表
	FilePaths []*string `json:"FilePaths,omitnil,omitempty" name:"FilePaths"`
}

type DeleteResourceFilesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 使用状态
	UseStatus *bool `json:"UseStatus,omitnil,omitempty" name:"UseStatus"`

	// 资源id列表
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 资源路径列表
	FilePaths []*string `json:"FilePaths,omitnil,omitempty" name:"FilePaths"`
}

func (r *DeleteResourceFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "UseStatus")
	delete(f, "ResourceIds")
	delete(f, "FilePaths")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteResourceFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceFilesResponseParams struct {
	// 资源批量删除结果
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteResourceFilesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteResourceFilesResponseParams `json:"Response"`
}

func (r *DeleteResourceFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteResourceFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteResourceRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
}

type DeleteResourceRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源ID
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteRuleRequestParams struct {
	// 质量规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DeleteRuleRequest struct {
	*tchttp.BaseRequest
	
	// 质量规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DeleteRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleResponseParams struct {
	// 是否删除成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleResponseParams `json:"Response"`
}

func (r *DeleteRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleTemplateRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 模板Id列表
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

type DeleteRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 模板Id列表
	Ids []*uint64 `json:"Ids,omitnil,omitempty" name:"Ids"`
}

func (r *DeleteRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Ids")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRuleTemplateResponseParams struct {
	// 删除成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRuleTemplateResponseParams `json:"Response"`
}

func (r *DeleteRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskAlarmRegularRequestParams struct {
	// 主键ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型(201表示实时任务，202表示离线任务)
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DeleteTaskAlarmRegularRequest struct {
	*tchttp.BaseRequest
	
	// 主键ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型(201表示实时任务，202表示离线任务)
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteTaskDsRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否删除脚本
	// true：删除
	// false：不删除
	// 不传默认false
	DeleteScript *bool `json:"DeleteScript,omitnil,omitempty" name:"DeleteScript"`

	// 任务操作是否消息通知下游任务责任人true：通知
	// false：不通知
	// 不传默认false
	OperateInform *bool `json:"OperateInform,omitnil,omitempty" name:"OperateInform"`

	// 任务ID
	// 和VirtualTaskId选填一个
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 虚拟任务id
	// 和TaskId选填一个
	VirtualTaskId *string `json:"VirtualTaskId,omitnil,omitempty" name:"VirtualTaskId"`

	// 虚拟任务标记
	// true：是虚拟任务
	// false：不是虚拟任务
	// 不传默认false
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`

	// 任务删除方式
	// true：不针对下游任务实例进行强制失败
	// false：针对下游任务实例进行强制失败
	// 不传默认false
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
}

type DeleteTaskDsRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否删除脚本
	// true：删除
	// false：不删除
	// 不传默认false
	DeleteScript *bool `json:"DeleteScript,omitnil,omitempty" name:"DeleteScript"`

	// 任务操作是否消息通知下游任务责任人true：通知
	// false：不通知
	// 不传默认false
	OperateInform *bool `json:"OperateInform,omitnil,omitempty" name:"OperateInform"`

	// 任务ID
	// 和VirtualTaskId选填一个
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 虚拟任务id
	// 和TaskId选填一个
	VirtualTaskId *string `json:"VirtualTaskId,omitnil,omitempty" name:"VirtualTaskId"`

	// 虚拟任务标记
	// true：是虚拟任务
	// false：不是虚拟任务
	// 不传默认false
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`

	// 任务删除方式
	// true：不针对下游任务实例进行强制失败
	// false：针对下游任务实例进行强制失败
	// 不传默认false
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
}

func (r *DeleteTaskDsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskDsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "DeleteScript")
	delete(f, "OperateInform")
	delete(f, "TaskId")
	delete(f, "VirtualTaskId")
	delete(f, "VirtualFlag")
	delete(f, "DeleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTaskDsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTaskDsResponseParams struct {
	// 是否删除成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteTaskDsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTaskDsResponseParams `json:"Response"`
}

func (r *DeleteTaskDsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTaskDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowByIdRequestParams struct {
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 删除后下游任务的处理方式，true:下游任务均正常运行 false:下游任务均运行失败
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`

	// 删除任务后是否通知下游任务责任人, true:通知 false:不通知
	EnableNotify *bool `json:"EnableNotify,omitnil,omitempty" name:"EnableNotify"`
}

type DeleteWorkflowByIdRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 删除后下游任务的处理方式，true:下游任务均正常运行 false:下游任务均运行失败
	DeleteMode *bool `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`

	// 删除任务后是否通知下游任务责任人, true:通知 false:不通知
	EnableNotify *bool `json:"EnableNotify,omitnil,omitempty" name:"EnableNotify"`
}

func (r *DeleteWorkflowByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	delete(f, "ProjectId")
	delete(f, "DeleteMode")
	delete(f, "EnableNotify")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWorkflowByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWorkflowByIdResponseParams struct {
	// 删除结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWorkflowByIdResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWorkflowByIdResponseParams `json:"Response"`
}

func (r *DeleteWorkflowByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWorkflowByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DependencyConfig struct {
	// 仅五种周期运行依赖配置： HOUR,DAY,WEEK,MONTH,YEAR,CRONTAB,MINUTE
	DependConfType *string `json:"DependConfType,omitnil,omitempty" name:"DependConfType"`

	// 依赖配置从属周期类型，CURRENT_HOUR，PREVIOUS_HOUR，CURRENT_DAY，PREVIOUS_DAY，PREVIOUS_WEEK，PREVIOUS_FRIDAY，PREVIOUS_WEEKEND，CURRENT_MONTH，PREVIOUS_MONTH，PREVIOUS_END_OF_MONTH
	//      * PREVIOUS_BEGIN_OF_MONTH，ALL_MONTH_OF_YEAR，ALL_DAY_OF_YEAR，CURRENT_YEAR，CURRENT，CURRENT_MINUTE，PREVIOUS_MINUTE_CYCLE，PREVIOUS_HOUR_CYCLE
	SubordinateCyclicType *string `json:"SubordinateCyclicType,omitnil,omitempty" name:"SubordinateCyclicType"`

	// WAITING，等待（默认策略）EXECUTING:执行
	DependencyStrategy *string `json:"DependencyStrategy,omitnil,omitempty" name:"DependencyStrategy"`

	// 父任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentTask *TaskInnerInfo `json:"ParentTask,omitnil,omitempty" name:"ParentTask"`

	// 子任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SonTask *TaskInnerInfo `json:"SonTask,omitnil,omitempty" name:"SonTask"`
}

// Predefined struct for user
type DescribeAlarmEventsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件(key可以是：AlarmLevel,AlarmIndicator,KeyWord)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段（AlarmTime）
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 类型(201表示实时，202表示离线)
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 监控对象类型(1:所有任务,2:指定任务,3:指定责任人,4:指定资源组)
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

type DescribeAlarmEventsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件(key可以是：AlarmLevel,AlarmIndicator,KeyWord)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段（AlarmTime）
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 类型(201表示实时，202表示离线)
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 监控对象类型(1:所有任务,2:指定任务,3:指定责任人,4:指定资源组)
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
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
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "TaskType")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "MonitorType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmEventsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmEventsResponseParams struct {
	// 告警事件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmEventInfoList []*AlarmEventInfo `json:"AlarmEventInfoList,omitnil,omitempty" name:"AlarmEventInfoList"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 类型
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 告警接收人ID(逗号分隔)
	AlarmRecipient *string `json:"AlarmRecipient,omitnil,omitempty" name:"AlarmRecipient"`

	// 告警接收人姓名(逗号分隔)
	AlarmRecipientName *string `json:"AlarmRecipientName,omitnil,omitempty" name:"AlarmRecipientName"`

	// 告警时间
	AlarmTime *string `json:"AlarmTime,omitnil,omitempty" name:"AlarmTime"`

	// 消息ID
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 告警记录id
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 监控对象类型(1:所有任务,2:指定任务,3:指定责任人,4:指定资源组)
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

type DescribeAlarmReceiverRequest struct {
	*tchttp.BaseRequest
	
	// 告警ID
	AlarmId *string `json:"AlarmId,omitnil,omitempty" name:"AlarmId"`

	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 类型
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 告警接收人ID(逗号分隔)
	AlarmRecipient *string `json:"AlarmRecipient,omitnil,omitempty" name:"AlarmRecipient"`

	// 告警接收人姓名(逗号分隔)
	AlarmRecipientName *string `json:"AlarmRecipientName,omitnil,omitempty" name:"AlarmRecipientName"`

	// 告警时间
	AlarmTime *string `json:"AlarmTime,omitnil,omitempty" name:"AlarmTime"`

	// 消息ID
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// 告警记录id
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 监控对象类型(1:所有任务,2:指定任务,3:指定责任人,4:指定资源组)
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
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
	delete(f, "TaskType")
	delete(f, "AlarmRecipient")
	delete(f, "AlarmRecipientName")
	delete(f, "AlarmTime")
	delete(f, "MessageId")
	delete(f, "RecordId")
	delete(f, "MonitorType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmReceiverRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmReceiverResponseParams struct {
	// 告警接收人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmReceiverInfoList []*AlarmReceiverInfo `json:"AlarmReceiverInfoList,omitnil,omitempty" name:"AlarmReceiverInfoList"`

	// 总记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeAllByFolderNewRequestParams struct {
	// 文件夹属性
	Folder *FolderOpsDto `json:"Folder,omitnil,omitempty" name:"Folder"`

	// 工作流列表
	Workflows []*WorkflowCanvasOpsDto `json:"Workflows,omitnil,omitempty" name:"Workflows"`

	// 目标文件id
	TargetFolderId *string `json:"TargetFolderId,omitnil,omitempty" name:"TargetFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 父文件id
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`

	// 拉取文件夹列表
	IsAddWorkflow *string `json:"IsAddWorkflow,omitnil,omitempty" name:"IsAddWorkflow"`

	// 任务状态
	TaskStates []*string `json:"TaskStates,omitnil,omitempty" name:"TaskStates"`

	// 搜索类型
	FindType *string `json:"FindType,omitnil,omitempty" name:"FindType"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

type DescribeAllByFolderNewRequest struct {
	*tchttp.BaseRequest
	
	// 文件夹属性
	Folder *FolderOpsDto `json:"Folder,omitnil,omitempty" name:"Folder"`

	// 工作流列表
	Workflows []*WorkflowCanvasOpsDto `json:"Workflows,omitnil,omitempty" name:"Workflows"`

	// 目标文件id
	TargetFolderId *string `json:"TargetFolderId,omitnil,omitempty" name:"TargetFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 父文件id
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`

	// 拉取文件夹列表
	IsAddWorkflow *string `json:"IsAddWorkflow,omitnil,omitempty" name:"IsAddWorkflow"`

	// 任务状态
	TaskStates []*string `json:"TaskStates,omitnil,omitempty" name:"TaskStates"`

	// 搜索类型
	FindType *string `json:"FindType,omitnil,omitempty" name:"FindType"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

func (r *DescribeAllByFolderNewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllByFolderNewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Folder")
	delete(f, "Workflows")
	delete(f, "TargetFolderId")
	delete(f, "KeyWords")
	delete(f, "ParentsFolderId")
	delete(f, "IsAddWorkflow")
	delete(f, "TaskStates")
	delete(f, "FindType")
	delete(f, "OptType")
	delete(f, "OperatorName")
	delete(f, "OperatorId")
	delete(f, "ProjectId")
	delete(f, "ProjectIdent")
	delete(f, "ProjectName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Count")
	delete(f, "RequestBaseInfo")
	delete(f, "IsCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllByFolderNewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllByFolderNewResponseParams struct {
	// 结果集
	Data *CollectionFolderOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllByFolderNewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllByFolderNewResponseParams `json:"Response"`
}

func (r *DescribeAllByFolderNewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllByFolderNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApply struct {
	// 申请列表详情
	Rows []*Apply `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 总分页页码
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`

	// 总个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeApproveListRequestParams struct {
	// 审批分类key
	ApproveClassification *string `json:"ApproveClassification,omitnil,omitempty" name:"ApproveClassification"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 自定义条件查询
	Filters []*FilterOptional `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderFieldOptional `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

type DescribeApproveListRequest struct {
	*tchttp.BaseRequest
	
	// 审批分类key
	ApproveClassification *string `json:"ApproveClassification,omitnil,omitempty" name:"ApproveClassification"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 自定义条件查询
	Filters []*FilterOptional `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderFieldOptional `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

func (r *DescribeApproveListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApproveListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApproveClassification")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Filters")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApproveListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApproveListResponseParams struct {
	// 待审批列表详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeApply `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApproveListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApproveListResponseParams `json:"Response"`
}

func (r *DescribeApproveListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApproveListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApproveTypeListRequestParams struct {
	// 类型key
	Classification *string `json:"Classification,omitnil,omitempty" name:"Classification"`
}

type DescribeApproveTypeListRequest struct {
	*tchttp.BaseRequest
	
	// 类型key
	Classification *string `json:"Classification,omitnil,omitempty" name:"Classification"`
}

func (r *DescribeApproveTypeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApproveTypeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Classification")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApproveTypeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApproveTypeListResponseParams struct {
	// 获取审批分类列表
	Data []*ApproveType `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApproveTypeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApproveTypeListResponseParams `json:"Response"`
}

func (r *DescribeApproveTypeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApproveTypeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchOperateTaskDTO struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 工作流Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// 文件夹名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 是否提交
	// 注意：此字段可能返回 null，表示取不到有效值。
	Submit *uint64 `json:"Submit,omitnil,omitempty" name:"Submit"`

	// 引擎：
	// presto\SparkJob\SparkSql
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataEngine *string `json:"DataEngine,omitnil,omitempty" name:"DataEngine"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 创造时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 调度计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceType *string `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`
}

type DescribeBatchOperateTaskPage struct {
	// 总页码数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DescribeBatchOperateTaskDTO `json:"Items,omitnil,omitempty" name:"Items"`

	// 总个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type DescribeBatchOperateTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码
	Page *string `json:"Page,omitnil,omitempty" name:"Page"`

	// 页号
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 状态列表
	// 草稿：'NS'，'N','P','R'
	// 运行：''Y'
	// 停止：'F'
	// 冻结：'O'
	// 停止中：'T'
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// 责任人名列表
	OwnerNameList []*string `json:"OwnerNameList,omitnil,omitempty" name:"OwnerNameList"`

	// 工作流列表
	WorkflowIdList []*string `json:"WorkflowIdList,omitnil,omitempty" name:"WorkflowIdList"`

	// 任务名称搜索
	TaskNameFilter *string `json:"TaskNameFilter,omitnil,omitempty" name:"TaskNameFilter"`

	// 任务类型列表
	TaskTypeList []*string `json:"TaskTypeList,omitnil,omitempty" name:"TaskTypeList"`

	// 文件夹列表
	FordIdList []*string `json:"FordIdList,omitnil,omitempty" name:"FordIdList"`

	// 任务Id搜索
	TaskIdFilter *string `json:"TaskIdFilter,omitnil,omitempty" name:"TaskIdFilter"`

	// 责任人搜索
	OwnerNameFilter *string `json:"OwnerNameFilter,omitnil,omitempty" name:"OwnerNameFilter"`

	// 排序字段：
	// UpdateTime
	// CreateTime
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// asc:升序
	// desc:降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 引擎类型列表：三种
	// SparkJob
	// SparkSql
	// presto
	DataEngineList []*string `json:"DataEngineList,omitnil,omitempty" name:"DataEngineList"`

	// 操作人名
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 1
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`

	// 1
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 数据源ID列表
	DatasourceIdList []*string `json:"DatasourceIdList,omitnil,omitempty" name:"DatasourceIdList"`

	// 数据源类型列表
	DatasourceTypeList []*string `json:"DatasourceTypeList,omitnil,omitempty" name:"DatasourceTypeList"`

	// 调度单位类型列表
	CycleUnitList []*string `json:"CycleUnitList,omitnil,omitempty" name:"CycleUnitList"`

	// 是否筛选出可提交的任务
	CanSubmit *bool `json:"CanSubmit,omitnil,omitempty" name:"CanSubmit"`
}

type DescribeBatchOperateTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码
	Page *string `json:"Page,omitnil,omitempty" name:"Page"`

	// 页号
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 状态列表
	// 草稿：'NS'，'N','P','R'
	// 运行：''Y'
	// 停止：'F'
	// 冻结：'O'
	// 停止中：'T'
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// 责任人名列表
	OwnerNameList []*string `json:"OwnerNameList,omitnil,omitempty" name:"OwnerNameList"`

	// 工作流列表
	WorkflowIdList []*string `json:"WorkflowIdList,omitnil,omitempty" name:"WorkflowIdList"`

	// 任务名称搜索
	TaskNameFilter *string `json:"TaskNameFilter,omitnil,omitempty" name:"TaskNameFilter"`

	// 任务类型列表
	TaskTypeList []*string `json:"TaskTypeList,omitnil,omitempty" name:"TaskTypeList"`

	// 文件夹列表
	FordIdList []*string `json:"FordIdList,omitnil,omitempty" name:"FordIdList"`

	// 任务Id搜索
	TaskIdFilter *string `json:"TaskIdFilter,omitnil,omitempty" name:"TaskIdFilter"`

	// 责任人搜索
	OwnerNameFilter *string `json:"OwnerNameFilter,omitnil,omitempty" name:"OwnerNameFilter"`

	// 排序字段：
	// UpdateTime
	// CreateTime
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// asc:升序
	// desc:降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 引擎类型列表：三种
	// SparkJob
	// SparkSql
	// presto
	DataEngineList []*string `json:"DataEngineList,omitnil,omitempty" name:"DataEngineList"`

	// 操作人名
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 1
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`

	// 1
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 数据源ID列表
	DatasourceIdList []*string `json:"DatasourceIdList,omitnil,omitempty" name:"DatasourceIdList"`

	// 数据源类型列表
	DatasourceTypeList []*string `json:"DatasourceTypeList,omitnil,omitempty" name:"DatasourceTypeList"`

	// 调度单位类型列表
	CycleUnitList []*string `json:"CycleUnitList,omitnil,omitempty" name:"CycleUnitList"`

	// 是否筛选出可提交的任务
	CanSubmit *bool `json:"CanSubmit,omitnil,omitempty" name:"CanSubmit"`
}

func (r *DescribeBatchOperateTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperateTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Page")
	delete(f, "Size")
	delete(f, "StatusList")
	delete(f, "OwnerNameList")
	delete(f, "WorkflowIdList")
	delete(f, "TaskNameFilter")
	delete(f, "TaskTypeList")
	delete(f, "FordIdList")
	delete(f, "TaskIdFilter")
	delete(f, "OwnerNameFilter")
	delete(f, "SortItem")
	delete(f, "SortType")
	delete(f, "DataEngineList")
	delete(f, "UserId")
	delete(f, "OwnerId")
	delete(f, "TenantId")
	delete(f, "DatasourceIdList")
	delete(f, "DatasourceTypeList")
	delete(f, "CycleUnitList")
	delete(f, "CanSubmit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBatchOperateTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBatchOperateTaskResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DescribeBatchOperateTaskPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeBatchOperateTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBatchOperateTaskResponseParams `json:"Response"`
}

func (r *DescribeBatchOperateTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBatchOperateTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeColumnLineageRequestParams struct {
	// 查询方向，INPUT,OUTPUT,BOTH枚举值
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 字段信息
	Data *ColumnLineageInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 单次查询入度
	InputDepth *int64 `json:"InputDepth,omitnil,omitempty" name:"InputDepth"`

	// 单次查询出度
	OutputDepth *int64 `json:"OutputDepth,omitnil,omitempty" name:"OutputDepth"`

	// 额外参数（传递调用方信息）
	ExtParams []*RecordField `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 是否过滤临时表 默认值为true
	IgnoreTemp *bool `json:"IgnoreTemp,omitnil,omitempty" name:"IgnoreTemp"`
}

type DescribeColumnLineageRequest struct {
	*tchttp.BaseRequest
	
	// 查询方向，INPUT,OUTPUT,BOTH枚举值
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 字段信息
	Data *ColumnLineageInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 单次查询入度
	InputDepth *int64 `json:"InputDepth,omitnil,omitempty" name:"InputDepth"`

	// 单次查询出度
	OutputDepth *int64 `json:"OutputDepth,omitnil,omitempty" name:"OutputDepth"`

	// 额外参数（传递调用方信息）
	ExtParams []*RecordField `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 是否过滤临时表 默认值为true
	IgnoreTemp *bool `json:"IgnoreTemp,omitnil,omitempty" name:"IgnoreTemp"`
}

func (r *DescribeColumnLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeColumnLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Data")
	delete(f, "InputDepth")
	delete(f, "OutputDepth")
	delete(f, "ExtParams")
	delete(f, "IgnoreTemp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeColumnLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeColumnLineageResponseParams struct {
	// 字段血缘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnAggregationLineage *ColumnAggregationLineage `json:"ColumnAggregationLineage,omitnil,omitempty" name:"ColumnAggregationLineage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeColumnLineageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeColumnLineageResponseParams `json:"Response"`
}

func (r *DescribeColumnLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeColumnLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeColumnsMetaRequestParams struct {
	// 表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤器
	FilterSet []*Filter `json:"FilterSet,omitnil,omitempty" name:"FilterSet"`

	// 排序字段
	OrderFieldSet []*OrderField `json:"OrderFieldSet,omitnil,omitempty" name:"OrderFieldSet"`

	// 是否查询分区字段，默认false
	IsPartitionQuery *bool `json:"IsPartitionQuery,omitnil,omitempty" name:"IsPartitionQuery"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitnil,omitempty" name:"ComplianceId"`
}

type DescribeColumnsMetaRequest struct {
	*tchttp.BaseRequest
	
	// 表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤器
	FilterSet []*Filter `json:"FilterSet,omitnil,omitempty" name:"FilterSet"`

	// 排序字段
	OrderFieldSet []*OrderField `json:"OrderFieldSet,omitnil,omitempty" name:"OrderFieldSet"`

	// 是否查询分区字段，默认false
	IsPartitionQuery *bool `json:"IsPartitionQuery,omitnil,omitempty" name:"IsPartitionQuery"`

	// 合规组ID
	ComplianceId *int64 `json:"ComplianceId,omitnil,omitempty" name:"ComplianceId"`
}

func (r *DescribeColumnsMetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeColumnsMetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "FilterSet")
	delete(f, "OrderFieldSet")
	delete(f, "IsPartitionQuery")
	delete(f, "ComplianceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeColumnsMetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeColumnsMetaResponseParams struct {
	// 分页返回的
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnMetaSet []*ColumnMeta `json:"ColumnMetaSet,omitnil,omitempty" name:"ColumnMetaSet"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeColumnsMetaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeColumnsMetaResponseParams `json:"Response"`
}

func (r *DescribeColumnsMetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeColumnsMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCheckStatRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeDataCheckStatRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeDataCheckStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCheckStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataCheckStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataCheckStatResponseParams struct {
	// 结果
	Data *DataCheckStat `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataCheckStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataCheckStatResponseParams `json:"Response"`
}

func (r *DescribeDataCheckStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataCheckStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataServicePublishedApiDetailRequestParams struct {
	// 服务Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeDataServicePublishedApiDetailRequest struct {
	*tchttp.BaseRequest
	
	// 服务Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeDataServicePublishedApiDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataServicePublishedApiDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataServicePublishedApiDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataServicePublishedApiDetailResp struct {
	// 服务Api名称
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// 服务请求Path
	PathUrl *string `json:"PathUrl,omitnil,omitempty" name:"PathUrl"`

	// 服务责任人名称
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 服务请求方式
	RequestType *string `json:"RequestType,omitnil,omitempty" name:"RequestType"`

	// 服务标签名称集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiTagNames *string `json:"ApiTagNames,omitnil,omitempty" name:"ApiTagNames"`

	// 服务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiDescription *string `json:"ApiDescription,omitnil,omitempty" name:"ApiDescription"`

	// 服务请求返回示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestExample *string `json:"RequestExample,omitnil,omitempty" name:"RequestExample"`

	// 服务请求成功返回示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestSuccess *string `json:"RequestSuccess,omitnil,omitempty" name:"RequestSuccess"`

	// 服务请求失败返回示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestError *string `json:"RequestError,omitnil,omitempty" name:"RequestError"`

	// 服务请求参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestParam []*DataServiceRequestParam `json:"RequestParam,omitnil,omitempty" name:"RequestParam"`

	// 服务响应参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseParam []*DataServiceResponseParam `json:"ResponseParam,omitnil,omitempty" name:"ResponseParam"`

	// 最大qps
	MaxAllowQps *int64 `json:"MaxAllowQps,omitnil,omitempty" name:"MaxAllowQps"`

	// 最大记录数
	MaxAllowPageSize *int64 `json:"MaxAllowPageSize,omitnil,omitempty" name:"MaxAllowPageSize"`

	// 超时时间，单位ms
	TimeoutPeriod *uint64 `json:"TimeoutPeriod,omitnil,omitempty" name:"TimeoutPeriod"`

	// ApiId
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiId *string `json:"ApiId,omitnil,omitempty" name:"ApiId"`

	// 0:免认证 1:应用认证
	AuthType *uint64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 请求地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayApiUrl *string `json:"GatewayApiUrl,omitnil,omitempty" name:"GatewayApiUrl"`

	// 服务Api状态 1:已上线  3:已下线
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiStatus *uint64 `json:"ApiStatus,omitnil,omitempty" name:"ApiStatus"`
}

// Predefined struct for user
type DescribeDataServicePublishedApiDetailResponseParams struct {
	// 服务详情
	Data *DescribeDataServicePublishedApiDetailResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataServicePublishedApiDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataServicePublishedApiDetailResponseParams `json:"Response"`
}

func (r *DescribeDataServicePublishedApiDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataServicePublishedApiDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataServicePublishedApiListRequestParams struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询参数
	Filters *DataServicePublishedApiListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序配置
	OrderFields []*DataServiceRequestListOrder `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

type DescribeDataServicePublishedApiListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询参数
	Filters *DataServicePublishedApiListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序配置
	OrderFields []*DataServiceRequestListOrder `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

func (r *DescribeDataServicePublishedApiListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataServicePublishedApiListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "Filters")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDataServicePublishedApiListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataServicePublishedApiListResp struct {
	// 服务id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 服务Api名称
	ApiName *string `json:"ApiName,omitnil,omitempty" name:"ApiName"`

	// 所属目录名称
	ApiFolderName *string `json:"ApiFolderName,omitnil,omitempty" name:"ApiFolderName"`

	// 服务Api标签名称集合
	ApiTagNames *string `json:"ApiTagNames,omitnil,omitempty" name:"ApiTagNames"`

	// 服务负责人
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 服务创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// Api的id
	ApiId *string `json:"ApiId,omitnil,omitempty" name:"ApiId"`

	// 服务Api认证方式 0:免认证 1:应用认证
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// 服务Api状态 0:创建 1:已上线 2:已删除 3:已下线
	ApiStatus *int64 `json:"ApiStatus,omitnil,omitempty" name:"ApiStatus"`

	// 配置方式 0:向导、1、脚本、2、注册Api
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigType *uint64 `json:"ConfigType,omitnil,omitempty" name:"ConfigType"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

// Predefined struct for user
type DescribeDataServicePublishedApiListResponseParams struct {
	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务列表
	DataSet []*DescribeDataServicePublishedApiListResp `json:"DataSet,omitnil,omitempty" name:"DataSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDataServicePublishedApiListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDataServicePublishedApiListResponseParams `json:"Response"`
}

func (r *DescribeDataServicePublishedApiListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDataServicePublishedApiListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDataSourceInfoListRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件（暂不支持）
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序配置
	OrderFields *OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 数据源类型，必选（如MYSQL、DLC等）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源名称过滤
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`
}

type DescribeDataSourceInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件（暂不支持）
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序配置
	OrderFields *OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 数据源类型，必选（如MYSQL、DLC等）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源名称过滤
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 数据源信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceSet []*DatasourceBaseInfo `json:"DatasourceSet,omitnil,omitempty" name:"DatasourceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 返回数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序配置
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 可选过滤条件，Filter可选配置(参考): "Name": { "type": "string", "description": "数据源名称" }, "Type": { "type": "string", "description": "类型" }, "ClusterId": { "type": "string", "description": "集群id" }, "CategoryId": { "type": "string", "description": "分类，项目或空间id" }
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDataSourceListRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 返回数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序配置
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 可选过滤条件，Filter可选配置(参考): "Name": { "type": "string", "description": "数据源名称" }, "Type": { "type": "string", "description": "类型" }, "ClusterId": { "type": "string", "description": "集群id" }, "CategoryId": { "type": "string", "description": "分类，项目或空间id" }
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Data *DataSourceInfoPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDatabaseInfoListRequestParams struct {
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 连接类型
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`
}

type DescribeDatabaseInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 连接类型
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`
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
	delete(f, "Filters")
	delete(f, "ConnectionType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseInfoListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseInfoListResponseParams struct {
	// 数据库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseInfo []*DatabaseInfo `json:"DatabaseInfo,omitnil,omitempty" name:"DatabaseInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDatabaseMetasRequestParams struct {
	// 过滤字段，projectIds/msTypes/createTime/modifiedTime
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，如name
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// pagesize
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// pageNumber
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeDatabaseMetasRequest struct {
	*tchttp.BaseRequest
	
	// 过滤字段，projectIds/msTypes/createTime/modifiedTime
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段，如name
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// pagesize
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// pageNumber
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeDatabaseMetasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseMetasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatabaseMetasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatabaseMetasResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseMeta []*DatabaseMeta `json:"DatabaseMeta,omitnil,omitempty" name:"DatabaseMeta"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDatabaseMetasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDatabaseMetasResponseParams `json:"Response"`
}

func (r *DescribeDatabaseMetasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDatabaseMetasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceRequestParams struct {
	// 对象唯一ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// production：生产，development开发
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
}

type DescribeDatasourceRequest struct {
	*tchttp.BaseRequest
	
	// 对象唯一ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// production：生产，development开发
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`
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
	delete(f, "Env")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDatasourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDatasourceResponseParams struct {
	// 数据源对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DataSourceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeDependOpsTasksRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 上游/下游层级1-6级
	Deep *uint64 `json:"Deep,omitnil,omitempty" name:"Deep"`

	// 1: 表示查询上游节点；0:表示查询下游节点；2：表示查询上游和下游节点
	Up *uint64 `json:"Up,omitnil,omitempty" name:"Up"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DescribeDependOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 上游/下游层级1-6级
	Deep *uint64 `json:"Deep,omitnil,omitempty" name:"Deep"`

	// 1: 表示查询上游节点；0:表示查询下游节点；2：表示查询上游和下游节点
	Up *uint64 `json:"Up,omitnil,omitempty" name:"Up"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DescribeDependOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependOpsTasksRequest) FromJsonString(s string) error {
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
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDependOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDependOpsTasksResponseParams struct {
	// 画布任务和链接信息
	Data *OpsTaskCanvasInfoList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDependOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDependOpsTasksResponseParams `json:"Response"`
}

func (r *DescribeDependOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDependTaskListsRequestParams struct {
	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeDependTaskListsRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id列表
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeDependTaskListsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependTaskListsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDependTaskListsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDependTaskListsResponseParams struct {
	// 删除结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDependTaskListsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDependTaskListsResponseParams `json:"Response"`
}

func (r *DescribeDependTaskListsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDependTaskListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDimensionScoreRequestParams struct {
	// 统计日期 时间戳
	StatisticsDate *int64 `json:"StatisticsDate,omitnil,omitempty" name:"StatisticsDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDimensionScoreRequest struct {
	*tchttp.BaseRequest
	
	// 统计日期 时间戳
	StatisticsDate *int64 `json:"StatisticsDate,omitnil,omitempty" name:"StatisticsDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDimensionScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDimensionScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StatisticsDate")
	delete(f, "ProjectId")
	delete(f, "DatasourceId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDimensionScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDimensionScoreResponseParams struct {
	// 维度评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DimensionScore `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDimensionScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDimensionScoreResponseParams `json:"Response"`
}

func (r *DescribeDimensionScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDimensionScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrInstancePageRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务来源 ADHOC || WORKFLOW
	TaskSource *string `json:"TaskSource,omitnil,omitempty" name:"TaskSource"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 提交开始时间 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 提交结束时间 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 文件夹id
	FolderIds []*string `json:"FolderIds,omitnil,omitempty" name:"FolderIds"`

	// 工作流id
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 只看我的
	JustMe *bool `json:"JustMe,omitnil,omitempty" name:"JustMe"`

	// 任务类型
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// 试运行提交人userId列表
	SubmitUsers []*string `json:"SubmitUsers,omitnil,omitempty" name:"SubmitUsers"`

	// 试运行状态
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

type DescribeDrInstancePageRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务来源 ADHOC || WORKFLOW
	TaskSource *string `json:"TaskSource,omitnil,omitempty" name:"TaskSource"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 提交开始时间 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 提交结束时间 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 文件夹id
	FolderIds []*string `json:"FolderIds,omitnil,omitempty" name:"FolderIds"`

	// 工作流id
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 只看我的
	JustMe *bool `json:"JustMe,omitnil,omitempty" name:"JustMe"`

	// 任务类型
	TaskTypes []*string `json:"TaskTypes,omitnil,omitempty" name:"TaskTypes"`

	// 试运行提交人userId列表
	SubmitUsers []*string `json:"SubmitUsers,omitnil,omitempty" name:"SubmitUsers"`

	// 试运行状态
	StatusList []*string `json:"StatusList,omitnil,omitempty" name:"StatusList"`
}

func (r *DescribeDrInstancePageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrInstancePageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskSource")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "TaskName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "FolderIds")
	delete(f, "WorkflowIds")
	delete(f, "JustMe")
	delete(f, "TaskTypes")
	delete(f, "SubmitUsers")
	delete(f, "StatusList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDrInstancePageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDrInstancePageResponseParams struct {
	// 结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *DrInstanceOpsDtoPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDrInstancePageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDrInstancePageResponseParams `json:"Response"`
}

func (r *DescribeDrInstancePageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDrInstancePageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDsFolderTreeRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否一级拉取 true 是 
	// false 否
	FirstLevelPull *bool `json:"FirstLevelPull,omitnil,omitempty" name:"FirstLevelPull"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 关键字搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 是否包含工作流 true 是 
	// false 否
	IncludeWorkflow *bool `json:"IncludeWorkflow,omitnil,omitempty" name:"IncludeWorkflow"`

	// 是否包含任务 true 是 
	// false 否
	IncludeTask *bool `json:"IncludeTask,omitnil,omitempty" name:"IncludeTask"`

	// 是否包含虚拟任务，当 IncludeTask 为 true 的时候，该参数才生效，默认为 true
	IncludeVirtualTask *bool `json:"IncludeVirtualTask,omitnil,omitempty" name:"IncludeVirtualTask"`

	// 任务目录id
	TaskFolderId *string `json:"TaskFolderId,omitnil,omitempty" name:"TaskFolderId"`

	// classification.分类展示  catalog.目录展示
	DisplayType *string `json:"DisplayType,omitnil,omitempty" name:"DisplayType"`

	// 是否包含任务目录 true 是 
	// false 否
	IncludeTaskFolder *bool `json:"IncludeTaskFolder,omitnil,omitempty" name:"IncludeTaskFolder"`
}

type DescribeDsFolderTreeRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否一级拉取 true 是 
	// false 否
	FirstLevelPull *bool `json:"FirstLevelPull,omitnil,omitempty" name:"FirstLevelPull"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 关键字搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 是否包含工作流 true 是 
	// false 否
	IncludeWorkflow *bool `json:"IncludeWorkflow,omitnil,omitempty" name:"IncludeWorkflow"`

	// 是否包含任务 true 是 
	// false 否
	IncludeTask *bool `json:"IncludeTask,omitnil,omitempty" name:"IncludeTask"`

	// 是否包含虚拟任务，当 IncludeTask 为 true 的时候，该参数才生效，默认为 true
	IncludeVirtualTask *bool `json:"IncludeVirtualTask,omitnil,omitempty" name:"IncludeVirtualTask"`

	// 任务目录id
	TaskFolderId *string `json:"TaskFolderId,omitnil,omitempty" name:"TaskFolderId"`

	// classification.分类展示  catalog.目录展示
	DisplayType *string `json:"DisplayType,omitnil,omitempty" name:"DisplayType"`

	// 是否包含任务目录 true 是 
	// false 否
	IncludeTaskFolder *bool `json:"IncludeTaskFolder,omitnil,omitempty" name:"IncludeTaskFolder"`
}

func (r *DescribeDsFolderTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDsFolderTreeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FirstLevelPull")
	delete(f, "FolderId")
	delete(f, "WorkflowId")
	delete(f, "Keyword")
	delete(f, "IncludeWorkflow")
	delete(f, "IncludeTask")
	delete(f, "IncludeVirtualTask")
	delete(f, "TaskFolderId")
	delete(f, "DisplayType")
	delete(f, "IncludeTaskFolder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDsFolderTreeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDsFolderTreeResponseParams struct {
	// 统一树结构返回属性列表
	Data []*PathNodeDsVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDsFolderTreeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDsFolderTreeResponseParams `json:"Response"`
}

func (r *DescribeDsFolderTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDsFolderTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDsParentFolderTreeRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务展示形式, 示例取值
	// -    classification:分类展示
	// -    catalog:目录展示
	DisplayType *string `json:"DisplayType,omitnil,omitempty" name:"DisplayType"`
}

type DescribeDsParentFolderTreeRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹ID
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务展示形式, 示例取值
	// -    classification:分类展示
	// -    catalog:目录展示
	DisplayType *string `json:"DisplayType,omitnil,omitempty" name:"DisplayType"`
}

func (r *DescribeDsParentFolderTreeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDsParentFolderTreeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderId")
	delete(f, "WorkflowId")
	delete(f, "TaskId")
	delete(f, "DisplayType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDsParentFolderTreeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDsParentFolderTreeResponseParams struct {
	// 统一树结构返回属性列表
	Data []*PathNodeDsVO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDsParentFolderTreeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDsParentFolderTreeResponseParams `json:"Response"`
}

func (r *DescribeDsParentFolderTreeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDsParentFolderTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDutyScheduleDetailsRequestParams struct {
	// 值班表id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 查询时间
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 扩展字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeDutyScheduleDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 值班表id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 查询时间
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 扩展字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeDutyScheduleDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDutyScheduleDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "QueryDate")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDutyScheduleDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDutyScheduleDetailsResponseParams struct {
	// 值班日历信息
	Data []*DutyScheduleDetailsInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDutyScheduleDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDutyScheduleDetailsResponseParams `json:"Response"`
}

func (r *DescribeDutyScheduleDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDutyScheduleDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDutyScheduleListRequestParams struct {
	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 值班表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeDutyScheduleListRequest struct {
	*tchttp.BaseRequest
	
	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 值班表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeDutyScheduleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDutyScheduleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDutyScheduleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDutyScheduleListResponseParams struct {
	// 无
	Data *DutySchedule `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDutyScheduleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDutyScheduleListResponseParams `json:"Response"`
}

func (r *DescribeDutyScheduleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDutyScheduleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventCasesRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件实例目录,示例取值:
	// - 已过期: expired
	// - 未过期: consuming
	// - 全部: all
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件类型
	//
	// Deprecated: EventType is deprecated.
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件分割类型
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 事件广播类型
	//
	// Deprecated: EventBroadcastType is deprecated.
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// 事件实例状态,示例取值:
	// - 已消费: COMSUMED
	// - 已过期: EXPIRED
	// - 待消费: ACTIVE
	// - 消费中: CONSUMING
	//
	// Deprecated: Status is deprecated.
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件实例最小创建时间
	CreationTimeStart *string `json:"CreationTimeStart,omitnil,omitempty" name:"CreationTimeStart"`

	// 事件实例最大创建时间
	CreationTimeEnd *string `json:"CreationTimeEnd,omitnil,omitempty" name:"CreationTimeEnd"`

	// 事件实例最小触发时间
	EventTriggeredTimeStart *string `json:"EventTriggeredTimeStart,omitnil,omitempty" name:"EventTriggeredTimeStart"`

	// 事件实例最大触发时间
	EventTriggeredTimeEnd *string `json:"EventTriggeredTimeEnd,omitnil,omitempty" name:"EventTriggeredTimeEnd"`

	// 事件实例最小消费时间
	LogTimeStart *string `json:"LogTimeStart,omitnil,omitempty" name:"LogTimeStart"`

	// 事件实例最大消费时间
	LogTimeEnd *string `json:"LogTimeEnd,omitnil,omitempty" name:"LogTimeEnd"`

	// 事件实例数据时间
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 事件实例有效时间
	TimeToLive *string `json:"TimeToLive,omitnil,omitempty" name:"TimeToLive"`

	// 排序字段
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序顺序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type DescribeEventCasesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件实例目录,示例取值:
	// - 已过期: expired
	// - 未过期: consuming
	// - 全部: all
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件类型
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件分割类型
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 事件广播类型
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// 事件实例状态,示例取值:
	// - 已消费: COMSUMED
	// - 已过期: EXPIRED
	// - 待消费: ACTIVE
	// - 消费中: CONSUMING
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件实例最小创建时间
	CreationTimeStart *string `json:"CreationTimeStart,omitnil,omitempty" name:"CreationTimeStart"`

	// 事件实例最大创建时间
	CreationTimeEnd *string `json:"CreationTimeEnd,omitnil,omitempty" name:"CreationTimeEnd"`

	// 事件实例最小触发时间
	EventTriggeredTimeStart *string `json:"EventTriggeredTimeStart,omitnil,omitempty" name:"EventTriggeredTimeStart"`

	// 事件实例最大触发时间
	EventTriggeredTimeEnd *string `json:"EventTriggeredTimeEnd,omitnil,omitempty" name:"EventTriggeredTimeEnd"`

	// 事件实例最小消费时间
	LogTimeStart *string `json:"LogTimeStart,omitnil,omitempty" name:"LogTimeStart"`

	// 事件实例最大消费时间
	LogTimeEnd *string `json:"LogTimeEnd,omitnil,omitempty" name:"LogTimeEnd"`

	// 事件实例数据时间
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 事件实例有效时间
	TimeToLive *string `json:"TimeToLive,omitnil,omitempty" name:"TimeToLive"`

	// 排序字段
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序顺序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

func (r *DescribeEventCasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventCasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Category")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "EventName")
	delete(f, "EventType")
	delete(f, "EventSubType")
	delete(f, "EventBroadcastType")
	delete(f, "Status")
	delete(f, "CreationTimeStart")
	delete(f, "CreationTimeEnd")
	delete(f, "EventTriggeredTimeStart")
	delete(f, "EventTriggeredTimeEnd")
	delete(f, "LogTimeStart")
	delete(f, "LogTimeEnd")
	delete(f, "Dimension")
	delete(f, "TimeToLive")
	delete(f, "SortItem")
	delete(f, "SortType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventCasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventCasesResponseParams struct {
	// 事件实例分页查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *EventCaseAuditLogVOCollection `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventCasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventCasesResponseParams `json:"Response"`
}

func (r *DescribeEventCasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventCasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventConsumeTasksRequestParams struct {
	// 事件实例ID
	EventCaseId *string `json:"EventCaseId,omitnil,omitempty" name:"EventCaseId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeEventConsumeTasksRequest struct {
	*tchttp.BaseRequest
	
	// 事件实例ID
	EventCaseId *string `json:"EventCaseId,omitnil,omitempty" name:"EventCaseId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeEventConsumeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventConsumeTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EventCaseId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventConsumeTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventConsumeTasksResponseParams struct {
	// 事件消费任务记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *EventCaseConsumeLogOptDtoCollection `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventConsumeTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventConsumeTasksResponseParams `json:"Response"`
}

func (r *DescribeEventConsumeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventConsumeTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`
}

type DescribeEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`
}

func (r *DescribeEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EventName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEventResponseParams struct {
	// 事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *EventOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeEventResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEventResponseParams `json:"Response"`
}

func (r *DescribeEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecStrategyRequestParams struct {
	// 规则组Id
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeExecStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 规则组Id
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeExecStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleGroupId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExecStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExecStrategyResponseParams struct {
	// 规则组执行策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleGroupExecStrategy `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExecStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExecStrategyResponseParams `json:"Response"`
}

func (r *DescribeExecStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExecStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFieldBasicInfoRequestParams struct {
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

type DescribeFieldBasicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

func (r *DescribeFieldBasicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFieldBasicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeFieldBasicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeFieldBasicInfoResponseParams struct {
	// 字段元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnBasicInfoList []*ColumnBasicInfo `json:"ColumnBasicInfoList,omitnil,omitempty" name:"ColumnBasicInfoList"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeFieldBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeFieldBasicInfoResponseParams `json:"Response"`
}

func (r *DescribeFieldBasicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeFieldBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFolderWorkflowListData struct {
	// 工作流信息列表
	Items []*Workflow `json:"Items,omitnil,omitempty" name:"Items"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

// Predefined struct for user
type DescribeFolderWorkflowListRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeFolderWorkflowListRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`

	// 关键字
	KeyWords *string `json:"KeyWords,omitnil,omitempty" name:"KeyWords"`

	// 页码，默认1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小，默认10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
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
	Data *DescribeFolderWorkflowListData `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Kinds []*FunctionTypeOrKind `json:"Kinds,omitnil,omitempty" name:"Kinds"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Types []*FunctionTypeOrKind `json:"Types,omitnil,omitempty" name:"Types"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeInstanceByCycleRequestParams struct {
	// 1
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 1
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`
}

type DescribeInstanceByCycleRequest struct {
	*tchttp.BaseRequest
	
	// 1
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 1
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`
}

func (r *DescribeInstanceByCycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceByCycleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TenantId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceByCycleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceByCycleResponseParams struct {
	// 统计结果
	Data []*TaskByCycle `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceByCycleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceByCycleResponseParams `json:"Response"`
}

func (r *DescribeInstanceByCycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceByCycleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLastLogRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
}

type DescribeInstanceLastLogRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 一页展示的条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 周期列表（如天，一次性），可选
	CycleList []*string `json:"CycleList,omitnil,omitempty" name:"CycleList"`

	// 责任人
	OwnerList []*string `json:"OwnerList,omitnil,omitempty" name:"OwnerList"`

	// 跟之前保持一致
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitnil,omitempty" name:"SortCol"`

	// 类型列表（如python任务类型：30
	// pyspark任务类型：31
	// hivesql任务类型：34
	// shell任务类型：35
	// sparksql任务类型：36 jdbcsql任务类型：21 dlc任务类型：32），可选
	TaskTypeList []*int64 `json:"TaskTypeList,omitnil,omitempty" name:"TaskTypeList"`

	// 状态列表（如成功 2，正在执行 1），可选
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`

	// 任务名称
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 一页展示的条数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 周期列表（如天，一次性），可选
	CycleList []*string `json:"CycleList,omitnil,omitempty" name:"CycleList"`

	// 责任人
	OwnerList []*string `json:"OwnerList,omitnil,omitempty" name:"OwnerList"`

	// 跟之前保持一致
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitnil,omitempty" name:"SortCol"`

	// 类型列表（如python任务类型：30
	// pyspark任务类型：31
	// hivesql任务类型：34
	// shell任务类型：35
	// sparksql任务类型：36 jdbcsql任务类型：21 dlc任务类型：32），可选
	TaskTypeList []*int64 `json:"TaskTypeList,omitnil,omitempty" name:"TaskTypeList"`

	// 状态列表（如成功 2，正在执行 1），可选
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`

	// 任务名称
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceList []*InstanceList `json:"InstanceList,omitnil,omitempty" name:"InstanceList"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeInstanceLogDetailRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 服务器Ip
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 文件Name
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`

	// 起始行
	StartCount *int64 `json:"StartCount,omitnil,omitempty" name:"StartCount"`

	// 每次查询行数
	LineCount *int64 `json:"LineCount,omitnil,omitempty" name:"LineCount"`
}

type DescribeInstanceLogDetailRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 服务器Ip
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 文件Name
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`

	// 起始行
	StartCount *int64 `json:"StartCount,omitnil,omitempty" name:"StartCount"`

	// 每次查询行数
	LineCount *int64 `json:"LineCount,omitnil,omitempty" name:"LineCount"`
}

func (r *DescribeInstanceLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	delete(f, "BrokerIp")
	delete(f, "OriginFileName")
	delete(f, "StartCount")
	delete(f, "LineCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogDetailResponseParams struct {
	// 日志结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *InstanceLogInfoOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogDetailResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogFileRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 执行机IP
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 日志文件名
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`
}

type DescribeInstanceLogFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 执行机IP
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 日志文件名
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`
}

func (r *DescribeInstanceLogFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	delete(f, "BrokerIp")
	delete(f, "OriginFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceLogFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogFileResponseParams struct {
	// 下载文件详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *InstanceDownloadLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceLogFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceLogFileResponseParams `json:"Response"`
}

func (r *DescribeInstanceLogFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceLogFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceLogListRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
}

type DescribeInstanceLogListRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLogList []*InstanceLogList `json:"InstanceLogList,omitnil,omitempty" name:"InstanceLogList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 服务器Ip
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 文件Name
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`
}

type DescribeInstanceLogRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 服务器Ip
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 文件Name
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`
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
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLogInfo *IntegrationInstanceLog `json:"InstanceLogInfo,omitnil,omitempty" name:"InstanceLogInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeIntegrationNodeRequestParams struct {
	// 节点id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DescribeIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 节点id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型，201为实时任务，202为离线任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// 上游节点是否已经修改。true 已修改，需要提示；false 没有修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceCheckFlag *bool `json:"SourceCheckFlag,omitnil,omitempty" name:"SourceCheckFlag"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeIntegrationStatisticsInstanceTrendRequestParams struct {
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`
}

type DescribeIntegrationStatisticsInstanceTrendRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`
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
	TrendsData []*IntegrationStatisticsTrendResult `json:"TrendsData,omitnil,omitempty" name:"TrendsData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`
}

type DescribeIntegrationStatisticsRecordsTrendRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`
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
	TrendsData []*IntegrationStatisticsTrendResult `json:"TrendsData,omitnil,omitempty" name:"TrendsData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`
}

type DescribeIntegrationStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`
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
	TotalTask *int64 `json:"TotalTask,omitnil,omitempty" name:"TotalTask"`

	// 生产态任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProdTask *int64 `json:"ProdTask,omitnil,omitempty" name:"ProdTask"`

	// 开发态任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DevTask *int64 `json:"DevTask,omitnil,omitempty" name:"DevTask"`

	// 总读取条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalReadRecords *int64 `json:"TotalReadRecords,omitnil,omitempty" name:"TotalReadRecords"`

	// 总写入条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalWriteRecords *int64 `json:"TotalWriteRecords,omitnil,omitempty" name:"TotalWriteRecords"`

	// 总脏数据条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalErrorRecords *int64 `json:"TotalErrorRecords,omitnil,omitempty" name:"TotalErrorRecords"`

	// 总告警事件数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalAlarmEvent *int64 `json:"TotalAlarmEvent,omitnil,omitempty" name:"TotalAlarmEvent"`

	// 当天读取增长条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseReadRecords *int64 `json:"IncreaseReadRecords,omitnil,omitempty" name:"IncreaseReadRecords"`

	// 当天写入增长条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseWriteRecords *int64 `json:"IncreaseWriteRecords,omitnil,omitempty" name:"IncreaseWriteRecords"`

	// 当天脏数据增长条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseErrorRecords *int64 `json:"IncreaseErrorRecords,omitnil,omitempty" name:"IncreaseErrorRecords"`

	// 当天告警事件增长数
	// 注意：此字段可能返回 null，表示取不到有效值。
	IncreaseAlarmEvent *int64 `json:"IncreaseAlarmEvent,omitnil,omitempty" name:"IncreaseAlarmEvent"`

	// 告警事件统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmEvent *string `json:"AlarmEvent,omitnil,omitempty" name:"AlarmEvent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`
}

type DescribeIntegrationStatisticsTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`
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
	StatusData *string `json:"StatusData,omitnil,omitempty" name:"StatusData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`
}

type DescribeIntegrationStatisticsTaskStatusTrendRequest struct {
	*tchttp.BaseRequest
	
	// 任务类型（实时：201，离线：202）
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 查询日期
	QueryDate *string `json:"QueryDate,omitnil,omitempty" name:"QueryDate"`

	// 资源组id
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`
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
	TrendsData []*IntegrationStatisticsTrendResult `json:"TrendsData,omitnil,omitempty" name:"TrendsData"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型，201: 实时集成任务,   202：离线集成任务，不传默认值为201 实时任务类型
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 提交版本号
	InstanceVersion *int64 `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`
}

type DescribeIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型，201: 实时集成任务,   202：离线集成任务，不传默认值为201 实时任务类型
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 提交版本号
	InstanceVersion *int64 `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`
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
	delete(f, "InstanceVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeIntegrationTaskResponseParams struct {
	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 采集器统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentStatus *AgentStatus `json:"AgentStatus,omitnil,omitempty" name:"AgentStatus"`

	// 任务版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskVersion *TaskVersionInstance `json:"TaskVersion,omitnil,omitempty" name:"TaskVersion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页第n页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询filter;默认查询任务的开发态，如需查询生产态任务需添加{"Values":["true"],"Name":"ProductionState"};如需查询查询任务状态需要查询生产态任务列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段信息
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 201. 实时同步, 202. 离线同步 默认实时
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DescribeIntegrationTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页第n页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询filter;默认查询任务的开发态，如需查询生产态任务需添加{"Values":["true"],"Name":"ProductionState"};如需查询查询任务状态需要查询生产态任务列表
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段信息
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 201. 实时同步, 202. 离线同步 默认实时
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	TaskInfoSet []*IntegrationTaskInfo `json:"TaskInfoSet,omitnil,omitempty" name:"TaskInfoSet"`

	// 任务总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 该任务选定版本的存储路径：
	// DescribeDsTaskVersionList 或者 DescribeDsTaskVersionInfo 返回的对应的 TaskInfo.TaskExt.Properties 下 Base64.encode($region | $bucket | $ftp.file.name) 值
	TaskVersionPath *string `json:"TaskVersionPath,omitnil,omitempty" name:"TaskVersionPath"`

	// 该任务选定版本id：
	// DescribeDsTaskVersionList 或者 DescribeDsTaskVersionInfo 返回的对应的 VersionId 取值
	TaskVersion *string `json:"TaskVersion,omitnil,omitempty" name:"TaskVersion"`
}

type DescribeIntegrationVersionNodesInfoRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 该任务选定版本的存储路径：
	// DescribeDsTaskVersionList 或者 DescribeDsTaskVersionInfo 返回的对应的 TaskInfo.TaskExt.Properties 下 Base64.encode($region | $bucket | $ftp.file.name) 值
	TaskVersionPath *string `json:"TaskVersionPath,omitnil,omitempty" name:"TaskVersionPath"`

	// 该任务选定版本id：
	// DescribeDsTaskVersionList 或者 DescribeDsTaskVersionInfo 返回的对应的 VersionId 取值
	TaskVersion *string `json:"TaskVersion,omitnil,omitempty" name:"TaskVersion"`
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
	Nodes []*IntegrationNodeInfo `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 任务映射信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mappings []*IntegrationNodeMapping `json:"Mappings,omitnil,omitempty" name:"Mappings"`

	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 长连接临时token。与Token相同含义，优先取Data，Data为空时，取Token。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeOperateOpsTasksRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹id，多个文件夹以逗号分隔
	FolderIdList *string `json:"FolderIdList,omitnil,omitempty" name:"FolderIdList"`

	// 工作流id，多个工作流id之间以英文字符逗号分隔
	WorkFlowIdList *string `json:"WorkFlowIdList,omitnil,omitempty" name:"WorkFlowIdList"`

	// 工作流名称，多个工作流名称之间以英文字符逗号分隔
	WorkFlowNameList *string `json:"WorkFlowNameList,omitnil,omitempty" name:"WorkFlowNameList"`

	// 任务名称，多个任务名称之间以英文字符逗号分隔
	TaskNameList *string `json:"TaskNameList,omitnil,omitempty" name:"TaskNameList"`

	// 任务id，多个任务id之间以英文字符逗号分隔
	TaskIdList *string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 页号
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序字段，支持字段为FirstSubmitTime和FirstRunTime，标识最近提交和首次执行事件
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序类型。两种取值 DESC、ASC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 责任人，多个责任人之间以英文字符逗号分隔
	InChargeList *string `json:"InChargeList,omitnil,omitempty" name:"InChargeList"`

	// 任务类型Id字符串，多个任务类型id之间以英文字符逗号分隔
	TaskTypeIdList *string `json:"TaskTypeIdList,omitnil,omitempty" name:"TaskTypeIdList"`

	// 任务状态字符串，多个任务状态之间以英文字符逗号分隔
	StatusList *string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// 任务周期类型字符串，多个任务周期之间以英文字符逗号分隔
	TaskCycleUnitList *string `json:"TaskCycleUnitList,omitnil,omitempty" name:"TaskCycleUnitList"`

	// 任务所属产品类型
	ProductNameList *string `json:"ProductNameList,omitnil,omitempty" name:"ProductNameList"`

	// 数据源id或（仅针对离线同步任务）来源数据源id
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 数据源类型或（仅针对离线同步任务）来源数据源类型
	SourceServiceType *string `json:"SourceServiceType,omitnil,omitempty" name:"SourceServiceType"`

	// （仅针对离线同步任务）目标数据源id
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// （仅针对离线同步任务）目标数据源类型
	TargetServiceType *string `json:"TargetServiceType,omitnil,omitempty" name:"TargetServiceType"`

	// 告警类型，多个类型以逗号分隔
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// 资源组id,多个资源组id之间以英文字符逗号分隔
	ExecutorGroupIdList *string `json:"ExecutorGroupIdList,omitnil,omitempty" name:"ExecutorGroupIdList"`

	// 任务标签
	TaskTags []*TaskTag `json:"TaskTags,omitnil,omitempty" name:"TaskTags"`

	// 查询关键字
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 实例生成方式
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`

	// 额外请求的资源类型
	RequestResourceTypes []*string `json:"RequestResourceTypes,omitnil,omitempty" name:"RequestResourceTypes"`
}

type DescribeOperateOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹id，多个文件夹以逗号分隔
	FolderIdList *string `json:"FolderIdList,omitnil,omitempty" name:"FolderIdList"`

	// 工作流id，多个工作流id之间以英文字符逗号分隔
	WorkFlowIdList *string `json:"WorkFlowIdList,omitnil,omitempty" name:"WorkFlowIdList"`

	// 工作流名称，多个工作流名称之间以英文字符逗号分隔
	WorkFlowNameList *string `json:"WorkFlowNameList,omitnil,omitempty" name:"WorkFlowNameList"`

	// 任务名称，多个任务名称之间以英文字符逗号分隔
	TaskNameList *string `json:"TaskNameList,omitnil,omitempty" name:"TaskNameList"`

	// 任务id，多个任务id之间以英文字符逗号分隔
	TaskIdList *string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 页号
	PageNumber *string `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *string `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序字段，支持字段为FirstSubmitTime和FirstRunTime，标识最近提交和首次执行事件
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序类型。两种取值 DESC、ASC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`

	// 责任人，多个责任人之间以英文字符逗号分隔
	InChargeList *string `json:"InChargeList,omitnil,omitempty" name:"InChargeList"`

	// 任务类型Id字符串，多个任务类型id之间以英文字符逗号分隔
	TaskTypeIdList *string `json:"TaskTypeIdList,omitnil,omitempty" name:"TaskTypeIdList"`

	// 任务状态字符串，多个任务状态之间以英文字符逗号分隔
	StatusList *string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// 任务周期类型字符串，多个任务周期之间以英文字符逗号分隔
	TaskCycleUnitList *string `json:"TaskCycleUnitList,omitnil,omitempty" name:"TaskCycleUnitList"`

	// 任务所属产品类型
	ProductNameList *string `json:"ProductNameList,omitnil,omitempty" name:"ProductNameList"`

	// 数据源id或（仅针对离线同步任务）来源数据源id
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 数据源类型或（仅针对离线同步任务）来源数据源类型
	SourceServiceType *string `json:"SourceServiceType,omitnil,omitempty" name:"SourceServiceType"`

	// （仅针对离线同步任务）目标数据源id
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// （仅针对离线同步任务）目标数据源类型
	TargetServiceType *string `json:"TargetServiceType,omitnil,omitempty" name:"TargetServiceType"`

	// 告警类型，多个类型以逗号分隔
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// 资源组id,多个资源组id之间以英文字符逗号分隔
	ExecutorGroupIdList *string `json:"ExecutorGroupIdList,omitnil,omitempty" name:"ExecutorGroupIdList"`

	// 任务标签
	TaskTags []*TaskTag `json:"TaskTags,omitnil,omitempty" name:"TaskTags"`

	// 查询关键字
	KeyWord *string `json:"KeyWord,omitnil,omitempty" name:"KeyWord"`

	// 实例生成方式
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`

	// 额外请求的资源类型
	RequestResourceTypes []*string `json:"RequestResourceTypes,omitnil,omitempty" name:"RequestResourceTypes"`
}

func (r *DescribeOperateOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOperateOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderIdList")
	delete(f, "WorkFlowIdList")
	delete(f, "WorkFlowNameList")
	delete(f, "TaskNameList")
	delete(f, "TaskIdList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortItem")
	delete(f, "SortType")
	delete(f, "InChargeList")
	delete(f, "TaskTypeIdList")
	delete(f, "StatusList")
	delete(f, "TaskCycleUnitList")
	delete(f, "ProductNameList")
	delete(f, "SourceServiceId")
	delete(f, "SourceServiceType")
	delete(f, "TargetServiceId")
	delete(f, "TargetServiceType")
	delete(f, "AlarmType")
	delete(f, "ExecutorGroupIdList")
	delete(f, "TaskTags")
	delete(f, "KeyWord")
	delete(f, "InitStrategy")
	delete(f, "RequestResourceTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOperateOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOperateOpsTasksResponseParams struct {
	// 任务列表信息
	Data *OpsTaskInfoPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOperateOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOperateOpsTasksResponseParams `json:"Response"`
}

func (r *DescribeOperateOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOperateOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsInstanceLogListRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
}

type DescribeOpsInstanceLogListRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
}

func (r *DescribeOpsInstanceLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsInstanceLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpsInstanceLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsInstanceLogListResponseParams struct {
	// 实例日志列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*InstanceLogInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOpsInstanceLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpsInstanceLogListResponseParams `json:"Response"`
}

func (r *DescribeOpsInstanceLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsInstanceLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsMakePlanInstancesRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 补录任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页页码，默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认值10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 实例状态列表
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`
}

type DescribeOpsMakePlanInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 补录任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 分页页码，默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认值10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 实例状态列表
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`
}

func (r *DescribeOpsMakePlanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsMakePlanInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PlanId")
	delete(f, "TaskId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "StateList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpsMakePlanInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsMakePlanInstancesResponseParams struct {
	// 补录计划实例分页查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *MakePlanInstanceOpsDtoCollection `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOpsMakePlanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpsMakePlanInstancesResponseParams `json:"Response"`
}

func (r *DescribeOpsMakePlanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsMakePlanInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsMakePlanTasksRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 分页页码，默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认值10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeOpsMakePlanTasksRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 分页页码，默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认值10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeOpsMakePlanTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsMakePlanTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PlanId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpsMakePlanTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsMakePlanTasksResponseParams struct {
	// 补录计划任务分页查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *MakePlanTaskOpsDtoCollection `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOpsMakePlanTasksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpsMakePlanTasksResponseParams `json:"Response"`
}

func (r *DescribeOpsMakePlanTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsMakePlanTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsMakePlansRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页数，默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认值10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 补录计划名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 补录任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 补录任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 补录计划创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 补录计划最小创建时间
	MinCreateTime *string `json:"MinCreateTime,omitnil,omitempty" name:"MinCreateTime"`

	// 补录计划最大创建时间
	MaxCreateTime *string `json:"MaxCreateTime,omitnil,omitempty" name:"MaxCreateTime"`
}

type DescribeOpsMakePlansRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页数，默认值1
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，默认值10
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 补录计划名称
	PlanName *string `json:"PlanName,omitnil,omitempty" name:"PlanName"`

	// 补录任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 补录任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 补录计划创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 补录计划最小创建时间
	MinCreateTime *string `json:"MinCreateTime,omitnil,omitempty" name:"MinCreateTime"`

	// 补录计划最大创建时间
	MaxCreateTime *string `json:"MaxCreateTime,omitnil,omitempty" name:"MaxCreateTime"`
}

func (r *DescribeOpsMakePlansRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsMakePlansRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "PlanId")
	delete(f, "PlanName")
	delete(f, "TaskName")
	delete(f, "TaskId")
	delete(f, "Creator")
	delete(f, "MinCreateTime")
	delete(f, "MaxCreateTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpsMakePlansRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsMakePlansResponseParams struct {
	// 补录计划分页查询结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *MakePlanOpsDtoCollection `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOpsMakePlansResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpsMakePlansResponseParams `json:"Response"`
}

func (r *DescribeOpsMakePlansResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsMakePlansResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsWorkflowsRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务产品类型名称列表，以 ',' 号分割
	ProductNameList *string `json:"ProductNameList,omitnil,omitempty" name:"ProductNameList"`

	// 文件id列表，以 ',' 号分割
	FolderIdList *string `json:"FolderIdList,omitnil,omitempty" name:"FolderIdList"`

	// 工作流id，以 ',' 号分割
	WorkFlowIdList *string `json:"WorkFlowIdList,omitnil,omitempty" name:"WorkFlowIdList"`

	// 工作流名称列表，以 ',' 号分割
	WorkFlowNameList *string `json:"WorkFlowNameList,omitnil,omitempty" name:"WorkFlowNameList"`

	// 任务名称列表，以 ',' 号分割
	TaskNameList *string `json:"TaskNameList,omitnil,omitempty" name:"TaskNameList"`

	// 任务id列表，以 ',' 号分割
	TaskIdList *string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 状态列表，以 ',' 号分割
	StatusList *string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// 负责人列表，以 ',' 号分割
	InChargeList *string `json:"InChargeList,omitnil,omitempty" name:"InChargeList"`

	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序项
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序方式，DESC或ASC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type DescribeOpsWorkflowsRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务产品类型名称列表，以 ',' 号分割
	ProductNameList *string `json:"ProductNameList,omitnil,omitempty" name:"ProductNameList"`

	// 文件id列表，以 ',' 号分割
	FolderIdList *string `json:"FolderIdList,omitnil,omitempty" name:"FolderIdList"`

	// 工作流id，以 ',' 号分割
	WorkFlowIdList *string `json:"WorkFlowIdList,omitnil,omitempty" name:"WorkFlowIdList"`

	// 工作流名称列表，以 ',' 号分割
	WorkFlowNameList *string `json:"WorkFlowNameList,omitnil,omitempty" name:"WorkFlowNameList"`

	// 任务名称列表，以 ',' 号分割
	TaskNameList *string `json:"TaskNameList,omitnil,omitempty" name:"TaskNameList"`

	// 任务id列表，以 ',' 号分割
	TaskIdList *string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 状态列表，以 ',' 号分割
	StatusList *string `json:"StatusList,omitnil,omitempty" name:"StatusList"`

	// 负责人列表，以 ',' 号分割
	InChargeList *string `json:"InChargeList,omitnil,omitempty" name:"InChargeList"`

	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 排序项
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 排序方式，DESC或ASC
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

func (r *DescribeOpsWorkflowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsWorkflowsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "ProductNameList")
	delete(f, "FolderIdList")
	delete(f, "WorkFlowIdList")
	delete(f, "WorkFlowNameList")
	delete(f, "TaskNameList")
	delete(f, "TaskIdList")
	delete(f, "StatusList")
	delete(f, "InChargeList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "SortItem")
	delete(f, "SortType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOpsWorkflowsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOpsWorkflowsResponseParams struct {
	// 工作流列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkflowExtOpsDtoPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeOpsWorkflowsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOpsWorkflowsResponseParams `json:"Response"`
}

func (r *DescribeOpsWorkflowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOpsWorkflowsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationalFunctionsRequestParams struct {
	// 场景类型：开发、使用
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 项目 ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 函数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 标准模式开发环境：DEV
	// 标准模式生产环境：PROD
	// 简单模式：ALL
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 过滤条件
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件
	OrderFields *OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

type DescribeOrganizationalFunctionsRequest struct {
	*tchttp.BaseRequest
	
	// 场景类型：开发、使用
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 项目 ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 函数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 标准模式开发环境：DEV
	// 标准模式生产环境：PROD
	// 简单模式：ALL
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`

	// 过滤条件
	Filters *Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件
	OrderFields *OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
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
	delete(f, "EnvType")
	delete(f, "Filters")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOrganizationalFunctionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOrganizationalFunctionsResponseParams struct {
	// 函数信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content []*OrganizationalFunction `json:"Content,omitnil,omitempty" name:"Content"`

	// 操作失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type DescribePendingSubmitTaskInfo struct {
	// 任务编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 修改类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyType *string `json:"ModifyType,omitnil,omitempty" name:"ModifyType"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 提交预检查（Y/N）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitPreCheck *string `json:"SubmitPreCheck,omitnil,omitempty" name:"SubmitPreCheck"`

	// 提交预检查提交可能会失败的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitPreCheckDetailList []*TaskSubmitPreCheckDetailInfo `json:"SubmitPreCheckDetailList,omitnil,omitempty" name:"SubmitPreCheckDetailList"`

	// 资源组编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`
}

// Predefined struct for user
type DescribePendingSubmitTaskListRequestParams struct {
	// 项目编号
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流编号
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务编号列表
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`
}

type DescribePendingSubmitTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 项目编号
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流编号
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务编号列表
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`
}

func (r *DescribePendingSubmitTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePendingSubmitTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "TaskIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePendingSubmitTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePendingSubmitTaskListResponseParams struct {
	// 待提交任务信息
	Data []*DescribePendingSubmitTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePendingSubmitTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePendingSubmitTaskListResponseParams `json:"Response"`
}

func (r *DescribePendingSubmitTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePendingSubmitTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectRequestParams struct {
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否展示关联集群信息
	DescribeClusters *bool `json:"DescribeClusters,omitnil,omitempty" name:"DescribeClusters"`

	// 是否展示关联执行组的信息，仅部分信息。
	DescribeExecutors *bool `json:"DescribeExecutors,omitnil,omitempty" name:"DescribeExecutors"`

	// 默认不展示项目管理员信息
	DescribeAdminUsers *bool `json:"DescribeAdminUsers,omitnil,omitempty" name:"DescribeAdminUsers"`

	// 默认不统计项目人员数量
	DescribeMemberCount *bool `json:"DescribeMemberCount,omitnil,omitempty" name:"DescribeMemberCount"`

	// 默认不查询创建者的信息
	DescribeCreator *bool `json:"DescribeCreator,omitnil,omitempty" name:"DescribeCreator"`

	// 项目名只在租户内唯一，一般用来转化为项目ID。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
}

type DescribeProjectRequest struct {
	*tchttp.BaseRequest
	
	// 项目id。一般使用项目Id来查询，与projectName必须存在一个。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否展示关联集群信息
	DescribeClusters *bool `json:"DescribeClusters,omitnil,omitempty" name:"DescribeClusters"`

	// 是否展示关联执行组的信息，仅部分信息。
	DescribeExecutors *bool `json:"DescribeExecutors,omitnil,omitempty" name:"DescribeExecutors"`

	// 默认不展示项目管理员信息
	DescribeAdminUsers *bool `json:"DescribeAdminUsers,omitnil,omitempty" name:"DescribeAdminUsers"`

	// 默认不统计项目人员数量
	DescribeMemberCount *bool `json:"DescribeMemberCount,omitnil,omitempty" name:"DescribeMemberCount"`

	// 默认不查询创建者的信息
	DescribeCreator *bool `json:"DescribeCreator,omitnil,omitempty" name:"DescribeCreator"`

	// 项目名只在租户内唯一，一般用来转化为项目ID。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`
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
	// 项目信息
	Data *Project `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeProjectUsersRequestParams struct {
	// 分页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 【过滤参数】自定义条件查询
	Filters []*FilterOptional `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 【排序参数】排序字段
	OrderFields []*OrderFieldOptional `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 是否项目管理员
	IsProjectAdmin *bool `json:"IsProjectAdmin,omitnil,omitempty" name:"IsProjectAdmin"`
}

type DescribeProjectUsersRequest struct {
	*tchttp.BaseRequest
	
	// 分页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 【过滤参数】自定义条件查询
	Filters []*FilterOptional `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 【排序参数】排序字段
	OrderFields []*OrderFieldOptional `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 是否项目管理员
	IsProjectAdmin *bool `json:"IsProjectAdmin,omitnil,omitempty" name:"IsProjectAdmin"`
}

func (r *DescribeProjectUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "IsProjectAdmin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeProjectUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeProjectUsersResponseParams struct {
	// 项目列表
	Data *ProjectUsersPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeProjectUsersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeProjectUsersResponseParams `json:"Response"`
}

func (r *DescribeProjectUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeProjectUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityScoreRequestParams struct {
	// 统计日期
	StatisticsDate *int64 `json:"StatisticsDate,omitnil,omitempty" name:"StatisticsDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeQualityScoreRequest struct {
	*tchttp.BaseRequest
	
	// 统计日期
	StatisticsDate *int64 `json:"StatisticsDate,omitnil,omitempty" name:"StatisticsDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeQualityScoreRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityScoreRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StatisticsDate")
	delete(f, "ProjectId")
	delete(f, "DatasourceId")
	delete(f, "ScoreType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQualityScoreRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityScoreResponseParams struct {
	// 质量评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *QualityScore `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQualityScoreResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQualityScoreResponseParams `json:"Response"`
}

func (r *DescribeQualityScoreResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityScoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityScoreTrendRequestParams struct {
	// 统计开始日期
	StatisticsStartDate *int64 `json:"StatisticsStartDate,omitnil,omitempty" name:"StatisticsStartDate"`

	// 统计结束日期
	StatisticsEndDate *int64 `json:"StatisticsEndDate,omitnil,omitempty" name:"StatisticsEndDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeQualityScoreTrendRequest struct {
	*tchttp.BaseRequest
	
	// 统计开始日期
	StatisticsStartDate *int64 `json:"StatisticsStartDate,omitnil,omitempty" name:"StatisticsStartDate"`

	// 统计结束日期
	StatisticsEndDate *int64 `json:"StatisticsEndDate,omitnil,omitempty" name:"StatisticsEndDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeQualityScoreTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityScoreTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StatisticsStartDate")
	delete(f, "StatisticsEndDate")
	delete(f, "ProjectId")
	delete(f, "DatasourceId")
	delete(f, "ScoreType")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQualityScoreTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQualityScoreTrendResponseParams struct {
	// 质量评分趋势视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *QualityScoreTrend `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQualityScoreTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQualityScoreTrendResponseParams `json:"Response"`
}

func (r *DescribeQualityScoreTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQualityScoreTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskInstanceNodeInfoRequestParams struct {
	// 实时任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRealTimeTaskInstanceNodeInfoRequest struct {
	*tchttp.BaseRequest
	
	// 实时任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	RealTimeTaskInstanceNodeInfo *RealTimeTaskInstanceNodeInfo `json:"RealTimeTaskInstanceNodeInfo,omitnil,omitempty" name:"RealTimeTaskInstanceNodeInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 要查看的实时任务的任务ID，可在任务列表页面中获得
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 要查看的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeRealTimeTaskMetricOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 要查看的实时任务的任务ID，可在任务列表页面中获得
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 要查看的项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`
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
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealTimeTaskMetricOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealTimeTaskMetricOverviewResponseParams struct {
	// 总读取记录数
	TotalRecordNumOfRead *uint64 `json:"TotalRecordNumOfRead,omitnil,omitempty" name:"TotalRecordNumOfRead"`

	// 总读取字节数
	TotalRecordByteNumOfRead *uint64 `json:"TotalRecordByteNumOfRead,omitnil,omitempty" name:"TotalRecordByteNumOfRead"`

	// 总写入记录数
	TotalRecordNumOfWrite *uint64 `json:"TotalRecordNumOfWrite,omitnil,omitempty" name:"TotalRecordNumOfWrite"`

	// 总写入字节数 单位字节
	TotalRecordByteNumOfWrite *uint64 `json:"TotalRecordByteNumOfWrite,omitnil,omitempty" name:"TotalRecordByteNumOfWrite"`

	// 总的脏记录数据
	TotalDirtyRecordNum *uint64 `json:"TotalDirtyRecordNum,omitnil,omitempty" name:"TotalDirtyRecordNum"`

	// 总的脏字节数 单位字节
	TotalDirtyRecordByte *uint64 `json:"TotalDirtyRecordByte,omitnil,omitempty" name:"TotalDirtyRecordByte"`

	// 运行时长 单位s
	TotalDuration *uint64 `json:"TotalDuration,omitnil,omitempty" name:"TotalDuration"`

	// 开始运行时间
	BeginRunTime *string `json:"BeginRunTime,omitnil,omitempty" name:"BeginRunTime"`

	// 目前运行到的时间
	EndRunTime *string `json:"EndRunTime,omitnil,omitempty" name:"EndRunTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 带毫秒的时间戳
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 带毫秒的时间戳
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 粒度，1或者5
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRealTimeTaskSpeedRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 带毫秒的时间戳
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 带毫秒的时间戳
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 粒度，1或者5
	Granularity *uint64 `json:"Granularity,omitnil,omitempty" name:"Granularity"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	RecordsSpeedList []*RecordsSpeed `json:"RecordsSpeedList,omitnil,omitempty" name:"RecordsSpeedList"`

	// 同步速度字节/s列表
	BytesSpeedList []*BytesSpeed `json:"BytesSpeedList,omitnil,omitempty" name:"BytesSpeedList"`

	// 同步速度，包括了RecordsSpeedList和BytesSpeedList
	Data *RealTimeTaskSpeed `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeResourceManagePathTreesRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 名字，供搜索
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 文件夹类型
	// personal 个人
	// project 项目
	// resource 资源
	DirType *string `json:"DirType,omitnil,omitempty" name:"DirType"`
}

type DescribeResourceManagePathTreesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 名字，供搜索
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文件路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 文件夹类型
	// personal 个人
	// project 项目
	// resource 资源
	DirType *string `json:"DirType,omitnil,omitempty" name:"DirType"`
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
	Data []*ResourcePathTree `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeRuleDimStatRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeRuleDimStatRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeRuleDimStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDimStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleDimStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleDimStatResponseParams struct {
	// 结果
	Data *RuleDimStat `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleDimStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleDimStatResponseParams `json:"Response"`
}

func (r *DescribeRuleDimStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleDimStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecDetailRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则执行id
	RuleExecId *uint64 `json:"RuleExecId,omitnil,omitempty" name:"RuleExecId"`
}

type DescribeRuleExecDetailRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则执行id
	RuleExecId *uint64 `json:"RuleExecId,omitnil,omitempty" name:"RuleExecId"`
}

func (r *DescribeRuleExecDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RuleExecId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleExecDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecDetailResponseParams struct {
	// 规则执行结果详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleExecResultDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleExecDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleExecDetailResponseParams `json:"Response"`
}

func (r *DescribeRuleExecDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecLogRequestParams struct {
	// 规则执行Id
	RuleExecId *uint64 `json:"RuleExecId,omitnil,omitempty" name:"RuleExecId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组执行id
	RuleGroupExecId *uint64 `json:"RuleGroupExecId,omitnil,omitempty" name:"RuleGroupExecId"`
}

type DescribeRuleExecLogRequest struct {
	*tchttp.BaseRequest
	
	// 规则执行Id
	RuleExecId *uint64 `json:"RuleExecId,omitnil,omitempty" name:"RuleExecId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组执行id
	RuleGroupExecId *uint64 `json:"RuleGroupExecId,omitnil,omitempty" name:"RuleGroupExecId"`
}

func (r *DescribeRuleExecLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleExecId")
	delete(f, "ProjectId")
	delete(f, "RuleGroupExecId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleExecLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecLogResponseParams struct {
	// 规则执行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleExecLog `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleExecLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleExecLogResponseParams `json:"Response"`
}

func (r *DescribeRuleExecLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecResultsRequestParams struct {
	// 规则组执行Id
	RuleGroupExecId *uint64 `json:"RuleGroupExecId,omitnil,omitempty" name:"RuleGroupExecId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRuleExecResultsRequest struct {
	*tchttp.BaseRequest
	
	// 规则组执行Id
	RuleGroupExecId *uint64 `json:"RuleGroupExecId,omitnil,omitempty" name:"RuleGroupExecId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeRuleExecResultsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecResultsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleGroupExecId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleExecResultsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecResultsResponseParams struct {
	// 规则执行结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleExecResultPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleExecResultsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleExecResultsResponseParams `json:"Response"`
}

func (r *DescribeRuleExecResultsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecResultsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecStatRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeRuleExecStatRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeRuleExecStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleExecStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleExecStatResponseParams struct {
	// 结果
	Data *RuleExecStat `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleExecStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleExecStatResponseParams `json:"Response"`
}

func (r *DescribeRuleExecStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleExecStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupExecResultsByPageRequestParams struct {
	// 分页序号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRuleGroupExecResultsByPageRequest struct {
	*tchttp.BaseRequest
	
	// 分页序号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeRuleGroupExecResultsByPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupExecResultsByPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleGroupExecResultsByPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupExecResultsByPageResponseParams struct {
	// 规则组执行结果列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleGroupExecResultPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleGroupExecResultsByPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleGroupExecResultsByPageResponseParams `json:"Response"`
}

func (r *DescribeRuleGroupExecResultsByPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupExecResultsByPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupRequestParams struct {
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 数据来源ID
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据表Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库ID
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`
}

type DescribeRuleGroupRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 数据来源ID
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据表Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库ID
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`
}

func (r *DescribeRuleGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleGroupId")
	delete(f, "DatasourceId")
	delete(f, "TableId")
	delete(f, "ProjectId")
	delete(f, "DatabaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupResponseParams struct {
	// 数据质量规则组详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleGroup `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleGroupResponseParams `json:"Response"`
}

func (r *DescribeRuleGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupSubscriptionRequestParams struct {
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRuleGroupSubscriptionRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeRuleGroupSubscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupSubscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleGroupId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleGroupSubscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupSubscriptionResponseParams struct {
	// 规则组订阅信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleGroupSubscribe `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleGroupSubscriptionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleGroupSubscriptionResponseParams `json:"Response"`
}

func (r *DescribeRuleGroupSubscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupTableRequestParams struct {
	// 表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`
}

type DescribeRuleGroupTableRequest struct {
	*tchttp.BaseRequest
	
	// 表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`
}

func (r *DescribeRuleGroupTableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupTableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleGroupTableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupTableResponseParams struct {
	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleGroupTable `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleGroupTableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleGroupTableResponseParams `json:"Response"`
}

func (r *DescribeRuleGroupTableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupsByPageRequestParams struct {
	// 分页序号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件,每次请求的Filters的上限为10，Filter.Values的上限为5
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRuleGroupsByPageRequest struct {
	*tchttp.BaseRequest
	
	// 分页序号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件,每次请求的Filters的上限为10，Filter.Values的上限为5
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeRuleGroupsByPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupsByPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleGroupsByPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleGroupsByPageResponseParams struct {
	// 规则组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleGroupPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleGroupsByPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleGroupsByPageResponseParams `json:"Response"`
}

func (r *DescribeRuleGroupsByPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleGroupsByPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleRequestParams struct {
	// 质量规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRuleRequest struct {
	*tchttp.BaseRequest
	
	// 质量规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleResponseParams struct {
	// 规则详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *Rule `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleResponseParams `json:"Response"`
}

func (r *DescribeRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleTemplateRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则模板Id
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

type DescribeRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则模板Id
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`
}

func (r *DescribeRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TemplateId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleTemplateResponseParams struct {
	// 模板详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RuleTemplate `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleTemplateResponseParams `json:"Response"`
}

func (r *DescribeRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleTemplatesByPageRequestParams struct {
	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 工作空间ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 通用排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 通用过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeRuleTemplatesByPageRequest struct {
	*tchttp.BaseRequest
	
	// 当前页
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 工作空间ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 通用排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 通用过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeRuleTemplatesByPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleTemplatesByPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "ProjectId")
	delete(f, "OrderFields")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleTemplatesByPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleTemplatesByPageResponseParams struct {
	// 结果
	Data *RuleTemplatePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleTemplatesByPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleTemplatesByPageResponseParams `json:"Response"`
}

func (r *DescribeRuleTemplatesByPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleTemplatesByPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleTemplatesRequestParams struct {
	// 模板类型 1.系统模板 2.自定义模板
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 1.常量 2.离线表级 2.离线字段级
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 源端对应的引擎类型
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`
}

type DescribeRuleTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// 模板类型 1.系统模板 2.自定义模板
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 1.常量 2.离线表级 2.离线字段级
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 源端对应的引擎类型
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`
}

func (r *DescribeRuleTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "SourceObjectType")
	delete(f, "ProjectId")
	delete(f, "SourceEngineTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRuleTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRuleTemplatesResponseParams struct {
	// 规则模板列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*RuleTemplate `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRuleTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRuleTemplatesResponseParams `json:"Response"`
}

func (r *DescribeRuleTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRuleTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesByPageRequestParams struct {
	// 分页序号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeRulesByPageRequest struct {
	*tchttp.BaseRequest
	
	// 分页序号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeRulesByPageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesByPageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesByPageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesByPageResponseParams struct {
	// 规则质量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *RulePage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRulesByPageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesByPageResponseParams `json:"Response"`
}

func (r *DescribeRulesByPageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesByPageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组id
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 该规则运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

type DescribeRulesRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组id
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 该规则运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`
}

func (r *DescribeRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RuleGroupId")
	delete(f, "EngineType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRulesResponseParams struct {
	// 规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*Rule `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRulesResponseParams `json:"Response"`
}

func (r *DescribeRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduleInstancesRequestParams struct {
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

type DescribeScheduleInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

func (r *DescribeScheduleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduleInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	delete(f, "SearchCondition")
	delete(f, "OptType")
	delete(f, "OperatorName")
	delete(f, "OperatorId")
	delete(f, "ProjectId")
	delete(f, "ProjectIdent")
	delete(f, "ProjectName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Count")
	delete(f, "RequestBaseInfo")
	delete(f, "IsCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScheduleInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduleInstancesResponseParams struct {
	// 实例结果集
	Data *CollectionInstanceOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScheduleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScheduleInstancesResponseParams `json:"Response"`
}

func (r *DescribeScheduleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerInstanceStatusRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型ID
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 执行资源组ID
	ExecutionGroupId *string `json:"ExecutionGroupId,omitnil,omitempty" name:"ExecutionGroupId"`

	// 执行资源组名字
	ExecutionGroupName *string `json:"ExecutionGroupName,omitnil,omitempty" name:"ExecutionGroupName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DescribeSchedulerInstanceStatusRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型ID
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 执行资源组ID
	ExecutionGroupId *string `json:"ExecutionGroupId,omitnil,omitempty" name:"ExecutionGroupId"`

	// 执行资源组名字
	ExecutionGroupName *string `json:"ExecutionGroupName,omitnil,omitempty" name:"ExecutionGroupName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DescribeSchedulerInstanceStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerInstanceStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskTypeId")
	delete(f, "ExecutionGroupId")
	delete(f, "ExecutionGroupName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "InCharge")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSchedulerInstanceStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerInstanceStatusResponseParams struct {
	// 响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ScreenInstanceInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSchedulerInstanceStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSchedulerInstanceStatusResponseParams `json:"Response"`
}

func (r *DescribeSchedulerInstanceStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerInstanceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerRunTimeInstanceCntByStatusRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 周期类型
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 时间单元 eg: 12h
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 开始日期：2023-03-02
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日前：2023-03-20
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 排序字段
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 升序降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type DescribeSchedulerRunTimeInstanceCntByStatusRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 周期类型
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 时间单元 eg: 12h
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 开始日期：2023-03-02
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日前：2023-03-20
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 任务类型
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 排序字段
	SortItem *string `json:"SortItem,omitnil,omitempty" name:"SortItem"`

	// 升序降序
	SortType *string `json:"SortType,omitnil,omitempty" name:"SortType"`
}

func (r *DescribeSchedulerRunTimeInstanceCntByStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerRunTimeInstanceCntByStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "CycleUnit")
	delete(f, "TimeUnit")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskType")
	delete(f, "InCharge")
	delete(f, "WorkflowId")
	delete(f, "SortItem")
	delete(f, "SortType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSchedulerRunTimeInstanceCntByStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerRunTimeInstanceCntByStatusResponseParams struct {
	// 响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*RuntimeInstanceCntTop `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSchedulerRunTimeInstanceCntByStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSchedulerRunTimeInstanceCntByStatusResponseParams `json:"Response"`
}

func (r *DescribeSchedulerRunTimeInstanceCntByStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerRunTimeInstanceCntByStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerTaskCntByStatusRequestParams struct {
	// 1
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Y
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 111
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 1
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DescribeSchedulerTaskCntByStatusRequest struct {
	*tchttp.BaseRequest
	
	// 1
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// Y
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 111
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 1
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DescribeSchedulerTaskCntByStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerTaskCntByStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskType")
	delete(f, "TypeName")
	delete(f, "ProjectId")
	delete(f, "InCharge")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSchedulerTaskCntByStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerTaskCntByStatusResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ScreenTaskInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSchedulerTaskCntByStatusResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSchedulerTaskCntByStatusResponseParams `json:"Response"`
}

func (r *DescribeSchedulerTaskCntByStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerTaskCntByStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerTaskTypeCntRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`
}

type DescribeSchedulerTaskTypeCntRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`
}

func (r *DescribeSchedulerTaskTypeCntRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerTaskTypeCntRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InCharge")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSchedulerTaskTypeCntRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSchedulerTaskTypeCntResponseParams struct {
	// data
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskTypeCnt `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSchedulerTaskTypeCntResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSchedulerTaskTypeCntResponseParams `json:"Response"`
}

func (r *DescribeSchedulerTaskTypeCntResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSchedulerTaskTypeCntResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStatisticInstanceStatusTrendOpsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型Id
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 时间类型
	TimeType *string `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 任务类型名称
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 资源组ID
	ExecutionGroupId *string `json:"ExecutionGroupId,omitnil,omitempty" name:"ExecutionGroupId"`

	// 资源组名称
	ExecutionGroupName *string `json:"ExecutionGroupName,omitnil,omitempty" name:"ExecutionGroupName"`

	// 1
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 1
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 1
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`

	// D代表天，H代表小时
	AggregationUnit *string `json:"AggregationUnit,omitnil,omitempty" name:"AggregationUnit"`

	// 1
	AverageWindowSize *int64 `json:"AverageWindowSize,omitnil,omitempty" name:"AverageWindowSize"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DescribeStatisticInstanceStatusTrendOpsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型Id
	TaskTypeId *string `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 时间类型
	TimeType *string `json:"TimeType,omitnil,omitempty" name:"TimeType"`

	// 任务类型名称
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 资源组ID
	ExecutionGroupId *string `json:"ExecutionGroupId,omitnil,omitempty" name:"ExecutionGroupId"`

	// 资源组名称
	ExecutionGroupName *string `json:"ExecutionGroupName,omitnil,omitempty" name:"ExecutionGroupName"`

	// 1
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 1
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 1
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`

	// D代表天，H代表小时
	AggregationUnit *string `json:"AggregationUnit,omitnil,omitempty" name:"AggregationUnit"`

	// 1
	AverageWindowSize *int64 `json:"AverageWindowSize,omitnil,omitempty" name:"AverageWindowSize"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DescribeStatisticInstanceStatusTrendOpsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticInstanceStatusTrendOpsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "TaskTypeId")
	delete(f, "TimeType")
	delete(f, "TypeName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ExecutionGroupId")
	delete(f, "ExecutionGroupName")
	delete(f, "InCharge")
	delete(f, "TaskType")
	delete(f, "StateList")
	delete(f, "AggregationUnit")
	delete(f, "AverageWindowSize")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStatisticInstanceStatusTrendOpsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStatisticInstanceStatusTrendOpsResponseParams struct {
	// 实例状态统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*InstanceStatisticInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStatisticInstanceStatusTrendOpsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStatisticInstanceStatusTrendOpsResponseParams `json:"Response"`
}

func (r *DescribeStatisticInstanceStatusTrendOpsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStatisticInstanceStatusTrendOpsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamTaskLogListRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// container名字
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序类型 desc asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 作业运行的实例ID
	RunningOrderId *uint64 `json:"RunningOrderId,omitnil,omitempty" name:"RunningOrderId"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeStreamTaskLogListRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 作业ID
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// container名字
	Container *string `json:"Container,omitnil,omitempty" name:"Container"`

	// 条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 排序类型 desc asc
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 作业运行的实例ID
	RunningOrderId *uint64 `json:"RunningOrderId,omitnil,omitempty" name:"RunningOrderId"`

	// 关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
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
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStreamTaskLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStreamTaskLogListResponseParams struct {
	// 是否是全量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ListOver *bool `json:"ListOver,omitnil,omitempty" name:"ListOver"`

	// 日志集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogContentList []*LogContentInfo `json:"LogContentList,omitnil,omitempty" name:"LogContentList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeSuccessorOpsTaskInfosRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeSuccessorOpsTaskInfosRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeSuccessorOpsTaskInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuccessorOpsTaskInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSuccessorOpsTaskInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSuccessorOpsTaskInfosResponseParams struct {
	// 下游任务列表
	Data []*TaskOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSuccessorOpsTaskInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSuccessorOpsTaskInfosResponseParams `json:"Response"`
}

func (r *DescribeSuccessorOpsTaskInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSuccessorOpsTaskInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableBasicInfoRequestParams struct {
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

type DescribeTableBasicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

func (r *DescribeTableBasicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableBasicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableBasicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableBasicInfoResponseParams struct {
	// 表元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableBasicInfoList []*TableBasicInfo `json:"TableBasicInfoList,omitnil,omitempty" name:"TableBasicInfoList"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableBasicInfoResponseParams `json:"Response"`
}

func (r *DescribeTableBasicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableInfoListRequestParams struct {
	// 表名
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 如果是hive这里写rpc，如果是其他类型不传
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 数据库源类型
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`
}

type DescribeTableInfoListRequest struct {
	*tchttp.BaseRequest
	
	// 表名
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 如果是hive这里写rpc，如果是其他类型不传
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 数据库源类型
	Catalog *string `json:"Catalog,omitnil,omitempty" name:"Catalog"`
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
	TableInfo []*TableInfo `json:"TableInfo,omitnil,omitempty" name:"TableInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTableLineageInfoRequestParams struct {
	// 查询方向，INPUT,OUTPUT,BOTH枚举值
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 表信息
	Data *TableLineageInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 单次查询入度,默认 1
	InputDepth *int64 `json:"InputDepth,omitnil,omitempty" name:"InputDepth"`

	// 单次查询出度,默认 1
	OutputDepth *int64 `json:"OutputDepth,omitnil,omitempty" name:"OutputDepth"`

	// 额外参数（传递调用方信息）
	ExtParams []*LineageParamRecord `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 是否过滤临时表,默认true
	IgnoreTemp *bool `json:"IgnoreTemp,omitnil,omitempty" name:"IgnoreTemp"`

	// 是否递归查询二级节点数目，默认为true
	RecursiveSecond *bool `json:"RecursiveSecond,omitnil,omitempty" name:"RecursiveSecond"`
}

type DescribeTableLineageInfoRequest struct {
	*tchttp.BaseRequest
	
	// 查询方向，INPUT,OUTPUT,BOTH枚举值
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 表信息
	Data *TableLineageInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 单次查询入度,默认 1
	InputDepth *int64 `json:"InputDepth,omitnil,omitempty" name:"InputDepth"`

	// 单次查询出度,默认 1
	OutputDepth *int64 `json:"OutputDepth,omitnil,omitempty" name:"OutputDepth"`

	// 额外参数（传递调用方信息）
	ExtParams []*LineageParamRecord `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 是否过滤临时表,默认true
	IgnoreTemp *bool `json:"IgnoreTemp,omitnil,omitempty" name:"IgnoreTemp"`

	// 是否递归查询二级节点数目，默认为true
	RecursiveSecond *bool `json:"RecursiveSecond,omitnil,omitempty" name:"RecursiveSecond"`
}

func (r *DescribeTableLineageInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableLineageInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Data")
	delete(f, "InputDepth")
	delete(f, "OutputDepth")
	delete(f, "ExtParams")
	delete(f, "IgnoreTemp")
	delete(f, "RecursiveSecond")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableLineageInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableLineageInfoResponseParams struct {
	// 表血缘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableLineageBasicInfo *TableLineageBaseInfo `json:"TableLineageBasicInfo,omitnil,omitempty" name:"TableLineageBasicInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableLineageInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableLineageInfoResponseParams `json:"Response"`
}

func (r *DescribeTableLineageInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableLineageInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableLineageRequestParams struct {
	// 查询方向，INPUT,OUTPUT,BOTH枚举值
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 表信息
	Data *TableLineageInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 单次查询入度,默认 1
	InputDepth *int64 `json:"InputDepth,omitnil,omitempty" name:"InputDepth"`

	// 单次查询出度,默认 1
	OutputDepth *int64 `json:"OutputDepth,omitnil,omitempty" name:"OutputDepth"`

	// 额外参数（传递调用方信息）
	ExtParams []*LineageParamRecord `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 是否过滤临时表,默认true
	IgnoreTemp *bool `json:"IgnoreTemp,omitnil,omitempty" name:"IgnoreTemp"`

	// 是否递归查询二级节点数目，默认为true
	RecursiveSecond *bool `json:"RecursiveSecond,omitnil,omitempty" name:"RecursiveSecond"`
}

type DescribeTableLineageRequest struct {
	*tchttp.BaseRequest
	
	// 查询方向，INPUT,OUTPUT,BOTH枚举值
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`

	// 表信息
	Data *TableLineageInfo `json:"Data,omitnil,omitempty" name:"Data"`

	// 单次查询入度,默认 1
	InputDepth *int64 `json:"InputDepth,omitnil,omitempty" name:"InputDepth"`

	// 单次查询出度,默认 1
	OutputDepth *int64 `json:"OutputDepth,omitnil,omitempty" name:"OutputDepth"`

	// 额外参数（传递调用方信息）
	ExtParams []*LineageParamRecord `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 是否过滤临时表,默认true
	IgnoreTemp *bool `json:"IgnoreTemp,omitnil,omitempty" name:"IgnoreTemp"`

	// 是否递归查询二级节点数目，默认为true
	RecursiveSecond *bool `json:"RecursiveSecond,omitnil,omitempty" name:"RecursiveSecond"`
}

func (r *DescribeTableLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Direction")
	delete(f, "Data")
	delete(f, "InputDepth")
	delete(f, "OutputDepth")
	delete(f, "ExtParams")
	delete(f, "IgnoreTemp")
	delete(f, "RecursiveSecond")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableLineageResponseParams struct {
	// 表血缘信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableLineage *TableLineageInfo `json:"TableLineage,omitnil,omitempty" name:"TableLineage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableLineageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableLineageResponseParams `json:"Response"`
}

func (r *DescribeTableLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableMetaRequestParams struct {
	// 表唯一id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 按名称查询的条件
	TableNameFilter *TableNameFilter `json:"TableNameFilter,omitnil,omitempty" name:"TableNameFilter"`

	// 查询条件类型0按id，1按名称，默认为0
	TableFilterType *uint64 `json:"TableFilterType,omitnil,omitempty" name:"TableFilterType"`

	// 查询字段列表
	SearchNames []*string `json:"SearchNames,omitnil,omitempty" name:"SearchNames"`
}

type DescribeTableMetaRequest struct {
	*tchttp.BaseRequest
	
	// 表唯一id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 按名称查询的条件
	TableNameFilter *TableNameFilter `json:"TableNameFilter,omitnil,omitempty" name:"TableNameFilter"`

	// 查询条件类型0按id，1按名称，默认为0
	TableFilterType *uint64 `json:"TableFilterType,omitnil,omitempty" name:"TableFilterType"`

	// 查询字段列表
	SearchNames []*string `json:"SearchNames,omitnil,omitempty" name:"SearchNames"`
}

func (r *DescribeTableMetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableMetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TableId")
	delete(f, "TableNameFilter")
	delete(f, "TableFilterType")
	delete(f, "SearchNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableMetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableMetaResponseParams struct {
	// 表的元数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMeta *TableMeta `json:"TableMeta,omitnil,omitempty" name:"TableMeta"`

	// 生命周期信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecycleInfo *LifecycleInfo `json:"LifecycleInfo,omitnil,omitempty" name:"LifecycleInfo"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagVoteSumList *TagVoteSum `json:"TagVoteSumList,omitnil,omitempty" name:"TagVoteSumList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableMetaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableMetaResponseParams `json:"Response"`
}

func (r *DescribeTableMetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableMetasRequestParams struct {
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

type DescribeTableMetasRequest struct {
	*tchttp.BaseRequest
	
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤字段
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序字段
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`
}

func (r *DescribeTableMetasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableMetasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableMetasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableMetasResponseParams struct {
	// 表元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableMetas []*TableMeta `json:"TableMetas,omitnil,omitempty" name:"TableMetas"`

	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableMetasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableMetasResponseParams `json:"Response"`
}

func (r *DescribeTableMetasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableMetasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablePartitionsRequestParams struct {
	// 表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 分页number
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页size
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤器
	FilterSet []*Filter `json:"FilterSet,omitnil,omitempty" name:"FilterSet"`

	// 排序字段
	OrderFieldSet []*OrderField `json:"OrderFieldSet,omitnil,omitempty" name:"OrderFieldSet"`
}

type DescribeTablePartitionsRequest struct {
	*tchttp.BaseRequest
	
	// 表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 分页number
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页size
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤器
	FilterSet []*Filter `json:"FilterSet,omitnil,omitempty" name:"FilterSet"`

	// 排序字段
	OrderFieldSet []*OrderField `json:"OrderFieldSet,omitnil,omitempty" name:"OrderFieldSet"`
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
	delete(f, "TableId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "FilterSet")
	delete(f, "OrderFieldSet")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTablePartitionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTablePartitionsResponseParams struct {
	// 分区详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TablePartitionSet []*TablePartition `json:"TablePartitionSet,omitnil,omitempty" name:"TablePartitionSet"`

	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

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
type DescribeTableQualityDetailsRequestParams struct {
	// 统计日期
	StatisticsDate *int64 `json:"StatisticsDate,omitnil,omitempty" name:"StatisticsDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页序号
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤参数TableName、DatabaseId 、DatabaseName、OwnerUserName
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序参数 排序方式 DESC 或者 ASC，表得分排序 TableScore
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`
}

type DescribeTableQualityDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 统计日期
	StatisticsDate *int64 `json:"StatisticsDate,omitnil,omitempty" name:"StatisticsDate"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 分页序号
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤参数TableName、DatabaseId 、DatabaseName、OwnerUserName
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序参数 排序方式 DESC 或者 ASC，表得分排序 TableScore
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 数据来源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`
}

func (r *DescribeTableQualityDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableQualityDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StatisticsDate")
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "DatasourceId")
	delete(f, "ScoreType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableQualityDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableQualityDetailsResponseParams struct {
	// 表质量分详情结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *TableQualityDetailPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableQualityDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableQualityDetailsResponseParams `json:"Response"`
}

func (r *DescribeTableQualityDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableQualityDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableSchemaInfoRequestParams struct {
	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据源的类型（例如MYSQL、HIVE、KAFKA等）
	MsType *string `json:"MsType,omitnil,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 连接类型（示例值rpc）
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 元数据Database下的Schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 项目空间ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 环境信息
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 空间模式
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 开发态的datasourceId
	DevDatasourceId *string `json:"DevDatasourceId,omitnil,omitempty" name:"DevDatasourceId"`
}

type DescribeTableSchemaInfoRequest struct {
	*tchttp.BaseRequest
	
	// 表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据源的类型（例如MYSQL、HIVE、KAFKA等）
	MsType *string `json:"MsType,omitnil,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 连接类型（示例值rpc）
	ConnectionType *string `json:"ConnectionType,omitnil,omitempty" name:"ConnectionType"`

	// 元数据Database下的Schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 项目空间ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 环境信息
	Env *string `json:"Env,omitnil,omitempty" name:"Env"`

	// 空间模式
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`

	// 开发态的datasourceId
	DevDatasourceId *string `json:"DevDatasourceId,omitnil,omitempty" name:"DevDatasourceId"`
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
	delete(f, "ProjectId")
	delete(f, "Env")
	delete(f, "Model")
	delete(f, "DevDatasourceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableSchemaInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableSchemaInfoResponseParams struct {
	// 查询到的SchemaInfo信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaInfoList []*SchemaDetail `json:"SchemaInfoList,omitnil,omitempty" name:"SchemaInfoList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTableScoreTrendRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间 秒级时间戳
	StatisticsStartDate *int64 `json:"StatisticsStartDate,omitnil,omitempty" name:"StatisticsStartDate"`

	// 结束时间 秒级时间戳
	StatisticsEndDate *int64 `json:"StatisticsEndDate,omitnil,omitempty" name:"StatisticsEndDate"`

	// 表id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`
}

type DescribeTableScoreTrendRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间 秒级时间戳
	StatisticsStartDate *int64 `json:"StatisticsStartDate,omitnil,omitempty" name:"StatisticsStartDate"`

	// 结束时间 秒级时间戳
	StatisticsEndDate *int64 `json:"StatisticsEndDate,omitnil,omitempty" name:"StatisticsEndDate"`

	// 表id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 1:按全维度权重计算,2:按已配置维度权重计算,3:不按维度权重计算,默认1
	ScoreType *string `json:"ScoreType,omitnil,omitempty" name:"ScoreType"`
}

func (r *DescribeTableScoreTrendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableScoreTrendRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "StatisticsStartDate")
	delete(f, "StatisticsEndDate")
	delete(f, "TableId")
	delete(f, "ScoreType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTableScoreTrendRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTableScoreTrendResponseParams struct {
	// 表得分趋势
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *QualityScoreTrend `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTableScoreTrendResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTableScoreTrendResponseParams `json:"Response"`
}

func (r *DescribeTableScoreTrendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTableScoreTrendResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskAlarmRegulationsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 当前页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件(name有RegularStatus、AlarmLevel、AlarmIndicator、RegularName)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件(RegularId)
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型(201代表实时任务，202代表离线任务)
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DescribeTaskAlarmRegulationsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 当前页
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页记录数
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件(name有RegularStatus、AlarmLevel、AlarmIndicator、RegularName)
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序条件(RegularId)
	OrderFields []*OrderField `json:"OrderFields,omitnil,omitempty" name:"OrderFields"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型(201代表实时任务，202代表离线任务)
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	delete(f, "ProjectId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Filters")
	delete(f, "OrderFields")
	delete(f, "TaskId")
	delete(f, "TaskType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskAlarmRegulationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskAlarmRegulationsResponseParams struct {
	// 告警规则信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAlarmInfos []*TaskAlarmInfo `json:"TaskAlarmInfos,omitnil,omitempty" name:"TaskAlarmInfos"`

	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTaskByCycleReportRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务周期类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeTaskByCycleReportRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务周期类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeTaskByCycleReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskByCycleReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Type")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskByCycleReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskByCycleReportResponseParams struct {
	// 任务周期增长趋势统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskByStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskByCycleReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskByCycleReportResponseParams `json:"Response"`
}

func (r *DescribeTaskByCycleReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskByCycleReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskByCycleRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 1
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DescribeTaskByCycleRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 1
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DescribeTaskByCycleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskByCycleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "InCharge")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskByCycleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskByCycleResponseParams struct {
	// 周期任务统计值
	Data []*TaskByCycle `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskByCycleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskByCycleResponseParams `json:"Response"`
}

func (r *DescribeTaskByCycleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskByCycleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskByStatusReportRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 时间类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 类型名称
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合单元，H小时
	AggregationUnit *string `json:"AggregationUnit,omitnil,omitempty" name:"AggregationUnit"`

	// 周期
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DescribeTaskByStatusReportRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 时间类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 类型名称
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合单元，H小时
	AggregationUnit *string `json:"AggregationUnit,omitnil,omitempty" name:"AggregationUnit"`

	// 周期
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DescribeTaskByStatusReportRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskByStatusReportRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "Type")
	delete(f, "TaskType")
	delete(f, "TypeName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AggregationUnit")
	delete(f, "CycleUnit")
	delete(f, "Status")
	delete(f, "InCharge")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskByStatusReportRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskByStatusReportResponseParams struct {
	// 任务上报趋势指标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*TaskByStatus `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskByStatusReportResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskByStatusReportResponseParams `json:"Response"`
}

func (r *DescribeTaskByStatusReportResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskByStatusReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLineageRequestParams struct {
	// 请求来源，WEB 前端；CLIENT 客户端
	RequestFromSource *string `json:"RequestFromSource,omitnil,omitempty" name:"RequestFromSource"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskLineageRequest struct {
	*tchttp.BaseRequest
	
	// 请求来源，WEB 前端；CLIENT 客户端
	RequestFromSource *string `json:"RequestFromSource,omitnil,omitempty" name:"RequestFromSource"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DescribeTaskLineageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLineageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RequestFromSource")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskLineageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLineageResponseParams struct {
	// 请求来源，WEB 前端；CLIENT 客户端
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestFromSource *string `json:"RequestFromSource,omitnil,omitempty" name:"RequestFromSource"`

	// 通过任务ID查询集成任务信息列表
	TaskLineageInfos []*TaskLineageInfoPair `json:"TaskLineageInfos,omitnil,omitempty" name:"TaskLineageInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskLineageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskLineageResponseParams `json:"Response"`
}

func (r *DescribeTaskLineageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskLineageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskLockStatusRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type DescribeTaskLockStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	TaskLockStatus *TaskLockStatus `json:"TaskLockStatus,omitnil,omitempty" name:"TaskLockStatus"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTaskRunHistoryRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 检索条件
	SearchCondition *InstanceSearchCondition `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeTaskRunHistoryRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 检索条件
	SearchCondition *InstanceSearchCondition `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeTaskRunHistoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRunHistoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "SearchCondition")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTaskRunHistoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskRunHistoryResponseParams struct {
	// 分页查询任务运行历史结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *InstanceOpsInfoPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTaskRunHistoryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTaskRunHistoryResponseParams `json:"Response"`
}

func (r *DescribeTaskRunHistoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTaskRunHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTaskScriptRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DescribeTaskScriptRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
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
	Data *TaskScriptContent `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeTemplateDimCountRequestParams struct {
	// 模板类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeTemplateDimCountRequest struct {
	*tchttp.BaseRequest
	
	// 模板类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeTemplateDimCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateDimCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTemplateDimCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTemplateDimCountResponseParams struct {
	// 维度统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DimensionCount `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTemplateDimCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTemplateDimCountResponseParams `json:"Response"`
}

func (r *DescribeTemplateDimCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTemplateDimCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdTaskRunLogRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
}

type DescribeThirdTaskRunLogRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
}

func (r *DescribeThirdTaskRunLogRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdTaskRunLogRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "CurRunDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeThirdTaskRunLogRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeThirdTaskRunLogResponseParams struct {
	// 获取第三方运行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeThirdTaskRunLogResponse struct {
	*tchttp.BaseResponse
	Response *DescribeThirdTaskRunLogResponseParams `json:"Response"`
}

func (r *DescribeThirdTaskRunLogResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeThirdTaskRunLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopTableStatRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeTopTableStatRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeTopTableStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopTableStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopTableStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopTableStatResponseParams struct {
	// 结果
	Data *TopTableStat `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTopTableStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopTableStatResponseParams `json:"Response"`
}

func (r *DescribeTopTableStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopTableStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrendStatRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

type DescribeTrendStatRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 开始时间，时间戳到秒
	BeginDate *string `json:"BeginDate,omitnil,omitempty" name:"BeginDate"`

	// 结束时间，时间戳到秒
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`
}

func (r *DescribeTrendStatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrendStatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "BeginDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTrendStatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTrendStatResponseParams struct {
	// 结果
	Data []*RuleExecDateStat `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTrendStatResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTrendStatResponseParams `json:"Response"`
}

func (r *DescribeTrendStatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTrendStatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowCanvasInfoRequestParams struct {
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeWorkflowCanvasInfoRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeWorkflowCanvasInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowCanvasInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowCanvasInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowCanvasInfoResponseParams struct {
	// 工作流调度详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkflowCanvasOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowCanvasInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowCanvasInfoResponseParams `json:"Response"`
}

func (r *DescribeWorkflowCanvasInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowCanvasInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowExecuteByIdRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkFlowIdList *string `json:"WorkFlowIdList,omitnil,omitempty" name:"WorkFlowIdList"`

	// 分页大小
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页索引
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeWorkflowExecuteByIdRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkFlowIdList *string `json:"WorkFlowIdList,omitnil,omitempty" name:"WorkFlowIdList"`

	// 分页大小
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页索引
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeWorkflowExecuteByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowExecuteByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkFlowIdList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowExecuteByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowExecuteByIdResponseParams struct {
	// 工作流运行时间信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkFlowExecuteDtoByPage `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowExecuteByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowExecuteByIdResponseParams `json:"Response"`
}

func (r *DescribeWorkflowExecuteByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowExecuteByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowInfoByIdRequestParams struct {
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeWorkflowInfoByIdRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeWorkflowInfoByIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowInfoByIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowInfoByIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowInfoByIdResponseParams struct {
	// 工作流调度详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkflowSchedulerOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowInfoByIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowInfoByIdResponseParams `json:"Response"`
}

func (r *DescribeWorkflowInfoByIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowInfoByIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowListByProjectIdRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeWorkflowListByProjectIdRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeWorkflowListByProjectIdRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowListByProjectIdRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowListByProjectIdRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowListByProjectIdResponseParams struct {
	// 根据项目id获取项目下所有工作流列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*WorkflowCanvasOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowListByProjectIdResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowListByProjectIdResponseParams `json:"Response"`
}

func (r *DescribeWorkflowListByProjectIdResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowListByProjectIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowSchedulerInfoDsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

type DescribeWorkflowSchedulerInfoDsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

func (r *DescribeWorkflowSchedulerInfoDsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowSchedulerInfoDsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowSchedulerInfoDsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowSchedulerInfoDsResponseParams struct {
	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkflowScheduleDtoDs `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowSchedulerInfoDsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowSchedulerInfoDsResponseParams `json:"Response"`
}

func (r *DescribeWorkflowSchedulerInfoDsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowSchedulerInfoDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowTaskCountRequestParams struct {
	// 工作流列表
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type DescribeWorkflowTaskCountRequest struct {
	*tchttp.BaseRequest
	
	// 工作流列表
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *DescribeWorkflowTaskCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowTaskCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowId")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowTaskCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowTaskCountResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *WorkflowTaskCountOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowTaskCountResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowTaskCountResponseParams `json:"Response"`
}

func (r *DescribeWorkflowTaskCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowTaskCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DiagnoseProRequestParams struct {
	// 查询条件（当前接口TaskId和CurRunDate需要填充在该字段方可诊断）
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

type DiagnoseProRequest struct {
	*tchttp.BaseRequest
	
	// 查询条件（当前接口TaskId和CurRunDate需要填充在该字段方可诊断）
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

func (r *DiagnoseProRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DiagnoseProRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SearchCondition")
	delete(f, "ProjectId")
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	delete(f, "OptType")
	delete(f, "OperatorName")
	delete(f, "OperatorId")
	delete(f, "ProjectIdent")
	delete(f, "ProjectName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Count")
	delete(f, "RequestBaseInfo")
	delete(f, "IsCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DiagnoseProRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DiagnoseProResponseParams struct {
	// 结果
	Data *DiagnoseRep `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DiagnoseProResponse struct {
	*tchttp.BaseResponse
	Response *DiagnoseProResponseParams `json:"Response"`
}

func (r *DiagnoseProResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DiagnoseProResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DiagnoseRep struct {
	// 诊断信息内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *Content `json:"Content,omitnil,omitempty" name:"Content"`

	// 诊断结果相关信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Table *Table `json:"Table,omitnil,omitempty" name:"Table"`
}

type DimensionCount struct {
	// 维度类型1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: DimType is deprecated.
	DimType *uint64 `json:"DimType,omitnil,omitempty" name:"DimType"`

	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 维度类型1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`
}

type DimensionScore struct {
	// 维度评分列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DimensionScoreList []*DimensionScoreInfo `json:"DimensionScoreList,omitnil,omitempty" name:"DimensionScoreList"`
}

type DimensionScoreInfo struct {
	// 维度名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *float64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 设置人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *int64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 设置人名
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 更新时间 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 参与评估表数量
	JoinTableNumber *int64 `json:"JoinTableNumber,omitnil,omitempty" name:"JoinTableNumber"`

	// 评分
	Score *float64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 设置人idStr
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIdStr *string `json:"UserIdStr,omitnil,omitempty" name:"UserIdStr"`
}

type DlcDataGovernPolicy struct {
	// 数据排布治理项
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteDataPolicy *DlcRewriteDataInfo `json:"RewriteDataPolicy,omitnil,omitempty" name:"RewriteDataPolicy"`

	// 快照过期治理项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredSnapshotsPolicy *DlcExpiredSnapshotsInfo `json:"ExpiredSnapshotsPolicy,omitnil,omitempty" name:"ExpiredSnapshotsPolicy"`

	// 移除孤立文件治理项
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoveOrphanFilesPolicy *DlcRemoveOrphanFilesInfo `json:"RemoveOrphanFilesPolicy,omitnil,omitempty" name:"RemoveOrphanFilesPolicy"`

	// 合并元数据Manifests治理项
	// 注意：此字段可能返回 null，表示取不到有效值。
	MergeManifestsPolicy *DlcMergeManifestsInfo `json:"MergeManifestsPolicy,omitnil,omitempty" name:"MergeManifestsPolicy"`

	// 是否集成库规则：default（默认继承）、none（不继承）
	// 注意：此字段可能返回 null，表示取不到有效值。
	InheritDataBase *string `json:"InheritDataBase,omitnil,omitempty" name:"InheritDataBase"`

	// 治理规则类型，Customize: 自定义；Intelligence: 智能治理
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 治理引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernEngine *string `json:"GovernEngine,omitnil,omitempty" name:"GovernEngine"`
}

type DlcExpiredSnapshotsInfo struct {
	// 是否启用快照过期治理项：enable、none
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredSnapshotsEnable *string `json:"ExpiredSnapshotsEnable,omitnil,omitempty" name:"ExpiredSnapshotsEnable"`

	// 用于运行快照过期治理项的引擎名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 需要保留的最近快照个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetainLast *uint64 `json:"RetainLast,omitnil,omitempty" name:"RetainLast"`

	// 过期指定天前的快照
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeforeDays *uint64 `json:"BeforeDays,omitnil,omitempty" name:"BeforeDays"`

	// 清理过期快照的并行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrentDeletes *uint64 `json:"MaxConcurrentDeletes,omitnil,omitempty" name:"MaxConcurrentDeletes"`

	// 快照过期治理运行周期，单位为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalMin *uint64 `json:"IntervalMin,omitnil,omitempty" name:"IntervalMin"`
}

type DlcMergeManifestsInfo struct {
	// 是否启用合并元数据Manifests文件治理项：enable、none
	// 注意：此字段可能返回 null，表示取不到有效值。
	MergeManifestsEnable *string `json:"MergeManifestsEnable,omitnil,omitempty" name:"MergeManifestsEnable"`

	// 用于运行合并元数据Manifests文件治理项的引擎名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 合并元数据Manifests文件治理运行周期，单位为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalMin *uint64 `json:"IntervalMin,omitnil,omitempty" name:"IntervalMin"`
}

type DlcRemoveOrphanFilesInfo struct {
	// 是否启用移除孤立文件治理项：enable、none
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemoveOrphanFilesEnable *string `json:"RemoveOrphanFilesEnable,omitnil,omitempty" name:"RemoveOrphanFilesEnable"`

	// 用于运行移除孤立文件治理项的引擎名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 移除指定天前的孤立文件
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeforeDays *uint64 `json:"BeforeDays,omitnil,omitempty" name:"BeforeDays"`

	// 移除孤立文件的并行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxConcurrentDeletes *uint64 `json:"MaxConcurrentDeletes,omitnil,omitempty" name:"MaxConcurrentDeletes"`

	// 移除孤立文件治理运行周期，单位为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalMin *uint64 `json:"IntervalMin,omitnil,omitempty" name:"IntervalMin"`
}

type DlcRewriteDataInfo struct {
	// 是否启用数据重排布治理项：enable（启动）、disable（不启用，默认）
	// 注意：此字段可能返回 null，表示取不到有效值。
	RewriteDataEnable *string `json:"RewriteDataEnable,omitnil,omitempty" name:"RewriteDataEnable"`

	// 用于运行数据重排布治理项的引擎名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Engine *string `json:"Engine,omitnil,omitempty" name:"Engine"`

	// 重排布任务执行的文件个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinInputFiles *uint64 `json:"MinInputFiles,omitnil,omitempty" name:"MinInputFiles"`

	// 数据重排布写后的数据文件大小，单位为字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetFileSizeBytes *uint64 `json:"TargetFileSizeBytes,omitnil,omitempty" name:"TargetFileSizeBytes"`

	// 数据重排布治理运行周期，单位为分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntervalMin *uint64 `json:"IntervalMin,omitnil,omitempty" name:"IntervalMin"`
}

type DrInstanceOpsDto struct {
	// 任务来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskSource *string `json:"TaskSource,omitnil,omitempty" name:"TaskSource"`

	// 编排空间jobId
	// 注意：此字段可能返回 null，表示取不到有效值。
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 任务提交记录Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *uint64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 子任务记录id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SonRecordId *uint64 `json:"SonRecordId,omitnil,omitempty" name:"SonRecordId"`

	// 任务实例Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 编排空间为任务id, 开发空间为脚本id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 脚本cos地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`

	// 试运行内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 任务提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 运行时长(秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *string `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 试运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 编排空间为任务名称，开发空间为脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 试运行提交人
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitUserName *string `json:"SubmitUserName,omitnil,omitempty" name:"SubmitUserName"`

	// 试运行提交人userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitUserId *string `json:"SubmitUserId,omitnil,omitempty" name:"SubmitUserId"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 是否含有结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasResultSet *bool `json:"HasResultSet,omitnil,omitempty" name:"HasResultSet"`
}

type DrInstanceOpsDtoPage struct {
	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*DrInstanceOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type DryRunDIOfflineTaskRequestParams struct {
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源组Id
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 默认 27
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`
}

type DryRunDIOfflineTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源组Id
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 默认 27
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`
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
	CurrentRunDate *string `json:"CurrentRunDate,omitnil,omitempty" name:"CurrentRunDate"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务实例唯一key
	TaskInstanceKey *string `json:"TaskInstanceKey,omitnil,omitempty" name:"TaskInstanceKey"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type Duty struct {
	// 值班Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DutyScheduleId *int64 `json:"DutyScheduleId,omitnil,omitempty" name:"DutyScheduleId"`

	// 值班开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 值班结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 值班人员
	// 注意：此字段可能返回 null，表示取不到有效值。
	DutyPersons []*DutyPerson `json:"DutyPersons,omitnil,omitempty" name:"DutyPersons"`
}

type DutyPerson struct {
	// 值班人子账号id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 值班人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 值班人员主账号用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserId *string `json:"OwnerUserId,omitnil,omitempty" name:"OwnerUserId"`

	// 值班人tenantId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *int64 `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 2023-11-02 08:00:00
	// 注意：此字段可能返回 null，表示取不到有效值。
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 扩展字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type DutySchedule struct {
	// 值班表列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*DutyScheduleData `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 页号
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 总页数
	TotalPageNumber *int64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type DutyScheduleData struct {
	// 值班表Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 值班表名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`
}

type DutyScheduleDetailsInfo struct {
	// 值班日期
	Day *string `json:"Day,omitnil,omitempty" name:"Day"`

	// 值班信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duty []*Duty `json:"Duty,omitnil,omitempty" name:"Duty"`
}

type EventBatchCaseDTO struct {
	// 事件实例id
	CaseId *string `json:"CaseId,omitnil,omitempty" name:"CaseId"`

	// 事件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件触发起始时间
	StartDimension *string `json:"StartDimension,omitnil,omitempty" name:"StartDimension"`

	// 创建时间
	CreationTs *string `json:"CreationTs,omitnil,omitempty" name:"CreationTs"`

	// 消费者id
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 事件触发结束时间
	EndDimension *string `json:"EndDimension,omitnil,omitempty" name:"EndDimension"`

	// 事件周期
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`
}

type EventCaseAuditLogOptDto struct {
	// 事件实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseId *string `json:"CaseId,omitnil,omitempty" name:"CaseId"`

	// 事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件分割类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 事件广播类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// 事件实例存活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TTL *uint64 `json:"TTL,omitnil,omitempty" name:"TTL"`

	// 事件实例存活时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 事件实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 事件实例触发时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventTriggerTimestamp *string `json:"EventTriggerTimestamp,omitnil,omitempty" name:"EventTriggerTimestamp"`

	// 事件实例消费时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTimestamp *string `json:"LogTimestamp,omitnil,omitempty" name:"LogTimestamp"`

	// 事件实例描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type EventCaseAuditLogVOCollection struct {
	// 结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前页记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*EventCaseAuditLogOptDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type EventCaseConsumeLogOptDto struct {
	// 消费ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumeLogId *string `json:"ConsumeLogId,omitnil,omitempty" name:"ConsumeLogId"`

	// 事件案例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCaseId *string `json:"EventCaseId,omitnil,omitempty" name:"EventCaseId"`

	// 消费者ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// 消费时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTimestamp *string `json:"CreationTimestamp,omitnil,omitempty" name:"CreationTimestamp"`

	// 任务详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerDetail *TaskOpsDto `json:"ConsumerDetail,omitnil,omitempty" name:"ConsumerDetail"`
}

type EventCaseConsumeLogOptDtoCollection struct {
	// 结果总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 结果总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前页结果数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数目
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*EventCaseConsumeLogOptDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type EventCaseDTO struct {
	// 事件实例id
	CaseId *string `json:"CaseId,omitnil,omitempty" name:"CaseId"`

	// 事件名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件格式
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 创建时间
	CreationTs *string `json:"CreationTs,omitnil,omitempty" name:"CreationTs"`

	// 消费者id
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type EventCaseOpsDto struct {
	// 案例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CaseId *string `json:"CaseId,omitnil,omitempty" name:"CaseId"`

	// 案例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 时间格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTimestamp *string `json:"CreationTimestamp,omitnil,omitempty" name:"CreationTimestamp"`

	// 消费者id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// 描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type EventListenerOpsDto struct {
	// 事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 关键字，如果是任务，则是任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 触发方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 事件属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTimestamp *string `json:"CreationTimestamp,omitnil,omitempty" name:"CreationTimestamp"`
}

type EventOpsDto struct {
	// 事件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 事件分割类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 事件广播类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// 数据时间格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DimensionFormat *string `json:"DimensionFormat,omitnil,omitempty" name:"DimensionFormat"`

	// 存活时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeToLive *uint64 `json:"TimeToLive,omitnil,omitempty" name:"TimeToLive"`

	// 存活时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 创建时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreationTimestamp *string `json:"CreationTimestamp,omitnil,omitempty" name:"CreationTimestamp"`

	// 所属者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 监听器
	// 注意：此字段可能返回 null，表示取不到有效值。
	Listeners []*EventListenerOpsDto `json:"Listeners,omitnil,omitempty" name:"Listeners"`

	// 事件案例
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventCases []*EventCaseOpsDto `json:"EventCases,omitnil,omitempty" name:"EventCases"`
}

type ExtResourceFlagDto struct {
	// 父任务信息获取标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentTask *bool `json:"ParentTask,omitnil,omitempty" name:"ParentTask"`

	// 生产事件获取标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventListener *string `json:"EventListener,omitnil,omitempty" name:"EventListener"`

	// Dlc相关配置获取标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DlcResourceConfig *string `json:"DlcResourceConfig,omitnil,omitempty" name:"DlcResourceConfig"`

	// 脚本信息获取标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Script *string `json:"Script,omitnil,omitempty" name:"Script"`

	// 离线任务信息获取标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineSyncTask *string `json:"OfflineSyncTask,omitnil,omitempty" name:"OfflineSyncTask"`
}

type FailMessage struct {
	// 数据唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type FieldConfig struct {
	// 字段key
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldKey *string `json:"FieldKey,omitnil,omitempty" name:"FieldKey"`

	// 字段值
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldValue *string `json:"FieldValue,omitnil,omitempty" name:"FieldValue"`

	// 字段值类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldDataType *string `json:"FieldDataType,omitnil,omitempty" name:"FieldDataType"`
}

type Filter struct {
	// 过滤字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type FilterOptional struct {
	// 过滤字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type FindAllFolderRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type FindAllFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *FindAllFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FindAllFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FindAllFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FindAllFolderResponseParams struct {
	// 文件夹列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderList []*FolderDsDto `json:"FolderList,omitnil,omitempty" name:"FolderList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FindAllFolderResponse struct {
	*tchttp.BaseResponse
	Response *FindAllFolderResponseParams `json:"Response"`
}

func (r *FindAllFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FindAllFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FolderDsDto struct {
	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所属项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 父文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`

	// 工作流总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 工作流列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workflows []*WorkflowCanvasOpsDto `json:"Workflows,omitnil,omitempty" name:"Workflows"`

	// 子文件夹总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalFolders *int64 `json:"TotalFolders,omitnil,omitempty" name:"TotalFolders"`

	// 子文件夹列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Folders []*FolderDsDto `json:"Folders,omitnil,omitempty" name:"Folders"`

	// 搜索类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FindType *string `json:"FindType,omitnil,omitempty" name:"FindType"`
}

type FolderOpsDto struct {
	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所属项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 父文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`

	// 工作流总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 工作流列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workflows []*WorkflowCanvasOpsDto `json:"Workflows,omitnil,omitempty" name:"Workflows"`

	// 子文件夹总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalFolders *int64 `json:"TotalFolders,omitnil,omitempty" name:"TotalFolders"`

	// 子文件夹列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	FoldersList *string `json:"FoldersList,omitnil,omitempty" name:"FoldersList"`

	// 搜索类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FindType *string `json:"FindType,omitnil,omitempty" name:"FindType"`
}

// Predefined struct for user
type FreezeOpsTasksRequestParams struct {
	// 任务列表
	Tasks []*SimpleTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务操作是否消息通知下游任务责任人
	OperateIsInform *bool `json:"OperateIsInform,omitnil,omitempty" name:"OperateIsInform"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type FreezeOpsTasksRequest struct {
	*tchttp.BaseRequest
	
	// 任务列表
	Tasks []*SimpleTaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务操作是否消息通知下游任务责任人
	OperateIsInform *bool `json:"OperateIsInform,omitnil,omitempty" name:"OperateIsInform"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *FreezeOpsTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeOpsTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tasks")
	delete(f, "OperateIsInform")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeOpsTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeOpsTasksResponseParams struct {
	// 操作结果
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FreezeOpsTasksResponse struct {
	*tchttp.BaseResponse
	Response *FreezeOpsTasksResponseParams `json:"Response"`
}

func (r *FreezeOpsTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeOpsTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksByWorkflowIdsRequestParams struct {
	// 工作流id列表
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

type FreezeTasksByWorkflowIdsRequest struct {
	*tchttp.BaseRequest
	
	// 工作流id列表
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否终止已生成的实例
	KillInstance *bool `json:"KillInstance,omitnil,omitempty" name:"KillInstance"`
}

func (r *FreezeTasksByWorkflowIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksByWorkflowIdsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowIds")
	delete(f, "ProjectId")
	delete(f, "KillInstance")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FreezeTasksByWorkflowIdsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FreezeTasksByWorkflowIdsResponseParams struct {
	// 操作返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FreezeTasksByWorkflowIdsResponse struct {
	*tchttp.BaseResponse
	Response *FreezeTasksByWorkflowIdsResponseParams `json:"Response"`
}

func (r *FreezeTasksByWorkflowIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FreezeTasksByWorkflowIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionResource struct {
	// 资源路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 资源唯一标识
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 资源 MD5 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 默认是 hdfs
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type FunctionTypeOrKind struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 函数分类英文名
	ZhName *string `json:"ZhName,omitnil,omitempty" name:"ZhName"`

	// 函数分类中文名
	EnName *string `json:"EnName,omitnil,omitempty" name:"EnName"`
}

type FunctionVersion struct {
	// 版本号：V0 V1 V2
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 提交人 ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 变更类型：ADD、MODIFY
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 备注
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 提交时间: UTC 秒数
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 提交人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 版本内容：json string 格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

// Predefined struct for user
type GenHiveTableDDLSqlRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标数据库
	SinkDatabase *string `json:"SinkDatabase,omitnil,omitempty" name:"SinkDatabase"`

	// 节点id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 元数据类型(MYSQL、ORACLE)
	MsType *string `json:"MsType,omitnil,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 来源库名
	SourceDatabase *string `json:"SourceDatabase,omitnil,omitempty" name:"SourceDatabase"`

	// 来源表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 目标表元数据类型(HIVE、GBASE)
	SinkType *string `json:"SinkType,omitnil,omitempty" name:"SinkType"`

	// 源端schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 上游节点的字段信息
	SourceFieldInfoList []*SourceFieldInfo `json:"SourceFieldInfoList,omitnil,omitempty" name:"SourceFieldInfoList"`

	// 分区字段
	Partitions []*Partition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 建表属性
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 建表模式，0:向导模式，1:ddl
	TableMode *int64 `json:"TableMode,omitnil,omitempty" name:"TableMode"`

	// DLC表版本，v1/v2
	TableVersion *string `json:"TableVersion,omitnil,omitempty" name:"TableVersion"`

	// 是否upsert写入
	UpsertFlag *bool `json:"UpsertFlag,omitnil,omitempty" name:"UpsertFlag"`

	// 表描述信息
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// 增加的文件数量阈值, 超过值将触发小文件合并
	AddDataFiles *int64 `json:"AddDataFiles,omitnil,omitempty" name:"AddDataFiles"`

	// 增加的Equality delete数量阈值, 超过值将触发小文件合并
	AddEqualityDeletes *int64 `json:"AddEqualityDeletes,omitnil,omitempty" name:"AddEqualityDeletes"`

	// 增加的Position delete数量阈值, 超过值将触发小文件合并
	AddPositionDeletes *int64 `json:"AddPositionDeletes,omitnil,omitempty" name:"AddPositionDeletes"`

	// 增加的delete file数量阈值
	AddDeleteFiles *int64 `json:"AddDeleteFiles,omitnil,omitempty" name:"AddDeleteFiles"`

	// 下游节点数据源ID
	TargetDatasourceId *string `json:"TargetDatasourceId,omitnil,omitempty" name:"TargetDatasourceId"`

	// dlc upsert主键
	UpsertKeys []*string `json:"UpsertKeys,omitnil,omitempty" name:"UpsertKeys"`

	// dlc表治理信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 目标端schema名称
	SinkSchemaName *string `json:"SinkSchemaName,omitnil,omitempty" name:"SinkSchemaName"`
}

type GenHiveTableDDLSqlRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 目标数据库
	SinkDatabase *string `json:"SinkDatabase,omitnil,omitempty" name:"SinkDatabase"`

	// 节点id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 元数据类型(MYSQL、ORACLE)
	MsType *string `json:"MsType,omitnil,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 来源库名
	SourceDatabase *string `json:"SourceDatabase,omitnil,omitempty" name:"SourceDatabase"`

	// 来源表名
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 目标表元数据类型(HIVE、GBASE)
	SinkType *string `json:"SinkType,omitnil,omitempty" name:"SinkType"`

	// 源端schema名称
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 上游节点的字段信息
	SourceFieldInfoList []*SourceFieldInfo `json:"SourceFieldInfoList,omitnil,omitempty" name:"SourceFieldInfoList"`

	// 分区字段
	Partitions []*Partition `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// 建表属性
	Properties []*Property `json:"Properties,omitnil,omitempty" name:"Properties"`

	// 建表模式，0:向导模式，1:ddl
	TableMode *int64 `json:"TableMode,omitnil,omitempty" name:"TableMode"`

	// DLC表版本，v1/v2
	TableVersion *string `json:"TableVersion,omitnil,omitempty" name:"TableVersion"`

	// 是否upsert写入
	UpsertFlag *bool `json:"UpsertFlag,omitnil,omitempty" name:"UpsertFlag"`

	// 表描述信息
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// 增加的文件数量阈值, 超过值将触发小文件合并
	AddDataFiles *int64 `json:"AddDataFiles,omitnil,omitempty" name:"AddDataFiles"`

	// 增加的Equality delete数量阈值, 超过值将触发小文件合并
	AddEqualityDeletes *int64 `json:"AddEqualityDeletes,omitnil,omitempty" name:"AddEqualityDeletes"`

	// 增加的Position delete数量阈值, 超过值将触发小文件合并
	AddPositionDeletes *int64 `json:"AddPositionDeletes,omitnil,omitempty" name:"AddPositionDeletes"`

	// 增加的delete file数量阈值
	AddDeleteFiles *int64 `json:"AddDeleteFiles,omitnil,omitempty" name:"AddDeleteFiles"`

	// 下游节点数据源ID
	TargetDatasourceId *string `json:"TargetDatasourceId,omitnil,omitempty" name:"TargetDatasourceId"`

	// dlc upsert主键
	UpsertKeys []*string `json:"UpsertKeys,omitnil,omitempty" name:"UpsertKeys"`

	// dlc表治理信息
	TableBaseInfo *TableBaseInfo `json:"TableBaseInfo,omitnil,omitempty" name:"TableBaseInfo"`

	// 目标端schema名称
	SinkSchemaName *string `json:"SinkSchemaName,omitnil,omitempty" name:"SinkSchemaName"`
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
	delete(f, "SourceFieldInfoList")
	delete(f, "Partitions")
	delete(f, "Properties")
	delete(f, "TableMode")
	delete(f, "TableVersion")
	delete(f, "UpsertFlag")
	delete(f, "TableComment")
	delete(f, "AddDataFiles")
	delete(f, "AddEqualityDeletes")
	delete(f, "AddPositionDeletes")
	delete(f, "AddDeleteFiles")
	delete(f, "TargetDatasourceId")
	delete(f, "UpsertKeys")
	delete(f, "TableBaseInfo")
	delete(f, "SinkSchemaName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenHiveTableDDLSqlRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenHiveTableDDLSqlResponseParams struct {
	// 生成的ddl语句
	DDLSql *string `json:"DDLSql,omitnil,omitempty" name:"DDLSql"`

	// 生成的ddl语句。与DDLSql相同含义，优先取Data，如果Data为空，则取DDLSql。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 通用任务参数内容,直接作用于任务的参数。不同参数用;
	// 分割
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type GetCosTokenRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求域名
	OriginDomain *string `json:"OriginDomain,omitnil,omitempty" name:"OriginDomain"`

	// 是否需要跨域
	CrossFlag *bool `json:"CrossFlag,omitnil,omitempty" name:"CrossFlag"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 远程地址
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`
}

type GetCosTokenRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 请求域名
	OriginDomain *string `json:"OriginDomain,omitnil,omitempty" name:"OriginDomain"`

	// 是否需要跨域
	CrossFlag *bool `json:"CrossFlag,omitnil,omitempty" name:"CrossFlag"`

	// 桶名
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 远程地址
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`
}

func (r *GetCosTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCosTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "OriginDomain")
	delete(f, "CrossFlag")
	delete(f, "BucketName")
	delete(f, "RemotePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCosTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCosTokenResponseParams struct {
	// cos地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// Token信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *CosTokenResponse `json:"Token,omitnil,omitempty" name:"Token"`

	// 桶名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 终止点（针对私有云环境）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPoint *string `json:"EndPoint,omitnil,omitempty" name:"EndPoint"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetCosTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetCosTokenResponseParams `json:"Response"`
}

func (r *GetCosTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCosTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFileInfoRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件路径
	// 项目区1470575647377821696项目，f1目录下： /datastudio/project/1470575647377821696/f1/sql1234.sql 个人区： /datastudio/personal/sqlTTT.sql
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

type GetFileInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件路径
	// 项目区1470575647377821696项目，f1目录下： /datastudio/project/1470575647377821696/f1/sql1234.sql 个人区： /datastudio/personal/sqlTTT.sql
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

func (r *GetFileInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetFileInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetFileInfoResponseParams struct {
	// 当前脚本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserFileInfo *UserFileInfo `json:"UserFileInfo,omitnil,omitempty" name:"UserFileInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetFileInfoResponse struct {
	*tchttp.BaseResponse
	Response *GetFileInfoResponseParams `json:"Response"`
}

func (r *GetFileInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetFileInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetIntegrationNodeColumnSchemaRequestParams struct {
	// 字段示例（json格式）
	ColumnContent *string `json:"ColumnContent,omitnil,omitempty" name:"ColumnContent"`

	// 数据源类型 MYSQL|HIVE|KAFKA|ES|MONGODB|REST_API|SYBASE|TIDB|DORIS|DM|
	DatasourceType *string `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`
}

type GetIntegrationNodeColumnSchemaRequest struct {
	*tchttp.BaseRequest
	
	// 字段示例（json格式）
	ColumnContent *string `json:"ColumnContent,omitnil,omitempty" name:"ColumnContent"`

	// 数据源类型 MYSQL|HIVE|KAFKA|ES|MONGODB|REST_API|SYBASE|TIDB|DORIS|DM|
	DatasourceType *string `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`
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
	Schemas []*IntegrationNodeSchema `json:"Schemas,omitnil,omitempty" name:"Schemas"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchConditionNew `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`
}

type GetOfflineDIInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 第几页
	PageIndex *uint64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchConditionNew `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 实例详情
	List []*OfflineInstance `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	PageIndex *string `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchCondition `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`
}

type GetOfflineInstanceListRequest struct {
	*tchttp.BaseRequest
	
	// 第几页
	PageIndex *string `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 每页几条
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 无
	SearchCondition *SearchCondition `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`
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
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 实例详情
	List []*OfflineInstance `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type InstanceApiOpsRequest struct {
	// 单个查询条件
	Instance *InstanceOpsDto `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 排序字段，目前包含：重试次数，实例数据时间，运行耗时
	SortCol *string `json:"SortCol,omitnil,omitempty" name:"SortCol"`

	// 任务id列表
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 按照taskName模糊查询
	TaskNameList []*string `json:"TaskNameList,omitnil,omitempty" name:"TaskNameList"`

	// 文件夹列表
	FolderList []*string `json:"FolderList,omitnil,omitempty" name:"FolderList"`

	// 升序或者降序
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 实例状态列表
	StateList []*int64 `json:"StateList,omitnil,omitempty" name:"StateList"`

	// 实例类型列表
	TaskTypeList []*int64 `json:"TaskTypeList,omitnil,omitempty" name:"TaskTypeList"`

	// 周期类型
	CycleList []*string `json:"CycleList,omitnil,omitempty" name:"CycleList"`

	// 责任人
	OwnerList []*string `json:"OwnerList,omitnil,omitempty" name:"OwnerList"`

	// 数据时间
	DateFrom *string `json:"DateFrom,omitnil,omitempty" name:"DateFrom"`

	// 数据时间
	DateTo *string `json:"DateTo,omitnil,omitempty" name:"DateTo"`

	// 实例入库时间
	CreateTimeFrom *string `json:"CreateTimeFrom,omitnil,omitempty" name:"CreateTimeFrom"`

	// 实例入库时间
	CreateTimeTo *string `json:"CreateTimeTo,omitnil,omitempty" name:"CreateTimeTo"`

	//  开始执行时间
	StartFrom *string `json:"StartFrom,omitnil,omitempty" name:"StartFrom"`

	//  开始执行时间
	StartTo *string `json:"StartTo,omitnil,omitempty" name:"StartTo"`

	// 所属工作流
	WorkflowIdList []*string `json:"WorkflowIdList,omitnil,omitempty" name:"WorkflowIdList"`

	// 按照workflowName模糊查询
	WorkflowNameList []*string `json:"WorkflowNameList,omitnil,omitempty" name:"WorkflowNameList"`

	// 关键字模糊查询
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// searchColumns是搜索的字段名列表
	SearchColumns []*string `json:"SearchColumns,omitnil,omitempty" name:"SearchColumns"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 任务类型映射关系, 存储任务类型id和任务类型描述信息
	TaskTypeMap []*TaskTypeMap `json:"TaskTypeMap,omitnil,omitempty" name:"TaskTypeMap"`

	// 0 补录类型 1 周期实例 2 非周期实例
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 是否dag
	DagDeal *bool `json:"DagDeal,omitnil,omitempty" name:"DagDeal"`

	//  1 父实例 2 子实例
	DagType *string `json:"DagType,omitnil,omitempty" name:"DagType"`

	// 1 自依赖 2 任务依赖  3 所有依赖
	DagDependent *string `json:"DagDependent,omitnil,omitempty" name:"DagDependent"`

	// dag深度 默认为1，取值 1-6
	DagDepth *int64 `json:"DagDepth,omitnil,omitempty" name:"DagDepth"`

	// 租户id
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 根据当前数据时间或者是下一个数据时间查询, 默认当前数据时间
	DataTimeCycle *string `json:"DataTimeCycle,omitnil,omitempty" name:"DataTimeCycle"`

	// 资源组id,多个资源组id用英文逗号分隔
	ExecutorGroupIdList []*string `json:"ExecutorGroupIdList,omitnil,omitempty" name:"ExecutorGroupIdList"`
}

type InstanceCondition struct {
	// 执行类型
	ExecutionSpace *string `json:"ExecutionSpace,omitnil,omitempty" name:"ExecutionSpace"`

	// 任务产品类型
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`
}

type InstanceDownloadLogInfo struct {
	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`
}

type InstanceLifeCycleOpsDto struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 实例生命次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeRound *int64 `json:"LifeRound,omitnil,omitempty" name:"LifeRound"`

	// 运行类型 重跑/补录/周期/非周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunType *string `json:"RunType,omitnil,omitempty" name:"RunType"`

	// 重跑次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *int64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 实例生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLifeDetailDtoList []*InstanceLifeDetailDto `json:"InstanceLifeDetailDtoList,omitnil,omitempty" name:"InstanceLifeDetailDtoList"`

	// Runner运行状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunnerState *string `json:"RunnerState,omitnil,omitempty" name:"RunnerState"`

	// 错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`

	// 错误告警级别
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCodeLevel *string `json:"ErrorCodeLevel,omitnil,omitempty" name:"ErrorCodeLevel"`

	// 实例日志简略信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLogListOpsDto *InstanceLogInfo `json:"InstanceLogListOpsDto,omitnil,omitempty" name:"InstanceLogListOpsDto"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceState *string `json:"InstanceState,omitnil,omitempty" name:"InstanceState"`
}

type InstanceLifeDetailDto struct {
	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 该状态开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 实例生命周期阶段状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DetailState *string `json:"DetailState,omitnil,omitempty" name:"DetailState"`

	// 该状态结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type InstanceList struct {
	// 耗费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *string `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 是否补录
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoFlag *int64 `json:"DoFlag,omitnil,omitempty" name:"DoFlag"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLog *string `json:"LastLog,omitnil,omitempty" name:"LastLog"`

	// 调度计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitnil,omitempty" name:"SchedulerDesc"`

	// 开始启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 实例状态  EVENT_LISTENING|DEPENDENCE|BEFORE_ASPECT|ALLOCATED|LAUNCHED|KILL|SNAP_STATE_SAVING|ISSUED|RUNNING|AFTER_ASPECT|PENDING|KILLING|FINAL_STATE_SAVING|FAILED|KILL_FAILED| COMPLETED|EXPIRED|KILL_EXPIRED|DELETED
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 尝试运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *int64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`
}

type InstanceLogInfo struct {
	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 尝试运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *string `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 日志更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitnil,omitempty" name:"LastUpdate"`

	// 日志所在节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文件名 含全路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`

	// 日志创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例日志类型, run: 运行; kill: 终止
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLogType *string `json:"InstanceLogType,omitnil,omitempty" name:"InstanceLogType"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 运行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *string `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 实例状态 COMPLETED 完成 FAILED失败重试 EXPIRED失败 RUNNING运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例代码文件，为空表示对应代码文件不存在，可能是因为执行机未升级/对应类型任务无代码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CodeFileName *string `json:"CodeFileName,omitnil,omitempty" name:"CodeFileName"`

	// 扩展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtensionInfo []*AttributeItemDTO `json:"ExtensionInfo,omitnil,omitempty" name:"ExtensionInfo"`
}

type InstanceLogInfoOpsDto struct {
	// 实例运行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogInfo *string `json:"LogInfo,omitnil,omitempty" name:"LogInfo"`

	// 实例运行提交的yarn日志地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	YarnLogInfo []*string `json:"YarnLogInfo,omitnil,omitempty" name:"YarnLogInfo"`

	// 实例运行产生的datax日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataLogInfo *string `json:"DataLogInfo,omitnil,omitempty" name:"DataLogInfo"`

	// 第三方任务运行日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdTaskRunLogInfo *string `json:"ThirdTaskRunLogInfo,omitnil,omitempty" name:"ThirdTaskRunLogInfo"`

	// 第三方任务日志链接描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThirdTaskLogUrlDesc *string `json:"ThirdTaskLogUrlDesc,omitnil,omitempty" name:"ThirdTaskLogUrlDesc"`

	// 日志行数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LineCount *int64 `json:"LineCount,omitnil,omitempty" name:"LineCount"`
}

type InstanceLogList struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *string `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 最后更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitnil,omitempty" name:"LastUpdate"`

	// 节点ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 原始文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginFileName *string `json:"OriginFileName,omitnil,omitempty" name:"OriginFileName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例日志类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLogType *string `json:"InstanceLogType,omitnil,omitempty" name:"InstanceLogType"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 耗费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *string `json:"CostTime,omitnil,omitempty" name:"CostTime"`
}

type InstanceNodeInfo struct {
	// 读取节点SOURCE 写入节点SINK
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type InstanceOpsDto struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 下一个数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCurDate *string `json:"NextCurDate,omitnil,omitempty" name:"NextCurDate"`

	// 运行优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunPriority *uint64 `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// 尝试运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *uint64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// 当前运行次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tries *uint64 `json:"Tries,omitnil,omitempty" name:"Tries"`

	// 重跑总次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRunNum *uint64 `json:"TotalRunNum,omitnil,omitempty" name:"TotalRunNum"`

	// 是否补录
	// 注意：此字段可能返回 null，表示取不到有效值。
	DoFlag *uint64 `json:"DoFlag,omitnil,omitempty" name:"DoFlag"`

	// 是否是重跑
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedoFlag *uint64 `json:"RedoFlag,omitnil,omitempty" name:"RedoFlag"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 运行节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuntimeBroker *string `json:"RuntimeBroker,omitnil,omitempty" name:"RuntimeBroker"`

	// 失败的原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *TaskTypeOpsDto `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 依赖判断完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependenceFulfillTime *string `json:"DependenceFulfillTime,omitnil,omitempty" name:"DependenceFulfillTime"`

	// 首次依赖判断通过时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstDependenceFulfillTime *string `json:"FirstDependenceFulfillTime,omitnil,omitempty" name:"FirstDependenceFulfillTime"`

	// 首次启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstStartTime *string `json:"FirstStartTime,omitnil,omitempty" name:"FirstStartTime"`

	// 开始启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 运行完成时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 耗费时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostTime *string `json:"CostTime,omitnil,omitempty" name:"CostTime"`

	// 耗费时间(ms)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostMillisecond *uint64 `json:"CostMillisecond,omitnil,omitempty" name:"CostMillisecond"`

	// 最大运行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCostTime *uint64 `json:"MaxCostTime,omitnil,omitempty" name:"MaxCostTime"`

	// 最小运行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinCostTime *uint64 `json:"MinCostTime,omitnil,omitempty" name:"MinCostTime"`

	// 平均运行耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvgCostTime *float64 `json:"AvgCostTime,omitnil,omitempty" name:"AvgCostTime"`

	// 最近日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLog *string `json:"LastLog,omitnil,omitempty" name:"LastLog"`

	// 调度时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDateTime *string `json:"SchedulerDateTime,omitnil,omitempty" name:"SchedulerDateTime"`

	// 上次调度时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSchedulerDateTime *string `json:"LastSchedulerDateTime,omitnil,omitempty" name:"LastSchedulerDateTime"`

	// 最后更新事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitnil,omitempty" name:"LastUpdate"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分支，依赖关系 and、or
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyRel *string `json:"DependencyRel,omitnil,omitempty" name:"DependencyRel"`

	// 执行空间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionSpace *string `json:"ExecutionSpace,omitnil,omitempty" name:"ExecutionSpace"`

	// 忽略事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreEvent *bool `json:"IgnoreEvent,omitnil,omitempty" name:"IgnoreEvent"`

	// 虚拟任务实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`

	// 文件夹ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 递归实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SonList *string `json:"SonList,omitnil,omitempty" name:"SonList"`

	// 产品业务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 资源组指定执行节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceInstanceId *string `json:"ResourceInstanceId,omitnil,omitempty" name:"ResourceInstanceId"`

	// 资源队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// 调度计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitnil,omitempty" name:"SchedulerDesc"`

	// 最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 首次执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstRunTime *string `json:"FirstRunTime,omitnil,omitempty" name:"FirstRunTime"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 租户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 实例标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// 关联实例信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedInstanceList []*InstanceOpsDto `json:"RelatedInstanceList,omitnil,omitempty" name:"RelatedInstanceList"`

	// 关联实例信息数量，不和RelatedInstanceList强关联。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedInstanceSize *int64 `json:"RelatedInstanceSize,omitnil,omitempty" name:"RelatedInstanceSize"`

	// ownerId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 实例生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceLifeCycleOpsDto *InstanceLifeCycleOpsDto `json:"InstanceLifeCycleOpsDto,omitnil,omitempty" name:"InstanceLifeCycleOpsDto"`

	// 自动重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryAttempts *uint64 `json:"RetryAttempts,omitnil,omitempty" name:"RetryAttempts"`

	// 紧急去除的依赖父实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeletedFatherList []*string `json:"DeletedFatherList,omitnil,omitempty" name:"DeletedFatherList"`

	// 循环依赖关联的实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	CirculateInstanceList []*InstanceOpsDto `json:"CirculateInstanceList,omitnil,omitempty" name:"CirculateInstanceList"`
}

type InstanceOpsInfoPage struct {
	// 总记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InstanceOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type InstanceSearchCondition struct {
	// 任务调度周期类型
	CycleList []*string `json:"CycleList,omitnil,omitempty" name:"CycleList"`

	// 起始时间
	DateFrom *string `json:"DateFrom,omitnil,omitempty" name:"DateFrom"`

	// 截止时间
	DateTo *string `json:"DateTo,omitnil,omitempty" name:"DateTo"`

	// 实例过滤条件
	Instance *InstanceCondition `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 模糊查询关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序方式
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序字段
	SortCol *string `json:"SortCol,omitnil,omitempty" name:"SortCol"`

	// 实例状态类型
	StateList []*string `json:"StateList,omitnil,omitempty" name:"StateList"`
}

type InstanceStatisticInfo struct {
	// 实例状态趋势状态统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountList []*uint64 `json:"CountList,omitnil,omitempty" name:"CountList"`

	// 实例状态趋势时间分割
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeList []*string `json:"TimeList,omitnil,omitempty" name:"TimeList"`

	// 实例状态标识：WAITING_RUNNING、KILLING、FAILED、FAILED_TRYING、SUCCEED 分别表示等待执行、正在终止、失败、失败重试、成功，用于实例状态分布和实例状态趋势
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 用于实例状态分布计数
	InstanceCount *uint64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 当前展示时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowTime *string `json:"ShowTime,omitnil,omitempty" name:"ShowTime"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReportTime *string `json:"ReportTime,omitnil,omitempty" name:"ReportTime"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type IntegrationInstanceLog struct {
	// 任务日志信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogInfo *string `json:"LogInfo,omitnil,omitempty" name:"LogInfo"`
}

type IntegrationNodeDetail struct {
	// 集成节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集成节点类型
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点数据源类型
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 节点配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitnil,omitempty" name:"Config"`

	// 节点扩展配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema []*IntegrationNodeSchema `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 节点映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeMapping *IntegrationNodeMapping `json:"NodeMapping,omitnil,omitempty" name:"NodeMapping"`

	// owner uin
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`
}

type IntegrationNodeInfo struct {
	// 集成节点id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 集成节点所属任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 集成节点名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 集成节点类型,INPUT: 输入节点，OUTPUT:输出节点 
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点数据源类型：MYSQL|POSTGRE|ORACLE|SQLSERVER|FTP|HIVE|HDFS|ICEBERG|KAFKA|HBASE|SPARK|VIRTUAL|TBASE|DB2|DM|GAUSSDB|GBASE|IMPALA|ES|S3_DATAINSIGHT|GREENPLUM|PHOENIX|SAP_HANA|SFTP|OCEANBASE|CLICKHOUSE|KUDU|VERTICA|REDIS|COS|DLC|DLCV1|DORIS|CKAFKA|DTS_KAFKA|S3|CDW|LOCAL|TDSQLC|TDSQL|TDSQL_MYSQL|MONGODB|INFORMIX|SYBASE|REST_API|SuperSQL|PRESTO|DR_SUM|TiDB|StarRocks|Trino|Kyuubi|GDB|TCHOUSE_X|TCHOUSE_P|TDSQL_POSTGRE
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// 节点描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 节点配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitnil,omitempty" name:"Config"`

	// 节点扩展配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema []*IntegrationNodeSchema `json:"Schema,omitnil,omitempty" name:"Schema"`

	// 节点映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeMapping *IntegrationNodeMapping `json:"NodeMapping,omitnil,omitempty" name:"NodeMapping"`

	// 归属用户AppId,展示字段 非传入
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建人uin,展示字段 非传入
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 操作人uin，展示字段 非传入
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// owner uin 展示字段 非传入
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type IntegrationNodeMapping struct {
	// 源节点id
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// 目标节点id
	SinkId *string `json:"SinkId,omitnil,omitempty" name:"SinkId"`

	// 源节点schema
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceSchema []*IntegrationNodeSchema `json:"SourceSchema,omitnil,omitempty" name:"SourceSchema"`

	// 节点schema映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaMappings []*IntegrationNodeSchemaMapping `json:"SchemaMappings,omitnil,omitempty" name:"SchemaMappings"`

	// 节点映射扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`
}

type IntegrationNodeSchema struct {
	// schema id 随机唯一
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// schema名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// schema类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// schema值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// schema拓展属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Properties []*RecordField `json:"Properties,omitnil,omitempty" name:"Properties"`

	// schema别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 字段备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// category
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`
}

type IntegrationNodeSchemaMapping struct {
	// 任务节点的源schema id
	SourceSchemaId *string `json:"SourceSchemaId,omitnil,omitempty" name:"SourceSchemaId"`

	// 任务节点目标schema id
	SinkSchemaId *string `json:"SinkSchemaId,omitnil,omitempty" name:"SinkSchemaId"`
}

type IntegrationStatisticsTrendResult struct {
	// 统计属性名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticName []*string `json:"StatisticName,omitnil,omitempty" name:"StatisticName"`

	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticValue []*int64 `json:"StatisticValue,omitnil,omitempty" name:"StatisticValue"`

	// 统计项目
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticType *string `json:"StatisticType,omitnil,omitempty" name:"StatisticType"`
}

type IntegrationTaskInfo struct {
	// 任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 同步类型1.解决方案(整库迁移),2.单表同步
	SyncType *int64 `json:"SyncType,omitnil,omitempty" name:"SyncType"`

	// 201.实时,202.离线
	TaskType *int64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务所属工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务调度id(oceanus or us等作业id)，非填项
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleTaskId *string `json:"ScheduleTaskId,omitnil,omitempty" name:"ScheduleTaskId"`

	// inlong任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskGroupId *string `json:"TaskGroupId,omitnil,omitempty" name:"TaskGroupId"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 操作人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// owner uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 应用id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 0:新建(任务开发态默认状态)|1:未开始|2:操作中|3:运行中|4:暂停|5:任务停止中|6:停止|7:执行失败|20:异常|21:未知|
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Nodes []*IntegrationNodeInfo `json:"Nodes,omitnil,omitempty" name:"Nodes"`

	// 执行资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorId *string `json:"ExecutorId,omitnil,omitempty" name:"ExecutorId"`

	// 任务配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitnil,omitempty" name:"Config"`

	// 任务扩展配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 任务执行context信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecuteContext []*RecordField `json:"ExecuteContext,omitnil,omitempty" name:"ExecuteContext"`

	// 节点映射
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mappings []*IntegrationNodeMapping `json:"Mappings,omitnil,omitempty" name:"Mappings"`

	// 任务配置模式，0:画布 1:表单 3:脚本
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskMode *string `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Incharge *string `json:"Incharge,omitnil,omitempty" name:"Incharge"`

	// 离线新增参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTaskAddEntity *OfflineTaskAddParam `json:"OfflineTaskAddEntity,omitnil,omitempty" name:"OfflineTaskAddEntity"`

	// group name
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// inlong manager url
	// 注意：此字段可能返回 null，表示取不到有效值。
	InLongManagerUrl *string `json:"InLongManagerUrl,omitnil,omitempty" name:"InLongManagerUrl"`

	// stream id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InLongStreamId *string `json:"InLongStreamId,omitnil,omitempty" name:"InLongStreamId"`

	// version
	// 注意：此字段可能返回 null，表示取不到有效值。
	InLongManagerVersion *string `json:"InLongManagerVersion,omitnil,omitempty" name:"InLongManagerVersion"`

	// inlong dataproxy url
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataProxyUrl []*string `json:"DataProxyUrl,omitnil,omitempty" name:"DataProxyUrl"`

	// 任务版本是否已提交运维
	// 注意：此字段可能返回 null，表示取不到有效值。
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// 数据源类型：MYSQL|POSTGRE|ORACLE|SQLSERVER|FTP|HIVE|HDFS|ICEBERG|KAFKA|HBASE|SPARK|VIRTUAL|TBASE|DB2|DM|GAUSSDB|GBASE|IMPALA|ES|S3_DATAINSIGHT|GREENPLUM|PHOENIX|SAP_HANA|SFTP|OCEANBASE|CLICKHOUSE|KUDU|VERTICA|REDIS|COS|DLC|DLCV1|DORIS|CKAFKA|DTS_KAFKA|S3|CDW|LOCAL|TDSQLC|TDSQL|TDSQL_MYSQL|MONGODB|INFORMIX|SYBASE|REST_API|SuperSQL|PRESTO|DR_SUM|TiDB|StarRocks|Trino|Kyuubi|GDB|TCHOUSE_X|TCHOUSE_P|TDSQL_POSTGRE
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputDatasourceType *string `json:"InputDatasourceType,omitnil,omitempty" name:"InputDatasourceType"`

	// 数据源类型：MYSQL|POSTGRE|ORACLE|SQLSERVER|FTP|HIVE|HDFS|ICEBERG|KAFKA|HBASE|SPARK|VIRTUAL|TBASE|DB2|DM|GAUSSDB|GBASE|IMPALA|ES|S3_DATAINSIGHT|GREENPLUM|PHOENIX|SAP_HANA|SFTP|OCEANBASE|CLICKHOUSE|KUDU|VERTICA|REDIS|COS|DLC|DLCV1|DORIS|CKAFKA|DTS_KAFKA|S3|CDW|LOCAL|TDSQLC|TDSQL|TDSQL_MYSQL|MONGODB|INFORMIX|SYBASE|REST_API|SuperSQL|PRESTO|DR_SUM|TiDB|StarRocks|Trino|Kyuubi|GDB|TCHOUSE_X|TCHOUSE_P|TDSQL_POSTGRE
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputDatasourceType *string `json:"OutputDatasourceType,omitnil,omitempty" name:"OutputDatasourceType"`

	// 读取条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumRecordsIn *int64 `json:"NumRecordsIn,omitnil,omitempty" name:"NumRecordsIn"`

	// 写入条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumRecordsOut *int64 `json:"NumRecordsOut,omitnil,omitempty" name:"NumRecordsOut"`

	// 读取延迟
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReaderDelay *float64 `json:"ReaderDelay,omitnil,omitempty" name:"ReaderDelay"`

	// 重启次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumRestarts *int64 `json:"NumRestarts,omitnil,omitempty" name:"NumRestarts"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 任务更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 任务最后一次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastRunTime *string `json:"LastRunTime,omitnil,omitempty" name:"LastRunTime"`

	// 任务停止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StopTime *string `json:"StopTime,omitnil,omitempty" name:"StopTime"`

	// 作业是否已提交
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasVersion *bool `json:"HasVersion,omitnil,omitempty" name:"HasVersion"`

	// 任务是否被锁定
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locked *bool `json:"Locked,omitnil,omitempty" name:"Locked"`

	// 任务锁定人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Locker *string `json:"Locker,omitnil,omitempty" name:"Locker"`

	// 耗费资源量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningCu *float64 `json:"RunningCu,omitnil,omitempty" name:"RunningCu"`

	// 该任务关联的告警规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAlarmRegularList []*string `json:"TaskAlarmRegularList,omitnil,omitempty" name:"TaskAlarmRegularList"`

	// 实时任务资源分层情况： 0：进行中,1：成功 ,2：失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	SwitchResource *int64 `json:"SwitchResource,omitnil,omitempty" name:"SwitchResource"`

	// 实时任务读取阶段：0：全部全量,1：部分全量,2：全部增量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReadPhase *int64 `json:"ReadPhase,omitnil,omitempty" name:"ReadPhase"`

	// 实时任务版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceVersion *int64 `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 离线任务导入到编排空间的任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArrangeSpaceTaskId *string `json:"ArrangeSpaceTaskId,omitnil,omitempty" name:"ArrangeSpaceTaskId"`

	// 离线任务状态区分1.未提交2.已提交3.已导出
	// 注意：此字段可能返回 null，表示取不到有效值。
	OfflineTaskStatus *int64 `json:"OfflineTaskStatus,omitnil,omitempty" name:"OfflineTaskStatus"`
}

// Predefined struct for user
type JudgeResourceFileRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

type JudgeResourceFileRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`
}

func (r *JudgeResourceFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JudgeResourceFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FilePath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "JudgeResourceFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type JudgeResourceFileResponseParams struct {
	// 资源文件完整路径
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type JudgeResourceFileResponse struct {
	*tchttp.BaseResponse
	Response *JudgeResourceFileResponseParams `json:"Response"`
}

func (r *JudgeResourceFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *JudgeResourceFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillOpsMakePlanInstancesRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

type KillOpsMakePlanInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录计划ID
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`
}

func (r *KillOpsMakePlanInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillOpsMakePlanInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "PlanId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillOpsMakePlanInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillOpsMakePlanInstancesResponseParams struct {
	// 批量操作结果
	Data *BatchOperateResultOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillOpsMakePlanInstancesResponse struct {
	*tchttp.BaseResponse
	Response *KillOpsMakePlanInstancesResponseParams `json:"Response"`
}

func (r *KillOpsMakePlanInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillOpsMakePlanInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillScheduleInstancesRequestParams struct {
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

type KillScheduleInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

func (r *KillScheduleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillScheduleInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	delete(f, "SearchCondition")
	delete(f, "OptType")
	delete(f, "OperatorName")
	delete(f, "OperatorId")
	delete(f, "ProjectId")
	delete(f, "ProjectIdent")
	delete(f, "ProjectName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Count")
	delete(f, "RequestBaseInfo")
	delete(f, "IsCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "KillScheduleInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type KillScheduleInstancesResponseParams struct {
	// 结果
	Data *BatchOperateResultOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type KillScheduleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *KillScheduleInstancesResponseParams `json:"Response"`
}

func (r *KillScheduleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *KillScheduleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LifecycleInfo struct {
	// 生命周期值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Lifecycle *int64 `json:"Lifecycle,omitnil,omitempty" name:"Lifecycle"`

	// 列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*string `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 日期格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	DateFormat []*string `json:"DateFormat,omitnil,omitempty" name:"DateFormat"`
}

type LineageParamRecord struct {
	// 字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type LinkOpsDto struct {
	// 边的id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 边的key
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkKey *string `json:"LinkKey,omitnil,omitempty" name:"LinkKey"`

	// 边的源节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFrom *string `json:"TaskFrom,omitnil,omitempty" name:"TaskFrom"`

	// 边的目标节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTo *string `json:"TaskTo,omitnil,omitempty" name:"TaskTo"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 父子任务之间的依赖关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkDependencyType *string `json:"LinkDependencyType,omitnil,omitempty" name:"LinkDependencyType"`

	// 父子任务之间依赖偏移量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *string `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 边的类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkType *string `json:"LinkType,omitnil,omitempty" name:"LinkType"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`
}

// Predefined struct for user
type LockIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type LockIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志包id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 日志内容
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`
}

type LogContentInfo struct {
	// 日志内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 日志组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgId *string `json:"PkgId,omitnil,omitempty" name:"PkgId"`

	// 日志Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PkgLogId *string `json:"PkgLogId,omitnil,omitempty" name:"PkgLogId"`

	// 时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 日志所属的容器名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContainerName *string `json:"ContainerName,omitnil,omitempty" name:"ContainerName"`
}

type MakePlanInstanceOpsDtoCollection struct {
	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前页记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*InstanceOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type MakePlanOpsDto struct {
	// 补录计划ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlanId *string `json:"PlanId,omitnil,omitempty" name:"PlanId"`

	// 补录计划名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MakeName *string `json:"MakeName,omitnil,omitempty" name:"MakeName"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 补录是否检查父任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckParent *bool `json:"CheckParent,omitnil,omitempty" name:"CheckParent"`

	// 是否使用任务原有自依赖配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	SameSelfDependType *bool `json:"SameSelfDependType,omitnil,omitempty" name:"SameSelfDependType"`

	// 并行度，在SameSelfDependType为false时生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParallelNum *int64 `json:"ParallelNum,omitnil,omitempty" name:"ParallelNum"`

	// 补录实例生成周期是否修改
	// 注意：此字段可能返回 null，表示取不到有效值。
	SameCycle *bool `json:"SameCycle,omitnil,omitempty" name:"SameCycle"`

	// 调度周期转换方式-原始周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceTaskCycle *string `json:"SourceTaskCycle,omitnil,omitempty" name:"SourceTaskCycle"`

	// 调度周期转换方式-目标周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetTaskCycle *string `json:"TargetTaskCycle,omitnil,omitempty" name:"TargetTaskCycle"`

	// 调度周期转换方式-目标周期类型指定时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetTaskAction *int64 `json:"TargetTaskAction,omitnil,omitempty" name:"TargetTaskAction"`

	// 补录实例自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MapParamList []*StrToStrMap `json:"MapParamList,omitnil,omitempty" name:"MapParamList"`

	// 创建人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitnil,omitempty" name:"CreatorId"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 补录任务ID集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIdList []*string `json:"TaskIdList,omitnil,omitempty" name:"TaskIdList"`

	// 补录计划日期范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	MakeDatetimeList []*CreateMakeDatetimeInfo `json:"MakeDatetimeList,omitnil,omitempty" name:"MakeDatetimeList"`

	// 补录计划说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 补录指定的调度资源组（ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerResourceGroup *string `json:"SchedulerResourceGroup,omitnil,omitempty" name:"SchedulerResourceGroup"`

	// 补录指定的调度资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerResourceGroupName *string `json:"SchedulerResourceGroupName,omitnil,omitempty" name:"SchedulerResourceGroupName"`

	// 补录指定的集成资源组（ID）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntegrationResourceGroup *string `json:"IntegrationResourceGroup,omitnil,omitempty" name:"IntegrationResourceGroup"`

	// 补录指定的集成资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntegrationResourceGroupName *string `json:"IntegrationResourceGroupName,omitnil,omitempty" name:"IntegrationResourceGroupName"`

	// 补录计划任务数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCount *int64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 补录计划实例完成百分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompletePercent *int64 `json:"CompletePercent,omitnil,omitempty" name:"CompletePercent"`

	// 补录计划实例成功百分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessPercent *int64 `json:"SuccessPercent,omitnil,omitempty" name:"SuccessPercent"`

	// 补录检查父任务类型。取值范围：
	// <li> NONE: 全部不检查 </li>
	// <li> ALL: 检查全部上游父任务 </li>
	// <li> MAKE_SCOPE: 只在（当前补录计划）选中任务中检查 </li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CheckParentType *string `json:"CheckParentType,omitnil,omitempty" name:"CheckParentType"`
}

type MakePlanOpsDtoCollection struct {
	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前页记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*MakePlanOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type MakePlanTaskOpsDto struct {
	// 任务基本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskBaseInfo *TaskOpsDto `json:"TaskBaseInfo,omitnil,omitempty" name:"TaskBaseInfo"`

	// 补录该任务当前已生成的实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 补录任务实例完成百分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompletePercent *int64 `json:"CompletePercent,omitnil,omitempty" name:"CompletePercent"`

	// 补录任务实例成功百分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SuccessPercent *int64 `json:"SuccessPercent,omitnil,omitempty" name:"SuccessPercent"`

	// 预计生成的总实例个数，由于是异步生成，-1代表实例还未完完全生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceTotalCount *int64 `json:"InstanceTotalCount,omitnil,omitempty" name:"InstanceTotalCount"`
}

type MakePlanTaskOpsDtoCollection struct {
	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前页记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageCount *int64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*MakePlanTaskOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type ModifyApproveStatusRequestParams struct {
	// 审批单ids
	ApproveIds []*string `json:"ApproveIds,omitnil,omitempty" name:"ApproveIds"`

	// 审批备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyApproveStatusRequest struct {
	*tchttp.BaseRequest
	
	// 审批单ids
	ApproveIds []*string `json:"ApproveIds,omitnil,omitempty" name:"ApproveIds"`

	// 审批备注
	Remark *string `json:"Remark,omitnil,omitempty" name:"Remark"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyApproveStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApproveStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApproveIds")
	delete(f, "Remark")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApproveStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApproveStatusResponseParams struct {
	// 修改审批单状态
	Data []*ApproveModify `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApproveStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApproveStatusResponseParams `json:"Response"`
}

func (r *ModifyApproveStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApproveStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataSourceRequestParams struct {
	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitnil,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitnil,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitnil,omitempty" name:"Collect"`

	// 项目id
	OwnerProjectId *string `json:"OwnerProjectId,omitnil,omitempty" name:"OwnerProjectId"`

	// 项目名称
	OwnerProjectName *string `json:"OwnerProjectName,omitnil,omitempty" name:"OwnerProjectName"`

	// 项目中文名
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitnil,omitempty" name:"OwnerProjectIdent"`

	// cos bucket
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitnil,omitempty" name:"COSRegion"`

	// 操作项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源开发环境配置
	DevelopmentParams *string `json:"DevelopmentParams,omitnil,omitempty" name:"DevelopmentParams"`
}

type ModifyDataSourceRequest struct {
	*tchttp.BaseRequest
	
	// 数据源名称，在相同SpaceName下，数据源名称不能为空
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据源类别：绑定引擎、绑定数据库
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 数据源类型:枚举值
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源ID
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 业务侧数据源的配置信息扩展
	BizParams *string `json:"BizParams,omitnil,omitempty" name:"BizParams"`

	// 数据源的配置信息，以JSON KV存储，根据每个数据源类型不同，而KV存储信息不同
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 数据源描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 数据源展示名，为了可视化查看
	Display *string `json:"Display,omitnil,omitempty" name:"Display"`

	// 若数据源列表为绑定数据库，则为db名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 数据源引擎的实例ID，如CDB实例ID
	Instance *string `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 数据源数据源的可见性，1为可见、0为不可见。默认为1
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 数据源所属的业务空间名称
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 是否采集
	Collect *string `json:"Collect,omitnil,omitempty" name:"Collect"`

	// 项目id
	OwnerProjectId *string `json:"OwnerProjectId,omitnil,omitempty" name:"OwnerProjectId"`

	// 项目名称
	OwnerProjectName *string `json:"OwnerProjectName,omitnil,omitempty" name:"OwnerProjectName"`

	// 项目中文名
	OwnerProjectIdent *string `json:"OwnerProjectIdent,omitnil,omitempty" name:"OwnerProjectIdent"`

	// cos bucket
	COSBucket *string `json:"COSBucket,omitnil,omitempty" name:"COSBucket"`

	// cos region
	COSRegion *string `json:"COSRegion,omitnil,omitempty" name:"COSRegion"`

	// 操作项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源开发环境配置
	DevelopmentParams *string `json:"DevelopmentParams,omitnil,omitempty" name:"DevelopmentParams"`
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
	delete(f, "ProjectId")
	delete(f, "DevelopmentParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDataSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDataSourceResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyDimensionWeightRequestParams struct {
	// 权重信息列表
	WeightInfoList []*WeightInfo `json:"WeightInfoList,omitnil,omitempty" name:"WeightInfoList"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否重刷历史数据
	Refresh *bool `json:"Refresh,omitnil,omitempty" name:"Refresh"`
}

type ModifyDimensionWeightRequest struct {
	*tchttp.BaseRequest
	
	// 权重信息列表
	WeightInfoList []*WeightInfo `json:"WeightInfoList,omitnil,omitempty" name:"WeightInfoList"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否重刷历史数据
	Refresh *bool `json:"Refresh,omitnil,omitempty" name:"Refresh"`
}

func (r *ModifyDimensionWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDimensionWeightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WeightInfoList")
	delete(f, "ProjectId")
	delete(f, "Refresh")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDimensionWeightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDimensionWeightResponseParams struct {
	// 更新权重是否成功
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDimensionWeightResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDimensionWeightResponseParams `json:"Response"`
}

func (r *ModifyDimensionWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDimensionWeightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDsFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`
}

type ModifyDsFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 文件夹Id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 父文件夹ID
	ParentsFolderId *string `json:"ParentsFolderId,omitnil,omitempty" name:"ParentsFolderId"`
}

func (r *ModifyDsFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDsFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "FolderName")
	delete(f, "FolderId")
	delete(f, "ParentsFolderId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDsFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDsFolderResponseParams struct {
	// true代表成功，false代表失败
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDsFolderResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDsFolderResponseParams `json:"Response"`
}

func (r *ModifyDsFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDsFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExecStrategyRequestParams struct {
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 监控类型 1.未配置, 2.关联生产调度, 3.离线周期检测
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 计算队列
	ExecQueue *string `json:"ExecQueue,omitnil,omitempty" name:"ExecQueue"`

	// 执行资源组ID
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 执行资源组名称
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// 关联的生产调度任务列表
	Tasks []*ProdSchedulerTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 离线周期模式下,生效日期-开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 离线周期模式下,生效日期-结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 离线周期模式下,调度周期 
	// MINUTE_CYCLE:I,
	// HOUR_CYCLE:H,
	// DAY_CYCLE:D,
	// WEEK_CYCLE:W,
	// MONTH_CYCLE:M
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 离线周期模式下,调度步长
	CycleStep *uint64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 离线周期模式下,指定时间
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 延时执行时间，单位分钟，可选: <0-1439
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 数据库Id
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据表Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	ExecEngineType *string `json:"ExecEngineType,omitnil,omitempty" name:"ExecEngineType"`

	// 触发场景
	TriggerTypes []*string `json:"TriggerTypes,omitnil,omitempty" name:"TriggerTypes"`
}

type ModifyExecStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 监控类型 1.未配置, 2.关联生产调度, 3.离线周期检测
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 计算队列
	ExecQueue *string `json:"ExecQueue,omitnil,omitempty" name:"ExecQueue"`

	// 执行资源组ID
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 执行资源组名称
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// 关联的生产调度任务列表
	Tasks []*ProdSchedulerTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 离线周期模式下,生效日期-开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 离线周期模式下,生效日期-结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 离线周期模式下,调度周期 
	// MINUTE_CYCLE:I,
	// HOUR_CYCLE:H,
	// DAY_CYCLE:D,
	// WEEK_CYCLE:W,
	// MONTH_CYCLE:M
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 离线周期模式下,调度步长
	CycleStep *uint64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 离线周期模式下,指定时间
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 延时执行时间，单位分钟，可选: <0-1439
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 数据库Id
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据表Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	ExecEngineType *string `json:"ExecEngineType,omitnil,omitempty" name:"ExecEngineType"`

	// 触发场景
	TriggerTypes []*string `json:"TriggerTypes,omitnil,omitempty" name:"TriggerTypes"`
}

func (r *ModifyExecStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExecStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleGroupId")
	delete(f, "MonitorType")
	delete(f, "ExecQueue")
	delete(f, "ExecutorGroupId")
	delete(f, "ExecutorGroupName")
	delete(f, "Tasks")
	delete(f, "ProjectId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "CycleType")
	delete(f, "CycleStep")
	delete(f, "TaskAction")
	delete(f, "DelayTime")
	delete(f, "DatabaseId")
	delete(f, "DatasourceId")
	delete(f, "TableId")
	delete(f, "ExecEngineType")
	delete(f, "TriggerTypes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyExecStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyExecStrategyResponseParams struct {
	// 规则组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *uint64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyExecStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyExecStrategyResponseParams `json:"Response"`
}

func (r *ModifyExecStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyExecStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyIntegrationNodeRequestParams struct {
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型 "201. stream, 202. offline"
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 区分画布模式和表单 "1. 画布, 2. 表单"模式
	TaskMode *uint64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`
}

type ModifyIntegrationNodeRequest struct {
	*tchttp.BaseRequest
	
	// 集成节点信息
	NodeInfo *IntegrationNodeInfo `json:"NodeInfo,omitnil,omitempty" name:"NodeInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型 "201. stream, 202. offline"
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 区分画布模式和表单 "1. 画布, 2. 表单"模式
	TaskMode *uint64 `json:"TaskMode,omitnil,omitempty" name:"TaskMode"`
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
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 默认false . 为true时表示走回滚节点逻辑
	RollbackFlag *bool `json:"RollbackFlag,omitnil,omitempty" name:"RollbackFlag"`
}

type ModifyIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务信息
	TaskInfo *IntegrationTaskInfo `json:"TaskInfo,omitnil,omitempty" name:"TaskInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 默认false . 为true时表示走回滚节点逻辑
	RollbackFlag *bool `json:"RollbackFlag,omitnil,omitempty" name:"RollbackFlag"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyMonitorStatusRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 监控开关状态
	MonitorStatus *bool `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`
}

type ModifyMonitorStatusRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 监控开关状态
	MonitorStatus *bool `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`
}

func (r *ModifyMonitorStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMonitorStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RuleGroupId")
	delete(f, "MonitorStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMonitorStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMonitorStatusResponseParams struct {
	// 监控状态修改成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMonitorStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMonitorStatusResponseParams `json:"Response"`
}

func (r *ModifyMonitorStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMonitorStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleGroupSubscriptionRequestParams struct {
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 订阅人信息
	Receivers []*SubscribeReceiver `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 订阅类型
	SubscribeType []*uint64 `json:"SubscribeType,omitnil,omitempty" name:"SubscribeType"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库Id
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据表Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 群机器人webhook信息
	WebHooks []*SubscribeWebHook `json:"WebHooks,omitnil,omitempty" name:"WebHooks"`
}

type ModifyRuleGroupSubscriptionRequest struct {
	*tchttp.BaseRequest
	
	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 订阅人信息
	Receivers []*SubscribeReceiver `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 订阅类型
	SubscribeType []*uint64 `json:"SubscribeType,omitnil,omitempty" name:"SubscribeType"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库Id
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据源Id
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据表Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 群机器人webhook信息
	WebHooks []*SubscribeWebHook `json:"WebHooks,omitnil,omitempty" name:"WebHooks"`
}

func (r *ModifyRuleGroupSubscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleGroupSubscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleGroupId")
	delete(f, "Receivers")
	delete(f, "SubscribeType")
	delete(f, "ProjectId")
	delete(f, "DatabaseId")
	delete(f, "DatasourceId")
	delete(f, "TableId")
	delete(f, "WebHooks")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleGroupSubscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleGroupSubscriptionResponseParams struct {
	// 规则组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *uint64 `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRuleGroupSubscriptionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleGroupSubscriptionResponseParams `json:"Response"`
}

func (r *ModifyRuleGroupSubscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleGroupSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 规则模板ID
	RuleTemplateId *uint64 `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则类型 1.系统模版, 2.自定义模版, 3.自定义SQL
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则所属质量维度（1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性）
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源字段详细类型，int、string
	SourceObjectDataTypeName *string `json:"SourceObjectDataTypeName,omitnil,omitempty" name:"SourceObjectDataTypeName"`

	// 源字段名称
	SourceObjectValue *string `json:"SourceObjectValue,omitnil,omitempty" name:"SourceObjectValue"`

	// 检测范围 1.全表   2.条件扫描
	ConditionType *uint64 `json:"ConditionType,omitnil,omitempty" name:"ConditionType"`

	// 条件扫描WHERE条件表达式
	ConditionExpression *string `json:"ConditionExpression,omitnil,omitempty" name:"ConditionExpression"`

	// 自定义SQL
	CustomSql *string `json:"CustomSql,omitnil,omitempty" name:"CustomSql"`

	// 报警触发条件
	CompareRule *CompareRule `json:"CompareRule,omitnil,omitempty" name:"CompareRule"`

	// 报警触发级别 1.低, 2.中, 3.高
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 目标库Id
	TargetDatabaseId *string `json:"TargetDatabaseId,omitnil,omitempty" name:"TargetDatabaseId"`

	// 目标表Id
	TargetTableId *string `json:"TargetTableId,omitnil,omitempty" name:"TargetTableId"`

	// 目标过滤条件表达式
	TargetConditionExpr *string `json:"TargetConditionExpr,omitnil,omitempty" name:"TargetConditionExpr"`

	// 源字段与目标字段关联条件on表达式
	RelConditionExpr *string `json:"RelConditionExpr,omitnil,omitempty" name:"RelConditionExpr"`

	// 自定义模版sql表达式字段替换参数
	FieldConfig *RuleFieldConfig `json:"FieldConfig,omitnil,omitempty" name:"FieldConfig"`

	// 目标字段名称  CITY
	TargetObjectValue *string `json:"TargetObjectValue,omitnil,omitempty" name:"TargetObjectValue"`

	// 该规则适配的执行引擎
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 目标库名
	TargetDatabaseName *string `json:"TargetDatabaseName,omitnil,omitempty" name:"TargetDatabaseName"`

	// 目标模式名
	TargetSchemaName *string `json:"TargetSchemaName,omitnil,omitempty" name:"TargetSchemaName"`

	// 目标表名
	TargetTableName *string `json:"TargetTableName,omitnil,omitempty" name:"TargetTableName"`
}

type ModifyRuleRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 规则ID
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则组ID
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数据表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 规则模板ID
	RuleTemplateId *uint64 `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则类型 1.系统模版, 2.自定义模版, 3.自定义SQL
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则所属质量维度（1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性）
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源字段详细类型，int、string
	SourceObjectDataTypeName *string `json:"SourceObjectDataTypeName,omitnil,omitempty" name:"SourceObjectDataTypeName"`

	// 源字段名称
	SourceObjectValue *string `json:"SourceObjectValue,omitnil,omitempty" name:"SourceObjectValue"`

	// 检测范围 1.全表   2.条件扫描
	ConditionType *uint64 `json:"ConditionType,omitnil,omitempty" name:"ConditionType"`

	// 条件扫描WHERE条件表达式
	ConditionExpression *string `json:"ConditionExpression,omitnil,omitempty" name:"ConditionExpression"`

	// 自定义SQL
	CustomSql *string `json:"CustomSql,omitnil,omitempty" name:"CustomSql"`

	// 报警触发条件
	CompareRule *CompareRule `json:"CompareRule,omitnil,omitempty" name:"CompareRule"`

	// 报警触发级别 1.低, 2.中, 3.高
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 规则描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 目标库Id
	TargetDatabaseId *string `json:"TargetDatabaseId,omitnil,omitempty" name:"TargetDatabaseId"`

	// 目标表Id
	TargetTableId *string `json:"TargetTableId,omitnil,omitempty" name:"TargetTableId"`

	// 目标过滤条件表达式
	TargetConditionExpr *string `json:"TargetConditionExpr,omitnil,omitempty" name:"TargetConditionExpr"`

	// 源字段与目标字段关联条件on表达式
	RelConditionExpr *string `json:"RelConditionExpr,omitnil,omitempty" name:"RelConditionExpr"`

	// 自定义模版sql表达式字段替换参数
	FieldConfig *RuleFieldConfig `json:"FieldConfig,omitnil,omitempty" name:"FieldConfig"`

	// 目标字段名称  CITY
	TargetObjectValue *string `json:"TargetObjectValue,omitnil,omitempty" name:"TargetObjectValue"`

	// 该规则适配的执行引擎
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 目标库名
	TargetDatabaseName *string `json:"TargetDatabaseName,omitnil,omitempty" name:"TargetDatabaseName"`

	// 目标模式名
	TargetSchemaName *string `json:"TargetSchemaName,omitnil,omitempty" name:"TargetSchemaName"`

	// 目标表名
	TargetTableName *string `json:"TargetTableName,omitnil,omitempty" name:"TargetTableName"`
}

func (r *ModifyRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "RuleId")
	delete(f, "RuleGroupId")
	delete(f, "Name")
	delete(f, "TableId")
	delete(f, "RuleTemplateId")
	delete(f, "Type")
	delete(f, "QualityDim")
	delete(f, "SourceObjectDataTypeName")
	delete(f, "SourceObjectValue")
	delete(f, "ConditionType")
	delete(f, "ConditionExpression")
	delete(f, "CustomSql")
	delete(f, "CompareRule")
	delete(f, "AlarmLevel")
	delete(f, "Description")
	delete(f, "TargetDatabaseId")
	delete(f, "TargetTableId")
	delete(f, "TargetConditionExpr")
	delete(f, "RelConditionExpr")
	delete(f, "FieldConfig")
	delete(f, "TargetObjectValue")
	delete(f, "SourceEngineTypes")
	delete(f, "TargetDatabaseName")
	delete(f, "TargetSchemaName")
	delete(f, "TargetTableName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleResponseParams struct {
	// 是否更新成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleResponseParams `json:"Response"`
}

func (r *ModifyRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleTemplateRequestParams struct {
	// 模板ID
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板类型  1.系统模板   2.自定义模板
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 质量检测维度 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源端数据对象类型 1.常量  2.离线表级   2.离线字段级
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 源端对应的引擎类型
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 是否关联其它库表
	MultiSourceFlag *bool `json:"MultiSourceFlag,omitnil,omitempty" name:"MultiSourceFlag"`

	// SQL 表达式
	SqlExpression *string `json:"SqlExpression,omitnil,omitempty" name:"SqlExpression"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否添加where参数
	WhereFlag *bool `json:"WhereFlag,omitnil,omitempty" name:"WhereFlag"`
}

type ModifyRuleTemplateRequest struct {
	*tchttp.BaseRequest
	
	// 模板ID
	TemplateId *uint64 `json:"TemplateId,omitnil,omitempty" name:"TemplateId"`

	// 模板类型  1.系统模板   2.自定义模板
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 质量检测维度 1.准确性 2.唯一性 3.完整性 4.一致性 5.及时性 6.有效性
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 源端数据对象类型 1.常量  2.离线表级   2.离线字段级
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 源端对应的引擎类型
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 是否关联其它库表
	MultiSourceFlag *bool `json:"MultiSourceFlag,omitnil,omitempty" name:"MultiSourceFlag"`

	// SQL 表达式
	SqlExpression *string `json:"SqlExpression,omitnil,omitempty" name:"SqlExpression"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 是否添加where参数
	WhereFlag *bool `json:"WhereFlag,omitnil,omitempty" name:"WhereFlag"`
}

func (r *ModifyRuleTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleTemplateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TemplateId")
	delete(f, "Type")
	delete(f, "Name")
	delete(f, "QualityDim")
	delete(f, "SourceObjectType")
	delete(f, "Description")
	delete(f, "SourceEngineTypes")
	delete(f, "MultiSourceFlag")
	delete(f, "SqlExpression")
	delete(f, "ProjectId")
	delete(f, "WhereFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRuleTemplateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRuleTemplateResponseParams struct {
	// 修改成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRuleTemplateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRuleTemplateResponseParams `json:"Response"`
}

func (r *ModifyRuleTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRuleTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskAlarmRegularRequestParams struct {
	// 规则ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitnil,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type ModifyTaskAlarmRegularRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则信息
	TaskAlarmInfo *TaskAlarmInfo `json:"TaskAlarmInfo,omitnil,omitempty" name:"TaskAlarmInfo"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 执行时间，单位分钟，天/周/月/年调度才有。比如天调度，每天的02:00点执行一次，delayTime就是120分钟
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 新的任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 失败重试间隔,单位分钟，创建任务的时候已经给了默认值
	RetryWait *int64 `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// 失败重试次数，创建任务的时候已经给了默认值
	TryLimit *int64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// 是否可重试，1代表可以重试
	Retriable *int64 `json:"Retriable,omitnil,omitempty" name:"Retriable"`

	// 运行优先级，4高 5中 6低
	RunPriority *int64 `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// 任务的扩展配置
	TaskExt []*TaskExtInfo `json:"TaskExt,omitnil,omitempty" name:"TaskExt"`

	// 执行资源组id，需要去资源管理服务上创建调度资源组，并且绑定cvm机器
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 资源池队列名称
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// 资源组下具体执行机，any 表示可以跑在任意一台。
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 责任人
	//
	// Deprecated: InCharge is deprecated.
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 任务备注
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// 任务参数
	TaskParamInfos []*ParamInfo `json:"TaskParamInfos,omitnil,omitempty" name:"TaskParamInfos"`

	// 源数据源
	SourceServer *string `json:"SourceServer,omitnil,omitempty" name:"SourceServer"`

	// 目标数据源
	TargetServer *string `json:"TargetServer,omitnil,omitempty" name:"TargetServer"`

	// 是否支持工作流依赖 yes / no 默认 no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 依赖配置
	DependencyConfigDTOs []*DependencyConfig `json:"DependencyConfigDTOs,omitnil,omitempty" name:"DependencyConfigDTOs"`

	// 执行耗时
	ExecutionTTL *int64 `json:"ExecutionTTL,omitnil,omitempty" name:"ExecutionTTL"`

	// 脚本是否改变
	ScriptChange *bool `json:"ScriptChange,omitnil,omitempty" name:"ScriptChange"`

	// 责任人id
	InChargeIds []*string `json:"InChargeIds,omitnil,omitempty" name:"InChargeIds"`
}

type ModifyTaskInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 执行时间，单位分钟，天/周/月/年调度才有。比如天调度，每天的02:00点执行一次，delayTime就是120分钟
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 新的任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 失败重试间隔,单位分钟，创建任务的时候已经给了默认值
	RetryWait *int64 `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// 失败重试次数，创建任务的时候已经给了默认值
	TryLimit *int64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// 是否可重试，1代表可以重试
	Retriable *int64 `json:"Retriable,omitnil,omitempty" name:"Retriable"`

	// 运行优先级，4高 5中 6低
	RunPriority *int64 `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// 任务的扩展配置
	TaskExt []*TaskExtInfo `json:"TaskExt,omitnil,omitempty" name:"TaskExt"`

	// 执行资源组id，需要去资源管理服务上创建调度资源组，并且绑定cvm机器
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 资源池队列名称
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// 资源组下具体执行机，any 表示可以跑在任意一台。
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 责任人
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 任务备注
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// 任务参数
	TaskParamInfos []*ParamInfo `json:"TaskParamInfos,omitnil,omitempty" name:"TaskParamInfos"`

	// 源数据源
	SourceServer *string `json:"SourceServer,omitnil,omitempty" name:"SourceServer"`

	// 目标数据源
	TargetServer *string `json:"TargetServer,omitnil,omitempty" name:"TargetServer"`

	// 是否支持工作流依赖 yes / no 默认 no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 依赖配置
	DependencyConfigDTOs []*DependencyConfig `json:"DependencyConfigDTOs,omitnil,omitempty" name:"DependencyConfigDTOs"`

	// 执行耗时
	ExecutionTTL *int64 `json:"ExecutionTTL,omitnil,omitempty" name:"ExecutionTTL"`

	// 脚本是否改变
	ScriptChange *bool `json:"ScriptChange,omitnil,omitempty" name:"ScriptChange"`

	// 责任人id
	InChargeIds []*string `json:"InChargeIds,omitnil,omitempty" name:"InChargeIds"`
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
	delete(f, "ExecutionTTL")
	delete(f, "ScriptChange")
	delete(f, "InChargeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTaskInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyTaskInfoResponseParams struct {
	// 执行结果
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父任务ID
	TaskFrom *string `json:"TaskFrom,omitnil,omitempty" name:"TaskFrom"`

	// 子任务ID
	TaskTo *string `json:"TaskTo,omitnil,omitempty" name:"TaskTo"`

	// 子任务工作流
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 父任务工作流
	RealFromWorkflowId *string `json:"RealFromWorkflowId,omitnil,omitempty" name:"RealFromWorkflowId"`

	// 父子任务之间的依赖关系
	LinkDependencyType *string `json:"LinkDependencyType,omitnil,omitempty" name:"LinkDependencyType"`
}

type ModifyTaskLinksRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 父任务ID
	TaskFrom *string `json:"TaskFrom,omitnil,omitempty" name:"TaskFrom"`

	// 子任务ID
	TaskTo *string `json:"TaskTo,omitnil,omitempty" name:"TaskTo"`

	// 子任务工作流
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 父任务工作流
	RealFromWorkflowId *string `json:"RealFromWorkflowId,omitnil,omitempty" name:"RealFromWorkflowId"`

	// 父子任务之间的依赖关系
	LinkDependencyType *string `json:"LinkDependencyType,omitnil,omitempty" name:"LinkDependencyType"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 备注
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
}

type ModifyTaskNameRequest struct {
	*tchttp.BaseRequest
	
	// 名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目/工作空间id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 备注
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 脚本内容 base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 集成任务脚本配置
	IntegrationNodeDetails []*IntegrationNodeDetail `json:"IntegrationNodeDetails,omitnil,omitempty" name:"IntegrationNodeDetails"`
}

type ModifyTaskScriptRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 脚本内容 base64编码
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 集成任务脚本配置
	IntegrationNodeDetails []*IntegrationNodeDetail `json:"IntegrationNodeDetails,omitnil,omitempty" name:"IntegrationNodeDetails"`
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
	Data *CommonContent `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 责任人
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 责任人id
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流所属用户分组id  若有多个,分号隔开: a;b;c
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 用于配置优化参数（线程、内存、CPU核数等），仅作用于Spark SQL节点。多个参数用英文分号分隔。
	GeneralTaskParams []*GeneralTaskParam `json:"GeneralTaskParams,omitnil,omitempty" name:"GeneralTaskParams"`
}

type ModifyWorkflowInfoRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 责任人
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 责任人id
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`

	// 备注
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流所属用户分组id  若有多个,分号隔开: a;b;c
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`

	// 工作流参数列表
	WorkflowParams []*ParamInfo `json:"WorkflowParams,omitnil,omitempty" name:"WorkflowParams"`

	// 用于配置优化参数（线程、内存、CPU核数等），仅作用于Spark SQL节点。多个参数用英文分号分隔。
	GeneralTaskParams []*GeneralTaskParam `json:"GeneralTaskParams,omitnil,omitempty" name:"GeneralTaskParams"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 延迟时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 工作流依赖 ,yes 或者no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`
}

type ModifyWorkflowScheduleRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 延迟时间，单位分钟
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *int64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 自依赖类型  1:有序串行 一次一个 排队, 2: 无序串行 一次一个 不排队， 3:并行 一次多个
	SelfDepend *int64 `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// "周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 步长，间隔时间，最小1
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 生效开始时间，格式 yyyy-MM-dd HH:mm:ss
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 生效结束时间，格式 yyyy-MM-dd HH:mm:ss
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// cron表达式  周期类型为crontab调度才需要
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间，格式：HH:mm  小时调度才有，例如小时任务, 每日固定区间生效
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 工作流依赖 ,yes 或者no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`
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
	Data *BatchResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type MoveTasksToFolderRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务文件夹ID
	TaskFolderId *string `json:"TaskFolderId,omitnil,omitempty" name:"TaskFolderId"`

	// 任务ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 虚拟任务ID
	VirtualTaskIds []*string `json:"VirtualTaskIds,omitnil,omitempty" name:"VirtualTaskIds"`
}

type MoveTasksToFolderRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务文件夹ID
	TaskFolderId *string `json:"TaskFolderId,omitnil,omitempty" name:"TaskFolderId"`

	// 任务ID
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 虚拟任务ID
	VirtualTaskIds []*string `json:"VirtualTaskIds,omitnil,omitempty" name:"VirtualTaskIds"`
}

func (r *MoveTasksToFolderRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveTasksToFolderRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "TaskFolderId")
	delete(f, "TaskIds")
	delete(f, "VirtualTaskIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "MoveTasksToFolderRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type MoveTasksToFolderResponseParams struct {
	// true代表成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type MoveTasksToFolderResponse struct {
	*tchttp.BaseResponse
	Response *MoveTasksToFolderResponseParams `json:"Response"`
}

func (r *MoveTasksToFolderResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *MoveTasksToFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OfflineInstance struct {
	// 创建账号sub uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *string `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// 操作账号sub uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// 主账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUin *string `json:"OwnerUin,omitnil,omitempty" name:"OwnerUin"`

	// 账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 任务Id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 数据时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`

	// 下发时间
	IssueId *string `json:"IssueId,omitnil,omitempty" name:"IssueId"`

	// 下发资源组id，非传入项
	// 注意：此字段可能返回 null，表示取不到有效值。
	InlongTaskId *string `json:"InlongTaskId,omitnil,omitempty" name:"InlongTaskId"`

	// 资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 任务类型(1 调试运行,2 调度执行)
	TaskRunType *uint64 `json:"TaskRunType,omitnil,omitempty" name:"TaskRunType"`

	// 实例状态 EVENT_LISTENING|DEPENDENCE|BEFORE_ASPECT|ALLOCATED|LAUNCHED|KILL|SNAP_STATE_SAVING|ISSUED|RUNNING|AFTER_ASPECT|PENDING|KILLING|FINAL_STATE_SAVING|FAILED|KILL_FAILED| COMPLETED|EXPIRED|KILL_EXPIRED|DELETED
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最后更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一key
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`
}

type OfflineTaskAddParam struct {
	// 名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 依赖：yes、no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 任务开始数据时间。非空。默认当前时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务结束数据时间。非空。默认当前时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 周期类型。一次性任务:6、分钟任务：1、小时任务：2、天任务：3、周任务：4、月任务：5、crontab任务：0
	CycleType *uint64 `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 间隔，可选，默认1。非空。默认 1
	CycleStep *uint64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 延时执行时间，单位分钟
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 任务cron表达式，仅cron任务使用，其他时候默认为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 重试等待
	RetryWait *uint64 `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// 是否可以重试
	Retriable *uint64 `json:"Retriable,omitnil,omitempty" name:"Retriable"`

	// 重试限制
	TryLimit *uint64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// 优先级
	RunPriority *uint64 `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// 产品名称
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 1 有序串行 一次一个，排队 orderly 
	// 2 无序串行 一次一个，不排队 serial  
	// 3 并行 一次多个 parallel
	SelfDepend *uint64 `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 时间指定，如月任务指定1，3号，则填入 1，3。非空。默认 ""
	// 月任务：如具体1，3号则写 "1,3"，指定月末不可和具体号数一起输入，仅能为 "L"
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 调度执行结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 调度执行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 是否自动提交
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAutoSubmit *bool `json:"TaskAutoSubmit,omitnil,omitempty" name:"TaskAutoSubmit"`

	// 实例生成方式，T_PLUS_0 当天任务当天调度 / T_PLUS_1 当天任务后一天调度
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`
}

type OperationOpsDto struct {
	// 操作是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 操作结果详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultMsg *string `json:"ResultMsg,omitnil,omitempty" name:"ResultMsg"`

	// 操作失败类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitnil,omitempty" name:"ErrorId"`

	// 操作失败描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`
}

type OpsTaskCanvasDto struct {
	// 任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 任务状态，'Y','F','O','T','INVALID' 分别表示调度中、已停止、已暂停、停止中、已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *uint64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

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
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 文件夹名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 首次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstRunTime *string `json:"FirstRunTime,omitnil,omitempty" name:"FirstRunTime"`

	// 调度计划展示描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// 负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 调度周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 画布x轴坐标点
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// 画布y轴坐标点
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// 跨工作流虚拟任务标识；true标识跨工作流任务；false标识本工作流任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`

	// 弹性周期配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 延迟时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 执行开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Layer *string `json:"Layer,omitnil,omitempty" name:"Layer"`

	// 来源数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 来源数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceType *string `json:"SourceServiceType,omitnil,omitempty" name:"SourceServiceType"`

	// 目标数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// 目标数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceType *string `json:"TargetServiceType,omitnil,omitempty" name:"TargetServiceType"`

	// 任务告警类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmType *string `json:"AlarmType,omitnil,omitempty" name:"AlarmType"`

	// 任务创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type OpsTaskCanvasInfoList struct {
	// 画布任务信息
	TasksList []*OpsTaskCanvasDto `json:"TasksList,omitnil,omitempty" name:"TasksList"`

	// 画布任务链接信息
	LinksList []*OpsTaskLinkInfoDto `json:"LinksList,omitnil,omitempty" name:"LinksList"`

	// 画布循环依赖任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CirculateTaskList []*OpsTaskCanvasDto `json:"CirculateTaskList,omitnil,omitempty" name:"CirculateTaskList"`
}

type OpsTaskInfoPage struct {
	// 页号
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 任务列表信息
	Items []*TaskOpsDto `json:"Items,omitnil,omitempty" name:"Items"`

	// 总页数
	TotalPage *uint64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 页数
	PageCount *uint64 `json:"PageCount,omitnil,omitempty" name:"PageCount"`

	// 总条数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type OpsTaskLinkInfoDto struct {
	// 下游任务id
	TaskTo *string `json:"TaskTo,omitnil,omitempty" name:"TaskTo"`

	// 上游任务id
	TaskFrom *string `json:"TaskFrom,omitnil,omitempty" name:"TaskFrom"`

	// 依赖边类型 1、“real_real”表示任务->任务；2、"virtual_real" 跨工作流任务->任务
	LinkType *string `json:"LinkType,omitnil,omitempty" name:"LinkType"`

	// 依赖边id
	LinkId *string `json:"LinkId,omitnil,omitempty" name:"LinkId"`

	// 为了区分新增的循环依赖新增的类型。默认是normal，循环依赖则是circulate
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkStyle *string `json:"LinkStyle,omitnil,omitempty" name:"LinkStyle"`
}

type OrderField struct {
	// 排序字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序方向：ASC|DESC
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type OrderFieldOptional struct {
	// 排序字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 排序方向：ASC|DESC
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type OrganizationalFunction struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 层级路径
	LayerPath *string `json:"LayerPath,omitnil,omitempty" name:"LayerPath"`

	// 上级层级路径
	ParentLayerPath *string `json:"ParentLayerPath,omitnil,omitempty" name:"ParentLayerPath"`

	// 函数类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 函数分类：窗口函数、聚合函数、日期函数......
	// 注意：此字段可能返回 null，表示取不到有效值。
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 函数种类：系统函数、自定义函数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// 函数状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 函数说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 函数用法
	// 注意：此字段可能返回 null，表示取不到有效值。
	Usage *string `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 函数参数说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamDesc *string `json:"ParamDesc,omitnil,omitempty" name:"ParamDesc"`

	// 函数返回值说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReturnDesc *string `json:"ReturnDesc,omitnil,omitempty" name:"ReturnDesc"`

	// 函数示例
	// 注意：此字段可能返回 null，表示取不到有效值。
	Example *string `json:"Example,omitnil,omitempty" name:"Example"`

	// 集群实例引擎 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 函数 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FuncId *string `json:"FuncId,omitnil,omitempty" name:"FuncId"`

	// 函数类名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// 函数资源列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceList []*FunctionVersion `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`

	// 操作人 ID 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserIds []*int64 `json:"OperatorUserIds,omitnil,omitempty" name:"OperatorUserIds"`

	// 公有云 Owner ID 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserIds []*int64 `json:"OwnerUserIds,omitnil,omitempty" name:"OwnerUserIds"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 提交失败错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitErrorMsg *string `json:"SubmitErrorMsg,omitnil,omitempty" name:"SubmitErrorMsg"`

	// 模式名
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 函数命令格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommandFormat *string `json:"CommandFormat,omitnil,omitempty" name:"CommandFormat"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitTimestamp *string `json:"SubmitTimestamp,omitnil,omitempty" name:"SubmitTimestamp"`

	// 函数tag
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tag *string `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 操作人 ID 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUserIdsStr []*string `json:"OperatorUserIdsStr,omitnil,omitempty" name:"OperatorUserIdsStr"`

	// 公有云 Owner ID 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserIdsStr []*string `json:"OwnerUserIdsStr,omitnil,omitempty" name:"OwnerUserIdsStr"`

	// 数据库环境
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvType *string `json:"EnvType,omitnil,omitempty" name:"EnvType"`
}

type Pair struct {
	// 键名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PairDto struct {
	// 键名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ParamInfo struct {
	// 参数名
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type ParamInfoDs struct {
	// 参数名
	ParamKey *string `json:"ParamKey,omitnil,omitempty" name:"ParamKey"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`
}

type Partition struct {
	// 分区转换策略
	Transform *string `json:"Transform,omitnil,omitempty" name:"Transform"`

	// 分区字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 策略参数
	TransformArgs []*string `json:"TransformArgs,omitnil,omitempty" name:"TransformArgs"`
}

type PathNodeDsVO struct {
	// PathNode ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// PathNode 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// PathNode 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 父节点唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`

	// 是否叶子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 子节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*PathNodeDsVO `json:"Children,omitnil,omitempty" name:"Children"`

	// 业务参数 ,base64编译的json串，获取具体参数需要base64反编译
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type ProdSchedulerTask struct {
	// 生产调度任务工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 生产调度任务Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 生产调度任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 生产调度任务任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *int64 `json:"CycleType,omitnil,omitempty" name:"CycleType"`
}

type Project struct {
	// 项目的所在租户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识，英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 项目显示名称，可以为中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *BaseUser `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 租户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tenant *BaseTenant `json:"Tenant,omitnil,omitempty" name:"Tenant"`

	// 项目的管理员
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminUsers []*BaseUser `json:"AdminUsers,omitnil,omitempty" name:"AdminUsers"`

	// 项目关联的集群信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Clusters []*BaseClusterInfo `json:"Clusters,omitnil,omitempty" name:"Clusters"`

	// 项目的额外配置参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 项目状态：0：禁用，1：启用，-3:禁用中，2：启用中
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 项目类型，SIMPLE：简单模式 STANDARD：标准模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *string `json:"Model,omitnil,omitempty" name:"Model"`
}

type ProjectBaseInfoOpsRequest struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 是否admin
	IsAdmin *bool `json:"IsAdmin,omitnil,omitempty" name:"IsAdmin"`
}

type ProjectUserRole struct {
	// 用户角色对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	Roles []*SystemRole `json:"Roles,omitnil,omitempty" name:"Roles"`

	// mc
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否创建者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *bool `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 显示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// 是否项目管理员
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsProjectAdmin *bool `json:"IsProjectAdmin,omitnil,omitempty" name:"IsProjectAdmin"`
}

type ProjectUsersPage struct {
	// 用户集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rows []*ProjectUserRole `json:"Rows,omitnil,omitempty" name:"Rows"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 总数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总分页页码
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPageNumber *uint64 `json:"TotalPageNumber,omitnil,omitempty" name:"TotalPageNumber"`
}

type Property struct {
	// key值
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// value值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QualityScore struct {
	// 综合分数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompositeScore *float64 `json:"CompositeScore,omitnil,omitempty" name:"CompositeScore"`

	// 评分分布
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScoringDistribution []*TableScoreStatisticsInfo `json:"ScoringDistribution,omitnil,omitempty" name:"ScoringDistribution"`

	// 总表数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTableNumber *int64 `json:"TotalTableNumber,omitnil,omitempty" name:"TotalTableNumber"`
}

type QualityScoreTrend struct {
	// 周期平均分
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageScore *float64 `json:"AverageScore,omitnil,omitempty" name:"AverageScore"`

	// 日评分列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	DailyScoreList []*DailyScoreInfo `json:"DailyScoreList,omitnil,omitempty" name:"DailyScoreList"`
}

type QuietPeriod struct {
	// 代表一周里的哪些天，1代表周一，7代表周日，以此类推
	// 注意：此字段可能返回 null，表示取不到有效值。
	DaysOfWeek []*uint64 `json:"DaysOfWeek,omitnil,omitempty" name:"DaysOfWeek"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type RealTimeTaskInstanceNodeInfo struct {
	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实时任务实例节点信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNodeInfoList []*InstanceNodeInfo `json:"InstanceNodeInfoList,omitnil,omitempty" name:"InstanceNodeInfoList"`
}

type RealTimeTaskSpeed struct {
	// 同步速度条/s列表
	RecordsSpeedList []*RecordsSpeed `json:"RecordsSpeedList,omitnil,omitempty" name:"RecordsSpeedList"`

	// 同步速度字节/s列表
	BytesSpeedList []*BytesSpeed `json:"BytesSpeedList,omitnil,omitempty" name:"BytesSpeedList"`
}

type RecordField struct {
	// 字段名称，拓展字段名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段值，拓展字段值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type RecordsSpeed struct {
	// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 速度值列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*SpeedValue `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type RegisterEventListenerRequestParams struct {
	// 关键字，如果是任务，则传任务Id
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型，默认 REST_API
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 配置信息，比如最长等待时间1天配置json：{"maxWaitEventTime":1,"maxWaitEventTimeUnit":"DAYS"}
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`
}

type RegisterEventListenerRequest struct {
	*tchttp.BaseRequest
	
	// 关键字，如果是任务，则传任务Id
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 事件名称
	EventName *string `json:"EventName,omitnil,omitempty" name:"EventName"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型，默认 REST_API
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 配置信息，比如最长等待时间1天配置json：{"maxWaitEventTime":1,"maxWaitEventTimeUnit":"DAYS"}
	Properties *string `json:"Properties,omitnil,omitempty" name:"Properties"`
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
	Data *BatchReturn `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件名称，支持英文、数字和下划线，最长20个字符, 不能以数字下划线开头。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件分割类型，周期类型: DAY，HOUR，MIN，SECOND
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 广播：BROADCAST,单播：SINGLE
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// 周期类型为天和小时为HOURS ，周期类型为分钟 ：MINUTES,周期类型为秒：SECONDS
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 事件所属人
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 事件类型，默认值：TIME_SERIES
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 对应day： yyyyMMdd，对应HOUR：yyyyMMddHH，对应MIN：yyyyMMddHHmm，对应SECOND：yyyyMMddHHmmss
	DimensionFormat *string `json:"DimensionFormat,omitnil,omitempty" name:"DimensionFormat"`

	// 存活时间
	TimeToLive *int64 `json:"TimeToLive,omitnil,omitempty" name:"TimeToLive"`

	// 事件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type RegisterEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件名称，支持英文、数字和下划线，最长20个字符, 不能以数字下划线开头。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 事件分割类型，周期类型: DAY，HOUR，MIN，SECOND
	EventSubType *string `json:"EventSubType,omitnil,omitempty" name:"EventSubType"`

	// 广播：BROADCAST,单播：SINGLE
	EventBroadcastType *string `json:"EventBroadcastType,omitnil,omitempty" name:"EventBroadcastType"`

	// 周期类型为天和小时为HOURS ，周期类型为分钟 ：MINUTES,周期类型为秒：SECONDS
	TimeUnit *string `json:"TimeUnit,omitnil,omitempty" name:"TimeUnit"`

	// 事件所属人
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 事件类型，默认值：TIME_SERIES
	EventType *string `json:"EventType,omitnil,omitempty" name:"EventType"`

	// 对应day： yyyyMMdd，对应HOUR：yyyyMMddHH，对应MIN：yyyyMMddHHmm，对应SECOND：yyyyMMddHHmmss
	DimensionFormat *string `json:"DimensionFormat,omitnil,omitempty" name:"DimensionFormat"`

	// 存活时间
	TimeToLive *int64 `json:"TimeToLive,omitnil,omitempty" name:"TimeToLive"`

	// 事件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	Data *BatchReturn `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type RemoveWorkflowDsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 是否删除脚本
	DeleteScript *string `json:"DeleteScript,omitnil,omitempty" name:"DeleteScript"`

	// 删除是否通知下游
	OperateIsInform *string `json:"OperateIsInform,omitnil,omitempty" name:"OperateIsInform"`

	// 是否终止进行中的任务
	DeleteMode *string `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
}

type RemoveWorkflowDsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 是否删除脚本
	DeleteScript *string `json:"DeleteScript,omitnil,omitempty" name:"DeleteScript"`

	// 删除是否通知下游
	OperateIsInform *string `json:"OperateIsInform,omitnil,omitempty" name:"OperateIsInform"`

	// 是否终止进行中的任务
	DeleteMode *string `json:"DeleteMode,omitnil,omitempty" name:"DeleteMode"`
}

func (r *RemoveWorkflowDsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkflowDsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "DeleteScript")
	delete(f, "OperateIsInform")
	delete(f, "DeleteMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveWorkflowDsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveWorkflowDsResponseParams struct {
	// 是否删除成功
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveWorkflowDsResponse struct {
	*tchttp.BaseResponse
	Response *RemoveWorkflowDsResponseParams `json:"Response"`
}

func (r *RemoveWorkflowDsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveWorkflowDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewWorkflowSchedulerInfoDsRequestParams struct {
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 延迟时间
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *uint64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 任务依赖
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定运行时间
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 调度周期时间单位
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 调度周期
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 执行时间左闭区间
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 任务初始化策略，T_PLUS_1、T_PLUS_0、T_MINUS_1
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`

	// 工作流依赖，yes or no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// CrontabExpression
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`
}

type RenewWorkflowSchedulerInfoDsRequest struct {
	*tchttp.BaseRequest
	
	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 延迟时间
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间
	StartupTime *uint64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 任务依赖
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定运行时间
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 调度周期时间单位
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 调度周期
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 执行时间左闭区间
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 任务初始化策略，T_PLUS_1、T_PLUS_0、T_MINUS_1
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`

	// 工作流依赖，yes or no
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// CrontabExpression
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`
}

func (r *RenewWorkflowSchedulerInfoDsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewWorkflowSchedulerInfoDsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowId")
	delete(f, "DelayTime")
	delete(f, "StartupTime")
	delete(f, "SelfDepend")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "TaskAction")
	delete(f, "CycleType")
	delete(f, "CycleStep")
	delete(f, "ExecutionStartTime")
	delete(f, "ExecutionEndTime")
	delete(f, "InstanceInitStrategy")
	delete(f, "DependencyWorkflow")
	delete(f, "CrontabExpression")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenewWorkflowSchedulerInfoDsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenewWorkflowSchedulerInfoDsResponseParams struct {
	// 数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchResultDs `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenewWorkflowSchedulerInfoDsResponse struct {
	*tchttp.BaseResponse
	Response *RenewWorkflowSchedulerInfoDsResponseParams `json:"Response"`
}

func (r *RenewWorkflowSchedulerInfoDsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenewWorkflowSchedulerInfoDsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourcePathTree struct {
	// 资源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否为叶子节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsLeaf *bool `json:"IsLeaf,omitnil,omitempty" name:"IsLeaf"`

	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 本地路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// 远程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 文件MD5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Value *string `json:"Md5Value,omitnil,omitempty" name:"Md5Value"`

	// 文件拥有者名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 更新人
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUser *string `json:"UpdateUser,omitnil,omitempty" name:"UpdateUser"`

	// 文件更新人uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserId *string `json:"UpdateUserId,omitnil,omitempty" name:"UpdateUserId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *uint64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// Cos存储桶名
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosBucket *string `json:"CosBucket,omitnil,omitempty" name:"CosBucket"`

	// Cos地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosRegion *string `json:"CosRegion,omitnil,omitempty" name:"CosRegion"`

	// 额外信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`
}

// Predefined struct for user
type ResumeIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, RESUME, COMMIT, TIMESTAMP)
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 额外参数
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 前端操作类型描述
	EventDesc *string `json:"EventDesc,omitnil,omitempty" name:"EventDesc"`
}

type ResumeIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, RESUME, COMMIT, TIMESTAMP)
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 额外参数
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 前端操作类型描述
	EventDesc *string `json:"EventDesc,omitnil,omitempty" name:"EventDesc"`
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
	delete(f, "Event")
	delete(f, "ExtConfig")
	delete(f, "EventDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type RobAndLockIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务类型：201. stream,   202. offline
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	RobLockState *RobLockState `json:"RobLockState,omitnil,omitempty" name:"RobLockState"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	IsRob *bool `json:"IsRob,omitnil,omitempty" name:"IsRob"`

	// 当前持锁人
	Locker *string `json:"Locker,omitnil,omitempty" name:"Locker"`
}

type RolePrivilege struct {
	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivilegeId *string `json:"PrivilegeId,omitnil,omitempty" name:"PrivilegeId"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivilegeName *string `json:"PrivilegeName,omitnil,omitempty" name:"PrivilegeName"`

	// 路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestPath *string `json:"RestPath,omitnil,omitempty" name:"RestPath"`

	// 方法
	// 注意：此字段可能返回 null，表示取不到有效值。
	RestMethod *string `json:"RestMethod,omitnil,omitempty" name:"RestMethod"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模块id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModuleId *string `json:"ModuleId,omitnil,omitempty" name:"ModuleId"`

	// 权限类型 N、R、RW、RWD
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type Rule struct {
	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 数据表Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则类型 1.系统模版, 2.自定义模版, 3.自定义SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则模板Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTemplateId *uint64 `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则模板概述
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleTemplateContent *string `json:"RuleTemplateContent,omitnil,omitempty" name:"RuleTemplateContent"`

	// 规则所属质量维度 1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 规则适用的源数据对象类型（1：常量，2：离线表级，3：离线字段级别）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 规则适用的源数据对象类型（1：数值，2：字符串）
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceObjectDataType *uint64 `json:"SourceObjectDataType,omitnil,omitempty" name:"SourceObjectDataType"`

	// 源字段详细类型，INT、STRING
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceObjectDataTypeName *string `json:"SourceObjectDataTypeName,omitnil,omitempty" name:"SourceObjectDataTypeName"`

	// 源字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceObjectValue *string `json:"SourceObjectValue,omitnil,omitempty" name:"SourceObjectValue"`

	// 检测范围 1.全表, 2.条件扫描
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionType *uint64 `json:"ConditionType,omitnil,omitempty" name:"ConditionType"`

	// 条件扫描WHERE条件表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionExpression *string `json:"ConditionExpression,omitnil,omitempty" name:"ConditionExpression"`

	// 自定义SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomSql *string `json:"CustomSql,omitnil,omitempty" name:"CustomSql"`

	// 报警触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareRule *CompareRule `json:"CompareRule,omitnil,omitempty" name:"CompareRule"`

	// 报警触发级别 1.低, 2.中, 3.高
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 规则配置人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 目标库Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetDatabaseId *string `json:"TargetDatabaseId,omitnil,omitempty" name:"TargetDatabaseId"`

	// 目标库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetDatabaseName *string `json:"TargetDatabaseName,omitnil,omitempty" name:"TargetDatabaseName"`

	// 目标表Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetTableId *string `json:"TargetTableId,omitnil,omitempty" name:"TargetTableId"`

	// 目标表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetTableName *string `json:"TargetTableName,omitnil,omitempty" name:"TargetTableName"`

	// 目标字段过滤条件表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetConditionExpr *string `json:"TargetConditionExpr,omitnil,omitempty" name:"TargetConditionExpr"`

	// 源字段与目标字段关联条件on表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelConditionExpr *string `json:"RelConditionExpr,omitnil,omitempty" name:"RelConditionExpr"`

	// 自定义模版sql表达式参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldConfig *RuleFieldConfig `json:"FieldConfig,omitnil,omitempty" name:"FieldConfig"`

	// 是否关联多表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiSourceFlag *bool `json:"MultiSourceFlag,omitnil,omitempty" name:"MultiSourceFlag"`

	// 是否where参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhereFlag *bool `json:"WhereFlag,omitnil,omitempty" name:"WhereFlag"`

	// 模版原始SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateSql *string `json:"TemplateSql,omitnil,omitempty" name:"TemplateSql"`

	// 模版子维度：0.父维度类型,1.一致性: 枚举范围一致性,2.一致性：数值范围一致性,3.一致性：字段数据相关性
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubQualityDim *uint64 `json:"SubQualityDim,omitnil,omitempty" name:"SubQualityDim"`

	// 规则适用的目标数据对象类型（1：常量，2：离线表级，3：离线字段级别）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetObjectType *uint64 `json:"TargetObjectType,omitnil,omitempty" name:"TargetObjectType"`

	// 规则适用的目标数据对象类型（1：数值，2：字符串）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetObjectDataType *uint64 `json:"TargetObjectDataType,omitnil,omitempty" name:"TargetObjectDataType"`

	// 目标字段详细类型，INT、STRING
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetObjectDataTypeName *string `json:"TargetObjectDataTypeName,omitnil,omitempty" name:"TargetObjectDataTypeName"`

	// 目标字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetObjectValue *string `json:"TargetObjectValue,omitnil,omitempty" name:"TargetObjectValue"`

	// 源端对应的引擎类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表负责人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerName *string `json:"TableOwnerName,omitnil,omitempty" name:"TableOwnerName"`

	// 执行策略信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecStrategy *RuleGroupExecStrategy `json:"ExecStrategy,omitnil,omitempty" name:"ExecStrategy"`

	// 订阅信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subscription *RuleGroupSubscribe `json:"Subscription,omitnil,omitempty" name:"Subscription"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 数据源 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *uint64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 监控是否开启.0false,1true
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorStatus *int64 `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`

	// 触发条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerCondition *string `json:"TriggerCondition,omitnil,omitempty" name:"TriggerCondition"`

	// 0或者未返回或者null：未定义，1：生产，2：开发
	// 注意：此字段可能返回 null，表示取不到有效值。
	DsEnvType *int64 `json:"DsEnvType,omitnil,omitempty" name:"DsEnvType"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceType *int64 `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`

	// 模式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 目标模式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetSchemaName *string `json:"TargetSchemaName,omitnil,omitempty" name:"TargetSchemaName"`
}

type RuleConfig struct {
	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则检测范围类型 1.全表  2.条件扫描
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionType *uint64 `json:"ConditionType,omitnil,omitempty" name:"ConditionType"`

	// 检测范围表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 目标检测范围表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetCondition *string `json:"TargetCondition,omitnil,omitempty" name:"TargetCondition"`
}

type RuleDimCnt struct {
	// 1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性
	Dim *uint64 `json:"Dim,omitnil,omitempty" name:"Dim"`

	// count 数
	Cnt *uint64 `json:"Cnt,omitnil,omitempty" name:"Cnt"`
}

type RuleDimStat struct {
	// 总数
	TotalCnt *uint64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 维度统计数
	DimCntList []*RuleDimCnt `json:"DimCntList,omitnil,omitempty" name:"DimCntList"`
}

type RuleExecConfig struct {
	// 计算队列名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 执行资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// DLC执行引擎资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DlcGroupName *string `json:"DlcGroupName,omitnil,omitempty" name:"DlcGroupName"`
}

type RuleExecDateStat struct {
	// 统计日期
	StatDate *string `json:"StatDate,omitnil,omitempty" name:"StatDate"`

	// 告警数
	AlarmCnt *uint64 `json:"AlarmCnt,omitnil,omitempty" name:"AlarmCnt"`

	// 阻塞数
	PipelineCnt *uint64 `json:"PipelineCnt,omitnil,omitempty" name:"PipelineCnt"`
}

type RuleExecLog struct {
	// 是否完成
	// 注意：此字段可能返回 null，表示取不到有效值。
	Finished *bool `json:"Finished,omitnil,omitempty" name:"Finished"`

	// 内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`
}

type RuleExecResult struct {
	// 规则执行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleExecId *uint64 `json:"RuleExecId,omitnil,omitempty" name:"RuleExecId"`

	// 规则组执行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupExecId *uint64 `json:"RuleGroupExecId,omitnil,omitempty" name:"RuleGroupExecId"`

	// 规则组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 规则ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 规则类型 1.系统模版, 2.自定义模版, 3.自定义SQL
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleType *uint64 `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 源字段详细类型，int string
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceObjectDataTypeName *string `json:"SourceObjectDataTypeName,omitnil,omitempty" name:"SourceObjectDataTypeName"`

	// 源字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceObjectValue *string `json:"SourceObjectValue,omitnil,omitempty" name:"SourceObjectValue"`

	// 条件扫描WHERE条件表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionExpression *string `json:"ConditionExpression,omitnil,omitempty" name:"ConditionExpression"`

	// 检测结果（1:检测通过，2：触发规则，3：检测失败）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecResultStatus *uint64 `json:"ExecResultStatus,omitnil,omitempty" name:"ExecResultStatus"`

	// 触发结果，告警发送成功, 阻断任务成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerResult *string `json:"TriggerResult,omitnil,omitempty" name:"TriggerResult"`

	// 对比结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareResult *CompareResult `json:"CompareResult,omitnil,omitempty" name:"CompareResult"`

	// 模版名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TemplateName *string `json:"TemplateName,omitnil,omitempty" name:"TemplateName"`

	// 质量维度
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 目标表-库表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetDBTableName *string `json:"TargetDBTableName,omitnil,omitempty" name:"TargetDBTableName"`

	// 目标表-字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetObjectValue *string `json:"TargetObjectValue,omitnil,omitempty" name:"TargetObjectValue"`

	// 目标表-字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetObjectDataType *string `json:"TargetObjectDataType,omitnil,omitempty" name:"TargetObjectDataType"`

	// 自定义模版sql表达式参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldConfig *RuleFieldConfig `json:"FieldConfig,omitnil,omitempty" name:"FieldConfig"`

	// 源字段与目标字段关联条件on表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelConditionExpr *string `json:"RelConditionExpr,omitnil,omitempty" name:"RelConditionExpr"`

	// 执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 1/2/3:低/中/高
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`
}

type RuleExecResultDetail struct {
	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *uint64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据库guid
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 表guid
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 规则执行记录
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleExecResult *RuleExecResult `json:"RuleExecResult,omitnil,omitempty" name:"RuleExecResult"`

	// 表负责人userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerUserId *uint64 `json:"TableOwnerUserId,omitnil,omitempty" name:"TableOwnerUserId"`

	// 2.HIVE 3.DLC
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceType *uint64 `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`

	// 集群部署类型，CVM/TKE
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDeployType *string `json:"ClusterDeployType,omitnil,omitempty" name:"ClusterDeployType"`
}

type RuleExecResultPage struct {
	// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则执行结果
	Items []*RuleExecResult `json:"Items,omitnil,omitempty" name:"Items"`
}

type RuleExecStat struct {
	// 规则运行总数
	TotalCnt *uint64 `json:"TotalCnt,omitnil,omitempty" name:"TotalCnt"`

	// 环比规则运行总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTotalCnt *uint64 `json:"LastTotalCnt,omitnil,omitempty" name:"LastTotalCnt"`

	// 规则运行总数占比
	TotalCntRatio *float64 `json:"TotalCntRatio,omitnil,omitempty" name:"TotalCntRatio"`

	// 规则运行总数环比变化
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTotalCntRatio *float64 `json:"LastTotalCntRatio,omitnil,omitempty" name:"LastTotalCntRatio"`

	// 规则触发数
	TriggerCnt *uint64 `json:"TriggerCnt,omitnil,omitempty" name:"TriggerCnt"`

	// 环比规则触发数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerCnt *uint64 `json:"LastTriggerCnt,omitnil,omitempty" name:"LastTriggerCnt"`

	// 触发占总数占比
	TriggerCntRatio *float64 `json:"TriggerCntRatio,omitnil,omitempty" name:"TriggerCntRatio"`

	// 环比规则触发数变化
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastTriggerCntRatio *float64 `json:"LastTriggerCntRatio,omitnil,omitempty" name:"LastTriggerCntRatio"`

	// 规则报警数
	AlarmCnt *uint64 `json:"AlarmCnt,omitnil,omitempty" name:"AlarmCnt"`

	// 环比规则报警数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAlarmCnt *uint64 `json:"LastAlarmCnt,omitnil,omitempty" name:"LastAlarmCnt"`

	// 报警占总数占比
	AlarmCntRatio *float64 `json:"AlarmCntRatio,omitnil,omitempty" name:"AlarmCntRatio"`

	// 环比报警数变化
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAlarmCntRatio *float64 `json:"LastAlarmCntRatio,omitnil,omitempty" name:"LastAlarmCntRatio"`

	// 阻塞发生数
	PipelineCnt *uint64 `json:"PipelineCnt,omitnil,omitempty" name:"PipelineCnt"`

	// 环比阻塞发生数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastPipelineCnt *uint64 `json:"LastPipelineCnt,omitnil,omitempty" name:"LastPipelineCnt"`

	// 阻塞占总数占比
	PipelineCntRatio *float64 `json:"PipelineCntRatio,omitnil,omitempty" name:"PipelineCntRatio"`

	// 环比阻塞发生数变化
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastPipelineCntRatio *float64 `json:"LastPipelineCntRatio,omitnil,omitempty" name:"LastPipelineCntRatio"`
}

type RuleFieldConfig struct {
	// where变量
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhereConfig []*FieldConfig `json:"WhereConfig,omitnil,omitempty" name:"WhereConfig"`

	// 库表变量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableConfig []*TableConfig `json:"TableConfig,omitnil,omitempty" name:"TableConfig"`
}

type RuleGroup struct {
	// 规则组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 数据源Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceType *uint64 `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`

	// 监控类型 1.未配置, 2.关联生产调度, 3.离线周期检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 关联数据表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 关联数据表Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 关联数据表负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerName *string `json:"TableOwnerName,omitnil,omitempty" name:"TableOwnerName"`

	// 执行策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecStrategy *RuleGroupExecStrategy `json:"ExecStrategy,omitnil,omitempty" name:"ExecStrategy"`

	// 执行策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subscription *RuleGroupSubscribe `json:"Subscription,omitnil,omitempty" name:"Subscription"`

	// 数据库id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 模式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 是否有权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *bool `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 已经配置的规则数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleCount *uint64 `json:"RuleCount,omitnil,omitempty" name:"RuleCount"`

	// 监控状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorStatus *bool `json:"MonitorStatus,omitnil,omitempty" name:"MonitorStatus"`

	// 表负责人UserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerUserId *uint64 `json:"TableOwnerUserId,omitnil,omitempty" name:"TableOwnerUserId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否已配置执行策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyConfig *bool `json:"StrategyConfig,omitnil,omitempty" name:"StrategyConfig"`

	// 是否已配置执行策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscribeConfig *bool `json:"SubscribeConfig,omitnil,omitempty" name:"SubscribeConfig"`

	// 数据源环境：0或者未返回.未定义，1.生产 2.开发
	// 注意：此字段可能返回 null，表示取不到有效值。
	DsEnvType *int64 `json:"DsEnvType,omitnil,omitempty" name:"DsEnvType"`

	// EMR集群部署方式：CVM/TKE
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDeployType *string `json:"ClusterDeployType,omitnil,omitempty" name:"ClusterDeployType"`
}

type RuleGroupExecResult struct {
	// 规则组执行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupExecId *uint64 `json:"RuleGroupExecId,omitnil,omitempty" name:"RuleGroupExecId"`

	// 规则组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 执行触发类型（1：手动触发， 2：调度事中触发，3：周期调度触发）
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 执行时间 yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecTime *string `json:"ExecTime,omitnil,omitempty" name:"ExecTime"`

	// 执行状态（1.已提交 2.检测中 3.正常 4.异常）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 异常规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmRuleCount *uint64 `json:"AlarmRuleCount,omitnil,omitempty" name:"AlarmRuleCount"`

	// 总规则数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalRuleCount *uint64 `json:"TotalRuleCount,omitnil,omitempty" name:"TotalRuleCount"`

	// 源表负责人
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerName *string `json:"TableOwnerName,omitnil,omitempty" name:"TableOwnerName"`

	// 源表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 数据库id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 有无权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Permission *bool `json:"Permission,omitnil,omitempty" name:"Permission"`

	// 执行详情，调度计划或者关联生产任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecDetail *string `json:"ExecDetail,omitnil,omitempty" name:"ExecDetail"`

	// 实际执行引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 规则执行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleExecResultVOList []*RuleExecResult `json:"RuleExecResultVOList,omitnil,omitempty" name:"RuleExecResultVOList"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 本地规则表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupTableId *string `json:"RuleGroupTableId,omitnil,omitempty" name:"RuleGroupTableId"`

	// 集群部署类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterDeployType *string `json:"ClusterDeployType,omitnil,omitempty" name:"ClusterDeployType"`

	// 实例id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据库所属环境，0.未定义，1.生产 2.开发
	// 注意：此字段可能返回 null，表示取不到有效值。
	DsEnvType *string `json:"DsEnvType,omitnil,omitempty" name:"DsEnvType"`
}

type RuleGroupExecResultPage struct {
	// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则组执行结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RuleGroupExecResult `json:"Items,omitnil,omitempty" name:"Items"`
}

type RuleGroupExecStrategy struct {
	// 规则组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 监控类型 1.未配置, 2.关联生产调度, 3.离线周期检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorType *uint64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 计算队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecQueue *string `json:"ExecQueue,omitnil,omitempty" name:"ExecQueue"`

	// 执行资源组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 执行资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// 关联的生产调度任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*ProdSchedulerTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 周期开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 周期结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 调度周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 延迟调度时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 间隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *uint64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 时间指定
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 运行的执行引擎，不传时会请求该数据源下默认的执行引擎
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecEngineType *string `json:"ExecEngineType,omitnil,omitempty" name:"ExecEngineType"`

	// 执行计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecPlan *string `json:"ExecPlan,omitnil,omitempty" name:"ExecPlan"`

	// 规则id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`

	// 触发类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerTypes []*string `json:"TriggerTypes,omitnil,omitempty" name:"TriggerTypes"`

	// DLC资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	DlcGroupName *string `json:"DlcGroupName,omitnil,omitempty" name:"DlcGroupName"`
}

type RuleGroupPage struct {
	// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则组列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*RuleGroup `json:"Items,omitnil,omitempty" name:"Items"`
}

type RuleGroupSchedulerInfo struct {
	// 规则组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 1:未配置 2:关联生产调度 3:离线周期检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorType *int64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 循环类型简写
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 循环步长
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 循环类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleDesc *string `json:"CycleDesc,omitnil,omitempty" name:"CycleDesc"`

	// 离线周期检测下指定时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 离线周期检测下延迟时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 离线周期检测下注册到任务调度的任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleTaskId *string `json:"CycleTaskId,omitnil,omitempty" name:"CycleTaskId"`

	// 关联生产调度下关联的任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AssociateTaskIds []*string `json:"AssociateTaskIds,omitnil,omitempty" name:"AssociateTaskIds"`
}

type RuleGroupSubscribe struct {
	// 规则组Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupId *uint64 `json:"RuleGroupId,omitnil,omitempty" name:"RuleGroupId"`

	// 订阅接收人列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Receivers []*SubscribeReceiver `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 订阅方式 1.邮件email  2.短信sms
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscribeType []*uint64 `json:"SubscribeType,omitnil,omitempty" name:"SubscribeType"`

	// 群机器人配置的webhook信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebHooks []*SubscribeWebHook `json:"WebHooks,omitnil,omitempty" name:"WebHooks"`

	// 规则Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleId *uint64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleName *string `json:"RuleName,omitnil,omitempty" name:"RuleName"`
}

type RuleGroupTable struct {
	// 表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableInfo *RuleGroupTableInnerInfo `json:"TableInfo,omitnil,omitempty" name:"TableInfo"`

	// 规则组调度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroups []*RuleGroupSchedulerInfo `json:"RuleGroups,omitnil,omitempty" name:"RuleGroups"`

	// 订阅者信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Subscriptions []*RuleGroupSubscribe `json:"Subscriptions,omitnil,omitempty" name:"Subscriptions"`
}

type RuleGroupTableInnerInfo struct {
	// 表ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceType *int64 `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`

	// 数据库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 责任人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *int64 `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type RulePage struct {
	// 记录数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 规则列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*Rule `json:"Items,omitnil,omitempty" name:"Items"`
}

type RuleTemplate struct {
	// 规则模版ID
	RuleTemplateId *uint64 `json:"RuleTemplateId,omitnil,omitempty" name:"RuleTemplateId"`

	// 规则模版名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则模版描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模版类型（1：系统模版，2：自定义）
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 规则适用的源数据对象类型（1：常量，2：离线表级，3：离线字段级别）
	SourceObjectType *uint64 `json:"SourceObjectType,omitnil,omitempty" name:"SourceObjectType"`

	// 规则适用的源数据对象类型（1：数值，2：字符串）
	SourceObjectDataType *uint64 `json:"SourceObjectDataType,omitnil,omitempty" name:"SourceObjectDataType"`

	// 规则模版源侧内容，区分引擎，JSON 结构
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceContent *string `json:"SourceContent,omitnil,omitempty" name:"SourceContent"`

	// 源数据适用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceEngineTypes []*uint64 `json:"SourceEngineTypes,omitnil,omitempty" name:"SourceEngineTypes"`

	// 规则所属质量维度（1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性）
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityDim *uint64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`

	// 规则支持的比较方式类型（1：固定值比较，大于、小于，大于等于等 2：波动值比较，绝对值、上升、下降）
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareType *uint64 `json:"CompareType,omitnil,omitempty" name:"CompareType"`

	// 引用次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CitationCount *uint64 `json:"CitationCount,omitnil,omitempty" name:"CitationCount"`

	// 创建人id
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *uint64 `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 创建人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`

	// 更新时间yyyy-MM-dd HH:mm:ss
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否添加where参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	WhereFlag *bool `json:"WhereFlag,omitnil,omitempty" name:"WhereFlag"`

	// 是否关联多个库表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiSourceFlag *bool `json:"MultiSourceFlag,omitnil,omitempty" name:"MultiSourceFlag"`

	// 自定义模板SQL表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	SqlExpression *string `json:"SqlExpression,omitnil,omitempty" name:"SqlExpression"`

	// 模版子维度，0.父维度类型,1.一致性: 枚举范围一致性,2.一致性：数值范围一致性,3.一致性：字段数据相关性
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubQualityDim *uint64 `json:"SubQualityDim,omitnil,omitempty" name:"SubQualityDim"`

	// sql表达式解析对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResolvedSqlExpression *SqlExpression `json:"ResolvedSqlExpression,omitnil,omitempty" name:"ResolvedSqlExpression"`

	// 支持的数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceTypes []*int64 `json:"DatasourceTypes,omitnil,omitempty" name:"DatasourceTypes"`

	// 创建人IdStr
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserIdStr *string `json:"UserIdStr,omitnil,omitempty" name:"UserIdStr"`
}

type RuleTemplatePage struct {
	// 记录数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 模版列表
	Items []*RuleTemplate `json:"Items,omitnil,omitempty" name:"Items"`
}

// Predefined struct for user
type RunForceSucScheduleInstancesRequestParams struct {
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

type RunForceSucScheduleInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

func (r *RunForceSucScheduleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunForceSucScheduleInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	delete(f, "SearchCondition")
	delete(f, "OptType")
	delete(f, "OperatorName")
	delete(f, "OperatorId")
	delete(f, "ProjectId")
	delete(f, "ProjectIdent")
	delete(f, "ProjectName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Count")
	delete(f, "RequestBaseInfo")
	delete(f, "IsCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunForceSucScheduleInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunForceSucScheduleInstancesResponseParams struct {
	// 结果
	Data *BatchOperateResultOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunForceSucScheduleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RunForceSucScheduleInstancesResponseParams `json:"Response"`
}

func (r *RunForceSucScheduleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunForceSucScheduleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunRerunScheduleInstancesRequestParams struct {
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

type RunRerunScheduleInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例列表
	Instances []*InstanceOpsDto `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 检查父任务类型, true: 检查父任务; false: 不检查父任务 
	CheckFather *bool `json:"CheckFather,omitnil,omitempty" name:"CheckFather"`

	// 重跑类型, 1: 自身; 3: 孩子; 2: 自身以及孩子 
	RerunType *string `json:"RerunType,omitnil,omitempty" name:"RerunType"`

	// 实例依赖方式, 1: 自依赖; 2: 任务依赖; 3: 自依赖及父子依赖 
	DependentWay *string `json:"DependentWay,omitnil,omitempty" name:"DependentWay"`

	// 重跑忽略事件监听与否 
	SkipEventListening *bool `json:"SkipEventListening,omitnil,omitempty" name:"SkipEventListening"`

	// 下游实例范围 1: 所在工作流 2: 所在项目 3: 所有跨工作流依赖的项目
	SonInstanceType *string `json:"SonInstanceType,omitnil,omitempty" name:"SonInstanceType"`

	// 查询条件
	SearchCondition *InstanceApiOpsRequest `json:"SearchCondition,omitnil,omitempty" name:"SearchCondition"`

	// 访问类型
	OptType *string `json:"OptType,omitnil,omitempty" name:"OptType"`

	// 操作者名称
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 操作者id
	OperatorId *string `json:"OperatorId,omitnil,omitempty" name:"OperatorId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标志
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 索引页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页面大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 数据总数
	Count *int64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 基础请求信息
	RequestBaseInfo *ProjectBaseInfoOpsRequest `json:"RequestBaseInfo,omitnil,omitempty" name:"RequestBaseInfo"`

	// 是否计算总数
	IsCount *bool `json:"IsCount,omitnil,omitempty" name:"IsCount"`
}

func (r *RunRerunScheduleInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunRerunScheduleInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Instances")
	delete(f, "CheckFather")
	delete(f, "RerunType")
	delete(f, "DependentWay")
	delete(f, "SkipEventListening")
	delete(f, "SonInstanceType")
	delete(f, "SearchCondition")
	delete(f, "OptType")
	delete(f, "OperatorName")
	delete(f, "OperatorId")
	delete(f, "ProjectId")
	delete(f, "ProjectIdent")
	delete(f, "ProjectName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "Count")
	delete(f, "RequestBaseInfo")
	delete(f, "IsCount")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunRerunScheduleInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunRerunScheduleInstancesResponseParams struct {
	// 结果
	Data *BatchOperateResultOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunRerunScheduleInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RunRerunScheduleInstancesResponseParams `json:"Response"`
}

func (r *RunRerunScheduleInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunRerunScheduleInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunTasksByMultiWorkflowRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id列表
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 是否补录中间实例 0.不补录 1.补录实例
	EnableMakeUp *uint64 `json:"EnableMakeUp,omitnil,omitempty" name:"EnableMakeUp"`
}

type RunTasksByMultiWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id列表
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 是否补录中间实例 0.不补录 1.补录实例
	EnableMakeUp *uint64 `json:"EnableMakeUp,omitnil,omitempty" name:"EnableMakeUp"`
}

func (r *RunTasksByMultiWorkflowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTasksByMultiWorkflowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowIds")
	delete(f, "EnableMakeUp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunTasksByMultiWorkflowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunTasksByMultiWorkflowResponseParams struct {
	// 操作返回结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *OperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunTasksByMultiWorkflowResponse struct {
	*tchttp.BaseResponse
	Response *RunTasksByMultiWorkflowResponseParams `json:"Response"`
}

func (r *RunTasksByMultiWorkflowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunTasksByMultiWorkflowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuntimeInstanceCntTop struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 任务周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitnil,omitempty" name:"State"`

	// 耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunTime *uint64 `json:"RunTime,omitnil,omitempty" name:"RunTime"`

	// 实例运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurRunTime *string `json:"CurRunTime,omitnil,omitempty" name:"CurRunTime"`

	// 等待调度耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitScheduleTime *uint64 `json:"WaitScheduleTime,omitnil,omitempty" name:"WaitScheduleTime"`
}

// Predefined struct for user
type SaveCustomFunctionRequestParams struct {
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 分类：窗口函数、聚合函数、日期函数......
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 集群引擎实例
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 类名
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// 资源列表
	ResourceList []*FunctionResource `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`

	// 函数说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用法
	Usage *string `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 参数说明
	ParamDesc *string `json:"ParamDesc,omitnil,omitempty" name:"ParamDesc"`

	// 返回值说明
	ReturnDesc *string `json:"ReturnDesc,omitnil,omitempty" name:"ReturnDesc"`

	// 示例
	Example *string `json:"Example,omitnil,omitempty" name:"Example"`

	// 项目id，不支持修改，dlc侧创建函数保存时用
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库名称，不支持修改，dlc侧创建函数保存时用
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 函数名称，不支持修改，dlc侧创建函数保存时用
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type SaveCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 分类：窗口函数、聚合函数、日期函数......
	Kind *string `json:"Kind,omitnil,omitempty" name:"Kind"`

	// 集群引擎实例
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 类名
	ClassName *string `json:"ClassName,omitnil,omitempty" name:"ClassName"`

	// 资源列表
	ResourceList []*FunctionResource `json:"ResourceList,omitnil,omitempty" name:"ResourceList"`

	// 函数说明
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 用法
	Usage *string `json:"Usage,omitnil,omitempty" name:"Usage"`

	// 参数说明
	ParamDesc *string `json:"ParamDesc,omitnil,omitempty" name:"ParamDesc"`

	// 返回值说明
	ReturnDesc *string `json:"ReturnDesc,omitnil,omitempty" name:"ReturnDesc"`

	// 示例
	Example *string `json:"Example,omitnil,omitempty" name:"Example"`

	// 项目id，不支持修改，dlc侧创建函数保存时用
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库名称，不支持修改，dlc侧创建函数保存时用
	DbName *string `json:"DbName,omitnil,omitempty" name:"DbName"`

	// 函数名称，不支持修改，dlc侧创建函数保存时用
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "ProjectId")
	delete(f, "DbName")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveCustomFunctionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveCustomFunctionResponseParams struct {
	// 函数唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 实例运行时间
	CurRunDate *string `json:"CurRunDate,omitnil,omitempty" name:"CurRunDate"`
}

type SchemaDetail struct {
	// 列
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnKey *string `json:"ColumnKey,omitnil,omitempty" name:"ColumnKey"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ScreenInstanceInfo struct {
	// 统计标示 0：全部、1：当前天、2：昨天
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountTag *uint64 `json:"CountTag,omitnil,omitempty" name:"CountTag"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningNum *uint64 `json:"RunningNum,omitnil,omitempty" name:"RunningNum"`

	// 等待运行
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitRunningNum *uint64 `json:"WaitRunningNum,omitnil,omitempty" name:"WaitRunningNum"`

	// 等待上游
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyNum *uint64 `json:"DependencyNum,omitnil,omitempty" name:"DependencyNum"`

	// 等待事件
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaitEventNum *string `json:"WaitEventNum,omitnil,omitempty" name:"WaitEventNum"`

	// 正在终止
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppingNum *uint64 `json:"StoppingNum,omitnil,omitempty" name:"StoppingNum"`

	// 成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	SucceedNum *uint64 `json:"SucceedNum,omitnil,omitempty" name:"SucceedNum"`

	// 失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailedNum *uint64 `json:"FailedNum,omitnil,omitempty" name:"FailedNum"`
}

type ScreenTaskInfo struct {
	// 统计标示 0：全部、1：当前天、2：昨天
	CountTag *uint64 `json:"CountTag,omitnil,omitempty" name:"CountTag"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalNum *uint64 `json:"TotalNum,omitnil,omitempty" name:"TotalNum"`

	// 运行中
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunningNum *uint64 `json:"RunningNum,omitnil,omitempty" name:"RunningNum"`

	// 停止中
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppingNum *uint64 `json:"StoppingNum,omitnil,omitempty" name:"StoppingNum"`

	// 已停止
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppedNum *uint64 `json:"StoppedNum,omitnil,omitempty" name:"StoppedNum"`

	// 暂停
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrozenNum *uint64 `json:"FrozenNum,omitnil,omitempty" name:"FrozenNum"`

	// 年任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	YearNum *uint64 `json:"YearNum,omitnil,omitempty" name:"YearNum"`

	// 月任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonthNum *uint64 `json:"MonthNum,omitnil,omitempty" name:"MonthNum"`

	// 周任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeekNum *uint64 `json:"WeekNum,omitnil,omitempty" name:"WeekNum"`

	// 天任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayNum *uint64 `json:"DayNum,omitnil,omitempty" name:"DayNum"`

	// 小时任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	HourNum *uint64 `json:"HourNum,omitnil,omitempty" name:"HourNum"`

	// 分钟任务
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinuteNum *uint64 `json:"MinuteNum,omitnil,omitempty" name:"MinuteNum"`
}

type ScriptInfoResponse struct {
	// 资源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 脚本名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件扩展名类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// md5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Value *string `json:"Md5Value,omitnil,omitempty" name:"Md5Value"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *float64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 本地路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// 远程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 路径深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathDepth *int64 `json:"PathDepth,omitnil,omitempty" name:"PathDepth"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 本地临时文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalTempPath *string `json:"LocalTempPath,omitnil,omitempty" name:"LocalTempPath"`

	// 本地压缩文件路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZipPath *string `json:"ZipPath,omitnil,omitempty" name:"ZipPath"`

	// cos桶名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// cos地区
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type ScriptRequestInfo struct {
	// 脚本路径
	// 项目区1470575647377821696项目，f1目录下：
	// /datastudio/project/1470575647377821696/f1/sql1234.sql
	// 个人区：
	// /datastudio/personal/sqlTTT.sql
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 脚本版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 操作类型
	Operation *string `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 额外信息
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 所属地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 文件扩展类型
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`
}

type SearchColumnDocVO struct {
	// 字段名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段中文名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChineseName *string `json:"ChineseName,omitnil,omitempty" name:"ChineseName"`

	// 字段类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 字段类型长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Length *uint64 `json:"Length,omitnil,omitempty" name:"Length"`

	// 字段类型精度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Precision *uint64 `json:"Precision,omitnil,omitempty" name:"Precision"`

	// 字段类型scale
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// 字段默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 字段的顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 是否为分区字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartition *bool `json:"IsPartition,omitnil,omitempty" name:"IsPartition"`

	// 列上的索引类型 主键: PRI,唯一索引: UNI,一般索引: MUL
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnKey *string `json:"ColumnKey,omitnil,omitempty" name:"ColumnKey"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`
}

type SearchCondition struct {
	// 查询框架，必选
	Instance *SearchConditionInstanceNew `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 查询关键字（任务Id精确匹配，任务名称模糊匹配），可选
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitnil,omitempty" name:"SortCol"`
}

type SearchConditionInstanceNew struct {
	// 执行空间，可选 
	//  "DRY_RUN"
	ExecutionSpace *string `json:"ExecutionSpace,omitnil,omitempty" name:"ExecutionSpace"`

	// 业务产品，可选: <DATA_DEV / DATA_QUALITY / DATA_INTEGRATION，默认DATA_DEV。非空。默认 自身项目
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 资源组信息，可选
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`
}

type SearchConditionNew struct {
	// 查询框架，必选
	Instance *SearchConditionInstanceNew `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 查询关键字（任务Id精确匹配，任务名称模糊匹配），可选
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 排序顺序（asc，desc）
	Sort *string `json:"Sort,omitnil,omitempty" name:"Sort"`

	// 排序列（costTime 运行耗时，startTime 开始时间，state 实例状态，curRunDate 数据时间）
	SortCol *string `json:"SortCol,omitnil,omitempty" name:"SortCol"`
}

// Predefined struct for user
type SetTaskAlarmNewRequestParams struct {
	// 设置任务超时告警和失败告警信息
	AlarmInfoList []*AlarmInfo `json:"AlarmInfoList,omitnil,omitempty" name:"AlarmInfoList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type SetTaskAlarmNewRequest struct {
	*tchttp.BaseRequest
	
	// 设置任务超时告警和失败告警信息
	AlarmInfoList []*AlarmInfo `json:"AlarmInfoList,omitnil,omitempty" name:"AlarmInfoList"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Data *BatchOperateResult `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SimpleColumnInfo struct {
	// 列ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 限定名
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualifiedName *string `json:"QualifiedName,omitnil,omitempty" name:"QualifiedName"`

	// 列名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnName *string `json:"ColumnName,omitnil,omitempty" name:"ColumnName"`

	// 列中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnNameCn *string `json:"ColumnNameCn,omitnil,omitempty" name:"ColumnNameCn"`

	// 列类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnType *string `json:"ColumnType,omitnil,omitempty" name:"ColumnType"`

	// 列描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 前缀路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrefixPath *string `json:"PrefixPath,omitnil,omitempty" name:"PrefixPath"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 下游数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownStreamCount *int64 `json:"DownStreamCount,omitnil,omitempty" name:"DownStreamCount"`

	// 上游数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpStreamCount *int64 `json:"UpStreamCount,omitnil,omitempty" name:"UpStreamCount"`

	// 关系参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelationParams *string `json:"RelationParams,omitnil,omitempty" name:"RelationParams"`

	// 参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`

	// 任务集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*string `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

type SimpleTaskInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`
}

type SourceFieldInfo struct {
	// 字段名称
	FieldName *string `json:"FieldName,omitnil,omitempty" name:"FieldName"`

	// 字段类型
	FieldType *string `json:"FieldType,omitnil,omitempty" name:"FieldType"`

	// 字段别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 字段描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type SpeedValue struct {
	// 带毫秒的时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	Time *uint64 `json:"Time,omitnil,omitempty" name:"Time"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Speed *float64 `json:"Speed,omitnil,omitempty" name:"Speed"`
}

type SqlExpression struct {
	// sql表达式表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableExpressions []*SqlExpressionTable `json:"TableExpressions,omitnil,omitempty" name:"TableExpressions"`

	// sql表达式字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParamExpressions []*string `json:"ParamExpressions,omitnil,omitempty" name:"ParamExpressions"`
}

type SqlExpressionTable struct {
	// sql表达式表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableExpression *string `json:"TableExpression,omitnil,omitempty" name:"TableExpression"`

	// sql表达式字段名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnExpression []*string `json:"ColumnExpression,omitnil,omitempty" name:"ColumnExpression"`
}

type StageCloudApiRequest struct {
	// 无
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 无
	StageId *string `json:"StageId,omitnil,omitempty" name:"StageId"`

	// 无
	JobId *string `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 无
	StageName *string `json:"StageName,omitnil,omitempty" name:"StageName"`

	// 无
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 无
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 无
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 无
	Queue *string `json:"Queue,omitnil,omitempty" name:"Queue"`

	// 无
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 无
	Parameters []*Property `json:"Parameters,omitnil,omitempty" name:"Parameters"`

	// 无
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 无
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 无
	JobType *string `json:"JobType,omitnil,omitempty" name:"JobType"`

	// 无
	WorkFlowId *string `json:"WorkFlowId,omitnil,omitempty" name:"WorkFlowId"`
}

// Predefined struct for user
type StartIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, RESUME, COMMIT, TIMESTAMP)
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 额外参数
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 操作类型描述
	EventDesc *string `json:"EventDesc,omitnil,omitempty" name:"EventDesc"`
}

type StartIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, RESUME, COMMIT, TIMESTAMP)
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 额外参数
	ExtConfig []*RecordField `json:"ExtConfig,omitnil,omitempty" name:"ExtConfig"`

	// 操作类型描述
	EventDesc *string `json:"EventDesc,omitnil,omitempty" name:"EventDesc"`
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
	delete(f, "Event")
	delete(f, "ExtConfig")
	delete(f, "EventDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type StartTaskInfo struct {
	// 批量运行任务类型，比如START，TIMESTAMP，RESTORE，RESUME等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 任务Id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 批量运行任务配置，目前仅用与实时集成基于时间位点启动。基于时间位点启动，需要设置一个name=timestamp, value=具体时间戳的RecordField的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config []*RecordField `json:"Config,omitnil,omitempty" name:"Config"`

	// 操作类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type StopIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type StopIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type StrToStrMap struct {
	// k
	// 注意：此字段可能返回 null，表示取不到有效值。
	K *string `json:"K,omitnil,omitempty" name:"K"`

	// v
	// 注意：此字段可能返回 null，表示取不到有效值。
	V *string `json:"V,omitnil,omitempty" name:"V"`
}

// Predefined struct for user
type SubmitCustomFunctionRequestParams struct {
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 备注信息
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type SubmitCustomFunctionRequest struct {
	*tchttp.BaseRequest
	
	// 函数唯一标识
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 集群实例 ID
	ClusterIdentifier *string `json:"ClusterIdentifier,omitnil,omitempty" name:"ClusterIdentifier"`

	// 备注信息
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 项目ID
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	FunctionId *string `json:"FunctionId,omitnil,omitempty" name:"FunctionId"`

	// 错误提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type SubmitSqlTaskRequestParams struct {
	// 数据库类型
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 数据源Id
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 资源组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 脚本文件id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 执行引擎实例ID
	EngineId *string `json:"EngineId,omitnil,omitempty" name:"EngineId"`

	// 脚本内容
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 资源队列
	ResourceQueue *string `json:"ResourceQueue,omitnil,omitempty" name:"ResourceQueue"`

	// 数据库类型
	DatasourceType *string `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`

	// 计算资源名称
	ComputeResource *string `json:"ComputeResource,omitnil,omitempty" name:"ComputeResource"`

	// 高级运行参数
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// 高级设置
	ConfParams *string `json:"ConfParams,omitnil,omitempty" name:"ConfParams"`

	// 脚本内容是否base64加密
	ScriptEncryption *bool `json:"ScriptEncryption,omitnil,omitempty" name:"ScriptEncryption"`
}

type SubmitSqlTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数据库类型
	DatabaseType *string `json:"DatabaseType,omitnil,omitempty" name:"DatabaseType"`

	// 数据源Id
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 资源组Id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 脚本文件id
	ScriptId *string `json:"ScriptId,omitnil,omitempty" name:"ScriptId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 执行引擎实例ID
	EngineId *string `json:"EngineId,omitnil,omitempty" name:"EngineId"`

	// 脚本内容
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 资源队列
	ResourceQueue *string `json:"ResourceQueue,omitnil,omitempty" name:"ResourceQueue"`

	// 数据库类型
	DatasourceType *string `json:"DatasourceType,omitnil,omitempty" name:"DatasourceType"`

	// 计算资源名称
	ComputeResource *string `json:"ComputeResource,omitnil,omitempty" name:"ComputeResource"`

	// 高级运行参数
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// 高级设置
	ConfParams *string `json:"ConfParams,omitnil,omitempty" name:"ConfParams"`

	// 脚本内容是否base64加密
	ScriptEncryption *bool `json:"ScriptEncryption,omitnil,omitempty" name:"ScriptEncryption"`
}

func (r *SubmitSqlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitSqlTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DatabaseType")
	delete(f, "DatasourceId")
	delete(f, "GroupId")
	delete(f, "ScriptId")
	delete(f, "ProjectId")
	delete(f, "DatabaseName")
	delete(f, "EngineId")
	delete(f, "ScriptContent")
	delete(f, "ResourceQueue")
	delete(f, "DatasourceType")
	delete(f, "ComputeResource")
	delete(f, "RunParams")
	delete(f, "ConfParams")
	delete(f, "ScriptEncryption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitSqlTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitSqlTaskResponseParams struct {
	// 任务提交记录
	Record *AdhocRecord `json:"Record,omitnil,omitempty" name:"Record"`

	// 子任务记录列表
	Details []*AdhocDetail `json:"Details,omitnil,omitempty" name:"Details"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitSqlTaskResponse struct {
	*tchttp.BaseResponse
	Response *SubmitSqlTaskResponseParams `json:"Response"`
}

func (r *SubmitSqlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitSqlTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitnil,omitempty" name:"StartScheduling"`
}

type SubmitTaskRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 版本备注
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitnil,omitempty" name:"StartScheduling"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type SubmitTaskTestRunRequestParams struct {
	// taskId列表
	TaskIds *string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitnil,omitempty" name:"WorkFlowId"`

	// 工作流名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作流任务列表
	Tasks []*StageCloudApiRequest `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 运行参数，map的Json形式
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// 脚本内容
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 版本号
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type SubmitTaskTestRunRequest struct {
	*tchttp.BaseRequest
	
	// taskId列表
	TaskIds *string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkFlowId *string `json:"WorkFlowId,omitnil,omitempty" name:"WorkFlowId"`

	// 工作流名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 工作流任务列表
	Tasks []*StageCloudApiRequest `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 备注
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 运行参数，map的Json形式
	RunParams *string `json:"RunParams,omitnil,omitempty" name:"RunParams"`

	// 脚本内容
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`

	// 版本号
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *SubmitTaskTestRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskTestRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskIds")
	delete(f, "ProjectId")
	delete(f, "WorkFlowId")
	delete(f, "Name")
	delete(f, "Tasks")
	delete(f, "Description")
	delete(f, "RunParams")
	delete(f, "ScriptContent")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SubmitTaskTestRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SubmitTaskTestRunResponseParams struct {
	// 提交运行jobid
	JobId *int64 `json:"JobId,omitnil,omitempty" name:"JobId"`

	// 运行记录id
	RecordId []*int64 `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SubmitTaskTestRunResponse struct {
	*tchttp.BaseResponse
	Response *SubmitTaskTestRunResponseParams `json:"Response"`
}

func (r *SubmitTaskTestRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SubmitTaskTestRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubmitWorkflow struct {
	// 被提交的任务id集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// 执行结果
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 执行情况备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorDesc *string `json:"ErrorDesc,omitnil,omitempty" name:"ErrorDesc"`

	// 执行情况id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorId *string `json:"ErrorId,omitnil,omitempty" name:"ErrorId"`
}

// Predefined struct for user
type SubmitWorkflowRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 提交的版本备注
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitnil,omitempty" name:"StartScheduling"`
}

type SubmitWorkflowRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 提交的版本备注
	VersionRemark *string `json:"VersionRemark,omitnil,omitempty" name:"VersionRemark"`

	// 是否启动调度
	StartScheduling *bool `json:"StartScheduling,omitnil,omitempty" name:"StartScheduling"`
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
	Data *SubmitWorkflow `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SubscribeReceiver struct {
	// 接收人Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverUserId *uint64 `json:"ReceiverUserId,omitnil,omitempty" name:"ReceiverUserId"`

	// 接收人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverName *string `json:"ReceiverName,omitnil,omitempty" name:"ReceiverName"`

	// 接收人Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReceiverUserIdStr *string `json:"ReceiverUserIdStr,omitnil,omitempty" name:"ReceiverUserIdStr"`
}

type SubscribeWebHook struct {
	// 群机器人类型，当前支持飞书
	// 注意：此字段可能返回 null，表示取不到有效值。
	HookType *string `json:"HookType,omitnil,omitempty" name:"HookType"`

	// 群机器人webhook地址，配置方式参考https://cloud.tencent.com/document/product/1254/70736
	// 注意：此字段可能返回 null，表示取不到有效值。
	HookAddress *string `json:"HookAddress,omitnil,omitempty" name:"HookAddress"`
}

// Predefined struct for user
type SuspendIntegrationTaskRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, SUSPEND_WITHOUT_SP,RESUME, COMMIT, TIMESTAMP)	
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`
}

type SuspendIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件类型(START, STOP, SUSPEND, SUSPEND_WITHOUT_SP,RESUME, COMMIT, TIMESTAMP)	
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`
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
	delete(f, "Event")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SuspendIntegrationTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SuspendIntegrationTaskResponseParams struct {
	// 操作成功与否标识
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type SystemRole struct {
	// 角色id
	RoleId *string `json:"RoleId,omitnil,omitempty" name:"RoleId"`

	// 角色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 角色昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NameCn *string `json:"NameCn,omitnil,omitempty" name:"NameCn"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 角色权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	Privileges []*RolePrivilege `json:"Privileges,omitnil,omitempty" name:"Privileges"`

	// 方法路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	MethodPaths []*string `json:"MethodPaths,omitnil,omitempty" name:"MethodPaths"`

	// 角色类型, 分为System,Tenant,Project,Commercial
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleType *string `json:"RoleType,omitnil,omitempty" name:"RoleType"`

	// 是否系统默认
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemInit *bool `json:"SystemInit,omitnil,omitempty" name:"SystemInit"`

	// 自定义参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *string `json:"Params,omitnil,omitempty" name:"Params"`
}

type Table struct {
	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Column []*ColumnItem `json:"Column,omitnil,omitempty" name:"Column"`

	// 1
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*ColumnData `json:"Data,omitnil,omitempty" name:"Data"`
}

type TableBaseInfo struct {
	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 数据表所属数据源名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceConnectionName *string `json:"DatasourceConnectionName,omitnil,omitempty" name:"DatasourceConnectionName"`

	// 表备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableComment *string `json:"TableComment,omitnil,omitempty" name:"TableComment"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据格式类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableFormat *string `json:"TableFormat,omitnil,omitempty" name:"TableFormat"`

	// 用户昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserAlias *string `json:"UserAlias,omitnil,omitempty" name:"UserAlias"`

	// 建表用户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserSubUin *string `json:"UserSubUin,omitnil,omitempty" name:"UserSubUin"`

	// 数据治理配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	GovernPolicy *DlcDataGovernPolicy `json:"GovernPolicy,omitnil,omitempty" name:"GovernPolicy"`

	// 库数据治理是否关闭，关闭：true，开启：false
	// 注意：此字段可能返回 null，表示取不到有效值。
	DbGovernPolicyIsDisable *string `json:"DbGovernPolicyIsDisable,omitnil,omitempty" name:"DbGovernPolicyIsDisable"`
}

type TableBasicInfo struct {
	// 表的全局唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 数据源全局唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据源名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 引擎/存储类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 表类型，视图，外部表等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 所属项目英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 所属项目英中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectDisplayName *string `json:"ProjectDisplayName,omitnil,omitempty" name:"ProjectDisplayName"`

	// 责任人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerId *string `json:"TableOwnerId,omitnil,omitempty" name:"TableOwnerId"`

	// 责任人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerName *string `json:"TableOwnerName,omitnil,omitempty" name:"TableOwnerName"`

	// 存储位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageLocation *int64 `json:"StorageLocation,omitnil,omitempty" name:"StorageLocation"`

	// 表描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否分区表，0-全量表 1-分区表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartitionTable *int64 `json:"IsPartitionTable,omitnil,omitempty" name:"IsPartitionTable"`

	// 分区字段list
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionColumns []*string `json:"PartitionColumns,omitnil,omitempty" name:"PartitionColumns"`

	// 存储格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageFormat *string `json:"StorageFormat,omitnil,omitempty" name:"StorageFormat"`

	// 存储量，字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *uint64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 存储量，单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSizeWithUnit *string `json:"StorageSizeWithUnit,omitnil,omitempty" name:"StorageSizeWithUnit"`

	// 累计存储【MB】
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSizeMb *uint64 `json:"TotalSizeMb,omitnil,omitempty" name:"TotalSizeMb"`

	// 副本数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicaCount *int64 `json:"ReplicaCount,omitnil,omitempty" name:"ReplicaCount"`

	// 文件数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileCount *int64 `json:"FileCount,omitnil,omitempty" name:"FileCount"`

	// 分区总数（包含hive，iceberg）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionCount *int64 `json:"PartitionCount,omitnil,omitempty" name:"PartitionCount"`

	// 分区字段数量（包含hive，iceberg）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionFieldCount *int64 `json:"PartitionFieldCount,omitnil,omitempty" name:"PartitionFieldCount"`

	// 生命周期-分区保留天数【分区保留策略时有效】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionExpireDays *int64 `json:"PartitionExpireDays,omitnil,omitempty" name:"PartitionExpireDays"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 存储位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`
}

type TableConfig struct {
	// 数据库Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableKey *string `json:"TableKey,omitnil,omitempty" name:"TableKey"`

	// 字段变量
	// 注意：此字段可能返回 null，表示取不到有效值。
	FieldConfig []*FieldConfig `json:"FieldConfig,omitnil,omitempty" name:"FieldConfig"`
}

type TableHeat struct {
	// 表ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 统计日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	DayTime *string `json:"DayTime,omitnil,omitempty" name:"DayTime"`

	// 表热度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Heat *float64 `json:"Heat,omitnil,omitempty" name:"Heat"`

	// 表热度最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxHeat *float64 `json:"MaxHeat,omitnil,omitempty" name:"MaxHeat"`
}

type TableInfo struct {
	// 表Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表类型，view/table
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 表databaseName
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginDatabaseName *string `json:"OriginDatabaseName,omitnil,omitempty" name:"OriginDatabaseName"`

	// 表schemaName
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginSchemaName *string `json:"OriginSchemaName,omitnil,omitempty" name:"OriginSchemaName"`
}

type TableLineageBaseInfo struct {
	// 元数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreType *string `json:"MetastoreType,omitnil,omitempty" name:"MetastoreType"`

	// 由中心节点到该节点的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrefixPath *string `json:"PrefixPath,omitnil,omitempty" name:"PrefixPath"`

	// 空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表血缘参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*LineageParamRecord `json:"Params,omitnil,omitempty" name:"Params"`

	// 父节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentSet *string `json:"ParentSet,omitnil,omitempty" name:"ParentSet"`

	// 子节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildSet *string `json:"ChildSet,omitnil,omitempty" name:"ChildSet"`

	// 额外参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtParams []*RecordField `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 血缘id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 元数据类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreTypeName *string `json:"MetastoreTypeName,omitnil,omitempty" name:"MetastoreTypeName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表全称
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualifiedName *string `json:"QualifiedName,omitnil,omitempty" name:"QualifiedName"`

	// 血缘下游节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownStreamCount *int64 `json:"DownStreamCount,omitnil,omitempty" name:"DownStreamCount"`

	// 血缘上游节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpStreamCount *int64 `json:"UpStreamCount,omitnil,omitempty" name:"UpStreamCount"`

	// 血缘描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 血缘创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 血缘更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 修改血缘的任务id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*string `json:"Tasks,omitnil,omitempty" name:"Tasks"`
}

type TableLineageInfo struct {
	// 元数据类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreType *string `json:"MetastoreType,omitnil,omitempty" name:"MetastoreType"`

	// 空间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 数据源id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *string `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表血缘参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*LineageParamRecord `json:"Params,omitnil,omitempty" name:"Params"`

	// 父节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentSet *string `json:"ParentSet,omitnil,omitempty" name:"ParentSet"`

	// 子节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChildSet *string `json:"ChildSet,omitnil,omitempty" name:"ChildSet"`

	// 额外参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtParams []*RecordField `json:"ExtParams,omitnil,omitempty" name:"ExtParams"`

	// 血缘id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 元数据类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreTypeName *string `json:"MetastoreTypeName,omitnil,omitempty" name:"MetastoreTypeName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表全称
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualifiedName *string `json:"QualifiedName,omitnil,omitempty" name:"QualifiedName"`

	// 血缘下游节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownStreamCount *int64 `json:"DownStreamCount,omitnil,omitempty" name:"DownStreamCount"`

	// 血缘上游节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpStreamCount *int64 `json:"UpStreamCount,omitnil,omitempty" name:"UpStreamCount"`

	// 血缘描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 由中心节点到该节点的路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrefixPath *string `json:"PrefixPath,omitnil,omitempty" name:"PrefixPath"`

	// 血缘创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 血缘更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 修改血缘的任务id列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*string `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 模块/应用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelType *string `json:"ChannelType,omitnil,omitempty" name:"ChannelType"`

	// 展示类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayType *string `json:"DisplayType,omitnil,omitempty" name:"DisplayType"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 表类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// datasourceName
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// DatabaseName
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// DatabaseId
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`
}

type TableMeta struct {
	// 表的全局唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerName *string `json:"TableOwnerName,omitnil,omitempty" name:"TableOwnerName"`

	// 数据源全局唯一ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 所属集群名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterName *string `json:"ClusterName,omitnil,omitempty" name:"ClusterName"`

	// 数据源名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceName *string `json:"DatasourceName,omitnil,omitempty" name:"DatasourceName"`

	// 数据库名
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	TablePath *string `json:"TablePath,omitnil,omitempty" name:"TablePath"`

	// 表中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNameCn *string `json:"TableNameCn,omitnil,omitempty" name:"TableNameCn"`

	// 元数据租户ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreId *int64 `json:"MetastoreId,omitnil,omitempty" name:"MetastoreId"`

	// 技术类型，可用值:HIVE,MYSQL,KAFKA, HBASE
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetastoreType *string `json:"MetastoreType,omitnil,omitempty" name:"MetastoreType"`

	// 表描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 列分隔符
	// 注意：此字段可能返回 null，表示取不到有效值。
	ColumnSeparator *string `json:"ColumnSeparator,omitnil,omitempty" name:"ColumnSeparator"`

	// 存储格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageFormat *string `json:"StorageFormat,omitnil,omitempty" name:"StorageFormat"`

	// 存储量，字节数
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSize *int64 `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 表类型，如hive MANAGED_TABLE;EXTERNAL_TABLE
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableType *string `json:"TableType,omitnil,omitempty" name:"TableType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近数据变更时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 最近DDL变更时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DdlModifyTime *string `json:"DdlModifyTime,omitnil,omitempty" name:"DdlModifyTime"`

	// 数据最后访问时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastAccessTime *string `json:"LastAccessTime,omitnil,omitempty" name:"LastAccessTime"`

	// 所属项目英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 所属数据目录id（可能多个）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizCatalogIds []*string `json:"BizCatalogIds,omitnil,omitempty" name:"BizCatalogIds"`

	// 所属数据目录（可能多个）
	// 注意：此字段可能返回 null，表示取不到有效值。
	BizCatalogNames []*string `json:"BizCatalogNames,omitnil,omitempty" name:"BizCatalogNames"`

	// true已收藏/false表示未收藏状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasFavorite *bool `json:"HasFavorite,omitnil,omitempty" name:"HasFavorite"`

	// 生命周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifeCycleTime *int64 `json:"LifeCycleTime,omitnil,omitempty" name:"LifeCycleTime"`

	// 存储量，已转为适合的单位展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageSizeWithUnit *string `json:"StorageSizeWithUnit,omitnil,omitempty" name:"StorageSizeWithUnit"`

	// 数据源引擎的实例ID：如EMR集群实例ID/数据源实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 数据来源技术类型：HIVE/MYSQL/HBASE/KAFKA等
	// 注意：此字段可能返回 null，表示取不到有效值。
	TechnologyType *string `json:"TechnologyType,omitnil,omitempty" name:"TechnologyType"`

	// 表英文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNameEn *string `json:"TableNameEn,omitnil,omitempty" name:"TableNameEn"`

	// 项目Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// Kafka Topic 分区数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *string `json:"Partitions,omitnil,omitempty" name:"Partitions"`

	// Kafka Topic 副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplicationFactor *string `json:"ReplicationFactor,omitnil,omitempty" name:"ReplicationFactor"`

	// 所属项目英中文名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectDisplayName *string `json:"ProjectDisplayName,omitnil,omitempty" name:"ProjectDisplayName"`

	// 数据最后修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataModifyTime *string `json:"DataModifyTime,omitnil,omitempty" name:"DataModifyTime"`

	// 集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 当前用户是否有管理员权限
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasAdminAuthority *bool `json:"HasAdminAuthority,omitnil,omitempty" name:"HasAdminAuthority"`

	// 数据源展示名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatasourceDisplayName *string `json:"DatasourceDisplayName,omitnil,omitempty" name:"DatasourceDisplayName"`

	// 数据库ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 租户下对表的收藏总次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FavoriteCount *int64 `json:"FavoriteCount,omitnil,omitempty" name:"FavoriteCount"`

	// 租户下对表的点赞总次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	LikeCount *int64 `json:"LikeCount,omitnil,omitempty" name:"LikeCount"`

	// true已点赞/false表示未点赞状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	HasLike *bool `json:"HasLike,omitnil,omitempty" name:"HasLike"`

	// 表的资产评分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TablePropertyScore *TablePropertyScore `json:"TablePropertyScore,omitnil,omitempty" name:"TablePropertyScore"`

	// 表的热度值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableHeat *TableHeat `json:"TableHeat,omitnil,omitempty" name:"TableHeat"`

	// 数据源ownerProjectId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerProjectId *string `json:"OwnerProjectId,omitnil,omitempty" name:"OwnerProjectId"`

	// 表负责人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableOwnerId *string `json:"TableOwnerId,omitnil,omitempty" name:"TableOwnerId"`

	// 系统源-CLUSTER, DB-自定义源
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceCategory *string `json:"DataSourceCategory,omitnil,omitempty" name:"DataSourceCategory"`

	// 表字段信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Columns []*SearchColumnDocVO `json:"Columns,omitnil,omitempty" name:"Columns"`

	// 表采集类型
	// TABLE, VIEW, MANAGED_TABLE(Hive管理表), EXTERNAL_TABLE(Hive外部表), VIRTUAL_VIEW(虚拟视图), MATERIALIZED_VIEW(物化视图), LATERAL_VIEW, INDEX_TABLE(索引表), END_SELECT(查询结构), INSTANCE(中间临时表类型(数据血缘)), CDW(CDW表类型)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaCrawlType *string `json:"MetaCrawlType,omitnil,omitempty" name:"MetaCrawlType"`

	// 是否视图
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsView *bool `json:"IsView,omitnil,omitempty" name:"IsView"`

	// 存储位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Location *string `json:"Location,omitnil,omitempty" name:"Location"`

	// 判断是否是分区表1 是 0否
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsPartitionTable *int64 `json:"IsPartitionTable,omitnil,omitempty" name:"IsPartitionTable"`

	// 分区字段 key
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionColumns []*string `json:"PartitionColumns,omitnil,omitempty" name:"PartitionColumns"`

	// 生命周期-分区保留天数【分区保留策略时有效】
	// 注意：此字段可能返回 null，表示取不到有效值。
	PartitionExpireDays *int64 `json:"PartitionExpireDays,omitnil,omitempty" name:"PartitionExpireDays"`

	// 表附属信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableProperties []*TableMetaProperty `json:"TableProperties,omitnil,omitempty" name:"TableProperties"`

	// 环境，取值 prod或者 dev
	// 注意：此字段可能返回 null，表示取不到有效值。
	Environment *string `json:"Environment,omitnil,omitempty" name:"Environment"`

	// 数据库模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Schema *string `json:"Schema,omitnil,omitempty" name:"Schema"`
}

type TableMetaProperty struct {
	// 属性的key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 属性的值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TableNameFilter struct {
	// 数据源类型
	MsType *string `json:"MsType,omitnil,omitempty" name:"MsType"`

	// 数据源id
	DatasourceId *int64 `json:"DatasourceId,omitnil,omitempty" name:"DatasourceId"`

	// 数据库名称
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// schema
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 表名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type TablePartition struct {
	// 分区名称
	PartitionName *string `json:"PartitionName,omitnil,omitempty" name:"PartitionName"`

	// 分区记录数
	RecordCount *int64 `json:"RecordCount,omitnil,omitempty" name:"RecordCount"`

	// 分区数据存储大小，字节数
	StorageSize *string `json:"StorageSize,omitnil,omitempty" name:"StorageSize"`

	// 分区创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 分区修改时间
	ModifiedTime *string `json:"ModifiedTime,omitnil,omitempty" name:"ModifiedTime"`

	// 分区数据存储大小，已转为适合的单位
	StorageSizeWithUnit *string `json:"StorageSizeWithUnit,omitnil,omitempty" name:"StorageSizeWithUnit"`

	// 文件数
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumFiles *int64 `json:"NumFiles,omitnil,omitempty" name:"NumFiles"`

	// 平均文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageFileSizeWithUnit *string `json:"AverageFileSizeWithUnit,omitnil,omitempty" name:"AverageFileSizeWithUnit"`
}

type TablePropertyScore struct {
	// 表ID
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 统计日期
	DayTime *string `json:"DayTime,omitnil,omitempty" name:"DayTime"`

	// 表完整性评分
	Integrity *float64 `json:"Integrity,omitnil,omitempty" name:"Integrity"`

	// 表保障性评分
	Safety *float64 `json:"Safety,omitnil,omitempty" name:"Safety"`

	// 表及时性评分
	Timeliness *float64 `json:"Timeliness,omitnil,omitempty" name:"Timeliness"`

	// 表稳定性评分
	Stability *float64 `json:"Stability,omitnil,omitempty" name:"Stability"`

	// 表规范性评分
	Normative *float64 `json:"Normative,omitnil,omitempty" name:"Normative"`

	// 资产评分平均分
	Average *float64 `json:"Average,omitnil,omitempty" name:"Average"`
}

type TableQualityDetail struct {
	// 数据库id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseId *string `json:"DatabaseId,omitnil,omitempty" name:"DatabaseId"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 表id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表责任人ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserId *int64 `json:"OwnerUserId,omitnil,omitempty" name:"OwnerUserId"`

	// 表责任人名
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerUserName *string `json:"OwnerUserName,omitnil,omitempty" name:"OwnerUserName"`

	// 库得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseScore *float64 `json:"DatabaseScore,omitnil,omitempty" name:"DatabaseScore"`

	// 表得分
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableScore *float64 `json:"TableScore,omitnil,omitempty" name:"TableScore"`

	// 表环比
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastPeriodRatio *float64 `json:"LastPeriodRatio,omitnil,omitempty" name:"LastPeriodRatio"`

	// 0或者未返回或者null：未定义，1：生产，2：开发
	// 注意：此字段可能返回 null，表示取不到有效值。
	DsEnvType *int64 `json:"DsEnvType,omitnil,omitempty" name:"DsEnvType"`

	// 模式名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`

	// 规则表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleGroupTableId *string `json:"RuleGroupTableId,omitnil,omitempty" name:"RuleGroupTableId"`
}

type TableQualityDetailPage struct {
	// 总条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 表质量列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*TableQualityDetail `json:"Items,omitnil,omitempty" name:"Items"`
}

type TableScoreStatisticsInfo struct {
	// 等级 1、2、3、4、5
	// 注意：此字段可能返回 null，表示取不到有效值。
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scale *int64 `json:"Scale,omitnil,omitempty" name:"Scale"`

	// 表数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableNumber *int64 `json:"TableNumber,omitnil,omitempty" name:"TableNumber"`
}

type TagVoteSum struct {
	// 标签id
	TagId *int64 `json:"TagId,omitnil,omitempty" name:"TagId"`

	// 该表该标签投票次数
	VoteSum *int64 `json:"VoteSum,omitnil,omitempty" name:"VoteSum"`

	// 当前用户对这张表是否加了该标签 true 已添加 false 未添加
	Status *bool `json:"Status,omitnil,omitempty" name:"Status"`

	// 标签名
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`
}

type TaskAlarmInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 规则名称
	RegularName *string `json:"RegularName,omitnil,omitempty" name:"RegularName"`

	// 规则状态(0表示关闭，1表示打开)
	RegularStatus *uint64 `json:"RegularStatus,omitnil,omitempty" name:"RegularStatus"`

	// 告警级别(0表示普通，1表示重要，2表示紧急)
	AlarmLevel *uint64 `json:"AlarmLevel,omitnil,omitempty" name:"AlarmLevel"`

	// 告警方式,多个用逗号隔开（1:邮件，2:短信，3:微信，4:语音，5:代表企业微信，6:http）
	AlarmWay *string `json:"AlarmWay,omitnil,omitempty" name:"AlarmWay"`

	// 任务类型(201表示实时，202表示离线)
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 规则ID
	RegularId *string `json:"RegularId,omitnil,omitempty" name:"RegularId"`

	// 告警指标,0表示任务失败，1表示任务运行超时，2表示任务停止，3表示任务暂停
	// ，4写入速度，5读取速度，6读取吞吐，7写入吞吐, 8脏数据字节数，9脏数据条数
	AlarmIndicator *uint64 `json:"AlarmIndicator,omitnil,omitempty" name:"AlarmIndicator"`

	// 指标阈值(1表示离线任务第一次运行失败，2表示离线任务所有重试完成后失败)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *uint64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// 预计的超时时间(分钟级别)
	// 注意：此字段可能返回 null，表示取不到有效值。
	EstimatedTime *uint64 `json:"EstimatedTime,omitnil,omitempty" name:"EstimatedTime"`

	// 告警接收人ID，多个用逗号隔开
	AlarmRecipientId *string `json:"AlarmRecipientId,omitnil,omitempty" name:"AlarmRecipientId"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creater *string `json:"Creater,omitnil,omitempty" name:"Creater"`

	// 告警接收人昵称，多个用逗号隔开
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmRecipientName *string `json:"AlarmRecipientName,omitnil,omitempty" name:"AlarmRecipientName"`

	// 告警指标描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicatorDesc *string `json:"AlarmIndicatorDesc,omitnil,omitempty" name:"AlarmIndicatorDesc"`

	// 实时任务告警需要的参数，1是大于2是小于
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *uint64 `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 节点id，多个逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称，多个逗号分隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmIndicatorInfos []*AlarmIndicatorInfo `json:"AlarmIndicatorInfos,omitnil,omitempty" name:"AlarmIndicatorInfos"`

	// 告警接收人类型，0指定人员；1任务责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmRecipientType *uint64 `json:"AlarmRecipientType,omitnil,omitempty" name:"AlarmRecipientType"`

	// 免打扰时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuietPeriods []*QuietPeriod `json:"QuietPeriods,omitnil,omitempty" name:"QuietPeriods"`

	// 企业微信群Hook地址，多个hook地址使用,隔开
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeComHook *string `json:"WeComHook,omitnil,omitempty" name:"WeComHook"`

	// 最近操作时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最近操作人Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorUin *string `json:"OperatorUin,omitnil,omitempty" name:"OperatorUin"`

	// 关联任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCount *int64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 监控对象类型,1:所有任务,2:指定任务,3:指定责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorType *int64 `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 监控对象列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	MonitorObjectIds []*string `json:"MonitorObjectIds,omitnil,omitempty" name:"MonitorObjectIds"`

	// 最近一次告警的实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestAlarmInstanceId *string `json:"LatestAlarmInstanceId,omitnil,omitempty" name:"LatestAlarmInstanceId"`

	// 最近一次告警时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestAlarmTime *string `json:"LatestAlarmTime,omitnil,omitempty" name:"LatestAlarmTime"`

	// 告警规则描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 飞书群Hook地址，多个hook地址使用,隔开
	// 注意：此字段可能返回 null，表示取不到有效值。
	LarkWebHooks *string `json:"LarkWebHooks,omitnil,omitempty" name:"LarkWebHooks"`

	// 钉钉群Hook地址，多个hook地址使用,隔开
	// 注意：此字段可能返回 null，表示取不到有效值。
	DingDingWebHooks *string `json:"DingDingWebHooks,omitnil,omitempty" name:"DingDingWebHooks"`
}

type TaskByCycle struct {
	// num
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *uint64 `json:"Number,omitnil,omitempty" name:"Number"`

	// 周期单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type TaskByStatus struct {
	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountGroup *string `json:"CountGroup,omitnil,omitempty" name:"CountGroup"`

	// 日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowTimeGroup *string `json:"ShowTimeGroup,omitnil,omitempty" name:"ShowTimeGroup"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 周期单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 1
	ReportTime *string `json:"ReportTime,omitnil,omitempty" name:"ReportTime"`

	// 1
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`
}

type TaskExtInfo struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TaskInnerInfo struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 周期类型  0:crontab类型, 1:分钟，2:小时，3:天，4:周，5:月，6:一次性，7:用户驱动，10:弹性周期 周,11:弹性周期 月,12:年,13:即时触发Instant类型，与正常周期调度任务逻辑隔离
	CycleType *int64 `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 虚拟任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskId *string `json:"VirtualTaskId,omitnil,omitempty" name:"VirtualTaskId"`

	// 虚拟任务标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`

	// 真实任务工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealWorkflowId *string `json:"RealWorkflowId,omitnil,omitempty" name:"RealWorkflowId"`
}

type TaskInstanceCountDto struct {
	// 成功的实例数
	Success *uint64 `json:"Success,omitnil,omitempty" name:"Success"`

	// 执行中实例数
	Running *uint64 `json:"Running,omitnil,omitempty" name:"Running"`

	// 等待中的实例数
	Waiting *uint64 `json:"Waiting,omitnil,omitempty" name:"Waiting"`

	// 等待上游实例数
	Depend *uint64 `json:"Depend,omitnil,omitempty" name:"Depend"`

	// 失败实例数
	Failed *uint64 `json:"Failed,omitnil,omitempty" name:"Failed"`

	// 永久终止实例数
	Stopped *uint64 `json:"Stopped,omitnil,omitempty" name:"Stopped"`
}

type TaskLineageInfo struct {
	// 任务 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 表名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 表 URI，格式：dsn.name
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableUri *string `json:"TableUri,omitnil,omitempty" name:"TableUri"`

	// 表方向类型
	//     - 0 - table作为源表
	//     - 1 - table作为目标表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceId *int64 `json:"DataSourceId,omitnil,omitempty" name:"DataSourceId"`

	// 数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSourceType *string `json:"DataSourceType,omitnil,omitempty" name:"DataSourceType"`

	// 数据库名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DatabaseName *string `json:"DatabaseName,omitnil,omitempty" name:"DatabaseName"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchemaName *string `json:"SchemaName,omitnil,omitempty" name:"SchemaName"`
}

type TaskLineageInfoPair struct {
	// 表血缘-源表
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceTable *TaskLineageInfo `json:"SourceTable,omitnil,omitempty" name:"SourceTable"`

	// 表血缘-目标表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetTable *TaskLineageInfo `json:"TargetTable,omitnil,omitempty" name:"TargetTable"`
}

type TaskLockStatus struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 持锁者
	Locker *string `json:"Locker,omitnil,omitempty" name:"Locker"`

	// 当前操作用户是否为持锁者，1表示为持锁者，0表示为不为持锁者
	IsLocker *int64 `json:"IsLocker,omitnil,omitempty" name:"IsLocker"`

	// 是否可以抢锁，1表示可以抢锁，0表示不可以抢锁
	IsRob *int64 `json:"IsRob,omitnil,omitempty" name:"IsRob"`
}

// Predefined struct for user
type TaskLogRequestParams struct {
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 起始时间戳，单位毫秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳，单位毫秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 拉取日志数量，默认100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志排序 desc 倒序 asc 顺序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type TaskLogRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 起始时间戳，单位毫秒
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳，单位毫秒
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 拉取日志数量，默认100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 日志排序 desc 倒序 asc 顺序
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`

	// 实时任务 201   离线任务 202  默认实时任务
	TaskType *uint64 `json:"TaskType,omitnil,omitempty" name:"TaskType"`
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
	// 任务详细日志
	LogContentList []*LogContent `json:"LogContentList,omitnil,omitempty" name:"LogContentList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type TaskOpsDto struct {
	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 虚拟任务id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskId *string `json:"VirtualTaskId,omitnil,omitempty" name:"VirtualTaskId"`

	// 虚拟任务标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualFlag *bool `json:"VirtualFlag,omitnil,omitempty" name:"VirtualFlag"`

	// 任务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务为虚拟任务时，任务所在的真实工作流Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealWorkflowId *string `json:"RealWorkflowId,omitnil,omitempty" name:"RealWorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 文件夹名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastUpdate *string `json:"LastUpdate,omitnil,omitempty" name:"LastUpdate"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`

	// 责任人用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	InChargeId *string `json:"InChargeId,omitnil,omitempty" name:"InChargeId"`

	// 调度生效日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 调度结束日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 执行时间左闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 周期类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 步长
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *uint64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 调度cron表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 延时执行时间，unit=分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 开始执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupTime *uint64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 重试等待时间, unit=分钟
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryWait *uint64 `json:"RetryWait,omitnil,omitempty" name:"RetryWait"`

	// 是否可重试，1 代表可以重试
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryAble *uint64 `json:"RetryAble,omitnil,omitempty" name:"RetryAble"`

	// 调度配置-弹性周期配置，小时/周/月/年调度才有，小时任务指定每天的0点3点4点跑，则为'0,3,4'
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 运行次数限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	TryLimit *uint64 `json:"TryLimit,omitnil,omitempty" name:"TryLimit"`

	// 运行优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunPriority *uint64 `json:"RunPriority,omitnil,omitempty" name:"RunPriority"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *TaskTypeOpsDto `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 指定的运行节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	BrokerIp *string `json:"BrokerIp,omitnil,omitempty" name:"BrokerIp"`

	// 集群name
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitnil,omitempty" name:"ClusterId"`

	// 最小数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MinDateTime *string `json:"MinDateTime,omitnil,omitempty" name:"MinDateTime"`

	// 最大数据时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxDateTime *string `json:"MaxDateTime,omitnil,omitempty" name:"MaxDateTime"`

	// 运行耗时超时时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionTTL *int64 `json:"ExecutionTTL,omitnil,omitempty" name:"ExecutionTTL"`

	// 自依赖类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 左侧坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	LeftCoordinate *float64 `json:"LeftCoordinate,omitnil,omitempty" name:"LeftCoordinate"`

	// 顶部坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopCoordinate *float64 `json:"TopCoordinate,omitnil,omitempty" name:"TopCoordinate"`

	// 任务备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Notes *string `json:"Notes,omitnil,omitempty" name:"Notes"`

	// 任务初始化策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`

	// 计算队列
	// 注意：此字段可能返回 null，表示取不到有效值。
	YarnQueue *string `json:"YarnQueue,omitnil,omitempty" name:"YarnQueue"`

	// 最新调度提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastSchedulerCommitTime *string `json:"LastSchedulerCommitTime,omitnil,omitempty" name:"LastSchedulerCommitTime"`

	// 按cron表达式计算的任务开始执行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	NormalizedJobStartTime *string `json:"NormalizedJobStartTime,omitnil,omitempty" name:"NormalizedJobStartTime"`

	// 调度计划描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitnil,omitempty" name:"SchedulerDesc"`

	// 计算资源组
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceGroup *string `json:"ResourceGroup,omitnil,omitempty" name:"ResourceGroup"`

	// 任务创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 任务依赖类型 and、or
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyRel *string `json:"DependencyRel,omitnil,omitempty" name:"DependencyRel"`

	// 任务工作流依赖 yes、no
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 事件监听配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventListenerConfig *string `json:"EventListenerConfig,omitnil,omitempty" name:"EventListenerConfig"`

	// 事件驱动配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventPublisherConfig *string `json:"EventPublisherConfig,omitnil,omitempty" name:"EventPublisherConfig"`

	// 虚拟任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualTaskStatus *string `json:"VirtualTaskStatus,omitnil,omitempty" name:"VirtualTaskStatus"`

	// 任务依赖边详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskLinkInfo *LinkOpsDto `json:"TaskLinkInfo,omitnil,omitempty" name:"TaskLinkInfo"`

	// 任务产品类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 主账户userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnId *string `json:"OwnId,omitnil,omitempty" name:"OwnId"`

	// 用户userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 租户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TenantId *string `json:"TenantId,omitnil,omitempty" name:"TenantId"`

	// 更新人昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUser *string `json:"UpdateUser,omitnil,omitempty" name:"UpdateUser"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 更新人userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateUserId *string `json:"UpdateUserId,omitnil,omitempty" name:"UpdateUserId"`

	// 任务类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeId *int64 `json:"TaskTypeId,omitnil,omitempty" name:"TaskTypeId"`

	// 任务类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskTypeDesc *string `json:"TaskTypeDesc,omitnil,omitempty" name:"TaskTypeDesc"`

	// 是否展示工作流
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowWorkflow *bool `json:"ShowWorkflow,omitnil,omitempty" name:"ShowWorkflow"`

	// 首次提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 首次运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstRunTime *string `json:"FirstRunTime,omitnil,omitempty" name:"FirstRunTime"`

	// 调度描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScheduleDesc *string `json:"ScheduleDesc,omitnil,omitempty" name:"ScheduleDesc"`

	// 周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleNum *int64 `json:"CycleNum,omitnil,omitempty" name:"CycleNum"`

	// 表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Crontab *string `json:"Crontab,omitnil,omitempty" name:"Crontab"`

	// 开始日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartDate *string `json:"StartDate,omitnil,omitempty" name:"StartDate"`

	// 结束日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndDate *string `json:"EndDate,omitnil,omitempty" name:"EndDate"`

	// 周期单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleUnit *string `json:"CycleUnit,omitnil,omitempty" name:"CycleUnit"`

	// 初始化策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	InitStrategy *string `json:"InitStrategy,omitnil,omitempty" name:"InitStrategy"`

	// 层级
	// 注意：此字段可能返回 null，表示取不到有效值。
	Layer *string `json:"Layer,omitnil,omitempty" name:"Layer"`

	// 来源数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceId *string `json:"SourceServiceId,omitnil,omitempty" name:"SourceServiceId"`

	// 来源数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceServiceType *string `json:"SourceServiceType,omitnil,omitempty" name:"SourceServiceType"`

	// 目标数据源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceId *string `json:"TargetServiceId,omitnil,omitempty" name:"TargetServiceId"`

	// 目标数据源类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetServiceType *string `json:"TargetServiceType,omitnil,omitempty" name:"TargetServiceType"`

	// 子任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	TasksStr *string `json:"TasksStr,omitnil,omitempty" name:"TasksStr"`

	// 任务版本是否已提交
	// 注意：此字段可能返回 null，表示取不到有效值。
	Submit *bool `json:"Submit,omitnil,omitempty" name:"Submit"`

	// 资源组id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupId *string `json:"ExecutorGroupId,omitnil,omitempty" name:"ExecutorGroupId"`

	// 资源组名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutorGroupName *string `json:"ExecutorGroupName,omitnil,omitempty" name:"ExecutorGroupName"`

	// 任务扩展信息(目前返沪离线同步的任务详情)
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskExtInfo *string `json:"TaskExtInfo,omitnil,omitempty" name:"TaskExtInfo"`

	// 任务绑定的事件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EventListenerInfos []*AiOpsEventListenerDTO `json:"EventListenerInfos,omitnil,omitempty" name:"EventListenerInfos"`

	// 脚本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptInfo *AiopsScriptInfo `json:"ScriptInfo,omitnil,omitempty" name:"ScriptInfo"`

	// DLC资源配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DLCResourceConfig *AiopsDLCResourceConfigDto `json:"DLCResourceConfig,omitnil,omitempty" name:"DLCResourceConfig"`

	// 父任务simple信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentTaskInfos *AiopsSimpleTaskDto `json:"ParentTaskInfos,omitnil,omitempty" name:"ParentTaskInfos"`

	// 资源获取标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtResourceFlag *ExtResourceFlagDto `json:"ExtResourceFlag,omitnil,omitempty" name:"ExtResourceFlag"`

	// 父任务simple信息(新)
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewParentTaskInfos []*AiopsSimpleTaskDto `json:"NewParentTaskInfos,omitnil,omitempty" name:"NewParentTaskInfos"`
}

type TaskScriptContent struct {
	// 脚本内容 base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptContent *string `json:"ScriptContent,omitnil,omitempty" name:"ScriptContent"`
}

type TaskSubmitPreCheckDetailInfo struct {
	// 任务编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 项目编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 责任人编号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InChargeId *string `json:"InChargeId,omitnil,omitempty" name:"InChargeId"`

	// 责任人名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InCharge *string `json:"InCharge,omitnil,omitempty" name:"InCharge"`
}

type TaskTag struct {
	// 标签名称
	TagName *string `json:"TagName,omitnil,omitempty" name:"TagName"`

	// 标签值列表
	TagValues []*string `json:"TagValues,omitnil,omitempty" name:"TagValues"`
}

type TaskTypeCnt struct {
	// 统计值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *string `json:"Number,omitnil,omitempty" name:"Number"`

	// 任务类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeName *string `json:"TypeName,omitnil,omitempty" name:"TypeName"`
}

type TaskTypeMap struct {
	// key
	Key *int64 `json:"Key,omitnil,omitempty" name:"Key"`

	// value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TaskTypeOpsDto struct {
	// 任务类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeDesc *string `json:"TypeDesc,omitnil,omitempty" name:"TypeDesc"`

	// 任务类型id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeId *int64 `json:"TypeId,omitnil,omitempty" name:"TypeId"`

	// 任务类型归类
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeSort *string `json:"TypeSort,omitnil,omitempty" name:"TypeSort"`
}

type TaskVersionInstance struct {
	// 实例版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceVersion *int64 `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 实例描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	VersionDesc *string `json:"VersionDesc,omitnil,omitempty" name:"VersionDesc"`

	// 0, "新增"，1, "修改"
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChangeType *int64 `json:"ChangeType,omitnil,omitempty" name:"ChangeType"`

	// 版本提交人UIN
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubmitterUin *string `json:"SubmitterUin,omitnil,omitempty" name:"SubmitterUin"`

	// 提交日期
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceDate *string `json:"InstanceDate,omitnil,omitempty" name:"InstanceDate"`

	// 0, "未启用"，1, "启用(生产态)"
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatus *int64 `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

type ThresholdValue struct {
	// 阈值类型  1.低阈值  2.高阈值   3.普通阈值  4.枚举值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueType *uint64 `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// 阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TopTableStat struct {
	// 告警表列表
	AlarmTables []*TopTableStatItem `json:"AlarmTables,omitnil,omitempty" name:"AlarmTables"`

	// 阻塞表列表
	PipelineTables []*TopTableStatItem `json:"PipelineTables,omitnil,omitempty" name:"PipelineTables"`
}

type TopTableStatItem struct {
	// 表Id
	TableId *string `json:"TableId,omitnil,omitempty" name:"TableId"`

	// 表名
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableName *string `json:"TableName,omitnil,omitempty" name:"TableName"`

	// 数
	Cnt *uint64 `json:"Cnt,omitnil,omitempty" name:"Cnt"`
}

// Predefined struct for user
type TriggerDsEventRequestParams struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件实例信息
	EventCaseList []*EventCaseDTO `json:"EventCaseList,omitnil,omitempty" name:"EventCaseList"`

	// 事件实例信息(连续时间)
	EventBatchCaseList []*EventBatchCaseDTO `json:"EventBatchCaseList,omitnil,omitempty" name:"EventBatchCaseList"`
}

type TriggerDsEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 事件实例信息
	EventCaseList []*EventCaseDTO `json:"EventCaseList,omitnil,omitempty" name:"EventCaseList"`

	// 事件实例信息(连续时间)
	EventBatchCaseList []*EventBatchCaseDTO `json:"EventBatchCaseList,omitnil,omitempty" name:"EventBatchCaseList"`
}

func (r *TriggerDsEventRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerDsEventRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "EventCaseList")
	delete(f, "EventBatchCaseList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerDsEventRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerDsEventResponseParams struct {
	// 操作结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchOpsDTO `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerDsEventResponse struct {
	*tchttp.BaseResponse
	Response *TriggerDsEventResponseParams `json:"Response"`
}

func (r *TriggerDsEventResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerDsEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerEventRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 案例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 时间格式：如果选择触发时间：2022年6月21，则设置为20220621
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type TriggerEventRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 案例名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 时间格式：如果选择触发时间：2022年6月21，则设置为20220621
	Dimension *string `json:"Dimension,omitnil,omitempty" name:"Dimension"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	Data *BatchReturn `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type UnlockIntegrationTaskRequest struct {
	*tchttp.BaseRequest
	
	// 任务id
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
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
	Data *bool `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type UpdateWorkflowOwnerRequestParams struct {
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流Ids
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 责任人，多个以';'号分割
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 责任人UserId，多个以';'号分割
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`
}

type UpdateWorkflowOwnerRequest struct {
	*tchttp.BaseRequest
	
	// 项目Id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流Ids
	WorkflowIds []*string `json:"WorkflowIds,omitnil,omitempty" name:"WorkflowIds"`

	// 责任人，多个以';'号分割
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 责任人UserId，多个以';'号分割
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`
}

func (r *UpdateWorkflowOwnerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowOwnerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProjectId")
	delete(f, "WorkflowIds")
	delete(f, "Owner")
	delete(f, "OwnerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateWorkflowOwnerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateWorkflowOwnerResponseParams struct {
	// 响应数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *BatchOperationOpsDto `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateWorkflowOwnerResponse struct {
	*tchttp.BaseResponse
	Response *UpdateWorkflowOwnerResponseParams `json:"Response"`
}

func (r *UpdateWorkflowOwnerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateWorkflowOwnerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadContentRequestParams struct {
	// 脚本上传信息
	ScriptRequestInfo *ScriptRequestInfo `json:"ScriptRequestInfo,omitnil,omitempty" name:"ScriptRequestInfo"`

	// 请求来源，WEB 前端；CLIENT 客户端
	RequestFromSource *string `json:"RequestFromSource,omitnil,omitempty" name:"RequestFromSource"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type UploadContentRequest struct {
	*tchttp.BaseRequest
	
	// 脚本上传信息
	ScriptRequestInfo *ScriptRequestInfo `json:"ScriptRequestInfo,omitnil,omitempty" name:"ScriptRequestInfo"`

	// 请求来源，WEB 前端；CLIENT 客户端
	RequestFromSource *string `json:"RequestFromSource,omitnil,omitempty" name:"RequestFromSource"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *UploadContentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadContentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScriptRequestInfo")
	delete(f, "RequestFromSource")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadContentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadContentResponseParams struct {
	// 脚本信息响应
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScriptInfo *ScriptInfoResponse `json:"ScriptInfo,omitnil,omitempty" name:"ScriptInfo"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadContentResponse struct {
	*tchttp.BaseResponse
	Response *UploadContentResponseParams `json:"Response"`
}

func (r *UploadContentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadResourceRequestParams struct {
	// 资源上传请求信息
	UploadResourceRequestInfo *UploadResourceRequestInfo `json:"UploadResourceRequestInfo,omitnil,omitempty" name:"UploadResourceRequestInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

type UploadResourceRequest struct {
	*tchttp.BaseRequest
	
	// 资源上传请求信息
	UploadResourceRequestInfo *UploadResourceRequestInfo `json:"UploadResourceRequestInfo,omitnil,omitempty" name:"UploadResourceRequestInfo"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`
}

func (r *UploadResourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadResourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UploadResourceRequestInfo")
	delete(f, "ProjectId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadResourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UploadResourceRequestInfo struct {
	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 资源路径
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 桶名称
	BucketName *string `json:"BucketName,omitnil,omitempty" name:"BucketName"`

	// 所属地区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 是否为新资源
	NewFile *bool `json:"NewFile,omitnil,omitempty" name:"NewFile"`

	// 资源列表
	FileList []*string `json:"FileList,omitnil,omitempty" name:"FileList"`

	// 资源大小列表
	FileSizeList []*string `json:"FileSizeList,omitnil,omitempty" name:"FileSizeList"`

	// File Md5（适配私有化，公有云可以不传）
	FileMd5 *string `json:"FileMd5,omitnil,omitempty" name:"FileMd5"`
}

// Predefined struct for user
type UploadResourceResponseParams struct {
	// 资源文件信息列表
	Data []*UserFileDTONew `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadResourceResponse struct {
	*tchttp.BaseResponse
	Response *UploadResourceResponseParams `json:"Response"`
}

func (r *UploadResourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadResourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserFileDTONew struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型，如 jar zip 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// 文件上传类型，资源管理为 resource
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件MD5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Value *string `json:"Md5Value,omitnil,omitempty" name:"Md5Value"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文件大小，单位为字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 本地路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// 本地临时路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalTempPath *string `json:"LocalTempPath,omitnil,omitempty" name:"LocalTempPath"`

	// 远程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`

	// 文件拥有者名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 文件拥有者uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 文件深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathDepth *string `json:"PathDepth,omitnil,omitempty" name:"PathDepth"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 本地临时压缩文件绝对路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZipPath *string `json:"ZipPath,omitnil,omitempty" name:"ZipPath"`

	// 文件所属存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 文件所属存储桶的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 删除用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteName *string `json:"DeleteName,omitnil,omitempty" name:"DeleteName"`

	// 删除用户id
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteOwner *string `json:"DeleteOwner,omitnil,omitempty" name:"DeleteOwner"`

	// 操作者id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 操作者名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 全路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	FullPath *string `json:"FullPath,omitnil,omitempty" name:"FullPath"`
}

type UserFileInfo struct {
	// 资源ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型，如 jar zip 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileExtensionType *string `json:"FileExtensionType,omitnil,omitempty" name:"FileExtensionType"`

	// 文件上传类型，资源管理为 resource
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 文件MD5值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5Value *string `json:"Md5Value,omitnil,omitempty" name:"Md5Value"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文件大小，单位为字节
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *uint64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 本地路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalPath *string `json:"LocalPath,omitnil,omitempty" name:"LocalPath"`

	// 本地临时路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	LocalTempPath *string `json:"LocalTempPath,omitnil,omitempty" name:"LocalTempPath"`

	// 远程路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	RemotePath *string `json:"RemotePath,omitnil,omitempty" name:"RemotePath"`

	// 文件拥有者名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerName *string `json:"OwnerName,omitnil,omitempty" name:"OwnerName"`

	// 文件拥有者uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 文件深度
	// 注意：此字段可能返回 null，表示取不到有效值。
	PathDepth *int64 `json:"PathDepth,omitnil,omitempty" name:"PathDepth"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo []*ParamInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 本地临时压缩文件绝对路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZipPath *string `json:"ZipPath,omitnil,omitempty" name:"ZipPath"`

	// 文件所属存储桶
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 文件所属存储桶的地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteName *string `json:"DeleteName,omitnil,omitempty" name:"DeleteName"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteOwner *string `json:"DeleteOwner,omitnil,omitempty" name:"DeleteOwner"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatorName *string `json:"OperatorName,omitnil,omitempty" name:"OperatorName"`

	// 附加信息 base64编码
	// 注意：此字段可能返回 null，表示取不到有效值。
	EncodeExtraInfo *string `json:"EncodeExtraInfo,omitnil,omitempty" name:"EncodeExtraInfo"`
}

type WeightInfo struct {
	// 权重
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 维度类型 1：准确性，2：唯一性，3：完整性，4：一致性，5：及时性，6：有效性
	QualityDim *int64 `json:"QualityDim,omitnil,omitempty" name:"QualityDim"`
}

type WorkFlowExecuteDto struct {
	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 工作流运行状态 0：等待运行、1：运行中、2：运行完成、3：运行出错
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type WorkFlowExecuteDtoByPage struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// data
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkFlowExecuteDto `json:"Items,omitnil,omitempty" name:"Items"`

	// 分页大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type Workflow struct {
	// 工作流id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 责任人Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`

	// 项目id
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流所属用户分组id 若有多个,分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称  若有多个,分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`
}

type WorkflowCanvasOpsDto struct {
	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流详情描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 所属文件夹id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 所属文件夹ids
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderIds []*string `json:"FolderIds,omitnil,omitempty" name:"FolderIds"`

	// 任务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*TaskOpsDto `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务依赖边列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Links []*LinkOpsDto `json:"Links,omitnil,omitempty" name:"Links"`

	// 工作流所属用户分组id,若有多个分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupId *string `json:"UserGroupId,omitnil,omitempty" name:"UserGroupId"`

	// 工作流所属用户分组名称,若有多个分号隔开: a;b;c
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserGroupName *string `json:"UserGroupName,omitnil,omitempty" name:"UserGroupName"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 责任人UserId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`
}

type WorkflowExtOpsDto struct {
	// 任务数量count
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskCount *uint64 `json:"TaskCount,omitnil,omitempty" name:"TaskCount"`

	// 文件名
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderName *string `json:"FolderName,omitnil,omitempty" name:"FolderName"`

	// 工作流id
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkFlowId *string `json:"WorkFlowId,omitnil,omitempty" name:"WorkFlowId"`

	// 责任人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 责任人userId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OwnerId *string `json:"OwnerId,omitnil,omitempty" name:"OwnerId"`

	// 项目id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 项目标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectIdent *string `json:"ProjectIdent,omitnil,omitempty" name:"ProjectIdent"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName *string `json:"ProjectName,omitnil,omitempty" name:"ProjectName"`

	// 工作流描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkFlowDesc *string `json:"WorkFlowDesc,omitnil,omitempty" name:"WorkFlowDesc"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkFlowName *string `json:"WorkFlowName,omitnil,omitempty" name:"WorkFlowName"`

	// 工作流文件id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FolderId *string `json:"FolderId,omitnil,omitempty" name:"FolderId"`

	// 工作流状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 工作流创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 最近更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type WorkflowExtOpsDtoPage struct {
	// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 记录列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Items []*WorkflowExtOpsDto `json:"Items,omitnil,omitempty" name:"Items"`
}

type WorkflowScheduleDtoDs struct {
	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId *string `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creater *string `json:"Creater,omitnil,omitempty" name:"Creater"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 延迟时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *int64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupTime *int64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 任务依赖
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定运行时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 调度周期时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 调度周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *int64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// Cron表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 实例初始化策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`

	// 工作流依赖
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 调度计划
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitnil,omitempty" name:"SchedulerDesc"`

	// 工作流首次提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 工作流最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestSubmitTime *string `json:"LatestSubmitTime,omitnil,omitempty" name:"LatestSubmitTime"`
}

type WorkflowSchedulerOpsDto struct {
	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 延迟时间, unit=minute
	// 注意：此字段可能返回 null，表示取不到有效值。
	DelayTime *uint64 `json:"DelayTime,omitnil,omitempty" name:"DelayTime"`

	// 启动时间,unit=minute
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartupTime *uint64 `json:"StartupTime,omitnil,omitempty" name:"StartupTime"`

	// 任务自依赖类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelfDepend *string `json:"SelfDepend,omitnil,omitempty" name:"SelfDepend"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 指定运行时间, 指定时间：如周一：1
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskAction *string `json:"TaskAction,omitnil,omitempty" name:"TaskAction"`

	// 调度周期类型，时间单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleType *string `json:"CycleType,omitnil,omitempty" name:"CycleType"`

	// 调度周期，间隔步长 unit=minute
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleStep *uint64 `json:"CycleStep,omitnil,omitempty" name:"CycleStep"`

	// 调度cron表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrontabExpression *string `json:"CrontabExpression,omitnil,omitempty" name:"CrontabExpression"`

	// 执行时间左闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionStartTime *string `json:"ExecutionStartTime,omitnil,omitempty" name:"ExecutionStartTime"`

	// 执行时间右闭区间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionEndTime *string `json:"ExecutionEndTime,omitnil,omitempty" name:"ExecutionEndTime"`

	// 任务实例初始化策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceInitStrategy *string `json:"InstanceInitStrategy,omitnil,omitempty" name:"InstanceInitStrategy"`

	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流自依赖
	// 注意：此字段可能返回 null，表示取不到有效值。
	DependencyWorkflow *string `json:"DependencyWorkflow,omitnil,omitempty" name:"DependencyWorkflow"`

	// 调度计划释义
	// 注意：此字段可能返回 null，表示取不到有效值。
	SchedulerDesc *string `json:"SchedulerDesc,omitnil,omitempty" name:"SchedulerDesc"`

	// 工作流首次提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstSubmitTime *string `json:"FirstSubmitTime,omitnil,omitempty" name:"FirstSubmitTime"`

	// 工作流最近提交时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LatestSubmitTime *string `json:"LatestSubmitTime,omitnil,omitempty" name:"LatestSubmitTime"`
}

type WorkflowTaskCountOpsDto struct {
	// 工作流任务数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 任务类型维度统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	TypeCount []*PairDto `json:"TypeCount,omitnil,omitempty" name:"TypeCount"`

	// 任务周期类型维度统计
	// 注意：此字段可能返回 null，表示取不到有效值。
	CycleCount []*PairDto `json:"CycleCount,omitnil,omitempty" name:"CycleCount"`
}