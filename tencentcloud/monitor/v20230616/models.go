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

package v20230616

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIWorkbenchSREDigitalTwinTask struct {
	// 任务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务配置
	TaskConfig *string `json:"TaskConfig,omitnil,omitempty" name:"TaskConfig"`

	// 唯一标识
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 所属数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`
}

type AIWorkbenchSREDigitalTwinTaskList struct {
	// 任务列表
	Tasks []*AIWorkbenchSREDigitalTwinTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// 任务总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type AIWorkbenchSREDigitalTwinWorkLog struct {
	// 唯一标识符
	ID *uint64 `json:"ID,omitnil,omitempty" name:"ID"`

	// 创建时间
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// 所属数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 所属数字分身任务ID
	TaskID *uint64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`

	// 分析时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 分析状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 分析结果摘要
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 所属任务名称
	TaskName *string `json:"TaskName,omitnil,omitempty" name:"TaskName"`

	// 所属任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`
}

type AIWorkbenchSREDigitalTwinWorkLogDetail struct {
	// 工作日志详细内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 工作日志任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 工作日志相关对话ID
	DialogID *int64 `json:"DialogID,omitnil,omitempty" name:"DialogID"`
}

type AIWorkbenchSREDigitalTwinWorkLogList struct {
	// 工作日志列表
	WorkLogs []*AIWorkbenchSREDigitalTwinWorkLog `json:"WorkLogs,omitnil,omitempty" name:"WorkLogs"`

	// 总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`
}

