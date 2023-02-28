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

package v20220401

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type AddClusterStorageOptionRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群存储选项。
	StorageOption *StorageOption `json:"StorageOption,omitempty" name:"StorageOption"`
}

type AddClusterStorageOptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群存储选项。
	StorageOption *StorageOption `json:"StorageOption,omitempty" name:"StorageOption"`
}

func (r *AddClusterStorageOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterStorageOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "StorageOption")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddClusterStorageOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddClusterStorageOptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddClusterStorageOptionResponse struct {
	*tchttp.BaseResponse
	Response *AddClusterStorageOptionResponseParams `json:"Response"`
}

func (r *AddClusterStorageOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddClusterStorageOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodesRequestParams struct {
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 添加节点数量。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜像和特定自定义镜像。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。<br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk []*SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 节点显示名称。
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 队列名称。不指定则为默认队列。<li>SLURM默认队列为：compute。<li>SGE默认队列为：all.q。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 添加节点角色。默认值：Compute<br><li>Compute：计算节点。<br><li>Login：登录节点。
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 添加节点类型。默认取值：STATIC。<li>STATIC：静态节点，不会参与弹性伸缩流程。<li>DYNAMIC：弹性节点，会被弹性缩容的节点。管控节点和登录节点不支持此参数。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

type AddNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 添加节点数量。
	Count *int64 `json:"Count,omitempty" name:"Count"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜像和特定自定义镜像。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。<br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk []*SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 节点显示名称。
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 队列名称。不指定则为默认队列。<li>SLURM默认队列为：compute。<li>SGE默认队列为：all.q。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 添加节点角色。默认值：Compute<br><li>Compute：计算节点。<br><li>Login：登录节点。
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 添加节点类型。默认取值：STATIC。<li>STATIC：静态节点，不会参与弹性伸缩流程。<li>DYNAMIC：弹性节点，会被弹性缩容的节点。管控节点和登录节点不支持此参数。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

func (r *AddNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "ClusterId")
	delete(f, "VirtualPrivateCloud")
	delete(f, "Count")
	delete(f, "ImageId")
	delete(f, "InstanceChargeType")
	delete(f, "InstanceChargePrepaid")
	delete(f, "InstanceType")
	delete(f, "SystemDisk")
	delete(f, "DataDisks")
	delete(f, "InternetAccessible")
	delete(f, "InstanceName")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "ClientToken")
	delete(f, "QueueName")
	delete(f, "NodeRole")
	delete(f, "DryRun")
	delete(f, "NodeType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddNodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddNodesResponse struct {
	*tchttp.BaseResponse
	Response *AddNodesResponseParams `json:"Response"`
}

func (r *AddNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddQueueRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type AddQueueRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *AddQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AddQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type AddQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type AddQueueResponse struct {
	*tchttp.BaseResponse
	Response *AddQueueResponseParams `json:"Response"`
}

func (r *AddQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AddQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoScalingGroupRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 弹性伸缩启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 弹性伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 队列名称。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。默认值120。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitempty" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。单位秒。默认值300。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitempty" name:"ShrinkIdleTime"`

	// 是否开启自动扩容，默认值true。
	EnableAutoExpansion *bool `json:"EnableAutoExpansion,omitempty" name:"EnableAutoExpansion"`

	// 是否开启自动缩容，默认值true。
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" name:"EnableAutoShrink"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会绑定弹性伸缩组。检查项包括是否填写了必需参数，请求格式，业务限制。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId。
	// false（默认）：发送正常请求，通过检查后直接绑定弹性伸缩组。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type BindAutoScalingGroupRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 弹性伸缩启动配置ID。
	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`

	// 弹性伸缩组ID。
	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`

	// 队列名称。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。默认值120。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitempty" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。单位秒。默认值300。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitempty" name:"ShrinkIdleTime"`

	// 是否开启自动扩容，默认值true。
	EnableAutoExpansion *bool `json:"EnableAutoExpansion,omitempty" name:"EnableAutoExpansion"`

	// 是否开启自动缩容，默认值true。
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" name:"EnableAutoShrink"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会绑定弹性伸缩组。检查项包括是否填写了必需参数，请求格式，业务限制。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId。
	// false（默认）：发送正常请求，通过检查后直接绑定弹性伸缩组。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *BindAutoScalingGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoScalingGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LaunchConfigurationId")
	delete(f, "AutoScalingGroupId")
	delete(f, "QueueName")
	delete(f, "ExpansionBusyTime")
	delete(f, "ShrinkIdleTime")
	delete(f, "EnableAutoExpansion")
	delete(f, "EnableAutoShrink")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "BindAutoScalingGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type BindAutoScalingGroupResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type BindAutoScalingGroupResponse struct {
	*tchttp.BaseResponse
	Response *BindAutoScalingGroupResponseParams `json:"Response"`
}

func (r *BindAutoScalingGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *BindAutoScalingGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CFSOption struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// 文件系统远程挂载ip及路径。
	RemotePath *string `json:"RemotePath,omitempty" name:"RemotePath"`

	// 文件系统协议类型，默认值NFS 3.0。
	// <li>NFS 3.0。
	// <li>NFS 4.0。
	// <li>TURBO。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值SD；其中 SD 为通用标准型标准型存储， HP为通用性能型存储， TB为turbo标准型， TP 为turbo性能型。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

type CFSOptionOverview struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// 文件系统远程挂载ip及路径。
	RemotePath *string `json:"RemotePath,omitempty" name:"RemotePath"`

	// 文件系统协议类型。
	// <li>NFS 3.0。
	// <li>NFS 4.0。
	// <li>TURBO。
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 文件系统存储类型，默认值SD；其中 SD 为通用标准型标准型存储， HP为通用性能型存储， TB为turbo标准型， TP 为turbo性能型。
	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
}

type ClusterActivity struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群活动ID。
	ActivityId *string `json:"ActivityId,omitempty" name:"ActivityId"`

	// 集群活动类型。取值范围：<br><li>CreateAndAddNodes：创建实例并添加进集群<br><li>RemoveNodesFromCluster：从集群移除实例<br><li>TerminateNodes：销毁实例<br><li>MountStorageOption：增加挂载选项并进行挂载<br><li>UmountStorageOption：删除集群挂载存储选项并解挂载
	ActivityType *string `json:"ActivityType,omitempty" name:"ActivityType"`

	// 集群活动状态。取值范围：<br><li>PENDING：等待运行<br><li>RUNNING：运行中<br><li>SUCCESSFUL：活动成功<br><li>PARTIALLY_SUCCESSFUL：活动部分成功<br><li>FAILED：活动失败
	ActivityStatus *string `json:"ActivityStatus,omitempty" name:"ActivityStatus"`

	// 集群活动状态码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActivityStatusCode *string `json:"ActivityStatusCode,omitempty" name:"ActivityStatusCode"`

	// 集群活动结果详情。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResultDetail *string `json:"ResultDetail,omitempty" name:"ResultDetail"`

	// 集群活动起因。
	Cause *string `json:"Cause,omitempty" name:"Cause"`

	// 集群活动描述。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 集群活动相关节点活动集合。
	RelatedNodeActivitySet []*NodeActivity `json:"RelatedNodeActivitySet,omitempty" name:"RelatedNodeActivitySet"`

	// 集群活动开始时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 集群活动结束时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type ClusterOverview struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 集群状态。取值范围：<br><li>PENDING：创建中<br><li>INITING：初始化中<br><li>INIT_FAILED：初始化失败<br><li>RUNNING：运行中<br><li>TERMINATING：销毁中
	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`

	// 集群名称。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群位置信息。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 集群创建时间。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 集群调度器。
	SchedulerType *string `json:"SchedulerType,omitempty" name:"SchedulerType"`

	// 计算节点数量。
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty" name:"ComputeNodeCount"`

	// 计算节点概览。
	ComputeNodeSet []*ComputeNodeOverview `json:"ComputeNodeSet,omitempty" name:"ComputeNodeSet"`

	// 管控节点数量。
	ManagerNodeCount *int64 `json:"ManagerNodeCount,omitempty" name:"ManagerNodeCount"`

	// 管控节点概览。
	ManagerNodeSet []*ManagerNodeOverview `json:"ManagerNodeSet,omitempty" name:"ManagerNodeSet"`

	// 登录节点概览。
	LoginNodeSet []*LoginNodeOverview `json:"LoginNodeSet,omitempty" name:"LoginNodeSet"`

	// 登录节点数量。
	LoginNodeCount *int64 `json:"LoginNodeCount,omitempty" name:"LoginNodeCount"`

	// 集群所属私有网络ID。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type ComputeNode struct {
	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 节点显示名称。<br><li>
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ComputeNodeOverview struct {
	// 计算节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 指定管理节点。
	ManagerNode *ManagerNode `json:"ManagerNode,omitempty" name:"ManagerNode"`

	// 指定管理节点的数量。默认取值：1。取值范围：1～2。
	ManagerNodeCount *int64 `json:"ManagerNodeCount,omitempty" name:"ManagerNodeCount"`

	// 指定计算节点。
	ComputeNode *ComputeNode `json:"ComputeNode,omitempty" name:"ComputeNode"`

	// 指定计算节点的数量。默认取值：0。
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty" name:"ComputeNodeCount"`

	// 调度器类型。默认取值：SLURM。<br><li>SGE：SGE调度器。<br><li>SLURM：SLURM调度器。
	SchedulerType *string `json:"SchedulerType,omitempty" name:"SchedulerType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜像。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 域名字服务类型。默认取值：NIS。
	// <li>NIS：NIS域名字服务。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 集群显示名称。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群存储选项
	StorageOption *StorageOption `json:"StorageOption,omitempty" name:"StorageOption"`

	// 指定登录节点。
	LoginNode *LoginNode `json:"LoginNode,omitempty" name:"LoginNode"`

	// 指定登录节点的数量。默认取值：0。取值范围：0～10。
	LoginNodeCount *int64 `json:"LoginNodeCount,omitempty" name:"LoginNodeCount"`

	// 创建集群时同时绑定的标签对说明。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 弹性伸缩类型。<br><li>AS：集群自动扩缩容由[弹性伸缩](https://cloud.tencent.com/document/product/377/3154)产品实现。<br><li>THPC_AS：集群自动扩缩容由THPC产品内部实现。
	AutoScalingType *string `json:"AutoScalingType,omitempty" name:"AutoScalingType"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群中实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 指定管理节点。
	ManagerNode *ManagerNode `json:"ManagerNode,omitempty" name:"ManagerNode"`

	// 指定管理节点的数量。默认取值：1。取值范围：1～2。
	ManagerNodeCount *int64 `json:"ManagerNodeCount,omitempty" name:"ManagerNodeCount"`

	// 指定计算节点。
	ComputeNode *ComputeNode `json:"ComputeNode,omitempty" name:"ComputeNode"`

	// 指定计算节点的数量。默认取值：0。
	ComputeNodeCount *int64 `json:"ComputeNodeCount,omitempty" name:"ComputeNodeCount"`

	// 调度器类型。默认取值：SLURM。<br><li>SGE：SGE调度器。<br><li>SLURM：SLURM调度器。
	SchedulerType *string `json:"SchedulerType,omitempty" name:"SchedulerType"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜像。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 集群登录设置。
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`

	// 集群中实例所属安全组。该参数可以通过调用 [DescribeSecurityGroups](https://cloud.tencent.com/document/api/215/15808) 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`

	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `json:"ClientToken,omitempty" name:"ClientToken"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会创建实例。检查项包括是否填写了必需参数，请求格式，业务限制和云服务器库存。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId.
	// false（默认）：发送正常请求，通过检查后直接创建实例
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`

	// 域名字服务类型。默认取值：NIS。
	// <li>NIS：NIS域名字服务。
	AccountType *string `json:"AccountType,omitempty" name:"AccountType"`

	// 集群显示名称。
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// 集群存储选项
	StorageOption *StorageOption `json:"StorageOption,omitempty" name:"StorageOption"`

	// 指定登录节点。
	LoginNode *LoginNode `json:"LoginNode,omitempty" name:"LoginNode"`

	// 指定登录节点的数量。默认取值：0。取值范围：0～10。
	LoginNodeCount *int64 `json:"LoginNodeCount,omitempty" name:"LoginNodeCount"`

	// 创建集群时同时绑定的标签对说明。
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// 弹性伸缩类型。<br><li>AS：集群自动扩缩容由[弹性伸缩](https://cloud.tencent.com/document/product/377/3154)产品实现。<br><li>THPC_AS：集群自动扩缩容由THPC产品内部实现。
	AutoScalingType *string `json:"AutoScalingType,omitempty" name:"AutoScalingType"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Placement")
	delete(f, "ManagerNode")
	delete(f, "ManagerNodeCount")
	delete(f, "ComputeNode")
	delete(f, "ComputeNodeCount")
	delete(f, "SchedulerType")
	delete(f, "ImageId")
	delete(f, "VirtualPrivateCloud")
	delete(f, "LoginSettings")
	delete(f, "SecurityGroupIds")
	delete(f, "ClientToken")
	delete(f, "DryRun")
	delete(f, "AccountType")
	delete(f, "ClusterName")
	delete(f, "StorageOption")
	delete(f, "LoginNode")
	delete(f, "LoginNodeCount")
	delete(f, "Tags")
	delete(f, "AutoScalingType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateClusterResponseParams struct {
	// 集群ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：[存储概述](https://cloud.tencent.com/document/product/213/4952)。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`

	// 数据盘类型。数据盘类型限制详见[存储概述](https://cloud.tencent.com/document/product/213/4952)。取值范围：<br><li>LOCAL_BASIC：本地硬盘<br><li>LOCAL_SSD：本地SSD硬盘<br><li>LOCAL_NVME：本地NVME硬盘，与InstanceType强相关，不支持指定<br><li>LOCAL_PRO：本地HDD硬盘，与InstanceType强相关，不支持指定<br><li>CLOUD_BASIC：普通云硬盘<br><li>CLOUD_PREMIUM：高性能云硬盘<br><li>CLOUD_SSD：SSD云硬盘<br><li>CLOUD_HSSD：增强型SSD云硬盘<br><li>CLOUD_TSSD：极速型SSD云硬盘<br><br>默认取值：LOCAL_BASIC。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterStorageOptionRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 本地挂载路径。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`
}

type DeleteClusterStorageOptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 本地挂载路径。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`
}

func (r *DeleteClusterStorageOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterStorageOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "LocalPath")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterStorageOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteClusterStorageOptionResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteClusterStorageOptionResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterStorageOptionResponseParams `json:"Response"`
}

func (r *DeleteClusterStorageOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterStorageOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodesRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 节点ID。
	NodeIds []*string `json:"NodeIds,omitempty" name:"NodeIds"`
}

type DeleteNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 节点ID。
	NodeIds []*string `json:"NodeIds,omitempty" name:"NodeIds"`
}

func (r *DeleteNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NodeIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteNodesResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteNodesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteNodesResponseParams `json:"Response"`
}

func (r *DeleteNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQueueRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type DeleteQueueRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 队列名称。<br><li>最多支持32个字符。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DeleteQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteQueueResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteQueueResponseParams `json:"Response"`
}

func (r *DeleteQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingConfigurationRequestParams struct {
	// 集群ID。	
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeAutoScalingConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。	
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAutoScalingConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAutoScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAutoScalingConfigurationResponseParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitempty" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitempty" name:"ShrinkIdleTime"`

	// 扩容队列配置概览列表。
	QueueConfigs []*QueueConfigOverview `json:"QueueConfigs,omitempty" name:"QueueConfigs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAutoScalingConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAutoScalingConfigurationResponseParams `json:"Response"`
}

func (r *DescribeAutoScalingConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAutoScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterActivitiesRequestParams struct {
	// 集群ID。通过该参数指定需要查询活动历史记录的集群。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeClusterActivitiesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。通过该参数指定需要查询活动历史记录的集群。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterActivitiesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterActivitiesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterActivitiesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterActivitiesResponseParams struct {
	// 集群活动历史记录列表。
	ClusterActivitySet []*ClusterActivity `json:"ClusterActivitySet,omitempty" name:"ClusterActivitySet"`

	// 集群活动历史记录数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterActivitiesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterActivitiesResponseParams `json:"Response"`
}

func (r *DescribeClusterActivitiesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterActivitiesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStorageOptionRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type DescribeClusterStorageOptionRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterStorageOptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStorageOptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterStorageOptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClusterStorageOptionResponseParams struct {
	// 集群存储选项信息概览。
	StorageOption *StorageOptionOverview `json:"StorageOption,omitempty" name:"StorageOption"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClusterStorageOptionResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterStorageOptionResponseParams `json:"Response"`
}

