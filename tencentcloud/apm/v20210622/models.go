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

package v20210622

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type APMKV struct {
	// Key 值定义
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value 值定义
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`
}

type APMKVItem struct {
	// Key 值定义
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// Value 值定义
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type AgentOperationConfigView struct {
	// 当前接口配置是否开启了接口白名单配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionValid *bool `json:"RetentionValid,omitnil,omitempty" name:"RetentionValid"`

	// RetentionValid为false时生效，接口配置中的黑名单配置，配置中的接口不采集
	// 注意：此字段可能返回 null，表示取不到有效值。
	IgnoreOperation *string `json:"IgnoreOperation,omitnil,omitempty" name:"IgnoreOperation"`

	// RetentionValid为true时生效，接口配置中的白名单配置，仅采集配置中的接口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RetentionOperation *string `json:"RetentionOperation,omitnil,omitempty" name:"RetentionOperation"`
}

type ApmAgentInfo struct {
	// Agent 下载地址
	AgentDownloadURL *string `json:"AgentDownloadURL,omitnil,omitempty" name:"AgentDownloadURL"`

	// Collector 上报地址
	CollectorURL *string `json:"CollectorURL,omitnil,omitempty" name:"CollectorURL"`

	// Token 信息
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// 外网上报地址
	PublicCollectorURL *string `json:"PublicCollectorURL,omitnil,omitempty" name:"PublicCollectorURL"`

	// 自研 VPC 上报地址
	InnerCollectorURL *string `json:"InnerCollectorURL,omitnil,omitempty" name:"InnerCollectorURL"`

	// 内网上报地址( Private Link 上报地址)
	PrivateLinkCollectorURL *string `json:"PrivateLinkCollectorURL,omitnil,omitempty" name:"PrivateLinkCollectorURL"`
}

type ApmAppConfig struct {
	// 实例ID
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// URL收敛开关
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// URL收敛阈值
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// URL收敛正则
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// 异常过滤正则
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// 错误码过滤
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// 服务组件类型
	Components *string `json:"Components,omitnil,omitempty" name:"Components"`

	// URL排除正则
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// 日志来源
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 日志所在地域
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// 是否开启日志 0 关 1 开
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// 日志主题ID
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// 需过滤的接口名
	IgnoreOperationName *string `json:"IgnoreOperationName,omitnil,omitempty" name:"IgnoreOperationName"`

	// CLS日志集 | ES集群ID
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 探针每秒上报trace数
	TraceRateLimit *int64 `json:"TraceRateLimit,omitnil,omitempty" name:"TraceRateLimit"`

	// 是否开启线程剖析
	EnableSnapshot *bool `json:"EnableSnapshot,omitnil,omitempty" name:"EnableSnapshot"`

	// 线程剖析超时阈值
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// 是否开启agent
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// 组件列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// 是否开启链路压缩
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// 是否开启应用诊断开关
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// 探针接口相关配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitnil,omitempty" name:"AgentOperationConfigView"`

	// 是否开启应用日志配置
	EnableLogConfig *bool `json:"EnableLogConfig,omitnil,omitempty" name:"EnableLogConfig"`

	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 应用是否开启dashboard配置： false 关（与业务系统保持一致）/true 开（应用级配置）
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitnil,omitempty" name:"EnableDashboardConfig"`

	// 是否关联dashboard： 0 关 1 开
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// dashboard ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// 是否开启应用级别配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitnil,omitempty" name:"EnableSecurityConfig"`

	// 是否开启组件漏洞检测
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// 是否开启SQL注入分析
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// 是否开启远程命令执行分析
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// 是否开启内存马检测分析
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// CLS索引类型(0=全文索引，1=键值索引)
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// traceId的索引key: 当CLS索引类型为键值索引时生效
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// 是否开启删除任意文件检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// 是否开启读取任意文件检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// 是否开启上传任意文件检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// 是否开启包含任意文件检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// 是否开启目录遍历检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// 是否开启模板引擎注入检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// 是否开启脚本引擎注入检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// 是否开启表达式注入检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// 是否开启JNDI注入检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// 是否开启JNI注入检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// 是否开启Webshell后门检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// 是否开启反序列化检测（0-关闭，1-开启）
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// 接口名称自动收敛开关（0-关闭，1-开启）
	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitnil,omitempty" name:"UrlAutoConvergenceEnable"`

	// URL长分段收敛阈值
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// URL数字分段收敛阈值
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// 探针熔断内存阈值
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// 探针熔断CPU阈值
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// 是否开启SQL参数获取
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// 慢SQL阈值
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`
}

type ApmApplicationConfigView struct {
	// 业务系统 ID
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 应用名	
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 接口过滤
	OperationNameFilter *string `json:"OperationNameFilter,omitnil,omitempty" name:"OperationNameFilter"`

	// 错误类型过滤
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// HTTP 状态码过滤
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// 应用诊断开关（已废弃）
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// URL 收敛开关 0 关 1 开
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// URL 收敛阈值	
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// URL 收敛规则正则	
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// URL 排除规则正则
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// 是否开启日志 0 关 1 开
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// 日志源	
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 日志集 
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 日志主题
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// 方法栈快照开关 true 开启 false 关闭
	SnapshotEnable *bool `json:"SnapshotEnable,omitnil,omitempty" name:"SnapshotEnable"`

	// 慢调用监听触发阈值
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// 探针总开关
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// 组件列表开关（已废弃）
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// 链路压缩开关（已废弃）
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// 探针熔断内存阈值
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// 探针熔断CPU阈值
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// 是否开启SQL参数获取
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// 慢SQL阈值
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`
}

type ApmAssociation struct {
	// 关联产品的实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`

	// 关联关系状态：1（启用）、2（不启用）、3（已失效）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// CKafka消息主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type ApmField struct {
	// 指标名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 指标数值
	Value *float64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 指标所对应的单位
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 同比结果数组，推荐使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareVals []*APMKVItem `json:"CompareVals,omitnil,omitempty" name:"CompareVals"`

	// 同比上一个周期的具体指标数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastPeriodValue []*APMKV `json:"LastPeriodValue,omitnil,omitempty" name:"LastPeriodValue"`

	// 同比指标值，已弃用，不建议使用
	CompareVal *string `json:"CompareVal,omitnil,omitempty" name:"CompareVal"`

	// 指标中文名
	NameCN *string `json:"NameCN,omitnil,omitempty" name:"NameCN"`

	// 指标英文名
	NameEN *string `json:"NameEN,omitnil,omitempty" name:"NameEN"`
}

