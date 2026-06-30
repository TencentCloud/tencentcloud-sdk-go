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

package v20230418

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AIGWACLSubject struct {
	// <p>鉴权主体ID</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>鉴权主体名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type AIGWConsumerGroupBrief struct {
	// <p>消费者组名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>消费者组Id</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`
}

type AIGWCrossServiceFallbackConfig struct {
	// <p>触发条件</p><p>枚举值：</p><ul><li>ServiceUnavailable： 服务不可用</li><li>ConnectionTimeout： 连接超时</li><li>RateLimited： 限流</li></ul>
	TriggerConditions []*string `json:"TriggerConditions,omitnil,omitempty" name:"TriggerConditions"`

	// <p>fallback 服务链</p>
	FallbackServiceChain []*AIGWFallbackServiceItem `json:"FallbackServiceChain,omitnil,omitempty" name:"FallbackServiceChain"`

	// <p>额度降级触发配置</p>
	QuotaFallbackTrigger *AIGWLLMQuotaFallbackTrigger `json:"QuotaFallbackTrigger,omitnil,omitempty" name:"QuotaFallbackTrigger"`
}

type AIGWCustomDesensitizeRule struct {
	// <p>自定义脱敏规则名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>自定义脱敏规则匹配正则</p>
	Pattern *string `json:"Pattern,omitnil,omitempty" name:"Pattern"`

	// <p>自定义脱敏规则掩码</p>
	MaskFormat *string `json:"MaskFormat,omitnil,omitempty" name:"MaskFormat"`

	// <p>自定义脱敏规则开关</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type AIGWFallbackServiceItem struct {
	// <p>模型服务 Id</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// <p>模型服务名</p>
	ModelServiceName *string `json:"ModelServiceName,omitnil,omitempty" name:"ModelServiceName"`
}

type AIGWForwardDesensitizeConfig struct {
	// <p>转发脱敏开关</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>预定义规则类型</p><p>枚举值：</p><ul><li>Phone： 电话号码</li><li>IdCard： 身份证号</li><li>BankCard： 银行卡号</li><li>Email： 电子邮箱地址</li><li>IP： IP地址</li><li>Name： 姓名</li></ul>
	PredefinedRuleTypes []*string `json:"PredefinedRuleTypes,omitnil,omitempty" name:"PredefinedRuleTypes"`

	// <p>自定义脱敏规则</p>
	CustomRules []*AIGWCustomDesensitizeRule `json:"CustomRules,omitnil,omitempty" name:"CustomRules"`

	// <p>掩码</p>
	PlaceholderFormat *string `json:"PlaceholderFormat,omitnil,omitempty" name:"PlaceholderFormat"`

	// <p>脱敏异常处理</p><p>枚举值：</p><ul><li>Reject： 拒绝请求</li><li>Skip： 跳过</li></ul>
	OnFailure *string `json:"OnFailure,omitnil,omitempty" name:"OnFailure"`
}

type AIGWHealthCheckSetting struct {
	// <p>健康检查类型</p><p>枚举值：</p><ul><li>MCP： 标准mcp</li><li>HTTP： http</li></ul>
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// <p>检查间隔</p>
	HealthCheckIntervalSecond *uint64 `json:"HealthCheckIntervalSecond,omitnil,omitempty" name:"HealthCheckIntervalSecond"`

	// <p>检查超时时间</p>
	HealthCheckTimeout *uint64 `json:"HealthCheckTimeout,omitnil,omitempty" name:"HealthCheckTimeout"`

	// <p>检查失败阈值</p>
	HealthCheckFailThreshold *uint64 `json:"HealthCheckFailThreshold,omitnil,omitempty" name:"HealthCheckFailThreshold"`

	// <p>检查恢复阈值</p>
	HealthCheckRecoverThreshold *uint64 `json:"HealthCheckRecoverThreshold,omitnil,omitempty" name:"HealthCheckRecoverThreshold"`

	// <p>检查路径</p>
	HealthCheckPath *string `json:"HealthCheckPath,omitnil,omitempty" name:"HealthCheckPath"`

	// <p>检查方法</p>
	HealthCheckMethod *string `json:"HealthCheckMethod,omitnil,omitempty" name:"HealthCheckMethod"`
}

type AIGWIntentRoute struct {
	// <p>意图识别模型id</p>
	IntentModelServiceId *string `json:"IntentModelServiceId,omitnil,omitempty" name:"IntentModelServiceId"`

	// <p>置信度</p>
	ConfidenceThreshold *float64 `json:"ConfidenceThreshold,omitnil,omitempty" name:"ConfidenceThreshold"`

	// <p>默认服务id</p>
	DefaultModelServiceId *string `json:"DefaultModelServiceId,omitnil,omitempty" name:"DefaultModelServiceId"`

	// <p>规则</p>
	Rules []*AIGWIntentRouteRule `json:"Rules,omitnil,omitempty" name:"Rules"`
}

type AIGWIntentRouteRule struct {
	// <p>意图编码</p><p>枚举值：</p><ul><li>Coder： 代码编写</li><li>Math： 数学计算</li><li>Translation： 翻译</li><li>Flash： 快速问答</li><li>Complex： 复杂推理</li></ul>
	IntentCode *string `json:"IntentCode,omitnil,omitempty" name:"IntentCode"`

	// <p>模型服务id</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`
}

type AIGWKVMatch struct {
	// <p>键</p>
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// <p>值</p>
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// <p>操作类型</p>
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`
}

type AIGWLLMQuotaFallbackTrigger struct {
	// <p>配额感知阈值百分比（RPM 与 TPM 共用）</p><p>取值范围：[0, 99]</p>
	ThresholdPercent *int64 `json:"ThresholdPercent,omitnil,omitempty" name:"ThresholdPercent"`

	// <p>检查维度策略</p><p>枚举值：</p><ul><li>AnyInsufficient：  RPM 或 TPM 任一不足即触发</li><li>AllInsufficient： RPM 和 TPM 同时不足才触发</li></ul>
	CheckDimension *string `json:"CheckDimension,omitnil,omitempty" name:"CheckDimension"`
}

type AIGWLLMQuotaLimit struct {
	// <p>该模型服务每分钟请求数上限，0 表示该维度不限</p>
	RPMLimit *int64 `json:"RPMLimit,omitnil,omitempty" name:"RPMLimit"`

	// <p>该模型服务每分钟 Token 数上限，0 表示该维度不限</p>
	TPMLimit *int64 `json:"TPMLimit,omitnil,omitempty" name:"TPMLimit"`
}

type AIGWLLMTokenUsageItem struct {
	// <p>消费者Id</p>
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// <p>消费者名称</p>
	ConsumerName *string `json:"ConsumerName,omitnil,omitempty" name:"ConsumerName"`

	// <p>消费者组信息列表</p>
	ConsumerGroups []*AIGWConsumerGroupBrief `json:"ConsumerGroups,omitnil,omitempty" name:"ConsumerGroups"`

	// <p>模型服务Id</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// <p>模型服务名称</p>
	ModelServiceName *string `json:"ModelServiceName,omitnil,omitempty" name:"ModelServiceName"`

	// <p>输入Token数（包含缓存命中Token数）</p>
	InputTokens *int64 `json:"InputTokens,omitnil,omitempty" name:"InputTokens"`

	// <p>命中缓存输入Token数</p>
	CacheReadInputTokens *int64 `json:"CacheReadInputTokens,omitnil,omitempty" name:"CacheReadInputTokens"`

	// <p>输出Token数</p>
	OutputTokens *int64 `json:"OutputTokens,omitnil,omitempty" name:"OutputTokens"`

	// <p>消耗总Token数</p>
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`

	// <p>请求总数</p>
	RequestCount *int64 `json:"RequestCount,omitnil,omitempty" name:"RequestCount"`

	// <p>花费成本</p>
	Cost *string `json:"Cost,omitnil,omitempty" name:"Cost"`

	// <p>成本货币单位</p><p>枚举值：</p><ul><li>CNY： 人民币</li></ul><p>当前仅支持CNY</p>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`
}

