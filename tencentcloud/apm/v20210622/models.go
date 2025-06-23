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

	// 业务系统鉴权 token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type ApmMetricRecord struct {
	// field数组，用于指标的查询结果
	Fields []*ApmField `json:"Fields,omitnil,omitempty" name:"Fields"`

	// tag数组，用于区分 Groupby 的对象
	Tags []*ApmTag `json:"Tags,omitnil,omitempty" name:"Tags"`
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