type ApmInstanceDetail struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 业务系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 业务系统描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 业务系统状态。{
	// 1: 初始化中; 2: 运行中; 4: 限流}
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 业务系统所属地域
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 业务系统 Tag 列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// AppID 信息
	AppId *int64 `json:"AppId,omitnil,omitempty" name:"AppId"`

	// 创建人 Uin
	CreateUin *string `json:"CreateUin,omitnil,omitempty" name:"CreateUin"`

	// 存储使用量(单位：MB)
	AmountOfUsedStorage *float64 `json:"AmountOfUsedStorage,omitnil,omitempty" name:"AmountOfUsedStorage"`

	// 该业务系统服务端应用数量
	ServiceCount *int64 `json:"ServiceCount,omitnil,omitempty" name:"ServiceCount"`

	// 日均上报 Span 数
	CountOfReportSpanPerDay *int64 `json:"CountOfReportSpanPerDay,omitnil,omitempty" name:"CountOfReportSpanPerDay"`

	// Trace 数据保存时长（单位：天）
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// 业务系统上报额度
	SpanDailyCounters *int64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// 业务系统是否已开通计费（0=未开通，1=已开通）
	BillingInstance *int64 `json:"BillingInstance,omitnil,omitempty" name:"BillingInstance"`

	// 错误警示线（单位：%）
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// 采样率（单位：%）
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 是否开启错误采样（0=关, 1=开）
	ErrorSample *int64 `json:"ErrorSample,omitnil,omitempty" name:"ErrorSample"`

	// 采样慢调用保存阈值（单位：ms）
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil,omitempty" name:"SlowRequestSavedThreshold"`

	// CLS 日志所在地域
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// 日志源
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 日志功能开关（0=关， 1=开）
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// 日志主题 ID
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// 该业务系统客户端应用数量
	ClientCount *int64 `json:"ClientCount,omitnil,omitempty" name:"ClientCount"`

	// 该业务系统最近2天活跃应用数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// CLS 日志集
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// Metric 数据保存时长（单位：天）
	MetricDuration *int64 `json:"MetricDuration,omitnil,omitempty" name:"MetricDuration"`

	// 用户自定义展示标签列表
	CustomShowTags []*string `json:"CustomShowTags,omitnil,omitempty" name:"CustomShowTags"`

	// 业务系统计费模式（1为预付费，0为按量付费）
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 业务系统计费模式是否生效
	PayModeEffective *bool `json:"PayModeEffective,omitnil,omitempty" name:"PayModeEffective"`

	// 响应时间警示线（单位：ms）
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// 是否免费（0=否，1=限额免费，2=完全免费），默认0
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`

	// 是否 TSF 默认业务系统（0=否，1=是）
	DefaultTSF *int64 `json:"DefaultTSF,omitnil,omitempty" name:"DefaultTSF"`

	// 是否关联 Dashboard（0=关, 1=开）
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// 关联的 Dashboard ID
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// 是否开启组件漏洞检测（0=关， 1=开）
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// 是否开启 SQL 注入分析（0=关， 1=开）
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// 限流原因。{
	// 1: 正式版限额;
	// 2: 试用版限额;
	// 4: 试用版到期;
	// 8: 账号欠费
	// }
	StopReason *int64 `json:"StopReason,omitnil,omitempty" name:"StopReason"`

	// 是否开远程命令执行检测（0=关， 1=开）
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// 是否开内存马执行检测（0=关， 1=开）
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// CLS索引类型(0=全文索引，1=键值索引)
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// traceId的索引key: 当CLS索引类型为键值索引时生效
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// 是否开启删除任意文件检测（0-关闭，1-开启）
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// 是否开启读取任意文件检测（0-关闭，1-开启）
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// 是否开启上传任意文件检测（0-关闭，1-开启）
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// 是否开启包含任意文件检测（0-关闭，1-开启）
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// 是否开启目录遍历检测（0-关闭，1-开启）
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// 是否开启模板引擎注入检测（0-关闭，1-开启）
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// 是否开启脚本引擎注入检测（0-关闭，1-开启）
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// 是否开启表达式注入检测（0-关闭，1-开启）
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// 是否开启JNDI注入检测（0-关闭，1-开启）
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// 是否开启JNI注入检测（0-关闭，1-开启）
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// 是否开启Webshell后门检测（0-关闭，1-开启）
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// 是否开启反序列化检测（0-关闭，1-开启）
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// 业务系统鉴权 token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`

	// URL长分段收敛阈值
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// URL数字分段收敛阈值
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`
}

type ApmMetricRecord struct {
	// field数组，用于指标的查询结果
	Fields []*ApmField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// tag数组，用于区分 Groupby 的对象
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ApmPrometheusRules struct {
	// 指标匹配规则ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 指标匹配规则名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则生效的应用。生效于全部应用就传空字符串
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 指标匹配规则状态：1(启用)、2（不启用）
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 指标匹配规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`

	// 匹配类型：0精准匹配，1前缀匹配，2后缀匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`
}

type ApmSampleConfig struct {
	// 实例ID
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 采样名字
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// 接口名
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 采样的span数
	SpanNum *int64 `json:"SpanNum,omitnil,omitempty" name:"SpanNum"`

	// 采样配置开关 0 关 1 开
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// tags数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 采样率
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 0=精确匹配（默认）；1=前缀匹配；2=后缀匹配
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`

	// 配置Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type ApmServiceMetric struct {
	// filed数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fields []*ApmField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// tag数组
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 应用信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDetail *ServiceDetail `json:"ServiceDetail,omitnil,omitempty" name:"ServiceDetail"`
}

type ApmTag struct {
	// 维度Key(列名，标签Key)
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 维度值（标签值）
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type CreateApmInstanceRequestParams struct {
	// 业务系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 业务系统描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Trace 数据保存时长（单位：天，默认存储时长为3天）
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// 业务系统 Tag 列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 业务系统上报额度值，默认赋值为0表示不限制上报额度，已废弃
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// 业务系统的计费模式（0=按量付费，1=预付费）
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否为免费版业务系统（0=付费版；1=TSF 受限免费版；2=免费版）
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`
}

type CreateApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 业务系统描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Trace 数据保存时长（单位：天，默认存储时长为3天）
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// 业务系统 Tag 列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 业务系统上报额度值，默认赋值为0表示不限制上报额度，已废弃
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// 业务系统的计费模式（0=按量付费，1=预付费）
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 是否为免费版业务系统（0=付费版；1=TSF 受限免费版；2=免费版）
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`
}

func (r *CreateApmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "TraceDuration")
	delete(f, "Tags")
	delete(f, "SpanDailyCounters")
	delete(f, "PayMode")
	delete(f, "Free")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmInstanceResponseParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateApmInstanceResponseParams `json:"Response"`
}

func (r *CreateApmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmPrometheusRuleRequestParams struct {
	// 指标匹配规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则生效的应用。作用全部应用就传空字符串
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 指标匹配类型：0精准匹配，1前缀匹配，2后缀匹配
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// 客户定义的命中指标名规则。
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type CreateApmPrometheusRuleRequest struct {
	*tchttp.BaseRequest
	
	// 指标匹配规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则生效的应用。作用全部应用就传空字符串
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 指标匹配类型：0精准匹配，1前缀匹配，2后缀匹配
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// 客户定义的命中指标名规则。
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *CreateApmPrometheusRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmPrometheusRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "ServiceName")
	delete(f, "MetricMatchType")
	delete(f, "MetricNameRule")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApmPrometheusRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmPrometheusRuleResponseParams struct {
	// 指标匹配规则的ID
	RuleId *int64 `json:"RuleId,omitnil,omitempty" name:"RuleId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApmPrometheusRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateApmPrometheusRuleResponseParams `json:"Response"`
}

