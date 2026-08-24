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

package v20260520

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AICallConfig struct {
	// 数智人配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	DigitalHuman *DigitalHumanConfig `json:"DigitalHuman,omitnil,omitempty" name:"DigitalHuman"`

	// 启用数智人
	EnableDigitalHuman *bool `json:"EnableDigitalHuman,omitnil,omitempty" name:"EnableDigitalHuman"`

	// 启用语音通话
	EnableVoiceCall *bool `json:"EnableVoiceCall,omitnil,omitempty" name:"EnableVoiceCall"`

	// 启用语音互动功能
	EnableVoiceInteract *bool `json:"EnableVoiceInteract,omitnil,omitempty" name:"EnableVoiceInteract"`

	// 音色配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Voice *VoiceConfig `json:"Voice,omitnil,omitempty" name:"Voice"`
}

type AIOptimizeModel struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type AccountInfo struct {
	// <p>员工子账号id</p>
	AccountUin *string `json:"AccountUin,omitnil,omitempty" name:"AccountUin"`

	// <p>员工昵称</p>
	NickName *string `json:"NickName,omitnil,omitempty" name:"NickName"`

	// <p>员工头像</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`
}

type AgentAdvancedConfig struct {
	// <p>最大推理轮数</p>
	MaxReasoningRound *uint64 `json:"MaxReasoningRound,omitnil,omitempty" name:"MaxReasoningRound"`
}

type AgentCollaborationConfig struct {
	// 协同方式。枚举值: 1:自由转交：Agent之间可自由传递任务, 2:工作流编排：基于预定义流程的协同, 3:Plan-and-Execute：规划与执行分离的协同模式
	AgentCollaborationMode *int64 `json:"AgentCollaborationMode,omitnil,omitempty" name:"AgentCollaborationMode"`

	// 工作流Id
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// 工作流名称
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`
}

type AgentDetail struct {
	// <p>Agent ID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent基本配置</p>
	Profile *AgentProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>系统提示词</p>
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`

	// <p>模型信息</p>
	Model *AgentModelConfig `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>工具详情</p>
	ToolList []*AgentTool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>插件配置</p>
	PluginList []*AgentPlugin `json:"PluginList,omitnil,omitempty" name:"PluginList"`

	// <p>技能详情</p>
	SkillList []*AgentSkill `json:"SkillList,omitnil,omitempty" name:"SkillList"`

	// <p>高级配置</p>
	AdvancedConfig *AgentAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type AgentInput struct {
	// <p>输入来源类型：0 用户输入，3 自定义变量（API参数）</p>
	InputType *int64 `json:"InputType,omitnil,omitempty" name:"InputType"`

	// <p>用户手写输入</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserInputValue *AgentUserInputValue `json:"UserInputValue,omitnil,omitempty" name:"UserInputValue"`

	// <p>系统参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemVariable *AgentSystemVariable `json:"SystemVariable,omitnil,omitempty" name:"SystemVariable"`

	// <p>自定义变量（API参数）</p>
	CustomVariableId *string `json:"CustomVariableId,omitnil,omitempty" name:"CustomVariableId"`

	// <p>环境变量参数</p>
	EnvVariableId *string `json:"EnvVariableId,omitnil,omitempty" name:"EnvVariableId"`

	// <p>应用变量参数</p>
	AppVariableId *string `json:"AppVariableId,omitnil,omitempty" name:"AppVariableId"`
}

type AgentModelConfig struct {
	// <p>模型唯一id</p>
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// <p>模型别名</p>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// <p>模型上下文长度字符限制</p>
	ContextWordsLimit *uint64 `json:"ContextWordsLimit,omitnil,omitempty" name:"ContextWordsLimit"`

	// <p>指令长度字符限制</p>
	InstructionsWordsLimit *uint64 `json:"InstructionsWordsLimit,omitnil,omitempty" name:"InstructionsWordsLimit"`

	// <p>模型参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelParameters *ModelParams `json:"ModelParameters,omitnil,omitempty" name:"ModelParameters"`
}

type AgentPlugin struct {
	// <p>插件基本配置</p>
	Config *AgentPluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>插件名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>插件图标url</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>插件描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>插件产品分类</p><p>枚举值：</p><ul><li>0： 普通插件</li><li>1： 连接器类插件</li></ul>
	PluginClass *int64 `json:"PluginClass,omitnil,omitempty" name:"PluginClass"`

	// <p>插件状态</p><p>枚举值：</p><ul><li>0： 未知</li><li>1： 可用</li><li>2： 不可用</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>插件鉴权配置状态</p><p>枚举值：</p><ul><li>0： 不需要授权</li><li>1： 未配置</li><li>2： 已配置</li></ul>
	AuthConfigStatus *int64 `json:"AuthConfigStatus,omitnil,omitempty" name:"AuthConfigStatus"`
}

type AgentPluginConfig struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>插件 Header 参数</p>
	HeaderParameterList []*AgentPluginParameter `json:"HeaderParameterList,omitnil,omitempty" name:"HeaderParameterList"`

	// <p>插件 Query 参数</p>
	QueryParameterList []*AgentPluginParameter `json:"QueryParameterList,omitnil,omitempty" name:"QueryParameterList"`

	// <p>是否使用CAM一键授权，仅 auth_type=2时生效</p>
	EnableCamRoleAuth *bool `json:"EnableCamRoleAuth,omitnil,omitempty" name:"EnableCamRoleAuth"`

	// <p>授权类型</p><p>枚举值：</p><ul><li>0： 无鉴权</li><li>1： API Key</li><li>2： CAM授权</li><li>3： OAuth2.0授权</li></ul>
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// <p>OAuth 授权同意模式；0-开发者授权；1-使用者授权（仅在auth_type=3时生效）</p>
	OAuthConsent *int64 `json:"OAuthConsent,omitnil,omitempty" name:"OAuthConsent"`
}

type AgentPluginParameter struct {
	// <p>参数名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>是否必填</p>
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// <p>输入的值</p>
	Input *AgentInput `json:"Input,omitnil,omitempty" name:"Input"`
}

type AgentProfile struct {
	// <p>Agent名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>图标URL</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Agent 角色：0=主 / 1=子</p>
	Role *int64 `json:"Role,omitnil,omitempty" name:"Role"`

	// <p>Agent 描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>应用名称</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>开发者</p>
	Developer *string `json:"Developer,omitnil,omitempty" name:"Developer"`

	// <p>主AgentId，只读，不可通过修改接口进行变更</p>
	ParentAgentId *string `json:"ParentAgentId,omitnil,omitempty" name:"ParentAgentId"`
}

type AgentRelease struct {
	// <p>名称</p>
	ItemName *string `json:"ItemName,omitnil,omitempty" name:"ItemName"`

	// <p>更新时间, unix 秒时间戳 (s)</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>动作描述</p>
	ActionDescription *string `json:"ActionDescription,omitnil,omitempty" name:"ActionDescription"`

	// <p>变更为 测试</p>
	ReleaseMessage *string `json:"ReleaseMessage,omitnil,omitempty" name:"ReleaseMessage"`
}

type AgentReleasePreview struct {
	// <p>AgentID</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>更新时间, unix 秒时间戳 (s)</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>状态, 状态值：1:新增, 2:修改, 3:删除</p>
	Action *int64 `json:"Action,omitnil,omitempty" name:"Action"`

	// <p>动作描述</p>
	ActionDescription *string `json:"ActionDescription,omitnil,omitempty" name:"ActionDescription"`

	// <p>发布消息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// <p>发布详情</p>
	ReleaseList []*AgentRelease `json:"ReleaseList,omitnil,omitempty" name:"ReleaseList"`
}

type AgentSkill struct {
	// <p>skillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>skill名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>技能描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>skill展示名称</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>技能展示描述</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill图标url</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>Skill来源</p>
	SourceType *int64 `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// <p>Skill版本</p>
	CurrentVersion *string `json:"CurrentVersion,omitnil,omitempty" name:"CurrentVersion"`
}

type AgentSkillConfig struct {
	// <p>技能ID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`
}

type AgentSpec struct {
	// <p>Agent基本配置</p>
	Profile *AgentProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 系统提示词
	Instructions *string `json:"Instructions,omitnil,omitempty" name:"Instructions"`

	// 主模型配置
	Model *AgentModelConfig `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>工具信息</p>
	ToolList []*AgentToolConfig `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>插件信息</p>
	PluginList []*AgentPluginConfig `json:"PluginList,omitnil,omitempty" name:"PluginList"`

	// <p>技能信息</p>
	SkillList []*AgentSkillConfig `json:"SkillList,omitnil,omitempty" name:"SkillList"`

	// 高级设置
	AdvancedConfig *AgentAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type AgentSummary struct {
	// <p>AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>Agent 身份画像</p>
	Profile *AgentProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>高级设置;scope=0 时返回</p>
	AdvancedConfig *AgentAdvancedConfig `json:"AdvancedConfig,omitnil,omitempty" name:"AdvancedConfig"`
}

type AgentSystemVariable struct {
	// <p>系统参数名</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>对话历史轮数的配置；如果Input是系统变量中的“对话历史”时才使用；</p>
	DialogHistoryLimit *int64 `json:"DialogHistoryLimit,omitnil,omitempty" name:"DialogHistoryLimit"`
}

type AgentTool struct {
	// <p>工具配置字段</p>
	Config *AgentToolBasicConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>工具名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>工具状态</p><p>枚举值：</p><ul><li>1： 可用</li><li>2： 不可用</li><li>3： 已失效</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>调用方式</p><p>枚举值：</p><ul><li>0： 非流式</li><li>1： 流式</li></ul>
	StreamMode *int64 `json:"StreamMode,omitnil,omitempty" name:"StreamMode"`

	// <p>工具访问模式</p><p>枚举值：</p><ul><li>0： 未指定</li><li>1： 只读</li><li>2： 写/删除</li></ul>
	ToolAccessMode *int64 `json:"ToolAccessMode,omitnil,omitempty" name:"ToolAccessMode"`
}

type AgentToolBasicConfig struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>工具id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>工具输入参数列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputList []*AgentToolInputParameter `json:"InputList,omitnil,omitempty" name:"InputList"`

	// <p>工具输出参数列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputList []*AgentToolOutputParameter `json:"OutputList,omitnil,omitempty" name:"OutputList"`

	// <p>工具Header参数列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeaderParameterList []*AgentPluginParameter `json:"HeaderParameterList,omitnil,omitempty" name:"HeaderParameterList"`

	// <p>工具Query参数列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueryParameterList []*AgentPluginParameter `json:"QueryParameterList,omitnil,omitempty" name:"QueryParameterList"`

	// <p>工具来源: 0-来自插件，1-来自工作流</p>
	ToolSource *uint64 `json:"ToolSource,omitnil,omitempty" name:"ToolSource"`

	// <p>是否禁用</p>
	IsDisabled *bool `json:"IsDisabled,omitnil,omitempty" name:"IsDisabled"`
}

type AgentToolConfig struct {
	// <p>工具配置</p>
	Config *AgentToolBasicConfig `json:"Config,omitnil,omitempty" name:"Config"`
}

type AgentToolInputParameter struct {
	// <p>工具名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>工具描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>工具参数类型</p><p>枚举值：</p><ul><li>0： STRING</li><li>1： INT</li><li>2： FLOAT</li><li>3： BOOL</li><li>4： OBJECT</li><li>5： ARRAY_STRING</li><li>6： ARRAY_INT</li><li>7： ARRAY_FLOAT</li><li>8： ARRAY_BOOL</li><li>9： ARRAY_OBJECT</li><li>20： ARRAY_ARRAY</li><li>99： NULL</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>是否必填</p>
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// <p>子参数，仅 OBJECT 或 ARRAY&lt;&gt; 类型时使用</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubParameterList []*AgentToolInputParameter `json:"SubParameterList,omitnil,omitempty" name:"SubParameterList"`

	// <p>模式下是否对模型隐藏</p>
	IsHidden *bool `json:"IsHidden,omitnil,omitempty" name:"IsHidden"`

	// <p>OneOf类型的参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OneOfList []*AgentToolInputParameter `json:"OneOfList,omitnil,omitempty" name:"OneOfList"`

	// <p>AnyOf类型的参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnyOfList []*AgentToolInputParameter `json:"AnyOfList,omitnil,omitempty" name:"AnyOfList"`

	// <p>参数取值来源</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *AgentInput `json:"Input,omitnil,omitempty" name:"Input"`
}

type AgentToolOutputParameter struct {
	// <p>参数名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>变量描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>参数类型</p><p>枚举值：</p><ul><li>0： STRING</li><li>1： INT</li><li>2： FLOAT</li><li>3： BOOL</li><li>4： OBJECT</li><li>5： ARRAY_STRING</li><li>6： ARRAY_INT</li><li>7： ARRAY_FLOAT</li><li>8： ARRAY_BOOL</li><li>9： ARRAY_OBJECT</li><li>20： ARRAY_ARRAY</li><li>99： NULL</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>子参数，仅 OBJECT 或 ARRAY_OBJECT 类型时使用</p>
	SubParameterList []*AgentToolOutputParameter `json:"SubParameterList,omitnil,omitempty" name:"SubParameterList"`

	// <p>解析方式</p>
	RenderMode *int64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`
}

type AgentUserInputValue struct {
	// <p>用户输入参数值</p>
	ValueList []*string `json:"ValueList,omitnil,omitempty" name:"ValueList"`
}

type ApiKeyAuthConfig struct {
	// 密钥位置 HEADER/QUERY
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | Header鉴权 |
	// | 1 | Query鉴权 |
	KeyLocation *int64 `json:"KeyLocation,omitnil,omitempty" name:"KeyLocation"`

	// 密钥参数名
	KeyParamName *string `json:"KeyParamName,omitnil,omitempty" name:"KeyParamName"`

	// 密钥参数值
	KeyParamValue *string `json:"KeyParamValue,omitnil,omitempty" name:"KeyParamValue"`
}

type ApiPluginConfig struct {
	// 授权配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuthConfig *AuthConfig `json:"AuthConfig,omitnil,omitempty" name:"AuthConfig"`
}

type ApiToolConfig struct {
	// <p>请求体参数</p>
	Body []*RequestParam `json:"Body,omitnil,omitempty" name:"Body"`

	// <p>示例</p>
	Example *ToolExample `json:"Example,omitnil,omitempty" name:"Example"`

	// <p>API插件外部调用地址</p>
	ExternalApiUrl *string `json:"ExternalApiUrl,omitnil,omitempty" name:"ExternalApiUrl"`

	// <p>Header</p>
	Header []*RequestParam `json:"Header,omitnil,omitempty" name:"Header"`

	// <p>请求方式</p>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>输出</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`

	// <p>查询参数</p>
	Query []*RequestParam `json:"Query,omitnil,omitempty" name:"Query"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>STREAM_MODE_UNARY</td><td>0</td><td>非流式</td></tr><tr><td>STREAM_MODE_STREAMING</td><td>1</td><td>流式</td></tr></tbody></table>
	StreamMode *int64 `json:"StreamMode,omitnil,omitempty" name:"StreamMode"`

	// <p>地址</p>
	Url *string `json:"Url,omitnil,omitempty" name:"Url"`
}

type App struct {
	// <p>辅助信息(子状态/审批/申诉/搜索资源/特殊状态等)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AuxiliaryInfo *AppAuxiliaryInfo `json:"AuxiliaryInfo,omitnil,omitempty" name:"AuxiliaryInfo"`

	// <p>配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *AppConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>元数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadata *AppMetadata `json:"Metadata,omitnil,omitempty" name:"Metadata"`

	// <p>应用密钥信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SecretInfo *AppSecretInfo `json:"SecretInfo,omitnil,omitempty" name:"SecretInfo"`

	// <p>分享链接信息(含访问控制)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ShareUrlInfo *AppShareURLInfo `json:"ShareUrlInfo,omitnil,omitempty" name:"ShareUrlInfo"`

	// <p>状态</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *AppStatusInfo `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>应用引用的共享知识库列表</p>
	SharedKbList []*AppSharedKbInfo `json:"SharedKbList,omitnil,omitempty" name:"SharedKbList"`

	// <p>企业共享配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`
}

type AppAdvancedConf struct {
	// <p>是否开启上下文改写</p>
	EnableContextRewrite *bool `json:"EnableContextRewrite,omitnil,omitempty" name:"EnableContextRewrite"`

	// <p>是否开启图文检索</p>
	EnableImageTextRetrieval *bool `json:"EnableImageTextRetrieval,omitnil,omitempty" name:"EnableImageTextRetrieval"`

	// <p>回复灵活度</p>
	ReplyFlexibility *uint64 `json:"ReplyFlexibility,omitnil,omitempty" name:"ReplyFlexibility"`

	// <p>对话端自定义配置(所有模式共用,允许对话中动态修改配置)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DialogCustomConfig *DialogCustomConfig `json:"DialogCustomConfig,omitnil,omitempty" name:"DialogCustomConfig"`

	// <p>意图达成优先级</p>
	IntentAchievement []*IntentAchievementInfo `json:"IntentAchievement,omitnil,omitempty" name:"IntentAchievement"`
}

type AppAppeal struct {
	// 申诉中的配置项
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppealingStatus *AppealingStatus `json:"AppealingStatus,omitnil,omitempty" name:"AppealingStatus"`
}

type AppAuxiliaryInfo struct {
	// 申诉信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Appeal *AppAppeal `json:"Appeal,omitnil,omitempty" name:"Appeal"`

	// 搜索资源状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	SearchResourceStatus *SearchResourceStatusInfo `json:"SearchResourceStatus,omitnil,omitempty" name:"SearchResourceStatus"`

	// 特殊状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecialStatusInfo *SpecialStatusInfo `json:"SpecialStatusInfo,omitnil,omitempty" name:"SpecialStatusInfo"`

	// 子状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubStatus *AppSubStatusInfo `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`
}

type AppConfig struct {
	// 体验配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Experience *AppExperienceConfig `json:"Experience,omitnil,omitempty" name:"Experience"`

	// 欢迎语配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Greeting *AppGreetingConfig `json:"Greeting,omitnil,omitempty" name:"Greeting"`

	// 记忆配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Memory *AppMemoryConfig `json:"Memory,omitnil,omitempty" name:"Memory"`

	// 模式相关配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Mode *AppModeConfig `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *AppModelConfig `json:"Model,omitnil,omitempty" name:"Model"`

	// 联网搜索配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebSearch *AppWebSearchConfig `json:"WebSearch,omitnil,omitempty" name:"WebSearch"`

	// 工作流配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workflow *AppWorkflowConfig `json:"Workflow,omitnil,omitempty" name:"Workflow"`
}

type AppExperienceConfig struct {
	// 高级配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Advanced *AppAdvancedConf `json:"Advanced,omitnil,omitempty" name:"Advanced"`

	// 对话体验配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Conversation *ConversationExperience `json:"Conversation,omitnil,omitempty" name:"Conversation"`

	// 角色配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Role *RoleConfig `json:"Role,omitnil,omitempty" name:"Role"`
}

type AppGreetingConfig struct {
	// <p>欢迎语内容</p>
	Greeting *string `json:"Greeting,omitnil,omitempty" name:"Greeting"`

	// <p>开场问题列表</p>
	OpeningQuestionList []*string `json:"OpeningQuestionList,omitnil,omitempty" name:"OpeningQuestionList"`
}

type AppMemoryConfig struct {
	// 是否开启长记忆
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 长记忆时长
	LongMemoryDay *uint64 `json:"LongMemoryDay,omitnil,omitempty" name:"LongMemoryDay"`

	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`

	// prompt内容
	PromptContent *string `json:"PromptContent,omitnil,omitempty" name:"PromptContent"`

	// 提示词模式。枚举值: 1:自定义
	PromptMode *int64 `json:"PromptMode,omitnil,omitempty" name:"PromptMode"`
}

type AppMetadata struct {
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 应用模式。枚举值: 1:标准模式, 2:Agent模式, 3:单工作流模式, 4:ClawAgent模式
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// 应用头像
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 创建时间 (Unix时间戳,秒级)
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 应用描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 更新时间 (Unix时间戳,秒级)
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type AppModeConfig struct {
	// 多智能体配置(Agent模式)
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiAgentConfig *MultiAgentConfig `json:"MultiAgentConfig,omitnil,omitempty" name:"MultiAgentConfig"`

	// 单工作流配置(单工作流模式)
	// 注意：此字段可能返回 null，表示取不到有效值。
	SingleWorkflowConfig *SingleWorkflowConfig `json:"SingleWorkflowConfig,omitnil,omitempty" name:"SingleWorkflowConfig"`

	// ClawAgent配置(ClawAgent模式)
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClawAgentConfig *ClawAgentConfig `json:"ClawAgentConfig,omitnil,omitempty" name:"ClawAgentConfig"`
}

type AppModelConfig struct {
	// AI一键优化模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiOptimizeModel *AIOptimizeModel `json:"AiOptimizeModel,omitnil,omitempty" name:"AiOptimizeModel"`

	// 实时文件解析模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileParseModel *FileParseModel `json:"FileParseModel,omitnil,omitempty" name:"FileParseModel"`

	// 生成模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	GenerateModel *GenerateModel `json:"GenerateModel,omitnil,omitempty" name:"GenerateModel"`

	// 多模态问答模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiModalQaModel *MultiModalQAModel `json:"MultiModalQaModel,omitnil,omitempty" name:"MultiModalQaModel"`

	// 多模态理解模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiModalUnderstandingModel *MultiModalUnderstandingModel `json:"MultiModalUnderstandingModel,omitnil,omitempty" name:"MultiModalUnderstandingModel"`

	// Prompt改写模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PromptRewriteModel *PromptRewriteModel `json:"PromptRewriteModel,omitnil,omitempty" name:"PromptRewriteModel"`

	// 思考模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThinkModel *ThinkModel `json:"ThinkModel,omitnil,omitempty" name:"ThinkModel"`
}

type AppOperation struct {
	// 创建人
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// 创建人UIN
	CreatorUin *string `json:"CreatorUin,omitnil,omitempty" name:"CreatorUin"`

	// 创建人账号(私有化场景使用)
	CreatorUserAccount *string `json:"CreatorUserAccount,omitnil,omitempty" name:"CreatorUserAccount"`

	// 修改时间 (Unix时间戳,秒级)
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 最后修改人
	Updater *string `json:"Updater,omitnil,omitempty" name:"Updater"`

	// 修改人UIN
	UpdaterUin *string `json:"UpdaterUin,omitnil,omitempty" name:"UpdaterUin"`
}

type AppPluginConfig struct {
	// 基于发布应用创建插件的应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type AppSecretInfo struct {
	// 应用密钥
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`
}

type AppShareAccessControl struct {
	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_SHARE_ACCESS_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>APP_SHARE_ACCESS_TYPE_PUBLIC</td><td>1</td><td>公开访问(所有用户都可访问)</td></tr><tr><td>APP_SHARE_ACCESS_TYPE_INTERNAL</td><td>2</td><td>内部访问(仅企业用户可访问)</td></tr><tr><td>APP_SHARE_ACCESS_TYPE_ACCOUNT_WHITELIST</td><td>3</td><td>账号白名单(指定UIN/手机/邮箱/IP可访问)</td></tr></tbody></table>
	AccessType *int64 `json:"AccessType,omitnil,omitempty" name:"AccessType"`

	// <p>是否开启访问控制</p><p>枚举值：</p><ul><li>true： 启用</li><li>false： 禁用</li></ul>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>白名单信息</p>
	Whitelist []*AppShareWhitelistItem `json:"Whitelist,omitnil,omitempty" name:"Whitelist"`
}

type AppShareURLInfo struct {
	// 当前生效的访问控制配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessControl *AppShareAccessControl `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`

	// 分享URL
	ShareUrl *string `json:"ShareUrl,omitnil,omitempty" name:"ShareUrl"`
}

