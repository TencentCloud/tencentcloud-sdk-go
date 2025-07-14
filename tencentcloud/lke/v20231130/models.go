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

package v20231130

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AICallConfig struct {
	// 启用语音互动功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableVoiceInteract *bool `json:"EnableVoiceInteract,omitnil,omitempty" name:"EnableVoiceInteract"`

	// 启用语音通话
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableVoiceCall *bool `json:"EnableVoiceCall,omitnil,omitempty" name:"EnableVoiceCall"`

	// 启用数智人
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDigitalHuman *bool `json:"EnableDigitalHuman,omitnil,omitempty" name:"EnableDigitalHuman"`

	// 音色配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Voice *VoiceConfig `json:"Voice,omitnil,omitempty" name:"Voice"`

	// 数智人配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DigitalHuman *DigitalHumanConfig `json:"DigitalHuman,omitnil,omitempty" name:"DigitalHuman"`
}

type Agent struct {
	// AgentID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// WorkflowID，非空则当前Agent从workflow转换而来
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// Agent名称，同一个应用内，Agent名称不能重复
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 插件图标url
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// Agent指令；当该Agent被调用时，将作为“系统提示词”使用，描述Agent应执行的操作和响应方式
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`

	// 当Agent作为转交目标时的描述，用于让其他Agent的LLM理解其功能和转交时机
	HandoffDescription *string `json:"HandoffDescription,omitnil,omitempty" name:"HandoffDescription"`

	// Agent可转交的子AgentId列表
	Handoffs []*string `json:"Handoffs,omitnil,omitempty" name:"Handoffs"`

	// Agent调用LLM时使用的模型配置
	Model *AgentModelInfo `json:"Model,omitnil,omitempty" name:"Model"`

	// Agent可使用的工具列表
	Tools []*AgentToolInfo `json:"Tools,omitnil,omitempty" name:"Tools"`

	// Agent可使用的插件列表
	Plugins []*AgentPluginInfo `json:"Plugins,omitnil,omitempty" name:"Plugins"`

	// 当前Agent是否是启动Agent
	IsStartingAgent *bool `json:"IsStartingAgent,omitnil,omitempty" name:"IsStartingAgent"`

	// Agent类型; 0: 未指定类型; 1: 知识库检索Agent
	AgentType *uint64 `json:"AgentType,omitnil,omitempty" name:"AgentType"`
}

type AgentDebugInfo struct {
	// 工具、大模型的输入信息，json
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 工具、大模型的输出信息，json
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`
}

type AgentInput struct {
	// 输入来源类型：0 用户输入，3 自定义变量（API参数）
	InputType *uint64 `json:"InputType,omitnil,omitempty" name:"InputType"`

	// 用户手写输入
	UserInputValue *AgentInputUserInputValue `json:"UserInputValue,omitnil,omitempty" name:"UserInputValue"`

	// 自定义变量（API参数）
	CustomVarId *string `json:"CustomVarId,omitnil,omitempty" name:"CustomVarId"`
}

type AgentInputUserInputValue struct {
	// 用户输入的值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AgentKnowledge struct {
	// 知识库id
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 0-应用内知识库
	// 1-共享知识库
	KnowledgeType *int64 `json:"KnowledgeType,omitnil,omitempty" name:"KnowledgeType"`

	// 0-全部知识
	// 1-按文档和问答
	// 2-按标签
	Filter *int64 `json:"Filter,omitnil,omitempty" name:"Filter"`

	// 文档id
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// true:包含所有问答
	// false:不包含问答
	AllQa *bool `json:"AllQa,omitnil,omitempty" name:"AllQa"`

	// 文档标签过滤器
	Tag *AgentKnowledgeFilterTag `json:"Tag,omitnil,omitempty" name:"Tag"`
}

type AgentKnowledgeAttrLabel struct {
	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 标签值，标签值之间是或的关系，只有匹配的，才会进行知识检索，否则报检索不到
	Inputs []*AgentInput `json:"Inputs,omitnil,omitempty" name:"Inputs"`
}

type AgentKnowledgeFilter struct {
	// 知识检索筛选方式; 0: 全部知识; 1:按文档和问答; 2: 按标签
	FilterType *uint64 `json:"FilterType,omitnil,omitempty" name:"FilterType"`

	// 文档和问答过滤器
	DocAndAnswer *AgentKnowledgeFilterDocAndAnswer `json:"DocAndAnswer,omitnil,omitempty" name:"DocAndAnswer"`

	// 标签过滤器
	Tag *AgentKnowledgeFilterTag `json:"Tag,omitnil,omitempty" name:"Tag"`

	// 知识库列表
	KnowledgeList []*AgentKnowledge `json:"KnowledgeList,omitnil,omitempty" name:"KnowledgeList"`

	// 是否检索全部知识
	AllKnowledge *bool `json:"AllKnowledge,omitnil,omitempty" name:"AllKnowledge"`
}

type AgentKnowledgeFilterDocAndAnswer struct {
	// 文档ID列表
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 问答
	AllQa *bool `json:"AllQa,omitnil,omitempty" name:"AllQa"`
}

type AgentKnowledgeFilterTag struct {
	// 标签之间的关系;0:AND, 1:OR
	Operator *uint64 `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 标签
	Labels []*AgentKnowledgeAttrLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type AgentKnowledgeQAPlugin struct {
	// 知识检索筛选范围
	Filter *AgentKnowledgeFilter `json:"Filter,omitnil,omitempty" name:"Filter"`
}

type AgentMCPServerInfo struct {
	// mcp server URL地址
	McpServerUrl *string `json:"McpServerUrl,omitnil,omitempty" name:"McpServerUrl"`

	// mcp server header信息
	Headers []*AgentPluginHeader `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 超时时间，单位秒
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// sse服务超时时间，单位秒
	SseReadTimeout *int64 `json:"SseReadTimeout,omitnil,omitempty" name:"SseReadTimeout"`
}

type AgentModelInfo struct {
	// 模型名称
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型别名
	ModelAliasName *string `json:"ModelAliasName,omitnil,omitempty" name:"ModelAliasName"`

	// 模型温度
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 模型TopP
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 模型是否可用
	IsEnabled *bool `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 对话历史条数限制
	HistoryLimit *uint64 `json:"HistoryLimit,omitnil,omitempty" name:"HistoryLimit"`

	// 模型上下文长度字符限制
	ModelContextWordsLimit *string `json:"ModelContextWordsLimit,omitnil,omitempty" name:"ModelContextWordsLimit"`

	// 指令长度字符限制
	InstructionsWordsLimit *uint64 `json:"InstructionsWordsLimit,omitnil,omitempty" name:"InstructionsWordsLimit"`

	// 单次会话最大推理轮数
	MaxReasoningRound *uint64 `json:"MaxReasoningRound,omitnil,omitempty" name:"MaxReasoningRound"`
}

type AgentPluginHeader struct {
	// 参数名称
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// 参数值
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`

	// header参数配置是否隐藏不可见
	GlobalHidden *bool `json:"GlobalHidden,omitnil,omitempty" name:"GlobalHidden"`

	// 输入的值
	Input *AgentInput `json:"Input,omitnil,omitempty" name:"Input"`

	// 参数是否可以为空
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`
}

type AgentPluginInfo struct {
	// 插件id
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// 应用配置的插件header信息
	Headers []*AgentPluginHeader `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 插件调用LLM时使用的模型配置，一般用于指定知识库问答插件的生成模型
	Model *AgentModelInfo `json:"Model,omitnil,omitempty" name:"Model"`

	// 插件信息类型; 0: 未指定类型; 1: 知识库问答插件
	PluginInfoType *uint64 `json:"PluginInfoType,omitnil,omitempty" name:"PluginInfoType"`

	// 知识库问答插件配置
	KnowledgeQa *AgentKnowledgeQAPlugin `json:"KnowledgeQa,omitnil,omitempty" name:"KnowledgeQa"`
}

type AgentProcedure struct {
	// 索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 执行过程英语名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 中文名, 用于展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 状态常量: 使用中: processing, 成功: success, 失败: failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// Agent调试信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Debugging *AgentProcedureDebugging `json:"Debugging,omitnil,omitempty" name:"Debugging"`

	// 是否切换Agent，取值为"main"或者"workflow",不切换为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	Switch *string `json:"Switch,omitnil,omitempty" name:"Switch"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 当前请求执行时间, 单位 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elapsed *uint64 `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// 工作流节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 用于展示思考放在哪个回复气泡中
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyIndex *uint64 `json:"ReplyIndex,omitnil,omitempty" name:"ReplyIndex"`

	// 主agent
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceAgentName *string `json:"SourceAgentName,omitnil,omitempty" name:"SourceAgentName"`

	// 挂号agent
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetAgentName *string `json:"TargetAgentName,omitnil,omitempty" name:"TargetAgentName"`

	// Agent的图标
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentIcon *string `json:"AgentIcon,omitnil,omitempty" name:"AgentIcon"`
}

type AgentProcedureDebugging struct {
	// 模型思考内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 展示的具体文本内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayContent *string `json:"DisplayContent,omitnil,omitempty" name:"DisplayContent"`

	// 展示类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayType *uint64 `json:"DisplayType,omitnil,omitempty" name:"DisplayType"`

	// 搜索引擎展示的索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuoteInfos []*QuoteInfo `json:"QuoteInfos,omitnil,omitempty" name:"QuoteInfos"`

	// 具体的参考来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	References []*AgentReference `json:"References,omitnil,omitempty" name:"References"`

	// 展示正在执行的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayStatus *string `json:"DisplayStatus,omitnil,omitempty" name:"DisplayStatus"`

	// 云桌面的URL地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	SandboxUrl *string `json:"SandboxUrl,omitnil,omitempty" name:"SandboxUrl"`

	// 云桌面里面通过浏览器打开的URL地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisplayUrl *string `json:"DisplayUrl,omitnil,omitempty" name:"DisplayUrl"`
}

type AgentReference struct {
	// 来源文档ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 文档业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文档名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocName *string `json:"DocName,omitnil,omitempty" name:"DocName"`

	// 问答业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 搜索引擎索引
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *uint64 `json:"Index,omitnil,omitempty" name:"Index"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type AgentThought struct {
	// 会话 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 请求 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 对应哪条会话, 会话 ID, 用于回答的消息存储使用, 可提前生成, 保存消息时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 当前请求执行时间, 单位 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elapsed *uint64 `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// 当前是否为工作流
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWorkflow *bool `json:"IsWorkflow,omitnil,omitempty" name:"IsWorkflow"`

	// 如果当前是工作流，工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 具体思考过程详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedures []*AgentProcedure `json:"Procedures,omitnil,omitempty" name:"Procedures"`

	// TraceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// 文件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Files []*FileInfo `json:"Files,omitnil,omitempty" name:"Files"`
}

type AgentToolInfo struct {
	// 插件id
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// 插件名称
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// 插件图标url
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// 0 自定义插件
	// 1 官方插件
	// 2 第三方插件 目前用于第三方实现的mcp server
	PluginType *uint64 `json:"PluginType,omitnil,omitempty" name:"PluginType"`

	// 工具id
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 工具名称
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 工具描述
	ToolDesc *string `json:"ToolDesc,omitnil,omitempty" name:"ToolDesc"`

	// 输入参数
	Inputs []*AgentToolReqParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// 输出参数
	Outputs []*AgentToolRspParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// 创建方式，0:服务创建，1:代码创建，2:MCP创建	
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// MCP插件的配置信息
	McpServer *AgentMCPServerInfo `json:"McpServer,omitnil,omitempty" name:"McpServer"`

	// 该工具是否和知识库绑定
	IsBindingKnowledge *bool `json:"IsBindingKnowledge,omitnil,omitempty" name:"IsBindingKnowledge"`

	// 插件状态，1:可用，2:不可用	
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// header信息
	Headers []*AgentPluginHeader `json:"Headers,omitnil,omitempty" name:"Headers"`

	// NON_STREAMING: 非流式  STREAMIN: 流式
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallingMethod *string `json:"CallingMethod,omitnil,omitempty" name:"CallingMethod"`
}

type AgentToolReqParam struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 参数类型，0:string, 1:int, 2:float，3:bool 4:object 5:array_string, 6:array_int, 7:array_float, 8:array_bool, 9:array_object
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数是否必填
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 子参数,ParamType 是OBJECT 或 ARRAY<>类型有用
	SubParams []*AgentToolReqParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// 是否隐藏不可见
	GlobalHidden *bool `json:"GlobalHidden,omitnil,omitempty" name:"GlobalHidden"`

	// agent模式下模型是否可见
	AgentHidden *bool `json:"AgentHidden,omitnil,omitempty" name:"AgentHidden"`

	// 其中任意
	AnyOf []*AgentToolReqParam `json:"AnyOf,omitnil,omitempty" name:"AnyOf"`

	// 其中一个
	OneOf []*AgentToolReqParam `json:"OneOf,omitnil,omitempty" name:"OneOf"`

	// 输入
	Input *AgentInput `json:"Input,omitnil,omitempty" name:"Input"`
}

type AgentToolRspParam struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 参数类型，0:string, 1:int, 2:float，3:bool 4:object 5:array_string, 6:array_int, 7:array_float, 8:array_bool, 9:array_object
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 子参数,ParamType 是OBJECT 或 ARRAY<>类型有用
	SubParams []*AgentToolRspParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// agent模式下模型是否可见
	AgentHidden *bool `json:"AgentHidden,omitnil,omitempty" name:"AgentHidden"`

	// 是否隐藏不可见
	GlobalHidden *bool `json:"GlobalHidden,omitnil,omitempty" name:"GlobalHidden"`

	// COVER: 覆盖解析 INCREMENT:增量解析
	AnalysisMethod *string `json:"AnalysisMethod,omitnil,omitempty" name:"AnalysisMethod"`
}

type ApiVarAttrInfo struct {
	// 自定义变量id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiVarId *string `json:"ApiVarId,omitnil,omitempty" name:"ApiVarId"`

	// 标签id
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`
}

type AppBaseInfo struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`
}

type AppConfig struct {
	// 知识问答管理应用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	KnowledgeQa *KnowledgeQaConfig `json:"KnowledgeQa,omitnil,omitempty" name:"KnowledgeQa"`

	// 知识摘要应用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Summary *SummaryConfig `json:"Summary,omitnil,omitempty" name:"Summary"`

	// 标签提取应用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Classify *ClassifyConfig `json:"Classify,omitnil,omitempty" name:"Classify"`
}

type AppInfo struct {
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用类型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppTypeDesc *string `json:"AppTypeDesc,omitnil,omitempty" name:"AppTypeDesc"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用头像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 应用描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 应用状态，1：未上线，2：运行中，3：停用
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppStatus *uint64 `json:"AppStatus,omitnil,omitempty" name:"AppStatus"`

	// 状态说明
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppStatusDesc *string `json:"AppStatusDesc,omitnil,omitempty" name:"AppStatusDesc"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最后修改人
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 生成模型别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelAliasName *string `json:"ModelAliasName,omitnil,omitempty" name:"ModelAliasName"`

	// 应用模式 standard:标准模式, agent: agent模式，single_workflow：单工作流模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 思考模型别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThoughtModelAliasName *string `json:"ThoughtModelAliasName,omitnil,omitempty" name:"ThoughtModelAliasName"`
}

type AppModel struct {
	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 上下文指代轮次
	// 注意：此字段可能返回 null，表示取不到有效值。
	ContextLimit *uint64 `json:"ContextLimit,omitnil,omitempty" name:"ContextLimit"`

	// 模型别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// token余量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenBalance *float64 `json:"TokenBalance,omitnil,omitempty" name:"TokenBalance"`

	// 是否使用上下文指代轮次
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUseContext *bool `json:"IsUseContext,omitnil,omitempty" name:"IsUseContext"`

	// 上下文记忆轮数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HistoryLimit *uint64 `json:"HistoryLimit,omitnil,omitempty" name:"HistoryLimit"`

	// 使用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsageType *string `json:"UsageType,omitnil,omitempty" name:"UsageType"`

	// 模型温度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Temperature *string `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 模型TopP
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopP *string `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 模型资源状态 1：资源可用；2：资源已用尽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type AttrLabel struct {
	// 标签来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`

	// 标签标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*Label `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type AttrLabelDetail struct {
	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`

	// 标签标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签值名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelNames []*string `json:"LabelNames,omitnil,omitempty" name:"LabelNames"`

	// 标签是否在更新中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUpdating *bool `json:"IsUpdating,omitnil,omitempty" name:"IsUpdating"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 标签值总数
	LabelTotalCount *string `json:"LabelTotalCount,omitnil,omitempty" name:"LabelTotalCount"`
}

type AttrLabelRefer struct {
	// 标签来源，1：标签
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 标签ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 标签值ID
	LabelBizIds []*string `json:"LabelBizIds,omitnil,omitempty" name:"LabelBizIds"`
}

type AttributeFilters struct {
	// 检索，属性或标签名称
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type AttributeLabel struct {
	// 标准词ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 标准词名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 同义词名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimilarLabels []*string `json:"SimilarLabels,omitnil,omitempty" name:"SimilarLabels"`
}

type AttributeLabelRefByWorkflow struct {
	// 标签值id
	AttributeLabelBizId *string `json:"AttributeLabelBizId,omitnil,omitempty" name:"AttributeLabelBizId"`

	// 标签值引用的工作流列表
	WorkflowList []*WorkflowRef `json:"WorkflowList,omitnil,omitempty" name:"WorkflowList"`
}

type BackgroundImageConfig struct {
	// 横图(pc)
	LandscapeImageUrl *string `json:"LandscapeImageUrl,omitnil,omitempty" name:"LandscapeImageUrl"`

	// 原始图
	OriginalImageUrl *string `json:"OriginalImageUrl,omitnil,omitempty" name:"OriginalImageUrl"`

	// 长图(手机)
	PortraitImageUrl *string `json:"PortraitImageUrl,omitnil,omitempty" name:"PortraitImageUrl"`

	// 主题色
	ThemeColor *string `json:"ThemeColor,omitnil,omitempty" name:"ThemeColor"`

	// 亮度值
	Brightness *int64 `json:"Brightness,omitnil,omitempty" name:"Brightness"`
}

type BaseConfig struct {
	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 应用头像url，在CreateApp和ModifyApp中作为入参必填。
	// 作为入参传入说明：
	// 1. 传入的url图片限制为jpeg和png，大小限制为500KB，url链接需允许head请求。
	// 2. 如果用户没有对象存储，可使用“获取文件上传临时密钥”(DescribeStorageCredential)接口，获取cos临时密钥和上传路径，自行上传头像至cos中并获取访问链接。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 应用描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type CallDetail struct {
	// 关联ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 调用时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallTime *string `json:"CallTime,omitnil,omitempty" name:"CallTime"`

	// 总token消耗
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTokenUsage *float64 `json:"TotalTokenUsage,omitnil,omitempty" name:"TotalTokenUsage"`

	// 输入token消耗
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputTokenUsage *float64 `json:"InputTokenUsage,omitnil,omitempty" name:"InputTokenUsage"`

	// 输出token消耗
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputTokenUsage *float64 `json:"OutputTokenUsage,omitnil,omitempty" name:"OutputTokenUsage"`

	// 搜索服务调用次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchUsage *uint64 `json:"SearchUsage,omitnil,omitempty" name:"SearchUsage"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 调用类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// 账号
	// 注意：此字段可能返回 null，表示取不到有效值。
	UinAccount *string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 总消耗页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageUsage *uint64 `json:"PageUsage,omitnil,omitempty" name:"PageUsage"`

	// 筛选子场景
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubScene *string `json:"SubScene,omitnil,omitempty" name:"SubScene"`

	// 账单明细对应的自定义tag
	BillingTag *string `json:"BillingTag,omitnil,omitempty" name:"BillingTag"`
}

type CateInfo struct {
	// 分类ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 分类名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类下的Record（如文档、同义词等）数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 是否可新增
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// 是否可编辑
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// 是否可删除
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// 子分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*CateInfo `json:"Children,omitnil,omitempty" name:"Children"`
}

// Predefined struct for user
type CheckAttributeLabelExistRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性名称
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 滚动加载，最后一个属性标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`
}

type CheckAttributeLabelExistRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性名称
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 滚动加载，最后一个属性标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`
}

func (r *CheckAttributeLabelExistRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelExistRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "LabelName")
	delete(f, "AttributeBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "LastLabelBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAttributeLabelExistRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAttributeLabelExistResponseParams struct {
	// 是否存在
	IsExist *bool `json:"IsExist,omitnil,omitempty" name:"IsExist"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAttributeLabelExistResponse struct {
	*tchttp.BaseResponse
	Response *CheckAttributeLabelExistResponseParams `json:"Response"`
}

func (r *CheckAttributeLabelExistResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelExistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAttributeLabelReferRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性标签
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 属性ID
	AttributeBizId []*string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`
}

type CheckAttributeLabelReferRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性标签
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 属性ID
	AttributeBizId []*string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`
}

func (r *CheckAttributeLabelReferRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelReferRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "LabelBizId")
	delete(f, "AttributeBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CheckAttributeLabelReferRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CheckAttributeLabelReferResponseParams struct {
	// 是否引用
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 引用的工作流详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*AttributeLabelRefByWorkflow `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CheckAttributeLabelReferResponse struct {
	*tchttp.BaseResponse
	Response *CheckAttributeLabelReferResponseParams `json:"Response"`
}

func (r *CheckAttributeLabelReferResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CheckAttributeLabelReferResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClassifyConfig struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// 标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Labels []*ClassifyLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 欢迎语，200字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`
}

type ClassifyLabel struct {
	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签取值范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type Context struct {
	// 消息记录ID信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordBizId *string `json:"RecordBizId,omitnil,omitempty" name:"RecordBizId"`

	// 是否为用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsVisitor *bool `json:"IsVisitor,omitnil,omitempty" name:"IsVisitor"`

	// 昵称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// 头像
	// 注意：此字段可能返回 null，表示取不到有效值。
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 消息内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 文档信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileInfos []*MsgFileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 回复方式，15：澄清确认回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyMethod *uint64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`
}

// Predefined struct for user
type CreateAgentRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 要增加的Agent的信息
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type CreateAgentRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 要增加的Agent的信息
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *CreateAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentResponseParams struct {
	// 新建的AgentID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAgentResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgentResponseParams `json:"Response"`
}

func (r *CreateAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// 应用类型；knowledge_qa-知识问答管理
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 应用模式 standard:标准模式, agent: agent模式，single_workflow：单工作流模式
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用类型；knowledge_qa-知识问答管理
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 应用模式 standard:标准模式, agent: agent模式，single_workflow：单工作流模式
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "BaseConfig")
	delete(f, "Pattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 判断账户应用列表权限是否是自定义的，用户交互提示
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCustomList *bool `json:"IsCustomList,omitnil,omitempty" name:"IsCustomList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppResponseParams `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAttributeLabelRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 标签名
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签值
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 标签标识（不生效，无需填写） 已作废
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type CreateAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 标签名
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签值
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 标签标识（不生效，无需填写） 已作废
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *CreateAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttrName")
	delete(f, "Labels")
	delete(f, "AttrKey")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAttributeLabelResponseParams struct {
	// 标签ID
	AttrBizId *string `json:"AttrBizId,omitnil,omitempty" name:"AttrBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *CreateAttributeLabelResponseParams `json:"Response"`
}

func (r *CreateAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocCateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 父级业务ID
	ParentBizId *string `json:"ParentBizId,omitnil,omitempty" name:"ParentBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateDocCateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 父级业务ID
	ParentBizId *string `json:"ParentBizId,omitnil,omitempty" name:"ParentBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreateDocCateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocCateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ParentBizId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateDocCateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateDocCateResponseParams struct {
	// 是否可新增
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// 是否可编辑
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// 是否可删除
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateDocCateResponse struct {
	*tchttp.BaseResponse
	Response *CreateDocCateResponseParams `json:"Response"`
}

func (r *CreateDocCateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateDocCateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQACateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 父级业务ID，创建顶级分类时传字符串"0"
	ParentBizId *string `json:"ParentBizId,omitnil,omitempty" name:"ParentBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 父级业务ID，创建顶级分类时传字符串"0"
	ParentBizId *string `json:"ParentBizId,omitnil,omitempty" name:"ParentBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *CreateQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ParentBizId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQACateResponseParams struct {
	// 是否可新增
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// 是否可编辑
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// 是否可删除
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQACateResponse struct {
	*tchttp.BaseResponse
	Response *CreateQACateResponseParams `json:"Response"`
}

func (r *CreateQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQARequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 相似问内容
	SimilarQuestions []*string `json:"SimilarQuestions,omitnil,omitempty" name:"SimilarQuestions"`

	// 问题描述
	QuestionDesc *string `json:"QuestionDesc,omitnil,omitempty" name:"QuestionDesc"`
}

type CreateQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 相似问内容
	SimilarQuestions []*string `json:"SimilarQuestions,omitnil,omitempty" name:"SimilarQuestions"`

	// 问题描述
	QuestionDesc *string `json:"QuestionDesc,omitnil,omitempty" name:"QuestionDesc"`
}

func (r *CreateQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Question")
	delete(f, "Answer")
	delete(f, "AttrRange")
	delete(f, "CustomParam")
	delete(f, "AttrLabels")
	delete(f, "DocBizId")
	delete(f, "CateBizId")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "SimilarQuestions")
	delete(f, "QuestionDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateQAResponseParams struct {
	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateQAResponse struct {
	*tchttp.BaseResponse
	Response *CreateQAResponseParams `json:"Response"`
}

func (r *CreateQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRejectedQuestionRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id， - 拒答来源于不满意回复  2 - 拒答来源于手动添加
	BusinessSource *uint64 `json:"BusinessSource,omitnil,omitempty" name:"BusinessSource"`

	// 拒答问题来源的数据源唯一id
	// 
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

type CreateRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id， - 拒答来源于不满意回复  2 - 拒答来源于手动添加
	BusinessSource *uint64 `json:"BusinessSource,omitnil,omitempty" name:"BusinessSource"`

	// 拒答问题来源的数据源唯一id
	// 
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`
}

func (r *CreateRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Question")
	delete(f, "BusinessSource")
	delete(f, "BusinessId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRejectedQuestionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *CreateRejectedQuestionResponseParams `json:"Response"`
}

func (r *CreateRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 渠道业务ID
	ChannelBizIds []*string `json:"ChannelBizIds,omitnil,omitempty" name:"ChannelBizIds"`
}

type CreateReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 渠道业务ID
	ChannelBizIds []*string `json:"ChannelBizIds,omitnil,omitempty" name:"ChannelBizIds"`
}

func (r *CreateReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Desc")
	delete(f, "ChannelBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseResponseParams struct {
	// 发布ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateReleaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateReleaseResponseParams `json:"Response"`
}

func (r *CreateReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSharedKnowledgeRequestParams struct {
	// 共享知识库名称，字符数量范围：[1, 50]
	KnowledgeName *string `json:"KnowledgeName,omitnil,omitempty" name:"KnowledgeName"`

	// 共享知识库描述，字符数量上限2000
	KnowledgeDescription *string `json:"KnowledgeDescription,omitnil,omitempty" name:"KnowledgeDescription"`

	// Embedding模型，字符数量上限128
	EmbeddingModel *string `json:"EmbeddingModel,omitnil,omitempty" name:"EmbeddingModel"`
}

type CreateSharedKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 共享知识库名称，字符数量范围：[1, 50]
	KnowledgeName *string `json:"KnowledgeName,omitnil,omitempty" name:"KnowledgeName"`

	// 共享知识库描述，字符数量上限2000
	KnowledgeDescription *string `json:"KnowledgeDescription,omitnil,omitempty" name:"KnowledgeDescription"`

	// Embedding模型，字符数量上限128
	EmbeddingModel *string `json:"EmbeddingModel,omitnil,omitempty" name:"EmbeddingModel"`
}

func (r *CreateSharedKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeName")
	delete(f, "KnowledgeDescription")
	delete(f, "EmbeddingModel")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSharedKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSharedKnowledgeResponseParams struct {
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSharedKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *CreateSharedKnowledgeResponseParams `json:"Response"`
}

func (r *CreateSharedKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSharedKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVarRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量名称，不允许重复，最大支持50个字符
	VarName *string `json:"VarName,omitnil,omitempty" name:"VarName"`

	// 变量描述，最大支持120个字符
	VarDesc *string `json:"VarDesc,omitnil,omitempty" name:"VarDesc"`

	// 变量类型定义，支持类型如下：(STRING,INT,FLOAT,BOOL,OBJECT,ARRAY_STRING,ARRAY_INT,ARRAY_FLOAT,ARRAY_BOOL,ARRAY_OBJECT,FILE,DOCUMENT,IMAGE,AUDIO);传输过程是json字符串，标签中仅支持"STRING"类型使用
	VarType *string `json:"VarType,omitnil,omitempty" name:"VarType"`

	// 自定义变量默认值
	VarDefaultValue *string `json:"VarDefaultValue,omitnil,omitempty" name:"VarDefaultValue"`

	// 自定义变量文件默认名称
	VarDefaultFileName *string `json:"VarDefaultFileName,omitnil,omitempty" name:"VarDefaultFileName"`
}

type CreateVarRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量名称，不允许重复，最大支持50个字符
	VarName *string `json:"VarName,omitnil,omitempty" name:"VarName"`

	// 变量描述，最大支持120个字符
	VarDesc *string `json:"VarDesc,omitnil,omitempty" name:"VarDesc"`

	// 变量类型定义，支持类型如下：(STRING,INT,FLOAT,BOOL,OBJECT,ARRAY_STRING,ARRAY_INT,ARRAY_FLOAT,ARRAY_BOOL,ARRAY_OBJECT,FILE,DOCUMENT,IMAGE,AUDIO);传输过程是json字符串，标签中仅支持"STRING"类型使用
	VarType *string `json:"VarType,omitnil,omitempty" name:"VarType"`

	// 自定义变量默认值
	VarDefaultValue *string `json:"VarDefaultValue,omitnil,omitempty" name:"VarDefaultValue"`

	// 自定义变量文件默认名称
	VarDefaultFileName *string `json:"VarDefaultFileName,omitnil,omitempty" name:"VarDefaultFileName"`
}

func (r *CreateVarRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVarRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "VarName")
	delete(f, "VarDesc")
	delete(f, "VarType")
	delete(f, "VarDefaultValue")
	delete(f, "VarDefaultFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVarRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVarResponseParams struct {
	// 变量ID
	VarId *string `json:"VarId,omitnil,omitempty" name:"VarId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVarResponse struct {
	*tchttp.BaseResponse
	Response *CreateVarResponseParams `json:"Response"`
}

func (r *CreateVarResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVarResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRunRequestParams struct {
	// 运行环境。0: 测试环境； 1: 正式环境
	RunEnv *uint64 `json:"RunEnv,omitnil,omitempty" name:"RunEnv"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 用户输入的内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// API参数配置
	CustomVariables []*CustomVariable `json:"CustomVariables,omitnil,omitempty" name:"CustomVariables"`
}

type CreateWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// 运行环境。0: 测试环境； 1: 正式环境
	RunEnv *uint64 `json:"RunEnv,omitnil,omitempty" name:"RunEnv"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 用户输入的内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// API参数配置
	CustomVariables []*CustomVariable `json:"CustomVariables,omitnil,omitempty" name:"CustomVariables"`
}

func (r *CreateWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunEnv")
	delete(f, "AppBizId")
	delete(f, "Query")
	delete(f, "CustomVariables")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkflowRunResponseParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 工作流运行实例的ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 运行环境。0: 测试环境； 1: 正式环境
	RunEnv *uint64 `json:"RunEnv,omitnil,omitempty" name:"RunEnv"`

	// 用户输入的内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// API参数配置
	CustomVariables []*CustomVariable `json:"CustomVariables,omitnil,omitempty" name:"CustomVariables"`

	// 创建时间（毫秒时间戳）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkflowRunResponseParams `json:"Response"`
}

func (r *CreateWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {
	// token
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 临时证书密钥ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretId *string `json:"TmpSecretId,omitnil,omitempty" name:"TmpSecretId"`

	// 临时证书密钥Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TmpSecretKey *string `json:"TmpSecretKey,omitnil,omitempty" name:"TmpSecretKey"`

	// 临时证书appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppId *uint64 `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type CustomVariable struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数的值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type DeleteAgentRequestParams struct {
	// Agent的ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

type DeleteAgentRequest struct {
	*tchttp.BaseRequest
	
	// Agent的ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

func (r *DeleteAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AgentId")
	delete(f, "AppBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentResponseParams struct {
	// Agent的ID
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAgentResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAgentResponseParams `json:"Response"`
}

func (r *DeleteAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

func (r *DeleteAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAppResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppResponseParams `json:"Response"`
}

func (r *DeleteAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttributeLabelRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 标签ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DeleteAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 标签ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DeleteAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAttributeLabelResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAttributeLabelResponseParams `json:"Response"`
}

func (r *DeleteAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocCateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type DeleteDocCateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *DeleteDocCateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocCateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocCateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocCateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDocCateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocCateResponseParams `json:"Response"`
}

func (r *DeleteDocCateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocCateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocRequestParams struct {
	// 文档业务ID列表
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DeleteDocRequest struct {
	*tchttp.BaseRequest
	
	// 文档业务ID列表
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DeleteDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocBizIds")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteDocResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteDocResponse struct {
	*tchttp.BaseResponse
	Response *DeleteDocResponseParams `json:"Response"`
}

func (r *DeleteDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQACateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type DeleteQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *DeleteQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQACateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQACateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQACateResponseParams `json:"Response"`
}

func (r *DeleteQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQARequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

type DeleteQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`
}

func (r *DeleteQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteQAResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQAResponseParams `json:"Response"`
}

func (r *DeleteQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRejectedQuestionRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizIds []*string `json:"RejectedBizIds,omitnil,omitempty" name:"RejectedBizIds"`
}

type DeleteRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizIds []*string `json:"RejectedBizIds,omitnil,omitempty" name:"RejectedBizIds"`
}

func (r *DeleteRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "RejectedBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRejectedQuestionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRejectedQuestionResponseParams `json:"Response"`
}

func (r *DeleteRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSharedKnowledgeRequestParams struct {
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`
}

type DeleteSharedKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`
}

func (r *DeleteSharedKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSharedKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSharedKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSharedKnowledgeResponseParams struct {
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSharedKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSharedKnowledgeResponseParams `json:"Response"`
}

func (r *DeleteSharedKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSharedKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVarRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量ID
	VarId *string `json:"VarId,omitnil,omitempty" name:"VarId"`
}

type DeleteVarRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量ID
	VarId *string `json:"VarId,omitnil,omitempty" name:"VarId"`
}

func (r *DeleteVarRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVarRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "VarId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVarRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVarResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVarResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVarResponseParams `json:"Response"`
}

func (r *DeleteVarResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVarResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppAgentListRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

type DescribeAppAgentListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

func (r *DescribeAppAgentListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppAgentListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppAgentListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppAgentListResponseParams struct {
	// 入口启动AgentID
	StaringAgentId *string `json:"StaringAgentId,omitnil,omitempty" name:"StaringAgentId"`

	// 应用Agent信息列表
	Agents []*Agent `json:"Agents,omitnil,omitempty" name:"Agents"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppAgentListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppAgentListResponseParams `json:"Response"`
}

func (r *DescribeAppAgentListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 是否发布后的配置
	IsRelease *bool `json:"IsRelease,omitnil,omitempty" name:"IsRelease"`
}

type DescribeAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 是否发布后的配置
	IsRelease *bool `json:"IsRelease,omitnil,omitempty" name:"IsRelease"`
}

func (r *DescribeAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	delete(f, "IsRelease")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppResponseParams struct {
	// 应用 ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用类型说明
	AppTypeDesc *string `json:"AppTypeDesc,omitnil,omitempty" name:"AppTypeDesc"`

	// 应用类型说明
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`

	// 头像是否在申诉中
	AvatarInAppeal *bool `json:"AvatarInAppeal,omitnil,omitempty" name:"AvatarInAppeal"`

	// 角色描述是否在申诉中
	RoleInAppeal *bool `json:"RoleInAppeal,omitnil,omitempty" name:"RoleInAppeal"`

	// 名称是否在申诉中
	NameInAppeal *bool `json:"NameInAppeal,omitnil,omitempty" name:"NameInAppeal"`

	// 欢迎语是否在申诉中
	GreetingInAppeal *bool `json:"GreetingInAppeal,omitnil,omitempty" name:"GreetingInAppeal"`

	// 未知问题回复语是否在申诉中
	BareAnswerInAppeal *bool `json:"BareAnswerInAppeal,omitnil,omitempty" name:"BareAnswerInAppeal"`

	// 应用appKey
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 应用状态，1：未上线，2：运行中，3：停用
	AppStatus *uint64 `json:"AppStatus,omitnil,omitempty" name:"AppStatus"`

	// 状态说明
	AppStatusDesc *string `json:"AppStatusDesc,omitnil,omitempty" name:"AppStatusDesc"`

	// 应用是否在复制中
	IsCopying *bool `json:"IsCopying,omitnil,omitempty" name:"IsCopying"`

	// 智能体类型 dialogue 对话式智能体，wechat 公众号智能体
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppResponseParams `json:"Response"`
}

func (r *DescribeAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttributeLabelRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 每次加载的数量 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询标签或相似标签
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 滚动加载游标的标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`

	// 查询范围 all(或者传空):标准词和相似词 standard:标准词 similar:相似词
	QueryScope *string `json:"QueryScope,omitnil,omitempty" name:"QueryScope"`
}

type DescribeAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 每次加载的数量 
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询标签或相似标签
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 滚动加载游标的标签ID
	LastLabelBizId *string `json:"LastLabelBizId,omitnil,omitempty" name:"LastLabelBizId"`

	// 查询范围 all(或者传空):标准词和相似词 standard:标准词 similar:相似词
	QueryScope *string `json:"QueryScope,omitnil,omitempty" name:"QueryScope"`
}

func (r *DescribeAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizId")
	delete(f, "Limit")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	delete(f, "LastLabelBizId")
	delete(f, "QueryScope")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAttributeLabelResponseParams struct {
	// 属性ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 属性标识
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 属性名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签数量
	LabelNumber *string `json:"LabelNumber,omitnil,omitempty" name:"LabelNumber"`

	// 标签名称
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAttributeLabelResponseParams `json:"Response"`
}

func (r *DescribeAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallStatsGraphRequestParams struct {
	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 筛选子场景(文档解析场景使用)
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type DescribeCallStatsGraphRequest struct {
	*tchttp.BaseRequest
	
	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 筛选子场景(文档解析场景使用)
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

func (r *DescribeCallStatsGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallStatsGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	delete(f, "SubScenes")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCallStatsGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCallStatsGraphResponseParams struct {
	// 接口调用次数统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Stat `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCallStatsGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCallStatsGraphResponseParams `json:"Response"`
}

func (r *DescribeCallStatsGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCallStatsGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageGraphRequestParams struct {
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeConcurrencyUsageGraphRequest struct {
	*tchttp.BaseRequest
	
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// uin
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeConcurrencyUsageGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyUsageGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageGraphResponseParams struct {
	// X轴: 时间区域；根据查询条件的粒度返回“分/小时/日”两种区间范围
	X []*string `json:"X,omitnil,omitempty" name:"X"`

	// 可用并发y轴坐标
	AvailableY []*int64 `json:"AvailableY,omitnil,omitempty" name:"AvailableY"`

	// 成功调用并发y轴坐标
	SuccessCallY []*int64 `json:"SuccessCallY,omitnil,omitempty" name:"SuccessCallY"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyUsageGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyUsageGraphResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyUsageGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageRequestParams struct {
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeConcurrencyUsageRequest struct {
	*tchttp.BaseRequest
	
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeConcurrencyUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyUsageResponseParams struct {
	// 可用并发数上限
	AvailableConcurrency *uint64 `json:"AvailableConcurrency,omitnil,omitempty" name:"AvailableConcurrency"`

	// 并发峰值
	ConcurrencyPeak *uint64 `json:"ConcurrencyPeak,omitnil,omitempty" name:"ConcurrencyPeak"`

	// 超出可用并发数上限的次数
	ExceedUsageTime *uint64 `json:"ExceedUsageTime,omitnil,omitempty" name:"ExceedUsageTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyUsageResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type DescribeDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *DescribeDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeDocResponseParams struct {
	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文档状态： 1-未生成 2-生成中 3-生成成功 4-生成失败 5-删除中 6-删除成功 7-审核中 8-审核失败 9-审核成功 10-待发布 11-发布中 12-已发布 13-学习中 14-学习失败 15-更新中 16-更新失败 17-解析中 18-解析失败 19-导入失败 20-已过期 21-超量失效 22-超量失效恢复
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 文档状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 生成失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 答案中是否引用
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 问答对数量
	QaNum *int64 `json:"QaNum,omitnil,omitempty" name:"QaNum"`

	// 是否删除
	IsDeleted *bool `json:"IsDeleted,omitnil,omitempty" name:"IsDeleted"`

	// 文档来源
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 文档来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 是否允许重新生成
	IsAllowRestart *bool `json:"IsAllowRestart,omitnil,omitempty" name:"IsAllowRestart"`

	// qa是否已删除
	IsDeletedQa *bool `json:"IsDeletedQa,omitnil,omitempty" name:"IsDeletedQa"`

	// 问答是否生成中
	IsCreatingQa *bool `json:"IsCreatingQa,omitnil,omitempty" name:"IsCreatingQa"`

	// 是否允许删除
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许操作引用开关
	IsAllowRefer *bool `json:"IsAllowRefer,omitnil,omitempty" name:"IsAllowRefer"`

	// 是否生成过问答
	IsCreatedQa *bool `json:"IsCreatedQa,omitnil,omitempty" name:"IsCreatedQa"`

	// 文档字符量
	DocCharSize *string `json:"DocCharSize,omitnil,omitempty" name:"DocCharSize"`

	// 是否允许编辑
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 标签适用范围 1：全部，2：按条件范围
	AttrRange *int64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 标签
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 文档是否停用，false:未停用，true:已停用
	IsDisabled *bool `json:"IsDisabled,omitnil,omitempty" name:"IsDisabled"`

	// 是否支持下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDownload *bool `json:"IsDownload,omitnil,omitempty" name:"IsDownload"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeDocResponse struct {
	*tchttp.BaseResponse
	Response *DescribeDocResponseParams `json:"Response"`
}

func (r *DescribeDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsagePieGraphRequestParams struct {
	// 应用ID数组
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeKnowledgeUsagePieGraphRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID数组
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeKnowledgeUsagePieGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsagePieGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeUsagePieGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsagePieGraphResponseParams struct {
	// 所有应用已用的字符总数
	AvailableCharSize *string `json:"AvailableCharSize,omitnil,omitempty" name:"AvailableCharSize"`

	// 应用饼图详情列表
	List []*KnowledgeCapacityPieGraphDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeUsagePieGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeUsagePieGraphResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeUsagePieGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsagePieGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsageRequestParams struct {

}

type DescribeKnowledgeUsageRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeKnowledgeUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeKnowledgeUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeKnowledgeUsageResponseParams struct {
	// 可用字符数上限
	AvailableCharSize *string `json:"AvailableCharSize,omitnil,omitempty" name:"AvailableCharSize"`

	// 超过可用字符数上限的字符数
	ExceedCharSize *string `json:"ExceedCharSize,omitnil,omitempty" name:"ExceedCharSize"`

	// 知识库使用字符总数
	UsedCharSize *string `json:"UsedCharSize,omitnil,omitempty" name:"UsedCharSize"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeKnowledgeUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeKnowledgeUsageResponseParams `json:"Response"`
}

func (r *DescribeKnowledgeUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeKnowledgeUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeRunRequestParams struct {
	// 节点运行实例ID
	NodeRunId *string `json:"NodeRunId,omitnil,omitempty" name:"NodeRunId"`
}

type DescribeNodeRunRequest struct {
	*tchttp.BaseRequest
	
	// 节点运行实例ID
	NodeRunId *string `json:"NodeRunId,omitnil,omitempty" name:"NodeRunId"`
}

func (r *DescribeNodeRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "NodeRunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodeRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodeRunResponseParams struct {
	// 节点运行实例详情
	NodeRun *NodeRunDetail `json:"NodeRun,omitnil,omitempty" name:"NodeRun"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNodeRunResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodeRunResponseParams `json:"Response"`
}

func (r *DescribeNodeRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodeRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQARequestParams struct {
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DescribeQARequest struct {
	*tchttp.BaseRequest
	
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DescribeQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QaBizId")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQAResponseParams struct {
	// QA业务ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 来源 1-文档生成问答对  2-批量导入问答对  3-单条手动录入问答对
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态 <br>1-未校验  2-未发布 3-发布中 4-已发布  5-发布失败 6-不采纳 7-审核中  8-审核失败  9-审核失败申诉后人工审核中  11-审核失败申诉后人工审核不通过  12-已过期  13-超量失效  14-超量失效恢复 19-学习中  20-学习失败
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 是否允许校验
	IsAllowAccept *bool `json:"IsAllowAccept,omitnil,omitempty" name:"IsAllowAccept"`

	// 是否允许删除
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许编辑
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 文档id
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 分片ID
	SegmentBizId *string `json:"SegmentBizId,omitnil,omitempty" name:"SegmentBizId"`

	// 分片内容
	PageContent *string `json:"PageContent,omitnil,omitempty" name:"PageContent"`

	// 分片高亮内容
	Highlights []*Highlight `json:"Highlights,omitnil,omitempty" name:"Highlights"`

	// 分片内容
	OrgData *string `json:"OrgData,omitnil,omitempty" name:"OrgData"`

	// 标签适用范围
	AttrRange *int64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 标签
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 相似问列表信息
	SimilarQuestions []*SimilarQuestion `json:"SimilarQuestions,omitnil,omitempty" name:"SimilarQuestions"`

	// 问题和答案文本审核状态 1审核失败
	QaAuditStatus *uint64 `json:"QaAuditStatus,omitnil,omitempty" name:"QaAuditStatus"`

	// 答案中的图片审核状态 1审核失败
	PicAuditStatus *uint64 `json:"PicAuditStatus,omitnil,omitempty" name:"PicAuditStatus"`

	// 答案中的视频审核状态 1审核失败
	VideoAuditStatus *uint64 `json:"VideoAuditStatus,omitnil,omitempty" name:"VideoAuditStatus"`

	// 问题描述
	QuestionDesc *string `json:"QuestionDesc,omitnil,omitempty" name:"QuestionDesc"`

	// 问答是否停用，false:未停用，true已停用
	IsDisabled *bool `json:"IsDisabled,omitnil,omitempty" name:"IsDisabled"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeQAResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQAResponseParams `json:"Response"`
}

func (r *DescribeQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReferRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 引用ID
	ReferBizIds []*string `json:"ReferBizIds,omitnil,omitempty" name:"ReferBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DescribeReferRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 引用ID
	ReferBizIds []*string `json:"ReferBizIds,omitnil,omitempty" name:"ReferBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DescribeReferRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReferRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReferBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReferRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReferResponseParams struct {
	// 引用列表
	List []*ReferDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReferResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReferResponseParams `json:"Response"`
}

func (r *DescribeReferResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReferResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseInfoRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type DescribeReleaseInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *DescribeReleaseInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseInfoResponseParams struct {
	// 最后发布时间
	LastTime *string `json:"LastTime,omitnil,omitempty" name:"LastTime"`

	// 发布状态 ， 1-待发布 , 2-发布中 , 3-发布成功 , 4-发布失败 , 5-审核中 , 6-审核成功 , 7-审核失败 , 8-发布成功回调处理中 , 9-发布暂停 , 10-申诉审核中 , 11-申诉审核通过 , 12-申诉审核不通过
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 是否编辑过, 当为true的时候表示可以发布
	IsUpdated *bool `json:"IsUpdated,omitnil,omitempty" name:"IsUpdated"`

	// 失败原因
	Msg *string `json:"Msg,omitnil,omitempty" name:"Msg"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseInfoResponseParams `json:"Response"`
}

func (r *DescribeReleaseInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布详情
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

type DescribeReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布详情
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

func (r *DescribeReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReleaseBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseResponseParams struct {
	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 发布描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 发布状态(1待发布 2发布中 3发布成功 4发布失败 5发布中(审核中) 6发布中(审核完成) 7发布失败(审核失败) 9发布暂停)
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 发布状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseResponseParams `json:"Response"`
}

func (r *DescribeReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRobotBizIDByAppKeyRequestParams struct {
	// 应用appkey
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`
}

type DescribeRobotBizIDByAppKeyRequest struct {
	*tchttp.BaseRequest
	
	// 应用appkey
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`
}

func (r *DescribeRobotBizIDByAppKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRobotBizIDByAppKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRobotBizIDByAppKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRobotBizIDByAppKeyResponseParams struct {
	// 应用业务ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRobotBizIDByAppKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRobotBizIDByAppKeyResponseParams `json:"Response"`
}

func (r *DescribeRobotBizIDByAppKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRobotBizIDByAppKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchStatsGraphRequestParams struct {
	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// uin列表
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type DescribeSearchStatsGraphRequest struct {
	*tchttp.BaseRequest
	
	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// uin列表
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 子业务类型
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *DescribeSearchStatsGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchStatsGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "UinAccount")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSearchStatsGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSearchStatsGraphResponseParams struct {
	// 统计结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Stat `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSearchStatsGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSearchStatsGraphResponseParams `json:"Response"`
}

func (r *DescribeSearchStatsGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSearchStatsGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentsRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档片段ID
	SegBizId []*string `json:"SegBizId,omitnil,omitempty" name:"SegBizId"`
}

type DescribeSegmentsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档片段ID
	SegBizId []*string `json:"SegBizId,omitnil,omitempty" name:"SegBizId"`
}

func (r *DescribeSegmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "SegBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSegmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSegmentsResponseParams struct {
	// 片段列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*DocSegment `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSegmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSegmentsResponseParams `json:"Response"`
}

func (r *DescribeSegmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSegmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSharedKnowledgeRequestParams struct {
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`
}

type DescribeSharedKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`
}

func (r *DescribeSharedKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSharedKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSharedKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSharedKnowledgeResponseParams struct {
	// 知识库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *KnowledgeDetailInfo `json:"Info,omitnil,omitempty" name:"Info"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSharedKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSharedKnowledgeResponseParams `json:"Response"`
}

func (r *DescribeSharedKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSharedKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageCredentialRequestParams struct {
	// 应用ID，参数非必填不代表不需要填写，下面不同的参数组合会获取到不同的权限，具体请参考 https://cloud.tencent.com/document/product/1759/116238
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件类型,正常的文件名类型后缀，例如 xlsx、pdf、 docx、png 等
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// IsPublic用于上传文件或图片时选择场景，当上传对话端图片时IsPublic为true，上传文件（包括文档库文件/图片等和对话端文件）时IsPublic为false
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
}

type DescribeStorageCredentialRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，参数非必填不代表不需要填写，下面不同的参数组合会获取到不同的权限，具体请参考 https://cloud.tencent.com/document/product/1759/116238
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件类型,正常的文件名类型后缀，例如 xlsx、pdf、 docx、png 等
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// IsPublic用于上传文件或图片时选择场景，当上传对话端图片时IsPublic为true，上传文件（包括文档库文件/图片等和对话端文件）时IsPublic为false
	IsPublic *bool `json:"IsPublic,omitnil,omitempty" name:"IsPublic"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
}

func (r *DescribeStorageCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileType")
	delete(f, "IsPublic")
	delete(f, "TypeKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeStorageCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeStorageCredentialResponseParams struct {
	// 密钥信息
	Credentials *Credentials `json:"Credentials,omitnil,omitempty" name:"Credentials"`

	// 失效时间
	ExpiredTime *int64 `json:"ExpiredTime,omitnil,omitempty" name:"ExpiredTime"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 对象存储桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 对象存储可用区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 文件存储目录
	FilePath *string `json:"FilePath,omitnil,omitempty" name:"FilePath"`

	// 存储类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 主号
	CorpUin *string `json:"CorpUin,omitnil,omitempty" name:"CorpUin"`

	// 图片存储目录
	ImagePath *string `json:"ImagePath,omitnil,omitempty" name:"ImagePath"`

	// 上传存储路径，到具体文件
	UploadPath *string `json:"UploadPath,omitnil,omitempty" name:"UploadPath"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeStorageCredentialResponse struct {
	*tchttp.BaseResponse
	Response *DescribeStorageCredentialResponseParams `json:"Response"`
}

func (r *DescribeStorageCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeStorageCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageGraphRequestParams struct {
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type DescribeTokenUsageGraphRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

func (r *DescribeTokenUsageGraphRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageGraphRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenUsageGraphRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageGraphResponseParams struct {
	// Token消耗总量
	Total []*Stat `json:"Total,omitnil,omitempty" name:"Total"`

	// 输入Token消耗量
	Input []*Stat `json:"Input,omitnil,omitempty" name:"Input"`

	// 输出Token消耗量
	Output []*Stat `json:"Output,omitnil,omitempty" name:"Output"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenUsageGraphResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenUsageGraphResponseParams `json:"Response"`
}

func (r *DescribeTokenUsageGraphResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageGraphResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageRequestParams struct {
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒(默认值0)
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒(默认值0， 必须大于开始时间戳)
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 筛选子场景(文档解析场景使用)
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

type DescribeTokenUsageRequest struct {
	*tchttp.BaseRequest
	
	// 腾讯云主账号
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 知识引擎子业务类型:  FileParse(文档解析)、Embedding、Rewrite(多轮改写)、 Concurrency(并发)、KnowledgeSummary(知识总结)   KnowledgeQA(知识问答)、KnowledgeCapacity(知识库容量)、SearchEngine(搜索引擎)
	SubBizType *string `json:"SubBizType,omitnil,omitempty" name:"SubBizType"`

	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间戳, 单位为秒(默认值0)
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳, 单位为秒(默认值0， 必须大于开始时间戳)
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 筛选子场景(文档解析场景使用)
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`
}

func (r *DescribeTokenUsageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "UinAccount")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "SubBizType")
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizIds")
	delete(f, "SubScenes")
	delete(f, "AppType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTokenUsageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTokenUsageResponseParams struct {
	// 总token消耗量
	TotalTokenUsage *float64 `json:"TotalTokenUsage,omitnil,omitempty" name:"TotalTokenUsage"`

	// 输入token消耗
	InputTokenUsage *float64 `json:"InputTokenUsage,omitnil,omitempty" name:"InputTokenUsage"`

	// 输出token消耗
	OutputTokenUsage *float64 `json:"OutputTokenUsage,omitnil,omitempty" name:"OutputTokenUsage"`

	// 接口调用次数
	ApiCallStats *uint64 `json:"ApiCallStats,omitnil,omitempty" name:"ApiCallStats"`

	// 搜索服务调用次数
	SearchUsage *float64 `json:"SearchUsage,omitnil,omitempty" name:"SearchUsage"`

	// 文档解析消耗页数
	PageUsage *uint64 `json:"PageUsage,omitnil,omitempty" name:"PageUsage"`

	// 拆分token消耗量
	SplitTokenUsage *float64 `json:"SplitTokenUsage,omitnil,omitempty" name:"SplitTokenUsage"`

	// Rag检索次数
	RagSearchUsage *float64 `json:"RagSearchUsage,omitnil,omitempty" name:"RagSearchUsage"`

	// 联网搜索次数
	InternetSearchUsage *float64 `json:"InternetSearchUsage,omitnil,omitempty" name:"InternetSearchUsage"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTokenUsageResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTokenUsageResponseParams `json:"Response"`
}

func (r *DescribeTokenUsageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTokenUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnsatisfiedReplyContextRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 回复ID
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DescribeUnsatisfiedReplyContextRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 回复ID
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DescribeUnsatisfiedReplyContextRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnsatisfiedReplyContextRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUnsatisfiedReplyContextRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUnsatisfiedReplyContextResponseParams struct {
	// 不满意回复上下文
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*Context `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUnsatisfiedReplyContextResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUnsatisfiedReplyContextResponseParams `json:"Response"`
}

func (r *DescribeUnsatisfiedReplyContextResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUnsatisfiedReplyContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowRunRequestParams struct {
	// 工作流运行实例ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type DescribeWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// 工作流运行实例ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

func (r *DescribeWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowRunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWorkflowRunResponseParams struct {
	// 总数
	WorkflowRun *WorkflowRunDetail `json:"WorkflowRun,omitnil,omitempty" name:"WorkflowRun"`

	// 节点列表
	NodeRuns []*NodeRunBase `json:"NodeRuns,omitnil,omitempty" name:"NodeRuns"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWorkflowRunResponseParams `json:"Response"`
}

func (r *DescribeWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DigitalHumanConfig struct {
	// 数智人资产key
	AssetKey *string `json:"AssetKey,omitnil,omitempty" name:"AssetKey"`

	// 数智人名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 图像
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 预览图
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`
}

type DocFilterFlag struct {
	// 标识位
	Flag *string `json:"Flag,omitnil,omitempty" name:"Flag"`

	// 标识值
	Value *bool `json:"Value,omitnil,omitempty" name:"Value"`
}

type DocSegment struct {
	// 片段ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	BusinessId *string `json:"BusinessId,omitnil,omitempty" name:"BusinessId"`

	// 文件类型(markdown,word,txt)
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文档切片类型(segment-文档切片 table-表格)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentType *string `json:"SegmentType,omitnil,omitempty" name:"SegmentType"`

	// 标题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 段落内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageContent *string `json:"PageContent,omitnil,omitempty" name:"PageContent"`

	// 段落原文
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgData *string `json:"OrgData,omitnil,omitempty" name:"OrgData"`

	// 文章ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// 文档业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文档链接
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocUrl *string `json:"DocUrl,omitnil,omitempty" name:"DocUrl"`

	// 文档的自定义链接
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 页码信息
	PageInfos []*uint64 `json:"PageInfos,omitnil,omitempty" name:"PageInfos"`
}

// Predefined struct for user
type ExportAttributeLabelRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 根据筛选数据导出
	Filters *AttributeFilters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 属性ID
	AttributeBizIds []*string `json:"AttributeBizIds,omitnil,omitempty" name:"AttributeBizIds"`

	// 根据筛选数据导出
	Filters *AttributeFilters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "AttributeBizIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportAttributeLabelResponseParams struct {
	// 导出任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ExportAttributeLabelResponseParams `json:"Response"`
}

func (r *ExportAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportQAListRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QA业务ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 查询参数
	Filters *QAQuery `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportQAListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QA业务ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 查询参数
	Filters *QAQuery `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportQAListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQAListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportQAListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportQAListResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportQAListResponse struct {
	*tchttp.BaseResponse
	Response *ExportQAListResponseParams `json:"Response"`
}

func (r *ExportQAListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportQAListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportUnsatisfiedReplyRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 勾选导出ID列表
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 检索过滤器
	Filters *Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type ExportUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 勾选导出ID列表
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 检索过滤器
	Filters *Filters `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *ExportUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExportUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExportUnsatisfiedReplyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExportUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *ExportUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *ExportUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExportUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtraInfo struct {
	// ECharts信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EChartsInfo []*string `json:"EChartsInfo,omitnil,omitempty" name:"EChartsInfo"`
}

type FileInfo struct {
	// 文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文件的URL地址，COS地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 解析后返回的DocID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`
}

type Filters struct {
	// 检索，用户问题或答案
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 错误类型检索
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type GenerateQARequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`
}

type GenerateQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`
}

func (r *GenerateQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GenerateQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GenerateQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GenerateQAResponse struct {
	*tchttp.BaseResponse
	Response *GenerateQAResponseParams `json:"Response"`
}

func (r *GenerateQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GenerateQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAnswerTypeDataCountRequestParams struct {
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API  3、对话测试  4、应用评测)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type GetAnswerTypeDataCountRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API  3、对话测试  4、应用评测)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *GetAnswerTypeDataCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAnswerTypeDataCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizId")
	delete(f, "Type")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAnswerTypeDataCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAnswerTypeDataCountResponseParams struct {
	// 总消息数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 大模型直接回复总数
	ModelReplyCount *uint64 `json:"ModelReplyCount,omitnil,omitempty" name:"ModelReplyCount"`

	// 知识型回复总数
	KnowledgeCount *uint64 `json:"KnowledgeCount,omitnil,omitempty" name:"KnowledgeCount"`

	// 任务流回复总数
	TaskFlowCount *uint64 `json:"TaskFlowCount,omitnil,omitempty" name:"TaskFlowCount"`

	// 搜索引擎回复总数
	SearchEngineCount *uint64 `json:"SearchEngineCount,omitnil,omitempty" name:"SearchEngineCount"`

	// 图片理解回复总数
	ImageUnderstandingCount *uint64 `json:"ImageUnderstandingCount,omitnil,omitempty" name:"ImageUnderstandingCount"`

	// 拒答回复总数
	RejectCount *uint64 `json:"RejectCount,omitnil,omitempty" name:"RejectCount"`

	// 敏感回复总数
	SensitiveCount *uint64 `json:"SensitiveCount,omitnil,omitempty" name:"SensitiveCount"`

	// 并发超限回复总数
	ConcurrentLimitCount *uint64 `json:"ConcurrentLimitCount,omitnil,omitempty" name:"ConcurrentLimitCount"`

	// 未知问题回复总数
	UnknownIssuesCount *uint64 `json:"UnknownIssuesCount,omitnil,omitempty" name:"UnknownIssuesCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAnswerTypeDataCountResponse struct {
	*tchttp.BaseResponse
	Response *GetAnswerTypeDataCountResponseParams `json:"Response"`
}

func (r *GetAnswerTypeDataCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAnswerTypeDataCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppKnowledgeCountRequestParams struct {
	// 类型：doc-文档；qa-问答对
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type GetAppKnowledgeCountRequest struct {
	*tchttp.BaseRequest
	
	// 类型：doc-文档；qa-问答对
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *GetAppKnowledgeCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppKnowledgeCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAppKnowledgeCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppKnowledgeCountResponseParams struct {
	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAppKnowledgeCountResponse struct {
	*tchttp.BaseResponse
	Response *GetAppKnowledgeCountResponseParams `json:"Response"`
}

func (r *GetAppKnowledgeCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppKnowledgeCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppSecretRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

type GetAppSecretRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`
}

func (r *GetAppSecretRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSecretRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetAppSecretRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetAppSecretResponseParams struct {
	// 应用密钥
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否发布
	IsRelease *bool `json:"IsRelease,omitnil,omitempty" name:"IsRelease"`

	// 是否有查看权限
	HasPermission *bool `json:"HasPermission,omitnil,omitempty" name:"HasPermission"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetAppSecretResponse struct {
	*tchttp.BaseResponse
	Response *GetAppSecretResponseParams `json:"Response"`
}

func (r *GetAppSecretResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetAppSecretResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDocPreviewRequestParams struct {
	// 文档BizID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
}

type GetDocPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 文档BizID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 存储类型: offline:离线文件，realtime:实时文件；为空默认为offline
	TypeKey *string `json:"TypeKey,omitnil,omitempty" name:"TypeKey"`
}

func (r *GetDocPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DocBizId")
	delete(f, "BotBizId")
	delete(f, "TypeKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetDocPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetDocPreviewResponseParams struct {
	// 文件名, 发布端固定使用这个名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// cos临时地址
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// cos桶
	Bucket *string `json:"Bucket,omitnil,omitempty" name:"Bucket"`

	// 存在文档重命名情况下的新名称, 评测端优先使用这个名称
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`

	// 文件md结果cos临时地址
	ParseResultCosUrl *string `json:"ParseResultCosUrl,omitnil,omitempty" name:"ParseResultCosUrl"`

	// 是否可下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDownload *bool `json:"IsDownload,omitnil,omitempty" name:"IsDownload"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetDocPreviewResponse struct {
	*tchttp.BaseResponse
	Response *GetDocPreviewResponseParams `json:"Response"`
}

func (r *GetDocPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetDocPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLikeDataCountRequestParams struct {
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type GetLikeDataCountRequest struct {
	*tchttp.BaseRequest
	
	// 开始日期
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束日期
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 应用id
	AppBizId []*string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 消息来源(1、分享用户端  2、对话API)
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *GetLikeDataCountRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLikeDataCountRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "AppBizId")
	delete(f, "Type")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetLikeDataCountRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetLikeDataCountResponseParams struct {
	// 可评价消息数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 评价数
	AppraisalTotal *uint64 `json:"AppraisalTotal,omitnil,omitempty" name:"AppraisalTotal"`

	// 参评率
	ParticipationRate *float64 `json:"ParticipationRate,omitnil,omitempty" name:"ParticipationRate"`

	// 点赞数
	LikeTotal *uint64 `json:"LikeTotal,omitnil,omitempty" name:"LikeTotal"`

	// 点赞率
	LikeRate *float64 `json:"LikeRate,omitnil,omitempty" name:"LikeRate"`

	// 点踩数
	DislikeTotal *uint64 `json:"DislikeTotal,omitnil,omitempty" name:"DislikeTotal"`

	// 点踩率
	DislikeRate *float64 `json:"DislikeRate,omitnil,omitempty" name:"DislikeRate"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetLikeDataCountResponse struct {
	*tchttp.BaseResponse
	Response *GetLikeDataCountResponseParams `json:"Response"`
}

func (r *GetLikeDataCountResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetLikeDataCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMsgRecordRequestParams struct {
	// 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数量,  数量需大于2, 最大1000
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 会话sessionid
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 应用AppKey, 当Type=5[API访客]时, 该字段必填  :</br>  获取方式:</br>   1、应用发布后在应用页面[发布管理]-[调用信息]-[API管理]处获取</br>   2、参考 https://cloud.tencent.com/document/product/1759/109469 第二项
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 场景, 体验: 1; 正式: 2
	Scene *uint64 `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 最后一条记录ID， 消息从后往前获取
	// 
	// MidRecordId与LastRecordId只能选择一个
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// 传该值，代表拉取该记录id的前后总共count条消息记录
	// 
	// MidRecordId与LastRecordId只能选择一个
	MidRecordId *string `json:"MidRecordId,omitnil,omitempty" name:"MidRecordId"`
}

type GetMsgRecordRequest struct {
	*tchttp.BaseRequest
	
	// 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 数量,  数量需大于2, 最大1000
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 会话sessionid
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 应用AppKey, 当Type=5[API访客]时, 该字段必填  :</br>  获取方式:</br>   1、应用发布后在应用页面[发布管理]-[调用信息]-[API管理]处获取</br>   2、参考 https://cloud.tencent.com/document/product/1759/109469 第二项
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 场景, 体验: 1; 正式: 2
	Scene *uint64 `json:"Scene,omitnil,omitempty" name:"Scene"`

	// 最后一条记录ID， 消息从后往前获取
	// 
	// MidRecordId与LastRecordId只能选择一个
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// 传该值，代表拉取该记录id的前后总共count条消息记录
	// 
	// MidRecordId与LastRecordId只能选择一个
	MidRecordId *string `json:"MidRecordId,omitnil,omitempty" name:"MidRecordId"`
}

func (r *GetMsgRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMsgRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "Count")
	delete(f, "SessionId")
	delete(f, "BotAppKey")
	delete(f, "Scene")
	delete(f, "LastRecordId")
	delete(f, "MidRecordId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetMsgRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetMsgRecordResponseParams struct {
	// 会话记录
	Records []*MsgRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// session 清除关联上下文时间, 单位 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionDisassociatedTimestamp *string `json:"SessionDisassociatedTimestamp,omitnil,omitempty" name:"SessionDisassociatedTimestamp"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetMsgRecordResponse struct {
	*tchttp.BaseResponse
	Response *GetMsgRecordResponseParams `json:"Response"`
}

func (r *GetMsgRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetMsgRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusRequestParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type GetTaskStatusRequest struct {
	*tchttp.BaseRequest
	
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *GetTaskStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TaskId")
	delete(f, "TaskType")
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetTaskStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetTaskStatusResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 任务类型
	TaskType *string `json:"TaskType,omitnil,omitempty" name:"TaskType"`

	// 任务状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 任务参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params *TaskParams `json:"Params,omitnil,omitempty" name:"Params"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetTaskStatusResponse struct {
	*tchttp.BaseResponse
	Response *GetTaskStatusResponseParams `json:"Response"`
}

func (r *GetTaskStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVarListRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量ID数组
	VarIds []*string `json:"VarIds,omitnil,omitempty" name:"VarIds"`

	// 按变量名称关键词搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 起始偏移量（默认0）
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限定数量（默认15）
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按变量类型过滤，默认查询所有类型(STRING,INT,FLOAT,BOOL,OBJECT,ARRAY_STRING,ARRAY_INT,ARRAY_FLOAT,ARRAY_BOOL,ARRAY_OBJECT,FILE,DOCUMENT,IMAGE,AUDIO)
	VarType *string `json:"VarType,omitnil,omitempty" name:"VarType"`

	// 是否需要内部变量(默认false)
	NeedInternalVar *bool `json:"NeedInternalVar,omitnil,omitempty" name:"NeedInternalVar"`
}

type GetVarListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量ID数组
	VarIds []*string `json:"VarIds,omitnil,omitempty" name:"VarIds"`

	// 按变量名称关键词搜索
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 起始偏移量（默认0）
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 限定数量（默认15）
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 按变量类型过滤，默认查询所有类型(STRING,INT,FLOAT,BOOL,OBJECT,ARRAY_STRING,ARRAY_INT,ARRAY_FLOAT,ARRAY_BOOL,ARRAY_OBJECT,FILE,DOCUMENT,IMAGE,AUDIO)
	VarType *string `json:"VarType,omitnil,omitempty" name:"VarType"`

	// 是否需要内部变量(默认false)
	NeedInternalVar *bool `json:"NeedInternalVar,omitnil,omitempty" name:"NeedInternalVar"`
}

func (r *GetVarListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVarListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "VarIds")
	delete(f, "Keyword")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "VarType")
	delete(f, "NeedInternalVar")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetVarListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetVarListResponseParams struct {
	// 变量总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 变量信息列表
	List []*TaskFLowVar `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetVarListResponse struct {
	*tchttp.BaseResponse
	Response *GetVarListResponseParams `json:"Response"`
}

func (r *GetVarListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetVarListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetWsTokenReq_Label struct {
	// 标签名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

// Predefined struct for user
type GetWsTokenRequestParams struct {
	// 接入类型， 5-API 访客
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	//   应用AppKey </br>   获取方式:</br>   1、应用发布后在应用页面[发布管理]-[调用信息]-[API管理]处获取</br>   2、参考 https://cloud.tencent.com/document/product/1759/109469 第二项
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 访客ID（外部输入，建议唯一，标识当前接入会话的用户）
	// 长度限制： string(64)
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// 知识标签，用于知识库中知识的检索过滤。该字段即将下线，请使用对话端接口中的 custom_variables 字段替代该字段。
	VisitorLabels []*GetWsTokenReq_Label `json:"VisitorLabels,omitnil,omitempty" name:"VisitorLabels"`
}

type GetWsTokenRequest struct {
	*tchttp.BaseRequest
	
	// 接入类型， 5-API 访客
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	//   应用AppKey </br>   获取方式:</br>   1、应用发布后在应用页面[发布管理]-[调用信息]-[API管理]处获取</br>   2、参考 https://cloud.tencent.com/document/product/1759/109469 第二项
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 访客ID（外部输入，建议唯一，标识当前接入会话的用户）
	// 长度限制： string(64)
	VisitorBizId *string `json:"VisitorBizId,omitnil,omitempty" name:"VisitorBizId"`

	// 知识标签，用于知识库中知识的检索过滤。该字段即将下线，请使用对话端接口中的 custom_variables 字段替代该字段。
	VisitorLabels []*GetWsTokenReq_Label `json:"VisitorLabels,omitnil,omitempty" name:"VisitorLabels"`
}

func (r *GetWsTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWsTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "BotAppKey")
	delete(f, "VisitorBizId")
	delete(f, "VisitorLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetWsTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetWsTokenResponseParams struct {
	// token值（有效期60s，仅一次有效，多次校验会报错）
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 余额; 余额大于 0 时表示有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Balance *float64 `json:"Balance,omitnil,omitempty" name:"Balance"`

	// 对话窗输入字符限制
	InputLenLimit *int64 `json:"InputLenLimit,omitnil,omitempty" name:"InputLenLimit"`

	// 应用模式，standard:标准模式, agent: agent模式，single_workflow：单工作流模式
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// SingleWorkflow
	SingleWorkflow *KnowledgeQaSingleWorkflow `json:"SingleWorkflow,omitnil,omitempty" name:"SingleWorkflow"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GetWsTokenResponse struct {
	*tchttp.BaseResponse
	Response *GetWsTokenResponseParams `json:"Response"`
}

func (r *GetWsTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetWsTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 操作对象的业务ID列表
	BizIds []*string `json:"BizIds,omitnil,omitempty" name:"BizIds"`

	// 分组 ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type GroupDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 操作对象的业务ID列表
	BizIds []*string `json:"BizIds,omitnil,omitempty" name:"BizIds"`

	// 分组 ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *GroupDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "BizIds")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GroupDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupDocResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GroupDocResponse struct {
	*tchttp.BaseResponse
	Response *GroupDocResponseParams `json:"Response"`
}

func (r *GroupDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupQARequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QaBizID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 分组 ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type GroupQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// QaBizID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 分组 ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *GroupQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GroupQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GroupQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type GroupQAResponse struct {
	*tchttp.BaseResponse
	Response *GroupQAResponseParams `json:"Response"`
}

func (r *GroupQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GroupQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Highlight struct {
	// 高亮起始位置
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartPos *string `json:"StartPos,omitnil,omitempty" name:"StartPos"`

	// 高亮结束位置
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndPos *string `json:"EndPos,omitnil,omitempty" name:"EndPos"`

	// 高亮子文本
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`
}

type HistorySummary struct {
	// 助手
	// 注意：此字段可能返回 null，表示取不到有效值。
	Assistant *string `json:"Assistant,omitnil,omitempty" name:"Assistant"`

	// 用户
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *string `json:"User,omitnil,omitempty" name:"User"`
}

// Predefined struct for user
type IgnoreUnsatisfiedReplyRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 不满意回复ID
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type IgnoreUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 不满意回复ID
	ReplyBizIds []*string `json:"ReplyBizIds,omitnil,omitempty" name:"ReplyBizIds"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *IgnoreUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReplyBizIds")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IgnoreUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IgnoreUnsatisfiedReplyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IgnoreUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *IgnoreUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *IgnoreUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IgnoreUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntentAchievement struct {
	// 意图达成方式，qa:问答回复、doc：文档回复、workflow：工作流回复，llm：大模型回复
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 意图达成方式描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`
}

type InvokeAPI struct {
	// 请求方法，如GET/POST等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 请求地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// header参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderValues []*StrValue `json:"HeaderValues,omitnil,omitempty" name:"HeaderValues"`

	// 入参Query
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryValues []*StrValue `json:"QueryValues,omitnil,omitempty" name:"QueryValues"`

	// Post请求的原始数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestPostBody *string `json:"RequestPostBody,omitnil,omitempty" name:"RequestPostBody"`

	// 返回的原始数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseBody *string `json:"ResponseBody,omitnil,omitempty" name:"ResponseBody"`

	// 出参
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResponseValues []*ValueInfo `json:"ResponseValues,omitnil,omitempty" name:"ResponseValues"`

	// 异常信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`
}

// Predefined struct for user
type IsTransferIntentRequestParams struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 应用appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`
}

type IsTransferIntentRequest struct {
	*tchttp.BaseRequest
	
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 应用appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`
}

func (r *IsTransferIntentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsTransferIntentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Content")
	delete(f, "BotAppKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "IsTransferIntentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type IsTransferIntentResponseParams struct {
	// 是否意图转人工
	Hit *bool `json:"Hit,omitnil,omitempty" name:"Hit"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type IsTransferIntentResponse struct {
	*tchttp.BaseResponse
	Response *IsTransferIntentResponseParams `json:"Response"`
}

func (r *IsTransferIntentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *IsTransferIntentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KnowledgeBaseInfo struct {
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 共享知识库名称
	KnowledgeName *string `json:"KnowledgeName,omitnil,omitempty" name:"KnowledgeName"`

	// 共享知识库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	KnowledgeDescription *string `json:"KnowledgeDescription,omitnil,omitempty" name:"KnowledgeDescription"`

	// Embedding模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmbeddingModel *string `json:"EmbeddingModel,omitnil,omitempty" name:"EmbeddingModel"`

	// 问答提取模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaExtractModel *string `json:"QaExtractModel,omitnil,omitempty" name:"QaExtractModel"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type KnowledgeCapacityPieGraphDetail struct {
	// 当前应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 当前应用使用的字符数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedCharSize *string `json:"UsedCharSize,omitnil,omitempty" name:"UsedCharSize"`

	// 当前应用对于总用量的占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proportion *float64 `json:"Proportion,omitnil,omitempty" name:"Proportion"`

	// 知识库类型:0默认1共享
	KnowledgeType *int64 `json:"KnowledgeType,omitnil,omitempty" name:"KnowledgeType"`
}

type KnowledgeDetail struct {
	// 应用名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// 已用字符数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedCharSize *string `json:"UsedCharSize,omitnil,omitempty" name:"UsedCharSize"`

	// 使用占比
	// 注意：此字段可能返回 null，表示取不到有效值。
	Proportion *float64 `json:"Proportion,omitnil,omitempty" name:"Proportion"`

	// 超量字符数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExceedCharSize *string `json:"ExceedCharSize,omitnil,omitempty" name:"ExceedCharSize"`

	// 废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSharedKnowledge *bool `json:"IsSharedKnowledge,omitnil,omitempty" name:"IsSharedKnowledge"`

	// 知识库类型:0默认1共享
	KnowledgeType *int64 `json:"KnowledgeType,omitnil,omitempty" name:"KnowledgeType"`
}

type KnowledgeDetailInfo struct {
	// 知识库信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Knowledge *KnowledgeBaseInfo `json:"Knowledge,omitnil,omitempty" name:"Knowledge"`

	// 应用列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppList []*AppBaseInfo `json:"AppList,omitnil,omitempty" name:"AppList"`

	// 用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	User *UserBaseInfo `json:"User,omitnil,omitempty" name:"User"`
}

type KnowledgeQaConfig struct {
	// 欢迎语，200字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`

	// 角色描述，4000字符以内。通过填写描述，设定应用的 #角色名称、 #风格特点 及可达成的#意图。建议按照下面的模板填写，且自定义意图建议不超过5个。
	// 
	// #角色名称：
	// #风格特点：
	// #输出要求：
	// #能力限制：
	// 
	// 能够达成以下用户意图
	// ##意图名称：
	// ##意图描述：
	// ##意图示例：
	// ##意图实现：
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoleDescription *string `json:"RoleDescription,omitnil,omitempty" name:"RoleDescription"`

	// 生成模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// 知识搜索配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Search []*KnowledgeQaSearch `json:"Search,omitnil,omitempty" name:"Search"`

	// 知识管理输出配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *KnowledgeQaOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 工作流程配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workflow *KnowledgeWorkflow `json:"Workflow,omitnil,omitempty" name:"Workflow"`

	// 检索范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchRange *SearchRange `json:"SearchRange,omitnil,omitempty" name:"SearchRange"`

	// 应用模式，standard:标准模式, agent: agent模式，single_workflow：单工作流模式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 检索策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchStrategy *SearchStrategy `json:"SearchStrategy,omitnil,omitempty" name:"SearchStrategy"`

	// 单工作流ID，Pattern为single_workflow时传入
	// 注意：此字段可能返回 null，表示取不到有效值。
	SingleWorkflow *KnowledgeQaSingleWorkflow `json:"SingleWorkflow,omitnil,omitempty" name:"SingleWorkflow"`

	// 应用关联插件
	// 注意：此字段可能返回 null，表示取不到有效值。
	Plugins []*KnowledgeQaPlugin `json:"Plugins,omitnil,omitempty" name:"Plugins"`

	// 思考模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThoughtModel *AppModel `json:"ThoughtModel,omitnil,omitempty" name:"ThoughtModel"`

	// 意图达成方式优先级
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentAchievements []*IntentAchievement `json:"IntentAchievements,omitnil,omitempty" name:"IntentAchievements"`

	// 是否开启图文检索
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTextRetrieval *bool `json:"ImageTextRetrieval,omitnil,omitempty" name:"ImageTextRetrieval"`

	// 配置语音通话参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiCall *AICallConfig `json:"AiCall,omitnil,omitempty" name:"AiCall"`

	// 共享知识库关联配置
	ShareKnowledgeBases []*ShareKnowledgeBase `json:"ShareKnowledgeBases,omitnil,omitempty" name:"ShareKnowledgeBases"`

	// 背景图相关信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackgroundImage *BackgroundImageConfig `json:"BackgroundImage,omitnil,omitempty" name:"BackgroundImage"`

	// 开场问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpeningQuestions []*string `json:"OpeningQuestions,omitnil,omitempty" name:"OpeningQuestions"`
}

type KnowledgeQaOutput struct {
	// 输出方式 1：流式 2：非流式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 通用模型回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseGeneralKnowledge *bool `json:"UseGeneralKnowledge,omitnil,omitempty" name:"UseGeneralKnowledge"`

	// 未知回复语，300字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	BareAnswer *string `json:"BareAnswer,omitnil,omitempty" name:"BareAnswer"`

	// 是否展示问题澄清开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowQuestionClarify *bool `json:"ShowQuestionClarify,omitnil,omitempty" name:"ShowQuestionClarify"`

	// 是否打开问题澄清
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseQuestionClarify *bool `json:"UseQuestionClarify,omitnil,omitempty" name:"UseQuestionClarify"`

	// 问题澄清关键词列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuestionClarifyKeywords []*string `json:"QuestionClarifyKeywords,omitnil,omitempty" name:"QuestionClarifyKeywords"`

	// 是否打开推荐问题开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseRecommended *bool `json:"UseRecommended,omitnil,omitempty" name:"UseRecommended"`
}

type KnowledgeQaPlugin struct {
	// 插件ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// 插件名称
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// 插件图标
	PluginIcon *string `json:"PluginIcon,omitnil,omitempty" name:"PluginIcon"`

	// 工具ID
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// 工具名称
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// 工具描述
	ToolDesc *string `json:"ToolDesc,omitnil,omitempty" name:"ToolDesc"`

	// 工具输入参数
	Inputs []*PluginToolReqParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// 插件是否和知识库绑定
	IsBindingKnowledge *bool `json:"IsBindingKnowledge,omitnil,omitempty" name:"IsBindingKnowledge"`
}

type KnowledgeQaSearch struct {
	// 知识来源 doc：文档，qa：问答  taskflow：业务流程，search：搜索增强，database:数据库
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 问答-回复灵活度 1：已采纳答案直接回复 2：已采纳润色后回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyFlexibility *uint64 `json:"ReplyFlexibility,omitnil,omitempty" name:"ReplyFlexibility"`

	// 搜索增强-搜索引擎状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	UseSearchEngine *bool `json:"UseSearchEngine,omitnil,omitempty" name:"UseSearchEngine"`

	// 是否显示搜索引擎检索状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShowSearchEngine *bool `json:"ShowSearchEngine,omitnil,omitempty" name:"ShowSearchEngine"`

	// 知识来源，是否选择
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEnabled *bool `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 问答最大召回数量, 默认2，限制5
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaTopN *uint64 `json:"QaTopN,omitnil,omitempty" name:"QaTopN"`

	// 文档最大召回数量, 默认3，限制5
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocTopN *uint64 `json:"DocTopN,omitnil,omitempty" name:"DocTopN"`

	// 检索置信度，针对文档和问答有效，最小0.01，最大0.99
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 资源状态 1：资源可用；2：资源已用尽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type KnowledgeQaSingleWorkflow struct {
	// 工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流描述
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 工作流状态，发布状态(UNPUBLISHED: 待发布 PUBLISHING: 发布中 PUBLISHED: 已发布 FAIL:发布失败)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 工作流是否启用
	IsEnable *bool `json:"IsEnable,omitnil,omitempty" name:"IsEnable"`

	// 是否开启异步调用工作流
	AsyncWorkflow *bool `json:"AsyncWorkflow,omitnil,omitempty" name:"AsyncWorkflow"`
}

type KnowledgeSummary struct {
	// 1是问答 2是文档片段
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 知识内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`
}

type KnowledgeUpdateInfo struct {
	// 共享知识库名称
	KnowledgeName *string `json:"KnowledgeName,omitnil,omitempty" name:"KnowledgeName"`

	// 共享知识库描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	KnowledgeDescription *string `json:"KnowledgeDescription,omitnil,omitempty" name:"KnowledgeDescription"`

	// Embedding模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmbeddingModel *string `json:"EmbeddingModel,omitnil,omitempty" name:"EmbeddingModel"`

	// 问答提取模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaExtractModel *string `json:"QaExtractModel,omitnil,omitempty" name:"QaExtractModel"`
}

type KnowledgeWorkflow struct {
	// 是否启用工作流
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsEnabled *bool `json:"IsEnabled,omitnil,omitempty" name:"IsEnabled"`

	// 是否启用PDL
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsePdl *bool `json:"UsePdl,omitnil,omitempty" name:"UsePdl"`
}

type Label struct {
	// 标签ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelBizId *string `json:"LabelBizId,omitnil,omitempty" name:"LabelBizId"`

	// 标签名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	LabelName *string `json:"LabelName,omitnil,omitempty" name:"LabelName"`
}

// Predefined struct for user
type ListAppCategoryRequestParams struct {

}

type ListAppCategoryRequest struct {
	*tchttp.BaseRequest
	
}

func (r *ListAppCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAppCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAppCategoryResponseParams struct {
	// 应用类型列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*ListAppCategoryRspOption `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAppCategoryResponse struct {
	*tchttp.BaseResponse
	Response *ListAppCategoryResponseParams `json:"Response"`
}

func (r *ListAppCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListAppCategoryRspOption struct {
	// 类型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 类型值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 类型log
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logo *string `json:"Logo,omitnil,omitempty" name:"Logo"`
}

// Predefined struct for user
type ListAppKnowledgeDetailRequestParams struct {
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页面大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用ID列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

type ListAppKnowledgeDetailRequest struct {
	*tchttp.BaseRequest
	
	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 页面大小
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用ID列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`
}

func (r *ListAppKnowledgeDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppKnowledgeDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "AppBizIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAppKnowledgeDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAppKnowledgeDetailResponseParams struct {
	// 列表总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 应用使用知识库容量详情
	List []*KnowledgeDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAppKnowledgeDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListAppKnowledgeDetailResponseParams `json:"Response"`
}

func (r *ListAppKnowledgeDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppKnowledgeDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAppRequestParams struct {
	// 应用类型；knowledge_qa - 知识问答管理 
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 每页数目，整型
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，整型
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 关键词：应用/修改人
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ListAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用类型；knowledge_qa - 知识问答管理 
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 每页数目，整型
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 页码，整型
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 关键词：应用/修改人
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ListAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	delete(f, "Keyword")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAppResponseParams struct {
	// 数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 标签列表
	List []*AppInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAppResponse struct {
	*tchttp.BaseResponse
	Response *ListAppResponseParams `json:"Response"`
}

func (r *ListAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttributeLabelRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 每个属性同步拉取的标签值数量
	LabelSize *uint64 `json:"LabelSize,omitnil,omitempty" name:"LabelSize"`
}

type ListAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 每个属性同步拉取的标签值数量
	LabelSize *uint64 `json:"LabelSize,omitnil,omitempty" name:"LabelSize"`
}

func (r *ListAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	delete(f, "LabelSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListAttributeLabelResponseParams struct {
	// 总数
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	List []*AttrLabelDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ListAttributeLabelResponseParams `json:"Response"`
}

func (r *ListAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDocCateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type ListDocCateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *ListDocCateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocCateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDocCateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDocCateResponseParams struct {
	// 列表
	List []*CateInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDocCateResponse struct {
	*tchttp.BaseResponse
	Response *ListDocCateResponseParams `json:"Response"`
}

func (r *ListDocCateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocCateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListDocItem struct {
	// 文档ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 文件名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 重命名的新文档名称，在重命名提交之后，文档发布之前都是这个名称
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// cos路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 文档状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 文档状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 答案中是否引用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 问答对数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	QaNum *int64 `json:"QaNum,omitnil,omitempty" name:"QaNum"`

	// 是否已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDeleted *bool `json:"IsDeleted,omitnil,omitempty" name:"IsDeleted"`

	// 文档来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 文档来源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 是否允许重新生成
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowRestart *bool `json:"IsAllowRestart,omitnil,omitempty" name:"IsAllowRestart"`

	// qa是否已删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDeletedQa *bool `json:"IsDeletedQa,omitnil,omitempty" name:"IsDeletedQa"`

	// 问答是否生成中
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCreatingQa *bool `json:"IsCreatingQa,omitnil,omitempty" name:"IsCreatingQa"`

	// 是否允许删除
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许操作引用开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowRefer *bool `json:"IsAllowRefer,omitnil,omitempty" name:"IsAllowRefer"`

	// 问答是否生成过
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsCreatedQa *bool `json:"IsCreatedQa,omitnil,omitempty" name:"IsCreatedQa"`

	// 文档字符量
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocCharSize *string `json:"DocCharSize,omitnil,omitempty" name:"DocCharSize"`

	// 属性标签适用范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 是否允许编辑
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 网页(或自定义链接)地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 有效开始时间，unix时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 是否允许重试，0：否，1：是
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowRetry *bool `json:"IsAllowRetry,omitnil,omitempty" name:"IsAllowRetry"`

	// 0:文档比对处理 1:文档生成问答
	// 注意：此字段可能返回 null，表示取不到有效值。
	Processing []*int64 `json:"Processing,omitnil,omitempty" name:"Processing"`

	// 文档创建落库时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 文档所属分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 文档的用户自定义ID
	CustomerKnowledgeId *string `json:"CustomerKnowledgeId,omitnil,omitempty" name:"CustomerKnowledgeId"`

	// 文档的属性标记，0: 不做用户外部权限校验
	AttributeFlags []*uint64 `json:"AttributeFlags,omitnil,omitempty" name:"AttributeFlags"`

	// false:未停用，ture:已停用
	IsDisabled *bool `json:"IsDisabled,omitnil,omitempty" name:"IsDisabled"`

	// 员工名称
	StaffName *string `json:"StaffName,omitnil,omitempty" name:"StaffName"`
}

// Predefined struct for user
type ListDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	// 
	// 输入特定标识 lke:system:untagged  将查询所有未关联标签的文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档状态： 1-未生成 2-生成中 3-生成成功 4-生成失败 5-删除中 6-删除成功  7-审核中  8-审核失败 9-审核成功  10-待发布  11-发布中  12-已发布  13-学习中  14-学习失败  15-更新中  16-更新失败  17-解析中  18-解析失败  19-导入失败   20-已过期 21-超量失效 22-超量失效恢复
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 查询类型 filename 文档、 attribute 标签
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 文件类型分类筛选
	FileTypes []*string `json:"FileTypes,omitnil,omitempty" name:"FileTypes"`

	// 文档列表筛选标识位
	FilterFlag []*DocFilterFlag `json:"FilterFlag,omitnil,omitempty" name:"FilterFlag"`

	// 是否只展示当前分类的数据 0不是，1是
	ShowCurrCate *uint64 `json:"ShowCurrCate,omitnil,omitempty" name:"ShowCurrCate"`
}

type ListDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	// 
	// 输入特定标识 lke:system:untagged  将查询所有未关联标签的文档
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 文档状态： 1-未生成 2-生成中 3-生成成功 4-生成失败 5-删除中 6-删除成功  7-审核中  8-审核失败 9-审核成功  10-待发布  11-发布中  12-已发布  13-学习中  14-学习失败  15-更新中  16-更新失败  17-解析中  18-解析失败  19-导入失败   20-已过期 21-超量失效 22-超量失效恢复
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 查询类型 filename 文档、 attribute 标签
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 文件类型分类筛选
	FileTypes []*string `json:"FileTypes,omitnil,omitempty" name:"FileTypes"`

	// 文档列表筛选标识位
	FilterFlag []*DocFilterFlag `json:"FilterFlag,omitnil,omitempty" name:"FilterFlag"`

	// 是否只展示当前分类的数据 0不是，1是
	ShowCurrCate *uint64 `json:"ShowCurrCate,omitnil,omitempty" name:"ShowCurrCate"`
}

func (r *ListDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "Status")
	delete(f, "QueryType")
	delete(f, "CateBizId")
	delete(f, "FileTypes")
	delete(f, "FilterFlag")
	delete(f, "ShowCurrCate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListDocResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ListDocItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListDocResponse struct {
	*tchttp.BaseResponse
	Response *ListDocResponseParams `json:"Response"`
}

func (r *ListDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListModelRequestParams struct {
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用模式 standard:标准模式, agent: agent模式，single_workflow：单工作流模式
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 模型类别 generate：生成模型，thought：思考模型
	ModelCategory *string `json:"ModelCategory,omitnil,omitempty" name:"ModelCategory"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ListModelRequest struct {
	*tchttp.BaseRequest
	
	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classifys-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用模式 standard:标准模式, agent: agent模式，single_workflow：单工作流模式
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// 模型类别 generate：生成模型，thought：思考模型
	ModelCategory *string `json:"ModelCategory,omitnil,omitempty" name:"ModelCategory"`

	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ListModelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListModelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppType")
	delete(f, "Pattern")
	delete(f, "ModelCategory")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListModelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListModelResponseParams struct {
	// 模型列表
	List []*ModelInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListModelResponse struct {
	*tchttp.BaseResponse
	Response *ListModelResponseParams `json:"Response"`
}

func (r *ListModelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListModelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQACateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

type ListQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`
}

func (r *ListQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQACateResponseParams struct {
	// 列表
	List []*QACate `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListQACateResponse struct {
	*tchttp.BaseResponse
	Response *ListQACateResponseParams `json:"Response"`
}

func (r *ListQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQARequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询问题
	// 
	// 输入特定标识 lke:system:untagged  将查询所有未关联标签的问答
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 校验状态(1未校验2采纳3不采纳)
	AcceptStatus []*int64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态(2待发布 3发布中 4已发布 7审核中 8审核失败 9人工申述中 11人工申述失败 12已过期 13超量失效 14超量失效恢复)
	ReleaseStatus []*int64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 来源(1 文档生成 2 批量导入 3 手动添加)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 查询答案
	QueryAnswer *string `json:"QueryAnswer,omitnil,omitempty" name:"QueryAnswer"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// QA业务ID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 查询类型 filename 名称、 attribute 标签
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 是否只展示当前分类的数据 0不是，1是
	ShowCurrCate *uint64 `json:"ShowCurrCate,omitnil,omitempty" name:"ShowCurrCate"`
}

type ListQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询问题
	// 
	// 输入特定标识 lke:system:untagged  将查询所有未关联标签的问答
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 校验状态(1未校验2采纳3不采纳)
	AcceptStatus []*int64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态(2待发布 3发布中 4已发布 7审核中 8审核失败 9人工申述中 11人工申述失败 12已过期 13超量失效 14超量失效恢复)
	ReleaseStatus []*int64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 来源(1 文档生成 2 批量导入 3 手动添加)
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 查询答案
	QueryAnswer *string `json:"QueryAnswer,omitnil,omitempty" name:"QueryAnswer"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// QA业务ID列表
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 查询类型 filename 名称、 attribute 标签
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 是否只展示当前分类的数据 0不是，1是
	ShowCurrCate *uint64 `json:"ShowCurrCate,omitnil,omitempty" name:"ShowCurrCate"`
}

func (r *ListQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "AcceptStatus")
	delete(f, "ReleaseStatus")
	delete(f, "DocBizId")
	delete(f, "Source")
	delete(f, "QueryAnswer")
	delete(f, "CateBizId")
	delete(f, "QaBizIds")
	delete(f, "QueryType")
	delete(f, "ShowCurrCate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListQAResponseParams struct {
	// 问答数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 待校验问答数量
	WaitVerifyTotal *string `json:"WaitVerifyTotal,omitnil,omitempty" name:"WaitVerifyTotal"`

	// 未采纳问答数量
	NotAcceptedTotal *string `json:"NotAcceptedTotal,omitnil,omitempty" name:"NotAcceptedTotal"`

	// 已采纳问答数量
	AcceptedTotal *string `json:"AcceptedTotal,omitnil,omitempty" name:"AcceptedTotal"`

	// 页码
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 问答详情
	List []*ListQaItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListQAResponse struct {
	*tchttp.BaseResponse
	Response *ListQAResponseParams `json:"Response"`
}

func (r *ListQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListQaItem struct {
	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 来源
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否允许编辑
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 是否允许删除
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`

	// 是否允许校验
	IsAllowAccept *bool `json:"IsAllowAccept,omitnil,omitempty" name:"IsAllowAccept"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 问答字符数
	QaCharSize *string `json:"QaCharSize,omitnil,omitempty" name:"QaCharSize"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *int64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签
	AttrLabels []*AttrLabel `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 相似问个数
	SimilarQuestionNum *uint64 `json:"SimilarQuestionNum,omitnil,omitempty" name:"SimilarQuestionNum"`

	// 返回问答关联的相似问,联动搜索,仅展示一条
	SimilarQuestionTips *string `json:"SimilarQuestionTips,omitnil,omitempty" name:"SimilarQuestionTips"`

	// 问答是否停用，false:未停用，ture:已停用
	IsDisabled *bool `json:"IsDisabled,omitnil,omitempty" name:"IsDisabled"`

	// 员工名称
	StaffName *string `json:"StaffName,omitnil,omitempty" name:"StaffName"`
}

// Predefined struct for user
type ListReferShareKnowledgeRequestParams struct {
	// 应用业务id
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ListReferShareKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 应用业务id
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ListReferShareKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReferShareKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReferShareKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReferShareKnowledgeResponseParams struct {
	// 共享知识库信息列表
	List []*KnowledgeBaseInfo `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReferShareKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *ListReferShareKnowledgeResponseParams `json:"Response"`
}

func (r *ListReferShareKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReferShareKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionPreviewRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type ListRejectedQuestionPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *ListRejectedQuestionPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "Actions")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRejectedQuestionPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionPreviewResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ReleaseRejectedQuestion `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRejectedQuestionPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListRejectedQuestionPreviewResponseParams `json:"Response"`
}

func (r *ListRejectedQuestionPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type ListRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *ListRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListRejectedQuestionResponseParams struct {
	// 总数
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 拒答问题列表
	List []*RejectedQuestion `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *ListRejectedQuestionResponseParams `json:"Response"`
}

func (r *ListRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseConfigPreviewRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 发布状态
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

type ListReleaseConfigPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 状态(1新增2更新3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 发布状态
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

func (r *ListReleaseConfigPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseConfigPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "Actions")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "ReleaseStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseConfigPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseConfigPreviewResponseParams struct {
	// 数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 配置项列表
	List []*ReleaseConfigs `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseConfigPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseConfigPreviewResponseParams `json:"Response"`
}

func (r *ListReleaseConfigPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseConfigPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseDocPreviewRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`
}

type ListReleaseDocPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`
}

func (r *ListReleaseDocPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseDocPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Actions")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseDocPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseDocPreviewResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ReleaseDoc `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseDocPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseDocPreviewResponseParams `json:"Response"`
}

func (r *ListReleaseDocPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseDocPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListReleaseItem struct {
	// 版本ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 发布人
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 发布描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 发布状态
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 发布状态描述
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 发布成功数
	SuccessCount *int64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 发布失败数
	FailCount *int64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`
}

// Predefined struct for user
type ListReleaseQAPreviewRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 发布状态(4发布成功5发布失败)
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

type ListReleaseQAPreviewRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 发布单ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 状态(1新增2修改3删除)
	Actions []*uint64 `json:"Actions,omitnil,omitempty" name:"Actions"`

	// 发布状态(4发布成功5发布失败)
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`
}

func (r *ListReleaseQAPreviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseQAPreviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "ReleaseBizId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Actions")
	delete(f, "ReleaseStatus")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseQAPreviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseQAPreviewResponseParams struct {
	// 文档数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 文档列表
	List []*ReleaseQA `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseQAPreviewResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseQAPreviewResponseParams `json:"Response"`
}

func (r *ListReleaseQAPreviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseQAPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type ListReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *ListReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListReleaseResponseParams struct {
	// 发布列表数量
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 发布列表
	List []*ListReleaseItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListReleaseResponse struct {
	*tchttp.BaseResponse
	Response *ListReleaseResponseParams `json:"Response"`
}

func (r *ListReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSelectDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档状态： 7 审核中、8 审核失败、10 待发布、11 发布中、12 已发布、13 学习中、14 学习失败 20 已过期
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type ListSelectDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档状态： 7 审核中、8 审核失败、10 待发布、11 发布中、12 已发布、13 学习中、14 学习失败 20 已过期
	Status []*int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ListSelectDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSelectDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSelectDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSelectDocResponseParams struct {
	// 下拉框内容
	List []*Option `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSelectDocResponse struct {
	*tchttp.BaseResponse
	Response *ListSelectDocResponseParams `json:"Response"`
}

func (r *ListSelectDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSelectDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSharedKnowledgeRequestParams struct {
	// 分页序号，编码从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，有效范围为[1,200]
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type ListSharedKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 分页序号，编码从1开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页大小，有效范围为[1,200]
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 搜索关键字
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *ListSharedKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSharedKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListSharedKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListSharedKnowledgeResponseParams struct {
	// 累计数量
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 知识库列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	KnowledgeList []*KnowledgeDetailInfo `json:"KnowledgeList,omitnil,omitempty" name:"KnowledgeList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListSharedKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *ListSharedKnowledgeResponseParams `json:"Response"`
}

func (r *ListSharedKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListSharedKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUnsatisfiedReplyRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 用户请求(问题或答案)
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 错误类型检索
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

type ListUnsatisfiedReplyRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 页码
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 用户请求(问题或答案)
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 错误类型检索
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

func (r *ListUnsatisfiedReplyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUnsatisfiedReplyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "Query")
	delete(f, "Reasons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUnsatisfiedReplyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUnsatisfiedReplyResponseParams struct {
	// 总数
	Total *string `json:"Total,omitnil,omitempty" name:"Total"`

	// 不满意回复列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	List []*UnsatisfiedReply `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUnsatisfiedReplyResponse struct {
	*tchttp.BaseResponse
	Response *ListUnsatisfiedReplyResponseParams `json:"Response"`
}

func (r *ListUnsatisfiedReplyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUnsatisfiedReplyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsageCallDetailRequestParams struct {
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码（从1开始）
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数量(最大值1000)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// uin列表
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 应用ID列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 调用类型列表
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// 筛选子场景
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 账单明细对应的自定义tag
	BillingTag *string `json:"BillingTag,omitnil,omitempty" name:"BillingTag"`
}

type ListUsageCallDetailRequest struct {
	*tchttp.BaseRequest
	
	// 模型标识
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 开始时间
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 页码（从1开始）
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 分页数量(最大值1000)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// uin列表
	UinAccount []*string `json:"UinAccount,omitnil,omitempty" name:"UinAccount"`

	// 应用ID列表
	AppBizIds []*string `json:"AppBizIds,omitnil,omitempty" name:"AppBizIds"`

	// 调用类型列表
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// 筛选子场景
	SubScenes []*string `json:"SubScenes,omitnil,omitempty" name:"SubScenes"`

	// 应用类型(knowledge_qa应用管理， shared_knowlege 共享知识库)
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 账单明细对应的自定义tag
	BillingTag *string `json:"BillingTag,omitnil,omitempty" name:"BillingTag"`
}

func (r *ListUsageCallDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsageCallDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelName")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "UinAccount")
	delete(f, "AppBizIds")
	delete(f, "CallType")
	delete(f, "SubScenes")
	delete(f, "AppType")
	delete(f, "BillingTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListUsageCallDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListUsageCallDetailResponseParams struct {
	// 列表总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 列表
	List []*CallDetail `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListUsageCallDetailResponse struct {
	*tchttp.BaseResponse
	Response *ListUsageCallDetailResponseParams `json:"Response"`
}

func (r *ListUsageCallDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListUsageCallDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowRunsRequestParams struct {
	// 运行环境。0: 测试环境； 1: 正式环境
	RunEnv *uint64 `json:"RunEnv,omitnil,omitempty" name:"RunEnv"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 页码
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ListWorkflowRunsRequest struct {
	*tchttp.BaseRequest
	
	// 运行环境。0: 测试环境； 1: 正式环境
	RunEnv *uint64 `json:"RunEnv,omitnil,omitempty" name:"RunEnv"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 页码
	Page *uint64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ListWorkflowRunsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowRunsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RunEnv")
	delete(f, "AppBizId")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ListWorkflowRunsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ListWorkflowRunsResponseParams struct {
	// 总数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 工作流运行列表
	WorkflowRuns []*WorkflowRunBase `json:"WorkflowRuns,omitnil,omitempty" name:"WorkflowRuns"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ListWorkflowRunsResponse struct {
	*tchttp.BaseResponse
	Response *ListWorkflowRunsResponseParams `json:"Response"`
}

func (r *ListWorkflowRunsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ListWorkflowRunsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModelInfo struct {
	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 模型描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelDesc *string `json:"ModelDesc,omitnil,omitempty" name:"ModelDesc"`

	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// 资源状态 1：资源可用；2：资源已用尽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`

	// 提示词内容字符限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	PromptWordsLimit *string `json:"PromptWordsLimit,omitnil,omitempty" name:"PromptWordsLimit"`

	// 通过核心采样控制内容生成的多样性，较高的Top P值会导致生成更多样的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopP *ModelParameter `json:"TopP,omitnil,omitempty" name:"TopP"`

	// 温度控制随机性
	// 注意：此字段可能返回 null，表示取不到有效值。
	Temperature *ModelParameter `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// 最多能生成的token数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTokens *ModelParameter `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`

	// 模型来源 Hunyuan：腾讯混元大模型,Industry：腾讯云行业大模型,Experience：新模型体验,Custom自定义模型
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 模型图标
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// 是否免费
	IsFree *bool `json:"IsFree,omitnil,omitempty" name:"IsFree"`

	// 模型对话框可输入的上限
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputLenLimit *uint64 `json:"InputLenLimit,omitnil,omitempty" name:"InputLenLimit"`

	// 支持工作流的类型 0:模型不支持; 1: 模型支持工作流； 2： 模型支持效果不佳；
	// 注意：此字段可能返回 null，表示取不到有效值。
	SupportWorkflowStatus *uint64 `json:"SupportWorkflowStatus,omitnil,omitempty" name:"SupportWorkflowStatus"`

	// 模型类别 generate：生成模型，thought：思考模型
	ModelCategory *string `json:"ModelCategory,omitnil,omitempty" name:"ModelCategory"`

	// 是否默认模型
	IsDefault *bool `json:"IsDefault,omitnil,omitempty" name:"IsDefault"`

	// 角色提示词输入长度限制
	RoleLenLimit *uint64 `json:"RoleLenLimit,omitnil,omitempty" name:"RoleLenLimit"`

	// 是否专属并发模型
	IsExclusive *bool `json:"IsExclusive,omitnil,omitempty" name:"IsExclusive"`

	// 模型支持智能通话效果
	SupportAiCallStatus *uint64 `json:"SupportAiCallStatus,omitnil,omitempty" name:"SupportAiCallStatus"`

	// 专属并发数
	Concurrency *uint64 `json:"Concurrency,omitnil,omitempty" name:"Concurrency"`
}

type ModelParameter struct {
	// 默认值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Default *float64 `json:"Default,omitnil,omitempty" name:"Default"`

	// 最小值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Min *float64 `json:"Min,omitnil,omitempty" name:"Min"`

	// 最大值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Max *float64 `json:"Max,omitnil,omitempty" name:"Max"`
}

// Predefined struct for user
type ModifyAgentRequestParams struct {
	// 需要修改的应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 修改后的Agent的信息
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

type ModifyAgentRequest struct {
	*tchttp.BaseRequest
	
	// 需要修改的应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 修改后的Agent的信息
	Agent *Agent `json:"Agent,omitnil,omitempty" name:"Agent"`
}

func (r *ModifyAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "Agent")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentResponseParams struct {
	// 修改的AgentId
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAgentResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAgentResponseParams `json:"Response"`
}

func (r *ModifyAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppRequestParams struct {
	// 应用 ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classify-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 应用类型；knowledge_qa-知识问答管理；summary-知识摘要；classify-知识标签提取
	AppType *string `json:"AppType,omitnil,omitempty" name:"AppType"`

	// 应用基础配置
	BaseConfig *BaseConfig `json:"BaseConfig,omitnil,omitempty" name:"BaseConfig"`

	// 应用配置
	AppConfig *AppConfig `json:"AppConfig,omitnil,omitempty" name:"AppConfig"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ModifyAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "AppType")
	delete(f, "BaseConfig")
	delete(f, "AppConfig")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// 应用App
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAppResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppResponseParams `json:"Response"`
}

func (r *ModifyAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttributeLabelRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 标签ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 标签名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签标识 （已作废）
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 删除的标签值
	DeleteLabelBizIds []*string `json:"DeleteLabelBizIds,omitnil,omitempty" name:"DeleteLabelBizIds"`

	// 新增或编辑的标签
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

type ModifyAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 标签ID
	AttributeBizId *string `json:"AttributeBizId,omitnil,omitempty" name:"AttributeBizId"`

	// 标签名称
	AttrName *string `json:"AttrName,omitnil,omitempty" name:"AttrName"`

	// 标签标识 （已作废）
	AttrKey *string `json:"AttrKey,omitnil,omitempty" name:"AttrKey"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 删除的标签值
	DeleteLabelBizIds []*string `json:"DeleteLabelBizIds,omitnil,omitempty" name:"DeleteLabelBizIds"`

	// 新增或编辑的标签
	Labels []*AttributeLabel `json:"Labels,omitnil,omitempty" name:"Labels"`
}

func (r *ModifyAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "AttributeBizId")
	delete(f, "AttrName")
	delete(f, "AttrKey")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "DeleteLabelBizIds")
	delete(f, "Labels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAttributeLabelResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAttributeLabelResponseParams `json:"Response"`
}

func (r *ModifyAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocAttrRangeRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

type ModifyDocAttrRangeRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizIds []*string `json:"DocBizIds,omitnil,omitempty" name:"DocBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

func (r *ModifyDocAttrRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocAttrRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizIds")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocAttrRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocAttrRangeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocAttrRangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocAttrRangeResponseParams `json:"Response"`
}

func (r *ModifyDocAttrRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocAttrRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocCateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type ModifyDocCateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *ModifyDocCateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocCateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Name")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocCateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocCateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocCateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocCateResponseParams `json:"Response"`
}

func (r *ModifyDocCateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocCateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 标签适用范围，需要传参为1
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 关联的标签
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 网页(或自定义链接)地址
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 是否可下载，IsRefer为true并且ReferUrlType为0时，该值才有意义
	IsDownload *bool `json:"IsDownload,omitnil,omitempty" name:"IsDownload"`
}

type ModifyDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 标签适用范围，需要传参为1
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 关联的标签
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 网页(或自定义链接)地址
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 是否可下载，IsRefer为true并且ReferUrlType为0时，该值才有意义
	IsDownload *bool `json:"IsDownload,omitnil,omitempty" name:"IsDownload"`
}

func (r *ModifyDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	delete(f, "IsRefer")
	delete(f, "AttrRange")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "AttrLabels")
	delete(f, "WebUrl")
	delete(f, "ReferUrlType")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "CateBizId")
	delete(f, "IsDownload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDocResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDocResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDocResponseParams `json:"Response"`
}

func (r *ModifyDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAAttrRangeRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

type ModifyQAAttrRangeRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizIds []*string `json:"QaBizIds,omitnil,omitempty" name:"QaBizIds"`

	// 属性标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 属性标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`
}

func (r *ModifyQAAttrRangeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAAttrRangeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizIds")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQAAttrRangeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAAttrRangeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQAAttrRangeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQAAttrRangeResponseParams `json:"Response"`
}

func (r *ModifyQAAttrRangeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAAttrRangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQACateRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

type ModifyQACateRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 分类名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类业务ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`
}

func (r *ModifyQACateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQACateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Name")
	delete(f, "CateBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQACateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQACateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQACateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQACateResponseParams `json:"Response"`
}

func (r *ModifyQACateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQACateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQARequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 相似问修改信息(相似问没有修改则不传)
	SimilarQuestionModify *SimilarQuestionModify `json:"SimilarQuestionModify,omitnil,omitempty" name:"SimilarQuestionModify"`

	// 问题描述
	QuestionDesc *string `json:"QuestionDesc,omitnil,omitempty" name:"QuestionDesc"`
}

type ModifyQARequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 自定义参数
	CustomParam *string `json:"CustomParam,omitnil,omitempty" name:"CustomParam"`

	// 标签适用范围 1：全部，2：按条件
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 相似问修改信息(相似问没有修改则不传)
	SimilarQuestionModify *SimilarQuestionModify `json:"SimilarQuestionModify,omitnil,omitempty" name:"SimilarQuestionModify"`

	// 问题描述
	QuestionDesc *string `json:"QuestionDesc,omitnil,omitempty" name:"QuestionDesc"`
}

func (r *ModifyQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "QaBizId")
	delete(f, "Question")
	delete(f, "Answer")
	delete(f, "CustomParam")
	delete(f, "AttrRange")
	delete(f, "AttrLabels")
	delete(f, "DocBizId")
	delete(f, "CateBizId")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "SimilarQuestionModify")
	delete(f, "QuestionDesc")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyQAResponse struct {
	*tchttp.BaseResponse
	Response *ModifyQAResponseParams `json:"Response"`
}

func (r *ModifyQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRejectedQuestionRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`
}

type ModifyRejectedQuestionRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 拒答问题
	// 
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 拒答问题来源的数据源唯一id
	// 
	// 
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`
}

func (r *ModifyRejectedQuestionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRejectedQuestionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "Question")
	delete(f, "RejectedBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRejectedQuestionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRejectedQuestionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyRejectedQuestionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRejectedQuestionResponseParams `json:"Response"`
}

func (r *ModifyRejectedQuestionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRejectedQuestionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgFileInfo struct {
	// 文档名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档大小
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *string `json:"FileSize,omitnil,omitempty" name:"FileSize"`

	// 文档URL
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// 文档类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 文档ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`
}

type MsgRecord struct {
	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 当前记录所对应的 Session ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 关联记录ID
	RelatedRecordId *string `json:"RelatedRecordId,omitnil,omitempty" name:"RelatedRecordId"`

	// 是否来自自己
	IsFromSelf *bool `json:"IsFromSelf,omitnil,omitempty" name:"IsFromSelf"`

	// 发送者名称
	FromName *string `json:"FromName,omitnil,omitempty" name:"FromName"`

	// 发送者头像
	FromAvatar *string `json:"FromAvatar,omitnil,omitempty" name:"FromAvatar"`

	// 时间戳
	Timestamp *string `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 是否已读
	HasRead *bool `json:"HasRead,omitnil,omitempty" name:"HasRead"`

	// 评价
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 是否评分
	CanRating *bool `json:"CanRating,omitnil,omitempty" name:"CanRating"`

	// 是否展示反馈按钮
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanFeedback *bool `json:"CanFeedback,omitnil,omitempty" name:"CanFeedback"`

	// 记录类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 引用来源
	References []*MsgRecordReference `json:"References,omitnil,omitempty" name:"References"`

	// 评价原因
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`

	// 是否大模型
	IsLlmGenerated *bool `json:"IsLlmGenerated,omitnil,omitempty" name:"IsLlmGenerated"`

	// 图片链接，可公有读
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrls []*string `json:"ImageUrls,omitnil,omitempty" name:"ImageUrls"`

	// 当次 token 统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenStat *TokenStat `json:"TokenStat,omitnil,omitempty" name:"TokenStat"`

	// 回复方式
	// 1:大模型直接回复;
	// 2:保守回复, 未知问题回复;
	// 3:拒答问题回复;
	// 4:敏感回复;
	// 5:问答对直接回复, 已采纳问答对优先回复;
	// 6:欢迎语回复;
	// 7:并发超限回复;
	// 8:全局干预知识;
	// 9:任务流程过程回复, 当历史记录中 task_flow.type = 0 时, 为大模型回复;
	// 10:任务流程答案回复;
	// 11:搜索引擎回复;
	// 12:知识润色后回复;
	// 13:图片理解回复;
	// 14:实时文档回复;
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyMethod *uint64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`

	// 选项卡, 用于多轮对话
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// 任务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlow *TaskFlowInfo `json:"TaskFlow,omitnil,omitempty" name:"TaskFlow"`

	// 用户传入的文件信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileInfos []*FileInfo `json:"FileInfos,omitnil,omitempty" name:"FileInfos"`

	// 参考来源引用位置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	QuoteInfos []*QuoteInfo `json:"QuoteInfos,omitnil,omitempty" name:"QuoteInfos"`

	// Agent的思考过程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentThought *AgentThought `json:"AgentThought,omitnil,omitempty" name:"AgentThought"`

	// 扩展信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *ExtraInfo `json:"ExtraInfo,omitnil,omitempty" name:"ExtraInfo"`

	// 工作流信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkFlow *WorkflowInfo `json:"WorkFlow,omitnil,omitempty" name:"WorkFlow"`
}

type MsgRecordReference struct {
	// id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 链接
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`

	// 类型
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 来源文档ID
	DocId *string `json:"DocId,omitnil,omitempty" name:"DocId"`

	// 知识库名称
	KnowledgeName *string `json:"KnowledgeName,omitnil,omitempty" name:"KnowledgeName"`

	// 知识库业务id
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 文档业务id
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 问答业务id
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`
}

type NodeRunBase struct {
	// 节点运行的ID
	NodeRunId *string `json:"NodeRunId,omitnil,omitempty" name:"NodeRunId"`

	// 节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 工作流运行实例的ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点类型。
	// 1： 开始节点
	// 2：参数提取节点
	// 3：大模型节点
	// 4：知识问答节点
	// 5：知识检索节点
	// 6：标签提取节点
	// 7：代码执行节点
	// 8：工具节点
	// 9：逻辑判断节点
	// 10：回复节点
	// 11：选项卡节点
	// 12：循环节点
	// 13：意图识别节点
	// 14：工作流节点
	// 15：插件节点
	// 16：结束节点
	// 17: 变量聚合节点数据
	// 18: 批处理节点
	// 19: 消息队列节点
	NodeType *uint64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 运行状态。0: 初始状态；1: 运行中；2: 运行成功； 3: 运行失败； 4: 已取消
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 错误码
	FailCode *string `json:"FailCode,omitnil,omitempty" name:"FailCode"`

	// 错误信息
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// 消耗时间（毫秒）
	CostMilliseconds *uint64 `json:"CostMilliseconds,omitnil,omitempty" name:"CostMilliseconds"`

	// 消耗的token总数
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

type NodeRunDetail struct {
	// 节点运行的ID
	NodeRunId *string `json:"NodeRunId,omitnil,omitempty" name:"NodeRunId"`

	// 节点ID
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 工作流运行实例的ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 节点名称
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 节点类型。
	// 1： 开始节点
	// 2：参数提取节点
	// 3：大模型节点
	// 4：知识问答节点
	// 5：知识检索节点
	// 6：标签提取节点
	// 7：代码执行节点
	// 8：工具节点
	// 9：逻辑判断节点
	// 10：回复节点
	// 11：选项卡节点
	// 12：循环节点
	// 13：意图识别节点
	// 14：工作流节点
	// 15：插件节点
	// 16：结束节点
	// 17: 变量聚合节点数据
	// 18: 批处理节点
	// 19: 消息队列节点
	NodeType *uint64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 运行状态。0: 初始状态；1: 运行中；2: 运行成功； 3: 运行失败； 4: 已取消
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 错误码
	FailCode *string `json:"FailCode,omitnil,omitempty" name:"FailCode"`

	// 错误信息
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// 消耗时间（毫秒）
	CostMilliseconds *uint64 `json:"CostMilliseconds,omitnil,omitempty" name:"CostMilliseconds"`

	// 消耗的token总数
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// 输入变量信息
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 节点的输入的完整内容的链接。（当Input内容超过限制的时候此字段才有值）
	InputRef *string `json:"InputRef,omitnil,omitempty" name:"InputRef"`

	// 输出变量信息
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 节点的输出的完整内容的链接。（当Output内容超过限制的时候此字段才有值）
	OutputRef *string `json:"OutputRef,omitnil,omitempty" name:"OutputRef"`

	// 原始输出信息。部分节点才有值，如工具节点、代码节点
	TaskOutput *string `json:"TaskOutput,omitnil,omitempty" name:"TaskOutput"`

	// 任务的原始输出的完整内容的链接。（当TaskOutput内容超过限制的时候此字段才有值）
	TaskOutputRef *string `json:"TaskOutputRef,omitnil,omitempty" name:"TaskOutputRef"`

	// 节点的日志
	Log *string `json:"Log,omitnil,omitempty" name:"Log"`

	// 节点的日志的完整内容的链接志（当Log内容超过限制的时候才有值）
	LogRef *string `json:"LogRef,omitnil,omitempty" name:"LogRef"`

	// 开始时间戳（毫秒）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间戳（毫秒）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// LLM统计信息。
	StatisticInfos []*StatisticInfo `json:"StatisticInfos,omitnil,omitempty" name:"StatisticInfos"`
}

type Option struct {
	// 文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 文件字符数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CharSize *string `json:"CharSize,omitnil,omitempty" name:"CharSize"`

	// 文件类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`
}

type PluginToolReqParam struct {
	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数描述
	Desc *string `json:"Desc,omitnil,omitempty" name:"Desc"`

	// 参数类型，0:string, 1:int, 2:float，3:bool 4:object 5:array_string, 6:array_int, 7:array_float, 8:array_bool, 9:array_object, 99:null, 100:upspecified
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数是否必填
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// 参数默认值
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// 子参数,ParamType 是OBJECT 或 ARRAY<>类型有用
	SubParams []*PluginToolReqParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// 插件参数配置是否隐藏不可见，true-隐藏不可见，false-可见
	GlobalHidden *bool `json:"GlobalHidden,omitnil,omitempty" name:"GlobalHidden"`

	// OneOf类型参数
	OneOf []*PluginToolReqParam `json:"OneOf,omitnil,omitempty" name:"OneOf"`

	// AnyOf类型参数
	AnyOf []*PluginToolReqParam `json:"AnyOf,omitnil,omitempty" name:"AnyOf"`
}

type Procedure struct {
	// 执行过程英语名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 中文名, 用于展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// 状态常量: 使用中: processing, 成功: success, 失败: failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 消耗 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Count *uint64 `json:"Count,omitnil,omitempty" name:"Count"`

	// 调试信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Debugging *ProcedureDebugging `json:"Debugging,omitnil,omitempty" name:"Debugging"`

	// 计费资源状态，1：可用，2：不可用
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceStatus *uint64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type ProcedureDebugging struct {
	// 检索query
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 系统prompt
	// 注意：此字段可能返回 null，表示取不到有效值。
	System *string `json:"System,omitnil,omitempty" name:"System"`

	// 多轮历史信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Histories []*HistorySummary `json:"Histories,omitnil,omitempty" name:"Histories"`

	// 检索知识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Knowledge []*KnowledgeSummary `json:"Knowledge,omitnil,omitempty" name:"Knowledge"`

	// 任务流程
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlow *TaskFlowSummary `json:"TaskFlow,omitnil,omitempty" name:"TaskFlow"`

	// 工作流调试信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkFlow *WorkFlowSummary `json:"WorkFlow,omitnil,omitempty" name:"WorkFlow"`

	// Agent调试信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Agent *AgentDebugInfo `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 自定义参数
	CustomVariables []*string `json:"CustomVariables,omitnil,omitempty" name:"CustomVariables"`
}

type QACate struct {
	// QA分类的业务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 分类名称
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 分类下QA数量
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 是否可新增
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// 是否可编辑
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`

	// 是否可删除
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// 子分类
	// 注意：此字段可能返回 null，表示取不到有效值。
	Children []*QACate `json:"Children,omitnil,omitempty" name:"Children"`
}

type QAList struct {
	// 问答ID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 是否采纳
	IsAccepted *bool `json:"IsAccepted,omitnil,omitempty" name:"IsAccepted"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`
}

type QAQuery struct {
	// 页码
	// 
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 查询内容
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 校验状态
	AcceptStatus []*uint64 `json:"AcceptStatus,omitnil,omitempty" name:"AcceptStatus"`

	// 发布状态
	ReleaseStatus []*uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// QAID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 来源
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 查询答案
	QueryAnswer *string `json:"QueryAnswer,omitnil,omitempty" name:"QueryAnswer"`

	// 查询类型 filename 名称、 attribute 标签
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`
}

type QuoteInfo struct {
	// 参考来源位置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Position *uint64 `json:"Position,omitnil,omitempty" name:"Position"`

	// 参考来源索引顺序
	// 注意：此字段可能返回 null，表示取不到有效值。
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`
}

// Predefined struct for user
type RateMsgRecordRequestParams struct {
	// 应用appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 消息ID 【大模型回复答案的RecordID】
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 1点赞2点踩
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 原因
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

type RateMsgRecordRequest struct {
	*tchttp.BaseRequest
	
	// 应用appKey
	BotAppKey *string `json:"BotAppKey,omitnil,omitempty" name:"BotAppKey"`

	// 消息ID 【大模型回复答案的RecordID】
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 1点赞2点踩
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 原因
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

func (r *RateMsgRecordRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RateMsgRecordRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotAppKey")
	delete(f, "RecordId")
	delete(f, "Score")
	delete(f, "Reasons")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RateMsgRecordRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RateMsgRecordResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RateMsgRecordResponse struct {
	*tchttp.BaseResponse
	Response *RateMsgRecordResponseParams `json:"Response"`
}

func (r *RateMsgRecordResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RateMsgRecordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReferDetail struct {
	// 引用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReferBizId *string `json:"ReferBizId,omitnil,omitempty" name:"ReferBizId"`

	// 文档类型 (1 QA, 2 文档段)
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocType *uint64 `json:"DocType,omitnil,omitempty" name:"DocType"`

	// 文档名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocName *string `json:"DocName,omitnil,omitempty" name:"DocName"`

	// 分片内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageContent *string `json:"PageContent,omitnil,omitempty" name:"PageContent"`

	// 问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 答案
	// 注意：此字段可能返回 null，表示取不到有效值。
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 置信度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitnil,omitempty" name:"Confidence"`

	// 标记
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mark *uint64 `json:"Mark,omitnil,omitempty" name:"Mark"`

	// 分片高亮内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Highlights []*Highlight `json:"Highlights,omitnil,omitempty" name:"Highlights"`

	// 原始内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrgData *string `json:"OrgData,omitnil,omitempty" name:"OrgData"`

	// 页码信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PageInfos []*uint64 `json:"PageInfos,omitnil,omitempty" name:"PageInfos"`

	// sheet信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SheetInfos []*string `json:"SheetInfos,omitnil,omitempty" name:"SheetInfos"`

	// 文档ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 知识库ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`
}

// Predefined struct for user
type ReferShareKnowledgeRequestParams struct {
	// 应用业务id
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 共享知识库业务id列表
	KnowledgeBizId []*string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ReferShareKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 应用业务id
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 共享知识库业务id列表
	KnowledgeBizId []*string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ReferShareKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReferShareKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "KnowledgeBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReferShareKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReferShareKnowledgeResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReferShareKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *ReferShareKnowledgeResponseParams `json:"Response"`
}

func (r *ReferShareKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReferShareKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RejectedQuestion struct {
	// 拒答问题ID
	// 
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RejectedBizId *string `json:"RejectedBizId,omitnil,omitempty" name:"RejectedBizId"`

	// 被拒答的问题
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// 更新时间
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 是否允许编辑
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowEdit *bool `json:"IsAllowEdit,omitnil,omitempty" name:"IsAllowEdit"`

	// 是否允许删除
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsAllowDelete *bool `json:"IsAllowDelete,omitnil,omitempty" name:"IsAllowDelete"`
}

type ReleaseConfigs struct {
	// 配置项描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigItem *string `json:"ConfigItem,omitnil,omitempty" name:"ConfigItem"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 变更后的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 变更前的内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastValue *string `json:"LastValue,omitnil,omitempty" name:"LastValue"`

	// 变更内容(优先级展示content内容,content为空取value内容)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ReleaseDoc struct {
	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文件类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 状态描述
	ActionDesc *string `json:"ActionDesc,omitnil,omitempty" name:"ActionDesc"`

	// 失败原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 文档业务ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type ReleaseQA struct {
	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 更新时间
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 状态描述
	ActionDesc *string `json:"ActionDesc,omitnil,omitempty" name:"ActionDesc"`

	// 来源1:文档生成，2：批量导入，3：手动添加
	Source *uint64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 来源描述
	SourceDesc *string `json:"SourceDesc,omitnil,omitempty" name:"SourceDesc"`

	// 文件名字
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档类型
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 失败原因
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 发布状态
	ReleaseStatus *uint64 `json:"ReleaseStatus,omitnil,omitempty" name:"ReleaseStatus"`

	// QAID
	QaBizId *string `json:"QaBizId,omitnil,omitempty" name:"QaBizId"`

	// 文档业务ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type ReleaseRejectedQuestion struct {
	// 问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Action *uint64 `json:"Action,omitnil,omitempty" name:"Action"`

	// 状态描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActionDesc *string `json:"ActionDesc,omitnil,omitempty" name:"ActionDesc"`

	// 失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

// Predefined struct for user
type RenameDocRequestParams struct {
	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 新文档名，需要带上后缀
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`
}

type RenameDocRequest struct {
	*tchttp.BaseRequest
	
	// 登录用户主账号(集成商模式必填)	
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)	
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 新文档名，需要带上后缀
	NewName *string `json:"NewName,omitnil,omitempty" name:"NewName"`
}

func (r *RenameDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	delete(f, "NewName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RenameDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RenameDocResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RenameDocResponse struct {
	*tchttp.BaseResponse
	Response *RenameDocResponseParams `json:"Response"`
}

func (r *RenameDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RenameDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocAuditRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocAuditRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *RetryDocAuditRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocAuditRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDocAuditRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocAuditResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryDocAuditResponse struct {
	*tchttp.BaseResponse
	Response *RetryDocAuditResponseParams `json:"Response"`
}

func (r *RetryDocAuditResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocParseRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type RetryDocParseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *RetryDocParseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocParseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryDocParseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryDocParseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryDocParseResponse struct {
	*tchttp.BaseResponse
	Response *RetryDocParseResponseParams `json:"Response"`
}

func (r *RetryDocParseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryDocParseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryReleaseRequestParams struct {
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布业务ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

type RetryReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 机器人ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 发布业务ID
	ReleaseBizId *string `json:"ReleaseBizId,omitnil,omitempty" name:"ReleaseBizId"`
}

func (r *RetryReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "ReleaseBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RetryReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryReleaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RetryReleaseResponse struct {
	*tchttp.BaseResponse
	Response *RetryReleaseResponseParams `json:"Response"`
}

func (r *RetryReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RetryReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunNodeInfo struct {
	// 节点类型，0:未指定，1:开始节点，2:API节点，3:询问节点，4:答案节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *int64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 请求的API
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvokeApi *InvokeAPI `json:"InvokeApi,omitnil,omitempty" name:"InvokeApi"`

	// 当前节点的所有槽位的值，key：SlotID。没有值的时候也要返回空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlotValues []*ValueInfo `json:"SlotValues,omitnil,omitempty" name:"SlotValues"`
}

// Predefined struct for user
type SaveDocRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档支持下面类型
	// pdf、doc、docx、ppt、mhtml、pptx、wps、ppsx，单个文件不超过200MB；
	// xlsx、xls、md、txt、csv、html，单个文件不超过20MB；
	// 
	// 图片支持下面类型：
	// jpg、png、jpeg、tiff、bmp、gif，单个文件不超过50MB
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 平台cos路径，与DescribeStorageCredential接口查询UploadPath参数保持一致
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// ETag 全称为 Entity Tag，是对象被创建时标识对象内容的信息标签，可用于检查对象的内容是否发生变化 成功上传cos后，从返回头中获取
	ETag *string `json:"ETag,omitnil,omitempty" name:"ETag"`

	// cos_hash x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性  
	// 成功上传cos后，从返回头中获取
	// 
	// 请注意：
	// cos_hash为文档唯一性标识，与文件名无关 相同的cos_hash会被判定为重复文档
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 标签适用范围，需要传参为1
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 来源（0 从本地文档导入），默认值为0
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 自定义链接地址, IsRefer为true的时候，该值才有意义
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 文档操作类型：1：批量导入（批量导入问答对）；2:文档导入（正常导入单个文档） 默认为1  <br> 请注意，opt=1的时候请从腾讯云智能体开发平台页面下载excel模板
	Opt *uint64 `json:"Opt,omitnil,omitempty" name:"Opt"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 是否可下载，IsRefer为true并且ReferUrlType为0时，该值才有意义
	IsDownload *bool `json:"IsDownload,omitnil,omitempty" name:"IsDownload"`
}

type SaveDocRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 文档支持下面类型
	// pdf、doc、docx、ppt、mhtml、pptx、wps、ppsx，单个文件不超过200MB；
	// xlsx、xls、md、txt、csv、html，单个文件不超过20MB；
	// 
	// 图片支持下面类型：
	// jpg、png、jpeg、tiff、bmp、gif，单个文件不超过50MB
	FileType *string `json:"FileType,omitnil,omitempty" name:"FileType"`

	// 平台cos路径，与DescribeStorageCredential接口查询UploadPath参数保持一致
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// ETag 全称为 Entity Tag，是对象被创建时标识对象内容的信息标签，可用于检查对象的内容是否发生变化 成功上传cos后，从返回头中获取
	ETag *string `json:"ETag,omitnil,omitempty" name:"ETag"`

	// cos_hash x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性  
	// 成功上传cos后，从返回头中获取
	// 
	// 请注意：
	// cos_hash为文档唯一性标识，与文件名无关 相同的cos_hash会被判定为重复文档
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 标签适用范围，需要传参为1
	AttrRange *uint64 `json:"AttrRange,omitnil,omitempty" name:"AttrRange"`

	// 来源（0 从本地文档导入），默认值为0
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// 自定义链接地址, IsRefer为true的时候，该值才有意义
	WebUrl *string `json:"WebUrl,omitnil,omitempty" name:"WebUrl"`

	// 标签引用
	AttrLabels []*AttrLabelRefer `json:"AttrLabels,omitnil,omitempty" name:"AttrLabels"`

	// 外部引用链接类型 0：系统链接 1：自定义链接
	// 值为1时，WebUrl 字段不能为空，否则不生效。
	ReferUrlType *uint64 `json:"ReferUrlType,omitnil,omitempty" name:"ReferUrlType"`

	// 有效开始时间，unix时间戳
	ExpireStart *string `json:"ExpireStart,omitnil,omitempty" name:"ExpireStart"`

	// 有效结束时间，unix时间戳，0代表永久有效
	ExpireEnd *string `json:"ExpireEnd,omitnil,omitempty" name:"ExpireEnd"`

	// 是否引用链接
	IsRefer *bool `json:"IsRefer,omitnil,omitempty" name:"IsRefer"`

	// 文档操作类型：1：批量导入（批量导入问答对）；2:文档导入（正常导入单个文档） 默认为1  <br> 请注意，opt=1的时候请从腾讯云智能体开发平台页面下载excel模板
	Opt *uint64 `json:"Opt,omitnil,omitempty" name:"Opt"`

	// 分类ID
	CateBizId *string `json:"CateBizId,omitnil,omitempty" name:"CateBizId"`

	// 是否可下载，IsRefer为true并且ReferUrlType为0时，该值才有意义
	IsDownload *bool `json:"IsDownload,omitnil,omitempty" name:"IsDownload"`
}

func (r *SaveDocRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDocRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "FileType")
	delete(f, "CosUrl")
	delete(f, "ETag")
	delete(f, "CosHash")
	delete(f, "Size")
	delete(f, "AttrRange")
	delete(f, "Source")
	delete(f, "WebUrl")
	delete(f, "AttrLabels")
	delete(f, "ReferUrlType")
	delete(f, "ExpireStart")
	delete(f, "ExpireEnd")
	delete(f, "IsRefer")
	delete(f, "Opt")
	delete(f, "CateBizId")
	delete(f, "IsDownload")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SaveDocRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SaveDocResponseParams struct {
	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`

	// 导入错误信息
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 错误链接
	ErrorLink *string `json:"ErrorLink,omitnil,omitempty" name:"ErrorLink"`

	// 错误链接文本
	ErrorLinkText *string `json:"ErrorLinkText,omitnil,omitempty" name:"ErrorLinkText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SaveDocResponse struct {
	*tchttp.BaseResponse
	Response *SaveDocResponseParams `json:"Response"`
}

func (r *SaveDocResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SaveDocResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchRange struct {
	// 检索条件and/or
	// 注意：此字段可能返回 null，表示取不到有效值。
	Condition *string `json:"Condition,omitnil,omitempty" name:"Condition"`

	// 自定义变量和标签关系数据	
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiVarAttrInfos []*ApiVarAttrInfo `json:"ApiVarAttrInfos,omitnil,omitempty" name:"ApiVarAttrInfos"`
}

type SearchStrategy struct {
	// 检索策略类型 0:混合检索，1：语义检索
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyType *uint64 `json:"StrategyType,omitnil,omitempty" name:"StrategyType"`

	// Excel检索增强开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	TableEnhancement *bool `json:"TableEnhancement,omitnil,omitempty" name:"TableEnhancement"`
}

type ShareKnowledgeBase struct {
	// 共享知识库ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 检索范围
	SearchRange *SearchRange `json:"SearchRange,omitnil,omitempty" name:"SearchRange"`
}

type SimilarQuestion struct {
	// 相似问ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SimBizId *string `json:"SimBizId,omitnil,omitempty" name:"SimBizId"`

	// 相似问内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 相似问审核状态，1审核失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuditStatus *uint64 `json:"AuditStatus,omitnil,omitempty" name:"AuditStatus"`
}

type SimilarQuestionModify struct {
	// 需要添加的相似问(内容)列表
	AddQuestions []*string `json:"AddQuestions,omitnil,omitempty" name:"AddQuestions"`

	// 需要更新的相似问列表
	UpdateQuestions []*SimilarQuestion `json:"UpdateQuestions,omitnil,omitempty" name:"UpdateQuestions"`

	// 需要删除的相似问列表
	DeleteQuestions []*SimilarQuestion `json:"DeleteQuestions,omitnil,omitempty" name:"DeleteQuestions"`
}

type Stat struct {
	// X轴: 时间区域；根据查询条件的粒度返回“分/小时/日”三种区间范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	X *string `json:"X,omitnil,omitempty" name:"X"`

	// Y轴: 该时间区域内的统计值，如token消耗量，调用次数或使用量等信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Y *float64 `json:"Y,omitnil,omitempty" name:"Y"`
}

type StatisticInfo struct {
	// 模型名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// 首Token耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstTokenCost *uint64 `json:"FirstTokenCost,omitnil,omitempty" name:"FirstTokenCost"`

	// 总耗时
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCost *uint64 `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// 输入Token数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputTokens *uint64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// 输出Token数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputTokens *uint64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// 总Token数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

// Predefined struct for user
type StopDocParseRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

type StopDocParseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文档ID
	DocBizId *string `json:"DocBizId,omitnil,omitempty" name:"DocBizId"`
}

func (r *StopDocParseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDocParseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "DocBizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopDocParseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopDocParseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopDocParseResponse struct {
	*tchttp.BaseResponse
	Response *StopDocParseResponseParams `json:"Response"`
}

func (r *StopDocParseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopDocParseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkflowRunRequestParams struct {
	// 工作流运行实例ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type StopWorkflowRunRequest struct {
	*tchttp.BaseRequest
	
	// 工作流运行实例ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

func (r *StopWorkflowRunRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkflowRunRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "WorkflowRunId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopWorkflowRunRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopWorkflowRunResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopWorkflowRunResponse struct {
	*tchttp.BaseResponse
	Response *StopWorkflowRunResponseParams `json:"Response"`
}

func (r *StopWorkflowRunResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopWorkflowRunResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrValue struct {
	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type SummaryConfig struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModel `json:"Model,omitnil,omitempty" name:"Model"`

	// 知识摘要输出配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *SummaryOutput `json:"Output,omitnil,omitempty" name:"Output"`

	// 欢迎语，200字符以内
	// 注意：此字段可能返回 null，表示取不到有效值。
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`
}

type SummaryOutput struct {
	// 输出方式 1：流式 2：非流式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Method *uint64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 输出要求 1：文本总结 2：自定义要求
	// 注意：此字段可能返回 null，表示取不到有效值。
	Requirement *uint64 `json:"Requirement,omitnil,omitempty" name:"Requirement"`

	// 自定义要求指令
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequireCommand *string `json:"RequireCommand,omitnil,omitempty" name:"RequireCommand"`
}

type TaskFLowVar struct {
	// 变量ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VarId *string `json:"VarId,omitnil,omitempty" name:"VarId"`

	// 变量名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VarName *string `json:"VarName,omitnil,omitempty" name:"VarName"`

	// 变量描述（默认为"-"）
	VarDesc *string `json:"VarDesc,omitnil,omitempty" name:"VarDesc"`

	// 变量类型 (STRING,INT,FLOAT,BOOL,OBJECT,ARRAY_STRING,ARRAY_INT,ARRAY_FLOAT,ARRAY_BOOL,ARRAY_OBJECT,FILE,DOCUMENT,IMAGE,AUDIO)
	VarType *string `json:"VarType,omitnil,omitempty" name:"VarType"`

	// 自定义变量默认值
	VarDefaultValue *string `json:"VarDefaultValue,omitnil,omitempty" name:"VarDefaultValue"`

	// 自定义变量文件默认名称
	VarDefaultFileName *string `json:"VarDefaultFileName,omitnil,omitempty" name:"VarDefaultFileName"`
}

type TaskFlowInfo struct {
	// 任务流程ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlowId *string `json:"TaskFlowId,omitnil,omitempty" name:"TaskFlowId"`

	// 任务流程名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskFlowName *string `json:"TaskFlowName,omitnil,omitempty" name:"TaskFlowName"`

	// Query 重写结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryRewrite *string `json:"QueryRewrite,omitnil,omitempty" name:"QueryRewrite"`

	// 命中意图
	// 注意：此字段可能返回 null，表示取不到有效值。
	HitIntent *string `json:"HitIntent,omitnil,omitempty" name:"HitIntent"`

	// 任务流程回复类型
	// 0: 任务流回复
	// 1: 任务流静默
	// 2: 任务流拉回话术
	// 3: 任务流自定义回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *uint64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type TaskFlowSummary struct {
	// 任务流程名
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntentName *string `json:"IntentName,omitnil,omitempty" name:"IntentName"`

	// 实体列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedSlotValues []*ValueInfo `json:"UpdatedSlotValues,omitnil,omitempty" name:"UpdatedSlotValues"`

	// 节点列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunNodes []*RunNodeInfo `json:"RunNodes,omitnil,omitempty" name:"RunNodes"`

	// 意图判断
	// 注意：此字段可能返回 null，表示取不到有效值。
	Purposes []*string `json:"Purposes,omitnil,omitempty" name:"Purposes"`
}

type TaskParams struct {
	// 下载地址,需要通过cos桶临时密钥去下载
	// 注意：此字段可能返回 null，表示取不到有效值。
	CosPath *string `json:"CosPath,omitnil,omitempty" name:"CosPath"`
}

type TokenStat struct {
	// 会话 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 请求 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// 对应哪条会话, 会话 ID, 用于回答的消息存储使用, 可提前生成, 保存消息时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// token 已使用数
	// 注意：此字段可能返回 null，表示取不到有效值。
	UsedCount *uint64 `json:"UsedCount,omitnil,omitempty" name:"UsedCount"`

	// 免费 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	FreeCount *uint64 `json:"FreeCount,omitnil,omitempty" name:"FreeCount"`

	// 订单总 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	OrderCount *uint64 `json:"OrderCount,omitnil,omitempty" name:"OrderCount"`

	// 当前执行状态汇总, 常量: 使用中: processing, 成功: success, 失败: failed
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusSummary *string `json:"StatusSummary,omitnil,omitempty" name:"StatusSummary"`

	// 当前执行状态汇总后中文展示
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusSummaryTitle *string `json:"StatusSummaryTitle,omitnil,omitempty" name:"StatusSummaryTitle"`

	// 当前请求执行时间, 单位 ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	Elapsed *uint64 `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// 当前请求消耗 token 数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenCount *uint64 `json:"TokenCount,omitnil,omitempty" name:"TokenCount"`

	// 执行过程信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Procedures []*Procedure `json:"Procedures,omitnil,omitempty" name:"Procedures"`

	// 执行过程信息TraceId
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`
}

type UnsatisfiedReply struct {
	// 不满意回复ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplyBizId *string `json:"ReplyBizId,omitnil,omitempty" name:"ReplyBizId"`

	// 消息记录ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordBizId *string `json:"RecordBizId,omitnil,omitempty" name:"RecordBizId"`

	// 用户问题
	// 注意：此字段可能返回 null，表示取不到有效值。
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 应用回复
	// 注意：此字段可能返回 null，表示取不到有效值。
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 错误类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Reasons []*string `json:"Reasons,omitnil,omitempty" name:"Reasons"`
}

// Predefined struct for user
type UpdateSharedKnowledgeRequestParams struct {
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 共享知识库更新信息
	Info *KnowledgeUpdateInfo `json:"Info,omitnil,omitempty" name:"Info"`
}

type UpdateSharedKnowledgeRequest struct {
	*tchttp.BaseRequest
	
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 共享知识库更新信息
	Info *KnowledgeUpdateInfo `json:"Info,omitnil,omitempty" name:"Info"`
}

func (r *UpdateSharedKnowledgeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSharedKnowledgeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "KnowledgeBizId")
	delete(f, "Info")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateSharedKnowledgeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateSharedKnowledgeResponseParams struct {
	// 共享知识库业务ID
	KnowledgeBizId *string `json:"KnowledgeBizId,omitnil,omitempty" name:"KnowledgeBizId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateSharedKnowledgeResponse struct {
	*tchttp.BaseResponse
	Response *UpdateSharedKnowledgeResponseParams `json:"Response"`
}

func (r *UpdateSharedKnowledgeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateSharedKnowledgeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateVarRequestParams struct {
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量ID
	VarId *string `json:"VarId,omitnil,omitempty" name:"VarId"`

	// 变量名称，最大支持50个字符
	VarName *string `json:"VarName,omitnil,omitempty" name:"VarName"`

	// 参数描述
	VarDesc *string `json:"VarDesc,omitnil,omitempty" name:"VarDesc"`

	// 参数类型
	VarType *string `json:"VarType,omitnil,omitempty" name:"VarType"`

	// 自定义变量默认值
	VarDefaultValue *string `json:"VarDefaultValue,omitnil,omitempty" name:"VarDefaultValue"`

	// 自定义变量文件默认名称
	VarDefaultFileName *string `json:"VarDefaultFileName,omitnil,omitempty" name:"VarDefaultFileName"`
}

type UpdateVarRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 变量ID
	VarId *string `json:"VarId,omitnil,omitempty" name:"VarId"`

	// 变量名称，最大支持50个字符
	VarName *string `json:"VarName,omitnil,omitempty" name:"VarName"`

	// 参数描述
	VarDesc *string `json:"VarDesc,omitnil,omitempty" name:"VarDesc"`

	// 参数类型
	VarType *string `json:"VarType,omitnil,omitempty" name:"VarType"`

	// 自定义变量默认值
	VarDefaultValue *string `json:"VarDefaultValue,omitnil,omitempty" name:"VarDefaultValue"`

	// 自定义变量文件默认名称
	VarDefaultFileName *string `json:"VarDefaultFileName,omitnil,omitempty" name:"VarDefaultFileName"`
}

func (r *UpdateVarRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateVarRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppBizId")
	delete(f, "VarId")
	delete(f, "VarName")
	delete(f, "VarDesc")
	delete(f, "VarType")
	delete(f, "VarDefaultValue")
	delete(f, "VarDefaultFileName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateVarRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateVarResponseParams struct {
	// 变量ID
	VarId *string `json:"VarId,omitnil,omitempty" name:"VarId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateVarResponse struct {
	*tchttp.BaseResponse
	Response *UpdateVarResponseParams `json:"Response"`
}

func (r *UpdateVarResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateVarResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadAttributeLabelRequestParams struct {
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type UploadAttributeLabelRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// cos路径
	CosUrl *string `json:"CosUrl,omitnil,omitempty" name:"CosUrl"`

	// x-cos-hash-crc64ecma 头部中的 CRC64编码进行校验上传到云端的文件和本地文件的一致性
	CosHash *string `json:"CosHash,omitnil,omitempty" name:"CosHash"`

	// 文件大小
	Size *string `json:"Size,omitnil,omitempty" name:"Size"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *UploadAttributeLabelRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAttributeLabelRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BotBizId")
	delete(f, "FileName")
	delete(f, "CosUrl")
	delete(f, "CosHash")
	delete(f, "Size")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UploadAttributeLabelRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UploadAttributeLabelResponseParams struct {
	// 导入错误
	ErrorMsg *string `json:"ErrorMsg,omitnil,omitempty" name:"ErrorMsg"`

	// 错误链接
	ErrorLink *string `json:"ErrorLink,omitnil,omitempty" name:"ErrorLink"`

	// 错误链接文本
	ErrorLinkText *string `json:"ErrorLinkText,omitnil,omitempty" name:"ErrorLinkText"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UploadAttributeLabelResponse struct {
	*tchttp.BaseResponse
	Response *UploadAttributeLabelResponseParams `json:"Response"`
}

func (r *UploadAttributeLabelResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UploadAttributeLabelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserBaseInfo struct {
	// 用户ID
	UserBizId *string `json:"UserBizId,omitnil,omitempty" name:"UserBizId"`

	// 用户名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type ValueInfo struct {
	// 值ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 值类型：0:未知或者空, 1:string, 2:int, 3:float, 4:bool, 5:array(字符串数组), 6: object_array(结构体数组), 7: object(结构体)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueType *int64 `json:"ValueType,omitnil,omitempty" name:"ValueType"`

	// string
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueStr *string `json:"ValueStr,omitnil,omitempty" name:"ValueStr"`

	// int（避免精度丢失使用字符串返回）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueInt *string `json:"ValueInt,omitnil,omitempty" name:"ValueInt"`

	// float
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueFloat *float64 `json:"ValueFloat,omitnil,omitempty" name:"ValueFloat"`

	// bool
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueBool *bool `json:"ValueBool,omitnil,omitempty" name:"ValueBool"`

	// array
	// 注意：此字段可能返回 null，表示取不到有效值。
	ValueStrArray []*string `json:"ValueStrArray,omitnil,omitempty" name:"ValueStrArray"`
}

// Predefined struct for user
type VerifyQARequestParams struct {
	// 问答列表
	List []*QAList `json:"List,omitnil,omitempty" name:"List"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type VerifyQARequest struct {
	*tchttp.BaseRequest
	
	// 问答列表
	List []*QAList `json:"List,omitnil,omitempty" name:"List"`

	// 应用ID
	BotBizId *string `json:"BotBizId,omitnil,omitempty" name:"BotBizId"`

	// 登录用户主账号(集成商模式必填)
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// 登录用户子账号(集成商模式必填)
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *VerifyQARequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyQARequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "List")
	delete(f, "BotBizId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "VerifyQARequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type VerifyQAResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type VerifyQAResponse struct {
	*tchttp.BaseResponse
	Response *VerifyQAResponseParams `json:"Response"`
}

func (r *VerifyQAResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *VerifyQAResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VoiceConfig struct {
	// 公有云音色id
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceType *uint64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`

	// 音色key
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimbreKey *string `json:"TimbreKey,omitnil,omitempty" name:"TimbreKey"`

	// 音色名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`
}

type WorkFlowSummary struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流运行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 节点信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RunNodes []*WorkflowRunNodeInfo `json:"RunNodes,omitnil,omitempty" name:"RunNodes"`

	// 选项卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// 多气泡的输出结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outputs []*string `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// 工作流发布时间，unix时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowReleaseTime *string `json:"WorkflowReleaseTime,omitnil,omitempty" name:"WorkflowReleaseTime"`
}

type WorkflowInfo struct {
	// 工作流ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 工作流运行ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 选项卡
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// 多气泡的输出结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Outputs []*string `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// 工作流发布时间，unix时间戳
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowReleaseTime *string `json:"WorkflowReleaseTime,omitnil,omitempty" name:"WorkflowReleaseTime"`
}

type WorkflowRef struct {
	// 任务流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 任务流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// 任务流描述
	WorkflowDesc *string `json:"WorkflowDesc,omitnil,omitempty" name:"WorkflowDesc"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 更新时间
	UpdateTime *uint64 `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type WorkflowRunBase struct {
	// 运行环境。0: 测试环境； 1: 正式环境
	RunEnv *uint64 `json:"RunEnv,omitnil,omitempty" name:"RunEnv"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 工作流运行实例的ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 所属工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 运行状态。0: 排队中；1: 运行中；2: 运行成功；3: 运行失败； 4: 已取消
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 错误信息
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// 消耗的token总数
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// 创建时间（毫秒时间戳）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 开始时间（毫秒时间戳）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（毫秒时间戳）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type WorkflowRunDetail struct {
	// 运行环境。0: 测试环境； 1: 正式环境
	RunEnv *uint64 `json:"RunEnv,omitnil,omitempty" name:"RunEnv"`

	// 应用ID
	AppBizId *string `json:"AppBizId,omitnil,omitempty" name:"AppBizId"`

	// 工作流运行实例的ID
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`

	// 所属工作流ID
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 运行状态。0: 排队中；1: 运行中；2: 运行成功；3: 运行失败； 4: 已取消
	State *uint64 `json:"State,omitnil,omitempty" name:"State"`

	// 错误信息
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// 消耗的token总数
	TotalTokens *uint64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// 创建时间（毫秒时间戳）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 开始时间（毫秒时间戳）
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（毫秒时间戳）
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 工作流画布Json
	DialogJson *string `json:"DialogJson,omitnil,omitempty" name:"DialogJson"`

	// 用户的输入
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// 主模型名称
	MainModelName *string `json:"MainModelName,omitnil,omitempty" name:"MainModelName"`

	// API参数配置
	CustomVariables []*CustomVariable `json:"CustomVariables,omitnil,omitempty" name:"CustomVariables"`
}

type WorkflowRunNodeInfo struct {
	// 节点ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *uint64 `json:"NodeType,omitnil,omitempty" name:"NodeType"`

	// 节点名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 输入
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *string `json:"Input,omitnil,omitempty" name:"Input"`

	// 输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *string `json:"Output,omitnil,omitempty" name:"Output"`

	// 任务输出
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskOutput *string `json:"TaskOutput,omitnil,omitempty" name:"TaskOutput"`

	// 错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	FailMessage *string `json:"FailMessage,omitnil,omitempty" name:"FailMessage"`

	// 花费时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	CostMilliSeconds *uint64 `json:"CostMilliSeconds,omitnil,omitempty" name:"CostMilliSeconds"`

	// 大模型输出信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatisticInfos []*StatisticInfo `json:"StatisticInfos,omitnil,omitempty" name:"StatisticInfos"`
}