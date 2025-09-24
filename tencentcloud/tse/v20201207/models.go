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

package v20201207

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type AccurateQpsThreshold struct {
	// qps阈值控制维度,包含:second、minute、hour、day、month、year
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 全局配置ID
	GlobalConfigId *string `json:"GlobalConfigId,omitnil,omitempty" name:"GlobalConfigId"`
}

type ApolloEnvParam struct {
	// 环境名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 环境内引擎的节点规格 ID
	// -1C2G
	// -2C4G
	// 兼容原spec-xxxxxx形式的规格ID
	EngineResourceSpec *string `json:"EngineResourceSpec,omitnil,omitempty" name:"EngineResourceSpec"`

	// 环境内引擎的节点数量
	EngineNodeNum *int64 `json:"EngineNodeNum,omitnil,omitempty" name:"EngineNodeNum"`

	// 配置存储空间大小，以GB为单位
	StorageCapacity *int64 `json:"StorageCapacity,omitnil,omitempty" name:"StorageCapacity"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为 ConfigServer 的访问地址
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 环境描述
	EnvDesc *string `json:"EnvDesc,omitnil,omitempty" name:"EnvDesc"`
}

type AutoScalerBehavior struct {
	// 扩容行为配置
	ScaleUp *AutoScalerRules `json:"ScaleUp,omitnil,omitempty" name:"ScaleUp"`

	// 缩容行为配置
	ScaleDown *AutoScalerRules `json:"ScaleDown,omitnil,omitempty" name:"ScaleDown"`
}

type AutoScalerPolicy struct {
	// 类型，Pods
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 数量
	Value *int64 `json:"Value,omitnil,omitempty" name:"Value"`

	// 扩容周期
	PeriodSeconds *int64 `json:"PeriodSeconds,omitnil,omitempty" name:"PeriodSeconds"`
}

type AutoScalerRules struct {
	// 稳定窗口时间，扩容时默认0，缩容时默认300
	StabilizationWindowSeconds *int64 `json:"StabilizationWindowSeconds,omitnil,omitempty" name:"StabilizationWindowSeconds"`

	// 选择策略依据
	SelectPolicy *string `json:"SelectPolicy,omitnil,omitempty" name:"SelectPolicy"`

	// 扩缩容策略
	Policies []*AutoScalerPolicy `json:"Policies,omitnil,omitempty" name:"Policies"`
}

// Predefined struct for user
type BindAutoScalerResourceStrategyToGroupsRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 网关分组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type BindAutoScalerResourceStrategyToGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 网关分组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *BindAutoScalerResourceStrategyToGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoScalerResourceStrategyToGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StrategyId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAutoScalerResourceStrategyToGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoScalerResourceStrategyToGroupsResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type BindAutoScalerResourceStrategyToGroupsResponse struct {
	*tchttp.BaseResponse
	Response *BindAutoScalerResourceStrategyToGroupsResponseParams `json:"Response"`
}

func (r *BindAutoScalerResourceStrategyToGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoScalerResourceStrategyToGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BoundK8SInfo struct {
	// 绑定的kubernetes集群ID
	BoundClusterId *string `json:"BoundClusterId,omitnil,omitempty" name:"BoundClusterId"`

	// 绑定的kubernetes的集群类型，分tke和eks两种
	BoundClusterType *string `json:"BoundClusterType,omitnil,omitempty" name:"BoundClusterType"`

	// 服务同步模式，all为全量同步，demand为按需同步
	SyncMode *string `json:"SyncMode,omitnil,omitempty" name:"SyncMode"`

	// 绑定的kubernetes集群所在地域
	BindRegion *string `json:"BindRegion,omitnil,omitempty" name:"BindRegion"`
}

type CLBMultiRegion struct {
	// 是否启用多可用区
	CLBMultiZoneFlag *bool `json:"CLBMultiZoneFlag,omitnil,omitempty" name:"CLBMultiZoneFlag"`

	// 主可用区信息
	CLBMasterZone *string `json:"CLBMasterZone,omitnil,omitempty" name:"CLBMasterZone"`

	// 备可用区信息
	CLBSlaveZone *string `json:"CLBSlaveZone,omitnil,omitempty" name:"CLBSlaveZone"`
}

type CanaryPriorityRule struct {
	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil,omitempty" name:"CanaryRule"`
}

type CertificateInfo struct {
	// 唯一id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

// Predefined struct for user
type CloseWafProtectionRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	//  防护资源的类型。
	// - Global  实例
	// - Service  服务
	// - Route  路由
	// - Object  对象
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当资源类型 Type 是 Service 或 Route 的时候，传入的服务或路由的列表
	List []*string `json:"List,omitnil,omitempty" name:"List"`
}

type CloseWafProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	//  防护资源的类型。
	// - Global  实例
	// - Service  服务
	// - Route  路由
	// - Object  对象
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当资源类型 Type 是 Service 或 Route 的时候，传入的服务或路由的列表
	List []*string `json:"List,omitnil,omitempty" name:"List"`
}