func (r *CreateApmPrometheusRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmPrometheusRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmSampleConfigRequestParams struct {
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样率
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// 采样Tags
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 接口名
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 0=精确匹配（默认）；1=前缀匹配；2=后缀匹配
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type CreateApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样率
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// 采样Tags
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 接口名
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 0=精确匹配（默认）；1=前缀匹配；2=后缀匹配
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *CreateApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleRate")
	delete(f, "ServiceName")
	delete(f, "SampleName")
	delete(f, "Tags")
	delete(f, "OperationName")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmSampleConfigResponseParams struct {
	// 采样配置参数
	ApmSampleConfig *ApmSampleConfig `json:"ApmSampleConfig,omitnil,omitempty" name:"ApmSampleConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *CreateApmSampleConfigResponseParams `json:"Response"`
}

func (r *CreateApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProfileTaskRequestParams struct {
	// 应用名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// APM业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用实例（在线）
	ServiceInstance *string `json:"ServiceInstance,omitnil,omitempty" name:"ServiceInstance"`

	// 事件类型（cpu、alloc）
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 任务持续时长(单位：毫秒)，范围限制在5~180秒
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 执行次数，范围限制在1~100次
	AllTimes *int64 `json:"AllTimes,omitnil,omitempty" name:"AllTimes"`

	// 开始时间戳，0代表从当前开始(单位：秒)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务执行间隔(单位：毫秒)，范围限制在10~600秒，不可小于1.5倍的Duration
	TaskInterval *int64 `json:"TaskInterval,omitnil,omitempty" name:"TaskInterval"`
}

type CreateProfileTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// APM业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用实例（在线）
	ServiceInstance *string `json:"ServiceInstance,omitnil,omitempty" name:"ServiceInstance"`

	// 事件类型（cpu、alloc）
	Event *string `json:"Event,omitnil,omitempty" name:"Event"`

	// 任务持续时长(单位：毫秒)，范围限制在5~180秒
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// 执行次数，范围限制在1~100次
	AllTimes *int64 `json:"AllTimes,omitnil,omitempty" name:"AllTimes"`

	// 开始时间戳，0代表从当前开始(单位：秒)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 任务执行间隔(单位：毫秒)，范围限制在10~600秒，不可小于1.5倍的Duration
	TaskInterval *int64 `json:"TaskInterval,omitnil,omitempty" name:"TaskInterval"`
}

func (r *CreateProfileTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProfileTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
	delete(f, "InstanceId")
	delete(f, "ServiceInstance")
	delete(f, "Event")
	delete(f, "Duration")
	delete(f, "AllTimes")
	delete(f, "StartTime")
	delete(f, "TaskInterval")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateProfileTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateProfileTaskResponseParams struct {
	// 任务ID
	TaskId *int64 `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateProfileTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateProfileTaskResponseParams `json:"Response"`
}

func (r *CreateProfileTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateProfileTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApmSampleConfigRequestParams struct {
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

type DeleteApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

func (r *DeleteApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteApmSampleConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *DeleteApmSampleConfigResponseParams `json:"Response"`
}

func (r *DeleteApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAgentRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 接入方式，现支持 skywalking, ot, ebpf 方式接入上报，不填默认为 ot
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// 上报环境，现支持 pl (内网上报), public (外网), inner (自研 VPC )环境上报，不传默认为 public
	NetworkMode *string `json:"NetworkMode,omitnil,omitempty" name:"NetworkMode"`

	// 语言，现支持 java, golang, php, python, dotNet, nodejs 语言上报，不传默认为 golang
	LanguageEnvironment *string `json:"LanguageEnvironment,omitnil,omitempty" name:"LanguageEnvironment"`

	// 上报方式，已弃用
	ReportMethod *string `json:"ReportMethod,omitnil,omitempty" name:"ReportMethod"`
}

type DescribeApmAgentRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 接入方式，现支持 skywalking, ot, ebpf 方式接入上报，不填默认为 ot
	AgentType *string `json:"AgentType,omitnil,omitempty" name:"AgentType"`

	// 上报环境，现支持 pl (内网上报), public (外网), inner (自研 VPC )环境上报，不传默认为 public
	NetworkMode *string `json:"NetworkMode,omitnil,omitempty" name:"NetworkMode"`

	// 语言，现支持 java, golang, php, python, dotNet, nodejs 语言上报，不传默认为 golang
	LanguageEnvironment *string `json:"LanguageEnvironment,omitnil,omitempty" name:"LanguageEnvironment"`

	// 上报方式，已弃用
	ReportMethod *string `json:"ReportMethod,omitnil,omitempty" name:"ReportMethod"`
}

func (r *DescribeApmAgentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "AgentType")
	delete(f, "NetworkMode")
	delete(f, "LanguageEnvironment")
	delete(f, "ReportMethod")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmAgentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAgentResponseParams struct {
	// Agent 信息
	ApmAgent *ApmAgentInfo `json:"ApmAgent,omitnil,omitempty" name:"ApmAgent"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmAgentResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmAgentResponseParams `json:"Response"`
}

func (r *DescribeApmAgentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmApplicationConfigRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type DescribeApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

func (r *DescribeApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmApplicationConfigResponseParams struct {
	// Apm应用配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApmAppConfig *ApmAppConfig `json:"ApmAppConfig,omitnil,omitempty" name:"ApmAppConfig"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmApplicationConfigResponseParams `json:"Response"`
}

func (r *DescribeApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAssociationRequestParams struct {
	// 关联的产品名，当前只支持Prometheus
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 业务系统名
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeApmAssociationRequest struct {
	*tchttp.BaseRequest
	
	// 关联的产品名，当前只支持Prometheus
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 业务系统名
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeApmAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAssociationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmAssociationResponseParams struct {
	// 关联的产品实例ID
	ApmAssociation *ApmAssociation `json:"ApmAssociation,omitnil,omitempty" name:"ApmAssociation"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmAssociationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmAssociationResponseParams `json:"Response"`
}

func (r *DescribeApmAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmInstancesRequestParams struct {
	// Tag 列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 按业务系统名过滤，支持模糊检索
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 按业务系统 ID 过滤，支持模糊检索
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按业务系统 ID 过滤
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否查询官方 Demo 业务系统（0=非 Demo 业务系统，1=Demo 业务系统，默认为0）
	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitnil,omitempty" name:"DemoInstanceFlag"`

	// 是否查询全地域业务系统（0=不查询全地域，1=查询全地域，默认为0）
	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitnil,omitempty" name:"AllRegionsFlag"`
}

type DescribeApmInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Tag 列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 按业务系统名过滤，支持模糊检索
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 按业务系统 ID 过滤，支持模糊检索
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 按业务系统 ID 过滤
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 是否查询官方 Demo 业务系统（0=非 Demo 业务系统，1=Demo 业务系统，默认为0）
	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitnil,omitempty" name:"DemoInstanceFlag"`

	// 是否查询全地域业务系统（0=不查询全地域，1=查询全地域，默认为0）
	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitnil,omitempty" name:"AllRegionsFlag"`
}

func (r *DescribeApmInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Tags")
	delete(f, "InstanceName")
	delete(f, "InstanceId")
	delete(f, "InstanceIds")
	delete(f, "DemoInstanceFlag")
	delete(f, "AllRegionsFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmInstancesResponseParams struct {
	// APM 业务系统列表
	Instances []*ApmInstanceDetail `json:"Instances,omitnil,omitempty" name:"Instances"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmInstancesResponseParams `json:"Response"`
}

func (r *DescribeApmInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmPrometheusRuleRequestParams struct {
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeApmPrometheusRuleRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeApmPrometheusRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmPrometheusRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmPrometheusRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmPrometheusRuleResponseParams struct {
	// 指标匹配规则
	ApmPrometheusRules []*ApmPrometheusRules `json:"ApmPrometheusRules,omitnil,omitempty" name:"ApmPrometheusRules"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmPrometheusRuleResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmPrometheusRuleResponseParams `json:"Response"`
}

func (r *DescribeApmPrometheusRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmPrometheusRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmSampleConfigRequestParams struct {
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

type DescribeApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`
}

func (r *DescribeApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmSampleConfigResponseParams struct {
	// 采样配置列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApmSampleConfigs []*ApmSampleConfig `json:"ApmSampleConfigs,omitnil,omitempty" name:"ApmSampleConfigs"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmSampleConfigResponseParams `json:"Response"`
}

func (r *DescribeApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmServiceMetricRequestParams struct {
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 应用ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 是否demo模式
	Demo *bool `json:"Demo,omitnil,omitempty" name:"Demo"`

	// 应用状态筛选，可枚举的值为：health、warning、error。如果选中多个状态用逗号隔开，例如："warning,error"
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// 标签列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeApmServiceMetricRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 应用ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 开始时间
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 是否demo模式
	Demo *bool `json:"Demo,omitnil,omitempty" name:"Demo"`

	// 应用状态筛选，可枚举的值为：health、warning、error。如果选中多个状态用逗号隔开，例如："warning,error"
	ServiceStatus *string `json:"ServiceStatus,omitnil,omitempty" name:"ServiceStatus"`

	// 标签列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeApmServiceMetricRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmServiceMetricRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	delete(f, "ServiceID")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "OrderBy")
	delete(f, "Demo")
	delete(f, "ServiceStatus")
	delete(f, "Tags")
	delete(f, "Page")
	delete(f, "PageSize")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApmServiceMetricRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApmServiceMetricResponseParams struct {
	// 应用指标列表
	ServiceMetricList []*ApmServiceMetric `json:"ServiceMetricList,omitnil,omitempty" name:"ServiceMetricList"`

	// 符合筛选条件的应用数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 警示异常应用数
	WarningErrorCount *int64 `json:"WarningErrorCount,omitnil,omitempty" name:"WarningErrorCount"`

	// 应用总数
	ApplicationCount *int64 `json:"ApplicationCount,omitnil,omitempty" name:"ApplicationCount"`

	// 页码
	Page *int64 `json:"Page,omitnil,omitempty" name:"Page"`

	// 页大小
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`

	// 异常应用数
	ErrorCount *int64 `json:"ErrorCount,omitnil,omitempty" name:"ErrorCount"`

	// 警示应用数
	WarningCount *int64 `json:"WarningCount,omitnil,omitempty" name:"WarningCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeApmServiceMetricResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApmServiceMetricResponseParams `json:"Response"`
}

func (r *DescribeApmServiceMetricResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApmServiceMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralApmApplicationConfigRequestParams struct {
	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeGeneralApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeGeneralApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ServiceName")
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralApmApplicationConfigResponseParams struct {
	// 应用配置项
	ApmApplicationConfigView *ApmApplicationConfigView `json:"ApmApplicationConfigView,omitnil,omitempty" name:"ApmApplicationConfigView"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralApmApplicationConfigResponseParams `json:"Response"`
}

func (r *DescribeGeneralApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralMetricDataRequestParams struct {
	// 需要查询的指标名称，不可自定义输入，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 视图名称，不可自定义输入。[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// 要过滤的维度信息，不同视图有对应的指标维度，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	Filters []*GeneralFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 聚合维度，不同视图有对应的指标维度，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 起始时间的时间戳，支持查询30天内的指标数据。（单位：秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间的时间戳，支持查询30天内的指标数据。（单位：秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否按固定时间跨度聚合，填入1及大于1的值按1处理，不填按0处理。
	// - 填入0，则计算开始时间到截止时间的指标数据。
	// - 填入1，则会按照开始时间到截止时间的时间跨度选择聚合粒度：
	//  - 时间跨度 (0,12) 小时，则按一分钟粒度聚合。
	//  - 时间跨度 [12,48] 小时，则按五分钟粒度聚合。
	//  - 时间跨度 (48, +∞) 小时，则按一小时粒度聚合。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 对查询指标进行排序：
	// Key 填写云 API 指标名称，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	// Value 填写排序方式：     
	// - asc：对查询指标进行升序排序
	// - desc：对查询指标进行降序排序
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 查询指标的限制条数，目前最多展示50条数据，PageSize取值为1-50，上送PageSize则根据PageSize的值展示限制条数。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeGeneralMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 需要查询的指标名称，不可自定义输入，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	Metrics []*string `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 视图名称，不可自定义输入。[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	ViewName *string `json:"ViewName,omitnil,omitempty" name:"ViewName"`

	// 要过滤的维度信息，不同视图有对应的指标维度，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	Filters []*GeneralFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 聚合维度，不同视图有对应的指标维度，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 起始时间的时间戳，支持查询30天内的指标数据。（单位：秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间的时间戳，支持查询30天内的指标数据。（单位：秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 是否按固定时间跨度聚合，填入1及大于1的值按1处理，不填按0处理。
	// - 填入0，则计算开始时间到截止时间的指标数据。
	// - 填入1，则会按照开始时间到截止时间的时间跨度选择聚合粒度：
	//  - 时间跨度 (0,12) 小时，则按一分钟粒度聚合。
	//  - 时间跨度 [12,48] 小时，则按五分钟粒度聚合。
	//  - 时间跨度 (48, +∞) 小时，则按一小时粒度聚合。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 对查询指标进行排序：
	// Key 填写云 API 指标名称，[详情请见。](https://cloud.tencent.com/document/product/248/101681)
	// Value 填写排序方式：     
	// - asc：对查询指标进行升序排序
	// - desc：对查询指标进行降序排序
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 查询指标的限制条数，目前最多展示50条数据，PageSize取值为1-50，上送PageSize则根据PageSize的值展示限制条数。
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeGeneralMetricDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralMetricDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Metrics")
	delete(f, "InstanceId")
	delete(f, "ViewName")
	delete(f, "Filters")
	delete(f, "GroupBy")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Period")
	delete(f, "OrderBy")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralMetricDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralMetricDataResponseParams struct {
	// 指标结果集
	Records []*Line `json:"Records,omitnil,omitempty" name:"Records"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralMetricDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralMetricDataResponseParams `json:"Response"`
}

func (r *DescribeGeneralMetricDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralOTSpanListRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span 查询开始时间戳（单位：秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span 查询结束时间戳（单位：秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通用过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序
	// 现支持的 Key 有：
	// 
	// - startTime(开始时间)
	// - endTime(结束时间)
	// - duration(响应时间)
	// 
	// 现支持的 Value 有：
	// 
	// - desc(降序排序)
	// - asc(升序排序)
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 业务自身服务名，控制台用户请填写taw
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// 单页项目个数，默认为10000，合法取值范围为0～10000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeGeneralOTSpanListRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span 查询开始时间戳（单位：秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span 查询结束时间戳（单位：秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通用过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序
	// 现支持的 Key 有：
	// 
	// - startTime(开始时间)
	// - endTime(结束时间)
	// - duration(响应时间)
	// 
	// 现支持的 Value 有：
	// 
	// - desc(降序排序)
	// - asc(升序排序)
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 业务自身服务名，控制台用户请填写taw
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// 单页项目个数，默认为10000，合法取值范围为0～10000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeGeneralOTSpanListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralOTSpanListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "BusinessName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralOTSpanListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralOTSpanListResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Spans字段中包含了链路数据的全部内容，由于数据经过了压缩，需要对结果进行如下三步转换，以还原始的文本。
	// 1. 将Spans字段中的文本进行 Base64 解码，得到经过压缩后字节数组。
	// 2. 使用 gzip 对压缩后的字节数组进行解压，得到压缩前的字节数组。
	// 3. 使用 UTF-8 字符集，将压缩前的字节数组转换为文本。
	Spans *string `json:"Spans,omitnil,omitempty" name:"Spans"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralOTSpanListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralOTSpanListResponseParams `json:"Response"`
}

func (r *DescribeGeneralOTSpanListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralOTSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralSpanListRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span 查询开始时间戳（单位：秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span 查询结束时间戳（单位：秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通用过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序
	// 现支持的 Key 有：
	// 
	// - startTime(开始时间)
	// - endTime(结束时间)
	// - duration(响应时间)
	// 
	// 现支持的 Value 有：
	// 
	// - desc(降序排序)
	// - asc(升序排序)
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 业务自身服务名，控制台用户请填写taw
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// 单页项目个数，默认为10000，合法取值范围为0～10000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeGeneralSpanListRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// Span 查询开始时间戳（单位：秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// Span 查询结束时间戳（单位：秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 通用过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序
	// 现支持的 Key 有：
	// 
	// - startTime(开始时间)
	// - endTime(结束时间)
	// - duration(响应时间)
	// 
	// 现支持的 Value 有：
	// 
	// - desc(降序排序)
	// - asc(升序排序)
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 业务自身服务名，控制台用户请填写taw
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// 单页项目个数，默认为10000，合法取值范围为0～10000
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeGeneralSpanListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralSpanListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "BusinessName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralSpanListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralSpanListResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// Span 分页列表
	Spans []*Span `json:"Spans,omitnil,omitempty" name:"Spans"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGeneralSpanListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGeneralSpanListResponseParams `json:"Response"`
}

func (r *DescribeGeneralSpanListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGeneralSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricRecordsRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 开始时间（单位为秒）
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（单位为秒）
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or 过滤条件
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// 排序
	// 现支持的 Key 有：
	// 
	// - startTime(开始时间)
	// - endTime(结束时间)
	// - duration(响应时间)
	// 
	// 现支持的 Value 有：
	// 
	// - desc(降序排序)
	// - asc(升序排序)
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 业务名称，控制台用户请填写taw。
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// 特殊处理查询结果
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 每页大小，默认为1000，合法取值范围为0~1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页长
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

type DescribeMetricRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 开始时间（单位为秒）
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（单位为秒）
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or 过滤条件
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// 排序
	// 现支持的 Key 有：
	// 
	// - startTime(开始时间)
	// - endTime(结束时间)
	// - duration(响应时间)
	// 
	// 现支持的 Value 有：
	// 
	// - desc(降序排序)
	// - asc(升序排序)
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 业务名称，控制台用户请填写taw。
	BusinessName *string `json:"BusinessName,omitnil,omitempty" name:"BusinessName"`

	// 特殊处理查询结果
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 每页大小，默认为1000，合法取值范围为0~1000
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil,omitempty" name:"PageIndex"`

	// 页长
	PageSize *int64 `json:"PageSize,omitnil,omitempty" name:"PageSize"`
}

func (r *DescribeMetricRecordsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricRecordsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Metrics")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "GroupBy")
	delete(f, "Filters")
	delete(f, "OrFilters")
	delete(f, "OrderBy")
	delete(f, "BusinessName")
	delete(f, "Type")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMetricRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricRecordsResponseParams struct {
	// 指标结果集
	Records []*ApmMetricRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 查询指标结果集条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeMetricRecordsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeMetricRecordsResponseParams `json:"Response"`
}

func (r *DescribeMetricRecordsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMetricRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceOverviewRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 开始时间（单位：秒）
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（单位：秒）
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式
	// Value 填写：
	// - asc：对查询指标进行升序排序
	// - desc：对查询指标进行降序排序
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 每页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeServiceOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 开始时间（单位：秒）
	StartTime *uint64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（单位：秒）
	EndTime *uint64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil,omitempty" name:"GroupBy"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 排序方式
	// Value 填写：
	// - asc：对查询指标进行升序排序
	// - desc：对查询指标进行降序排序
	OrderBy *OrderBy `json:"OrderBy,omitnil,omitempty" name:"OrderBy"`

	// 每页大小
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeServiceOverviewRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceOverviewRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Metrics")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "GroupBy")
	delete(f, "Filters")
	delete(f, "OrderBy")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceOverviewResponseParams struct {
	// 指标结果集
	Records []*ApmMetricRecord `json:"Records,omitnil,omitempty" name:"Records"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeServiceOverviewResponse struct {
	*tchttp.BaseResponse
	Response *DescribeServiceOverviewResponseParams `json:"Response"`
}

func (r *DescribeServiceOverviewResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeServiceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维度名
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 开始时间（单位为秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（单位为秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or 过滤条件
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// 使用类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 维度名
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 开始时间（单位为秒）
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 结束时间（单位为秒）
	EndTime *int64 `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// Or 过滤条件
	OrFilters []*Filter `json:"OrFilters,omitnil,omitempty" name:"OrFilters"`

	// 使用类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

func (r *DescribeTagValuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "TagKey")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Filters")
	delete(f, "OrFilters")
	delete(f, "Type")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTagValuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTagValuesResponseParams struct {
	// 维度值列表
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTagValuesResponseParams `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 过滤方式（=, !=, in）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 过滤维度名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 过滤值，in过滤方式用逗号分割多个值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type GeneralFilter struct {
	// 过滤维度名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 过滤值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Instrument struct {
	// 组件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 组件开关
	Enable *bool `json:"Enable,omitnil,omitempty" name:"Enable"`
}

type Line struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 指标中文名
	MetricNameCN *string `json:"MetricNameCN,omitnil,omitempty" name:"MetricNameCN"`

	// 时间序列
	TimeSerial []*int64 `json:"TimeSerial,omitnil,omitempty" name:"TimeSerial"`

	// 数据序列
	DataSerial []*float64 `json:"DataSerial,omitnil,omitempty" name:"DataSerial"`

	// 维度列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 指标数据单位
	MetricUnit *string `json:"MetricUnit,omitnil,omitempty" name:"MetricUnit"`
}

// Predefined struct for user
type ModifyApmApplicationConfigRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// URL收敛开关,0 关 | 1 开
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// URL收敛阈值
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// 异常过滤正则规则，逗号分隔
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// URL收敛正则规则，逗号分隔
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// 错误码过滤，逗号分隔
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// URL排除正则规则，逗号分隔
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// 日志开关 0 关 1 开
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// 日志地域
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// 日志主题ID
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// CLS 日志集 | ES 集群ID
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 日志来源 CLS | ES
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 需过滤的接口
	IgnoreOperationName *string `json:"IgnoreOperationName,omitnil,omitempty" name:"IgnoreOperationName"`

	// 是否开启线程剖析
	EnableSnapshot *bool `json:"EnableSnapshot,omitnil,omitempty" name:"EnableSnapshot"`

	// 线程剖析超时阈值
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// 是否开启agent
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// 是否开启链路压缩
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// 是否开启应用诊断的开关
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// 组件列表
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// 探针接口相关配置
	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitnil,omitempty" name:"AgentOperationConfigView"`

	// 是否开启应用日志配置
	EnableLogConfig *bool `json:"EnableLogConfig,omitnil,omitempty" name:"EnableLogConfig"`

	// 应用是否开启dashboard配置： false 关（与业务系统保持一致）/true 开（应用级配置）
	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitnil,omitempty" name:"EnableDashboardConfig"`

	// 是否关联dashboard： 0 关 1 开
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// dashboard ID
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// CLS索引类型(0=全文索引，1=键值索引)
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// traceId的索引key: 当CLS索引类型为键值索引时生效
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// 是否开启应用安全配置
	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitnil,omitempty" name:"EnableSecurityConfig"`

	// 是否开启SQL注入分析
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// 是否开启组件漏洞检测
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// 是否开启远程命令检测
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// 是否开启内存马检测
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// 是否开启删除任意文件检测（0-关闭，1-开启）
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// 是否开启读取任意文件检测（0-关闭，1-开启）
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// 是否开启上传任意文件检测（0-关闭，1-开启）
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// 是否开启包含任意文件检测（0-关闭，1-开启）
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// 是否开启目录遍历检测（0-关闭，1-开启）
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// 是否开启模板引擎注入检测（0-关闭，1-开启）
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// 是否开启脚本引擎注入检测（0-关闭，1-开启）
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// 是否开启表达式注入检测（0-关闭，1-开启）
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// 是否开启JNDI注入检测（0-关闭，1-开启）
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// 是否开启JNI注入检测（0-关闭，1-开启）
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// 是否开启Webshell后门检测（0-关闭，1-开启）
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// 是否开启反序列化检测（0-关闭，1-开启）
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// 接口自动收敛开关,0 关 | 1 开
	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitnil,omitempty" name:"UrlAutoConvergenceEnable"`

	// URL长分段收敛阈值
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// URL数字分段收敛阈值
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// 探针熔断内存阈值
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// 探针熔断CPU阈值
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// 是否开启SQL参数获取
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// 慢SQL阈值
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`
}

type ModifyApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 应用名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// URL收敛开关,0 关 | 1 开
	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitnil,omitempty" name:"UrlConvergenceSwitch"`

	// URL收敛阈值
	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitnil,omitempty" name:"UrlConvergenceThreshold"`

	// 异常过滤正则规则，逗号分隔
	ExceptionFilter *string `json:"ExceptionFilter,omitnil,omitempty" name:"ExceptionFilter"`

	// URL收敛正则规则，逗号分隔
	UrlConvergence *string `json:"UrlConvergence,omitnil,omitempty" name:"UrlConvergence"`

	// 错误码过滤，逗号分隔
	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitnil,omitempty" name:"ErrorCodeFilter"`

	// URL排除正则规则，逗号分隔
	UrlExclude *string `json:"UrlExclude,omitnil,omitempty" name:"UrlExclude"`

	// 日志开关 0 关 1 开
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// 日志地域
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// 日志主题ID
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// CLS 日志集 | ES 集群ID
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 日志来源 CLS | ES
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 需过滤的接口
	IgnoreOperationName *string `json:"IgnoreOperationName,omitnil,omitempty" name:"IgnoreOperationName"`

	// 是否开启线程剖析
	EnableSnapshot *bool `json:"EnableSnapshot,omitnil,omitempty" name:"EnableSnapshot"`

	// 线程剖析超时阈值
	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitnil,omitempty" name:"SnapshotTimeout"`

	// 是否开启agent
	AgentEnable *bool `json:"AgentEnable,omitnil,omitempty" name:"AgentEnable"`

	// 是否开启链路压缩
	TraceSquash *bool `json:"TraceSquash,omitnil,omitempty" name:"TraceSquash"`

	// 是否开启应用诊断的开关
	EventEnable *bool `json:"EventEnable,omitnil,omitempty" name:"EventEnable"`

	// 组件列表
	InstrumentList []*Instrument `json:"InstrumentList,omitnil,omitempty" name:"InstrumentList"`

	// 探针接口相关配置
	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitnil,omitempty" name:"AgentOperationConfigView"`

	// 是否开启应用日志配置
	EnableLogConfig *bool `json:"EnableLogConfig,omitnil,omitempty" name:"EnableLogConfig"`

	// 应用是否开启dashboard配置： false 关（与业务系统保持一致）/true 开（应用级配置）
	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitnil,omitempty" name:"EnableDashboardConfig"`

	// 是否关联dashboard： 0 关 1 开
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// dashboard ID
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// CLS索引类型(0=全文索引，1=键值索引)
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// traceId的索引key: 当CLS索引类型为键值索引时生效
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// 是否开启应用安全配置
	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitnil,omitempty" name:"EnableSecurityConfig"`

	// 是否开启SQL注入分析
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// 是否开启组件漏洞检测
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// 是否开启远程命令检测
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// 是否开启内存马检测
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// 是否开启删除任意文件检测（0-关闭，1-开启）
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// 是否开启读取任意文件检测（0-关闭，1-开启）
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// 是否开启上传任意文件检测（0-关闭，1-开启）
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// 是否开启包含任意文件检测（0-关闭，1-开启）
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// 是否开启目录遍历检测（0-关闭，1-开启）
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// 是否开启模板引擎注入检测（0-关闭，1-开启）
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// 是否开启脚本引擎注入检测（0-关闭，1-开启）
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// 是否开启表达式注入检测（0-关闭，1-开启）
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// 是否开启JNDI注入检测（0-关闭，1-开启）
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// 是否开启JNI注入检测（0-关闭，1-开启）
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// 是否开启Webshell后门检测（0-关闭，1-开启）
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// 是否开启反序列化检测（0-关闭，1-开启）
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// 接口自动收敛开关,0 关 | 1 开
	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitnil,omitempty" name:"UrlAutoConvergenceEnable"`

	// URL长分段收敛阈值
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// URL数字分段收敛阈值
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`

	// 探针熔断内存阈值
	DisableMemoryUsed *int64 `json:"DisableMemoryUsed,omitnil,omitempty" name:"DisableMemoryUsed"`

	// 探针熔断CPU阈值
	DisableCpuUsed *int64 `json:"DisableCpuUsed,omitnil,omitempty" name:"DisableCpuUsed"`

	// 是否开启SQL参数获取
	DbStatementParametersEnabled *bool `json:"DbStatementParametersEnabled,omitnil,omitempty" name:"DbStatementParametersEnabled"`

	// 慢SQL阈值
	SlowSQLThresholds []*ApmTag `json:"SlowSQLThresholds,omitnil,omitempty" name:"SlowSQLThresholds"`
}

func (r *ModifyApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ServiceName")
	delete(f, "UrlConvergenceSwitch")
	delete(f, "UrlConvergenceThreshold")
	delete(f, "ExceptionFilter")
	delete(f, "UrlConvergence")
	delete(f, "ErrorCodeFilter")
	delete(f, "UrlExclude")
	delete(f, "IsRelatedLog")
	delete(f, "LogRegion")
	delete(f, "LogTopicID")
	delete(f, "LogSet")
	delete(f, "LogSource")
	delete(f, "IgnoreOperationName")
	delete(f, "EnableSnapshot")
	delete(f, "SnapshotTimeout")
	delete(f, "AgentEnable")
	delete(f, "TraceSquash")
	delete(f, "EventEnable")
	delete(f, "InstrumentList")
	delete(f, "AgentOperationConfigView")
	delete(f, "EnableLogConfig")
	delete(f, "EnableDashboardConfig")
	delete(f, "IsRelatedDashboard")
	delete(f, "DashboardTopicID")
	delete(f, "LogIndexType")
	delete(f, "LogTraceIdKey")
	delete(f, "EnableSecurityConfig")
	delete(f, "IsSqlInjectionAnalysis")
	delete(f, "IsInstrumentationVulnerabilityScan")
	delete(f, "IsRemoteCommandExecutionAnalysis")
	delete(f, "IsMemoryHijackingAnalysis")
	delete(f, "IsDeleteAnyFileAnalysis")
	delete(f, "IsReadAnyFileAnalysis")
	delete(f, "IsUploadAnyFileAnalysis")
	delete(f, "IsIncludeAnyFileAnalysis")
	delete(f, "IsDirectoryTraversalAnalysis")
	delete(f, "IsTemplateEngineInjectionAnalysis")
	delete(f, "IsScriptEngineInjectionAnalysis")
	delete(f, "IsExpressionInjectionAnalysis")
	delete(f, "IsJNDIInjectionAnalysis")
	delete(f, "IsJNIInjectionAnalysis")
	delete(f, "IsWebshellBackdoorAnalysis")
	delete(f, "IsDeserializationAnalysis")
	delete(f, "UrlAutoConvergenceEnable")
	delete(f, "UrlLongSegmentThreshold")
	delete(f, "UrlNumberSegmentThreshold")
	delete(f, "DisableMemoryUsed")
	delete(f, "DisableCpuUsed")
	delete(f, "DbStatementParametersEnabled")
	delete(f, "SlowSQLThresholds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmApplicationConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmApplicationConfigResponseParams `json:"Response"`
}

func (r *ModifyApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmAssociationRequestParams struct {
	// 关联的产品名，当前只支持Prometheus、CKafka
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 关联关系的状态：// 关联关系状态：1（启用）、2（不启用）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 关联的产品实例ID
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`

	// CKafka消息主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

type ModifyApmAssociationRequest struct {
	*tchttp.BaseRequest
	
	// 关联的产品名，当前只支持Prometheus、CKafka
	ProductName *string `json:"ProductName,omitnil,omitempty" name:"ProductName"`

	// 关联关系的状态：// 关联关系状态：1（启用）、2（不启用）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 关联的产品实例ID
	PeerId *string `json:"PeerId,omitnil,omitempty" name:"PeerId"`

	// CKafka消息主题
	Topic *string `json:"Topic,omitnil,omitempty" name:"Topic"`
}

func (r *ModifyApmAssociationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmAssociationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ProductName")
	delete(f, "Status")
	delete(f, "InstanceId")
	delete(f, "PeerId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmAssociationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmAssociationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmAssociationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmAssociationResponseParams `json:"Response"`
}

func (r *ModifyApmAssociationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmAssociationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmInstanceRequestParams struct {
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 业务系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Tag 列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 业务系统描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Trace 数据保存时长（单位：天）
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// 是否开启计费
	OpenBilling *bool `json:"OpenBilling,omitnil,omitempty" name:"OpenBilling"`

	// 业务系统上报额度
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// 错误率警示线，当应用的平均错误率超出该阈值时，系统会给出异常提示。
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// 采样率（单位：%）
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 是否开启错误采样（0=关, 1=开）
	ErrorSample *int64 `json:"ErrorSample,omitnil,omitempty" name:"ErrorSample"`

	// 采样慢调用保存阈值（单位：ms）
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil,omitempty" name:"SlowRequestSavedThreshold"`

	// 是否开启日志功能（0=关, 1=开）
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// 日志地域，开启日志功能后才会生效
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// CLS 日志主题 ID，开启日志功能后才会生效
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// 日志集，开启日志功能后才会生效
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 日志源，开启日志功能后才会生效
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 用户自定义展示标签列表
	CustomShowTags []*string `json:"CustomShowTags,omitnil,omitempty" name:"CustomShowTags"`

	// 修改计费模式（1为预付费，0为按量付费）
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 响应时间警示线
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// 是否免费（0=付费版；1=TSF 受限免费版；2=免费版），默认0
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`

	// 是否关联 Dashboard（0=关,1=开）
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// 关联的 Dashboard ID，开启关联 Dashboard 后才会生效
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// 是否开启 SQL 注入检测（0=关,1=开）
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// 是否开启组件漏洞检测（0=关,1=开）
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// 是否开启远程命令攻击检测
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// 是否开启内存马检测
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// CLS索引类型(0=全文索引，1=键值索引)
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// traceId的索引key: 当CLS索引类型为键值索引时生效
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// 是否开启删除任意文件检测（0-关闭，1-开启）
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// 是否开启读取任意文件检测（0-关闭，1-开启）
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// 是否开启上传任意文件检测（0-关闭，1-开启）
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// 是否开启包含任意文件检测（0-关闭，1-开启）
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// 是否开启目录遍历检测（0-关闭，1-开启）
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// 是否开启模板引擎注入检测（0-关闭，1-开启）
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// 是否开启脚本引擎注入检测（0-关闭，1-开启）
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// 是否开启表达式注入检测（0-关闭，1-开启）
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// 是否开启JNDI注入检测（0-关闭，1-开启）
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// 是否开启JNI注入检测（0-关闭，1-开启）
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// 是否开启Webshell后门检测（0-关闭，1-开启）
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// 是否开启反序列化检测（0-关闭，1-开启）
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// URL长分段收敛阈值
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// URL数字分段收敛阈值
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`
}

type ModifyApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 业务系统名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Tag 列表
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 业务系统描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// Trace 数据保存时长（单位：天）
	TraceDuration *int64 `json:"TraceDuration,omitnil,omitempty" name:"TraceDuration"`

	// 是否开启计费
	OpenBilling *bool `json:"OpenBilling,omitnil,omitempty" name:"OpenBilling"`

	// 业务系统上报额度
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil,omitempty" name:"SpanDailyCounters"`

	// 错误率警示线，当应用的平均错误率超出该阈值时，系统会给出异常提示。
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil,omitempty" name:"ErrRateThreshold"`

	// 采样率（单位：%）
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 是否开启错误采样（0=关, 1=开）
	ErrorSample *int64 `json:"ErrorSample,omitnil,omitempty" name:"ErrorSample"`

	// 采样慢调用保存阈值（单位：ms）
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil,omitempty" name:"SlowRequestSavedThreshold"`

	// 是否开启日志功能（0=关, 1=开）
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil,omitempty" name:"IsRelatedLog"`

	// 日志地域，开启日志功能后才会生效
	LogRegion *string `json:"LogRegion,omitnil,omitempty" name:"LogRegion"`

	// CLS 日志主题 ID，开启日志功能后才会生效
	LogTopicID *string `json:"LogTopicID,omitnil,omitempty" name:"LogTopicID"`

	// 日志集，开启日志功能后才会生效
	LogSet *string `json:"LogSet,omitnil,omitempty" name:"LogSet"`

	// 日志源，开启日志功能后才会生效
	LogSource *string `json:"LogSource,omitnil,omitempty" name:"LogSource"`

	// 用户自定义展示标签列表
	CustomShowTags []*string `json:"CustomShowTags,omitnil,omitempty" name:"CustomShowTags"`

	// 修改计费模式（1为预付费，0为按量付费）
	PayMode *int64 `json:"PayMode,omitnil,omitempty" name:"PayMode"`

	// 响应时间警示线
	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitnil,omitempty" name:"ResponseDurationWarningThreshold"`

	// 是否免费（0=付费版；1=TSF 受限免费版；2=免费版），默认0
	Free *int64 `json:"Free,omitnil,omitempty" name:"Free"`

	// 是否关联 Dashboard（0=关,1=开）
	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitnil,omitempty" name:"IsRelatedDashboard"`

	// 关联的 Dashboard ID，开启关联 Dashboard 后才会生效
	DashboardTopicID *string `json:"DashboardTopicID,omitnil,omitempty" name:"DashboardTopicID"`

	// 是否开启 SQL 注入检测（0=关,1=开）
	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitnil,omitempty" name:"IsSqlInjectionAnalysis"`

	// 是否开启组件漏洞检测（0=关,1=开）
	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitnil,omitempty" name:"IsInstrumentationVulnerabilityScan"`

	// 是否开启远程命令攻击检测
	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitnil,omitempty" name:"IsRemoteCommandExecutionAnalysis"`

	// 是否开启内存马检测
	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitnil,omitempty" name:"IsMemoryHijackingAnalysis"`

	// CLS索引类型(0=全文索引，1=键值索引)
	LogIndexType *int64 `json:"LogIndexType,omitnil,omitempty" name:"LogIndexType"`

	// traceId的索引key: 当CLS索引类型为键值索引时生效
	LogTraceIdKey *string `json:"LogTraceIdKey,omitnil,omitempty" name:"LogTraceIdKey"`

	// 是否开启删除任意文件检测（0-关闭，1-开启）
	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitnil,omitempty" name:"IsDeleteAnyFileAnalysis"`

	// 是否开启读取任意文件检测（0-关闭，1-开启）
	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitnil,omitempty" name:"IsReadAnyFileAnalysis"`

	// 是否开启上传任意文件检测（0-关闭，1-开启）
	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitnil,omitempty" name:"IsUploadAnyFileAnalysis"`

	// 是否开启包含任意文件检测（0-关闭，1-开启）
	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitnil,omitempty" name:"IsIncludeAnyFileAnalysis"`

	// 是否开启目录遍历检测（0-关闭，1-开启）
	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitnil,omitempty" name:"IsDirectoryTraversalAnalysis"`

	// 是否开启模板引擎注入检测（0-关闭，1-开启）
	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitnil,omitempty" name:"IsTemplateEngineInjectionAnalysis"`

	// 是否开启脚本引擎注入检测（0-关闭，1-开启）
	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitnil,omitempty" name:"IsScriptEngineInjectionAnalysis"`

	// 是否开启表达式注入检测（0-关闭，1-开启）
	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitnil,omitempty" name:"IsExpressionInjectionAnalysis"`

	// 是否开启JNDI注入检测（0-关闭，1-开启）
	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitnil,omitempty" name:"IsJNDIInjectionAnalysis"`

	// 是否开启JNI注入检测（0-关闭，1-开启）
	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitnil,omitempty" name:"IsJNIInjectionAnalysis"`

	// 是否开启Webshell后门检测（0-关闭，1-开启）
	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitnil,omitempty" name:"IsWebshellBackdoorAnalysis"`

	// 是否开启反序列化检测（0-关闭，1-开启）
	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitnil,omitempty" name:"IsDeserializationAnalysis"`

	// URL长分段收敛阈值
	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitnil,omitempty" name:"UrlLongSegmentThreshold"`

	// URL数字分段收敛阈值
	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitnil,omitempty" name:"UrlNumberSegmentThreshold"`
}

func (r *ModifyApmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Tags")
	delete(f, "Description")
	delete(f, "TraceDuration")
	delete(f, "OpenBilling")
	delete(f, "SpanDailyCounters")
	delete(f, "ErrRateThreshold")
	delete(f, "SampleRate")
	delete(f, "ErrorSample")
	delete(f, "SlowRequestSavedThreshold")
	delete(f, "IsRelatedLog")
	delete(f, "LogRegion")
	delete(f, "LogTopicID")
	delete(f, "LogSet")
	delete(f, "LogSource")
	delete(f, "CustomShowTags")
	delete(f, "PayMode")
	delete(f, "ResponseDurationWarningThreshold")
	delete(f, "Free")
	delete(f, "IsRelatedDashboard")
	delete(f, "DashboardTopicID")
	delete(f, "IsSqlInjectionAnalysis")
	delete(f, "IsInstrumentationVulnerabilityScan")
	delete(f, "IsRemoteCommandExecutionAnalysis")
	delete(f, "IsMemoryHijackingAnalysis")
	delete(f, "LogIndexType")
	delete(f, "LogTraceIdKey")
	delete(f, "IsDeleteAnyFileAnalysis")
	delete(f, "IsReadAnyFileAnalysis")
	delete(f, "IsUploadAnyFileAnalysis")
	delete(f, "IsIncludeAnyFileAnalysis")
	delete(f, "IsDirectoryTraversalAnalysis")
	delete(f, "IsTemplateEngineInjectionAnalysis")
	delete(f, "IsScriptEngineInjectionAnalysis")
	delete(f, "IsExpressionInjectionAnalysis")
	delete(f, "IsJNDIInjectionAnalysis")
	delete(f, "IsJNIInjectionAnalysis")
	delete(f, "IsWebshellBackdoorAnalysis")
	delete(f, "IsDeserializationAnalysis")
	delete(f, "UrlLongSegmentThreshold")
	delete(f, "UrlNumberSegmentThreshold")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmInstanceResponseParams `json:"Response"`
}

func (r *ModifyApmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmPrometheusRuleRequestParams struct {
	// 规则ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 所要修改的规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则状态：1(启用)、2（不启用）、3（删除）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则生效的应用。生效于全部应用就传空（这个如果不修改也要传旧的规则）
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 匹配类型：0精准匹配，1前缀匹配，2后缀匹配（这个如果不修改也要传旧的规则）
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// 客户定义的命中指标名规则。
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`
}

type ModifyApmPrometheusRuleRequest struct {
	*tchttp.BaseRequest
	
	// 规则ID
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 所要修改的规则名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 规则状态：1(启用)、2（不启用）、3（删除）
	Status *uint64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 规则生效的应用。生效于全部应用就传空（这个如果不修改也要传旧的规则）
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 匹配类型：0精准匹配，1前缀匹配，2后缀匹配（这个如果不修改也要传旧的规则）
	MetricMatchType *int64 `json:"MetricMatchType,omitnil,omitempty" name:"MetricMatchType"`

	// 客户定义的命中指标名规则。
	MetricNameRule *string `json:"MetricNameRule,omitnil,omitempty" name:"MetricNameRule"`
}

func (r *ModifyApmPrometheusRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmPrometheusRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Id")
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "Status")
	delete(f, "ServiceName")
	delete(f, "MetricMatchType")
	delete(f, "MetricNameRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmPrometheusRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmPrometheusRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmPrometheusRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmPrometheusRuleResponseParams `json:"Response"`
}

func (r *ModifyApmPrometheusRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmPrometheusRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmSampleConfigRequestParams struct {
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// 采样率
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 应用名，生效于所有应用则填空
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 接口名
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 采样tag
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 采样开关 0关 1开 2删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 配置Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 0=精确匹配（默认）；1=前缀匹配；2=后缀匹配
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

type ModifyApmSampleConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 采样规则名
	SampleName *string `json:"SampleName,omitnil,omitempty" name:"SampleName"`

	// 采样率
	SampleRate *int64 `json:"SampleRate,omitnil,omitempty" name:"SampleRate"`

	// 应用名，生效于所有应用则填空
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 接口名
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 采样tag
	Tags []*APMKVItem `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 采样开关 0关 1开 2删除
	Status *int64 `json:"Status,omitnil,omitempty" name:"Status"`

	// 配置Id
	Id *int64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 0=精确匹配（默认）；1=前缀匹配；2=后缀匹配
	OperationType *int64 `json:"OperationType,omitnil,omitempty" name:"OperationType"`
}

func (r *ModifyApmSampleConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmSampleConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SampleName")
	delete(f, "SampleRate")
	delete(f, "ServiceName")
	delete(f, "OperationName")
	delete(f, "Tags")
	delete(f, "Status")
	delete(f, "Id")
	delete(f, "OperationType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmSampleConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmSampleConfigResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyApmSampleConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyApmSampleConfigResponseParams `json:"Response"`
}

func (r *ModifyApmSampleConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyApmSampleConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGeneralApmApplicationConfigRequestParams struct {
	// 业务系统Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要修改的字段key value分别指定字段名、字段值
	// [具体字段请见](https://cloud.tencent.com/document/product/248/111241)
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 需要修改配置的应用列表名称	
	ServiceNames []*string `json:"ServiceNames,omitnil,omitempty" name:"ServiceNames"`
}

type ModifyGeneralApmApplicationConfigRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 需要修改的字段key value分别指定字段名、字段值
	// [具体字段请见](https://cloud.tencent.com/document/product/248/111241)
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 需要修改配置的应用列表名称	
	ServiceNames []*string `json:"ServiceNames,omitnil,omitempty" name:"ServiceNames"`
}

func (r *ModifyGeneralApmApplicationConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGeneralApmApplicationConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Tags")
	delete(f, "ServiceNames")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGeneralApmApplicationConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGeneralApmApplicationConfigResponseParams struct {
	// 返回值描述
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGeneralApmApplicationConfigResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGeneralApmApplicationConfigResponseParams `json:"Response"`
}

func (r *ModifyGeneralApmApplicationConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGeneralApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OrderBy struct {
	// 需要排序的字段，现支持 startTIme, endTime, duration
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// asc 顺序排序 / desc 倒序排序
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type QueryMetricItem struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 同比，现支持 CompareByYesterday (与昨天相比)和CompareByLastWeek (与上周相比) 
	Compares []*string `json:"Compares,omitnil,omitempty" name:"Compares"`

	// 同比，已弃用，不建议使用
	Compare *string `json:"Compare,omitnil,omitempty" name:"Compare"`
}

type ServiceDetail struct {
	// 应用ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 业务系统ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceKey *string `json:"InstanceKey,omitnil,omitempty" name:"InstanceKey"`

	// 用户appid
	// 注意：此字段可能返回 null，表示取不到有效值。
	AppID *int64 `json:"AppID,omitnil,omitempty" name:"AppID"`

	// 主账号uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUIN *string `json:"CreateUIN,omitnil,omitempty" name:"CreateUIN"`

	// 应用名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 应用描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceDescription *string `json:"ServiceDescription,omitnil,omitempty" name:"ServiceDescription"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 业务系统名称
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`
}

type Span struct {
	// Trace ID
	TraceID *string `json:"TraceID,omitnil,omitempty" name:"TraceID"`

	// 日志
	Logs []*SpanLog `json:"Logs,omitnil,omitempty" name:"Logs"`

	// 标签
	Tags []*SpanTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 上报应用服务信息
	Process *SpanProcess `json:"Process,omitnil,omitempty" name:"Process"`

	// 产生时间戳(毫秒)
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// Span 名称
	OperationName *string `json:"OperationName,omitnil,omitempty" name:"OperationName"`

	// 关联关系
	References []*SpanReference `json:"References,omitnil,omitempty" name:"References"`

	// 产生时间戳(微秒)
	StartTime *int64 `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 持续耗时(微妙)
	Duration *int64 `json:"Duration,omitnil,omitempty" name:"Duration"`

	// Span ID
	SpanID *string `json:"SpanID,omitnil,omitempty" name:"SpanID"`

	// 产生时间戳(毫秒)
	StartTimeMillis *int64 `json:"StartTimeMillis,omitnil,omitempty" name:"StartTimeMillis"`

	// Parent Span ID
	ParentSpanID *string `json:"ParentSpanID,omitnil,omitempty" name:"ParentSpanID"`
}

type SpanLog struct {
	// 日志时间戳
	Timestamp *int64 `json:"Timestamp,omitnil,omitempty" name:"Timestamp"`

	// 标签
	Fields []*SpanTag `json:"Fields,omitnil,omitempty" name:"Fields"`
}

type SpanProcess struct {
	// 应用服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Tags 标签数组
	Tags []*SpanTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type SpanReference struct {
	// 关联关系类型
	RefType *string `json:"RefType,omitnil,omitempty" name:"RefType"`

	// Span ID
	SpanID *string `json:"SpanID,omitnil,omitempty" name:"SpanID"`

	// Trace ID
	TraceID *string `json:"TraceID,omitnil,omitempty" name:"TraceID"`
}

type SpanTag struct {
	// 标签类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 标签Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type TerminateApmInstanceRequestParams struct {
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type TerminateApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *TerminateApmInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateApmInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "TerminateApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type TerminateApmInstanceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type TerminateApmInstanceResponse struct {
	*tchttp.BaseResponse
	Response *TerminateApmInstanceResponseParams `json:"Response"`
}

func (r *TerminateApmInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *TerminateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}