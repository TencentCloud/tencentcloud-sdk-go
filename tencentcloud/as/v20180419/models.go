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

package v20180419

import (
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
)

type Activity struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 伸缩活动ID。
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 伸缩活动类型。取值如下：<br>
	// <li>SCALE_OUT：扩容活动<li>SCALE_IN：缩容活动<li>ATTACH_INSTANCES：添加实例<li>REMOVE_INSTANCES：销毁实例<li>DETACH_INSTANCES：移出实例<li>TERMINATE_INSTANCES_UNEXPECTEDLY：实例在CVM控制台被销毁<li>REPLACE_UNHEALTHY_INSTANCE：替换不健康实例
	// <li>START_INSTANCES：开启实例
	// <li>STOP_INSTANCES：关闭实例
	// <li>INVOKE_COMMAND：执行命令
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// 伸缩活动状态。取值如下：<br>
	// <li>INIT：初始化中
	// <li>RUNNING：运行中
	// <li>SUCCESSFUL：活动成功
	// <li>PARTIALLY_SUCCESSFUL：活动部分成功
	// <li>FAILED：活动失败
	// <li>CANCELLED：活动取消
	StatusCode *string `json:"StatusCode,omitnil,omitempty" name:"StatusCode"`

	// 伸缩活动状态描述。
	StatusMessage *string `json:"StatusMessage,omitnil,omitempty" name:"StatusMessage"`

	// 伸缩活动起因。
	Cause *string `json:"Cause,omitnil,omitempty" name:"Cause"`

	// 伸缩活动描述。
	Description *string `json:"Description,omitnil,omitempty" name:"Description"`

	// 伸缩活动开始时间。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 伸缩活动结束时间。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 伸缩活动创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 该参数已废弃，请勿使用。
	//
	// Deprecated: ActivityRelatedInstanceSet is deprecated.
	ActivityRelatedInstanceSet []*ActivtyRelatedInstance `json:"ActivityRelatedInstanceSet,omitnil,omitempty" name:"ActivityRelatedInstanceSet"`

	// 伸缩活动状态简要描述。
	StatusMessageSimplified *string `json:"StatusMessageSimplified,omitnil,omitempty" name:"StatusMessageSimplified"`

	// 伸缩活动中生命周期挂钩的执行结果。
	LifecycleActionResultSet []*LifecycleActionResultInfo `json:"LifecycleActionResultSet,omitnil,omitempty" name:"LifecycleActionResultSet"`

	// 伸缩活动状态详细描述。
	DetailedStatusMessageSet []*DetailedStatusMessage `json:"DetailedStatusMessageSet,omitnil,omitempty" name:"DetailedStatusMessageSet"`

	// 执行命令结果。
	InvocationResultSet []*InvocationResult `json:"InvocationResultSet,omitnil,omitempty" name:"InvocationResultSet"`

	// 伸缩活动相关实例信息集合。
	RelatedInstanceSet []*RelatedInstance `json:"RelatedInstanceSet,omitnil,omitempty" name:"RelatedInstanceSet"`
}

type ActivtyRelatedInstance struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例在伸缩活动中的状态。取值如下：
	// <li>INIT：初始化中
	// <li>RUNNING：实例操作中
	// <li>SUCCESSFUL：活动成功
	// <li>FAILED：活动失败
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

type Advice struct {
	// 问题描述。
	Problem *string `json:"Problem,omitnil,omitempty" name:"Problem"`

	// 问题详情。
	Detail *string `json:"Detail,omitnil,omitempty" name:"Detail"`

	// 建议解决方案。
	Solution *string `json:"Solution,omitnil,omitempty" name:"Solution"`

	// 伸缩建议警告级别。取值范围：<br>
	// <li>WARNING：警告级别<br>
	// <li>CRITICAL：严重级别<br>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`
}

// Predefined struct for user
type AttachInstancesRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type AttachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *AttachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachInstancesResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *AttachInstancesResponseParams `json:"Response"`
}

