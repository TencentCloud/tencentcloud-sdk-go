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
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Activity struct {

	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 伸缩活动ID。
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 伸缩活动类型。取值如下：<br>
	// <li>SCALE_OUT：扩容活动<li>SCALE_IN：缩容活动<li>ATTACH_INSTANCES：添加实例<li>REMOVE_INSTANCES：销毁实例<li>DETACH_INSTANCES：移出实例<li>TERMINATE_INSTANCES_UNEXPECTEDLY：实例在CVM控制台被销毁<li>REPLACE_UNHEALTHY_INSTANCE：替换不健康实例）
	ActivityType *string `json:"ActivityType,omitempty" name:"ActivityType"`

	// 伸缩活动状态。取值如下：<br>
	// <li>INIT：初始化中
	// <li>RUNNING：运行中
	// <li>SUCCESSFUL：活动成功
	// <li>PARTIALLY_SUCCESSFUL：活动部分成功
	// <li>FAILED：活动失败
	// <li>CANCELLED：活动取消
	StatusCode *string `json:"StatusCode,omitempty" name:"StatusCode"`

	// 伸缩活动状态描述。
	StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`

	// 伸缩活动起因。
	Cause *string `json:"Cause,omitempty" name:"Cause"`

	// 伸缩活动描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 伸缩活动开始时间。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 伸缩活动结束时间。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 伸缩活动创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 伸缩活动相关实例信息集合。
	ActivityRelatedInstanceSet []*ActivtyRelatedInstance `json:"ActivityRelatedInstanceSet,omitempty" name:"ActivityRelatedInstanceSet" list`

	// 伸缩活动状态简要描述。
	StatusMessageSimplified *string `json:"StatusMessageSimplified,omitempty" name:"StatusMessageSimplified"`
}

type ActivtyRelatedInstance struct {

	// 实例ID。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例在伸缩活动中的状态。取值如下：
	// <li>INIT：初始化中
	// <li>RUNNING：实例操作中
	// <li>SUCCESSFUL：活动成功
	// <li>FAILED：活动失败
	InstanceStatus *string `json:"InstanceStatus,omitempty" name:"InstanceStatus"`
}

type AttachInstancesRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *AttachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 伸缩活动ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AutoScalingGroup struct {

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组名称
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`

	// 伸缩组当前状态。取值范围：<br><li>NORMAL：正常<br><li>CVM_ABNORMAL：启动配置异常<br><li>LB_ABNORMAL：负载均衡器异常<br><li>VPC_ABNORMAL：VPC网络异常<br><li>INSUFFICIENT_BALANCE：余额不足<br>
	AutoScalingGroupStatus *string `json:"AutoScalingGroupStatus,omitempty" name:"AutoScalingGroupStatus"`

	// 创建时间，采用UTC标准计时
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 默认冷却时间，单位秒
	DefaultCooldown *int64 `json:"DefaultCooldown,omitempty" name:"DefaultCooldown"`

	// 期望实例数
	DesiredCapacity *int64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// 启用状态，取值包括`ENABLED`和`DISABLED`
	EnabledStatus *string `json:"EnabledStatus,omitempty" name:"EnabledStatus"`

	// 应用型负载均衡器列表
	ForwardLoadBalancerSet []*ForwardLoadBalancer `json:"ForwardLoadBalancerSet,omitempty" name:"ForwardLoadBalancerSet" list`

	// 实例数量
	InstanceCount *int64 `json:"InstanceCount,omitempty" name:"InstanceCount"`

	// 状态为`IN_SERVICE`实例的数量
	InServiceInstanceCount *int64 `json:"InServiceInstanceCount,omitempty" name:"InServiceInstanceCount"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 启动配置名称
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// 传统型负载均衡器ID列表
	LoadBalancerIdSet []*string `json:"LoadBalancerIdSet,omitempty" name:"LoadBalancerIdSet" list`

	// 最大实例数
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 最小实例数
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子网ID列表
	SubnetIdSet []*string `json:"SubnetIdSet,omitempty" name:"SubnetIdSet" list`

	// 销毁策略
	TerminationPolicySet []*string `json:"TerminationPolicySet,omitempty" name:"TerminationPolicySet" list`

	// VPC标识
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 可用区列表
	ZoneSet []*string `json:"ZoneSet,omitempty" name:"ZoneSet" list`

	// 重试策略
	RetryPolicy *string `json:"RetryPolicy,omitempty" name:"RetryPolicy"`

	// 伸缩组是否处于伸缩活动中，`IN_ACTIVITY`表示处于伸缩活动中，`NOT_IN_ACTIVITY`表示不处于伸缩活动中。
	InActivityStatus *string `json:"InActivityStatus,omitempty" name:"InActivityStatus"`

	// 伸缩组标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// 服务设置
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitempty" name:"ServiceSettings"`
}

type AutoScalingGroupAbstract struct {

	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组名称。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`
}

type AutoScalingNotification struct {

	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 用户组ID列表。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`

	// 通知事件列表。
	NotificationTypes []*string `json:"NotificationTypes,omitempty" name:"NotificationTypes" list`

	// 事件通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`
}

type CompleteLifecycleActionRequest struct {
	*tchttp.BaseRequest

	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

	// 生命周期动作的结果，取值范围为“CONTINUE”或“ABANDON”
	LifecycleActionResult *string `json:"LifecycleActionResult,omitempty" name:"LifecycleActionResult"`

	// 实例ID，“InstanceId”和“LifecycleActionToken”必须填充其中一个
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// “InstanceId”和“LifecycleActionToken”必须填充其中一个
	LifecycleActionToken *string `json:"LifecycleActionToken,omitempty" name:"LifecycleActionToken"`
}

