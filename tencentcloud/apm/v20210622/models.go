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
	// Key值定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// Value值定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitnil" name:"Value"`
}

type APMKVItem struct {
	// Key值定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// Value值定义
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type ApmAgentInfo struct {
	// Agent下载地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	AgentDownloadURL *string `json:"AgentDownloadURL,omitnil" name:"AgentDownloadURL"`

	// Collector上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	CollectorURL *string `json:"CollectorURL,omitnil" name:"CollectorURL"`

	// Token信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Token *string `json:"Token,omitnil" name:"Token"`

	// 外网上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicCollectorURL *string `json:"PublicCollectorURL,omitnil" name:"PublicCollectorURL"`

	// 自研VPC上报地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	InnerCollectorURL *string `json:"InnerCollectorURL,omitnil" name:"InnerCollectorURL"`

	// 内网上报地址(Private Link上报地址)
	// 注意：此字段可能返回 null，表示取不到有效值。
	PrivateLinkCollectorURL *string `json:"PrivateLinkCollectorURL,omitnil" name:"PrivateLinkCollectorURL"`
}

type ApmField struct {
	// 昨日同比指标值，已弃用，不建议使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareVal *string `json:"CompareVal,omitnil" name:"CompareVal"`

	// Compare值结果数组，推荐使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompareVals []*APMKVItem `json:"CompareVals,omitnil" name:"CompareVals"`

	// 指标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *float64 `json:"Value,omitnil" name:"Value"`

	// 指标所对应的单位
	// 注意：此字段可能返回 null，表示取不到有效值。
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 请求数
	Key *string `json:"Key,omitnil" name:"Key"`

	// 同环比上周期具体数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastPeriodValue []*APMKV `json:"LastPeriodValue,omitnil" name:"LastPeriodValue"`
}

type ApmInstanceDetail struct {
	// 存储使用量(MB)
	// 注意：此字段可能返回 null，表示取不到有效值。
	AmountOfUsedStorage *float64 `json:"AmountOfUsedStorage,omitnil" name:"AmountOfUsedStorage"`

	// 实例名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 实例所属tag列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 创建人Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateUin *string `json:"CreateUin,omitnil" name:"CreateUin"`

	// 该实例已上报的服务端应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceCount *int64 `json:"ServiceCount,omitnil" name:"ServiceCount"`

	// 日均上报Span数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountOfReportSpanPerDay *int64 `json:"CountOfReportSpanPerDay,omitnil" name:"CountOfReportSpanPerDay"`

	// AppId信息
	AppId *int64 `json:"AppId,omitnil" name:"AppId"`

	// Trace数据保存时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceDuration *int64 `json:"TraceDuration,omitnil" name:"TraceDuration"`

	// 实例描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// 实例所属地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region *string `json:"Region,omitnil" name:"Region"`

	// 实例上报额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpanDailyCounters *int64 `json:"SpanDailyCounters,omitnil" name:"SpanDailyCounters"`

	// 实例是否开通计费
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingInstance *int64 `json:"BillingInstance,omitnil" name:"BillingInstance"`

	// 错误率阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil" name:"ErrRateThreshold"`

	// 采样率阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleRate *int64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 是否开启错误采样 0  关 1 开
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorSample *int64 `json:"ErrorSample,omitnil" name:"ErrorSample"`

	// 慢调用保存阈值
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil" name:"SlowRequestSavedThreshold"`

	// cls日志所在地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogRegion *string `json:"LogRegion,omitnil" name:"LogRegion"`

	// 日志来源
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSource *string `json:"LogSource,omitnil" name:"LogSource"`

	// 日志功能开关 0 关 | 1 开
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil" name:"IsRelatedLog"`

	// 日志主题ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogTopicID *string `json:"LogTopicID,omitnil" name:"LogTopicID"`

	// 该实例已上报的客户端应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientCount *int64 `json:"ClientCount,omitnil" name:"ClientCount"`

	// 该实例已上报的总应用数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// CLS日志集 | ES集群ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	LogSet *string `json:"LogSet,omitnil" name:"LogSet"`

	// Metric数据保存时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricDuration *int64 `json:"MetricDuration,omitnil" name:"MetricDuration"`

	// 用户自定义展示标签列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomShowTags []*string `json:"CustomShowTags,omitnil" name:"CustomShowTags"`

	// 实例计费模式
	// 1为预付费
	// 0为按量付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`

	// 实例计费模式是否生效
	// 注意：此字段可能返回 null，表示取不到有效值。
	PayModeEffective *bool `json:"PayModeEffective,omitnil" name:"PayModeEffective"`
}

type ApmMetricRecord struct {
	// field数组
	Fields []*ApmField `json:"Fields,omitnil" name:"Fields"`

	// tag数组
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`
}

type ApmTag struct {
	// 维度Key(列名，标签Key)
	Key *string `json:"Key,omitnil" name:"Key"`

	// 维度值（标签值）
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type CreateApmInstanceRequestParams struct {
	// 实例名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 实例描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// Trace数据保存时长
	TraceDuration *int64 `json:"TraceDuration,omitnil" name:"TraceDuration"`

	// 标签列表
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`

	// 实例上报额度值
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil" name:"SpanDailyCounters"`

	// 实例的计费模式
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`
}

type CreateApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 实例描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// Trace数据保存时长
	TraceDuration *int64 `json:"TraceDuration,omitnil" name:"TraceDuration"`

	// 标签列表
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`

	// 实例上报额度值
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil" name:"SpanDailyCounters"`

	// 实例的计费模式
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateApmInstanceResponseParams struct {
	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 接入方式
	AgentType *string `json:"AgentType,omitnil" name:"AgentType"`

	// 环境
	NetworkMode *string `json:"NetworkMode,omitnil" name:"NetworkMode"`

	// 语言
	LanguageEnvironment *string `json:"LanguageEnvironment,omitnil" name:"LanguageEnvironment"`

	// 上报方式
	ReportMethod *string `json:"ReportMethod,omitnil" name:"ReportMethod"`
}

type DescribeApmAgentRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 接入方式
	AgentType *string `json:"AgentType,omitnil" name:"AgentType"`

	// 环境
	NetworkMode *string `json:"NetworkMode,omitnil" name:"NetworkMode"`

	// 语言
	LanguageEnvironment *string `json:"LanguageEnvironment,omitnil" name:"LanguageEnvironment"`

	// 上报方式
	ReportMethod *string `json:"ReportMethod,omitnil" name:"ReportMethod"`
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
	// Agent信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ApmAgent *ApmAgentInfo `json:"ApmAgent,omitnil" name:"ApmAgent"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// Tag列表
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`

	// 搜索实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 过滤实例ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 是否查询官方demo实例
	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitnil" name:"DemoInstanceFlag"`

	// 是否查询全地域实例
	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitnil" name:"AllRegionsFlag"`
}

type DescribeApmInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Tag列表
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`

	// 搜索实例名
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// 过滤实例ID
	InstanceIds []*string `json:"InstanceIds,omitnil" name:"InstanceIds"`

	// 是否查询官方demo实例
	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitnil" name:"DemoInstanceFlag"`

	// 是否查询全地域实例
	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitnil" name:"AllRegionsFlag"`
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
	// apm实例列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instances []*ApmInstanceDetail `json:"Instances,omitnil" name:"Instances"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeGeneralMetricDataRequestParams struct {
	// 要过滤的维度信息：
	// service_metric视图支持：service.name（服务名）、span.kind（客户端/服务端视角）为维度进行过滤，service.name（服务名）必填。
	// span.kind:
	// 	server:服务端视角
	// 	client:客户端视角
	// 默认为服务端视角进行查询。
	// runtime_metric视图支持：service.name（服务名）维度进行过滤，service.name（服务名）必填。
	// sql_metric视图支持：service.name（服务名）、db.instance（数据库名称）、db.ip（数据库实例ip）维度进行过滤，查询service_slow_sql_count（慢sql）指标时service.name必填，查询sql_duration_avg（耗时）指标时db.instance（数据库名称）必填。
	Filters []*GeneralFilter `json:"Filters,omitnil" name:"Filters"`

	// 需要查询的指标，不可自定义输入。
	// service_metric视图支持：service_request_count（总请求）、service_duration（平均响应时间）、service_error_req_rate（平均错误率）、service_slow_call_count（慢调用）、service_error_request_count（异常数量）。
	// runtime_metric视图支持：service_gc_full_count（Full GC）。
	// sql_metric视图支持：service_slow_sql_count（慢sql）、sql_duration_avg（耗时）。
	Metrics []*string `json:"Metrics,omitnil" name:"Metrics"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 视图名称，不可自定义输入。支持：service_metric、runtime_metric、sql_metric。
	ViewName *string `json:"ViewName,omitnil" name:"ViewName"`

	// 聚合维度：
	// service_metric视图支持：service.name（服务名）、span.kind （客户端/服务端视角）维度进行聚合，service.name（服务名）必填。
	// runtime_metric视图支持：service.name（服务名）维度进行聚合，service.name（服务名）必填。
	// sql_metric视图支持：service.name（服务名）、db.statement（sql语句）维度进行聚合，查询service_slow_sql_count（慢sql）时service.name（服务名）必填，查询sql_duration_avg（耗时）指标时service.name（服务名）、db.statement（sql语句）必填。
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 起始时间的时间戳，单位为秒，只支持查询2天内最多1小时的指标数据。
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间的时间戳，单位为秒，只支持查询2天内最多1小时的指标数据。
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 聚合粒度，单位为秒，最小为60s，即一分钟的聚合粒度；如果为空或0则计算开始时间到截止时间的指标数据，上报其他值会报错。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 对查询指标进行排序：
	// service_metric视图支持：service_request_count（总请求）、service_duration（平均响应时间）、service_error_req_rate（平均错误率）、service_slow_call_count（慢调用）、service_error_request_count（异常数量）。
	// runtime_metric视图支持：service_gc_full_count（Full GC）。
	// sql_metric视图支持：service_slow_sql_count（慢sql）、sql_duration_avg（耗时）。
	// asc:对查询指标进行升序排序
	// desc：对查询指标进行降序排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// 查询指标的限制条数，目前最多展示50条数据，PageSize取值为1-50，上送PageSize则根据PageSize的值展示限制条数。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
}

type DescribeGeneralMetricDataRequest struct {
	*tchttp.BaseRequest
	
	// 要过滤的维度信息：
	// service_metric视图支持：service.name（服务名）、span.kind（客户端/服务端视角）为维度进行过滤，service.name（服务名）必填。
	// span.kind:
	// 	server:服务端视角
	// 	client:客户端视角
	// 默认为服务端视角进行查询。
	// runtime_metric视图支持：service.name（服务名）维度进行过滤，service.name（服务名）必填。
	// sql_metric视图支持：service.name（服务名）、db.instance（数据库名称）、db.ip（数据库实例ip）维度进行过滤，查询service_slow_sql_count（慢sql）指标时service.name必填，查询sql_duration_avg（耗时）指标时db.instance（数据库名称）必填。
	Filters []*GeneralFilter `json:"Filters,omitnil" name:"Filters"`

	// 需要查询的指标，不可自定义输入。
	// service_metric视图支持：service_request_count（总请求）、service_duration（平均响应时间）、service_error_req_rate（平均错误率）、service_slow_call_count（慢调用）、service_error_request_count（异常数量）。
	// runtime_metric视图支持：service_gc_full_count（Full GC）。
	// sql_metric视图支持：service_slow_sql_count（慢sql）、sql_duration_avg（耗时）。
	Metrics []*string `json:"Metrics,omitnil" name:"Metrics"`

	// 业务系统ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 视图名称，不可自定义输入。支持：service_metric、runtime_metric、sql_metric。
	ViewName *string `json:"ViewName,omitnil" name:"ViewName"`

	// 聚合维度：
	// service_metric视图支持：service.name（服务名）、span.kind （客户端/服务端视角）维度进行聚合，service.name（服务名）必填。
	// runtime_metric视图支持：service.name（服务名）维度进行聚合，service.name（服务名）必填。
	// sql_metric视图支持：service.name（服务名）、db.statement（sql语句）维度进行聚合，查询service_slow_sql_count（慢sql）时service.name（服务名）必填，查询sql_duration_avg（耗时）指标时service.name（服务名）、db.statement（sql语句）必填。
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 起始时间的时间戳，单位为秒，只支持查询2天内最多1小时的指标数据。
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 结束时间的时间戳，单位为秒，只支持查询2天内最多1小时的指标数据。
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`

	// 聚合粒度，单位为秒，最小为60s，即一分钟的聚合粒度；如果为空或0则计算开始时间到截止时间的指标数据，上报其他值会报错。
	Period *int64 `json:"Period,omitnil" name:"Period"`

	// 对查询指标进行排序：
	// service_metric视图支持：service_request_count（总请求）、service_duration（平均响应时间）、service_error_req_rate（平均错误率）、service_slow_call_count（慢调用）、service_error_request_count（异常数量）。
	// runtime_metric视图支持：service_gc_full_count（Full GC）。
	// sql_metric视图支持：service_slow_sql_count（慢sql）、sql_duration_avg（耗时）。
	// asc:对查询指标进行升序排序
	// desc：对查询指标进行降序排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// 查询指标的限制条数，目前最多展示50条数据，PageSize取值为1-50，上送PageSize则根据PageSize的值展示限制条数。
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`
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
	delete(f, "Filters")
	delete(f, "Metrics")
	delete(f, "InstanceId")
	delete(f, "ViewName")
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
	// 注意：此字段可能返回 null，表示取不到有效值。
	Records []*Line `json:"Records,omitnil" name:"Records"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
type DescribeGeneralSpanListRequestParams struct {
	// 分页
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 列表项个数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// span查询开始时间戳（单位:秒）
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 实例名
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 通用过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 业务自身服务名
	BusinessName *string `json:"BusinessName,omitnil" name:"BusinessName"`

	// span查询结束时间戳（单位:秒）
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeGeneralSpanListRequest struct {
	*tchttp.BaseRequest
	
	// 分页
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 列表项个数
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// span查询开始时间戳（单位:秒）
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 实例名
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 通用过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 业务自身服务名
	BusinessName *string `json:"BusinessName,omitnil" name:"BusinessName"`

	// span查询结束时间戳（单位:秒）
	EndTime *int64 `json:"EndTime,omitnil" name:"EndTime"`
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
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "OrderBy")
	delete(f, "StartTime")
	delete(f, "InstanceId")
	delete(f, "Filters")
	delete(f, "BusinessName")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGeneralSpanListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGeneralSpanListResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Span分页列表
	Spans []*Span `json:"Spans,omitnil" name:"Spans"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil" name:"Metrics"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 业务名称（默认值：taw）
	BusinessName *string `json:"BusinessName,omitnil" name:"BusinessName"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil" name:"PageIndex"`

	// 页长
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// Or过滤条件
	OrFilters []*Filter `json:"OrFilters,omitnil" name:"OrFilters"`
}

type DescribeMetricRecordsRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil" name:"Metrics"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 每页大小
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`

	// 业务名称（默认值：taw）
	BusinessName *string `json:"BusinessName,omitnil" name:"BusinessName"`

	// 页码
	PageIndex *int64 `json:"PageIndex,omitnil" name:"PageIndex"`

	// 页长
	PageSize *int64 `json:"PageSize,omitnil" name:"PageSize"`

	// Or过滤条件
	OrFilters []*Filter `json:"OrFilters,omitnil" name:"OrFilters"`
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
	delete(f, "Filters")
	delete(f, "Metrics")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "Offset")
	delete(f, "EndTime")
	delete(f, "BusinessName")
	delete(f, "PageIndex")
	delete(f, "PageSize")
	delete(f, "OrFilters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMetricRecordsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeMetricRecordsResponseParams struct {
	// 指标结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Records []*ApmMetricRecord `json:"Records,omitnil" name:"Records"`

	// 查询指标结果集条数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil" name:"Metrics"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 每页大小
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
}

type DescribeServiceOverviewRequest struct {
	*tchttp.BaseRequest
	
	// 过滤条件
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 指标列表
	Metrics []*QueryMetricItem `json:"Metrics,omitnil" name:"Metrics"`

	// 聚合维度
	GroupBy []*string `json:"GroupBy,omitnil" name:"GroupBy"`

	// 排序
	OrderBy *OrderBy `json:"OrderBy,omitnil" name:"OrderBy"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 每页大小
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 开始时间
	StartTime *uint64 `json:"StartTime,omitnil" name:"StartTime"`

	// 分页起始点
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 结束时间
	EndTime *uint64 `json:"EndTime,omitnil" name:"EndTime"`
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
	delete(f, "Filters")
	delete(f, "Metrics")
	delete(f, "GroupBy")
	delete(f, "OrderBy")
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "StartTime")
	delete(f, "Offset")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeServiceOverviewRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeServiceOverviewResponseParams struct {
	// 指标结果集
	// 注意：此字段可能返回 null，表示取不到有效值。
	Records []*ApmMetricRecord `json:"Records,omitnil" name:"Records"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type Filter struct {
	// 过滤方式（=, !=, in）
	Type *string `json:"Type,omitnil" name:"Type"`

	// 过滤维度名
	Key *string `json:"Key,omitnil" name:"Key"`

	// 过滤值，in过滤方式用逗号分割多个值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type GeneralFilter struct {
	// 过滤维度名
	Key *string `json:"Key,omitnil" name:"Key"`

	// 过滤值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type Line struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 指标中文名
	MetricNameCN *string `json:"MetricNameCN,omitnil" name:"MetricNameCN"`

	// 时间序列
	TimeSerial []*int64 `json:"TimeSerial,omitnil" name:"TimeSerial"`

	// 数据序列
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataSerial []*float64 `json:"DataSerial,omitnil" name:"DataSerial"`

	// 维度列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`
}