func (r *CloseWafProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWafProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Type")
	delete(f, "List")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CloseWafProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CloseWafProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CloseWafProtectionResponse struct {
	*tchttp.BaseResponse
	Response *CloseWafProtectionResponseParams `json:"Response"`
}

func (r *CloseWafProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CloseWafProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloudAPIGatewayCanaryRuleList struct {
	// 灰度规则
	CanaryRuleList []*CloudNativeAPIGatewayCanaryRule `json:"CanaryRuleList,omitnil,omitempty" name:"CanaryRuleList"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type CloudNativeAPIGatewayBalancedService struct {
	// 服务 ID，作为入参时，必填
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 服务名称，作为入参时，无意义
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// Upstream 名称，作为入参时，无意义
	UpstreamName *string `json:"UpstreamName,omitnil,omitempty" name:"UpstreamName"`

	// 百分比，10 即 10%，范围0-100
	Percent *float64 `json:"Percent,omitnil,omitempty" name:"Percent"`
}

type CloudNativeAPIGatewayCanaryRule struct {
	// 优先级，值范围为 0 到 100；值越大，优先级越高；不同规则间优先级不可重复
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 是否启用规则
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 参数匹配条件
	ConditionList []*CloudNativeAPIGatewayCanaryRuleCondition `json:"ConditionList,omitnil,omitempty" name:"ConditionList"`

	// 服务的流量百分比配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	BalancedServiceList []*CloudNativeAPIGatewayBalancedService `json:"BalancedServiceList,omitnil,omitempty" name:"BalancedServiceList"`

	// 归属服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 归属服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 灰度规则类别
	// Standard｜Lane
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 全链路灰度策略多个条件之间的匹配方式，与AND，或OR
	MatchType *string `json:"MatchType,omitnil,omitempty" name:"MatchType"`

	// 泳道组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 泳道组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 泳道ID
	LaneId *string `json:"LaneId,omitnil,omitempty" name:"LaneId"`

	// 泳道名称
	LaneName *string `json:"LaneName,omitnil,omitempty" name:"LaneName"`

	// 泳道匹配规则：严格STRICT｜宽松PERMISSIVE
	MatchMode *string `json:"MatchMode,omitnil,omitempty" name:"MatchMode"`

	// 泳道标签
	LaneTag *string `json:"LaneTag,omitnil,omitempty" name:"LaneTag"`
}

type CloudNativeAPIGatewayCanaryRuleCondition struct {
	// 条件类型，支持 path, method, query, header, cookie, body 和 system。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 参数名
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 操作符，支持 "le", "eq", "lt", "ne", "ge", "gt", "regex", "exists", "in", "not in",  "prefix" ,"exact", "regex" 等
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// 目标参数值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 分隔符，当 Operator 为 in 或者 not in 时生效。支持值为英文逗号，英文分号，空格，换行符。
	Delimiter *string `json:"Delimiter,omitnil,omitempty" name:"Delimiter"`

	// 全局配置 Id
	GlobalConfigId *string `json:"GlobalConfigId,omitnil,omitempty" name:"GlobalConfigId"`

	// 全局配置名称
	GlobalConfigName *string `json:"GlobalConfigName,omitnil,omitempty" name:"GlobalConfigName"`
}

type CloudNativeAPIGatewayConfig struct {
	// 控制台类型。
	ConsoleType *string `json:"ConsoleType,omitnil,omitempty" name:"ConsoleType"`

	// HTTP链接地址。
	HttpUrl *string `json:"HttpUrl,omitnil,omitempty" name:"HttpUrl"`

	// HTTPS链接地址。
	HttpsUrl *string `json:"HttpsUrl,omitnil,omitempty" name:"HttpsUrl"`

	// 网络类型, Open|Internal。
	NetType *string `json:"NetType,omitnil,omitempty" name:"NetType"`

	// 管理员用户名。
	AdminUser *string `json:"AdminUser,omitnil,omitempty" name:"AdminUser"`

	// 管理员密码。
	AdminPassword *string `json:"AdminPassword,omitnil,omitempty" name:"AdminPassword"`

	// 网络状态, Open|Closed|Updating
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 网络访问策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccessControl *NetworkAccessControl `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`

	// 内网子网 ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 内网VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 负载均衡的描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 负载均衡的规格类型
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// clb规格名称
	SlaName *string `json:"SlaName,omitnil,omitempty" name:"SlaName"`

	// clb vip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 带宽
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否多可用区
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// 主可用区
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// 备可用区
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`

	// 主可用区名称
	MasterZoneName *string `json:"MasterZoneName,omitnil,omitempty" name:"MasterZoneName"`

	// 备可用区名称
	SlaveZoneName *string `json:"SlaveZoneName,omitnil,omitempty" name:"SlaveZoneName"`

	// 网络 id
	NetworkId *string `json:"NetworkId,omitnil,omitempty" name:"NetworkId"`

	// 是否为新 ipv6 CLB
	IPV6FullChain *bool `json:"IPV6FullChain,omitnil,omitempty" name:"IPV6FullChain"`

	// 负载均衡个性化配置内容
	CustomizedConfigContent *string `json:"CustomizedConfigContent,omitnil,omitempty" name:"CustomizedConfigContent"`
}

type CloudNativeAPIGatewayNode struct {
	// 云原生网关节点 id
	NodeId *string `json:"NodeId,omitnil,omitempty" name:"NodeId"`

	// 节点 ip
	NodeIp *string `json:"NodeIp,omitnil,omitempty" name:"NodeIp"`

	// Zone id
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// Zone
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 节点权重
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 是否默认权重
	IsDefaultWeight *bool `json:"IsDefaultWeight,omitnil,omitempty" name:"IsDefaultWeight"`
}

type CloudNativeAPIGatewayNodeConfig struct {
	// 节点配置, 1c2g|2c4g|4c8g|8c16g。
	Specification *string `json:"Specification,omitnil,omitempty" name:"Specification"`

	// 节点数量，2-9。
	Number *int64 `json:"Number,omitnil,omitempty" name:"Number"`
}

type CloudNativeAPIGatewayRateLimitDetail struct {
	// 插件启用状态
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// qps阈值
	QpsThresholds []*QpsThreshold `json:"QpsThresholds,omitnil,omitempty" name:"QpsThresholds"`

	// 需要进行流量控制的请求路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 需要进行流量控制的请求头Key
	Header *string `json:"Header,omitnil,omitempty" name:"Header"`

	// 限流依据
	// ip service consumer credential path header
	LimitBy *string `json:"LimitBy,omitnil,omitempty" name:"LimitBy"`

	// 外部redis配置
	ExternalRedis *ExternalRedis `json:"ExternalRedis,omitnil,omitempty" name:"ExternalRedis"`

	// 计数器策略 
	// local 单机
	// redis  默认redis
	// external_redis 外部redis
	Policy *string `json:"Policy,omitnil,omitempty" name:"Policy"`

	// 响应配置，响应策略为text
	RateLimitResponse *RateLimitResponse `json:"RateLimitResponse,omitnil,omitempty" name:"RateLimitResponse"`

	// 请求转发地址
	RateLimitResponseUrl *string `json:"RateLimitResponseUrl,omitnil,omitempty" name:"RateLimitResponseUrl"`

	// 响应策略
	// url请求转发
	// text 响应配置
	// default 直接返回
	ResponseType *string `json:"ResponseType,omitnil,omitempty" name:"ResponseType"`

	// 是否隐藏限流客户端响应头
	HideClientHeaders *bool `json:"HideClientHeaders,omitnil,omitempty" name:"HideClientHeaders"`

	// 排队时间
	LineUpTime *int64 `json:"LineUpTime,omitnil,omitempty" name:"LineUpTime"`

	// 是否开启请求排队
	IsDelay *bool `json:"IsDelay,omitnil,omitempty" name:"IsDelay"`

	// 基础限流
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicLimitQpsThresholds []*QpsThreshold `json:"BasicLimitQpsThresholds,omitnil,omitempty" name:"BasicLimitQpsThresholds"`

	// 参数限流的规则
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimitRules []*LimitRule `json:"LimitRules,omitnil,omitempty" name:"LimitRules"`
}

type CloudNativeAPIGatewayStrategy struct {
	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 更新时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 弹性伸缩配置
	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 定时伸缩配置
	CronConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronConfig,omitnil,omitempty" name:"CronConfig"`

	// 最大节点数
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: MaxReplicas is deprecated.
	MaxReplicas *uint64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`
}

type CloudNativeAPIGatewayStrategyAutoScalerConfig struct {
	// 最大副本数
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 指标列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metrics []*CloudNativeAPIGatewayStrategyAutoScalerConfigMetric `json:"Metrics,omitnil,omitempty" name:"Metrics"`

	// 是否开启指标伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Enabled is deprecated.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: CreateTime is deprecated.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModifyTime is deprecated.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 弹性策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: StrategyId is deprecated.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 指标配置ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: AutoScalerId is deprecated.
	AutoScalerId *string `json:"AutoScalerId,omitnil,omitempty" name:"AutoScalerId"`

	// 指标伸缩行为配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Behavior *AutoScalerBehavior `json:"Behavior,omitnil,omitempty" name:"Behavior"`
}

type CloudNativeAPIGatewayStrategyAutoScalerConfigMetric struct {
	// 指标类型
	// - Resource
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 指标资源名称
	// - cpu
	// - memory
	ResourceName *string `json:"ResourceName,omitnil,omitempty" name:"ResourceName"`

	// 指标目标类型，目前只支持百分比Utilization
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 指标目标值
	TargetValue *int64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`
}

type CloudNativeAPIGatewayStrategyBindingGroupInfo struct {
	// 网关分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 节点配置
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// 绑定时间
	BindTime *string `json:"BindTime,omitnil,omitempty" name:"BindTime"`

	// 网关分组名称
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 绑定状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type CloudNativeAPIGatewayStrategyCronScalerConfig struct {
	// 是否开启定时伸缩
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Enabled is deprecated.
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 定时伸缩配置参数列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Params []*CloudNativeAPIGatewayStrategyCronScalerConfigParam `json:"Params,omitnil,omitempty" name:"Params"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: CreateTime is deprecated.
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ModifyTime is deprecated.
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 弹性策略ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: StrategyId is deprecated.
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type CloudNativeAPIGatewayStrategyCronScalerConfigParam struct {
	// 定时伸缩周期
	Period *string `json:"Period,omitnil,omitempty" name:"Period"`

	// 定时伸缩开始时间
	StartAt *string `json:"StartAt,omitnil,omitempty" name:"StartAt"`

	// 定时伸缩目标节点数，不超过指标伸缩中定义的最大节点数
	TargetReplicas *int64 `json:"TargetReplicas,omitnil,omitempty" name:"TargetReplicas"`

	// 定时伸缩cron表达式，无需输入
	Crontab *string `json:"Crontab,omitnil,omitempty" name:"Crontab"`
}

type CloudNativeAPIGatewayVpcConfig struct {
	// 私有网络ID。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`
}

type ConfigFile struct {
	// 配置文件id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置文件组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 配置文件格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 配置文件注释
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 配置文件状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 配置文件标签数组
	Tags []*ConfigFileTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 配置文件创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 配置文件创建者
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 配置文件修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 配置文件修改者
	ModifyBy *string `json:"ModifyBy,omitnil,omitempty" name:"ModifyBy"`

	// 配置文件发布时间
	ReleaseTime *string `json:"ReleaseTime,omitnil,omitempty" name:"ReleaseTime"`

	// 配置文件发布者
	ReleaseBy *string `json:"ReleaseBy,omitnil,omitempty" name:"ReleaseBy"`

	// 配置文件类型
	ConfigFileSupportedClient *int64 `json:"ConfigFileSupportedClient,omitnil,omitempty" name:"ConfigFileSupportedClient"`

	// 配置文件持久化
	ConfigFilePersistent *ConfigFilePersistent `json:"ConfigFilePersistent,omitnil,omitempty" name:"ConfigFilePersistent"`

	// 是否开启加密算法
	Encrypted *bool `json:"Encrypted,omitnil,omitempty" name:"Encrypted"`

	// 加密算法
	EncryptAlgo *string `json:"EncryptAlgo,omitnil,omitempty" name:"EncryptAlgo"`
}

type ConfigFileGroup struct {
	// 配置文件组id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 配置文件组名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 备注
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 创建者
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 修改者
	ModifyBy *string `json:"ModifyBy,omitnil,omitempty" name:"ModifyBy"`

	// 文件数
	FileCount *uint64 `json:"FileCount,omitnil,omitempty" name:"FileCount"`

	// 关联用户，link_users
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 组id，link_groups
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// remove_link_users
	RemoveUserIds []*string `json:"RemoveUserIds,omitnil,omitempty" name:"RemoveUserIds"`

	// remove_link_groups
	RemoveGroupIds []*string `json:"RemoveGroupIds,omitnil,omitempty" name:"RemoveGroupIds"`

	// 是否可编辑
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 归属者
	Owner *string `json:"Owner,omitnil,omitempty" name:"Owner"`

	// 部门
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// 业务
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// 配置文件组标签
	ConfigFileGroupTags []*ConfigFileGroupTag `json:"ConfigFileGroupTags,omitnil,omitempty" name:"ConfigFileGroupTags"`
}

type ConfigFileGroupTag struct {
	// key-value 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// key-value 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ConfigFilePersistent struct {
	// 文件编码
	Encoding *string `json:"Encoding,omitnil,omitempty" name:"Encoding"`

	// 文件下发路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 文件后置命令
	PostCmd *string `json:"PostCmd,omitnil,omitempty" name:"PostCmd"`
}

type ConfigFilePublishInfo struct {
	// 发布名称
	ReleaseName *string `json:"ReleaseName,omitnil,omitempty" name:"ReleaseName"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 发布组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 创建者
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 修改者
	ModifyBy *string `json:"ModifyBy,omitnil,omitempty" name:"ModifyBy"`

	// 标签
	Tags []*ConfigFileTag `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type ConfigFileRelease struct {
	// 配置文件发布id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 配置文件发布名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件发布命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置文件发布组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件发布文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件发布内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 配置文件发布注释
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 配置文件发布Md5
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 配置文件发布版本
	Version *uint64 `json:"Version,omitnil,omitempty" name:"Version"`

	// 配置文件发布创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 配置文件发布创建者
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 配置文件发布修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 配置文件发布修改者
	ModifyBy *string `json:"ModifyBy,omitnil,omitempty" name:"ModifyBy"`

	// 发布描述
	ReleaseDescription *string `json:"ReleaseDescription,omitnil,omitempty" name:"ReleaseDescription"`

	// 是否生效
	Active *bool `json:"Active,omitnil,omitempty" name:"Active"`

	// 格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 配置文件ID
	ConfigFileId *string `json:"ConfigFileId,omitnil,omitempty" name:"ConfigFileId"`

	// 配置文件类型
	ConfigFileSupportedClient *int64 `json:"ConfigFileSupportedClient,omitnil,omitempty" name:"ConfigFileSupportedClient"`

	// 配置文件持久化
	ConfigFilePersistent *ConfigFilePersistent `json:"ConfigFilePersistent,omitnil,omitempty" name:"ConfigFilePersistent"`
}

type ConfigFileReleaseDeletion struct {
	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 发布版本
	ReleaseVersion *string `json:"ReleaseVersion,omitnil,omitempty" name:"ReleaseVersion"`

	// 配置发布ID
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`
}

type ConfigFileReleaseHistory struct {
	// 配置文件发布历史记录id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 配置文件发布历史名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件发布历史命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置文件发布历史组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件发布历史名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件发布历史内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 配置文件发布历史格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 配置文件发布历史注释
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 配置文件发布历史Md5
	Md5 *string `json:"Md5,omitnil,omitempty" name:"Md5"`

	// 配置文件发布历史类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 配置文件发布历史状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 配置文件发布历史标签组
	Tags []*ConfigFileTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 配置文件发布创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 配置文件发布创建者
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 配置文件发布修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 配置文件发布修改者
	ModifyBy *string `json:"ModifyBy,omitnil,omitempty" name:"ModifyBy"`

	// 发布描述
	ReleaseDescription *string `json:"ReleaseDescription,omitnil,omitempty" name:"ReleaseDescription"`

	// 原因，用于失败时原因展示
	ReleaseReason *string `json:"ReleaseReason,omitnil,omitempty" name:"ReleaseReason"`

	// 配置文件类型
	ConfigFileSupportedClient *int64 `json:"ConfigFileSupportedClient,omitnil,omitempty" name:"ConfigFileSupportedClient"`

	// 配置文件持久化
	ConfigFilePersistent *ConfigFilePersistent `json:"ConfigFilePersistent,omitnil,omitempty" name:"ConfigFilePersistent"`
}

type ConfigFileTag struct {
	// key-value 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// key-value 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type ConfigFileTemplate struct {
	// 配置文件模板id
	Id *uint64 `json:"Id,omitnil,omitempty" name:"Id"`

	// 配置文件模板名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件模板内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 配置文件模板格式
	Format *string `json:"Format,omitnil,omitempty" name:"Format"`

	// 配置文件模板注释
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 配置文件模板创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 配置文件模板创建者
	CreateBy *string `json:"CreateBy,omitnil,omitempty" name:"CreateBy"`

	// 配置文件模板修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 配置文件模板修改者
	ModifyBy *string `json:"ModifyBy,omitnil,omitempty" name:"ModifyBy"`
}

// Predefined struct for user
type CreateAutoScalerResourceStrategyRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 指标伸缩配置
	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 定时伸缩配置列表
	//
	// Deprecated: CronScalerConfig is deprecated.
	CronScalerConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronScalerConfig,omitnil,omitempty" name:"CronScalerConfig"`

	// 最大节点数
	//
	// Deprecated: MaxReplicas is deprecated.
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 定时伸缩配置
	CronConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronConfig,omitnil,omitempty" name:"CronConfig"`
}

type CreateAutoScalerResourceStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 指标伸缩配置
	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 定时伸缩配置列表
	CronScalerConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronScalerConfig,omitnil,omitempty" name:"CronScalerConfig"`

	// 最大节点数
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 定时伸缩配置
	CronConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronConfig,omitnil,omitempty" name:"CronConfig"`
}

func (r *CreateAutoScalerResourceStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalerResourceStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StrategyName")
	delete(f, "Description")
	delete(f, "Config")
	delete(f, "CronScalerConfig")
	delete(f, "MaxReplicas")
	delete(f, "CronConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoScalerResourceStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalerResourceStrategyResponseParams struct {
	// 是否成功
	//
	// Deprecated: Result is deprecated.
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 策略Id
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoScalerResourceStrategyResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoScalerResourceStrategyResponseParams `json:"Response"`
}

func (r *CreateAutoScalerResourceStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalerResourceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayCanaryRuleRequestParams struct {
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil,omitempty" name:"CanaryRule"`

	// 灰度规则配置列表，如果配置了此参数，将以此参数为准，忽略CanaryRule参数
	CanaryRuleList []*CloudNativeAPIGatewayCanaryRule `json:"CanaryRuleList,omitnil,omitempty" name:"CanaryRuleList"`
}

type CreateCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil,omitempty" name:"CanaryRule"`

	// 灰度规则配置列表，如果配置了此参数，将以此参数为准，忽略CanaryRule参数
	CanaryRuleList []*CloudNativeAPIGatewayCanaryRule `json:"CanaryRuleList,omitnil,omitempty" name:"CanaryRuleList"`
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
	delete(f, "CanaryRuleList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayCanaryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayCanaryRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 绑定的域名
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// ssl平台证书 Id
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证书私钥
	//
	// Deprecated: Key is deprecated.
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 证书pem格式
	//
	// Deprecated: Crt is deprecated.
	Crt *string `json:"Crt,omitnil,omitempty" name:"Crt"`
}

type CreateCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 绑定的域名
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// ssl平台证书 Id
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证书私钥
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 证书pem格式
	Crt *string `json:"Crt,omitnil,omitempty" name:"Crt"`
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
	Result *CertificateInfo `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateCloudNativeAPIGatewayPublicNetworkRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 公网负载均衡配置。
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil,omitempty" name:"InternetConfig"`
}

type CreateCloudNativeAPIGatewayPublicNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id。
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 公网负载均衡配置。
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil,omitempty" name:"InternetConfig"`
}

func (r *CreateCloudNativeAPIGatewayPublicNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayPublicNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "InternetConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayPublicNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayPublicNetworkResponseParams struct {
	// 返回结果
	Result *CreatePublicNetworkResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateCloudNativeAPIGatewayPublicNetworkResponse struct {
	*tchttp.BaseResponse
	Response *CreateCloudNativeAPIGatewayPublicNetworkResponseParams `json:"Response"`
}

func (r *CreateCloudNativeAPIGatewayPublicNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCloudNativeAPIGatewayPublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRequestParams struct {
	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云原生API网关类型, 目前只支持kong。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 云原生API网关版本。参考值：
	// - 2.4.1
	// - 2.5.1
	GatewayVersion *string `json:"GatewayVersion,omitnil,omitempty" name:"GatewayVersion"`

	// 云原生API网关节点配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// 云原生API网关vpc配置。
	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitnil,omitempty" name:"VpcConfig"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签列表
	Tags []*InstanceTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启 CLS 日志。默认值：fasle
	EnableCls *bool `json:"EnableCls,omitnil,omitempty" name:"EnableCls"`

	// 产品版本。参考值：
	// - TRIAL：开发版
	// - STANDARD：标准版 （默认值）
	// - PROFESSIONAL：专业版
	FeatureVersion *string `json:"FeatureVersion,omitnil,omitempty" name:"FeatureVersion"`

	// 公网出流量带宽，[1,2048]Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 实例实际的地域信息，默认值：ap-guangzhou
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// ingress Class名称
	IngressClassName *string `json:"IngressClassName,omitnil,omitempty" name:"IngressClassName"`

	// 付费类型。参考值：
	// 0：后付费（默认值）
	// 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil,omitempty" name:"TradeType"`

	// 公网相关配置
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil,omitempty" name:"InternetConfig"`

	// 关联的prometheus ID
	PromId *string `json:"PromId,omitnil,omitempty" name:"PromId"`
}

type CreateCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云原生API网关类型, 目前只支持kong。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 云原生API网关版本。参考值：
	// - 2.4.1
	// - 2.5.1
	GatewayVersion *string `json:"GatewayVersion,omitnil,omitempty" name:"GatewayVersion"`

	// 云原生API网关节点配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// 云原生API网关vpc配置。
	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitnil,omitempty" name:"VpcConfig"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 标签列表
	Tags []*InstanceTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启 CLS 日志。默认值：fasle
	EnableCls *bool `json:"EnableCls,omitnil,omitempty" name:"EnableCls"`

	// 产品版本。参考值：
	// - TRIAL：开发版
	// - STANDARD：标准版 （默认值）
	// - PROFESSIONAL：专业版
	FeatureVersion *string `json:"FeatureVersion,omitnil,omitempty" name:"FeatureVersion"`

	// 公网出流量带宽，[1,2048]Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 实例实际的地域信息，默认值：ap-guangzhou
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// ingress Class名称
	IngressClassName *string `json:"IngressClassName,omitnil,omitempty" name:"IngressClassName"`

	// 付费类型。参考值：
	// 0：后付费（默认值）
	// 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil,omitempty" name:"TradeType"`

	// 公网相关配置
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil,omitempty" name:"InternetConfig"`

	// 关联的prometheus ID
	PromId *string `json:"PromId,omitnil,omitempty" name:"PromId"`
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
	delete(f, "PromId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayResponseParams struct {
	// 创建云原生API网关实例响应结果。
	Result *CreateCloudNativeAPIGatewayResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 云原生网关状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRouteRateLimitRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
}

type CreateCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

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
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil,omitempty" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil,omitempty" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil,omitempty" name:"StripPath"`

	// 是否开启强制HTTPS
	//
	// Deprecated: ForceHttps is deprecated.
	ForceHttps *bool `json:"ForceHttps,omitnil,omitempty" name:"ForceHttps"`

	// 四层匹配的目的端口
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil,omitempty" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 是否缓存请求body，默认true
	RequestBuffering *bool `json:"RequestBuffering,omitnil,omitempty" name:"RequestBuffering"`

	// 是否缓存响应body，默认true
	ResponseBuffering *bool `json:"ResponseBuffering,omitnil,omitempty" name:"ResponseBuffering"`

	// 正则优先级
	RegexPriority *int64 `json:"RegexPriority,omitnil,omitempty" name:"RegexPriority"`
}

type CreateCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

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
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil,omitempty" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil,omitempty" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil,omitempty" name:"StripPath"`

	// 是否开启强制HTTPS
	ForceHttps *bool `json:"ForceHttps,omitnil,omitempty" name:"ForceHttps"`

	// 四层匹配的目的端口
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil,omitempty" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 是否缓存请求body，默认true
	RequestBuffering *bool `json:"RequestBuffering,omitnil,omitempty" name:"RequestBuffering"`

	// 是否缓存响应body，默认true
	ResponseBuffering *bool `json:"ResponseBuffering,omitnil,omitempty" name:"ResponseBuffering"`

	// 正则优先级
	RegexPriority *int64 `json:"RegexPriority,omitnil,omitempty" name:"RegexPriority"`
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
	delete(f, "RequestBuffering")
	delete(f, "ResponseBuffering")
	delete(f, "RegexPriority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayServiceRateLimitRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
}

type CreateCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 请求协议：
	// - https
	// - http
	// - tcp
	// - udp
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 服务配置信息
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// 请求路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type CreateCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 请求协议：
	// - https
	// - http
	// - tcp
	// - udp
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 服务配置信息
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// 请求路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
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
	delete(f, "Timeout")
	delete(f, "Retries")
	delete(f, "UpstreamType")
	delete(f, "UpstreamInfo")
	delete(f, "Path")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCloudNativeAPIGatewayServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCloudNativeAPIGatewayServiceResponseParams struct {
	// 网关服务创建结果
	Result *CreateGatewayServiceResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateConfigFileGroupRequestParams struct {
	// tse 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件组实体
	ConfigFileGroup *ConfigFileGroup `json:"ConfigFileGroup,omitnil,omitempty" name:"ConfigFileGroup"`
}

type CreateConfigFileGroupRequest struct {
	*tchttp.BaseRequest
	
	// tse 实例 id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件组实体
	ConfigFileGroup *ConfigFileGroup `json:"ConfigFileGroup,omitnil,omitempty" name:"ConfigFileGroup"`
}

func (r *CreateConfigFileGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigFileGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigFileGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigFileGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigFileGroupResponseParams struct {
	// 是否创建成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigFileGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigFileGroupResponseParams `json:"Response"`
}

func (r *CreateConfigFileGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigFileGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigFileRequestParams struct {
	// TSE 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件列表详情
	ConfigFile *ConfigFile `json:"ConfigFile,omitnil,omitempty" name:"ConfigFile"`
}

type CreateConfigFileRequest struct {
	*tchttp.BaseRequest
	
	// TSE 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件列表详情
	ConfigFile *ConfigFile `json:"ConfigFile,omitnil,omitempty" name:"ConfigFile"`
}

func (r *CreateConfigFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigFile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateConfigFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateConfigFileResponseParams struct {
	// 是否创建成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 创建的配置文件Id
	ConfigFileId *string `json:"ConfigFileId,omitnil,omitempty" name:"ConfigFileId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateConfigFileResponse struct {
	*tchttp.BaseResponse
	Response *CreateConfigFileResponseParams `json:"Response"`
}

func (r *CreateConfigFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateConfigFileResponse) FromJsonString(s string) error {
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
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 引擎的开源版本。每种引擎支持的开源版本不同，请参考产品文档或者控制台购买页
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 引擎的产品版本。参考值：
	// - STANDARD： 标准版
	// - PROFESSIONAL: 专业版（Zookeeper）/企业版（PolarisMesh）
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
	EngineProductVersion *string `json:"EngineProductVersion,omitnil,omitempty" name:"EngineProductVersion"`

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
	// - na-siliconvalley：硅谷
	// - na-ashburn: 弗吉尼亚
	// 金融专区 参考值
	// - ap-beijing-fsi：北京金融
	// - ap-shanghai-fsi：上海金融
	// - ap-shenzhen-fsi：深圳金融
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// 引擎名称。参考值：
	// - eurek-test
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 付费类型。参考值：
	// - 0：后付费
	// - 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil,omitempty" name:"TradeType"`

	// 引擎的节点规格 ID。参见EngineProductVersion字段说明
	EngineResourceSpec *string `json:"EngineResourceSpec,omitnil,omitempty" name:"EngineResourceSpec"`

	// 引擎的节点数量。参见EngineProductVersion字段说明
	EngineNodeNum *int64 `json:"EngineNodeNum,omitnil,omitempty" name:"EngineNodeNum"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - vpc-conz6aix
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - subnet-ahde9me9
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Apollo 环境配置参数列表。参数说明：
	// 如果创建Apollo类型，此参数为必填的环境信息列表，最多可选4个环境。环境信息参数说明：
	// - Name：环境名。参考值：prod, dev, fat, uat
	// - EngineResourceSpec：环境内引擎的节点规格ID。参见EngineProductVersion参数说明
	// - EngineNodeNum：环境内引擎的节点数量。参见EngineProductVersion参数说明，其中prod环境支持的节点数为2，3，4，5
	// - StorageCapacity：配置存储空间大小，以GB为单位，步长为5.参考值：35
	// - VpcId：VPC ID。参考值：vpc-conz6aix
	// - SubnetId：子网 ID。参考值：subnet-ahde9me9
	ApolloEnvParams []*ApolloEnvParam `json:"ApolloEnvParams,omitnil,omitempty" name:"ApolloEnvParams"`

	// 引擎的标签列表。用户自定义的key/value形式，无参考值
	EngineTags []*InstanceTagInfo `json:"EngineTags,omitnil,omitempty" name:"EngineTags"`

	// 引擎的初始账号信息。可设置参数：
	// - Name：控制台初始用户名
	// - Password：控制台初始密码
	// - Token：引擎接口的管理员 Token
	EngineAdmin *EngineAdmin `json:"EngineAdmin,omitnil,omitempty" name:"EngineAdmin"`

	// 预付费时长，以月为单位
	PrepaidPeriod *int64 `json:"PrepaidPeriod,omitnil,omitempty" name:"PrepaidPeriod"`

	// 自动续费标记，仅预付费使用。参考值：
	// - 0：不自动续费
	// - 1：自动续费
	PrepaidRenewFlag *int64 `json:"PrepaidRenewFlag,omitnil,omitempty" name:"PrepaidRenewFlag"`

	// 跨地域部署的引擎地域配置详情
	// zk标准版没有跨地域部署，请不要填写
	// zk专业版跨地域部署开启了固定Leader所在地域，需要满足以下条件
	// - 固定Leader所在地域当前仅支持跨两个地域
	// - leader地域的副本数必须是3/2 + 1，5/2+1，7/2+1，也就是 2，3，4
	EngineRegionInfos []*EngineRegionInfo `json:"EngineRegionInfos,omitnil,omitempty" name:"EngineRegionInfos"`

	// zk标准版请填CLOUD_PREMIUM，zk标准版无法选择磁盘类型和磁盘容量，默认为CLOUD_PREMIUM
	// zk专业版可以为：CLOUD_SSD,CLOUD_SSD_PLUS,CLOUD_PREMIUM
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// zk标准版请填50，zk标准版无法选择磁盘类型和磁盘容量，磁盘容量默认为50
	StorageCapacity *int64 `json:"StorageCapacity,omitnil,omitempty" name:"StorageCapacity"`

	// zk专业版至多有两个盘，且磁盘的容量在50-3200之间
	// 如果只有一个磁盘，storageCapacity与storageOption里面的capacity应该一致
	StorageOption []*StorageOption `json:"StorageOption,omitnil,omitempty" name:"StorageOption"`

	// ZK引擎实例，可用区分布约束，STRICT:强约束，PERMISSIVE: 弱约束
	AffinityConstraint *string `json:"AffinityConstraint,omitnil,omitempty" name:"AffinityConstraint"`

	// 指定zone id列表
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 地域特殊标签，用于区分相同地域，不通的业务属性
	EngineRegionTag *string `json:"EngineRegionTag,omitnil,omitempty" name:"EngineRegionTag"`
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
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 引擎的开源版本。每种引擎支持的开源版本不同，请参考产品文档或者控制台购买页
	EngineVersion *string `json:"EngineVersion,omitnil,omitempty" name:"EngineVersion"`

	// 引擎的产品版本。参考值：
	// - STANDARD： 标准版
	// - PROFESSIONAL: 专业版（Zookeeper）/企业版（PolarisMesh）
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
	EngineProductVersion *string `json:"EngineProductVersion,omitnil,omitempty" name:"EngineProductVersion"`

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
	// - na-siliconvalley：硅谷
	// - na-ashburn: 弗吉尼亚
	// 金融专区 参考值
	// - ap-beijing-fsi：北京金融
	// - ap-shanghai-fsi：上海金融
	// - ap-shenzhen-fsi：深圳金融
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// 引擎名称。参考值：
	// - eurek-test
	EngineName *string `json:"EngineName,omitnil,omitempty" name:"EngineName"`

	// 付费类型。参考值：
	// - 0：后付费
	// - 1：预付费（接口暂不支持创建预付费实例）
	TradeType *int64 `json:"TradeType,omitnil,omitempty" name:"TradeType"`

	// 引擎的节点规格 ID。参见EngineProductVersion字段说明
	EngineResourceSpec *string `json:"EngineResourceSpec,omitnil,omitempty" name:"EngineResourceSpec"`

	// 引擎的节点数量。参见EngineProductVersion字段说明
	EngineNodeNum *int64 `json:"EngineNodeNum,omitnil,omitempty" name:"EngineNodeNum"`

	// VPC ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - vpc-conz6aix
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网 ID。在 VPC 的子网内分配一个 IP 作为引擎的访问地址。参考值：
	// - subnet-ahde9me9
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// Apollo 环境配置参数列表。参数说明：
	// 如果创建Apollo类型，此参数为必填的环境信息列表，最多可选4个环境。环境信息参数说明：
	// - Name：环境名。参考值：prod, dev, fat, uat
	// - EngineResourceSpec：环境内引擎的节点规格ID。参见EngineProductVersion参数说明
	// - EngineNodeNum：环境内引擎的节点数量。参见EngineProductVersion参数说明，其中prod环境支持的节点数为2，3，4，5
	// - StorageCapacity：配置存储空间大小，以GB为单位，步长为5.参考值：35
	// - VpcId：VPC ID。参考值：vpc-conz6aix
	// - SubnetId：子网 ID。参考值：subnet-ahde9me9
	ApolloEnvParams []*ApolloEnvParam `json:"ApolloEnvParams,omitnil,omitempty" name:"ApolloEnvParams"`

	// 引擎的标签列表。用户自定义的key/value形式，无参考值
	EngineTags []*InstanceTagInfo `json:"EngineTags,omitnil,omitempty" name:"EngineTags"`

	// 引擎的初始账号信息。可设置参数：
	// - Name：控制台初始用户名
	// - Password：控制台初始密码
	// - Token：引擎接口的管理员 Token
	EngineAdmin *EngineAdmin `json:"EngineAdmin,omitnil,omitempty" name:"EngineAdmin"`

	// 预付费时长，以月为单位
	PrepaidPeriod *int64 `json:"PrepaidPeriod,omitnil,omitempty" name:"PrepaidPeriod"`

	// 自动续费标记，仅预付费使用。参考值：
	// - 0：不自动续费
	// - 1：自动续费
	PrepaidRenewFlag *int64 `json:"PrepaidRenewFlag,omitnil,omitempty" name:"PrepaidRenewFlag"`

	// 跨地域部署的引擎地域配置详情
	// zk标准版没有跨地域部署，请不要填写
	// zk专业版跨地域部署开启了固定Leader所在地域，需要满足以下条件
	// - 固定Leader所在地域当前仅支持跨两个地域
	// - leader地域的副本数必须是3/2 + 1，5/2+1，7/2+1，也就是 2，3，4
	EngineRegionInfos []*EngineRegionInfo `json:"EngineRegionInfos,omitnil,omitempty" name:"EngineRegionInfos"`

	// zk标准版请填CLOUD_PREMIUM，zk标准版无法选择磁盘类型和磁盘容量，默认为CLOUD_PREMIUM
	// zk专业版可以为：CLOUD_SSD,CLOUD_SSD_PLUS,CLOUD_PREMIUM
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// zk标准版请填50，zk标准版无法选择磁盘类型和磁盘容量，磁盘容量默认为50
	StorageCapacity *int64 `json:"StorageCapacity,omitnil,omitempty" name:"StorageCapacity"`

	// zk专业版至多有两个盘，且磁盘的容量在50-3200之间
	// 如果只有一个磁盘，storageCapacity与storageOption里面的capacity应该一致
	StorageOption []*StorageOption `json:"StorageOption,omitnil,omitempty" name:"StorageOption"`

	// ZK引擎实例，可用区分布约束，STRICT:强约束，PERMISSIVE: 弱约束
	AffinityConstraint *string `json:"AffinityConstraint,omitnil,omitempty" name:"AffinityConstraint"`

	// 指定zone id列表
	ZoneIds []*int64 `json:"ZoneIds,omitnil,omitempty" name:"ZoneIds"`

	// 地域特殊标签，用于区分相同地域，不通的业务属性
	EngineRegionTag *string `json:"EngineRegionTag,omitnil,omitempty" name:"EngineRegionTag"`
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
	delete(f, "StorageType")
	delete(f, "StorageCapacity")
	delete(f, "StorageOption")
	delete(f, "AffinityConstraint")
	delete(f, "ZoneIds")
	delete(f, "EngineRegionTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEngineRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEngineResponseParams struct {
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

type CreateGatewayServiceResult struct {
	// 网关服务ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

// Predefined struct for user
type CreateGovernanceAliasRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 服务别名命名空间
	AliasNamespace *string `json:"AliasNamespace,omitnil,omitempty" name:"AliasNamespace"`

	// 服务别名所指向的服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务别名所指向的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务别名描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type CreateGovernanceAliasRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 服务别名命名空间
	AliasNamespace *string `json:"AliasNamespace,omitnil,omitempty" name:"AliasNamespace"`

	// 服务别名所指向的服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务别名所指向的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务别名描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *CreateGovernanceAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Alias")
	delete(f, "AliasNamespace")
	delete(f, "Service")
	delete(f, "Namespace")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGovernanceAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGovernanceAliasResponseParams struct {
	// 创建是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGovernanceAliasResponse struct {
	*tchttp.BaseResponse
	Response *CreateGovernanceAliasResponseParams `json:"Response"`
}

func (r *CreateGovernanceAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGovernanceInstancesRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务实例信息。
	GovernanceInstances []*GovernanceInstanceInput `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

type CreateGovernanceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务实例信息。
	GovernanceInstances []*GovernanceInstanceInput `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

func (r *CreateGovernanceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGovernanceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGovernanceInstancesResponseParams struct {
	// 创建是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGovernanceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *CreateGovernanceInstancesResponseParams `json:"Response"`
}

func (r *CreateGovernanceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGovernanceNamespacesRequestParams struct {
	// tse 实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间信息。
	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitnil,omitempty" name:"GovernanceNamespaces"`
}

type CreateGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// tse 实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间信息。
	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitnil,omitempty" name:"GovernanceNamespaces"`
}

func (r *CreateGovernanceNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceNamespaces")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGovernanceNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGovernanceNamespacesResponseParams struct {
	// 操作是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *CreateGovernanceNamespacesResponseParams `json:"Response"`
}

func (r *CreateGovernanceNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGovernanceServicesRequestParams struct {
	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务信息。
	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitnil,omitempty" name:"GovernanceServices"`
}

type CreateGovernanceServicesRequest struct {
	*tchttp.BaseRequest
	
	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务信息。
	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitnil,omitempty" name:"GovernanceServices"`
}

func (r *CreateGovernanceServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceServices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateGovernanceServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateGovernanceServicesResponseParams struct {
	// 创建是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateGovernanceServicesResponse struct {
	*tchttp.BaseResponse
	Response *CreateGovernanceServicesResponseParams `json:"Response"`
}

func (r *CreateGovernanceServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNativeGatewayServerGroupRequestParams struct {
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点配置
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 公网带宽信息
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 公网配置。
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil,omitempty" name:"InternetConfig"`
}

type CreateNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 节点配置
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// 子网id
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 公网带宽信息
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 公网配置。
	InternetConfig *InternetConfig `json:"InternetConfig,omitnil,omitempty" name:"InternetConfig"`
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
	Result *CreateCloudNativeAPIGatewayServerGroupResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type CreateNativeGatewayServiceSourceRequestParams struct {
	// 网关实例ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 服务来源类型，参考值：
	// - TSE-Nacos 
	// - TSE-Consul 
	// - TSE-PolarisMesh
	// - Customer-Nacos
	// - Customer-Consul
	// - Customer-PolarisMesh
	// - TSF
	// - TKE
	// - EKS
	// - PrivateDNS
	// - Customer-DNS
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 服务来源实例ID，当SourceType的值不为PrivateDNS或Customer-DNS时，必填
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 服务来源实例名称，当SourceType的值不为PrivateDNS时，必填
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 服务来源实例额外信息
	SourceInfo *SourceInfo `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

type CreateNativeGatewayServiceSourceRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 服务来源类型，参考值：
	// - TSE-Nacos 
	// - TSE-Consul 
	// - TSE-PolarisMesh
	// - Customer-Nacos
	// - Customer-Consul
	// - Customer-PolarisMesh
	// - TSF
	// - TKE
	// - EKS
	// - PrivateDNS
	// - Customer-DNS
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 服务来源实例ID，当SourceType的值不为PrivateDNS或Customer-DNS时，必填
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 服务来源实例名称，当SourceType的值不为PrivateDNS时，必填
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 服务来源实例额外信息
	SourceInfo *SourceInfo `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

func (r *CreateNativeGatewayServiceSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNativeGatewayServiceSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayID")
	delete(f, "SourceType")
	delete(f, "SourceID")
	delete(f, "SourceName")
	delete(f, "SourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNativeGatewayServiceSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNativeGatewayServiceSourceResponseParams struct {
	// 创建是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 服务来源ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNativeGatewayServiceSourceResponse struct {
	*tchttp.BaseResponse
	Response *CreateNativeGatewayServiceSourceResponseParams `json:"Response"`
}

func (r *CreateNativeGatewayServiceSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNativeGatewayServiceSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrUpdateConfigFileAndReleaseRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件列表详情	
	ConfigFilePublishInfo *ConfigFilePublishInfo `json:"ConfigFilePublishInfo,omitnil,omitempty" name:"ConfigFilePublishInfo"`

	// 控制开启校验配置版本是否已经存在
	StrictEnable *bool `json:"StrictEnable,omitnil,omitempty" name:"StrictEnable"`
}

type CreateOrUpdateConfigFileAndReleaseRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件列表详情	
	ConfigFilePublishInfo *ConfigFilePublishInfo `json:"ConfigFilePublishInfo,omitnil,omitempty" name:"ConfigFilePublishInfo"`

	// 控制开启校验配置版本是否已经存在
	StrictEnable *bool `json:"StrictEnable,omitnil,omitempty" name:"StrictEnable"`
}

func (r *CreateOrUpdateConfigFileAndReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrUpdateConfigFileAndReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigFilePublishInfo")
	delete(f, "StrictEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateOrUpdateConfigFileAndReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateOrUpdateConfigFileAndReleaseResponseParams struct {
	// 操作是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 配置发布Id
	ConfigFileReleaseId *string `json:"ConfigFileReleaseId,omitnil,omitempty" name:"ConfigFileReleaseId"`

	// 配置文件Id
	ConfigFileId *string `json:"ConfigFileId,omitnil,omitempty" name:"ConfigFileId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateOrUpdateConfigFileAndReleaseResponse struct {
	*tchttp.BaseResponse
	Response *CreateOrUpdateConfigFileAndReleaseResponseParams `json:"Response"`
}

func (r *CreateOrUpdateConfigFileAndReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateOrUpdateConfigFileAndReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePublicNetworkResult struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 客户端公网网络ID
	NetworkId *string `json:"NetworkId,omitnil,omitempty" name:"NetworkId"`
}

// Predefined struct for user
type CreateWafDomainsRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// WAF 防护域名列表
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type CreateWafDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// WAF 防护域名列表
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

func (r *CreateWafDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWafDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateWafDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateWafDomainsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateWafDomainsResponse struct {
	*tchttp.BaseResponse
	Response *CreateWafDomainsResponseParams `json:"Response"`
}

func (r *CreateWafDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateWafDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoScalerResourceStrategyRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DeleteAutoScalerResourceStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *DeleteAutoScalerResourceStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScalerResourceStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoScalerResourceStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoScalerResourceStrategyResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoScalerResourceStrategyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoScalerResourceStrategyResponseParams `json:"Response"`
}

func (r *DeleteAutoScalerResourceStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScalerResourceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayCanaryRuleRequestParams struct {
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 优先级列表，如果配置了此参数，将以此参数为准，忽略Priority参数
	PriorityList []*int64 `json:"PriorityList,omitnil,omitempty" name:"PriorityList"`
}

type DeleteCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 优先级
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 优先级列表，如果配置了此参数，将以此参数为准，忽略Priority参数
	PriorityList []*int64 `json:"PriorityList,omitnil,omitempty" name:"PriorityList"`
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
	delete(f, "PriorityList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayCanaryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayCanaryRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteCloudNativeAPIGatewayPublicNetworkRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id，kong类型时必填
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 公网类型
	// - IPV4 （默认值）
	// - IPV6
	InternetAddressVersion *string `json:"InternetAddressVersion,omitnil,omitempty" name:"InternetAddressVersion"`

	// 公网ip，存在多个公网时必填
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

type DeleteCloudNativeAPIGatewayPublicNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id，kong类型时必填
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 公网类型
	// - IPV4 （默认值）
	// - IPV6
	InternetAddressVersion *string `json:"InternetAddressVersion,omitnil,omitempty" name:"InternetAddressVersion"`

	// 公网ip，存在多个公网时必填
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`
}

func (r *DeleteCloudNativeAPIGatewayPublicNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayPublicNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "InternetAddressVersion")
	delete(f, "Vip")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayPublicNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayPublicNetworkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteCloudNativeAPIGatewayPublicNetworkResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCloudNativeAPIGatewayPublicNetworkResponseParams `json:"Response"`
}

func (r *DeleteCloudNativeAPIGatewayPublicNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCloudNativeAPIGatewayPublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 是否删除实例关联的 CLS 日志主题。
	DeleteClsTopic *bool `json:"DeleteClsTopic,omitnil,omitempty" name:"DeleteClsTopic"`
}

type DeleteCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 是否删除实例关联的 CLS 日志主题。
	DeleteClsTopic *bool `json:"DeleteClsTopic,omitnil,omitempty" name:"DeleteClsTopic"`
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
	Result *DeleteCloudNativeAPIGatewayResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 云原生网关状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayRouteRateLimitRequestParams struct {
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由的ID或名字，不支持名称“未命名”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由的ID或名字，不支持名称“未命名”
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DeleteCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名字，服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否同步删除服务上绑定的路由
	DeleteRoutes *bool `json:"DeleteRoutes,omitnil,omitempty" name:"DeleteRoutes"`
}

type DeleteCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名字，服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否同步删除服务上绑定的路由
	DeleteRoutes *bool `json:"DeleteRoutes,omitnil,omitempty" name:"DeleteRoutes"`
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
	delete(f, "DeleteRoutes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCloudNativeAPIGatewayServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCloudNativeAPIGatewayServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteConfigFileGroupRequestParams struct {
	// tse 实例 id。	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

type DeleteConfigFileGroupRequest struct {
	*tchttp.BaseRequest
	
	// tse 实例 id。	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`
}

func (r *DeleteConfigFileGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFileGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigFileGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFileGroupResponseParams struct {
	// 是否删除成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigFileGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigFileGroupResponseParams `json:"Response"`
}

func (r *DeleteConfigFileGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFileGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFileReleasesRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待删除配置发布详情
	ConfigFileReleases []*ConfigFileReleaseDeletion `json:"ConfigFileReleases,omitnil,omitempty" name:"ConfigFileReleases"`
}

type DeleteConfigFileReleasesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 待删除配置发布详情
	ConfigFileReleases []*ConfigFileReleaseDeletion `json:"ConfigFileReleases,omitnil,omitempty" name:"ConfigFileReleases"`
}

func (r *DeleteConfigFileReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFileReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigFileReleases")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigFileReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFileReleasesResponseParams struct {
	// 删除配置发布结果
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigFileReleasesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigFileReleasesResponseParams `json:"Response"`
}

func (r *DeleteConfigFileReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFileReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFilesRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DeleteConfigFilesRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DeleteConfigFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "Name")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteConfigFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteConfigFilesResponseParams struct {
	// 修改是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteConfigFilesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteConfigFilesResponseParams `json:"Response"`
}

func (r *DeleteConfigFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteEngineRequestParams struct {
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DeleteEngineRequest struct {
	*tchttp.BaseRequest
	
	// 引擎实例 ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DeleteGovernanceAliasesRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名列表
	GovernanceAliases []*GovernanceAlias `json:"GovernanceAliases,omitnil,omitempty" name:"GovernanceAliases"`
}

type DeleteGovernanceAliasesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名列表
	GovernanceAliases []*GovernanceAlias `json:"GovernanceAliases,omitnil,omitempty" name:"GovernanceAliases"`
}

func (r *DeleteGovernanceAliasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceAliasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceAliases")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGovernanceAliasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceAliasesResponseParams struct {
	// 创建是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGovernanceAliasesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGovernanceAliasesResponseParams `json:"Response"`
}

func (r *DeleteGovernanceAliasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceAliasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceInstancesByHostRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要删除的服务实例信息。
	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

type DeleteGovernanceInstancesByHostRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要删除的服务实例信息。
	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

func (r *DeleteGovernanceInstancesByHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceInstancesByHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGovernanceInstancesByHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceInstancesByHostResponseParams struct {
	// 操作是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGovernanceInstancesByHostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGovernanceInstancesByHostResponseParams `json:"Response"`
}

func (r *DeleteGovernanceInstancesByHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceInstancesByHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceInstancesRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要删除的服务实例信息。
	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

type DeleteGovernanceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 要删除的服务实例信息。
	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

func (r *DeleteGovernanceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGovernanceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceInstancesResponseParams struct {
	// 操作是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGovernanceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGovernanceInstancesResponseParams `json:"Response"`
}

func (r *DeleteGovernanceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceNamespacesRequestParams struct {
	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间信息。
	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitnil,omitempty" name:"GovernanceNamespaces"`
}

type DeleteGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间信息。
	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitnil,omitempty" name:"GovernanceNamespaces"`
}

func (r *DeleteGovernanceNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceNamespaces")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGovernanceNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceNamespacesResponseParams struct {
	// 删除是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGovernanceNamespacesResponseParams `json:"Response"`
}

func (r *DeleteGovernanceNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceServicesRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务信息。
	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitnil,omitempty" name:"GovernanceServices"`
}

type DeleteGovernanceServicesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务信息。
	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitnil,omitempty" name:"GovernanceServices"`
}

func (r *DeleteGovernanceServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceServices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteGovernanceServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteGovernanceServicesResponseParams struct {
	// 删除服务结果。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteGovernanceServicesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteGovernanceServicesResponseParams `json:"Response"`
}

func (r *DeleteGovernanceServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNativeGatewayServerGroupRequestParams struct {
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DeleteNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	Result *DeleteNativeGatewayServerGroupResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 删除状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type DeleteNativeGatewayServiceSourceRequestParams struct {
	// 网关实例 ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 服务来源实例 ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`
}

type DeleteNativeGatewayServiceSourceRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例 ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 服务来源实例 ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`
}

func (r *DeleteNativeGatewayServiceSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNativeGatewayServiceSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayID")
	delete(f, "SourceID")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNativeGatewayServiceSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNativeGatewayServiceSourceResponseParams struct {
	// 结果
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNativeGatewayServiceSourceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNativeGatewayServiceSourceResponseParams `json:"Response"`
}

func (r *DeleteNativeGatewayServiceSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNativeGatewayServiceSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWafDomainsRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// WAF 防护域名列表
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

type DeleteWafDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// WAF 防护域名列表
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

func (r *DeleteWafDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWafDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Domains")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteWafDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteWafDomainsResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteWafDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteWafDomainsResponseParams `json:"Response"`
}

func (r *DeleteWafDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteWafDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllConfigFileTemplatesRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeAllConfigFileTemplatesRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeAllConfigFileTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllConfigFileTemplatesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAllConfigFileTemplatesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAllConfigFileTemplatesResponseParams struct {
	// 数据总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 配置文件模板列表
	ConfigFileTemplates []*ConfigFileTemplate `json:"ConfigFileTemplates,omitnil,omitempty" name:"ConfigFileTemplates"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAllConfigFileTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAllConfigFileTemplatesResponseParams `json:"Response"`
}

func (r *DescribeAllConfigFileTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAllConfigFileTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalerResourceStrategiesRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

type DescribeAutoScalerResourceStrategiesRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`
}

func (r *DescribeAutoScalerResourceStrategiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalerResourceStrategiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StrategyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalerResourceStrategiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalerResourceStrategiesResponseParams struct {
	// 获取云原生API网关实例弹性伸缩策略列表响应结果。
	Result *ListCloudNativeAPIGatewayStrategyResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalerResourceStrategiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalerResourceStrategiesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalerResourceStrategiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalerResourceStrategiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalerResourceStrategyBindingGroupsRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAutoScalerResourceStrategyBindingGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 查询偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询数量限制
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAutoScalerResourceStrategyBindingGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalerResourceStrategyBindingGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StrategyId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalerResourceStrategyBindingGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalerResourceStrategyBindingGroupsResponseParams struct {
	// 云原生API网关实例策略绑定网关分组列表响应结果
	Result *ListCloudNativeAPIGatewayStrategyBindingGroupInfoResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalerResourceStrategyBindingGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalerResourceStrategyBindingGroupsResponseParams `json:"Response"`
}

func (r *DescribeAutoScalerResourceStrategyBindingGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalerResourceStrategyBindingGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayCanaryRulesRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 灰度规则类别 Standard｜Lane
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCloudNativeAPIGatewayCanaryRulesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 灰度规则类别 Standard｜Lane
	RuleType *string `json:"RuleType,omitnil,omitempty" name:"RuleType"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	delete(f, "RuleType")
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
	Result *CloudAPIGatewayCanaryRuleList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeCloudNativeAPIGatewayCertificateDetailsRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
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
	Result *KongCertificate `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持BindDomain ，Name
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayCertificatesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持BindDomain ，Name
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Result *KongCertificatesList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id，不填时为默认分组
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribeCloudNativeAPIGatewayConfigRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id，不填时为默认分组
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
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
	Result *DescribeCloudNativeAPIGatewayConfigResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组网络配置列表。
	ConfigList []*CloudNativeAPIGatewayConfig `json:"ConfigList,omitnil,omitempty" name:"ConfigList"`

	// 分组子网信息
	GroupSubnetId *string `json:"GroupSubnetId,omitnil,omitempty" name:"GroupSubnetId"`

	// 分组VPC信息
	GroupVpcId *string `json:"GroupVpcId,omitnil,omitempty" name:"GroupVpcId"`

	// 分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayInfoByIpRequestParams struct {
	// 云原生网关的公网ip
	PublicNetworkIP *string `json:"PublicNetworkIP,omitnil,omitempty" name:"PublicNetworkIP"`
}

type DescribeCloudNativeAPIGatewayInfoByIpRequest struct {
	*tchttp.BaseRequest
	
	// 云原生网关的公网ip
	PublicNetworkIP *string `json:"PublicNetworkIP,omitnil,omitempty" name:"PublicNetworkIP"`
}

func (r *DescribeCloudNativeAPIGatewayInfoByIpRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayInfoByIpRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "PublicNetworkIP")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayInfoByIpRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayInfoByIpResponseParams struct {
	// 出参
	Result *DescribeInstanceInfoByIpResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayInfoByIpResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayInfoByIpResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayInfoByIpResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayInfoByIpResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayNodesRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 实例分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 翻页获取多少个
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页从第几个开始获取
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeCloudNativeAPIGatewayNodesRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 实例分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 翻页获取多少个
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页从第几个开始获取
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	Result *DescribeCloudNativeAPIGatewayNodesResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云原生API网关节点列表。
	NodeList []*CloudNativeAPIGatewayNode `json:"NodeList,omitnil,omitempty" name:"NodeList"`
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayPortsRequestParams struct {
	// 云原生API网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DescribeCloudNativeAPIGatewayPortsRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
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
	Result *DescribeGatewayInstancePortResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DescribeCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
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
	Result *DescribeCloudNativeAPIGatewayResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 云原生API网关状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 云原生API网关名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云原生API网关类型。
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 实例版本：
	// - 2.4.1
	// - 2.5.1
	GatewayVersion *string `json:"GatewayVersion,omitnil,omitempty" name:"GatewayVersion"`

	// 云原生API网关节点信息。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// 云原生API网关vpc配置。
	VpcConfig *CloudNativeAPIGatewayVpcConfig `json:"VpcConfig,omitnil,omitempty" name:"VpcConfig"`

	// 云原生API网关描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 云原生API网关创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例的标签信息
	Tags []*InstanceTagInfo `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 是否开启 cls 日志
	EnableCls *bool `json:"EnableCls,omitnil,omitempty" name:"EnableCls"`

	// 付费模式，0表示后付费，1预付费
	TradeType *int64 `json:"TradeType,omitnil,omitempty" name:"TradeType"`

	// 实例版本，当前支持开发版、标准版、专业版【TRIAL、STANDARD、PROFESSIONAL】
	FeatureVersion *string `json:"FeatureVersion,omitnil,omitempty" name:"FeatureVersion"`

	// 公网出流量带宽，[1,2048]Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 自动续费标记，0表示默认状态(用户未设置，即初始状态)；
	// 1表示自动续费，2表示明确不自动续费(用户设置)，若业务无续费概念或无需自动续费，需要设置为0
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 到期时间，预付费时使用
	CurDeadline *string `json:"CurDeadline,omitnil,omitempty" name:"CurDeadline"`

	// 隔离时间，实例隔离时使用
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 是否开启客户端公网。
	EnableInternet *bool `json:"EnableInternet,omitnil,omitempty" name:"EnableInternet"`

	// 实例实际的地域信息
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// Ingress class名称
	IngressClassName *string `json:"IngressClassName,omitnil,omitempty" name:"IngressClassName"`

	// 公网计费方式。可选取值 BANDWIDTH | TRAFFIC ，表示按带宽和按流量计费。
	InternetPayMode *string `json:"InternetPayMode,omitnil,omitempty" name:"InternetPayMode"`

	// 云原生API网关小版本号
	GatewayMinorVersion *string `json:"GatewayMinorVersion,omitnil,omitempty" name:"GatewayMinorVersion"`

	// 实例监听的端口信息
	InstancePort *InstancePort `json:"InstancePort,omitnil,omitempty" name:"InstancePort"`

	// 公网CLB默认类型
	LoadBalancerType *string `json:"LoadBalancerType,omitnil,omitempty" name:"LoadBalancerType"`

	// 公网IP地址列表
	PublicIpAddresses []*string `json:"PublicIpAddresses,omitnil,omitempty" name:"PublicIpAddresses"`

	// 是否开启删除保护
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayRouteRateLimitRequestParams struct {
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由Id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
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
	Result *CloudNativeAPIGatewayRateLimitDetail `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务的名字，精确匹配
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 路由的名字，精确匹配
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name, path, host, method, service, protocol
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayRoutesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务的名字，精确匹配
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 路由的名字，精确匹配
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name, path, host, method, service, protocol
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Result *KongServiceRouteList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关Id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	Result *CloudNativeAPIGatewayRateLimitDetail `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCloudNativeAPIGatewayServicesLightRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表 offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 id、name、upstreamType
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayServicesLightRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表 offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 id、name、upstreamType
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

func (r *DescribeCloudNativeAPIGatewayServicesLightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayServicesLightRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayServicesLightRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayServicesLightResponseParams struct {
	// 无
	Result *GatewayServices `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayServicesLightResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayServicesLightResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayServicesLightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayServicesLightResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayServicesRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表 offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name,upstreamType
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCloudNativeAPIGatewayServicesRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 列表数量
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 列表 offset
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 过滤条件，多个过滤条件之间是与的关系，支持 name,upstreamType
	Filters []*ListFilter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Result *KongServices `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeCloudNativeAPIGatewayUpstreamRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名字
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

type DescribeCloudNativeAPIGatewayUpstreamRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名字
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`
}

func (r *DescribeCloudNativeAPIGatewayUpstreamRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayUpstreamRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCloudNativeAPIGatewayUpstreamRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewayUpstreamResponseParams struct {
	// 无
	Result *KongUpstreamList `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeCloudNativeAPIGatewayUpstreamResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCloudNativeAPIGatewayUpstreamResponseParams `json:"Response"`
}

func (r *DescribeCloudNativeAPIGatewayUpstreamResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCloudNativeAPIGatewayUpstreamResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCloudNativeAPIGatewaysRequestParams struct {
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 请求过滤参数，支持按照实例名称、ID和标签键值（Name、GatewayId、Tag）筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeCloudNativeAPIGatewaysRequest struct {
	*tchttp.BaseRequest
	
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为 0。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 请求过滤参数，支持按照实例名称、ID和标签键值（Name、GatewayId、Tag）筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Result *ListCloudNativeAPIGatewayResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type DescribeConfigFileGroupsRequestParams struct {
	// tse实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 根据命名空间过滤
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 根据配置文件组名过滤
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 根据配置文件组名过滤
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 返回数量，默认为20，最大值为100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeConfigFileGroupsRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 根据命名空间过滤
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 根据配置文件组名过滤
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 根据配置文件组名过滤
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 返回数量，默认为20，最大值为100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeConfigFileGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "FileName")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFileGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileGroupsResponseParams struct {
	// 列表总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 配置文件组列表
	ConfigFileGroups []*ConfigFileGroup `json:"ConfigFileGroups,omitnil,omitempty" name:"ConfigFileGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFileGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFileGroupsResponseParams `json:"Response"`
}

func (r *DescribeConfigFileGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleaseHistoriesRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布历史记录id，用于分页优化，一般指定 EndId，就不用指定 Offset，否则分页可能不连续
	EndId *uint64 `json:"EndId,omitnil,omitempty" name:"EndId"`

	// 配置文件ID
	ConfigFileId *string `json:"ConfigFileId,omitnil,omitempty" name:"ConfigFileId"`

	// 返回数量，默认为20，最大值为100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeConfigFileReleaseHistoriesRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 发布历史记录id，用于分页优化，一般指定 EndId，就不用指定 Offset，否则分页可能不连续
	EndId *uint64 `json:"EndId,omitnil,omitempty" name:"EndId"`

	// 配置文件ID
	ConfigFileId *string `json:"ConfigFileId,omitnil,omitempty" name:"ConfigFileId"`

	// 返回数量，默认为20，最大值为100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeConfigFileReleaseHistoriesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleaseHistoriesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "Name")
	delete(f, "EndId")
	delete(f, "ConfigFileId")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFileReleaseHistoriesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleaseHistoriesResponseParams struct {
	// 数据总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 配置文件发布历史列表
	ConfigFileReleaseHistories []*ConfigFileReleaseHistory `json:"ConfigFileReleaseHistories,omitnil,omitempty" name:"ConfigFileReleaseHistories"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFileReleaseHistoriesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFileReleaseHistoriesResponseParams `json:"Response"`
}

func (r *DescribeConfigFileReleaseHistoriesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleaseHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleaseRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件发布名称
	ReleaseName *string `json:"ReleaseName,omitnil,omitempty" name:"ReleaseName"`

	// 配置文件发布Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeConfigFileReleaseRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组名称
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件发布名称
	ReleaseName *string `json:"ReleaseName,omitnil,omitempty" name:"ReleaseName"`

	// 配置文件发布Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeConfigFileReleaseRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleaseRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "Name")
	delete(f, "ReleaseName")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFileReleaseRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleaseResponseParams struct {
	// 配置文件发布详情
	ConfigFileRelease *ConfigFileRelease `json:"ConfigFileRelease,omitnil,omitempty" name:"ConfigFileRelease"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFileReleaseResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFileReleaseResponseParams `json:"Response"`
}

func (r *DescribeConfigFileReleaseResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleaseResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleaseVersionsRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件ID
	ConfigFileId *string `json:"ConfigFileId,omitnil,omitempty" name:"ConfigFileId"`
}

type DescribeConfigFileReleaseVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 配置文件ID
	ConfigFileId *string `json:"ConfigFileId,omitnil,omitempty" name:"ConfigFileId"`
}

func (r *DescribeConfigFileReleaseVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleaseVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "FileName")
	delete(f, "ConfigFileId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFileReleaseVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleaseVersionsResponseParams struct {
	// 版本信息
	ReleaseVersions []*ReleaseVersion `json:"ReleaseVersions,omitnil,omitempty" name:"ReleaseVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFileReleaseVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFileReleaseVersionsResponseParams `json:"Response"`
}

func (r *DescribeConfigFileReleaseVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleaseVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleasesRequestParams struct {
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 只保护处于使用状态
	OnlyUse *bool `json:"OnlyUse,omitnil,omitempty" name:"OnlyUse"`

	// 发布名称
	ReleaseName *string `json:"ReleaseName,omitnil,omitempty" name:"ReleaseName"`

	// 排序字段，mtime/version/name
	// ，默认version
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序，asc/desc，默认 desc
	OrderDesc *string `json:"OrderDesc,omitnil,omitempty" name:"OrderDesc"`

	// 配置发布ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeConfigFileReleasesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 文件名称
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`

	// 只保护处于使用状态
	OnlyUse *bool `json:"OnlyUse,omitnil,omitempty" name:"OnlyUse"`

	// 发布名称
	ReleaseName *string `json:"ReleaseName,omitnil,omitempty" name:"ReleaseName"`

	// 排序字段，mtime/version/name
	// ，默认version
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序，asc/desc，默认 desc
	OrderDesc *string `json:"OrderDesc,omitnil,omitempty" name:"OrderDesc"`

	// 配置发布ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeConfigFileReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "FileName")
	delete(f, "OnlyUse")
	delete(f, "ReleaseName")
	delete(f, "OrderField")
	delete(f, "OrderDesc")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFileReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileReleasesResponseParams struct {
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 发布列表
	Releases []*ConfigFileRelease `json:"Releases,omitnil,omitempty" name:"Releases"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFileReleasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFileReleasesResponseParams `json:"Response"`
}

func (r *DescribeConfigFileReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeConfigFileRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 配置文件Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeConfigFileRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "Name")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFileRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFileResponseParams struct {
	// 配置文件
	ConfigFile *ConfigFile `json:"ConfigFile,omitnil,omitempty" name:"ConfigFile"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFileResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFileResponseParams `json:"Response"`
}

func (r *DescribeConfigFileResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFileResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFilesByGroupRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间名
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组名
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 返回数量，默认为20，最大值为100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeConfigFilesByGroupRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间名
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 组名
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 返回数量，默认为20，最大值为100。	
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeConfigFilesByGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFilesByGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Group")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFilesByGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFilesByGroupResponseParams struct {
	// 记录总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 配置文件列表
	ConfigFiles []*ConfigFile `json:"ConfigFiles,omitnil,omitempty" name:"ConfigFiles"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFilesByGroupResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFilesByGroupResponseParams `json:"Response"`
}

func (r *DescribeConfigFilesByGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFilesByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFilesRequestParams struct {
	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 组名
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签列表
	Tags []*ConfigFileTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 返回数量，默认为20，最大值为100。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 配置文件ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

type DescribeConfigFilesRequest struct {
	*tchttp.BaseRequest
	
	// 命名空间名称
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 组名
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置文件名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签列表
	Tags []*ConfigFileTag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 返回数量，默认为20，最大值为100。	
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。	
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 配置文件ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`
}

func (r *DescribeConfigFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Namespace")
	delete(f, "InstanceId")
	delete(f, "Group")
	delete(f, "Name")
	delete(f, "Tags")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "Id")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeConfigFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeConfigFilesResponseParams struct {
	// 分页总数量
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 配置文件列表
	ConfigFiles []*ConfigFile `json:"ConfigFiles,omitnil,omitempty" name:"ConfigFiles"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeConfigFilesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeConfigFilesResponseParams `json:"Response"`
}

func (r *DescribeConfigFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGatewayInstancePortResult struct {
	// 云原生API网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关实例协议端口列表
	GatewayInstancePortList []*GatewayInstanceSchemeAndPorts `json:"GatewayInstancePortList,omitnil,omitempty" name:"GatewayInstancePortList"`
}

// Predefined struct for user
type DescribeGovernanceAliasesRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名所指向的服务名。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务别名所指向的命名空间名。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务别名。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 服务别名命名空间。
	AliasNamespace *string `json:"AliasNamespace,omitnil,omitempty" name:"AliasNamespace"`

	// 服务别名描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGovernanceAliasesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名所指向的服务名。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务别名所指向的命名空间名。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务别名。
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 服务别名命名空间。
	AliasNamespace *string `json:"AliasNamespace,omitnil,omitempty" name:"AliasNamespace"`

	// 服务别名描述。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGovernanceAliasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceAliasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Service")
	delete(f, "Namespace")
	delete(f, "Alias")
	delete(f, "AliasNamespace")
	delete(f, "Comment")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGovernanceAliasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceAliasesResponseParams struct {
	// 服务别名总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务别名列表。
	Content []*GovernanceAlias `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGovernanceAliasesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGovernanceAliasesResponseParams `json:"Response"`
}

func (r *DescribeGovernanceAliasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceAliasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceInstancesRequestParams struct {
	// 实例所在的服务名。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 实例所在命名空间名。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 根据实例ip过滤，多个ip使用英文逗号分隔。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 根据实例版本过滤。
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 根据实例协议过滤。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 根据实例健康状态过滤。false：表示不健康，true：表示健康。
	HealthStatus *bool `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 根据实例隔离状态过滤。false：表示非隔离，true：表示隔离中。
	Isolate *bool `json:"Isolate,omitnil,omitempty" name:"Isolate"`

	// 根据元数据信息过滤。目前只支持一组元数据键值，若传了多个键值对，只会以第一个过滤。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 地域
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`
}

type DescribeGovernanceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 实例所在的服务名。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 实例所在命名空间名。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 根据实例ip过滤，多个ip使用英文逗号分隔。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 根据实例版本过滤。
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 根据实例协议过滤。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 根据实例健康状态过滤。false：表示不健康，true：表示健康。
	HealthStatus *bool `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 根据实例隔离状态过滤。false：表示非隔离，true：表示隔离中。
	Isolate *bool `json:"Isolate,omitnil,omitempty" name:"Isolate"`

	// 根据元数据信息过滤。目前只支持一组元数据键值，若传了多个键值对，只会以第一个过滤。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 地域
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`
}

func (r *DescribeGovernanceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Service")
	delete(f, "Namespace")
	delete(f, "InstanceId")
	delete(f, "Host")
	delete(f, "InstanceVersion")
	delete(f, "Protocol")
	delete(f, "HealthStatus")
	delete(f, "Isolate")
	delete(f, "Metadatas")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Location")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGovernanceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceInstancesResponseParams struct {
	// 服务实例总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务里实例列表。
	Content []*GovernanceInstance `json:"Content,omitnil,omitempty" name:"Content"`

	// 地域
	Location *Location `json:"Location,omitnil,omitempty" name:"Location"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGovernanceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGovernanceInstancesResponseParams `json:"Response"`
}

func (r *DescribeGovernanceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceNamespacesRequestParams struct {
	// tse实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 根据命名空间名称过滤。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否开启同步到全局注册中心	
	SyncToGlobalRegistry *string `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 根据命名空间名称过滤。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否开启同步到全局注册中心	
	SyncToGlobalRegistry *string `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeGovernanceNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Name")
	delete(f, "SyncToGlobalRegistry")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGovernanceNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceNamespacesResponseParams struct {
	// 列表总数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 治理中心命名空间实例列表。
	Content []*GovernanceNamespace `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGovernanceNamespacesResponseParams `json:"Response"`
}

func (r *DescribeGovernanceNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceServiceContractVersionsRequestParams struct {
	// 引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`
}

type DescribeGovernanceServiceContractVersionsRequest struct {
	*tchttp.BaseRequest
	
	// 引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`
}

func (r *DescribeGovernanceServiceContractVersionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceServiceContractVersionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Namespace")
	delete(f, "Service")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGovernanceServiceContractVersionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceServiceContractVersionsResponseParams struct {
	// 服务契约版本列表
	GovernanceServiceContractVersions []*GovernanceServiceContractVersion `json:"GovernanceServiceContractVersions,omitnil,omitempty" name:"GovernanceServiceContractVersions"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGovernanceServiceContractVersionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGovernanceServiceContractVersionsResponseParams `json:"Response"`
}

func (r *DescribeGovernanceServiceContractVersionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceServiceContractVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceServiceContractsRequestParams struct {
	// 北极星引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 契约名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 契约版本
	ContractVersion *string `json:"ContractVersion,omitnil,omitempty" name:"ContractVersion"`

	// 契约协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 是否只展示基本信息
	Brief *bool `json:"Brief,omitnil,omitempty" name:"Brief"`
}

type DescribeGovernanceServiceContractsRequest struct {
	*tchttp.BaseRequest
	
	// 北极星引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 分页偏移量
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 分页条数
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 契约名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 契约版本
	ContractVersion *string `json:"ContractVersion,omitnil,omitempty" name:"ContractVersion"`

	// 契约协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 是否只展示基本信息
	Brief *bool `json:"Brief,omitnil,omitempty" name:"Brief"`
}

func (r *DescribeGovernanceServiceContractsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceServiceContractsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Namespace")
	delete(f, "Service")
	delete(f, "Name")
	delete(f, "ContractVersion")
	delete(f, "Protocol")
	delete(f, "Brief")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGovernanceServiceContractsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceServiceContractsResponseParams struct {
	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 返回条数
	Size *int64 `json:"Size,omitnil,omitempty" name:"Size"`

	// 契约定义列表
	ServiceContracts []*GovernanceServiceContract `json:"ServiceContracts,omitnil,omitempty" name:"ServiceContracts"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGovernanceServiceContractsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGovernanceServiceContractsResponseParams `json:"Response"`
}

func (r *DescribeGovernanceServiceContractsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceServiceContractsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceServicesRequestParams struct {
	// 按照服务名过滤，精确匹配。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 按照命名空间过滤，精确匹配。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 使用元数据过滤，目前只支持一组元组数，若传了多条，只会使用第一条元数据过滤。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务所属部门。
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// 服务所属业务。
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// 服务中实例的ip，用来过滤服务。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 是否只查询存在健康实例的服务
	OnlyExistHealthyInstance *bool `json:"OnlyExistHealthyInstance,omitnil,omitempty" name:"OnlyExistHealthyInstance"`

	// 是否开启同步到全局注册中心	
	SyncToGlobalRegistry *string `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`
}

type DescribeGovernanceServicesRequest struct {
	*tchttp.BaseRequest
	
	// 按照服务名过滤，精确匹配。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 按照命名空间过滤，精确匹配。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 使用元数据过滤，目前只支持一组元组数，若传了多条，只会使用第一条元数据过滤。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务所属部门。
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// 服务所属业务。
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// 服务中实例的ip，用来过滤服务。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 是否只查询存在健康实例的服务
	OnlyExistHealthyInstance *bool `json:"OnlyExistHealthyInstance,omitnil,omitempty" name:"OnlyExistHealthyInstance"`

	// 是否开启同步到全局注册中心	
	SyncToGlobalRegistry *string `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`
}

func (r *DescribeGovernanceServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Namespace")
	delete(f, "Metadatas")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceId")
	delete(f, "Department")
	delete(f, "Business")
	delete(f, "Host")
	delete(f, "OnlyExistHealthyInstance")
	delete(f, "SyncToGlobalRegistry")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeGovernanceServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeGovernanceServicesResponseParams struct {
	// 服务数总量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 服务信息详情。
	Content []*GovernanceService `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeGovernanceServicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeGovernanceServicesResponseParams `json:"Response"`
}

func (r *DescribeGovernanceServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceInfoByIpResult struct {
	// 实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribeInstanceRegionInfo struct {
	// 引擎部署地域信息
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// 引擎在该地域的副本数
	Replica *int64 `json:"Replica,omitnil,omitempty" name:"Replica"`

	// 引擎在该地域的规格id
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 客户端内网的网络信息
	IntranetVpcInfos []*VpcInfo `json:"IntranetVpcInfos,omitnil,omitempty" name:"IntranetVpcInfos"`

	// 控制台内网的网络信息
	ConsoleIntranetVpcInfos []*VpcInfo `json:"ConsoleIntranetVpcInfos,omitnil,omitempty" name:"ConsoleIntranetVpcInfos"`

	// 是否开公网
	EnableClientInternet *bool `json:"EnableClientInternet,omitnil,omitempty" name:"EnableClientInternet"`

	// 限流客户端内网的网络信息
	LimiterIntranetVpcInfos []*VpcInfo `json:"LimiterIntranetVpcInfos,omitnil,omitempty" name:"LimiterIntranetVpcInfos"`

	// 是否为主地域，仅在服务治理中心多地域有效
	MainRegion *bool `json:"MainRegion,omitnil,omitempty" name:"MainRegion"`

	// 该地域所在的EKS集群
	EKSClusterID *string `json:"EKSClusterID,omitnil,omitempty" name:"EKSClusterID"`
}

// Predefined struct for user
type DescribeInstanceTagInfosRequestParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

type DescribeInstanceTagInfosRequest struct {
	*tchttp.BaseRequest
	
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceTagInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTagInfosRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstanceTagInfosRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstanceTagInfosResponseParams struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例标签集合
	TagInfos []*InstanceTagInfo `json:"TagInfos,omitnil,omitempty" name:"TagInfos"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeInstanceTagInfosResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstanceTagInfosResponseParams `json:"Response"`
}

func (r *DescribeInstanceTagInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstanceTagInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNacosReplicasRequestParams struct {
	// 引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeNacosReplicasRequest struct {
	*tchttp.BaseRequest
	
	// 引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	Replicas []*NacosReplica `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 副本个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeNacosServerInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 接口列表
	Content []*NacosServerInterface `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持按照分组名称、分组ID（Name、GroupId）筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
}

type DescribeNativeGatewayServerGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 偏移量，默认为 0。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为 20。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 过滤参数，支持按照分组名称、分组ID（Name、GroupId）筛选
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`
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
	Result *NativeGatewayServerGroups `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeNativeGatewayServiceSourcesRequestParams struct {
	// 网关实例ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 单页条数，最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务来源ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 服务来源实例名称，模糊搜索
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 微服务引擎类型：TSE-Nacos｜TSE-Consul｜TSE-PolarisMesh｜Customer-Nacos｜Customer-Consul｜Customer-PolarisMesh
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// 排序字段类型，当前仅支持SourceName
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序类型，AES/DESC
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

type DescribeNativeGatewayServiceSourcesRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 单页条数，最大100
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 分页偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 服务来源ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 服务来源实例名称，模糊搜索
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 微服务引擎类型：TSE-Nacos｜TSE-Consul｜TSE-PolarisMesh｜Customer-Nacos｜Customer-Consul｜Customer-PolarisMesh
	SourceTypes []*string `json:"SourceTypes,omitnil,omitempty" name:"SourceTypes"`

	// 排序字段类型，当前仅支持SourceName
	OrderField *string `json:"OrderField,omitnil,omitempty" name:"OrderField"`

	// 排序类型，AES/DESC
	OrderType *string `json:"OrderType,omitnil,omitempty" name:"OrderType"`
}

func (r *DescribeNativeGatewayServiceSourcesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNativeGatewayServiceSourcesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayID")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceID")
	delete(f, "SourceName")
	delete(f, "SourceTypes")
	delete(f, "OrderField")
	delete(f, "OrderType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNativeGatewayServiceSourcesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNativeGatewayServiceSourcesResponseParams struct {
	// 总实例数
	Total *uint64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 服务来源实例列表
	List []*NativeGatewayServiceSourceItem `json:"List,omitnil,omitempty" name:"List"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNativeGatewayServiceSourcesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNativeGatewayServiceSourcesResponseParams `json:"Response"`
}

func (r *DescribeNativeGatewayServiceSourcesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNativeGatewayServiceSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeOneCloudNativeAPIGatewayServiceRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名字，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeOneCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名字，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	Result *KongServiceDetail `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribePublicAddressConfigRequestParams struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 查询该分组的公网信息，不传则查询实例所有的公网负载均衡信息
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

type DescribePublicAddressConfigRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 查询该分组的公网信息，不传则查询实例所有的公网负载均衡信息
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`
}

func (r *DescribePublicAddressConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicAddressConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicAddressConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicAddressConfigResponseParams struct {
	// 公网地址信息
	Result *DescribePublicAddressConfigResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicAddressConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicAddressConfigResponseParams `json:"Response"`
}

func (r *DescribePublicAddressConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicAddressConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicAddressConfigResult struct {
	// 网关实例id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 公网地址信息
	ConfigList []*PublicAddressConfig `json:"ConfigList,omitnil,omitempty" name:"ConfigList"`

	// 总个数	
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

// Predefined struct for user
type DescribePublicNetworkRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网络ID
	NetworkId *string `json:"NetworkId,omitnil,omitempty" name:"NetworkId"`
}

type DescribePublicNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网络ID
	NetworkId *string `json:"NetworkId,omitnil,omitempty" name:"NetworkId"`
}

func (r *DescribePublicNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "NetworkId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublicNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublicNetworkResponseParams struct {
	// 获取云原生API网关公网详情响应结果。
	Result *DescribePublicNetworkResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribePublicNetworkResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublicNetworkResponseParams `json:"Response"`
}

func (r *DescribePublicNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublicNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublicNetworkResult struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组ID
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 客户端公网信息
	PublicNetwork *CloudNativeAPIGatewayConfig `json:"PublicNetwork,omitnil,omitempty" name:"PublicNetwork"`
}

// Predefined struct for user
type DescribeSREInstanceAccessAddressRequestParams struct {
	// 注册引擎实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 引擎其他组件名称（pushgateway、polaris-limiter）
	Workload *string `json:"Workload,omitnil,omitempty" name:"Workload"`

	// 部署地域
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`
}

type DescribeSREInstanceAccessAddressRequest struct {
	*tchttp.BaseRequest
	
	// 注册引擎实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 引擎其他组件名称（pushgateway、polaris-limiter）
	Workload *string `json:"Workload,omitnil,omitempty" name:"Workload"`

	// 部署地域
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`
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
	IntranetAddress *string `json:"IntranetAddress,omitnil,omitempty" name:"IntranetAddress"`

	// 公网访问地址
	InternetAddress *string `json:"InternetAddress,omitnil,omitempty" name:"InternetAddress"`

	// apollo多环境公网ip
	EnvAddressInfos []*EnvAddressInfo `json:"EnvAddressInfos,omitnil,omitempty" name:"EnvAddressInfos"`

	// 控制台公网访问地址
	ConsoleInternetAddress *string `json:"ConsoleInternetAddress,omitnil,omitempty" name:"ConsoleInternetAddress"`

	// 控制台内网访问地址
	ConsoleIntranetAddress *string `json:"ConsoleIntranetAddress,omitnil,omitempty" name:"ConsoleIntranetAddress"`

	// 客户端公网带宽
	InternetBandWidth *int64 `json:"InternetBandWidth,omitnil,omitempty" name:"InternetBandWidth"`

	// 控制台公网带宽
	ConsoleInternetBandWidth *int64 `json:"ConsoleInternetBandWidth,omitnil,omitempty" name:"ConsoleInternetBandWidth"`

	// 北极星限流server节点接入IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	LimiterAddressInfos []*PolarisLimiterAddress `json:"LimiterAddressInfos,omitnil,omitempty" name:"LimiterAddressInfos"`

	// InternetAddress 的公网 CLB 多可用区信息
	CLBMultiRegion *CLBMultiRegion `json:"CLBMultiRegion,omitnil,omitempty" name:"CLBMultiRegion"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询类型
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 调用方来源
	QuerySource *string `json:"QuerySource,omitnil,omitempty" name:"QuerySource"`
}

type DescribeSREInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 请求过滤参数
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 翻页单页查询限制数量[0,1000], 默认值0
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 翻页单页偏移量，默认值0
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 查询类型
	QueryType *string `json:"QueryType,omitnil,omitempty" name:"QueryType"`

	// 调用方来源
	QuerySource *string `json:"QuerySource,omitnil,omitempty" name:"QuerySource"`
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
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 实例记录
	Content []*SREInstance `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type DescribeUpstreamHealthCheckConfigRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type DescribeUpstreamHealthCheckConfigRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

func (r *DescribeUpstreamHealthCheckConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamHealthCheckConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUpstreamHealthCheckConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUpstreamHealthCheckConfigResponseParams struct {
	// 健康检查配置
	Result *UpstreamHealthCheckConfig `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeUpstreamHealthCheckConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUpstreamHealthCheckConfigResponseParams `json:"Response"`
}

func (r *DescribeUpstreamHealthCheckConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUpstreamHealthCheckConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafDomainsRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

type DescribeWafDomainsRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`
}

func (r *DescribeWafDomainsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafDomainsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafDomainsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafDomainsResponseParams struct {
	// 已被 WAF 防护域名
	Result *DescribeWafDomainsResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWafDomainsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafDomainsResponseParams `json:"Response"`
}

func (r *DescribeWafDomainsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafDomainsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafDomainsResult struct {
	// WAF防护域名列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Domains []*string `json:"Domains,omitnil,omitempty" name:"Domains"`
}

// Predefined struct for user
type DescribeWafProtectionRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	//  防护资源的类型。
	// - Global  实例
	// - Service  服务
	// - Route  路由
	// - Object  对象
	//
	// Deprecated: Type is deprecated.
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 防护资源类型列表，支持查询多个类型（Global、Service、Route、Object）。为空时，默认查询Global类型。
	TypeList []*string `json:"TypeList,omitnil,omitempty" name:"TypeList"`
}

type DescribeWafProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	//  防护资源的类型。
	// - Global  实例
	// - Service  服务
	// - Route  路由
	// - Object  对象
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 防护资源类型列表，支持查询多个类型（Global、Service、Route、Object）。为空时，默认查询Global类型。
	TypeList []*string `json:"TypeList,omitnil,omitempty" name:"TypeList"`
}

func (r *DescribeWafProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Type")
	delete(f, "TypeList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeWafProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeWafProtectionResponseParams struct {
	// 保护状态
	Result *DescribeWafProtectionResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeWafProtectionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeWafProtectionResponseParams `json:"Response"`
}

func (r *DescribeWafProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeWafProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWafProtectionResult struct {
	// 全局防护状态
	GlobalStatus *string `json:"GlobalStatus,omitnil,omitempty" name:"GlobalStatus"`

	// 服务防护状态
	ServicesStatus []*ServiceWafStatus `json:"ServicesStatus,omitnil,omitempty" name:"ServicesStatus"`

	// 路由防护状态
	RouteStatus []*RouteWafStatus `json:"RouteStatus,omitnil,omitempty" name:"RouteStatus"`

	// 对象防护状态
	ObjectStatus *string `json:"ObjectStatus,omitnil,omitempty" name:"ObjectStatus"`
}

// Predefined struct for user
type DescribeZookeeperReplicasRequestParams struct {
	// 注册引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeZookeeperReplicasRequest struct {
	*tchttp.BaseRequest
	
	// 注册引擎实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 副本列表Limit
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 副本列表Offset
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	Replicas []*ZookeeperReplica `json:"Replicas,omitnil,omitempty" name:"Replicas"`

	// 副本个数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeZookeeperServerInterfacesRequest struct {
	*tchttp.BaseRequest
	
	// 实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 返回的列表个数
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 返回的列表起始偏移量
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
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
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 接口列表
	Content []*ZookeeperServerInterface `json:"Content,omitnil,omitempty" name:"Content"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 控制台初始密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 引擎接口的管理员 Token
	Token *string `json:"Token,omitnil,omitempty" name:"Token"`
}

type EngineRegionInfo struct {
	// 引擎节点所在地域
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// 此地域节点分配数量
	Replica *int64 `json:"Replica,omitnil,omitempty" name:"Replica"`

	// 集群网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil,omitempty" name:"VpcInfos"`

	// Polaris: 是否为主地域
	// Zookeeper: 是否为Leader固定地域
	MainRegion *bool `json:"MainRegion,omitnil,omitempty" name:"MainRegion"`

	// 引擎规格ID
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`
}

type EnvAddressInfo struct {
	// 环境名
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 是否开启config公网
	EnableConfigInternet *bool `json:"EnableConfigInternet,omitnil,omitempty" name:"EnableConfigInternet"`

	// config公网ip
	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitnil,omitempty" name:"ConfigInternetServiceIp"`

	// config内网访问地址
	ConfigIntranetAddress *string `json:"ConfigIntranetAddress,omitnil,omitempty" name:"ConfigIntranetAddress"`

	// 是否开启config内网clb
	EnableConfigIntranet *bool `json:"EnableConfigIntranet,omitnil,omitempty" name:"EnableConfigIntranet"`

	// 客户端公网带宽
	InternetBandWidth *int64 `json:"InternetBandWidth,omitnil,omitempty" name:"InternetBandWidth"`

	// 客户端公网CLB多可用区信息
	CLBMultiRegion *CLBMultiRegion `json:"CLBMultiRegion,omitnil,omitempty" name:"CLBMultiRegion"`
}

type EnvInfo struct {
	// 环境名称
	EnvName *string `json:"EnvName,omitnil,omitempty" name:"EnvName"`

	// 环境对应的网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil,omitempty" name:"VpcInfos"`

	// 云硬盘容量
	StorageCapacity *int64 `json:"StorageCapacity,omitnil,omitempty" name:"StorageCapacity"`

	// 运行状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// Admin service 访问地址
	AdminServiceIp *string `json:"AdminServiceIp,omitnil,omitempty" name:"AdminServiceIp"`

	// Config service访问地址
	ConfigServiceIp *string `json:"ConfigServiceIp,omitnil,omitempty" name:"ConfigServiceIp"`

	// 是否开启config-server公网
	EnableConfigInternet *bool `json:"EnableConfigInternet,omitnil,omitempty" name:"EnableConfigInternet"`

	// config-server公网访问地址
	ConfigInternetServiceIp *string `json:"ConfigInternetServiceIp,omitnil,omitempty" name:"ConfigInternetServiceIp"`

	// 规格ID
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 环境的节点数
	EnvReplica *int64 `json:"EnvReplica,omitnil,omitempty" name:"EnvReplica"`

	// 环境运行的节点数
	RunningCount *int64 `json:"RunningCount,omitnil,omitempty" name:"RunningCount"`

	// 环境别名
	AliasEnvName *string `json:"AliasEnvName,omitnil,omitempty" name:"AliasEnvName"`

	// 环境描述
	EnvDesc *string `json:"EnvDesc,omitnil,omitempty" name:"EnvDesc"`

	// 客户端带宽
	ClientBandWidth *uint64 `json:"ClientBandWidth,omitnil,omitempty" name:"ClientBandWidth"`

	// 客户端内网开关
	EnableConfigIntranet *bool `json:"EnableConfigIntranet,omitnil,omitempty" name:"EnableConfigIntranet"`
}

type ExternalRedis struct {
	// redis ip
	RedisHost *string `json:"RedisHost,omitnil,omitempty" name:"RedisHost"`

	// redis密码
	RedisPassword *string `json:"RedisPassword,omitnil,omitempty" name:"RedisPassword"`

	// redis端口
	RedisPort *int64 `json:"RedisPort,omitnil,omitempty" name:"RedisPort"`

	// 超时时间  ms
	RedisTimeout *int64 `json:"RedisTimeout,omitnil,omitempty" name:"RedisTimeout"`
}

type Filter struct {
	// 过滤参数名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 过滤参数值
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type GatewayInstanceSchemeAndPorts struct {
	// 端口协议，可选HTTP、HTTPS、TCP和UDP
	Scheme *string `json:"Scheme,omitnil,omitempty" name:"Scheme"`

	// 端口列表
	PortList []*uint64 `json:"PortList,omitnil,omitempty" name:"PortList"`
}

type GatewayServices struct {
	// 服务列表
	ServiceList []*KongServiceLightPreview `json:"ServiceList,omitnil,omitempty" name:"ServiceList"`

	// 结果总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type GovernanceAlias struct {
	// 服务别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 服务别名命名空间
	AliasNamespace *string `json:"AliasNamespace,omitnil,omitempty" name:"AliasNamespace"`

	// 服务别名指向的服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务别名指向的服务命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务别名的描述信息
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 服务别名创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 服务别名修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 服务别名ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 该服务别名是否可以编辑
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 元数据信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`
}

type GovernanceInstance struct {
	// 实例id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 实例所在服务名。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 实例所在命名空间名。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 实例ip地址。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 实例端口信息。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 通信协议。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 版本信息。
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 负载均衡权重。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 是否开启健康检查。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// 实例是否健康。
	Healthy *bool `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// 实例是否隔离。
	Isolate *bool `json:"Isolate,omitnil,omitempty" name:"Isolate"`

	// 实例创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 实例修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 元数据数组。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// 上报心跳间隔。
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 版本信息。
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 状态信息
	HealthStatus *string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type GovernanceInstanceInput struct {
	// 实例所在服务名。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 实例服务所在命名空间。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 实例负载均衡权重信息。不填默认为100。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 实例默认健康信息。不填默认为健康。
	Healthy *bool `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// 实例隔离信息。不填默认为非隔离。
	Isolate *bool `json:"Isolate,omitnil,omitempty" name:"Isolate"`

	// 实例ip。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 实例监听端口。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例使用协议。不填默认为空。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 实例版本。不填默认为空。
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 是否启用健康检查。不填默认不启用。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// 上报心跳时间间隔。若 EnableHealthCheck 为不启用，则此参数不生效；若 EnableHealthCheck 启用，此参数不填，则默认 ttl 为 5s。
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`
}

type GovernanceInstanceUpdate struct {
	// 实例所在服务名。
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 实例服务所在命名空间。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 实例负载均衡权重信息。不填默认为100。
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 实例默认健康信息。不填默认为健康。
	Healthy *bool `json:"Healthy,omitnil,omitempty" name:"Healthy"`

	// 实例隔离信息。不填默认为非隔离。
	Isolate *bool `json:"Isolate,omitnil,omitempty" name:"Isolate"`

	// 实例ip。
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 实例监听端口。
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 实例使用协议。不填默认为空。
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 实例版本。不填默认为空。
	InstanceVersion *string `json:"InstanceVersion,omitnil,omitempty" name:"InstanceVersion"`

	// 是否启用健康检查。不填默认不启用。
	EnableHealthCheck *bool `json:"EnableHealthCheck,omitnil,omitempty" name:"EnableHealthCheck"`

	// 上报心跳时间间隔。若 EnableHealthCheck 为不启用，则此参数不生效；若 EnableHealthCheck 启用，此参数不填，则默认 ttl 为 5s。
	Ttl *uint64 `json:"Ttl,omitnil,omitempty" name:"Ttl"`

	// 治理中心服务实例id。
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 元数据信息。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`
}

type GovernanceInterfaceDescription struct {
	// 契约接口ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 方法名称
	Method *string `json:"Method,omitnil,omitempty" name:"Method"`

	// 路径/接口名称
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 内容
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 创建来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// 信息摘要
	Revision *string `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 接口名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type GovernanceNamespace struct {
	// 命名空间名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 命名空间下总服务数据量
	TotalServiceCount *int64 `json:"TotalServiceCount,omitnil,omitempty" name:"TotalServiceCount"`

	// 命名空间下总健康实例数量
	TotalHealthInstanceCount *int64 `json:"TotalHealthInstanceCount,omitnil,omitempty" name:"TotalHealthInstanceCount"`

	// 命名空间下总实例数量
	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitnil,omitempty" name:"TotalInstanceCount"`

	// 命名空间ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 是否可以编辑
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 可以操作此命名空间的用户ID列表
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 可以操作此命名空间的用户组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 移除可以操作此命名空间的用户ID列表
	RemoveUserIds []*string `json:"RemoveUserIds,omitnil,omitempty" name:"RemoveUserIds"`

	// 移除可以操作此命名空间的用户组ID列表
	RemoveGroupIds []*string `json:"RemoveGroupIds,omitnil,omitempty" name:"RemoveGroupIds"`

	// 该命名空间下的服务对哪些命名空间可见
	ServiceExportTo []*string `json:"ServiceExportTo,omitnil,omitempty" name:"ServiceExportTo"`

	// 是否开启同步到全局注册中心	
	SyncToGlobalRegistry *bool `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`

	// 元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`
}

type GovernanceNamespaceInput struct {
	// 命名空间名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 新增的可以操作此命名空间的用户ID列表
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 新增的可以操作此命名空间的用户组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 移除可以操作此命名空间的用户ID列表
	RemoveUserIds []*string `json:"RemoveUserIds,omitnil,omitempty" name:"RemoveUserIds"`

	// 移除可以操作此命名空间的用户组ID列表
	RemoveGroupIds []*string `json:"RemoveGroupIds,omitnil,omitempty" name:"RemoveGroupIds"`

	// 该命名空间下的服务对哪些命名空间下可见，
	// 1、为空或者不填写，表示仅当前命名空间可见
	// 2、列表内容仅一个元素，且为字符 *，表示所有命名空间可见（包括新增）
	// 3、列表内容为部份命名空间名称，则只对这些命名空间下可见
	ServiceExportTo []*string `json:"ServiceExportTo,omitnil,omitempty" name:"ServiceExportTo"`

	// 是否开启同步到全局注册中心
	SyncToGlobalRegistry *bool `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`
}

type GovernanceService struct {
	// 服务名称。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 命名空间名称。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 元数据信息数组。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// 描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 创建时间。
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间。
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 服务所属部门。
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// 服务所属业务。
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// 健康服务实例数
	HealthyInstanceCount *uint64 `json:"HealthyInstanceCount,omitnil,omitempty" name:"HealthyInstanceCount"`

	// 服务实例总数
	TotalInstanceCount *uint64 `json:"TotalInstanceCount,omitnil,omitempty" name:"TotalInstanceCount"`

	// 服务ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 是否可以编辑
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 可以编辑该资源的用户ID
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 可以编辑该资源的用户组ID
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 移除可以编辑该资源的用户ID
	RemoveUserIds []*string `json:"RemoveUserIds,omitnil,omitempty" name:"RemoveUserIds"`

	// 移除可以编辑该资源的用户组ID
	RemoveGroupIds []*string `json:"RemoveGroupIds,omitnil,omitempty" name:"RemoveGroupIds"`

	// 该服务对哪些命名空间可见	
	ExportTo []*string `json:"ExportTo,omitnil,omitempty" name:"ExportTo"`

	// 该服务信息摘要签名
	Revision *string `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 是否开启同步到全局注册中心
	SyncToGlobalRegistry *bool `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`
}

type GovernanceServiceContract struct {
	// 契约名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 所属服务命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 契约ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 所属服务名称
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 信息摘要
	Revision *string `json:"Revision,omitnil,omitempty" name:"Revision"`

	// 额外内容描述
	Content *string `json:"Content,omitnil,omitempty" name:"Content"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 契约接口列表
	Interfaces []*GovernanceInterfaceDescription `json:"Interfaces,omitnil,omitempty" name:"Interfaces"`

	// 元数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`
}

type GovernanceServiceContractVersion struct {
	// 契约版本
	Version *string `json:"Version,omitnil,omitempty" name:"Version"`

	// 契约名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 唯一名称
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`
}

type GovernanceServiceInput struct {
	// 服务名。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 服务所属命名空间。
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务描述信息。
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`

	// 服务元数据。
	Metadatas []*Metadata `json:"Metadatas,omitnil,omitempty" name:"Metadatas"`

	// 服务所属部门。
	Department *string `json:"Department,omitnil,omitempty" name:"Department"`

	// 服务所属业务。
	Business *string `json:"Business,omitnil,omitempty" name:"Business"`

	// 被添加进来可以操作此命名空间的用户ID列表
	UserIds []*string `json:"UserIds,omitnil,omitempty" name:"UserIds"`

	// 被添加进来可以操作此命名空间的用户组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`

	// 从操作此命名空间的用户组ID列表被移除的ID列表
	RemoveUserIds []*string `json:"RemoveUserIds,omitnil,omitempty" name:"RemoveUserIds"`

	// 从可以操作此命名空间的用户组ID列表中被移除的ID列表
	RemoveGroupIds []*string `json:"RemoveGroupIds,omitnil,omitempty" name:"RemoveGroupIds"`

	// 该服务对哪些命名空间可见
	ExportTo []*string `json:"ExportTo,omitnil,omitempty" name:"ExportTo"`

	// 是否开启同步到全局注册中心
	SyncToGlobalRegistry *bool `json:"SyncToGlobalRegistry,omitnil,omitempty" name:"SyncToGlobalRegistry"`
}

type InstancePort struct {
	// 监听的 http 端口范围。
	HttpPort *string `json:"HttpPort,omitnil,omitempty" name:"HttpPort"`

	// 监听的 https 端口范围。
	HttpsPort *string `json:"HttpsPort,omitnil,omitempty" name:"HttpsPort"`

	// 监听的 tcp 端口范围。
	TcpPort *string `json:"TcpPort,omitnil,omitempty" name:"TcpPort"`

	// 监听的 udp 端口范围。
	UdpPort *string `json:"UdpPort,omitnil,omitempty" name:"UdpPort"`
}

type InstanceTagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitnil,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitnil,omitempty" name:"TagValue"`
}

type InternetConfig struct {
	// 公网地址版本，可选："IPV4" | "IPV6" 。不填默认 IPV4 。
	InternetAddressVersion *string `json:"InternetAddressVersion,omitnil,omitempty" name:"InternetAddressVersion"`

	// 公网付费类型，当前仅可选："BANDWIDTH"。不填默认为 "BANDWIDTH"
	InternetPayMode *string `json:"InternetPayMode,omitnil,omitempty" name:"InternetPayMode"`

	// 公网带宽。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 负载均衡描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 负载均衡的规格类型，支持clb.c2.medium、clb.c3.small、clb.c3.medium、clb.c4.small、clb.c4.medium、clb.c4.large、clb.c4.xlarge，不传为共享型。
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`

	// 负载均衡是否多可用区
	MultiZoneFlag *bool `json:"MultiZoneFlag,omitnil,omitempty" name:"MultiZoneFlag"`

	// 主可用区
	MasterZoneId *string `json:"MasterZoneId,omitnil,omitempty" name:"MasterZoneId"`

	// 备可用区
	SlaveZoneId *string `json:"SlaveZoneId,omitnil,omitempty" name:"SlaveZoneId"`
}

type KVMapping struct {
	// key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KVPair struct {
	// 键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KeyValue struct {
	// 条件的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 条件的Value
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type KongActiveHealthCheck struct {
	// 主动健康检查健康探测间隔，单位：秒，0表示不开启
	HealthyInterval *uint64 `json:"HealthyInterval,omitnil,omitempty" name:"HealthyInterval"`

	// 主动健康检查异常探测间隔，单位：秒，0表示不开启
	UnHealthyInterval *uint64 `json:"UnHealthyInterval,omitnil,omitempty" name:"UnHealthyInterval"`

	// 在 GET HTTP 请求中使用的路径，以作为主动运行状况检查的探测器运行。默认： ”/”。
	HttpPath *string `json:"HttpPath,omitnil,omitempty" name:"HttpPath"`

	// GET HTTP 请求的超时时间，单位：秒。默认 60。
	Timeout *float64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type KongCertificate struct {
	// 无
	Cert *KongCertificatesPreview `json:"Cert,omitnil,omitempty" name:"Cert"`
}

type KongCertificatesList struct {
	// 证书列表总数
	Total *int64 `json:"Total,omitnil,omitempty" name:"Total"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	CertificatesList []*KongCertificatesPreview `json:"CertificatesList,omitnil,omitempty" name:"CertificatesList"`

	// 证书列表总页数
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: Pages is deprecated.
	Pages *int64 `json:"Pages,omitnil,omitempty" name:"Pages"`
}

type KongCertificatesPreview struct {
	// 证书名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// Id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 绑定的域名
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// 证书状态：expired(已过期)
	//                    active(生效中)
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 证书pem格式
	Crt *string `json:"Crt,omitnil,omitempty" name:"Crt"`

	// 证书私钥
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 证书过期时间
	ExpireTime *string `json:"ExpireTime,omitnil,omitempty" name:"ExpireTime"`

	// 证书上传时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 证书签发时间
	IssueTime *string `json:"IssueTime,omitnil,omitempty" name:"IssueTime"`

	// 证书来源：native(kong自定义证书)
	//                     ssl(ssl平台证书)
	CertSource *string `json:"CertSource,omitnil,omitempty" name:"CertSource"`

	// ssl平台证书Id
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`
}

type KongPassiveHealthCheck struct {
	// 后端target协议类型，被动健康检查支持http和tcp，主动健康检查支持http
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`
}

type KongRoutePreview struct {
	// 服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 服务名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`

	// 无
	// 注意：此字段可能返回 null，表示取不到有效值。
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 无
	Protocols []*string `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// 无
	PreserveHost *bool `json:"PreserveHost,omitnil,omitempty" name:"PreserveHost"`

	// 无
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil,omitempty" name:"HttpsRedirectStatusCode"`

	// 无
	StripPath *bool `json:"StripPath,omitnil,omitempty" name:"StripPath"`

	// 无
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 是否开启了强制HTTPS
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: ForceHttps is deprecated.
	ForceHttps *bool `json:"ForceHttps,omitnil,omitempty" name:"ForceHttps"`

	// 服务名
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 服务ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 目的端口
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil,omitempty" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 是否缓存请求body，默认true
	RequestBuffering *bool `json:"RequestBuffering,omitnil,omitempty" name:"RequestBuffering"`

	// 是否缓存响应body，默认true
	ResponseBuffering *bool `json:"ResponseBuffering,omitnil,omitempty" name:"ResponseBuffering"`

	// 正则优先级
	RegexPriority *int64 `json:"RegexPriority,omitnil,omitempty" name:"RegexPriority"`
}

type KongServiceDetail struct {
	// 服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 服务名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 后端协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 后端路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 后端延时，单位ms
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// 标签
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 后端配置
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// 后端类型
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 是否可编辑
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type KongServiceLightPreview struct {
	// 服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 服务名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 后端配置
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// 后端类型
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`

	// 后端协议
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 重试次数
	Retries *uint64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// 后端延时，单位ms
	Timeout *uint64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`
}

type KongServicePreview struct {
	// 服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 服务名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 标签
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 后端配置
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// 后端类型
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 是否可编辑
	Editable *bool `json:"Editable,omitnil,omitempty" name:"Editable"`

	// 请求路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type KongServiceRouteList struct {
	// 无
	RouteList []*KongRoutePreview `json:"RouteList,omitnil,omitempty" name:"RouteList"`

	// 总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type KongServices struct {
	// kong实例的服务列表
	ServiceList []*KongServicePreview `json:"ServiceList,omitnil,omitempty" name:"ServiceList"`

	// 列表总数
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`
}

type KongTarget struct {
	// Host
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 权重
	Weight *int64 `json:"Weight,omitnil,omitempty" name:"Weight"`

	// 健康状态
	Health *string `json:"Health,omitnil,omitempty" name:"Health"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// Target的来源
	Source *string `json:"Source,omitnil,omitempty" name:"Source"`

	// CVM实例ID
	CvmInstanceId *string `json:"CvmInstanceId,omitnil,omitempty" name:"CvmInstanceId"`

	// CVM实例名称
	CvmInstanceName *string `json:"CvmInstanceName,omitnil,omitempty" name:"CvmInstanceName"`

	// target标签
	Tags []*string `json:"Tags,omitnil,omitempty" name:"Tags"`
}

type KongUpstreamInfo struct {
	// IP或域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 服务来源ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务（注册中心或Kubernetes中的服务）名字
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 服务后端类型是IPList时提供
	// 注意：此字段可能返回 null，表示取不到有效值。
	Targets []*KongTarget `json:"Targets,omitnil,omitempty" name:"Targets"`

	// 服务来源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// SCF函数类型
	ScfType *string `json:"ScfType,omitnil,omitempty" name:"ScfType"`

	// SCF函数命名空间
	ScfNamespace *string `json:"ScfNamespace,omitnil,omitempty" name:"ScfNamespace"`

	// SCF函数名
	ScfLambdaName *string `json:"ScfLambdaName,omitnil,omitempty" name:"ScfLambdaName"`

	// SCF函数版本
	ScfLambdaQualifier *string `json:"ScfLambdaQualifier,omitnil,omitempty" name:"ScfLambdaQualifier"`

	// 冷启动时间，单位秒
	SlowStart *int64 `json:"SlowStart,omitnil,omitempty" name:"SlowStart"`

	// 负载均衡算法，默认为 round-robin，还支持 least-connections，consisten_hashing
	Algorithm *string `json:"Algorithm,omitnil,omitempty" name:"Algorithm"`

	// CVM弹性伸缩组ID
	AutoScalingGroupID *string `json:"AutoScalingGroupID,omitnil,omitempty" name:"AutoScalingGroupID"`

	// CVM弹性伸缩组端口
	AutoScalingCvmPort *uint64 `json:"AutoScalingCvmPort,omitnil,omitempty" name:"AutoScalingCvmPort"`

	// CVM弹性伸缩组使用的CVM TAT命令状态
	AutoScalingTatCmdStatus *string `json:"AutoScalingTatCmdStatus,omitnil,omitempty" name:"AutoScalingTatCmdStatus"`

	// CVM弹性伸缩组生命周期挂钩状态
	AutoScalingHookStatus *string `json:"AutoScalingHookStatus,omitnil,omitempty" name:"AutoScalingHookStatus"`

	// 服务来源的名字
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 精确的服务来源类型，新建服务来源时候传入的类型
	RealSourceType *string `json:"RealSourceType,omitnil,omitempty" name:"RealSourceType"`

	// upstream健康状态HEALTHY（健康）, UNHEALTHY（异常）, HEALTHCHECKS_OFF（未开启）和NONE（不支持健康检查）
	HealthStatus *string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 云函数是否开启CAM鉴权，不填时默认为开启(true)
	ScfCamAuthEnable *bool `json:"ScfCamAuthEnable,omitnil,omitempty" name:"ScfCamAuthEnable"`

	// 云函数是否开启Base64编码，默认为false
	ScfIsBase64Encoded *bool `json:"ScfIsBase64Encoded,omitnil,omitempty" name:"ScfIsBase64Encoded"`

	// 云函数是否开启响应集成，默认为false
	ScfIsIntegratedResponse *bool `json:"ScfIsIntegratedResponse,omitnil,omitempty" name:"ScfIsIntegratedResponse"`
}

type KongUpstreamList struct {
	// 无
	UpstreamList []*KongUpstreamPreview `json:"UpstreamList,omitnil,omitempty" name:"UpstreamList"`
}

type KongUpstreamPreview struct {
	// 服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 服务名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 后端配置
	Target []*KongTarget `json:"Target,omitnil,omitempty" name:"Target"`
}

type LimitRule struct {
	// 请求匹配条件
	Filters []*RuleFilter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 参数限流依据组合
	LimitBy []*KeyValue `json:"LimitBy,omitnil,omitempty" name:"LimitBy"`

	// 限流阈值
	QpsThresholds []*QpsThreshold `json:"QpsThresholds,omitnil,omitempty" name:"QpsThresholds"`

	// 精确限流阈值
	AccurateQpsThresholds []*AccurateQpsThreshold `json:"AccurateQpsThresholds,omitnil,omitempty" name:"AccurateQpsThresholds"`
}

type ListCloudNativeAPIGatewayResult struct {
	// 总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云原生API网关实例列表。
	GatewayList []*DescribeCloudNativeAPIGatewayResult `json:"GatewayList,omitnil,omitempty" name:"GatewayList"`
}

type ListCloudNativeAPIGatewayStrategyBindingGroupInfoResult struct {
	// 数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云原生API网关实例策略绑定网关分组列表
	GroupInfos []*CloudNativeAPIGatewayStrategyBindingGroupInfo `json:"GroupInfos,omitnil,omitempty" name:"GroupInfos"`
}

type ListCloudNativeAPIGatewayStrategyResult struct {
	// 总数。
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 云原生API网关实例策略列表。
	StrategyList []*CloudNativeAPIGatewayStrategy `json:"StrategyList,omitnil,omitempty" name:"StrategyList"`
}

type ListFilter struct {
	// 过滤字段
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 过滤值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type Location struct {
	// 大区
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 机房
	Campus *string `json:"Campus,omitnil,omitempty" name:"Campus"`
}

type Metadata struct {
	// 元数据键名。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 元数据键值。不填则默认为空字符串。
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

// Predefined struct for user
type ModifyAutoScalerResourceStrategyRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 指标伸缩配置
	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 定时伸缩配置
	//
	// Deprecated: CronScalerConfig is deprecated.
	CronScalerConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronScalerConfig,omitnil,omitempty" name:"CronScalerConfig"`

	// 最大节点数
	//
	// Deprecated: MaxReplicas is deprecated.
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 指标伸缩配置
	CronConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronConfig,omitnil,omitempty" name:"CronConfig"`
}

type ModifyAutoScalerResourceStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 策略名称
	StrategyName *string `json:"StrategyName,omitnil,omitempty" name:"StrategyName"`

	// 策略描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 指标伸缩配置
	Config *CloudNativeAPIGatewayStrategyAutoScalerConfig `json:"Config,omitnil,omitempty" name:"Config"`

	// 定时伸缩配置
	CronScalerConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronScalerConfig,omitnil,omitempty" name:"CronScalerConfig"`

	// 最大节点数
	MaxReplicas *int64 `json:"MaxReplicas,omitnil,omitempty" name:"MaxReplicas"`

	// 指标伸缩配置
	CronConfig *CloudNativeAPIGatewayStrategyCronScalerConfig `json:"CronConfig,omitnil,omitempty" name:"CronConfig"`
}

func (r *ModifyAutoScalerResourceStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScalerResourceStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StrategyId")
	delete(f, "StrategyName")
	delete(f, "Description")
	delete(f, "Config")
	delete(f, "CronScalerConfig")
	delete(f, "MaxReplicas")
	delete(f, "CronConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoScalerResourceStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoScalerResourceStrategyResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoScalerResourceStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoScalerResourceStrategyResponseParams `json:"Response"`
}

func (r *ModifyAutoScalerResourceStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScalerResourceStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayCanaryRuleRequestParams struct {
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 优先级，同一个服务的灰度规则优先级是唯一的
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil,omitempty" name:"CanaryRule"`

	// 灰度规则配置列表，如果配置了此参数，将以此参数为准，忽略Priority和CanaryRule参数
	CanaryRuleList []*CanaryPriorityRule `json:"CanaryRuleList,omitnil,omitempty" name:"CanaryRuleList"`
}

type ModifyCloudNativeAPIGatewayCanaryRuleRequest struct {
	*tchttp.BaseRequest
	
	// 网关 ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务 ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`

	// 优先级，同一个服务的灰度规则优先级是唯一的
	Priority *int64 `json:"Priority,omitnil,omitempty" name:"Priority"`

	// 灰度规则配置
	CanaryRule *CloudNativeAPIGatewayCanaryRule `json:"CanaryRule,omitnil,omitempty" name:"CanaryRule"`

	// 灰度规则配置列表，如果配置了此参数，将以此参数为准，忽略Priority和CanaryRule参数
	CanaryRuleList []*CanaryPriorityRule `json:"CanaryRuleList,omitnil,omitempty" name:"CanaryRuleList"`
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
	delete(f, "CanaryRuleList")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayCanaryRuleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayCanaryRuleResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyCloudNativeAPIGatewayCertificateRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 证书名称，即将废弃
	//
	// Deprecated: Name is deprecated.
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证书私钥，CertSource为native时必填。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 证书pem格式，CertSource为native时必填。
	Crt *string `json:"Crt,omitnil,omitempty" name:"Crt"`

	// 绑定的域名，即将废弃
	//
	// Deprecated: BindDomains is deprecated.
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// ssl平台证书 Id，CertSource为ssl时必填。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书来源
	// - ssl (ssl平台证书)，默认值
	// - native (kong自定义证书) 
	CertSource *string `json:"CertSource,omitnil,omitempty" name:"CertSource"`
}

type ModifyCloudNativeAPIGatewayCertificateRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 证书名称，即将废弃
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 证书私钥，CertSource为native时必填。
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 证书pem格式，CertSource为native时必填。
	Crt *string `json:"Crt,omitnil,omitempty" name:"Crt"`

	// 绑定的域名，即将废弃
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// ssl平台证书 Id，CertSource为ssl时必填。
	CertId *string `json:"CertId,omitnil,omitempty" name:"CertId"`

	// 证书来源
	// - ssl (ssl平台证书)，默认值
	// - native (kong自定义证书) 
	CertSource *string `json:"CertSource,omitnil,omitempty" name:"CertSource"`
}

func (r *ModifyCloudNativeAPIGatewayCertificateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayCertificateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Id")
	delete(f, "Name")
	delete(f, "Key")
	delete(f, "Crt")
	delete(f, "BindDomains")
	delete(f, "CertId")
	delete(f, "CertSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayCertificateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayCertificateResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyCloudNativeAPIGatewayCertificateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCloudNativeAPIGatewayCertificateResponseParams `json:"Response"`
}

func (r *ModifyCloudNativeAPIGatewayCertificateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCloudNativeAPIGatewayCertificateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否开启 CLS 日志。暂时取值只能是 true，即只能从关闭状态变成开启状态。
	EnableCls *bool `json:"EnableCls,omitnil,omitempty" name:"EnableCls"`

	// 公网计费模式。可选取值 BANDWIDTH | TRAFFIC ，表示按带宽和按流量计费。
	InternetPayMode *string `json:"InternetPayMode,omitnil,omitempty" name:"InternetPayMode"`

	// 是否开启实例删除保护,默认false
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`
}

type ModifyCloudNativeAPIGatewayRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 是否开启 CLS 日志。暂时取值只能是 true，即只能从关闭状态变成开启状态。
	EnableCls *bool `json:"EnableCls,omitnil,omitempty" name:"EnableCls"`

	// 公网计费模式。可选取值 BANDWIDTH | TRAFFIC ，表示按带宽和按流量计费。
	InternetPayMode *string `json:"InternetPayMode,omitnil,omitempty" name:"InternetPayMode"`

	// 是否开启实例删除保护,默认false
	DeleteProtect *bool `json:"DeleteProtect,omitnil,omitempty" name:"DeleteProtect"`
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
	delete(f, "DeleteProtect")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
}

type ModifyCloudNativeAPIGatewayRouteRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 路由id，或路由名称。
	// 不支持“未命名”
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 路由的ID，实例级别唯一
	RouteID *string `json:"RouteID,omitnil,omitempty" name:"RouteID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

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
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil,omitempty" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil,omitempty" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil,omitempty" name:"StripPath"`

	// 是否开启强制HTTPS
	//
	// Deprecated: ForceHttps is deprecated.
	ForceHttps *bool `json:"ForceHttps,omitnil,omitempty" name:"ForceHttps"`

	// 四层匹配的目的端口	
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil,omitempty" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 是否缓存请求body，默认true
	RequestBuffering *bool `json:"RequestBuffering,omitnil,omitempty" name:"RequestBuffering"`

	// 是否缓存响应body，默认true
	ResponseBuffering *bool `json:"ResponseBuffering,omitnil,omitempty" name:"ResponseBuffering"`

	// 增加优先级
	RegexPriority *int64 `json:"RegexPriority,omitnil,omitempty" name:"RegexPriority"`
}

type ModifyCloudNativeAPIGatewayRouteRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 所属服务的ID
	ServiceID *string `json:"ServiceID,omitnil,omitempty" name:"ServiceID"`

	// 路由的ID，实例级别唯一
	RouteID *string `json:"RouteID,omitnil,omitempty" name:"RouteID"`

	// 路由的名字，实例级别唯一，可以不提供
	RouteName *string `json:"RouteName,omitnil,omitempty" name:"RouteName"`

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
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 路由的域名
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 路由的路径
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`

	// 路由的协议，可选
	// - https
	// - http
	Protocols []*string `json:"Protocols,omitnil,omitempty" name:"Protocols"`

	// 转发到后端时是否保留Host
	PreserveHost *bool `json:"PreserveHost,omitnil,omitempty" name:"PreserveHost"`

	// https重定向状态码
	HttpsRedirectStatusCode *int64 `json:"HttpsRedirectStatusCode,omitnil,omitempty" name:"HttpsRedirectStatusCode"`

	// 转发到后端时是否StripPath
	StripPath *bool `json:"StripPath,omitnil,omitempty" name:"StripPath"`

	// 是否开启强制HTTPS
	ForceHttps *bool `json:"ForceHttps,omitnil,omitempty" name:"ForceHttps"`

	// 四层匹配的目的端口	
	DestinationPorts []*uint64 `json:"DestinationPorts,omitnil,omitempty" name:"DestinationPorts"`

	// 路由的Headers
	Headers []*KVMapping `json:"Headers,omitnil,omitempty" name:"Headers"`

	// 是否缓存请求body，默认true
	RequestBuffering *bool `json:"RequestBuffering,omitnil,omitempty" name:"RequestBuffering"`

	// 是否缓存响应body，默认true
	ResponseBuffering *bool `json:"ResponseBuffering,omitnil,omitempty" name:"ResponseBuffering"`

	// 增加优先级
	RegexPriority *int64 `json:"RegexPriority,omitnil,omitempty" name:"RegexPriority"`
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
	delete(f, "RequestBuffering")
	delete(f, "ResponseBuffering")
	delete(f, "RegexPriority")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayRouteRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayRouteResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
}

type ModifyCloudNativeAPIGatewayServiceRateLimitRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称，或服务ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 限流配置
	LimitDetail *CloudNativeAPIGatewayRateLimitDetail `json:"LimitDetail,omitnil,omitempty" name:"LimitDetail"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 请求协议： 
	// - https 
	// - http 
	// - tcp 
	// - udp
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 服务配置
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// 服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 请求路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
}

type ModifyCloudNativeAPIGatewayServiceRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 请求协议： 
	// - https 
	// - http 
	// - tcp 
	// - udp
	Protocol *string `json:"Protocol,omitnil,omitempty" name:"Protocol"`

	// 超时时间，单位ms
	Timeout *int64 `json:"Timeout,omitnil,omitempty" name:"Timeout"`

	// 重试次数
	Retries *int64 `json:"Retries,omitnil,omitempty" name:"Retries"`

	// 服务类型: 
	// - Kubernetes 
	// - Registry
	// - IPList
	// - HostIP
	// - Scf	
	UpstreamType *string `json:"UpstreamType,omitnil,omitempty" name:"UpstreamType"`

	// 服务配置
	UpstreamInfo *KongUpstreamInfo `json:"UpstreamInfo,omitnil,omitempty" name:"UpstreamInfo"`

	// 服务ID
	ID *string `json:"ID,omitnil,omitempty" name:"ID"`

	// 请求路径
	Path *string `json:"Path,omitnil,omitempty" name:"Path"`
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
	delete(f, "Timeout")
	delete(f, "Retries")
	delete(f, "UpstreamType")
	delete(f, "UpstreamInfo")
	delete(f, "ID")
	delete(f, "Path")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCloudNativeAPIGatewayServiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCloudNativeAPIGatewayServiceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
type ModifyConfigFileGroupRequestParams struct {
	// tse实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件组
	ConfigFileGroup *ConfigFileGroup `json:"ConfigFileGroup,omitnil,omitempty" name:"ConfigFileGroup"`
}

type ModifyConfigFileGroupRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件组
	ConfigFileGroup *ConfigFileGroup `json:"ConfigFileGroup,omitnil,omitempty" name:"ConfigFileGroup"`
}

func (r *ModifyConfigFileGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigFileGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigFileGroup")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigFileGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigFileGroupResponseParams struct {
	// 修改是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConfigFileGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigFileGroupResponseParams `json:"Response"`
}

func (r *ModifyConfigFileGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigFileGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigFilesRequestParams struct {
	// ins-df344df5	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件列表
	ConfigFile *ConfigFile `json:"ConfigFile,omitnil,omitempty" name:"ConfigFile"`
}

type ModifyConfigFilesRequest struct {
	*tchttp.BaseRequest
	
	// ins-df344df5	
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件列表
	ConfigFile *ConfigFile `json:"ConfigFile,omitnil,omitempty" name:"ConfigFile"`
}

func (r *ModifyConfigFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigFile")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConfigFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConfigFilesResponseParams struct {
	// 修改是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConfigFilesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConfigFilesResponseParams `json:"Response"`
}

func (r *ModifyConfigFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsoleNetworkRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网络类型：
	// - Open 公网
	// - Internal 内网（暂不支持）
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 开启Konga网络，不填时默认为Open
	// - Open，开启
	// - Close，关闭
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 访问控制策略
	AccessControl *NetworkAccessControl `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`
}

type ModifyConsoleNetworkRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网络类型：
	// - Open 公网
	// - Internal 内网（暂不支持）
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// 开启Konga网络，不填时默认为Open
	// - Open，开启
	// - Close，关闭
	Operate *string `json:"Operate,omitnil,omitempty" name:"Operate"`

	// 访问控制策略
	AccessControl *NetworkAccessControl `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`
}

func (r *ModifyConsoleNetworkRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsoleNetworkRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "NetworkType")
	delete(f, "Operate")
	delete(f, "AccessControl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyConsoleNetworkRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyConsoleNetworkResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyConsoleNetworkResponse struct {
	*tchttp.BaseResponse
	Response *ModifyConsoleNetworkResponseParams `json:"Response"`
}

func (r *ModifyConsoleNetworkResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyConsoleNetworkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceAliasRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 服务别名命名空间
	AliasNamespace *string `json:"AliasNamespace,omitnil,omitempty" name:"AliasNamespace"`

	// 服务别名所指向的服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务别名所指向的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务别名描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

type ModifyGovernanceAliasRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务别名
	Alias *string `json:"Alias,omitnil,omitempty" name:"Alias"`

	// 服务别名命名空间
	AliasNamespace *string `json:"AliasNamespace,omitnil,omitempty" name:"AliasNamespace"`

	// 服务别名所指向的服务名
	Service *string `json:"Service,omitnil,omitempty" name:"Service"`

	// 服务别名所指向的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 服务别名描述
	Comment *string `json:"Comment,omitnil,omitempty" name:"Comment"`
}

func (r *ModifyGovernanceAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceAliasRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Alias")
	delete(f, "AliasNamespace")
	delete(f, "Service")
	delete(f, "Namespace")
	delete(f, "Comment")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGovernanceAliasRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceAliasResponseParams struct {
	// 创建是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGovernanceAliasResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGovernanceAliasResponseParams `json:"Response"`
}

func (r *ModifyGovernanceAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceAliasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceInstancesRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务实例信息。
	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

type ModifyGovernanceInstancesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务实例信息。
	GovernanceInstances []*GovernanceInstanceUpdate `json:"GovernanceInstances,omitnil,omitempty" name:"GovernanceInstances"`
}

func (r *ModifyGovernanceInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceInstances")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGovernanceInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceInstancesResponseParams struct {
	// 修改是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGovernanceInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGovernanceInstancesResponseParams `json:"Response"`
}

func (r *ModifyGovernanceInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceNamespacesRequestParams struct {
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间信息。
	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitnil,omitempty" name:"GovernanceNamespaces"`
}

type ModifyGovernanceNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// tse实例id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 命名空间信息。
	GovernanceNamespaces []*GovernanceNamespaceInput `json:"GovernanceNamespaces,omitnil,omitempty" name:"GovernanceNamespaces"`
}

func (r *ModifyGovernanceNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceNamespaces")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGovernanceNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceNamespacesResponseParams struct {
	// 操作是否成功。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGovernanceNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGovernanceNamespacesResponseParams `json:"Response"`
}

func (r *ModifyGovernanceNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceServicesRequestParams struct {
	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务信息。
	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitnil,omitempty" name:"GovernanceServices"`
}

type ModifyGovernanceServicesRequest struct {
	*tchttp.BaseRequest
	
	// tse 实例 id。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 服务信息。
	GovernanceServices []*GovernanceServiceInput `json:"GovernanceServices,omitnil,omitempty" name:"GovernanceServices"`
}

func (r *ModifyGovernanceServicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceServicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "GovernanceServices")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyGovernanceServicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyGovernanceServicesResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyGovernanceServicesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyGovernanceServicesResponseParams `json:"Response"`
}

func (r *ModifyGovernanceServicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyGovernanceServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNativeGatewayServerGroupRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组 id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

type ModifyNativeGatewayServerGroupRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组 id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 云原生API网关名字, 最多支持60个字符。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 云原生API网关描述信息, 最多支持120个字符。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type ModifyNativeGatewayServiceSourceRequestParams struct {
	// 网关实例ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 服务来源实例ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 服务来源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 服务来源实例额外信息
	SourceInfo *SourceInfo `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

type ModifyNativeGatewayServiceSourceRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 服务来源实例ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 服务来源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 服务来源实例额外信息
	SourceInfo *SourceInfo `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`
}

func (r *ModifyNativeGatewayServiceSourceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNativeGatewayServiceSourceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayID")
	delete(f, "SourceID")
	delete(f, "SourceName")
	delete(f, "SourceInfo")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNativeGatewayServiceSourceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNativeGatewayServiceSourceResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNativeGatewayServiceSourceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNativeGatewayServiceSourceResponseParams `json:"Response"`
}

func (r *ModifyNativeGatewayServiceSourceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNativeGatewayServiceSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAccessStrategyRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网络类型： 
	// - Open 公网
	// - Internal 内网	（暂不支持）
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// ip地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 访问控制策略
	AccessControl *NetworkAccessControl `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`
}

type ModifyNetworkAccessStrategyRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网络类型： 
	// - Open 公网
	// - Internal 内网	（暂不支持）
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// ip地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 访问控制策略
	AccessControl *NetworkAccessControl `json:"AccessControl,omitnil,omitempty" name:"AccessControl"`
}

func (r *ModifyNetworkAccessStrategyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAccessStrategyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "NetworkType")
	delete(f, "Vip")
	delete(f, "AccessControl")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkAccessStrategyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkAccessStrategyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkAccessStrategyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkAccessStrategyResponseParams `json:"Response"`
}

func (r *ModifyNetworkAccessStrategyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkAccessStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkBasicInfoRequestParams struct {
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网络类型：
	// - Open 公网ipv4
	// - Open-IPv6 公网ipv6
	// - Internal 内网
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// ip地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 公网出流量带宽[1,2048]Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 负载均衡描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 负载均衡的规格类型，支持：
	// - 不传为共享型。
	// - clb.c2.medium：标准型规格
	// - clb.c3.small：高阶型1规格
	// - clb.c3.medium：高阶型2规格
	// - clb.c4.small：超强型1规格
	// - clb.c4.medium：超强型2规格
	// - clb.c4.large：超强型3规格
	// - clb.c4.xlarge：超强型4规格
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`
}

type ModifyNetworkBasicInfoRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网络类型：
	// - Open 公网ipv4
	// - Open-IPv6 公网ipv6
	// - Internal 内网
	NetworkType *string `json:"NetworkType,omitnil,omitempty" name:"NetworkType"`

	// ip地址
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 公网出流量带宽[1,2048]Mbps
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 负载均衡描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 负载均衡的规格类型，支持：
	// - 不传为共享型。
	// - clb.c2.medium：标准型规格
	// - clb.c3.small：高阶型1规格
	// - clb.c3.medium：高阶型2规格
	// - clb.c4.small：超强型1规格
	// - clb.c4.medium：超强型2规格
	// - clb.c4.large：超强型3规格
	// - clb.c4.xlarge：超强型4规格
	SlaType *string `json:"SlaType,omitnil,omitempty" name:"SlaType"`
}

func (r *ModifyNetworkBasicInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkBasicInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "GroupId")
	delete(f, "NetworkType")
	delete(f, "Vip")
	delete(f, "InternetMaxBandwidthOut")
	delete(f, "Description")
	delete(f, "SlaType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNetworkBasicInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNetworkBasicInfoResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNetworkBasicInfoResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNetworkBasicInfoResponseParams `json:"Response"`
}

func (r *ModifyNetworkBasicInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNetworkBasicInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUpstreamNodeStatusRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 访问IP地址或域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 访问端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// HEALTHY或UNHEALTHY
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type ModifyUpstreamNodeStatusRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 访问IP地址或域名
	Host *string `json:"Host,omitnil,omitempty" name:"Host"`

	// 访问端口
	Port *int64 `json:"Port,omitnil,omitempty" name:"Port"`

	// HEALTHY或UNHEALTHY
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

func (r *ModifyUpstreamNodeStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUpstreamNodeStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "ServiceName")
	delete(f, "Host")
	delete(f, "Port")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUpstreamNodeStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUpstreamNodeStatusResponseParams struct {
	// 是否成功
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyUpstreamNodeStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUpstreamNodeStatusResponseParams `json:"Response"`
}

func (r *ModifyUpstreamNodeStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUpstreamNodeStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NacosReplica struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 可用区ID
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// VPC ID	
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type NacosServerInterface struct {
	// 接口名
	Interface *string `json:"Interface,omitnil,omitempty" name:"Interface"`
}

type NativeGatewayServerGroup struct {
	// 云原生网关分组唯一id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 分组名
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 描述信息
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 节点规格、节点数信息
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`

	// 网关分组状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 是否是默认分组。
	// 0：否。
	// 1：是。
	IsFirstGroup *int64 `json:"IsFirstGroup,omitnil,omitempty" name:"IsFirstGroup"`

	// 关联策略信息
	BindingStrategy *CloudNativeAPIGatewayStrategy `json:"BindingStrategy,omitnil,omitempty" name:"BindingStrategy"`

	// 网关实例 id
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 带宽
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`

	// 子网id
	SubnetIds *string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 分组默认权重
	DefaultWeight *int64 `json:"DefaultWeight,omitnil,omitempty" name:"DefaultWeight"`

	// 弹性节点
	ElasticNumber *uint64 `json:"ElasticNumber,omitnil,omitempty" name:"ElasticNumber"`

	// 是否支持TOA
	SupportTOA *bool `json:"SupportTOA,omitnil,omitempty" name:"SupportTOA"`

	// 是否支持IPV6
	SupportIPV6 *bool `json:"SupportIPV6,omitnil,omitempty" name:"SupportIPV6"`
}

type NativeGatewayServerGroups struct {
	// 总数
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 分组信息数组。
	GatewayGroupList []*NativeGatewayServerGroup `json:"GatewayGroupList,omitnil,omitempty" name:"GatewayGroupList"`
}

type NativeGatewayServiceSourceItem struct {
	// 网关实例ID
	GatewayID *string `json:"GatewayID,omitnil,omitempty" name:"GatewayID"`

	// 服务来源ID
	SourceID *string `json:"SourceID,omitnil,omitempty" name:"SourceID"`

	// 服务来源名称
	SourceName *string `json:"SourceName,omitnil,omitempty" name:"SourceName"`

	// 服务来源类型
	SourceType *string `json:"SourceType,omitnil,omitempty" name:"SourceType"`

	// 服务来源额外信息
	SourceInfo *SourceInfo `json:"SourceInfo,omitnil,omitempty" name:"SourceInfo"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 修改时间
	ModifyTime *string `json:"ModifyTime,omitnil,omitempty" name:"ModifyTime"`
}

type NetworkAccessControl struct {
	// 访问模式：Whitelist|Blacklist
	Mode *string `json:"Mode,omitnil,omitempty" name:"Mode"`

	// 白名单列表
	CidrWhiteList []*string `json:"CidrWhiteList,omitnil,omitempty" name:"CidrWhiteList"`

	// 黑名单列表
	CidrBlackList []*string `json:"CidrBlackList,omitnil,omitempty" name:"CidrBlackList"`
}

// Predefined struct for user
type OpenWafProtectionRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	//  防护资源的类型。
	// - Global  实例
	// - Service  服务
	// - Route  路由
	// - Object  对象（接口暂不支持）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当资源类型 Type 是 Service 或 Route 的时候，传入的服务或路由的列表
	List []*string `json:"List,omitnil,omitempty" name:"List"`
}

type OpenWafProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	//  防护资源的类型。
	// - Global  实例
	// - Service  服务
	// - Route  路由
	// - Object  对象（接口暂不支持）
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 当资源类型 Type 是 Service 或 Route 的时候，传入的服务或路由的列表
	List []*string `json:"List,omitnil,omitempty" name:"List"`
}

func (r *OpenWafProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWafProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Type")
	delete(f, "List")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "OpenWafProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type OpenWafProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type OpenWafProtectionResponse struct {
	*tchttp.BaseResponse
	Response *OpenWafProtectionResponseParams `json:"Response"`
}

func (r *OpenWafProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *OpenWafProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolarisCLSTopicInfo struct {
	// 日志集ID
	LogSetId *string `json:"LogSetId,omitnil,omitempty" name:"LogSetId"`

	// 日志集名称
	LogSetName *string `json:"LogSetName,omitnil,omitempty" name:"LogSetName"`

	// 日志主题ID
	TopicId *string `json:"TopicId,omitnil,omitempty" name:"TopicId"`

	// 日志主题名称
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type PolarisLimiterAddress struct {
	// VPC接入IP列表
	IntranetAddress *string `json:"IntranetAddress,omitnil,omitempty" name:"IntranetAddress"`
}

type PublicAddressConfig struct {
	// 公网 ip
	Vip *string `json:"Vip,omitnil,omitempty" name:"Vip"`

	// 公网最大带宽
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 公网所属分组 id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 公网所属分组名
	GroupName *string `json:"GroupName,omitnil,omitempty" name:"GroupName"`

	// 公网负载均衡 id
	NetworkId *string `json:"NetworkId,omitnil,omitempty" name:"NetworkId"`

	// 公网负载均衡描述
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`
}

// Predefined struct for user
type PublishConfigFilesRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件发布
	ConfigFileReleases *ConfigFileRelease `json:"ConfigFileReleases,omitnil,omitempty" name:"ConfigFileReleases"`

	// 控制开启校验配置版本是否已经存在
	StrictEnable *bool `json:"StrictEnable,omitnil,omitempty" name:"StrictEnable"`
}

type PublishConfigFilesRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 配置文件发布
	ConfigFileReleases *ConfigFileRelease `json:"ConfigFileReleases,omitnil,omitempty" name:"ConfigFileReleases"`

	// 控制开启校验配置版本是否已经存在
	StrictEnable *bool `json:"StrictEnable,omitnil,omitempty" name:"StrictEnable"`
}

func (r *PublishConfigFilesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishConfigFilesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "ConfigFileReleases")
	delete(f, "StrictEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishConfigFilesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type PublishConfigFilesResponseParams struct {
	// 配置文件发布是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 配置文件发布Id
	ConfigFileReleaseId *string `json:"ConfigFileReleaseId,omitnil,omitempty" name:"ConfigFileReleaseId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type PublishConfigFilesResponse struct {
	*tchttp.BaseResponse
	Response *PublishConfigFilesResponseParams `json:"Response"`
}

func (r *PublishConfigFilesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishConfigFilesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QpsThreshold struct {
	// qps阈值控制维度,包含:second、minute、hour、day、month、year
	Unit *string `json:"Unit,omitnil,omitempty" name:"Unit"`

	// 阈值
	Max *int64 `json:"Max,omitnil,omitempty" name:"Max"`
}

type RateLimitResponse struct {
	// 自定义响应体
	Body *string `json:"Body,omitnil,omitempty" name:"Body"`

	// Headers
	Headers []*KVMapping `json:"Headers,omitnil,omitempty" name:"Headers"`

	// http状态码
	HttpStatus *int64 `json:"HttpStatus,omitnil,omitempty" name:"HttpStatus"`
}

type ReleaseVersion struct {
	// 配置发布的版本
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 是否生效
	Active *bool `json:"Active,omitnil,omitempty" name:"Active"`

	// 配置发布的ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 配置发布的命名空间
	Namespace *string `json:"Namespace,omitnil,omitempty" name:"Namespace"`

	// 配置发布的分组
	Group *string `json:"Group,omitnil,omitempty" name:"Group"`

	// 配置发布的文件名
	FileName *string `json:"FileName,omitnil,omitempty" name:"FileName"`
}

// Predefined struct for user
type RestartSREInstanceRequestParams struct {
	// 微服务引擎实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重启的环境类型（PROD，DEV，UAT等）
	EnvTypes []*string `json:"EnvTypes,omitnil,omitempty" name:"EnvTypes"`

	// 指定需要重启的实例节点（当前仅支持zk单节点重启）
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

type RestartSREInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 微服务引擎实例Id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 重启的环境类型（PROD，DEV，UAT等）
	EnvTypes []*string `json:"EnvTypes,omitnil,omitempty" name:"EnvTypes"`

	// 指定需要重启的实例节点（当前仅支持zk单节点重启）
	NodeName *string `json:"NodeName,omitnil,omitempty" name:"NodeName"`
}

func (r *RestartSREInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartSREInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "EnvTypes")
	delete(f, "NodeName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RestartSREInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RestartSREInstanceResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RestartSREInstanceResponse struct {
	*tchttp.BaseResponse
	Response *RestartSREInstanceResponseParams `json:"Response"`
}

func (r *RestartSREInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RestartSREInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackConfigFileReleasesRequestParams struct {
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 回滚发布
	RollbackConfigFileReleases []*ConfigFileRelease `json:"RollbackConfigFileReleases,omitnil,omitempty" name:"RollbackConfigFileReleases"`
}

type RollbackConfigFileReleasesRequest struct {
	*tchttp.BaseRequest
	
	// TSE实例id
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 回滚发布
	RollbackConfigFileReleases []*ConfigFileRelease `json:"RollbackConfigFileReleases,omitnil,omitempty" name:"RollbackConfigFileReleases"`
}

func (r *RollbackConfigFileReleasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackConfigFileReleasesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "RollbackConfigFileReleases")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackConfigFileReleasesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackConfigFileReleasesResponseParams struct {
	// 回滚结果
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackConfigFileReleasesResponse struct {
	*tchttp.BaseResponse
	Response *RollbackConfigFileReleasesResponseParams `json:"Response"`
}

func (r *RollbackConfigFileReleasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackConfigFileReleasesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteWafStatus struct {
	// 路由的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 路由的 ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	//  路由是否开启 WAF 防护
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 方法
	Methods []*string `json:"Methods,omitnil,omitempty" name:"Methods"`

	// 路径
	Paths []*string `json:"Paths,omitnil,omitempty" name:"Paths"`

	// 域名
	Hosts []*string `json:"Hosts,omitnil,omitempty" name:"Hosts"`

	// 路由对应服务的名字
	ServiceName *string `json:"ServiceName,omitnil,omitempty" name:"ServiceName"`

	// 路由对应服务的ID
	ServiceId *string `json:"ServiceId,omitnil,omitempty" name:"ServiceId"`
}

type RuleFilter struct {
	// 限流条件的Key
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 限流条件的Values
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`

	// 操作符
	Operator *string `json:"Operator,omitnil,omitempty" name:"Operator"`

	// header或query对应的name
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type SREInstance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 版本号
	Edition *string `json:"Edition,omitnil,omitempty" name:"Edition"`

	// 状态, 枚举值:creating/create_fail/running/updating/update_fail/restarting/restart_fail/destroying/destroy_fail
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 规格ID
	SpecId *string `json:"SpecId,omitnil,omitempty" name:"SpecId"`

	// 副本数
	Replica *int64 `json:"Replica,omitnil,omitempty" name:"Replica"`

	// 类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// Vpc iD
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 是否开启持久化存储
	EnableStorage *bool `json:"EnableStorage,omitnil,omitempty" name:"EnableStorage"`

	// 数据存储方式
	StorageType *string `json:"StorageType,omitnil,omitempty" name:"StorageType"`

	// 云硬盘容量
	StorageCapacity *int64 `json:"StorageCapacity,omitnil,omitempty" name:"StorageCapacity"`

	// 计费方式
	Paymode *string `json:"Paymode,omitnil,omitempty" name:"Paymode"`

	// EKS集群的ID
	EKSClusterID *string `json:"EKSClusterID,omitnil,omitempty" name:"EKSClusterID"`

	// 集群创建时间
	CreateTime *string `json:"CreateTime,omitnil,omitempty" name:"CreateTime"`

	// 环境配置信息列表
	EnvInfos []*EnvInfo `json:"EnvInfos,omitnil,omitempty" name:"EnvInfos"`

	// 引擎所在的区域
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// 注册引擎是否开启公网
	EnableInternet *bool `json:"EnableInternet,omitnil,omitempty" name:"EnableInternet"`

	// 私有网络列表信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil,omitempty" name:"VpcInfos"`

	// 服务治理相关信息列表
	ServiceGovernanceInfos []*ServiceGovernanceInfo `json:"ServiceGovernanceInfos,omitnil,omitempty" name:"ServiceGovernanceInfos"`

	// 实例的标签信息
	Tags []*KVPair `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 引擎实例是否开启控制台公网访问地址
	EnableConsoleInternet *bool `json:"EnableConsoleInternet,omitnil,omitempty" name:"EnableConsoleInternet"`

	// 引擎实例是否开启控制台内网访问地址
	EnableConsoleIntranet *bool `json:"EnableConsoleIntranet,omitnil,omitempty" name:"EnableConsoleIntranet"`

	// 引擎实例是否展示参数配置页面
	ConfigInfoVisible *bool `json:"ConfigInfoVisible,omitnil,omitempty" name:"ConfigInfoVisible"`

	// 引擎实例控制台默认密码
	ConsoleDefaultPwd *string `json:"ConsoleDefaultPwd,omitnil,omitempty" name:"ConsoleDefaultPwd"`

	// 交易付费类型，0后付费/1预付费
	TradeType *int64 `json:"TradeType,omitnil,omitempty" name:"TradeType"`

	// 自动续费标记：0表示默认状态(用户未设置，即初始状态)， 1表示自动续费，2表示明确不自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitnil,omitempty" name:"AutoRenewFlag"`

	// 预付费到期时间
	CurDeadline *string `json:"CurDeadline,omitnil,omitempty" name:"CurDeadline"`

	// 隔离开始时间
	IsolateTime *string `json:"IsolateTime,omitnil,omitempty" name:"IsolateTime"`

	// 实例地域相关的描述信息
	RegionInfos []*DescribeInstanceRegionInfo `json:"RegionInfos,omitnil,omitempty" name:"RegionInfos"`

	// 所在EKS环境，分为common和yunti
	EKSType *string `json:"EKSType,omitnil,omitempty" name:"EKSType"`

	// 引擎的产品版本
	FeatureVersion *string `json:"FeatureVersion,omitnil,omitempty" name:"FeatureVersion"`

	// 引擎实例是否开启客户端内网访问地址
	EnableClientIntranet *bool `json:"EnableClientIntranet,omitnil,omitempty" name:"EnableClientIntranet"`

	// 存储额外配置选项
	StorageOption []*StorageOption `json:"StorageOption,omitnil,omitempty" name:"StorageOption"`

	// Zookeeper的额外环境数据信息
	ZookeeperRegionInfo *ZookeeperRegionInfo `json:"ZookeeperRegionInfo,omitnil,omitempty" name:"ZookeeperRegionInfo"`

	// 部署架构
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 全局属性
	GlobalType *string `json:"GlobalType,omitnil,omitempty" name:"GlobalType"`

	// 所属组类型
	GroupType *string `json:"GroupType,omitnil,omitempty" name:"GroupType"`

	// 组id
	GroupId []*string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 是否为主地域
	IsMainRegion *bool `json:"IsMainRegion,omitnil,omitempty" name:"IsMainRegion"`
}

type ServiceGovernanceInfo struct {
	// 引擎所在的地域
	EngineRegion *string `json:"EngineRegion,omitnil,omitempty" name:"EngineRegion"`

	// 服务治理引擎绑定的kubernetes集群信息
	BoundK8SInfos []*BoundK8SInfo `json:"BoundK8SInfos,omitnil,omitempty" name:"BoundK8SInfos"`

	// 服务治理引擎绑定的网络信息
	VpcInfos []*VpcInfo `json:"VpcInfos,omitnil,omitempty" name:"VpcInfos"`

	// 当前实例鉴权是否开启
	AuthOpen *bool `json:"AuthOpen,omitnil,omitempty" name:"AuthOpen"`

	// 该实例支持的功能，鉴权就是 Auth
	Features []*string `json:"Features,omitnil,omitempty" name:"Features"`

	// 主账户名默认为 polaris，该值为主账户的默认密码
	MainPassword *string `json:"MainPassword,omitnil,omitempty" name:"MainPassword"`

	// 服务治理pushgateway引擎绑定的网络信息
	PgwVpcInfos []*VpcInfo `json:"PgwVpcInfos,omitnil,omitempty" name:"PgwVpcInfos"`

	// 服务治理限流server引擎绑定的网络信息
	LimiterVpcInfos []*VpcInfo `json:"LimiterVpcInfos,omitnil,omitempty" name:"LimiterVpcInfos"`

	// 引擎关联CLS日志主题信息
	CLSTopics []*PolarisCLSTopicInfo `json:"CLSTopics,omitnil,omitempty" name:"CLSTopics"`

	// 子用户密码
	SubPassword *string `json:"SubPassword,omitnil,omitempty" name:"SubPassword"`
}

type ServiceWafStatus struct {
	//  服务的名字
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 服务的 ID
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 服务的类型
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	//  服务是否开启 WAF 防护
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`
}

type SourceInfo struct {
	// 微服务引擎接入IP地址信息
	Addresses []*string `json:"Addresses,omitnil,omitempty" name:"Addresses"`

	// 微服务引擎VPC信息
	VpcInfo *SourceInstanceVpcInfo `json:"VpcInfo,omitnil,omitempty" name:"VpcInfo"`

	// 微服务引擎鉴权信息
	Auth *SourceInstanceAuth `json:"Auth,omitnil,omitempty" name:"Auth"`
}

type SourceInstanceAuth struct {
	// 用户名
	Username *string `json:"Username,omitnil,omitempty" name:"Username"`

	// 账户密码
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 访问凭据 token
	AccessToken *string `json:"AccessToken,omitnil,omitempty" name:"AccessToken"`
}

type SourceInstanceVpcInfo struct {
	// 微服务引擎VPC信息
	VpcID *string `json:"VpcID,omitnil,omitempty" name:"VpcID"`

	// 微服务引擎子网信息
	SubnetID *string `json:"SubnetID,omitnil,omitempty" name:"SubnetID"`
}

type StorageOption struct {
	// 存储对象，分为snap和txn两种
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 存储类型，分为三类CLOUD_PREMIUM/CLOUD_SSD/CLOUD_SSD_PLUS，分别对应高性能云硬盘、SSD云硬盘、增强型SSD云硬盘
	Type *string `json:"Type,omitnil,omitempty" name:"Type"`

	// 存储容量，[50, 3200]的范围
	Capacity *uint64 `json:"Capacity,omitnil,omitempty" name:"Capacity"`
}

// Predefined struct for user
type UnbindAutoScalerResourceStrategyFromGroupsRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 网关分组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

type UnbindAutoScalerResourceStrategyFromGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 策略ID
	StrategyId *string `json:"StrategyId,omitnil,omitempty" name:"StrategyId"`

	// 网关分组ID列表
	GroupIds []*string `json:"GroupIds,omitnil,omitempty" name:"GroupIds"`
}

func (r *UnbindAutoScalerResourceStrategyFromGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoScalerResourceStrategyFromGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "StrategyId")
	delete(f, "GroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindAutoScalerResourceStrategyFromGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UnbindAutoScalerResourceStrategyFromGroupsResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UnbindAutoScalerResourceStrategyFromGroupsResponse struct {
	*tchttp.BaseResponse
	Response *UnbindAutoScalerResourceStrategyFromGroupsResponseParams `json:"Response"`
}

func (r *UnbindAutoScalerResourceStrategyFromGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindAutoScalerResourceStrategyFromGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateCloudNativeAPIGatewayCertificateInfoRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 绑定的域名列表
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// 证书名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
}

type UpdateCloudNativeAPIGatewayCertificateInfoRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 证书id
	Id *string `json:"Id,omitnil,omitempty" name:"Id"`

	// 绑定的域名列表
	BindDomains []*string `json:"BindDomains,omitnil,omitempty" name:"BindDomains"`

	// 证书名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 云原生网关状态。
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 任务ID
	TaskId *string `json:"TaskId,omitnil,omitempty" name:"TaskId"`
}

// Predefined struct for user
type UpdateCloudNativeAPIGatewaySpecRequestParams struct {
	// 云原生API网关实例ID。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网关分组节点规格配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`
}

type UpdateCloudNativeAPIGatewaySpecRequest struct {
	*tchttp.BaseRequest
	
	// 云原生API网关实例ID。
	// 只支持后付费实例
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关分组id
	GroupId *string `json:"GroupId,omitnil,omitempty" name:"GroupId"`

	// 网关分组节点规格配置。
	NodeConfig *CloudNativeAPIGatewayNodeConfig `json:"NodeConfig,omitnil,omitempty" name:"NodeConfig"`
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
	Result *UpdateCloudNativeAPIGatewayResult `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 引擎类型
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 是否开启客户端公网访问，true开 false关
	EnableClientInternetAccess *bool `json:"EnableClientInternetAccess,omitnil,omitempty" name:"EnableClientInternetAccess"`
}

type UpdateEngineInternetAccessRequest struct {
	*tchttp.BaseRequest
	
	// 引擎ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 引擎类型
	EngineType *string `json:"EngineType,omitnil,omitempty" name:"EngineType"`

	// 是否开启客户端公网访问，true开 false关
	EnableClientInternetAccess *bool `json:"EnableClientInternetAccess,omitnil,omitempty" name:"EnableClientInternetAccess"`
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
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
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

// Predefined struct for user
type UpdateUpstreamHealthCheckConfigRequestParams struct {
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 健康检查配置
	HealthCheckConfig *UpstreamHealthCheckConfig `json:"HealthCheckConfig,omitnil,omitempty" name:"HealthCheckConfig"`
}

type UpdateUpstreamHealthCheckConfigRequest struct {
	*tchttp.BaseRequest
	
	// 网关ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 网关服务名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 健康检查配置
	HealthCheckConfig *UpstreamHealthCheckConfig `json:"HealthCheckConfig,omitnil,omitempty" name:"HealthCheckConfig"`
}

func (r *UpdateUpstreamHealthCheckConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUpstreamHealthCheckConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "HealthCheckConfig")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUpstreamHealthCheckConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUpstreamHealthCheckConfigResponseParams struct {
	// 是否成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUpstreamHealthCheckConfigResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUpstreamHealthCheckConfigResponseParams `json:"Response"`
}

func (r *UpdateUpstreamHealthCheckConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUpstreamHealthCheckConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUpstreamTargetsRequestParams struct {
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称或ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例列表
	Targets []*KongTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

type UpdateUpstreamTargetsRequest struct {
	*tchttp.BaseRequest
	
	// 网关实例ID
	GatewayId *string `json:"GatewayId,omitnil,omitempty" name:"GatewayId"`

	// 服务名称或ID
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 实例列表
	Targets []*KongTarget `json:"Targets,omitnil,omitempty" name:"Targets"`
}

func (r *UpdateUpstreamTargetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUpstreamTargetsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GatewayId")
	delete(f, "Name")
	delete(f, "Targets")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateUpstreamTargetsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateUpstreamTargetsResponseParams struct {
	// 是否更新成功
	Result *bool `json:"Result,omitnil,omitempty" name:"Result"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpdateUpstreamTargetsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateUpstreamTargetsResponseParams `json:"Response"`
}

func (r *UpdateUpstreamTargetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateUpstreamTargetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpstreamHealthCheckConfig struct {
	// 开启主动健康检查
	EnableActiveHealthCheck *bool `json:"EnableActiveHealthCheck,omitnil,omitempty" name:"EnableActiveHealthCheck"`

	// 主动健康检查配置
	ActiveHealthCheck *KongActiveHealthCheck `json:"ActiveHealthCheck,omitnil,omitempty" name:"ActiveHealthCheck"`

	// 开启被动健康检查
	EnablePassiveHealthCheck *bool `json:"EnablePassiveHealthCheck,omitnil,omitempty" name:"EnablePassiveHealthCheck"`

	// 被动健康检查配置
	PassiveHealthCheck *KongPassiveHealthCheck `json:"PassiveHealthCheck,omitnil,omitempty" name:"PassiveHealthCheck"`

	// 连续健康阈值，单位：次
	Successes *uint64 `json:"Successes,omitnil,omitempty" name:"Successes"`

	// 连续异常阈值，单位：次	
	Failures *uint64 `json:"Failures,omitnil,omitempty" name:"Failures"`

	// 超时阈值，单位：次
	Timeouts *uint64 `json:"Timeouts,omitnil,omitempty" name:"Timeouts"`

	// 健康HTTP状态码
	HealthyHttpStatuses []*uint64 `json:"HealthyHttpStatuses,omitnil,omitempty" name:"HealthyHttpStatuses"`

	// 异常HTTP状态码
	UnhealthyHttpStatuses []*uint64 `json:"UnhealthyHttpStatuses,omitnil,omitempty" name:"UnhealthyHttpStatuses"`

	// 健康检查监控上报的数据屏蔽权重为0的节点
	// 注意：此字段可能返回 null，表示取不到有效值。
	//
	// Deprecated: IgnoreZeroWeightNodes is deprecated.
	IgnoreZeroWeightNodes *bool `json:"IgnoreZeroWeightNodes,omitnil,omitempty" name:"IgnoreZeroWeightNodes"`

	// 健康检查支持权重为0节点
	ZeroWeightHeathCheck *bool `json:"ZeroWeightHeathCheck,omitnil,omitempty" name:"ZeroWeightHeathCheck"`
}

type VpcInfo struct {
	// Vpc Id
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 内网访问地址
	IntranetAddress *string `json:"IntranetAddress,omitnil,omitempty" name:"IntranetAddress"`

	// 负载均衡均衡接入点子网ID
	LbSubnetId *string `json:"LbSubnetId,omitnil,omitempty" name:"LbSubnetId"`
}

type ZookeeperRegionInfo struct {
	// 部署架构信息
	// 
	// - SingleRegion: 普通单地域
	// - MultiRegion: 普通多地域场景
	// - MasterSlave: 两地域，主备地域场景
	DeployMode *string `json:"DeployMode,omitnil,omitempty" name:"DeployMode"`

	// 主地域的额外信息
	MainRegion *ZookeeperRegionMyIdInfo `json:"MainRegion,omitnil,omitempty" name:"MainRegion"`

	// 其他地域的额外信息
	OtherRegions []*ZookeeperRegionMyIdInfo `json:"OtherRegions,omitnil,omitempty" name:"OtherRegions"`
}

type ZookeeperRegionMyIdInfo struct {
	// 地域信息
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`

	// myid 的起始号段
	MyIdStart *int64 `json:"MyIdStart,omitnil,omitempty" name:"MyIdStart"`

	// myid 的结束号段
	MyIdEnd *int64 `json:"MyIdEnd,omitnil,omitempty" name:"MyIdEnd"`
}

type ZookeeperReplica struct {
	// 名称
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 角色
	Role *string `json:"Role,omitnil,omitempty" name:"Role"`

	// 状态
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 子网ID
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 可用区ID
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 可用区ID
	ZoneId *string `json:"ZoneId,omitnil,omitempty" name:"ZoneId"`

	// 别名
	AliasName *string `json:"AliasName,omitnil,omitempty" name:"AliasName"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`
}

type ZookeeperServerInterface struct {
	// 接口名
	Interface *string `json:"Interface,omitnil,omitempty" name:"Interface"`
}