type AgentInfo struct {
	// <p>Agent ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent 名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent 描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent 分类</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>状态: draft/configured/running/standby/disabled</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>关联技能 ID 列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>关联的资源地图 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>关联的mcp id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>资源标签</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CamTags []*Tag `json:"CamTags,omitnil,omitempty" name:"CamTags"`

	// <p>agent运行时所需环境变量</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type AlarmLable struct {
	// label name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// label value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AlarmNotifyHistory struct {
	// 通知的唯一ID
	NotifyId *string `json:"NotifyId,omitnil,omitempty" name:"NotifyId"`

	// 告警策略ID
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`

	// 告警周期iD
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 通知时间 unix秒级时间戳
	NotifyTime *int64 `json:"NotifyTime,omitnil,omitempty" name:"NotifyTime"`

	// 触发时间 unix秒级时间戳
	TriggerTime *int64 `json:"TriggerTime,omitnil,omitempty" name:"TriggerTime"`

	// 告警级别 None 非分级告警级别; Note 提示级别; Warn 严重级别; Serious 紧急级别
	TriggerLevel *string `json:"TriggerLevel,omitnil,omitempty" name:"TriggerLevel"`

	// 告警内容
	AlarmContent *string `json:"AlarmContent,omitnil,omitempty" name:"AlarmContent"`

	// 告警对象
	AlarmObject *string `json:"AlarmObject,omitnil,omitempty" name:"AlarmObject"`

	// 本次告警通知涉及到的渠道合集
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelSet []*string `json:"ChannelSet,omitnil,omitempty" name:"ChannelSet"`

	// 渠道的接收人信息
	ChannelsReceivers []*ChannelsReceivers `json:"ChannelsReceivers,omitnil,omitempty" name:"ChannelsReceivers"`

	// 告警策略名称
	PolicyName *string `json:"PolicyName,omitnil,omitempty" name:"PolicyName"`

	// Prometheus实例ID, 仅当 MT_PROME 时有效
	PromeInstanceID *string `json:"PromeInstanceID,omitnil,omitempty" name:"PromeInstanceID"`

	// Prometheus实例所在的地域, 仅当 MT_PROME 时有效
	PromeInstanceRegion *string `json:"PromeInstanceRegion,omitnil,omitempty" name:"PromeInstanceRegion"`

	// 通知模板相关的配置信息
	Notices []*NotifyRelatedNotice `json:"Notices,omitnil,omitempty" name:"Notices"`

	// 告警触发状态  Trigger 告警状态触发; Recovery 告警状态恢复
	TriggerStatus *string `json:"TriggerStatus,omitnil,omitempty" name:"TriggerStatus"`

	// 与当前Prometheus通知历史相关控制台页面地址，仅当 MR_PROME 时有效
	PromeConsoleURL *string `json:"PromeConsoleURL,omitnil,omitempty" name:"PromeConsoleURL"`

	// 告警的lable
	Labels []*AlarmLable `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ArtifactInfo struct {
	// <p>产物 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>产物名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>物理类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MimeType *string `json:"MimeType,omitnil,omitempty" name:"MimeType"`

	// <p>文件大小(字节)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SizeBytes *int64 `json:"SizeBytes,omitnil,omitempty" name:"SizeBytes"`

	// <p>是否公共</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsGlobal *bool `json:"IsGlobal,omitnil,omitempty" name:"IsGlobal"`

	// <p>创建时间 Unix 秒时间戳</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *int64 `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>修改时间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedAt *int64 `json:"UpdatedAt,omitnil,omitempty" name:"UpdatedAt"`

	// <p>产生该制品的 Agent ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>产生该制品的 Skill ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>用于解析调用下载接口</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoragePath *string `json:"StoragePath,omitnil,omitempty" name:"StoragePath"`
}

// Predefined struct for user
type CancelAIWorkbenchChatRequestParams struct {
	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type CancelAIWorkbenchChatRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话id</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *CancelAIWorkbenchChatRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAIWorkbenchChatRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelAIWorkbenchChatRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelAIWorkbenchChatResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelAIWorkbenchChatResponse struct {
	*tchttp.BaseResponse
	Response *CancelAIWorkbenchChatResponseParams `json:"Response"`
}

func (r *CancelAIWorkbenchChatResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelAIWorkbenchChatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelsReceivers struct {
	// 通知渠道名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ChannelName *string `json:"ChannelName,omitnil,omitempty" name:"ChannelName"`

	// 接收者
	// 注意：此字段可能返回 null，表示取不到有效值。
	Receivers []*string `json:"Receivers,omitnil,omitempty" name:"Receivers"`

	// 发送结果,0-无效,1-成功,2-失败,3-无需发送
	// 注意：此字段可能返回 null，表示取不到有效值。
	SendStatus *string `json:"SendStatus,omitnil,omitempty" name:"SendStatus"`
}

type ContentBlockInfo struct {
	// <p>类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>数据内容</p>
	Data *string `json:"Data,omitnil,omitempty" name:"Data"`
}

// Predefined struct for user
type CreateAIWorkbenchAgentRequestParams struct {
	// <p>Agent 名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent 描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent 分类</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent 标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent 提示词</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>关联技能 ID 列表</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>来源: builtin / custom</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>关联的资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>关联的mcp工具</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>资源标签</p>
	CamTags []*Tag `json:"CamTags,omitnil,omitempty" name:"CamTags"`

	// <p>agent运行时环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type CreateAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent 名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent 描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent 分类</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent 标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent 提示词</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>关联技能 ID 列表</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>来源: builtin / custom</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>关联的资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>关联的mcp工具</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>资源标签</p>
	CamTags []*Tag `json:"CamTags,omitnil,omitempty" name:"CamTags"`

	// <p>agent运行时环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *CreateAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Category")
	delete(f, "Tags")
	delete(f, "Instruction")
	delete(f, "SkillIds")
	delete(f, "Source")
	delete(f, "ResourceMapId")
	delete(f, "MCPIds")
	delete(f, "CamTags")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIWorkbenchAgentResponseParams struct {
	// <p>Agent ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *CreateAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIWorkbenchTaskRequestParams struct {
	// <p>任务名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>任务描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>关联 Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>提示词模板</p>
	PromptTemplate *string `json:"PromptTemplate,omitnil,omitempty" name:"PromptTemplate"`

	// <p>输出格式: markdown / json</p>
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// <p>触发类型: manual / cron / webhook</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Cron 表达式</p>
	CronExpr *string `json:"CronExpr,omitnil,omitempty" name:"CronExpr"`

	// <p>Cron 时区</p>
	CronTimezone *string `json:"CronTimezone,omitnil,omitempty" name:"CronTimezone"`

	// <p>关联资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>技能 ID 列表</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>MCP 端点 ID 列表</p>
	McpEndpointIds []*string `json:"McpEndpointIds,omitnil,omitempty" name:"McpEndpointIds"`

	// <p>超时时间(秒)</p>
	TimeoutSec *int64 `json:"TimeoutSec,omitnil,omitempty" name:"TimeoutSec"`

	// <p>重试次数</p>
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CreateAIWorkbenchTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>任务描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>关联 Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>提示词模板</p>
	PromptTemplate *string `json:"PromptTemplate,omitnil,omitempty" name:"PromptTemplate"`

	// <p>输出格式: markdown / json</p>
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// <p>触发类型: manual / cron / webhook</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Cron 表达式</p>
	CronExpr *string `json:"CronExpr,omitnil,omitempty" name:"CronExpr"`

	// <p>Cron 时区</p>
	CronTimezone *string `json:"CronTimezone,omitnil,omitempty" name:"CronTimezone"`

	// <p>关联资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>技能 ID 列表</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>MCP 端点 ID 列表</p>
	McpEndpointIds []*string `json:"McpEndpointIds,omitnil,omitempty" name:"McpEndpointIds"`

	// <p>超时时间(秒)</p>
	TimeoutSec *int64 `json:"TimeoutSec,omitnil,omitempty" name:"TimeoutSec"`

	// <p>重试次数</p>
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *CreateAIWorkbenchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "AgentId")
	delete(f, "PromptTemplate")
	delete(f, "OutputFormat")
	delete(f, "TriggerType")
	delete(f, "CronExpr")
	delete(f, "CronTimezone")
	delete(f, "ResourceMapId")
	delete(f, "SkillIds")
	delete(f, "McpEndpointIds")
	delete(f, "TimeoutSec")
	delete(f, "RetryCount")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAIWorkbenchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAIWorkbenchTaskResponseParams struct {
	// <p>任务 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAIWorkbenchTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAIWorkbenchTaskResponseParams `json:"Response"`
}

func (r *CreateAIWorkbenchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAIWorkbenchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDispenseExternalRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云监控对外命名空间
	ExtNamespace *string `json:"ExtNamespace,omitnil,omitempty" name:"ExtNamespace"`

	// 转发目标消信息
	Producer *Producer `json:"Producer,omitnil,omitempty" name:"Producer"`

	// 转发部署地域列表
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`

	// 云监控对外指标
	ExtMetrics []*string `json:"ExtMetrics,omitnil,omitempty" name:"ExtMetrics"`

	// 指标统计周期
	Period []*int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 转发过滤条件信息
	DispenseConditions []*DispenseCondition `json:"DispenseConditions,omitnil,omitempty" name:"DispenseConditions"`
}

type CreateDispenseExternalRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云监控对外命名空间
	ExtNamespace *string `json:"ExtNamespace,omitnil,omitempty" name:"ExtNamespace"`

	// 转发目标消信息
	Producer *Producer `json:"Producer,omitnil,omitempty" name:"Producer"`

	// 转发部署地域列表
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`

	// 云监控对外指标
	ExtMetrics []*string `json:"ExtMetrics,omitnil,omitempty" name:"ExtMetrics"`

	// 指标统计周期
	Period []*int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 转发过滤条件信息
	DispenseConditions []*DispenseCondition `json:"DispenseConditions,omitnil,omitempty" name:"DispenseConditions"`
}

func (r *CreateDispenseExternalRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDispenseExternalRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ExtNamespace")
	delete(f, "Producer")
	delete(f, "DispenseRegions")
	delete(f, "ExtMetrics")
	delete(f, "Period")
	delete(f, "DispenseConditions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDispenseExternalRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDispenseExternalRuleResponseParams struct {
	// 转发规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDispenseExternalRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateDispenseExternalRuleResponseParams `json:"Response"`
}

func (r *CreateDispenseExternalRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDispenseExternalRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNoticeContentTmplRequestParams struct {
	// <p>模板名称</p>
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// <p>监控类型</p>
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// <p>模板内容</p>
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// <p>模板语言 en/zh</p>
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`
}

type CreateNoticeContentTmplRequest struct {
	*tchttp.BaseRequest
	
	// <p>模板名称</p>
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// <p>监控类型</p>
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// <p>模板内容</p>
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// <p>模板语言 en/zh</p>
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`
}

func (r *CreateNoticeContentTmplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNoticeContentTmplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TmplName")
	delete(f, "MonitorType")
	delete(f, "TmplContents")
	delete(f, "TmplLanguage")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNoticeContentTmplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNoticeContentTmplResponseParams struct {
	// <p>自定义内容模板ID</p>
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNoticeContentTmplResponse struct {
	*tchttp.BaseResponse
	Response *CreateNoticeContentTmplResponseParams `json:"Response"`
}

func (r *CreateNoticeContentTmplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchAgentRequestParams struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DeleteAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DeleteAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchAgentResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *DeleteAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchTaskRequestParams struct {
	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type DeleteAIWorkbenchTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *DeleteAIWorkbenchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAIWorkbenchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAIWorkbenchTaskResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAIWorkbenchTaskResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAIWorkbenchTaskResponseParams `json:"Response"`
}

func (r *DeleteAIWorkbenchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAIWorkbenchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDispenseExternalRuleRequestParams struct {
	// 需要删除的规则Id
	RuleIdList []*int64 `json:"RuleIdList,omitnil,omitempty" name:"RuleIdList"`
}

type DeleteDispenseExternalRuleRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的规则Id
	RuleIdList []*int64 `json:"RuleIdList,omitnil,omitempty" name:"RuleIdList"`
}

func (r *DeleteDispenseExternalRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDispenseExternalRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDispenseExternalRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDispenseExternalRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDispenseExternalRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDispenseExternalRuleResponseParams `json:"Response"`
}

func (r *DeleteDispenseExternalRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDispenseExternalRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNoticeContentTmplsRequestParams struct {
	// <p>要删除的模板id</p>
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`
}

type DeleteNoticeContentTmplsRequest struct {
	*tchttp.BaseRequest
	
	// <p>要删除的模板id</p>
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`
}

func (r *DeleteNoticeContentTmplsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNoticeContentTmplsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TmplIDs")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNoticeContentTmplsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNoticeContentTmplsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNoticeContentTmplsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNoticeContentTmplsResponseParams `json:"Response"`
}

func (r *DeleteNoticeContentTmplsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNoticeContentTmplsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchAgentRequestParams struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchAgentResponseParams struct {
	// <p>Agent 信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Agent *AgentInfo `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchArtifactRequestParams struct {
	// <p>产物 ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>是否需要下载 URL</p><p><code>1</code> = 需要，<code>0</code> 或不传 = 不需要</p>
	NeedDownloadURL *int64 `json:"NeedDownloadURL,omitnil,omitempty" name:"NeedDownloadURL"`
}

type DescribeAIWorkbenchArtifactRequest struct {
	*tchttp.BaseRequest
	
	// <p>产物 ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`

	// <p>是否需要下载 URL</p><p><code>1</code> = 需要，<code>0</code> 或不传 = 不需要</p>
	NeedDownloadURL *int64 `json:"NeedDownloadURL,omitnil,omitempty" name:"NeedDownloadURL"`
}

func (r *DescribeAIWorkbenchArtifactRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchArtifactRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ArtifactId")
	delete(f, "NeedDownloadURL")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchArtifactRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchArtifactResponseParams struct {
	// <p>产物信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Artifact *ArtifactInfo `json:"Artifact,omitnil,omitempty" name:"Artifact"`

	// <p>COS 预签名下载 URL</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadURL *string `json:"DownloadURL,omitnil,omitempty" name:"DownloadURL"`

	// <p>下载 URL 过期时间（RFC3339 格式）</p>
	DownloadURLExpiredAt *string `json:"DownloadURLExpiredAt,omitnil,omitempty" name:"DownloadURLExpiredAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchArtifactResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchArtifactResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchArtifactResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchArtifactResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchExecutionRequestParams struct {
	// <p>执行 ID</p>
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`
}

type DescribeAIWorkbenchExecutionRequest struct {
	*tchttp.BaseRequest
	
	// <p>执行 ID</p>
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`
}

func (r *DescribeAIWorkbenchExecutionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchExecutionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExecutionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchExecutionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchExecutionResponseParams struct {
	// <p>执行记录</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Execution *ExecutionInfo `json:"Execution,omitnil,omitempty" name:"Execution"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchExecutionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchExecutionResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchExecutionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchExecutionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinTaskListRequestParams struct {
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAIWorkbenchSREDigitalTwinTaskListRequest struct {
	*tchttp.BaseRequest
	
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 数量限制
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinTaskListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinTaskListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TwinID")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSREDigitalTwinTaskListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinTaskListResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身任务列表
	Data *AIWorkbenchSREDigitalTwinTaskList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSREDigitalTwinTaskListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSREDigitalTwinTaskListResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinTaskListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinTaskListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequestParams struct {
	// 工作日志ID
	WorkLogID *int64 `json:"WorkLogID,omitnil,omitempty" name:"WorkLogID"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest struct {
	*tchttp.BaseRequest
	
	// 工作日志ID
	WorkLogID *int64 `json:"WorkLogID,omitnil,omitempty" name:"WorkLogID"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkLogID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSREDigitalTwinWorkLogDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身详细信息
	Data *AIWorkbenchSREDigitalTwinWorkLogDetail `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogListRequestParams struct {
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest struct {
	*tchttp.BaseRequest
	
	// 数字分身ID
	TwinID *uint64 `json:"TwinID,omitnil,omitempty" name:"TwinID"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页限制条数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TwinID")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSREDigitalTwinWorkLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSREDigitalTwinWorkLogListResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身工作日志列表
	Data *AIWorkbenchSREDigitalTwinWorkLogList `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSREDigitalTwinWorkLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSREDigitalTwinWorkLogListResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSREDigitalTwinWorkLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSessionRequestParams struct {
	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

type DescribeAIWorkbenchSessionRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`
}

func (r *DescribeAIWorkbenchSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSessionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSessionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSessionResponseParams struct {
	// <p>会话信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Session *SessionInfo `json:"Session,omitnil,omitempty" name:"Session"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSessionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSessionResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSessionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSkillRequestParams struct {
	// <p>技能 ID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`
}

type DescribeAIWorkbenchSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>技能 ID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`
}

func (r *DescribeAIWorkbenchSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAIWorkbenchSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAIWorkbenchSkillResponseParams struct {
	// <p>技能信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Skill *SkillInfo `json:"Skill,omitnil,omitempty" name:"Skill"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAIWorkbenchSkillResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAIWorkbenchSkillResponseParams `json:"Response"`
}

func (r *DescribeAIWorkbenchSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAIWorkbenchSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesRequestParams struct {
	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 起始时间点，unix秒级时间戳
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// 从 QueryBaseTime 开始，需要查询往前多久的时间，单位秒
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// 分页参数
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// 当监控类型为 MT_QCE 时候需要填写，归属的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 当监控类型为 MT_QCE 时候需要填写， 告警策略类型
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询某个策略的通知历史
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

type DescribeAlarmNotifyHistoriesRequest struct {
	*tchttp.BaseRequest
	
	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// 起始时间点，unix秒级时间戳
	QueryBaseTime *int64 `json:"QueryBaseTime,omitnil,omitempty" name:"QueryBaseTime"`

	// 从 QueryBaseTime 开始，需要查询往前多久的时间，单位秒
	QueryBeforeSeconds *int64 `json:"QueryBeforeSeconds,omitnil,omitempty" name:"QueryBeforeSeconds"`

	// 分页参数
	PageParams *PageByNoParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`

	// 当监控类型为 MT_QCE 时候需要填写，归属的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 当监控类型为 MT_QCE 时候需要填写， 告警策略类型
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 查询某个策略的通知历史
	PolicyId *string `json:"PolicyId,omitnil,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmNotifyHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MonitorType")
	delete(f, "QueryBaseTime")
	delete(f, "QueryBeforeSeconds")
	delete(f, "PageParams")
	delete(f, "Namespace")
	delete(f, "ModelName")
	delete(f, "PolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAlarmNotifyHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAlarmNotifyHistoriesResponseParams struct {
	// 告警历史
	AlarmNotifyHistoryList []*AlarmNotifyHistory `json:"AlarmNotifyHistoryList,omitnil,omitempty" name:"AlarmNotifyHistoryList"`

	// 分页情况
	PageResult *PageByNoResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAlarmNotifyHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAlarmNotifyHistoriesResponseParams `json:"Response"`
}

func (r *DescribeAlarmNotifyHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAlarmNotifyHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDispenseExternalRuleListRequestParams struct {
	// 页数
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页面大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 转发部署地域
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`

	// 关键字搜索规则名
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeDispenseExternalRuleListRequest struct {
	*tchttp.BaseRequest
	
	// 页数
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页面大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 转发部署地域
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`

	// 关键字搜索规则名
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeDispenseExternalRuleListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDispenseExternalRuleListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "DispenseRegions")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDispenseExternalRuleListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDispenseExternalRuleListResponseParams struct {
	// 指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RuleList []*Rule `json:"RuleList,omitnil,omitempty" name:"RuleList"`

	// 列表大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDispenseExternalRuleListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDispenseExternalRuleListResponseParams `json:"Response"`
}

func (r *DescribeDispenseExternalRuleListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDispenseExternalRuleListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDispenseExternalRuleRequestParams struct {
	// 规则id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

type DescribeDispenseExternalRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`
}

func (r *DescribeDispenseExternalRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDispenseExternalRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDispenseExternalRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDispenseExternalRuleResponseParams struct {
	// 规则
	Rule *Rule `json:"Rule,omitnil,omitempty" name:"Rule"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDispenseExternalRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDispenseExternalRuleResponseParams `json:"Response"`
}

func (r *DescribeDispenseExternalRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDispenseExternalRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDispenseRegionRequestParams struct {

}

type DescribeDispenseRegionRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeDispenseRegionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDispenseRegionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDispenseRegionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDispenseRegionResponseParams struct {
	// 转发地域列表
	RegionList []*DispenseRegion `json:"RegionList,omitnil,omitempty" name:"RegionList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDispenseRegionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDispenseRegionResponseParams `json:"Response"`
}

func (r *DescribeDispenseRegionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDispenseRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtMetricRequestParams struct {
	// 对外命名空间
	ExtNamespace *string `json:"ExtNamespace,omitnil,omitempty" name:"ExtNamespace"`
}

type DescribeExtMetricRequest struct {
	*tchttp.BaseRequest
	
	// 对外命名空间
	ExtNamespace *string `json:"ExtNamespace,omitnil,omitempty" name:"ExtNamespace"`
}

func (r *DescribeExtMetricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtMetricRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ExtNamespace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtMetricRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtMetricResponseParams struct {
	// 对外指标
	ExtMetricList []*ExtMetric `json:"ExtMetricList,omitnil,omitempty" name:"ExtMetricList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtMetricResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtMetricResponseParams `json:"Response"`
}

func (r *DescribeExtMetricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtNamespaceRequestParams struct {

}

type DescribeExtNamespaceRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeExtNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeExtNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeExtNamespaceResponseParams struct {
	// 对外命名空间列表
	ExtNamespaceList []*string `json:"ExtNamespaceList,omitnil,omitempty" name:"ExtNamespaceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeExtNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeExtNamespaceResponseParams `json:"Response"`
}

func (r *DescribeExtNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeExtNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaRequestParams struct {
	// kafka地址
	Brokers *string `json:"Brokers,omitnil,omitempty" name:"Brokers"`

	// 转发部署地域列表
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`
}

type DescribeKafkaRequest struct {
	*tchttp.BaseRequest
	
	// kafka地址
	Brokers *string `json:"Brokers,omitnil,omitempty" name:"Brokers"`

	// 转发部署地域列表
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`
}

func (r *DescribeKafkaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Brokers")
	delete(f, "DispenseRegions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKafkaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKafkaResponseParams struct {
	// 连通性列表
	KafkaConnectivityList []*KafkaConnectivity `json:"KafkaConnectivityList,omitnil,omitempty" name:"KafkaConnectivityList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKafkaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKafkaResponseParams `json:"Response"`
}

func (r *DescribeKafkaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKafkaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeContentTmplRequestParams struct {
	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 指定模板ID查询，查询参数都为空则默认查询账号下所有模板
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`

	// 指定模板名称查询，查询参数都为空则默认查询账号下所有模板
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 指定通知模板ID查询，查询参数都为空则默认查询账号下所有模板
	NoticeID *string `json:"NoticeID,omitnil,omitempty" name:"NoticeID"`

	// 模板语言 en/zh 缺省不过滤
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

type DescribeNoticeContentTmplRequest struct {
	*tchttp.BaseRequest
	
	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 指定模板ID查询，查询参数都为空则默认查询账号下所有模板
	TmplIDs []*string `json:"TmplIDs,omitnil,omitempty" name:"TmplIDs"`

	// 指定模板名称查询，查询参数都为空则默认查询账号下所有模板
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 指定通知模板ID查询，查询参数都为空则默认查询账号下所有模板
	NoticeID *string `json:"NoticeID,omitnil,omitempty" name:"NoticeID"`

	// 模板语言 en/zh 缺省不过滤
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`

	// 监控类型
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`
}

func (r *DescribeNoticeContentTmplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeContentTmplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "TmplIDs")
	delete(f, "TmplName")
	delete(f, "NoticeID")
	delete(f, "TmplLanguage")
	delete(f, "MonitorType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNoticeContentTmplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNoticeContentTmplResponseParams struct {
	// 自定义通知内容模板
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoticeContentTmpls []*NoticeContentTmpl `json:"NoticeContentTmpls,omitnil,omitempty" name:"NoticeContentTmpls"`

	// 通知内容模板绑定的告警策略数量
	NoticeContentTmplBindPolicyCounts []*NoticeContentTmplBindPolicyCount `json:"NoticeContentTmplBindPolicyCounts,omitnil,omitempty" name:"NoticeContentTmplBindPolicyCounts"`

	// 分页数
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNoticeContentTmplResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNoticeContentTmplResponseParams `json:"Response"`
}

func (r *DescribeNoticeContentTmplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DingDingRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`

	// 标题模板
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type DingDingRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *DingDingRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type DispenseCondition struct {
	// 对外指标名
	ExtMetric *string `json:"ExtMetric,omitnil,omitempty" name:"ExtMetric"`

	// 过滤条件表
	DispenseFilters []*DispenseFilter `json:"DispenseFilters,omitnil,omitempty" name:"DispenseFilters"`

	// 过滤条件id
	ConditionId *int64 `json:"ConditionId,omitnil,omitempty" name:"ConditionId"`
}

type DispenseFilter struct {
	// 维度名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 维度值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 表示式
	Expression *string `json:"Expression,omitnil,omitempty" name:"Expression"`
}

type DispenseGlobalTag struct {
	// 维度key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 维度值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type DispenseRegion struct {
	// 地域缩写
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 地域中文名
	RegionCnName *string `json:"RegionCnName,omitnil,omitempty" name:"RegionCnName"`

	// 地域英文名
	RegionEnName *string `json:"RegionEnName,omitnil,omitempty" name:"RegionEnName"`

	// 规则数量
	RuleNumber *int64 `json:"RuleNumber,omitnil,omitempty" name:"RuleNumber"`
}

type EnvEntry struct {
	// <p>环境变量value</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>是否脱敏</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sensitive *bool `json:"Sensitive,omitnil,omitempty" name:"Sensitive"`
}

type EnvVar struct {
	// <p>环境变量key</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>环境变量value</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *EnvEntry `json:"Value,omitnil,omitempty" name:"Value"`
}

type ExecutionInfo struct {
	// <p>任务名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>任务 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>执行 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// <p>Agent ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>会话 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>触发类型: manual / cron / webhook</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>状态: pending/running/completed/failed/timeout/cancelled</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>执行摘要</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *string `json:"Summary,omitnil,omitempty" name:"Summary"`

	// <p>执行耗时(毫秒)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DurationMs *int64 `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`
}

type ExtMetric struct {
	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 中文指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricCName *string `json:"MetricCName,omitnil,omitempty" name:"MetricCName"`

	// 中文含义
	// 注意：此字段可能返回 null，表示取不到有效值。
	CNMeaning *string `json:"CNMeaning,omitnil,omitempty" name:"CNMeaning"`

	// 英文含义
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnMeaning *string `json:"EnMeaning,omitnil,omitempty" name:"EnMeaning"`

	// 单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 是否配置对外维度
	DimensionFlag *bool `json:"DimensionFlag,omitnil,omitempty" name:"DimensionFlag"`
}

type FeiShuRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`

	// 标题模板
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type FeiShuRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *FeiShuRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

// Predefined struct for user
type GetAIWorkbenchArtifactDownloadURLRequestParams struct {
	// <p>会话ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>制品ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`
}

type GetAIWorkbenchArtifactDownloadURLRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>制品ID</p>
	ArtifactId *string `json:"ArtifactId,omitnil,omitempty" name:"ArtifactId"`
}

func (r *GetAIWorkbenchArtifactDownloadURLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAIWorkbenchArtifactDownloadURLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "ArtifactId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAIWorkbenchArtifactDownloadURLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAIWorkbenchArtifactDownloadURLResponseParams struct {
	// <p>COS 预签名 HTTPS 下载 URL</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadURL *string `json:"DownloadURL,omitnil,omitempty" name:"DownloadURL"`

	// <p>URL 过期时间（RFC3339 格式）</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpiredAt *string `json:"ExpiredAt,omitnil,omitempty" name:"ExpiredAt"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAIWorkbenchArtifactDownloadURLResponse struct {
	*tchttp.BaseResponse
	Response *GetAIWorkbenchArtifactDownloadURLResponseParams `json:"Response"`
}

func (r *GetAIWorkbenchArtifactDownloadURLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAIWorkbenchArtifactDownloadURLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoogleChatRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`
}

type GoogleChatRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *GoogleChatRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type InstructionConfig struct {
	// <p>角色定义</p>
	RolePosition *string `json:"RolePosition,omitnil,omitempty" name:"RolePosition"`

	// <p>核心职责</p>
	CoreDuty *string `json:"CoreDuty,omitnil,omitempty" name:"CoreDuty"`

	// <p>核心原则</p>
	CoreTruths *string `json:"CoreTruths,omitnil,omitempty" name:"CoreTruths"`

	// <p>风格约束</p>
	Vibe *string `json:"Vibe,omitnil,omitempty" name:"Vibe"`

	// <p>注意事项</p>
	Boundaries *string `json:"Boundaries,omitnil,omitempty" name:"Boundaries"`
}

type KafkaConnectivity struct {
	// 地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 连通
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`
}

// Predefined struct for user
type ListAIWorkbenchAgentsRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>状态筛选</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>分类筛选</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>来源筛选</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Agent ID 列表筛选</p>
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`
}

type ListAIWorkbenchAgentsRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>状态筛选</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>分类筛选</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>来源筛选</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>Agent ID 列表筛选</p>
	AgentIds []*string `json:"AgentIds,omitnil,omitempty" name:"AgentIds"`
}

func (r *ListAIWorkbenchAgentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchAgentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Status")
	delete(f, "Category")
	delete(f, "Keyword")
	delete(f, "Source")
	delete(f, "AgentIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchAgentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchAgentsResponseParams struct {
	// <p>Agent 列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Agents []*AgentInfo `json:"Agents,omitnil,omitempty" name:"Agents"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchAgentsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchAgentsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchAgentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchAgentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchArtifactsRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>会话ID</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`

	// <p>消息内容类型</p>
	MimeTypes []*string `json:"MimeTypes,omitnil,omitempty" name:"MimeTypes"`

	// <p>排序</p><p>枚举值：</p><ul><li>ASC： 正序</li><li>DESC： 倒序</li></ul>
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

type ListAIWorkbenchArtifactsRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>会话ID</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`

	// <p>消息内容类型</p>
	MimeTypes []*string `json:"MimeTypes,omitnil,omitempty" name:"MimeTypes"`

	// <p>排序</p><p>枚举值：</p><ul><li>ASC： 正序</li><li>DESC： 倒序</li></ul>
	OrderDirection *string `json:"OrderDirection,omitnil,omitempty" name:"OrderDirection"`
}

func (r *ListAIWorkbenchArtifactsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchArtifactsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "SessionIds")
	delete(f, "MimeTypes")
	delete(f, "OrderDirection")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchArtifactsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchArtifactsResponseParams struct {
	// <p>产物列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Artifacts []*ArtifactInfo `json:"Artifacts,omitnil,omitempty" name:"Artifacts"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchArtifactsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchArtifactsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchArtifactsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchArtifactsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchExecutionsRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按 Agent 筛选</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>按状态筛选</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>执行 ID 列表筛选</p>
	ExecutionIds []*string `json:"ExecutionIds,omitnil,omitempty" name:"ExecutionIds"`

	// <p>任务id</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>触发方式</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>关键值</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ListAIWorkbenchExecutionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按 Agent 筛选</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>按状态筛选</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>执行 ID 列表筛选</p>
	ExecutionIds []*string `json:"ExecutionIds,omitnil,omitempty" name:"ExecutionIds"`

	// <p>任务id</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>触发方式</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>关键值</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ListAIWorkbenchExecutionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchExecutionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "AgentId")
	delete(f, "Status")
	delete(f, "ExecutionIds")
	delete(f, "TaskIds")
	delete(f, "TriggerType")
	delete(f, "Keyword")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchExecutionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchExecutionsResponseParams struct {
	// <p>执行列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Executions []*ExecutionInfo `json:"Executions,omitnil,omitempty" name:"Executions"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchExecutionsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchExecutionsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchExecutionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchExecutionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMCPsRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按传输协议筛选</p>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>是否启用筛选</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>关联的mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>MCP类型（内置/私有）</p><p>枚举值：</p><ul><li>builtin： 平台内置</li><li>private： 用户自定义</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ListAIWorkbenchMCPsRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按传输协议筛选</p>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>是否启用筛选</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>关联的mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>MCP类型（内置/私有）</p><p>枚举值：</p><ul><li>builtin： 平台内置</li><li>private： 用户自定义</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *ListAIWorkbenchMCPsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMCPsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Transport")
	delete(f, "Keyword")
	delete(f, "Enabled")
	delete(f, "MCPIds")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchMCPsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMCPsResponseParams struct {
	// <p>MCP 列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCPs []*MCPInfo `json:"MCPs,omitnil,omitempty" name:"MCPs"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchMCPsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchMCPsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchMCPsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMCPsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMessagesRequestParams struct {
	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>游标分页的定位标记</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>窗口大小</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>拉取顺序</p>
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type ListAIWorkbenchMessagesRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话 ID</p>
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>游标分页的定位标记</p>
	Cursor *string `json:"Cursor,omitnil,omitempty" name:"Cursor"`

	// <p>窗口大小</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>拉取顺序</p>
	Direction *string `json:"Direction,omitnil,omitempty" name:"Direction"`
}

func (r *ListAIWorkbenchMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SessionId")
	delete(f, "Cursor")
	delete(f, "Limit")
	delete(f, "Direction")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchMessagesResponseParams struct {
	// <p>消息列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Messages []*MessageInfo `json:"Messages,omitnil,omitempty" name:"Messages"`

	// <p>下一个游标</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// <p>还有后续吗</p>
	HasMore *bool `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchMessagesResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchMessagesResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceInstancesRequestParams struct {
	// <p>资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>分页参数</p>
	PageParams *PageByNumParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`
}

type ListAIWorkbenchResourceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// <p>资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>分页参数</p>
	PageParams *PageByNumParams `json:"PageParams,omitnil,omitempty" name:"PageParams"`
}

func (r *ListAIWorkbenchResourceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceMapId")
	delete(f, "PageParams")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchResourceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceInstancesResponseParams struct {
	// <p>资源实例列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*ResourceInstance `json:"Instances,omitnil,omitempty" name:"Instances"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchResourceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchResourceInstancesResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchResourceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceMapsRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按名称搜索</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListAIWorkbenchResourceMapsRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按名称搜索</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListAIWorkbenchResourceMapsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceMapsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchResourceMapsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchResourceMapsResponseParams struct {
	// <p>资源地图列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceMaps []*ResourceMapInfo `json:"ResourceMaps,omitnil,omitempty" name:"ResourceMaps"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchResourceMapsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchResourceMapsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchResourceMapsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchResourceMapsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSessionsRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按 Agent 筛选</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>会话 ID 列表筛选</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

type ListAIWorkbenchSessionsRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按 Agent 筛选</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>会话 ID 列表筛选</p>
	SessionIds []*string `json:"SessionIds,omitnil,omitempty" name:"SessionIds"`
}

func (r *ListAIWorkbenchSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSessionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "AgentId")
	delete(f, "Keyword")
	delete(f, "SessionIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchSessionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSessionsResponseParams struct {
	// <p>会话列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sessions []*SessionInfo `json:"Sessions,omitnil,omitempty" name:"Sessions"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchSessionsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchSessionsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSessionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSkillsRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按类型筛选</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>是否启用筛选</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>技能 ID 列表筛选</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`
}

type ListAIWorkbenchSkillsRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按类型筛选</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>是否启用筛选</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>技能 ID 列表筛选</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`
}

func (r *ListAIWorkbenchSkillsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSkillsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "Type")
	delete(f, "Keyword")
	delete(f, "Enabled")
	delete(f, "SkillIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchSkillsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchSkillsResponseParams struct {
	// <p>技能列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Skills []*SkillInfo `json:"Skills,omitnil,omitempty" name:"Skills"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchSkillsResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchSkillsResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchSkillsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchSkillsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchTasksRequestParams struct {
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按 Agent 筛选</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>按触发类型筛选</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>任务 ID 列表筛选</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>是否启用筛选</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type ListAIWorkbenchTasksRequest struct {
	*tchttp.BaseRequest
	
	// <p>每页数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>页码</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`

	// <p>按 Agent 筛选</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>按触发类型筛选</p>
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>任务 ID 列表筛选</p>
	TaskIds []*string `json:"TaskIds,omitnil,omitempty" name:"TaskIds"`

	// <p>是否启用筛选</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

func (r *ListAIWorkbenchTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchTasksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PerPage")
	delete(f, "PageNo")
	delete(f, "AgentId")
	delete(f, "TriggerType")
	delete(f, "Keyword")
	delete(f, "TaskIds")
	delete(f, "Enabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAIWorkbenchTasksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAIWorkbenchTasksResponseParams struct {
	// <p>任务列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tasks []*TaskInfo `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>分页结果</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageResult *PageByNumResult `json:"PageResult,omitnil,omitempty" name:"PageResult"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAIWorkbenchTasksResponse struct {
	*tchttp.BaseResponse
	Response *ListAIWorkbenchTasksResponseParams `json:"Response"`
}

func (r *ListAIWorkbenchTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAIWorkbenchTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MCPInfo struct {
	// <p>mcp的ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCPId *string `json:"MCPId,omitnil,omitempty" name:"MCPId"`

	// <p>MCP 名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>MCP 描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>MCP URL</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// <p>传输协议: sse / streamable_http / stdio</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>认证类型: none / bearer / basic / api_key</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// <p>认证密钥(响应时脱敏)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthSecret *string `json:"AuthSecret,omitnil,omitempty" name:"AuthSecret"`

	// <p>超时时间(秒)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>重试次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>请求头 JSON</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers *string `json:"Headers,omitnil,omitempty" name:"Headers"`

	// <p>是否启用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type MessageInfo struct {
	// <p>实体id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EntryId *string `json:"EntryId,omitnil,omitempty" name:"EntryId"`

	// <p>会话 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>角色: user / assistant</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// <p>消息内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>块内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContentBlocks []*ContentBlockInfo `json:"ContentBlocks,omitnil,omitempty" name:"ContentBlocks"`
}

// Predefined struct for user
type ModifyDispenseExternalRuleRequestParams struct {
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云监控对外命名空间
	ExtNamespace *string `json:"ExtNamespace,omitnil,omitempty" name:"ExtNamespace"`

	// 转发目标消信息
	Producer *Producer `json:"Producer,omitnil,omitempty" name:"Producer"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发部署地域列表
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`

	// 云监控对外指标
	ExtMetrics []*string `json:"ExtMetrics,omitnil,omitempty" name:"ExtMetrics"`

	// 指标统计周期
	Period []*int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 转发过滤信息
	DispenseConditions []*DispenseCondition `json:"DispenseConditions,omitnil,omitempty" name:"DispenseConditions"`
}

type ModifyDispenseExternalRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云监控对外命名空间
	ExtNamespace *string `json:"ExtNamespace,omitnil,omitempty" name:"ExtNamespace"`

	// 转发目标消信息
	Producer *Producer `json:"Producer,omitnil,omitempty" name:"Producer"`

	// 规则ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 转发部署地域列表
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`

	// 云监控对外指标
	ExtMetrics []*string `json:"ExtMetrics,omitnil,omitempty" name:"ExtMetrics"`

	// 指标统计周期
	Period []*int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 转发过滤信息
	DispenseConditions []*DispenseCondition `json:"DispenseConditions,omitnil,omitempty" name:"DispenseConditions"`
}

func (r *ModifyDispenseExternalRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDispenseExternalRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ExtNamespace")
	delete(f, "Producer")
	delete(f, "RuleId")
	delete(f, "DispenseRegions")
	delete(f, "ExtMetrics")
	delete(f, "Period")
	delete(f, "DispenseConditions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDispenseExternalRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDispenseExternalRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDispenseExternalRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDispenseExternalRuleResponseParams `json:"Response"`
}

func (r *ModifyDispenseExternalRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDispenseExternalRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDispenseExternalRuleStatusRequestParams struct {
	// 规则id列表
	RuleIdList []*int64 `json:"RuleIdList,omitnil,omitempty" name:"RuleIdList"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyDispenseExternalRuleStatusRequest struct {
	*tchttp.BaseRequest
	
	// 规则id列表
	RuleIdList []*int64 `json:"RuleIdList,omitnil,omitempty" name:"RuleIdList"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyDispenseExternalRuleStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDispenseExternalRuleStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RuleIdList")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDispenseExternalRuleStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDispenseExternalRuleStatusResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDispenseExternalRuleStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDispenseExternalRuleStatusResponseParams `json:"Response"`
}

func (r *ModifyDispenseExternalRuleStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDispenseExternalRuleStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNoticeContentTmplRequestParams struct {
	// 模板名称
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 模板内容
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// 需要修改的模板ID
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`
}

type ModifyNoticeContentTmplRequest struct {
	*tchttp.BaseRequest
	
	// 模板名称
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// 模板内容
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// 需要修改的模板ID
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`
}

func (r *ModifyNoticeContentTmplRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNoticeContentTmplRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TmplName")
	delete(f, "TmplContents")
	delete(f, "TmplID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNoticeContentTmplRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNoticeContentTmplResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNoticeContentTmplResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNoticeContentTmplResponseParams `json:"Response"`
}

func (r *ModifyNoticeContentTmplResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNoticeContentTmplResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeContentTmpl struct {
	// <p>自定义通知内容模板id，唯一id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmplID *string `json:"TmplID,omitnil,omitempty" name:"TmplID"`

	// <p>自定义通知内容模板名</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmplName *string `json:"TmplName,omitnil,omitempty" name:"TmplName"`

	// <p>通知内容</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmplContents *NoticeContentTmplItem `json:"TmplContents,omitnil,omitempty" name:"TmplContents"`

	// <p>Unix时间戳，秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *int64 `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>Unix时间戳，秒</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>最后修改人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastModifier *string `json:"LastModifier,omitnil,omitempty" name:"LastModifier"`

	// <p>创建人</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// <p>监控类型</p>
	MonitorType *string `json:"MonitorType,omitnil,omitempty" name:"MonitorType"`

	// <p>模板语言 en/zh</p>
	TmplLanguage *string `json:"TmplLanguage,omitnil,omitempty" name:"TmplLanguage"`
}

type NoticeContentTmplBindPolicyCount struct {
	// 通知内容模板ID
	NoticeContentTmplID *string `json:"NoticeContentTmplID,omitnil,omitempty" name:"NoticeContentTmplID"`

	// 绑定告警策略数量
	BindCount *uint64 `json:"BindCount,omitnil,omitempty" name:"BindCount"`
}

type NoticeContentTmplItem struct {
	// <p>官网通知渠道配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	QCloudYehe []*QCloudYeheNoticeTmplMatcher `json:"QCloudYehe,omitnil,omitempty" name:"QCloudYehe"`

	// <p>企业微信机器人通知渠道配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WeWorkRobot []*WeWorkRobotNoticeTmplMatcher `json:"WeWorkRobot,omitnil,omitempty" name:"WeWorkRobot"`

	// <p>钉钉机器人通知渠道配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DingDingRobot []*DingDingRobotNoticeTmplMatcher `json:"DingDingRobot,omitnil,omitempty" name:"DingDingRobot"`

	// <p>飞书机器人通知渠道配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeiShuRobot []*FeiShuRobotNoticeTmplMatcher `json:"FeiShuRobot,omitnil,omitempty" name:"FeiShuRobot"`

	// <p>自定义Webhook通知渠道配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Webhook []*WebhookNoticeTmplMatcher `json:"Webhook,omitnil,omitempty" name:"Webhook"`

	// <p>Teams机器人通知渠道配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TeamsRobot []*TeamsRobotNoticeTmplMatcher `json:"TeamsRobot,omitnil,omitempty" name:"TeamsRobot"`

	// <p>PagerDutyRobot机器人通知渠道配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PagerDutyRobot []*PagerDutyRobotNoticeTmplMatcher `json:"PagerDutyRobot,omitnil,omitempty" name:"PagerDutyRobot"`

	// <p>GoogleChat</p>
	GoogleChatRobot []*GoogleChatRobotNoticeTmplMatcher `json:"GoogleChatRobot,omitnil,omitempty" name:"GoogleChatRobot"`

	// <p>Slack</p>
	SlackRobot []*SlackRobotNoticeTmplMatcher `json:"SlackRobot,omitnil,omitempty" name:"SlackRobot"`

	// <p>Teams 工作流渠道</p>
	TeamsWorkflowRobot []*TeamsWorkflowRobotNoticeTmplMatcher `json:"TeamsWorkflowRobot,omitnil,omitempty" name:"TeamsWorkflowRobot"`
}

type NotifyRelatedNotice struct {
	// 通知模板ID
	NoticeId *string `json:"NoticeId,omitnil,omitempty" name:"NoticeId"`

	// 通知模板的名称
	NoticeName *string `json:"NoticeName,omitnil,omitempty" name:"NoticeName"`
}

type PageByNoParams struct {
	// 每个分页的数量是多少
	// 注意：此字段可能返回 null，表示取不到有效值。
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// 第几个分页，从1开始
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageNo *string `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

type PageByNoResult struct {
	// 总共有多少数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 总共有多少个分页
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// 当前的分页号
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPageNo *int64 `json:"CurrentPageNo,omitnil,omitempty" name:"CurrentPageNo"`

	// 【已弃用】是否遍历到末尾
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IsEnd is deprecated.
	IsEnd *bool `json:"IsEnd,omitnil,omitempty" name:"IsEnd"`

	// 是否遍历到末尾
	End *bool `json:"End,omitnil,omitempty" name:"End"`
}

type PageByNumParams struct {
	// <p>每个分页的数量</p>
	PerPage *int64 `json:"PerPage,omitnil,omitempty" name:"PerPage"`

	// <p>第几个分页，从1开始</p>
	PageNo *int64 `json:"PageNo,omitnil,omitempty" name:"PageNo"`
}

type PageByNumResult struct {
	// <p>总共有多少数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>总共有多少个分页</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalPage *int64 `json:"TotalPage,omitnil,omitempty" name:"TotalPage"`

	// <p>当前的分页号</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPageNo *int64 `json:"CurrentPageNo,omitnil,omitempty" name:"CurrentPageNo"`
}

type PagerDutyRobotNoticeTmpl struct {
	// 请求体模板 仅支持json
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 请求头 暂时未支持
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*PagerDutyRobotNoticeTmplHeader `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 标题模板
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type PagerDutyRobotNoticeTmplHeader struct {
	// http请求中header的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// http请求中header的value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type PagerDutyRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid; Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 自定义PagerDutyRobot内容模板
	Template *PagerDutyRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type Producer struct {
	// 转发协议类型，0-stormRetPb, 1-tcbDispensePb, 2-stormRetJson, 3-ADPPb(废弃)，4-中台pb
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtocolType *int64 `json:"ProtocolType,omitnil,omitempty" name:"ProtocolType"`

	// 目标类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 转发kafka地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Brokers *string `json:"Brokers,omitnil,omitempty" name:"Brokers"`

	// 转发kafka topic
	// 注意：此字段可能返回 null，表示取不到有效值。
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`

	// 是否合并指标,默认是1，合并
	Merge *int64 `json:"Merge,omitnil,omitempty" name:"Merge"`

	// 全局维度组
	GlobalTags []*DispenseGlobalTag `json:"GlobalTags,omitnil,omitempty" name:"GlobalTags"`

	// 默认维度组，只提供维度即可
	DefaultTags []*string `json:"DefaultTags,omitnil,omitempty" name:"DefaultTags"`

	// Kafka用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// Kafka密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type QCloudYeheNoticeTmpl struct {
	// 邮件通知渠道
	Email *QCloudYeheNoticeTmplItem `json:"Email,omitnil,omitempty" name:"Email"`

	// 企业微信通知渠道
	QYWX *QCloudYeheNoticeTmplItem `json:"QYWX,omitnil,omitempty" name:"QYWX"`

	// 短信通知渠道
	SMS *QCloudYeheNoticeTmplItem `json:"SMS,omitnil,omitempty" name:"SMS"`

	// 语音通知渠道
	Voice *QCloudYeheNoticeTmplItem `json:"Voice,omitnil,omitempty" name:"Voice"`

	// 微信通知渠道
	WeChat *QCloudYeheWeChatNoticeTmplItem `json:"WeChat,omitnil,omitempty" name:"WeChat"`

	// 站内信通知渠道
	Site *QCloudYeheNoticeTmplItem `json:"Site,omitnil,omitempty" name:"Site"`

	// 安灯通知渠道
	Andon *QCloudYeheNoticeTmplItem `json:"Andon,omitnil,omitempty" name:"Andon"`
}

type QCloudYeheNoticeTmplItem struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`

	// 标题
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type QCloudYeheNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *QCloudYeheNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type QCloudYeheWeChatNoticeTmplItem struct {
	// 告警内容模板
	AlarmContentTmpl *string `json:"AlarmContentTmpl,omitnil,omitempty" name:"AlarmContentTmpl"`

	// 告警对象模板
	AlarmObjectTmpl *string `json:"AlarmObjectTmpl,omitnil,omitempty" name:"AlarmObjectTmpl"`

	// 告警地域模板
	AlarmRegionTmpl *string `json:"AlarmRegionTmpl,omitnil,omitempty" name:"AlarmRegionTmpl"`

	// 告警时间模板
	AlarmTimeTmpl *string `json:"AlarmTimeTmpl,omitnil,omitempty" name:"AlarmTimeTmpl"`
}

type ResourceInstance struct {
	// <p>实例 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>服务名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// <p>地域</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// <p>是否就绪</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsReady *bool `json:"IsReady,omitnil,omitempty" name:"IsReady"`
}

type ResourceMapInfo struct {
	// <p>资源地图 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>资源地图名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>资源地图描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>总实例数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`
}

type Rule struct {
	// 规则Id
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 规则名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 对外namespace
	ExtNamespace *string `json:"ExtNamespace,omitnil,omitempty" name:"ExtNamespace"`

	// 对外指标列表
	ExtMetric []*ExtMetric `json:"ExtMetric,omitnil,omitempty" name:"ExtMetric"`

	// 输出信息
	Producer *Producer `json:"Producer,omitnil,omitempty" name:"Producer"`

	// 更新时间
	UpdateTime *int64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 规则触发状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 指标粒度周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period []*int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 转发过滤条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	DispenseConditions []*DispenseCondition `json:"DispenseConditions,omitnil,omitempty" name:"DispenseConditions"`

	// 转发地域列表
	DispenseRegions []*string `json:"DispenseRegions,omitnil,omitempty" name:"DispenseRegions"`
}

type SessionInfo struct {
	// <p>会话 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// <p>Agent ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>会话标题</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>状态: active / archived / deleted</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>如果该会话由任务触发，则携带触发其会话的任务ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type SkillInfo struct {
	// <p>技能 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>技能名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>技能描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否启用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type SlackRobotNoticeTmpl struct {
	// <p>内容模板</p>
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`
}

type SlackRobotNoticeTmplMatcher struct {
	// <p>匹配状态 Invalid;<br>Trigger 告警触发; Recovery 告警恢复</p>
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// <p>模板配置</p>
	Template *SlackRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type Tag struct {
	// 标签key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type TaskInfo struct {
	// <p>任务 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// <p>任务名称</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>任务描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>关联 Agent ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>提示词模板</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PromptTemplate *string `json:"PromptTemplate,omitnil,omitempty" name:"PromptTemplate"`

	// <p>输出格式: markdown / json</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputFormat *string `json:"OutputFormat,omitnil,omitempty" name:"OutputFormat"`

	// <p>触发类型: manual / cron / webhook</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TriggerType *string `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>Cron 表达式</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronExpr *string `json:"CronExpr,omitnil,omitempty" name:"CronExpr"`

	// <p>Cron 时区</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronTimezone *string `json:"CronTimezone,omitnil,omitempty" name:"CronTimezone"`

	// <p>关联技能 ID 列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>关联 MCP 端点 ID 列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	McpEndpointIds []*string `json:"McpEndpointIds,omitnil,omitempty" name:"McpEndpointIds"`

	// <p>超时时间(秒)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeoutSec *int64 `json:"TimeoutSec,omitnil,omitempty" name:"TimeoutSec"`

	// <p>重试次数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetryCount *int64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>通知id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyIds []*string `json:"NotifyIds,omitnil,omitempty" name:"NotifyIds"`

	// <p>是否启用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type TeamsRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`
}

type TeamsRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *TeamsRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type TeamsWorkflowRobotNoticeTmpl struct {
	// <p>内容模板</p>
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`

	// <p>区分 TeamsWorkflow 是自定义内容还是自定义 POST BODY</p><p>枚举值：</p><ul><li>WorkflowText： 自定义内容</li><li>WorkflowJson： 自定义 POST BODY</li></ul>
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// <p>标题模版</p>
	TitleTmpl *string `json:"TitleTmpl,omitnil,omitempty" name:"TitleTmpl"`
}

type TeamsWorkflowRobotNoticeTmplMatcher struct {
	// <p>匹配状态 Invalid; Trigger 告警触发; Recovery 告警恢复</p><p>枚举值：</p><ul><li>Trigger： 告警触发</li><li>Recovery： 告警恢复</li></ul>
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// <p>模板配置</p>
	Template *TeamsWorkflowRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

// Predefined struct for user
type TriggerAIWorkbenchSREDigitalTwinTaskRequestParams struct {
	// 数字分身任务ID
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

type TriggerAIWorkbenchSREDigitalTwinTaskRequest struct {
	*tchttp.BaseRequest
	
	// 数字分身任务ID
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

func (r *TriggerAIWorkbenchSREDigitalTwinTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchSREDigitalTwinTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerAIWorkbenchSREDigitalTwinTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerAIWorkbenchSREDigitalTwinTaskResponseParams struct {
	// Json序列化路径
	JSONStrPaths []*string `json:"JSONStrPaths,omitnil,omitempty" name:"JSONStrPaths"`

	// 数字分身任务信息
	Data *TriggerDigitalTwinTaskResp `json:"Data,omitnil,omitempty" name:"Data"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerAIWorkbenchSREDigitalTwinTaskResponse struct {
	*tchttp.BaseResponse
	Response *TriggerAIWorkbenchSREDigitalTwinTaskResponseParams `json:"Response"`
}

func (r *TriggerAIWorkbenchSREDigitalTwinTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchSREDigitalTwinTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerAIWorkbenchTaskRequestParams struct {
	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

type TriggerAIWorkbenchTaskRequest struct {
	*tchttp.BaseRequest
	
	// <p>任务 ID</p>
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

func (r *TriggerAIWorkbenchTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TriggerAIWorkbenchTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TriggerAIWorkbenchTaskResponseParams struct {
	// <p>执行 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExecutionId *string `json:"ExecutionId,omitnil,omitempty" name:"ExecutionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TriggerAIWorkbenchTaskResponse struct {
	*tchttp.BaseResponse
	Response *TriggerAIWorkbenchTaskResponseParams `json:"Response"`
}

func (r *TriggerAIWorkbenchTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TriggerAIWorkbenchTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TriggerDigitalTwinTaskResp struct {
	// 数字分身任务ID
	TaskID *int64 `json:"TaskID,omitnil,omitempty" name:"TaskID"`
}

// Predefined struct for user
type UpdateAIWorkbenchAgentRequestParams struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent 名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent 描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent 分类</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent 标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent 提示词</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>关联技能 ID 列表</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>关联的资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>关联的mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>agent运行时环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

type UpdateAIWorkbenchAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent 名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>Agent 描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>Agent 分类</p>
	Category *string `json:"Category,omitnil,omitempty" name:"Category"`

	// <p>Agent 标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>Agent 提示词</p>
	Instruction *InstructionConfig `json:"Instruction,omitnil,omitempty" name:"Instruction"`

	// <p>关联技能 ID 列表</p>
	SkillIds []*string `json:"SkillIds,omitnil,omitempty" name:"SkillIds"`

	// <p>来源</p>
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>状态</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>关联的资源地图 ID</p>
	ResourceMapId *string `json:"ResourceMapId,omitnil,omitempty" name:"ResourceMapId"`

	// <p>关联的mcp</p>
	MCPIds []*string `json:"MCPIds,omitnil,omitempty" name:"MCPIds"`

	// <p>agent运行时环境变量</p>
	EnvVars []*EnvVar `json:"EnvVars,omitnil,omitempty" name:"EnvVars"`
}

func (r *UpdateAIWorkbenchAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIWorkbenchAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "Category")
	delete(f, "Tags")
	delete(f, "Instruction")
	delete(f, "SkillIds")
	delete(f, "Source")
	delete(f, "Status")
	delete(f, "ResourceMapId")
	delete(f, "MCPIds")
	delete(f, "EnvVars")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateAIWorkbenchAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateAIWorkbenchAgentResponseParams struct {
	// <p>更新后的 Agent 信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Agent *AgentInfo `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateAIWorkbenchAgentResponse struct {
	*tchttp.BaseResponse
	Response *UpdateAIWorkbenchAgentResponseParams `json:"Response"`
}

func (r *UpdateAIWorkbenchAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateAIWorkbenchAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WeWorkRobotNoticeTmpl struct {
	// 内容模板
	ContentTmpl *string `json:"ContentTmpl,omitnil,omitempty" name:"ContentTmpl"`
}

type WeWorkRobotNoticeTmplMatcher struct {
	// 匹配状态 Invalid;
	// Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 模板配置
	Template *WeWorkRobotNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}

type WebhookNoticeTmpl struct {
	// 请求体
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// 请求体的类型，非必填、默认为JSON
	// 注意：此字段可能返回 null，表示取不到有效值。
	BodyContentType *string `json:"BodyContentType,omitnil,omitempty" name:"BodyContentType"`

	// 请求头
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*WebhookNoticeTmplHeader `json:"Headers,omitnil,omitempty" name:"Headers"`
}

type WebhookNoticeTmplHeader struct {
	// http请求中header的key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// http请求中header的value
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type WebhookNoticeTmplMatcher struct {
	// 匹配状态 Invalid; Trigger 告警触发; Recovery 告警恢复
	MatchingStatus []*string `json:"MatchingStatus,omitnil,omitempty" name:"MatchingStatus"`

	// 自定义Webhook内容模板
	Template *WebhookNoticeTmpl `json:"Template,omitnil,omitempty" name:"Template"`
}