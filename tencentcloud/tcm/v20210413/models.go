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

package v20210413

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AccessLogConfig struct {

	// 是否启用
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 采用的模板，可选值：istio（默认）、trace
	Template *string `json:"Template,omitempty" name:"Template"`

	// 选中的范围
	SelectedRange *SelectedRange `json:"SelectedRange,omitempty" name:"SelectedRange"`

	// 腾讯云日志服务相关参数
	CLS *CLS `json:"CLS,omitempty" name:"CLS"`

	// 编码格式，可选值：TEXT、JSON
	Encoding *string `json:"Encoding,omitempty" name:"Encoding"`

	// 日志格式
	Format *string `json:"Format,omitempty" name:"Format"`
}

type ActiveOperation struct {

	// 操作Id
	OperationId *string `json:"OperationId,omitempty" name:"OperationId"`

	// 操作类型，取值范围：
	// - LINK_CLUSTERS: 关联集群
	// - RELINK_CLUSTERS: 重新关联集群
	// - UNLINK_CLUSTERS: 解关联集群
	// - INSTALL_MESH: 安装网格
	Type *string `json:"Type,omitempty" name:"Type"`
}

type AutoInjectionNamespaceState struct {

	// 命名空间名称
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 注入状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`
}

type CLS struct {

	// 是否启用
	Enable *bool `json:"Enable,omitempty" name:"Enable"`

	// 日志集
	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`

	// 日志主题
	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

type Cluster struct {

	// 集群Id
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 集群角色，取值范围：
	// - MASTER：控制面所在的主集群
	// - REMOTE：主集群管理的远端集群
	Role *string `json:"Role,omitempty" name:"Role"`

	// 私有网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 名称，只读
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 状态，只读
	State *string `json:"State,omitempty" name:"State"`

	// 关联时间，只读
	LinkedTime *string `json:"LinkedTime,omitempty" name:"LinkedTime"`

	// 集群配置
	Config *ClusterConfig `json:"Config,omitempty" name:"Config"`

	// 详细状态，只读
	Status *ClusterStatus `json:"Status,omitempty" name:"Status"`

	// 类型，取值范围：
	// - TKE
	// - EKS
	Type *string `json:"Type,omitempty" name:"Type"`
}

type ClusterConfig struct {

	// 自动注入SideCar的NameSpace
	AutoInjectionNamespaceList []*string `json:"AutoInjectionNamespaceList,omitempty" name:"AutoInjectionNamespaceList"`

	// Ingress配置列表
	IngressGatewayList []*IngressGateway `json:"IngressGatewayList,omitempty" name:"IngressGatewayList"`

	// Egress配置列表
	EgressGatewayList []*EgressGateway `json:"EgressGatewayList,omitempty" name:"EgressGatewayList"`

	// Istiod配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Istiod *IstiodConfig `json:"Istiod,omitempty" name:"Istiod"`

	// 部署配置
	DeployConfig *DeployConfig `json:"DeployConfig,omitempty" name:"DeployConfig"`

	// 自动注入命名空间状态列表
	AutoInjectionNamespaceStateList []*AutoInjectionNamespaceState `json:"AutoInjectionNamespaceStateList,omitempty" name:"AutoInjectionNamespaceStateList"`
}

type ClusterStatus struct {

	// 关联状态，取值范围：
	// - LINKING: 关联中
	// - LINKED: 已关联
	// - UNLINKING: 解关联中
	// - LINK_FAILED: 关联失败
	// - UNLINK_FAILED: 解关联失败
	LinkState *string `json:"LinkState,omitempty" name:"LinkState"`

	// 关联错误详情
	// 注意：此字段可能返回 null，表示取不到有效值。
	LinkErrorDetail *string `json:"LinkErrorDetail,omitempty" name:"LinkErrorDetail"`
}

type DeployConfig struct {

	// 部署类型，取值范围：
	// - SPECIFIC：专有模式
	// - AUTO：普通模式
	NodeSelectType *string `json:"NodeSelectType,omitempty" name:"NodeSelectType"`

	// 指定的节点
	Nodes []*string `json:"Nodes,omitempty" name:"Nodes"`
}

type DescribeMeshListRequest struct {
	*tchttp.BaseRequest

	// 过滤条件
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 分页限制
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 分页偏移
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMeshListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMeshListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMeshListResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 查询到的网格信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		MeshList []*Mesh `json:"MeshList,omitempty" name:"MeshList"`

		// 总数
		Total *int64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMeshListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMeshRequest struct {
	*tchttp.BaseRequest

	// 需要查询的网格 Id
	MeshId *string `json:"MeshId,omitempty" name:"MeshId"`
}

func (r *DescribeMeshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MeshId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeMeshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMeshResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Mesh详细信息
		Mesh *Mesh `json:"Mesh,omitempty" name:"Mesh"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMeshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeMeshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EgressGateway struct {

	// Egress名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 所在的Namespace
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 工作负载配置
	Workload *WorkloadConfig `json:"Workload,omitempty" name:"Workload"`
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type GrafanaInfo struct {

	// 是否开启
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`

	// 内网地址
	InternalURL *string `json:"InternalURL,omitempty" name:"InternalURL"`

	// 公网地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicURL *string `json:"PublicURL,omitempty" name:"PublicURL"`

	// 公网失败原因
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicFailedReason *string `json:"PublicFailedReason,omitempty" name:"PublicFailedReason"`

	// 公网失败详细信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicFailedMessage *string `json:"PublicFailedMessage,omitempty" name:"PublicFailedMessage"`
}

type HorizontalPodAutoscalerSpec struct {

	// 最小副本数
	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`

	// 最大副本数
	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`

	// 用于计算副本数的指标
	Metrics []*MetricSpec `json:"Metrics,omitempty" name:"Metrics"`
}

type IngressGateway struct {

	// IngressGateway 实例名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 集群 ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Service 配置
	Service *Service `json:"Service,omitempty" name:"Service"`

	// Workload 配置
	Workload *WorkloadConfig `json:"Workload,omitempty" name:"Workload"`

	// 负载均衡配置，自动创建 CLB 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	LoadBalancer *LoadBalancer `json:"LoadBalancer,omitempty" name:"LoadBalancer"`

	// IngressGateway 状态信息，只读
	Status *IngressGatewayStatus `json:"Status,omitempty" name:"Status"`

	// 负载均衡实例ID，使用已有 CLB 时返回
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`
}

type IngressGatewayStatus struct {

	// 负载均衡实例状态
	LoadBalancer *LoadBalancerStatus `json:"LoadBalancer,omitempty" name:"LoadBalancer"`
}

type InjectConfig struct {

	// 不需要进行代理的 ip 地址范围
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExcludeIPRanges []*string `json:"ExcludeIPRanges,omitempty" name:"ExcludeIPRanges"`

	// 是否等待sidecar启动
	// 注意：此字段可能返回 null，表示取不到有效值。
	HoldApplicationUntilProxyStarts *bool `json:"HoldApplicationUntilProxyStarts,omitempty" name:"HoldApplicationUntilProxyStarts"`
}

type IstioConfig struct {

	// 外部流量策略
	OutboundTrafficPolicy *string `json:"OutboundTrafficPolicy,omitempty" name:"OutboundTrafficPolicy"`

	// 调用链配置
	Tracing *TracingConfig `json:"Tracing,omitempty" name:"Tracing"`

	// 禁用策略检查功能
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisablePolicyChecks *bool `json:"DisablePolicyChecks,omitempty" name:"DisablePolicyChecks"`
}

type IstiodConfig struct {

	// 工作负载配置
	Workload *WorkloadConfig `json:"Workload,omitempty" name:"Workload"`
}

type LoadBalancer struct {

	// 负载均衡实例的网络类型：
	// OPEN：公网属性， INTERNAL：内网属性。
	// 只读。
	LoadBalancerType *string `json:"LoadBalancerType,omitempty" name:"LoadBalancerType"`

	// 负载均衡实例所在的子网（仅对内网VPC型LB有意义），只读。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// TRAFFIC_POSTPAID_BY_HOUR 按流量按小时后计费 ; BANDWIDTH_POSTPAID_BY_HOUR 按带宽按小时后计费;只读。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 最大出带宽，单位Mbps，范围支持0到2048，仅对公网属性的LB生效，默认值 10
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type LoadBalancerStatus struct {

	// 负载均衡实例 ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 负载均衡实例名字
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" name:"LoadBalancerName"`

	// 负载均衡实例 VIP
	LoadBalancerVip *string `json:"LoadBalancerVip,omitempty" name:"LoadBalancerVip"`
}

type Mesh struct {

	// Mesh实例Id
	MeshId *string `json:"MeshId,omitempty" name:"MeshId"`

	// Mesh名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// Mesh类型，取值范围：
	// - STANDALONE：独立网格
	// - HOSTED：托管网格
	Type *string `json:"Type,omitempty" name:"Type"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 版本
	Version *string `json:"Version,omitempty" name:"Version"`

	// Mesh状态，取值范围：
	// - PENDING：等待中
	// - CREATING：创建中
	// - RUNNING：运行中
	// - ABNORMAL：异常
	// - UPGRADING：升级中
	// - CANARY_UPGRADED：升级灰度完成
	// - ROLLBACKING：升级回滚
	// - DELETING：删除中
	// - CREATE_FAILED：安装失败
	// - DELETE_FAILED：删除失败
	// - UPGRADE_FAILED：升级失败
	// - ROLLBACK_FAILED：回滚失败
	State *string `json:"State,omitempty" name:"State"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 修改时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// 集群列表
	ClusterList []*Cluster `json:"ClusterList,omitempty" name:"ClusterList"`

	// Mesh配置
	Config *MeshConfig `json:"Config,omitempty" name:"Config"`

	// Mesh详细状态
	Status *MeshStatus `json:"Status,omitempty" name:"Status"`
}

type MeshConfig struct {

	// Istio配置
	Istio *IstioConfig `json:"Istio,omitempty" name:"Istio"`

	// AccessLog配置
	AccessLog *AccessLogConfig `json:"AccessLog,omitempty" name:"AccessLog"`

	// Prometheus配置
	Prometheus *PrometheusConfig `json:"Prometheus,omitempty" name:"Prometheus"`

	// 自动注入配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	Inject *InjectConfig `json:"Inject,omitempty" name:"Inject"`
}

type MeshStatus struct {

	// 服务数量
	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`

	// 灰度升级的版本
	// 注意：此字段可能返回 null，表示取不到有效值。
	CanaryVersion *string `json:"CanaryVersion,omitempty" name:"CanaryVersion"`

	// 已废弃
	// 注意：此字段可能返回 null，表示取不到有效值。
	Prometheus []*PrometheusStatus `json:"Prometheus,omitempty" name:"Prometheus"`

	// 状态附带信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StateMessage *string `json:"StateMessage,omitempty" name:"StateMessage"`

	// 正在执行的异步操作
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveOperationList []*ActiveOperation `json:"ActiveOperationList,omitempty" name:"ActiveOperationList"`

	// 获取TPS信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	TPS *PrometheusStatus `json:"TPS,omitempty" name:"TPS"`
}

type MetricSpec struct {

	// 指标来源类型，支持 Pods/Resource
	Type *string `json:"Type,omitempty" name:"Type"`

	// 使用自定义指标扩进行自动扩缩容
	Pods *PodsMetricSource `json:"Pods,omitempty" name:"Pods"`

	// 使用资源指标扩进行自动扩缩容
	Resource *ResourceMetricSource `json:"Resource,omitempty" name:"Resource"`
}

type PodsMetricSource struct {

	// 指标名
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 目标值
	TargetAverageValue *string `json:"TargetAverageValue,omitempty" name:"TargetAverageValue"`
}

type PrometheusConfig struct {

	// 虚拟网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网Id
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`

	// 地域
	Region *string `json:"Region,omitempty" name:"Region"`

	// 关联已存在实例Id，不填则默认创建
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type PrometheusStatus struct {

	// Prometheus Id
	PrometheusId *string `json:"PrometheusId,omitempty" name:"PrometheusId"`

	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 虚拟网络Id
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 状态
	State *string `json:"State,omitempty" name:"State"`

	// 地区
	Region *string `json:"Region,omitempty" name:"Region"`

	// Grafana信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Grafana *GrafanaInfo `json:"Grafana,omitempty" name:"Grafana"`
}

type Resource struct {

	// 资源类型 cpu/memory
	Name *string `json:"Name,omitempty" name:"Name"`

	// 资源数量
	Quantity *string `json:"Quantity,omitempty" name:"Quantity"`
}

type ResourceMetricSource struct {

	// 资源名称 cpu/memory
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目标平均利用率
	TargetAverageUtilization *int64 `json:"TargetAverageUtilization,omitempty" name:"TargetAverageUtilization"`

	// 目标平均值
	TargetAverageValue *string `json:"TargetAverageValue,omitempty" name:"TargetAverageValue"`
}

type ResourceRequirements struct {

	// Limits 描述了允许的最大计算资源量。
	Limits []*Resource `json:"Limits,omitempty" name:"Limits"`

	// Requests 描述所需的最小计算资源量。
	Requests []*Resource `json:"Requests,omitempty" name:"Requests"`
}

type SelectedItems struct {

	// 命名空间
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 选中项目名字
	ItemName *string `json:"ItemName,omitempty" name:"ItemName"`

	// ingress gw的名称列表
	Gateways []*string `json:"Gateways,omitempty" name:"Gateways"`
}

type SelectedRange struct {

	// 选中的项目详细内容
	Items []*SelectedItems `json:"Items,omitempty" name:"Items"`

	// 是否全选
	All *bool `json:"All,omitempty" name:"All"`
}

type Service struct {

	// ClusterIP/NodePort/LoadBalancer
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否开启LB直通Pod
	CLBDirectAccess *bool `json:"CLBDirectAccess,omitempty" name:"CLBDirectAccess"`

	// 服务是否希望将外部流量路由到节点本地或集群范围的端点。 有两个可用选项：Cluster（默认）和 Local。Cluster 隐藏了客户端源 IP，可能导致第二跳到另一个节点；Local 保留客户端源 IP 并避免 LoadBalancer 和 NodePort 类型服务的第二跳。
	ExternalTrafficPolicy *string `json:"ExternalTrafficPolicy,omitempty" name:"ExternalTrafficPolicy"`
}

type TracingConfig struct {

	// 调用链采样率，百分比
	Sampling *float64 `json:"Sampling,omitempty" name:"Sampling"`
}

type WorkloadConfig struct {

	// 工作副本数
	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`

	// 资源限制
	Resources *ResourceRequirements `json:"Resources,omitempty" name:"Resources"`

	// HPA策略
	HorizontalPodAutoscaler *HorizontalPodAutoscalerSpec `json:"HorizontalPodAutoscaler,omitempty" name:"HorizontalPodAutoscaler"`

	// 部署到指定节点
	SelectedNodeList []*string `json:"SelectedNodeList,omitempty" name:"SelectedNodeList"`
}