type AppShareWhitelistItem struct {
	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_SHARE_WHITELIST_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>APP_SHARE_WHITELIST_TYPE_UIN</td><td>1</td><td>UIN账号</td></tr><tr><td>APP_SHARE_WHITELIST_TYPE_PHONE</td><td>2</td><td>手机号码</td></tr><tr><td>APP_SHARE_WHITELIST_TYPE_EMAIL</td><td>3</td><td>邮箱地址</td></tr><tr><td>APP_SHARE_WHITELIST_TYPE_IP</td><td>4</td><td>IP地址</td></tr><tr><td>APP_SHARE_WHITELIST_TYPE_RTX</td><td>5</td><td>RTX账号</td></tr></tbody></table>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>白名单数组信息</p><p>参数格式：白名单值</p>
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type AppSharedKbInfo struct {
	// 共享知识库ID
	KbId *string `json:"KbId,omitnil,omitempty" name:"KbId"`

	// 共享知识库名称
	KbName *string `json:"KbName,omitnil,omitempty" name:"KbName"`
}

type AppStatusInfo struct {
	// <p>应用状态</p><p>枚举值：</p><ul><li>1： 未上线</li><li>2： 运行中</li><li>3： 停用</li><li>4： 导入中</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>状态描述</p>
	StatusDescription *string `json:"StatusDescription,omitnil,omitempty" name:"StatusDescription"`
}

type AppSubStatusInfo struct {
	// 审批记录ID (当sub_status_list包含PUBLISH_APPROVING时有效)
	ApprovalId *string `json:"ApprovalId,omitnil,omitempty" name:"ApprovalId"`

	// 应用子状态列表 (可能同时处于多个子状态)
	SubStatusList []*int64 `json:"SubStatusList,omitnil,omitempty" name:"SubStatusList"`
}

type AppSummary struct {
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 应用模式。枚举值: 1:标准模式, 2:Agent模式, 3:单工作流模式, 4:ClawAgent模式
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// 应用头像
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 操作信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationInfo *AppOperation `json:"OperationInfo,omitnil,omitempty" name:"OperationInfo"`

	// 状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *AppStatusInfo `json:"Status,omitnil,omitempty" name:"Status"`

	// 子状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubStatus *AppSubStatusInfo `json:"SubStatus,omitnil,omitempty" name:"SubStatus"`

	// 资源操作权限
	PermissionIdList []*string `json:"PermissionIdList,omitnil,omitempty" name:"PermissionIdList"`
}

type AppToolConfig struct {
	// <p>输入参数</p>
	Inputs []*RequestParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// <p>输出参数</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`
}

type AppTrigger struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>执行配置</p>
	ExecuteConfig *ExecuteConfig `json:"ExecuteConfig,omitnil,omitempty" name:"ExecuteConfig"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td><td>1</td><td>指令执行</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td><td>2</td><td>工作流执行</td></tr></tbody></table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>失败次数</p>
	FailedCount *string `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// <p>推送渠道配置</p>
	PushConfig *TimerPushConfig `json:"PushConfig,omitnil,omitempty" name:"PushConfig"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_STATUS_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_STATUS_ENABLED</td><td>1</td><td>启用</td></tr><tr><td>APP_TRIGGER_STATUS_PAUSED</td><td>2</td><td>暂停</td></tr><tr><td>APP_TRIGGER_STATUS_DELETED</td><td>3</td><td>已删除</td></tr></tbody></table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>成功次数</p>
	SuccessCount *string `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// <p>触发器配置</p>
	TriggerConfig *TriggerConfig `json:"TriggerConfig,omitnil,omitempty" name:"TriggerConfig"`

	// <p>触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>触发器名称</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <p>触发器状态</p>
	TriggerStatus *TriggerStatus `json:"TriggerStatus,omitnil,omitempty" name:"TriggerStatus"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_TYPE_SCHEDULED</td><td>1</td><td>定时触发</td></tr><tr><td>APP_TRIGGER_TYPE_WEBHOOK</td><td>2</td><td>Webhook 触发</td></tr></tbody></table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type AppTriggerInstance struct {
	// <p>应用id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>会话id</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>触发器创建时间</p>
	CreatedAt *string `json:"CreatedAt,omitnil,omitempty" name:"CreatedAt"`

	// <p>结束时间</p>
	FinishedAt *string `json:"FinishedAt,omitnil,omitempty" name:"FinishedAt"`

	// <p>触发器运行实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <p>请求ID</p>
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`

	// <p>结果码</p>
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// <p>结果摘要</p>
	ResultSummary *string `json:"ResultSummary,omitnil,omitempty" name:"ResultSummary"`

	// <p>单次对话id</p>
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_INSTANCE_SOURCE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_INSTANCE_SOURCE_APP_TRIGGER</td><td>1</td><td>来源于应用触发器</td></tr></tbody></table>
	Source *int64 `json:"Source,omitnil,omitempty" name:"Source"`

	// <p>触发器开始执行时间</p>
	StartedAt *string `json:"StartedAt,omitnil,omitempty" name:"StartedAt"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>TIMER_RUN_STATUS_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>TIMER_RUN_STATUS_PENDING</td><td>1</td><td>等待执行</td></tr><tr><td>TIMER_RUN_STATUS_RUNNING</td><td>2</td><td>执行中</td></tr><tr><td>TIMER_RUN_STATUS_RETRY_WAIT</td><td>3</td><td>等待重试</td></tr><tr><td>TIMER_RUN_STATUS_SUCCESS</td><td>4</td><td>成功</td></tr><tr><td>TIMER_RUN_STATUS_DEAD</td><td>5</td><td>失败终态 (重试耗尽 / 不可重试)</td></tr><tr><td>TIMER_RUN_STATUS_CANCELLED</td><td>6</td><td>被任务暂停/删除/修改取消</td></tr></tbody></table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>TraceId，用于日志记录</p>
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// <p>触发器id</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>工作流运行ID</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type AppTriggerParamBinding struct {
	// <p>参数名字</p>
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// <p>参数类型</p><p>枚举值：</p><ul><li>0： 字符串</li><li>1： 整数</li><li>2： 浮点数</li><li>4： 对象</li><li>5： 字符串数组</li><li>6： 整数数组</li><li>7： 浮点数数组</li><li>8： 布尔值数组</li><li>3： 布尔值</li><li>9： 对象数组</li><li>10： 文件</li><li>11： 文档</li><li>12： 图片</li><li>13： 音频</li><li>14： 视频</li><li>15： 文件数组</li><li>16： 文档数组</li><li>17： 图片数组</li><li>18： 音频数组</li><li>19： 视频数组</li><li>20： 数组嵌套</li><li>22： 密钥</li><li>99： 空值</li><li>100： 未指定类型，用于OneOf和AnyOf场景</li></ul>
	ParamType *int64 `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// <p>参数值</p>
	Value *AppTriggerParamBindingValue `json:"Value,omitnil,omitempty" name:"Value"`
}

type AppTriggerParamBindingConfig struct {
	// <p>绑定参数列表</p>
	ParamList []*AppTriggerParamBinding `json:"ParamList,omitnil,omitempty" name:"ParamList"`
}

type AppTriggerParamBindingValue struct {
	// <p>参数值</p>
	ParamValue *string `json:"ParamValue,omitnil,omitempty" name:"ParamValue"`

	// <p>应用变量名</p>
	VariableName *string `json:"VariableName,omitnil,omitempty" name:"VariableName"`
}

type AppTriggerParamSchema struct {
	// <p>参数名</p>
	ParamName *string `json:"ParamName,omitnil,omitempty" name:"ParamName"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>PARAM_TYPE_STRING</td><td>0</td><td>字符串</td></tr><tr><td>PARAM_TYPE_INT</td><td>1</td><td>整数</td></tr><tr><td>PARAM_TYPE_FLOAT</td><td>2</td><td>浮点数</td></tr><tr><td>PARAM_TYPE_BOOL</td><td>3</td><td>布尔值</td></tr><tr><td>PARAM_TYPE_OBJECT</td><td>4</td><td>对象</td></tr><tr><td>PARAM_TYPE_ARRAY_STRING</td><td>5</td><td>字符串数组</td></tr><tr><td>PARAM_TYPE_ARRAY_INT</td><td>6</td><td>整数数组</td></tr><tr><td>PARAM_TYPE_ARRAY_FLOAT</td><td>7</td><td>浮点数数组</td></tr><tr><td>PARAM_TYPE_ARRAY_BOOL</td><td>8</td><td>布尔值数组</td></tr><tr><td>PARAM_TYPE_ARRAY_OBJECT</td><td>9</td><td>对象数组</td></tr><tr><td>PARAM_TYPE_ARRAY_ARRAY</td><td>20</td><td>数组嵌套</td></tr><tr><td>PARAM_TYPE_NULL</td><td>99</td><td>空值</td></tr><tr><td>PARAM_TYPE_UNSPECIFIED</td><td>100</td><td>未指定类型，用于OneOf和AnyOf场景</td></tr></tbody></table>
	ParamType *int64 `json:"ParamType,omitnil,omitempty" name:"ParamType"`

	// <p>是否必选</p>
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// <p>子参数列表</p>
	SubParamList []*AppTriggerParamSchema `json:"SubParamList,omitnil,omitempty" name:"SubParamList"`
}

type AppTriggerPromptExecuteConfig struct {
	// <p>触发器执行提示词</p>
	ExecutePrompt *string `json:"ExecutePrompt,omitnil,omitempty" name:"ExecutePrompt"`

	// <p>api参数绑定</p>
	ParamBindingsApi *AppTriggerParamBindingConfig `json:"ParamBindingsApi,omitnil,omitempty" name:"ParamBindingsApi"`
}

type AppTriggerRunLog struct {
	// <p>会话id</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>执行时间</p>
	DurationMs *string `json:"DurationMs,omitnil,omitempty" name:"DurationMs"`

	// <p>结束时间</p><p>参数格式：YYYY-MM-DD hh:mm:ss</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_FIRE_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_FIRE_TYPE_SCHEDULED</td><td>1</td><td>定时触发</td></tr><tr><td>APP_TRIGGER_FIRE_TYPE_WEBHOOK</td><td>2</td><td>Webhook 触发</td></tr><tr><td>APP_TRIGGER_FIRE_TYPE_MANUAL_RUN</td><td>3</td><td>手动立即执行</td></tr><tr><td>APP_TRIGGER_FIRE_TYPE_TEST_RUN</td><td>4</td><td>测试执行</td></tr></tbody></table>
	FireType *int64 `json:"FireType,omitnil,omitempty" name:"FireType"`

	// <p>触发实例id</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>TIMER_RUN_PUSH_STATUS_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>TIMER_RUN_PUSH_STATUS_NONE</td><td>1</td><td>未配置推送</td></tr><tr><td>TIMER_RUN_PUSH_STATUS_WAITING</td><td>2</td><td>等待推送</td></tr><tr><td>TIMER_RUN_PUSH_STATUS_SUCCESS</td><td>3</td><td>推送成功</td></tr><tr><td>TIMER_RUN_PUSH_STATUS_FAILED</td><td>4</td><td>推送失败</td></tr></tbody></table>
	PushStatus *int64 `json:"PushStatus,omitnil,omitempty" name:"PushStatus"`

	// <p>结果码</p>
	ResultCode *string `json:"ResultCode,omitnil,omitempty" name:"ResultCode"`

	// <p>结果概要</p>
	ResultSummary *string `json:"ResultSummary,omitnil,omitempty" name:"ResultSummary"`

	// <p>单次对话id</p>
	RunId *string `json:"RunId,omitnil,omitempty" name:"RunId"`

	// <p>触发时间</p><p>参数格式：YYYY:MM:DD hh:mm:ss</p>
	ScheduledFireTime *string `json:"ScheduledFireTime,omitnil,omitempty" name:"ScheduledFireTime"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>开始执行时间</p><p>参数格式：YYYY:MM:DD hh:mm:ss</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>TIMER_RUN_STATUS_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>TIMER_RUN_STATUS_PENDING</td><td>1</td><td>等待执行</td></tr><tr><td>TIMER_RUN_STATUS_RUNNING</td><td>2</td><td>执行中</td></tr><tr><td>TIMER_RUN_STATUS_RETRY_WAIT</td><td>3</td><td>等待重试</td></tr><tr><td>TIMER_RUN_STATUS_SUCCESS</td><td>4</td><td>成功</td></tr><tr><td>TIMER_RUN_STATUS_DEAD</td><td>5</td><td>失败终态 (重试耗尽 / 不可重试)</td></tr><tr><td>TIMER_RUN_STATUS_CANCELLED</td><td>6</td><td>被任务暂停/删除/修改取消</td></tr></tbody></table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>触发器id</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>是否已读</p>
	Unread *bool `json:"Unread,omitnil,omitempty" name:"Unread"`

	// <p>访客Id</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>工作流运行id</p>
	WorkflowRunId *string `json:"WorkflowRunId,omitnil,omitempty" name:"WorkflowRunId"`
}

type AppTriggerScheduleConfig struct {
	// <p>触发器定时配置</p>
	Schedule *TimerScheduleConfig `json:"Schedule,omitnil,omitempty" name:"Schedule"`
}

type AppTriggerScheduleStatus struct {
	// <p>最近一次触发时间</p><p>参数格式：格式为YYYY-MM-DD hh:mm:ss</p>
	LastFireTime *string `json:"LastFireTime,omitnil,omitempty" name:"LastFireTime"`

	// <p>下一次触发时间</p><p>参数格式：格式为YYYY-MM-DD hh:mm:ss</p>
	NextFireTime *string `json:"NextFireTime,omitnil,omitempty" name:"NextFireTime"`

	// <p>触发方式</p>
	PolicySummary *string `json:"PolicySummary,omitnil,omitempty" name:"PolicySummary"`
}

type AppTriggerSummary struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td><td>1</td><td>指令执行</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td><td>2</td><td>工作流执行</td></tr></tbody></table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>失败次数</p>
	FailedCount *string `json:"FailedCount,omitnil,omitempty" name:"FailedCount"`

	// <p>最近一次会话id</p>
	LastSessionId *string `json:"LastSessionId,omitnil,omitempty" name:"LastSessionId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table><p>取值范围：[0, 2]</p>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_STATUS_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_STATUS_ENABLED</td><td>1</td><td>启用</td></tr><tr><td>APP_TRIGGER_STATUS_PAUSED</td><td>2</td><td>暂停</td></tr><tr><td>APP_TRIGGER_STATUS_DELETED</td><td>3</td><td>已删除</td></tr></tbody></table>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>成功次数</p>
	SuccessCount *string `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// <p>触发器id</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>触发器名称</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <p>触发器执行状态</p>
	TriggerStatus *TriggerStatus `json:"TriggerStatus,omitnil,omitempty" name:"TriggerStatus"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_TYPE_SCHEDULED</td><td>1</td><td>定时触发</td></tr><tr><td>APP_TRIGGER_TYPE_WEBHOOK</td><td>2</td><td>Webhook 触发</td></tr></tbody></table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>未读日志的数量</p>
	UnreadRunLogCount *string `json:"UnreadRunLogCount,omitnil,omitempty" name:"UnreadRunLogCount"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type AppTriggerWebhookConfig struct {
	// <p>触发器webhook参数配置</p>
	ParamSchemaConfig *AppTriggerWebhookParamSchemaConfig `json:"ParamSchemaConfig,omitnil,omitempty" name:"ParamSchemaConfig"`

	// <p>webhook的key</p>
	WebhookKey *string `json:"WebhookKey,omitnil,omitempty" name:"WebhookKey"`

	// <p>webhook的密钥</p>
	WebhookToken *string `json:"WebhookToken,omitnil,omitempty" name:"WebhookToken"`

	// <p>webhook的地址</p>
	WebhookUrl *string `json:"WebhookUrl,omitnil,omitempty" name:"WebhookUrl"`
}

type AppTriggerWebhookParamSchemaConfig struct {
	// <p>触发器API参数列表</p>
	SchemaList []*AppTriggerParamSchema `json:"SchemaList,omitnil,omitempty" name:"SchemaList"`
}

type AppTriggerWebhookStatus struct {
	// <p>推送Webbook地址</p>
	WebhookUrl *string `json:"WebhookUrl,omitnil,omitempty" name:"WebhookUrl"`
}

type AppTriggerWorkflowExecuteConfig struct {
	// <p>工作流API参数绑定</p>
	ParamBindingsApi *AppTriggerParamBindingConfig `json:"ParamBindingsApi,omitnil,omitempty" name:"ParamBindingsApi"`

	// <p>工作流参数绑定</p>
	ParamBindingsWorkflow *AppTriggerParamBindingConfig `json:"ParamBindingsWorkflow,omitnil,omitempty" name:"ParamBindingsWorkflow"`

	// <p>工作流ID</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流名字</p>
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`
}

type AppWebSearchConfig struct {
	// API密钥
	ApiKey *string `json:"ApiKey,omitnil,omitempty" name:"ApiKey"`

	// 是否开启
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 服务提供商
	Provider *string `json:"Provider,omitnil,omitempty" name:"Provider"`

	// 返回结果数量
	TopN *uint64 `json:"TopN,omitnil,omitempty" name:"TopN"`
}

type AppWorkflowConfig struct {
	// 是否使用PDL
	EnablePDL *bool `json:"EnablePDL,omitnil,omitempty" name:"EnablePDL"`
}

type AppealingStatus struct {
	// 头像是否在申诉中
	AvatarInAppeal *bool `json:"AvatarInAppeal,omitnil,omitempty" name:"AvatarInAppeal"`

	// 兜底回复语是否在申诉中
	FallbackReplyInAppeal *bool `json:"FallbackReplyInAppeal,omitnil,omitempty" name:"FallbackReplyInAppeal"`

	// 欢迎语是否在申诉中
	GreetingInAppeal *bool `json:"GreetingInAppeal,omitnil,omitempty" name:"GreetingInAppeal"`

	// 应用名称是否在申诉中
	NameInAppeal *bool `json:"NameInAppeal,omitnil,omitempty" name:"NameInAppeal"`

	// 角色描述是否在申诉中
	RoleInAppeal *bool `json:"RoleInAppeal,omitnil,omitempty" name:"RoleInAppeal"`
}

type AuditLog struct {
	// <p>员工信息</p>
	AccountInfo *AccountInfo `json:"AccountInfo,omitnil,omitempty" name:"AccountInfo"`

	// <p>应用业务id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用名称</p><p>操作日志触发时的名称</p>
	AppName *string `json:"AppName,omitnil,omitempty" name:"AppName"`

	// <p>操作时间</p><p>参数格式：秒时间戳</p>
	OperateTime *string `json:"OperateTime,omitnil,omitempty" name:"OperateTime"`

	// <p>操作类型</p>
	Action *string `json:"Action,omitnil,omitempty" name:"Action"`

	// <p>操作对象</p>
	Biz *string `json:"Biz,omitnil,omitempty" name:"Biz"`

	// <p>操作内容</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>操作唯一ID</p>
	UniqueId *string `json:"UniqueId,omitnil,omitempty" name:"UniqueId"`
}

type AuditLogMetaField struct {
	// <p>操作日志元数据key</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>操作日志元数据Name</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AuthConfig struct {
	// <p>授权方式。</p><p>枚举值：</p><ul><li>0：无鉴权</li><li>1：API Key 鉴权</li><li>2：CAM 授权</li><li>3：OAuth 2.0 授权</li></ul>
	AuthType *int64 `json:"AuthType,omitnil,omitempty" name:"AuthType"`

	// API Key授权配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiKeyAuthConfig *ApiKeyAuthConfig `json:"ApiKeyAuthConfig,omitnil,omitempty" name:"ApiKeyAuthConfig"`

	// CAM授权配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CamAuthConfig *CamAuthConfig `json:"CamAuthConfig,omitnil,omitempty" name:"CamAuthConfig"`

	// OAuth2.0授权配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	OAuthConfig *OAuthConfig `json:"OAuthConfig,omitnil,omitempty" name:"OAuthConfig"`
}

type BackgroundImage struct {
	// 亮度值
	Brightness *uint64 `json:"Brightness,omitnil,omitempty" name:"Brightness"`

	// 横图(pc)
	LandscapeImageUrl *string `json:"LandscapeImageUrl,omitnil,omitempty" name:"LandscapeImageUrl"`

	// 原始图
	OriginalImageUrl *string `json:"OriginalImageUrl,omitnil,omitempty" name:"OriginalImageUrl"`

	// 长图(手机)
	PortraitImageUrl *string `json:"PortraitImageUrl,omitnil,omitempty" name:"PortraitImageUrl"`

	// 主题色
	ThemeColor *string `json:"ThemeColor,omitnil,omitempty" name:"ThemeColor"`
}

type BasicBilling struct {
	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>UNKNOW</td><td>0</td><td></td></tr><tr><td>TOKEN</td><td>1</td><td>按token</td></tr><tr><td>PAGE_COUNT</td><td>2</td><td>按页数</td></tr><tr><td>TIMES</td><td>3</td><td>按次数</td></tr><tr><td>TIMES_THOUSAND</td><td>4</td><td>按千次数</td></tr><tr><td>SECOND</td><td>5</td><td>按时长</td></tr><tr><td>CHARACTER</td><td>6</td><td>按字符数</td></tr><tr><td>CHARACTER_THOUSAND</td><td>7</td><td>按千字符数</td></tr><tr><td>SHEET</td><td>8</td><td>按张</td></tr><tr><td>NUMBER</td><td>9</td><td>按个数</td></tr></tbody></table>
	BillingUnit *int64 `json:"BillingUnit,omitnil,omitempty" name:"BillingUnit"`

	// <p>现金价格</p><p>单位：元</p>
	CashPrice *float64 `json:"CashPrice,omitnil,omitempty" name:"CashPrice"`

	// <p>PU价格</p><p>单位：pu</p>
	PuPrice *float64 `json:"PuPrice,omitnil,omitempty" name:"PuPrice"`
}

type BillingAttribute struct {
	// <p>属性名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>属性值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type CallSource struct {
	// <p>调用主体 ID，含义由 subject_type 决定（如 app_id、kb_id 等）</p>
	SubjectId *string `json:"SubjectId,omitnil,omitempty" name:"SubjectId"`

	// <p>调用主体名称</p>
	SubjectName *string `json:"SubjectName,omitnil,omitempty" name:"SubjectName"`

	// <p>调用主体类型：APP/KB/WIDGET/OPEN_CLAW/KB_RECALL_TEST/WORKBENCH/MODEL_API</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>METRIC_SOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>METRIC_SOURCE_TYPE_APP</td><td>1</td><td>应用开发</td></tr><tr><td>METRIC_SOURCE_TYPE_KB</td><td>2</td><td>知识库</td></tr><tr><td>METRIC_SOURCE_TYPE_WIDGET</td><td>3</td><td>Widget</td></tr><tr><td>METRIC_SOURCE_TYPE_OPEN_CLAW</td><td>4</td><td>ClawPro</td></tr><tr><td>METRIC_SOURCE_TYPE_KB_RECALL_TEST</td><td>5</td><td>知识库召回测试</td></tr><tr><td>METRIC_SOURCE_TYPE_WORKBENCH</td><td>6</td><td>智能工作台</td></tr><tr><td>METRIC_SOURCE_TYPE_MODEL_API</td><td>7</td><td>模型 API 调用</td></tr></tbody></table>
	SubjectType *int64 `json:"SubjectType,omitnil,omitempty" name:"SubjectType"`
}

type CamAuthConfig struct {
	// 角色名称
	RoleName *string `json:"RoleName,omitnil,omitempty" name:"RoleName"`

	// 密钥位置 HEADER/QUERY
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 头鉴权 |
	// | 1 | 请求信息鉴权 |
	KeyLocation *int64 `json:"KeyLocation,omitnil,omitempty" name:"KeyLocation"`

	// SecretId字段名称
	SecretIdName *string `json:"SecretIdName,omitnil,omitempty" name:"SecretIdName"`

	// SecretKey字段名称
	SecretKeyName *string `json:"SecretKeyName,omitnil,omitempty" name:"SecretKeyName"`
}

type CategoryPermission struct {
	// <p>当前用户是否可新增子分类</p>
	CanAdd *bool `json:"CanAdd,omitnil,omitempty" name:"CanAdd"`

	// <p>当前用户是否可删除该分类</p>
	CanDelete *bool `json:"CanDelete,omitnil,omitempty" name:"CanDelete"`

	// <p>当前用户是否可编辑该分类</p>
	CanEdit *bool `json:"CanEdit,omitnil,omitempty" name:"CanEdit"`
}

type ClawAgentAgentTeamConfig struct {
	// <p>是否开启Agent团队协作</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>prompt内容</p>
	PromptContent *string `json:"PromptContent,omitnil,omitempty" name:"PromptContent"`
}

type ClawAgentConfig struct {
	// Agent团队协作配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentTeamConfig *ClawAgentAgentTeamConfig `json:"AgentTeamConfig,omitnil,omitempty" name:"AgentTeamConfig"`

	// 长期记忆配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	LongMemoryConfig *ClawAgentLongMemoryConfig `json:"LongMemoryConfig,omitnil,omitempty" name:"LongMemoryConfig"`
}

type ClawAgentLongMemoryConfig struct {
	// <p>是否开启长期记忆</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type CodeToolConfig struct {
	// <p>代码</p>
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>示例</p>
	Example *ToolExample `json:"Example,omitnil,omitempty" name:"Example"`

	// <p>输入参数</p>
	Inputs []*RequestParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// <p>输出参数</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`
}

type ComplexBilling struct {
	// <p>复合计费列表</p>
	ComplexList []*ComplexBillingItem `json:"ComplexList,omitnil,omitempty" name:"ComplexList"`
}