func (r *AttachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachLoadBalancersRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 传统型负载均衡器ID列表，每个伸缩组绑定传统型负载均衡器数量上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 应用型负载均衡器列表，每个伸缩组绑定应用型负载均衡器数量上限为100，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

type AttachLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 传统型负载均衡器ID列表，每个伸缩组绑定传统型负载均衡器数量上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 应用型负载均衡器列表，每个伸缩组绑定应用型负载均衡器数量上限为100，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

func (r *AttachLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LoadBalancerIds")
	delete(f, "ForwardLoadBalancers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AttachLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AttachLoadBalancersResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type AttachLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *AttachLoadBalancersResponseParams `json:"Response"`
}

func (r *AttachLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AttachLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoScalingAdvice struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组警告级别。取值范围：<br>
	// <li>NORMAL：正常<br>
	// <li>WARNING：警告级别<br>
	// <li>CRITICAL：严重级别<br>
	Level *string `json:"Level,omitnil,omitempty" name:"Level"`

	// 伸缩组配置建议集合。
	Advices []*Advice `json:"Advices,omitnil,omitempty" name:"Advices"`
}

type AutoScalingGroup struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组名称
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 伸缩组当前状态。取值范围：<br>
	// <li>NORMAL：正常<br>
	// <li>CVM_ABNORMAL：启动配置异常<br>
	// <li>LB_ABNORMAL：负载均衡器异常<br>
	// <li>LB_LISTENER_ABNORMAL：负载均衡器监听器异常<br>
	// <li>LB_LOCATION_ABNORMAL：负载均衡器监听器转发配置异常<br>
	// <li>VPC_ABNORMAL：VPC网络异常<br>
	// <li>SUBNET_ABNORMAL：VPC子网异常<br>
	// <li>INSUFFICIENT_BALANCE：余额不足<br>
	// <li>LB_BACKEND_REGION_NOT_MATCH：CLB实例后端地域与AS服务所在地域不匹配<br>
	// <li>LB_BACKEND_VPC_NOT_MATCH：CLB实例VPC与伸缩组VPC不匹配
	AutoScalingGroupStatus *string `json:"AutoScalingGroupStatus,omitnil,omitempty" name:"AutoScalingGroupStatus"`

	// 创建时间，采用UTC标准计时
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 默认冷却时间，单位秒
	DefaultCooldown *int64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// 期望实例数
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 启用状态，取值包括`ENABLED`和`DISABLED`
	EnabledStatus *string `json:"EnabledStatus,omitnil,omitempty" name:"EnabledStatus"`

	// 应用型负载均衡器列表
	ForwardLoadBalancerSet []*ForwardLoadBalancer `json:"ForwardLoadBalancerSet,omitnil,omitempty" name:"ForwardLoadBalancerSet"`

	// 实例数量
	InstanceCount *int64 `json:"InstanceCount,omitnil,omitempty" name:"InstanceCount"`

	// 状态为`IN_SERVICE`实例的数量
	InServiceInstanceCount *int64 `json:"InServiceInstanceCount,omitnil,omitempty" name:"InServiceInstanceCount"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 启动配置名称
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 传统型负载均衡器ID列表
	LoadBalancerIdSet []*string `json:"LoadBalancerIdSet,omitnil,omitempty" name:"LoadBalancerIdSet"`

	// 最大实例数
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 最小实例数
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 子网ID列表
	SubnetIdSet []*string `json:"SubnetIdSet,omitnil,omitempty" name:"SubnetIdSet"`

	// 销毁策略
	TerminationPolicySet []*string `json:"TerminationPolicySet,omitnil,omitempty" name:"TerminationPolicySet"`

	// VPC标识
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 可用区列表
	ZoneSet []*string `json:"ZoneSet,omitnil,omitempty" name:"ZoneSet"`

	// 重试策略
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// 伸缩组是否处于伸缩活动中，`IN_ACTIVITY`表示处于伸缩活动中，`NOT_IN_ACTIVITY`表示不处于伸缩活动中。
	InActivityStatus *string `json:"InActivityStatus,omitnil,omitempty" name:"InActivityStatus"`

	// 伸缩组标签列表
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务设置
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// 实例具有IPv6地址数量的配置
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// 多可用区/子网策略。
	// <br><li> PRIORITY，按照可用区/子网列表的顺序，作为优先级来尝试创建实例，如果优先级最高的可用区/子网可以创建成功，则总在该可用区/子网创建。
	// <br><li> EQUALITY：每次选择当前实例数最少的可用区/子网进行扩容，使得每个可用区/子网都有机会发生扩容，多次扩容出的实例会打散到多个可用区/子网。
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// 伸缩组实例健康检查类型，取值如下：<br><li>CVM：根据实例网络状态判断实例是否处于不健康状态，不健康的网络状态即发生实例 PING 不可达事件，详细判断标准可参考[实例健康检查](https://cloud.tencent.com/document/product/377/8553)<br><li>CLB：根据 CLB 的健康检查状态判断实例是否处于不健康状态，CLB健康检查原理可参考[健康检查](https://cloud.tencent.com/document/product/214/6097)
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// CLB健康检查宽限期
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// 实例分配策略，取值包括 LAUNCH_CONFIGURATION 和 SPOT_MIXED。
	// <br><li> LAUNCH_CONFIGURATION，代表传统的按照启动配置模式。
	// <br><li> SPOT_MIXED，代表竞价混合模式。目前仅支持启动配置为按量计费模式时使用混合模式，混合模式下，伸缩组将根据设定扩容按量或竞价机型。使用混合模式时，关联的启动配置的计费类型不可被修改。
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// 竞价混合模式下，各计费类型实例的分配策略。
	// 仅当 InstanceAllocationPolicy 取 SPOT_MIXED 时才会返回有效值。
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// 容量重平衡功能，仅对伸缩组内的竞价实例有效。取值范围：
	// <br><li> TRUE，开启该功能，当伸缩组内的竞价实例即将被竞价实例服务自动回收前，AS 主动发起竞价实例销毁流程，如果有配置过缩容 hook，则销毁前 hook 会生效。销毁流程启动后，AS 会异步开启一个扩容活动，用于补齐期望实例数。
	// <br><li> FALSE，不开启该功能，则 AS 等待竞价实例被销毁后才会去扩容补齐伸缩组期望实例数。
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// 实例名称序号相关设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
}

type AutoScalingGroupAbstract struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组名称。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`
}

type AutoScalingNotification struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 用户组ID列表。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// 通知事件列表。
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// 事件通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// 通知接收端类型。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CMQ 队列名。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

// Predefined struct for user
type CancelInstanceRefreshRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新活动ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

type CancelInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新活动ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

func (r *CancelInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CancelInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CancelInstanceRefreshResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CancelInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *CancelInstanceRefreshResponseParams `json:"Response"`
}

func (r *CancelInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CancelInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearLaunchConfigurationAttributesRequestParams struct {
	// 启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 是否清空数据盘信息，非必填，默认为 false。
	// 填 true 代表清空“数据盘”信息，清空后基于此新创建的云主机将不含有任何数据盘。
	ClearDataDisks *bool `json:"ClearDataDisks,omitnil,omitempty" name:"ClearDataDisks"`

	// 是否清空云服务器主机名相关设置信息，非必填，默认为 false。
	// 填 true 代表清空主机名设置信息，清空后基于此新创建的云主机将不设置主机名。
	ClearHostNameSettings *bool `json:"ClearHostNameSettings,omitnil,omitempty" name:"ClearHostNameSettings"`

	// 是否清空云服务器实例名相关设置信息，非必填，默认为 false。
	// 填 true 代表清空主机名设置信息，清空后基于此新创建的云主机将按照“as-{{ 伸缩组AutoScalingGroupName }}”进行设置。
	ClearInstanceNameSettings *bool `json:"ClearInstanceNameSettings,omitnil,omitempty" name:"ClearInstanceNameSettings"`

	// 是否清空置放群组信息，非必填，默认为 false。
	// 填 true 代表清空置放群组信息，清空后基于此新创建的云主机将不指定任何置放群组。
	ClearDisasterRecoverGroupIds *bool `json:"ClearDisasterRecoverGroupIds,omitnil,omitempty" name:"ClearDisasterRecoverGroupIds"`

	// 是否清空实例标签列表，非必填，默认为 false。
	// 填 true 代表清空实例标签列表，清空后基于此新创建的云主机将不会绑定列表中的标签。
	ClearInstanceTags *bool `json:"ClearInstanceTags,omitnil,omitempty" name:"ClearInstanceTags"`
}

type ClearLaunchConfigurationAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 是否清空数据盘信息，非必填，默认为 false。
	// 填 true 代表清空“数据盘”信息，清空后基于此新创建的云主机将不含有任何数据盘。
	ClearDataDisks *bool `json:"ClearDataDisks,omitnil,omitempty" name:"ClearDataDisks"`

	// 是否清空云服务器主机名相关设置信息，非必填，默认为 false。
	// 填 true 代表清空主机名设置信息，清空后基于此新创建的云主机将不设置主机名。
	ClearHostNameSettings *bool `json:"ClearHostNameSettings,omitnil,omitempty" name:"ClearHostNameSettings"`

	// 是否清空云服务器实例名相关设置信息，非必填，默认为 false。
	// 填 true 代表清空主机名设置信息，清空后基于此新创建的云主机将按照“as-{{ 伸缩组AutoScalingGroupName }}”进行设置。
	ClearInstanceNameSettings *bool `json:"ClearInstanceNameSettings,omitnil,omitempty" name:"ClearInstanceNameSettings"`

	// 是否清空置放群组信息，非必填，默认为 false。
	// 填 true 代表清空置放群组信息，清空后基于此新创建的云主机将不指定任何置放群组。
	ClearDisasterRecoverGroupIds *bool `json:"ClearDisasterRecoverGroupIds,omitnil,omitempty" name:"ClearDisasterRecoverGroupIds"`

	// 是否清空实例标签列表，非必填，默认为 false。
	// 填 true 代表清空实例标签列表，清空后基于此新创建的云主机将不会绑定列表中的标签。
	ClearInstanceTags *bool `json:"ClearInstanceTags,omitnil,omitempty" name:"ClearInstanceTags"`
}

func (r *ClearLaunchConfigurationAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearLaunchConfigurationAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	delete(f, "ClearDataDisks")
	delete(f, "ClearHostNameSettings")
	delete(f, "ClearInstanceNameSettings")
	delete(f, "ClearDisasterRecoverGroupIds")
	delete(f, "ClearInstanceTags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearLaunchConfigurationAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ClearLaunchConfigurationAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ClearLaunchConfigurationAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ClearLaunchConfigurationAttributesResponseParams `json:"Response"`
}

func (r *ClearLaunchConfigurationAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearLaunchConfigurationAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteLifecycleActionRequestParams struct {
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 生命周期动作的结果，取值范围为“CONTINUE”或“ABANDON”
	LifecycleActionResult *string `json:"LifecycleActionResult,omitnil,omitempty" name:"LifecycleActionResult"`

	// 实例ID，“InstanceId”和“LifecycleActionToken”必须填充其中一个
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// “InstanceId”和“LifecycleActionToken”必须填充其中一个
	LifecycleActionToken *string `json:"LifecycleActionToken,omitnil,omitempty" name:"LifecycleActionToken"`
}

type CompleteLifecycleActionRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 生命周期动作的结果，取值范围为“CONTINUE”或“ABANDON”
	LifecycleActionResult *string `json:"LifecycleActionResult,omitnil,omitempty" name:"LifecycleActionResult"`

	// 实例ID，“InstanceId”和“LifecycleActionToken”必须填充其中一个
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// “InstanceId”和“LifecycleActionToken”必须填充其中一个
	LifecycleActionToken *string `json:"LifecycleActionToken,omitnil,omitempty" name:"LifecycleActionToken"`
}

func (r *CompleteLifecycleActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteLifecycleActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	delete(f, "LifecycleActionResult")
	delete(f, "InstanceId")
	delete(f, "LifecycleActionToken")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CompleteLifecycleActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CompleteLifecycleActionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CompleteLifecycleActionResponse struct {
	*tchttp.BaseResponse
	Response *CompleteLifecycleActionResponseParams `json:"Response"`
}

func (r *CompleteLifecycleActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CompleteLifecycleActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupFromInstanceRequestParams struct {
	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 最小实例数，取值范围为0-2000。
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 期望实例数，大小介于最小实例数和最大实例数之间。
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 是否继承实例标签，默认值为False
	InheritInstanceTag *bool `json:"InheritInstanceTag,omitnil,omitempty" name:"InheritInstanceTag"`
}

type CreateAutoScalingGroupFromInstanceRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 最小实例数，取值范围为0-2000。
	MinSize *int64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *int64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 期望实例数，大小介于最小实例数和最大实例数之间。
	DesiredCapacity *int64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 是否继承实例标签，默认值为False
	InheritInstanceTag *bool `json:"InheritInstanceTag,omitnil,omitempty" name:"InheritInstanceTag"`
}

func (r *CreateAutoScalingGroupFromInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupFromInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupName")
	delete(f, "InstanceId")
	delete(f, "MinSize")
	delete(f, "MaxSize")
	delete(f, "DesiredCapacity")
	delete(f, "InheritInstanceTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoScalingGroupFromInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupFromInstanceResponseParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoScalingGroupFromInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoScalingGroupFromInstanceResponseParams `json:"Response"`
}

func (r *CreateAutoScalingGroupFromInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupFromInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupRequestParams struct {
	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// VPC ID，基础网络则填空字符串
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 默认冷却时间，单位秒，默认值为300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// 期望实例数，大小介于最小实例数和最大实例数之间
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 传统负载均衡器ID列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 伸缩组内实例所属项目ID。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 应用型负载均衡器列表，目前长度上限为100，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// 子网ID列表，VPC场景下必须指定子网。多个子网以填写顺序为优先级，依次进行尝试，直至可以成功创建实例。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 销毁策略，目前长度上限为1。取值包括 OLDEST_INSTANCE 和 NEWEST_INSTANCE，默认取值为 OLDEST_INSTANCE。
	// <br><li> OLDEST_INSTANCE 优先销毁伸缩组中最旧的实例。
	// <br><li> NEWEST_INSTANCE，优先销毁伸缩组中最新的实例。
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// 可用区列表，基础网络场景下必须指定可用区。多个可用区以填写顺序为优先级，依次进行尝试，直至可以成功创建实例。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 重试策略，取值包括 IMMEDIATE_RETRY、 INCREMENTAL_INTERVALS、NO_RETRY，默认取值为 IMMEDIATE_RETRY。部分成功的伸缩活动判定为一次失败活动。
	// <br><li> IMMEDIATE_RETRY，立即重试，在较短时间内快速重试，连续失败超过一定次数（5次）后不再重试。
	// <br><li> INCREMENTAL_INTERVALS，间隔递增重试，随着连续失败次数的增加，重试间隔逐渐增大，重试间隔从秒级到1天不等。
	// <br><li> NO_RETRY，不进行重试，直到再次收到用户调用或者告警信息后才会重试。
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// 可用区校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有可用区（Zone）或子网（SubnetId）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个可用区（Zone）或子网（SubnetId）可用则通过校验，否则校验报错。
	// 
	// 可用区或子网不可用的常见原因包括该可用区CVM实例类型售罄、该可用区CBS云盘售罄、该可用区配额不足、该子网IP不足等。
	// 如果 Zones/SubnetIds 中可用区或者子网不存在，则无论 ZonesCheckPolicy 采用何种取值，都会校验报错。
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// 标签描述列表。通过指定该参数可以支持绑定标签到伸缩组。同时绑定标签到相应的资源实例。每个伸缩组最多支持30个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务设置，包括云监控不健康替换等服务设置。
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// 实例具有IPv6地址数量的配置，取值包括 0、1，默认值为0。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// 多可用区/子网策略，取值包括 PRIORITY 和 EQUALITY，默认为 PRIORITY。
	// <br><li> PRIORITY，按照可用区/子网列表的顺序，作为优先级来尝试创建实例，如果优先级最高的可用区/子网可以创建成功，则总在该可用区/子网创建。
	// <br><li> EQUALITY：扩容出的实例会打散到多个可用区/子网，保证扩容后的各个可用区/子网实例数相对均衡。
	// 
	// 与本策略相关的注意点：
	// <br><li> 当伸缩组为基础网络时，本策略适用于多可用区；当伸缩组为VPC网络时，本策略适用于多子网，此时不再考虑可用区因素，例如四个子网ABCD，其中ABC处于可用区1，D处于可用区2，此时考虑子网ABCD进行排序，而不考虑可用区1、2。
	// <br><li> 本策略适用于多可用区/子网，不适用于启动配置的多机型。多机型按照优先级策略进行选择。
	// <br><li> 按照 PRIORITY 策略创建实例时，先保证多机型的策略，后保证多可用区/子网的策略。例如多机型A、B，多子网1、2、3，会按照A1、A2、A3、B1、B2、B3 进行尝试，如果A1售罄，会尝试A2（而非B1）。
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// 伸缩组实例健康检查类型，取值如下：<br><li>CVM：根据实例网络状态判断实例是否处于不健康状态，不健康的网络状态即发生实例 PING 不可达事件，详细判断标准可参考[实例健康检查](https://cloud.tencent.com/document/product/377/8553)<br><li>CLB：根据 CLB 的健康检查状态判断实例是否处于不健康状态，CLB健康检查原理可参考[健康检查](https://cloud.tencent.com/document/product/214/6097) <br>如果选择了`CLB`类型，伸缩组将同时检查实例网络状态与CLB健康检查状态，如果出现实例网络状态不健康，实例将被标记为 UNHEALTHY 状态；如果出现 CLB 健康检查状态异常，实例将被标记为CLB_UNHEALTHY 状态，如果两个异常状态同时出现，实例`HealthStatus`字段将返回 UNHEALTHY|CLB_UNHEALTHY。默认值：CLB
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// CLB健康检查宽限期，当扩容的实例进入`IN_SERVICE`后，在宽限期时间范围内将不会被标记为不健康`CLB_UNHEALTHY`。<br>默认值：0。取值范围[0, 7200]，单位：秒。
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// 实例分配策略，取值包括 LAUNCH_CONFIGURATION 和 SPOT_MIXED，默认取 LAUNCH_CONFIGURATION。
	// <br><li> LAUNCH_CONFIGURATION，代表传统的按照启动配置模式。
	// <br><li> SPOT_MIXED，代表竞价混合模式。目前仅支持启动配置为按量计费模式时使用混合模式，混合模式下，伸缩组将根据设定扩容按量或竞价机型。使用混合模式时，关联的启动配置的计费类型不可被修改。
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// 竞价混合模式下，各计费类型实例的分配策略。
	// 仅当 InstanceAllocationPolicy 取 SPOT_MIXED 时可用。
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// 容量重平衡功能，仅对伸缩组内的竞价实例有效。取值范围：
	// <br><li> TRUE，开启该功能，当伸缩组内的竞价实例即将被竞价实例服务自动回收前，AS 主动发起竞价实例销毁流程，如果有配置过缩容 hook，则销毁前 hook 会生效。销毁流程启动后，AS 会异步开启一个扩容活动，用于补齐期望实例数。
	// <br><li> FALSE，不开启该功能，则 AS 等待竞价实例被销毁后才会去扩容补齐伸缩组期望实例数。
	// 
	// 默认取 FALSE。
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// 实例名称序号相关设置。若不指定该参数，则默认不开启。开启后为伸缩组内自动创建的实例名称添加递增的数字序号。
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
}

type CreateAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// VPC ID，基础网络则填空字符串
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 默认冷却时间，单位秒，默认值为300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// 期望实例数，大小介于最小实例数和最大实例数之间
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 传统负载均衡器ID列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 伸缩组内实例所属项目ID。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 应用型负载均衡器列表，目前长度上限为100，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// 子网ID列表，VPC场景下必须指定子网。多个子网以填写顺序为优先级，依次进行尝试，直至可以成功创建实例。
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 销毁策略，目前长度上限为1。取值包括 OLDEST_INSTANCE 和 NEWEST_INSTANCE，默认取值为 OLDEST_INSTANCE。
	// <br><li> OLDEST_INSTANCE 优先销毁伸缩组中最旧的实例。
	// <br><li> NEWEST_INSTANCE，优先销毁伸缩组中最新的实例。
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// 可用区列表，基础网络场景下必须指定可用区。多个可用区以填写顺序为优先级，依次进行尝试，直至可以成功创建实例。
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 重试策略，取值包括 IMMEDIATE_RETRY、 INCREMENTAL_INTERVALS、NO_RETRY，默认取值为 IMMEDIATE_RETRY。部分成功的伸缩活动判定为一次失败活动。
	// <br><li> IMMEDIATE_RETRY，立即重试，在较短时间内快速重试，连续失败超过一定次数（5次）后不再重试。
	// <br><li> INCREMENTAL_INTERVALS，间隔递增重试，随着连续失败次数的增加，重试间隔逐渐增大，重试间隔从秒级到1天不等。
	// <br><li> NO_RETRY，不进行重试，直到再次收到用户调用或者告警信息后才会重试。
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// 可用区校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有可用区（Zone）或子网（SubnetId）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个可用区（Zone）或子网（SubnetId）可用则通过校验，否则校验报错。
	// 
	// 可用区或子网不可用的常见原因包括该可用区CVM实例类型售罄、该可用区CBS云盘售罄、该可用区配额不足、该子网IP不足等。
	// 如果 Zones/SubnetIds 中可用区或者子网不存在，则无论 ZonesCheckPolicy 采用何种取值，都会校验报错。
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// 标签描述列表。通过指定该参数可以支持绑定标签到伸缩组。同时绑定标签到相应的资源实例。每个伸缩组最多支持30个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 服务设置，包括云监控不健康替换等服务设置。
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// 实例具有IPv6地址数量的配置，取值包括 0、1，默认值为0。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// 多可用区/子网策略，取值包括 PRIORITY 和 EQUALITY，默认为 PRIORITY。
	// <br><li> PRIORITY，按照可用区/子网列表的顺序，作为优先级来尝试创建实例，如果优先级最高的可用区/子网可以创建成功，则总在该可用区/子网创建。
	// <br><li> EQUALITY：扩容出的实例会打散到多个可用区/子网，保证扩容后的各个可用区/子网实例数相对均衡。
	// 
	// 与本策略相关的注意点：
	// <br><li> 当伸缩组为基础网络时，本策略适用于多可用区；当伸缩组为VPC网络时，本策略适用于多子网，此时不再考虑可用区因素，例如四个子网ABCD，其中ABC处于可用区1，D处于可用区2，此时考虑子网ABCD进行排序，而不考虑可用区1、2。
	// <br><li> 本策略适用于多可用区/子网，不适用于启动配置的多机型。多机型按照优先级策略进行选择。
	// <br><li> 按照 PRIORITY 策略创建实例时，先保证多机型的策略，后保证多可用区/子网的策略。例如多机型A、B，多子网1、2、3，会按照A1、A2、A3、B1、B2、B3 进行尝试，如果A1售罄，会尝试A2（而非B1）。
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// 伸缩组实例健康检查类型，取值如下：<br><li>CVM：根据实例网络状态判断实例是否处于不健康状态，不健康的网络状态即发生实例 PING 不可达事件，详细判断标准可参考[实例健康检查](https://cloud.tencent.com/document/product/377/8553)<br><li>CLB：根据 CLB 的健康检查状态判断实例是否处于不健康状态，CLB健康检查原理可参考[健康检查](https://cloud.tencent.com/document/product/214/6097) <br>如果选择了`CLB`类型，伸缩组将同时检查实例网络状态与CLB健康检查状态，如果出现实例网络状态不健康，实例将被标记为 UNHEALTHY 状态；如果出现 CLB 健康检查状态异常，实例将被标记为CLB_UNHEALTHY 状态，如果两个异常状态同时出现，实例`HealthStatus`字段将返回 UNHEALTHY|CLB_UNHEALTHY。默认值：CLB
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// CLB健康检查宽限期，当扩容的实例进入`IN_SERVICE`后，在宽限期时间范围内将不会被标记为不健康`CLB_UNHEALTHY`。<br>默认值：0。取值范围[0, 7200]，单位：秒。
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// 实例分配策略，取值包括 LAUNCH_CONFIGURATION 和 SPOT_MIXED，默认取 LAUNCH_CONFIGURATION。
	// <br><li> LAUNCH_CONFIGURATION，代表传统的按照启动配置模式。
	// <br><li> SPOT_MIXED，代表竞价混合模式。目前仅支持启动配置为按量计费模式时使用混合模式，混合模式下，伸缩组将根据设定扩容按量或竞价机型。使用混合模式时，关联的启动配置的计费类型不可被修改。
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// 竞价混合模式下，各计费类型实例的分配策略。
	// 仅当 InstanceAllocationPolicy 取 SPOT_MIXED 时可用。
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// 容量重平衡功能，仅对伸缩组内的竞价实例有效。取值范围：
	// <br><li> TRUE，开启该功能，当伸缩组内的竞价实例即将被竞价实例服务自动回收前，AS 主动发起竞价实例销毁流程，如果有配置过缩容 hook，则销毁前 hook 会生效。销毁流程启动后，AS 会异步开启一个扩容活动，用于补齐期望实例数。
	// <br><li> FALSE，不开启该功能，则 AS 等待竞价实例被销毁后才会去扩容补齐伸缩组期望实例数。
	// 
	// 默认取 FALSE。
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// 实例名称序号相关设置。若不指定该参数，则默认不开启。开启后为伸缩组内自动创建的实例名称添加递增的数字序号。
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
}

func (r *CreateAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupName")
	delete(f, "LaunchConfigurationId")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "VpcId")
	delete(f, "DefaultCooldown")
	delete(f, "DesiredCapacity")
	delete(f, "LoadBalancerIds")
	delete(f, "ProjectId")
	delete(f, "ForwardLoadBalancers")
	delete(f, "SubnetIds")
	delete(f, "TerminationPolicies")
	delete(f, "Zones")
	delete(f, "RetryPolicy")
	delete(f, "ZonesCheckPolicy")
	delete(f, "Tags")
	delete(f, "ServiceSettings")
	delete(f, "Ipv6AddressCount")
	delete(f, "MultiZoneSubnetPolicy")
	delete(f, "HealthCheckType")
	delete(f, "LoadBalancerHealthCheckGracePeriod")
	delete(f, "InstanceAllocationPolicy")
	delete(f, "SpotMixedAllocationPolicy")
	delete(f, "CapacityRebalance")
	delete(f, "InstanceNameIndexSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAutoScalingGroupResponseParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateAutoScalingGroupResponseParams `json:"Response"`
}

func (r *CreateAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchConfigurationRequestParams struct {
	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像Id与镜像族名称，二者必填一个且只能填写一个。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 启动配置所属项目ID。不填为默认项目。
	// 注意：伸缩组内实例所属项目ID取伸缩组项目ID，与这里取值无关。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749) 来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。
	// `InstanceType`和`InstanceTypes`参数互斥，二者必填一个且只能填写一个。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，最多支持指定11块数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。
	// <li>POSTPAID_BY_HOUR：按小时后付费</li>
	// <li>SPOTPAID：竞价付费</li>
	// <li>PREPAID：预付费，即包年包月</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 实例机型列表，不同实例机型指定了不同的资源规格，最多支持10种实例机型。
	// `InstanceType`和`InstanceTypes`参数互斥，二者必填一个且只能填写一个。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 实例类型校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。</li>
	// <li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。</li>
	// 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// 标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 标签描述列表。通过指定该参数可以支持绑定标签到启动配置。每个启动配置最多支持30个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// 云服务器实例名（InstanceName）的相关设置。
	// 如果用户在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 参照此字段进行设置，并传递给 CVM；如果用户未在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 按照“as-{{ 伸缩组AutoScalingGroupName }}”进行设置，并传递给 CVM。
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 云盘类型选择策略，默认取值 ORIGINAL，取值范围：
	// <li>ORIGINAL：使用设置的云盘类型</li>
	// <li>AUTOMATIC：自动选择当前可用的云盘类型</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// 高性能计算集群ID。<br>
	// 注意：此字段默认为空。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6公网带宽相关信息设置。若新建实例包含IPv6地址，该参数可为新建实例的IPv6地址分配公网带宽。关联启动配置的伸缩组Ipv6AddressCount参数为0时，该参数不会生效。
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// 置放群组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 镜像族名称。镜像Id与镜像族名称，二者必填一个且只能填写一个。
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

type CreateLaunchConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像Id与镜像族名称，二者必填一个且只能填写一个。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 启动配置所属项目ID。不填为默认项目。
	// 注意：伸缩组内实例所属项目ID取伸缩组项目ID，与这里取值无关。
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749) 来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。
	// `InstanceType`和`InstanceTypes`参数互斥，二者必填一个且只能填写一个。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，最多支持指定11块数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。
	// <li>POSTPAID_BY_HOUR：按小时后付费</li>
	// <li>SPOTPAID：竞价付费</li>
	// <li>PREPAID：预付费，即包年包月</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 实例机型列表，不同实例机型指定了不同的资源规格，最多支持10种实例机型。
	// `InstanceType`和`InstanceTypes`参数互斥，二者必填一个且只能填写一个。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 实例类型校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。</li>
	// <li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。</li>
	// 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// 标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 标签描述列表。通过指定该参数可以支持绑定标签到启动配置。每个启动配置最多支持30个标签。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// 云服务器实例名（InstanceName）的相关设置。
	// 如果用户在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 参照此字段进行设置，并传递给 CVM；如果用户未在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 按照“as-{{ 伸缩组AutoScalingGroupName }}”进行设置，并传递给 CVM。
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 云盘类型选择策略，默认取值 ORIGINAL，取值范围：
	// <li>ORIGINAL：使用设置的云盘类型</li>
	// <li>AUTOMATIC：自动选择当前可用的云盘类型</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// 高性能计算集群ID。<br>
	// 注意：此字段默认为空。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6公网带宽相关信息设置。若新建实例包含IPv6地址，该参数可为新建实例的IPv6地址分配公网带宽。关联启动配置的伸缩组Ipv6AddressCount参数为0时，该参数不会生效。
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// 置放群组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 镜像族名称。镜像Id与镜像族名称，二者必填一个且只能填写一个。
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

func (r *CreateLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationName")
	delete(f, "ImageId")
	delete(f, "ProjectId")
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "InternetAccessible")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "EnhancedService")
	delete(f, "UserData")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceMarketOptions")
	delete(f, "InstanceTypes")
	delete(f, "CamRoleName")
	delete(f, "InstanceTypesCheckPolicy")
	delete(f, "InstanceTags")
	delete(f, "Tags")
	delete(f, "HostNameSettings")
	delete(f, "InstanceNameSettings")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DiskTypePolicy")
	delete(f, "HpcClusterId")
	delete(f, "IPv6InternetAccessible")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "ImageFamily")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLaunchConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLaunchConfigurationResponseParams struct {
	// 当通过本接口来创建启动配置时会返回该参数，表示启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *CreateLaunchConfigurationResponseParams `json:"Response"`
}

func (r *CreateLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLaunchConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecycleHookRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 生命周期挂钩名称。名称仅支持中文、英文、数字、下划线（_）、短横线（-）、小数点（.），最大长度不能超128个字节。
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// 进行生命周期挂钩的场景，取值范围包括 INSTANCE_LAUNCHING 和 INSTANCE_TERMINATING
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值范围是 CONTINUE 或 ABANDON，默认值为 CONTINUE
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从30到7200秒，默认值为300秒
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息，配置通知时使用,默认值为空字符串""。最大长度不能超过1024个字节。
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// 通知目标。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// 进行生命周期挂钩的场景类型，取值范围包括NORMAL 和 EXTENSION。说明：设置为EXTENSION值，在AttachInstances、DetachInstances、RemoveInstaces接口时会触发生命周期挂钩操作，值为NORMAL则不会在这些接口中触发生命周期挂钩。
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// 远程命令执行对象。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type CreateLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 生命周期挂钩名称。名称仅支持中文、英文、数字、下划线（_）、短横线（-）、小数点（.），最大长度不能超128个字节。
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// 进行生命周期挂钩的场景，取值范围包括 INSTANCE_LAUNCHING 和 INSTANCE_TERMINATING
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值范围是 CONTINUE 或 ABANDON，默认值为 CONTINUE
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从30到7200秒，默认值为300秒
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息，配置通知时使用,默认值为空字符串""。最大长度不能超过1024个字节。
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// 通知目标。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// 进行生命周期挂钩的场景类型，取值范围包括NORMAL 和 EXTENSION。说明：设置为EXTENSION值，在AttachInstances、DetachInstances、RemoveInstaces接口时会触发生命周期挂钩操作，值为NORMAL则不会在这些接口中触发生命周期挂钩。
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// 远程命令执行对象。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

func (r *CreateLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LifecycleHookName")
	delete(f, "LifecycleTransition")
	delete(f, "DefaultResult")
	delete(f, "HeartbeatTimeout")
	delete(f, "NotificationMetadata")
	delete(f, "NotificationTarget")
	delete(f, "LifecycleTransitionType")
	delete(f, "LifecycleCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateLifecycleHookResponseParams struct {
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *CreateLifecycleHookResponseParams `json:"Response"`
}

func (r *CreateLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotificationConfigurationRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 通知类型，即为需要订阅的通知类型集合，取值范围如下：
	// <li>SCALE_OUT_SUCCESSFUL：扩容成功</li>
	// <li>SCALE_OUT_FAILED：扩容失败</li>
	// <li>SCALE_IN_SUCCESSFUL：缩容成功</li>
	// <li>SCALE_IN_FAILED：缩容失败</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL：替换不健康子机成功</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED：替换不健康子机失败</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[ListGroups](https://cloud.tencent.com/document/product/598/34589)查询。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// 通知接收端类型，取值如下
	// <br><li>USER_GROUP：用户组
	// <br><li>CMQ_QUEUE：CMQ 队列
	// <br><li>CMQ_TOPIC：CMQ 主题
	// <br><li>TDMQ_CMQ_TOPIC：TDMQ CMQ 主题
	// <br><li>TDMQ_CMQ_QUEUE：TDMQ CMQ 队列
	// 
	// 默认值为：`USER_GROUP`。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CMQ 队列名称，如 TargetType 取值为 `CMQ_QUEUE` 或 `TDMQ_CMQ_QUEUE` 时，该字段必填。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ 主题名称，如 TargetType 取值为 `CMQ_TOPIC` 或 `TDMQ_CMQ_TOPIC` 时，该字段必填。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type CreateNotificationConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 通知类型，即为需要订阅的通知类型集合，取值范围如下：
	// <li>SCALE_OUT_SUCCESSFUL：扩容成功</li>
	// <li>SCALE_OUT_FAILED：扩容失败</li>
	// <li>SCALE_IN_SUCCESSFUL：缩容成功</li>
	// <li>SCALE_IN_FAILED：缩容失败</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL：替换不健康子机成功</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED：替换不健康子机失败</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[ListGroups](https://cloud.tencent.com/document/product/598/34589)查询。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// 通知接收端类型，取值如下
	// <br><li>USER_GROUP：用户组
	// <br><li>CMQ_QUEUE：CMQ 队列
	// <br><li>CMQ_TOPIC：CMQ 主题
	// <br><li>TDMQ_CMQ_TOPIC：TDMQ CMQ 主题
	// <br><li>TDMQ_CMQ_QUEUE：TDMQ CMQ 队列
	// 
	// 默认值为：`USER_GROUP`。
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// CMQ 队列名称，如 TargetType 取值为 `CMQ_QUEUE` 或 `TDMQ_CMQ_QUEUE` 时，该字段必填。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ 主题名称，如 TargetType 取值为 `CMQ_TOPIC` 或 `TDMQ_CMQ_TOPIC` 时，该字段必填。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *CreateNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotificationConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "NotificationTypes")
	delete(f, "NotificationUserGroupIds")
	delete(f, "TargetType")
	delete(f, "QueueName")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateNotificationConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateNotificationConfigurationResponseParams struct {
	// 通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *CreateNotificationConfigurationResponseParams `json:"Response"`
}

func (r *CreateNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateNotificationConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScalingPolicyRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 告警触发策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// 告警触发策略类型，默认类型为SIMPLE。取值范围：<br><li>SIMPLE：简单策略</li><li>TARGET_TRACKING：目标追踪策略</li>
	ScalingPolicyType *string `json:"ScalingPolicyType,omitnil,omitempty" name:"ScalingPolicyType"`

	// 告警触发后，期望实例数修改方式，仅适用于简单策略。取值范围：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值，仅适用于简单策略。<br><li>当 AdjustmentType 为 CHANGE_IN_CAPACITY 时，AdjustmentValue 为正数表示告警触发后增加实例，为负数表示告警触发后减少实例 </li> <li> 当 AdjustmentType 为 EXACT_CAPACITY 时，AdjustmentValue 的值即为告警触发后新的期望实例数，需要大于或等于0 </li> <li> 当 AdjustmentType 为 PERCENT_CHANGE_IN_CAPACITY 时，AdjusmentValue 为正数表示告警触发后按百分比增加实例，为负数表示告警触发后按百分比减少实例，单位是：%。
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// 冷却时间，单位为秒，仅适用于简单策略。默认冷却时间300秒。
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// 告警监控指标，仅适用于简单策略。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// 预定义监控项，仅适用于目标追踪策略。取值范围：<br><li>ASG_AVG_CPU_UTILIZATION：平均CPU使用率</li><li>ASG_AVG_LAN_TRAFFIC_OUT：平均内网出带宽</li><li>ASG_AVG_LAN_TRAFFIC_IN：平均内网入带宽</li><li>ASG_AVG_WAN_TRAFFIC_OUT：平均外网出带宽</li><li>ASG_AVG_WAN_TRAFFIC_IN：平均外网出带宽</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// 目标值，仅适用于目标追踪策略。<br><li>ASG_AVG_CPU_UTILIZATION：[1, 100)，单位：%</li><li>ASG_AVG_LAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_IN：>0，单位：Mbps</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// 实例预热时间，单位为秒，仅适用于目标追踪策略。取值范围为0-3600，默认预热时间300秒。
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// 是否禁用缩容，仅适用于目标追踪策略，默认值为 false。取值范围：<br><li>true：目标追踪策略仅触发扩容</li><li>false：目标追踪策略触发扩容和缩容</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// 此参数已不再生效，请使用[创建通知](https://cloud.tencent.com/document/api/377/33185)。
	// 通知组ID，即为用户组ID集合。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

type CreateScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 告警触发策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// 告警触发策略类型，默认类型为SIMPLE。取值范围：<br><li>SIMPLE：简单策略</li><li>TARGET_TRACKING：目标追踪策略</li>
	ScalingPolicyType *string `json:"ScalingPolicyType,omitnil,omitempty" name:"ScalingPolicyType"`

	// 告警触发后，期望实例数修改方式，仅适用于简单策略。取值范围：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值，仅适用于简单策略。<br><li>当 AdjustmentType 为 CHANGE_IN_CAPACITY 时，AdjustmentValue 为正数表示告警触发后增加实例，为负数表示告警触发后减少实例 </li> <li> 当 AdjustmentType 为 EXACT_CAPACITY 时，AdjustmentValue 的值即为告警触发后新的期望实例数，需要大于或等于0 </li> <li> 当 AdjustmentType 为 PERCENT_CHANGE_IN_CAPACITY 时，AdjusmentValue 为正数表示告警触发后按百分比增加实例，为负数表示告警触发后按百分比减少实例，单位是：%。
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// 冷却时间，单位为秒，仅适用于简单策略。默认冷却时间300秒。
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// 告警监控指标，仅适用于简单策略。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// 预定义监控项，仅适用于目标追踪策略。取值范围：<br><li>ASG_AVG_CPU_UTILIZATION：平均CPU使用率</li><li>ASG_AVG_LAN_TRAFFIC_OUT：平均内网出带宽</li><li>ASG_AVG_LAN_TRAFFIC_IN：平均内网入带宽</li><li>ASG_AVG_WAN_TRAFFIC_OUT：平均外网出带宽</li><li>ASG_AVG_WAN_TRAFFIC_IN：平均外网出带宽</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// 目标值，仅适用于目标追踪策略。<br><li>ASG_AVG_CPU_UTILIZATION：[1, 100)，单位：%</li><li>ASG_AVG_LAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_IN：>0，单位：Mbps</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// 实例预热时间，单位为秒，仅适用于目标追踪策略。取值范围为0-3600，默认预热时间300秒。
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// 是否禁用缩容，仅适用于目标追踪策略，默认值为 false。取值范围：<br><li>true：目标追踪策略仅触发扩容</li><li>false：目标追踪策略触发扩容和缩容</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// 此参数已不再生效，请使用[创建通知](https://cloud.tencent.com/document/api/377/33185)。
	// 通知组ID，即为用户组ID集合。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

func (r *CreateScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScalingPolicyName")
	delete(f, "ScalingPolicyType")
	delete(f, "AdjustmentType")
	delete(f, "AdjustmentValue")
	delete(f, "Cooldown")
	delete(f, "MetricAlarm")
	delete(f, "PredefinedMetricType")
	delete(f, "TargetValue")
	delete(f, "EstimatedInstanceWarmup")
	delete(f, "DisableScaleIn")
	delete(f, "NotificationUserGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScalingPolicyResponseParams struct {
	// 告警触发策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *CreateScalingPolicyResponseParams `json:"Response"`
}

func (r *CreateScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduledActionRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 定时任务的首次触发时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 定时任务的结束时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。<br><br>此参数与`Recurrence`需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 定时任务的重复方式。为标准 Cron 格式<br><br>此参数与`EndTime`需要同时指定。
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

type CreateScheduledActionRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 定时任务的首次触发时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 定时任务的结束时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。<br><br>此参数与`Recurrence`需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 定时任务的重复方式。为标准 Cron 格式<br><br>此参数与`EndTime`需要同时指定。
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

func (r *CreateScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduledActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScheduledActionName")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "DesiredCapacity")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Recurrence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScheduledActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScheduledActionResponseParams struct {
	// 定时任务ID
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type CreateScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *CreateScheduledActionResponseParams `json:"Response"`
}

func (r *CreateScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScheduledActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// 数据盘类型。数据盘类型限制详见[云硬盘类型](https://cloud.tencent.com/document/product/362/2353)。取值范围：
	// <li>LOCAL_BASIC：本地硬盘</li>
	// <li>LOCAL_SSD：本地SSD硬盘</li>
	// <li>CLOUD_BASIC：普通云硬盘</li>
	// <li>CLOUD_PREMIUM：高性能云硬盘</li>
	// <li>CLOUD_SSD：SSD云硬盘</li>
	// <li>CLOUD_HSSD：增强型SSD云硬盘</li>
	// <li>CLOUD_TSSD：极速型SSD云硬盘</li>
	// 默认取值与系统盘类型（SystemDisk.DiskType）保持一致。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：[CVM实例配置](https://cloud.tencent.com/document/product/213/2177)。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`

	// 数据盘快照 ID，类似 `snap-l8psqwnt`。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotId *string `json:"SnapshotId,omitnil,omitempty" name:"SnapshotId"`

	// 数据盘是否随子机销毁。取值范围：
	// <li>TRUE：子机销毁时，销毁数据盘，只支持按小时后付费云盘</li>
	// <li>FALSE：子机销毁时，保留数据盘</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitnil,omitempty" name:"DeleteWithInstance"`

	// 数据盘是否加密。取值范围：
	// <li>TRUE：加密</li>
	// <li>FALSE：不加密</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Encrypt *bool `json:"Encrypt,omitnil,omitempty" name:"Encrypt"`

	// 云硬盘性能，单位：MB/s。使用此参数可给云硬盘购买额外的性能，功能介绍和类型限制详见：[增强型 SSD 云硬盘额外性能说明](https://cloud.tencent.com/document/product/362/51896#.E5.A2.9E.E5.BC.BA.E5.9E.8B-ssd-.E4.BA.91.E7.A1.AC.E7.9B.98.E9.A2.9D.E5.A4.96.E6.80.A7.E8.83.BD)。
	// 当前仅支持极速型云盘（CLOUD_TSSD）和增强型SSD云硬盘（CLOUD_HSSD）且 需容量 > 460GB。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitnil,omitempty" name:"ThroughputPerformance"`

	// 突发性能。是否开启突发性能，默认取值为 false。
	// 
	// 注：内测中，需提单申请后使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BurstPerformance *bool `json:"BurstPerformance,omitnil,omitempty" name:"BurstPerformance"`
}

// Predefined struct for user
type DeleteAutoScalingGroupRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

type DeleteAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

func (r *DeleteAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteAutoScalingGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteAutoScalingGroupResponseParams `json:"Response"`
}

func (r *DeleteAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchConfigurationRequestParams struct {
	// 需要删除的启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`
}

type DeleteLaunchConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 需要删除的启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`
}

func (r *DeleteLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLaunchConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLaunchConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLaunchConfigurationResponseParams `json:"Response"`
}

func (r *DeleteLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLaunchConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifecycleHookRequestParams struct {
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`
}

type DeleteLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`
}

func (r *DeleteLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteLifecycleHookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *DeleteLifecycleHookResponseParams `json:"Response"`
}

func (r *DeleteLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotificationConfigurationRequestParams struct {
	// 待删除的通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`
}

type DeleteNotificationConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`
}

func (r *DeleteNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotificationConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingNotificationId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNotificationConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNotificationConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNotificationConfigurationResponseParams `json:"Response"`
}

func (r *DeleteNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNotificationConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScalingPolicyRequestParams struct {
	// 待删除的告警策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`
}

type DeleteScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的告警策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScalingPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScalingPolicyResponseParams `json:"Response"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduledActionRequestParams struct {
	// 待删除的定时任务ID。
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`
}

type DeleteScheduledActionRequest struct {
	*tchttp.BaseRequest
	
	// 待删除的定时任务ID。
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`
}

func (r *DeleteScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduledActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduledActionId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScheduledActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScheduledActionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DeleteScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScheduledActionResponseParams `json:"Response"`
}

func (r *DeleteScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScheduledActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountLimitsRequestParams struct {

}

type DescribeAccountLimitsRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeAccountLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountLimitsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAccountLimitsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAccountLimitsResponseParams struct {
	// 用户账户被允许创建的启动配置最大数量
	MaxNumberOfLaunchConfigurations *int64 `json:"MaxNumberOfLaunchConfigurations,omitnil,omitempty" name:"MaxNumberOfLaunchConfigurations"`

	// 用户账户启动配置的当前数量
	NumberOfLaunchConfigurations *int64 `json:"NumberOfLaunchConfigurations,omitnil,omitempty" name:"NumberOfLaunchConfigurations"`

	// 用户账户被允许创建的伸缩组最大数量
	MaxNumberOfAutoScalingGroups *int64 `json:"MaxNumberOfAutoScalingGroups,omitnil,omitempty" name:"MaxNumberOfAutoScalingGroups"`

	// 用户账户伸缩组的当前数量
	NumberOfAutoScalingGroups *int64 `json:"NumberOfAutoScalingGroups,omitnil,omitempty" name:"NumberOfAutoScalingGroups"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAccountLimitsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAccountLimitsResponseParams `json:"Response"`
}

func (r *DescribeAccountLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAccountLimitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingActivitiesRequestParams struct {
	// 按照一个或者多个伸缩活动ID查询。伸缩活动ID形如：`asa-5l2ejpfo`。每次请求的上限为100。参数不支持同时指定`ActivityIds`和`Filters`。
	ActivityIds []*string `json:"ActivityIds,omitnil,omitempty" name:"ActivityIds"`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> activity-status-code - String - 是否必填：否 -（过滤条件）按照伸缩活动状态过滤。（INIT：初始化中|RUNNING：运行中|SUCCESSFUL：活动成功|PARTIALLY_SUCCESSFUL：活动部分成功|FAILED：活动失败|CANCELLED：活动取消）</li>
	// <li> activity-type - String - 是否必填：否 -（过滤条件）按照伸缩活动类型过滤。（SCALE_OUT：扩容活动|SCALE_IN：缩容活动|ATTACH_INSTANCES：添加实例|REMOVE_INSTANCES：销毁实例|DETACH_INSTANCES：移出实例|TERMINATE_INSTANCES_UNEXPECTEDLY：实例在CVM控制台被销毁|REPLACE_UNHEALTHY_INSTANCE：替换不健康实例|UPDATE_LOAD_BALANCERS：更新负载均衡器）</li>
	// <li> activity-id - String - 是否必填：否 -（过滤条件）按照伸缩活动ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`ActivityIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 伸缩活动最早的开始时间，如果指定了ActivityIds，此参数将被忽略。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 伸缩活动最晚的结束时间，如果指定了ActivityIds，此参数将被忽略。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type DescribeAutoScalingActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个伸缩活动ID查询。伸缩活动ID形如：`asa-5l2ejpfo`。每次请求的上限为100。参数不支持同时指定`ActivityIds`和`Filters`。
	ActivityIds []*string `json:"ActivityIds,omitnil,omitempty" name:"ActivityIds"`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> activity-status-code - String - 是否必填：否 -（过滤条件）按照伸缩活动状态过滤。（INIT：初始化中|RUNNING：运行中|SUCCESSFUL：活动成功|PARTIALLY_SUCCESSFUL：活动部分成功|FAILED：活动失败|CANCELLED：活动取消）</li>
	// <li> activity-type - String - 是否必填：否 -（过滤条件）按照伸缩活动类型过滤。（SCALE_OUT：扩容活动|SCALE_IN：缩容活动|ATTACH_INSTANCES：添加实例|REMOVE_INSTANCES：销毁实例|DETACH_INSTANCES：移出实例|TERMINATE_INSTANCES_UNEXPECTEDLY：实例在CVM控制台被销毁|REPLACE_UNHEALTHY_INSTANCE：替换不健康实例|UPDATE_LOAD_BALANCERS：更新负载均衡器）</li>
	// <li> activity-id - String - 是否必填：否 -（过滤条件）按照伸缩活动ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`ActivityIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 伸缩活动最早的开始时间，如果指定了ActivityIds，此参数将被忽略。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 伸缩活动最晚的结束时间，如果指定了ActivityIds，此参数将被忽略。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

func (r *DescribeAutoScalingActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ActivityIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "StartTime")
	delete(f, "EndTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingActivitiesResponseParams struct {
	// 符合条件的伸缩活动数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的伸缩活动信息集合。
	ActivitySet []*Activity `json:"ActivitySet,omitnil,omitempty" name:"ActivitySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingActivitiesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingAdvicesRequestParams struct {
	// 待查询的伸缩组列表，上限100。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

type DescribeAutoScalingAdvicesRequest struct {
	*tchttp.BaseRequest
	
	// 待查询的伸缩组列表，上限100。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

func (r *DescribeAutoScalingAdvicesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingAdvicesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingAdvicesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingAdvicesResponseParams struct {
	// 伸缩组配置建议集合。
	AutoScalingAdviceSet []*AutoScalingAdvice `json:"AutoScalingAdviceSet,omitnil,omitempty" name:"AutoScalingAdviceSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingAdvicesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingAdvicesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingAdvicesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingAdvicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupLastActivitiesRequestParams struct {
	// 伸缩组ID列表
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

type DescribeAutoScalingGroupLastActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID列表
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`
}

func (r *DescribeAutoScalingGroupLastActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupLastActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingGroupLastActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupLastActivitiesResponseParams struct {
	// 符合条件的伸缩活动信息集合。说明：伸缩组伸缩活动不存在的则不返回，如传50个伸缩组ID，返回45条数据，说明其中有5个伸缩组伸缩活动不存在。
	ActivitySet []*Activity `json:"ActivitySet,omitnil,omitempty" name:"ActivitySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingGroupLastActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingGroupLastActivitiesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingGroupLastActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupLastActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupsRequestParams struct {
	// 按照一个或者多个伸缩组ID查询。伸缩组ID形如：`asg-nkdwoui0`。每次请求的上限为100。参数不支持同时指定`AutoScalingGroupIds`和`Filters`。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> auto-scaling-group-name - String - 是否必填：否 -（过滤条件）按照伸缩组名称过滤。</li>
	// <li> vague-auto-scaling-group-name - String - 是否必填：否 -（过滤条件）按照伸缩组名称模糊搜索。</li>
	// <li> launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingGroupIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeAutoScalingGroupsRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个伸缩组ID查询。伸缩组ID形如：`asg-nkdwoui0`。每次请求的上限为100。参数不支持同时指定`AutoScalingGroupIds`和`Filters`。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitnil,omitempty" name:"AutoScalingGroupIds"`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> auto-scaling-group-name - String - 是否必填：否 -（过滤条件）按照伸缩组名称过滤。</li>
	// <li> vague-auto-scaling-group-name - String - 是否必填：否 -（过滤条件）按照伸缩组名称模糊搜索。</li>
	// <li> launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingGroupIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeAutoScalingGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingGroupsResponseParams struct {
	// 伸缩组详细信息列表。
	AutoScalingGroupSet []*AutoScalingGroup `json:"AutoScalingGroupSet,omitnil,omitempty" name:"AutoScalingGroupSet"`

	// 符合条件的伸缩组数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingGroupsResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingInstancesRequestParams struct {
	// 待查询云服务器（CVM）的实例ID。每次请求的上限为100。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤条件。
	// <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 待查询云服务器（CVM）的实例ID。每次请求的上限为100。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 过滤条件。
	// <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingInstancesResponseParams struct {
	// 实例详细信息列表。
	AutoScalingInstanceSet []*Instance `json:"AutoScalingInstanceSet,omitnil,omitempty" name:"AutoScalingInstanceSet"`

	// 符合条件的实例数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingInstancesResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchConfigurationsRequestParams struct {
	// 按照一个或者多个启动配置ID查询。启动配置ID形如：`asc-ouy1ax38`。每次请求的上限为100。参数不支持同时指定`LaunchConfigurationIds`和`Filters`
	LaunchConfigurationIds []*string `json:"LaunchConfigurationIds,omitnil,omitempty" name:"LaunchConfigurationIds"`

	// 过滤条件。
	// <li> launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。</li>
	// <li> launch-configuration-name - String - 是否必填：否 -（过滤条件）按照启动配置名称过滤。</li>
	// <li> vague-launch-configuration-name - String - 是否必填：否 -（过滤条件）按照启动配置名称模糊搜索。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例3</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LaunchConfigurationIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeLaunchConfigurationsRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个启动配置ID查询。启动配置ID形如：`asc-ouy1ax38`。每次请求的上限为100。参数不支持同时指定`LaunchConfigurationIds`和`Filters`
	LaunchConfigurationIds []*string `json:"LaunchConfigurationIds,omitnil,omitempty" name:"LaunchConfigurationIds"`

	// 过滤条件。
	// <li> launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。</li>
	// <li> launch-configuration-name - String - 是否必填：否 -（过滤条件）按照启动配置名称过滤。</li>
	// <li> vague-launch-configuration-name - String - 是否必填：否 -（过滤条件）按照启动配置名称模糊搜索。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例3</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LaunchConfigurationIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeLaunchConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchConfigurationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLaunchConfigurationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLaunchConfigurationsResponseParams struct {
	// 符合条件的启动配置数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 启动配置详细信息列表。
	LaunchConfigurationSet []*LaunchConfiguration `json:"LaunchConfigurationSet,omitnil,omitempty" name:"LaunchConfigurationSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLaunchConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLaunchConfigurationsResponseParams `json:"Response"`
}

func (r *DescribeLaunchConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLaunchConfigurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecycleHooksRequestParams struct {
	// 按照一个或者多个生命周期挂钩ID查询。生命周期挂钩ID形如：`ash-8azjzxcl`。每次请求的上限为100。参数不支持同时指定`LifecycleHookIds`和`Filters`。
	LifecycleHookIds []*string `json:"LifecycleHookIds,omitnil,omitempty" name:"LifecycleHookIds"`

	// 过滤条件。
	// <li> lifecycle-hook-id - String - 是否必填：否 -（过滤条件）按照生命周期挂钩ID过滤。</li>
	// <li> lifecycle-hook-name - String - 是否必填：否 -（过滤条件）按照生命周期挂钩名称过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LifecycleHookIds `和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeLifecycleHooksRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个生命周期挂钩ID查询。生命周期挂钩ID形如：`ash-8azjzxcl`。每次请求的上限为100。参数不支持同时指定`LifecycleHookIds`和`Filters`。
	LifecycleHookIds []*string `json:"LifecycleHookIds,omitnil,omitempty" name:"LifecycleHookIds"`

	// 过滤条件。
	// <li> lifecycle-hook-id - String - 是否必填：否 -（过滤条件）按照生命周期挂钩ID过滤。</li>
	// <li> lifecycle-hook-name - String - 是否必填：否 -（过滤条件）按照生命周期挂钩名称过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LifecycleHookIds `和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeLifecycleHooksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecycleHooksRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeLifecycleHooksRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeLifecycleHooksResponseParams struct {
	// 生命周期挂钩数组
	LifecycleHookSet []*LifecycleHook `json:"LifecycleHookSet,omitnil,omitempty" name:"LifecycleHookSet"`

	// 总体数量
	TotalCount *int64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeLifecycleHooksResponse struct {
	*tchttp.BaseResponse
	Response *DescribeLifecycleHooksResponseParams `json:"Response"`
}

func (r *DescribeLifecycleHooksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeLifecycleHooksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotificationConfigurationsRequestParams struct {
	// 按照一个或者多个通知ID查询。实例ID形如：asn-2sestqbr。每次请求的实例的上限为100。参数不支持同时指定`AutoScalingNotificationIds`和`Filters`。
	AutoScalingNotificationIds []*string `json:"AutoScalingNotificationIds,omitnil,omitempty" name:"AutoScalingNotificationIds"`

	// 过滤条件。
	// <li> auto-scaling-notification-id - String - 是否必填：否 -（过滤条件）按照通知ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingNotificationIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeNotificationConfigurationsRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个通知ID查询。实例ID形如：asn-2sestqbr。每次请求的实例的上限为100。参数不支持同时指定`AutoScalingNotificationIds`和`Filters`。
	AutoScalingNotificationIds []*string `json:"AutoScalingNotificationIds,omitnil,omitempty" name:"AutoScalingNotificationIds"`

	// 过滤条件。
	// <li> auto-scaling-notification-id - String - 是否必填：否 -（过滤条件）按照通知ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingNotificationIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeNotificationConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotificationConfigurationsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingNotificationIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNotificationConfigurationsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNotificationConfigurationsResponseParams struct {
	// 符合条件的通知数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 弹性伸缩事件通知详细信息列表。
	AutoScalingNotificationSet []*AutoScalingNotification `json:"AutoScalingNotificationSet,omitnil,omitempty" name:"AutoScalingNotificationSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeNotificationConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNotificationConfigurationsResponseParams `json:"Response"`
}

func (r *DescribeNotificationConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNotificationConfigurationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRefreshActivitiesRequestParams struct {
	// 刷新活动ID列表。ID形如：`asr-5l2ejpfo`。每次请求的上限为100。参数不支持同时指定`RefreshActivityIds`和`Filters`。
	RefreshActivityIds []*string `json:"RefreshActivityIds,omitnil,omitempty" name:"RefreshActivityIds"`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> refresh-activity-status-code - String - 是否必填：否 -（过滤条件）按照刷新活动状态过滤。（INIT：初始化中 | RUNNING：运行中 | SUCCESSFUL：活动成功 | FAILED_PAUSE：失败暂停 | AUTO_PAUSE：自动暂停 | MANUAL_PAUSE：手动暂停 | CANCELLED：活动取消 | FAILED：活动失败）</li>
	// <li> refresh-activity-type - String - 是否必填：否 -（过滤条件）按照刷新活动类型过滤。（NORMAL：正常刷新活动 | ROLLBACK：回滚刷新活动）</li>
	// <li> refresh-activity-id - String - 是否必填：否 -（过滤条件）按照刷新活动ID过滤。</li>
	// <li> 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定RefreshActivityIds和Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeRefreshActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 刷新活动ID列表。ID形如：`asr-5l2ejpfo`。每次请求的上限为100。参数不支持同时指定`RefreshActivityIds`和`Filters`。
	RefreshActivityIds []*string `json:"RefreshActivityIds,omitnil,omitempty" name:"RefreshActivityIds"`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> refresh-activity-status-code - String - 是否必填：否 -（过滤条件）按照刷新活动状态过滤。（INIT：初始化中 | RUNNING：运行中 | SUCCESSFUL：活动成功 | FAILED_PAUSE：失败暂停 | AUTO_PAUSE：自动暂停 | MANUAL_PAUSE：手动暂停 | CANCELLED：活动取消 | FAILED：活动失败）</li>
	// <li> refresh-activity-type - String - 是否必填：否 -（过滤条件）按照刷新活动类型过滤。（NORMAL：正常刷新活动 | ROLLBACK：回滚刷新活动）</li>
	// <li> refresh-activity-id - String - 是否必填：否 -（过滤条件）按照刷新活动ID过滤。</li>
	// <li> 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定RefreshActivityIds和Filters。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeRefreshActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRefreshActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RefreshActivityIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRefreshActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRefreshActivitiesResponseParams struct {
	// 符合条件的刷新活动数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 符合条件的刷新活动信息集合。
	RefreshActivitySet []*RefreshActivity `json:"RefreshActivitySet,omitnil,omitempty" name:"RefreshActivitySet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeRefreshActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRefreshActivitiesResponseParams `json:"Response"`
}

func (r *DescribeRefreshActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRefreshActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScalingPoliciesRequestParams struct {
	// 按照一个或者多个告警策略ID查询。告警策略ID形如：asp-i9vkg894。每次请求的实例的上限为100。参数不支持同时指定`AutoScalingPolicyIds`和`Filters`。
	AutoScalingPolicyIds []*string `json:"AutoScalingPolicyIds,omitnil,omitempty" name:"AutoScalingPolicyIds"`

	// 过滤条件。
	// <li> auto-scaling-policy-id - String - 是否必填：否 -（过滤条件）按照告警策略ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> scaling-policy-name - String - 是否必填：否 -（过滤条件）按照告警策略名称过滤。</li>
	// <li> scaling-policy-type - String - 是否必填：否 -（过滤条件）按照告警策略类型过滤，取值范围为SIMPLE，TARGET_TRACKING。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingPolicyIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

type DescribeScalingPoliciesRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个告警策略ID查询。告警策略ID形如：asp-i9vkg894。每次请求的实例的上限为100。参数不支持同时指定`AutoScalingPolicyIds`和`Filters`。
	AutoScalingPolicyIds []*string `json:"AutoScalingPolicyIds,omitnil,omitempty" name:"AutoScalingPolicyIds"`

	// 过滤条件。
	// <li> auto-scaling-policy-id - String - 是否必填：否 -（过滤条件）按照告警策略ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> scaling-policy-name - String - 是否必填：否 -（过滤条件）按照告警策略名称过滤。</li>
	// <li> scaling-policy-type - String - 是否必填：否 -（过滤条件）按照告警策略类型过滤，取值范围为SIMPLE，TARGET_TRACKING。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingPolicyIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`
}

func (r *DescribeScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScalingPoliciesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyIds")
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScalingPoliciesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScalingPoliciesResponseParams struct {
	// 弹性伸缩告警触发策略详细信息列表。
	ScalingPolicySet []*ScalingPolicy `json:"ScalingPolicySet,omitnil,omitempty" name:"ScalingPolicySet"`

	// 符合条件的通知数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScalingPoliciesResponseParams `json:"Response"`
}

func (r *DescribeScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScalingPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduledActionsRequestParams struct {
	// 按照一个或者多个定时任务ID查询。实例ID形如：asst-am691zxo。每次请求的实例的上限为100。参数不支持同时指定ScheduledActionIds和Filters。
	ScheduledActionIds []*string `json:"ScheduledActionIds,omitnil,omitempty" name:"ScheduledActionIds"`

	// 过滤条件。
	// <li> scheduled-action-id - String - 是否必填：否 -（过滤条件）按照定时任务ID过滤。</li>
	// <li> scheduled-action-name - String - 是否必填：否 - （过滤条件） 按照定时任务名称过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 - （过滤条件） 按照伸缩组ID过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

type DescribeScheduledActionsRequest struct {
	*tchttp.BaseRequest
	
	// 按照一个或者多个定时任务ID查询。实例ID形如：asst-am691zxo。每次请求的实例的上限为100。参数不支持同时指定ScheduledActionIds和Filters。
	ScheduledActionIds []*string `json:"ScheduledActionIds,omitnil,omitempty" name:"ScheduledActionIds"`

	// 过滤条件。
	// <li> scheduled-action-id - String - 是否必填：否 -（过滤条件）按照定时任务ID过滤。</li>
	// <li> scheduled-action-name - String - 是否必填：否 - （过滤条件） 按照定时任务名称过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 - （过滤条件） 按照伸缩组ID过滤。</li>
	Filters []*Filter `json:"Filters,omitnil,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitnil,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitnil,omitempty" name:"Limit"`
}

func (r *DescribeScheduledActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduledActionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduledActionIds")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScheduledActionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScheduledActionsResponseParams struct {
	// 符合条件的定时任务数量。
	TotalCount *uint64 `json:"TotalCount,omitnil,omitempty" name:"TotalCount"`

	// 定时任务详细信息列表。
	ScheduledActionSet []*ScheduledAction `json:"ScheduledActionSet,omitnil,omitempty" name:"ScheduledActionSet"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DescribeScheduledActionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScheduledActionsResponseParams `json:"Response"`
}

func (r *DescribeScheduledActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScheduledActionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type DetachInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *DetachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachInstancesResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DetachInstancesResponseParams `json:"Response"`
}

func (r *DetachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachLoadBalancersRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 传统负载均衡器ID列表，列表长度上限为20，LoadBalancerIds 和 ForwardLoadBalancerIdentifications 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 应用型负载均衡器标识信息列表，列表长度上限为100，LoadBalancerIds 和 ForwardLoadBalancerIdentifications二者同时最多只能指定一个
	ForwardLoadBalancerIdentifications []*ForwardLoadBalancerIdentification `json:"ForwardLoadBalancerIdentifications,omitnil,omitempty" name:"ForwardLoadBalancerIdentifications"`
}

type DetachLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 传统负载均衡器ID列表，列表长度上限为20，LoadBalancerIds 和 ForwardLoadBalancerIdentifications 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 应用型负载均衡器标识信息列表，列表长度上限为100，LoadBalancerIds 和 ForwardLoadBalancerIdentifications二者同时最多只能指定一个
	ForwardLoadBalancerIdentifications []*ForwardLoadBalancerIdentification `json:"ForwardLoadBalancerIdentifications,omitnil,omitempty" name:"ForwardLoadBalancerIdentifications"`
}

func (r *DetachLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LoadBalancerIds")
	delete(f, "ForwardLoadBalancerIdentifications")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DetachLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DetachLoadBalancersResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DetachLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *DetachLoadBalancersResponseParams `json:"Response"`
}

func (r *DetachLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DetachLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DetailedStatusMessage struct {
	// 错误类型。
	Code *string `json:"Code,omitnil,omitempty" name:"Code"`

	// 可用区信息。
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例计费类型。
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 子网ID。
	SubnetId *string `json:"SubnetId,omitnil,omitempty" name:"SubnetId"`

	// 错误描述。
	Message *string `json:"Message,omitnil,omitempty" name:"Message"`

	// 实例类型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`
}

// Predefined struct for user
type DisableAutoScalingGroupRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

type DisableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

func (r *DisableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DisableAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DisableAutoScalingGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type DisableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *DisableAutoScalingGroupResponseParams `json:"Response"`
}

func (r *DisableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DisableAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableAutoScalingGroupRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

type EnableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`
}

func (r *EnableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "EnableAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type EnableAutoScalingGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type EnableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *EnableAutoScalingGroupResponseParams `json:"Response"`
}

func (r *EnableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *EnableAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitnil,omitempty" name:"SecurityService"`

	// 开启云监控服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitnil,omitempty" name:"MonitorService"`

	// 该参数已废弃，查询时会返回空值，请勿使用。
	//
	// Deprecated: AutomationService is deprecated.
	AutomationService []*RunAutomationServiceEnabled `json:"AutomationService,omitnil,omitempty" name:"AutomationService"`

	// 开启自动化助手服务。若不指定该参数，则默认逻辑与CVM保持一致。注意：此字段可能返回 null，表示取不到有效值。
	AutomationToolsService *RunAutomationServiceEnabled `json:"AutomationToolsService,omitnil,omitempty" name:"AutomationToolsService"`
}

// Predefined struct for user
type ExecuteScalingPolicyRequestParams struct {
	// 告警伸缩策略ID，不支持目标追踪策略。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// 是否检查伸缩组活动处于冷却时间内，默认值为false
	HonorCooldown *bool `json:"HonorCooldown,omitnil,omitempty" name:"HonorCooldown"`

	// 执行伸缩策略的触发来源，取值包括 API 和 CLOUD_MONITOR，默认值为 API。CLOUD_MONITOR 专门供云监控触发调用。
	TriggerSource *string `json:"TriggerSource,omitnil,omitempty" name:"TriggerSource"`
}

type ExecuteScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 告警伸缩策略ID，不支持目标追踪策略。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// 是否检查伸缩组活动处于冷却时间内，默认值为false
	HonorCooldown *bool `json:"HonorCooldown,omitnil,omitempty" name:"HonorCooldown"`

	// 执行伸缩策略的触发来源，取值包括 API 和 CLOUD_MONITOR，默认值为 API。CLOUD_MONITOR 专门供云监控触发调用。
	TriggerSource *string `json:"TriggerSource,omitnil,omitempty" name:"TriggerSource"`
}

func (r *ExecuteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyId")
	delete(f, "HonorCooldown")
	delete(f, "TriggerSource")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExecuteScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExecuteScalingPolicyResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExecuteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ExecuteScalingPolicyResponseParams `json:"Response"`
}

func (r *ExecuteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExecuteScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExitStandbyRequestParams struct {
	// 伸缩组 ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 备用中状态 CVM 实例列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type ExitStandbyRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组 ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 备用中状态 CVM 实例列表。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *ExitStandbyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExitStandbyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ExitStandbyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ExitStandbyResponseParams struct {
	// 伸缩活动ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ExitStandbyResponse struct {
	*tchttp.BaseResponse
	Response *ExitStandbyResponseParams `json:"Response"`
}

func (r *ExitStandbyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ExitStandbyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {
	// 需要过滤的字段。
	Name *string `json:"Name,omitnil,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitnil,omitempty" name:"Values"`
}

type ForwardLoadBalancer struct {
	// 负载均衡器ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 应用型负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 目标规则属性列表
	TargetAttributes []*TargetAttribute `json:"TargetAttributes,omitnil,omitempty" name:"TargetAttributes"`

	// 转发规则ID，注意：针对七层监听器此参数必填
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`

	// 负载均衡实例所属地域，默认取AS服务所在地域。格式与公共参数Region相同，如："ap-guangzhou"。
	Region *string `json:"Region,omitnil,omitempty" name:"Region"`
}

type ForwardLoadBalancerIdentification struct {
	// 负载均衡器ID
	LoadBalancerId *string `json:"LoadBalancerId,omitnil,omitempty" name:"LoadBalancerId"`

	// 应用型负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitnil,omitempty" name:"ListenerId"`

	// 转发规则ID，注意：针对七层监听器此参数必填
	LocationId *string `json:"LocationId,omitnil,omitempty" name:"LocationId"`
}

type HostNameSettings struct {
	// 云服务器的主机名。
	// <li> 点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。</li> 
	// <li> 不支持 Windows 实例。</li> 
	// <li> 其他类型（Linux 等）实例：字符长度为[2, 40]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。不允许为纯数字。</li> 
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitnil,omitempty" name:"HostName"`

	// 云服务器主机名的风格，取值范围包括 ORIGINAL 和  UNIQUE，默认为 ORIGINAL。
	// <li> ORIGINAL，AS 直接将入参中所填的 HostName 传递给 CVM，CVM 可能会对 HostName 追加序列号，伸缩组中实例的 HostName 会出现冲突的情况。</li> 
	// <li> UNIQUE，入参所填的 HostName 相当于主机名前缀，AS 和 CVM 会对其进行拓展，伸缩组中实例的 HostName 可以保证唯一。</li> 
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostNameStyle *string `json:"HostNameStyle,omitnil,omitempty" name:"HostNameStyle"`

	// 云服务器的主机名后缀。
	// <li> 点号（.）和短横线（-）不能作为 HostNameSuffix 的首尾字符，不能连续使用。</li> 
	// <li> 不支持 Windows 实例。</li> 
	// <li>其他类型（Linux 等）实例：字符长度为[1, 37]，且与 HostName 的长度和不能超过 39，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。</li> 
	// 假设后缀名称为 suffix，原主机名为 test.0，最终主机名为 test.0.suffix。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostNameSuffix *string `json:"HostNameSuffix,omitnil,omitempty" name:"HostNameSuffix"`
}

type IPv6InternetAccessible struct {
	// 网络计费模式。取值包括TRAFFIC_POSTPAID_BY_HOUR、BANDWIDTH_PACKAGE，默认取值为TRAFFIC_POSTPAID_BY_HOUR。查看当前账户类型可参考[账户类型说明](https://cloud.tencent.com/document/product/1199/49090#judge)。
	// <br><li> IPv6对标准账户类型支持TRAFFIC_POSTPAID_BY_HOUR。
	// <br><li> IPv6对传统账户类型支持BANDWIDTH_PACKAGE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。<br>默认值：0，此时不为IPv6分配公网带宽。不同机型、可用区、计费模式的带宽上限范围不一致，具体限制详见[公网带宽上限](https://cloud.tencent.com/document/product/213/12523)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 带宽包ID。可通过[DescribeBandwidthPackages](https://cloud.tencent.com/document/api/215/19209)接口返回值中的`BandwidthPackageId`获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type Instance struct {
	// 实例ID
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 启动配置名称
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 生命周期状态，取值如下：<br>
	// <li>IN_SERVICE：运行中
	// <li>CREATING：创建中
	// <li>CREATION_FAILED：创建失败
	// <li>TERMINATING：中止中
	// <li>TERMINATION_FAILED：中止失败
	// <li>ATTACHING：绑定中
	// <li>ATTACH_FAILED：绑定失败
	// <li>DETACHING：解绑中
	// <li>DETACH_FAILED：解绑失败
	// <li>ATTACHING_LB：绑定LB中
	// <li>DETACHING_LB：解绑LB中
	// <li>MODIFYING_LB：修改LB中
	// <li>STARTING：开机中
	// <li>START_FAILED：开机失败
	// <li>STOPPING：关机中
	// <li>STOP_FAILED：关机失败
	// <li>STOPPED：已关机
	// <li>IN_LAUNCHING_HOOK：扩容生命周期挂钩中
	// <li>IN_TERMINATING_HOOK：缩容生命周期挂钩中
	LifeCycleState *string `json:"LifeCycleState,omitnil,omitempty" name:"LifeCycleState"`

	// 健康状态，取值包括HEALTHY和UNHEALTHY
	HealthStatus *string `json:"HealthStatus,omitnil,omitempty" name:"HealthStatus"`

	// 是否加入缩容保护
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`

	// 可用区
	Zone *string `json:"Zone,omitnil,omitempty" name:"Zone"`

	// 创建类型，取值包括AUTO_CREATION, MANUAL_ATTACHING。
	CreationType *string `json:"CreationType,omitnil,omitempty" name:"CreationType"`

	// 实例加入时间
	AddTime *string `json:"AddTime,omitnil,omitempty" name:"AddTime"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 版本号
	VersionNumber *int64 `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// 伸缩组名称
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 预热状态，取值如下：
	// <li>WAITING_ENTER_WARMUP：等待进入预热
	// <li>NO_NEED_WARMUP：无需预热
	// <li>IN_WARMUP：预热中
	// <li>AFTER_WARMUP：完成预热
	WarmupStatus *string `json:"WarmupStatus,omitnil,omitempty" name:"WarmupStatus"`

	// 置放群组id，仅支持指定一个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *int64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费</li><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费</li><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费</li>默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitnil,omitempty" name:"RenewFlag"`
}

type InstanceMarketOptionsRequest struct {
	// 竞价相关选项
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitnil,omitempty" name:"SpotOptions"`

	// 市场选项类型，当前只支持取值：spot
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarketType *string `json:"MarketType,omitnil,omitempty" name:"MarketType"`
}

type InstanceNameIndexSettings struct {
	// 是否开启实例创建序号，默认不开启。取值范围：<li>TRUE：表示开启实例创建序号<li>FALSE：表示不开启实例创建序号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`

	// 初始序号，取值范围为 [0, 99999999]。当序号递增后超出取值范围时，扩容活动会失败。<li>首次开启实例名称序号：默认值为 0。<li>非首次开启实例名称序号：若不指定该参数，沿用历史序号。下调初始序号可能会造成伸缩组内实例名称序号重复。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginIndex *int64 `json:"BeginIndex,omitnil,omitempty" name:"BeginIndex"`
}

type InstanceNameSettings struct {
	// 云服务器的实例名。字符长度为[2, 108]。
	InstanceName *string `json:"InstanceName,omitnil,omitempty" name:"InstanceName"`

	// 云服务器实例名的风格，取值范围包括 ORIGINAL 和 UNIQUE，默认为 ORIGINAL。
	// 
	// ORIGINAL，AS 直接将入参中所填的 InstanceName 传递给 CVM，CVM 可能会对 InstanceName 追加序列号，伸缩组中实例的 InstanceName 会出现冲突的情况。
	// 
	// UNIQUE，入参所填的 InstanceName 相当于实例名前缀，AS 和 CVM 会对其进行拓展，伸缩组中实例的 InstanceName 可以保证唯一。
	InstanceNameStyle *string `json:"InstanceNameStyle,omitnil,omitempty" name:"InstanceNameStyle"`

	// 云服务器实例名后缀。字符长度为[1,105]，且与 InstanceName 的长度和不能超过107。
	// 
	// 假设后缀名称为 suffix，原实例名为 test.0，最终实例名为 test.0.suffix。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceNameSuffix *string `json:"InstanceNameSuffix,omitnil,omitempty" name:"InstanceNameSuffix"`
}

type InstanceTag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`
}

type InternetAccessible struct {
	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：TRAFFIC_POSTPAID_BY_HOUR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetChargeType *string `json:"InternetChargeType,omitnil,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见[购买网络带宽](https://cloud.tencent.com/document/product/213/509)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitnil,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitnil,omitempty" name:"PublicIpAssigned"`

	// 带宽包ID。可通过[DescribeBandwidthPackages](https://cloud.tencent.com/document/api/215/19209)接口返回值中的`BandwidthPackageId`获取。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BandwidthPackageId *string `json:"BandwidthPackageId,omitnil,omitempty" name:"BandwidthPackageId"`
}

type InvocationResult struct {
	// 实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 执行活动ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationId *string `json:"InvocationId,omitnil,omitempty" name:"InvocationId"`

	// 执行任务ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InvocationTaskId *string `json:"InvocationTaskId,omitnil,omitempty" name:"InvocationTaskId"`

	// 命令ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommandId *string `json:"CommandId,omitnil,omitempty" name:"CommandId"`

	// 执行任务状态。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskStatus *string `json:"TaskStatus,omitnil,omitempty" name:"TaskStatus"`

	// 执行异常信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorMessage *string `json:"ErrorMessage,omitnil,omitempty" name:"ErrorMessage"`
}

type LaunchConfiguration struct {
	// 实例所属项目ID。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 启动配置名称。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitnil,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 实例登录设置。
	LoginSettings *LimitedLoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 公网带宽相关信息设置。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例所属安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 启动配置关联的伸缩组。
	AutoScalingGroupAbstractSet []*AutoScalingGroupAbstract `json:"AutoScalingGroupAbstractSet,omitnil,omitempty" name:"AutoScalingGroupAbstractSet"`

	// 自定义数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 启动配置创建时间。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 实例的增强服务启用情况与其设置。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 镜像ID。
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 启动配置当前状态。取值范围：<li>NORMAL：正常</li><li>IMAGE_ABNORMAL：启动配置镜像异常</li><li>CBS_SNAP_ABNORMAL：启动配置数据盘快照异常</li><li>SECURITY_GROUP_ABNORMAL：启动配置安全组异常</li>
	LaunchConfigurationStatus *string `json:"LaunchConfigurationStatus,omitnil,omitempty" name:"LaunchConfigurationStatus"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。<li>POSTPAID_BY_HOUR：按小时后付费</li><li>SPOTPAID：竞价付费</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 实例机型列表。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 实例标签列表。扩容出来的实例会自动带上标签，最多支持10个标签。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 标签列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*Tag `json:"Tags,omitnil,omitempty" name:"Tags"`

	// 版本号。
	VersionNumber *int64 `json:"VersionNumber,omitnil,omitempty" name:"VersionNumber"`

	// 更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitnil,omitempty" name:"UpdatedTime"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 上次操作时，InstanceTypesCheckPolicy 取值。
	LastOperationInstanceTypesCheckPolicy *string `json:"LastOperationInstanceTypesCheckPolicy,omitnil,omitempty" name:"LastOperationInstanceTypesCheckPolicy"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// 云服务器实例名（InstanceName）的相关设置。
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 云盘类型选择策略。取值范围：<li>ORIGINAL：使用设置的云盘类型</li><li>AUTOMATIC：自动选择当前可用区下可用的云盘类型</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// 高性能计算集群ID。<br>
	// 注意：此字段默认为空。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6公网带宽相关信息设置。
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// 置放群组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 镜像族名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

type LifecycleActionResultInfo struct {
	// 生命周期挂钩标识。
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 实例标识。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 执行活动ID。可通过TAT的[查询执行活动](https://cloud.tencent.com/document/api/1340/52679)API查询具体的执行结果。
	InvocationId *string `json:"InvocationId,omitnil,omitempty" name:"InvocationId"`

	// 命令调用的结果，表示执行TAT命令是否成功。<br>
	// <li>SUCCESSFUL 命令调用成功，不代表命令执行成功，执行的具体情况可根据InvocationId进行查询</li>
	// <li>FAILED 命令调用失败</li>
	// <li>NONE</li>
	InvokeCommandResult *string `json:"InvokeCommandResult,omitnil,omitempty" name:"InvokeCommandResult"`

	// 通知的结果，表示通知CMQ/TDMQ是否成功。<br>
	// <li>SUCCESSFUL 通知成功</li>
	// <li>FAILED 通知失败</li>
	// <li>NONE</li>
	NotificationResult *string `json:"NotificationResult,omitnil,omitempty" name:"NotificationResult"`

	// 生命周期挂钩动作的执行结果，取值包括 CONTINUE、ABANDON。
	LifecycleActionResult *string `json:"LifecycleActionResult,omitnil,omitempty" name:"LifecycleActionResult"`

	// 结果的原因。<br>
	// <li>HEARTBEAT_TIMEOUT 由于心跳超时，结果根据DefaultResult设置。</li>
	// <li>NOTIFICATION_FAILURE 由于发送通知失败，结果根据DefaultResult设置。</li>
	// <li>CALL_INTERFACE 调用了接口CompleteLifecycleAction设置结果。</li>
	// <li>ANOTHER_ACTION_ABANDON 另一个生命周期操作的结果已设置为“ABANDON”。</li>
	// <li>COMMAND_CALL_FAILURE  由于命令调用失败，结果根据DefaultResult设置。</li>
	// <li>COMMAND_EXEC_FINISH  命令执行完成。</li>
	// <li>COMMAND_EXEC_FAILURE 由于命令执行失败，结果根据DefaultResult设置。</li>
	// <li>COMMAND_EXEC_RESULT_CHECK_FAILURE 由于命令结果检查失败，结果根据DefaultResult设置。</li>
	ResultReason *string `json:"ResultReason,omitnil,omitempty" name:"ResultReason"`
}

type LifecycleCommand struct {
	// 远程命令ID。若选择执行命令，则此项必填。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CommandId *string `json:"CommandId,omitnil,omitempty" name:"CommandId"`

	// 自定义参数。字段类型为 json encoded string。如：{"varA": "222"}。
	// key为自定义参数名称，value为该参数的默认取值。kv均为字符串型。
	// 如果未提供该参数取值，将使用 Command 的 DefaultParameters 进行替换。
	// 自定义参数最多20个。自定义参数名称需符合以下规范：字符数目上限64，可选范围【a-zA-Z0-9-_】。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Parameters *string `json:"Parameters,omitnil,omitempty" name:"Parameters"`
}

type LifecycleHook struct {
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 生命周期挂钩名称
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 生命周期挂钩默认结果
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// 生命周期挂钩等待超时时间
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// 生命周期挂钩适用场景
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// 通知目标的附加信息
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 通知目标
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// 生命周期挂钩适用场景
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// 远程命令执行对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type LimitedLoginSettings struct {
	// 密钥ID列表。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitnil,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitnil,omitempty" name:"KeyIds"`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	KeepImageLogin *bool `json:"KeepImageLogin,omitnil,omitempty" name:"KeepImageLogin"`
}

type MetricAlarm struct {
	// 比较运算符，可选值：<br><li>GREATER_THAN：大于</li><li>GREATER_THAN_OR_EQUAL_TO：大于或等于</li><li>LESS_THAN：小于</li><li> LESS_THAN_OR_EQUAL_TO：小于或等于</li><li> EQUAL_TO：等于</li> <li>NOT_EQUAL_TO：不等于</li>
	ComparisonOperator *string `json:"ComparisonOperator,omitnil,omitempty" name:"ComparisonOperator"`

	// 指标名称，可选字段如下：<br><li>CPU_UTILIZATION：CPU利用率</li><li>MEM_UTILIZATION：内存利用率</li><li>LAN_TRAFFIC_OUT：内网出带宽</li><li>LAN_TRAFFIC_IN：内网入带宽</li><li>WAN_TRAFFIC_OUT：外网出带宽</li><li>WAN_TRAFFIC_IN：外网入带宽</li><li>TCP_CURR_ESTAB：TCP连接数</li>
	MetricName *string `json:"MetricName,omitnil,omitempty" name:"MetricName"`

	// 告警阈值：<br><li>CPU_UTILIZATION：[1, 100]，单位：%</li><li>MEM_UTILIZATION：[1, 100]，单位：%</li><li>LAN_TRAFFIC_OUT：>0，单位：Mbps </li><li>LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>WAN_TRAFFIC_IN：>0，单位：Mbps</li><li>TCP_CURR_ESTAB：>0, 单位：Count</li>
	Threshold *uint64 `json:"Threshold,omitnil,omitempty" name:"Threshold"`

	// 时间周期，单位：秒，取值枚举值为60、300。
	Period *uint64 `json:"Period,omitnil,omitempty" name:"Period"`

	// 重复次数。取值范围 [1, 10]
	ContinuousTime *uint64 `json:"ContinuousTime,omitnil,omitempty" name:"ContinuousTime"`

	// 统计类型，可选字段如下：<br><li>AVERAGE：平均值</li><li>MAXIMUM：最大值<li>MINIMUM：最小值</li><br> 默认取值：AVERAGE
	Statistic *string `json:"Statistic,omitnil,omitempty" name:"Statistic"`

	// 精确告警阈值，本参数不作为入参输入，仅用作查询接口出参：<br><li>CPU_UTILIZATION：(0, 100]，单位：%</li><li>MEM_UTILIZATION：(0, 100]，单位：%</li><li>LAN_TRAFFIC_OUT：>0，单位：Mbps </li><li>LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>WAN_TRAFFIC_IN：>0，单位：Mbps</li><li>TCP_CURR_ESTAB：>0, 单位：Count</li>
	PreciseThreshold *float64 `json:"PreciseThreshold,omitnil,omitempty" name:"PreciseThreshold"`
}

// Predefined struct for user
type ModifyAutoScalingGroupRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 默认冷却时间，单位秒，默认值为300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// 期望实例数，大小介于最小实例数和最大实例数之间
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 销毁策略，目前长度上限为1。取值包括 OLDEST_INSTANCE 和 NEWEST_INSTANCE。
	// <li> OLDEST_INSTANCE 优先销毁伸缩组中最旧的实例。</li>
	// <li> NEWEST_INSTANCE，优先销毁伸缩组中最新的实例。</li>
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// VPC ID，基础网络则填空字符串。修改为具体VPC ID时，需指定相应的SubnetIds；修改为基础网络时，需指定相应的Zones。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 可用区列表
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 重试策略，取值包括 IMMEDIATE_RETRY、 INCREMENTAL_INTERVALS、NO_RETRY，默认取值为 IMMEDIATE_RETRY。部分成功的伸缩活动判定为一次失败活动。
	// <li>IMMEDIATE_RETRY，立即重试，在较短时间内快速重试，连续失败超过一定次数（5次）后不再重试。</li>
	// <li>INCREMENTAL_INTERVALS，间隔递增重试，随着连续失败次数的增加，重试间隔逐渐增大，重试间隔从秒级到1天不等。</li>
	// <li> NO_RETRY，不进行重试，直到再次收到用户调用或者告警信息后才会重试。</li>
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// 可用区校验策略，取值包括 ALL 和 ANY，默认取值为ANY。在伸缩组实际变更资源相关字段时（启动配置、可用区、子网）发挥作用。
	// <li> ALL，所有可用区（Zone）或子网（SubnetId）都可用则通过校验，否则校验报错。</li>
	// <li> ANY，存在任何一个可用区（Zone）或子网（SubnetId）可用则通过校验，否则校验报错。</li>
	// 
	// 可用区或子网不可用的常见原因包括该可用区CVM实例类型售罄、该可用区CBS云盘售罄、该可用区配额不足、该子网IP不足等。
	// 如果 Zones/SubnetIds 中可用区或者子网不存在，则无论 ZonesCheckPolicy 采用何种取值，都会校验报错。
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// 服务设置，包括云监控不健康替换等服务设置。
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// 实例具有IPv6地址数量的配置，取值包括0、1。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// 多可用区/子网策略，取值包括 PRIORITY 和 EQUALITY，默认为 PRIORITY。
	// <li> PRIORITY，按照可用区/子网列表的顺序，作为优先级来尝试创建实例，如果优先级最高的可用区/子网可以创建成功，则总在该可用区/子网创建。</li> 
	// <li> EQUALITY：扩容出的实例会打散到多个可用区/子网，保证扩容后的各个可用区/子网实例数相对均衡。</li> 
	// 
	// 与本策略相关的注意点：
	// <li> 当伸缩组为基础网络时，本策略适用于多可用区；当伸缩组为VPC网络时，本策略适用于多子网，此时不再考虑可用区因素，例如四个子网ABCD，其中ABC处于可用区1，D处于可用区2，此时考虑子网ABCD进行排序，而不考虑可用区1、2。</li> 
	// <li> 本策略适用于多可用区/子网，不适用于启动配置的多机型。多机型按照优先级策略进行选择。</li> 
	// <li> 按照 PRIORITY 策略创建实例时，先保证多机型的策略，后保证多可用区/子网的策略。例如多机型A、B，多子网1、2、3，会按照A1、A2、A3、B1、B2、B3 进行尝试，如果A1售罄，会尝试A2（而非B1）。</li> 
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// 伸缩组实例健康检查类型，取值如下：
	// <li>CVM：根据实例网络状态判断实例是否处于不健康状态，不健康的网络状态即发生实例 PING 不可达事件，详细判断标准可参考[实例健康检查](https://cloud.tencent.com/document/product/377/8553)</li>
	// <li>CLB：根据 CLB 的健康检查状态判断实例是否处于不健康状态，CLB健康检查原理可参考[健康检查](https://cloud.tencent.com/document/product/214/6097)</li>
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// CLB健康检查宽限期。
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// 实例分配策略，取值包括 LAUNCH_CONFIGURATION 和 SPOT_MIXED。
	// <li> LAUNCH_CONFIGURATION，代表传统的按照启动配置模式。</li>
	// <li> SPOT_MIXED，代表竞价混合模式。目前仅支持启动配置为按量计费模式时使用混合模式，混合模式下，伸缩组将根据设定扩容按量或竞价机型。使用混合模式时，关联的启动配置的计费类型不可被修改。</li>
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// 竞价混合模式下，各计费类型实例的分配策略。
	// 仅当 InstanceAllocationPolicy 取 SPOT_MIXED 时可用。
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// 容量重平衡功能，仅对伸缩组内的竞价实例有效。取值范围：
	// <li> TRUE，开启该功能，当伸缩组内的竞价实例即将被竞价实例服务自动回收前，AS 主动发起竞价实例销毁流程，如果有配置过缩容 hook，则销毁前 hook 会生效。销毁流程启动后，AS 会异步开启一个扩容活动，用于补齐期望实例数。</li> 
	// <li> FALSE，不开启该功能，则 AS 等待竞价实例被销毁后才会去扩容补齐伸缩组期望实例数。</li> 
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// 实例名称序号相关设置。开启后为伸缩组内自动创建的实例名称添加递增的数字序号。
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
}

type ModifyAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitnil,omitempty" name:"AutoScalingGroupName"`

	// 默认冷却时间，单位秒，默认值为300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitnil,omitempty" name:"DefaultCooldown"`

	// 期望实例数，大小介于最小实例数和最大实例数之间
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitnil,omitempty" name:"SubnetIds"`

	// 销毁策略，目前长度上限为1。取值包括 OLDEST_INSTANCE 和 NEWEST_INSTANCE。
	// <li> OLDEST_INSTANCE 优先销毁伸缩组中最旧的实例。</li>
	// <li> NEWEST_INSTANCE，优先销毁伸缩组中最新的实例。</li>
	TerminationPolicies []*string `json:"TerminationPolicies,omitnil,omitempty" name:"TerminationPolicies"`

	// VPC ID，基础网络则填空字符串。修改为具体VPC ID时，需指定相应的SubnetIds；修改为基础网络时，需指定相应的Zones。
	VpcId *string `json:"VpcId,omitnil,omitempty" name:"VpcId"`

	// 可用区列表
	Zones []*string `json:"Zones,omitnil,omitempty" name:"Zones"`

	// 重试策略，取值包括 IMMEDIATE_RETRY、 INCREMENTAL_INTERVALS、NO_RETRY，默认取值为 IMMEDIATE_RETRY。部分成功的伸缩活动判定为一次失败活动。
	// <li>IMMEDIATE_RETRY，立即重试，在较短时间内快速重试，连续失败超过一定次数（5次）后不再重试。</li>
	// <li>INCREMENTAL_INTERVALS，间隔递增重试，随着连续失败次数的增加，重试间隔逐渐增大，重试间隔从秒级到1天不等。</li>
	// <li> NO_RETRY，不进行重试，直到再次收到用户调用或者告警信息后才会重试。</li>
	RetryPolicy *string `json:"RetryPolicy,omitnil,omitempty" name:"RetryPolicy"`

	// 可用区校验策略，取值包括 ALL 和 ANY，默认取值为ANY。在伸缩组实际变更资源相关字段时（启动配置、可用区、子网）发挥作用。
	// <li> ALL，所有可用区（Zone）或子网（SubnetId）都可用则通过校验，否则校验报错。</li>
	// <li> ANY，存在任何一个可用区（Zone）或子网（SubnetId）可用则通过校验，否则校验报错。</li>
	// 
	// 可用区或子网不可用的常见原因包括该可用区CVM实例类型售罄、该可用区CBS云盘售罄、该可用区配额不足、该子网IP不足等。
	// 如果 Zones/SubnetIds 中可用区或者子网不存在，则无论 ZonesCheckPolicy 采用何种取值，都会校验报错。
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitnil,omitempty" name:"ZonesCheckPolicy"`

	// 服务设置，包括云监控不健康替换等服务设置。
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitnil,omitempty" name:"ServiceSettings"`

	// 实例具有IPv6地址数量的配置，取值包括0、1。
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitnil,omitempty" name:"Ipv6AddressCount"`

	// 多可用区/子网策略，取值包括 PRIORITY 和 EQUALITY，默认为 PRIORITY。
	// <li> PRIORITY，按照可用区/子网列表的顺序，作为优先级来尝试创建实例，如果优先级最高的可用区/子网可以创建成功，则总在该可用区/子网创建。</li> 
	// <li> EQUALITY：扩容出的实例会打散到多个可用区/子网，保证扩容后的各个可用区/子网实例数相对均衡。</li> 
	// 
	// 与本策略相关的注意点：
	// <li> 当伸缩组为基础网络时，本策略适用于多可用区；当伸缩组为VPC网络时，本策略适用于多子网，此时不再考虑可用区因素，例如四个子网ABCD，其中ABC处于可用区1，D处于可用区2，此时考虑子网ABCD进行排序，而不考虑可用区1、2。</li> 
	// <li> 本策略适用于多可用区/子网，不适用于启动配置的多机型。多机型按照优先级策略进行选择。</li> 
	// <li> 按照 PRIORITY 策略创建实例时，先保证多机型的策略，后保证多可用区/子网的策略。例如多机型A、B，多子网1、2、3，会按照A1、A2、A3、B1、B2、B3 进行尝试，如果A1售罄，会尝试A2（而非B1）。</li> 
	MultiZoneSubnetPolicy *string `json:"MultiZoneSubnetPolicy,omitnil,omitempty" name:"MultiZoneSubnetPolicy"`

	// 伸缩组实例健康检查类型，取值如下：
	// <li>CVM：根据实例网络状态判断实例是否处于不健康状态，不健康的网络状态即发生实例 PING 不可达事件，详细判断标准可参考[实例健康检查](https://cloud.tencent.com/document/product/377/8553)</li>
	// <li>CLB：根据 CLB 的健康检查状态判断实例是否处于不健康状态，CLB健康检查原理可参考[健康检查](https://cloud.tencent.com/document/product/214/6097)</li>
	HealthCheckType *string `json:"HealthCheckType,omitnil,omitempty" name:"HealthCheckType"`

	// CLB健康检查宽限期。
	LoadBalancerHealthCheckGracePeriod *uint64 `json:"LoadBalancerHealthCheckGracePeriod,omitnil,omitempty" name:"LoadBalancerHealthCheckGracePeriod"`

	// 实例分配策略，取值包括 LAUNCH_CONFIGURATION 和 SPOT_MIXED。
	// <li> LAUNCH_CONFIGURATION，代表传统的按照启动配置模式。</li>
	// <li> SPOT_MIXED，代表竞价混合模式。目前仅支持启动配置为按量计费模式时使用混合模式，混合模式下，伸缩组将根据设定扩容按量或竞价机型。使用混合模式时，关联的启动配置的计费类型不可被修改。</li>
	InstanceAllocationPolicy *string `json:"InstanceAllocationPolicy,omitnil,omitempty" name:"InstanceAllocationPolicy"`

	// 竞价混合模式下，各计费类型实例的分配策略。
	// 仅当 InstanceAllocationPolicy 取 SPOT_MIXED 时可用。
	SpotMixedAllocationPolicy *SpotMixedAllocationPolicy `json:"SpotMixedAllocationPolicy,omitnil,omitempty" name:"SpotMixedAllocationPolicy"`

	// 容量重平衡功能，仅对伸缩组内的竞价实例有效。取值范围：
	// <li> TRUE，开启该功能，当伸缩组内的竞价实例即将被竞价实例服务自动回收前，AS 主动发起竞价实例销毁流程，如果有配置过缩容 hook，则销毁前 hook 会生效。销毁流程启动后，AS 会异步开启一个扩容活动，用于补齐期望实例数。</li> 
	// <li> FALSE，不开启该功能，则 AS 等待竞价实例被销毁后才会去扩容补齐伸缩组期望实例数。</li> 
	CapacityRebalance *bool `json:"CapacityRebalance,omitnil,omitempty" name:"CapacityRebalance"`

	// 实例名称序号相关设置。开启后为伸缩组内自动创建的实例名称添加递增的数字序号。
	InstanceNameIndexSettings *InstanceNameIndexSettings `json:"InstanceNameIndexSettings,omitnil,omitempty" name:"InstanceNameIndexSettings"`
}

func (r *ModifyAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "AutoScalingGroupName")
	delete(f, "DefaultCooldown")
	delete(f, "DesiredCapacity")
	delete(f, "LaunchConfigurationId")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "ProjectId")
	delete(f, "SubnetIds")
	delete(f, "TerminationPolicies")
	delete(f, "VpcId")
	delete(f, "Zones")
	delete(f, "RetryPolicy")
	delete(f, "ZonesCheckPolicy")
	delete(f, "ServiceSettings")
	delete(f, "Ipv6AddressCount")
	delete(f, "MultiZoneSubnetPolicy")
	delete(f, "HealthCheckType")
	delete(f, "LoadBalancerHealthCheckGracePeriod")
	delete(f, "InstanceAllocationPolicy")
	delete(f, "SpotMixedAllocationPolicy")
	delete(f, "CapacityRebalance")
	delete(f, "InstanceNameIndexSettings")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyAutoScalingGroupResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAutoScalingGroupResponseParams `json:"Response"`
}

func (r *ModifyAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDesiredCapacityRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 期望实例数
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

type ModifyDesiredCapacityRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 期望实例数
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`
}

func (r *ModifyDesiredCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDesiredCapacityRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "DesiredCapacity")
	delete(f, "MinSize")
	delete(f, "MaxSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyDesiredCapacityRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyDesiredCapacityResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyDesiredCapacityResponse struct {
	*tchttp.BaseResponse
	Response *ModifyDesiredCapacityResponseParams `json:"Response"`
}

func (r *ModifyDesiredCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyDesiredCapacityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchConfigurationAttributesRequestParams struct {
	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例类型列表，不同实例机型指定了不同的资源规格，最多支持10种实例机型。
	// InstanceType 指定单一实例类型，通过设置 InstanceTypes可以指定多实例类型，并使原有的InstanceType失效。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 实例类型校验策略，在实际修改 InstanceTypes 时发挥作用，取值包括 ALL 和 ANY，默认取值为ANY。
	// <li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。</li> 
	// <li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。</li> 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。如果要清空UserData，则指定其为空字符串。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。
	// 若指定该参数，请至少提供一个安全组，列表顺序有先后。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 公网带宽相关信息设置。
	// 当公网出带宽上限为0Mbps时，不支持修改为开通分配公网IP；相应的，当前为开通分配公网IP时，修改的公网出带宽上限值必须大于0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例计费类型。具体取值范围如下：
	// <li>POSTPAID_BY_HOUR：按小时后付费</li>
	// <li>SPOTPAID：竞价付费</li>
	// <li>PREPAID：预付费，即包年包月</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	// 若修改实例的付费模式为预付费，则该参数必传；从预付费修改为其他付费模式时，本字段原信息会自动丢弃。
	// 当新增该字段时，必须传递购买实例的时长，其它未传递字段会设置为默认值。
	// 当修改本字段时，当前付费模式必须为预付费。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例的市场相关选项，如竞价实例相关参数。
	// 若修改实例的付费模式为竞价付费，则该参数必传；从竞价付费修改为其他付费模式时，本字段原信息会自动丢弃。
	// 当新增该字段时，必须传递竞价相关选项下的竞价出价，其它未传递字段会设置为默认值。
	// 当修改本字段时，当前付费模式必须为竞价付费。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 云盘类型选择策略，取值范围：
	// <li>ORIGINAL：使用设置的云盘类型。</li>
	// <li>AUTOMATIC：自动选择当前可用的云盘类型。</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// 实例系统盘配置信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。
	// 最多支持指定11块数据盘。采取整体修改，因此请提供修改后的全部值。
	// 数据盘类型默认与系统盘类型保持一致。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 云服务器主机名（HostName）的相关设置。
	// 不支持windows实例设置主机名。
	// 新增该属性时，必须传递云服务器的主机名，其它未传递字段会设置为默认值。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// 云服务器（InstanceName）实例名的相关设置。 
	// 如果用户在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 参照此字段进行设置，并传递给 CVM；如果用户未在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 按照“as-{{ 伸缩组AutoScalingGroupName }}”进行设置，并传递给 CVM。
	// 新增该属性时，必须传递云服务器的实例名称，其它未传递字段会设置为默认值。
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。<br>
	// 注意：此字段默认为空。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6公网带宽相关信息设置。若新建实例包含IPv6地址，该参数可为新建实例的IPv6地址分配公网带宽。关联启动配置的伸缩组Ipv6AddressCount参数为0时，该参数不会生效。
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// 置放群组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 实例登录设置，包括密码、密钥或保持镜像的原始登录设置。<br>请注意，指定新的登录设置会覆盖原有登录设置。例如，如果您之前使用密码登录，使用该参数将登录设置修改为密钥，则原有密码被清除。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	// 该参数会覆盖原有的实例标签列表，如需新增标签，需将新标签和原有标签一并传入。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 镜像族名称。
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

type ModifyLaunchConfigurationAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例类型列表，不同实例机型指定了不同的资源规格，最多支持10种实例机型。
	// InstanceType 指定单一实例类型，通过设置 InstanceTypes可以指定多实例类型，并使原有的InstanceType失效。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 实例类型校验策略，在实际修改 InstanceTypes 时发挥作用，取值包括 ALL 和 ANY，默认取值为ANY。
	// <li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。</li> 
	// <li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。</li> 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。如果要清空UserData，则指定其为空字符串。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。
	// 若指定该参数，请至少提供一个安全组，列表顺序有先后。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 公网带宽相关信息设置。
	// 当公网出带宽上限为0Mbps时，不支持修改为开通分配公网IP；相应的，当前为开通分配公网IP时，修改的公网出带宽上限值必须大于0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 实例计费类型。具体取值范围如下：
	// <li>POSTPAID_BY_HOUR：按小时后付费</li>
	// <li>SPOTPAID：竞价付费</li>
	// <li>PREPAID：预付费，即包年包月</li>
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。
	// 若修改实例的付费模式为预付费，则该参数必传；从预付费修改为其他付费模式时，本字段原信息会自动丢弃。
	// 当新增该字段时，必须传递购买实例的时长，其它未传递字段会设置为默认值。
	// 当修改本字段时，当前付费模式必须为预付费。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 实例的市场相关选项，如竞价实例相关参数。
	// 若修改实例的付费模式为竞价付费，则该参数必传；从竞价付费修改为其他付费模式时，本字段原信息会自动丢弃。
	// 当新增该字段时，必须传递竞价相关选项下的竞价出价，其它未传递字段会设置为默认值。
	// 当修改本字段时，当前付费模式必须为竞价付费。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 云盘类型选择策略，取值范围：
	// <li>ORIGINAL：使用设置的云盘类型。</li>
	// <li>AUTOMATIC：自动选择当前可用的云盘类型。</li>
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// 实例系统盘配置信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。
	// 最多支持指定11块数据盘。采取整体修改，因此请提供修改后的全部值。
	// 数据盘类型默认与系统盘类型保持一致。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 云服务器主机名（HostName）的相关设置。
	// 不支持windows实例设置主机名。
	// 新增该属性时，必须传递云服务器的主机名，其它未传递字段会设置为默认值。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// 云服务器（InstanceName）实例名的相关设置。 
	// 如果用户在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 参照此字段进行设置，并传递给 CVM；如果用户未在启动配置中设置此字段，则伸缩组创建出的实例 InstanceName 按照“as-{{ 伸缩组AutoScalingGroupName }}”进行设置，并传递给 CVM。
	// 新增该属性时，必须传递云服务器的实例名称，其它未传递字段会设置为默认值。
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 高性能计算集群ID。<br>
	// 注意：此字段默认为空。
	HpcClusterId *string `json:"HpcClusterId,omitnil,omitempty" name:"HpcClusterId"`

	// IPv6公网带宽相关信息设置。若新建实例包含IPv6地址，该参数可为新建实例的IPv6地址分配公网带宽。关联启动配置的伸缩组Ipv6AddressCount参数为0时，该参数不会生效。
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`

	// 置放群组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `json:"DisasterRecoverGroupIds,omitnil,omitempty" name:"DisasterRecoverGroupIds"`

	// 实例登录设置，包括密码、密钥或保持镜像的原始登录设置。<br>请注意，指定新的登录设置会覆盖原有登录设置。例如，如果您之前使用密码登录，使用该参数将登录设置修改为密钥，则原有密码被清除。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	// 该参数会覆盖原有的实例标签列表，如需新增标签，需将新标签和原有标签一并传入。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// 镜像族名称。
	ImageFamily *string `json:"ImageFamily,omitnil,omitempty" name:"ImageFamily"`
}

func (r *ModifyLaunchConfigurationAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchConfigurationAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	delete(f, "ImageId")
	delete(f, "InstanceTypes")
	delete(f, "InstanceTypesCheckPolicy")
	delete(f, "LaunchConfigurationName")
	delete(f, "UserData")
	delete(f, "SecurityGroupIds")
	delete(f, "InternetAccessible")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceMarketOptions")
	delete(f, "DiskTypePolicy")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "HostNameSettings")
	delete(f, "InstanceNameSettings")
	delete(f, "EnhancedService")
	delete(f, "CamRoleName")
	delete(f, "HpcClusterId")
	delete(f, "IPv6InternetAccessible")
	delete(f, "DisasterRecoverGroupIds")
	delete(f, "LoginSettings")
	delete(f, "InstanceTags")
	delete(f, "ImageFamily")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLaunchConfigurationAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLaunchConfigurationAttributesResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLaunchConfigurationAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLaunchConfigurationAttributesResponseParams `json:"Response"`
}

func (r *ModifyLaunchConfigurationAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLaunchConfigurationAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifecycleHookRequestParams struct {
	// 生命周期挂钩ID。
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 生命周期挂钩名称。
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// 进入生命周期挂钩场景，取值包括：
	// <li> INSTANCE_LAUNCHING：实例启动后
	// <li> INSTANCE_TERMINATING：实例销毁前
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值包括：
	// <li> CONTINUE： 超时后继续伸缩活动
	// <li> ABANDON：超时后终止伸缩活动
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从 30 到 7200 秒。
	HeartbeatTimeout *uint64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息。
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// 进行生命周期挂钩的场景类型，取值范围包括`NORMAL`和 `EXTENSION`。说明：设置为`EXTENSION`值，在AttachInstances、DetachInstances、RemoveInstances 接口时会触发生命周期挂钩操作，值为`NORMAL`则不会在这些接口中触发生命周期挂钩。
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// 通知目标信息。
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// 远程命令执行对象。
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type ModifyLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期挂钩ID。
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 生命周期挂钩名称。
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// 进入生命周期挂钩场景，取值包括：
	// <li> INSTANCE_LAUNCHING：实例启动后
	// <li> INSTANCE_TERMINATING：实例销毁前
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值包括：
	// <li> CONTINUE： 超时后继续伸缩活动
	// <li> ABANDON：超时后终止伸缩活动
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从 30 到 7200 秒。
	HeartbeatTimeout *uint64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息。
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// 进行生命周期挂钩的场景类型，取值范围包括`NORMAL`和 `EXTENSION`。说明：设置为`EXTENSION`值，在AttachInstances、DetachInstances、RemoveInstances 接口时会触发生命周期挂钩操作，值为`NORMAL`则不会在这些接口中触发生命周期挂钩。
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// 通知目标信息。
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// 远程命令执行对象。
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

func (r *ModifyLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	delete(f, "LifecycleHookName")
	delete(f, "LifecycleTransition")
	delete(f, "DefaultResult")
	delete(f, "HeartbeatTimeout")
	delete(f, "NotificationMetadata")
	delete(f, "LifecycleTransitionType")
	delete(f, "NotificationTarget")
	delete(f, "LifecycleCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLifecycleHookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLifecycleHookResponseParams `json:"Response"`
}

func (r *ModifyLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerTargetAttributesRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 需修改目标规则属性的应用型负载均衡器列表，列表长度上限为100
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

type ModifyLoadBalancerTargetAttributesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 需修改目标规则属性的应用型负载均衡器列表，列表长度上限为100
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`
}

func (r *ModifyLoadBalancerTargetAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerTargetAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ForwardLoadBalancers")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancerTargetAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancerTargetAttributesResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancerTargetAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancerTargetAttributesResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancerTargetAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancerTargetAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancersRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 传统负载均衡器ID列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 应用型负载均衡器列表，目前长度上限为100，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// 负载均衡器校验策略，取值包括 ALL 和 DIFF，默认取值为 ALL。
	// <br><li> ALL，所有负载均衡器都合法则通过校验，否则校验报错。
	// <br><li> DIFF，仅校验负载均衡器参数中实际变化的部分，如果合法则通过校验，否则校验报错。
	LoadBalancersCheckPolicy *string `json:"LoadBalancersCheckPolicy,omitnil,omitempty" name:"LoadBalancersCheckPolicy"`
}

type ModifyLoadBalancersRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 传统负载均衡器ID列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitnil,omitempty" name:"LoadBalancerIds"`

	// 应用型负载均衡器列表，目前长度上限为100，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitnil,omitempty" name:"ForwardLoadBalancers"`

	// 负载均衡器校验策略，取值包括 ALL 和 DIFF，默认取值为 ALL。
	// <br><li> ALL，所有负载均衡器都合法则通过校验，否则校验报错。
	// <br><li> DIFF，仅校验负载均衡器参数中实际变化的部分，如果合法则通过校验，否则校验报错。
	LoadBalancersCheckPolicy *string `json:"LoadBalancersCheckPolicy,omitnil,omitempty" name:"LoadBalancersCheckPolicy"`
}

func (r *ModifyLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "LoadBalancerIds")
	delete(f, "ForwardLoadBalancers")
	delete(f, "LoadBalancersCheckPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyLoadBalancersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyLoadBalancersResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *ModifyLoadBalancersResponseParams `json:"Response"`
}

func (r *ModifyLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyLoadBalancersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotificationConfigurationRequestParams struct {
	// 待修改的通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// 通知类型，即为需要订阅的通知类型集合，取值范围如下：
	// <li>SCALE_OUT_SUCCESSFUL：扩容成功</li>
	// <li>SCALE_OUT_FAILED：扩容失败</li>
	// <li>SCALE_IN_SUCCESSFUL：缩容成功</li>
	// <li>SCALE_IN_FAILED：缩容失败</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL：替换不健康子机成功</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED：替换不健康子机失败</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[ListGroups](https://cloud.tencent.com/document/product/598/34589)查询。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// CMQ 队列或 TDMQ CMQ 队列名。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ 主题或 TDMQ CMQ 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type ModifyNotificationConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitnil,omitempty" name:"AutoScalingNotificationId"`

	// 通知类型，即为需要订阅的通知类型集合，取值范围如下：
	// <li>SCALE_OUT_SUCCESSFUL：扩容成功</li>
	// <li>SCALE_OUT_FAILED：扩容失败</li>
	// <li>SCALE_IN_SUCCESSFUL：缩容成功</li>
	// <li>SCALE_IN_FAILED：缩容失败</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL：替换不健康子机成功</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED：替换不健康子机失败</li>
	NotificationTypes []*string `json:"NotificationTypes,omitnil,omitempty" name:"NotificationTypes"`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[ListGroups](https://cloud.tencent.com/document/product/598/34589)查询。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`

	// CMQ 队列或 TDMQ CMQ 队列名。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// CMQ 主题或 TDMQ CMQ 主题名。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

func (r *ModifyNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotificationConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingNotificationId")
	delete(f, "NotificationTypes")
	delete(f, "NotificationUserGroupIds")
	delete(f, "QueueName")
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyNotificationConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyNotificationConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyNotificationConfigurationResponseParams `json:"Response"`
}

func (r *ModifyNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyNotificationConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScalingPolicyRequestParams struct {
	// 告警策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// 告警策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// 告警触发后，期望实例数修改方式，仅适用于简单策略。取值范围：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值，仅适用于简单策略。<li>当 AdjustmentType 为 CHANGE_IN_CAPACITY 时，AdjustmentValue 为正数表示告警触发后增加实例，为负数表示告警触发后减少实例 </li> <li> 当 AdjustmentType 为 EXACT_CAPACITY 时，AdjustmentValue 的值即为告警触发后新的期望实例数，需要大于或等于0 </li> <li> 当 AdjustmentType 为 PERCENT_CHANGE_IN_CAPACITY 时，AdjusmentValue 为正数表示告警触发后按百分比增加实例，为负数表示告警触发后按百分比减少实例，单位是：%。</li>
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// 冷却时间，仅适用于简单策略，单位为秒。
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// 告警监控指标，仅适用于简单策略。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// 预定义监控项，仅适用于目标追踪策略。取值范围：<br><li>ASG_AVG_CPU_UTILIZATION：平均CPU使用率</li><li>ASG_AVG_LAN_TRAFFIC_OUT：平均内网出带宽</li><li>ASG_AVG_LAN_TRAFFIC_IN：平均内网入带宽</li><li>ASG_AVG_WAN_TRAFFIC_OUT：平均外网出带宽</li><li>ASG_AVG_WAN_TRAFFIC_IN：平均外网出带宽</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// 目标值，仅适用于目标追踪策略。<br><li>ASG_AVG_CPU_UTILIZATION：[1, 100)，单位：%</li><li>ASG_AVG_LAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_IN：>0，单位：Mbps</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// 实例预热时间，单位为秒，仅适用于目标追踪策略。取值范围为0-3600。
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// 是否禁用缩容，仅适用于目标追踪策略。取值范围：<br><li>true：目标追踪策略仅触发扩容</li><li>false：目标追踪策略触发扩容和缩容</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// 此参数已不再生效，请使用[创建通知](https://cloud.tencent.com/document/api/377/33185)。
	// 通知组ID，即为用户组ID集合。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

type ModifyScalingPolicyRequest struct {
	*tchttp.BaseRequest
	
	// 告警策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// 告警策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// 告警触发后，期望实例数修改方式，仅适用于简单策略。取值范围：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值，仅适用于简单策略。<li>当 AdjustmentType 为 CHANGE_IN_CAPACITY 时，AdjustmentValue 为正数表示告警触发后增加实例，为负数表示告警触发后减少实例 </li> <li> 当 AdjustmentType 为 EXACT_CAPACITY 时，AdjustmentValue 的值即为告警触发后新的期望实例数，需要大于或等于0 </li> <li> 当 AdjustmentType 为 PERCENT_CHANGE_IN_CAPACITY 时，AdjusmentValue 为正数表示告警触发后按百分比增加实例，为负数表示告警触发后按百分比减少实例，单位是：%。</li>
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// 冷却时间，仅适用于简单策略，单位为秒。
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// 告警监控指标，仅适用于简单策略。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// 预定义监控项，仅适用于目标追踪策略。取值范围：<br><li>ASG_AVG_CPU_UTILIZATION：平均CPU使用率</li><li>ASG_AVG_LAN_TRAFFIC_OUT：平均内网出带宽</li><li>ASG_AVG_LAN_TRAFFIC_IN：平均内网入带宽</li><li>ASG_AVG_WAN_TRAFFIC_OUT：平均外网出带宽</li><li>ASG_AVG_WAN_TRAFFIC_IN：平均外网出带宽</li>
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// 目标值，仅适用于目标追踪策略。<br><li>ASG_AVG_CPU_UTILIZATION：[1, 100)，单位：%</li><li>ASG_AVG_LAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_IN：>0，单位：Mbps</li>
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// 实例预热时间，单位为秒，仅适用于目标追踪策略。取值范围为0-3600。
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// 是否禁用缩容，仅适用于目标追踪策略。取值范围：<br><li>true：目标追踪策略仅触发扩容</li><li>false：目标追踪策略触发扩容和缩容</li>
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// 此参数已不再生效，请使用[创建通知](https://cloud.tencent.com/document/api/377/33185)。
	// 通知组ID，即为用户组ID集合。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

func (r *ModifyScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScalingPolicyRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingPolicyId")
	delete(f, "ScalingPolicyName")
	delete(f, "AdjustmentType")
	delete(f, "AdjustmentValue")
	delete(f, "Cooldown")
	delete(f, "MetricAlarm")
	delete(f, "PredefinedMetricType")
	delete(f, "TargetValue")
	delete(f, "EstimatedInstanceWarmup")
	delete(f, "DisableScaleIn")
	delete(f, "NotificationUserGroupIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyScalingPolicyRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScalingPolicyResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *ModifyScalingPolicyResponseParams `json:"Response"`
}

func (r *ModifyScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScalingPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduledActionRequestParams struct {
	// 待修改的定时任务ID
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 定时任务的首次触发时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 定时任务的结束时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。<br>此参数与`Recurrence`需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 定时任务的重复方式。为标准 Cron 格式<br>此参数与`EndTime`需要同时指定。
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

type ModifyScheduledActionRequest struct {
	*tchttp.BaseRequest
	
	// 待修改的定时任务ID
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 定时任务的首次触发时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 定时任务的结束时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。<br>此参数与`Recurrence`需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 定时任务的重复方式。为标准 Cron 格式<br>此参数与`EndTime`需要同时指定。
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`
}

func (r *ModifyScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduledActionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ScheduledActionId")
	delete(f, "ScheduledActionName")
	delete(f, "MaxSize")
	delete(f, "MinSize")
	delete(f, "DesiredCapacity")
	delete(f, "StartTime")
	delete(f, "EndTime")
	delete(f, "Recurrence")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyScheduledActionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyScheduledActionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ModifyScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *ModifyScheduledActionResponseParams `json:"Response"`
}

func (r *ModifyScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyScheduledActionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NotificationTarget struct {
	// 目标类型，取值范围包括`CMQ_QUEUE`、`CMQ_TOPIC`、`TDMQ_CMQ_QUEUE`、`TDMQ_CMQ_TOPIC`。
	// <li> CMQ_QUEUE，指腾讯云消息队列-队列模型。</li>
	// <li> CMQ_TOPIC，指腾讯云消息队列-主题模型。</li>
	// <li> TDMQ_CMQ_QUEUE，指腾讯云 TDMQ 消息队列-队列模型。</li>
	// <li> TDMQ_CMQ_TOPIC，指腾讯云 TDMQ 消息队列-主题模型。</li>
	TargetType *string `json:"TargetType,omitnil,omitempty" name:"TargetType"`

	// 队列名称，如果`TargetType`取值为`CMQ_QUEUE` 或 `TDMQ_CMQ_QUEUE`，则本字段必填。
	QueueName *string `json:"QueueName,omitnil,omitempty" name:"QueueName"`

	// 主题名称，如果`TargetType`取值为`CMQ_TOPIC` 或 `TDMQ_CMQ_TOPIC`，则本字段必填。
	TopicName *string `json:"TopicName,omitnil,omitempty" name:"TopicName"`
}

type RefreshActivity struct {
	// 伸缩组 ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新活动 ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// 原始刷新活动ID，仅在回滚刷新活动中存在。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OriginRefreshActivityId *string `json:"OriginRefreshActivityId,omitnil,omitempty" name:"OriginRefreshActivityId"`

	// 刷新批次信息列表。
	RefreshBatchSet []*RefreshBatch `json:"RefreshBatchSet,omitnil,omitempty" name:"RefreshBatchSet"`

	// 刷新模式。
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`

	// 实例更新设置参数。
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// 刷新活动类型。取值如下：<br><li>NORMAL：正常刷新活动</li><li>ROLLBACK：回滚刷新活动
	ActivityType *string `json:"ActivityType,omitnil,omitempty" name:"ActivityType"`

	// 刷新活动状态。取值如下：<br><li>INIT：初始化中</li><li>RUNNING：运行中</li><li>SUCCESSFUL：活动成功</li><li>FAILED_PAUSE：因刷新批次失败暂停</li><li>AUTO_PAUSE：因暂停策略自动暂停</li><li>MANUAL_PAUSE：手动暂停</li><li>CANCELLED：活动取消</li><li>FAILED：活动失败
	Status *string `json:"Status,omitnil,omitempty" name:"Status"`

	// 当前刷新批次序号。例如，2 表示当前活动正在刷新第二批次的实例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentRefreshBatchNum *uint64 `json:"CurrentRefreshBatchNum,omitnil,omitempty" name:"CurrentRefreshBatchNum"`

	// 刷新活动开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 刷新活动结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 刷新活动创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`
}

type RefreshBatch struct {
	// 刷新批次序号。例如，2 表示当前批次实例会在第二批次进行实例刷新。
	RefreshBatchNum *uint64 `json:"RefreshBatchNum,omitnil,omitempty" name:"RefreshBatchNum"`

	// 刷新批次状态。取值如下：<br><li>WAITING：待刷新</li><li>INIT：初始化中</li><li>RUNNING：刷新中</li><li>FAILED:  刷新失败</li><li>PARTIALLY_SUCCESSFUL：批次部分成功</li><li>CANCELLED：已取消</li><li>SUCCESSFUL：刷新成功
	RefreshBatchStatus *string `json:"RefreshBatchStatus,omitnil,omitempty" name:"RefreshBatchStatus"`

	// 刷新批次关联实例列表。
	RefreshBatchRelatedInstanceSet []*RefreshBatchRelatedInstance `json:"RefreshBatchRelatedInstanceSet,omitnil,omitempty" name:"RefreshBatchRelatedInstanceSet"`

	// 刷新批次开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 刷新批次结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`
}

type RefreshBatchRelatedInstance struct {
	// 实例 ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 刷新实例状态。如果在刷新时实例被移出或销毁，状态会更新为 NOT_FOUND。取值如下：<br><li>WAITING：待刷新</li><li>INIT：初始化中</li><li>RUNNING：刷新中</li><li>FAILED：刷新失败</li><li>CANCELLED：已取消</li><li>SUCCESSFUL：刷新成功</li><li>NOT_FOUND：实例不存在
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`

	// 实例刷新中最近一次伸缩活动 ID，可通过 DescribeAutoScalingActivities 接口查询。
	// 需注意伸缩活动与实例刷新活动不同，一次实例刷新活动可能包括多次伸缩活动。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastActivityId *string `json:"LastActivityId,omitnil,omitempty" name:"LastActivityId"`

	// 实例刷新状态信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceStatusMessage *string `json:"InstanceStatusMessage,omitnil,omitempty" name:"InstanceStatusMessage"`
}

type RefreshSettings struct {
	// 滚动更新设置参数。RefreshMode 为滚动更新该参数必须填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	RollingUpdateSettings *RollingUpdateSettings `json:"RollingUpdateSettings,omitnil,omitempty" name:"RollingUpdateSettings"`

	// 实例后端服务健康状态检查，默认为 FALSE。仅针对绑定应用型负载均衡器的伸缩组生效，开启该检查后，如刷新后实例未通过检查，负载均衡器端口权重始终为 0，且标记为刷新失败。取值范围如下：<br><li>TRUE：开启检查</li><li>FALSE：不开启检查
	CheckInstanceTargetHealth *bool `json:"CheckInstanceTargetHealth,omitnil,omitempty" name:"CheckInstanceTargetHealth"`
}

type RelatedInstance struct {
	// 实例ID。
	InstanceId *string `json:"InstanceId,omitnil,omitempty" name:"InstanceId"`

	// 实例在伸缩活动中的状态。取值如下：
	// INIT：初始化中
	// RUNNING：实例操作中
	// SUCCESSFUL：活动成功
	// FAILED：活动失败
	InstanceStatus *string `json:"InstanceStatus,omitnil,omitempty" name:"InstanceStatus"`
}

// Predefined struct for user
type RemoveInstancesRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type RemoveInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *RemoveInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RemoveInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RemoveInstancesResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RemoveInstancesResponse struct {
	*tchttp.BaseResponse
	Response *RemoveInstancesResponseParams `json:"Response"`
}

func (r *RemoveInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RemoveInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeInstanceRefreshRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新活动ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// 当前批次刷新失败实例的恢复方式，如不存在失败实例，该参数无效。默认值为RETRY，取值范围如下：<br><li>RETRY: 重试当前批次刷新失败实例</li><li>CONTINUE: 跳过当前批次刷新失败实例
	ResumeMode *string `json:"ResumeMode,omitnil,omitempty" name:"ResumeMode"`
}

type ResumeInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新活动ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// 当前批次刷新失败实例的恢复方式，如不存在失败实例，该参数无效。默认值为RETRY，取值范围如下：<br><li>RETRY: 重试当前批次刷新失败实例</li><li>CONTINUE: 跳过当前批次刷新失败实例
	ResumeMode *string `json:"ResumeMode,omitnil,omitempty" name:"ResumeMode"`
}

func (r *ResumeInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshActivityId")
	delete(f, "ResumeMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResumeInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ResumeInstanceRefreshResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ResumeInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *ResumeInstanceRefreshResponseParams `json:"Response"`
}

func (r *ResumeInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResumeInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackInstanceRefreshRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新设置。
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// 原始刷新活动 ID。
	OriginRefreshActivityId *string `json:"OriginRefreshActivityId,omitnil,omitempty" name:"OriginRefreshActivityId"`

	// 刷新模式，目前仅支持滚动更新，默认值为 ROLLING_UPDATE_RESET。
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

type RollbackInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新设置。
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// 原始刷新活动 ID。
	OriginRefreshActivityId *string `json:"OriginRefreshActivityId,omitnil,omitempty" name:"OriginRefreshActivityId"`

	// 刷新模式，目前仅支持滚动更新，默认值为 ROLLING_UPDATE_RESET。
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

func (r *RollbackInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshSettings")
	delete(f, "OriginRefreshActivityId")
	delete(f, "RefreshMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RollbackInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type RollbackInstanceRefreshResponseParams struct {
	// 刷新活动 ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type RollbackInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *RollbackInstanceRefreshResponseParams `json:"Response"`
}

func (r *RollbackInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RollbackInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RollingUpdateSettings struct {
	// 批次数量。批次数量为大于 0 的正整数，但不能大于待刷新实例数量。
	BatchNumber *uint64 `json:"BatchNumber,omitnil,omitempty" name:"BatchNumber"`

	// 批次间暂停策略。默认值为 Automatic，取值范围如下：<br><li>FIRST_BATCH_PAUSE：第一批次更新完成后暂停</li><li>BATCH_INTERVAL_PAUSE：批次间暂停</li><li>AUTOMATIC：不暂停
	BatchPause *string `json:"BatchPause,omitnil,omitempty" name:"BatchPause"`
}

type RunAutomationServiceEnabled struct {
	// 是否开启[自动化助手](https://cloud.tencent.com/document/product/1340)服务。取值范围：<br><li>TRUE：表示开启自动化助手服务<br><li>FALSE：表示不开启自动化助手服务
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunMonitorServiceEnabled struct {
	// 是否开启[云监控](https://cloud.tencent.com/document/product/248)服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {
	// 是否开启[云安全](https://cloud.tencent.com/document/product/296)服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitnil,omitempty" name:"Enabled"`
}

// Predefined struct for user
type ScaleInInstancesRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 希望缩容的实例数量。
	ScaleInNumber *uint64 `json:"ScaleInNumber,omitnil,omitempty" name:"ScaleInNumber"`
}

type ScaleInInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 希望缩容的实例数量。
	ScaleInNumber *uint64 `json:"ScaleInNumber,omitnil,omitempty" name:"ScaleInNumber"`
}

func (r *ScaleInInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleInInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScaleInNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleInInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleInInstancesResponseParams struct {
	// 伸缩活动ID。
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleInInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ScaleInInstancesResponseParams `json:"Response"`
}

func (r *ScaleInInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleInInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstancesRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 希望扩容的实例数量。
	ScaleOutNumber *uint64 `json:"ScaleOutNumber,omitnil,omitempty" name:"ScaleOutNumber"`
}

type ScaleOutInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 希望扩容的实例数量。
	ScaleOutNumber *uint64 `json:"ScaleOutNumber,omitnil,omitempty" name:"ScaleOutNumber"`
}

func (r *ScaleOutInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "ScaleOutNumber")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScaleOutInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScaleOutInstancesResponseParams struct {
	// 伸缩活动ID。
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type ScaleOutInstancesResponse struct {
	*tchttp.BaseResponse
	Response *ScaleOutInstancesResponseParams `json:"Response"`
}

func (r *ScaleOutInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScaleOutInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScalingPolicy struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 告警触发策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitnil,omitempty" name:"AutoScalingPolicyId"`

	// 告警触发策略类型。取值：
	// - SIMPLE：简单策略
	// - TARGET_TRACKING：目标追踪策略
	ScalingPolicyType *string `json:"ScalingPolicyType,omitnil,omitempty" name:"ScalingPolicyType"`

	// 告警触发策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitnil,omitempty" name:"ScalingPolicyName"`

	// 告警触发后，期望实例数修改方式，仅适用于简单策略。取值范围：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitnil,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值，仅适用于简单策略。
	AdjustmentValue *int64 `json:"AdjustmentValue,omitnil,omitempty" name:"AdjustmentValue"`

	// 冷却时间，仅适用于简单策略。
	Cooldown *uint64 `json:"Cooldown,omitnil,omitempty" name:"Cooldown"`

	// 简单告警触发策略告警监控指标，仅适用于简单策略。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitnil,omitempty" name:"MetricAlarm"`

	// 预定义监控项，仅适用于目标追踪策略。取值范围：<br><li>ASG_AVG_CPU_UTILIZATION：平均CPU使用率</li><li>ASG_AVG_LAN_TRAFFIC_OUT：平均内网出带宽</li><li>ASG_AVG_LAN_TRAFFIC_IN：平均内网入带宽</li><li>ASG_AVG_WAN_TRAFFIC_OUT：平均外网出带宽</li><li>ASG_AVG_WAN_TRAFFIC_IN：平均外网出带宽</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	PredefinedMetricType *string `json:"PredefinedMetricType,omitnil,omitempty" name:"PredefinedMetricType"`

	// 目标值，仅适用于目标追踪策略。<br><li>ASG_AVG_CPU_UTILIZATION：[1, 100)，单位：%</li><li>ASG_AVG_LAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>ASG_AVG_WAN_TRAFFIC_IN：>0，单位：Mbps</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetValue *uint64 `json:"TargetValue,omitnil,omitempty" name:"TargetValue"`

	// 实例预热时间，单位为秒，仅适用于目标追踪策略。取值范围为0-3600。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EstimatedInstanceWarmup *uint64 `json:"EstimatedInstanceWarmup,omitnil,omitempty" name:"EstimatedInstanceWarmup"`

	// 是否禁用缩容，仅适用于目标追踪策略。取值范围：<br><li>true：目标追踪策略仅触发扩容</li><li>false：目标追踪策略触发扩容和缩容</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	DisableScaleIn *bool `json:"DisableScaleIn,omitnil,omitempty" name:"DisableScaleIn"`

	// 告警监控指标列表，仅适用于目标追踪策略。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricAlarms []*MetricAlarm `json:"MetricAlarms,omitnil,omitempty" name:"MetricAlarms"`

	// 通知组ID，即为用户组ID集合。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitnil,omitempty" name:"NotificationUserGroupIds"`
}

type ScheduledAction struct {
	// 定时任务ID。
	ScheduledActionId *string `json:"ScheduledActionId,omitnil,omitempty" name:"ScheduledActionId"`

	// 定时任务名称。
	ScheduledActionName *string `json:"ScheduledActionName,omitnil,omitempty" name:"ScheduledActionName"`

	// 定时任务所在伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 定时任务的开始时间。取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitnil,omitempty" name:"StartTime"`

	// 定时任务的重复方式。
	Recurrence *string `json:"Recurrence,omitnil,omitempty" name:"Recurrence"`

	// 定时任务的结束时间。取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	EndTime *string `json:"EndTime,omitnil,omitempty" name:"EndTime"`

	// 定时任务设置的最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitnil,omitempty" name:"MaxSize"`

	// 定时任务设置的期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitnil,omitempty" name:"DesiredCapacity"`

	// 定时任务设置的最小实例数。
	MinSize *uint64 `json:"MinSize,omitnil,omitempty" name:"MinSize"`

	// 定时任务的创建时间。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitnil,omitempty" name:"CreatedTime"`

	// 定时任务的执行类型。取值范围：<br><li>CRONTAB：代表定时任务为重复执行。<br><li>ONCE：代表定时任务为单次执行。
	ScheduledType *string `json:"ScheduledType,omitnil,omitempty" name:"ScheduledType"`
}

type ServiceSettings struct {
	// 开启监控不健康替换服务。若开启则对于云监控标记为不健康的实例，弹性伸缩服务会进行替换。若不指定该参数，则默认为 False。
	ReplaceMonitorUnhealthy *bool `json:"ReplaceMonitorUnhealthy,omitnil,omitempty" name:"ReplaceMonitorUnhealthy"`

	// 取值范围： 
	// CLASSIC_SCALING：经典方式，使用创建、销毁实例来实现扩缩容； 
	// WAKE_UP_STOPPED_SCALING：扩容优先开机。扩容时优先对已关机的实例执行开机操作，若开机后实例数仍低于期望实例数，则创建实例，缩容仍采用销毁实例的方式。用户可以使用StopAutoScalingInstances接口来关闭伸缩组内的实例。监控告警触发的扩容仍将创建实例
	// 默认取值：CLASSIC_SCALING
	ScalingMode *string `json:"ScalingMode,omitnil,omitempty" name:"ScalingMode"`

	// 开启负载均衡不健康替换服务。若开启则对于负载均衡健康检查判断不健康的实例，弹性伸缩服务会进行替换。若不指定该参数，则默认为 False。
	ReplaceLoadBalancerUnhealthy *bool `json:"ReplaceLoadBalancerUnhealthy,omitnil,omitempty" name:"ReplaceLoadBalancerUnhealthy"`

	// 不健康替换服务的替换模式。取值范围：
	// RECREATE：重建实例替代原有不健康实例；
	// RESET：对原有不健康实例进行重装系统操作，可保持数据盘、内网IP、实例id等信息不发生变化，实例登录设置、主机名、增强服务和 UserData 与当前启动配置保持一致。
	// 默认取值：RECREATE
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReplaceMode *string `json:"ReplaceMode,omitnil,omitempty" name:"ReplaceMode"`
}

// Predefined struct for user
type SetInstancesProtectionRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 实例ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例是否需要设置保护。
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`
}

type SetInstancesProtectionRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 实例ID。
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 实例是否需要设置保护。
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitnil,omitempty" name:"ProtectedFromScaleIn"`
}

func (r *SetInstancesProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetInstancesProtectionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	delete(f, "ProtectedFromScaleIn")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetInstancesProtectionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetInstancesProtectionResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type SetInstancesProtectionResponse struct {
	*tchttp.BaseResponse
	Response *SetInstancesProtectionResponseParams `json:"Response"`
}

func (r *SetInstancesProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetInstancesProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpotMarketOptions struct {
	// 竞价出价，例如“1.05”
	MaxPrice *string `json:"MaxPrice,omitnil,omitempty" name:"MaxPrice"`

	// 竞价请求类型，当前仅支持类型：one-time，默认值为one-time
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpotInstanceType *string `json:"SpotInstanceType,omitnil,omitempty" name:"SpotInstanceType"`
}

type SpotMixedAllocationPolicy struct {
	// 混合模式下，基础容量的大小，基础容量部分固定为按量计费实例。默认值 0，最大不可超过伸缩组的最大实例数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BaseCapacity *uint64 `json:"BaseCapacity,omitnil,omitempty" name:"BaseCapacity"`

	// 超出基础容量部分，按量计费实例所占的比例。取值范围 [0, 100]，0 代表超出基础容量的部分仅生产竞价实例，100 代表仅生产按量实例，默认值为 70。按百分比计算按量实例数时，向上取整。
	// 比如，总期望实例数取 3，基础容量取 1，超基础部分按量百分比取 1，则最终按量 2 台（1 台来自基础容量，1 台按百分比向上取整得到），竞价 1台。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OnDemandPercentageAboveBaseCapacity *uint64 `json:"OnDemandPercentageAboveBaseCapacity,omitnil,omitempty" name:"OnDemandPercentageAboveBaseCapacity"`

	// 混合模式下，竞价实例的分配策略。取值包括 COST_OPTIMIZED 和 CAPACITY_OPTIMIZED，默认取 COST_OPTIMIZED。
	// <br><li> COST_OPTIMIZED，成本优化策略。对于启动配置内的所有机型，按照各机型在各可用区的每核单价由小到大依次尝试。优先尝试购买每核单价最便宜的，如果购买失败则尝试购买次便宜的，以此类推。
	// <br><li> CAPACITY_OPTIMIZED，容量优化策略。对于启动配置内的所有机型，按照各机型在各可用区的库存情况由大到小依次尝试。优先尝试购买剩余库存最大的机型，这样可尽量降低竞价实例被动回收的发生概率。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpotAllocationStrategy *string `json:"SpotAllocationStrategy,omitnil,omitempty" name:"SpotAllocationStrategy"`

	// 按量实例替补功能。取值范围：
	// <br><li> TRUE，开启该功能，当所有竞价机型因库存不足等原因全部购买失败后，尝试购买按量实例。
	// <br><li> FALSE，不开启该功能，伸缩组在需要扩容竞价实例时仅尝试所配置的竞价机型。
	// 
	// 默认取值： TRUE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CompensateWithBaseInstance *bool `json:"CompensateWithBaseInstance,omitnil,omitempty" name:"CompensateWithBaseInstance"`
}

// Predefined struct for user
type StartAutoScalingInstancesRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 待开启的CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

type StartAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 待开启的CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`
}

func (r *StartAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAutoScalingInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartAutoScalingInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartAutoScalingInstancesResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StartAutoScalingInstancesResponseParams `json:"Response"`
}

func (r *StartAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartAutoScalingInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstanceRefreshRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新设置。
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// 刷新模式，目前仅支持滚动更新，默认值为 ROLLING_UPDATE_RESET。
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

type StartInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新设置。
	RefreshSettings *RefreshSettings `json:"RefreshSettings,omitnil,omitempty" name:"RefreshSettings"`

	// 刷新模式，目前仅支持滚动更新，默认值为 ROLLING_UPDATE_RESET。
	RefreshMode *string `json:"RefreshMode,omitnil,omitempty" name:"RefreshMode"`
}

func (r *StartInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshSettings")
	delete(f, "RefreshMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StartInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StartInstanceRefreshResponseParams struct {
	// 刷新活动 ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StartInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *StartInstanceRefreshResponseParams `json:"Response"`
}

func (r *StartInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StartInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAutoScalingInstancesRequestParams struct {
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 待关闭的CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 关闭的实例是否收费，取值为：  
	// KEEP_CHARGING：关机继续收费  
	// STOP_CHARGING：关机停止收费
	// 默认为 KEEP_CHARGING
	StoppedMode *string `json:"StoppedMode,omitnil,omitempty" name:"StoppedMode"`
}

type StopAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 待关闭的CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitnil,omitempty" name:"InstanceIds"`

	// 关闭的实例是否收费，取值为：  
	// KEEP_CHARGING：关机继续收费  
	// STOP_CHARGING：关机停止收费
	// 默认为 KEEP_CHARGING
	StoppedMode *string `json:"StoppedMode,omitnil,omitempty" name:"StoppedMode"`
}

func (r *StopAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoScalingInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "InstanceIds")
	delete(f, "StoppedMode")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopAutoScalingInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopAutoScalingInstancesResponseParams struct {
	// 伸缩活动ID
	ActivityId *string `json:"ActivityId,omitnil,omitempty" name:"ActivityId"`

	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *StopAutoScalingInstancesResponseParams `json:"Response"`
}

func (r *StopAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopAutoScalingInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceRefreshRequestParams struct {
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新活动ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

type StopInstanceRefreshRequest struct {
	*tchttp.BaseRequest
	
	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitnil,omitempty" name:"AutoScalingGroupId"`

	// 刷新活动ID。
	RefreshActivityId *string `json:"RefreshActivityId,omitnil,omitempty" name:"RefreshActivityId"`
}

func (r *StopInstanceRefreshRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRefreshRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AutoScalingGroupId")
	delete(f, "RefreshActivityId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "StopInstanceRefreshRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type StopInstanceRefreshResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type StopInstanceRefreshResponse struct {
	*tchttp.BaseResponse
	Response *StopInstanceRefreshResponseParams `json:"Response"`
}

func (r *StopInstanceRefreshResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *StopInstanceRefreshResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SystemDisk struct {
	// 系统盘类型。系统盘类型限制详见[云硬盘类型](https://cloud.tencent.com/document/product/362/2353)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：CLOUD_PREMIUM。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitnil,omitempty" name:"DiskType"`

	// 系统盘大小，单位：GB。默认值为 50
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *uint64 `json:"DiskSize,omitnil,omitempty" name:"DiskSize"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitnil,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitnil,omitempty" name:"Value"`

	// 标签绑定的资源类型，当前支持类型："auto-scaling-group
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitnil,omitempty" name:"ResourceType"`
}

type TargetAttribute struct {
	// 端口
	Port *uint64 `json:"Port,omitnil,omitempty" name:"Port"`

	// 权重
	Weight *uint64 `json:"Weight,omitnil,omitempty" name:"Weight"`
}

// Predefined struct for user
type UpgradeLaunchConfigurationRequestParams struct {
	// 启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例机型列表，不同实例机型指定了不同的资源规格，最多支持5种实例机型。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，最多支持指定11块数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。
	// <br><li>POSTPAID_BY_HOUR：按小时后付费
	// <br><li>SPOTPAID：竞价付费
	// <br><li>PREPAID：预付费，即包年包月
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 实例类型校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。
	// 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 该参数已失效，请勿使用。升级启动配置接口无法修改或覆盖 LoginSettings 参数，升级后 LoginSettings 不会发生变化。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属项目ID。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// 云服务器实例名（InstanceName）的相关设置。
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 云盘类型选择策略，取值范围：
	// <br><li>ORIGINAL：使用设置的云盘类型
	// <br><li>AUTOMATIC：自动选择当前可用的云盘类型
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// IPv6公网带宽相关信息设置。若新建实例包含IPv6地址，该参数可为新建实例的IPv6地址分配公网带宽。关联启动配置的伸缩组Ipv6AddressCount参数为0时，该参数不会生效。
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`
}

type UpgradeLaunchConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitnil,omitempty" name:"LaunchConfigurationId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitnil,omitempty" name:"ImageId"`

	// 实例机型列表，不同实例机型指定了不同的资源规格，最多支持5种实例机型。
	InstanceTypes []*string `json:"InstanceTypes,omitnil,omitempty" name:"InstanceTypes"`

	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitnil,omitempty" name:"LaunchConfigurationName"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，最多支持指定11块数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitnil,omitempty" name:"DataDisks"`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitnil,omitempty" name:"EnhancedService"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。
	// <br><li>POSTPAID_BY_HOUR：按小时后付费
	// <br><li>SPOTPAID：竞价付费
	// <br><li>PREPAID：预付费，即包年包月
	InstanceChargeType *string `json:"InstanceChargeType,omitnil,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitnil,omitempty" name:"InstanceMarketOptions"`

	// 实例类型校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。
	// 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitnil,omitempty" name:"InstanceTypesCheckPolicy"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitnil,omitempty" name:"InternetAccessible"`

	// 该参数已失效，请勿使用。升级启动配置接口无法修改或覆盖 LoginSettings 参数，升级后 LoginSettings 不会发生变化。
	LoginSettings *LoginSettings `json:"LoginSettings,omitnil,omitempty" name:"LoginSettings"`

	// 实例所属项目ID。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitnil,omitempty" name:"ProjectId"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitnil,omitempty" name:"SecurityGroupIds"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitnil,omitempty" name:"SystemDisk"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。
	UserData *string `json:"UserData,omitnil,omitempty" name:"UserData"`

	// 标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitnil,omitempty" name:"InstanceTags"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitnil,omitempty" name:"CamRoleName"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitnil,omitempty" name:"HostNameSettings"`

	// 云服务器实例名（InstanceName）的相关设置。
	InstanceNameSettings *InstanceNameSettings `json:"InstanceNameSettings,omitnil,omitempty" name:"InstanceNameSettings"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitnil,omitempty" name:"InstanceChargePrepaid"`

	// 云盘类型选择策略，取值范围：
	// <br><li>ORIGINAL：使用设置的云盘类型
	// <br><li>AUTOMATIC：自动选择当前可用的云盘类型
	DiskTypePolicy *string `json:"DiskTypePolicy,omitnil,omitempty" name:"DiskTypePolicy"`

	// IPv6公网带宽相关信息设置。若新建实例包含IPv6地址，该参数可为新建实例的IPv6地址分配公网带宽。关联启动配置的伸缩组Ipv6AddressCount参数为0时，该参数不会生效。
	IPv6InternetAccessible *IPv6InternetAccessible `json:"IPv6InternetAccessible,omitnil,omitempty" name:"IPv6InternetAccessible"`
}

func (r *UpgradeLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLaunchConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LaunchConfigurationId")
	delete(f, "ImageId")
	delete(f, "InstanceTypes")
	delete(f, "LaunchConfigurationName")
	delete(f, "DataDisks")
	delete(f, "EnhancedService")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceMarketOptions")
	delete(f, "InstanceTypesCheckPolicy")
	delete(f, "InternetAccessible")
	delete(f, "LoginSettings")
	delete(f, "ProjectId")
	delete(f, "SecurityGroupIds")
	delete(f, "SystemDisk")
	delete(f, "UserData")
	delete(f, "InstanceTags")
	delete(f, "CamRoleName")
	delete(f, "HostNameSettings")
	delete(f, "InstanceNameSettings")
	delete(f, "InstanceChargePrepaid")
	delete(f, "DiskTypePolicy")
	delete(f, "IPv6InternetAccessible")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeLaunchConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLaunchConfigurationResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeLaunchConfigurationResponseParams `json:"Response"`
}

func (r *UpgradeLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLaunchConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLifecycleHookRequestParams struct {
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 生命周期挂钩名称
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// 进行生命周期挂钩的场景，取值范围包括“INSTANCE_LAUNCHING”和“INSTANCE_TERMINATING”
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值范围是“CONTINUE”或“ABANDON”，默认值为“CONTINUE”
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从30到7200秒，默认值为300秒
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息，配置通知时使用，默认值为空字符串""
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// 通知目标。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// 进行生命周期挂钩的场景类型，取值范围包括NORMAL 和 EXTENSION。说明：设置为EXTENSION值，在AttachInstances、DetachInstances、RemoveInstaces接口时会触发生命周期挂钩操作，值为NORMAL则不会在这些接口中触发生命周期挂钩。
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// 远程命令执行对象。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

type UpgradeLifecycleHookRequest struct {
	*tchttp.BaseRequest
	
	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitnil,omitempty" name:"LifecycleHookId"`

	// 生命周期挂钩名称
	LifecycleHookName *string `json:"LifecycleHookName,omitnil,omitempty" name:"LifecycleHookName"`

	// 进行生命周期挂钩的场景，取值范围包括“INSTANCE_LAUNCHING”和“INSTANCE_TERMINATING”
	LifecycleTransition *string `json:"LifecycleTransition,omitnil,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值范围是“CONTINUE”或“ABANDON”，默认值为“CONTINUE”
	DefaultResult *string `json:"DefaultResult,omitnil,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从30到7200秒，默认值为300秒
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitnil,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息，配置通知时使用，默认值为空字符串""
	NotificationMetadata *string `json:"NotificationMetadata,omitnil,omitempty" name:"NotificationMetadata"`

	// 通知目标。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitnil,omitempty" name:"NotificationTarget"`

	// 进行生命周期挂钩的场景类型，取值范围包括NORMAL 和 EXTENSION。说明：设置为EXTENSION值，在AttachInstances、DetachInstances、RemoveInstaces接口时会触发生命周期挂钩操作，值为NORMAL则不会在这些接口中触发生命周期挂钩。
	LifecycleTransitionType *string `json:"LifecycleTransitionType,omitnil,omitempty" name:"LifecycleTransitionType"`

	// 远程命令执行对象。NotificationTarget和LifecycleCommand参数互斥，二者不可同时指定。
	LifecycleCommand *LifecycleCommand `json:"LifecycleCommand,omitnil,omitempty" name:"LifecycleCommand"`
}

func (r *UpgradeLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLifecycleHookRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "LifecycleHookId")
	delete(f, "LifecycleHookName")
	delete(f, "LifecycleTransition")
	delete(f, "DefaultResult")
	delete(f, "HeartbeatTimeout")
	delete(f, "NotificationMetadata")
	delete(f, "NotificationTarget")
	delete(f, "LifecycleTransitionType")
	delete(f, "LifecycleCommand")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpgradeLifecycleHookRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpgradeLifecycleHookResponseParams struct {
	// 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitnil,omitempty" name:"RequestId"`
}

type UpgradeLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *UpgradeLifecycleHookResponseParams `json:"Response"`
}

func (r *UpgradeLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpgradeLifecycleHookResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}