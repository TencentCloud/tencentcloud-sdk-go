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

package v20201207

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type ApolloEnvParam struct {
	// 环境名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 环境内引擎的节点规格 ID
	// -1C2G
	// -2C4G
	// 兼容原spec-xxxxxx形式的规格ID
	EngineResourceSpec *string `json:"EngineResourceSpec,omitnil" name:"EngineResourceSpec"`

	// 环境内引擎的节点数量
	EngineNodeNum *int64 `json:"EngineNodeNum,omitnil" name:"EngineNodeNum"`

	// 配置存储空间大小，以GB为单位
	StorageCapacity *int64 `json:"StorageCapacity,omitnil" name:"StorageCapacity"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 环境描述
	EnvDesc *string `json:"EnvDesc,omitnil" name:"EnvDesc"`
}

type AutoScalerBehavior struct {
	// 扩容行为配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleUp *AutoScalerRules `json:"ScaleUp,omitnil" name:"ScaleUp"`

	// 缩容行为配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScaleDown *AutoScalerRules `json:"ScaleDown,omitnil" name:"ScaleDown"`
}

type AutoScalerPolicy struct {
	// 类型，Pods或Percent
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 数量
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *int64 `json:"Value,omitnil" name:"Value"`

	// 扩容周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil" name:"PeriodSeconds"`
}

type AutoScalerRules struct {
	// 稳定窗口时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitnil" name:"StabilizationWindowSeconds"`

	// 选择策略依据
	// 注意：此字段可能返回 null，表示取不到有效值。
	SelectPolicy *string `json:"SelectPolicy,omitnil" name:"SelectPolicy"`

	// 扩容策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policies []*AutoScalerPolicy `json:"Policies,omitnil" name:"Policies"`
}

type BoundK8SInfo struct {
	// 绑定的kubernetes集群ID
	BoundClusterId *string `json:"BoundClusterId,omitnil" name:"BoundClusterId"`

	// 绑定的kubernetes的集群类型，分tke和eks两种
	// 注意：此字段可能返回 null，表示取不到有效值。
	BoundClusterType *string `json:"BoundClusterType,omitnil" name:"BoundClusterType"`

	// 服务同步模式，all为全量同步，demand为按需同步
	// 注意：此字段可能返回 null，表示取不到有效值。
	SyncMode *string `json:"SyncMode,omitnil" name:"SyncMode"`
}

type CertificateInfo struct {
	// 唯一id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`
}

type CloudAPIGatewayCanaryRuleList struct {
	// 灰度规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanaryRuleList []*CloudNativeAPIGatewayCanaryRule `json:"CanaryRuleList,omitnil" name:"CanaryRuleList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type CloudNativeAPIGatewayBalancedService struct {
	// 服务 ID，作为入参时，必填
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceID *string `json:"ServiceID,omitnil" name:"ServiceID"`

	// 服务名称，作为入参时，无意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// Upstream 名称，作为入参时，无意义
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamName *string `json:"UpstreamName,omitnil" name:"UpstreamName"`

	// 百分比，10 即 10%，范围0-100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Percent *float64 `json:"Percent,omitnil" name:"Percent"`
}

type CloudNativeAPIGatewayCanaryRule struct {
	// 优先级，值范围为 0 到 100；值越大，优先级越高；不同规则间优先级不可重复
	// 注意：此字段可能返回 null，表示取不到有效值。
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 是否启用规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil" name:"Enabled"`

	// 参数匹配条件
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConditionList []*CloudNativeAPIGatewayCanaryRuleCondition `json:"ConditionList,omitnil" name:"ConditionList"`

	// 服务的流量百分比配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BalancedServiceList []*CloudNativeAPIGatewayBalancedService `json:"BalancedServiceList,omitnil" name:"BalancedServiceList"`

	// 归属服务 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 归属服务名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`
}

type CloudNativeAPIGatewayCanaryRuleCondition struct {
	// 条件类型，支持 path, method, query, header, cookie, body 和 system。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 参数名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 操作符，支持 "le", "eq", "lt", "ne", "ge", "gt", "regex", "exists", "in", "not in",  "prefix" ,"exact", "regex" 等
	// 注意：此字段可能返回 null，表示取不到有效值。
	Operator *string `json:"Operator,omitnil" name:"Operator"`

	// 目标参数值
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`

	// 分隔符，当 Operator 为 in 或者 not in 时生效。支持值为英文逗号，英文分号，空格，换行符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Delimiter *string `json:"Delimiter,omitnil" name:"Delimiter"`

	// 全局配置 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalConfigId *string `json:"GlobalConfigId,omitnil" name:"GlobalConfigId"`

	// 全局配置名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GlobalConfigName *string `json:"GlobalConfigName,omitnil" name:"GlobalConfigName"`
}

type CloudNativeAPIGatewayConfig struct {
	// 控制台类型。
	ConsoleType *string `json:"ConsoleType,omitnil" name:"ConsoleType"`

	// HTTP链接地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpUrl *string `json:"HttpUrl,omitnil" name:"HttpUrl"`

	// HTTPS链接地址。
	HttpsUrl *string `json:"HttpsUrl,omitnil" name:"HttpsUrl"`

	// 网络类型, Open|Internal。
	NetType *string `json:"NetType,omitnil" name:"NetType"`

	// 管理员用户名。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminUser *string `json:"AdminUser,omitnil" name:"AdminUser"`

	// 管理员密码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdminPassword *string `json:"AdminPassword,omitnil" name:"AdminPassword"`

	// 网络状态, Open|Closed|Updating
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 网络访问策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessControl *NetworkAccessControl `json:"AccessControl,omitnil" name:"AccessControl"`

	// 内网子网 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 内网VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 负载均衡的描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 负载均衡的规格类型，传 "SLA" 表示性能容量型，返回空为共享型
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaType *string `json:"SlaType,omitnil" name:"SlaType"`

	// clb vip
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vip *string `json:"Vip,omitnil" name:"Vip"`

	// 带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 是否多可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil" name:"MultiZoneFlag"`

	// 主可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZoneId *string `json:"MasterZoneId,omitnil" name:"MasterZoneId"`

	// 备可用区
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZoneId *string `json:"SlaveZoneId,omitnil" name:"SlaveZoneId"`

	// 主可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	MasterZoneName *string `json:"MasterZoneName,omitnil" name:"MasterZoneName"`

	// 备可用区名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	SlaveZoneName *string `json:"SlaveZoneName,omitnil" name:"SlaveZoneName"`

	// 网络 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	NetworkId *string `json:"NetworkId,omitnil" name:"NetworkId"`
}

type CloudNativeAPIGatewayNode struct {
	// 云原生网关节点 id
	NodeId *string `json:"NodeId,omitnil" name:"NodeId"`

	// 节点 ip
	NodeIp *string `json:"NodeIp,omitnil" name:"NodeIp"`

	// Zone id
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// Zone
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 分组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName *string `json:"GroupName,omitnil" name:"GroupName"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`
}

type CloudNativeAPIGatewayNodeConfig struct {
	// 节点配置, 1c2g|2c4g|4c8g|8c16g。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// 节点数量，2-9。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Number *int64 `json:"Number,omitnil" name:"Number"`
}

type CloudNativeAPIGatewayRateLimitDetail struct {
	// 插件启用状态
	Enabled *bool `json:"Enabled,omitnil" name:"Enabled"`

	// qps阈值
	QpsThresholds []*QpsThreshold `json:"QpsThresholds,omitnil" name:"QpsThresholds"`

	// 限流依据
	// ip service consumer credential path header
	LimitBy *string `json:"LimitBy,omitnil" name:"LimitBy"`

	// 响应策略
	// url请求转发
	// text 响应配置
	// default 直接返回
	ResponseType *string `json:"ResponseType,omitnil" name:"ResponseType"`

	// 是否隐藏限流客户端响应头
	HideClientHeaders *bool `json:"HideClientHeaders,omitnil" name:"HideClientHeaders"`

	// 是否开启请求排队
	IsDelay *bool `json:"IsDelay,omitnil" name:"IsDelay"`

	// 需要进行流量控制的请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 需要进行流量控制的请求头Key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Header *string `json:"Header,omitnil" name:"Header"`

	// 外部redis配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExternalRedis *ExternalRedis `json:"ExternalRedis,omitnil" name:"ExternalRedis"`

	// 计数器策略 
	// local 单机
	// redis  默认redis
	// external_redis 外部redis
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Policy *string `json:"Policy,omitnil" name:"Policy"`

	// 响应配置，响应策略为text
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitResponse *RateLimitResponse `json:"RateLimitResponse,omitnil" name:"RateLimitResponse"`

	// 请求转发地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	RateLimitResponseUrl *string `json:"RateLimitResponseUrl,omitnil" name:"RateLimitResponseUrl"`

	// 排队时间
	LineUpTime *int64 `json:"LineUpTime,omitnil" name:"LineUpTime"`
}

type CloudNativeAPIGatewayStrategy struct {
	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil" name:"StrategyId"`

	// 策略名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrategyName *string `json:"StrategyName,omitnil" name:"StrategyName"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 策略描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 弹性伸缩配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitnil" name:"Config"`

	// 网关实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 定时伸缩配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CronConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronConfig,omitnil" name:"CronConfig"`

	// 最大节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: MaxReplicas is deprecated.
	MaxReplicas *uint64 `json:"MaxReplicas,omitnil" name:"MaxReplicas"`
}