func (r *CompleteLifecycleActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteLifecycleActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CompleteLifecycleActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CompleteLifecycleActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CompleteLifecycleActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// VPC ID，基础网络则填空字符串
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 默认冷却时间，单位秒，默认值为300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitempty" name:"DefaultCooldown"`

	// 期望实例数，大小介于最小实例数和最大实例数之间
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// 传统负载均衡器ID列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 应用型负载均衡器列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitempty" name:"ForwardLoadBalancers" list`

	// 子网ID列表，VPC场景下必须指定子网。多个子网以填写顺序为优先级，依次进行尝试，直至可以成功创建实例。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 销毁策略，目前长度上限为1。取值包括 OLDEST_INSTANCE 和 NEWEST_INSTANCE，默认取值为 OLDEST_INSTANCE。
	// <br><li> OLDEST_INSTANCE 优先销毁伸缩组中最旧的实例。
	// <br><li> NEWEST_INSTANCE，优先销毁伸缩组中最新的实例。
	TerminationPolicies []*string `json:"TerminationPolicies,omitempty" name:"TerminationPolicies" list`

	// 可用区列表，基础网络场景下必须指定可用区。多个可用区以填写顺序为优先级，依次进行尝试，直至可以成功创建实例。
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 重试策略，取值包括 IMMEDIATE_RETRY、 INCREMENTAL_INTERVALS、NO_RETRY，默认取值为 IMMEDIATE_RETRY。
	// <br><li> IMMEDIATE_RETRY，立即重试，在较短时间内快速重试，连续失败超过一定次数（5次）后不再重试。
	// <br><li> INCREMENTAL_INTERVALS，间隔递增重试，随着连续失败次数的增加，重试间隔逐渐增大，重试间隔从秒级到1天不等。
	// <br><li> NO_RETRY，不进行重试，直到再次收到用户调用或者告警信息后才会重试。
	RetryPolicy *string `json:"RetryPolicy,omitempty" name:"RetryPolicy"`

	// 可用区校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有可用区（Zone）或子网（SubnetId）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个可用区（Zone）或子网（SubnetId）可用则通过校验，否则校验报错。
	// 
	// 可用区或子网不可用的常见原因包括该可用区CVM实例类型售罄、该可用区CBS云盘售罄、该可用区配额不足、该子网IP不足等。
	// 如果 Zones/SubnetIds 中可用区或者子网不存在，则无论 ZonesCheckPolicy 采用何种取值，都会校验报错。
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitempty" name:"ZonesCheckPolicy"`

	// 标签描述列表。通过指定该参数可以支持绑定标签到伸缩组。同时绑定标签到相应的资源实例，
	Tags []*Tag `json:"Tags,omitempty" name:"Tags" list`

	// 服务设置，包括云监控不健康替换等服务设置。
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitempty" name:"ServiceSettings"`
}

func (r *CreateAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 伸缩组ID
		AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLaunchConfigurationRequest struct {
	*tchttp.BaseRequest

	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/378/4400) 的返回值中的`projectId`字段来获取。不填为默认项目。
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749) 来获得最新的规格表或参见[实例类型](https://cloud.tencent.com/document/product/213/11518)描述。
	// `InstanceType`和`InstanceTypes`参数互斥，二者必填一个且只能填写一个。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，最多支持指定11块数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。
	// <br><li>POSTPAID_BY_HOUR：按小时后付费
	// <br><li>SPOTPAID：竞价付费
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// 实例机型列表，不同实例机型指定了不同的资源规格，最多支持10种实例机型。
	// `InstanceType`和`InstanceTypes`参数互斥，二者必填一个且只能填写一个。
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// 实例类型校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。
	// 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitempty" name:"InstanceTypesCheckPolicy"`

	// 标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitempty" name:"HostNameSettings"`
}

func (r *CreateLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLaunchConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当通过本接口来创建启动配置时会返回该参数，表示启动配置ID。
		LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLaunchConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLifecycleHookRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 生命周期挂钩名称。名称仅支持中文、英文、数字、下划线（_）、短横线（-）、小数点（.），最大长度不能超128个字节。
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" name:"LifecycleHookName"`

	// 进行生命周期挂钩的场景，取值范围包括 INSTANCE_LAUNCHING 和 INSTANCE_TERMINATING
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值范围是 CONTINUE 或 ABANDON，默认值为 CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从30到3600秒，默认值为300秒
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息，默认值为空字符串“”。最大长度不能超过1024个字节。
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" name:"NotificationMetadata"`

	// 通知目标
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitempty" name:"NotificationTarget"`
}