type ComplexBillingItem struct {
	// <p>复合计费维度信息</p>
	BillingAttributeList []*BillingAttribute `json:"BillingAttributeList,omitnil,omitempty" name:"BillingAttributeList"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>UNKNOW</td><td>0</td><td></td></tr><tr><td>TOKEN</td><td>1</td><td>按token</td></tr><tr><td>PAGE_COUNT</td><td>2</td><td>按页数</td></tr><tr><td>TIMES</td><td>3</td><td>按次数</td></tr><tr><td>TIMES_THOUSAND</td><td>4</td><td>按千次数</td></tr><tr><td>SECOND</td><td>5</td><td>按时长</td></tr><tr><td>CHARACTER</td><td>6</td><td>按字符数</td></tr><tr><td>CHARACTER_THOUSAND</td><td>7</td><td>按千字符数</td></tr><tr><td>SHEET</td><td>8</td><td>按张</td></tr><tr><td>NUMBER</td><td>9</td><td>按个数</td></tr></tbody></table>
	BillingUnit *int64 `json:"BillingUnit,omitnil,omitempty" name:"BillingUnit"`

	// <p>现金价格</p><p>单位：元</p>
	CashPrice *float64 `json:"CashPrice,omitnil,omitempty" name:"CashPrice"`

	// <p>pu价格</p><p>单位：pu</p>
	PuPrice *float64 `json:"PuPrice,omitnil,omitempty" name:"PuPrice"`
}

type ConcurrencyLimitDetail struct {
	// <p>调用来源（subject_type 决定 subject_id/subject_name 的含义，如 APP 时 subject_id=app_id、subject_name=app_name）</p>
	CallSource *CallSource `json:"CallSource,omitnil,omitempty" name:"CallSource"`

	// <p>超限发生时间（Unix秒）</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>模型名称</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>请求内容（用户请求的原始查询文本）</p>
	RequestQuery *string `json:"RequestQuery,omitnil,omitempty" name:"RequestQuery"`

	// <p>空间 ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type ConsumptionClassification struct {
	// <p>消耗场景（如推理/训练/评测等）</p>
	ConsumptionScene *string `json:"ConsumptionScene,omitnil,omitempty" name:"ConsumptionScene"`

	// <p>消耗目标（如具体模型名/插件名/平台功能名）</p>
	ConsumptionTarget *string `json:"ConsumptionTarget,omitnil,omitempty" name:"ConsumptionTarget"`

	// <p>消耗类型，取值集合由业务方定义（如 model/plugin/platform 等）</p>
	ConsumptionType *string `json:"ConsumptionType,omitnil,omitempty" name:"ConsumptionType"`

	// <p>套餐包名称</p>
	PackageName *string `json:"PackageName,omitnil,omitempty" name:"PackageName"`
}

type ConsumptionDetail struct {
	// <p>消耗分类（类型/目标/场景/套餐包）</p>
	Classification *ConsumptionClassification `json:"Classification,omitnil,omitempty" name:"Classification"`

	// <p>消耗发生时间，Unix 秒</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>用量来源类型</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>METRIC_SOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>METRIC_SOURCE_TYPE_APP</td><td>1</td><td>应用开发</td></tr><tr><td>METRIC_SOURCE_TYPE_KB</td><td>2</td><td>知识库</td></tr><tr><td>METRIC_SOURCE_TYPE_WIDGET</td><td>3</td><td>Widget</td></tr><tr><td>METRIC_SOURCE_TYPE_OPEN_CLAW</td><td>4</td><td>ClawPro</td></tr><tr><td>METRIC_SOURCE_TYPE_KB_RECALL_TEST</td><td>5</td><td>知识库召回测试</td></tr><tr><td>METRIC_SOURCE_TYPE_WORKBENCH</td><td>6</td><td>智能工作台</td></tr><tr><td>METRIC_SOURCE_TYPE_MODEL_API</td><td>7</td><td>模型 API 调用</td></tr></tbody></table>
	MetricSourceType *int64 `json:"MetricSourceType,omitnil,omitempty" name:"MetricSourceType"`

	// <p>名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>空间名称</p>
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// <p>消耗用量（数值/单位/PU 消耗）</p>
	Usage *ConsumptionUsage `json:"Usage,omitnil,omitempty" name:"Usage"`

	// <p>用户名称</p>
	UserName *string `json:"UserName,omitnil,omitempty" name:"UserName"`
}

type ConsumptionUsage struct {
	// <p>消耗PU</p>
	ConsumptionPU *float64 `json:"ConsumptionPU,omitnil,omitempty" name:"ConsumptionPU"`

	// <p>用量数值</p>
	Usage *float64 `json:"Usage,omitnil,omitempty" name:"Usage"`

	// <p>用量单位，枚举值 DosageUnit</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>DOSAGE_UNIT_TOKEN</td><td>0</td><td>token（默认）</td></tr><tr><td>DOSAGE_UNIT_PAGE_COUNT</td><td>1</td><td>page_count（页数）</td></tr><tr><td>DOSAGE_UNIT_TIMES</td><td>2</td><td>times（次数）</td></tr><tr><td>DOSAGE_UNIT_SECOND</td><td>3</td><td>second（秒）</td></tr><tr><td>DOSAGE_UNIT_ITEM</td><td>4</td><td>item（条）</td></tr><tr><td>DOSAGE_UNIT_SHEET</td><td>5</td><td>sheet（张）</td></tr><tr><td>DOSAGE_UNIT_CHARACTER</td><td>6</td><td>character（字符）</td></tr><tr><td>DOSAGE_UNIT_GB</td><td>7</td><td>GB</td></tr><tr><td>DOSAGE_UNIT_NUMBER</td><td>8</td><td>number（个数）</td></tr><tr><td>DOSAGE_UNIT_MILL_SECOND</td><td>9</td><td>mill_second（毫秒）</td></tr></tbody></table>
	UsageUnit *int64 `json:"UsageUnit,omitnil,omitempty" name:"UsageUnit"`
}

type Conversation struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>会话标题</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>会话使用的用户端 AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type ConversationAgentTask struct {
	// <p>任务内容</p>
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// <p>任务序号</p>
	Index *string `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>任务状态，pending:待执行，processing:处理中，success:已完成，failed:处理失败，stop:已取消</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConversationContent struct {
	// <p>文本内容</p>
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// <p>内容类型, text：文本,image：图片,file：文件,custom_variables：自定义输入参数信息,widget_action：widget动作信息</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>自定义参数数据</p>
	CustomParamList []*string `json:"CustomParamList,omitnil,omitempty" name:"CustomParamList"`

	// <p>自定义参数数据</p>
	CustomParams []*string `json:"CustomParams,omitnil,omitempty" name:"CustomParams"`

	// <p>自定义变量数据</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomVariablesData *string `json:"CustomVariablesData,omitnil,omitempty" name:"CustomVariablesData"`

	// <p>企业表单</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnterpriseCharts *string `json:"EnterpriseCharts,omitnil,omitempty" name:"EnterpriseCharts"`

	// <p>选项卡列表</p>
	OptionCardList []*string `json:"OptionCardList,omitnil,omitempty" name:"OptionCardList"`

	// <p>选项卡列表</p>
	OptionCards []*string `json:"OptionCards,omitnil,omitempty" name:"OptionCards"`

	// <p>选项卡模式 枚举值: 0-OPTION_MODE_SINGLE(单选), 1-OPTION_MODE_MULTI(多选)</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	OptionMode *int64 `json:"OptionMode,omitnil,omitempty" name:"OptionMode"`

	// <p>引用角标信息列表</p>
	QuoteInfoList []*ConversationQuoteInfo `json:"QuoteInfoList,omitnil,omitempty" name:"QuoteInfoList"`

	// <p>引用角标信息列表</p>
	QuoteInfos []*ConversationQuoteInfo `json:"QuoteInfos,omitnil,omitempty" name:"QuoteInfos"`

	// <p>参考来源列表</p>
	ReferenceList []*ConversationReference `json:"ReferenceList,omitnil,omitempty" name:"ReferenceList"`

	// <p>参考来源列表</p>
	References []*ConversationReference `json:"References,omitnil,omitempty" name:"References"`

	// <p>关联记录 ID</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RelatedRecordId *string `json:"RelatedRecordId,omitnil,omitempty" name:"RelatedRecordId"`

	// <p>智能体任务列表</p>
	TaskList []*ConversationAgentTask `json:"TaskList,omitnil,omitempty" name:"TaskList"`

	// <p>智能体任务列表</p>
	Tasks []*ConversationAgentTask `json:"Tasks,omitnil,omitempty" name:"Tasks"`

	// <p>工作流输入参数</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WorkflowInput *string `json:"WorkflowInput,omitnil,omitempty" name:"WorkflowInput"`
}

type ConversationExperience struct {
	// AI通话配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiCall *AICallConfig `json:"AiCall,omitnil,omitempty" name:"AiCall"`

	// 背景图片配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BackgroundImage *BackgroundImage `json:"BackgroundImage,omitnil,omitempty" name:"BackgroundImage"`

	// 兜底回复开关
	EnableFallbackReply *bool `json:"EnableFallbackReply,omitnil,omitempty" name:"EnableFallbackReply"`

	// 是否使用推荐问
	EnableRecommended *bool `json:"EnableRecommended,omitnil,omitempty" name:"EnableRecommended"`

	// 是否使用联网搜索
	EnableWebSearch *bool `json:"EnableWebSearch,omitnil,omitempty" name:"EnableWebSearch"`

	// 兜底回复语
	FallbackReply *string `json:"FallbackReply,omitnil,omitempty" name:"FallbackReply"`

	// 输入框配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	InputBoxConfig *InputBoxConfig `json:"InputBoxConfig,omitnil,omitempty" name:"InputBoxConfig"`

	// 输出方式。枚举值: 1:流式, 2:非流式
	Method *int64 `json:"Method,omitnil,omitempty" name:"Method"`

	// 推荐问生成prompt模式。枚举值: 1:仅结合知识库输出推荐问的prompt
	RecommendPromptMode *int64 `json:"RecommendPromptMode,omitnil,omitempty" name:"RecommendPromptMode"`
}

type ConversationMessage struct {
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>消息图标</p>
	Icon *string `json:"Icon,omitnil,omitempty" name:"Icon"`

	// <p>消息 ID</p>
	MessageId *string `json:"MessageId,omitnil,omitempty" name:"MessageId"`

	// <p>消息名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>记录 ID</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>消息角色</p>
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// <p>消息状态，pending:待执行，processing:处理中，success:已完成，failed:处理失败，stop:已取消</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>状态描述</p>
	StatusDesc *string `json:"StatusDesc,omitnil,omitempty" name:"StatusDesc"`

	// <p>消息标题</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>消息内容列表</p>
	ContentList []*ConversationContent `json:"ContentList,omitnil,omitempty" name:"ContentList"`

	// <p>消息内容列表</p>
	Contents []*ConversationContent `json:"Contents,omitnil,omitempty" name:"Contents"`

	// <p>类型</p>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type ConversationQuoteInfo struct {
	// <p>参考来源的索引值</p>
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>参考来源位置</p>
	Position *int64 `json:"Position,omitnil,omitempty" name:"Position"`
}

type ConversationRecordErrorInfo struct {
	// <p>对话失败错误码</p>
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// <p>对话失败错误信息</p>
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`
}

type ConversationRecordSummary struct {
	// <p>回复记录 ID，对应 messages 中回复消息的 record_id</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>用户提问记录 ID，对应 messages 中用户消息的 record_id</p>
	RelatedRecordId *string `json:"RelatedRecordId,omitnil,omitempty" name:"RelatedRecordId"`

	// <p>单次对话耗时信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeUsage *ConversationRecordTimeUsage `json:"TimeUsage,omitnil,omitempty" name:"TimeUsage"`

	// <p>单次对话 token 消耗信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TokenUsage *ConversationRecordTokenUsage `json:"TokenUsage,omitnil,omitempty" name:"TokenUsage"`

	// <p>单次对话失败信息；成功时为空</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorInfo *ConversationRecordErrorInfo `json:"ErrorInfo,omitnil,omitempty" name:"ErrorInfo"`

	// <p>单次员工助理对话当前状态</p><p>枚举值：</p><ul><li>pending： 待处理</li><li>processing： 处理中</li><li>success： 成功</li><li>failed： 失败</li><li>stop： 停止</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ConversationRecordTimeUsage struct {
	// <p>单次对话总耗时，单位毫秒</p>
	Elapsed *string `json:"Elapsed,omitnil,omitempty" name:"Elapsed"`

	// <p>首 token 耗时，单位毫秒</p>
	FirstTokenCost *string `json:"FirstTokenCost,omitnil,omitempty" name:"FirstTokenCost"`

	// <p>模型推理总耗时，单位毫秒</p>
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`
}

type ConversationRecordTokenUsage struct {
	// <p>输入 token 总数</p>
	InputTokens *string `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// <p>输出 token 总数</p>
	OutputTokens *string `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// <p>消耗 token 总数</p>
	TotalTokens *string `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// <p>缓存命中 token 总数</p>
	CachedTokens *string `json:"CachedTokens,omitnil,omitempty" name:"CachedTokens"`

	// <p>推理 token 总数</p>
	ReasoningTokens *string `json:"ReasoningTokens,omitnil,omitempty" name:"ReasoningTokens"`
}

type ConversationReference struct {
	// <p>参考来源索引</p>
	Index *int64 `json:"Index,omitnil,omitempty" name:"Index"`

	// <p>参考来源名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>参考来源类型 枚举值: 0-APP_REFERENCE_TYPE_UNSPECIFIED(未指定), 1-APP_REFERENCE_TYPE_QA(问答), 2-APP_REFERENCE_TYPE_SEGMENT(分片), 3-APP_REFERENCE_TYPE_DOC(文档), 4-APP_REFERENCE_TYPE_WEB_SEARCH(Web 搜索), 5-APP_REFERENCE_TYPE_GRAPH_RAG(GraphRAG)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ConversationResetInfo struct {
	// <p>最近一次重置的毫秒级时间戳</p>
	ResetTime *string `json:"ResetTime,omitnil,omitempty" name:"ResetTime"`

	// <p>最近一次重置边界；该记录及更早的记录不再作为对话上下文</p>
	ResetThroughRecordId *string `json:"ResetThroughRecordId,omitnil,omitempty" name:"ResetThroughRecordId"`
}

type ConversationWorkspace struct {
	// <p>工作空间 ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>存储类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`
}

// Predefined struct for user
type CopyAgentFromAppRequestParams struct {
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>目标应用ID，kind=0时需传入</p>
	TargetAppId *string `json:"TargetAppId,omitnil,omitempty" name:"TargetAppId"`

	// <p>Agent 类型，区分 B 端配置态 Agent 与 C 端用户态 Agent</p><p>枚举值：</p><ul><li>0：  配置端Agent </li><li>1：  用户态 Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
}

type CopyAgentFromAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>目标应用ID，kind=0时需传入</p>
	TargetAppId *string `json:"TargetAppId,omitnil,omitempty" name:"TargetAppId"`

	// <p>Agent 类型，区分 B 端配置态 Agent 与 C 端用户态 Agent</p><p>枚举值：</p><ul><li>0：  配置端Agent </li><li>1：  用户态 Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
}

func (r *CopyAgentFromAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAgentFromAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TargetAppId")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAgentFromAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAgentFromAppResponseParams struct {
	// <p>主 Agent Id</p>
	ParentAgentId *string `json:"ParentAgentId,omitnil,omitempty" name:"ParentAgentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyAgentFromAppResponse struct {
	*tchttp.BaseResponse
	Response *CopyAgentFromAppResponseParams `json:"Response"`
}

func (r *CopyAgentFromAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAgentFromAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAppRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// target_space_id
	TargetSpaceId *string `json:"TargetSpaceId,omitnil,omitempty" name:"TargetSpaceId"`
}

type CopyAppRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// target_space_id
	TargetSpaceId *string `json:"TargetSpaceId,omitnil,omitempty" name:"TargetSpaceId"`
}

func (r *CopyAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "TargetSpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CopyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CopyAppResponseParams struct {
	// new_app_id
	NewAppId *string `json:"NewAppId,omitnil,omitempty" name:"NewAppId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CopyAppResponse struct {
	*tchttp.BaseResponse
	Response *CopyAppResponseParams `json:"Response"`
}