func (r *DescribeClusterStorageOptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterStorageOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// 集群ID列表。通过该参数可以指定需要查询信息的集群列表。<br>如果您不指定该参数，则返回Limit数量以内的集群信息。
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID列表。通过该参数可以指定需要查询信息的集群列表。<br>如果您不指定该参数，则返回Limit数量以内的集群信息。
	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterIds")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// 集群概览信息列表。
	ClusterSet []*ClusterOverview `json:"ClusterSet,omitempty" name:"ClusterSet"`

	// 集群数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodesRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// <li><strong>queue-name</strong></li> <p style="padding-left: 30px;">按照【<strong>队列名称</strong>】进行过滤。队列名称形如：compute。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-role</strong></li> <p style="padding-left: 30px;">按照【<strong>节点角色</strong>】进行过滤。节点角色形如：Manager。（Manager：管控节点。Compute：计算节点。Login：登录节点。ManagerBackup：备用管控节点。）</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-type</strong></li> <p style="padding-left: 30px;">按照【<strong>节点类型</strong>】进行过滤。节点类型形如：STATIC。(STATIC：静态节点。DYNAMIC：弹性节点。)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeNodesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// <li><strong>queue-name</strong></li> <p style="padding-left: 30px;">按照【<strong>队列名称</strong>】进行过滤。队列名称形如：compute。</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-role</strong></li> <p style="padding-left: 30px;">按照【<strong>节点角色</strong>】进行过滤。节点角色形如：Manager。（Manager：管控节点。Compute：计算节点。Login：登录节点。ManagerBackup：备用管控节点。）</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;"><li><strong>node-type</strong></li> <p style="padding-left: 30px;">按照【<strong>节点类型</strong>】进行过滤。节点类型形如：STATIC。(STATIC：静态节点。DYNAMIC：弹性节点。)</p><p style="padding-left: 30px;">类型：String</p><p style="padding-left: 30px;">必选：否</p><p style="padding-left: 30px;">每次请求的`Filters`的上限为10，`Filter.Values`的上限为5。
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNodesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeNodesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeNodesResponseParams struct {
	// 节点概览信息列表。
	NodeSet []*NodeOverview `json:"NodeSet,omitempty" name:"NodeSet"`

	// 符合条件的节点数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeNodesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeNodesResponseParams `json:"Response"`
}

func (r *DescribeNodesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueuesRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeQueuesRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 偏移量，默认为0。关于`Offset`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，默认为20，最大值为100。关于`Limit`的更进一步介绍请参考 API [简介](https://cloud.tencent.com/document/api/213/15688)中的相关小节。
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeQueuesResponseParams struct {
	// 队列概览信息列表。
	QueueSet []*QueueOverview `json:"QueueSet,omitempty" name:"QueueSet"`

	// 符合条件的节点数量。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeQueuesResponseParams `json:"Response"`
}

func (r *DescribeQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExpansionNodeConfig struct {
	// 扩容实例所在的位置。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br><li>SPOTPAID：竞价付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 私有网络相关信息配置。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`
}

type ExpansionNodeConfigOverview struct {
	// 节点机型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 扩容实例所在的位置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Placement *Placement `json:"Placement,omitempty" name:"Placement"`

	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 私有网络相关信息配置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty" name:"VirtualPrivateCloud"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 公网带宽相关信息设置。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 节点系统盘配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 节点数据盘配置信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

type Filter struct {
	// 需要过滤的字段。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 字段的过滤值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type GooseFSOption struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// 文件系统远程挂载路径。
	RemotePath *string `json:"RemotePath,omitempty" name:"RemotePath"`

	// 文件系统master的ip和端口。
	Masters []*string `json:"Masters,omitempty" name:"Masters"`
}

type GooseFSOptionOverview struct {
	// 文件系统本地挂载路径。
	LocalPath *string `json:"LocalPath,omitempty" name:"LocalPath"`

	// 文件系统远程挂载路径。
	RemotePath *string `json:"RemotePath,omitempty" name:"RemotePath"`

	// 文件系统master的ip和端口。
	Masters []*string `json:"Masters,omitempty" name:"Masters"`
}

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36, 48, 60。
	Period *int64 `json:"Period,omitempty" name:"Period"`

	// 自动续费标识。取值范围：
	// NOTIFY_AND_AUTO_RENEW：通知过期且自动续费
	// NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费
	// DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费
	// 
	// 默认取值：NOTIFY_AND_MANUAL_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type InternetAccessible struct {
	// 网络计费类型。取值范围：
	// BANDWIDTH_PREPAID：预付费按带宽结算
	// TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费
	// BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费
	// BANDWIDTH_PACKAGE：带宽包用户
	// 默认取值：非带宽包用户默认与子机付费类型保持一致。
	InternetChargeType *string `json:"InternetChargeType,omitempty" name:"InternetChargeType"`

	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty" name:"InternetMaxBandwidthOut"`
}

type LoginNode struct {
	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk []*SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible []*InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 节点显示名称。<br><li>
	// 不指定节点显示名称则默认显示‘未命名’。
	// 最多支持60个字符。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type LoginNodeOverview struct {
	// 登录节点ID。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

type LoginSettings struct {
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到30位，至少包括两项[a-z]，[A-Z]、[0-9] 和 [( ) \` ~ ! @ # $ % ^ & *  - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到30位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) \` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitempty" name:"Password"`
}

type ManagerNode struct {
	// 节点[计费类型](https://cloud.tencent.com/document/product/213/2180)。<br><li>PREPAID：预付费，即包年包月<br><li>POSTPAID_BY_HOUR：按小时后付费<br>默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`

	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月节点的购买时长、是否设置自动续费等属性。若指定节点的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`

	// 节点机型。不同实例机型指定了不同的资源规格。
	// <br><li>对于付费模式为PREPAID或POSTPAID\_BY\_HOUR的实例创建，具体取值可通过调用接口[DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)来获得最新的规格表或参见[实例规格](https://cloud.tencent.com/document/product/213/11518)描述。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 节点显示名称。<br><li>
	// 不指定节点显示名称则默认显示‘未命名’。
	// </li><li>购买多个节点，如果指定模式串`{R:x}`，表示生成数字[`[x, x+n-1]`，其中`n`表示购买节点的数量，例如`server_{R:3}`，购买1个时，节点显示名称为`server_3`；购买2个时，节点显示名称分别为`server_3`，`server_4`。支持指定多个模式串`{R:x}`。
	// 购买多个节点，如果不指定模式串，则在节点显示名称添加后缀`1、2...n`，其中`n`表示购买节点的数量，例如`server_`，购买2个时，节点显示名称分别为`server_1`，`server_2`。</li><li>
	// 最多支持60个字符（包含模式串）。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type ManagerNodeOverview struct {
	// 管控节点ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
}

type NodeActivity struct {
	// 节点活动所在的实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeInstanceId *string `json:"NodeInstanceId,omitempty" name:"NodeInstanceId"`

	// 节点活动状态。取值范围：<br><li>RUNNING：运行中<br><li>SUCCESSFUL：活动成功<br><li>FAILED：活动失败
	NodeActivityStatus *string `json:"NodeActivityStatus,omitempty" name:"NodeActivityStatus"`

	// 节点活动状态码。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeActivityStatusCode *string `json:"NodeActivityStatusCode,omitempty" name:"NodeActivityStatusCode"`

	// 节点活动状态原因。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeActivityStatusReason *string `json:"NodeActivityStatusReason,omitempty" name:"NodeActivityStatusReason"`
}

type NodeOverview struct {
	// 节点实例ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 节点所在可用区信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Zone *string `json:"Zone,omitempty" name:"Zone"`

	// 节点状态。<br><li>SUBMITTED：已完成提交。<br><li>CREATING：创建中。<br><li>CREATED：完成创建。<br><li>INITING：初始化中。<br><li>INIT_FAILED：初始化失败。<br><li>RUNNING：运行中。<br><li>DELETING：销毁中。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeState *string `json:"NodeState,omitempty" name:"NodeState"`

	// 镜像ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 节点所属队列名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 节点角色。<br><li>Manager：管控节点。<br><li>Compute：计算节点。<br><li>Login：登录节点。<br><li>ManagerBackup：备用管控节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`

	// 节点类型。<br><li>STATIC：静态节点。<br><li>DYNAMIC：弹性节点。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
}

type Placement struct {
	// 实例所属的可用区名称。该参数可以通过调用  [DescribeZones](https://cloud.tencent.com/document/product/213/15707) 的返回值中的Zone字段来获取。
	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type QueueConfig struct {
	// 队列名称。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 队列中弹性节点数量最小值。取值范围0～200。
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 队列中弹性节点数量最大值。取值范围0～200。
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 是否开启自动扩容。
	EnableAutoExpansion *bool `json:"EnableAutoExpansion,omitempty" name:"EnableAutoExpansion"`

	// 是否开启自动缩容。
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" name:"EnableAutoShrink"`

	// 指定有效的[镜像](https://cloud.tencent.com/document/product/213/4940)ID，格式形如`img-xxx`。目前仅支持公有镜和特定自定义镜像。
	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`

	// 节点系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty" name:"SystemDisk"`

	// 节点数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定21块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含20块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`

	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`

	// 扩容节点配置信息。
	ExpansionNodeConfigs []*ExpansionNodeConfig `json:"ExpansionNodeConfigs,omitempty" name:"ExpansionNodeConfigs"`
}

type QueueConfigOverview struct {
	// 队列名称。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// 队列中弹性节点数量最小值。取值范围0～200。
	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`

	// 队列中弹性节点数量最大值。取值范围0～200。
	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 是否开启自动扩容。
	EnableAutoExpansion *bool `json:"EnableAutoExpansion,omitempty" name:"EnableAutoExpansion"`

	// 是否开启自动缩容。
	EnableAutoShrink *bool `json:"EnableAutoShrink,omitempty" name:"EnableAutoShrink"`

	// 扩容节点配置信息。
	ExpansionNodeConfigs []*ExpansionNodeConfigOverview `json:"ExpansionNodeConfigs,omitempty" name:"ExpansionNodeConfigs"`
}

type QueueOverview struct {
	// 队列名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

// Predefined struct for user
type SetAutoScalingConfigurationRequestParams struct {
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。默认值120。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitempty" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。单位秒。默认值300。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitempty" name:"ShrinkIdleTime"`

	// 扩容队列配置列表。
	QueueConfigs []*QueueConfig `json:"QueueConfigs,omitempty" name:"QueueConfigs"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会绑定弹性伸缩组。检查项包括是否填写了必需参数，请求格式，业务限制。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId。
	// false（默认）：发送正常请求，通过检查后直接绑定弹性伸缩组。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

type SetAutoScalingConfigurationRequest struct {
	*tchttp.BaseRequest
	
	// 集群ID。
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 任务连续等待时间，队列的任务处于连续等待的时间。单位秒。默认值120。
	ExpansionBusyTime *int64 `json:"ExpansionBusyTime,omitempty" name:"ExpansionBusyTime"`

	// 节点连续空闲（未运行作业）时间，一个节点连续处于空闲状态时间。单位秒。默认值300。
	ShrinkIdleTime *int64 `json:"ShrinkIdleTime,omitempty" name:"ShrinkIdleTime"`

	// 扩容队列配置列表。
	QueueConfigs []*QueueConfig `json:"QueueConfigs,omitempty" name:"QueueConfigs"`

	// 是否只预检此次请求。
	// true：发送检查请求，不会绑定弹性伸缩组。检查项包括是否填写了必需参数，请求格式，业务限制。
	// 如果检查不通过，则返回对应错误码；
	// 如果检查通过，则返回RequestId。
	// false（默认）：发送正常请求，通过检查后直接绑定弹性伸缩组。
	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *SetAutoScalingConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoScalingConfigurationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ExpansionBusyTime")
	delete(f, "ShrinkIdleTime")
	delete(f, "QueueConfigs")
	delete(f, "DryRun")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SetAutoScalingConfigurationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SetAutoScalingConfigurationResponseParams struct {
	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type SetAutoScalingConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *SetAutoScalingConfigurationResponseParams `json:"Response"`
}

func (r *SetAutoScalingConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SetAutoScalingConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StorageOption struct {
	// 集群挂载CFS文件系统选项
	CFSOptions []*CFSOption `json:"CFSOptions,omitempty" name:"CFSOptions"`

	// 集群挂在GooseFS文件系统选项
	GooseFSOptions []*GooseFSOption `json:"GooseFSOptions,omitempty" name:"GooseFSOptions"`
}

type StorageOptionOverview struct {
	// CFS存储选项概览信息列表。
	CFSOptions []*CFSOptionOverview `json:"CFSOptions,omitempty" name:"CFSOptions"`

	// GooseFS存储选项概览信息列表。
	GooseFSOptions []*GooseFSOptionOverview `json:"GooseFSOptions,omitempty" name:"GooseFSOptions"`
}

type SystemDisk struct {
	// 系统盘类型。系统盘类型限制详见存储概述。取值范围：
	// LOCAL_BASIC：本地硬盘
	// LOCAL_SSD：本地SSD硬盘
	// CLOUD_BASIC：普通云硬盘
	// CLOUD_SSD：SSD云硬盘
	// CLOUD_PREMIUM：高性能云硬盘
	// 
	// 默认取值：当前有库存的硬盘类型。
	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`

	// 系统盘大小，单位：GB。默认值为 50
	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

type Tag struct {
	// 标签键
	Key *string `json:"Key,omitempty" name:"Key"`

	// 标签值
	Value *string `json:"Value,omitempty" name:"Value"`
}

type VirtualPrivateCloud struct {
	// 私有网络ID，形如`vpc-xxx`。有效的VpcId可通过登录[控制台](https://console.cloud.tencent.com/vpc/vpc?rid=1)查询；也可以调用接口 [DescribeVpcEx](/document/api/215/1372) ，从接口返回中的`unVpcId`字段获取。若在创建子机时VpcId与SubnetId同时传入`DEFAULT`，则强制使用默认vpc网络。
	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 私有网络子网ID，形如`subnet-xxx`。有效的私有网络子网ID可通过登录[控制台](https://console.cloud.tencent.com/vpc/subnet?rid=1)查询；也可以调用接口  [DescribeSubnets](/document/api/215/15784) ，从接口返回中的`unSubnetId`字段获取。若在创建子机时SubnetId与VpcId同时传入`DEFAULT`，则强制使用默认vpc网络。
	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}