func (r *CreateLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLifecycleHookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生命周期挂钩ID
		LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateLifecycleHookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotificationConfigurationRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 通知类型，即为需要订阅的通知类型集合，取值范围如下：
	// <li>SCALE_OUT_SUCCESSFUL：扩容成功</li>
	// <li>SCALE_OUT_FAILED：扩容失败</li>
	// <li>SCALE_IN_SUCCESSFUL：缩容成功</li>
	// <li>SCALE_IN_FAILED：缩容失败</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL：替换不健康子机成功</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED：替换不健康子机失败</li>
	NotificationTypes []*string `json:"NotificationTypes,omitempty" name:"NotificationTypes" list`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[DescribeUserGroup](https://cloud.tencent.com/document/api/378/4404)查询。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *CreateNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotificationConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 通知ID。
		AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateNotificationConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePaiInstanceRequest struct {
	*tchttp.BaseRequest

	// PAI实例的域名。
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

	// 公网带宽相关信息设置。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 启动脚本的base64编码字符串。
	InitScript *string `json:"InitScript,omitempty" name:"InitScript"`

	// 可用区列表。
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// VpcId。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网列表。
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 实例显示名称。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例机型列表。
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// 实例登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例计费类型。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *CreatePaiInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePaiInstanceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePaiInstanceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例`ID`。返回实例`ID`列表并不代表实例创建成功，可根据 [DescribeInstances](https://cloud.tencent.com/document/api/213/15728) 接口查询返回的InstancesSet中对应实例的`ID`的状态来判断创建是否完成；如果实例状态由“准备中”变为“正在运行”，则为创建成功。
		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePaiInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePaiInstanceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 告警触发策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`

	// 告警触发后，期望实例数修改方式。取值 ：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值。取值：<br><li>当 AdjustmentType 为 CHANGE_IN_CAPACITY 时，AdjustmentValue 为正数表示告警触发后增加实例，为负数表示告警触发后减少实例 </li> <li> 当 AdjustmentType 为 EXACT_CAPACITY 时，AdjustmentValue 的值即为告警触发后新的期望实例数，需要大于或等于0 </li> <li> 当 AdjustmentType 为 PERCENT_CHANGE_IN_CAPACITY 时，AdjusmentValue 为正数表示告警触发后按百分比增加实例，为负数表示告警触发后按百分比减少实例，单位是：%。
	AdjustmentValue *int64 `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`

	// 告警监控指标。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitempty" name:"MetricAlarm"`

	// 冷却时间，单位为秒。默认冷却时间300秒。
	Cooldown *uint64 `json:"Cooldown,omitempty" name:"Cooldown"`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[DescribeUserGroup](https://cloud.tencent.com/document/api/378/4404)查询。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *CreateScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 告警触发策略ID。
		AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScheduledActionRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName *string `json:"ScheduledActionName,omitempty" name:"ScheduledActionName"`

	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// 定时任务的首次触发时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 定时任务的结束时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。<br><br>此参数与`Recurrence`需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 定时任务的重复方式。为标准 Cron 格式<br><br>此参数与`EndTime`需要同时指定。
	Recurrence *string `json:"Recurrence,omitempty" name:"Recurrence"`
}

func (r *CreateScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScheduledActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 定时任务ID
		ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateScheduledActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 数据盘类型。数据盘类型限制详见[CVM实例配置](https://cloud.tencent.com/document/product/213/2177)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：[CVM实例配置](https://cloud.tencent.com/document/product/213/2177)。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘快照 ID，类似 `snap-l8psqwnt`。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
}

type DeleteAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
}

func (r *DeleteAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLaunchConfigurationRequest struct {
	*tchttp.BaseRequest

	// 需要删除的启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`
}

func (r *DeleteLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLaunchConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLaunchConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLifecycleHookRequest struct {
	*tchttp.BaseRequest

	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`
}

func (r *DeleteLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLifecycleHookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteLifecycleHookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNotificationConfigurationRequest struct {
	*tchttp.BaseRequest

	// 待删除的通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`
}

func (r *DeleteNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNotificationConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteNotificationConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 待删除的告警策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScheduledActionRequest struct {
	*tchttp.BaseRequest

	// 待删除的定时任务ID。
	ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`
}

func (r *DeleteScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScheduledActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScheduledActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountLimitsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountLimitsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountLimitsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountLimitsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 用户账户被允许创建的启动配置最大数量
		MaxNumberOfLaunchConfigurations *int64 `json:"MaxNumberOfLaunchConfigurations,omitempty" name:"MaxNumberOfLaunchConfigurations"`

		// 用户账户启动配置的当前数量
		NumberOfLaunchConfigurations *int64 `json:"NumberOfLaunchConfigurations,omitempty" name:"NumberOfLaunchConfigurations"`

		// 用户账户被允许创建的伸缩组最大数量
		MaxNumberOfAutoScalingGroups *int64 `json:"MaxNumberOfAutoScalingGroups,omitempty" name:"MaxNumberOfAutoScalingGroups"`

		// 用户账户伸缩组的当前数量
		NumberOfAutoScalingGroups *int64 `json:"NumberOfAutoScalingGroups,omitempty" name:"NumberOfAutoScalingGroups"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountLimitsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAccountLimitsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingActivitiesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个伸缩活动ID查询。伸缩活动ID形如：`asa-5l2ejpfo`。每次请求的上限为100。参数不支持同时指定`ActivityIds`和`Filters`。
	ActivityIds []*string `json:"ActivityIds,omitempty" name:"ActivityIds" list`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> activity-status-code - String - 是否必填：否 -（过滤条件）按照伸缩活动状态过滤。（INIT：初始化中|RUNNING：运行中|SUCCESSFUL：活动成功|PARTIALLY_SUCCESSFUL：活动部分成功|FAILED：活动失败|CANCELLED：活动取消）</li>
	// <li> activity-type - String - 是否必填：否 -（过滤条件）按照伸缩活动类型过滤。（SCALE_OUT：扩容活动|SCALE_IN：缩容活动|ATTACH_INSTANCES：添加实例|REMOVE_INSTANCES：销毁实例|DETACH_INSTANCES：移出实例|TERMINATE_INSTANCES_UNEXPECTEDLY：实例在CVM控制台被销毁|REPLACE_UNHEALTHY_INSTANCE：替换不健康实例|UPDATE_LOAD_BALANCERS：更新负载均衡器）</li>
	// <li> activity-id - String - 是否必填：否 -（过滤条件）按照伸缩活动ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`ActivityIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 伸缩活动最早的开始时间，如果指定了ActivityIds，此参数将被忽略。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 伸缩活动最晚的结束时间，如果指定了ActivityIds，此参数将被忽略。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAutoScalingActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingActivitiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的伸缩活动数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 符合条件的伸缩活动信息集合。
		ActivitySet []*Activity `json:"ActivitySet,omitempty" name:"ActivitySet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingActivitiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupLastActivitiesRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID列表
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds" list`
}

func (r *DescribeAutoScalingGroupLastActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupLastActivitiesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupLastActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的伸缩活动信息集合。说明：伸缩组伸缩活动不存在的则不返回，如传50个伸缩组ID，返回45条数据，说明其中有5个伸缩组伸缩活动不存在。
		ActivitySet []*Activity `json:"ActivitySet,omitempty" name:"ActivitySet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingGroupLastActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupLastActivitiesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个伸缩组ID查询。伸缩组ID形如：`asg-nkdwoui0`。每次请求的上限为100。参数不支持同时指定`AutoScalingGroupIds`和`Filters`。
	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds" list`

	// 过滤条件。
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> auto-scaling-group-name - String - 是否必填：否 -（过滤条件）按照伸缩组名称过滤。</li>
	// <li> vague-auto-scaling-group-name - String - 是否必填：否 -（过滤条件）按照伸缩组名称模糊搜索。</li>
	// <li> launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。</li>
	// <li> tag-key - String - 是否必填：否 -（过滤条件）按照标签键进行过滤。</li>
	// <li> tag-value - String - 是否必填：否 -（过滤条件）按照标签值进行过滤。</li>
	// <li> tag:tag-key - String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingGroupIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAutoScalingGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingGroupsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 伸缩组详细信息列表。
		AutoScalingGroupSet []*AutoScalingGroup `json:"AutoScalingGroupSet,omitempty" name:"AutoScalingGroupSet" list`

		// 符合条件的伸缩组数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingGroupsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingInstancesRequest struct {
	*tchttp.BaseRequest

	// 待查询云服务器（CVM）的实例ID。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤条件。
	// <li> instance-id - String - 是否必填：否 -（过滤条件）按照实例ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`InstanceIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAutoScalingInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAutoScalingInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例详细信息列表。
		AutoScalingInstanceSet []*Instance `json:"AutoScalingInstanceSet,omitempty" name:"AutoScalingInstanceSet" list`

		// 符合条件的实例数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAutoScalingInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAutoScalingInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLaunchConfigurationsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个启动配置ID查询。启动配置ID形如：`asc-ouy1ax38`。每次请求的上限为100。参数不支持同时指定`LaunchConfigurationIds`和`Filters`
	LaunchConfigurationIds []*string `json:"LaunchConfigurationIds,omitempty" name:"LaunchConfigurationIds" list`

	// 过滤条件。
	// <li> launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。</li>
	// <li> launch-configuration-name - String - 是否必填：否 -（过滤条件）按照启动配置名称过滤。</li>
	// <li> vague-launch-configuration-name - String - 是否必填：否 -（过滤条件）按照启动配置名称模糊搜索。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LaunchConfigurationIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLaunchConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLaunchConfigurationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLaunchConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的启动配置数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 启动配置详细信息列表。
		LaunchConfigurationSet []*LaunchConfiguration `json:"LaunchConfigurationSet,omitempty" name:"LaunchConfigurationSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLaunchConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLaunchConfigurationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleHooksRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个生命周期挂钩ID查询。生命周期挂钩ID形如：`ash-8azjzxcl`。每次请求的上限为100。参数不支持同时指定`LifecycleHookIds`和`Filters`。
	LifecycleHookIds []*string `json:"LifecycleHookIds,omitempty" name:"LifecycleHookIds" list`

	// 过滤条件。
	// <li> lifecycle-hook-id - String - 是否必填：否 -（过滤条件）按照生命周期挂钩ID过滤。</li>
	// <li> lifecycle-hook-name - String - 是否必填：否 -（过滤条件）按照生命周期挂钩名称过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 过滤条件。
	// <li> lifecycle-hook-id - String - 是否必填：否 -（过滤条件）按照生命周期挂钩ID过滤。</li>
	// <li> lifecycle-hook-name - String - 是否必填：否 -（过滤条件）按照生命周期挂钩名称过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`LifecycleHookIds `和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeLifecycleHooksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLifecycleHooksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeLifecycleHooksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生命周期挂钩数组
		LifecycleHookSet []*LifecycleHook `json:"LifecycleHookSet,omitempty" name:"LifecycleHookSet" list`

		// 总体数量
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLifecycleHooksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeLifecycleHooksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotificationConfigurationsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个通知ID查询。实例ID形如：asn-2sestqbr。每次请求的实例的上限为100。参数不支持同时指定`AutoScalingNotificationIds`和`Filters`。
	AutoScalingNotificationIds []*string `json:"AutoScalingNotificationIds,omitempty" name:"AutoScalingNotificationIds" list`

	// 过滤条件。
	// <li> auto-scaling-notification-id - String - 是否必填：否 -（过滤条件）按照通知ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingNotificationIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNotificationConfigurationsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotificationConfigurationsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeNotificationConfigurationsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的通知数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 弹性伸缩事件通知详细信息列表。
		AutoScalingNotificationSet []*AutoScalingNotification `json:"AutoScalingNotificationSet,omitempty" name:"AutoScalingNotificationSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNotificationConfigurationsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeNotificationConfigurationsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePaiInstancesRequest struct {
	*tchttp.BaseRequest

	// 依据PAI实例的实例ID进行查询。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 过滤条件。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为20，最大值为100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePaiInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePaiInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePaiInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的PAI实例数量
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// PAI实例详细信息
		PaiInstanceSet []*PaiInstance `json:"PaiInstanceSet,omitempty" name:"PaiInstanceSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePaiInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePaiInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个告警策略ID查询。告警策略ID形如：asp-i9vkg894。每次请求的实例的上限为100。参数不支持同时指定`AutoScalingPolicyIds`和`Filters`。
	AutoScalingPolicyIds []*string `json:"AutoScalingPolicyIds,omitempty" name:"AutoScalingPolicyIds" list`

	// 过滤条件。
	// <li> auto-scaling-policy-id - String - 是否必填：否 -（过滤条件）按照告警策略ID过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 -（过滤条件）按照伸缩组ID过滤。</li>
	// <li> scaling-policy-name - String - 是否必填：否 -（过滤条件）按照告警策略名称过滤。</li>
	// 每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。参数不支持同时指定`AutoScalingPolicyIds`和`Filters`。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 弹性伸缩告警触发策略详细信息列表。
		ScalingPolicySet []*ScalingPolicy `json:"ScalingPolicySet,omitempty" name:"ScalingPolicySet" list`

		// 符合条件的通知数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledActionsRequest struct {
	*tchttp.BaseRequest

	// 按照一个或者多个定时任务ID查询。实例ID形如：asst-am691zxo。每次请求的实例的上限为100。参数不支持同时指定ScheduledActionIds和Filters。
	ScheduledActionIds []*string `json:"ScheduledActionIds,omitempty" name:"ScheduledActionIds" list`

	// 过滤条件。
	// <li> scheduled-action-id - String - 是否必填：否 -（过滤条件）按照定时任务ID过滤。</li>
	// <li> scheduled-action-name - String - 是否必填：否 - （过滤条件） 按照定时任务名称过滤。</li>
	// <li> auto-scaling-group-id - String - 是否必填：否 - （过滤条件） 按照伸缩组ID过滤。</li>
	Filters []*Filter `json:"Filters,omitempty" name:"Filters" list`

	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScheduledActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScheduledActionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的定时任务数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 定时任务详细信息列表。
		ScheduledActionSet []*ScheduledAction `json:"ScheduledActionSet,omitempty" name:"ScheduledActionSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScheduledActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScheduledActionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachInstancesRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *DetachInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 伸缩活动ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
}

func (r *DisableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DisableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DisableAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
}

func (r *EnableAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnableAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EnableAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`

	// 开启云监控服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type ExecuteScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 告警伸缩策略ID
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

	// 是否检查伸缩组活动处于冷却时间内，默认值为false
	HonorCooldown *bool `json:"HonorCooldown,omitempty" name:"HonorCooldown"`
}

func (r *ExecuteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExecuteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 伸缩活动ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	Values []*string `json:"Values,omitempty" name:"Values" list`
}

type ForwardLoadBalancer struct {

	// 负载均衡器ID
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" name:"LoadBalancerId"`

	// 应用型负载均衡监听器 ID
	ListenerId *string `json:"ListenerId,omitempty" name:"ListenerId"`

	// 目标规则属性列表
	TargetAttributes []*TargetAttribute `json:"TargetAttributes,omitempty" name:"TargetAttributes" list`

	// 转发规则ID，注意：针对七层监听器此参数必填
	LocationId *string `json:"LocationId,omitempty" name:"LocationId"`
}

type HostNameSettings struct {

	// 云服务器的主机名。
	// <br><li> 点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。
	// <br><li> 不支持 Windows 实例。
	// <br><li> 其他类型（Linux 等）实例：字符长度为[2, 40]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostName *string `json:"HostName,omitempty" name:"HostName"`

	// 云服务器主机名的风格，取值范围包括 ORIGINAL 和  UNIQUE，默认为 ORIGINAL。
	// <br><li> ORIGINAL，AS 直接将入参中所填的 HostName 传递给 CVM，CVM 可能会对 HostName 追加序列号，伸缩组中实例的 HostName 会出现冲突的情况。
	// <br><li> UNIQUE，入参所填的 HostName 相当于主机名前缀，AS 和 CVM 会对其进行拓展，伸缩组中实例的 HostName 可以保证唯一。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HostNameStyle *string `json:"HostNameStyle,omitempty" name:"HostNameStyle"`
}

type Instance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 启动配置名称
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// 生命周期状态，取值包括IN_SERVICE, CREATING, TERMINATING, ATTACHING, DETACHING, ATTACHING_LB, DETACHING_LB等
	LifeCycleState *string `json:"LifeCycleState,omitempty" name:"LifeCycleState"`

	// 健康状态，取值包括HEALTHY和UNHEALTHY
	HealthStatus *string `json:"HealthStatus,omitempty" name:"HealthStatus"`

	// 是否加入缩容保护
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitempty" name:"ProtectedFromScaleIn"`

	// 可用区
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 创建类型，取值包括AUTO_CREATION, MANUAL_ATTACHING。
	CreationType *string `json:"CreationType,omitempty" name:"CreationType"`

	// 实例加入时间
	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`

	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 版本号
	VersionNumber *int64 `json:"VersionNumber,omitempty" name:"VersionNumber"`

	// 伸缩组名称
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`
}

type InstanceChargePrepaid struct {

	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：<br><li>NOTIFY_AND_AUTO_RENEW：通知过期且自动续费<br><li>NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费<br><li>DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费<br><br>默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InstanceMarketOptionsRequest struct {
	*tchttp.BaseRequest

	// 竞价相关选项
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitempty" name:"SpotOptions"`

	// 市场选项类型，当前只支持取值：spot
	// 注意：此字段可能返回 null，表示取不到有效值。
	MarketType *string `json:"MarketType,omitempty" name:"MarketType"`
}

func (r *InstanceMarketOptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *InstanceMarketOptionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceTag struct {

	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type InternetAccessible struct {

	// 网络计费类型。取值范围：<br><li>BANDWIDTH_PREPAID：预付费按带宽结算<br><li>TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费<br><li>BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费<br><li>BANDWIDTH_PACKAGE：带宽包用户<br>默认取值：TRAFFIC_POSTPAID_BY_HOUR。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见[购买网络带宽](https://cloud.tencent.com/document/product/213/509)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetMaxBandwidthOut *uint64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`

	// 是否分配公网IP。取值范围：<br><li>TRUE：表示分配公网IP<br><li>FALSE：表示不分配公网IP<br><br>当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty" name:"PublicIpAssigned"`
}

type LaunchConfiguration struct {

	// 实例所属项目ID。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 启动配置名称。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// 实例机型。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例系统盘配置信息。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 实例数据盘配置信息。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 实例登录设置。
	LoginSettings *LimitedLoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 公网带宽相关信息设置。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 实例所属安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 启动配置关联的伸缩组。
	AutoScalingGroupAbstractSet []*AutoScalingGroupAbstract `json:"AutoScalingGroupAbstractSet,omitempty" name:"AutoScalingGroupAbstractSet" list`

	// 自定义数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 启动配置创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 实例的增强服务启用情况与其设置。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 镜像ID。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 启动配置当前状态。取值范围：<br><li>NORMAL：正常<br><li>IMAGE_ABNORMAL：启动配置镜像异常<br><li>CBS_SNAP_ABNORMAL：启动配置数据盘快照异常<br><li>SECURITY_GROUP_ABNORMAL：启动配置安全组异常<br>
	LaunchConfigurationStatus *string `json:"LaunchConfigurationStatus,omitempty" name:"LaunchConfigurationStatus"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。
	// <br><li>POSTPAID_BY_HOUR：按小时后付费
	// <br><li>SPOTPAID：竞价付费
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// 实例机型列表。
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// 标签列表。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// 版本号。
	VersionNumber *int64 `json:"VersionNumber,omitempty" name:"VersionNumber"`

	// 更新时间。
	UpdatedTime *string `json:"UpdatedTime,omitempty" name:"UpdatedTime"`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// 上次操作时，InstanceTypesCheckPolicy 取值。
	LastOperationInstanceTypesCheckPolicy *string `json:"LastOperationInstanceTypesCheckPolicy,omitempty" name:"LastOperationInstanceTypesCheckPolicy"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitempty" name:"HostNameSettings"`
}

type LifecycleHook struct {

	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

	// 生命周期挂钩名称
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" name:"LifecycleHookName"`

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 生命周期挂钩默认结果
	DefaultResult *string `json:"DefaultResult,omitempty" name:"DefaultResult"`

	// 生命周期挂钩等待超时时间
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitempty" name:"HeartbeatTimeout"`

	// 生命周期挂钩适用场景
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" name:"LifecycleTransition"`

	// 通知目标的附加信息
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" name:"NotificationMetadata"`

	// 创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`

	// 通知目标
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitempty" name:"NotificationTarget"`
}

type LimitedLoginSettings struct {

	// 密钥ID列表。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Password *string `json:"Password,omitempty" name:"Password"`

	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds" list`

	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeepImageLogin *bool `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type MetricAlarm struct {

	// 比较运算符，可选值：<br><li>GREATER_THAN：大于</li><li>GREATER_THAN_OR_EQUAL_TO：大于或等于</li><li>LESS_THAN：小于</li><li> LESS_THAN_OR_EQUAL_TO：小于或等于</li><li> EQUAL_TO：等于</li> <li>NOT_EQUAL_TO：不等于</li>
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// 指标名称，可选字段如下：<br><li>CPU_UTILIZATION：CPU利用率</li><li>MEM_UTILIZATION：内存利用率</li><li>LAN_TRAFFIC_OUT：内网出带宽</li><li>LAN_TRAFFIC_IN：内网入带宽</li><li>WAN_TRAFFIC_OUT：外网出带宽</li><li>WAN_TRAFFIC_IN：外网入带宽</li>
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 告警阈值：<br><li>CPU_UTILIZATION：[1, 100]，单位：%</li><li>MEM_UTILIZATION：[1, 100]，单位：%</li><li>LAN_TRAFFIC_OUT：>0，单位：Mbps </li><li>LAN_TRAFFIC_IN：>0，单位：Mbps</li><li>WAN_TRAFFIC_OUT：>0，单位：Mbps</li><li>WAN_TRAFFIC_IN：>0，单位：Mbps</li>
	Threshold *uint64 `json:"Threshold,omitempty" name:"Threshold"`

	// 时间周期，单位：秒，取值枚举值为60、300。
	Period *uint64 `json:"Period,omitempty" name:"Period"`

	// 重复次数。取值范围 [1, 10]
	ContinuousTime *uint64 `json:"ContinuousTime,omitempty" name:"ContinuousTime"`

	// 统计类型，可选字段如下：<br><li>AVERAGE：平均值</li><li>MAXIMUM：最大值<li>MINIMUM：最小值</li><br> 默认取值：AVERAGE
	Statistic *string `json:"Statistic,omitempty" name:"Statistic"`
}

type ModifyAutoScalingGroupRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName *string `json:"AutoScalingGroupName,omitempty" name:"AutoScalingGroupName"`

	// 默认冷却时间，单位秒，默认值为300
	DefaultCooldown *uint64 `json:"DefaultCooldown,omitempty" name:"DefaultCooldown"`

	// 期望实例数，大小介于最小实例数和最大实例数之间
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 最大实例数，取值范围为0-2000。
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 最小实例数，取值范围为0-2000。
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 项目ID
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 子网ID列表
	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds" list`

	// 销毁策略，目前长度上限为1。取值包括 OLDEST_INSTANCE 和 NEWEST_INSTANCE。
	// <br><li> OLDEST_INSTANCE 优先销毁伸缩组中最旧的实例。
	// <br><li> NEWEST_INSTANCE，优先销毁伸缩组中最新的实例。
	TerminationPolicies []*string `json:"TerminationPolicies,omitempty" name:"TerminationPolicies" list`

	// VPC ID，基础网络则填空字符串。修改为具体VPC ID时，需指定相应的SubnetIds；修改为基础网络时，需指定相应的Zones。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 可用区列表
	Zones []*string `json:"Zones,omitempty" name:"Zones" list`

	// 重试策略，取值包括 IMMEDIATE_RETRY、 INCREMENTAL_INTERVALS、NO_RETRY，默认取值为 IMMEDIATE_RETRY。
	// <br><li> IMMEDIATE_RETRY，立即重试，在较短时间内快速重试，连续失败超过一定次数（5次）后不再重试。
	// <br><li> INCREMENTAL_INTERVALS，间隔递增重试，随着连续失败次数的增加，重试间隔逐渐增大，重试间隔从秒级到1天不等。
	// <br><li> NO_RETRY，不进行重试，直到再次收到用户调用或者告警信息后才会重试。
	RetryPolicy *string `json:"RetryPolicy,omitempty" name:"RetryPolicy"`

	// 可用区校验策略，取值包括 ALL 和 ANY，默认取值为ANY。在伸缩组实际变更资源相关字段时（启动配置、可用区、子网）发挥作用。
	// <br><li> ALL，所有可用区（Zone）或子网（SubnetId）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个可用区（Zone）或子网（SubnetId）可用则通过校验，否则校验报错。
	// 
	// 可用区或子网不可用的常见原因包括该可用区CVM实例类型售罄、该可用区CBS云盘售罄、该可用区配额不足、该子网IP不足等。
	// 如果 Zones/SubnetIds 中可用区或者子网不存在，则无论 ZonesCheckPolicy 采用何种取值，都会校验报错。
	ZonesCheckPolicy *string `json:"ZonesCheckPolicy,omitempty" name:"ZonesCheckPolicy"`

	// 服务设置，包括云监控不健康替换等服务设置。
	ServiceSettings *ServiceSettings `json:"ServiceSettings,omitempty" name:"ServiceSettings"`
}

func (r *ModifyAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoScalingGroupRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAutoScalingGroupResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDesiredCapacityRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 期望实例数
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`
}

func (r *ModifyDesiredCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDesiredCapacityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyDesiredCapacityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDesiredCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyDesiredCapacityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLaunchConfigurationAttributesRequest struct {
	*tchttp.BaseRequest

	// 启动配置ID
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例类型列表，不同实例机型指定了不同的资源规格，最多支持5种实例机型。
	// 启动配置，通过 InstanceType 表示单一实例类型，通过 InstanceTypes 表示多实例类型。指定 InstanceTypes 成功启动配置后，原有的 InstanceType 自动失效。
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// 实例类型校验策略，在实际修改 InstanceTypes 时发挥作用，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。
	// 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitempty" name:"InstanceTypesCheckPolicy"`

	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。如果要清空UserData，则指定其为空字符串
	UserData *string `json:"UserData,omitempty" name:"UserData"`
}

func (r *ModifyLaunchConfigurationAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLaunchConfigurationAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLaunchConfigurationAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLaunchConfigurationAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLaunchConfigurationAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancersRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 传统负载均衡器ID列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `json:"LoadBalancerIds,omitempty" name:"LoadBalancerIds" list`

	// 应用型负载均衡器列表，目前长度上限为20，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `json:"ForwardLoadBalancers,omitempty" name:"ForwardLoadBalancers" list`

	// 负载均衡器校验策略，取值包括 ALL 和 DIFF，默认取值为 ALL。
	// <br><li> ALL，所有负载均衡器都合法则通过校验，否则校验报错。
	// <br><li> DIFF，仅校验负载均衡器参数中实际变化的部分，如果合法则通过校验，否则校验报错。
	LoadBalancersCheckPolicy *string `json:"LoadBalancersCheckPolicy,omitempty" name:"LoadBalancersCheckPolicy"`
}

func (r *ModifyLoadBalancersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoadBalancersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyLoadBalancersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 伸缩活动ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLoadBalancersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyLoadBalancersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNotificationConfigurationRequest struct {
	*tchttp.BaseRequest

	// 待修改的通知ID。
	AutoScalingNotificationId *string `json:"AutoScalingNotificationId,omitempty" name:"AutoScalingNotificationId"`

	// 通知类型，即为需要订阅的通知类型集合，取值范围如下：
	// <li>SCALE_OUT_SUCCESSFUL：扩容成功</li>
	// <li>SCALE_OUT_FAILED：扩容失败</li>
	// <li>SCALE_IN_SUCCESSFUL：缩容成功</li>
	// <li>SCALE_IN_FAILED：缩容失败</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL：替换不健康子机成功</li>
	// <li>REPLACE_UNHEALTHY_INSTANCE_FAILED：替换不健康子机失败</li>
	NotificationTypes []*string `json:"NotificationTypes,omitempty" name:"NotificationTypes" list`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[DescribeUserGroup](https://cloud.tencent.com/document/api/378/4404)查询。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *ModifyNotificationConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNotificationConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyNotificationConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNotificationConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyNotificationConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

	// 告警策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`

	// 告警触发后，期望实例数修改方式。取值 ：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值。取值：<br><li>当 AdjustmentType 为 CHANGE_IN_CAPACITY 时，AdjustmentValue 为正数表示告警触发后增加实例，为负数表示告警触发后减少实例 </li> <li> 当 AdjustmentType 为 EXACT_CAPACITY 时，AdjustmentValue 的值即为告警触发后新的期望实例数，需要大于或等于0 </li> <li> 当 AdjustmentType 为 PERCENT_CHANGE_IN_CAPACITY 时，AdjusmentValue 为正数表示告警触发后按百分比增加实例，为负数表示告警触发后按百分比减少实例，单位是：%。
	AdjustmentValue *int64 `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`

	// 冷却时间，单位为秒。
	Cooldown *uint64 `json:"Cooldown,omitempty" name:"Cooldown"`

	// 告警监控指标。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitempty" name:"MetricAlarm"`

	// 通知组ID，即为用户组ID集合，用户组ID可以通过[DescribeUserGroup](https://cloud.tencent.com/document/api/378/4404)查询。
	// 如果需要清空通知用户组，需要在列表中传入特定字符串 "NULL"。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

func (r *ModifyScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScheduledActionRequest struct {
	*tchttp.BaseRequest

	// 待修改的定时任务ID
	ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`

	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName *string `json:"ScheduledActionName,omitempty" name:"ScheduledActionName"`

	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// 定时任务的首次触发时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 定时任务的结束时间，取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。<br>此参数与`Recurrence`需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 定时任务的重复方式。为标准 Cron 格式<br>此参数与`EndTime`需要同时指定。
	Recurrence *string `json:"Recurrence,omitempty" name:"Recurrence"`
}

func (r *ModifyScheduledActionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScheduledActionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyScheduledActionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyScheduledActionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyScheduledActionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type NotificationTarget struct {

	// 目标类型，取值范围包括`CMQ_QUEUE`、`CMQ_TOPIC`。
	// <li> CMQ_QUEUE，指腾讯云消息队列-队列模型。</li>
	// <li> CMQ_TOPIC，指腾讯云消息队列-主题模型。</li>
	TargetType *string `json:"TargetType,omitempty" name:"TargetType"`

	// 队列名称，如果`TargetType`取值为`CMQ_QUEUE`，则本字段必填。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 主题名称，如果`TargetType`取值为`CMQ_TOPIC`，则本字段必填。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type PaiInstance struct {

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例域名
	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
}

type PreviewPaiDomainNameRequest struct {
	*tchttp.BaseRequest

	// 域名类型
	DomainNameType *string `json:"DomainNameType,omitempty" name:"DomainNameType"`
}

func (r *PreviewPaiDomainNameRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreviewPaiDomainNameRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PreviewPaiDomainNameResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 可用的PAI域名
		DomainName *string `json:"DomainName,omitempty" name:"DomainName"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreviewPaiDomainNameResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PreviewPaiDomainNameResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstancesRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// CVM实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`
}

func (r *RemoveInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RemoveInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 伸缩活动ID
		ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *RemoveInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// 是否开启[云监控](https://cloud.tencent.com/document/product/248)服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type RunSecurityServiceEnabled struct {

	// 是否开启[云安全](https://cloud.tencent.com/document/product/296)服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type ScalingPolicy struct {

	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 告警触发策略ID。
	AutoScalingPolicyId *string `json:"AutoScalingPolicyId,omitempty" name:"AutoScalingPolicyId"`

	// 告警触发策略名称。
	ScalingPolicyName *string `json:"ScalingPolicyName,omitempty" name:"ScalingPolicyName"`

	// 告警触发后，期望实例数修改方式。取值 ：<br><li>CHANGE_IN_CAPACITY：增加或减少若干期望实例数</li><li>EXACT_CAPACITY：调整至指定期望实例数</li> <li>PERCENT_CHANGE_IN_CAPACITY：按百分比调整期望实例数</li>
	AdjustmentType *string `json:"AdjustmentType,omitempty" name:"AdjustmentType"`

	// 告警触发后，期望实例数的调整值。
	AdjustmentValue *int64 `json:"AdjustmentValue,omitempty" name:"AdjustmentValue"`

	// 冷却时间。
	Cooldown *uint64 `json:"Cooldown,omitempty" name:"Cooldown"`

	// 告警监控指标。
	MetricAlarm *MetricAlarm `json:"MetricAlarm,omitempty" name:"MetricAlarm"`

	// 通知组ID，即为用户组ID集合。
	NotificationUserGroupIds []*string `json:"NotificationUserGroupIds,omitempty" name:"NotificationUserGroupIds" list`
}

type ScheduledAction struct {

	// 定时任务ID。
	ScheduledActionId *string `json:"ScheduledActionId,omitempty" name:"ScheduledActionId"`

	// 定时任务名称。
	ScheduledActionName *string `json:"ScheduledActionName,omitempty" name:"ScheduledActionName"`

	// 定时任务所在伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 定时任务的开始时间。取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 定时任务的重复方式。
	Recurrence *string `json:"Recurrence,omitempty" name:"Recurrence"`

	// 定时任务的结束时间。取值为`北京时间`（UTC+8），按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ss+08:00`。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 定时任务设置的最大实例数。
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 定时任务设置的期望实例数。
	DesiredCapacity *uint64 `json:"DesiredCapacity,omitempty" name:"DesiredCapacity"`

	// 定时任务设置的最小实例数。
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 定时任务的创建时间。取值为`UTC`时间，按照`ISO8601`标准，格式：`YYYY-MM-DDThh:mm:ssZ`。
	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ServiceSettings struct {

	// 开启监控不健康替换服务。若开启则对于云监控标记为不健康的实例，弹性伸缩服务会进行替换。若不指定该参数，则默认为 False。
	ReplaceMonitorUnhealthy *bool `json:"ReplaceMonitorUnhealthy,omitempty" name:"ReplaceMonitorUnhealthy"`
}

type SetInstancesProtectionRequest struct {
	*tchttp.BaseRequest

	// 伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 实例ID。
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例是否需要移出保护。
	ProtectedFromScaleIn *bool `json:"ProtectedFromScaleIn,omitempty" name:"ProtectedFromScaleIn"`
}

func (r *SetInstancesProtectionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetInstancesProtectionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetInstancesProtectionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetInstancesProtectionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetInstancesProtectionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SpotMarketOptions struct {

	// 竞价出价，例如“1.05”
	MaxPrice *string `json:"MaxPrice,omitempty" name:"MaxPrice"`

	// 竞价请求类型，当前仅支持类型：one-time，默认值为one-time
	// 注意：此字段可能返回 null，表示取不到有效值。
	SpotInstanceType *string `json:"SpotInstanceType,omitempty" name:"SpotInstanceType"`
}

type SystemDisk struct {

	// 系统盘类型。系统盘类型限制详见[CVM实例配置](https://cloud.tencent.com/document/product/213/2177)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><br>默认取值：LOCAL_BASIC。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小，单位：GB。默认值为 50
	// 注意：此字段可能返回 null，表示取不到有效值。
	DiskSize *uint64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Tag struct {

	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`

	// 标签绑定的资源类型，当前支持类型："auto-scaling-group
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type TargetAttribute struct {

	// 端口
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 权重
	Weight *uint64 `json:"Weight,omitempty" name:"Weight"`
}

type UpgradeLaunchConfigurationRequest struct {
	*tchttp.BaseRequest

	// 启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-8toqc6s3`。镜像类型分为四种：<br/><li>公共镜像</li><li>自定义镜像</li><li>共享镜像</li><li>服务市场镜像</li><br/>可通过以下方式获取可用的镜像ID：<br/><li>`公共镜像`、`自定义镜像`、`共享镜像`的镜像ID可通过登录[控制台](https://console.cloud.tencent.com/cvm/image?rid=1&imageType=PUBLIC_IMAGE)查询；`服务镜像市场`的镜像ID可通过[云市场](https://market.cloud.tencent.com/list)查询。</li><li>通过调用接口 [DescribeImages](https://cloud.tencent.com/document/api/213/15715) ，取返回信息中的`ImageId`字段。</li>
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 实例机型列表，不同实例机型指定了不同的资源规格，最多支持5种实例机型。
	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes" list`

	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName *string `json:"LaunchConfigurationName,omitempty" name:"LaunchConfigurationName"`

	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，最多支持指定11块数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks" list`

	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`

	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。
	// <br><li>POSTPAID_BY_HOUR：按小时后付费
	// <br><li>SPOTPAID：竞价付费
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `json:"InstanceMarketOptions,omitempty" name:"InstanceMarketOptions"`

	// 实例类型校验策略，取值包括 ALL 和 ANY，默认取值为ANY。
	// <br><li> ALL，所有实例类型（InstanceType）都可用则通过校验，否则校验报错。
	// <br><li> ANY，存在任何一个实例类型（InstanceType）可用则通过校验，否则校验报错。
	// 
	// 实例类型不可用的常见原因包括该实例类型售罄、对应云盘售罄等。
	// 如果 InstanceTypes 中一款机型不存在或者已下线，则无论 InstanceTypesCheckPolicy 采用何种取值，都会校验报错。
	InstanceTypesCheckPolicy *string `json:"InstanceTypesCheckPolicy,omitempty" name:"InstanceTypesCheckPolicy"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 实例所属项目ID。该参数可以通过调用 [DescribeProject](https://cloud.tencent.com/document/api/378/4400) 的返回值中的`projectId`字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的`SecurityGroupId`字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds" list`

	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。
	UserData *string `json:"UserData,omitempty" name:"UserData"`

	// 标签列表。通过指定该参数，可以为扩容的实例绑定标签。最多支持指定10个标签。
	InstanceTags []*InstanceTag `json:"InstanceTags,omitempty" name:"InstanceTags" list`

	// CAM角色名称。可通过DescribeRoleList接口返回值中的roleName获取。
	CamRoleName *string `json:"CamRoleName,omitempty" name:"CamRoleName"`

	// 云服务器主机名（HostName）的相关设置。
	HostNameSettings *HostNameSettings `json:"HostNameSettings,omitempty" name:"HostNameSettings"`
}

func (r *UpgradeLaunchConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLaunchConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLaunchConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeLaunchConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLaunchConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLifecycleHookRequest struct {
	*tchttp.BaseRequest

	// 生命周期挂钩ID
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" name:"LifecycleHookId"`

	// 生命周期挂钩名称
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" name:"LifecycleHookName"`

	// 进行生命周期挂钩的场景，取值范围包括“INSTANCE_LAUNCHING”和“INSTANCE_TERMINATING”
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" name:"LifecycleTransition"`

	// 定义伸缩组在生命周期挂钩超时的情况下应采取的操作，取值范围是“CONTINUE”或“ABANDON”，默认值为“CONTINUE”
	DefaultResult *string `json:"DefaultResult,omitempty" name:"DefaultResult"`

	// 生命周期挂钩超时之前可以经过的最长时间（以秒为单位），范围从30到3600秒，默认值为300秒
	HeartbeatTimeout *int64 `json:"HeartbeatTimeout,omitempty" name:"HeartbeatTimeout"`

	// 弹性伸缩向通知目标发送的附加信息，默认值为''
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" name:"NotificationMetadata"`

	// 通知目标
	NotificationTarget *NotificationTarget `json:"NotificationTarget,omitempty" name:"NotificationTarget"`
}

func (r *UpgradeLifecycleHookRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLifecycleHookRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpgradeLifecycleHookResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeLifecycleHookResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpgradeLifecycleHookResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