type AIGWLLMTokenUsageListResult struct {
	// <p>Token用量明细返回结果列表</p>
	DataList []*AIGWLLMTokenUsageItem `json:"DataList,omitnil,omitempty" name:"DataList"`

	// <p>结果总数</p>
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type AIGWLLMTokenUsageStatisticsResult struct {
	// <p>查询时间范围内请求总数</p>
	TotalRequestCount *int64 `json:"TotalRequestCount,omitnil,omitempty" name:"TotalRequestCount"`

	// <p>查询时间范围内总输入Token数（包含命中缓存的Token数）</p>
	TotalInputTokens *int64 `json:"TotalInputTokens,omitnil,omitempty" name:"TotalInputTokens"`

	// <p>查询时间范围内总输出Token数</p>
	TotalOutputTokens *int64 `json:"TotalOutputTokens,omitnil,omitempty" name:"TotalOutputTokens"`

	// <p>查询时间范围内总命中缓存输入Token数</p>
	TotalCachedReadInputTokens *int64 `json:"TotalCachedReadInputTokens,omitnil,omitempty" name:"TotalCachedReadInputTokens"`

	// <p>查询时间范围内总成本</p>
	TotalCost *string `json:"TotalCost,omitnil,omitempty" name:"TotalCost"`

	// <p>成本货币单位</p>
	Currency *string `json:"Currency,omitnil,omitempty" name:"Currency"`

	// <p>查询时间范围内成本最高的消费者</p>
	TopConsumers []*AIGWTopConsumersItem `json:"TopConsumers,omitnil,omitempty" name:"TopConsumers"`
}

type AIGWLatencyPriorityConfig struct {
	// <p>路由规则列表</p>
	Rules []*AIGWLatencyPriorityRouteRule `json:"Rules,omitnil,omitempty" name:"Rules"`

	// <p>延迟指标</p><p>枚举值：</p><ul><li>LLMLatency： LLM 延迟</li><li>NetworkLatency： 网络延迟</li></ul>
	LatencyMetric *string `json:"LatencyMetric,omitnil,omitempty" name:"LatencyMetric"`

	// <p>路由策略</p><p>枚举值：</p><ul><li>FastMode： 快速模式</li><li>BalanceMode： 均衡模式</li></ul>
	RouteMode *string `json:"RouteMode,omitnil,omitempty" name:"RouteMode"`
}

type AIGWLatencyPriorityRouteRule struct {
	// <p>模型服务id</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`
}

type AIGWLogConfig struct {
	// <p>是否开启请求 payload 记录日志</p>
	EnableRequestLogPayloads *bool `json:"EnableRequestLogPayloads,omitnil,omitempty" name:"EnableRequestLogPayloads"`

	// <p>是否开启响应 payload 记录日志</p>
	EnableResponseLogPayloads *bool `json:"EnableResponseLogPayloads,omitnil,omitempty" name:"EnableResponseLogPayloads"`

	// <p>日志记录的请求body的最大字节数</p><p>取值范围：[512, 1048576]</p><p>EnableRequestLogPayloads 为true时必填</p>
	RequestLogPayloadMaxSize *int64 `json:"RequestLogPayloadMaxSize,omitnil,omitempty" name:"RequestLogPayloadMaxSize"`

	// <p>日志记录的响应body的最大字节数</p><p>取值范围：[512, 1048576]</p><p>EnableResponseLogPayloads 为true时必填</p>
	ResponseLogPayloadMaxSize *int64 `json:"ResponseLogPayloadMaxSize,omitnil,omitempty" name:"ResponseLogPayloadMaxSize"`
}

type AIGWLogDesensitizeConfig struct {
	// <p>日志脱敏开关</p>
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// <p>预定义规则类型</p><p>枚举值：</p><ul><li>Phone： 电话号码</li><li>IdCard： 身份证号</li><li>BankCard： 银行卡号</li><li>Email： 邮箱地址</li><li>IP： IP地址</li><li>Name： 姓名</li></ul>
	PredefinedRuleTypes []*string `json:"PredefinedRuleTypes,omitnil,omitempty" name:"PredefinedRuleTypes"`

	// <p>自定义脱敏规则</p>
	CustomRules []*AIGWCustomDesensitizeRule `json:"CustomRules,omitnil,omitempty" name:"CustomRules"`

	// <p>日志脱敏范围</p><p>枚举值：</p><ul><li>Request： 请求</li><li>Response： 响应</li></ul>
	Scope []*string `json:"Scope,omitnil,omitempty" name:"Scope"`
}

type AIGWMCPServer struct {
	// <p>MCP Server ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>MCP Server名称</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>MCP Server类型，取值：MCP/Rest2MCP</p>
	ServerType *string `json:"ServerType,omitnil,omitempty" name:"ServerType"`

	// <p>协议类型，取值: StreamableHttp</p>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>服务类型：</p><ul><li>Registry  </li><li>HostIP</li></ul>
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>MCP提供给客户端的Endpoint</p>
	MCPEndpoint *string `json:"MCPEndpoint,omitnil,omitempty" name:"MCPEndpoint"`

	// <p>注册中心来源信息</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamInfo *AIGWMCPUpstreamInfoDetail `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// <p>会话配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionConfig *AIGWMCPSessionConfig `json:"SessionConfig,omitnil,omitempty" name:"SessionConfig"`

	// <p>超时时间，单位ms</p>
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>失败重试次数</p>
	RetryCount *uint64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>运行状态</p><p>枚举值：</p><ul><li>Online： 在线</li><li>Offline： 离线</li><li>Error： 错误</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>是否启用健康检查</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>健康检查配置</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthCheck *AIGWHealthCheckSetting `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`

	// <p>Tool分组内工具数量限制</p>
	ToolCountLimit *uint64 `json:"ToolCountLimit,omitnil,omitempty" name:"ToolCountLimit"`

	// <p>Tool分组内工具命名冲突策略</p><p>枚举值：</p><ul><li>AutoPrefix： 自动前缀</li><li>Reject： 拒绝</li></ul>
	ConflictStrategy *string `json:"ConflictStrategy,omitnil,omitempty" name:"ConflictStrategy"`

	// <p>MCP 市场发布状态</p><p>枚举值：</p><ul><li>None： 未发布</li><li>Published： 已发布</li></ul>
	MarketStatus *string `json:"MarketStatus,omitnil,omitempty" name:"MarketStatus"`
}

type AIGWMCPServerACLResult struct {
	// <p>黑白名单鉴权类型</p><p>枚举值：</p><ul><li>None： 不鉴权</li><li>Allow： 白名单</li><li>Deny： 黑名单</li></ul>
	ACLType *string `json:"ACLType,omitnil,omitempty" name:"ACLType"`

	// <p>关联的消费者ID列表</p>
	ACLConsumers []*AIGWACLSubject `json:"ACLConsumers,omitnil,omitempty" name:"ACLConsumers"`

	// <p>关联的消费者组ID列表</p>
	ACLConsumerGroups []*AIGWACLSubject `json:"ACLConsumerGroups,omitnil,omitempty" name:"ACLConsumerGroups"`

	// <p>认证类型</p><p>枚举值：</p><ul><li>None： 无认证</li><li>ApiKey： API Key认证</li></ul>
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`
}

type AIGWMCPServerAuthResult struct {
	// <p>MCP服务认证类型</p><p>枚举值：</p><ul><li>None： 无认证</li><li>ApiKey： API Key认证</li></ul>
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`
}

type AIGWMCPServerList struct {
	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// <p>mcp server 数据列表</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataList []*AIGWMCPServer `json:"DataList,omitnil,omitempty" name:"DataList"`
}

type AIGWMCPSessionConfig struct {
	// <p>会话存储类型</p><p>枚举值：</p><ul><li>memory： 内存</li><li>redis： redis</li></ul>
	SessionStorage *string `json:"SessionStorage,omitnil,omitempty" name:"SessionStorage"`

	// <p>Redis配置</p>
	RedisConfig *AIGWRedisConfig `json:"RedisConfig,omitnil,omitempty" name:"RedisConfig"`
}

type AIGWMCPToolACLItem struct {
	// <p>MCP工具ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>MCP工具名称</p>
	ToolName *string `json:"ToolName,omitnil,omitempty" name:"ToolName"`

	// <p>MCP工具鉴权类型</p><p>枚举值：</p><ul><li>None： 继承server鉴权类型</li><li>Allow： 白名单</li><li>Deny： 黑名单</li></ul>
	ACLType *string `json:"ACLType,omitnil,omitempty" name:"ACLType"`

	// <p>关联的消费者ID列表</p>
	ACLConsumers []*AIGWACLSubject `json:"ACLConsumers,omitnil,omitempty" name:"ACLConsumers"`

	// <p>关联的消费者组ID列表</p>
	ACLConsumerGroups []*AIGWACLSubject `json:"ACLConsumerGroups,omitnil,omitempty" name:"ACLConsumerGroups"`
}

type AIGWMCPToolACLListResult struct {
	// <p>黑白名单鉴权类型</p><p>枚举值：</p><ul><li>None： 不鉴权</li><li>Allow： 白名单</li><li>Deny： 黑名单</li></ul>
	ACLType *string `json:"ACLType,omitnil,omitempty" name:"ACLType"`

	// <p>数据列表</p>
	DataList []*AIGWMCPToolACLItem `json:"DataList,omitnil,omitempty" name:"DataList"`

	// <p>计数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type AIGWMCPUpstreamInfo struct {
	// <p>注册中心来源ID</p>
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// <p>命名空间</p>
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// <p>MCP服务 id</p>
	MCPServerId *string `json:"MCPServerId,omitnil,omitempty" name:"MCPServerId"`

	// <p>协议，UpstreamType是Registry 时必传</p><ul><li>http</li><li>https</li></ul>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>域名或ip</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>端口</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>服务 id</p>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// <p>服务分组</p>
	ServiceGroup *string `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// <p>mcp endpoint</p>
	MCPEndpoint *string `json:"MCPEndpoint,omitnil,omitempty" name:"MCPEndpoint"`

	// <p>message端点路径，SSE协议时配置</p>
	MessageEndpoint *string `json:"MessageEndpoint,omitnil,omitempty" name:"MessageEndpoint"`
}

type AIGWMCPUpstreamInfoDetail struct {
	// <p>注册中心来源ID</p>
	SourceId *string `json:"SourceId,omitnil,omitempty" name:"SourceId"`

	// <p>注册中心来源名称, 入参不传，用于返回</p>
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// <p>命名空间</p>
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// <p>服务 id</p>
	MCPServerId *string `json:"MCPServerId,omitnil,omitempty" name:"MCPServerId"`

	// <p>协议，UpstreamType是Registry 时必传</p><ul><li>http</li><li>https</li></ul>
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// <p>域名或ip</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>端口</p>
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>服务 id</p>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// <p>服务名字</p>
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// <p>服务分组</p>
	ServiceGroup *string `json:"ServiceGroup,omitnil,omitempty" name:"ServiceGroup"`

	// <p>mcp endpoint</p>
	MCPEndpoint *string `json:"MCPEndpoint,omitnil,omitempty" name:"MCPEndpoint"`

	// <p>SSE message路径</p>
	MessageEndpoint *string `json:"MessageEndpoint,omitnil,omitempty" name:"MessageEndpoint"`
}

type AIGWRedisConfig struct {
	// <p>Host</p>
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// <p>端口</p>
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// <p>用户名</p>
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// <p>密码</p>
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`
}

type AIGWTagFilter struct {
	// <p>匹配策略</p><p>枚举值：</p><ul><li>AND： 并</li><li>OR： 或</li></ul>
	MatchStrategy *string `json:"MatchStrategy,omitnil,omitempty" name:"MatchStrategy"`

	// <p>标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type AIGWTopConsumersItem struct {
	// <p>消费者Id</p>
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// <p>消费者名称</p>
	ConsumerName *string `json:"ConsumerName,omitnil,omitempty" name:"ConsumerName"`

	// <p>该消费者花费的Token数</p>
	TotalTokens *int64 `json:"TotalTokens,omitnil,omitempty" name:"TotalTokens"`
}

// Predefined struct for user
type AddCloudNativeAPIGatewayConsumerGroupAuthRequestParams struct {
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>授权资源类型。</p><p>枚举值：</p><ul><li>ModelAPI：模型 API</li><li>MCPServer：MCP Server</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>对应资源的 ID。</p><ul><li>ResourceType=ModelAPI 时是模型 API ID</li><li>ResourceType=MCPServer 时是 MCP Server ID</li></ul>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>消费者组 ID 列表（每个 ID 以 cg- 开头），长度 1-10。</p>
	ConsumerGroupIds []*string `json:"ConsumerGroupIds,omitnil,omitempty" name:"ConsumerGroupIds"`
}

type AddCloudNativeAPIGatewayConsumerGroupAuthRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>授权资源类型。</p><p>枚举值：</p><ul><li>ModelAPI：模型 API</li><li>MCPServer：MCP Server</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>对应资源的 ID。</p><ul><li>ResourceType=ModelAPI 时是模型 API ID</li><li>ResourceType=MCPServer 时是 MCP Server ID</li></ul>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>消费者组 ID 列表（每个 ID 以 cg- 开头），长度 1-10。</p>
	ConsumerGroupIds []*string `json:"ConsumerGroupIds,omitnil,omitempty" name:"ConsumerGroupIds"`
}

func (r *AddCloudNativeAPIGatewayConsumerGroupAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCloudNativeAPIGatewayConsumerGroupAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "ConsumerGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCloudNativeAPIGatewayConsumerGroupAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCloudNativeAPIGatewayConsumerGroupAuthResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCloudNativeAPIGatewayConsumerGroupAuthResponse struct {
	*tchttp.BaseResponse
	Response *AddCloudNativeAPIGatewayConsumerGroupAuthResponseParams `json:"Response"`
}

func (r *AddCloudNativeAPIGatewayConsumerGroupAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCloudNativeAPIGatewayConsumerGroupAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCloudNativeAPIGatewayConsumerInGroupRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组 ID（以 cg- 开头）。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>消费者 ID 列表，长度 1-10。</p>
	ConsumerIds []*string `json:"ConsumerIds,omitnil,omitempty" name:"ConsumerIds"`
}

type AddCloudNativeAPIGatewayConsumerInGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组 ID（以 cg- 开头）。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>消费者 ID 列表，长度 1-10。</p>
	ConsumerIds []*string `json:"ConsumerIds,omitnil,omitempty" name:"ConsumerIds"`
}