type CloudNativeAPIGatewayStrategyAutoScalerConfig struct {
	// 最大副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxReplicas *int64 `json:"MaxReplicas,omitnil" name:"MaxReplicas"`

	// 指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*CloudNativeAPIGatewayStrategyAutoScalerConfigMetric `json:"Metrics,omitnil" name:"Metrics"`

	// 是否开启指标伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil" name:"Enabled"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: CreateTime is deprecated.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModifyTime is deprecated.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 弹性策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: StrategyId is deprecated.
	StrategyId *string `json:"StrategyId,omitnil" name:"StrategyId"`

	// 指标配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: AutoScalerId is deprecated.
	AutoScalerId *string `json:"AutoScalerId,omitnil" name:"AutoScalerId"`

	// 指标伸缩行为配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Behavior *AutoScalerBehavior `json:"Behavior,omitnil" name:"Behavior"`
}

type CloudNativeAPIGatewayStrategyAutoScalerConfigMetric struct {
	// 指标类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 指标资源名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceName *string `json:"ResourceName,omitnil" name:"ResourceName"`

	// 指标目标类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetType *string `json:"TargetType,omitnil" name:"TargetType"`

	// 指标目标值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetValue *int64 `json:"TargetValue,omitnil" name:"TargetValue"`
}

type CloudNativeAPIGatewayStrategyCronScalerConfig struct {
	// 是否开启定时伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil" name:"Enabled"`

	// 定时伸缩配置参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*CloudNativeAPIGatewayStrategyCronScalerConfigParam `json:"Params,omitnil" name:"Params"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: CreateTime is deprecated.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModifyTime is deprecated.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 弹性策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: StrategyId is deprecated.
	StrategyId *string `json:"StrategyId,omitnil" name:"StrategyId"`
}

type CloudNativeAPIGatewayStrategyCronScalerConfigParam struct {
	// 定时伸缩周期
	// 注意：此字段可能返回 null，表示取不到有效值。
	Period *string `json:"Period,omitnil" name:"Period"`

	// 定时伸缩开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartAt *string `json:"StartAt,omitnil" name:"StartAt"`

	// 定时伸缩目标节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetReplicas *int64 `json:"TargetReplicas,omitnil" name:"TargetReplicas"`

	// 定时伸缩cron表达式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Crontab *string `json:"Crontab,omitnil" name:"Crontab"`
}

type CloudNativeAPIGatewayVpcConfig struct {
	// 私有网络ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayCanaryRuleRequestParams struct {
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil" name:"CanaryRule"`
}

type CreateCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil" name:"CanaryRule"`
}

func (r *CreateCloudNativeAPIGatewayCanaryRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayCanaryRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceId")
	delete(f, "CanaryRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayCanaryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayCanaryRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayCanaryRuleResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayCanaryRuleResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayCanaryRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayCanaryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayCertificateRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 绑定的域名
	BindDomains []*string `json:"BindDomains,omitnil" name:"BindDomains"`

	// ssl平台证书 Id
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// 证书名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 证书私钥
	//
	// Deprecated: Key is deprecated.
	Key *string `json:"Key,omitnil" name:"Key"`

	// 证书pem格式
	//
	// Deprecated: Crt is deprecated.
	Crt *string `json:"Crt,omitnil" name:"Crt"`
}

type CreateCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 绑定的域名
	BindDomains []*string `json:"BindDomains,omitnil" name:"BindDomains"`

	// ssl平台证书 Id
	CertId *string `json:"CertId,omitnil" name:"CertId"`

	// 证书名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 证书私钥
	Key *string `json:"Key,omitnil" name:"Key"`

	// 证书pem格式
	Crt *string `json:"Crt,omitnil" name:"Crt"`
}

func (r *CreateCloudNativeAPIGatewayCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "BindDomains")
	delete(f, "CertId")
	delete(f, "Name")
	delete(f, "Key")
	delete(f, "Crt")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayCertificateResponseParams struct {
	// 创建证书结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CertificateInfo `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayCertificateResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayCertificateResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRequestParams struct {
	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 云原生API网关类型, 目前只支持kong。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 云原生API网关版本。参考值：
	// - 2.4.1
	// - 2.5.1
	GatewayVersion *string `json:"GatewayVersion,omitnil" name:"GatewayVersion"`

	// 云原生API网关节点配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`

	// 云原生API网关vpc配置。
	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitnil" name:"VpcConfig"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 标签列表
	Tags []*InstanceTagInfo `json:"Tags,omitnil" name:"Tags"`

	// 是否开启 CLS 日志。默认值：fasle
	EnableCls *bool `json:"EnableCls,omitnil" name:"EnableCls"`

	// 产品版本。参考值：
	// - TRIAL：开发版
	// - STANDARD：标准版 （默认值）
	// - PROFESSIONAL：专业版
	FeatureVersion *string `json:"FeatureVersion,omitnil" name:"FeatureVersion"`

	// 公网出流量带宽，[1,2048]Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 实例实际的地域信息，默认值：ap-guangzhou
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// ingress Class名称
	IngressClassName *string `json:"IngressClassName,omitnil" name:"IngressClassName"`

	// 付费类型。参考值：
	// 0：后付费（默认值）
	// 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil" name:"TradeType"`

	// 公网相关配置
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil" name:"InternetConfig"`
}

type CreateCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 云原生API网关类型, 目前只支持kong。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 云原生API网关版本。参考值：
	// - 2.4.1
	// - 2.5.1
	GatewayVersion *string `json:"GatewayVersion,omitnil" name:"GatewayVersion"`

	// 云原生API网关节点配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`

	// 云原生API网关vpc配置。
	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitnil" name:"VpcConfig"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 标签列表
	Tags []*InstanceTagInfo `json:"Tags,omitnil" name:"Tags"`

	// 是否开启 CLS 日志。默认值：fasle
	EnableCls *bool `json:"EnableCls,omitnil" name:"EnableCls"`

	// 产品版本。参考值：
	// - TRIAL：开发版
	// - STANDARD：标准版 （默认值）
	// - PROFESSIONAL：专业版
	FeatureVersion *string `json:"FeatureVersion,omitnil" name:"FeatureVersion"`

	// 公网出流量带宽，[1,2048]Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 实例实际的地域信息，默认值：ap-guangzhou
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// ingress Class名称
	IngressClassName *string `json:"IngressClassName,omitnil" name:"IngressClassName"`

	// 付费类型。参考值：
	// 0：后付费（默认值）
	// 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil" name:"TradeType"`

	// 公网相关配置
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil" name:"InternetConfig"`
}

func (r *CreateCloudNativeAPIGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Type")
	delete(f, "GatewayVersion")
	delete(f, "NodeConfig")
	delete(f, "VpcConfig")
	delete(f, "Description")
	delete(f, "Tags")
	delete(f, "EnableCls")
	delete(f, "FeatureVersion")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "EngineRegion")
	delete(f, "IngressClassName")
	delete(f, "TradeType")
	delete(f, "InternetConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayResponseParams struct {
	// 创建云原生API网关实例响应结果。
	Result *CreateCloudNativeAPIGatewayResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayResult struct {
	// 云原生API网关ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 云原生网关状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRouteRateLimitRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

type CreateCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

func (r *CreateCloudNativeAPIGatewayRouteRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayRouteRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	delete(f, "LimitDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayRouteRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRouteRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayRouteRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayRouteRateLimitResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayRouteRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayRouteRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRouteRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil" name:"ServiceID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 路由的方法，其中方法可选值：
	// - GET
	// - POST
	// - DELETE
	// - PUT
	// - OPTIONS
	// - PATCH
	// - HEAD
	// - ANY
	// - TRACE
	// - COPY
	// - MOVE
	// - PROPFIND
	// - PROPPATCH
	// - MKCOL
	// - LOCK
	// - UNLOCK
	Methods []*string `json:"Methods,omitnil" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil" name:"StripPath"`

	// 是否开启强制HTTPS
	//
	// Deprecated: ForceHttps is deprecated.
	ForceHttps *bool `json:"ForceHttps,omitnil" name:"ForceHttps"`

	// 四层匹配的目的端口
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil" name:"Headers"`
}

type CreateCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil" name:"ServiceID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 路由的方法，其中方法可选值：
	// - GET
	// - POST
	// - DELETE
	// - PUT
	// - OPTIONS
	// - PATCH
	// - HEAD
	// - ANY
	// - TRACE
	// - COPY
	// - MOVE
	// - PROPFIND
	// - PROPPATCH
	// - MKCOL
	// - LOCK
	// - UNLOCK
	Methods []*string `json:"Methods,omitnil" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil" name:"StripPath"`

	// 是否开启强制HTTPS
	ForceHttps *bool `json:"ForceHttps,omitnil" name:"ForceHttps"`

	// 四层匹配的目的端口
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil" name:"Headers"`
}

func (r *CreateCloudNativeAPIGatewayRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceID")
	delete(f, "RouteName")
	delete(f, "Methods")
	delete(f, "Hosts")
	delete(f, "Paths")
	delete(f, "Protocols")
	delete(f, "PreserveHost")
	delete(f, "HttpsRedirectStatusCode")
	delete(f, "StripPath")
	delete(f, "ForceHttps")
	delete(f, "DestinationPorts")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRouteResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayRouteResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayRouteResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCloudNativeAPIGatewayServerGroupResult struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayServiceRateLimitRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

type CreateCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

func (r *CreateCloudNativeAPIGatewayServiceRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayServiceRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "LimitDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayServiceRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayServiceRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayServiceRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayServiceRateLimitResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayServiceRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayServiceRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayServiceRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 请求协议：
	// - https
	// - http
	// - tcp
	// - udp
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 请求路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 服务配置信息
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil" name:"UpstreamInfo"`
}

type CreateCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 请求协议：
	// - https
	// - http
	// - tcp
	// - udp
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 请求路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 服务配置信息
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil" name:"UpstreamInfo"`
}

func (r *CreateCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Protocol")
	delete(f, "Path")
	delete(f, "Timeout")
	delete(f, "Retries")
	delete(f, "UpstreamType")
	delete(f, "UpstreamInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayServiceResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEngineRequestParams struct {
	// 引擎类型。参考值：
	// - zookeeper
	// - nacos
	// - consul
	// - apollo
	// - eureka
	// - polaris
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 引擎的开源版本。每种引擎支持的开源版本不同，请参考产品文档或者控制台购买页
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`

	// 引擎的产品版本。参考值：
	// - STANDARD： 标准版
	// 
	// 引擎各版本及可选择的规格、节点数说明：
	// apollo - STANDARD版本
	// 规格列表：1C2G、2C4G、4C8G、8C16G、16C32G
	// 节点数：1，2，3，4，5
	// 
	// eureka - STANDARD版本
	// 规格列表：1C2G、2C4G、4C8G、8C16G、16C32G
	// 节点数：3，4，5
	// 
	// polarismesh - STANDARD版本
	// 规格列表：NUM50、NUM100、NUM200、NUM500、NUM1000、NUM5000、NUM10000、NUM50000
	// 
	// 兼容原spec-xxxxxx形式的规格ID
	EngineProductVersion *string `json:"EngineProductVersion,omitnil" name:"EngineProductVersion"`

	// 引擎所在地域。参考值说明：
	// 中国区 参考值：
	// - ap-guangzhou：广州
	// - ap-beijing：北京
	// - ap-chengdu：成都
	// - ap-chongqing：重庆
	// - ap-nanjing：南京
	// - ap-shanghai：上海
	// - ap-hongkong：香港
	// - ap-taipei：台北
	// 亚太区 参考值：
	// - ap-jakarta：雅加达
	// - ap-singapore：新加坡
	// 北美区 参考值
	// - na-toronto：多伦多
	// 金融专区 参考值
	// - ap-beijing-fsi：北京金融
	// - ap-shanghai-fsi：上海金融
	// - ap-shenzhen-fsi：深圳金融
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// 引擎名称。参考值：
	// - eurek-test
	EngineName *string `json:"EngineName,omitnil" name:"EngineName"`

	// 付费类型。参考值：
	// - 0：后付费
	// - 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil" name:"TradeType"`

	// 引擎的节点规格 ID。参见EngineProductVersion字段说明
	EngineResourceSpec *string `json:"EngineResourceSpec,omitnil" name:"EngineResourceSpec"`

	// 引擎的节点数量。参见EngineProductVersion字段说明
	EngineNodeNum *int64 `json:"EngineNodeNum,omitnil" name:"EngineNodeNum"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - vpc-conz6aix
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - subnet-ahde9me9
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Apollo 环境配置参数列表。参数说明：
	// 如果创建Apollo类型，此参数为必填的环境信息列表，最多可选4个环境。环境信息参数说明：
	// - Name：环境名。参考值：prod, dev, fat, uat
	// - EngineResourceSpec：环境内引擎的节点规格ID。参见EngineProductVersion参数说明
	// - EngineNodeNum：环境内引擎的节点数量。参见EngineProductVersion参数说明，其中prod环境支持的节点数为2，3，4，5
	// - StorageCapacity：配置存储空间大小，以GB为单位，步长为5.参考值：35
	// - VpcId：VPC ID。参考值：vpc-conz6aix
	// - SubnetId：子网 ID。参考值：subnet-ahde9me9
	ApolloEnvParams []*ApolloEnvParam `json:"ApolloEnvParams,omitnil" name:"ApolloEnvParams"`

	// 引擎的标签列表。用户自定义的key/value形式，无参考值
	EngineTags []*InstanceTagInfo `json:"EngineTags,omitnil" name:"EngineTags"`

	// 引擎的初始帐号信息。可设置参数：
	// - Name：控制台初始用户名
	// - Password：控制台初始密码
	// - Token：引擎接口的管理员 Token
	EngineAdmin *EngineAdmin `json:"EngineAdmin,omitnil" name:"EngineAdmin"`

	// 预付费时长，以月为单位
	PrepaidPeriod *int64 `json:"PrepaidPeriod,omitnil" name:"PrepaidPeriod"`

	// 自动续费标记，仅预付费使用。参考值：
	// - 0：不自动续费
	// - 1：自动续费
	PrepaidRenewFlag *int64 `json:"PrepaidRenewFlag,omitnil" name:"PrepaidRenewFlag"`

	// 跨地域部署的引擎地域配置详情
	EngineRegionInfos []*EngineRegionInfo `json:"EngineRegionInfos,omitnil" name:"EngineRegionInfos"`
}

type CreateEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎类型。参考值：
	// - zookeeper
	// - nacos
	// - consul
	// - apollo
	// - eureka
	// - polaris
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 引擎的开源版本。每种引擎支持的开源版本不同，请参考产品文档或者控制台购买页
	EngineVersion *string `json:"EngineVersion,omitnil" name:"EngineVersion"`

	// 引擎的产品版本。参考值：
	// - STANDARD： 标准版
	// 
	// 引擎各版本及可选择的规格、节点数说明：
	// apollo - STANDARD版本
	// 规格列表：1C2G、2C4G、4C8G、8C16G、16C32G
	// 节点数：1，2，3，4，5
	// 
	// eureka - STANDARD版本
	// 规格列表：1C2G、2C4G、4C8G、8C16G、16C32G
	// 节点数：3，4，5
	// 
	// polarismesh - STANDARD版本
	// 规格列表：NUM50、NUM100、NUM200、NUM500、NUM1000、NUM5000、NUM10000、NUM50000
	// 
	// 兼容原spec-xxxxxx形式的规格ID
	EngineProductVersion *string `json:"EngineProductVersion,omitnil" name:"EngineProductVersion"`

	// 引擎所在地域。参考值说明：
	// 中国区 参考值：
	// - ap-guangzhou：广州
	// - ap-beijing：北京
	// - ap-chengdu：成都
	// - ap-chongqing：重庆
	// - ap-nanjing：南京
	// - ap-shanghai：上海
	// - ap-hongkong：香港
	// - ap-taipei：台北
	// 亚太区 参考值：
	// - ap-jakarta：雅加达
	// - ap-singapore：新加坡
	// 北美区 参考值
	// - na-toronto：多伦多
	// 金融专区 参考值
	// - ap-beijing-fsi：北京金融
	// - ap-shanghai-fsi：上海金融
	// - ap-shenzhen-fsi：深圳金融
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// 引擎名称。参考值：
	// - eurek-test
	EngineName *string `json:"EngineName,omitnil" name:"EngineName"`

	// 付费类型。参考值：
	// - 0：后付费
	// - 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil" name:"TradeType"`

	// 引擎的节点规格 ID。参见EngineProductVersion字段说明
	EngineResourceSpec *string `json:"EngineResourceSpec,omitnil" name:"EngineResourceSpec"`

	// 引擎的节点数量。参见EngineProductVersion字段说明
	EngineNodeNum *int64 `json:"EngineNodeNum,omitnil" name:"EngineNodeNum"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - vpc-conz6aix
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - subnet-ahde9me9
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Apollo 环境配置参数列表。参数说明：
	// 如果创建Apollo类型，此参数为必填的环境信息列表，最多可选4个环境。环境信息参数说明：
	// - Name：环境名。参考值：prod, dev, fat, uat
	// - EngineResourceSpec：环境内引擎的节点规格ID。参见EngineProductVersion参数说明
	// - EngineNodeNum：环境内引擎的节点数量。参见EngineProductVersion参数说明，其中prod环境支持的节点数为2，3，4，5
	// - StorageCapacity：配置存储空间大小，以GB为单位，步长为5.参考值：35
	// - VpcId：VPC ID。参考值：vpc-conz6aix
	// - SubnetId：子网 ID。参考值：subnet-ahde9me9
	ApolloEnvParams []*ApolloEnvParam `json:"ApolloEnvParams,omitnil" name:"ApolloEnvParams"`

	// 引擎的标签列表。用户自定义的key/value形式，无参考值
	EngineTags []*InstanceTagInfo `json:"EngineTags,omitnil" name:"EngineTags"`

	// 引擎的初始帐号信息。可设置参数：
	// - Name：控制台初始用户名
	// - Password：控制台初始密码
	// - Token：引擎接口的管理员 Token
	EngineAdmin *EngineAdmin `json:"EngineAdmin,omitnil" name:"EngineAdmin"`

	// 预付费时长，以月为单位
	PrepaidPeriod *int64 `json:"PrepaidPeriod,omitnil" name:"PrepaidPeriod"`

	// 自动续费标记，仅预付费使用。参考值：
	// - 0：不自动续费
	// - 1：自动续费
	PrepaidRenewFlag *int64 `json:"PrepaidRenewFlag,omitnil" name:"PrepaidRenewFlag"`

	// 跨地域部署的引擎地域配置详情
	EngineRegionInfos []*EngineRegionInfo `json:"EngineRegionInfos,omitnil" name:"EngineRegionInfos"`
}

func (r *CreateEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EngineType")
	delete(f, "EngineVersion")
	delete(f, "EngineProductVersion")
	delete(f, "EngineRegion")
	delete(f, "EngineName")
	delete(f, "TradeType")
	delete(f, "EngineResourceSpec")
	delete(f, "EngineNodeNum")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ApolloEnvParams")
	delete(f, "EngineTags")
	delete(f, "EngineAdmin")
	delete(f, "PrepaidPeriod")
	delete(f, "PrepaidRenewFlag")
	delete(f, "EngineRegionInfos")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEngineResponseParams struct {
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEngineResponse struct {
	*tchttp.BaseResponse
	Response *CreateEngineResponseParams `json:"Response"`
}

func (r *CreateEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNativeGatewayServerGroupRequestParams struct {
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 节点配置
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 公网带宽信息
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 公网配置。
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil" name:"InternetConfig"`
}

type CreateNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 节点配置
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 公网带宽信息
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 公网配置。
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil" name:"InternetConfig"`
}

func (r *CreateNativeGatewayServerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNativeGatewayServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "NodeConfig")
	delete(f, "SubnetId")
	delete(f, "Description")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "InternetConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNativeGatewayServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNativeGatewayServerGroupResponseParams struct {
	// 网关分组创建信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *CreateCloudNativeAPIGatewayServerGroupResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateNativeGatewayServerGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateNativeGatewayServerGroupResponseParams `json:"Response"`
}

func (r *CreateNativeGatewayServerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNativeGatewayServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayCanaryRuleRequestParams struct {
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`
}

type DeleteCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`
}

func (r *DeleteCloudNativeAPIGatewayCanaryRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayCanaryRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceId")
	delete(f, "Priority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayCanaryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayCanaryRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayCanaryRuleResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayCanaryRuleResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayCanaryRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayCanaryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayCertificateRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteCloudNativeAPIGatewayCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayCertificateResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayCertificateResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayCertificateResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 是否删除实例关联的 CLS 日志主题。
	DeleteClsTopic *bool `json:"DeleteClsTopic,omitnil" name:"DeleteClsTopic"`
}

type DeleteCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 是否删除实例关联的 CLS 日志主题。
	DeleteClsTopic *bool `json:"DeleteClsTopic,omitnil" name:"DeleteClsTopic"`
}

func (r *DeleteCloudNativeAPIGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "DeleteClsTopic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayResponseParams struct {
	// 删除云原生API网关实例响应结果。
	Result *DeleteCloudNativeAPIGatewayResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCloudNativeAPIGatewayResult struct {
	// 云原生网关ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 云原生网关状态。
	Status *string `json:"Status,omitnil" name:"Status"`
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayRouteRateLimitRequestParams struct {
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DeleteCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayRouteRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayRouteRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayRouteRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayRouteRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayRouteRateLimitResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayRouteRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayRouteRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayRouteRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由的ID或名字，不支持名称“未命名”
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DeleteCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由的ID或名字，不支持名称“未命名”
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DeleteCloudNativeAPIGatewayRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayRouteResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayRouteResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayRouteResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayServiceRateLimitRequestParams struct {
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DeleteCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayServiceRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayServiceRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayServiceRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayServiceRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayServiceRateLimitResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayServiceRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayServiceRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayServiceRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名字，服务ID
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DeleteCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名字，服务ID
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DeleteCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayServiceResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEngineRequestParams struct {
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DeleteEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DeleteEngineRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEngineRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEngineResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteEngineResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEngineResponseParams `json:"Response"`
}

func (r *DeleteEngineResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEngineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNativeGatewayServerGroupRequestParams struct {
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

type DeleteNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *DeleteNativeGatewayServerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNativeGatewayServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNativeGatewayServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNativeGatewayServerGroupResponseParams struct {
	// 删除信息
	Result *DeleteNativeGatewayServerGroupResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteNativeGatewayServerGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNativeGatewayServerGroupResponseParams `json:"Response"`
}

func (r *DeleteNativeGatewayServerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNativeGatewayServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNativeGatewayServerGroupResult struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 删除状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayCanaryRulesRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeCloudNativeAPIGatewayCanaryRulesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeCloudNativeAPIGatewayCanaryRulesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayCanaryRulesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayCanaryRulesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayCanaryRulesResponseParams struct {
	// 灰度规则列表
	Result *CloudAPIGatewayCanaryRuleList `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayCanaryRulesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayCanaryRulesResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayCanaryRulesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayCanaryRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayCertificateDetailsRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeCloudNativeAPIGatewayCertificateDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayCertificateDetailsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayCertificateDetailsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayCertificateDetailsResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *KongCertificate `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayCertificateDetailsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayCertificateDetailsResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayCertificateDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayCertificateDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayCertificatesRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持BindDomain ，Name
	Filters []*ListFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持BindDomain ，Name
	Filters []*ListFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeCloudNativeAPIGatewayCertificatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayCertificatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayCertificatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayCertificatesResponseParams struct {
	// 无
	Result *KongCertificatesList `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayCertificatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayCertificatesResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayCertificatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayConfigRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 分组id，不填时为默认分组
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

type DescribeCloudNativeAPIGatewayConfigRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 分组id，不填时为默认分组
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

func (r *DescribeCloudNativeAPIGatewayConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayConfigResponseParams struct {
	// 获取云原生API网关响应结果。
	Result *DescribeCloudNativeAPIGatewayConfigResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayConfigResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayConfigResult struct {
	// 网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 分组网络配置列表。
	ConfigList []*CloudNativeAPIGatewayConfig `json:"ConfigList,omitnil" name:"ConfigList"`

	// 分组子网信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupSubnetId *string `json:"GroupSubnetId,omitnil" name:"GroupSubnetId"`

	// 分组VPC信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupVpcId *string `json:"GroupVpcId,omitnil" name:"GroupVpcId"`

	// 分组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayNodesRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 实例分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 翻页获取多少个
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页从第几个开始获取
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeCloudNativeAPIGatewayNodesRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 实例分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 翻页获取多少个
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页从第几个开始获取
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeCloudNativeAPIGatewayNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayNodesResponseParams struct {
	// 获取云原生网关节点列表结果。
	Result *DescribeCloudNativeAPIGatewayNodesResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayNodesResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayNodesResult struct {
	// 获取云原生API网关节点列表响应结果。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 云原生API网关节点列表。
	NodeList []*CloudNativeAPIGatewayNode `json:"NodeList,omitnil" name:"NodeList"`
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayPortsRequestParams struct {
	// 云原生API网关实例ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type DescribeCloudNativeAPIGatewayPortsRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewayPortsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayPortsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayPortsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayPortsResponseParams struct {
	// 云原生API网关实例协议端口列表响应结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *DescribeGatewayInstancePortResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayPortsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayPortsResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayPortsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayPortsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayRequestParams struct {
	// 云原生API网关实例ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

type DescribeCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`
}

func (r *DescribeCloudNativeAPIGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayResponseParams struct {
	// 获取云原生API网关基础信息响应结果。
	Result *DescribeCloudNativeAPIGatewayResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCloudNativeAPIGatewayResult struct {
	// 云原生API网关ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 云原生API网关状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 云原生API网关名。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 云原生API网关类型。
	Type *string `json:"Type,omitnil" name:"Type"`

	// 实例版本：
	// - 2.4.1
	// - 2.5.1
	GatewayVersion *string `json:"GatewayVersion,omitnil" name:"GatewayVersion"`

	// 云原生API网关节点信息。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`

	// 云原生API网关vpc配置。
	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitnil" name:"VpcConfig"`

	// 云原生API网关描述。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 云原生API网关创建时间。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 实例的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*InstanceTagInfo `json:"Tags,omitnil" name:"Tags"`

	// 是否开启 cls 日志
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableCls *bool `json:"EnableCls,omitnil" name:"EnableCls"`

	// 付费模式，0表示后付费，1预付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeType *int64 `json:"TradeType,omitnil" name:"TradeType"`

	// 实例版本，当前支持开发版、标准版、专业版【TRIAL、STANDARD、PROFESSIONAL】
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureVersion *string `json:"FeatureVersion,omitnil" name:"FeatureVersion"`

	// 公网出流量带宽，[1,2048]Mbps
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 自动续费标记，0表示默认状态(用户未设置，即初始状态)；
	// 1表示自动续费，2表示明确不自动续费(用户设置)，若业务无续费概念或无需自动续费，需要设置为0
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 到期时间，预付费时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurDeadline *string `json:"CurDeadline,omitnil" name:"CurDeadline"`

	// 隔离时间，实例隔离时使用
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitnil" name:"IsolateTime"`

	// 是否开启客户端公网。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableInternet *bool `json:"EnableInternet,omitnil" name:"EnableInternet"`

	// 实例实际的地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// Ingress class名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	IngressClassName *string `json:"IngressClassName,omitnil" name:"IngressClassName"`

	// 公网计费方式。可选取值 BANDWIDTH | TRAFFIC ，表示按带宽和按流量计费。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetPayMode *string `json:"InternetPayMode,omitnil" name:"InternetPayMode"`

	// 云原生API网关小版本号
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayMinorVersion *string `json:"GatewayMinorVersion,omitnil" name:"GatewayMinorVersion"`

	// 实例监听的端口信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstancePort *InstancePort `json:"InstancePort,omitnil" name:"InstancePort"`

	// 公网CLB默认类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancerType *string `json:"LoadBalancerType,omitnil" name:"LoadBalancerType"`

	// 公网IP地址列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil" name:"PublicIpAddresses"`
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayRouteRateLimitRequestParams struct {
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`
}

type DescribeCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`
}

func (r *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayRouteRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayRouteRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayRouteRateLimitResponseParams struct {
	// 获取云原生网关限流插件(路由)
	Result *CloudNativeAPIGatewayRateLimitDetail `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayRouteRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayRouteRateLimitResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayRouteRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayRouteRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayRoutesRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 服务的名字，精确匹配
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 路由的名字，精确匹配
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name, path, host, method, service, protocol
	Filters []*ListFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 服务的名字，精确匹配
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 路由的名字，精确匹配
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name, path, host, method, service, protocol
	Filters []*ListFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeCloudNativeAPIGatewayRoutesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayRoutesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "ServiceName")
	delete(f, "RouteName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayRoutesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayRoutesResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *KongServiceRouteList `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayRoutesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayRoutesResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayRoutesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayServiceRateLimitRequestParams struct {
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID。
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DescribeCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID。
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayServiceRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayServiceRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayServiceRateLimitResponseParams struct {
	// 获取云原生网关限流插件(服务)
	Result *CloudNativeAPIGatewayRateLimitDetail `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayServiceRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayServiceRateLimitResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayServiceRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayServiceRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayServicesRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列表 offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name,upstreamType
	Filters []*ListFilter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayServicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 列表 offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name,upstreamType
	Filters []*ListFilter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeCloudNativeAPIGatewayServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayServicesResponseParams struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *KongServices `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayServicesResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewaysRequestParams struct {
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 请求过滤参数，支持按照实例名称、ID和标签键值（Name、GatewayId、Tag）筛选
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeCloudNativeAPIGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 请求过滤参数，支持按照实例名称、ID和标签键值（Name、GatewayId、Tag）筛选
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeCloudNativeAPIGatewaysRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewaysRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewaysRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewaysResponseParams struct {
	// 获取云原生API网关实例列表响应结果。
	Result *ListCloudNativeAPIGatewayResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewaysResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewaysResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewaysResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewaysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayInstancePortResult struct {
	// 云原生API网关ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关实例协议端口列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayInstancePortList []*GatewayInstanceSchemeAndPorts `json:"GatewayInstancePortList,omitnil" name:"GatewayInstancePortList"`
}

type DescribeInstanceRegionInfo struct {
	// 引擎部署地域信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// 引擎在该地域的副本数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Replica *int64 `json:"Replica,omitnil" name:"Replica"`

	// 引擎在该地域的规格id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpecId *string `json:"SpecId,omitnil" name:"SpecId"`

	// 内网的网络信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetVpcInfos []*VpcInfo `json:"IntranetVpcInfos,omitnil" name:"IntranetVpcInfos"`

	// 是否开公网
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableClientInternet *bool `json:"EnableClientInternet,omitnil" name:"EnableClientInternet"`
}

// Predefined struct for user
type DescribeNacosReplicasRequestParams struct {
	// 引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeNacosReplicasRequest struct {
	*tchttp.BaseRequest
	
	// 引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeNacosReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNacosReplicasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNacosReplicasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNacosReplicasResponseParams struct {
	// 引擎实例副本信息
	Replicas []*NacosReplica `json:"Replicas,omitnil" name:"Replicas"`

	// 副本个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNacosReplicasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNacosReplicasResponseParams `json:"Response"`
}

func (r *DescribeNacosReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNacosReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNacosServerInterfacesRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeNacosServerInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeNacosServerInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNacosServerInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNacosServerInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNacosServerInterfacesResponseParams struct {
	// 接口总个数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 接口列表
	Content []*NacosServerInterface `json:"Content,omitnil" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNacosServerInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNacosServerInterfacesResponseParams `json:"Response"`
}

func (r *DescribeNacosServerInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNacosServerInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNativeGatewayServerGroupsRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 翻页从第几个开始获取
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页获取多少个
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeNativeGatewayServerGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 翻页从第几个开始获取
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// 翻页获取多少个
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeNativeGatewayServerGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNativeGatewayServerGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNativeGatewayServerGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNativeGatewayServerGroupsResponseParams struct {
	// 分组列表信息
	Result *NativeGatewayServerGroups `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeNativeGatewayServerGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNativeGatewayServerGroupsResponseParams `json:"Response"`
}

func (r *DescribeNativeGatewayServerGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNativeGatewayServerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOneCloudNativeAPIGatewayServiceRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名字，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`
}

type DescribeOneCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名字，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *DescribeOneCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOneCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeOneCloudNativeAPIGatewayServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOneCloudNativeAPIGatewayServiceResponseParams struct {
	// 无
	Result *KongServiceDetail `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeOneCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse
	Response *DescribeOneCloudNativeAPIGatewayServiceResponseParams `json:"Response"`
}

func (r *DescribeOneCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeOneCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstanceAccessAddressRequestParams struct {
	// 注册引擎实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 引擎其他组件名称（pushgateway、polaris-limiter）
	Workload *string `json:"Workload,omitnil" name:"Workload"`

	// 部署地域
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`
}

type DescribeSREInstanceAccessAddressRequest struct {
	*tchttp.BaseRequest
	
	// 注册引擎实例Id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 引擎其他组件名称（pushgateway、polaris-limiter）
	Workload *string `json:"Workload,omitnil" name:"Workload"`

	// 部署地域
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`
}

func (r *DescribeSREInstanceAccessAddressRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "Workload")
	delete(f, "EngineRegion")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSREInstanceAccessAddressRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstanceAccessAddressResponseParams struct {
	// 内网访问地址
	IntranetAddress *string `json:"IntranetAddress,omitnil" name:"IntranetAddress"`

	// 公网访问地址
	InternetAddress *string `json:"InternetAddress,omitnil" name:"InternetAddress"`

	// apollo多环境公网ip
	EnvAddressInfos []*EnvAddressInfo `json:"EnvAddressInfos,omitnil" name:"EnvAddressInfos"`

	// 控制台公网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleInternetAddress *string `json:"ConsoleInternetAddress,omitnil" name:"ConsoleInternetAddress"`

	// 控制台内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleIntranetAddress *string `json:"ConsoleIntranetAddress,omitnil" name:"ConsoleIntranetAddress"`

	// 客户端公网带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetBandWidth *int64 `json:"InternetBandWidth,omitnil" name:"InternetBandWidth"`

	// 控制台公网带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleInternetBandWidth *int64 `json:"ConsoleInternetBandWidth,omitnil" name:"ConsoleInternetBandWidth"`

	// 北极星限流server节点接入IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimiterAddressInfos []*PolarisLimiterAddress `json:"LimiterAddressInfos,omitnil" name:"LimiterAddressInfos"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSREInstanceAccessAddressResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSREInstanceAccessAddressResponseParams `json:"Response"`
}

func (r *DescribeSREInstanceAccessAddressResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstanceAccessAddressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstancesRequestParams struct {
	// 请求过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询类型
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 调用方来源
	QuerySource *string `json:"QuerySource,omitnil" name:"QuerySource"`
}

type DescribeSREInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 请求过滤参数
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// 查询类型
	QueryType *string `json:"QueryType,omitnil" name:"QueryType"`

	// 调用方来源
	QuerySource *string `json:"QuerySource,omitnil" name:"QuerySource"`
}

func (r *DescribeSREInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "QueryType")
	delete(f, "QuerySource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSREInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSREInstancesResponseParams struct {
	// 总数量
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 实例记录
	Content []*SREInstance `json:"Content,omitnil" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSREInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSREInstancesResponseParams `json:"Response"`
}

func (r *DescribeSREInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSREInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZookeeperReplicasRequestParams struct {
	// 注册引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeZookeeperReplicasRequest struct {
	*tchttp.BaseRequest
	
	// 注册引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeZookeeperReplicasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZookeeperReplicasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZookeeperReplicasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZookeeperReplicasResponseParams struct {
	// 注册引擎实例副本信息
	Replicas []*ZookeeperReplica `json:"Replicas,omitnil" name:"Replicas"`

	// 副本个数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeZookeeperReplicasResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZookeeperReplicasResponseParams `json:"Response"`
}

func (r *DescribeZookeeperReplicasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZookeeperReplicasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZookeeperServerInterfacesRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeZookeeperServerInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeZookeeperServerInterfacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZookeeperServerInterfacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeZookeeperServerInterfacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeZookeeperServerInterfacesResponseParams struct {
	// 接口总个数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 接口列表
	Content []*ZookeeperServerInterface `json:"Content,omitnil" name:"Content"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeZookeeperServerInterfacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeZookeeperServerInterfacesResponseParams `json:"Response"`
}

func (r *DescribeZookeeperServerInterfacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeZookeeperServerInterfacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineAdmin struct {
	// 控制台初始用户名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 控制台初始密码
	Password *string `json:"Password,omitnil" name:"Password"`

	// 引擎接口的管理员 Token
	Token *string `json:"Token,omitnil" name:"Token"`
}

type EngineRegionInfo struct {
	// 引擎节点所在地域
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// 此地域节点分配数量
	Replica *int64 `json:"Replica,omitnil" name:"Replica"`

	// 集群网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil" name:"VpcInfos"`
}

type EnvAddressInfo struct {
	// 环境名
	EnvName *string `json:"EnvName,omitnil" name:"EnvName"`

	// 是否开启config公网
	EnableConfigInternet *bool `json:"EnableConfigInternet,omitnil" name:"EnableConfigInternet"`

	// config公网ip
	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitnil" name:"ConfigInternetServiceIp"`

	// config内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigIntranetAddress *string `json:"ConfigIntranetAddress,omitnil" name:"ConfigIntranetAddress"`

	// 是否开启config内网clb
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableConfigIntranet *bool `json:"EnableConfigIntranet,omitnil" name:"EnableConfigIntranet"`

	// 客户端公网带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetBandWidth *int64 `json:"InternetBandWidth,omitnil" name:"InternetBandWidth"`
}

type EnvInfo struct {
	// 环境名称
	EnvName *string `json:"EnvName,omitnil" name:"EnvName"`

	// 环境对应的网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil" name:"VpcInfos"`

	// 云硬盘容量
	StorageCapacity *int64 `json:"StorageCapacity,omitnil" name:"StorageCapacity"`

	// 运行状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// Admin service 访问地址
	AdminServiceIp *string `json:"AdminServiceIp,omitnil" name:"AdminServiceIp"`

	// Config service访问地址
	ConfigServiceIp *string `json:"ConfigServiceIp,omitnil" name:"ConfigServiceIp"`

	// 是否开启config-server公网
	EnableConfigInternet *bool `json:"EnableConfigInternet,omitnil" name:"EnableConfigInternet"`

	// config-server公网访问地址
	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitnil" name:"ConfigInternetServiceIp"`

	// 规格ID
	SpecId *string `json:"SpecId,omitnil" name:"SpecId"`

	// 环境的节点数
	EnvReplica *int64 `json:"EnvReplica,omitnil" name:"EnvReplica"`

	// 环境运行的节点数
	RunningCount *int64 `json:"RunningCount,omitnil" name:"RunningCount"`

	// 环境别名
	AliasEnvName *string `json:"AliasEnvName,omitnil" name:"AliasEnvName"`

	// 环境描述
	EnvDesc *string `json:"EnvDesc,omitnil" name:"EnvDesc"`

	// 客户端带宽
	ClientBandWidth *uint64 `json:"ClientBandWidth,omitnil" name:"ClientBandWidth"`

	// 客户端内网开关
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableConfigIntranet *bool `json:"EnableConfigIntranet,omitnil" name:"EnableConfigIntranet"`
}

type ExternalRedis struct {
	// redis ip
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedisHost *string `json:"RedisHost,omitnil" name:"RedisHost"`

	// redis密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedisPassword *string `json:"RedisPassword,omitnil" name:"RedisPassword"`

	// redis端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedisPort *int64 `json:"RedisPort,omitnil" name:"RedisPort"`

	// 超时时间  ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	RedisTimeout *int64 `json:"RedisTimeout,omitnil" name:"RedisTimeout"`
}

type Filter struct {
	// 过滤参数名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 过滤参数值
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type GatewayInstanceSchemeAndPorts struct {
	// 端口协议，可选HTTP、HTTPS、TCP和UDP
	// 注意：此字段可能返回 null，表示取不到有效值。
	Scheme *string `json:"Scheme,omitnil" name:"Scheme"`

	// 端口列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	PortList []*uint64 `json:"PortList,omitnil" name:"PortList"`
}

type InstancePort struct {
	// 监听的 http 端口范围。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpPort *string `json:"HttpPort,omitnil" name:"HttpPort"`

	// 监听的 https 端口范围。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpsPort *string `json:"HttpsPort,omitnil" name:"HttpsPort"`
}

type InstanceTagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type InternetConfig struct {
	// 公网地址版本，可选："IPV4" | "IPV6" 。不填默认 IPV4 。
	InternetAddressVersion *string `json:"InternetAddressVersion,omitnil" name:"InternetAddressVersion"`

	// 公网付费类型，当前仅可选："BANDWIDTH"。不填默认为 "BANDWIDTH"
	InternetPayMode *string `json:"InternetPayMode,omitnil" name:"InternetPayMode"`

	// 公网带宽。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 负载均衡描述
	Description *string `json:"Description,omitnil" name:"Description"`

	// 负载均衡的规格类型，传 "SLA" 表示性能容量型，不传为共享型。
	SlaType *string `json:"SlaType,omitnil" name:"SlaType"`

	// 负载均衡是否多可用区
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil" name:"MultiZoneFlag"`

	// 主可用区
	MasterZoneId *string `json:"MasterZoneId,omitnil" name:"MasterZoneId"`

	// 备可用区
	SlaveZoneId *string `json:"SlaveZoneId,omitnil" name:"SlaveZoneId"`
}

type KVMapping struct {
	// key
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// value
	// 注意：此字段可能返回 null，表示取不到有效值。
	Value *string `json:"Value,omitnil" name:"Value"`
}

type KVPair struct {
	// 键
	Key *string `json:"Key,omitnil" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil" name:"Value"`
}

type KongCertificate struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Cert *KongCertificatesPreview `json:"Cert,omitnil" name:"Cert"`
}

type KongCertificatesList struct {
	// 证书列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Total *int64 `json:"Total,omitnil" name:"Total"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificatesList []*KongCertificatesPreview `json:"CertificatesList,omitnil" name:"CertificatesList"`

	// 证书列表总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Pages is deprecated.
	Pages *int64 `json:"Pages,omitnil" name:"Pages"`
}

type KongCertificatesPreview struct {
	// 证书名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitnil" name:"Id"`

	// 绑定的域名
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindDomains []*string `json:"BindDomains,omitnil" name:"BindDomains"`

	// 证书状态：expired(已过期)
	//                    active(生效中)
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 证书pem格式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Crt *string `json:"Crt,omitnil" name:"Crt"`

	// 证书私钥
	// 注意：此字段可能返回 null，表示取不到有效值。
	Key *string `json:"Key,omitnil" name:"Key"`

	// 证书过期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// 证书上传时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 证书签发时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IssueTime *string `json:"IssueTime,omitnil" name:"IssueTime"`

	// 证书来源：native(kong自定义证书)
	//                     ssl(ssl平台证书)
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertSource *string `json:"CertSource,omitnil" name:"CertSource"`

	// ssl平台证书Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertId *string `json:"CertId,omitnil" name:"CertId"`
}

type KongRoutePreview struct {
	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil" name:"ID"`

	// 服务名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Methods []*string `json:"Methods,omitnil" name:"Methods"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paths []*string `json:"Paths,omitnil" name:"Paths"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocols []*string `json:"Protocols,omitnil" name:"Protocols"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreserveHost *bool `json:"PreserveHost,omitnil" name:"PreserveHost"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil" name:"HttpsRedirectStatusCode"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	StripPath *bool `json:"StripPath,omitnil" name:"StripPath"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 是否开启了强制HTTPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ForceHttps is deprecated.
	ForceHttps *bool `json:"ForceHttps,omitnil" name:"ForceHttps"`

	// 服务名
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceID *string `json:"ServiceID,omitnil" name:"ServiceID"`

	// 目的端口
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil" name:"DestinationPorts"`

	// 路由的Headers
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*KVMapping `json:"Headers,omitnil" name:"Headers"`
}

type KongServiceDetail struct {
	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil" name:"ID"`

	// 服务名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 后端协议
	// 注意：此字段可能返回 null，表示取不到有效值。
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 后端路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`

	// 后端延时，单位ms
	// 注意：此字段可能返回 null，表示取不到有效值。
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`

	// 重试次数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Retries *int64 `json:"Retries,omitnil" name:"Retries"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 后端配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil" name:"UpstreamInfo"`

	// 后端类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamType *string `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 是否可编辑
	// 注意：此字段可能返回 null，表示取不到有效值。
	Editable *bool `json:"Editable,omitnil" name:"Editable"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`
}

type KongServicePreview struct {
	// 服务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ID *string `json:"ID,omitnil" name:"ID"`

	// 服务名字
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// 后端配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil" name:"UpstreamInfo"`

	// 后端类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpstreamType *string `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// 是否可编辑
	// 注意：此字段可能返回 null，表示取不到有效值。
	Editable *bool `json:"Editable,omitnil" name:"Editable"`

	// 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil" name:"Path"`
}

type KongServiceRouteList struct {
	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	RouteList []*KongRoutePreview `json:"RouteList,omitnil" name:"RouteList"`

	// 总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type KongServices struct {
	// kong实例的服务列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceList []*KongServicePreview `json:"ServiceList,omitnil" name:"ServiceList"`

	// 列表总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`
}

type KongTarget struct {
	// Host
	Host *string `json:"Host,omitnil" name:"Host"`

	// 端口
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 权重
	Weight *int64 `json:"Weight,omitnil" name:"Weight"`

	// 健康状态
	Health *string `json:"Health,omitnil" name:"Health"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil" name:"CreatedTime"`

	// Target的来源
	Source *string `json:"Source,omitnil" name:"Source"`
}

type KongUpstreamInfo struct {
	// IP或域名
	Host *string `json:"Host,omitnil" name:"Host"`

	// 端口
	Port *int64 `json:"Port,omitnil" name:"Port"`

	// 服务来源ID
	SourceID *string `json:"SourceID,omitnil" name:"SourceID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// 服务（注册中心或Kubernetes中的服务）名字
	ServiceName *string `json:"ServiceName,omitnil" name:"ServiceName"`

	// 服务后端类型是IPList时提供
	Targets []*KongTarget `json:"Targets,omitnil" name:"Targets"`

	// 服务来源类型
	SourceType *string `json:"SourceType,omitnil" name:"SourceType"`

	// SCF函数类型
	ScfType *string `json:"ScfType,omitnil" name:"ScfType"`

	// SCF函数命名空间
	ScfNamespace *string `json:"ScfNamespace,omitnil" name:"ScfNamespace"`

	// SCF函数名
	ScfLambdaName *string `json:"ScfLambdaName,omitnil" name:"ScfLambdaName"`

	// SCF函数版本
	ScfLambdaQualifier *string `json:"ScfLambdaQualifier,omitnil" name:"ScfLambdaQualifier"`

	// 冷启动时间，单位秒
	SlowStart *int64 `json:"SlowStart,omitnil" name:"SlowStart"`

	// 负载均衡算法，默认为 round-robin，还支持 least-connections，consisten_hashing
	Algorithm *string `json:"Algorithm,omitnil" name:"Algorithm"`

	// CVM弹性伸缩组ID
	AutoScalingGroupID *string `json:"AutoScalingGroupID,omitnil" name:"AutoScalingGroupID"`

	// CVM弹性伸缩组端口
	AutoScalingCvmPort *uint64 `json:"AutoScalingCvmPort,omitnil" name:"AutoScalingCvmPort"`

	// CVM弹性伸缩组使用的CVM TAT命令状态
	AutoScalingTatCmdStatus *string `json:"AutoScalingTatCmdStatus,omitnil" name:"AutoScalingTatCmdStatus"`

	// CVM弹性伸缩组生命周期挂钩状态
	AutoScalingHookStatus *string `json:"AutoScalingHookStatus,omitnil" name:"AutoScalingHookStatus"`

	// 服务来源的名字
	SourceName *string `json:"SourceName,omitnil" name:"SourceName"`

	// 精确的服务来源类型，新建服务来源时候传入的类型
	RealSourceType *string `json:"RealSourceType,omitnil" name:"RealSourceType"`
}

type ListCloudNativeAPIGatewayResult struct {
	// 总数。
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 云原生API网关实例列表。
	GatewayList []*DescribeCloudNativeAPIGatewayResult `json:"GatewayList,omitnil" name:"GatewayList"`
}

type ListFilter struct {
	// 过滤字段
	Key *string `json:"Key,omitnil" name:"Key"`

	// 过滤值
	Value *string `json:"Value,omitnil" name:"Value"`
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayCanaryRuleRequestParams struct {
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 优先级，同一个服务的灰度规则优先级是唯一的
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil" name:"CanaryRule"`
}

type ModifyCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil" name:"ServiceId"`

	// 优先级，同一个服务的灰度规则优先级是唯一的
	Priority *int64 `json:"Priority,omitnil" name:"Priority"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil" name:"CanaryRule"`
}

func (r *ModifyCloudNativeAPIGatewayCanaryRuleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayCanaryRuleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceId")
	delete(f, "Priority")
	delete(f, "CanaryRule")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayCanaryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayCanaryRuleResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayCanaryRuleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayCanaryRuleResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayCanaryRuleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayCanaryRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否开启 CLS 日志。暂时取值只能是 true，即只能从关闭状态变成开启状态。
	EnableCls *bool `json:"EnableCls,omitnil" name:"EnableCls"`

	// 公网计费模式。可选取值 BANDWIDTH | TRAFFIC ，表示按带宽和按流量计费。
	InternetPayMode *string `json:"InternetPayMode,omitnil" name:"InternetPayMode"`
}

type ModifyCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil" name:"Description"`

	// 是否开启 CLS 日志。暂时取值只能是 true，即只能从关闭状态变成开启状态。
	EnableCls *bool `json:"EnableCls,omitnil" name:"EnableCls"`

	// 公网计费模式。可选取值 BANDWIDTH | TRAFFIC ，表示按带宽和按流量计费。
	InternetPayMode *string `json:"InternetPayMode,omitnil" name:"InternetPayMode"`
}

func (r *ModifyCloudNativeAPIGatewayRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Description")
	delete(f, "EnableCls")
	delete(f, "InternetPayMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayRouteRateLimitRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

type ModifyCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

func (r *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayRouteRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	delete(f, "LimitDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayRouteRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayRouteRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayRouteRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayRouteRateLimitResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayRouteRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayRouteRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayRouteRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil" name:"ServiceID"`

	// 路由的ID，实例级别唯一
	RouteID *string `json:"RouteID,omitnil" name:"RouteID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 路由的方法，其中方法可选值：
	// - GET
	// - POST
	// - DELETE
	// - PUT
	// - OPTIONS
	// - PATCH
	// - HEAD
	// - ANY
	// - TRACE
	// - COPY
	// - MOVE
	// - PROPFIND
	// - PROPPATCH
	// - MKCOL
	// - LOCK
	// - UNLOCK
	Methods []*string `json:"Methods,omitnil" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil" name:"StripPath"`

	// 是否开启强制HTTPS
	//
	// Deprecated: ForceHttps is deprecated.
	ForceHttps *bool `json:"ForceHttps,omitnil" name:"ForceHttps"`

	// 四层匹配的目的端口	
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil" name:"Headers"`
}

type ModifyCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil" name:"ServiceID"`

	// 路由的ID，实例级别唯一
	RouteID *string `json:"RouteID,omitnil" name:"RouteID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil" name:"RouteName"`

	// 路由的方法，其中方法可选值：
	// - GET
	// - POST
	// - DELETE
	// - PUT
	// - OPTIONS
	// - PATCH
	// - HEAD
	// - ANY
	// - TRACE
	// - COPY
	// - MOVE
	// - PROPFIND
	// - PROPPATCH
	// - MKCOL
	// - LOCK
	// - UNLOCK
	Methods []*string `json:"Methods,omitnil" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil" name:"StripPath"`

	// 是否开启强制HTTPS
	ForceHttps *bool `json:"ForceHttps,omitnil" name:"ForceHttps"`

	// 四层匹配的目的端口	
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil" name:"Headers"`
}

func (r *ModifyCloudNativeAPIGatewayRouteRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayRouteRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceID")
	delete(f, "RouteID")
	delete(f, "RouteName")
	delete(f, "Methods")
	delete(f, "Hosts")
	delete(f, "Paths")
	delete(f, "Protocols")
	delete(f, "PreserveHost")
	delete(f, "HttpsRedirectStatusCode")
	delete(f, "StripPath")
	delete(f, "ForceHttps")
	delete(f, "DestinationPorts")
	delete(f, "Headers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayRouteResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayRouteResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayRouteResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayRouteResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayServiceRateLimitRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

type ModifyCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil" name:"LimitDetail"`
}

func (r *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayServiceRateLimitRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "LimitDetail")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayServiceRateLimitRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayServiceRateLimitResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayServiceRateLimitResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayServiceRateLimitResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayServiceRateLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayServiceRateLimitResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayServiceRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 请求协议： 
	// - https 
	// - http 
	// - tcp 
	// - udp
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 请求路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 服务配置
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil" name:"UpstreamInfo"`

	// 服务ID
	ID *string `json:"ID,omitnil" name:"ID"`
}

type ModifyCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 请求协议： 
	// - https 
	// - http 
	// - tcp 
	// - udp
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// 请求路径
	Path *string `json:"Path,omitnil" name:"Path"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil" name:"UpstreamType"`

	// 服务配置
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil" name:"UpstreamInfo"`

	// 服务ID
	ID *string `json:"ID,omitnil" name:"ID"`
}

func (r *ModifyCloudNativeAPIGatewayServiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayServiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Protocol")
	delete(f, "Path")
	delete(f, "Timeout")
	delete(f, "Retries")
	delete(f, "UpstreamType")
	delete(f, "UpstreamInfo")
	delete(f, "ID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayServiceResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayServiceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayServiceResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayServiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNativeGatewayServerGroupRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组 id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil" name:"Description"`
}

type ModifyNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组 id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil" name:"Description"`
}

func (r *ModifyNativeGatewayServerGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNativeGatewayServerGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "Name")
	delete(f, "Description")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNativeGatewayServerGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNativeGatewayServerGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyNativeGatewayServerGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNativeGatewayServerGroupResponseParams `json:"Response"`
}

func (r *ModifyNativeGatewayServerGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNativeGatewayServerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NacosReplica struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 角色
	Role *string `json:"Role,omitnil" name:"Role"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// VPC ID	
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type NacosServerInterface struct {
	// 接口名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interface *string `json:"Interface,omitnil" name:"Interface"`
}

type NativeGatewayServerGroup struct {
	// 云原生网关分组唯一id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 分组名
	Name *string `json:"Name,omitnil" name:"Name"`

	// 描述信息
	Description *string `json:"Description,omitnil" name:"Description"`

	// 节点规格、节点数信息
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`

	// 网关分组状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 是否是默认分组。
	// 0：否。
	// 1：是。
	IsFirstGroup *int64 `json:"IsFirstGroup,omitnil" name:"IsFirstGroup"`

	// 关联策略信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	BindingStrategy *CloudNativeAPIGatewayStrategy `json:"BindingStrategy,omitnil" name:"BindingStrategy"`

	// 网关实例 id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 带宽
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil" name:"InternetMaxBandwidthOut"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// 子网id
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds *string `json:"SubnetIds,omitnil" name:"SubnetIds"`
}

type NativeGatewayServerGroups struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// 分组信息数组。
	GatewayGroupList []*NativeGatewayServerGroup `json:"GatewayGroupList,omitnil" name:"GatewayGroupList"`
}

type NetworkAccessControl struct {
	// 访问模式：Whitelist|Blacklist
	Mode *string `json:"Mode,omitnil" name:"Mode"`

	// 白名单列表
	CidrWhiteList []*string `json:"CidrWhiteList,omitnil" name:"CidrWhiteList"`

	// 黑名单列表
	CidrBlackList []*string `json:"CidrBlackList,omitnil" name:"CidrBlackList"`
}

type PolarisLimiterAddress struct {
	// VPC接入IP列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetAddress *string `json:"IntranetAddress,omitnil" name:"IntranetAddress"`
}

type QpsThreshold struct {
	// qps阈值控制维度,包含:second、minute、hour、day、month、year
	Unit *string `json:"Unit,omitnil" name:"Unit"`

	// 阈值
	Max *int64 `json:"Max,omitnil" name:"Max"`
}

type RateLimitResponse struct {
	// 自定义响应体
	// 注意：此字段可能返回 null，表示取不到有效值。
	Body *string `json:"Body,omitnil" name:"Body"`

	// Headers
	// 注意：此字段可能返回 null，表示取不到有效值。
	Headers []*KVMapping `json:"Headers,omitnil" name:"Headers"`

	// http状态码
	// 注意：此字段可能返回 null，表示取不到有效值。
	HttpStatus *int64 `json:"HttpStatus,omitnil" name:"HttpStatus"`
}

type SREInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 版本号
	Edition *string `json:"Edition,omitnil" name:"Edition"`

	// 状态, 枚举值:creating/create_fail/running/updating/update_fail/restarting/restart_fail/destroying/destroy_fail
	Status *string `json:"Status,omitnil" name:"Status"`

	// 规格ID
	SpecId *string `json:"SpecId,omitnil" name:"SpecId"`

	// 副本数
	Replica *int64 `json:"Replica,omitnil" name:"Replica"`

	// 类型
	Type *string `json:"Type,omitnil" name:"Type"`

	// Vpc iD
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetIds []*string `json:"SubnetIds,omitnil" name:"SubnetIds"`

	// 是否开启持久化存储
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableStorage *bool `json:"EnableStorage,omitnil" name:"EnableStorage"`

	// 数据存储方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageType *string `json:"StorageType,omitnil" name:"StorageType"`

	// 云硬盘容量
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageCapacity *int64 `json:"StorageCapacity,omitnil" name:"StorageCapacity"`

	// 计费方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paymode *string `json:"Paymode,omitnil" name:"Paymode"`

	// EKS集群的ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	EKSClusterID *string `json:"EKSClusterID,omitnil" name:"EKSClusterID"`

	// 集群创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// 环境配置信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil" name:"EnvInfos"`

	// 引擎所在的区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// 注册引擎是否开启公网
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableInternet *bool `json:"EnableInternet,omitnil" name:"EnableInternet"`

	// 私有网络列表信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil" name:"VpcInfos"`

	// 服务治理相关信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	ServiceGovernanceInfos []*ServiceGovernanceInfo `json:"ServiceGovernanceInfos,omitnil" name:"ServiceGovernanceInfos"`

	// 实例的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*KVPair `json:"Tags,omitnil" name:"Tags"`

	// 引擎实例是否开启控制台公网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableConsoleInternet *bool `json:"EnableConsoleInternet,omitnil" name:"EnableConsoleInternet"`

	// 引擎实例是否开启控制台内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableConsoleIntranet *bool `json:"EnableConsoleIntranet,omitnil" name:"EnableConsoleIntranet"`

	// 引擎实例是否展示参数配置页面
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConfigInfoVisible *bool `json:"ConfigInfoVisible,omitnil" name:"ConfigInfoVisible"`

	// 引擎实例控制台默认密码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsoleDefaultPwd *string `json:"ConsoleDefaultPwd,omitnil" name:"ConsoleDefaultPwd"`

	// 交易付费类型，0后付费/1预付费
	// 注意：此字段可能返回 null，表示取不到有效值。
	TradeType *int64 `json:"TradeType,omitnil" name:"TradeType"`

	// 自动续费标记：0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费
	// 注意：此字段可能返回 null，表示取不到有效值。
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// 预付费到期时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurDeadline *string `json:"CurDeadline,omitnil" name:"CurDeadline"`

	// 隔离开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsolateTime *string `json:"IsolateTime,omitnil" name:"IsolateTime"`

	// 实例地域相关的描述信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionInfos []*DescribeInstanceRegionInfo `json:"RegionInfos,omitnil" name:"RegionInfos"`

	// 所在EKS环境，分为common和yunti
	// 注意：此字段可能返回 null，表示取不到有效值。
	EKSType *string `json:"EKSType,omitnil" name:"EKSType"`

	// 引擎的产品版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	FeatureVersion *string `json:"FeatureVersion,omitnil" name:"FeatureVersion"`

	// 引擎实例是否开启客户端内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnableClientIntranet *bool `json:"EnableClientIntranet,omitnil" name:"EnableClientIntranet"`
}

type ServiceGovernanceInfo struct {
	// 引擎所在的地域
	EngineRegion *string `json:"EngineRegion,omitnil" name:"EngineRegion"`

	// 服务治理引擎绑定的kubernetes集群信息
	BoundK8SInfos []*BoundK8SInfo `json:"BoundK8SInfos,omitnil" name:"BoundK8SInfos"`

	// 服务治理引擎绑定的网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil" name:"VpcInfos"`

	// 当前实例鉴权是否开启
	AuthOpen *bool `json:"AuthOpen,omitnil" name:"AuthOpen"`

	// 该实例支持的功能，鉴权就是 Auth
	Features []*string `json:"Features,omitnil" name:"Features"`

	// 主账户名默认为 polaris，该值为主账户的默认密码
	MainPassword *string `json:"MainPassword,omitnil" name:"MainPassword"`

	// 服务治理pushgateway引擎绑定的网络信息
	PgwVpcInfos []*VpcInfo `json:"PgwVpcInfos,omitnil" name:"PgwVpcInfos"`

	// 服务治理限流server引擎绑定的网络信息
	LimiterVpcInfos []*VpcInfo `json:"LimiterVpcInfos,omitnil" name:"LimiterVpcInfos"`
}

// Predefined struct for user
type UpdateCloudNativeAPIGatewayCertificateInfoRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 证书id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 绑定的域名列表
	BindDomains []*string `json:"BindDomains,omitnil" name:"BindDomains"`

	// 证书名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

type UpdateCloudNativeAPIGatewayCertificateInfoRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 证书id
	Id *string `json:"Id,omitnil" name:"Id"`

	// 绑定的域名列表
	BindDomains []*string `json:"BindDomains,omitnil" name:"BindDomains"`

	// 证书名称
	Name *string `json:"Name,omitnil" name:"Name"`
}

func (r *UpdateCloudNativeAPIGatewayCertificateInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCloudNativeAPIGatewayCertificateInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	delete(f, "BindDomains")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCloudNativeAPIGatewayCertificateInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCloudNativeAPIGatewayCertificateInfoResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateCloudNativeAPIGatewayCertificateInfoResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCloudNativeAPIGatewayCertificateInfoResponseParams `json:"Response"`
}

func (r *UpdateCloudNativeAPIGatewayCertificateInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCloudNativeAPIGatewayCertificateInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateCloudNativeAPIGatewayResult struct {
	// 云原生API网关ID。
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 云原生网关状态。
	Status *string `json:"Status,omitnil" name:"Status"`

	// 任务ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskId *string `json:"TaskId,omitnil" name:"TaskId"`
}

// Predefined struct for user
type UpdateCloudNativeAPIGatewaySpecRequestParams struct {
	// 云原生API网关实例ID。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 网关分组节点规格配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`
}

type UpdateCloudNativeAPIGatewaySpecRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// 网关分组节点规格配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil" name:"NodeConfig"`
}

func (r *UpdateCloudNativeAPIGatewaySpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCloudNativeAPIGatewaySpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "NodeConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateCloudNativeAPIGatewaySpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCloudNativeAPIGatewaySpecResponseParams struct {
	// 更新云原生API网关实例规格的响应结果。
	Result *UpdateCloudNativeAPIGatewayResult `json:"Result,omitnil" name:"Result"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateCloudNativeAPIGatewaySpecResponse struct {
	*tchttp.BaseResponse
	Response *UpdateCloudNativeAPIGatewaySpecResponseParams `json:"Response"`
}

func (r *UpdateCloudNativeAPIGatewaySpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateCloudNativeAPIGatewaySpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEngineInternetAccessRequestParams struct {
	// 引擎ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 引擎类型
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 是否开启客户端公网访问，true开 false关
	EnableClientInternetAccess *bool `json:"EnableClientInternetAccess,omitnil" name:"EnableClientInternetAccess"`
}

type UpdateEngineInternetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// 引擎类型
	EngineType *string `json:"EngineType,omitnil" name:"EngineType"`

	// 是否开启客户端公网访问，true开 false关
	EnableClientInternetAccess *bool `json:"EnableClientInternetAccess,omitnil" name:"EnableClientInternetAccess"`
}

func (r *UpdateEngineInternetAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEngineInternetAccessRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EngineType")
	delete(f, "EnableClientInternetAccess")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateEngineInternetAccessRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateEngineInternetAccessResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UpdateEngineInternetAccessResponse struct {
	*tchttp.BaseResponse
	Response *UpdateEngineInternetAccessResponseParams `json:"Response"`
}

func (r *UpdateEngineInternetAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateEngineInternetAccessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {
	// Vpc Id
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 内网访问地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IntranetAddress *string `json:"IntranetAddress,omitnil" name:"IntranetAddress"`
}

type ZookeeperReplica struct {
	// 名称
	Name *string `json:"Name,omitnil" name:"Name"`

	// 角色
	Role *string `json:"Role,omitnil" name:"Role"`

	// 状态
	Status *string `json:"Status,omitnil" name:"Status"`

	// 子网ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitnil" name:"Zone"`

	// 可用区ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// 别名
	// 注意：此字段可能返回 null，表示取不到有效值。
	AliasName *string `json:"AliasName,omitnil" name:"AliasName"`

	// VPC ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`
}

type ZookeeperServerInterface struct {
	// 接口名
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interface *string `json:"Interface,omitnil" name:"Interface"`
}