// Predefined struct for user
type ModifyApmInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标签列表
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`

	// 实例详情
	Description *string `json:"Description,omitnil" name:"Description"`

	// Trace数据保存时长
	TraceDuration *int64 `json:"TraceDuration,omitnil" name:"TraceDuration"`

	// 是否开启计费
	OpenBilling *bool `json:"OpenBilling,omitnil" name:"OpenBilling"`

	// 实例上报额度
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil" name:"SpanDailyCounters"`

	// 错误率阈值
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil" name:"ErrRateThreshold"`

	// 采样率
	SampleRate *int64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 是否开启错误采样 0 关 1 开
	ErrorSample *int64 `json:"ErrorSample,omitnil" name:"ErrorSample"`

	// 慢请求阈值
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil" name:"SlowRequestSavedThreshold"`

	// 是否开启日志功能 0 关 1 开
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil" name:"IsRelatedLog"`

	// 日志地域
	LogRegion *string `json:"LogRegion,omitnil" name:"LogRegion"`

	// CLS日志主题ID | ES 索引名
	LogTopicID *string `json:"LogTopicID,omitnil" name:"LogTopicID"`

	// CLS日志集 | ES集群ID
	LogSet *string `json:"LogSet,omitnil" name:"LogSet"`

	// CLS | ES
	LogSource *string `json:"LogSource,omitnil" name:"LogSource"`

	// 用户自定义展示标签列表
	CustomShowTags []*string `json:"CustomShowTags,omitnil" name:"CustomShowTags"`

	// 修改计费模式
	// 1为预付费
	// 0为按量付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`
}

type ModifyApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 实例名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标签列表
	Tags []*ApmTag `json:"Tags,omitnil" name:"Tags"`

	// 实例详情
	Description *string `json:"Description,omitnil" name:"Description"`

	// Trace数据保存时长
	TraceDuration *int64 `json:"TraceDuration,omitnil" name:"TraceDuration"`

	// 是否开启计费
	OpenBilling *bool `json:"OpenBilling,omitnil" name:"OpenBilling"`

	// 实例上报额度
	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitnil" name:"SpanDailyCounters"`

	// 错误率阈值
	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitnil" name:"ErrRateThreshold"`

	// 采样率
	SampleRate *int64 `json:"SampleRate,omitnil" name:"SampleRate"`

	// 是否开启错误采样 0 关 1 开
	ErrorSample *int64 `json:"ErrorSample,omitnil" name:"ErrorSample"`

	// 慢请求阈值
	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitnil" name:"SlowRequestSavedThreshold"`

	// 是否开启日志功能 0 关 1 开
	IsRelatedLog *int64 `json:"IsRelatedLog,omitnil" name:"IsRelatedLog"`

	// 日志地域
	LogRegion *string `json:"LogRegion,omitnil" name:"LogRegion"`

	// CLS日志主题ID | ES 索引名
	LogTopicID *string `json:"LogTopicID,omitnil" name:"LogTopicID"`

	// CLS日志集 | ES集群ID
	LogSet *string `json:"LogSet,omitnil" name:"LogSet"`

	// CLS | ES
	LogSource *string `json:"LogSource,omitnil" name:"LogSource"`

	// 用户自定义展示标签列表
	CustomShowTags []*string `json:"CustomShowTags,omitnil" name:"CustomShowTags"`

	// 修改计费模式
	// 1为预付费
	// 0为按量付费
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`
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
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyApmInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyApmInstanceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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

type OrderBy struct {
	// 需要排序的字段
	Key *string `json:"Key,omitnil" name:"Key"`

	// 顺序排序/倒序排序
	Value *string `json:"Value,omitnil" name:"Value"`
}

type QueryMetricItem struct {
	// 指标名
	MetricName *string `json:"MetricName,omitnil" name:"MetricName"`

	// 同比，已弃用，不建议使用
	Compare *string `json:"Compare,omitnil" name:"Compare"`

	// 同比，支持多种同比方式
	Compares []*string `json:"Compares,omitnil" name:"Compares"`
}

type Span struct {
	// Trace Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	TraceID *string `json:"TraceID,omitnil" name:"TraceID"`

	// 日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	Logs []*SpanLog `json:"Logs,omitnil" name:"Logs"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*SpanTag `json:"Tags,omitnil" name:"Tags"`

	// 上报应用服务信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Process *SpanProcess `json:"Process,omitnil" name:"Process"`

	// 产生时间戳(毫秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// Span名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperationName *string `json:"OperationName,omitnil" name:"OperationName"`

	// 关联关系
	// 注意：此字段可能返回 null，表示取不到有效值。
	References []*SpanReference `json:"References,omitnil" name:"References"`

	// 产生时间戳(微秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *int64 `json:"StartTime,omitnil" name:"StartTime"`

	// 持续耗时(微妙)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitnil" name:"Duration"`

	// Span Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpanID *string `json:"SpanID,omitnil" name:"SpanID"`

	// 产生时间戳(毫秒)
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeMillis *int64 `json:"StartTimeMillis,omitnil" name:"StartTimeMillis"`

	// Parent Span Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ParentSpanID *string `json:"ParentSpanID,omitnil" name:"ParentSpanID"`
}

type SpanLog struct {
	// 日志时间戳
	Timestamp *int64 `json:"Timestamp,omitnil" name:"Timestamp"`

	// 标签
	Fields []*SpanTag `json:"Fields,omitnil" name:"Fields"`
}

type SpanProcess struct {
	// 应用服务名称
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// Tags 标签数组
	Tags []*SpanTag `json:"Tags,omitnil" name:"Tags"`
}

type SpanReference struct {
	// 关联关系类型
	RefType *string `json:"RefType,omitnil" name:"RefType"`

	// Span ID
	SpanID *string `json:"SpanID,omitnil" name:"SpanID"`

	// Trace ID
	TraceID *string `json:"TraceID,omitnil" name:"TraceID"`
}

type SpanTag struct {
	// 标签类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// 标签Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type TerminateApmInstanceRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type TerminateApmInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
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
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
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