func (r *AddCloudNativeAPIGatewayConsumerInGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCloudNativeAPIGatewayConsumerInGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerGroupId")
	delete(f, "ConsumerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddCloudNativeAPIGatewayConsumerInGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddCloudNativeAPIGatewayConsumerInGroupResponseParams struct {
	// <p>是否成功。</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AddCloudNativeAPIGatewayConsumerInGroupResponse struct {
	*tchttp.BaseResponse
	Response *AddCloudNativeAPIGatewayConsumerInGroupResponseParams `json:"Response"`
}

func (r *AddCloudNativeAPIGatewayConsumerInGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddCloudNativeAPIGatewayConsumerInGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudNativeAPIGatewaySecretKeyRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源ID，当前最多支持一个
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

type BindCloudNativeAPIGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源ID，当前最多支持一个
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

func (r *BindCloudNativeAPIGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudNativeAPIGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ResourceType")
	delete(f, "ResourceIds")
	delete(f, "SecretKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindCloudNativeAPIGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindCloudNativeAPIGatewaySecretKeyResponseParams struct {
	// 结果
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindCloudNativeAPIGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *BindCloudNativeAPIGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *BindCloudNativeAPIGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindCloudNativeAPIGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CNAPIGwConsumer struct {
	// <p>消费者 ID。</p>
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// <p>名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间 yyyy-MM-dd hh:mm:ss</p>
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>消费者分组</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerGroups []*CNAPIGwConsumerGroup `json:"ConsumerGroups,omitnil,omitempty" name:"ConsumerGroups"`
}

type CNAPIGwConsumerGroup struct {
	// 分组id
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// 名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 状态Disable/Enable
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间 yyyy-MM-dd hh:mm:ss
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 绑定的消费者数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindCount *uint64 `json:"BindCount,omitnil,omitempty" name:"BindCount"`
}

type CNAPIGwCreateCommonResult struct {
	// 是否成功
	Success *bool `json:"Success,omitnil,omitempty" name:"Success"`

	// 对应的id 值
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`
}

type CNAPIGwMCPTool struct {
	// <p>名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>方法</p><p>枚举值：</p><ul><li>GET： GET</li><li>POST： POST</li><li>PUT： PUT</li><li>DELETE： DELETE</li><li>PATCH： PATCH</li></ul>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>工具 id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>报文格式</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>服务 id</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// <p>路径</p>
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>参数</p>
	InputParams []*CNAPIGwMCPToolParam `json:"InputParams,omitnil,omitempty" name:"InputParams"`

	// <p>创建时间</p><p>参数格式：yyyy-MM-dd hh:mm:ss</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>更新时间</p><p>参数格式：yyyy-MM-dd hh:mm:ss</p>
	UpdateTime *string `json:"UpdateTime,omitnil,omitempty" name:"UpdateTime"`

	// <p>tool状态</p><p>枚举值：</p><ul><li>Enable： 启用</li><li>Disable： 禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>当前版本号</p>
	CurrentVersion *string `json:"CurrentVersion,omitnil,omitempty" name:"CurrentVersion"`
}

type CNAPIGwMCPToolList struct {
	// <p>MCPTool 列表</p>
	DataList []*CNAPIGwMCPTool `json:"DataList,omitnil,omitempty" name:"DataList"`

	// <p>总数</p>
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type CNAPIGwMCPToolParam struct {
	// <p>名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>参数类型</p><p>枚举值：</p><ul><li>string： 字符串</li><li>number： 数字</li><li>boolean： 布尔</li><li>array： 数组</li><li>object： 对象</li></ul>
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// <p>必填</p>
	Required *bool `json:"Required,omitnil,omitempty" name:"Required"`

	// <p>位置</p><p>枚举值：</p><ul><li>query： query</li><li>path： query</li><li>header： header</li><li>body： body</li></ul>
	Position *string `json:"Position,omitnil,omitempty" name:"Position"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>默认值</p>
	Default *string `json:"Default,omitnil,omitempty" name:"Default"`

	// <p>数组子项</p>
	Items *CNAPIGwMCPToolParam `json:"Items,omitnil,omitempty" name:"Items"`

	// <p>对象属性</p>
	Properties []*CNAPIGwMCPToolParam `json:"Properties,omitnil,omitempty" name:"Properties"`

	// <p>转发到后端的名称</p><p>不填则使用原始名称</p>
	BackendName *string `json:"BackendName,omitnil,omitempty" name:"BackendName"`
}

type CNAPIGwSecretKey struct {
	// <p>密钥id</p>
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`

	// <p>密钥名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>密钥协议类型。</p>
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// <p>状态。</p><p>枚举值：</p><ul><li>Enable： 启用</li><li>Disable： 禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>密钥生成方式。</p><p>枚举值：</p><ul><li>System： 系统自动生成</li><li>Custom： 用户自定义</li><li>KMS： 使用 KMS 密钥</li></ul>
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// <p>密钥明文</p>
	SecretValue *string `json:"SecretValue,omitnil,omitempty" name:"SecretValue"`

	// <p>KMS凭证名字</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	KmsKeyName *string `json:"KmsKeyName,omitnil,omitempty" name:"KmsKeyName"`

	// <p>KMS凭证版本</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	KmsKeyVersion *string `json:"KmsKeyVersion,omitnil,omitempty" name:"KmsKeyVersion"`

	// <p>描述</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否可以绑定</p>
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanBind *bool `json:"CanBind,omitnil,omitempty" name:"CanBind"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>修改时间</p>
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// <p>绑定数</p>
	BindCount *uint64 `json:"BindCount,omitnil,omitempty" name:"BindCount"`

	// <p>密钥归属资源类型。</p><p>枚举值：</p><ul><li>Consumer： 消费者</li><li>ModelService： 模型服务</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type CloudNativeAPIGatewayLLMModelAPI struct {
	// <p>模型 API ID。</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>创建时间</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>修改时间</p>
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// <p>AI 网关 LLM 模型 API 的唯一标识名称，格式规则：2-50 字符，支持英文、数字、下划线。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>选择业务场景,xa0 选项：Chat（聊天）。</p>
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// <p>业务场景对应的请求协议，选项：OpenAI（目前只支持 OpenAI）。</p>
	RequestProtocol *string `json:"RequestProtocol,omitnil,omitempty" name:"RequestProtocol"`

	// <p>路由列表</p>
	RouteList []*DefaultKongRoute `json:"RouteList,omitnil,omitempty" name:"RouteList"`

	// <p>为API设置统一的前缀，格式：以/开头，支持字母、数字、短横线。</p>
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// <p>路径简化，<br>启用时：客户端请求路径 → 移除Base Path → 后端接收简洁路径<br>禁用时：客户端请求路径 → 完整传递给后端。</p>
	StripPath *bool `json:"StripPath,omitnil,omitempty" name:"StripPath"`

	// <p>模型 API 的相关描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>模型服务Id</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// <p>模型服务名称</p>
	ModelServiceName *string `json:"ModelServiceName,omitnil,omitempty" name:"ModelServiceName"`

	// <p>模型服务路由策略（是指如何路由到模型服务）</p>
	ModelServiceRoute *CloudNativeAPIGatewayLLMModelServiceRoute `json:"ModelServiceRoute,omitnil,omitempty" name:"ModelServiceRoute"`

	// <p>无</p>
	MatchHeaders []*AIGWKVMatch `json:"MatchHeaders,omitnil,omitempty" name:"MatchHeaders"`

	// <p>是否开启跨服务fallback</p>
	EnableCrossServiceFallback *bool `json:"EnableCrossServiceFallback,omitnil,omitempty" name:"EnableCrossServiceFallback"`

	// <p>跨服务fallback配置详情</p>
	CrossServiceFallbackConfig *AIGWCrossServiceFallbackConfig `json:"CrossServiceFallbackConfig,omitnil,omitempty" name:"CrossServiceFallbackConfig"`

	// <p>是否展示模型API</p>
	DescribeCloudNativeAPIGatewayLLMModelAPI *bool `json:"DescribeCloudNativeAPIGatewayLLMModelAPI,omitnil,omitempty" name:"DescribeCloudNativeAPIGatewayLLMModelAPI"`

	// <p>标签</p>
	TagFilter *AIGWTagFilter `json:"TagFilter,omitnil,omitempty" name:"TagFilter"`

	// <p>日志显示相关开关</p>
	LogConfig *AIGWLogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`

	// <p>日志脱敏规则</p>
	LogDesensitizeConfig *AIGWLogDesensitizeConfig `json:"LogDesensitizeConfig,omitnil,omitempty" name:"LogDesensitizeConfig"`

	// <p>转发脱敏规则</p>
	ForwardDesensitizeConfig *AIGWForwardDesensitizeConfig `json:"ForwardDesensitizeConfig,omitnil,omitempty" name:"ForwardDesensitizeConfig"`
}

type CloudNativeAPIGatewayLLMModelFallbackRule struct {
	// 备选模型，主模型不可用时将依次按顺序尝试。
	FallbackModels []*string `json:"FallbackModels,omitnil,omitempty" name:"FallbackModels"`
}

type CloudNativeAPIGatewayLLMModelParamCheckInfo struct {
	// 允许的模型列表。
	AllowModelList []*string `json:"AllowModelList,omitnil,omitempty" name:"AllowModelList"`

	// 模型参数校验失败时的处理策略，选项：Return404（返回404）、FallBackToDefaultModel（使用默认模型降级）。
	ModelValidationFailureStrategy *string `json:"ModelValidationFailureStrategy,omitnil,omitempty" name:"ModelValidationFailureStrategy"`
}

type CloudNativeAPIGatewayLLMModelService struct {
	// <p>模型服务 ID。</p>
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// <p>模型服务名称。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>创建时间。</p>
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// <p>修改时间。</p>
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// <p>服务类型，目前只支持xa0LLMService。</p>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>选择模型提供商, 选项：OpenAI、Anthropic、Azure OpenAI、自定义HTTP。</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>API协议标准，根据供应商动态变化：OpenAI→OpenAI/v1，Anthropic→Anthropic/v1等</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>自定义的模型请求 URL。</p>
	UpstreamURL *string `json:"UpstreamURL,omitnil,omitempty" name:"UpstreamURL"`

	// <p>模型选择方式，选项：Specify（指定模型）、PassThrough（透传请求模型）。</p>
	ModelSelector *string `json:"ModelSelector,omitnil,omitempty" name:"ModelSelector"`

	// <p>默认模型，模型选择方式为 Specify 时必填。</p>
	DefaultModel *string `json:"DefaultModel,omitnil,omitempty" name:"DefaultModel"`

	// <p>开启模型降级，模型选择方式为 Specify 时必填。</p>
	EnableModelFallback *bool `json:"EnableModelFallback,omitnil,omitempty" name:"EnableModelFallback"`

	// <p>可以配置备用模型规则，EnableSpecifyModelFallbackxa0为 true 时必填。</p>
	ModelFallbackRule *CloudNativeAPIGatewayLLMModelFallbackRule `json:"ModelFallbackRule,omitnil,omitempty" name:"ModelFallbackRule"`

	// <p>开启模型参数校验，是否校验客户端传递的 model 参数,xa0模型选择方式为 PassThrough 时必填。</p>
	EnableModelParamCheck *bool `json:"EnableModelParamCheck,omitnil,omitempty" name:"EnableModelParamCheck"`

	// <p>模型检验信息，EnableModelParamCheckxa0为 true 时必填。</p>
	ModelParamCheckRule *CloudNativeAPIGatewayLLMModelParamCheckInfo `json:"ModelParamCheckRule,omitnil,omitempty" name:"ModelParamCheckRule"`

	// <p>描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>连接超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：10000</p>
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>写入超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	WriteTimeout *int64 `json:"WriteTimeout,omitnil,omitempty" name:"WriteTimeout"`

	// <p>读取超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p>
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// <p>重试次数</p><p>取值范围：[0, 5]</p><p>单位：次</p><p>默认值：0</p>
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// <p>路径拼接模式</p><p>枚举值：</p><ul><li>FixedPath： 固定路径</li><li>AutoConcat： 自动拼接</li></ul>
	UpstreamUrlMode *string `json:"UpstreamUrlMode,omitnil,omitempty" name:"UpstreamUrlMode"`

	// <p>sni</p>
	SNI *string `json:"SNI,omitnil,omitempty" name:"SNI"`

	// <p>配额限制</p>
	QuotaLimit *AIGWLLMQuotaLimit `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`

	// <p>标签</p>
	Tags *string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// <p>绑定的模型服务秘钥</p>
	SecretKeyIds []*string `json:"SecretKeyIds,omitnil,omitempty" name:"SecretKeyIds"`
}

type CloudNativeAPIGatewayLLMModelServiceRoute struct {
	// <p>生效的路由算法类型：权重路由，模型名称路由、参数路由等Weighted/ModelName/Query (预留多个，暂时只能填写一个)</p>
	SelectedTypes []*string `json:"SelectedTypes,omitnil,omitempty" name:"SelectedTypes"`

	// <p>权重路由配置，最多10个</p>
	WeightedConfig []*CloudNativeAPIGatewayLLMModelServiceRouteWeightedStrategy `json:"WeightedConfig,omitnil,omitempty" name:"WeightedConfig"`

	// <p>模型名称路由配置，最多10个</p>
	ModelNameConfig []*CloudNativeAPIGatewayLLMModelServiceRouteModelNameStrategy `json:"ModelNameConfig,omitnil,omitempty" name:"ModelNameConfig"`

	// <p>意图识别</p>
	IntentRouteConfig *AIGWIntentRoute `json:"IntentRouteConfig,omitnil,omitempty" name:"IntentRouteConfig"`

	// <p>延迟路由</p>
	LatencyPriorityConfig *AIGWLatencyPriorityConfig `json:"LatencyPriorityConfig,omitnil,omitempty" name:"LatencyPriorityConfig"`
}

type CloudNativeAPIGatewayLLMModelServiceRouteModelNameStrategy struct {
	// <p>模型服务id</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// <p>匹配模型服务</p>
	MatchModelName *string `json:"MatchModelName,omitnil,omitempty" name:"MatchModelName"`

	// <p>重写模型</p>
	RewriteModelName *string `json:"RewriteModelName,omitnil,omitempty" name:"RewriteModelName"`
}

type CloudNativeAPIGatewayLLMModelServiceRouteWeightedStrategy struct {
	// <p>模型服务id</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// <p>权重值</p>
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayConsumerGroupRequestParams struct {
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组名称，最长 60 字符。同一网关下唯一。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>启用状态。</p><p>枚举值：</p><ul><li>Enable：启用</li><li>Disable：禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>消费者组描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateCloudNativeAPIGatewayConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组名称，最长 60 字符。同一网关下唯一。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>启用状态。</p><p>枚举值：</p><ul><li>Enable：启用</li><li>Disable：禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>消费者组描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateCloudNativeAPIGatewayConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayConsumerGroupResponseParams struct {
	// <p>创建结果。包含成功标识与新建资源 ID。</p>
	Result *CNAPIGwCreateCommonResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayConsumerGroupResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayConsumerRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者名称，最长 60 字符。同一网关下唯一。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>消费者描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateCloudNativeAPIGatewayConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者名称，最长 60 字符。同一网关下唯一。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>消费者描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateCloudNativeAPIGatewayConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayConsumerResponseParams struct {
	// <p>创建结果。包含成功标识与新建资源 ID。</p>
	Result *CNAPIGwCreateCommonResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayConsumerResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayConsumerResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayLLMModelAPIRequestParams struct {
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型 API 名称，最长 60 字符。同一网关下唯一。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>业务场景。</p><p>枚举值：</p><ul><li>Chat：聊天</li><li>Image：图像（需要网关版本 ≥ 3.9.3）</li></ul>
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// <p>请求协议（小写）。当前仅支持：</p><ul><li>openai</li></ul>
	RequestProtocol *string `json:"RequestProtocol,omitnil,omitempty" name:"RequestProtocol"`

	// <p>关联的模型服务 ID 列表，长度 1-10。</p><p>注：字段名建议改为 ModelServiceIds，当前保留用于兼容。</p>
	ListModelServiceId []*string `json:"ListModelServiceId,omitnil,omitempty" name:"ListModelServiceId"`

	// <p>路由列表，至少 1 条。每条包含 Methods/Paths/Hosts 等 Kong 路由属性。</p>
	RouteList []*DefaultKongRoute `json:"RouteList,omitnil,omitempty" name:"RouteList"`

	// <p>统一前缀路径（可选）。例如 /v1/openai。</p>
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// <p>模型 API 描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>多模型服务路由策略。ListModelServiceId 多于 1 项时必填。</p>
	ModelServiceRoute *CloudNativeAPIGatewayLLMModelServiceRoute `json:"ModelServiceRoute,omitnil,omitempty" name:"ModelServiceRoute"`

	// <p>Header 路由匹配规则。当前仅支持 Operator=exact。</p>
	MatchHeaders []*AIGWKVMatch `json:"MatchHeaders,omitnil,omitempty" name:"MatchHeaders"`

	// <p>是否启用跨服务 Fallback。开启后需提供 CrossServiceFallbackConfig。</p>
	EnableCrossServiceFallback *bool `json:"EnableCrossServiceFallback,omitnil,omitempty" name:"EnableCrossServiceFallback"`

	// <p>跨服务 Fallback 配置。EnableCrossServiceFallback=true 时必填。</p>
	CrossServiceFallbackConfig *AIGWCrossServiceFallbackConfig `json:"CrossServiceFallbackConfig,omitnil,omitempty" name:"CrossServiceFallbackConfig"`

	// <p>标签过滤策略。需要网关版本 ≥ 3.9.4。</p>
	TagFilter *AIGWTagFilter `json:"TagFilter,omitnil,omitempty" name:"TagFilter"`

	// <p>日志输出配置（请求/响应 payload 落 LLM Log）。需要网关版本 ≥ 3.9.4。</p>
	LogConfig *AIGWLogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`
}

type CreateCloudNativeAPIGatewayLLMModelAPIRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型 API 名称，最长 60 字符。同一网关下唯一。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>业务场景。</p><p>枚举值：</p><ul><li>Chat：聊天</li><li>Image：图像（需要网关版本 ≥ 3.9.3）</li></ul>
	SceneType *string `json:"SceneType,omitnil,omitempty" name:"SceneType"`

	// <p>请求协议（小写）。当前仅支持：</p><ul><li>openai</li></ul>
	RequestProtocol *string `json:"RequestProtocol,omitnil,omitempty" name:"RequestProtocol"`

	// <p>关联的模型服务 ID 列表，长度 1-10。</p><p>注：字段名建议改为 ModelServiceIds，当前保留用于兼容。</p>
	ListModelServiceId []*string `json:"ListModelServiceId,omitnil,omitempty" name:"ListModelServiceId"`

	// <p>路由列表，至少 1 条。每条包含 Methods/Paths/Hosts 等 Kong 路由属性。</p>
	RouteList []*DefaultKongRoute `json:"RouteList,omitnil,omitempty" name:"RouteList"`

	// <p>统一前缀路径（可选）。例如 /v1/openai。</p>
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// <p>模型 API 描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>多模型服务路由策略。ListModelServiceId 多于 1 项时必填。</p>
	ModelServiceRoute *CloudNativeAPIGatewayLLMModelServiceRoute `json:"ModelServiceRoute,omitnil,omitempty" name:"ModelServiceRoute"`

	// <p>Header 路由匹配规则。当前仅支持 Operator=exact。</p>
	MatchHeaders []*AIGWKVMatch `json:"MatchHeaders,omitnil,omitempty" name:"MatchHeaders"`

	// <p>是否启用跨服务 Fallback。开启后需提供 CrossServiceFallbackConfig。</p>
	EnableCrossServiceFallback *bool `json:"EnableCrossServiceFallback,omitnil,omitempty" name:"EnableCrossServiceFallback"`

	// <p>跨服务 Fallback 配置。EnableCrossServiceFallback=true 时必填。</p>
	CrossServiceFallbackConfig *AIGWCrossServiceFallbackConfig `json:"CrossServiceFallbackConfig,omitnil,omitempty" name:"CrossServiceFallbackConfig"`

	// <p>标签过滤策略。需要网关版本 ≥ 3.9.4。</p>
	TagFilter *AIGWTagFilter `json:"TagFilter,omitnil,omitempty" name:"TagFilter"`

	// <p>日志输出配置（请求/响应 payload 落 LLM Log）。需要网关版本 ≥ 3.9.4。</p>
	LogConfig *AIGWLogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`
}

func (r *CreateCloudNativeAPIGatewayLLMModelAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayLLMModelAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "SceneType")
	delete(f, "RequestProtocol")
	delete(f, "ListModelServiceId")
	delete(f, "RouteList")
	delete(f, "BasePath")
	delete(f, "Description")
	delete(f, "ModelServiceRoute")
	delete(f, "MatchHeaders")
	delete(f, "EnableCrossServiceFallback")
	delete(f, "CrossServiceFallbackConfig")
	delete(f, "TagFilter")
	delete(f, "LogConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayLLMModelAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayLLMModelAPIResponseParams struct {
	// <p>是否成功。</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>模型 API ID，全局唯一标识。</p>
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayLLMModelAPIResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayLLMModelAPIResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayLLMModelAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayLLMModelAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayLLMModelServiceRequestParams struct {
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>服务名称，最长60个字符，支持中英文大小写、数字及分隔符（“-”、“_”)，不能以数字和分隔符开头，不能以分隔符结尾。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>服务类型。目前仅支持 LLMService。</p><p>枚举值：</p><ul><li>LLMService： 大语言模型服务</li></ul>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>选择模型提供商, 选项：OpenAI、Anthropic、Azure OpenAI等。</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>API协议标准，根据供应商动态变化：OpenAI→OpenAI/v1，Anthropic→Anthropic/v1等</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>模型选择方式，选项：Specify（指定模型）、PassThrough（透传请求模型）。</p>
	ModelSelector *string `json:"ModelSelector,omitnil,omitempty" name:"ModelSelector"`

	// <p>LLM 厂商颁发的认证信息 token 。</p>
	SecretKeyIds []*string `json:"SecretKeyIds,omitnil,omitempty" name:"SecretKeyIds"`

	// <p>默认模型，模型选择方式为 Specify 时必填。</p>
	DefaultModel *string `json:"DefaultModel,omitnil,omitempty" name:"DefaultModel"`

	// <p>开启模型降级，模型选择方式为 Specify 时必填。</p>
	EnableModelFallback *bool `json:"EnableModelFallback,omitnil,omitempty" name:"EnableModelFallback"`

	// <p>可以配置备用模型规则，EnableSpecifyModelFallbackxa0为 true 时必填。</p>
	ModelFallbackRule *CloudNativeAPIGatewayLLMModelFallbackRule `json:"ModelFallbackRule,omitnil,omitempty" name:"ModelFallbackRule"`

	// <p>开启模型参数校验，是否校验客户端传递的 model 参数,xa0模型选择方式为 PassThrough 时必填</p>
	EnableModelParamCheck *bool `json:"EnableModelParamCheck,omitnil,omitempty" name:"EnableModelParamCheck"`

	// <p>模型检验信息，EnableModelParamCheckxa0为 true 时必填。</p>
	ModelParamCheckRule *CloudNativeAPIGatewayLLMModelParamCheckInfo `json:"ModelParamCheckRule,omitnil,omitempty" name:"ModelParamCheckRule"`

	// <p>描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>服务提供商自定义 url</p>
	UpstreamURL *string `json:"UpstreamURL,omitnil,omitempty" name:"UpstreamURL"`

	// <p>连接超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：10000</p>
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>写入超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	WriteTimeout *int64 `json:"WriteTimeout,omitnil,omitempty" name:"WriteTimeout"`

	// <p>读取超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// <p>重试次数</p><p>取值范围：[0, 5]</p><p>单位：次</p><p>默认值：0</p>
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// <p>路径拼接模式</p><p>枚举值：</p><ul><li>FixedPath： 固定地址</li><li>AutoConcat： 自动拼接</li></ul>
	UpstreamUrlMode *string `json:"UpstreamUrlMode,omitnil,omitempty" name:"UpstreamUrlMode"`

	// <p>sni</p>
	SNI *string `json:"SNI,omitnil,omitempty" name:"SNI"`

	// <p>模型服务级别的配额上限（RPM/TPM）。需要网关版本 ≥ 3.9.4。</p>
	QuotaLimit *AIGWLLMQuotaLimit `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`

	// <p>标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type CreateCloudNativeAPIGatewayLLMModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>服务名称，最长60个字符，支持中英文大小写、数字及分隔符（“-”、“_”)，不能以数字和分隔符开头，不能以分隔符结尾。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>服务类型。目前仅支持 LLMService。</p><p>枚举值：</p><ul><li>LLMService： 大语言模型服务</li></ul>
	ServiceType *string `json:"ServiceType,omitnil,omitempty" name:"ServiceType"`

	// <p>选择模型提供商, 选项：OpenAI、Anthropic、Azure OpenAI等。</p>
	ModelProvider *string `json:"ModelProvider,omitnil,omitempty" name:"ModelProvider"`

	// <p>API协议标准，根据供应商动态变化：OpenAI→OpenAI/v1，Anthropic→Anthropic/v1等</p>
	ModelProtocol *string `json:"ModelProtocol,omitnil,omitempty" name:"ModelProtocol"`

	// <p>模型选择方式，选项：Specify（指定模型）、PassThrough（透传请求模型）。</p>
	ModelSelector *string `json:"ModelSelector,omitnil,omitempty" name:"ModelSelector"`

	// <p>LLM 厂商颁发的认证信息 token 。</p>
	SecretKeyIds []*string `json:"SecretKeyIds,omitnil,omitempty" name:"SecretKeyIds"`

	// <p>默认模型，模型选择方式为 Specify 时必填。</p>
	DefaultModel *string `json:"DefaultModel,omitnil,omitempty" name:"DefaultModel"`

	// <p>开启模型降级，模型选择方式为 Specify 时必填。</p>
	EnableModelFallback *bool `json:"EnableModelFallback,omitnil,omitempty" name:"EnableModelFallback"`

	// <p>可以配置备用模型规则，EnableSpecifyModelFallbackxa0为 true 时必填。</p>
	ModelFallbackRule *CloudNativeAPIGatewayLLMModelFallbackRule `json:"ModelFallbackRule,omitnil,omitempty" name:"ModelFallbackRule"`

	// <p>开启模型参数校验，是否校验客户端传递的 model 参数,xa0模型选择方式为 PassThrough 时必填</p>
	EnableModelParamCheck *bool `json:"EnableModelParamCheck,omitnil,omitempty" name:"EnableModelParamCheck"`

	// <p>模型检验信息，EnableModelParamCheckxa0为 true 时必填。</p>
	ModelParamCheckRule *CloudNativeAPIGatewayLLMModelParamCheckInfo `json:"ModelParamCheckRule,omitnil,omitempty" name:"ModelParamCheckRule"`

	// <p>描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>服务提供商自定义 url</p>
	UpstreamURL *string `json:"UpstreamURL,omitnil,omitempty" name:"UpstreamURL"`

	// <p>连接超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：10000</p>
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>写入超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	WriteTimeout *int64 `json:"WriteTimeout,omitnil,omitempty" name:"WriteTimeout"`

	// <p>读取超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// <p>重试次数</p><p>取值范围：[0, 5]</p><p>单位：次</p><p>默认值：0</p>
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// <p>路径拼接模式</p><p>枚举值：</p><ul><li>FixedPath： 固定地址</li><li>AutoConcat： 自动拼接</li></ul>
	UpstreamUrlMode *string `json:"UpstreamUrlMode,omitnil,omitempty" name:"UpstreamUrlMode"`

	// <p>sni</p>
	SNI *string `json:"SNI,omitnil,omitempty" name:"SNI"`

	// <p>模型服务级别的配额上限（RPM/TPM）。需要网关版本 ≥ 3.9.4。</p>
	QuotaLimit *AIGWLLMQuotaLimit `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`

	// <p>标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *CreateCloudNativeAPIGatewayLLMModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayLLMModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "ServiceType")
	delete(f, "ModelProvider")
	delete(f, "ModelProtocol")
	delete(f, "ModelSelector")
	delete(f, "SecretKeyIds")
	delete(f, "DefaultModel")
	delete(f, "EnableModelFallback")
	delete(f, "ModelFallbackRule")
	delete(f, "EnableModelParamCheck")
	delete(f, "ModelParamCheckRule")
	delete(f, "Description")
	delete(f, "UpstreamURL")
	delete(f, "ConnectTimeout")
	delete(f, "WriteTimeout")
	delete(f, "ReadTimeout")
	delete(f, "Retries")
	delete(f, "UpstreamUrlMode")
	delete(f, "SNI")
	delete(f, "QuotaLimit")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayLLMModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayLLMModelServiceResponseParams struct {
	// <p>是否成功</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// <p>模型服务 ID，全局唯一标识。</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayLLMModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayLLMModelServiceResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayLLMModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayLLMModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayMCPServerRequestParams struct {
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>MCP服务类型</p><ul><li>MCP</li><li>Rest2MCP</li></ul>
	ServerType *string `json:"ServerType,omitnil,omitempty" name:"ServerType"`

	// <p>传输协议：StreamableHttp或SSE</p><p>枚举值：</p><ul><li>StreamableHttp： Streamable HTTP</li><li>SSE： Server-Sent Events</li></ul>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>后端类型</p><p>枚举值：</p><ul><li>MCPRegistry： mcp 注册中心- Registry</li><li>Registry： 普通注册中心</li><li>HostIP： 域名或ip</li><li>VirtualMCPServer： 虚拟MCPServer</li></ul>
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// <p>注册中心来源信息</p>
	UpstreamInfo *AIGWMCPUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// <p>会话配置</p>
	SessionConfig *AIGWMCPSessionConfig `json:"SessionConfig,omitnil,omitempty" name:"SessionConfig"`

	// <p>超时时间，单位ms，最大60000</p>
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>重试次数，最大3次</p>
	RetryCount *uint64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否启用健康检查</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>健康检查配置</p>
	HealthCheck *AIGWHealthCheckSetting `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`
}

type CreateCloudNativeAPIGatewayMCPServerRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>MCP服务类型</p><ul><li>MCP</li><li>Rest2MCP</li></ul>
	ServerType *string `json:"ServerType,omitnil,omitempty" name:"ServerType"`

	// <p>传输协议：StreamableHttp或SSE</p><p>枚举值：</p><ul><li>StreamableHttp： Streamable HTTP</li><li>SSE： Server-Sent Events</li></ul>
	Transport *string `json:"Transport,omitnil,omitempty" name:"Transport"`

	// <p>后端类型</p><p>枚举值：</p><ul><li>MCPRegistry： mcp 注册中心- Registry</li><li>Registry： 普通注册中心</li><li>HostIP： 域名或ip</li><li>VirtualMCPServer： 虚拟MCPServer</li></ul>
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// <p>注册中心来源信息</p>
	UpstreamInfo *AIGWMCPUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// <p>会话配置</p>
	SessionConfig *AIGWMCPSessionConfig `json:"SessionConfig,omitnil,omitempty" name:"SessionConfig"`

	// <p>超时时间，单位ms，最大60000</p>
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>重试次数，最大3次</p>
	RetryCount *uint64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否启用健康检查</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>健康检查配置</p>
	HealthCheck *AIGWHealthCheckSetting `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`
}

func (r *CreateCloudNativeAPIGatewayMCPServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayMCPServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "DisplayName")
	delete(f, "ServerType")
	delete(f, "Transport")
	delete(f, "UpstreamType")
	delete(f, "UpstreamInfo")
	delete(f, "SessionConfig")
	delete(f, "Timeout")
	delete(f, "RetryCount")
	delete(f, "Description")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayMCPServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayMCPServerResponseParams struct {
	// <p>创建结果</p>
	Result *CNAPIGwCreateCommonResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayMCPServerResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayMCPServerResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayMCPServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayMCPServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayMCPToolRequestParams struct {
	// <p>内容类型</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>服务 id</p>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

type CreateCloudNativeAPIGatewayMCPToolRequest struct {
	*tchttp.BaseRequest
	
	// <p>内容类型</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>服务 id</p>
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

func (r *CreateCloudNativeAPIGatewayMCPToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayMCPToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ContentType")
	delete(f, "DisplayName")
	delete(f, "ServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayMCPToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayMCPToolResponseParams struct {
	// <p>创建结果</p>
	Result *CNAPIGwCreateCommonResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayMCPToolResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayMCPToolResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayMCPToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayMCPToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewaySecretKeyRequestParams struct {
	// 实例 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>密钥协议类型。</p><p>枚举值：</p><ul><li>ApiKey</li><li>Basic</li><li>Hmac</li><li>OAuth2</li><li>JWT</li></ul>
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// <p>密钥名称，2-60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>密钥生成方式。</p><p>枚举值：</p><ul><li>System：系统自动生成</li><li>Custom：用户自定义（需传 SecretValue）</li><li>KMS：使用 KMS 密钥（需传 KmsKeyName 与 KmsKeyVersion）</li></ul>
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// <p>密钥归属资源类型。</p><p>枚举值：</p><ul><li>Consumer：消费者</li><li>ModelService：模型服务</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>KMS 密钥名称。GenerateType=KMS 时必填。</p>
	KmsKeyName *string `json:"KmsKeyName,omitnil,omitempty" name:"KmsKeyName"`

	// <p>KMS 密钥版本。GenerateType=KMS 时必填。</p>
	KmsKeyVersion *string `json:"KmsKeyVersion,omitnil,omitempty" name:"KmsKeyVersion"`

	// <p>密钥值，长度 8-256。GenerateType=Custom 时必填。</p>
	SecretValue *string `json:"SecretValue,omitnil,omitempty" name:"SecretValue"`

	// <p>密钥描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type CreateCloudNativeAPIGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>密钥协议类型。</p><p>枚举值：</p><ul><li>ApiKey</li><li>Basic</li><li>Hmac</li><li>OAuth2</li><li>JWT</li></ul>
	SecretType *string `json:"SecretType,omitnil,omitempty" name:"SecretType"`

	// <p>密钥名称，2-60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>密钥生成方式。</p><p>枚举值：</p><ul><li>System：系统自动生成</li><li>Custom：用户自定义（需传 SecretValue）</li><li>KMS：使用 KMS 密钥（需传 KmsKeyName 与 KmsKeyVersion）</li></ul>
	GenerateType *string `json:"GenerateType,omitnil,omitempty" name:"GenerateType"`

	// <p>密钥归属资源类型。</p><p>枚举值：</p><ul><li>Consumer：消费者</li><li>ModelService：模型服务</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>KMS 密钥名称。GenerateType=KMS 时必填。</p>
	KmsKeyName *string `json:"KmsKeyName,omitnil,omitempty" name:"KmsKeyName"`

	// <p>KMS 密钥版本。GenerateType=KMS 时必填。</p>
	KmsKeyVersion *string `json:"KmsKeyVersion,omitnil,omitempty" name:"KmsKeyVersion"`

	// <p>密钥值，长度 8-256。GenerateType=Custom 时必填。</p>
	SecretValue *string `json:"SecretValue,omitnil,omitempty" name:"SecretValue"`

	// <p>密钥描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *CreateCloudNativeAPIGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "SecretType")
	delete(f, "Name")
	delete(f, "GenerateType")
	delete(f, "ResourceType")
	delete(f, "KmsKeyName")
	delete(f, "KmsKeyVersion")
	delete(f, "SecretValue")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewaySecretKeyResponseParams struct {
	// <p>创建结果。包含成功标识与新建资源 ID。</p>
	Result *CNAPIGwCreateCommonResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DefaultKongRoute struct {
	// <p>服务名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>服务ID</p>
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// <p>HTTP Method</p>
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// <p>Http Path</p>
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayConsumerGroupRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 消费者组ID
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`
}

type DeleteCloudNativeAPIGatewayConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 消费者组ID
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`
}

func (r *DeleteCloudNativeAPIGatewayConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayConsumerGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayConsumerGroupResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayConsumerRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 消费者ID
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`
}

type DeleteCloudNativeAPIGatewayConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 消费者ID
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`
}

func (r *DeleteCloudNativeAPIGatewayConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayConsumerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayConsumerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayConsumerResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayLLMModelAPIRequestParams struct {
	// 网关 id。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 模型 API ID，全局唯一标识。
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`
}

type DeleteCloudNativeAPIGatewayLLMModelAPIRequest struct {
	*tchttp.BaseRequest
	
	// 网关 id。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 模型 API ID，全局唯一标识。
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`
}

func (r *DeleteCloudNativeAPIGatewayLLMModelAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayLLMModelAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ModelAPIId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayLLMModelAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayLLMModelAPIResponseParams struct {
	// <p>是否成功。</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayLLMModelAPIResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayLLMModelAPIResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayLLMModelAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayLLMModelAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayLLMModelServiceRequestParams struct {
	// 网关 id。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 模型服务 ID，全局唯一标识。
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`
}

type DeleteCloudNativeAPIGatewayLLMModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关 id。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 模型服务 ID，全局唯一标识。
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`
}

func (r *DeleteCloudNativeAPIGatewayLLMModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayLLMModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ModelServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayLLMModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayLLMModelServiceResponseParams struct {
	// <p>是否成功。</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayLLMModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayLLMModelServiceResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayLLMModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayLLMModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayMCPServerRequestParams struct {
	// <p>云原生API网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

type DeleteCloudNativeAPIGatewayMCPServerRequest struct {
	*tchttp.BaseRequest
	
	// <p>云原生API网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

func (r *DeleteCloudNativeAPIGatewayMCPServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayMCPServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayMCPServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayMCPServerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayMCPServerResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayMCPServerResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayMCPServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayMCPServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayMCPToolRequestParams struct {
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>工具id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>MCP 服务 id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

type DeleteCloudNativeAPIGatewayMCPToolRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>工具id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>MCP 服务 id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

func (r *DeleteCloudNativeAPIGatewayMCPToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayMCPToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ToolId")
	delete(f, "ServerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayMCPToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayMCPToolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayMCPToolResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayMCPToolResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayMCPToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayMCPToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewaySecretKeyRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

type DeleteCloudNativeAPIGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

func (r *DeleteCloudNativeAPIGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "SecretKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewaySecretKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayConsumerGroupRequestParams struct {
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组ID</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`
}

type DescribeCloudNativeAPIGatewayConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组ID</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`
}

func (r *DescribeCloudNativeAPIGatewayConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayConsumerGroupResponseParams struct {
	// <p>消费者组详情。</p>
	Result *CNAPIGwConsumerGroup `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayConsumerGroupResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayConsumerRequestParams struct {
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者ID</p>
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`
}

type DescribeCloudNativeAPIGatewayConsumerRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者ID</p>
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`
}

func (r *DescribeCloudNativeAPIGatewayConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayConsumerResponseParams struct {
	// <p>消费者详情</p>
	Result *CNAPIGwConsumer `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayConsumerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayConsumerResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelAPIRequestParams struct {
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型 API ID，全局唯一标识。</p>
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`
}

type DescribeCloudNativeAPIGatewayLLMModelAPIRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型 API ID，全局唯一标识。</p>
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ModelAPIId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayLLMModelAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelAPIResponseParams struct {
	// <p>模型 API 信息。</p>
	Result *CloudNativeAPIGatewayLLMModelAPI `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayLLMModelAPIResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayLLMModelAPIResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelAPIsRequestParams struct {
	// 网关 id。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>每页条数，范围 [1, 1000]，默认 10。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>起始位置，从 0 开始。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>过滤条件。当前未启用具体字段。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>模糊匹配模型 API 名称。</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>消费者组 ID（以 cg- 开头），与 UseToBind 搭配使用。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>是否用于绑定场景。true 时仅返回可被绑定到指定消费者组的模型 API。</p>
	UseToBind *bool `json:"UseToBind,omitnil,omitempty" name:"UseToBind"`
}

type DescribeCloudNativeAPIGatewayLLMModelAPIsRequest struct {
	*tchttp.BaseRequest
	
	// 网关 id。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>每页条数，范围 [1, 1000]，默认 10。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>起始位置，从 0 开始。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>过滤条件。当前未启用具体字段。</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>模糊匹配模型 API 名称。</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`

	// <p>消费者组 ID（以 cg- 开头），与 UseToBind 搭配使用。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>是否用于绑定场景。true 时仅返回可被绑定到指定消费者组的模型 API。</p>
	UseToBind *bool `json:"UseToBind,omitnil,omitempty" name:"UseToBind"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelAPIsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelAPIsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "Keyword")
	delete(f, "ConsumerGroupId")
	delete(f, "UseToBind")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayLLMModelAPIsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelAPIsResponseParams struct {
	// 模型 API 列表。
	Result *ListCloudNativeAPIGatewayLLMModelAPI `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayLLMModelAPIsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayLLMModelAPIsResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelAPIsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelAPIsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelServiceRequestParams struct {
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型服务 ID，全局唯一标识。</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`
}

type DescribeCloudNativeAPIGatewayLLMModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型服务 ID，全局唯一标识。</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ModelServiceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayLLMModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelServiceResponseParams struct {
	// <p>模型服务。</p>
	Result *CloudNativeAPIGatewayLLMModelService `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayLLMModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayLLMModelServiceResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelServicesRequestParams struct {
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>返回数量，默认为 10，最大值为 1000。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量，默认为 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>过滤条件，多个过滤条件之间是“与”的关系，支持 Name</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>通过模型 API 筛选模型服务</p>
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`

	// <p>通过密匙查询绑定的模型服务</p>
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`

	// <p>搜索关键词，模糊匹配 name 和 description</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeCloudNativeAPIGatewayLLMModelServicesRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>返回数量，默认为 10，最大值为 1000。</p>
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>偏移量，默认为 0。</p>
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>过滤条件，多个过滤条件之间是“与”的关系，支持 Name</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>通过模型 API 筛选模型服务</p>
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`

	// <p>通过密匙查询绑定的模型服务</p>
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`

	// <p>搜索关键词，模糊匹配 name 和 description</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	delete(f, "ModelAPIId")
	delete(f, "SecretKeyId")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayLLMModelServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMModelServicesResponseParams struct {
	// <p>模型服务列表。</p>
	Result *ListCloudNativeAPIGatewayLLMModelService `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayLLMModelServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayLLMModelServicesResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayLLMModelServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMModelServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMTokenUsageListRequestParams struct {
	// <p>网关实例Id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>查询起始时间戳</p><p>单位：秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>查询结束时间戳</p><p>单位：秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>查询过滤条件，Name取值为ConsumerId或ConsumerGroupId</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分页条件，每页条数</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页条件，分页偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例Id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>查询起始时间戳</p><p>单位：秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>查询结束时间戳</p><p>单位：秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>查询过滤条件，Name取值为ConsumerId或ConsumerGroupId</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// <p>分页条件，每页条数</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页条件，分页偏移量</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayLLMTokenUsageListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMTokenUsageListResponseParams struct {
	// <p>查询Token用量明细结果</p>
	Result *AIGWLLMTokenUsageListResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayLLMTokenUsageListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayLLMTokenUsageListResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequestParams struct {
	// <p>网关实例Id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>查询开始时间戳</p><p>单位：秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>查询结束时间戳</p><p>单位：秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>查询过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例Id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>查询开始时间戳</p><p>单位：秒</p>
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// <p>查询结束时间戳</p><p>单位：秒</p>
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// <p>查询过滤条件</p>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponseParams struct {
	// <p>请求结果</p>
	Result *AIGWLLMTokenUsageStatisticsResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayLLMTokenUsageStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerACLRequestParams struct {
	// <p>网关实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

type DescribeCloudNativeAPIGatewayMCPServerACLRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayMCPServerACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerACLResponseParams struct {
	// <p>MCP 服务 ACL 列表结果</p>
	Result *AIGWMCPServerACLResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayMCPServerACLResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayMCPServerACLResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerAuthRequestParams struct {
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

type DescribeCloudNativeAPIGatewayMCPServerAuthRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayMCPServerAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerAuthResponseParams struct {
	// <p>MCP服务认证查询结果</p>
	Result *AIGWMCPServerAuthResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayMCPServerAuthResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayMCPServerAuthResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerListRequestParams struct {
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>分页大小</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCloudNativeAPIGatewayMCPServerListRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>分页大小</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayMCPServerListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerListResponseParams struct {
	// <p>MCP Server 列表结果</p>
	Result *AIGWMCPServerList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayMCPServerListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayMCPServerListResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerRequestParams struct {
	// <p>云原生API网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

type DescribeCloudNativeAPIGatewayMCPServerRequest struct {
	*tchttp.BaseRequest
	
	// <p>云原生API网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayMCPServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPServerResponseParams struct {
	// <p>mcp server详情</p>
	Result *AIGWMCPServer `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayMCPServerResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayMCPServerResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayMCPServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPToolACLListRequestParams struct {
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP 服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>分页大小</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

type DescribeCloudNativeAPIGatewayMCPToolACLListRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP 服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>分页大小</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>分页偏移</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// <p>搜索关键词</p>
	Keyword *string `json:"Keyword,omitnil,omitempty" name:"Keyword"`
}

func (r *DescribeCloudNativeAPIGatewayMCPToolACLListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPToolACLListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Keyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayMCPToolACLListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPToolACLListResponseParams struct {
	// <p>MCP 服务 Tool ACL 列表结果</p>
	Result *AIGWMCPToolACLListResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayMCPToolACLListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayMCPToolACLListResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayMCPToolACLListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPToolACLListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPToolListRequestParams struct {
	// <p>实例 id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务 id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>条数</p><p>取值范围：[1, 100]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>开始位置</p><p>取值范围：[1, 100000]</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCloudNativeAPIGatewayMCPToolListRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务 id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>条数</p><p>取值范围：[1, 100]</p>
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// <p>开始位置</p><p>取值范围：[1, 100000]</p>
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeCloudNativeAPIGatewayMCPToolListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPToolListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayMCPToolListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPToolListResponseParams struct {
	// <p>tool 列表</p>
	Result *CNAPIGwMCPToolList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayMCPToolListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayMCPToolListResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayMCPToolListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPToolListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPToolRequestParams struct {
	// <p>网关实例 id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP Server id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>工具 id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

type DescribeCloudNativeAPIGatewayMCPToolRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例 id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP Server id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>工具 id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

func (r *DescribeCloudNativeAPIGatewayMCPToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "ToolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayMCPToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayMCPToolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayMCPToolResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayMCPToolResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayMCPToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayMCPToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewaySecretKeyRequestParams struct {
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>密钥id</p>
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

type DescribeCloudNativeAPIGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>密钥id</p>
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

func (r *DescribeCloudNativeAPIGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "SecretKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewaySecretKeyResponseParams struct {
	// <p>密钥详情。</p>
	Result *CNAPIGwSecretKey `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewaySecretKeyValueRequestParams struct {
	// 实例 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

type DescribeCloudNativeAPIGatewaySecretKeyValueRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

func (r *DescribeCloudNativeAPIGatewaySecretKeyValueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewaySecretKeyValueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "SecretKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewaySecretKeyValueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewaySecretKeyValueResponseParams struct {
	// 密钥值
	Result *string `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewaySecretKeyValueResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewaySecretKeyValueResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaySecretKeyValueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewaySecretKeyValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤参数值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ListCloudNativeAPIGatewayLLMModelAPI struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// AI 网关模型 API 列表。
	DataList []*CloudNativeAPIGatewayLLMModelAPI `json:"DataList,omitnil,omitempty" name:"DataList"`
}

type ListCloudNativeAPIGatewayLLMModelService struct {
	// 模型服务数量。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 模型服务列表。
	DataList []*CloudNativeAPIGatewayLLMModelService `json:"DataList,omitnil,omitempty" name:"DataList"`
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayConsumerGroupRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组 ID（以 cg- 开头）。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>消费者组名称，最长 60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>启用状态。</p><p>枚举值：</p><ul><li>Enable：启用</li><li>Disable：禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>消费者组描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyCloudNativeAPIGatewayConsumerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组 ID（以 cg- 开头）。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>消费者组名称，最长 60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>启用状态。</p><p>枚举值：</p><ul><li>Enable：启用</li><li>Disable：禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>消费者组描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyCloudNativeAPIGatewayConsumerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayConsumerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerGroupId")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayConsumerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayConsumerGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayConsumerGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayConsumerGroupResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayConsumerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayConsumerRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者 ID。</p>
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// <p>消费者名称，最长 60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>消费者描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyCloudNativeAPIGatewayConsumerRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者 ID。</p>
	ConsumerId *string `json:"ConsumerId,omitnil,omitempty" name:"ConsumerId"`

	// <p>消费者名称，最长 60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>消费者描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyCloudNativeAPIGatewayConsumerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayConsumerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayConsumerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayConsumerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayConsumerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayConsumerResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayConsumerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayLLMModelAPIRequestParams struct {
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型 API ID，全局唯一标识。</p>
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`

	// <p>模型 API 名称，最长 60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>统一前缀路径（可选）。例如 /v1/openai。</p>
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// <p>模型 API 描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>关联的模型服务 ID 列表，长度 1-10。</p>
	ListModelServiceId []*string `json:"ListModelServiceId,omitnil,omitempty" name:"ListModelServiceId"`

	// <p>多模型服务路由策略。ListModelServiceId 多于 1 项时必填。</p>
	ModelServiceRoute *CloudNativeAPIGatewayLLMModelServiceRoute `json:"ModelServiceRoute,omitnil,omitempty" name:"ModelServiceRoute"`

	// <p>Header 路由匹配规则。当前仅支持 Operator=exact。</p>
	MatchHeaders []*AIGWKVMatch `json:"MatchHeaders,omitnil,omitempty" name:"MatchHeaders"`

	// <p>是否启用跨服务 Fallback。</p>
	EnableCrossServiceFallback *bool `json:"EnableCrossServiceFallback,omitnil,omitempty" name:"EnableCrossServiceFallback"`

	// <p>跨服务 Fallback 配置。EnableCrossServiceFallback=true 时必填。</p>
	CrossServiceFallbackConfig *AIGWCrossServiceFallbackConfig `json:"CrossServiceFallbackConfig,omitnil,omitempty" name:"CrossServiceFallbackConfig"`

	// <p>标签过滤策略。需要网关版本 ≥ 3.9.4。</p>
	TagFilter *AIGWTagFilter `json:"TagFilter,omitnil,omitempty" name:"TagFilter"`

	// <p>日志输出配置。需要网关版本 ≥ 3.9.4。</p>
	LogConfig *AIGWLogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`
}

type ModifyCloudNativeAPIGatewayLLMModelAPIRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型 API ID，全局唯一标识。</p>
	ModelAPIId *string `json:"ModelAPIId,omitnil,omitempty" name:"ModelAPIId"`

	// <p>模型 API 名称，最长 60 字符。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>统一前缀路径（可选）。例如 /v1/openai。</p>
	BasePath *string `json:"BasePath,omitnil,omitempty" name:"BasePath"`

	// <p>模型 API 描述。最长 200 字符。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>关联的模型服务 ID 列表，长度 1-10。</p>
	ListModelServiceId []*string `json:"ListModelServiceId,omitnil,omitempty" name:"ListModelServiceId"`

	// <p>多模型服务路由策略。ListModelServiceId 多于 1 项时必填。</p>
	ModelServiceRoute *CloudNativeAPIGatewayLLMModelServiceRoute `json:"ModelServiceRoute,omitnil,omitempty" name:"ModelServiceRoute"`

	// <p>Header 路由匹配规则。当前仅支持 Operator=exact。</p>
	MatchHeaders []*AIGWKVMatch `json:"MatchHeaders,omitnil,omitempty" name:"MatchHeaders"`

	// <p>是否启用跨服务 Fallback。</p>
	EnableCrossServiceFallback *bool `json:"EnableCrossServiceFallback,omitnil,omitempty" name:"EnableCrossServiceFallback"`

	// <p>跨服务 Fallback 配置。EnableCrossServiceFallback=true 时必填。</p>
	CrossServiceFallbackConfig *AIGWCrossServiceFallbackConfig `json:"CrossServiceFallbackConfig,omitnil,omitempty" name:"CrossServiceFallbackConfig"`

	// <p>标签过滤策略。需要网关版本 ≥ 3.9.4。</p>
	TagFilter *AIGWTagFilter `json:"TagFilter,omitnil,omitempty" name:"TagFilter"`

	// <p>日志输出配置。需要网关版本 ≥ 3.9.4。</p>
	LogConfig *AIGWLogConfig `json:"LogConfig,omitnil,omitempty" name:"LogConfig"`
}

func (r *ModifyCloudNativeAPIGatewayLLMModelAPIRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayLLMModelAPIRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ModelAPIId")
	delete(f, "Name")
	delete(f, "BasePath")
	delete(f, "Description")
	delete(f, "ListModelServiceId")
	delete(f, "ModelServiceRoute")
	delete(f, "MatchHeaders")
	delete(f, "EnableCrossServiceFallback")
	delete(f, "CrossServiceFallbackConfig")
	delete(f, "TagFilter")
	delete(f, "LogConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayLLMModelAPIRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayLLMModelAPIResponseParams struct {
	// <p>是否成功。</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayLLMModelAPIResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayLLMModelAPIResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayLLMModelAPIResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayLLMModelAPIResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayLLMModelServiceRequestParams struct {
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型服务 ID，全局唯一标识。</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// <p>修改服务名称，长度2-50字符，支持中英文、数字、下划线。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>修改默认模型，模型选择方式为 Specify 时必填。</p>
	DefaultModel *string `json:"DefaultModel,omitnil,omitempty" name:"DefaultModel"`

	// <p>修改模型选择方式，选项：Specify（指定模型）、PassThrough（透传请求模型）。</p>
	ModelSelector *string `json:"ModelSelector,omitnil,omitempty" name:"ModelSelector"`

	// <p>修改开启模型降级，模型选择方式为 Specify 时必填。</p>
	EnableModelFallback *bool `json:"EnableModelFallback,omitnil,omitempty" name:"EnableModelFallback"`

	// <p>修改可以配置备用模型规则，EnableSpecifyModelFallback 为 true 时必填。</p>
	ModelFallbackRule *CloudNativeAPIGatewayLLMModelFallbackRule `json:"ModelFallbackRule,omitnil,omitempty" name:"ModelFallbackRule"`

	// <p>修改开启模型参数校验，是否校验客户端传递的 model 参数, 模型选择方式为 PassThrough 时必填</p>
	EnableModelParamCheck *bool `json:"EnableModelParamCheck,omitnil,omitempty" name:"EnableModelParamCheck"`

	// <p>修改模型检验信息，EnableModelParamCheck 为 true 时必填。</p>
	ModelParamCheckRule *CloudNativeAPIGatewayLLMModelParamCheckInfo `json:"ModelParamCheckRule,omitnil,omitempty" name:"ModelParamCheckRule"`

	// <p>修改描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>修改模型服务地址</p>
	UpstreamURL *string `json:"UpstreamURL,omitnil,omitempty" name:"UpstreamURL"`

	// <p>连接超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：10000</p>
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>写入超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	WriteTimeout *int64 `json:"WriteTimeout,omitnil,omitempty" name:"WriteTimeout"`

	// <p>读取超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// <p>重试次数</p><p>取值范围：[0, 5]</p><p>单位：次</p><p>默认值：0</p>
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// <p>路径拼接模式</p><p>枚举值：</p><ul><li>FixedPath： 固定路径</li><li>AutoConcat： 自动拼接</li></ul>
	UpstreamUrlMode *string `json:"UpstreamUrlMode,omitnil,omitempty" name:"UpstreamUrlMode"`

	// <p>SNI</p>
	SNI *string `json:"SNI,omitnil,omitempty" name:"SNI"`

	// <p>模型服务级别的配额上限（RPM/TPM）。需要网关版本 ≥ 3.9.4。</p>
	QuotaLimit *AIGWLLMQuotaLimit `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`

	// <p>标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ModifyCloudNativeAPIGatewayLLMModelServiceRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关 id。</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>模型服务 ID，全局唯一标识。</p>
	ModelServiceId *string `json:"ModelServiceId,omitnil,omitempty" name:"ModelServiceId"`

	// <p>修改服务名称，长度2-50字符，支持中英文、数字、下划线。</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>修改默认模型，模型选择方式为 Specify 时必填。</p>
	DefaultModel *string `json:"DefaultModel,omitnil,omitempty" name:"DefaultModel"`

	// <p>修改模型选择方式，选项：Specify（指定模型）、PassThrough（透传请求模型）。</p>
	ModelSelector *string `json:"ModelSelector,omitnil,omitempty" name:"ModelSelector"`

	// <p>修改开启模型降级，模型选择方式为 Specify 时必填。</p>
	EnableModelFallback *bool `json:"EnableModelFallback,omitnil,omitempty" name:"EnableModelFallback"`

	// <p>修改可以配置备用模型规则，EnableSpecifyModelFallback 为 true 时必填。</p>
	ModelFallbackRule *CloudNativeAPIGatewayLLMModelFallbackRule `json:"ModelFallbackRule,omitnil,omitempty" name:"ModelFallbackRule"`

	// <p>修改开启模型参数校验，是否校验客户端传递的 model 参数, 模型选择方式为 PassThrough 时必填</p>
	EnableModelParamCheck *bool `json:"EnableModelParamCheck,omitnil,omitempty" name:"EnableModelParamCheck"`

	// <p>修改模型检验信息，EnableModelParamCheck 为 true 时必填。</p>
	ModelParamCheckRule *CloudNativeAPIGatewayLLMModelParamCheckInfo `json:"ModelParamCheckRule,omitnil,omitempty" name:"ModelParamCheckRule"`

	// <p>修改描述。</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>修改模型服务地址</p>
	UpstreamURL *string `json:"UpstreamURL,omitnil,omitempty" name:"UpstreamURL"`

	// <p>连接超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：10000</p>
	ConnectTimeout *int64 `json:"ConnectTimeout,omitnil,omitempty" name:"ConnectTimeout"`

	// <p>写入超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	WriteTimeout *int64 `json:"WriteTimeout,omitnil,omitempty" name:"WriteTimeout"`

	// <p>读取超时时间</p><p>取值范围：[1, 3600000]</p><p>单位：毫秒</p><p>默认值：60000</p>
	ReadTimeout *int64 `json:"ReadTimeout,omitnil,omitempty" name:"ReadTimeout"`

	// <p>重试次数</p><p>取值范围：[0, 5]</p><p>单位：次</p><p>默认值：0</p>
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// <p>路径拼接模式</p><p>枚举值：</p><ul><li>FixedPath： 固定路径</li><li>AutoConcat： 自动拼接</li></ul>
	UpstreamUrlMode *string `json:"UpstreamUrlMode,omitnil,omitempty" name:"UpstreamUrlMode"`

	// <p>SNI</p>
	SNI *string `json:"SNI,omitnil,omitempty" name:"SNI"`

	// <p>模型服务级别的配额上限（RPM/TPM）。需要网关版本 ≥ 3.9.4。</p>
	QuotaLimit *AIGWLLMQuotaLimit `json:"QuotaLimit,omitnil,omitempty" name:"QuotaLimit"`

	// <p>标签</p>
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

func (r *ModifyCloudNativeAPIGatewayLLMModelServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayLLMModelServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ModelServiceId")
	delete(f, "Name")
	delete(f, "DefaultModel")
	delete(f, "ModelSelector")
	delete(f, "EnableModelFallback")
	delete(f, "ModelFallbackRule")
	delete(f, "EnableModelParamCheck")
	delete(f, "ModelParamCheckRule")
	delete(f, "Description")
	delete(f, "UpstreamURL")
	delete(f, "ConnectTimeout")
	delete(f, "WriteTimeout")
	delete(f, "ReadTimeout")
	delete(f, "Retries")
	delete(f, "UpstreamUrlMode")
	delete(f, "SNI")
	delete(f, "QuotaLimit")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayLLMModelServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayLLMModelServiceResponseParams struct {
	// <p>是否成功</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayLLMModelServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayLLMModelServiceResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayLLMModelServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayLLMModelServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerACLRequestParams struct {
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>ACL类型</p><p>枚举值：</p><ul><li>None： 不开启</li><li>Allow： 白名单</li><li>Deny： 黑名单</li></ul><p>默认值：None</p>
	ACLType *string `json:"ACLType,omitnil,omitempty" name:"ACLType"`

	// <p>需要关联的消费者ID列表</p>
	ACLConsumerIds []*string `json:"ACLConsumerIds,omitnil,omitempty" name:"ACLConsumerIds"`

	// <p>需要关联的消费者组ID列表</p>
	ACLConsumerGroupIds []*string `json:"ACLConsumerGroupIds,omitnil,omitempty" name:"ACLConsumerGroupIds"`
}

type ModifyCloudNativeAPIGatewayMCPServerACLRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>ACL类型</p><p>枚举值：</p><ul><li>None： 不开启</li><li>Allow： 白名单</li><li>Deny： 黑名单</li></ul><p>默认值：None</p>
	ACLType *string `json:"ACLType,omitnil,omitempty" name:"ACLType"`

	// <p>需要关联的消费者ID列表</p>
	ACLConsumerIds []*string `json:"ACLConsumerIds,omitnil,omitempty" name:"ACLConsumerIds"`

	// <p>需要关联的消费者组ID列表</p>
	ACLConsumerGroupIds []*string `json:"ACLConsumerGroupIds,omitnil,omitempty" name:"ACLConsumerGroupIds"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "ACLType")
	delete(f, "ACLConsumerIds")
	delete(f, "ACLConsumerGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayMCPServerACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayMCPServerACLResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayMCPServerACLResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerAuthRequestParams struct {
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>认证类型</p><p>枚举值：</p><ul><li>None： 无认证</li><li>ApiKey： API Key认证</li></ul>
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`
}

type ModifyCloudNativeAPIGatewayMCPServerAuthRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>认证类型</p><p>枚举值：</p><ul><li>None： 无认证</li><li>ApiKey： API Key认证</li></ul>
	AuthType *string `json:"AuthType,omitnil,omitempty" name:"AuthType"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "AuthType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayMCPServerAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerAuthResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayMCPServerAuthResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayMCPServerAuthResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerRequestParams struct {
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>服务 id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>后端类型</p><p>枚举值：</p><ul><li>HostIP： 域名 ip</li><li>MCPRegistry： MCP 注册中心</li><li>VirtualMCPServer： 虚拟MCP 服务</li></ul>
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// <p>超时时间，单位ms，最大60000</p>
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>重试次数，最大3次</p>
	RetryCount *uint64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>注册中心来源信息</p>
	UpstreamInfo *AIGWMCPUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// <p>会话配置</p>
	SessionConfig *AIGWMCPSessionConfig `json:"SessionConfig,omitnil,omitempty" name:"SessionConfig"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否启用健康检查</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>健康检查配置</p>
	HealthCheck *AIGWHealthCheckSetting `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`
}

type ModifyCloudNativeAPIGatewayMCPServerRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>展示名字</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>服务 id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>后端类型</p><p>枚举值：</p><ul><li>HostIP： 域名 ip</li><li>MCPRegistry： MCP 注册中心</li><li>VirtualMCPServer： 虚拟MCP 服务</li></ul>
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// <p>超时时间，单位ms，最大60000</p>
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// <p>重试次数，最大3次</p>
	RetryCount *uint64 `json:"RetryCount,omitnil,omitempty" name:"RetryCount"`

	// <p>注册中心来源信息</p>
	UpstreamInfo *AIGWMCPUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// <p>会话配置</p>
	SessionConfig *AIGWMCPSessionConfig `json:"SessionConfig,omitnil,omitempty" name:"SessionConfig"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>是否启用健康检查</p>
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// <p>健康检查配置</p>
	HealthCheck *AIGWHealthCheckSetting `json:"HealthCheck,omitnil,omitempty" name:"HealthCheck"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "DisplayName")
	delete(f, "ServerId")
	delete(f, "UpstreamType")
	delete(f, "Timeout")
	delete(f, "RetryCount")
	delete(f, "UpstreamInfo")
	delete(f, "SessionConfig")
	delete(f, "Description")
	delete(f, "EnableHealthCheck")
	delete(f, "HealthCheck")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayMCPServerRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayMCPServerResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayMCPServerResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerStatusRequestParams struct {
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>mcp server id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>mcp server状态</p><p>枚举值：</p><ul><li>Online： 上线</li><li>Offline： 下线</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyCloudNativeAPIGatewayMCPServerStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>mcp server id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>mcp server状态</p><p>枚举值：</p><ul><li>Online： 上线</li><li>Offline： 下线</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayMCPServerStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPServerStatusResponseParams struct {
	// <p>创建结果</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayMCPServerStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayMCPServerStatusResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayMCPServerStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPServerStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPToolACLRequestParams struct {
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>MCP工具ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>鉴权类型</p><p>枚举值：</p><ul><li>None： 继承server鉴权类型</li><li>Allow： 白名单</li><li>Deny： 黑名单</li></ul>
	ACLType *string `json:"ACLType,omitnil,omitempty" name:"ACLType"`

	// <p>需要关联的消费者ID列表</p>
	ACLConsumerIds []*string `json:"ACLConsumerIds,omitnil,omitempty" name:"ACLConsumerIds"`

	// <p>需要关联的消费者组ID列表</p>
	ACLConsumerGroupIds []*string `json:"ACLConsumerGroupIds,omitnil,omitempty" name:"ACLConsumerGroupIds"`
}

type ModifyCloudNativeAPIGatewayMCPToolACLRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCP服务ID</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>MCP工具ID</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>鉴权类型</p><p>枚举值：</p><ul><li>None： 继承server鉴权类型</li><li>Allow： 白名单</li><li>Deny： 黑名单</li></ul>
	ACLType *string `json:"ACLType,omitnil,omitempty" name:"ACLType"`

	// <p>需要关联的消费者ID列表</p>
	ACLConsumerIds []*string `json:"ACLConsumerIds,omitnil,omitempty" name:"ACLConsumerIds"`

	// <p>需要关联的消费者组ID列表</p>
	ACLConsumerGroupIds []*string `json:"ACLConsumerGroupIds,omitnil,omitempty" name:"ACLConsumerGroupIds"`
}

func (r *ModifyCloudNativeAPIGatewayMCPToolACLRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPToolACLRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "ToolId")
	delete(f, "ACLType")
	delete(f, "ACLConsumerIds")
	delete(f, "ACLConsumerGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayMCPToolACLRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPToolACLResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayMCPToolACLResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayMCPToolACLResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayMCPToolACLResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPToolACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPToolRequestParams struct {
	// <p>网关实例 id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCPserverId</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>工具 id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>工具名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>路径</p>
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>报文格式</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>请求方法</p><p>枚举值：</p><ul><li>GET： GET</li><li>PUT： PUT</li><li>POST： POST</li><li>DELETE： DELETE</li><li>PATCH： PATCH</li></ul>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>展示</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>输入参数</p>
	InputParams []*CNAPIGwMCPToolParam `json:"InputParams,omitnil,omitempty" name:"InputParams"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>tool版本</p>
	ToolVersion *string `json:"ToolVersion,omitnil,omitempty" name:"ToolVersion"`
}

type ModifyCloudNativeAPIGatewayMCPToolRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例 id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>MCPserverId</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>工具 id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`

	// <p>工具名字</p>
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// <p>路径</p>
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// <p>报文格式</p>
	ContentType *string `json:"ContentType,omitnil,omitempty" name:"ContentType"`

	// <p>请求方法</p><p>枚举值：</p><ul><li>GET： GET</li><li>PUT： PUT</li><li>POST： POST</li><li>DELETE： DELETE</li><li>PATCH： PATCH</li></ul>
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// <p>展示</p>
	DisplayName *string `json:"DisplayName,omitnil,omitempty" name:"DisplayName"`

	// <p>输入参数</p>
	InputParams []*CNAPIGwMCPToolParam `json:"InputParams,omitnil,omitempty" name:"InputParams"`

	// <p>描述</p>
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// <p>tool版本</p>
	ToolVersion *string `json:"ToolVersion,omitnil,omitempty" name:"ToolVersion"`
}

func (r *ModifyCloudNativeAPIGatewayMCPToolRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPToolRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "ToolId")
	delete(f, "Name")
	delete(f, "Path")
	delete(f, "ContentType")
	delete(f, "Method")
	delete(f, "DisplayName")
	delete(f, "InputParams")
	delete(f, "Description")
	delete(f, "ToolVersion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayMCPToolRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPToolResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayMCPToolResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayMCPToolResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayMCPToolResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPToolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPToolStatusRequestParams struct {
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>mcp server id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>mcp tool 状态</p><p>枚举值：</p><ul><li>Enable： 启用</li><li>Disable： 禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>mcp tool id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

type ModifyCloudNativeAPIGatewayMCPToolStatusRequest struct {
	*tchttp.BaseRequest
	
	// <p>实例 ID</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>mcp server id</p>
	ServerId *string `json:"ServerId,omitnil,omitempty" name:"ServerId"`

	// <p>mcp tool 状态</p><p>枚举值：</p><ul><li>Enable： 启用</li><li>Disable： 禁用</li></ul>
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// <p>mcp tool id</p>
	ToolId *string `json:"ToolId,omitnil,omitempty" name:"ToolId"`
}

func (r *ModifyCloudNativeAPIGatewayMCPToolStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPToolStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServerId")
	delete(f, "Status")
	delete(f, "ToolId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayMCPToolStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayMCPToolStatusResponseParams struct {
	// <p>创建结果</p>
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayMCPToolStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayMCPToolStatusResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayMCPToolStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayMCPToolStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewaySecretKeyRequestParams struct {
	// 实例 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 密钥名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`

	// 描述,200字以内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyCloudNativeAPIGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 实例 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 密钥名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`

	// 描述,200字以内
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

func (r *ModifyCloudNativeAPIGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "SecretKeyId")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewaySecretKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveCloudNativeAPIGatewayConsumerGroupAuthRequestParams struct {
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>授权资源类型。</p><p>枚举值：</p><ul><li>ModelAPI：模型 API</li><li>MCPServer：MCP Server</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>对应资源的 ID。</p><ul><li>ResourceType=ModelAPI 时是模型 API ID</li><li>ResourceType=MCPServer 时是 MCP Server ID</li></ul>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>消费者组 ID 列表（每个 ID 以 cg- 开头），长度 1-10。</p>
	ConsumerGroupIds []*string `json:"ConsumerGroupIds,omitnil,omitempty" name:"ConsumerGroupIds"`
}

type RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest struct {
	*tchttp.BaseRequest
	
	// <p>网关实例id</p>
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>授权资源类型。</p><p>枚举值：</p><ul><li>ModelAPI：模型 API</li><li>MCPServer：MCP Server</li></ul>
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// <p>对应资源的 ID。</p><ul><li>ResourceType=ModelAPI 时是模型 API ID</li><li>ResourceType=MCPServer 时是 MCP Server ID</li></ul>
	ResourceId *string `json:"ResourceId,omitnil,omitempty" name:"ResourceId"`

	// <p>消费者组 ID 列表（每个 ID 以 cg- 开头），长度 1-10。</p>
	ConsumerGroupIds []*string `json:"ConsumerGroupIds,omitnil,omitempty" name:"ConsumerGroupIds"`
}

func (r *RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ResourceType")
	delete(f, "ResourceId")
	delete(f, "ConsumerGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveCloudNativeAPIGatewayConsumerGroupAuthRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveCloudNativeAPIGatewayConsumerGroupAuthResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveCloudNativeAPIGatewayConsumerGroupAuthResponse struct {
	*tchttp.BaseResponse
	Response *RemoveCloudNativeAPIGatewayConsumerGroupAuthResponseParams `json:"Response"`
}

func (r *RemoveCloudNativeAPIGatewayConsumerGroupAuthResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveCloudNativeAPIGatewayConsumerGroupAuthResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveCloudNativeAPIGatewayConsumerInGroupRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组 ID（以 cg- 开头）。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>消费者 ID 列表，长度 1-10。</p>
	ConsumerIds []*string `json:"ConsumerIds,omitnil,omitempty" name:"ConsumerIds"`
}

type RemoveCloudNativeAPIGatewayConsumerInGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// <p>消费者组 ID（以 cg- 开头）。</p>
	ConsumerGroupId *string `json:"ConsumerGroupId,omitnil,omitempty" name:"ConsumerGroupId"`

	// <p>消费者 ID 列表，长度 1-10。</p>
	ConsumerIds []*string `json:"ConsumerIds,omitnil,omitempty" name:"ConsumerIds"`
}

func (r *RemoveCloudNativeAPIGatewayConsumerInGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveCloudNativeAPIGatewayConsumerInGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ConsumerGroupId")
	delete(f, "ConsumerIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveCloudNativeAPIGatewayConsumerInGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveCloudNativeAPIGatewayConsumerInGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveCloudNativeAPIGatewayConsumerInGroupResponse struct {
	*tchttp.BaseResponse
	Response *RemoveCloudNativeAPIGatewayConsumerInGroupResponseParams `json:"Response"`
}

func (r *RemoveCloudNativeAPIGatewayConsumerInGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveCloudNativeAPIGatewayConsumerInGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindCloudNativeAPIGatewaySecretKeyRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源ID，当前最多支持一个
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

type UnbindCloudNativeAPIGatewaySecretKeyRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 资源类型
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`

	// 资源ID，当前最多支持一个
	ResourceIds []*string `json:"ResourceIds,omitnil,omitempty" name:"ResourceIds"`

	// 密钥id
	SecretKeyId *string `json:"SecretKeyId,omitnil,omitempty" name:"SecretKeyId"`
}

func (r *UnbindCloudNativeAPIGatewaySecretKeyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCloudNativeAPIGatewaySecretKeyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ResourceType")
	delete(f, "ResourceIds")
	delete(f, "SecretKeyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindCloudNativeAPIGatewaySecretKeyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindCloudNativeAPIGatewaySecretKeyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindCloudNativeAPIGatewaySecretKeyResponse struct {
	*tchttp.BaseResponse
	Response *UnbindCloudNativeAPIGatewaySecretKeyResponseParams `json:"Response"`
}

func (r *UnbindCloudNativeAPIGatewaySecretKeyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCloudNativeAPIGatewaySecretKeyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}