func (r *CopyAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CopyAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CorpShareConfig struct {
	// <p>企业共享开关</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>共享范围类型，1：企业全员，2：指定账户，3：指定空间</p>
	ShareScope *int64 `json:"ShareScope,omitnil,omitempty" name:"ShareScope"`

	// <p>企业共享应用标签</p>
	TagIdList []*string `json:"TagIdList,omitnil,omitempty" name:"TagIdList"`

	// <p>共享范围信息(用户时StrId为uin,Name为用户名称;空间时StrId为空间ID,Name为空间名称)</p>
	ShareScopeList []*Identity `json:"ShareScopeList,omitnil,omitempty" name:"ShareScopeList"`
}

// Predefined struct for user
type CreateAgentRequestParams struct {
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent 配置</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>Agent 类型，区分 B 端配置态 Agent 与 C 端用户态 Agent</p><p>枚举值：</p><ul><li>0： 配置端Agent</li><li>1： 用户态 Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
}

type CreateAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent 配置</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>Agent 类型，区分 B 端配置态 Agent 与 C 端用户态 Agent</p><p>枚举值：</p><ul><li>0： 配置端Agent</li><li>1： 用户态 Agent</li></ul>
	Kind *int64 `json:"Kind,omitnil,omitempty" name:"Kind"`
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
	delete(f, "AppId")
	delete(f, "Agent")
	delete(f, "Kind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgentResponseParams struct {
	// <p>Agent Id</p>
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
	// 空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 应用模式。枚举值: 1:标准模式, 2:Agent模式, 3:单工作流模式, 4:ClawAgent模式
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// 应用头像
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 应用描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 应用模式。枚举值: 1:标准模式, 2:Agent模式, 3:单工作流模式, 4:ClawAgent模式
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// 应用头像
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 应用描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 应用名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	delete(f, "SpaceId")
	delete(f, "AppMode")
	delete(f, "Avatar")
	delete(f, "Description")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

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
type CreateAppTriggerRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用触发器执行配置</p>
	ExecuteConfig *ExecuteConfig `json:"ExecuteConfig,omitnil,omitempty" name:"ExecuteConfig"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td><td>1</td><td>指令执行</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td><td>2</td><td>工作流执行</td></tr></tbody></table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>第三方推送配置</p>
	PushConfig *TimerPushConfig `json:"PushConfig,omitnil,omitempty" name:"PushConfig"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>触发器配置</p>
	TriggerConfig *TriggerConfig `json:"TriggerConfig,omitnil,omitempty" name:"TriggerConfig"`

	// <p>触发器名字</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_TYPE_SCHEDULED</td><td>1</td><td>定时触发</td></tr><tr><td>APP_TRIGGER_TYPE_WEBHOOK</td><td>2</td><td>Webhook 触发</td></tr></tbody></table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用触发器执行配置</p>
	ExecuteConfig *ExecuteConfig `json:"ExecuteConfig,omitnil,omitempty" name:"ExecuteConfig"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_PROMPT</td><td>1</td><td>指令执行</td></tr><tr><td>APP_TRIGGER_EXECUTE_TYPE_WORKFLOW</td><td>2</td><td>工作流执行</td></tr></tbody></table>
	ExecuteType *int64 `json:"ExecuteType,omitnil,omitempty" name:"ExecuteType"`

	// <p>第三方推送配置</p>
	PushConfig *TimerPushConfig `json:"PushConfig,omitnil,omitempty" name:"PushConfig"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>触发器配置</p>
	TriggerConfig *TriggerConfig `json:"TriggerConfig,omitnil,omitempty" name:"TriggerConfig"`

	// <p>触发器名字</p>
	TriggerName *string `json:"TriggerName,omitnil,omitempty" name:"TriggerName"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_TYPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_TYPE_SCHEDULED</td><td>1</td><td>定时触发</td></tr><tr><td>APP_TRIGGER_TYPE_WEBHOOK</td><td>2</td><td>Webhook 触发</td></tr></tbody></table>
	TriggerType *int64 `json:"TriggerType,omitnil,omitempty" name:"TriggerType"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ExecuteConfig")
	delete(f, "ExecuteType")
	delete(f, "PushConfig")
	delete(f, "Scope")
	delete(f, "TriggerConfig")
	delete(f, "TriggerName")
	delete(f, "TriggerType")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppTriggerResponseParams struct {
	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppTriggerResponseParams `json:"Response"`
}

func (r *CreateAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConversationRequestParams struct {
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户端 AgnetId，当Claw模式开启了“允许在对话中动态修改配置”时可用</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type CreateConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户端 AgnetId，当Claw模式开启了“允许在对话中动态修改配置”时可用</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *CreateConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConversationResponseParams struct {
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConversationResponse struct {
	*tchttp.BaseResponse
	Response *CreateConversationResponseParams `json:"Response"`
}

func (r *CreateConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMsgRecordCategoryRequestParams struct {
	// <p>分类名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>父分类业务 ID，0 表示一级分类（未分类）</p>
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

type CreateMsgRecordCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>分类名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>父分类业务 ID，0 表示一级分类（未分类）</p>
	ParentId *string `json:"ParentId,omitnil,omitempty" name:"ParentId"`
}

func (r *CreateMsgRecordCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMsgRecordCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "AppId")
	delete(f, "ParentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateMsgRecordCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateMsgRecordCategoryResponseParams struct {
	// <p>新建分类的业务 ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateMsgRecordCategoryResponse struct {
	*tchttp.BaseResponse
	Response *CreateMsgRecordCategoryResponseParams `json:"Response"`
}

func (r *CreateMsgRecordCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateMsgRecordCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginRequestParams struct {
	// <p>插件基础资料</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>插件类型配置</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>插件的工具列表</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type CreatePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>插件基础资料</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>插件类型配置</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>插件的工具列表</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *CreatePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Profile")
	delete(f, "Config")
	delete(f, "SpaceId")
	delete(f, "ToolList")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreatePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreatePluginResponseParams struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreatePluginResponse struct {
	*tchttp.BaseResponse
	Response *CreatePluginResponseParams `json:"Response"`
}

func (r *CreatePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreatePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用分享访问控制配置</p>
	AppShareAccessControl *AppShareAccessControl `json:"AppShareAccessControl,omitnil,omitempty" name:"AppShareAccessControl"`

	// <p>渠道ID列表</p>
	ChannelIdList []*string `json:"ChannelIdList,omitnil,omitempty" name:"ChannelIdList"`

	// <p>企业共享配置</p>
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`

	// <p>发布描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>将默认知识库中，仅调试生效的知识批量变更为&quot;调试/发布都生效&quot;</p>
	IsDevToRelease *bool `json:"IsDevToRelease,omitnil,omitempty" name:"IsDevToRelease"`

	// <p>是否同步发布为应用模板</p>
	IsPublishAsTemplate *bool `json:"IsPublishAsTemplate,omitnil,omitempty" name:"IsPublishAsTemplate"`
}

type CreateReleaseRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用分享访问控制配置</p>
	AppShareAccessControl *AppShareAccessControl `json:"AppShareAccessControl,omitnil,omitempty" name:"AppShareAccessControl"`

	// <p>渠道ID列表</p>
	ChannelIdList []*string `json:"ChannelIdList,omitnil,omitempty" name:"ChannelIdList"`

	// <p>企业共享配置</p>
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`

	// <p>发布描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>将默认知识库中，仅调试生效的知识批量变更为&quot;调试/发布都生效&quot;</p>
	IsDevToRelease *bool `json:"IsDevToRelease,omitnil,omitempty" name:"IsDevToRelease"`

	// <p>是否同步发布为应用模板</p>
	IsPublishAsTemplate *bool `json:"IsPublishAsTemplate,omitnil,omitempty" name:"IsPublishAsTemplate"`
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
	delete(f, "AppId")
	delete(f, "AppShareAccessControl")
	delete(f, "ChannelIdList")
	delete(f, "CorpShareConfig")
	delete(f, "Description")
	delete(f, "IsDevToRelease")
	delete(f, "IsPublishAsTemplate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateReleaseResponseParams struct {
	// <p>need_approval</p>
	NeedApproval *bool `json:"NeedApproval,omitnil,omitempty" name:"NeedApproval"`

	// <p>release_id</p>
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`

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
type CreateSkillRequestParams struct {
	// <p>Skill 创建方式，必填；仅允许</p><p>枚举值：</p><ul><li>1： FILE_UPLOAD（文件上传）</li><li>3： AIGC（AIGC生成）</li></ul>
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// <p>skill包文件地址（zip）；FILE_UPLOAD / AIGC 均必填</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill展示描述</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill展示名称</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>图标地址</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>skill业务唯一标识名（同企业下唯一）；未传时从skill包解析</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>版本号</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>版本变更说明</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

type CreateSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>Skill 创建方式，必填；仅允许</p><p>枚举值：</p><ul><li>1： FILE_UPLOAD（文件上传）</li><li>3： AIGC（AIGC生成）</li></ul>
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// <p>skill包文件地址（zip）；FILE_UPLOAD / AIGC 均必填</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill展示描述</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill展示名称</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>图标地址</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>skill业务唯一标识名（同企业下唯一）；未传时从skill包解析</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>版本号</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>版本变更说明</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

func (r *CreateSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "CreateType")
	delete(f, "FileUrl")
	delete(f, "SpaceId")
	delete(f, "DisplayDescription")
	delete(f, "DisplayName")
	delete(f, "IconUrl")
	delete(f, "Name")
	delete(f, "SkillVersion")
	delete(f, "UpdateDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillResponseParams struct {
	// <p>创建成功后的skillID</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>创建成功后的版本ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSkillResponse struct {
	*tchttp.BaseResponse
	Response *CreateSkillResponseParams `json:"Response"`
}

func (r *CreateSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillShareRequestParams struct {
	// <p>必填，申请备注（弹窗&quot;申请备注&quot;）</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>必填，原skill_id</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>必填，被共享的版本id（必须高于已共享版本）</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type CreateSkillShareRequest struct {
	*tchttp.BaseRequest
	
	// <p>必填，申请备注（弹窗&quot;申请备注&quot;）</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>必填，原skill_id</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>必填，被共享的版本id（必须高于已共享版本）</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *CreateSkillShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillShareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyRemark")
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSkillShareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSkillShareResponseParams struct {
	// <p>是否走了审批流（false表示无需审批已直接创建共享任务）</p>
	NeedApproval *bool `json:"NeedApproval,omitnil,omitempty" name:"NeedApproval"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSkillShareResponse struct {
	*tchttp.BaseResponse
	Response *CreateSkillShareResponseParams `json:"Response"`
}

func (r *CreateSkillShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSkillShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpaceRequestParams struct {
	// 工作空间名称,长度最大30个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 空间描述，长度最大150个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间名称,长度最大30个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 空间描述，长度最大150个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateSpaceResponseParams struct {
	// 空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateSpaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateSpaceResponseParams `json:"Response"`
}

func (r *CreateSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 变量信息
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

type CreateVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 变量信息
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

func (r *CreateVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Variable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateVariableResponseParams struct {
	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateVariableResponse struct {
	*tchttp.BaseResponse
	Response *CreateVariableResponseParams `json:"Response"`
}

func (r *CreateVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebSocketTokenRequestParams struct {
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateWebSocketTokenRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateWebSocketTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebSocketTokenRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWebSocketTokenRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWebSocketTokenResponseParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>WebSocket Token</p>
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWebSocketTokenResponse struct {
	*tchttp.BaseResponse
	Response *CreateWebSocketTokenResponseParams `json:"Response"`
}

func (r *CreateWebSocketTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWebSocketTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceCredentialRequestParams struct {
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>工作空间 ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type CreateWorkspaceCredentialRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>工作空间 ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *CreateWorkspaceCredentialRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceCredentialRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "WorkspaceId")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWorkspaceCredentialRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWorkspaceCredentialResponseParams struct {
	// <p>存储类型</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// <p>工作空间 ID</p>
	WorkspaceId *string `json:"WorkspaceId,omitnil,omitempty" name:"WorkspaceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWorkspaceCredentialResponse struct {
	*tchttp.BaseResponse
	Response *CreateWorkspaceCredentialResponseParams `json:"Response"`
}

func (r *CreateWorkspaceCredentialResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWorkspaceCredentialResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CronSchedule struct {
	// cron表达式
	Expression *string `json:"Expression,omitnil,omitempty" name:"Expression"`
}

type DailySchedule struct {
	// 时间
	TimeOfDay *string `json:"TimeOfDay,omitnil,omitempty" name:"TimeOfDay"`
}

// Predefined struct for user
type DeleteAgentRequestParams struct {
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>待删除AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 协作模式；0-Claw模式；1-Multi-Agent模式
	CollaborationMode *int64 `json:"CollaborationMode,omitnil,omitempty" name:"CollaborationMode"`
}

type DeleteAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>待删除AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 协作模式；0-Claw模式；1-Multi-Agent模式
	CollaborationMode *int64 `json:"CollaborationMode,omitnil,omitempty" name:"CollaborationMode"`
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
	delete(f, "AppId")
	delete(f, "AgentId")
	delete(f, "CollaborationMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAgentResponseParams struct {
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
	// <p>app_id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>删除原因(非必填,审批时展示)</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
}

type DeleteAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>app_id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>删除原因(非必填,审批时展示)</p>
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`
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
	delete(f, "AppId")
	delete(f, "Reason")
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
type DeleteAppTriggerRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DeleteAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DeleteAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAppTriggerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAppTriggerResponseParams `json:"Response"`
}

func (r *DeleteAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConversationRequestParams struct {
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`
}

type DeleteConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`
}

func (r *DeleteConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConversationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConversationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConversationResponseParams `json:"Response"`
}

func (r *DeleteConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMsgRecordCategoryRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>待删除的分类业务 ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

type DeleteMsgRecordCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>待删除的分类业务 ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`
}

func (r *DeleteMsgRecordCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMsgRecordCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "CategoryId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteMsgRecordCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteMsgRecordCategoryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteMsgRecordCategoryResponse struct {
	*tchttp.BaseResponse
	Response *DeleteMsgRecordCategoryResponseParams `json:"Response"`
}

func (r *DeleteMsgRecordCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteMsgRecordCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginRequestParams struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type DeletePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *DeletePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeletePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeletePluginResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeletePluginResponse struct {
	*tchttp.BaseResponse
	Response *DeletePluginResponseParams `json:"Response"`
}

func (r *DeletePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeletePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillRequestParams struct {
	// <p>Skill ID，必填</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DeleteSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>Skill ID，必填</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DeleteSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSkillResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSkillResponseParams `json:"Response"`
}

func (r *DeleteSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillShareRequestParams struct {
	// <p>申请备注，必填（弹窗&quot;申请备注&quot;）</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>原 Skill ID，必填（前端无须感知 _shared 后缀）</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>原版本 ID，必填（与 CreateSkillShare 上架时传的同一 version_id）</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type DeleteSkillShareRequest struct {
	*tchttp.BaseRequest
	
	// <p>申请备注，必填（弹窗&quot;申请备注&quot;）</p>
	ApplyRemark *string `json:"ApplyRemark,omitnil,omitempty" name:"ApplyRemark"`

	// <p>原 Skill ID，必填（前端无须感知 _shared 后缀）</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>原版本 ID，必填（与 CreateSkillShare 上架时传的同一 version_id）</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *DeleteSkillShareRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillShareRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ApplyRemark")
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSkillShareRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSkillShareResponseParams struct {
	// <p>是否走审批流（false 表示无需审批已直接执行下架）</p>
	NeedApproval *bool `json:"NeedApproval,omitnil,omitempty" name:"NeedApproval"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSkillShareResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSkillShareResponseParams `json:"Response"`
}

func (r *DeleteSkillShareResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSkillShareResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpaceRequestParams struct {
	// 空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DeleteSpaceRequest struct {
	*tchttp.BaseRequest
	
	// 空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DeleteSpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSpaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteSpaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSpaceResponseParams `json:"Response"`
}

func (r *DeleteSpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type。枚举值: 1:环境参数, 2:应用参数, 3:系统参数, -1:所有参数
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

type DeleteVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type。枚举值: 1:环境参数, 2:应用参数, 3:系统参数, -1:所有参数
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

func (r *DeleteVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "VariableId")
	delete(f, "ModuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteVariableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteVariableResponse struct {
	*tchttp.BaseResponse
	Response *DeleteVariableResponseParams `json:"Response"`
}

func (r *DeleteVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountListRequestParams struct {
	// <p>页码</p><p>从0开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>分页数量</p><p>取值范围：[1, 100]</p><p>单位：个</p><p>最大100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>参数过滤</p><p>支持SpaceId,NIckName 过滤查询</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeAccountListRequest struct {
	*tchttp.BaseRequest
	
	// <p>页码</p><p>从0开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>分页数量</p><p>取值范围：[1, 100]</p><p>单位：个</p><p>最大100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>参数过滤</p><p>支持SpaceId,NIckName 过滤查询</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeAccountListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountListResponseParams struct {
	// <p>总数</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>员工列表</p>
	AccountList []*AccountInfo `json:"AccountList,omitnil,omitempty" name:"AccountList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountListResponseParams `json:"Response"`
}

func (r *DescribeAccountListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDetailRequestParams struct {
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeAgentDetailRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeAgentDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentDetailResponseParams struct {
	// <p>Agent信息</p>
	Agent *AgentDetail `json:"Agent,omitnil,omitempty" name:"Agent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentDetailResponseParams `json:"Response"`
}

func (r *DescribeAgentDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentReleasePreviewListRequestParams struct {
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>页码</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量在1到200之间</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>查询关键字, 用于模糊匹配标题</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>过滤条件</p><p>入参限制：支持 StartTime、EndTime、ActionList、ReleaseStatusList</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeAgentReleasePreviewListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>页码</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量在1到200之间</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>查询关键字, 用于模糊匹配标题</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>过滤条件</p><p>入参限制：支持 StartTime、EndTime、ActionList、ReleaseStatusList</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeAgentReleasePreviewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentReleasePreviewListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentReleasePreviewListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentReleasePreviewListResponseParams struct {
	// <p>发布预览列表</p>
	ReleaseList []*AgentReleasePreview `json:"ReleaseList,omitnil,omitempty" name:"ReleaseList"`

	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentReleasePreviewListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentReleasePreviewListResponseParams `json:"Response"`
}

func (r *DescribeAgentReleasePreviewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentReleasePreviewListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSummaryListRequestParams struct {
	// <p>查询范围；0-单应用查询；1-跨应用查询</p>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用Id，Scope=0 时为目标应用ID（必填）；scope=1 时无需填写</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>过滤条件（name: "SearchWord", "SpaceId", "AgentSource", "AppId"）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>每页数目</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

type DescribeAgentSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>查询范围；0-单应用查询；1-跨应用查询</p>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用Id，Scope=0 时为目标应用ID（必填）；scope=1 时无需填写</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>过滤条件（name: "SearchWord", "SpaceId", "AgentSource", "AppId"）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>每页数目</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>页码</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`
}

func (r *DescribeAgentSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Scope")
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageSize")
	delete(f, "PageNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgentSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgentSummaryListResponseParams struct {
	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>Agent摘要信息</p>
	AgentList []*AgentSummary `json:"AgentList,omitnil,omitempty" name:"AgentList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAgentSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgentSummaryListResponseParams `json:"Response"`
}

func (r *DescribeAgentSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgentSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用域: ADP_DOMAIN_DEV(1)=开发域, ADP_DOMAIN_PROD(2)=发布域。枚举值: 1:开发域, 2:生产域</p>
	Domain *int64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>字段掩码，指定需要返回的字段(Paths为空则返回所有字段)。Paths枚举值：AppConfig(应用配置), SecretInfo(应用密钥信息), ShareUrlInfo(分享链接信息), SpecialStatusInfo(特殊状态信息), SearchResourceStatus(搜索资源状态), SharedKbList(应用引用的共享知识库列表),CorpShareConfig(企业共享配置)</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>特殊状态类型(当FieldMask包含SpecialStatusInfo时必填)。枚举值: 1:回滚状态, 2:首次导入状态</p>
	StatusType *int64 `json:"StatusType,omitnil,omitempty" name:"StatusType"`
}

type DescribeAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用域: ADP_DOMAIN_DEV(1)=开发域, ADP_DOMAIN_PROD(2)=发布域。枚举值: 1:开发域, 2:生产域</p>
	Domain *int64 `json:"Domain,omitnil,omitempty" name:"Domain"`

	// <p>字段掩码，指定需要返回的字段(Paths为空则返回所有字段)。Paths枚举值：AppConfig(应用配置), SecretInfo(应用密钥信息), ShareUrlInfo(分享链接信息), SpecialStatusInfo(特殊状态信息), SearchResourceStatus(搜索资源状态), SharedKbList(应用引用的共享知识库列表),CorpShareConfig(企业共享配置)</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>特殊状态类型(当FieldMask包含SpecialStatusInfo时必填)。枚举值: 1:回滚状态, 2:首次导入状态</p>
	StatusType *int64 `json:"StatusType,omitnil,omitempty" name:"StatusType"`
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
	delete(f, "AppId")
	delete(f, "Domain")
	delete(f, "FieldMask")
	delete(f, "StatusType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppResponseParams struct {
	// <p>应用详情</p>
	App *App `json:"App,omitnil,omitempty" name:"App"`

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
type DescribeAppSummaryListRequestParams struct {
	// 空间ID(必填)
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 过滤条件(多个Filter之间为AND关系,同一Filter的多个Values为OR关系): - AppStatus: 应用状态,枚举值,精确匹配(APP_STATUS_OFFLINE=1/APP_STATUS_RUNNING=2/APP_STATUS_DISABLED=3) - AppMode: 应用模式,枚举值,精确匹配(APP_MODE_STANDARD=1/APP_MODE_AGENT=2/APP_MODE_SINGLE_WORKFLOW=3/APP_MODE_CLAW_AGENT=4)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// 页码(从0开始)
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量(最大值:100)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 模糊查询
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeAppSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID(必填)
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 过滤条件(多个Filter之间为AND关系,同一Filter的多个Values为OR关系): - AppStatus: 应用状态,枚举值,精确匹配(APP_STATUS_OFFLINE=1/APP_STATUS_RUNNING=2/APP_STATUS_DISABLED=3) - AppMode: 应用模式,枚举值,精确匹配(APP_MODE_STANDARD=1/APP_MODE_AGENT=2/APP_MODE_SINGLE_WORKFLOW=3/APP_MODE_CLAW_AGENT=4)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// 页码(从0开始)
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量(最大值:100)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 模糊查询
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeAppSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppSummaryListResponseParams struct {
	// 应用摘要列表
	AppSummaryList []*AppSummary `json:"AppSummaryList,omitnil,omitempty" name:"AppSummaryList"`

	// total_count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppSummaryListResponseParams `json:"Response"`
}

func (r *DescribeAppSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerInstanceRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>触发器运行实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerInstanceRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>触发器运行实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "InstanceId")
	delete(f, "Scope")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerInstanceResponseParams struct {
	// <p>应用触发器实例</p>
	Instance *AppTriggerInstance `json:"Instance,omitnil,omitempty" name:"Instance"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerInstanceResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerResponseParams struct {
	// <p>应用触发器信息</p>
	Trigger *AppTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerRunLogListRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>过滤参数</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码</p><p>取值范围：[1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数据量</p><p>取值范围：[1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerRunLogListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>过滤参数</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码</p><p>取值范围：[1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数据量</p><p>取值范围：[1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerRunLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerRunLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerRunLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerRunLogListResponseParams struct {
	// <p>日志列表</p>
	RunLogList []*AppTriggerRunLog `json:"RunLogList,omitnil,omitempty" name:"RunLogList"`

	// <p>日志列表数量</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerRunLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerRunLogListResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerRunLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerRunLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerSummaryListRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>参数过滤列表</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码</p><p>取值范围：[1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小</p><p>取值范围：[1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>模糊查询字符串</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeAppTriggerSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>参数过滤列表</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码</p><p>取值范围：[1, 1000000]</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页大小</p><p>取值范围：[1, 100]</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>模糊查询字符串</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeAppTriggerSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "Scope")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppTriggerSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppTriggerSummaryListResponseParams struct {
	// <p>应用触发器数量</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>应用触发器列表</p>
	TriggerList []*AppTriggerSummary `json:"TriggerList,omitnil,omitempty" name:"TriggerList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAppTriggerSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppTriggerSummaryListResponseParams `json:"Response"`
}

func (r *DescribeAppTriggerSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppTriggerSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogListRequestParams struct {
	// <p>空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>每页数量</p><p>取值范围：[1, 100]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>es查询起始位置</p><p>对应接口返回SearchAfter</p>
	SearchAfter []*string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// <p>参数过滤</p><p>支持 Action,BizObject,Content<br>支持SpaceId,AccountUin,AppId(最多100个)<br>支持startTime,endTime(秒时间戳)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeAuditLogListRequest struct {
	*tchttp.BaseRequest
	
	// <p>空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>每页数量</p><p>取值范围：[1, 100]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>es查询起始位置</p><p>对应接口返回SearchAfter</p>
	SearchAfter []*string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// <p>参数过滤</p><p>支持 Action,BizObject,Content<br>支持SpaceId,AccountUin,AppId(最多100个)<br>支持startTime,endTime(秒时间戳)</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeAuditLogListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "Limit")
	delete(f, "SearchAfter")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogListResponseParams struct {
	// <p>操作日志列表</p>
	AuditLogList []*AuditLog `json:"AuditLogList,omitnil,omitempty" name:"AuditLogList"`

	// <p>es查询起始位置</p><p>用于入参查询下一页</p>
	SearchAfter []*string `json:"SearchAfter,omitnil,omitempty" name:"SearchAfter"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogListResponseParams `json:"Response"`
}

func (r *DescribeAuditLogListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogMetaRequestParams struct {

}

type DescribeAuditLogMetaRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAuditLogMetaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogMetaRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAuditLogMetaRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAuditLogMetaResponseParams struct {
	// <p>操作类型列表</p>
	Actions []*AuditLogMetaField `json:"Actions,omitnil,omitempty" name:"Actions"`

	// <p>操作对象列表</p>
	BizObjects []*AuditLogMetaField `json:"BizObjects,omitnil,omitempty" name:"BizObjects"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAuditLogMetaResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAuditLogMetaResponseParams `json:"Response"`
}

func (r *DescribeAuditLogMetaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAuditLogMetaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyLimitDetailListRequestParams struct {
	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图/ 应用视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：concurrency_type（qpm_tpm/dedicated，默认 qpm_tpm）、model_name（必填）、space_id、app_id/resource_id/source_id（应用ID，多选）、metric_source_type（METRIC_SOURCE_TYPE_* 枚举名或整数）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeConcurrencyLimitDetailListRequest struct {
	*tchttp.BaseRequest
	
	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图/ 应用视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：concurrency_type（qpm_tpm/dedicated，默认 qpm_tpm）、model_name（必填）、space_id、app_id/resource_id/source_id（应用ID，多选）、metric_source_type（METRIC_SOURCE_TYPE_* 枚举名或整数）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeConcurrencyLimitDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyLimitDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConcurrencyLimitDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConcurrencyLimitDetailListResponseParams struct {
	// <p>并发超限明细列表</p>
	ConcurrencyLimitDetailList []*ConcurrencyLimitDetail `json:"ConcurrencyLimitDetailList,omitnil,omitempty" name:"ConcurrencyLimitDetailList"`

	// <p>总记录数，用于前端分页</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConcurrencyLimitDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConcurrencyLimitDetailListResponseParams `json:"Response"`
}

func (r *DescribeConcurrencyLimitDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConcurrencyLimitDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumptionDetailListRequestParams struct {
	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：metric_source_type（METRIC_SOURCE_TYPE_* 或整数）、source_ids（多选来源ID）、resource_id/source_id（单选来源ID，source_ids 未传时生效）、space_id、user_id</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeConsumptionDetailListRequest struct {
	*tchttp.BaseRequest
	
	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：metric_source_type（METRIC_SOURCE_TYPE_* 或整数）、source_ids（多选来源ID）、resource_id/source_id（单选来源ID，source_ids 未传时生效）、space_id、user_id</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeConsumptionDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumptionDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConsumptionDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConsumptionDetailListResponseParams struct {
	// <p>资源消耗明细列表</p>
	ConsumptionDetailList []*ConsumptionDetail `json:"ConsumptionDetailList,omitnil,omitempty" name:"ConsumptionDetailList"`

	// <p>总记录数，用于前端分页</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConsumptionDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConsumptionDetailListResponseParams `json:"Response"`
}

func (r *DescribeConsumptionDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConsumptionDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationListRequestParams struct {
	// <p>会话类型，传 CONVERSATION_TYPE_UNSPECIFIED 表示全部 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>限制数目（整型），配合Offset使用</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>子账户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主账户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>偏移量（整型），配合Limit使用，从0开始</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户端 AgentId，当需要查询基于用户端 AgentId 创建的会话时使用</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

type DescribeConversationListRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话类型，传 CONVERSATION_TYPE_UNSPECIFIED 表示全部 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>限制数目（整型），配合Offset使用</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>子账户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主账户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>偏移量（整型），配合Limit使用，从0开始</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>用户端 AgentId，当需要查询基于用户端 AgentId 创建的会话时使用</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`
}

func (r *DescribeConversationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "Keyword")
	delete(f, "Limit")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "Offset")
	delete(f, "ShareCode")
	delete(f, "UserId")
	delete(f, "AgentId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConversationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationListResponseParams struct {
	// <p>会话列表</p>
	ConversationList []*Conversation `json:"ConversationList,omitnil,omitempty" name:"ConversationList"`

	// <p>会话列表</p>
	Conversations []*Conversation `json:"Conversations,omitnil,omitempty" name:"Conversations"`

	// <p>总数</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConversationListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConversationListResponseParams `json:"Response"`
}

func (r *DescribeConversationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationMessageListRequestParams struct {
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>返回记录总数量，默认 10，最大 50。向前或向后查询时，不包含record_id指定记录的消息，查询方向中心向前后查询时，包含record_id指定的记录消息，返回记录数量为前后各limit / 2条，向上取整</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>查询锚点记录 ID</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>相对于 record_id 的查询方向 枚举值: 0-RECORD_QUERY_DIRECTION_UNSPECIFIED(未指定，兼容旧逻辑，默认向前查询), 1-RECORD_QUERY_DIRECTION_BACKWARD(从 record_id 向前查询更早的消息), 2-RECORD_QUERY_DIRECTION_FORWARD(从 record_id 向后查询更新的消息), 3-RECORD_QUERY_DIRECTION_BIDIRECTIONAL(以 record_id 为中心，同时向前后查询)</p>
	RecordQueryDirection *int64 `json:"RecordQueryDirection,omitnil,omitempty" name:"RecordQueryDirection"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeConversationMessageListRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>返回记录总数量，默认 10，最大 50。向前或向后查询时，不包含record_id指定记录的消息，查询方向中心向前后查询时，包含record_id指定的记录消息，返回记录数量为前后各limit / 2条，向上取整</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>查询锚点记录 ID</p>
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// <p>相对于 record_id 的查询方向 枚举值: 0-RECORD_QUERY_DIRECTION_UNSPECIFIED(未指定，兼容旧逻辑，默认向前查询), 1-RECORD_QUERY_DIRECTION_BACKWARD(从 record_id 向前查询更早的消息), 2-RECORD_QUERY_DIRECTION_FORWARD(从 record_id 向后查询更新的消息), 3-RECORD_QUERY_DIRECTION_BIDIRECTIONAL(以 record_id 为中心，同时向前后查询)</p>
	RecordQueryDirection *int64 `json:"RecordQueryDirection,omitnil,omitempty" name:"RecordQueryDirection"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeConversationMessageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationMessageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "Limit")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "RecordId")
	delete(f, "RecordQueryDirection")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConversationMessageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationMessageListResponseParams struct {
	// <p>第一个记录 ID</p>
	FirstRecordId *string `json:"FirstRecordId,omitnil,omitempty" name:"FirstRecordId"`

	// <p>更新消息方向是否还有更多</p>
	HasMoreAfter *bool `json:"HasMoreAfter,omitnil,omitempty" name:"HasMoreAfter"`

	// <p>更早消息方向是否还有更多</p>
	HasMoreBefore *bool `json:"HasMoreBefore,omitnil,omitempty" name:"HasMoreBefore"`

	// <p>最后一个记录 ID</p>
	LastRecordId *string `json:"LastRecordId,omitnil,omitempty" name:"LastRecordId"`

	// <p>消息列表</p>
	MessageList []*ConversationMessage `json:"MessageList,omitnil,omitempty" name:"MessageList"`

	// <p>消息列表</p>
	//
	// Deprecated: Messages is deprecated.
	Messages []*ConversationMessage `json:"Messages,omitnil,omitempty" name:"Messages"`

	// <p>最近一次重置信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResetInfo *ConversationResetInfo `json:"ResetInfo,omitnil,omitempty" name:"ResetInfo"`

	// <p>单次对话记录统计列表，与 message_list 通过 record_id / related_record_id 关联</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	RecordSummaryList []*ConversationRecordSummary `json:"RecordSummaryList,omitnil,omitempty" name:"RecordSummaryList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConversationMessageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConversationMessageListResponseParams `json:"Response"`
}

func (r *DescribeConversationMessageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationMessageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationRequestParams struct {
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>主用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>子用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type DescribeConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>主用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>子用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *DescribeConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConversationResponseParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>工作空间</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Workspace *ConversationWorkspace `json:"Workspace,omitnil,omitempty" name:"Workspace"`

	// <p>会话标题</p>
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`

	// <p>会话使用的用户端 AgentId</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConversationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConversationResponseParams `json:"Response"`
}

func (r *DescribeConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestReleaseRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeLatestReleaseRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeLatestReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLatestReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLatestReleaseResponseParams struct {
	// 是否有发布变更
	IsChanged *bool `json:"IsChanged,omitnil,omitempty" name:"IsChanged"`

	// 发布信息
	ReleaseSummary *ReleaseSummary `json:"ReleaseSummary,omitnil,omitempty" name:"ReleaseSummary"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLatestReleaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLatestReleaseResponseParams `json:"Response"`
}

func (r *DescribeLatestReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLatestReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricOverviewListRequestParams struct {
	// <p>看板域，必填，决定返回哪个域的 KPI 数据</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>模型用量</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>插件用量</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>平台功能用量</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>模型并发超限</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>知识库容量</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>用量汇总</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>资源消耗（计费明细）</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤（resource_type=MODEL）。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：model_name（模型名）、user_id（用户ID）、space_id（空间ID）、resource_id/source_id（来源ID）、metric_source_type（METRIC_SOURCE_TYPE_* 枚举名或整数）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeMetricOverviewListRequest struct {
	*tchttp.BaseRequest
	
	// <p>看板域，必填，决定返回哪个域的 KPI 数据</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>模型用量</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>插件用量</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>平台功能用量</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>模型并发超限</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>知识库容量</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>用量汇总</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>资源消耗（计费明细）</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤（resource_type=MODEL）。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：model_name（模型名）、user_id（用户ID）、space_id（空间ID）、resource_id/source_id（来源ID）、metric_source_type（METRIC_SOURCE_TYPE_* 枚举名或整数）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeMetricOverviewListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricOverviewListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMetricOverviewListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricOverviewListResponseParams struct {
	// <p>所有域 Overview 统一出参：KPI 卡片列表，key 字符串标识指标，客户端按 resource_type 解析；key 白名单参考 platform.common.v2.MetricOverview 注释</p>
	MetricList []*MetricOverview `json:"MetricList,omitnil,omitempty" name:"MetricList"`

	// <p>总记录数，等于 MetricList 长度，仅为列表接口一致性预留</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMetricOverviewListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMetricOverviewListResponseParams `json:"Response"`
}

func (r *DescribeMetricOverviewListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricOverviewListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListRequestParams struct {
	// <p>模型场景。0-不区分场景, 1-标准生成, 2-标准思考, 3-Agent思考, 4-多模态理解, 5-多模态问答, 6-改写, 7-长期记忆, 8-自然语言转SQL, 9-AI优化, 10-实时文件解析, 11-文件解析, 12-GraphRAG, 13-OpenClaw, 14-多模态Embedding, 15-Rerank, 16-文本Embedding, 17-Widget, 18-Claw模式, 19-工作流代码生成, 20-工作流大模型节点, 21-工作流节点专用向量化, 22-工作流参数提取, 23-工作流大模型知识问答, 24-工作流标签提取, 25-工作流意图识别, 26-工作流选项卡, 27-工作流逻辑判断, 28-文档生成问答, 29-知识库Schema</p><p>枚举值：</p><ul><li>0： 不区分场景</li><li>1： 标准生成</li><li>2： 标准思考</li><li>3： Agent思考</li><li>4： 多模态理解</li><li>5： 多模态问答</li><li>6： 改写</li><li>7： 长期记忆</li><li>8： 自然语言转SQL</li><li>9： AI优化</li><li>10： 实时文件解析</li><li>11： 文件解析</li><li>12： GraphRAG</li><li>13： OpenClaw</li><li>14： 多模态Embedding</li><li>15： Rerank</li><li>16： 文本Embedding</li><li>17： Widget</li><li>18： Claw模式</li><li>19： 工作流代码生成</li><li>20： 工作流大模型节点</li><li>21： 工作流节点专用向量化</li><li>22： 工作流参数提取</li><li>23： 工作流大模型知识问答</li><li>24： 工作流标签提取</li><li>25： 工作流意图识别</li><li>26： 工作流选项卡</li><li>27： 工作流逻辑判断</li><li>28： 文档生成问答</li><li>29： 知识库Schema</li></ul>
	ModelScene *int64 `json:"ModelScene,omitnil,omitempty" name:"ModelScene"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>关键词模糊搜索</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>页码。从0开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，默认20，最大100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>过滤条件(多个 Filter 之间为 AND, 同一 Filter 多 Values 为 OR)<br>DeveloperName： 模型作者名称<br>ProviderName： 模型提供商名称<br>ProviderType：模型提供商类型</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

type DescribeModelListRequest struct {
	*tchttp.BaseRequest
	
	// <p>模型场景。0-不区分场景, 1-标准生成, 2-标准思考, 3-Agent思考, 4-多模态理解, 5-多模态问答, 6-改写, 7-长期记忆, 8-自然语言转SQL, 9-AI优化, 10-实时文件解析, 11-文件解析, 12-GraphRAG, 13-OpenClaw, 14-多模态Embedding, 15-Rerank, 16-文本Embedding, 17-Widget, 18-Claw模式, 19-工作流代码生成, 20-工作流大模型节点, 21-工作流节点专用向量化, 22-工作流参数提取, 23-工作流大模型知识问答, 24-工作流标签提取, 25-工作流意图识别, 26-工作流选项卡, 27-工作流逻辑判断, 28-文档生成问答, 29-知识库Schema</p><p>枚举值：</p><ul><li>0： 不区分场景</li><li>1： 标准生成</li><li>2： 标准思考</li><li>3： Agent思考</li><li>4： 多模态理解</li><li>5： 多模态问答</li><li>6： 改写</li><li>7： 长期记忆</li><li>8： 自然语言转SQL</li><li>9： AI优化</li><li>10： 实时文件解析</li><li>11： 文件解析</li><li>12： GraphRAG</li><li>13： OpenClaw</li><li>14： 多模态Embedding</li><li>15： Rerank</li><li>16： 文本Embedding</li><li>17： Widget</li><li>18： Claw模式</li><li>19： 工作流代码生成</li><li>20： 工作流大模型节点</li><li>21： 工作流节点专用向量化</li><li>22： 工作流参数提取</li><li>23： 工作流大模型知识问答</li><li>24： 工作流标签提取</li><li>25： 工作流意图识别</li><li>26： 工作流选项卡</li><li>27： 工作流逻辑判断</li><li>28： 文档生成问答</li><li>29： 知识库Schema</li></ul>
	ModelScene *int64 `json:"ModelScene,omitnil,omitempty" name:"ModelScene"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>关键词模糊搜索</p>
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>页码。从0开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，默认20，最大100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>过滤条件(多个 Filter 之间为 AND, 同一 Filter 多 Values 为 OR)<br>DeveloperName： 模型作者名称<br>ProviderName： 模型提供商名称<br>ProviderType：模型提供商类型</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`
}

func (r *DescribeModelListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelScene")
	delete(f, "SpaceId")
	delete(f, "Query")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "FilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeModelListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeModelListResponseParams struct {
	// <p>模型列表</p>
	ModelList []*Model `json:"ModelList,omitnil,omitempty" name:"ModelList"`

	// <p>模型总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeModelListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeModelListResponseParams `json:"Response"`
}

func (r *DescribeModelListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeModelListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordCategoryListRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeMsgRecordCategoryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeMsgRecordCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordCategoryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMsgRecordCategoryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordCategoryListResponseParams struct {
	// <p>消息记录分类树列表</p>
	CategoryList []*MsgRecordCategory `json:"CategoryList,omitnil,omitempty" name:"CategoryList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMsgRecordCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMsgRecordCategoryListResponseParams `json:"Response"`
}

func (r *DescribeMsgRecordCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordListRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <ul><li><strong>ChannelType</strong> :   0-表示全部 2-体验页面（腾讯云）3-调试页面（腾讯云）4-体验页面（手机号）5-对话端API接入 6-应用评测 7-调试API&#39; 10000-微信服务号 10001-微信订阅号 10002-企微应用 10004-微信客服 10005-微信小程序 10006-腾讯元器 10007-应用宝 10009-企微智能机器人 10014-企微智能机器人 10011-LINE 10012-Telegram 10013-钉钉机器人 10016-飞书机器人 30000-定时任务 30001-触发器 </li><li><strong>FeedbackType</strong> :   反馈类型，-1-表示全部 0-未评价 1-点赞 2-点踩 </li><li><strong>QueryType</strong> :    检索类型（按平台约定取值）, SessionId-sessionID Question-问题  Answer-回复 Intent-意图 User-用户 </li><li><strong>Query</strong> : 对应QueryType输入的过滤条件 </li><li><strong>CategoryId</strong> :  分类ID </li><li><strong>ReplyMethod</strong> :  回复类型 0-全部回复类型  1-拒答问题回复 2-问答直接回复 3-审核失败回复 4-知识润色回复 6-工作流回复 8-图片理解回复 9-搜索引擎回复 10-大模型直接回复 11-兜底回复 </li><li><strong>StartTime</strong> : 开始时间，秒级时间戳 </li><li><strong>EndTime</strong>:  结束时间，秒级时间戳 </li><li><strong>Cursor</strong> : 游标信息，上一页取响应 PrevCursor，下一页取响应 NextCursor  </li><li><strong>Direction</strong> : 方向，next 下一页，prev 上一页</li><li><strong>CallResult</strong> : 调用结果：0 全部 / 1 成功 / 2 失败 / 3 用户取消 &lt;</li><li><strong>FailReason</strong> : 失败原因 0-全部  1-审核失败 2-达到QPM上限 3-达到TPM上限 4-达到并发上限  5-系统异常 </li><li><strong>Intent</strong> :  意图</li></ul>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始；不传时按 0 处理</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100；不传或传 0 时按默认分页大小处理</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序条件，只支持按 CreateTime 排序</p>
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

type DescribeMsgRecordListRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <ul><li><strong>ChannelType</strong> :   0-表示全部 2-体验页面（腾讯云）3-调试页面（腾讯云）4-体验页面（手机号）5-对话端API接入 6-应用评测 7-调试API&#39; 10000-微信服务号 10001-微信订阅号 10002-企微应用 10004-微信客服 10005-微信小程序 10006-腾讯元器 10007-应用宝 10009-企微智能机器人 10014-企微智能机器人 10011-LINE 10012-Telegram 10013-钉钉机器人 10016-飞书机器人 30000-定时任务 30001-触发器 </li><li><strong>FeedbackType</strong> :   反馈类型，-1-表示全部 0-未评价 1-点赞 2-点踩 </li><li><strong>QueryType</strong> :    检索类型（按平台约定取值）, SessionId-sessionID Question-问题  Answer-回复 Intent-意图 User-用户 </li><li><strong>Query</strong> : 对应QueryType输入的过滤条件 </li><li><strong>CategoryId</strong> :  分类ID </li><li><strong>ReplyMethod</strong> :  回复类型 0-全部回复类型  1-拒答问题回复 2-问答直接回复 3-审核失败回复 4-知识润色回复 6-工作流回复 8-图片理解回复 9-搜索引擎回复 10-大模型直接回复 11-兜底回复 </li><li><strong>StartTime</strong> : 开始时间，秒级时间戳 </li><li><strong>EndTime</strong>:  结束时间，秒级时间戳 </li><li><strong>Cursor</strong> : 游标信息，上一页取响应 PrevCursor，下一页取响应 NextCursor  </li><li><strong>Direction</strong> : 方向，next 下一页，prev 上一页</li><li><strong>CallResult</strong> : 调用结果：0 全部 / 1 成功 / 2 失败 / 3 用户取消 &lt;</li><li><strong>FailReason</strong> : 失败原因 0-全部  1-审核失败 2-达到QPM上限 3-达到TPM上限 4-达到并发上限  5-系统异常 </li><li><strong>Intent</strong> :  意图</li></ul>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始；不传时按 0 处理</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100；不传或传 0 时按默认分页大小处理</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// <p>排序条件，只支持按 CreateTime 排序</p>
	Sort *Sort `json:"Sort,omitnil,omitempty" name:"Sort"`
}

func (r *DescribeMsgRecordListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMsgRecordListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMsgRecordListResponseParams struct {
	// <p>是否有更多页</p>
	HasMore *bool `json:"HasMore,omitnil,omitempty" name:"HasMore"`

	// <p>消息记录列表</p>
	MsgRecordList []*MsgRecord `json:"MsgRecordList,omitnil,omitempty" name:"MsgRecordList"`

	// <p>下一页游标信息</p>
	NextCursor *string `json:"NextCursor,omitnil,omitempty" name:"NextCursor"`

	// <p>上一页游标信息</p>
	PrevCursor *string `json:"PrevCursor,omitnil,omitempty" name:"PrevCursor"`

	// <p>符合条件的总记录数，用于前端分页显示</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMsgRecordListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMsgRecordListResponseParams `json:"Response"`
}

func (r *DescribeMsgRecordListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMsgRecordListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginRequestParams struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>获取指定字段</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>插件展示场景。不传或取 0 时不限定场景。</p><p>枚举值：</p><ul><li>0：不限定场景</li><li>1：Agent 模式</li><li>2：工作流</li><li>3：智能工作台</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`
}

type DescribePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>获取指定字段</p>
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`

	// <p>插件展示场景。不传或取 0 时不限定场景。</p><p>枚举值：</p><ul><li>0：不限定场景</li><li>1：Agent 模式</li><li>2：工作流</li><li>3：智能工作台</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`
}

func (r *DescribePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "SpaceId")
	delete(f, "FieldMask")
	delete(f, "Module")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginResponseParams struct {
	// <p>插件详情</p>
	Plugin *Plugin `json:"Plugin,omitnil,omitempty" name:"Plugin"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePluginResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginResponseParams `json:"Response"`
}

func (r *DescribePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginSummaryListRequestParams struct {
	// 空间ID，查询空间内的插件列表时使用
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 过滤条件列表 支持：PluginKind、CategoryKey、PluginSource、PluginId、PluginClass、BillingType
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>是否只返回已收藏插件。取 true 时，仅返回当前用户已收藏的插件；取 false 或不传时不按收藏状态过滤。</p>
	IsFavoriteOnly *bool `json:"IsFavoriteOnly,omitnil,omitempty" name:"IsFavoriteOnly"`

	// <p>插件展示场景。不传或取 0 时不限定场景。</p><p>枚举值：</p><ul><li>0：不限定场景</li><li>1：Agent 模式</li><li>2：工作流</li><li>3：智能工作台</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`

	// 页码 从0开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容 模糊匹配：插件名称/插件描述/工具名称/工具描述
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>排序方式。</p><p>枚举值：</p><ul><li>0：未指定，默认排序</li><li>1：按相关性排序</li><li>2：按更新时间排序</li><li>3：默认排序</li><li>4：按热度排序</li></ul>
	SortType *int64 `json:"SortType,omitnil,omitempty" name:"SortType"`
}

type DescribePluginSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID，查询空间内的插件列表时使用
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 过滤条件列表 支持：PluginKind、CategoryKey、PluginSource、PluginId、PluginClass、BillingType
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>是否只返回已收藏插件。取 true 时，仅返回当前用户已收藏的插件；取 false 或不传时不按收藏状态过滤。</p>
	IsFavoriteOnly *bool `json:"IsFavoriteOnly,omitnil,omitempty" name:"IsFavoriteOnly"`

	// <p>插件展示场景。不传或取 0 时不限定场景。</p><p>枚举值：</p><ul><li>0：不限定场景</li><li>1：Agent 模式</li><li>2：工作流</li><li>3：智能工作台</li></ul>
	Module *int64 `json:"Module,omitnil,omitempty" name:"Module"`

	// 页码 从0开始
	PageNumber *int64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询内容 模糊匹配：插件名称/插件描述/工具名称/工具描述
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`

	// <p>排序方式。</p><p>枚举值：</p><ul><li>0：未指定，默认排序</li><li>1：按相关性排序</li><li>2：按更新时间排序</li><li>3：默认排序</li><li>4：按热度排序</li></ul>
	SortType *int64 `json:"SortType,omitnil,omitempty" name:"SortType"`
}

func (r *DescribePluginSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FilterList")
	delete(f, "IsFavoriteOnly")
	delete(f, "Module")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	delete(f, "SortType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePluginSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePluginSummaryListResponseParams struct {
	// plugin_list
	PluginList []*PluginSummary `json:"PluginList,omitnil,omitempty" name:"PluginList"`

	// total_count
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePluginSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribePluginSummaryListResponseParams `json:"Response"`
}

func (r *DescribePluginSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePluginSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseListRequestParams struct {
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 页码(从0开始)
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量(最大值:100)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeReleaseListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 页码(从0开始)
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量(最大值:100)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeReleaseListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseListResponseParams struct {
	// release_list
	ReleaseList []*ReleaseRecord `json:"ReleaseList,omitnil,omitempty" name:"ReleaseList"`

	// total_count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseListResponseParams `json:"Response"`
}

func (r *DescribeReleaseListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseSummaryRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

type DescribeReleaseSummaryRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

func (r *DescribeReleaseSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ReleaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeReleaseSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeReleaseSummaryResponseParams struct {
	// 发布信息
	ReleaseSummary *ReleaseSummary `json:"ReleaseSummary,omitnil,omitempty" name:"ReleaseSummary"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeReleaseSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribeReleaseSummaryResponseParams `json:"Response"`
}

func (r *DescribeReleaseSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeReleaseSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillCategoryListRequestParams struct {

}

type DescribeSkillCategoryListRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeSkillCategoryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillCategoryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillCategoryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillCategoryListResponseParams struct {
	// Skill 分类列表
	CategoryList []*SkillCategory `json:"CategoryList,omitnil,omitempty" name:"CategoryList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillCategoryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillCategoryListResponseParams `json:"Response"`
}

func (r *DescribeSkillCategoryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillCategoryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillDetailRequestParams struct {
	// skillID
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// 空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 版本过滤条件(多个Filter之间为AND关系,同一Filter的多个Values为OR关系): - Perspective: 视角枚举,字符串单值,Values 长度必须为 1,多值视为非法;仅作用于详情返回的 version_list 裁剪,不决定接口本身可见性;不传默认 USER (USER=使用者视角,version_list 仅返回已上线版本 / EDITOR=编辑者视角,version_list 返回全部存活版本 / ALL=全量视角,同 EDITOR)
	VersionFilterList []*Filter `json:"VersionFilterList,omitnil,omitempty" name:"VersionFilterList"`
}

type DescribeSkillDetailRequest struct {
	*tchttp.BaseRequest
	
	// skillID
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// 空间ID
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 版本过滤条件(多个Filter之间为AND关系,同一Filter的多个Values为OR关系): - Perspective: 视角枚举,字符串单值,Values 长度必须为 1,多值视为非法;仅作用于详情返回的 version_list 裁剪,不决定接口本身可见性;不传默认 USER (USER=使用者视角,version_list 仅返回已上线版本 / EDITOR=编辑者视角,version_list 返回全部存活版本 / ALL=全量视角,同 EDITOR)
	VersionFilterList []*Filter `json:"VersionFilterList,omitnil,omitempty" name:"VersionFilterList"`
}

func (r *DescribeSkillDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionFilterList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillDetailResponseParams struct {
	// skill详情
	SkillDetail *SkillDetail `json:"SkillDetail,omitnil,omitempty" name:"SkillDetail"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillDetailResponseParams `json:"Response"`
}

func (r *DescribeSkillDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillReferenceListRequestParams struct {
	// <p>Skill ID，必填</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type DescribeSkillReferenceListRequest struct {
	*tchttp.BaseRequest
	
	// <p>Skill ID，必填</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID，必填</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *DescribeSkillReferenceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillReferenceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillReferenceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillReferenceListResponseParams struct {
	// <p>按 SkillRefType 分组的引用汇总：某类型 total_count = 0 时不入组（不返回空占位） 本期同时落 OPENCLAW / AGENT / CORP_ASSISTANT 三路</p>
	ReferenceList []*SkillReferenceGroup `json:"ReferenceList,omitnil,omitempty" name:"ReferenceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillReferenceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillReferenceListResponseParams `json:"Response"`
}

func (r *DescribeSkillReferenceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillReferenceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillSummaryListRequestParams struct {
	// 空间ID，必填
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 仅查询当前用户收藏的 Skill
	FavoriteOnly *bool `json:"FavoriteOnly,omitnil,omitempty" name:"FavoriteOnly"`

	//    过滤条件(多个Filter之间为AND关系,同一Filter的多个Values为OR关系):
	//    - SkillIdList: Skill ID列表,字符串数组,精确匹配
	//    - ProviderType: Skill 提供方类型,枚举值数组,精确匹配
	//      (SKILL_PROVIDER_TYPE_OFFICIAL=1/SKILL_PROVIDER_TYPE_THIRD_PARTY=2/SKILL_PROVIDER_TYPE_CUSTOM=3/SKILL_PROVIDER_TYPE_CUSTOM_SHARED=4)
	//    - CategoryKey: 分类标识,字符串数组,精确匹配
	//    - AnalysisStatus: 安全检测状态,枚举值数组,精确匹配
	//      (SKILL_ANALYSIS_PENDING=0/SKILL_ANALYSIS_RUNNING=1/SKILL_ANALYSIS_AVAILABLE=2/SKILL_ANALYSIS_UNAVAILABLE=3/SKILL_ANALYSIS_FAILED=4)
	//    - RiskLevel: 风险等级,枚举值数组,精确匹配
	//      (SKILL_RISK_NONE=0/SKILL_RISK_LOW=1/SKILL_RISK_MEDIUM=2/SKILL_RISK_HIGH=3)
	// - SkillStatus: Skill 维度发布状态,枚举值数组,精确匹配,多值之间 OR;仅在 Perspective=EDITOR/ALL 时有实际意义
	// (SKILL_STATUS_INITIALIZED=0/SKILL_STATUS_AUDITING=1/SKILL_STATUS_PENDING_RELEASE=2/SKILL_STATUS_RELEASED=3)
	//    - ShareStatus: 共享状态,枚举值数组,精确匹配,仅在ProviderType包含SKILL_PROVIDER_TYPE_CUSTOM/SKILL_PROVIDER_TYPE_CUSTOM_SHARED时生效
	//      (SHARE_STATUS_UNSHARED=0/SHARE_STATUS_SHARED=1/SHARE_STATUS_APPROVING=2)
	//    - Perspective: 视角枚举,字符串单值,Values 长度必须为 1,多值视为非法;仅在 ProviderType=SKILL_PROVIDER_TYPE_CUSTOM 时生效;不传默认 USER
	//      (USER=使用者视角,仅返回仅有使用权限的 Skill / EDITOR=编辑者视角,仅返回有编辑权限的 Skill / ALL=全量视角,返回有任一权限位的 Skill)
	//   - Creator: 创建者过滤,字符串单值,Values 长度必须为 1,多值视为非法;仅在 ProviderType=SKILL_PROVIDER_TYPE_CUSTOM 时生效
	//    当前仅支持占位符 "$self",表示仅返回当前调用者创建的 Skill
	//    后续如需扩展为指定身份,再在此处追加约定
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// 页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量，最大值 100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 名称/展示名称模糊搜索
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeSkillSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// 空间ID，必填
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 仅查询当前用户收藏的 Skill
	FavoriteOnly *bool `json:"FavoriteOnly,omitnil,omitempty" name:"FavoriteOnly"`

	//    过滤条件(多个Filter之间为AND关系,同一Filter的多个Values为OR关系):
	//    - SkillIdList: Skill ID列表,字符串数组,精确匹配
	//    - ProviderType: Skill 提供方类型,枚举值数组,精确匹配
	//      (SKILL_PROVIDER_TYPE_OFFICIAL=1/SKILL_PROVIDER_TYPE_THIRD_PARTY=2/SKILL_PROVIDER_TYPE_CUSTOM=3/SKILL_PROVIDER_TYPE_CUSTOM_SHARED=4)
	//    - CategoryKey: 分类标识,字符串数组,精确匹配
	//    - AnalysisStatus: 安全检测状态,枚举值数组,精确匹配
	//      (SKILL_ANALYSIS_PENDING=0/SKILL_ANALYSIS_RUNNING=1/SKILL_ANALYSIS_AVAILABLE=2/SKILL_ANALYSIS_UNAVAILABLE=3/SKILL_ANALYSIS_FAILED=4)
	//    - RiskLevel: 风险等级,枚举值数组,精确匹配
	//      (SKILL_RISK_NONE=0/SKILL_RISK_LOW=1/SKILL_RISK_MEDIUM=2/SKILL_RISK_HIGH=3)
	// - SkillStatus: Skill 维度发布状态,枚举值数组,精确匹配,多值之间 OR;仅在 Perspective=EDITOR/ALL 时有实际意义
	// (SKILL_STATUS_INITIALIZED=0/SKILL_STATUS_AUDITING=1/SKILL_STATUS_PENDING_RELEASE=2/SKILL_STATUS_RELEASED=3)
	//    - ShareStatus: 共享状态,枚举值数组,精确匹配,仅在ProviderType包含SKILL_PROVIDER_TYPE_CUSTOM/SKILL_PROVIDER_TYPE_CUSTOM_SHARED时生效
	//      (SHARE_STATUS_UNSHARED=0/SHARE_STATUS_SHARED=1/SHARE_STATUS_APPROVING=2)
	//    - Perspective: 视角枚举,字符串单值,Values 长度必须为 1,多值视为非法;仅在 ProviderType=SKILL_PROVIDER_TYPE_CUSTOM 时生效;不传默认 USER
	//      (USER=使用者视角,仅返回仅有使用权限的 Skill / EDITOR=编辑者视角,仅返回有编辑权限的 Skill / ALL=全量视角,返回有任一权限位的 Skill)
	//   - Creator: 创建者过滤,字符串单值,Values 长度必须为 1,多值视为非法;仅在 ProviderType=SKILL_PROVIDER_TYPE_CUSTOM 时生效
	//    当前仅支持占位符 "$self",表示仅返回当前调用者创建的 Skill
	//    后续如需扩展为指定身份,再在此处追加约定
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// 页码，从 0 开始
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量，最大值 100
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 名称/展示名称模糊搜索
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeSkillSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SpaceId")
	delete(f, "FavoriteOnly")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSkillSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSkillSummaryListResponseParams struct {
	// Skill 摘要列表
	SkillSummaryList []*SkillSummary `json:"SkillSummaryList,omitnil,omitempty" name:"SkillSummaryList"`

	// 总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSkillSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSkillSummaryListResponseParams `json:"Response"`
}

func (r *DescribeSkillSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSkillSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceListRequestParams struct {
	// 支持空间名称模糊搜索
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeSpaceListRequest struct {
	*tchttp.BaseRequest
	
	// 支持空间名称模糊搜索
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeSpaceListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSpaceListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSpaceListResponseParams struct {
	// 总数
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 空间列表
	SpaceList []*Space `json:"SpaceList,omitnil,omitempty" name:"SpaceList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSpaceListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSpaceListResponseParams `json:"Response"`
}

func (r *DescribeSpaceListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSpaceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemVariableListRequestParams struct {
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

type DescribeSystemVariableListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`
}

func (r *DescribeSystemVariableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemVariableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSystemVariableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSystemVariableListResponseParams struct {
	// system_variable_list
	SystemVariableList []*SystemVariable `json:"SystemVariableList,omitnil,omitempty" name:"SystemVariableList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeSystemVariableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSystemVariableListResponseParams `json:"Response"`
}

func (r *DescribeSystemVariableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSystemVariableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageDetailListRequestParams struct {
	// <p>资源类型，限定为 RESOURCE_TYPE_MODEL / RESOURCE_TYPE_PLUGIN</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>模型用量</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>插件用量</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>平台功能用量</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>模型并发超限</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>知识库容量</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>用量汇总</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>资源消耗（计费明细）</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图 / 应用视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤（resource_type=MODEL）。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：model_name、user_id、space_id、resource_id/source_id、metric_source_type（METRIC_SOURCE_TYPE_* 或整数）、call_type（调用类型）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeUsageDetailListRequest struct {
	*tchttp.BaseRequest
	
	// <p>资源类型，限定为 RESOURCE_TYPE_MODEL / RESOURCE_TYPE_PLUGIN</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>模型用量</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>插件用量</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>平台功能用量</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>模型并发超限</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>知识库容量</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>用量汇总</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>资源消耗（计费明细）</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图 / 应用视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤（resource_type=MODEL）。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：model_name、user_id、space_id、resource_id/source_id、metric_source_type（METRIC_SOURCE_TYPE_* 或整数）、call_type（调用类型）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeUsageDetailListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageDetailListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageDetailListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageDetailListResponseParams struct {
	// <p>总记录数，用于前端分页</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>资源调用时序明细列表</p>
	UsageDetailList []*UsageDetail `json:"UsageDetailList,omitnil,omitempty" name:"UsageDetailList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsageDetailListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageDetailListResponseParams `json:"Response"`
}

func (r *DescribeUsageDetailListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageDetailListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageSummaryListRequestParams struct {
	// <p>资源类型，限定为 MODEL / PLUGIN / PLATFORM</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>模型用量</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>插件用量</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>平台功能用量</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>模型并发超限</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>知识库容量</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>用量汇总</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>资源消耗（计费明细）</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图 / 应用视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤（resource_type=MODEL）。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：model_name（模型名）、user_id（用户ID）、space_id（空间ID）、resource_id/source_id（来源ID）、metric_source_type（METRIC_SOURCE_TYPE_* 枚举名或整数）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeUsageSummaryListRequest struct {
	*tchttp.BaseRequest
	
	// <p>资源类型，限定为 MODEL / PLUGIN / PLATFORM</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>RESOURCE_TYPE_UNSPECIFIED</td><td>0</td><td></td></tr><tr><td>RESOURCE_TYPE_MODEL</td><td>1</td><td>模型用量</td></tr><tr><td>RESOURCE_TYPE_PLUGIN</td><td>2</td><td>插件用量</td></tr><tr><td>RESOURCE_TYPE_PLATFORM</td><td>3</td><td>平台功能用量</td></tr><tr><td>RESOURCE_TYPE_MODEL_CONCURRENCY</td><td>4</td><td>模型并发超限</td></tr><tr><td>RESOURCE_TYPE_KB_CAPACITY</td><td>5</td><td>知识库容量</td></tr><tr><td>RESOURCE_TYPE_USAGE_SUMMARY</td><td>6</td><td>用量汇总</td></tr><tr><td>RESOURCE_TYPE_RESOURCE_CONSUME</td><td>7</td><td>资源消耗（计费明细）</td></tr></tbody></table>
	ResourceType *int64 `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>查询时间范围（Unix 秒）</p>
	TimeRange *TimeRange `json:"TimeRange,omitnil,omitempty" name:"TimeRange"`

	// <p>视图范围：企业视图 / 空间视图 / 应用视图</p>
	ViewScope *ViewScope `json:"ViewScope,omitnil,omitempty" name:"ViewScope"`

	// <p>扩展过滤（resource_type=MODEL）。Filter 组合规则：多项 AND，同项 value_list OR。支持 Name：model_name（模型名）、user_id（用户ID）、space_id（空间ID）、resource_id/source_id（来源ID）、metric_source_type（METRIC_SOURCE_TYPE_* 枚举名或整数）</p>
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// <p>页码，从 0 开始</p>
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// <p>每页数量，最大 100</p>
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeUsageSummaryListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageSummaryListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ResourceType")
	delete(f, "TimeRange")
	delete(f, "ViewScope")
	delete(f, "FilterList")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUsageSummaryListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUsageSummaryListResponseParams struct {
	// <p>总记录数，用于前端分页</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>资源用量聚合明细列表</p>
	UsageSummaryList []*UsageSummary `json:"UsageSummaryList,omitnil,omitempty" name:"UsageSummaryList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUsageSummaryListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUsageSummaryListResponseParams `json:"Response"`
}

func (r *DescribeUsageSummaryListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUsageSummaryListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableListRequestParams struct {
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 过滤条件(支持: VariableIdList-变量ID列表, VariableType-变量类型)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// 模块类型。枚举值: 1:环境参数, 2:应用参数, 3:系统参数, -1:所有参数
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`

	// 是否需要内部变量
	NeedInternalVariable *bool `json:"NeedInternalVariable,omitnil,omitempty" name:"NeedInternalVariable"`

	// 页码(从0开始)
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量(最大值:100)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询关键词
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

type DescribeVariableListRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 过滤条件(支持: VariableIdList-变量ID列表, VariableType-变量类型)
	FilterList []*Filter `json:"FilterList,omitnil,omitempty" name:"FilterList"`

	// 模块类型。枚举值: 1:环境参数, 2:应用参数, 3:系统参数, -1:所有参数
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`

	// 是否需要内部变量
	NeedInternalVariable *bool `json:"NeedInternalVariable,omitnil,omitempty" name:"NeedInternalVariable"`

	// 页码(从0开始)
	PageNumber *uint64 `json:"PageNumber,omitnil,omitempty" name:"PageNumber"`

	// 每页数量(最大值:100)
	PageSize *uint64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 查询关键词
	Query *string `json:"Query,omitnil,omitempty" name:"Query"`
}

func (r *DescribeVariableListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "FilterList")
	delete(f, "ModuleType")
	delete(f, "NeedInternalVariable")
	delete(f, "PageNumber")
	delete(f, "PageSize")
	delete(f, "Query")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVariableListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableListResponseParams struct {
	// total_count
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// variable_list
	VariableList []*Variable `json:"VariableList,omitnil,omitempty" name:"VariableList"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVariableListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVariableListResponseParams `json:"Response"`
}

func (r *DescribeVariableListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type。枚举值: 1:环境参数, 2:应用参数, 3:系统参数, -1:所有参数
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

type DescribeVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// variable_id
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// module_type。枚举值: 1:环境参数, 2:应用参数, 3:系统参数, -1:所有参数
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`
}

func (r *DescribeVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "VariableId")
	delete(f, "ModuleType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeVariableResponseParams struct {
	// 变量信息
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeVariableResponse struct {
	*tchttp.BaseResponse
	Response *DescribeVariableResponseParams `json:"Response"`
}

func (r *DescribeVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DialogCustomConfig struct {
	// <p>是否开启对话端动态修改配置</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type DigitalHumanConfig struct {
	// 数智人形象资产id
	AssetKey *string `json:"AssetKey,omitnil,omitempty" name:"AssetKey"`

	// 数智人图片
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// 数智人形象名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 数智人预览地址
	PreviewUrl *string `json:"PreviewUrl,omitnil,omitempty" name:"PreviewUrl"`
}

type DuplexBilling struct {
	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>UNKNOW</td><td>0</td><td></td></tr><tr><td>TOKEN</td><td>1</td><td>按token</td></tr><tr><td>PAGE_COUNT</td><td>2</td><td>按页数</td></tr><tr><td>TIMES</td><td>3</td><td>按次数</td></tr><tr><td>TIMES_THOUSAND</td><td>4</td><td>按千次数</td></tr><tr><td>SECOND</td><td>5</td><td>按时长</td></tr><tr><td>CHARACTER</td><td>6</td><td>按字符数</td></tr><tr><td>CHARACTER_THOUSAND</td><td>7</td><td>按千字符数</td></tr><tr><td>SHEET</td><td>8</td><td>按张</td></tr><tr><td>NUMBER</td><td>9</td><td>按个数</td></tr></tbody></table>
	BillingUnit *int64 `json:"BillingUnit,omitnil,omitempty" name:"BillingUnit"`

	// <p>输入现金价格</p><p>单位：元</p>
	InputCashPrice *float64 `json:"InputCashPrice,omitnil,omitempty" name:"InputCashPrice"`

	// <p>输入pu价格</p><p>单位：pu</p>
	InputPuPrice *float64 `json:"InputPuPrice,omitnil,omitempty" name:"InputPuPrice"`

	// <p>输出现金价格</p><p>单位：元</p>
	OutputCashPrice *float64 `json:"OutputCashPrice,omitnil,omitempty" name:"OutputCashPrice"`

	// <p>输出pu价格</p><p>单位：pu</p>
	OutputPuPrice *float64 `json:"OutputPuPrice,omitnil,omitempty" name:"OutputPuPrice"`
}

type ExecuteConfig struct {
	// <p>Prompt配置</p>
	PromptConfig *AppTriggerPromptExecuteConfig `json:"PromptConfig,omitnil,omitempty" name:"PromptConfig"`

	// <p>工作流配置</p>
	WorkflowConfig *AppTriggerWorkflowExecuteConfig `json:"WorkflowConfig,omitnil,omitempty" name:"WorkflowConfig"`
}

// Predefined struct for user
type FavoritePluginRequestParams struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type FavoritePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *FavoritePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoritePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FavoritePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FavoritePluginResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FavoritePluginResponse struct {
	*tchttp.BaseResponse
	Response *FavoritePluginResponseParams `json:"Response"`
}

func (r *FavoritePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoritePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FavoriteSkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type FavoriteSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *FavoriteSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoriteSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "FavoriteSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type FavoriteSkillResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type FavoriteSkillResponse struct {
	*tchttp.BaseResponse
	Response *FavoriteSkillResponseParams `json:"Response"`
}

func (r *FavoriteSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *FavoriteSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldMask struct {
	// <p>参数名称</p><p>参数格式：需要获取的指定字段路径</p>
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

type FileParseModel struct {
	// 模型别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 模型描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 增强模式
	EnhancementMode *string `json:"EnhancementMode,omitnil,omitempty" name:"EnhancementMode"`

	// 模型唯一ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型类型
	ModelProviderType *string `json:"ModelProviderType,omitnil,omitempty" name:"ModelProviderType"`

	// 是否启用公式增强
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableFormulaEnhancement *bool `json:"EnableFormulaEnhancement,omitnil,omitempty" name:"EnableFormulaEnhancement"`

	// 是否启用 LLM 增强
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableLLMEnhancement *bool `json:"EnableLLMEnhancement,omitnil,omitempty" name:"EnableLLMEnhancement"`

	// 是否输出 HTML 表格
	// 注意：此字段可能返回 null，表示取不到有效值。
	OutputHtmlTable *bool `json:"OutputHtmlTable,omitnil,omitempty" name:"OutputHtmlTable"`

	// 支持的文件类型列表
	SupportedFileList []*SupportedFileType `json:"SupportedFileList,omitnil,omitempty" name:"SupportedFileList"`
}

type Filter struct {
	// 过滤字段名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 操作符，默认 IN（向后兼容）<table><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>FILTER_OPERATOR_IN</td><td>0</td><td>属于 value_list（默认值，向后兼容；value_list 不可为空）</td></tr><tr><td>FILTER_OPERATOR_NOT_IN</td><td>1</td><td>不属于 value_list（value_list 不可为空）</td></tr></table>
	Operator *int64 `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 过滤值数组
	ValueList []*string `json:"ValueList,omitnil,omitempty" name:"ValueList"`
}

type GenerateModel struct {
	// 生成模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type Identity struct {
	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>数字 ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>字符串 ID</p>
	StrId *string `json:"StrId,omitnil,omitempty" name:"StrId"`
}

type InputBoxConfig struct {
	// 输入框按钮，1：上传图片、2：上传文档，3：腾讯文档，4：联网搜索
	InputBoxButtons []*int64 `json:"InputBoxButtons,omitnil,omitempty" name:"InputBoxButtons"`
}

type IntentAchievementInfo struct {
	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type IntervalSchedule struct {
	// 开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 
	// 枚举值:
	// | 枚举值 | uint |
	// | --- | --- |
	// | INTERVAL_UNIT_UNSPECIFIED | 0 |
	// | INTERVAL_UNIT_HOUR | 1 |
	// | INTERVAL_UNIT_DAY | 2 |
	Unit *int64 `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 值
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type MCPPluginConfig struct {
	// <p>MCP插件外部访问地址</p>
	ExternalMCPServerUrl *string `json:"ExternalMCPServerUrl,omitnil,omitempty" name:"ExternalMCPServerUrl"`

	// <p>MCP server地址</p>
	MCPServerUrl *string `json:"MCPServerUrl,omitnil,omitempty" name:"MCPServerUrl"`

	// <p>MCP传输类型: SSE/Streamable<br>枚举值:<br>| uint | 描述 |<br>| --- | --- |<br>| 0 | SSE + HTTP 模式 |<br>| 1 | Streamable HTTP 模式 |</p>
	MCPTransport *int64 `json:"MCPTransport,omitnil,omitempty" name:"MCPTransport"`

	// <p>MCP插件的header参数</p>
	PluginHeader []*PluginParam `json:"PluginHeader,omitnil,omitempty" name:"PluginHeader"`

	// <p>MCP插件的query参数</p>
	PluginQuery []*PluginParam `json:"PluginQuery,omitnil,omitempty" name:"PluginQuery"`

	// <p>SSE长连接超时时间，单位秒</p>
	SSEReadTimeout *int64 `json:"SSEReadTimeout,omitnil,omitempty" name:"SSEReadTimeout"`

	// <p>请求超时时间，单位秒</p>
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>授权信息</p>
	AuthConfig *AuthConfig `json:"AuthConfig,omitnil,omitempty" name:"AuthConfig"`
}

type MCPToolConfig struct {
	// <p>输入参数</p>
	Inputs []*RequestParam `json:"Inputs,omitnil,omitempty" name:"Inputs"`

	// <p>输出参数</p>
	Outputs []*ResponseParam `json:"Outputs,omitnil,omitempty" name:"Outputs"`
}

type ManualOnlySchedule struct {
	// 启用
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type MetricOverview struct {
	// <p>指标键，取值参考 MetricOverview 注释中的 key 白名单</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>环比百分比，无环比时填 0</p>
	Mom *float64 `json:"Mom,omitnil,omitempty" name:"Mom"`

	// <p>指标单位，枚举值 DosageUnit；key 与 unit 的对应关系参考 MetricOverview 注释白名单</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>DOSAGE_UNIT_TOKEN</td><td>0</td><td>token（默认）</td></tr><tr><td>DOSAGE_UNIT_PAGE_COUNT</td><td>1</td><td>page_count（页数）</td></tr><tr><td>DOSAGE_UNIT_TIMES</td><td>2</td><td>times（次数）</td></tr><tr><td>DOSAGE_UNIT_SECOND</td><td>3</td><td>second（秒）</td></tr><tr><td>DOSAGE_UNIT_ITEM</td><td>4</td><td>item（条）</td></tr><tr><td>DOSAGE_UNIT_SHEET</td><td>5</td><td>sheet（张）</td></tr><tr><td>DOSAGE_UNIT_CHARACTER</td><td>6</td><td>character（字符）</td></tr><tr><td>DOSAGE_UNIT_GB</td><td>7</td><td>GB</td></tr><tr><td>DOSAGE_UNIT_NUMBER</td><td>8</td><td>number（个数）</td></tr><tr><td>DOSAGE_UNIT_MILL_SECOND</td><td>9</td><td>mill_second（毫秒）</td></tr></tbody></table>
	Unit *int64 `json:"Unit,omitnil,omitempty" name:"Unit"`

	// <p>指标数值</p>
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type Model struct {
	// <p>模型徽章列表</p>
	BadgeList []*ModelBadge `json:"BadgeList,omitnil,omitempty" name:"BadgeList"`

	// <p>模型限制信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimitInfo *ModelLimit `json:"LimitInfo,omitnil,omitempty" name:"LimitInfo"`

	// <p>模型基本信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelBasic *ModelBasic `json:"ModelBasic,omitnil,omitempty" name:"ModelBasic"`

	// <p>模型超参配置</p>
	ParameterList []*ModelParameter `json:"ParameterList,omitnil,omitempty" name:"ParameterList"`

	// <p>模型属性配置</p>
	PropertyList []*ModelProperty `json:"PropertyList,omitnil,omitempty" name:"PropertyList"`

	// <p>模型提供商信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProviderInfo *ModelProviderBasic `json:"ProviderInfo,omitnil,omitempty" name:"ProviderInfo"`

	// <p>模型状态信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusInfo *ModelStatus `json:"StatusInfo,omitnil,omitempty" name:"StatusInfo"`

	// <p>模型标签列表</p>
	TagList []*string `json:"TagList,omitnil,omitempty" name:"TagList"`

	// <p>模型作者信息</p>
	DeveloperInfo *ModelDeveloperBasic `json:"DeveloperInfo,omitnil,omitempty" name:"DeveloperInfo"`
}

type ModelBadge struct {
	// 展示文案
	Text *string `json:"Text,omitnil,omitempty" name:"Text"`

	// 样式主题。1-信息（蓝色）, 2-成功（绿色）, 3-警告（橙色）, 4-危险（红色）
	Theme *int64 `json:"Theme,omitnil,omitempty" name:"Theme"`

	// tooltip文案，为空则不展示
	Tips *string `json:"Tips,omitnil,omitempty" name:"Tips"`

	// 徽章类型。1-限时免费, 2-即将下线, 3-新模型, 4-热门
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModelBasic struct {
	// 模型描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 模型图标地址
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// 模型ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型类型。1-LLM模型, 2-Rerank模型, 3-Embedding模型, 4-文档解析模型
	ModelType *int64 `json:"ModelType,omitnil,omitempty" name:"ModelType"`

	// 模型名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModelDetailInfo struct {
	// 模型别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 历史对话条数限制
	HistoryLimit *uint64 `json:"HistoryLimit,omitnil,omitempty" name:"HistoryLimit"`

	// 模型唯一 ID
	ModelId *string `json:"ModelId,omitnil,omitempty" name:"ModelId"`

	// 模型参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModelParams *ModelParams `json:"ModelParams,omitnil,omitempty" name:"ModelParams"`
}

type ModelDeveloperBasic struct {
	// <p>作者标识</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>作者显示名称</p>
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`
}

type ModelLimit struct {
	// 模型上下文长度展示文案（如 "128K"、"1000K"）
	ContextLengthDescription *string `json:"ContextLengthDescription,omitnil,omitempty" name:"ContextLengthDescription"`

	// 模型对话框输入长度字符数限制
	InputLengthLimit *int64 `json:"InputLengthLimit,omitnil,omitempty" name:"InputLengthLimit"`

	// 模型提示词长度字符数限制
	PromptLengthLimit *int64 `json:"PromptLengthLimit,omitnil,omitempty" name:"PromptLengthLimit"`
}

type ModelParameter struct {
	// <p>默认值</p>
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// <p>可选值列表</p>
	EnumValueList []*string `json:"EnumValueList,omitnil,omitempty" name:"EnumValueList"`

	// <p>最大值（仅数值类型有效）</p>
	MaxValue *float64 `json:"MaxValue,omitnil,omitempty" name:"MaxValue"`

	// <p>最小值（仅数值类型有效）</p>
	MinValue *float64 `json:"MinValue,omitnil,omitempty" name:"MinValue"`

	// <p>超参名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>超参类型。1-浮点数, 2-整数, 3-字符串</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type ModelParams struct {
	// 是否开启深度思考
	DeepThinking *string `json:"DeepThinking,omitnil,omitempty" name:"DeepThinking"`

	// 频率惩罚
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrequencyPenalty *float64 `json:"FrequencyPenalty,omitnil,omitempty" name:"FrequencyPenalty"`

	// 最大输出长度
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxTokens *int64 `json:"MaxTokens,omitnil,omitempty" name:"MaxTokens"`

	// 存在惩罚
	// 注意：此字段可能返回 null，表示取不到有效值。
	PresencePenalty *float64 `json:"PresencePenalty,omitnil,omitempty" name:"PresencePenalty"`

	// 深度思考效果
	ReasoningEffort *string `json:"ReasoningEffort,omitnil,omitempty" name:"ReasoningEffort"`

	// 重复惩罚
	// 注意：此字段可能返回 null，表示取不到有效值。
	RepetitionPenalty *float64 `json:"RepetitionPenalty,omitnil,omitempty" name:"RepetitionPenalty"`

	// 输出格式（text、json_object）
	ReplyFormat *string `json:"ReplyFormat,omitnil,omitempty" name:"ReplyFormat"`

	// seed 随机种子
	// 注意：此字段可能返回 null，表示取不到有效值。
	Seed *int64 `json:"Seed,omitnil,omitempty" name:"Seed"`

	// 停止序列
	StopSequenceList []*string `json:"StopSequenceList,omitnil,omitempty" name:"StopSequenceList"`

	// 温度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Temperature *float64 `json:"Temperature,omitnil,omitempty" name:"Temperature"`

	// top_p
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopP *float64 `json:"TopP,omitnil,omitempty" name:"TopP"`
}

type ModelProperty struct {
	// 属性名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 属性值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ModelProviderBasic struct {
	// 模型提供商别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 模型提供商名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 模型提供商类型。1-自有提供商, 2-自定义模型提供商, 3-第三方模型提供商
	ProviderType *int64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`
}

type ModelStatus struct {
	// 专属并发数
	Concurrency *int64 `json:"Concurrency,omitnil,omitempty" name:"Concurrency"`

	// 是否专属并发
	IsExclusive *bool `json:"IsExclusive,omitnil,omitempty" name:"IsExclusive"`

	// 资源状态。1-资源可用, 2-资源已用尽
	ResourceStatus *int64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type ModelUsageDetail struct {
	// <p>调用类型，来源于计费 scene_billing（与 filter.call_type 对应）</p>
	CallType *string `json:"CallType,omitnil,omitempty" name:"CallType"`

	// <p>是否默认知识库</p>
	IsDefaultKB *bool `json:"IsDefaultKB,omitnil,omitempty" name:"IsDefaultKB"`

	// <p>模型名称</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>MODEL 域单次调用的消耗计量列表（权威字段）：按单位+label 分项列出每类计量。unit=TOKEN 时 label 区分 Token 子类别（input/output/avg_*/cache_*），label 为空表示 total_tokens；unit=PAGE_COUNT 表示模型消耗页数</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`

	// <p>本次调用消耗 PU 量</p>
	ConsumptionPU *float64 `json:"ConsumptionPU,omitnil,omitempty" name:"ConsumptionPU"`
}

type ModelUsageSummary struct {
	// <p>调用次数（业务调用维度的顶层计数）</p>
	CallCount *float64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// <p>是否默认知识库</p>
	IsDefaultKB *bool `json:"IsDefaultKB,omitnil,omitempty" name:"IsDefaultKB"`

	// <p>模型名称，标识使用的 AI 模型</p>
	ModelName *string `json:"ModelName,omitnil,omitempty" name:"ModelName"`

	// <p>MODEL 域消耗计量列表（权威字段）：按单位+label 分项列出每类计量。unit=TOKEN 时 label 区分 Token 子类别（input/output/avg_*/cache_*），label 为空表示 total_tokens；unit=PAGE_COUNT 表示模型消耗页数</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`

	// <p>模型消耗 PU 总量（聚合维度内的 PU 消耗之和）</p>
	ConsumptionPU *float64 `json:"ConsumptionPU,omitnil,omitempty" name:"ConsumptionPU"`
}

// Predefined struct for user
type ModifyAgentRequestParams struct {
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent Id</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>修改后的Agent的信息</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>需要更新的字段路径，如 ["Profile.Name", "Profile.IconUrl", "Instructions", "Model", "ToolList", "PluginList", "SkillList", "AdvancedConfig"]</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyAgentRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用Id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Agent Id</p>
	AgentId *string `json:"AgentId,omitnil,omitempty" name:"AgentId"`

	// <p>修改后的Agent的信息</p>
	Agent *AgentSpec `json:"Agent,omitnil,omitempty" name:"Agent"`

	// <p>需要更新的字段路径，如 ["Profile.Name", "Profile.IconUrl", "Instructions", "Model", "ToolList", "PluginList", "SkillList", "AdvancedConfig"]</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
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
	delete(f, "AppId")
	delete(f, "AgentId")
	delete(f, "Agent")
	delete(f, "UpdateMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAgentResponseParams struct {
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
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用模式。枚举值: 1:标准模式, 2:Agent模式, 3:单工作流模式, 4:ClawAgent模式</p>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// <p>应用头像</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// <p>应用配置</p>
	Config *AppConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>应用描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>应用名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>引用的共享知识库ID列表(全量覆盖)</p>
	SharedKbIdList []*string `json:"SharedKbIdList,omitnil,omitempty" name:"SharedKbIdList"`

	// <p>字段掩码，指定需要更新的字段(Paths为空则不更新任何字段)。Paths枚举值：<br>【顶层】Name, Avatar, Description, AppMode, SharedKbIdList<br>【Greeting】Config.Greeting, Config.Greeting.Greeting, Config.Greeting.OpeningQuestionList<br>【Model】Config.Model, Config.Model.ThinkModel, Config.Model.GenerateModel, Config.Model.AiOptimizeModel, Config.Model.FileParseModel, Config.Model.PromptRewriteModel, Config.Model.MultiModalQaModel, Config.Model.MultiModalUnderstandingModel<br>【WebSearch】Config.WebSearch<br>【Memory】Config.Memory, Config.Memory.Enabled, Config.Memory.LongMemoryDay, Config.Memory.Model, Config.Memory.PromptMode, Config.Memory.PromptContent<br>【Mode】Config.Mode, Config.Mode.MultiAgentConfig, Config.Mode.SingleWorkflowConfig, Config.Mode.ClawAgentConfig<br>【Mode.ClawAgentConfig】Config.Mode.ClawAgentConfig.LongMemoryConfig, Config.Mode.ClawAgentConfig.AgentTeamConfig<br>【Experience】Config.Experience, Config.Experience.Conversation, Config.Experience.Role, Config.Experience.Advanced<br>【Experience.Conversation】Config.Experience.Conversation.AiCall, Config.Experience.Conversation.BackgroundImage, Config.Experience.Conversation.Method, Config.Experience.Conversation.FallbackReply, Config.Experience.Conversation.Recommended, Config.Experience.Conversation.InputBoxConfig, Config.Experience.Conversation.WebSearch<br>【Experience.Conversation.AiCall】Config.Experience.Conversation.AiCall.VoiceInteract, Config.Experience.Conversation.AiCall.VoiceCall, Config.Experience.Conversation.AiCall.DigitalHuman<br>【Experience.Advanced】Config.Experience.Advanced.ContextRewrite, Config.Experience.Advanced.ImageTextRetrieval, Config.Experience.Advanced.IntentAchievement, Config.Experience.Advanced.ReplyFlexibility, Config.Experience.Advanced.DialogCustomConfig</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
}

type ModifyAppRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>应用模式。枚举值: 1:标准模式, 2:Agent模式, 3:单工作流模式, 4:ClawAgent模式</p>
	AppMode *int64 `json:"AppMode,omitnil,omitempty" name:"AppMode"`

	// <p>应用头像</p>
	Avatar *string `json:"Avatar,omitnil,omitempty" name:"Avatar"`

	// <p>应用配置</p>
	Config *AppConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>应用描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>应用名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>引用的共享知识库ID列表(全量覆盖)</p>
	SharedKbIdList []*string `json:"SharedKbIdList,omitnil,omitempty" name:"SharedKbIdList"`

	// <p>字段掩码，指定需要更新的字段(Paths为空则不更新任何字段)。Paths枚举值：<br>【顶层】Name, Avatar, Description, AppMode, SharedKbIdList<br>【Greeting】Config.Greeting, Config.Greeting.Greeting, Config.Greeting.OpeningQuestionList<br>【Model】Config.Model, Config.Model.ThinkModel, Config.Model.GenerateModel, Config.Model.AiOptimizeModel, Config.Model.FileParseModel, Config.Model.PromptRewriteModel, Config.Model.MultiModalQaModel, Config.Model.MultiModalUnderstandingModel<br>【WebSearch】Config.WebSearch<br>【Memory】Config.Memory, Config.Memory.Enabled, Config.Memory.LongMemoryDay, Config.Memory.Model, Config.Memory.PromptMode, Config.Memory.PromptContent<br>【Mode】Config.Mode, Config.Mode.MultiAgentConfig, Config.Mode.SingleWorkflowConfig, Config.Mode.ClawAgentConfig<br>【Mode.ClawAgentConfig】Config.Mode.ClawAgentConfig.LongMemoryConfig, Config.Mode.ClawAgentConfig.AgentTeamConfig<br>【Experience】Config.Experience, Config.Experience.Conversation, Config.Experience.Role, Config.Experience.Advanced<br>【Experience.Conversation】Config.Experience.Conversation.AiCall, Config.Experience.Conversation.BackgroundImage, Config.Experience.Conversation.Method, Config.Experience.Conversation.FallbackReply, Config.Experience.Conversation.Recommended, Config.Experience.Conversation.InputBoxConfig, Config.Experience.Conversation.WebSearch<br>【Experience.Conversation.AiCall】Config.Experience.Conversation.AiCall.VoiceInteract, Config.Experience.Conversation.AiCall.VoiceCall, Config.Experience.Conversation.AiCall.DigitalHuman<br>【Experience.Advanced】Config.Experience.Advanced.ContextRewrite, Config.Experience.Advanced.ImageTextRetrieval, Config.Experience.Advanced.IntentAchievement, Config.Experience.Advanced.ReplyFlexibility, Config.Experience.Advanced.DialogCustomConfig</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`
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
	delete(f, "AppId")
	delete(f, "AppMode")
	delete(f, "Avatar")
	delete(f, "Config")
	delete(f, "Description")
	delete(f, "Name")
	delete(f, "SharedKbIdList")
	delete(f, "UpdateMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppResponseParams struct {
	// <p>app_id</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>更新时间 (Unix时间戳,秒级)</p>
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
type ModifyAppTriggerRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>触发器信息</p>
	Trigger *AppTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>触发器唯一ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>修改字段</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ModifyAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>触发器信息</p>
	Trigger *AppTrigger `json:"Trigger,omitnil,omitempty" name:"Trigger"`

	// <p>触发器唯一ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>修改字段</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ModifyAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "Trigger")
	delete(f, "TriggerId")
	delete(f, "UpdateMask")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAppTriggerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppTriggerResponseParams `json:"Response"`
}

func (r *ModifyAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConversationRequestParams struct {
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 会话ID
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// 会话标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

type ModifyConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 会话ID
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// 会话标题
	Title *string `json:"Title,omitnil,omitempty" name:"Title"`
}

func (r *ModifyConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Type")
	delete(f, "AppId")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	delete(f, "ConversationId")
	delete(f, "Title")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConversationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConversationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConversationResponseParams `json:"Response"`
}

func (r *ModifyConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMsgRecordCategoryRequestParams struct {
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>待修改的分类业务 ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>修改后的分类名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ModifyMsgRecordCategoryRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用 ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <p>待修改的分类业务 ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>修改后的分类名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *ModifyMsgRecordCategoryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMsgRecordCategoryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "CategoryId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyMsgRecordCategoryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyMsgRecordCategoryResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyMsgRecordCategoryResponse struct {
	*tchttp.BaseResponse
	Response *ModifyMsgRecordCategoryResponseParams `json:"Response"`
}

func (r *ModifyMsgRecordCategoryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyMsgRecordCategoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginRequestParams struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>插件版本号</p>
	PluginVersion *int64 `json:"PluginVersion,omitnil,omitempty" name:"PluginVersion"`

	// <p>插件基础资料</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>插件类型配置</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>指定需要更新的字段，避免全量覆盖</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>插件的工具列表，mcp插件不传</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

type ModifyPluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>插件版本号</p>
	PluginVersion *int64 `json:"PluginVersion,omitnil,omitempty" name:"PluginVersion"`

	// <p>插件基础资料</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>插件类型配置</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>指定需要更新的字段，避免全量覆盖</p>
	UpdateMask *FieldMask `json:"UpdateMask,omitnil,omitempty" name:"UpdateMask"`

	// <p>插件的工具列表，mcp插件不传</p>
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// <p>登录用户主账号(集成商模式必填)</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>登录用户子账号(集成商模式必填)</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`
}

func (r *ModifyPluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "PluginVersion")
	delete(f, "Profile")
	delete(f, "Config")
	delete(f, "UpdateMask")
	delete(f, "ToolList")
	delete(f, "LoginUin")
	delete(f, "LoginSubAccountUin")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyPluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyPluginResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyPluginResponse struct {
	*tchttp.BaseResponse
	Response *ModifyPluginResponseParams `json:"Response"`
}

func (r *ModifyPluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyPluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill描述</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill名称</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>skill包文件地址（zip）；传入则触发新版本生成，需与SkillVersion、UpdateDescription配套传入</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>图标地址</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>skill版本号（与FileUrl配套传入）</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>版本变更说明（与FileUrl配套传入）</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

type ModifySkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>skill描述</p>
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// <p>skill名称</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>skill包文件地址（zip）；传入则触发新版本生成，需与SkillVersion、UpdateDescription配套传入</p>
	FileUrl *string `json:"FileUrl,omitnil,omitempty" name:"FileUrl"`

	// <p>图标地址</p>
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// <p>skill版本号（与FileUrl配套传入）</p>
	SkillVersion *string `json:"SkillVersion,omitnil,omitempty" name:"SkillVersion"`

	// <p>版本变更说明（与FileUrl配套传入）</p>
	UpdateDescription *string `json:"UpdateDescription,omitnil,omitempty" name:"UpdateDescription"`
}

func (r *ModifySkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "DisplayDescription")
	delete(f, "DisplayName")
	delete(f, "FileUrl")
	delete(f, "IconUrl")
	delete(f, "SkillVersion")
	delete(f, "UpdateDescription")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySkillResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySkillResponse struct {
	*tchttp.BaseResponse
	Response *ModifySkillResponseParams `json:"Response"`
}

func (r *ModifySkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpaceRequestParams struct {
	// 工作空间名称,长度最大30个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 空间描述，长度最大150个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 指定需要更新的字段，支持Name和Description
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`
}

type ModifySpaceRequest struct {
	*tchttp.BaseRequest
	
	// 工作空间名称,长度最大30个字符
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 空间描述，长度最大150个字符
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 指定需要更新的字段，支持Name和Description
	FieldMask *FieldMask `json:"FieldMask,omitnil,omitempty" name:"FieldMask"`
}

func (r *ModifySpaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "SpaceId")
	delete(f, "FieldMask")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifySpaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifySpaceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifySpaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifySpaceResponseParams `json:"Response"`
}

func (r *ModifySpaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifySpaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVariableRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 变量信息
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

type ModifyVariableRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 变量信息
	Variable *Variable `json:"Variable,omitnil,omitempty" name:"Variable"`
}

func (r *ModifyVariableRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVariableRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Variable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyVariableRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyVariableResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyVariableResponse struct {
	*tchttp.BaseResponse
	Response *ModifyVariableResponseParams `json:"Response"`
}

func (r *ModifyVariableResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyVariableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgRecord struct {
	// 答案
	Answer *string `json:"Answer,omitnil,omitempty" name:"Answer"`

	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 分类ID
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 意图
	Intent *string `json:"Intent,omitnil,omitempty" name:"Intent"`

	// 意图分类
	IntentCategory *string `json:"IntentCategory,omitnil,omitempty" name:"IntentCategory"`

	// 是否是智能分类
	IsSmart *bool `json:"IsSmart,omitnil,omitempty" name:"IsSmart"`

	// 问题
	Question *string `json:"Question,omitnil,omitempty" name:"Question"`

	// 记录ID
	RecordId *string `json:"RecordId,omitnil,omitempty" name:"RecordId"`

	// 表示消息的回复方式，枚举 ReplyMethod：0=未指定, 1=大模型直接回复, 2=保守回复, 3=拒答, 4=敏感回复, 5=问答对优先回复, 6=欢迎语, 7=并发超限, 8=全局干预知识, 9=任务流程过程回复, 10=任务流程答案, 11=搜索引擎, 12=知识润色, 13=图片理解, 14=实时文档, 15=澄清确认, 16=工作流回复, 17=工作流结束, 18=智能体回复, 19=多意图, 20=中断, 21=智能体计划预览, 22=智能体计划结果, 23=智能体结构化输出。
	ReplyMethod *int64 `json:"ReplyMethod,omitnil,omitempty" name:"ReplyMethod"`

	// 返回结果
	Result *MsgRecordResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 分数
	Score *uint64 `json:"Score,omitnil,omitempty" name:"Score"`

	// 会话ID
	SessionId *string `json:"SessionId,omitnil,omitempty" name:"SessionId"`

	// 来源
	Source *MsgRecordSource `json:"Source,omitnil,omitempty" name:"Source"`

	// trace_id
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`
}

type MsgRecordCategory struct {
	// <p>分类的业务 ID</p>
	CategoryId *string `json:"CategoryId,omitnil,omitempty" name:"CategoryId"`

	// <p>子分类列表，树形嵌套</p>
	Children []*MsgRecordCategory `json:"Children,omitnil,omitempty" name:"Children"`

	// <p>分类名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>当前用户对该分类的操作权限</p>
	Permission *CategoryPermission `json:"Permission,omitnil,omitempty" name:"Permission"`

	// <p>该分类下消息记录的数量</p>
	TotalCount *string `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type MsgRecordResult struct {
	// 表示该条消息的调用结果：0=CALL_RESULT_UNKNOWN（全部/未知）, 1=CALL_RESULT_SUCCESS（成功）, 2=CALL_RESULT_FAILED（失败）；fail_reason（string）为调用失败时的失败原因描述。
	CallResult *int64 `json:"CallResult,omitnil,omitempty" name:"CallResult"`

	// 自定义变量，json字符串
	CustomerVariable *string `json:"CustomerVariable,omitnil,omitempty" name:"CustomerVariable"`

	// 失败原因
	FailReason *string `json:"FailReason,omitnil,omitempty" name:"FailReason"`

	// 首token耗时
	FirstTokenLatency *uint64 `json:"FirstTokenLatency,omitnil,omitempty" name:"FirstTokenLatency"`

	// 输入token数
	InputToken *uint64 `json:"InputToken,omitnil,omitempty" name:"InputToken"`

	// 输出token数
	OutputToken *uint64 `json:"OutputToken,omitnil,omitempty" name:"OutputToken"`

	// 总token数
	TotalToken *uint64 `json:"TotalToken,omitnil,omitempty" name:"TotalToken"`

	// 总token耗时
	TotalTokenLatency *uint64 `json:"TotalTokenLatency,omitnil,omitempty" name:"TotalTokenLatency"`
}

type MsgRecordSource struct {
	// 对话消息的接入渠道类型：0=未指定, 1=坐席, 2=体验页面(腾讯云), 3=评测端对话, 4=体验页面(手机号), 5=对话端API接入, 6=评测任务对话, 10=工作流调试, 10000=微信公众号, 10001=微信服务号, 10002=企微应用, 10003=网页组件, 10004=微信客服, 10005=微信小程序, 10006=元器, 10007=应用宝, 10008=元宝, 10009=企微智能机器人, 10010=元器API, 10011=LINE, 10012=Telegram, 10100=电脑管家, 20001=荣耀智能体平台, 20002=小米应用商店；user_id（string）为该渠道下的访客唯一标识。
	ChannelType *uint64 `json:"ChannelType,omitnil,omitempty" name:"ChannelType"`

	// 用户ID
	FromId *string `json:"FromId,omitnil,omitempty" name:"FromId"`

	// 消息发送者的用户来源类型：1=用户（访客/C端用户）, 2=机器人（AI回复）, 3=坐席（人工客服）；from_id（string）为该来源类型下的用户唯一标识 ID。
	FromType *uint64 `json:"FromType,omitnil,omitempty" name:"FromType"`

	// 用户头像
	UserAvatar *string `json:"UserAvatar,omitnil,omitempty" name:"UserAvatar"`

	// 访客ID
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// 访客名称
	UserNickname *string `json:"UserNickname,omitnil,omitempty" name:"UserNickname"`
}

type MultiAgentConfig struct {
	// Agent协同配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentCollaboration *AgentCollaborationConfig `json:"AgentCollaboration,omitnil,omitempty" name:"AgentCollaboration"`
}

type MultiModalQAModel struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type MultiModalUnderstandingModel struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type OAuthConfig struct {
	// OAuth服务方授权页url地址
	AuthorizationUrl *string `json:"AuthorizationUrl,omitnil,omitempty" name:"AuthorizationUrl"`

	// 客户端ID
	ClientId *string `json:"ClientId,omitnil,omitempty" name:"ClientId"`

	// 客户端密钥
	ClientSecret *string `json:"ClientSecret,omitnil,omitempty" name:"ClientSecret"`

	// 请求授权的数据范围
	ScopeList []*string `json:"ScopeList,omitnil,omitempty" name:"ScopeList"`

	// 获取access token的url地址
	TokenUrl *string `json:"TokenUrl,omitnil,omitempty" name:"TokenUrl"`
}

type OnceSchedule struct {
	// 触发时间
	FireTime *string `json:"FireTime,omitnil,omitempty" name:"FireTime"`
}

// Predefined struct for user
type PauseAppTriggerRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type PauseAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *PauseAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PauseAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PauseAppTriggerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PauseAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *PauseAppTriggerResponseParams `json:"Response"`
}

func (r *PauseAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PauseAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PlatformUsageSummary struct {
	// <p>PLATFORM 域消耗计量列表（权威字段）：按单位+label 分项列出每类计量，label 取 PlatformBizType 枚举名称字符串；典型如 unit=TIMES + label=PLATFORM_BIZ_TYPE_SECURITY_AUDIT/WEB_SEARCH/OPEN_CLAW/APP_INVOKE，unit=ITEM + label=PLATFORM_BIZ_TYPE_LONG_TERM_MEMORY</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`
}

type Plugin struct {
	// 插件配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 创建时间，unix时间戳
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 插件运营管理信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operation *PluginOperation `json:"Operation,omitnil,omitempty" name:"Operation"`

	// 插件id
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// 插件版本号
	PluginVersion *int64 `json:"PluginVersion,omitnil,omitempty" name:"PluginVersion"`

	// 插件基础信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// 插件统计信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Statistics *PluginStatistics `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// <p>插件状态，1:可用，2:不可用 </p><p>枚举值：</p><ul><li>1： 可用</li><li>2： 不可用</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 工具列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolList []*Tool `json:"ToolList,omitnil,omitempty" name:"ToolList"`

	// 更新时间，Unix时间戳
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 用户维度的插件状态信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserState *PluginUserState `json:"UserState,omitnil,omitempty" name:"UserState"`
}

type PluginConfig struct {
	// API插件配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApiPluginConfig *ApiPluginConfig `json:"ApiPluginConfig,omitnil,omitempty" name:"ApiPluginConfig"`

	// 应用插件配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppPluginConfig *AppPluginConfig `json:"AppPluginConfig,omitnil,omitempty" name:"AppPluginConfig"`

	// mcp插件配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	MCPPluginConfig *MCPPluginConfig `json:"MCPPluginConfig,omitnil,omitempty" name:"MCPPluginConfig"`
}

type PluginOperation struct {
	// 是否允许外部调用
	AllowExternalAccess *bool `json:"AllowExternalAccess,omitnil,omitempty" name:"AllowExternalAccess"`

	// <p>计费类型。</p><p>枚举值：</p><ul><li>0：免费</li><li>1：公测</li><li>2：官方收费</li></ul>
	BillingType *int64 `json:"BillingType,omitnil,omitempty" name:"BillingType"`

	// 插件分类标识
	CategoryKey *string `json:"CategoryKey,omitnil,omitempty" name:"CategoryKey"`

	// 插件概述
	Introduction *string `json:"Introduction,omitnil,omitempty" name:"Introduction"`

	// 是否精选
	IsRecommended *bool `json:"IsRecommended,omitnil,omitempty" name:"IsRecommended"`
}

type PluginParam struct {
	// 参数配置是否隐藏不可见
	IsGlobalHidden *bool `json:"IsGlobalHidden,omitnil,omitempty" name:"IsGlobalHidden"`

	// 参数是否必填
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// 参数名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type PluginProfile struct {
	// 插件作者
	Author *string `json:"Author,omitnil,omitempty" name:"Author"`

	// 插件描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 插件图标url
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// 插件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>插件产品分类</p><p>枚举值：</p><ul><li>0：普通插件</li><li>1：连接器类插件</li></ul>
	PluginClass *int64 `json:"PluginClass,omitnil,omitempty" name:"PluginClass"`

	// <p>插件类型</p><p>枚举值：</p><ul><li>0：API接口</li><li>1：代码</li><li>2：MCP</li><li>3：应用</li></ul>
	PluginKind *int64 `json:"PluginKind,omitnil,omitempty" name:"PluginKind"`

	// <p>插件来源</p><p>枚举值：</p><ul><li>0：自定义插件</li><li>1：官方插件</li><li>2：第三方插件</li></ul>
	PluginSource *int64 `json:"PluginSource,omitnil,omitempty" name:"PluginSource"`
}

type PluginStatistics struct {
	// 插件调用量
	CallCount *uint64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// 工具数量
	ToolCount *int64 `json:"ToolCount,omitnil,omitempty" name:"ToolCount"`
}

type PluginSummary struct {
	// <p>插件运营管理信息</p>
	Operation *PluginOperation `json:"Operation,omitnil,omitempty" name:"Operation"`

	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>插件基础信息</p>
	Profile *PluginProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// <p>插件统计信息</p>
	Statistics *PluginStatistics `json:"Statistics,omitnil,omitempty" name:"Statistics"`

	// <p>插件状态，1:可用，2:不可用 </p><p>枚举值：</p><ul><li>1： 可用</li><li>2： 不可用</li></ul>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>用户维度的插件状态信息</p>
	UserState *PluginUserState `json:"UserState,omitnil,omitempty" name:"UserState"`

	// <p>插件配置信息</p>
	Config *PluginConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// <p>工具信息</p>
	ToolList []*ToolSummary `json:"ToolList,omitnil,omitempty" name:"ToolList"`
}

type PluginUsageDetail struct {
	// <p>插件名称</p>
	PluginName *string `json:"PluginName,omitnil,omitempty" name:"PluginName"`

	// <p>PLUGIN 域单次调用的消耗计量列表（权威字段）：按单位+label 分项列出每类计量。unit=TOKEN 时 label 区分 Token 子类别（input/output/avg_*），label 为空表示 total_tokens</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`

	// <p>插件工具名（tool_name）</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`
}

type PluginUsageSummary struct {
	// <p>调用次数（业务调用维度的顶层计数）</p>
	CallCount *float64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// <p>PLUGIN 域消耗计量列表（权威字段）：按单位+label 分项列出每类计量。unit=TOKEN 时 label 区分 Token 子类别（input/output/avg_*），label 为空表示 total_tokens</p>
	ResourceConsumptionList []*ResourceConsumption `json:"ResourceConsumptionList,omitnil,omitempty" name:"ResourceConsumptionList"`
}

type PluginUserState struct {
	// 是否已收藏该插件
	IsFavorite *bool `json:"IsFavorite,omitnil,omitempty" name:"IsFavorite"`

	// 是否在插件白名单内
	IsInWhiteList *bool `json:"IsInWhiteList,omitnil,omitempty" name:"IsInWhiteList"`

	// <p>白名单类型，用于表示当前用户是否可直接使用该插件。</p><p>枚举值：</p><ul><li>0：非白名单插件，全量开放</li><li>1：当前用户在白名单内</li><li>2：当前用户不在白名单内，需提交申请</li></ul>
	WhiteListType *int64 `json:"WhiteListType,omitnil,omitempty" name:"WhiteListType"`
}

type PromptRewriteModel struct {
	// 模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type ReleaseRecord struct {
	// 是否可导出
	CanExport *bool `json:"CanExport,omitnil,omitempty" name:"CanExport"`

	// 是否可回滚
	CanRollback *bool `json:"CanRollback,omitnil,omitempty" name:"CanRollback"`

	// 发布描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 发布失败数
	FailCount *uint64 `json:"FailCount,omitnil,omitempty" name:"FailCount"`

	// 失败原因
	Reason *string `json:"Reason,omitnil,omitempty" name:"Reason"`

	// 发布ID
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`

	// 发布版本
	ReleaseVersion *string `json:"ReleaseVersion,omitnil,omitempty" name:"ReleaseVersion"`

	// 发布状态。枚举值: 1:待发布, 2:发布中, 3:发布成功, 4:发布失败, 5:审核中, 6:审核成功, 7:审核失败, 8:发布成功回调处理中, 9:发布暂停, 10:申诉审核中, 11:申诉审核通过, 12:申诉审核不通过
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 状态描述
	StatusDescription *string `json:"StatusDescription,omitnil,omitempty" name:"StatusDescription"`

	// 发布成功数
	SuccessCount *uint64 `json:"SuccessCount,omitnil,omitempty" name:"SuccessCount"`

	// 更新时间 (Unix时间戳,秒级)
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// 发布人
	Updater *string `json:"Updater,omitnil,omitempty" name:"Updater"`
}

// Predefined struct for user
type ReleaseSkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>版本ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

type ReleaseSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>版本ID</p>
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`
}

func (r *ReleaseSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	delete(f, "VersionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReleaseSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReleaseSkillResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ReleaseSkillResponse struct {
	*tchttp.BaseResponse
	Response *ReleaseSkillResponseParams `json:"Response"`
}

func (r *ReleaseSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReleaseSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseSummary struct {
	// <p>创建时间 (Unix时间戳,秒级)</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>发布描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>发布ID</p>
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`

	// <p>发布状态。枚举值: 1:待发布, 2:发布中, 3:发布成功, 4:发布失败, 5:审核中, 6:审核成功, 7:审核失败, 8:发布成功回调处理中, 9:发布暂停, 10:申诉审核中, 11:申诉审核通过, 12:申诉审核不通过</p>
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>状态描述</p>
	StatusDescription *string `json:"StatusDescription,omitnil,omitempty" name:"StatusDescription"`

	// <p>应用分享访问控制</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppShareAccessControl *AppShareAccessControl `json:"AppShareAccessControl,omitnil,omitempty" name:"AppShareAccessControl"`

	// <p>发布渠道ID列表</p>
	ChannelIdList []*string `json:"ChannelIdList,omitnil,omitempty" name:"ChannelIdList"`

	// <p>企业共享配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CorpShareConfig *CorpShareConfig `json:"CorpShareConfig,omitnil,omitempty" name:"CorpShareConfig"`
}

type RequestParam struct {
	// <p>AnyOf类型的参数</p>
	AnyOf []*RequestParam `json:"AnyOf,omitnil,omitempty" name:"AnyOf"`

	// <p>默认值</p>
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// <p>参数描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>全局隐藏不可见（区别于Agent场景的agent_hidden），true-全局隐藏不可见，false-可见</p>
	IsGlobalHidden *bool `json:"IsGlobalHidden,omitnil,omitempty" name:"IsGlobalHidden"`

	// <p>是否必选</p>
	IsRequired *bool `json:"IsRequired,omitnil,omitempty" name:"IsRequired"`

	// <p>参数名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>OneOf类型的参数</p>
	OneOf []*RequestParam `json:"OneOf,omitnil,omitempty" name:"OneOf"`

	// <p>子参数,ParamType 是OBJECT 或 ARRAY&lt;&gt;类型有用</p>
	SubParams []*RequestParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>PARAM_TYPE_STRING</td><td>0</td><td>字符串</td></tr><tr><td>PARAM_TYPE_INT</td><td>1</td><td>整数</td></tr><tr><td>PARAM_TYPE_FLOAT</td><td>2</td><td>浮点数</td></tr><tr><td>PARAM_TYPE_BOOL</td><td>3</td><td>布尔值</td></tr><tr><td>PARAM_TYPE_OBJECT</td><td>4</td><td>对象</td></tr><tr><td>PARAM_TYPE_ARRAY_STRING</td><td>5</td><td>字符串数组</td></tr><tr><td>PARAM_TYPE_ARRAY_INT</td><td>6</td><td>整数数组</td></tr><tr><td>PARAM_TYPE_ARRAY_FLOAT</td><td>7</td><td>浮点数数组</td></tr><tr><td>PARAM_TYPE_ARRAY_BOOL</td><td>8</td><td>布尔值数组</td></tr><tr><td>PARAM_TYPE_ARRAY_OBJECT</td><td>9</td><td>对象数组</td></tr><tr><td>PARAM_TYPE_ARRAY_ARRAY</td><td>20</td><td>数组嵌套</td></tr><tr><td>PARAM_TYPE_NULL</td><td>99</td><td>空值</td></tr><tr><td>PARAM_TYPE_UNSPECIFIED</td><td>100</td><td>未指定类型，用于OneOf和AnyOf场景</td></tr></tbody></table>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type ResetConversationRequestParams struct {
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ResetConversationRequest struct {
	*tchttp.BaseRequest
	
	// <p>会话 ID</p>
	ConversationId *string `json:"ConversationId,omitnil,omitempty" name:"ConversationId"`

	// <p>会话类型 枚举值: 0-CONVERSATION_TYPE_UNSPECIFIED(未指定；列表查询时表示全部), 1-CONVERSATION_TYPE_VISITOR(访客端体验), 2-CONVERSATION_TYPE_EVALUATION(评测), 5-CONVERSATION_TYPE_API(API 接入), 10-CONVERSATION_TYPE_WORKFLOW(工作流调试), 20-CONVERSATION_TYPE_SHARE(分享链接)</p>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，应用密钥</p>
	AppKey *string `json:"AppKey,omitnil,omitempty" name:"AppKey"`

	// <p>子用户Uin</p>
	LoginSubAccountUin *string `json:"LoginSubAccountUin,omitnil,omitempty" name:"LoginSubAccountUin"`

	// <p>主用户Uin</p>
	LoginUin *string `json:"LoginUin,omitnil,omitempty" name:"LoginUin"`

	// <p>Type=CONVERSATION_TYPE_SHARE 时必填，分享码</p>
	ShareCode *string `json:"ShareCode,omitnil,omitempty" name:"ShareCode"`

	// <p>Type=CONVERSATION_TYPE_API 时必填，访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ResetConversationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConversationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ConversationId")
	delete(f, "Type")
	delete(f, "AppKey")
	delete(f, "LoginSubAccountUin")
	delete(f, "LoginUin")
	delete(f, "ShareCode")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetConversationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResetConversationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResetConversationResponse struct {
	*tchttp.BaseResponse
	Response *ResetConversationResponseParams `json:"Response"`
}

func (r *ResetConversationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetConversationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceConsumption struct {
	// <p>功能标签，PLATFORM 场景取 PlatformBizType 枚举名称；MODEL/PLUGIN 场景为空</p>
	Label *string `json:"Label,omitnil,omitempty" name:"Label"`

	// <p>消耗计量单位</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>DOSAGE_UNIT_TOKEN</td><td>0</td><td>token（默认）</td></tr><tr><td>DOSAGE_UNIT_PAGE_COUNT</td><td>1</td><td>page_count（页数）</td></tr><tr><td>DOSAGE_UNIT_TIMES</td><td>2</td><td>times（次数）</td></tr><tr><td>DOSAGE_UNIT_SECOND</td><td>3</td><td>second（秒）</td></tr><tr><td>DOSAGE_UNIT_ITEM</td><td>4</td><td>item（条）</td></tr><tr><td>DOSAGE_UNIT_SHEET</td><td>5</td><td>sheet（张）</td></tr><tr><td>DOSAGE_UNIT_CHARACTER</td><td>6</td><td>character（字符）</td></tr><tr><td>DOSAGE_UNIT_GB</td><td>7</td><td>GB</td></tr><tr><td>DOSAGE_UNIT_NUMBER</td><td>8</td><td>number（个数）</td></tr><tr><td>DOSAGE_UNIT_MILL_SECOND</td><td>9</td><td>mill_second（毫秒）</td></tr></tbody></table>
	Unit *int64 `json:"Unit,omitnil,omitempty" name:"Unit"`

	// <p>消耗数值</p>
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type ResponseParam struct {
	// <p>变量描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>参数名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>OUTPUT_RENDER_REPLACE</td><td>0</td><td>覆盖（全量替换）</td></tr><tr><td>OUTPUT_RENDER_APPEND</td><td>1</td><td>增量追加</td></tr></tbody></table>
	RenderMode *int64 `json:"RenderMode,omitnil,omitempty" name:"RenderMode"`

	// <p>只对 OBJECT 或 ARRAY_OBJECT 类型有用</p>
	SubParams []*ResponseParam `json:"SubParams,omitnil,omitempty" name:"SubParams"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>PARAM_TYPE_STRING</td><td>0</td><td>字符串</td></tr><tr><td>PARAM_TYPE_INT</td><td>1</td><td>整数</td></tr><tr><td>PARAM_TYPE_FLOAT</td><td>2</td><td>浮点数</td></tr><tr><td>PARAM_TYPE_BOOL</td><td>3</td><td>布尔值</td></tr><tr><td>PARAM_TYPE_OBJECT</td><td>4</td><td>对象</td></tr><tr><td>PARAM_TYPE_ARRAY_STRING</td><td>5</td><td>字符串数组</td></tr><tr><td>PARAM_TYPE_ARRAY_INT</td><td>6</td><td>整数数组</td></tr><tr><td>PARAM_TYPE_ARRAY_FLOAT</td><td>7</td><td>浮点数数组</td></tr><tr><td>PARAM_TYPE_ARRAY_BOOL</td><td>8</td><td>布尔值数组</td></tr><tr><td>PARAM_TYPE_ARRAY_OBJECT</td><td>9</td><td>对象数组</td></tr><tr><td>PARAM_TYPE_ARRAY_ARRAY</td><td>20</td><td>数组嵌套</td></tr><tr><td>PARAM_TYPE_NULL</td><td>99</td><td>空值</td></tr><tr><td>PARAM_TYPE_UNSPECIFIED</td><td>100</td><td>未指定类型，用于OneOf和AnyOf场景</td></tr></tbody></table>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

// Predefined struct for user
type ResumeAppTriggerRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type ResumeAppTriggerRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *ResumeAppTriggerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeAppTriggerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeAppTriggerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeAppTriggerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeAppTriggerResponse struct {
	*tchttp.BaseResponse
	Response *ResumeAppTriggerResponseParams `json:"Response"`
}

func (r *ResumeAppTriggerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeAppTriggerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RetryReleaseRequestParams struct {
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 发布任务ID
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

type RetryReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 发布任务ID
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
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
	delete(f, "AppId")
	delete(f, "ReleaseId")
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

type RoleConfig struct {
	// 角色描述
	RoleDescription *string `json:"RoleDescription,omitnil,omitempty" name:"RoleDescription"`
}

// Predefined struct for user
type RollbackReleaseRequestParams struct {
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

type RollbackReleaseRequest struct {
	*tchttp.BaseRequest
	
	// app_id
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// release_id
	ReleaseId *string `json:"ReleaseId,omitnil,omitempty" name:"ReleaseId"`
}

func (r *RollbackReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "ReleaseId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackReleaseResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackReleaseResponse struct {
	*tchttp.BaseResponse
	Response *RollbackReleaseResponseParams `json:"Response"`
}

func (r *RollbackReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunAppTriggerNowRequestParams struct {
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type RunAppTriggerNowRequest struct {
	*tchttp.BaseRequest
	
	// <p>应用ID</p>
	AppId *string `json:"AppId,omitnil,omitempty" name:"AppId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>应用触发器ID</p>
	TriggerId *string `json:"TriggerId,omitnil,omitempty" name:"TriggerId"`

	// <p>访客ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

func (r *RunAppTriggerNowRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunAppTriggerNowRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppId")
	delete(f, "Scope")
	delete(f, "TriggerId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RunAppTriggerNowRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RunAppTriggerNowResponseParams struct {
	// <p>应用触发器实例ID</p>
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RunAppTriggerNowResponse struct {
	*tchttp.BaseResponse
	Response *RunAppTriggerNowResponseParams `json:"Response"`
}

func (r *RunAppTriggerNowResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RunAppTriggerNowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchResourceStatusInfo struct {
	// 搜索资源状态: AVAILABLE(1)=资源可用, EXHAUSTED(2)=资源已用尽。枚举值: 1:资源可用, 2:资源已用尽
	ResourceStatus *int64 `json:"ResourceStatus,omitnil,omitempty" name:"ResourceStatus"`
}

type SingleWorkflowConfig struct {
	// <p>是否开启异步工作流</p>
	AsyncWorkflow *bool `json:"AsyncWorkflow,omitnil,omitempty" name:"AsyncWorkflow"`

	// <p>状态 发布状态(UNPUBLISHED: 待发布 PUBLISHING: 发布中 PUBLISHED: 已发布 PUBLISHED_FAIL:发布失败；DRAFT：待调试)</p>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>工作流描述</p>
	WorkflowDescription *string `json:"WorkflowDescription,omitnil,omitempty" name:"WorkflowDescription"`

	// <p>工作流Id</p>
	WorkflowId *string `json:"WorkflowId,omitnil,omitempty" name:"WorkflowId"`

	// <p>工作流名称</p>
	WorkflowName *string `json:"WorkflowName,omitnil,omitempty" name:"WorkflowName"`

	// <p>工作流是否启用</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type SkillAnalysisInfo struct {
	// 安全检测状态
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 待检测 |
	// | 1 | 检测中 |
	// | 2 | 可用 |
	// | 3 | 不可用 |
	// | 4 | 检测失败 |
	AnalysisStatus *int64 `json:"AnalysisStatus,omitnil,omitempty" name:"AnalysisStatus"`

	// 风险描述
	RiskDescription *string `json:"RiskDescription,omitnil,omitempty" name:"RiskDescription"`

	// 风险等级
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 无风险 |
	// | 1 | 低风险 |
	// | 2 | 中风险 |
	// | 3 | 高风险 |
	RiskLevel *int64 `json:"RiskLevel,omitnil,omitempty" name:"RiskLevel"`

	// 安全报告跳转url;
	SecurityReportUrl *string `json:"SecurityReportUrl,omitnil,omitempty" name:"SecurityReportUrl"`
}

type SkillCategory struct {
	// 分类标识
	CategoryKey *string `json:"CategoryKey,omitnil,omitempty" name:"CategoryKey"`

	// 分类名称
	CategoryName *string `json:"CategoryName,omitnil,omitempty" name:"CategoryName"`
}

type SkillClassification struct {
	// Skill 计费类型
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 免费 |
	// | 1 | 付费 |
	BillingType *int64 `json:"BillingType,omitnil,omitempty" name:"BillingType"`

	// Skill 内置来源，仅在 create_type 为 SKILL_CREATE_TYPE_BUILTIN 时生效
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 占位 |
	// | 1 | ADP 专有 |
	// | 2 | 腾讯专有 |
	// | 3 | SkillHub |
	// | 99 | 其他 |
	BuiltinSource *int64 `json:"BuiltinSource,omitnil,omitempty" name:"BuiltinSource"`

	// Skill 分类
	CategoryKey *string `json:"CategoryKey,omitnil,omitempty" name:"CategoryKey"`

	// Skill 创建方式
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 占位 |
	// | 1 | 文件上传 |
	// | 2 | 由企业级共享流程生成 |
	// | 3 | AIGC 生成 |
	// | 99 | 内置 Skill |
	CreateType *int64 `json:"CreateType,omitnil,omitempty" name:"CreateType"`

	// Skill 提供方类型
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 占位 |
	// | 1 | 官方 |
	// | 2 | 第三方 |
	// | 3 | 自定义 |
	// | 4 | 自定义企业级共享 |
	ProviderType *int64 `json:"ProviderType,omitnil,omitempty" name:"ProviderType"`

	// Skill 来源链接
	SourceLink *string `json:"SourceLink,omitnil,omitempty" name:"SourceLink"`
}

type SkillDetail struct {
	// 调用情况摘要
	ReferenceSummaryList []*SkillReferenceSummary `json:"ReferenceSummaryList,omitnil,omitempty" name:"ReferenceSummaryList"`

	// Skill 摘要
	SkillSummary *SkillSummary `json:"SkillSummary,omitnil,omitempty" name:"SkillSummary"`

	// 版本列表
	VersionList []*SkillVersion `json:"VersionList,omitnil,omitempty" name:"VersionList"`
}

type SkillNotice struct {
	// 通知级别
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 占位 |
	// | 1 | 成功，字符串面："success" |
	// | 2 | 警告，字符串面："warning" |
	// | 3 | 错误，字符串面："error" |
	Level *int64 `json:"Level,omitnil,omitempty" name:"Level"`

	// 文案（i18n 后字符串）
	NoticeContent *string `json:"NoticeContent,omitnil,omitempty" name:"NoticeContent"`

	// 触发本通知的 Skill 版本ID
	TriggerVersionId *string `json:"TriggerVersionId,omitnil,omitempty" name:"TriggerVersionId"`

	// 通知类型 
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 占位 |
	// | 1 | 发布失败 |
	// | 2 | 共享审批被拒 |
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`
}

type SkillProfile struct {
	// 创建时间（Unix秒）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建者
	Creator *string `json:"Creator,omitnil,omitempty" name:"Creator"`

	// Skill 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Skill 展示描述
	DisplayDescription *string `json:"DisplayDescription,omitnil,omitempty" name:"DisplayDescription"`

	// Skill 展示名称
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// Skill 图标
	IconUrl *string `json:"IconUrl,omitnil,omitempty" name:"IconUrl"`

	// Skill 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 更新时间（Unix秒）
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`
}

type SkillReferenceGroup struct {
	// <p>该类型下的引用详情列表</p>
	ReferenceSummaryList []*SkillReferenceSummary `json:"ReferenceSummaryList,omitnil,omitempty" name:"ReferenceSummaryList"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>SKILL_REF_UNKNOWN</td><td>0</td><td>占位</td></tr><tr><td>SKILL_REF_OPENCLAW</td><td>1</td><td>openclaw</td></tr><tr><td>SKILL_REF_AGENT</td><td>2</td><td>agent</td></tr><tr><td>SKILL_REF_CORP_ASSISTANT</td><td>3</td><td>企业助手</td></tr></tbody></table>
	ReferenceType *int64 `json:"ReferenceType,omitnil,omitempty" name:"ReferenceType"`

	// <p>该类型下的引用总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type SkillReferenceSummary struct {
	// <p>关联ID</p>
	ReferenceId *string `json:"ReferenceId,omitnil,omitempty" name:"ReferenceId"`

	// <p>关联名称</p>
	ReferenceName *string `json:"ReferenceName,omitnil,omitempty" name:"ReferenceName"`

	// <p>关联类型</p><p>枚举值:<br>| uint | 描述 |<br>| --- | --- |<br>| 0 | 占位 |<br>| 1 | ClawPro |<br>| 2 | agent |</p>
	ReferenceType *int64 `json:"ReferenceType,omitnil,omitempty" name:"ReferenceType"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// <p>空间名称</p>
	SpaceName *string `json:"SpaceName,omitnil,omitempty" name:"SpaceName"`

	// <p>Reference实例拥有者</p>
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`
}

type SkillShare struct {
	// 审批ID
	ApprovalId *string `json:"ApprovalId,omitnil,omitempty" name:"ApprovalId"`

	// 共享后关联的新 skill_id
	ShareSkillId *string `json:"ShareSkillId,omitnil,omitempty" name:"ShareSkillId"`

	// 共享版本，如 1.0.0
	ShareVersion *string `json:"ShareVersion,omitnil,omitempty" name:"ShareVersion"`

	// 共享版本ID
	ShareVersionId *string `json:"ShareVersionId,omitnil,omitempty" name:"ShareVersionId"`

	// 原 skill_id
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// 共享状态
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 未共享 |
	// | 1 | 已共享 |
	// | 2 | 审批中 |
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type SkillSummary struct {
	// 分类信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationInfo *SkillClassification `json:"ClassificationInfo,omitnil,omitempty" name:"ClassificationInfo"`

	// 当前版本信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentVersionInfo *SkillVersion `json:"CurrentVersionInfo,omitnil,omitempty" name:"CurrentVersionInfo"`

	// 当前用户是否收藏
	IsFavorite *bool `json:"IsFavorite,omitnil,omitempty" name:"IsFavorite"`

	// 基础信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Profile *SkillProfile `json:"Profile,omitnil,omitempty" name:"Profile"`

	// Skill ID
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// Skill 异常通知列表
	NoticeList []*SkillNotice `json:"NoticeList,omitnil,omitempty" name:"NoticeList"`

	// 当前用户对该 Skill 的资源操作权限位列表；内置/共享 Skill 固定为空数组
	PermissionIdList []*string `json:"PermissionIdList,omitnil,omitempty" name:"PermissionIdList"`

	// 共享信息；可能有两条，一条是已共享的，一条是审核中的
	ShareList []*SkillShare `json:"ShareList,omitnil,omitempty" name:"ShareList"`

	// Skill状态 
	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 | 初始化（无任何已发布版本，且最新版本处于 INITIALIZED/UNRELEASED） |
	// | 1 | 安全检测中（无任何已发布版本，且最新版本处于 AUDITING） |
	// | 2 | 待发布（无任何已发布版本，且最新版本处于 PENDING_RELEASE） |
	// | 3 | 已发布（存在任一 RELEASED 版本，吸收态） |
	SkillStatus *int64 `json:"SkillStatus,omitnil,omitempty" name:"SkillStatus"`
}

type SkillVersion struct {
	// 检测信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnalysisInfo *SkillAnalysisInfo `json:"AnalysisInfo,omitnil,omitempty" name:"AnalysisInfo"`

	// 当前生效版本号
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 当前生效版本ID
	VersionId *string `json:"VersionId,omitnil,omitempty" name:"VersionId"`

	//     Skill 版本发布流程状态：
	//       - 0 INITIALIZED      初始化（版本初始态）
	//       - 1 AUDITING         审核中（f_analysis_status ∈ {PENDING, RUNNING}）
	//       - 2 PENDING_RELEASE  待发布（低/中风险，等用户确认上架）
	//       - 3 RELEASED         已发布
	//       - 4 UNRELEASED       未发布（HIGH / UNAVAILABLE / FAILED / 用户放弃，含历史"不通过"语义）
	//     与 SkillAnalysisStatus 解耦：前者是用户视角发布生命周期，后者是安全检测阶段。
	VersionStatus *int64 `json:"VersionStatus,omitnil,omitempty" name:"VersionStatus"`

	// Skill包的md5信息
	SkillMd5 *string `json:"SkillMd5,omitnil,omitempty" name:"SkillMd5"`

	// 版本包地址
	SkillUrl *string `json:"SkillUrl,omitnil,omitempty" name:"SkillUrl"`

	// 版本创建时间（Unix秒）
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// skill md文档
	SkillMarkdownUrl *string `json:"SkillMarkdownUrl,omitnil,omitempty" name:"SkillMarkdownUrl"`

	// 版本变更说明
	UpdateDesc *string `json:"UpdateDesc,omitnil,omitempty" name:"UpdateDesc"`
}

type Sort struct {
	// <p>排序字段名，如 create_time</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>排序方向，1 升序，2 降序</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>SORT_ORDER_INVALID</td><td>0</td><td>无效</td></tr><tr><td>SORT_ORDER_ASC</td><td>1</td><td>升序</td></tr><tr><td>SORT_ORDER_DESC</td><td>2</td><td>降序</td></tr></tbody></table>
	Direction *int64 `json:"Direction,omitnil,omitempty" name:"Direction"`
}

type Space struct {
	// 空间id
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`

	// 空间名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 空间描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 空间权限
	PermissionIdList []*string `json:"PermissionIdList,omitnil,omitempty" name:"PermissionIdList"`
}

type SpecialStatusInfo struct {
	// 状态 (0-不在特殊状态中, 1-在特殊状态中)。枚举值: 1:在特殊状态中
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`
}

type SupportedFileType struct {
	// 文件类型描述(如"文本文档")
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 文件类型(如 txt、pdf、jpg, 建议用扩展名)
	FileExt *string `json:"FileExt,omitnil,omitempty" name:"FileExt"`

	// 文件大小限制(单位: 字节)
	MaxSizeBytes *string `json:"MaxSizeBytes,omitnil,omitempty" name:"MaxSizeBytes"`
}

type SystemVariable struct {
	// 变量描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 变量名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type ThinkModel struct {
	// 思考模型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Model *ModelDetailInfo `json:"Model,omitnil,omitempty" name:"Model"`
}

type TimeRange struct {
	// <p>结束时间，Unix 秒</p>
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>开始时间，Unix 秒</p>
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`
}

type TimerPushConfig struct {
	// <p>枚举值:<br>| uint | 描述 |<br>| --- | --- |<br>| 0 |  |<br>| 1 | 不推送 |<br>| 2 | 微信公众号 |<br>| 3 | 企业微信 AI 机器人 |</p>
	PushChannel *int64 `json:"PushChannel,omitnil,omitempty" name:"PushChannel"`

	// <p>推送会话ID</p>
	PushTargetId *string `json:"PushTargetId,omitnil,omitempty" name:"PushTargetId"`

	// <p>枚举值:<br>| uint | 描述 |<br>| --- | --- |<br>| 0 |  |<br>| 1 | 用户 (微信公众号 openid) |<br>| 2 | 群聊 (企微机器人 chat_id) |</p>
	PushTargetType *int64 `json:"PushTargetType,omitnil,omitempty" name:"PushTargetType"`

	// <p>推送webhook的url</p>
	PushWebhookUrl *string `json:"PushWebhookUrl,omitnil,omitempty" name:"PushWebhookUrl"`
}

type TimerScheduleConfig struct {
	// cron配置
	Cron *CronSchedule `json:"Cron,omitnil,omitempty" name:"Cron"`

	// 每日触发
	Daily *DailySchedule `json:"Daily,omitnil,omitempty" name:"Daily"`

	// 固定间隔
	Interval *IntervalSchedule `json:"Interval,omitnil,omitempty" name:"Interval"`

	// 仅手动
	ManualOnly *ManualOnlySchedule `json:"ManualOnly,omitnil,omitempty" name:"ManualOnly"`

	// 单次
	Once *OnceSchedule `json:"Once,omitnil,omitempty" name:"Once"`

	// 
	// 枚举值:
	// | uint | 描述 |
	// | --- | --- |
	// | 0 |  |
	// | 1 | 仅手动 |
	// | 2 | 每天 |
	// | 3 | 每周 |
	// | 4 | 按间隔 |
	// | 5 | 一次性 |
	// | 6 | Cron |
	ScheduleType *int64 `json:"ScheduleType,omitnil,omitempty" name:"ScheduleType"`

	// 时区
	Timezone *string `json:"Timezone,omitnil,omitempty" name:"Timezone"`

	// 每周固定时间触发
	Weekly *WeeklySchedule `json:"Weekly,omitnil,omitempty" name:"Weekly"`
}

type Tool struct {
	// <p>工具计费信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Billing *ToolBilling `json:"Billing,omitnil,omitempty" name:"Billing"`

	// <p>工具调用次数</p><p>单位：次数</p>
	CallCount *uint64 `json:"CallCount,omitnil,omitempty" name:"CallCount"`

	// <p>工具描述信息</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>工具名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>插件ID</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>TOOL_ACCESS_MODE_UNKNOWN</td><td>0</td><td>未指定</td></tr><tr><td>TOOL_ACCESS_MODE_READ_ONLY</td><td>1</td><td>只读</td></tr><tr><td>TOOL_ACCESS_MODE_WRITE_DELETE</td><td>2</td><td>写/删除</td></tr></tbody></table>
	ToolAccessMode *int64 `json:"ToolAccessMode,omitnil,omitempty" name:"ToolAccessMode"`

	// <p>工具配置信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ToolConfig *ToolConfig `json:"ToolConfig,omitnil,omitempty" name:"ToolConfig"`

	// <p>工具ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

type ToolBilling struct {
	// <p>基础计费信息</p>
	BasicBilling *BasicBilling `json:"BasicBilling,omitnil,omitempty" name:"BasicBilling"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>BILLING_TYPE_FREE</td><td>0</td><td>免费</td></tr><tr><td>BILLING_TYPE_LIMITED_FREE</td><td>1</td><td>限时免费</td></tr><tr><td>BILLING_TYPE_OFFICIAL_PAID</td><td>2</td><td>官方收费</td></tr><tr><td>BILLING_TYPE_OFFICIAL_PAID_OLD_FREE</td><td>3</td><td>官方收费（新/升级用户收费，存量老用户限时免费）</td></tr></tbody></table>
	BillingType *int64 `json:"BillingType,omitnil,omitempty" name:"BillingType"`

	// <p>复合类型计费信息</p>
	ComplexBilling *ComplexBilling `json:"ComplexBilling,omitnil,omitempty" name:"ComplexBilling"`

	// <p>双向计费信息</p>
	DuplexBilling *DuplexBilling `json:"DuplexBilling,omitnil,omitempty" name:"DuplexBilling"`
}

type ToolConfig struct {
	// <p>API工具配置信息</p>
	ApiToolConfig *ApiToolConfig `json:"ApiToolConfig,omitnil,omitempty" name:"ApiToolConfig"`

	// <p>应用配置信息</p>
	AppToolConfig *AppToolConfig `json:"AppToolConfig,omitnil,omitempty" name:"AppToolConfig"`

	// <p>代码工具配置信息</p>
	CodeToolConfig *CodeToolConfig `json:"CodeToolConfig,omitnil,omitempty" name:"CodeToolConfig"`

	// <p>MCP工具配置信息</p>
	MCPToolConfig *MCPToolConfig `json:"MCPToolConfig,omitnil,omitempty" name:"MCPToolConfig"`
}

type ToolExample struct {
	// <p>请求参数</p>
	Request *string `json:"Request,omitnil,omitempty" name:"Request"`

	// <p>响应参数</p>
	Response *string `json:"Response,omitnil,omitempty" name:"Response"`
}

type ToolSummary struct {
	// <p>工具Id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

type TriggerConfig struct {
	// <p>定时器配置</p>
	ScheduledConfig *AppTriggerScheduleConfig `json:"ScheduledConfig,omitnil,omitempty" name:"ScheduledConfig"`

	// <p>Webhook配置</p>
	WebhookConfig *AppTriggerWebhookConfig `json:"WebhookConfig,omitnil,omitempty" name:"WebhookConfig"`
}

type TriggerStatus struct {
	// <p>定时器状态</p>
	ScheduledStatus *AppTriggerScheduleStatus `json:"ScheduledStatus,omitnil,omitempty" name:"ScheduledStatus"`

	// <table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>APP_TRIGGER_SCOPE_UNSPECIFIED</td><td>0</td><td>未指定</td></tr><tr><td>APP_TRIGGER_SCOPE_APP</td><td>1</td><td>B 端管理员</td></tr><tr><td>APP_TRIGGER_SCOPE_USER</td><td>2</td><td>C 端访客</td></tr></tbody></table>
	Scope *int64 `json:"Scope,omitnil,omitempty" name:"Scope"`

	// <p>访客id</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`

	// <p>Webhook状态</p>
	WebhookStatus *AppTriggerWebhookStatus `json:"WebhookStatus,omitnil,omitempty" name:"WebhookStatus"`
}

// Predefined struct for user
type UnfavoritePluginRequestParams struct {
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type UnfavoritePluginRequest struct {
	*tchttp.BaseRequest
	
	// <p>插件id</p>
	PluginId *string `json:"PluginId,omitnil,omitempty" name:"PluginId"`

	// <p>当前空间id</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *UnfavoritePluginRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoritePluginRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PluginId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnfavoritePluginRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfavoritePluginResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnfavoritePluginResponse struct {
	*tchttp.BaseResponse
	Response *UnfavoritePluginResponseParams `json:"Response"`
}

func (r *UnfavoritePluginResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoritePluginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfavoriteSkillRequestParams struct {
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

type UnfavoriteSkillRequest struct {
	*tchttp.BaseRequest
	
	// <p>SkillId</p>
	SkillId *string `json:"SkillId,omitnil,omitempty" name:"SkillId"`

	// <p>空间ID</p>
	SpaceId *string `json:"SpaceId,omitnil,omitempty" name:"SpaceId"`
}

func (r *UnfavoriteSkillRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoriteSkillRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SkillId")
	delete(f, "SpaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnfavoriteSkillRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnfavoriteSkillResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnfavoriteSkillResponse struct {
	*tchttp.BaseResponse
	Response *UnfavoriteSkillResponseParams `json:"Response"`
}

func (r *UnfavoriteSkillResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnfavoriteSkillResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UsageDetail struct {
	// <p>调用来源</p>
	CallSource *CallSource `json:"CallSource,omitnil,omitempty" name:"CallSource"`

	// <p>计量 ID，用于对账/回溯</p>
	DosageId *string `json:"DosageId,omitnil,omitempty" name:"DosageId"`

	// <p>调用时间戳（Unix 秒）</p>
	EventTime *string `json:"EventTime,omitnil,omitempty" name:"EventTime"`

	// <p>MODEL 域专属</p>
	Model *ModelUsageDetail `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>PLUGIN 域专属</p>
	Plugin *PluginUsageDetail `json:"Plugin,omitnil,omitempty" name:"Plugin"`

	// <p>调用链路追踪 ID</p>
	TraceId *string `json:"TraceId,omitnil,omitempty" name:"TraceId"`

	// <p>用户 ID</p>
	UserId *string `json:"UserId,omitnil,omitempty" name:"UserId"`
}

type UsageSummary struct {
	// <p>MODEL 域专属</p>
	Model *ModelUsageSummary `json:"Model,omitnil,omitempty" name:"Model"`

	// <p>PLATFORM 域专属</p>
	Platform *PlatformUsageSummary `json:"Platform,omitnil,omitempty" name:"Platform"`

	// <p>PLUGIN 域专属</p>
	Plugin *PluginUsageSummary `json:"Plugin,omitnil,omitempty" name:"Plugin"`

	// <p>来源 ID；CORP 视图=space_id（企业视图按 space 分组），SPACE 视图=app_id（uint64 字符串），APP 视图=app_id</p>
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// <p>来源名称；CORP 视图=space_name，SPACE 视图=app_name，APP 视图=app_name</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>视图类型，决定 SourceId/SourceName 的业务含义</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>VIEW_TYPE_UNSPECIFIED</td><td>0</td><td>未指定（无效值，请求勿传）</td></tr><tr><td>VIEW_TYPE_CORP</td><td>1</td><td>企业视图</td></tr><tr><td>VIEW_TYPE_SPACE</td><td>2</td><td>空间视图</td></tr><tr><td>VIEW_TYPE_APP</td><td>3</td><td>应用视图</td></tr></tbody></table>
	ViewType *int64 `json:"ViewType,omitnil,omitempty" name:"ViewType"`
}

type Variable struct {
	// <p>默认文件名称</p>
	DefaultFileName *string `json:"DefaultFileName,omitnil,omitempty" name:"DefaultFileName"`

	// <p>默认值</p>
	DefaultValue *string `json:"DefaultValue,omitnil,omitempty" name:"DefaultValue"`

	// <p>变量描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>模块类型。枚举值: 1:环境参数, 2:应用参数, 3:系统参数, -1:所有参数</p>
	ModuleType *int64 `json:"ModuleType,omitnil,omitempty" name:"ModuleType"`

	// <p>变量名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>变量类型</p><p>枚举值：</p><ul><li>0： 字符串</li><li>1： 整数</li><li>2： 浮点数</li><li>3： 布尔值</li><li>4： 对象</li><li>5： 字符串数组</li><li>6： 整数数组</li><li>7： 浮点数数组</li><li>8： 布尔值数组</li><li>9： 对象数组</li><li>10： 文件</li><li>11： 文档</li><li>12： 图片</li><li>13： 音频</li><li>14： 视频</li><li>15： 文件数组</li><li>16： 文档数组</li><li>17： 图片数组</li><li>18： 音频数组</li><li>19： 视频数组</li><li>20： 数组的数组</li><li>21： 密钥</li></ul>
	Type *int64 `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>变量ID</p>
	VariableId *string `json:"VariableId,omitnil,omitempty" name:"VariableId"`

	// <p>是否启用网络策略(仅环境变量生效)</p>
	EnableEndpoints *bool `json:"EnableEndpoints,omitnil,omitempty" name:"EnableEndpoints"`

	// <p>网络策略列表(支持: 精确域名、*.通配子域名、可带协议/端口/路径前缀)</p>
	EndpointList []*string `json:"EndpointList,omitnil,omitempty" name:"EndpointList"`
}

type ViewScope struct {
	// <p>视图类型；枚举值：VIEW_TYPE_CORP(1) 企业视图、VIEW_TYPE_SPACE(2) 空间视图、VIEW_TYPE_APP(3) 应用视图</p><table><tbody><tr><td>枚举项</td><td>枚举值</td><td>描述</td></tr><tr><td>VIEW_TYPE_UNSPECIFIED</td><td>0</td><td>未指定（无效值，请求勿传）</td></tr><tr><td>VIEW_TYPE_CORP</td><td>1</td><td>企业视图</td></tr><tr><td>VIEW_TYPE_SPACE</td><td>2</td><td>空间视图</td></tr><tr><td>VIEW_TYPE_APP</td><td>3</td><td>应用视图</td></tr></tbody></table>
	ViewType *int64 `json:"ViewType,omitnil,omitempty" name:"ViewType"`

	// <p>视图范围 ID；VIEW_TYPE_CORP 留空；VIEW_TYPE_SPACE 填 space_id；VIEW_TYPE_APP 填 app_id（uint64 雪花 ID 的十进制字符串）</p>
	ScopeId *string `json:"ScopeId,omitnil,omitempty" name:"ScopeId"`
}

type VoiceConfig struct {
	// 数智人音色key,需要和公有云音色id对齐
	TimbreKey *string `json:"TimbreKey,omitnil,omitempty" name:"TimbreKey"`

	// 音色名称
	VoiceName *string `json:"VoiceName,omitnil,omitempty" name:"VoiceName"`

	// 公有云音色id
	VoiceType *uint64 `json:"VoiceType,omitnil,omitempty" name:"VoiceType"`
}

type WeeklySchedule struct {
	// 定时配置（星期）
	Times []*WeeklyTime `json:"Times,omitnil,omitempty" name:"Times"`
}

type WeeklyTime struct {
	// 时间
	TimeOfDay *string `json:"TimeOfDay,omitnil,omitempty" name:"TimeOfDay"`

	// 周几
	Weekday *int64 `json:"Weekday,omitnil,omitempty" name:"Weekday"`
}